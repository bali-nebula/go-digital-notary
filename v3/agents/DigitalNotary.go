/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agents

import (
	byt "bytes"
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	com "github.com/bali-nebula/go-digital-notary/v3/components"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DigitalNotaryClass() DigitalNotaryClassLike {
	return digitalNotaryClass()
}

// Constructor Methods

func (c *digitalNotaryClass_) DigitalNotary(
	ssm Trusted,
	hsm Hardened,
) DigitalNotaryLike {
	// Validate the arguments.
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}

	// Reset the HSM.
	hsm.EraseKeys()

	// Create the new digital notary.
	var certificate com.DocumentLike
	var instance = &digitalNotary_{
		// Initialize the instance attributes.
		ssm_:         ssm,
		hsm_:         hsm,
		certificate_: certificate,
	}

	return instance
}

func (c *digitalNotaryClass_) DigitalNotaryWithCertificate(
	ssm Trusted,
	hsm Hardened,
	certificate com.DocumentLike,
) DigitalNotaryLike {
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}
	if uti.IsUndefined(certificate) || !certificate.IsNotarized() {
		panic("a notarized \"certificate\" attribute is required by this class.")
	}

	// Validate the seal on the certificate document.
	var seal = certificate.RemoveNotarySeal()
	var source = certificate.AsSource()
	var sourceBytes = []byte(source)
	certificate.SetNotarySeal(seal)
	var keyBytes = hsm.GetPublicKey()
	var signatureBytes = seal.GetSignature().AsIntrinsic()
	if !hsm.IsValid(keyBytes, sourceBytes, signatureBytes) {
		var message = fmt.Sprintf(
			"The \"certificate\" document is invalid: %s\n",
			certificate.AsSource(),
		)
		panic(message)
	}

	// Create the new digital notary.
	var instance = &digitalNotary_{
		// Initialize the instance attributes.
		ssm_:         ssm,
		hsm_:         hsm,
		certificate_: certificate,
	}

	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *digitalNotary_) GetClass() DigitalNotaryClassLike {
	return digitalNotaryClass()
}

func (v *digitalNotary_) CiteDocument(
	document com.DocumentLike,
) com.CitationLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to create a citation to a document",
	)

	// Create a citation to the document.
	var algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = []byte(source)
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	var content = document.GetContent()
	var tag = content.GetTag()
	var version = content.GetVersion()
	var citation = com.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)
	return citation
}

func (v *digitalNotary_) CitationMatches(
	citation com.CitationLike,
	document com.DocumentLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify a document citation",
	)

	// Compare the citation digest with a digest of the document.
	var citationAlgorithm = citation.GetAlgorithm()
	var citationDigest = citation.GetDigest()
	var ssmAlgorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = []byte(source)
	var documentDigest = doc.Binary(v.ssm_.DigestBytes(bytes))
	if citationAlgorithm.AsSource() != ssmAlgorithm.AsSource() {
		return false
	}
	if !byt.Equal(citationDigest.AsIntrinsic(), documentDigest.AsIntrinsic()) {
		return false
	}
	return true
}

func (v *digitalNotary_) GenerateKey(
	attributes doc.Composite,
) com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a new key pair",
	)

	// Make sure the digital notary has not been initialized.
	if uti.IsDefined(v.certificate_) {
		panic("The digital notary has already been initialized.")
	}

	// Generate a new key pair.
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)

	// Create the new certificate document.
	var tag = doc.Tag()         // Generate a new random tag.
	var version = doc.Version() // v1
	var previous doc.ResourceLike
	var identity = com.IdentityClass().Identity(
		algorithm,
		key,
		attributes,
		tag,
		version,
		previous,
	)
	var certificate = com.DocumentClass().Document(identity)

	// Notarize the document using its own key.
	v.notarizeDocument(certificate)
	v.certificate_ = certificate
	return certificate
}

func (v *digitalNotary_) RefreshKey() com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh the key pair",
	)

	// Make sure the digital notary has been initialized.
	if uti.IsUndefined(v.certificate_) {
		panic("The digital notary has not yet been initialized.")
	}

	// Generate a new key pair.
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)

	// Create the new certificate document.
	var content = v.certificate_.GetContent()
	var identity = com.IdentityClass().IdentityFromSource(
		content.AsSource(),
	)
	var attributes = identity.GetAttributes()
	var tag = content.GetTag()
	var version = doc.VersionClass().GetNextVersion(content.GetVersion(), 0)
	var previous = v.CiteDocument(v.certificate_).AsResource()
	var certificate = com.IdentityClass().Identity(
		algorithm,
		key,
		attributes,
		tag,
		version,
		previous,
	)
	var document = com.DocumentClass().Document(certificate)

	// Notarize the document using the previous key.
	v.notarizeDocument(document)
	v.certificate_ = document
	return document
}

func (v *digitalNotary_) ForgetKey() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to forget the private key",
	)

	// Erase the stored keys and certificate citation.
	v.certificate_ = nil
	v.hsm_.EraseKeys()
}

func (v *digitalNotary_) GenerateCredential(
	context any,
) com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a security credential",
	)

	// Make sure the digital notary has been initialized.
	if uti.IsUndefined(v.certificate_) {
		panic("The digital notary has not yet been initialized.")
	}

	// Create the credential document.
	var type_ = doc.Name("/bali/types/notary/Credential/v3")
	var tag = doc.Tag()
	var version = doc.Version()
	var permissions = doc.Name("/bali/permissions/Public/v3")
	var previous doc.ResourceLike
	var credential = com.ContentClass().Content(
		context,
		type_,
		tag,
		version,
		permissions,
		previous,
	)
	var document = com.DocumentClass().Document(credential)

	// Notarize the credential document.
	v.notarizeDocument(document)

	return document
}

func (v *digitalNotary_) RefreshCredential(
	context any,
	document com.DocumentLike,
) com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh a security credential",
	)

	// Make sure the digital notary has been initialized.
	if uti.IsUndefined(v.certificate_) {
		panic("The digital notary has not yet been initialized.")
	}

	// Create the next version of the credential document.
	var previous = v.CiteDocument(document).AsResource()
	var content = document.GetContent()
	var type_ = content.GetType()
	var tag = content.GetTag()
	var current = content.GetVersion()
	var version = doc.VersionClass().GetNextVersion(current, 0)
	var permissions = content.GetPermissions()
	var credential = com.ContentClass().Content(
		context,
		type_,
		tag,
		version,
		permissions,
		previous,
	)
	document = com.DocumentClass().Document(credential)

	// Notarize the credential document.
	v.notarizeDocument(document)

	return document
}

func (v *digitalNotary_) NotarizeDocument(
	document com.DocumentLike,
) {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to notarize a document",
	)

	// Make sure the digital notary has been initialized.
	if uti.IsUndefined(v.certificate_) {
		panic("The digital notary has not yet been initialized.")
	}

	// Notarize the document.
	v.notarizeDocument(document)
}

func (v *digitalNotary_) SealMatches(
	document com.DocumentLike,
	certificate com.DocumentLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to match a document seal",
	)

	// Compare the signature algorithms for the public certificate and SSM.
	var identity = com.IdentityClass().IdentityFromSource(
		certificate.GetContent().AsSource(),
	)
	var certificateAlgorithm = string(identity.GetAlgorithm().AsIntrinsic())
	var hsmAlgorithm = v.hsm_.GetSignatureAlgorithm()
	if certificateAlgorithm != hsmAlgorithm {
		var message = fmt.Sprintf(
			"The certificate algorithm %q is incompatible with the SSM algorithm %q.",
			certificateAlgorithm,
			hsmAlgorithm,
		)
		panic(message)
	}

	// Validate the seal on the notarized document.
	var publicKey = identity.GetKey()
	var seal = document.RemoveNotarySeal()
	var source = document.AsSource()
	var sourceBytes = []byte(source)
	document.SetNotarySeal(seal)
	var keyBytes = publicKey.AsIntrinsic()
	var signatureBytes = seal.GetSignature().AsIntrinsic()
	return v.hsm_.IsValid(keyBytes, sourceBytes, signatureBytes)
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

func (v *digitalNotary_) errorCheck(
	message string,
) {
	if e := recover(); e != nil {
		message = fmt.Sprintf(
			"DigitalNotary: %s:\n    %v",
			message,
			e,
		)
		panic(message)
	}
}

func (v *digitalNotary_) notarizeDocument(
	document com.DocumentLike,
) {
	// Check for new certificate document.
	var owner doc.TagLike
	var citation com.CitationLike
	if uti.IsDefined(v.certificate_) {
		owner = v.certificate_.GetContent().GetTag()
		citation = v.CiteDocument(v.certificate_)
	} else {
		owner = doc.Tag() // Generate a new random tag.
	}

	// Add the notary attribute to the document.
	var notary = com.NotaryClass().Notary(
		owner,
		citation,
	)
	document.AddNotary(notary)

	// Digitally sign the document.
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetNotarySeal(seal)
}

// Instance Structure

type digitalNotary_ struct {
	// Declare the instance attributes.
	ssm_         Trusted
	hsm_         Hardened
	certificate_ com.DocumentLike
}

// Class Structure

type digitalNotaryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func digitalNotaryClass() *digitalNotaryClass_ {
	return digitalNotaryClassReference_
}

var digitalNotaryClassReference_ = &digitalNotaryClass_{
	// Initialize the class constants.
}

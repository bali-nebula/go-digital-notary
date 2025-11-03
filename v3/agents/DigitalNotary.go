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
	owner doc.TagLike,
	ssm Trusted,
	hsm Hardened,
	certificate com.CitationLike,
) DigitalNotaryLike {
	if uti.IsUndefined(owner) {
		panic("The \"owner\" attribute is required by this class.")
	}
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}
	if uti.IsUndefined(certificate) {
		hsm.EraseKeys()
	}

	// Create the new digital notary.
	var instance = &digitalNotary_{
		// Initialize the instance attributes.
		owner_:       owner,
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

func (v *digitalNotary_) GenerateKey() com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a new key pair",
	)

	// Generate a new key pair.
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)

	// Create the new certificate.
	var tag = doc.Tag()         // Generate a new random tag.
	var version = doc.Version() // v1
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var previous doc.ResourceLike
	var certificate = com.CertificateClass().Certificate(
		algorithm,
		key,
		tag,
		version,
		previous,
	)
	var document = com.DocumentClass().Document(certificate)

	// Notarize the document using its own key.
	var owner = v.owner_
	var notary com.CitationLike
	document.SetNotary(owner, notary)
	var source = document.AsSource()
	bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	// Create a citation to the new certificate.
	algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(document.AsSource())
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	v.certificate_ = com.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)

	return document
}

func (v *digitalNotary_) RefreshKey() com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh the key pair",
	)

	// Generate a new key pair.
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)

	// Create the new certificate.
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var previous = v.getCertificate()
	var tag = previous.GetTag()
	var current = previous.GetVersion()
	var version = doc.VersionClass().GetNextVersion(current, 0)
	var certificate = com.CertificateClass().Certificate(
		algorithm,
		key,
		tag,
		version,
		previous.AsResource(),
	)
	var document = com.DocumentClass().Document(certificate)

	// Notarize the document using the previous key.
	var owner = v.owner_
	document.SetNotary(owner, previous)
	var source = document.AsSource()
	bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	// Create a citation to the new certificate.
	algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(document.AsSource())
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	v.certificate_ = com.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)

	return document
}

func (v *digitalNotary_) ForgetKey() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to forget the private key",
	)

	// Erase the stored keys and certificate citation.
	v.hsm_.EraseKeys()
	v.certificate_ = nil
}

func (v *digitalNotary_) GenerateCredential(
	context any,
) com.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a security credential",
	)

	// Create the credential.
	var tag = doc.Tag()
	var version = doc.Version()
	var previous doc.ResourceLike
	var credential = com.CredentialClass().Credential(
		context,
		tag,
		version,
		previous,
	)

	// Notarized the credential.
	var document = com.DocumentClass().Document(credential)
	var owner = v.owner_
	var notary = v.getCertificate()
	document.SetNotary(owner, notary)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

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

	// Create the next version of the credential.
	var previous = v.CiteDocument(document).AsResource()
	var content = document.GetContent()
	var tag = content.GetTag()
	var current = content.GetVersion()
	var version = doc.VersionClass().GetNextVersion(current, 0)
	var credential = com.CredentialClass().Credential(
		context,
		tag,
		version,
		previous,
	)

	// Notarized the credential.
	document = com.DocumentClass().Document(credential)
	var owner = v.owner_
	var notary = v.getCertificate()
	document.SetNotary(owner, notary)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	return document
}

func (v *digitalNotary_) NotarizeDocument(
	document com.DocumentLike,
) {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to notarize a document",
	)

	// Notarize the document.
	var owner = v.owner_
	var notary = v.getCertificate()
	document.SetNotary(owner, notary)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsSource()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = com.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)
}

func (v *digitalNotary_) SealMatches(
	document com.DocumentLike,
	certificate com.CertificateLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to match a document seal",
	)

	// Compare the signature algorithms for the public certificate and SSM.
	var certificateAlgorithm = string(certificate.GetAlgorithm().AsIntrinsic())
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
	var publicKey = certificate.GetKey()
	var seal = document.RemoveSeal()
	var source = document.AsSource()
	var sourceBytes = []byte(source)
	document.SetSeal(seal)
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

func (v *digitalNotary_) getCertificate() com.CitationLike {
	if uti.IsUndefined(v.certificate_) {
		panic("The digital notary has not yet been initialized.")
	}
	return v.certificate_
}

// Instance Structure

type digitalNotary_ struct {
	// Declare the instance attributes.
	owner_       doc.TagLike
	ssm_         Trusted
	hsm_         Hardened
	certificate_ com.CitationLike
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

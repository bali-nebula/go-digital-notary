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
	authority com.DocumentLike,
	ssm Trusted,
	hsm Hardened,
) DigitalNotaryLike {
	if uti.IsUndefined(authority) {
		panic("The \"authority\" attribute is required by this class.")
	}
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}
	if uti.IsUndefined(authority.GetOptionalNotary()) {
		hsm.EraseKeys()
	}

	// Create the new digital notary.
	var identity = com.IdentityClass().IdentityFromSource(
		authority.GetContent().AsSource(),
	)
	var owner = identity.GetTag()
	var instance = &digitalNotary_{
		// Initialize the instance attributes.
		authority_: authority,
		owner_:     owner,
		ssm_:       ssm,
		hsm_:       hsm,
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
	var document = com.DocumentClass().Document(
		com.CertificateClass().Certificate(
			algorithm,
			key,
			tag,
			version,
			previous,
		),
	)

	// Notarize the document using its own key.
	var citation com.CitationLike
	v.notarizeDocument(document, citation)
	citation = v.CiteDocument(document)
	var certificate = citation.AsResource()
	v.setAuthorityCertificate(certificate)
	v.notarizeDocument(v.authority_, citation)

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
	var previous = v.getAuthorityCertificate()
	var citation = com.CitationClass().CitationFromResource(previous)
	var tag = citation.GetTag()
	var current = citation.GetVersion()
	var version = doc.VersionClass().GetNextVersion(current, 0)
	var document = com.DocumentClass().Document(
		com.CertificateClass().Certificate(
			algorithm,
			key,
			tag,
			version,
			previous,
		),
	)

	// Notarize the document using the previous key.
	v.notarizeDocument(document, citation)
	citation = v.CiteDocument(document) // Cite the new certificate.
	var certificate = citation.AsResource()
	v.setAuthorityCertificate(certificate)
	v.notarizeDocument(v.authority_, citation)

	return document
}

func (v *digitalNotary_) ForgetKey() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to forget the private key",
	)

	// Erase the stored keys and certificate citation.
	v.hsm_.EraseKeys()
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

	// Notarized the credential document.
	var document = com.DocumentClass().Document(credential)
	var resource = v.getAuthorityCertificate()
	var citation = com.CitationClass().CitationFromResource(resource)
	v.notarizeDocument(document, citation)

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
	var resource = v.getAuthorityCertificate()
	var citation = com.CitationClass().CitationFromResource(resource)
	v.notarizeDocument(document, citation)

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
	var resource = v.getAuthorityCertificate()
	var citation = com.CitationClass().CitationFromResource(resource)
	v.notarizeDocument(document, citation)
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

func (v *digitalNotary_) setAuthorityCertificate(certificate doc.ResourceLike) {
	v.authority_.SetSubcomponent(
		certificate,
		doc.Symbol("$content"),
		doc.Symbol("$certificate"),
	)
}

func (v *digitalNotary_) getAuthorityCertificate() doc.ResourceLike {
	var subcomponent = v.authority_.GetSubcomponent(
		doc.Symbol("$content"),
		doc.Symbol("$certificate"),
	)
	if uti.IsUndefined(subcomponent) {
		panic("The digital notary has not yet been initialized.")
	}
	var certificate = doc.Resource(
		doc.FormatComponent(subcomponent.GetComposite()),
	)
	return certificate
}

func (v *digitalNotary_) notarizeDocument(
	document com.DocumentLike,
	citation com.CitationLike,
) {
	// Add the notary attribute to the document.
	var owner = v.owner_
	var notary = com.NotaryClass().Notary(
		owner,
		citation,
	)
	document.SetOptionalNotary(notary)

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
	authority_ com.DocumentLike
	owner_     doc.TagLike
	ssm_       Trusted
	hsm_       Hardened
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

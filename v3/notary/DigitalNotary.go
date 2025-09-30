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

package notary

import (
	byt "bytes"
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	not "github.com/bali-nebula/go-digital-notary/v3/documents"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func DigitalNotaryClass() DigitalNotaryClassLike {
	return digitalNotaryClass()
}

// Constructor Methods

func (c *digitalNotaryClass_) DigitalNotary(
	directory string,
	ssm Trusted,
	hsm Hardened,
) DigitalNotaryLike {
	if uti.IsUndefined(directory) {
		panic("The \"directory\" attribute is required by this class.")
	}
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}

	// Initialize the digital notary attributes.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += "notary/"
	uti.MakeDirectory(directory)
	var filename = directory + "Citation.bali"
	var account = doc.Tag(hsm.GetTag())
	if !uti.PathExists(filename) {
		// There is no way to retrieve a citation to the certificate.
		hsm.EraseKeys()
	}

	// Create the new digital notary.
	var instance = &digitalNotary_{
		// Initialize the instance attributes.
		directory_: directory,
		filename_:  filename,
		account_:   account,
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

func (v *digitalNotary_) GenerateKey() not.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a new key pair",
	)

	// Generate a new key pair.
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)

	// Create the new certificate.
	var account = v.account_
	var tag = doc.Tag()         // Generate a new random tag.
	var version = doc.Version() // v1
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var certificate = not.CertificateClass().Certificate(
		account,
		tag,
		version,
		algorithm,
		key,
	)
	var document = not.DocumentClass().Document(certificate)

	// Notarize the document using its own key.
	var notary not.CitationLike
	document.SetNotary(notary)
	var source = document.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = not.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	// Create a citation to the new certificate.
	algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(document.AsString())
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	var citation = not.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)

	// Save off the citation.
	source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	return document
}

func (v *digitalNotary_) RefreshKey() not.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh the key pair",
	)

	// Generate a new key pair.
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var key = doc.Binary(bytes)

	// Create the new certificate.
	var account = v.account_
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var previous = v.getCitation()
	var tag = previous.GetTag()
	var current = previous.GetVersion()
	var version = doc.VersionClass().GetNextVersion(current, 0)
	var certificate = not.CertificateClass().Certificate(
		account,
		tag,
		version,
		algorithm,
		key,
	)
	var document = not.DocumentClass().Document(certificate)

	// Notarize the document using the previous key.
	document.SetNotary(previous)
	var source = document.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = not.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	// Create a citation to the new certificate.
	algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(document.AsString())
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	var citation = not.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)

	// Save off the citation.
	source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	return document
}

func (v *digitalNotary_) ForgetKey() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to forget the private key",
	)

	// Erase the stored keys and certificate citation.
	v.hsm_.EraseKeys()
	uti.RemovePath(v.filename_)
}

func (v *digitalNotary_) GenerateCredential(
	tag doc.TagLike,
	version doc.VersionLike,
) not.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a security credential",
	)

	// Create the credential.
	var account = v.account_
	var credential = not.CredentialClass().Credential(
		account,
		tag,
		version,
	)

	// Notarized the credential.
	var document = not.DocumentClass().Document(
		credential,
	)
	var notary = v.getCitation()
	document.SetNotary(notary)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsString()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = not.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	return document
}

func (v *digitalNotary_) NotarizeDocument(
	document not.DocumentLike,
) not.DocumentLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to notarize a document",
	)

	// Notarize the document.
	var notary = v.getCitation()
	document.SetNotary(notary)
	var algorithm = doc.Quote(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = document.AsString()
	var bytes = v.hsm_.SignBytes([]byte(source))
	var signature = doc.Binary(bytes)
	var seal = not.SealClass().Seal(
		algorithm,
		signature,
	)
	document.SetSeal(seal)

	return document
}

func (v *digitalNotary_) SealMatches(
	document not.DocumentLike,
	certificate not.DocumentLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to match a document seal",
	)

	// Compare the signature algorithms for the public certificate and SSM.
	var content = not.CertificateClass().CertificateFromString(
		certificate.GetContent().AsString(),
	)
	var certificateAlgorithm = string(content.GetAlgorithm().AsIntrinsic())
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
	var publicKey = content.GetKey()
	var seal = document.RemoveSeal()
	var source = document.AsString()
	var sourceBytes = []byte(source)
	document.SetSeal(seal)
	var certificateBytes = publicKey.AsIntrinsic()
	var sealBytes = seal.GetSignature().AsIntrinsic()
	return v.ssm_.IsValid(certificateBytes, sealBytes, sourceBytes)
}

func (v *digitalNotary_) CiteDocument(
	document not.DocumentLike,
) not.CitationLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to create a citation to a document",
	)

	// Create a citation to the document.
	var algorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = document.AsString()
	var bytes = []byte(source)
	var digest = doc.Binary(v.ssm_.DigestBytes(bytes))
	var content = document.GetContent()
	var tag = content.GetTag()
	var version = content.GetVersion()
	var citation = not.CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)
	return citation
}

func (v *digitalNotary_) CitationMatches(
	citation not.CitationLike,
	document not.DocumentLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify a document citation",
	)

	// Compare the citation digest with a digest of the document.
	var citationAlgorithm = citation.GetAlgorithm()
	var citationDigest = citation.GetDigest()
	var ssmAlgorithm = doc.Quote(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = document.AsString()
	var bytes = []byte(source)
	var documentDigest = doc.Binary(v.ssm_.DigestBytes(bytes))
	if citationAlgorithm.AsString() != ssmAlgorithm.AsString() {
		return false
	}
	if !byt.Equal(citationDigest.AsIntrinsic(), documentDigest.AsIntrinsic()) {
		return false
	}
	return true
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

func (v *digitalNotary_) getCitation() not.CitationLike {
	if !uti.PathExists(v.filename_) {
		panic("The digital notary has not yet been initialized.")
	}
	var source = uti.ReadFile(v.filename_)
	var citation = not.CitationClass().CitationFromString(source)
	return citation
}

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

// Instance Structure

type digitalNotary_ struct {
	// Declare the instance attributes.
	directory_ string
	filename_  string
	account_   doc.TagLike
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

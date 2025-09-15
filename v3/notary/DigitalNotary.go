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
	doc "github.com/bali-nebula/go-digital-notary/v3/documents"
	fra "github.com/craterdog/go-component-framework/v7"
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
	var account = fra.TagFromString(hsm.GetTag())
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

func (v *digitalNotary_) GetCitation() fra.ResourceLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to retrieve the certificate citation",
	)

	if !uti.PathExists(v.filename_) {
		panic("The digital notary has not yet been initialized.")
	}
	var source = uti.ReadFile(v.filename_)
	var citation = doc.CitationClass().CitationFromString(source)
	return citation.AsResource()
}

func (v *digitalNotary_) GenerateKey() doc.CertificateLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a new key pair",
	)

	// Create a new key pair.
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var base64 = fra.Binary(bytes)
	var tag = fra.TagWithSize(20)
	var version = fra.VersionFromString("v1")
	var key = doc.KeyClass().Key(
		algorithm,
		base64,
		tag,
		version,
	)

	// Create a digest of the new public key.
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(key.AsString())
	base64 = fra.Binary(v.hsm_.DigestBytes(bytes))
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)

	// Create a citation to the new public key.
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		digest,
	)

	// Save off the citation.
	var source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	// Create the new certificate.
	var account = v.account_
	var signatory = citation.AsResource()
	var certificate = doc.CertificateClass().Certificate(
		key,
		account,
		signatory,
	)
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	source = certificate.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	base64 = fra.Binary(bytes)
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	certificate.SetSignature(signature)
	return certificate
}

func (v *digitalNotary_) RefreshKey() doc.CertificateLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh the key pair",
	)

	// Generate a new key pair.
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var base64 = fra.Binary(bytes)

	// Generate a the next version of the certificate.
	var previous = v.GetCitation()
	var citation = doc.CitationClass().CitationFromResource(previous)
	var tag = citation.GetTag()
	var current = citation.GetVersion()
	var version = fra.VersionClass().GetNextVersion(current, 0)
	var key = doc.KeyClass().Key(
		algorithm,
		base64,
		tag,
		version,
	)

	// Create a citation to the new public key.
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(key.AsString())
	base64 = fra.Binary(v.hsm_.DigestBytes(bytes))
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)
	citation = doc.CitationClass().Citation(
		tag,
		version,
		digest,
	)

	// Save off the citation.
	var source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	// Create the new certificate.
	var account = v.account_
	var signatory = previous
	var certificate = doc.CertificateClass().Certificate(
		key,
		account,
		signatory,
	)
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	source = certificate.AsString()
	bytes = v.hsm_.SignBytes([]byte(source)) // Signed using the previous key.
	base64 = fra.Binary(bytes)
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	certificate.SetSignature(signature)
	return certificate
}

func (v *digitalNotary_) ForgetKey() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to forget the private key",
	)

	v.hsm_.EraseKeys()
	uti.RemovePath(v.filename_)
}

func (v *digitalNotary_) GenerateCredential() doc.ContractLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a security credential",
	)

	// Create the credential document including timestamp component.
	var timestamp = fra.Now()
	var type_ = fra.ResourceFromString("<bali:/types/notary/Credential:v3>")
	var tag = fra.TagWithSize(20)
	var version = fra.VersionFromString("v1")
	var permissions = fra.ResourceFromString("<bali:/permissions/Public:v3>")
	var previous fra.ResourceLike
	var draft = doc.DraftClass().Draft(
		timestamp,
		type_,
		tag,
		version,
		permissions,
		previous,
	)

	// Digitally notarize the credential document.
	var citation = v.GetCitation()
	var contract = doc.ContractClass().Contract(
		draft,
		v.account_,
		citation,
	)
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = contract.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.hsm_.SignBytes(bytes))
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *digitalNotary_) NotarizeDraft(
	draft doc.Parameterized,
) doc.ContractLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to notarize a draft document",
	)

	// Digitally notarize the draft document.
	var citation = v.GetCitation()
	var contract = doc.ContractClass().Contract(
		draft,
		v.account_,
		citation,
	)
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var source = contract.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.hsm_.SignBytes(bytes))
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *digitalNotary_) SignatureMatches(
	document doc.Notarized,
	key doc.KeyLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to match a document signature",
	)

	// Validate the signature on the document using the public key.
	var keyAlgorithm = string(key.GetAlgorithm().AsIntrinsic())
	var ssmAlgorithm = v.ssm_.GetSignatureAlgorithm()
	if keyAlgorithm != ssmAlgorithm {
		var message = fmt.Sprintf(
			"The key signature algorithm %q is incompatible with the SSM algorithm %q.",
			keyAlgorithm,
			ssmAlgorithm,
		)
		panic(message)
	}
	var publicKey = key.GetBase64()
	var signature = document.GetSignature()
	document.RemoveSignature()
	var source = document.AsString()
	var sourceBytes = []byte(source)
	document.SetSignature(signature)
	var keyBytes = publicKey.AsIntrinsic()
	var signatureBytes = signature.GetBase64().AsIntrinsic()
	return v.ssm_.IsValid(keyBytes, signatureBytes, sourceBytes)
}

func (v *digitalNotary_) CiteDraft(
	draft doc.Parameterized,
) fra.ResourceLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to create a citation to a draft document",
	)

	// Create a citation to the draft document.
	var tag = draft.GetTag()
	var version = draft.GetVersion()
	var algorithm = fra.QuoteFromString(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = draft.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.ssm_.DigestBytes(bytes))
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		digest,
	)
	var resource = citation.AsResource()
	return resource
}

func (v *digitalNotary_) CitationMatches(
	citation fra.ResourceLike,
	draft doc.Parameterized,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify a document citation",
	)

	// Compare the citation digest with a digest of the draft document.
	var digest = doc.CitationClass().CitationFromResource(citation).GetDigest()
	var algorithm = fra.QuoteFromString(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = draft.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.ssm_.DigestBytes(bytes))
	if algorithm.AsString() != digest.GetAlgorithm().AsString() {
		return false
	}
	if !byt.Equal(base64.AsIntrinsic(), digest.GetBase64().AsIntrinsic()) {
		return false
	}
	return true
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

// Instance Structure

type digitalNotary_ struct {
	// Declare the instance attributes.
	directory_ string
	filename_  string
	account_   fra.TagLike
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

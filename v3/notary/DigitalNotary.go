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
		// There is no way to retrieve the citation to the certificate.
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

func (v *digitalNotary_) GenerateKey() doc.ContractLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate a new private key",
	)

	// Create a new certificate.
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var publicKey = fra.Binary(bytes)
	var tag = fra.TagWithSize(20)
	var version = fra.VersionFromString("v1")
	var previous = fra.PatternClass().None()
	var certificate = doc.CertificateClass().Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		previous,
	)

	// Create a digest of the new certificate.
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(certificate.AsString())
	var base64 = fra.Binary(v.hsm_.DigestBytes(bytes))
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)

	// Create a citation to the new certificate.
	var isNotarized = fra.Boolean(true)
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		isNotarized,
		digest,
	)

	// Save off the citation.
	var source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	// Digitally notarize the certificate.
	source = certificate.AsString()
	var draft = doc.DraftClass().DraftFromString(source)
	var contract = doc.ContractClass().Contract(
		draft,
		v.account_,
		citation,
	)
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	base64 = fra.Binary(bytes)
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *digitalNotary_) GetCitation() doc.CitationLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to retrieve the public certificate",
	)

	if !uti.PathExists(v.filename_) {
		panic("The digital notary has not yet been initialized.")
	}
	var source = uti.ReadFile(v.filename_)
	var citation = doc.CitationClass().CitationFromString(source)
	return citation
}

func (v *digitalNotary_) RefreshKey() doc.ContractLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to refresh the private key",
	)

	// Generate a new key pair.
	var algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var publicKey = fra.Binary(bytes)

	// Generate a the next version of the certificate.
	var citation = v.GetCitation()
	var tag = citation.GetTag()
	var current = citation.GetVersion()
	var version = fra.VersionClass().GetNextVersion(current, 0)
	var previous = citation.AsResource()
	var certificate = doc.CertificateClass().Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		previous,
	)

	// Create a citation to the new version of the certificate.
	var isNotarized = fra.Boolean(true)
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetDigestAlgorithm() + `"`)
	bytes = []byte(certificate.AsString())
	var base64 = fra.Binary(v.hsm_.DigestBytes(bytes))
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)
	citation = doc.CitationClass().Citation(
		tag,
		version,
		isNotarized,
		digest,
	)

	// Save off the citation.
	var source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	// Digitally notarize the certificate.
	source = certificate.AsString()
	var draft = doc.DraftClass().DraftFromString(source)
	var contract = doc.ContractClass().Contract(
		draft,
		v.account_,
		citation,
	)
	algorithm = fra.QuoteFromString(`"` + v.hsm_.GetSignatureAlgorithm() + `"`)
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	base64 = fra.Binary(bytes)
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
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
	var type_ = fra.ResourceFromString("<bali:/nebula/types/Credential:v3>")
	var tag = fra.TagWithSize(20)
	var version = fra.VersionFromString("v1")
	var permissions = fra.ResourceFromString("<bali:/nebula/permissions/public:v3>")
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
	draft doc.DraftLike,
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
	contract doc.ContractLike,
	certificate doc.CertificateLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to match a contract signature",
	)

	// Validate the signature on the contract using the public certificate.
	var certificateAlgorithm = string(certificate.GetAlgorithm().AsIntrinsic())
	var ssmAlgorithm = v.ssm_.GetSignatureAlgorithm()
	if certificateAlgorithm != ssmAlgorithm {
		var message = fmt.Sprintf(
			"The certificate signature algorithm %q is incompatible with the SSM algorithm %q.",
			certificateAlgorithm,
			ssmAlgorithm,
		)
		panic(message)
	}
	var publicKey = certificate.GetPublicKey()
	var signature = contract.GetSignature()
	contract.RemoveSignature()
	var source = contract.AsString()
	var sourceBytes = []byte(source)
	contract.SetSignature(signature)
	var keyBytes = publicKey.AsIntrinsic()
	var signatureBytes = signature.GetBase64().AsIntrinsic()
	return v.ssm_.IsValid(keyBytes, signatureBytes, sourceBytes)
}

func (v *digitalNotary_) CiteDraft(
	draft doc.DraftLike,
) doc.CitationLike {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to create a citation to a draft document",
	)

	// Create a citation to the draft document.
	var tag = draft.GetTag()
	var version = draft.GetVersion()
	var isNotarized = fra.Boolean(false)
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
		isNotarized,
		digest,
	)
	return citation
}

func (v *digitalNotary_) CitationMatches(
	citation doc.CitationLike,
	draft doc.DraftLike,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify a document citation",
	)

	// Compare the citation digest with a digest of the draft document.
	var algorithm = fra.QuoteFromString(`"` + v.ssm_.GetDigestAlgorithm() + `"`)
	var source = draft.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.ssm_.DigestBytes(bytes))
	if algorithm.AsString() != citation.GetDigest().GetAlgorithm().AsString() {
		return false
	}
	if base64.AsString() != citation.GetDigest().GetBase64().AsString() {
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

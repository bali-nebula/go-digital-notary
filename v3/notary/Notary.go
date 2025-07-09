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
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	bal "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NotaryClass() NotaryClassLike {
	return notaryClass()
}

// Constructor Methods

func (c *notaryClass_) Notary(
	ssm Trusted,
	hsm Hardened,
) NotaryLike {
	if uti.IsUndefined(ssm) {
		panic("The \"ssm\" attribute is required by this class.")
	}
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}

	// Initialize the notary attributes.
	var directory = uti.HomeDirectory()
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += ".bali/notary/"
	uti.MakeDirectory(directory)
	var filename = directory + "Citation.bali"
	var account = hsm.GetTag()
	if !uti.PathExists(filename) {
		// There is no way to retrieve the citation to the certificate.
		hsm.EraseKeys()
	}

	// Create the new notary.
	var instance = &notary_{
		// Initialize the instance attributes.
		filename_: filename,
		account_:  account,
		ssm_:      ssm,
		hsm_:      hsm,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *notary_) GetClass() NotaryClassLike {
	return notaryClass()
}

func (v *notary_) GenerateKey() doc.ContractLike {
	// Create a new certificate.
	var algorithm = v.hsm_.GetSignatureAlgorithm()
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var publicKey = fra.Binary(bytes).AsString()
	var tag = fra.TagWithSize(20).AsString()
	var version = "v1" // This is the first version of this certificate.
	var previous doc.CitationLike
	var certificate = doc.CertificateClass().Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		previous,
	)

	// Create a digest of the new certificate.
	algorithm = v.hsm_.GetDigestAlgorithm()
	bytes = []byte(certificate.AsString())
	var base64 = fra.Binary(v.hsm_.DigestBytes(bytes)).AsString()
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)

	// Create a citation to the new certificate.
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		digest,
	)

	// Save off the citation.
	var source = citation.AsString()
	uti.WriteFile(v.filename_, source)

	// Digitally notarize the certificate.
	source = certificate.AsString()
	var document = doc.DocumentClass().DocumentFromString(source)
	var contract = doc.ContractClass().Contract(
		document,
		v.account_,
		citation,
	)
	algorithm = v.hsm_.GetSignatureAlgorithm()
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	base64 = fra.Binary(bytes).AsString()
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) GetCitation() doc.CitationLike {
	if !uti.PathExists(v.filename_) {
		panic("The digital notary has not yet been initialized.")
	}
	var source = uti.ReadFile(v.filename_)
	var citation = doc.CitationClass().CitationFromString(source)
	return citation
}

func (v *notary_) RefreshKey() doc.ContractLike {
	// Generate a new key pair.
	var algorithm = v.hsm_.GetSignatureAlgorithm()
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var publicKey = fra.Binary(bytes).AsString()

	// Generate a the next version of the certificate.
	var citation = v.GetCitation()
	var tag = citation.GetTag()
	var version = citation.GetVersion()
	var current = fra.VersionFromString(version)
	version = fra.VersionClass().GetNextVersion(current, 0).AsString()
	var previous = citation
	var certificate = doc.CertificateClass().Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		previous,
	)

	// Create a citation to the new version of the certificate.
	algorithm = v.hsm_.GetDigestAlgorithm()
	bytes = []byte(certificate.AsString())
	var base64 = fra.Binary(v.hsm_.DigestBytes(bytes)).AsString()
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

	// Digitally notarize the certificate.
	source = certificate.AsString()
	var document = doc.DocumentClass().DocumentFromString(source)
	var contract = doc.ContractClass().Contract(
		document,
		v.account_,
		citation,
	)
	algorithm = v.hsm_.GetSignatureAlgorithm()
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	base64 = fra.Binary(bytes).AsString()
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) ForgetKey() {
	v.hsm_.EraseKeys()
	uti.RemovePath(v.filename_)
}

func (v *notary_) GenerateCredential() doc.ContractLike {
	// Create the credential document including timestamp component.
	var timestamp = fra.Now().AsString()
	var component = bal.Component(bal.Element(timestamp))
	var type_ = "<bali:/types/documents/Credential:v3>"
	var tag = fra.TagWithSize(20).AsString()
	var version = "v1"
	var permissions = "<bali:/permissions/Public:v3>"
	var previous doc.CitationLike
	var document = doc.DocumentClass().Document(
		component,
		type_,
		tag,
		version,
		permissions,
		previous,
	)

	// Digitally notarize the credential document.
	var citation = v.GetCitation()
	var contract = doc.ContractClass().Contract(
		document,
		v.account_,
		citation,
	)
	var algorithm = v.hsm_.GetSignatureAlgorithm()
	var source = contract.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.hsm_.SignBytes(bytes)).AsString()
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) NotarizeDocument(
	document doc.DocumentLike,
) doc.ContractLike {
	// Digitally notarize the document.
	var citation = v.GetCitation()
	var contract = doc.ContractClass().Contract(
		document,
		v.account_,
		citation,
	)
	var algorithm = v.hsm_.GetSignatureAlgorithm()
	var source = contract.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.hsm_.SignBytes(bytes)).AsString()
	var signature = doc.SignatureClass().Signature(
		algorithm,
		base64,
	)
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) SignatureMatches(
	contract doc.ContractLike,
	certificate doc.CertificateLike,
) bool {
	// Validate the signature on the contract using the public certificate.
	if certificate.GetAlgorithm() != v.ssm_.GetSignatureAlgorithm() {
		var message = fmt.Sprintf(
			"The certificate signature algorithm %q is incompatible with the SSM algorithm %q.",
			certificate.GetAlgorithm(),
			v.ssm_.GetSignatureAlgorithm(),
		)
		panic(message)
	}
	var publicKey = certificate.GetPublicKey()
	var signature = contract.GetSignature()
	contract.SetSignature(nil)
	var source = contract.AsString()
	var sourceBytes = []byte(source)
	contract.SetSignature(signature)
	var keyBytes = fra.BinaryFromString(publicKey).AsIntrinsic()
	var signatureBytes = fra.BinaryFromString(signature.GetBase64()).AsIntrinsic()
	return v.ssm_.IsValid(keyBytes, signatureBytes, sourceBytes)
}

func (v *notary_) CiteDocument(
	document doc.DocumentLike,
) doc.CitationLike {
	var tag = document.GetTag()
	var version = document.GetVersion()
	var algorithm = v.ssm_.GetDigestAlgorithm()
	var source = document.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.ssm_.DigestBytes(bytes)).AsString()
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		digest,
	)
	return citation
}

func (v *notary_) CitationMatches(
	citation doc.CitationLike,
	document doc.DocumentLike,
) bool {
	// Compare the citation digest with a digest of the record.
	var algorithm = v.ssm_.GetDigestAlgorithm()
	var source = document.AsString()
	var bytes = []byte(source)
	var base64 = fra.Binary(v.ssm_.DigestBytes(bytes)).AsString()
	var digest = doc.DigestClass().Digest(
		algorithm,
		base64,
	)
	if digest.GetAlgorithm() != citation.GetDigest().GetAlgorithm() {
		return false
	}
	if digest.GetBase64() != citation.GetDigest().GetBase64() {
		return false
	}
	return true
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type notary_ struct {
	// Declare the instance attributes.
	filename_ string
	account_  string
	ssm_      Trusted
	hsm_      Hardened
}

// Class Structure

type notaryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func notaryClass() *notaryClass_ {
	return notaryClassReference_
}

var notaryClassReference_ = &notaryClass_{
	// Initialize the class constants.
}

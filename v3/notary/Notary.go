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
	bal "github.com/bali-nebula/go-bali-documents/v3"
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssmv1"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func NotaryClass() NotaryClassLike {
	return notaryClass()
}

// Constructor Methods

func (c *notaryClass_) Notary(
	optionalDirectory string,
	hsm ssm.V1Secure,
) NotaryLike {
	if uti.IsUndefined(optionalDirectory) {
		optionalDirectory = uti.HomeDirectory()
	}
	var filename = optionalDirectory + "Citation.bali"
	if uti.IsUndefined(hsm) {
		panic("The \"hsm\" attribute is required by this class.")
	}
	var account = hsm.GetTag()

	// Initialize the modules catalog with all versions of software security
	// modules. The modules must be ordered with latest version first.
	var modules = fra.Catalog[string, ssm.V1Secure]()
	//modules.SetValue("v3", SsmV3(optionalDirectory))
	//modules.SetValue("v2", SsmV2(optionalDirectory))
	modules.SetValue("v1", hsm)

	// Replace the corresponding SSM with the HSM
	var protocol = hsm.GetProtocolVersion()
	modules.SetValue(protocol, hsm)

	// Create the new notary.
	var instance = &notary_{
		// Initialize the instance attributes.
		filename_: filename,
		account_:  account,
		protocol_: protocol,
		hsm_:      hsm,
		modules_:  modules,
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
	var digest =`"` + v.hsm_.GetDigestAlgorithm() + `"`
	var signature =`"` + v.hsm_.GetSignatureAlgorithm() + `"`
	var bytes = v.hsm_.GenerateKeys() // Returns the new public key.
	var key = fra.Binary(bytes).AsString()
	var tag = fra.TagWithSize(20).AsString()
	var version = "v1" // This is the first version of this certificate.
	var previous doc.CitationLike
	var certificate = doc.CertificateClass().Certificate(
		digest,
		signature,
		key,
		tag,
		version,
		previous,
	)

	// Create a citation to the new certificate.
	bytes = []byte(certificate.AsString())
	digest = fra.Binary(v.hsm_.DigestBytes(bytes)).AsString()
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		v.protocol_,
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
		v.protocol_,
		citation,
	)
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	signature = fra.Binary(bytes).AsString()
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
	var digest =`"` + v.hsm_.GetDigestAlgorithm() + `"`
	var signature =`"` + v.hsm_.GetSignatureAlgorithm() + `"`

	// Generate a new key pair.
	var bytes = v.hsm_.RotateKeys() // Returns the new public key.
	var key = fra.Binary(bytes).AsString()

	// Generate a the next version of the certificate.
	var citation = v.GetCitation()
	var tag = citation.GetTag()
	var version = citation.GetVersion()
	var current = fra.VersionFromString(version)
	version = fra.VersionClass().GetNextVersion(current, 0).AsString()
	var previous = citation
	var certificate = doc.CertificateClass().Certificate(
		digest,
		signature,
		key,
		tag,
		version,
		previous,
	)

	// Create a citation to the new version of the certificate.
	bytes = []byte(certificate.AsString())
	digest = fra.Binary(v.hsm_.DigestBytes(bytes)).AsString()
	citation = doc.CitationClass().Citation(
		tag,
		version,
		v.protocol_,
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
		v.protocol_,
		citation,
	)
	source = contract.AsString()
	bytes = v.hsm_.SignBytes([]byte(source))
	signature = fra.Binary(bytes).AsString()
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) ForgetKey() {
	v.hsm_.EraseKeys()
}

func (v *notary_) GenerateCredential() doc.ContractLike {
	// Create the credential document including timestamp component.
	var timestamp = fra.MomentClass().Now().AsString()
	var component = bal.Component(bal.Element(timestamp))
	var type_ = "<bali:/types/documents/Credential@v3>"
	var tag = fra.TagWithSize(20).AsString()
	var version = "v1"
	var permissions = "<bali:/permissions/Public@v3>"
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
		v.protocol_,
		citation,
	)
	var source = contract.AsString()
	var bytes = []byte(source)
	var signature = fra.Binary(v.hsm_.SignBytes(bytes)).AsString()
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
		v.protocol_,
		citation,
	)
	var source = contract.AsString()
	var bytes = []byte(source)
	var signature = fra.Binary(v.hsm_.SignBytes(bytes)).AsString()
	contract.SetSignature(signature)
	return contract
}

func (v *notary_) SignatureMatches(
	contract doc.ContractLike,
	certificate doc.CertificateLike,
) bool {
	// Retrieve the SSM that supports the required security protocol.
	var protocol = contract.GetProtocol()
	var ssm = v.modules_.GetValue(protocol)
	if ssm == nil {
		var message = fmt.Sprintf(
			"The required security protocol (%v) is not supported by this digital notary.\n",
			protocol)
		panic(message)
	}

	// Validate the signature on the contract using the public certificate.
	var key = certificate.GetKey()
	var signature = contract.GetSignature()
	contract.SetSignature("")
	var source = contract.AsString()
	var bytes = []byte(source)
	contract.SetSignature(signature)
	var keyBytes = fra.BinaryFromString(key).AsIntrinsic()
	var signatureBytes = fra.BinaryFromString(signature).AsIntrinsic()
	return ssm.IsValid(keyBytes, signatureBytes, bytes)
}

func (v *notary_) CiteDocument(
	document doc.DocumentLike,
) doc.CitationLike {
	var tag = document.GetTag()
	var version = document.GetVersion()
	var source = document.AsString()
	var bytes = []byte(source)
	var digest = fra.Binary(v.hsm_.DigestBytes(bytes)).AsString()
	var citation = doc.CitationClass().Citation(
		tag,
		version,
		v.protocol_,
		digest,
	)
	return citation
}

func (v *notary_) CitationMatches(
	citation doc.CitationLike,
	document doc.DocumentLike,
) bool {
	// Retrieve the SSM that supports the required security protocol.
	var protocol = citation.GetProtocol()
	var ssm = v.modules_.GetValue(protocol)
	if ssm == nil {
		var message = fmt.Sprintf(
			"The required security protocol (%v) is not supported by this digital notary.\n",
			protocol)
		panic(message)
	}

	// Compare the citation digest with a digest of the record.
	var source = document.AsString()
	var bytes = []byte(source)
	var digest = fra.Binary(ssm.DigestBytes(bytes)).AsString()
	return digest == citation.GetDigest()
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type notary_ struct {
	// Declare the instance attributes.
	filename_ string
	account_  string
	protocol_ string
	hsm_      ssm.V1Secure
	modules_  fra.CatalogLike[string, ssm.V1Secure]
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

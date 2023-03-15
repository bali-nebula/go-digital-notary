/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package notary

import (
	byt "bytes"
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	age "github.com/bali-nebula/go-component-framework/v2/agents"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/craterdog/go-collection-framework/v2"
)

// NOTARY INTERFACE

// This constructor creates a new digital notary. The notary may be used to
// digitally sign contract documents using the specified hardware security
// module (HSM). The notary may also be used to validate the signature on a
// contract that was signed using the current or any previous version of the
// security protocol used by digital notaries. The HSM will be used to validate
// all current version signatures while a software security module (SSM) will be
// used to validate previous version signatures.
func Notary(directory string, hsm SecurityModuleLike) NotaryLike {
	if hsm == nil {
		panic("A security module must be provided to the digital notary.")
	}

	// Initialize the modules catalog with all versions of software security
	// modules. The modules must be ordered with latest version first.
	var modules = col.Catalog[abs.VersionLike, SecurityModuleLike]()
	//modules.SetValue(v3, SSMv3(""))
	//modules.SetValue(v2, SSMv2(""))
	modules.SetValue(v1, SSMv1(""))

	// Replace the corresponding SSM with the HSM
	var protocol = bal.Version(hsm.GetProtocol())
	modules.SetValue(protocol, hsm)

	// Create the new notary.
	var account = bal.Tag(hsm.GetTag())
	var configurator = age.Configurator(directory, "citation.bali")
	return &notary{account, protocol, hsm, modules, configurator}
}

// NOTARY IMPLEMENTATION

// This type defines the structure and methods associated with a digital notary.
type notary struct {
	account      abs.TagLike
	protocol     abs.VersionLike
	hsm          SecurityModuleLike
	modules      col.CatalogLike[abs.VersionLike, SecurityModuleLike]
	configurator abs.ConfiguratorLike
}

// These constants define the supported versions of the security protocol.
var (
	v1 = bal.Version("v1")
	v2 = bal.Version("v2")
	v3 = bal.Version("v3")
)

// This constant captures the algorithms used in this version of the protocol.
var algorithms = bal.Catalog(`[
    $digest: "SHA512"
    $signature: "ED25519"
]`)

// These constants define the attribute names for the standard attribues.
var (
	accountAttribute     = bal.Symbol("$account")
	algorithmsAttribute  = bal.Symbol("$algorithms")
	certificateAttribute = bal.Symbol("$certificate")
	digestAttribute      = bal.Symbol("$digest")
	documentAttribute    = bal.Symbol("$document")
	keyAttribute         = bal.Symbol("$key")
	permissionsAttribute = bal.Symbol("$permissions")
	previousAttribute    = bal.Symbol("$previous")
	protocolAttribute    = bal.Symbol("$protocol")
	saltAttribute        = bal.Symbol("$salt")
	signatureAttribute   = bal.Symbol("$signature")
	timestampAttribute   = bal.Symbol("$timestamp")
	typeAttribute        = bal.Symbol("$type")
	versionAttribute     = bal.Symbol("$version")
)

// PRUDENT INTERFACE

// This method generates a new private notary key and returns the corresponding
// public notary certificate.
func (v *notary) GenerateKey() ContractLike {
	var key = bal.Binary(v.hsm.GenerateKeys()) // Returns the new public key.
	var tag = bal.NewTag()
	var version = v1              // The first version of this certificate.
	var previous CitationLike // No previous version.
	var certificate = Certificate(key, algorithms, tag, version, previous)
	var bytes = []byte(bal.FormatDocument(certificate))
	var digest = bal.Binary(v.hsm.DigestBytes(bytes))
	var citation = Citation(tag, version, v.protocol, digest)
	v.configurator.Store(bal.FormatDocument(citation))
	var contract = Contract(certificate, v.account, v.protocol, citation)
	bytes = bal.FormatDocument(contract)
	var signature = bal.Binary(v.hsm.SignBytes(bytes))
	contract.AddSignature(signature)
	return contract
}

// This method retrieves a citation to the public notary certificate for the
// current private notary key.
func (v *notary) GetCitation() CitationLike {
	if !v.configurator.Exists() {
		panic("The digital notary has not yet been initialized.")
	}
	var component = bal.ParseDocument(v.configurator.Load())
	return &citation{component}
}

// This method replaces an existing private notary key with a new one and
// returns the corresponding public notary certificate digitally signed by the
// old private notary key. The old private notary key is destroyed.
func (v *notary) RefreshKey() ContractLike {
	var citation = v.GetCitation()
	var key = bal.Binary(v.hsm.RotateKeys()) // Returns the new public key.
	var tag = citation.GetTag()
	var version = bal.NextVersion(citation.GetVersion())
	var previous = citation // Record the previous certificate citation.
	var certificate = Certificate(key, algorithms, tag, version, previous)
	var bytes = []byte(bal.FormatDocument(certificate))
	var digest = bal.Binary(v.hsm.DigestBytes(bytes))
	citation = Citation(tag, version, v.protocol, digest)
	v.configurator.Store(bal.FormatDocument(citation))
	var contract = Contract(certificate, v.account, v.protocol, previous)
	bytes = bal.FormatDocument(contract)
	var signature = bal.Binary(v.hsm.SignBytes(bytes))
	contract.AddSignature(signature)
	return contract
}

// This method destroys any existing private notary key.
func (v *notary) ForgetKey() {
	v.hsm.EraseKeys()
	v.configurator.Delete()
}

// CERTIFIED INTERFACE

// This method generates a new set of account credentials that can be used for
// authentication.
func (v *notary) GenerateCredentials(salt abs.BinaryLike) ContractLike {
	var citation = v.GetCitation()
	var credentials = Credentials(salt)
	var contract = Contract(credentials, v.account, v.protocol, citation)
	var bytes = bal.FormatDocument(contract)
	var signature = bal.Binary(v.hsm.SignBytes(bytes))
	contract.AddSignature(signature)
	return contract
}

// This method uses the current private notary key to notarized the specified
// document and returns the resulting contract.
func (v *notary) NotarizeDocument(document DocumentLike) ContractLike {
	var citation = v.GetCitation()
	var contract = Contract(document, v.account, v.protocol, citation)
	var bytes = bal.FormatDocument(contract)
	var signature = bal.Binary(v.hsm.SignBytes(bytes))
	contract.AddSignature(signature)
	return contract
}

// This method determines whether or not the signature on the specified contract
// is valid using the specified public notary certificate.
func (v *notary) SignatureMatches(contract ContractLike, certificate CertificateLike) bool {
	// Retrieve the SSM that supports the required security protocol.
	var protocol = contract.GetProtocol()
	var ssm = v.modules.GetValue(protocol)
	if ssm == nil {
		var message = fmt.Sprintf(
			"The required security protocol (%v) is not supported by this digital notary.\n",
			protocol)
		panic(message)
	}

	// Validate the signature on the contract using the public certificate.
	var key = certificate.GetKey()
	var signature = contract.RemoveSignature()
	var bytes = bal.FormatDocument(contract)
	contract.AddSignature(signature)
	var isValid = ssm.IsValid(key.AsArray(), signature.AsArray(), bytes)
	return isValid
}

// This method generates a citation to the specified document and returns that
// citation.
func (v *notary) CiteDocument(document DocumentLike) CitationLike {
	var tag = document.GetTag()
	var version = document.GetVersion()
	var bytes = bal.FormatDocument(document)
	var digest = bal.Binary(v.hsm.DigestBytes(bytes))
	var citation = Citation(tag, version, v.protocol, digest)
	return citation
}

// This method determines whether or not the specified document citation matches
// the specified document.
func (v *notary) CitationMatches(citation CitationLike, document DocumentLike) bool {
	// Retrieve the SSM that supports the required security protocol.
	var protocol = citation.GetProtocol()
	var ssm = v.modules.GetValue(protocol)
	if ssm == nil {
		var message = fmt.Sprintf(
			"The required security protocol (%v) is not supported by this digital notary.\n",
			protocol)
		panic(message)
	}

	// Compare the citation digest with a digest of the document.
	var bytes = bal.FormatDocument(document)
	var digest = bal.Binary(ssm.DigestBytes(bytes))
	return byt.Equal(digest.AsArray(), citation.GetDigest().AsArray())
}

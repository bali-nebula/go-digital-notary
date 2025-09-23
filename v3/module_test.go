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

package module_test

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	not "github.com/bali-nebula/go-digital-notary/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const directory = "./test/"

func TestParsingCitations(t *tes.T) {
	var filename = directory + "documents/Citation.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var citation = not.Citation(source)
	var algorithm = citation.GetAlgorithm()
	var digest = citation.GetDigest()
	var tag = citation.GetTag()
	var version = citation.GetVersion()
	citation = not.Citation(
		algorithm,
		digest,
		tag,
		version,
	)
	var formatted = citation.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCredentials(t *tes.T) {
	var filename = directory + "documents/Credential.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var credential = not.Credential(source)
	var timestamp = credential.GetTimestamp()
	var tag = credential.GetTag()
	var version = credential.GetVersion()
	credential = not.Credential(
		tag,
		version,
	)
	credential.SetObject(timestamp, doc.Symbol("$timestamp"))
	var formatted = credential.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCertificates(t *tes.T) {
	var filename = directory + "documents/Certificate.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var certificate = not.Certificate(source)
	var timestamp = certificate.GetTimestamp()
	var algorithm = certificate.GetAlgorithm()
	var key = certificate.GetKey()
	var tag = certificate.GetTag()
	var version = certificate.GetVersion()
	certificate = not.Certificate(
		algorithm,
		key,
		tag,
		version,
	)
	certificate.SetObject(timestamp, doc.Symbol("$timestamp"))
	var formatted = certificate.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingDrafts(t *tes.T) {
	var filename = directory + "documents/Draft.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var draft = not.Draft(source)
	var entity = draft.GetEntity()
	var type_ = draft.GetType()
	var tag = draft.GetTag()
	var version = draft.GetVersion()
	var permissions = draft.GetPermissions()
	var optionalPrevious = draft.GetOptionalPrevious()
	draft = not.Draft(
		entity,
		type_,
		tag,
		version,
		permissions,
		optionalPrevious,
	)
	var formatted = draft.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingContracts(t *tes.T) {
	var filename = directory + "documents/Contract.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var contract = not.Contract(source)
	var content = contract.GetContent()
	var account = contract.GetAccount()
	var notary = contract.GetNotary()
	var seal = contract.RemoveSeal()
	contract = not.Contract(
		content,
		account,
		notary,
	)
	contract.SetSeal(seal)
	var formatted = contract.AsString()
	ass.Equal(t, source, formatted)
}

// Create the security module and digital notary.
var module = not.Ssm(directory)
var notary = not.DigitalNotary(directory, module, module)

func TestSSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	ass.Equal(t, "SHA512", module.GetDigestAlgorithm())
	ass.Equal(t, "ED25519", module.GetSignatureAlgorithm())
	ass.Equal(t, 64, len(module.DigestBytes(bytes)))

	module.EraseKeys()
	var publicKey = module.GenerateKeys()

	var seal = module.SignBytes(bytes)
	ass.True(t, module.IsValid(publicKey, seal, bytes))

	var newPublicKey = module.RotateKeys()
	seal = module.SignBytes(newPublicKey)
	ass.True(t, module.IsValid(publicKey, seal, newPublicKey))

	module.EraseKeys()
}

func TestDigitalNotaryInitialization(t *tes.T) {
	// Should not be able to retrieve the certificate citation without any keys.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "DigitalNotary: An error occurred while attempting to generate a security credential:\n    The digital notary has not yet been initialized.", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
		notary.ForgetKey()
	}()
	notary.ForgetKey()
	var tag = doc.Tag()
	var version = doc.Version()
	notary.GenerateCredential(tag, version)
}

func TestDigitalNotaryGenerateKey(t *tes.T) {
	// Generate a new public-private key pair.
	notary.ForgetKey()
	notary.GenerateKey()

	// Should not be able to do this twice.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(
				t,
				"DigitalNotary: An error occurred while attempting to generate a new key pair:\n    Ssm: An error occurred while attempting to generate new keys:\n        Attempted to transition from state \"$LoneKey\" to an invalid state on event \"$GenerateKeys\".",
				message,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GenerateKey()
}

func TestDigitalNotaryLifecycle(t *tes.T) {
	// Generate and validate a new public-private key pair.
	notary.ForgetKey()
	var certificateV1 = notary.GenerateKey()
	var content = certificateV1.GetContent()
	var keyV1 = not.Certificate(content.AsString())
	ass.True(
		t,
		notary.SealMatches(
			certificateV1,
			keyV1,
		),
	)
	var filename = "./test/notary/CertificateV1.bali"
	var source = certificateV1.AsString()
	uti.WriteFile(filename, source)

	// Create and cite a new transaction document.
	var timestamp = doc.Moment().AsString()
	var transaction = not.Draft(
		`[
    $timestamp: ` + timestamp + `
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com/>
    $amount: 7.95($currency: $USD)
](
	$type: <bali:/examples/Transaction:v1.2.3>
	$tag: #BCASYZR1MC2J2QDPL03HG42W0M7P36TQ
	$version: v1
	$permissions: <bali:/permissions/Public:v3>
	$previous: none
)`,
	)
	filename = "./test/notary/Draft.bali"
	source = transaction.AsString()
	uti.WriteFile(filename, source)

	var citation = notary.CiteDocument(transaction)
	ass.True(t, notary.CitationMatches(citation, transaction))

	// Notarize the transaction document to create a notarized contract.
	var contract = notary.NotarizeDocument(transaction)
	ass.True(
		t,
		notary.SealMatches(
			contract,
			keyV1,
		),
	)
	filename = "./test/notary/Contract.bali"
	source = contract.AsString()
	uti.WriteFile(filename, source)

	// Pickup where we left off with a new security module and digital notary.
	module = not.Ssm(directory)
	notary = not.DigitalNotary(directory, module, module)

	// Refresh and validate the public-private key pair.
	var certificateV2 = notary.RefreshKey()
	ass.True(
		t,
		notary.SealMatches(
			certificateV2,
			keyV1,
		),
	)
	filename = "./test/notary/CertificateV2.bali"
	source = certificateV2.AsString()
	uti.WriteFile(filename, source)

	// Generate an authentication credential.
	var tag = doc.Tag()
	var version = doc.Version()
	var credential = notary.GenerateCredential(tag, version)
	content = certificateV2.GetContent()
	var keyV2 = not.Certificate(content.AsString())
	ass.True(
		t,
		notary.SealMatches(
			credential,
			keyV2,
		),
	)
	filename = "./test/notary/Credential.bali"
	source = credential.AsString()
	uti.WriteFile(filename, source)
}

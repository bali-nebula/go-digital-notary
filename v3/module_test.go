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
	var tag = citation.GetTag()
	var version = citation.GetVersion()
	var algorithm = citation.GetAlgorithm()
	var digest = citation.GetDigest()
	citation = not.Citation(
		tag,
		version,
		algorithm,
		digest,
	)
	var formatted = citation.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCredentials(t *tes.T) {
	var filename = directory + "documents/Credential.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var credential = not.Credential(source)
	credential.GetTag()
	credential.GetVersion()
	var formatted = credential.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCertificates(t *tes.T) {
	var filename = directory + "documents/Certificate.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var certificate = not.Certificate(source)
	var tag = certificate.GetTag()
	var version = certificate.GetVersion()
	var algorithm = certificate.GetAlgorithm()
	var key = certificate.GetKey()
	certificate = not.Certificate(
		tag,
		version,
		algorithm,
		key,
	)
	var formatted = certificate.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingContents(t *tes.T) {
	var filename = directory + "documents/Content.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var content = not.Content(source)
	var entity = content.GetEntity()
	var type_ = content.GetType()
	var tag = content.GetTag()
	var version = content.GetVersion()
	var optionalPrevious = content.GetOptionalPrevious()
	var permissions = content.GetPermissions()
	content = not.Content(
		entity,
		type_,
		tag,
		version,
		optionalPrevious,
		permissions,
	)
	var formatted = content.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingDocuments(t *tes.T) {
	var filename = directory + "documents/Document.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var document = not.Document(source)
	document.GetContent()
	document.GetTimestamp()
	document.GetNotary()
	var seal = document.RemoveSeal()
	document.SetSeal(seal)
	var formatted = document.AsString()
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
	var context = doc.Moment()
	notary.GenerateCredential(context)
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
	ass.True(
		t,
		notary.SealMatches(
			certificateV1,
			certificateV1,
		),
	)
	var filename = "./test/notary/CertificateV1.bali"
	var source = certificateV1.AsString()
	uti.WriteFile(filename, source)

	// Create and cite a new transaction document.
	var timestamp = doc.Moment().AsString()
	var transaction = not.Content(
		`[
    $timestamp: ` + timestamp + `
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com/>
    $amount: 7.95($currency: $USD)
](
	$type: <bali:/examples/Transaction:v1.2.3>
	$tag: #BCASYZR1MC2J2QDPL03HG42W0M7P36TQ
	$version: v1
	$previous: none
	$permissions: <bali:/permissions/Public:v3>
)`,
	)
	filename = "./test/notary/Content.bali"
	source = transaction.AsString()
	uti.WriteFile(filename, source)

	var document = not.Document(transaction)
	var citation = notary.CiteDocument(document)
	ass.True(t, notary.CitationMatches(citation, document))

	// Notarize the transaction document to create a notarized document.
	notary.NotarizeDocument(document)
	ass.True(
		t,
		notary.SealMatches(
			document,
			certificateV1,
		),
	)
	filename = "./test/notary/Document.bali"
	source = document.AsString()
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
			certificateV1,
		),
	)
	filename = "./test/notary/CertificateV2.bali"
	source = certificateV2.AsString()
	uti.WriteFile(filename, source)

	// Generate an authentication credential.
	var context = doc.Moment()
	var credential = notary.GenerateCredential(context)
	ass.True(
		t,
		notary.SealMatches(
			credential,
			certificateV2,
		),
	)
	credential = notary.RefreshCredential(credential, context)
	ass.True(
		t,
		notary.SealMatches(
			credential,
			certificateV2,
		),
	)
	filename = "./test/notary/Credential.bali"
	source = credential.AsString()
	uti.WriteFile(filename, source)
}

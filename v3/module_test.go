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
	bal "github.com/bali-nebula/go-bali-documents/v3"
	not "github.com/bali-nebula/go-digital-notary/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

// Create the security module and digital notary.
const directory = "./test/"
var module = not.SsmV1(directory)
var notary = not.Notary(directory, module)

func TestParsingCitations(t *tes.T) {
	var filename = directory + "Citation.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var citation = not.CitationFromString(source)
	var formatted = citation.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCertificates(t *tes.T) {
	var filename = directory + "Certificate.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var certificate = not.CertificateFromString(source)
	var formatted = certificate.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingDocuments(t *tes.T) {
	var filename = directory + "Document.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var document = not.DocumentFromString(source)
	var formatted = document.AsString()
	ass.Equal(t, source, formatted)
	var attribute = not.DocumentClass().ExtractAttribute(
		"$consumer",
		bal.ParseSource(source),
	)
	ass.Equal(t, `"Derk Norton"`, attribute)
}

func TestParsingContracts(t *tes.T) {
	var filename = directory + "Contract.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var contract = not.ContractFromString(source)
	var formatted = contract.AsString()
	ass.Equal(t, source, formatted)
}

func TestSSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	ass.Equal(t, "v1", module.GetProtocolVersion())
	ass.Equal(t, "SHA512", module.GetDigestAlgorithm())
	ass.Equal(t, "ED25519", module.GetSignatureAlgorithm())
	ass.Equal(t, 64, len(module.DigestBytes(bytes)))

	var publicKey = module.GenerateKeys()

	var signature = module.SignBytes(bytes)
	ass.True(t, module.IsValid(publicKey, signature, bytes))

	var newPublicKey = module.RotateKeys()
	signature = module.SignBytes(newPublicKey)
	ass.True(t, module.IsValid(publicKey, signature, newPublicKey))

	module.EraseKeys()
}

func TestNotaryInitialization(t *tes.T) {
	// Should not be able to retrieve the certificate citation without any keys.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "The digital notary has not yet been initialized.", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
		notary.ForgetKey()
	}()
	notary.ForgetKey()
	uti.RemovePath(directory + "Citation.bali")
	notary.GetCitation()
}

func TestNotaryGenerateKey(t *tes.T) {
	// Generate a new public-private key pair.
	notary.GenerateKey()

	// Should not be able to do this twice.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "Attempted to transition to an invalid state: 0", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GenerateKey()
}

func TestNotaryLifecycle(t *tes.T) {
	// Generate and validate a new public-private key pair.
	notary.ForgetKey()
	var contractV1 = notary.GenerateKey()
	uti.WriteFile(directory + "CertificateV1.bali", contractV1.AsString())
	var certificateV1 = not.CertificateFromString(
		contractV1.GetDocument().AsString(),
	)
	ass.True(
		t,
		notary.SignatureMatches(
			contractV1,
			certificateV1,
		),
	)

	// Extract the citation to the public certificate.
	notary.GetCitation()

	// Create and cite a new transaction document.
	var transaction = not.DocumentFromString(
		`[
    $timestamp: <2022-06-03T07:39:54>
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com>
    $amount: 4.95($currency: $USD)
](
	$type: <bali:/examples/Record@v1.2.3>
	$tag: #BCASYZR1MC2J2QDPL03HG42W0M7P36TQ
	$version: v1
	$permissions: <bali:/permissions/Public@v1>
)`,
)
	uti.WriteFile(directory + "/Transaction.bali", transaction.AsString())
	var citation = notary.CiteDocument(transaction)
	ass.True(t, notary.CitationMatches(citation, transaction))
	uti.WriteFile(directory + "/Citation.bali", citation.AsString())

	// Notarize the transaction document to create a signed contract.
	var contract = notary.NotarizeDocument(transaction)
	uti.WriteFile(directory + "Contract.bali", contract.AsString())
	ass.True(
		t,
		notary.SignatureMatches(
			contract,
			certificateV1,
		),
	)

	// Pickup where we left off with a new security module and digital notary.
	module = not.SsmV1(directory)
	notary = not.Notary(directory, module)

	// Refresh and validate the public-private key pair.
	var contractV2 = notary.RefreshKey()
	uti.WriteFile(directory + "CertificateV2.bali", contractV2.AsString())
	ass.True(
		t,
		notary.SignatureMatches(
			contractV2,
			certificateV1,
		),
	)

	// Generate an authentication credential.
	var credential = notary.GenerateCredential()
	uti.WriteFile(directory + "Credential.bali", credential.AsString())
	var certificateV2 = not.CertificateFromString(
		contractV2.GetDocument().AsString(),
	)
	ass.True(
		t,
		notary.SignatureMatches(
			credential,
			certificateV2,
		),
	)

	// Reset the security module and digital notary to an uninitialized state.
	notary.ForgetKey()
	notary.GetCitation()
}

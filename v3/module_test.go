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
	var citation = not.CitationFromString(source)
	var formatted = citation.AsString()
	ass.Equal(t, source, formatted)
	citation = not.Citation(
		citation.GetTag(),
		citation.GetVersion(),
		citation.IsNotarized(),
		citation.GetDigest(),
	)
	source = citation.AsString()
	uti.WriteFile(filename, source)
}

func TestParsingCertificates(t *tes.T) {
	var filename = directory + "documents/Certificate.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var certificate = not.CertificateFromString(source)
	var formatted = certificate.AsString()
	ass.Equal(t, source, formatted)
	certificate = not.Certificate(
		certificate.GetAlgorithm(),
		certificate.GetPublicKey(),
		certificate.GetTag(),
		certificate.GetVersion(),
		certificate.GetPrevious(),
	)
	source = certificate.AsString()
	uti.WriteFile(filename, source)
}

func TestParsingDrafts(t *tes.T) {
	var filename = directory + "documents/Draft.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var draft = not.DraftFromString(source)
	var formatted = draft.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingContracts(t *tes.T) {
	var filename = directory + "documents/Contract.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var contract = not.ContractFromString(source)
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

	var signature = module.SignBytes(bytes)
	ass.True(t, module.IsValid(publicKey, signature, bytes))

	var newPublicKey = module.RotateKeys()
	signature = module.SignBytes(newPublicKey)
	ass.True(t, module.IsValid(publicKey, signature, newPublicKey))

	module.EraseKeys()
}

func TestDigitalNotaryInitialization(t *tes.T) {
	// Should not be able to retrieve the certificate citation without any keys.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "DigitalNotary: An error occurred while attempting to retrieve the public certificate:\n    The digital notary has not yet been initialized.", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
		notary.ForgetKey()
	}()
	notary.ForgetKey()
	notary.GetCitation()
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
				"DigitalNotary: An error occurred while attempting to generate a new private key:\n    Ssm: An error occurred while attempting to generate new keys:\n        Attempted to transition from state \"$LoneKey\" to an invalid state on event \"$GenerateKeys\".",
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
	var contractV1 = notary.GenerateKey()
	var certificateV1 = not.CertificateFromString(
		contractV1.GetDraft().AsString(),
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
	var transaction = not.DraftFromString(
		`[
    $timestamp: <2022-06-03T07:39:54>
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com>
    $amount: 4.95($currency: $USD)
](
	$type: <bali:/examples/Record:v1.2.3>
	$tag: #BCASYZR1MC2J2QDPL03HG42W0M7P36TQ
	$version: v1
	$permissions: <bali:/permissions/public:v3>
)`,
	)
	var citation = notary.CiteDraft(transaction)
	ass.True(t, notary.CitationMatches(citation, transaction))

	// Notarize the transaction document to create a signed contract.
	var contract = notary.NotarizeDraft(transaction)
	ass.True(
		t,
		notary.SignatureMatches(
			contract,
			certificateV1,
		),
	)

	// Pickup where we left off with a new security module and digital notary.
	module = not.Ssm(directory)
	notary = not.DigitalNotary(directory, module, module)

	// Refresh and validate the public-private key pair.
	var contractV2 = notary.RefreshKey()
	ass.True(
		t,
		notary.SignatureMatches(
			contractV2,
			certificateV1,
		),
	)

	// Generate an authentication credential.
	var credential = notary.GenerateCredential()
	var certificateV2 = not.CertificateFromString(
		contractV2.GetDraft().AsString(),
	)
	ass.True(
		t,
		notary.SignatureMatches(
			credential,
			certificateV2,
		),
	)
}

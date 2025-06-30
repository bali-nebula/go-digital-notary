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

package notary_test

import (
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssmv1"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const directory = "../test/"

// Create the security module and digital notary.
var module = ssm.SsmV1Class().SsmV1(directory)
var notary = not.NotaryClass().Notary(directory, module)

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
	var certificateV1 = doc.CertificateClass().CertificateFromString(
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
	var transaction = doc.DocumentClass().DocumentFromString(
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
	module = ssm.SsmV1Class().SsmV1(directory)
	notary = not.NotaryClass().Notary(directory, module)

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
	var certificateV2 = doc.CertificateClass().CertificateFromString(
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

	// Confirm that an error is raised if we try to retrieve the certificate citation.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "The digital notary has not yet been initialized.", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GetCitation()
}

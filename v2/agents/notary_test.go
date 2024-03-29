/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
	age "github.com/bali-nebula/go-digital-notary/v2/agents"
	doc "github.com/bali-nebula/go-digital-notary/v2/documents"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const directory = "./"

func TestNotaryInitialization(t *tes.T) {
	// Initialize the security module and digital notary.
	var module = age.SSMv1(directory)
	var notary = age.Notary(directory, module)

	// Should not be able to retrieve the certificate citation without any keys.
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

func TestNotaryGenerateKey(t *tes.T) {
	// Initialize the security module and digital notary.
	var module = age.SSMv1(directory)
	var notary = age.Notary(directory, module)

	// Generate a new public-private key pair.
	notary.GenerateKey()

	// Should not be able to do this twice.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "Attempted to transition to an invalid state: 0", message)
			notary.ForgetKey()
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GenerateKey()
}

func TestNotaryLifecycle(t *tes.T) {
	// Initialize the security module and digital notary.
	var module = age.SSMv1(directory)
	var notary = age.Notary(directory, module)

	// Generate and validate a new public-private key pair.
	var certificateV1 = notary.GenerateKey()
	_ = osx.WriteFile("../examples/certificateV1.bali", bal.FormatDocument(certificateV1), 0600)
	ass.True(t, notary.SignatureMatches(certificateV1, certificateV1.GetComponent().(abs.CertificateLike)))

	// Extract the citation to the public certificate.
	_ = notary.GetCitation()

	// Create and cite a new transaction record.
	var attributes = bal.Catalog(`[
    $timestamp: <2022-06-03T07:39:54>
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com/>
    $amount: 4.95($currency: $USD)
]`)
	var type_ = doc.Type(bal.Moniker("/bali/examples/Record/v1.2.3"), nil)
	var tag = bal.NewTag()
	var version = bal.Version("v1")
	var permissions = bal.Moniker("/bali/permissions/public/v1")
	var previous abs.CitationLike
	var record = doc.Record(attributes, type_, tag, version, permissions, previous)
	_ = osx.WriteFile("../examples/record.bali", bal.FormatDocument(record), 0600)
	var citation = notary.CiteRecord(record)
	ass.True(t, notary.CitationMatches(citation, record))
	_ = osx.WriteFile("../examples/citation.bali", bal.FormatDocument(citation), 0600)

	// Notarize the transaction record to create a signed contract.
	var contract = notary.NotarizeComponent(record)
	_ = osx.WriteFile("../examples/contract.bali", bal.FormatDocument(contract), 0600)
	ass.True(t, notary.SignatureMatches(contract, certificateV1.GetComponent().(abs.CertificateLike)))

	// Pickup where we left off with a new security module and digital notary.
	module = age.SSMv1(directory)
	notary = age.Notary(directory, module)

	// Refresh and validate the public-private key pair.
	var certificateV2 = notary.RefreshKey()
	_ = osx.WriteFile("../examples/certificateV2.bali", bal.FormatDocument(certificateV2), 0600)
	ass.True(t, notary.SignatureMatches(certificateV2, certificateV1.GetComponent().(abs.CertificateLike)))

	// Generate an authentication credential.
	var salt = bal.Binary(64)
	var credential = notary.GenerateCredential(salt)
	_ = osx.WriteFile("../examples/credential.bali", bal.FormatDocument(credential), 0600)
	ass.True(t, notary.SignatureMatches(credential, certificateV2.GetComponent().(abs.CertificateLike)))

	// Reset the security module and digital notary to an uninitialized state.
	notary.ForgetKey()

	// Confirm that an error is raised if we try to retrieve the certificate citation.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.True(t, sts.HasPrefix(message, "The digital notary has not yet been initialized."))
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GetCitation()
}

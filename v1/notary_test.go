/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package notary_test

import (
	abs "github.com/bali-nebula/go-component-framework/v1/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v1/bali"
	doc "github.com/bali-nebula/go-component-framework/v1/documents"
	not "github.com/bali-nebula/go-digital-notary/v1"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const directory = "./"

func TestNotaryInitialization(t *tes.T) {
	// Initialize the security module and digital notary.
	var module = not.SSMv1(directory)
	var notary = doc.Notary(directory, module)

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
	var module = not.SSMv1(directory)
	var notary = doc.Notary(directory, module)

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
	var module = not.SSMv1(directory)
	var notary = doc.Notary(directory, module)

	// Generate and validate a new public-private key pair.
	var certificateV1 = notary.GenerateKey()
	osx.WriteFile("./examples/certificateV1.bali", bal.FormatDocument(certificateV1), 0600)
	ass.True(t, notary.SignatureMatches(certificateV1, certificateV1.GetDocument().(not.CertificateLike)))

	// Extract the citation to the public certificate.
	var citation = notary.GetCitation()
	osx.WriteFile("./examples/citation.bali", bal.FormatDocument(citation), 0600)

	// Create and cite a new transaction document.
	var attributes = bal.Catalog(`[
    $timestamp: <2022-06-03T07:39:54>
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com/>
    $amount: 4.95($currency: $USD)
]`)
	var type_ = bal.Component("/bali/examples/Document/v1.2.3")
	var tag = bal.NewTag()
	var version = bal.Version("v1.2")
	var permissions = bal.Moniker("/bali/permissions/public/v1")
	var previous not.CitationLike
	var document = doc.Document(attributes, type_, tag, version, permissions, previous)
	osx.WriteFile("./examples/document.bali", bal.FormatDocument(document), 0600)
	citation = notary.CiteDocument(document)
	ass.True(t, notary.CitationMatches(citation, document))

	// Notarize the transaction document to create a signed contract.
	var contract = notary.NotarizeDocument(document)
	osx.WriteFile("./examples/contract.bali", bal.FormatDocument(contract), 0600)
	ass.True(t, notary.SignatureMatches(contract, certificateV1.GetDocument().(not.CertificateLike)))

	// Pickup where we left off with a new security module and digital notary.
	module = not.SSMv1(directory)
	notary = doc.Notary(directory, module)

	// Refresh and validate the public-private key pair.
	var certificateV2 = notary.RefreshKey()
	osx.WriteFile("./examples/certificateV2.bali", bal.FormatDocument(certificateV2), 0600)
	ass.True(t, notary.SignatureMatches(certificateV2, certificateV1.GetDocument().(not.CertificateLike)))

	// Generate some authentication credentials.
	var salt = bal.Binary(64)
	var credentials = notary.GenerateCredentials(salt)
	osx.WriteFile("./examples/credentials.bali", bal.FormatDocument(credentials), 0600)
	ass.True(t, notary.SignatureMatches(credentials, certificateV2.GetDocument().(not.CertificateLike)))

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

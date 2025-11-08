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
	uti "github.com/craterdog/go-essential-utilities/v8"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingCitations(t *tes.T) {
	var filename = testDirectory + "components/Citation.bali"
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
	var formatted = citation.AsSource()
	ass.Equal(t, source, formatted)
	citation = not.Citation(citation.AsResource())
	formatted = citation.AsSource()
	ass.Equal(t, source, formatted)
}

func TestParsingCredentials(t *tes.T) {
	var filename = testDirectory + "components/Credential.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var credential = not.Credential(source)
	credential.GetTag()
	credential.GetVersion()
	var formatted = credential.AsSource()
	ass.Equal(t, source, formatted)
}

func TestParsingContents(t *tes.T) {
	var filename = testDirectory + "components/Content.bali"
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
		permissions,
		optionalPrevious,
	)
	var formatted = content.AsSource()
	ass.Equal(t, source, formatted)
}

func TestParsingDocuments(t *tes.T) {
	var filename = testDirectory + "components/Document.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var document = not.Document(source)
	document.GetContent()
	var seal = document.RemoveNotarySeal()
	document.SetNotarySeal(seal)
	var formatted = document.AsSource()
	ass.Equal(t, source, formatted)
}

var identity not.IdentityLike

func TestParsingIdentities(t *tes.T) {
	var filename = testDirectory + "components/Identity.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	identity = not.Identity(source)
	var algorithm = identity.GetAlgorithm()
	var key = identity.GetKey()
	var attributes = identity.GetAttributes()
	var tag = identity.GetTag()
	var version = identity.GetVersion()
	var previous = identity.GetOptionalPrevious()
	identity = not.Identity(
		algorithm,
		key,
		attributes,
		tag,
		version,
		previous,
	)
	var formatted = identity.AsSource()
	ass.Equal(t, source, formatted)
}

// Create the security module and digital notary.
var ssm = not.SsmSha512()
var secret = "#ACH22TPZL7QSSFFH6GGG8D21N3S6Y5RQ"
var hsm = HsmEd25519TestClass().HsmEd25519(secret)

func TestSSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	var digest = ssm.DigestBytes(bytes)
	ass.Equal(t, "SHA512", ssm.GetDigestAlgorithm())
	ass.Equal(t, 64, len(digest))
	ass.Equal(t, []byte{
		0xb7, 0xb7, 0xa, 0xb, 0x14, 0xd7, 0xfa, 0x21,
		0x3c, 0x6c, 0xcd, 0x3c, 0xbf, 0xfc, 0x8b, 0xb8,
		0xf8, 0xe1, 0x1a, 0x85, 0xf1, 0x11, 0x3b, 0xe,
		0xb2, 0x6a, 0x0, 0x20, 0x8f, 0x2b, 0x9b, 0x3a,
		0x1d, 0xd4, 0xaa, 0xf3, 0x99, 0x62, 0x86, 0x1e,
		0x16, 0xab, 0x6, 0x22, 0x74, 0x34, 0x2a, 0x1c,
		0xe1, 0xf9, 0xdb, 0xa3, 0x65, 0x4f, 0x36, 0xfc,
		0x33, 0x82, 0x45, 0x58, 0x9f, 0x29, 0x6c, 0x28,
	}, digest)
}

func TestHSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	ass.Equal(t, "ED25519", hsm.GetSignatureAlgorithm())
	hsm.EraseKeys()
	var publicKey = hsm.GenerateKeys()
	var signature = hsm.SignBytes(bytes)
	ass.True(t, hsm.IsValid(publicKey, bytes, signature))
	var newPublicKey = hsm.RotateKeys()
	signature = hsm.SignBytes(newPublicKey)
	ass.True(t, hsm.IsValid(publicKey, newPublicKey, signature))
	hsm.EraseKeys()
}

var notary not.DigitalNotaryLike

func TestDigitalNotaryInitialization(t *tes.T) {
	// Should not be able to retrieve the certificate citation without any keys.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(t, "DigitalNotary: An error occurred while attempting to generate a security credential:\n    The digital notary has not yet been initialized.", message)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary = not.DigitalNotary(ssm, hsm)
	var context = doc.Moment()
	notary.GenerateCredential(context)
}

func TestDigitalNotaryGenerateKey(t *tes.T) {
	// Generate a new public-private key pair.
	var attributes = identity.GetAttributes()
	notary.ForgetKey()
	notary.GenerateKey(attributes)

	// Should not be able to do this twice.
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.Equal(
				t,
				"DigitalNotary: An error occurred while attempting to generate a new key pair:\n    The digital notary has already been initialized.",
				message,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	notary.GenerateKey(attributes)
}

func TestDigitalNotaryLifecycle(t *tes.T) {
	// Generate and validate a new public-private key pair.
	notary.ForgetKey()
	var attributes = identity.GetAttributes()
	var document = notary.GenerateKey(attributes)
	var certificateV1 = document
	ass.True(t, notary.SealMatches(document, certificateV1))
	var filename = "./test/agents/CertificateV1.bali"
	var source = document.AsSource()
	uti.WriteFile(filename, source)

	// Generate and validate a new citation to the certificate.
	var citation = notary.CiteDocument(document)
	ass.True(t, notary.CitationMatches(citation, document))
	filename = "./test/agents/Citation.bali"
	source = citation.AsSource()
	uti.WriteFile(filename, source)

	// Create and cite a new transaction document.
	var timestamp = doc.Moment().AsSource()
	var transaction = not.Content(
		`[
    $timestamp: ` + timestamp + `
    $consumer: "Derk Norton"
    $merchant: <https://www.starbucks.com/>
    $amount: 7.95($currency: $USD)
](
	$type: /bali/examples/Transaction/v1.2.3
	$tag: #BCASYZR1MC2J2QDPL03HG42W0M7P36TQ
	$version: v1
	$permissions: /bali/permissions/Public/v3
	$previous: none
)`,
	)
	filename = "./test/agents/Content.bali"
	source = transaction.AsSource()
	uti.WriteFile(filename, source)

	document = not.Document(transaction)

	// Notarize the transaction document to create a notarized document.
	notary.NotarizeDocument(document)
	ass.True(t, notary.SealMatches(document, certificateV1))
	filename = "./test/agents/Document.bali"
	source = document.AsSource()
	uti.WriteFile(filename, source)

	// Pickup where we left off with a new security module and digital notary.
	ssm = not.SsmSha512()
	hsm = HsmEd25519TestClass().HsmEd25519(secret)
	notary = not.DigitalNotaryWithCertificate(ssm, hsm, certificateV1)

	// Refresh and validate the public-private key pair.
	document = notary.RefreshKey()
	var certificateV2 = document
	ass.True(t, notary.SealMatches(document, certificateV1))
	filename = "./test/agents/CertificateV2.bali"
	source = document.AsSource()
	uti.WriteFile(filename, source)

	// Generate an authentication credential.
	var context = doc.ParseComponent(`[
    $website: <https://bali-nebula.com/>
    $sessionId: "ABC123456789"
]`).GetEntity()
	document = notary.GenerateCredential(context)
	ass.True(t, notary.SealMatches(document, certificateV2))
	document = notary.RefreshCredential(context, document)
	ass.True(t, notary.SealMatches(document, certificateV2))
	filename = "./test/agents/Credential.bali"
	source = document.AsSource()
	uti.WriteFile(filename, source)
}

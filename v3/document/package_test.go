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

package document_test

import (
	fmt "fmt"
	bal "github.com/bali-nebula/go-bali-documents/v3"
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestParsingCitations(t *tes.T) {
	var filename = "../test/Citation.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var citation = doc.CitationClass().CitationFromString(source)
	var formatted = citation.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingCertificates(t *tes.T) {
	var filename = "../test/Certificate.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var certificate = doc.CertificateClass().CertificateFromString(source)
	var formatted = certificate.AsString()
	ass.Equal(t, source, formatted)
}

func TestParsingDocuments(t *tes.T) {
	var filename = "../test/Document.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var documentClass = doc.DocumentClass()
	var document = documentClass.DocumentFromString(source)
	var formatted = document.AsString()
	ass.Equal(t, source, formatted)
	var attribute = documentClass.ExtractAttribute(
		"$consumer",
		bal.ParseSource(source),
	)
	ass.Equal(t, `"Derk Norton"`, attribute)
}

func TestParsingContracts(t *tes.T) {
	var filename = "../test/Contract.bali"
	fmt.Println(filename)
	var source = uti.ReadFile(filename)
	var contract = doc.ContractClass().ContractFromString(source)
	var formatted = contract.AsString()
	ass.Equal(t, source, formatted)
}

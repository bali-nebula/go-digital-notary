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

package document

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ContractClass() ContractClassLike {
	return contractClass()
}

// Constructor Methods

func (c *contractClass_) Contract(
	document DocumentLike,
	account string,
	protocol string,
	certificate CitationLike,
	signature string,
) ContractLike {
	if uti.IsUndefined(document) {
		panic("The \"document\" attribute is required by this class.")
	}
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(protocol) {
		panic("The \"protocol\" attribute is required by this class.")
	}
	if uti.IsUndefined(certificate) {
		panic("The \"certificate\" attribute is required by this class.")
	}
	if uti.IsUndefined(signature) {
		panic("The \"signature\" attribute is required by this class.")
	}
	var instance = &contract_{
		// Initialize the instance attributes.
		document_:    document,
		account_:     account,
		protocol_:    protocol,
		certificate_: certificate,
		signature_:   signature,
	}
	return instance
}

func (c *contractClass_) ContractFromString(
	source string,
) ContractLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid contract was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var component = doc.ParseSource(source).GetComponent()
	var collection = component.GetAny().(doc.CollectionLike)
	var attributes = collection.GetAny().(doc.AttributesLike)
	var associations = attributes.GetAssociations()

	var association = associations.GetValue(1)
	var element = association.GetPrimitive().GetAny().(doc.ElementLike)
	var symbol = element.GetAny().(string)
	if symbol != "$document" {
		panic("Missing the $document attribute.")
	}
	var document = documentClass().DocumentFromString(
		doc.FormatDocument(association.GetDocument()),
	)

	association = associations.GetValue(2)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$account" {
		panic("Missing the $account attribute.")
	}
	var account = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(3)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$protocol" {
		panic("Missing the $protocol attribute.")
	}
	var protocol = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(4)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$certificate" {
		panic("Missing the $certificate attribute.")
	}
	var certificate = citationClass().CitationFromString(
		doc.FormatDocument(association.GetDocument()),
	)

	association = associations.GetValue(5)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$signature" {
		panic("Missing the $signature attribute.")
	}
	var signature = doc.FormatDocument(association.GetDocument())

	return c.Contract(document, account, protocol, certificate, signature)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *contract_) GetClass() ContractClassLike {
	return contractClass()
}

func (v *contract_) AsString() string {
	var string_ = `[
`
	string_ += `    $document: ` + v.GetDocument().AsString()
	string_ += `    $account: ` + v.GetAccount()
	string_ += `    $protocol: ` + v.GetProtocol()
	string_ += `    $certificate: ` + v.GetCertificate().AsString()
	string_ += `    $signature: ` + v.GetSignature()
	string_ += `]($type: <bali:/types/documents/Contract@v1>)
`
	var contract = doc.ParseSource(string_)
	string_ = doc.FormatDocument(contract)
	return string_
}

// Attribute Methods

func (v *contract_) GetDocument() DocumentLike {
	return v.document_
}

func (v *contract_) GetAccount() string {
	return v.account_
}

func (v *contract_) GetProtocol() string {
	return v.protocol_
}

func (v *contract_) GetCertificate() CitationLike {
	return v.certificate_
}

func (v *contract_) GetSignature() string {
	return v.signature_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type contract_ struct {
	// Declare the instance attributes.
	document_    DocumentLike
	account_     string
	protocol_    string
	certificate_ CitationLike
	signature_   string
}

// Class Structure

type contractClass_ struct {
	// Declare the class constants.
}

// Class Reference

func contractClass() *contractClass_ {
	return contractClassReference_
}

var contractClassReference_ = &contractClass_{
	// Initialize the class constants.
}

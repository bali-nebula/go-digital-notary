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
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ContractClass() ContractClassLike {
	return contractClass()
}

// Constructor Methods

func (c *contractClass_) Contract(
	draft DraftLike,
	account fra.TagLike,
	certificate CitationLike,
) ContractLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(certificate) {
		panic("The \"certificate\" attribute is required by this class.")
	}
	var instance = &contract_{
		// Initialize the instance attributes.
		draft_:       draft,
		account_:     account,
		certificate_: certificate,
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
	var document = not.ParseSource(source)
	var account = c.extractAccount(document)
	var certificate = c.extractCertificate(document)
	var signature = c.extractSignature(document)
	var component = c.extractDraft(document)
	var contract = c.Contract(
		component,
		account,
		certificate,
	)
	contract.SetSignature(signature)
	return contract
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
	string_ += `    $draft: ` + v.GetDraft().AsString()
	string_ += `    $account: ` + v.GetAccount().AsString()
	string_ += `    $certificate: ` + v.GetCertificate().AsString()
	if uti.IsDefined(v.signature_) {
		string_ += `    $signature: ` + v.GetSignature().AsString()
	}
	string_ += `]($type: <bali:/types/documents/Contract:v3>)
`
	var contract = not.ParseSource(string_)
	string_ = not.FormatDocument(contract)
	return string_
}

// Attribute Methods

func (v *contract_) GetDraft() DraftLike {
	return v.draft_
}

func (v *contract_) GetAccount() fra.TagLike {
	return v.account_
}

func (v *contract_) GetCertificate() CitationLike {
	return v.certificate_
}

func (v *contract_) GetSignature() SignatureLike {
	return v.signature_
}

func (v *contract_) SetSignature(
	signature SignatureLike,
) {
	v.signature_ = signature
}

// PROTECTED INTERFACE

// Private Methods

func (c *contractClass_) extractAccount(
	document not.DocumentLike,
) fra.TagLike {
	var attribute = c.extractAttribute("$account", document)
	var account = fra.TagFromString(attribute)
	return account
}

func (c *contractClass_) extractAttribute(
	name string,
	document not.DocumentLike,
) string {
	var attribute string
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(not.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			attribute = not.FormatDocument(association.GetDocument())
			attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
			break
		}
	}
	return attribute
}

func (c *contractClass_) extractCertificate(
	document not.DocumentLike,
) CitationLike {
	var attribute = c.extractAttribute("$certificate", document)
	var certificate = CitationClass().CitationFromString(attribute)
	return certificate
}

func (c *contractClass_) extractDraft(
	document not.DocumentLike,
) DraftLike {
	var attribute = c.extractAttribute("$draft", document)
	var draft = DraftClass().DraftFromString(attribute)
	return draft
}

func (c *contractClass_) extractSignature(
	document not.DocumentLike,
) SignatureLike {
	var attribute = c.extractAttribute("$signature", document)
	var signature = SignatureClass().SignatureFromString(attribute)
	return signature
}

// Instance Structure

type contract_ struct {
	// Declare the instance attributes.
	draft_       DraftLike
	account_     fra.TagLike
	certificate_ CitationLike
	signature_   SignatureLike
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

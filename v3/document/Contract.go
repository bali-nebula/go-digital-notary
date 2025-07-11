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
	bal "github.com/bali-nebula/go-document-notation/v3"
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
	var document = bal.ParseSource(source)
	var account = fra.TagFromString(
		DraftClass().ExtractAttribute("$account", document),
	)
	var certificate = DraftClass().ExtractCertificate(document)
	var signature = DraftClass().ExtractSignature(document)
	var component = DraftClass().ExtractDraft(document)
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
	var contract = bal.ParseSource(string_)
	string_ = bal.FormatDocument(contract)
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

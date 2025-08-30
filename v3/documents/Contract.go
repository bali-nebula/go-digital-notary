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

package documents

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
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
	certificate fra.ResourceLike,
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

	var component = doc.ParseSource(`[
    $draft: ` + draft.AsString() + `
    $account: ` + account.AsString() + `
    $certificate: ` + certificate.AsString() + `
]($type: <bali:/nebula/types/Contract:v3>)`,
	)

	var instance = &contract_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *contractClass_) ContractFromString(
	source string,
) ContractLike {
	var component = doc.ParseSource(source)
	var instance = &contract_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *contract_) GetClass() ContractClassLike {
	return contractClass()
}

func (v *contract_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *contract_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *contract_) GetDraft() DraftLike {
	var object = v.GetObject(fra.Symbol("draft"))
	return DraftClass().DraftFromString(doc.FormatComponent(object))
}

func (v *contract_) GetAccount() fra.TagLike {
	var object = v.GetObject(fra.Symbol("account"))
	return fra.TagFromString(doc.FormatComponent(object))
}

func (v *contract_) GetCertificate() fra.ResourceLike {
	var object = v.GetObject(fra.Symbol("certificate"))
	return fra.ResourceFromString(doc.FormatComponent(object))
}

func (v *contract_) GetSignature() SignatureLike {
	var signature SignatureLike
	var object = v.GetObject(fra.Symbol("signature"))
	if uti.IsDefined(object) {
		signature = SignatureClass().SignatureFromString(doc.FormatComponent(object))
	}
	return signature
}

func (v *contract_) SetSignature(
	signature SignatureLike,
) {
	var component = doc.ParseSource(signature.AsString())
	v.SetObject(component, fra.Symbol("signature"))
}

func (v *contract_) RemoveSignature() {
	v.RemoveObject(fra.Symbol("signature"))
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type contract_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
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

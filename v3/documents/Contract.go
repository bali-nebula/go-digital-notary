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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ContractClass() ContractClassLike {
	return contractClass()
}

// Constructor Methods

func (c *contractClass_) Contract(
	draft Parameterized,
	account doc.TagLike,
	notary doc.ResourceLike,
) ContractLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(notary) {
		panic("The \"notary\" attribute is required by this class.")
	}

	var component = doc.ParseSource(`[
    $content: ` + draft.AsString() + `
    $account: ` + account.AsString() + `
    $notary: ` + notary.AsString() + `
]($type: <bali:/types/notary/Contract:v3>)`,
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

// Attribute Methods

func (v *contract_) GetDraft() Parameterized {
	var object = v.GetObject(doc.Symbol("content"))
	return DraftClass().DraftFromString(doc.FormatComponent(object))
}

// Notarized Methods

func (v *contract_) GetContent() Parameterized {
	var object = v.GetObject(doc.Symbol("content"))
	return DraftClass().DraftFromString(doc.FormatComponent(object))
}

func (v *contract_) GetAccount() doc.TagLike {
	var object = v.GetObject(doc.Symbol("account"))
	return doc.Tag(doc.FormatComponent(object))
}

func (v *contract_) GetNotary() doc.ResourceLike {
	var object = v.GetObject(doc.Symbol("notary"))
	return doc.Resource(doc.FormatComponent(object))
}

func (v *contract_) SetSeal(
	seal SealLike,
) {
	var component = doc.ParseSource(seal.AsString())
	v.SetObject(component, doc.Symbol("seal"))
}

func (v *contract_) RemoveSeal() SealLike {
	var seal SealLike
	var symbol = doc.Symbol("seal")
	var object = v.GetObject(symbol)
	if uti.IsDefined(object) {
		v.RemoveObject(symbol)
		seal = SealClass().SealFromString(doc.FormatComponent(object))
	}
	return seal
}

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

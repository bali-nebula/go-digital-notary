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

func CredentialClass() CredentialClassLike {
	return credentialClass()
}

// Constructor Methods

func (c *credentialClass_) Credential(
	account fra.TagLike,
	notary fra.ResourceLike,
) CredentialLike {
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(notary) {
		panic("The \"notary\" attribute is required by this class.")
	}

	var timestamp = fra.Now()
	var component = doc.ParseSource(`[
    $content: ` + timestamp.AsString() + `
    $account: ` + account.AsString() + `
    $notary: ` + notary.AsString() + `
]($type: <bali:/types/notary/Credential:v3>)`,
	)

	var instance = &credential_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}

	return instance
}

func (c *credentialClass_) CredentialFromString(
	source string,
) CredentialLike {
	var component = doc.ParseSource(source)
	var instance = &credential_{
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

func (v *credential_) GetClass() CredentialClassLike {
	return credentialClass()
}

func (v *credential_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *credential_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

// Notarized Methods

func (v *credential_) GetContent() doc.ComponentLike {
	var object = v.GetObject(fra.Symbol("content"))
	return object.GetComponent()
}

func (v *credential_) GetAccount() fra.TagLike {
	var object = v.GetObject(fra.Symbol("account"))
	return fra.TagFromString(doc.FormatComponent(object))
}

func (v *credential_) GetNotary() fra.ResourceLike {
	var object = v.GetObject(fra.Symbol("notary"))
	return fra.ResourceFromString(doc.FormatComponent(object))
}

func (v *credential_) GetSeal() SealLike {
	var seal SealLike
	var object = v.GetObject(fra.Symbol("seal"))
	if uti.IsDefined(object) {
		seal = SealClass().SealFromString(doc.FormatComponent(object))
	}
	return seal
}

func (v *credential_) SetSeal(
	seal SealLike,
) {
	var component = doc.ParseSource(seal.AsString())
	v.SetObject(component, fra.Symbol("seal"))
}

func (v *credential_) RemoveSeal() SealLike {
	var seal SealLike
	var symbol = fra.Symbol("seal")
	var object = v.GetObject(symbol)
	if uti.IsDefined(object) {
		v.RemoveObject(symbol)
		seal = SealClass().SealFromString(doc.FormatComponent(object))
	}
	return seal
}

// Private Methods

// Instance Structure

type credential_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type credentialClass_ struct {
	// Declare the class constants.
}

// Class Reference

func credentialClass() *credentialClass_ {
	return credentialClassReference_
}

var credentialClassReference_ = &credentialClass_{
	// Initialize the class constants.
}

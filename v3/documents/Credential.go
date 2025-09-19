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
	signatory fra.ResourceLike,
) CredentialLike {
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(signatory) {
		panic("The \"signatory\" attribute is required by this class.")
	}

	var timestamp = fra.Now()
	var component = doc.ParseSource(`[
    $content: ` + timestamp.AsString() + `
    $account: ` + account.AsString() + `
    $signatory: ` + signatory.AsString() + `
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

func (v *credential_) GetSignatory() fra.ResourceLike {
	var object = v.GetObject(fra.Symbol("signatory"))
	return fra.ResourceFromString(doc.FormatComponent(object))
}

func (v *credential_) GetSignature() SignatureLike {
	var signature SignatureLike
	var object = v.GetObject(fra.Symbol("signature"))
	if uti.IsDefined(object) {
		signature = SignatureClass().SignatureFromString(doc.FormatComponent(object))
	}
	return signature
}

func (v *credential_) SetSignature(
	signature SignatureLike,
) {
	var component = doc.ParseSource(signature.AsString())
	v.SetObject(component, fra.Symbol("signature"))
}

func (v *credential_) RemoveSignature() {
	v.RemoveObject(fra.Symbol("signature"))
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

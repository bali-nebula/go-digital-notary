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

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	key KeyLike,
	account fra.TagLike,
	notary fra.ResourceLike,
) CertificateLike {
	if uti.IsUndefined(key) {
		panic("The \"key\" attribute is required by this class.")
	}
	if uti.IsUndefined(account) {
		panic("The \"account\" attribute is required by this class.")
	}
	if uti.IsUndefined(notary) {
		panic("The \"notary\" attribute is required by this class.")
	}

	var component = doc.ParseSource(`[
    $content: ` + key.AsString() + `
    $account: ` + account.AsString() + `
    $notary: ` + notary.AsString() + `
]($type: <bali:/types/notary/Certificate:v3>)`,
	)

	var instance = &certificate_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}

	return instance
}

func (c *certificateClass_) CertificateFromString(
	source string,
) CertificateLike {
	var component = doc.ParseSource(source)
	var instance = &certificate_{
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

func (v *certificate_) GetClass() CertificateClassLike {
	return certificateClass()
}

func (v *certificate_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *certificate_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

// Attribute Methods

func (v *certificate_) GetKey() KeyLike {
	var object = v.GetObject(fra.Symbol("content"))
	return KeyClass().KeyFromString(doc.FormatComponent(object))
}

// Notarized Methods

func (v *certificate_) GetContent() doc.ComponentLike {
	var object = v.GetObject(fra.Symbol("content"))
	return object.GetComponent()
}

func (v *certificate_) GetAccount() fra.TagLike {
	var object = v.GetObject(fra.Symbol("account"))
	return fra.TagFromString(doc.FormatComponent(object))
}

func (v *certificate_) GetNotary() fra.ResourceLike {
	var object = v.GetObject(fra.Symbol("notary"))
	return fra.ResourceFromString(doc.FormatComponent(object))
}

func (v *certificate_) GetSeal() SealLike {
	var seal SealLike
	var object = v.GetObject(fra.Symbol("seal"))
	if uti.IsDefined(object) {
		seal = SealClass().SealFromString(doc.FormatComponent(object))
	}
	return seal
}

func (v *certificate_) SetSeal(
	seal SealLike,
) {
	var component = doc.ParseSource(seal.AsString())
	v.SetObject(component, fra.Symbol("seal"))
}

func (v *certificate_) RemoveSeal() {
	v.RemoveObject(fra.Symbol("seal"))
}

// Private Methods

// Instance Structure

type certificate_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type certificateClass_ struct {
	// Declare the class constants.
}

// Class Reference

func certificateClass() *certificateClass_ {
	return certificateClassReference_
}

var certificateClassReference_ = &certificateClass_{
	// Initialize the class constants.
}

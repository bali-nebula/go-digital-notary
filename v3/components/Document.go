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

package components

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	content Parameterized,
) DocumentLike {
	if uti.IsUndefined(content) {
		panic("The \"content\" attribute is required by this class.")
	}

	var source = `[
    $content: ` + content.AsSource() + `
]($type: /bali/types/notary/Document/v3)`
	return c.DocumentFromSource(source)
}

func (c *documentClass_) DocumentFromSource(
	source string,
) DocumentLike {
	var component = doc.ParseComponent(source)
	var instance = &document_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Composite: component,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

func (v *document_) AsIntrinsic() doc.Composite {
	return v.Composite
}

func (v *document_) AsSource() string {
	return doc.FormatComponent(v.Composite) + "\n"
}

// Attribute Methods

func (v *document_) GetContent() Parameterized {
	var content = v.GetSubcomponent(doc.Symbol("$content"))
	return ContentClass().ContentFromSource(doc.FormatComponent(content))
}

func (v *document_) SetOptionalNotary(
	notary NotaryLike,
) {
	v.SetSubcomponent(
		notary,
		doc.Symbol("$notary"),
	)
}

func (v *document_) GetOptionalNotary() NotaryLike {
	var notary NotaryLike
	var component = v.GetSubcomponent(doc.Symbol("$notary"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		notary = NotaryClass().NotaryFromSource(source)
	}
	return notary
}

func (v *document_) SetNotarySeal(
	seal SealLike,
) {
	v.SetSubcomponent(
		seal,
		doc.Symbol("$notary"),
		doc.Symbol("$seal"),
	)
}

func (v *document_) RemoveNotarySeal() SealLike {
	var seal SealLike
	var component = v.RemoveSubcomponent(
		doc.Symbol("$notary"),
		doc.Symbol("$seal"),
	)
	if uti.IsDefined(component) {
		seal = SealClass().SealFromSource(doc.FormatComponent(component))
	}
	return seal
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type document_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Composite
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentClass() *documentClass_ {
	return documentClassReference_
}

var documentClassReference_ = &documentClass_{
	// Initialize the class constants.
}

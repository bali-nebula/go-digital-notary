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
		ComponentLike: component,
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

func (v *document_) AsIntrinsic() doc.ComponentLike {
	return v.ComponentLike
}

func (v *document_) AsSource() string {
	return doc.FormatComponent(v.ComponentLike) + "\n"
}

// Attribute Methods

func (v *document_) GetContent() Parameterized {
	var object = v.GetObject(doc.Symbol("$content"))
	return ContentClass().ContentFromSource(doc.FormatComponent(object))
}

func (v *document_) GetTimestamp() doc.MomentLike {
	var object = v.GetObject(doc.Symbol("$timestamp"))
	var timestamp doc.MomentLike
	if uti.IsDefined(object) {
		timestamp = doc.Moment(doc.FormatComponent(object))
	}
	return timestamp
}

func (v *document_) GetAccount() doc.TagLike {
	var object = v.GetObject(doc.Symbol("$account"))
	var account doc.TagLike
	if uti.IsDefined(object) {
		account = doc.Tag(doc.FormatComponent(object))
	}
	return account
}

func (v *document_) SetNotary(
	account doc.TagLike,
	notary CitationLike,
) {
	var component = doc.ParseComponent(account.AsSource())
	v.SetObject(component, doc.Symbol("$account"))
	component = doc.ParseComponent(doc.Moment().AsSource())
	v.SetObject(component, doc.Symbol("$timestamp"))
	component = doc.ParseComponent("none")
	if uti.IsDefined(notary) {
		component = doc.ParseComponent(notary.AsSource())
	}
	v.SetObject(component, doc.Symbol("$notary"))
}

func (v *document_) GetNotary() CitationLike {
	var object = v.GetObject(doc.Symbol("$notary"))
	var notary CitationLike
	if uti.IsDefined(object) && doc.FormatComponent(object) != "none" {
		notary = CitationClass().CitationFromSource(doc.FormatComponent(object))
	}
	return notary
}

func (v *document_) SetSeal(
	seal SealLike,
) {
	var component = doc.ParseComponent(seal.AsSource())
	v.SetObject(component, doc.Symbol("$seal"))
}

func (v *document_) HasSeal() bool {
	var symbol = doc.Symbol("$seal")
	var object = v.GetObject(symbol)
	return uti.IsDefined(object)
}

func (v *document_) RemoveSeal() SealLike {
	var seal SealLike
	var symbol = doc.Symbol("$seal")
	var object = v.GetObject(symbol)
	if uti.IsDefined(object) {
		v.RemoveObject(symbol)
		seal = SealClass().SealFromSource(doc.FormatComponent(object))
	}
	return seal
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type document_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.ComponentLike
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

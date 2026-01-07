/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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

func NotaryClass() NotaryClassLike {
	return notaryClass()
}

// Constructor Methods

func (c *notaryClass_) Notary(
	owner doc.TagLike,
	optionalCitation CitationLike,
) NotaryLike {
	if uti.IsUndefined(owner) {
		panic("The \"owner\" attribute is required by this class.")
	}

	var timestamp = doc.Moment() // The current date and time.
	var citation = "none"        // In case this is a self-signed citation.
	if uti.IsDefined(optionalCitation) {
		citation = optionalCitation.AsSource()
	}
	var source = `[
    $owner: ` + owner.AsSource() + `
    $timestamp: ` + timestamp.AsSource() + `
	$citation: ` + citation + `
]($type: /bali/types/notary/Notary/v3)`
	return c.NotaryFromSource(source)
}

func (c *notaryClass_) NotaryFromSource(
	source string,
) NotaryLike {
	var component = doc.ParseComponent(source)
	var instance = &notary_{
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

func (v *notary_) GetClass() NotaryClassLike {
	return notaryClass()
}

func (v *notary_) AsIntrinsic() doc.Composite {
	return v.Composite
}

func (v *notary_) AsSource() string {
	return doc.FormatComponent(v.Composite) + "\n"
}

// Attribute Methods

func (v *notary_) GetOwner() doc.TagLike {
	var component = v.GetSubcomponent(doc.Symbol("$owner"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *notary_) GetTimestamp() doc.MomentLike {
	var component = v.GetSubcomponent(doc.Symbol("$timestamp"))
	return doc.Moment(doc.FormatComponent(component))
}

func (v *notary_) GetOptionalCitation() CitationLike {
	var citation CitationLike
	var component = v.GetParameter(doc.Symbol("$citation"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		if source != "none" {
			citation = CitationClass().CitationFromSource(source)
		}
	}
	return citation
}

func (v *notary_) GetOptionalSeal() SealLike {
	var seal SealLike
	var component = v.GetSubcomponent(doc.Symbol("$seal"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		seal = SealClass().SealFromSource(source)
	}
	return seal
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type notary_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Composite
}

// Class Structure

type notaryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func notaryClass() *notaryClass_ {
	return notaryClassReference_
}

var notaryClassReference_ = &notaryClass_{
	// Initialize the class constants.
}

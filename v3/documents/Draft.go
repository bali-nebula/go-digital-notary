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

func DraftClass() DraftClassLike {
	return draftClass()
}

// Constructor Methods

func (c *draftClass_) Draft(
	entity any,
	type_ fra.ResourceLike,
	tag fra.TagLike,
	version fra.VersionLike,
	permissions fra.ResourceLike,
	optionalPrevious fra.ResourceLike,
) DraftLike {
	if uti.IsUndefined(entity) {
		panic("The \"entity\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(permissions) {
		panic("The \"permissions\" attribute is required by this class.")
	}
	var previous = "none"
	if uti.IsDefined(optionalPrevious) {
		previous = optionalPrevious.AsString()
	}

	var component = doc.ParseSource(
		doc.FormatComponent(entity) + `(
    $type: ` + type_.AsString() + `
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $permissions: ` + permissions.AsString() + `
    $previous: ` + previous + `
)`,
	)

	var instance = &draft_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *draftClass_) DraftFromString(
	source string,
) DraftLike {
	var component = doc.ParseSource(source)
	var instance = &draft_{
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

func (v *draft_) GetClass() DraftClassLike {
	return draftClass()
}

func (v *draft_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *draft_) GetComponent() any {
	return v.Declarative.(doc.ComponentLike)
}

// Attribute Methods

// Parameterized Methods

func (v *draft_) GetType() fra.ResourceLike {
	var component = v.GetParameter(fra.Symbol("type"))
	return fra.ResourceFromString(doc.FormatComponent(component))
}

func (v *draft_) GetTag() fra.TagLike {
	var component = v.GetParameter(fra.Symbol("tag"))
	return fra.TagFromString(doc.FormatComponent(component))
}

func (v *draft_) GetVersion() fra.VersionLike {
	var component = v.GetParameter(fra.Symbol("version"))
	return fra.VersionFromString(doc.FormatComponent(component))
}

func (v *draft_) GetPermissions() fra.ResourceLike {
	var component = v.GetParameter(fra.Symbol("permissions"))
	return fra.ResourceFromString(doc.FormatComponent(component))
}

func (v *draft_) GetPrevious() any {
	var component = v.GetParameter(fra.Symbol("previous"))
	var source = doc.FormatComponent(component)
	switch source {
	case "none":
		return fra.PatternFromString(source)
	default:
		return fra.ResourceFromString(source)
	}
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type draft_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type draftClass_ struct {
	// Declare the class constants.
}

// Class Reference

func draftClass() *draftClass_ {
	return draftClassReference_
}

var draftClassReference_ = &draftClass_{
	// Initialize the class constants.
}

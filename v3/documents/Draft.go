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

func DraftClass() DraftClassLike {
	return draftClass()
}

// Constructor Methods

func (c *draftClass_) Draft(
	entity any,
	type_ doc.ResourceLike,
	tag doc.TagLike,
	version doc.VersionLike,
	permissions doc.ResourceLike,
	optionalPrevious doc.ResourceLike,
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

	var component = doc.ParseSource(doc.FormatComponent(entity) + `(
    $type: ` + type_.AsString() + `
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $permissions: ` + permissions.AsString() + `
    $previous: ` + previous + `
)`)

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

func (v *draft_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *draft_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

// Attribute Methods

// Parameterized Methods

func (v *draft_) GetEntity() any {
	return v.Declarative.(doc.ComponentLike).GetEntity()
}

func (v *draft_) GetType() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *draft_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *draft_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *draft_) GetPermissions() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *draft_) GetOptionalPrevious() doc.ResourceLike {
	var previous doc.ResourceLike
	var component = v.GetParameter(doc.Symbol("$previous"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		if source != "none" {
			previous = doc.Resource(source)
		}
	}
	return previous
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

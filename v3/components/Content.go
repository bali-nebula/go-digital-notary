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

func ContentClass() ContentClassLike {
	return contentClass()
}

// Constructor Methods

func (c *contentClass_) Content(
	entity any,
	type_ doc.NameLike,
	tag doc.TagLike,
	version doc.VersionLike,
	permissions doc.NameLike,
	optionalPrevious doc.ResourceLike,
) ContentLike {
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
		previous = optionalPrevious.AsSource()
	}

	var source = doc.FormatComponent(entity) + `(
    $type: ` + type_.AsSource() + `
    $tag: ` + tag.AsSource() + `
    $version: ` + version.AsSource() + `
    $permissions: ` + permissions.AsSource() + `
    $previous: ` + previous + `
)`
	return c.ContentFromSource(source)
}

func (c *contentClass_) ContentFromSource(
	source string,
) ContentLike {
	var component = doc.ParseComponent(source)
	var instance = &content_{
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

func (v *content_) GetClass() ContentClassLike {
	return contentClass()
}

func (v *content_) AsIntrinsic() doc.ComponentLike {
	return v.ComponentLike
}

// Attribute Methods

// Parameterized Methods

func (v *content_) AsSource() string {
	return doc.FormatComponent(v.ComponentLike) + "\n"
}

func (v *content_) GetType() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *content_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *content_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *content_) GetPermissions() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *content_) GetOptionalPrevious() doc.ResourceLike {
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

type content_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.ComponentLike
}

// Class Structure

type contentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func contentClass() *contentClass_ {
	return contentClassReference_
}

var contentClassReference_ = &contentClass_{
	// Initialize the class constants.
}

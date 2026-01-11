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

func IdentityClass() IdentityClassLike {
	return identityClass()
}

// Constructor Methods

func (c *identityClass_) Identity(
	algorithm doc.QuoteLike,
	key doc.BinaryLike,
	attributes doc.Composite,
	tag doc.TagLike,
	version doc.VersionLike,
	optionalPrevious doc.ResourceLike,
) IdentityLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(key) {
		panic("The \"key\" attribute is required by this class.")
	}
	if uti.IsUndefined(attributes) {
		panic("The \"attributes\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var previous = "none"
	if uti.IsDefined(optionalPrevious) {
		previous = optionalPrevious.AsSource()
	}
	var source = `[
    $algorithm: ` + algorithm.AsSource() + `
    $key: ` + key.AsSource() + `
    $attributes: ` + doc.FormatComponent(attributes) + `
](
    $type: /bali/types/notary/Identity/v3
    $tag: ` + tag.AsSource() + `
    $version: ` + version.AsSource() + `
    $permissions: /bali/permissions/Public/v3
    $previous: ` + previous + `
)`
	return c.IdentityFromSource(source)
}

func (c *identityClass_) IdentityFromSource(
	source string,
) IdentityLike {
	var component = doc.ParseComponent(source)
	var instance = &identity_{
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

func (v *identity_) GetClass() IdentityClassLike {
	return identityClass()
}

func (v *identity_) AsIntrinsic() doc.Composite {
	return v.Composite
}

func (v *identity_) AsSource() string {
	return doc.FormatComponent(v.Composite) + "\n"
}

// Attribute Methods

func (v *identity_) GetAlgorithm() doc.QuoteLike {
	var component = v.GetSubcomponent(doc.Symbol("$algorithm"))
	return doc.Quote(doc.FormatComponent(component))
}

func (v *identity_) GetKey() doc.BinaryLike {
	var component = v.GetSubcomponent(doc.Symbol("$key"))
	return doc.Binary(doc.FormatComponent(component))
}

func (v *identity_) GetAttributes() doc.Composite {
	var component = v.GetSubcomponent(doc.Symbol("$attributes"))
	return component
}

// Parameterized Methods

func (v *identity_) GetType() doc.NameLike {
	var component = v.GetConstraint(doc.Symbol("$type"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *identity_) GetTag() doc.TagLike {
	var component = v.GetConstraint(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *identity_) GetVersion() doc.VersionLike {
	var component = v.GetConstraint(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *identity_) GetPermissions() doc.NameLike {
	var component = v.GetConstraint(doc.Symbol("$permissions"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *identity_) GetOptionalPrevious() doc.ResourceLike {
	var previous doc.ResourceLike
	var component = v.GetConstraint(doc.Symbol("$previous"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		if source != "none" {
			previous = doc.Resource(source)
		}
	}
	return previous
}

// Private Methods

// Instance Structure

type identity_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Composite
}

// Class Structure

type identityClass_ struct {
	// Declare the class constants.
}

// Class Reference

func identityClass() *identityClass_ {
	return identityClassReference_
}

var identityClassReference_ = &identityClass_{
	// Initialize the class constants.
}

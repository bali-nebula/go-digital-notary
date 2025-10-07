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

func CredentialClass() CredentialClassLike {
	return credentialClass()
}

// Constructor Methods

func (c *credentialClass_) Credential(
	context any,
	tag doc.TagLike,
	version doc.VersionLike,
	optionalPrevious doc.ResourceLike,
) CredentialLike {
	if uti.IsUndefined(context) {
		panic("The \"context\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var previous = "none"
	if uti.IsDefined(optionalPrevious) {
		previous = optionalPrevious.AsString()
	}
	var component = doc.Component(context, nil)
	var source = doc.FormatComponent(component) + `(
    $type: <bali:/types/notary/Credential:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $previous: ` + previous + `
    $permissions: <bali:/permissions/Public:v3>
)`
	return c.CredentialFromString(source)
}

func (c *credentialClass_) CredentialFromString(
	source string,
) CredentialLike {
	var component = doc.ParseSource(source)
	var instance = &credential_{
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

func (v *credential_) GetClass() CredentialClassLike {
	return credentialClass()
}

func (v *credential_) AsIntrinsic() doc.ComponentLike {
	return v.ComponentLike
}

func (v *credential_) AsString() string {
	return doc.FormatDocument(v.ComponentLike)
}

// Attribute Methods

func (v *credential_) GetContext() any {
	var object = v.GetEntity()
	return doc.Moment(doc.FormatComponent(object))
}

// Parameterized Methods

func (v *credential_) GetType() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *credential_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *credential_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *credential_) GetOptionalPrevious() doc.ResourceLike {
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

func (v *credential_) GetPermissions() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Resource(doc.FormatComponent(component))
}

// Private Methods

// Instance Structure

type credential_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.ComponentLike
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

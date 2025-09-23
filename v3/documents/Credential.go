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
	tag doc.TagLike,
	version doc.VersionLike,
) CredentialLike {
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var timestamp = doc.Moment() // The current moment in time as a salt.
	var previous = "none"
	var current = version.AsIntrinsic()[0]
	if current > 1 {
		previous = "<nebula:/" + tag.AsString()[1:] +
			":" + doc.Version([]uint{current - 1}).AsString() + ">"
	}
	var component = doc.ParseSource(`[
    $timestamp: ` + timestamp.AsString() + `
](
    $type: <bali:/types/notary/Credential:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $permissions: <bali:/permissions/Public:v3>
    $previous: ` + previous + `
)`,
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

// Attribute Methods

func (v *credential_) GetTimestamp() doc.MomentLike {
	var object = v.GetObject(doc.Symbol("$timestamp"))
	return doc.Moment(doc.FormatComponent(object))
}

// Parameterized Methods

func (v *credential_) GetEntity() any {
	return v.Declarative.(doc.ComponentLike).GetEntity()
}

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

func (v *credential_) GetPermissions() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Resource(doc.FormatComponent(component))
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

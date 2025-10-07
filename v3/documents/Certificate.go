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

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	tag doc.TagLike,
	version doc.VersionLike,
	algorithm doc.QuoteLike,
	key doc.BinaryLike,
	optionalPrevious doc.ResourceLike,
) CertificateLike {
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(key) {
		panic("The \"key\" attribute is required by this class.")
	}

	var previous = "none"
	if uti.IsDefined(optionalPrevious) {
		previous = optionalPrevious.AsString()
	}
	var source = `[
    $algorithm: ` + algorithm.AsString() + `
    $key: ` + key.AsString() + `
](
    $type: <bali:/types/notary/Certificate:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $previous: ` + previous + `
    $permissions: <bali:/permissions/Public:v3>
)`
	return c.CertificateFromString(source)
}

func (c *certificateClass_) CertificateFromString(
	source string,
) CertificateLike {
	var component = doc.ParseSource(source)
	var instance = &certificate_{
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

func (v *certificate_) GetClass() CertificateClassLike {
	return certificateClass()
}

func (v *certificate_) AsIntrinsic() doc.ComponentLike {
	return v.ComponentLike
}

func (v *certificate_) AsString() string {
	return doc.FormatDocument(v.ComponentLike)
}

// Attribute Methods

func (v *certificate_) GetAlgorithm() doc.QuoteLike {
	var object = v.GetObject(doc.Symbol("$algorithm"))
	return doc.Quote(doc.FormatComponent(object))
}

func (v *certificate_) GetKey() doc.BinaryLike {
	var object = v.GetObject(doc.Symbol("$key"))
	return doc.Binary(doc.FormatComponent(object))
}

// Parameterized Methods

func (v *certificate_) GetType() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *certificate_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *certificate_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *certificate_) GetOptionalPrevious() doc.ResourceLike {
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

func (v *certificate_) GetPermissions() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Resource(doc.FormatComponent(component))
}

// Private Methods

// Instance Structure

type certificate_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.ComponentLike
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

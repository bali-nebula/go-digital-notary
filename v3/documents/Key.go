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

func KeyClass() KeyClassLike {
	return keyClass()
}

// Constructor Methods

func (c *keyClass_) Key(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
	tag fra.TagLike,
	version fra.VersionLike,
) KeyLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(base64) {
		panic("The \"base64\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var created = fra.Now()
	var previous = "none"
	var current = version.AsIntrinsic()[0]
	if current > 1 {
		previous = "<bali:/nebula/certificates/" + tag.AsString()[1:] +
			":" + fra.Version([]uti.Ordinal{current - 1}).AsString() + ">"
	}
	var component = doc.ParseSource(`[
    $created: ` + created.AsString() + `
    $algorithm: ` + algorithm.AsString() + `
    $base64: ` + base64.AsString() + `
](
    $type: <bali:/nebula/types/Key:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
	$permissions: <bali:/nebula/permissions/public:v3>
    $previous: ` + previous + `
)`,
	)

	var instance = &key_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *keyClass_) KeyFromString(
	source string,
) KeyLike {
	var component = doc.ParseSource(source)
	var instance = &key_{
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

func (v *key_) GetClass() KeyClassLike {
	return keyClass()
}

func (v *key_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *key_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *key_) GetCreated() fra.MomentLike {
	var object = v.GetObject(fra.Symbol("created"))
	return fra.MomentFromString(doc.FormatComponent(object))
}

func (v *key_) GetAlgorithm() fra.QuoteLike {
	var object = v.GetObject(fra.Symbol("algorithm"))
	return fra.QuoteFromString(doc.FormatComponent(object))
}

func (v *key_) GetBase64() fra.BinaryLike {
	var object = v.GetObject(fra.Symbol("base64"))
	return fra.BinaryFromString(doc.FormatComponent(object))
}

// Attribute Methods

// Parameterized Methods

func (v *key_) GetEntity() any {
	return v.Declarative.(doc.ComponentLike).GetEntity()
}

func (v *key_) GetType() fra.ResourceLike {
	var component = v.GetParameter(fra.Symbol("type"))
	return fra.ResourceFromString(doc.FormatComponent(component))
}

func (v *key_) GetTag() fra.TagLike {
	var component = v.GetParameter(fra.Symbol("tag"))
	return fra.TagFromString(doc.FormatComponent(component))
}

func (v *key_) GetVersion() fra.VersionLike {
	var component = v.GetParameter(fra.Symbol("version"))
	return fra.VersionFromString(doc.FormatComponent(component))
}

func (v *key_) GetPermissions() fra.ResourceLike {
	var component = v.GetParameter(fra.Symbol("permissions"))
	return fra.ResourceFromString(doc.FormatComponent(component))
}

func (v *key_) GetOptionalPrevious() fra.ResourceLike {
	var previous fra.ResourceLike
	var component = v.GetParameter(fra.Symbol("previous"))
	var source = doc.FormatComponent(component)
	if source != "none" {
		previous = fra.ResourceFromString(source)
	}
	return previous
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type key_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type keyClass_ struct {
	// Declare the class constants.
}

// Class Reference

func keyClass() *keyClass_ {
	return keyClassReference_
}

var keyClassReference_ = &keyClass_{
	// Initialize the class constants.
}

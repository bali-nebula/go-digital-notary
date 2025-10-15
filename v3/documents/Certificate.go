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
		previous = optionalPrevious.AsSource()
	}
	var source = `[
    $algorithm: ` + algorithm.AsSource() + `
    $key: ` + key.AsSource() + `
](
	$type: /bali/types/notary/Certificate/v3
    $tag: ` + tag.AsSource() + `
    $version: ` + version.AsSource() + `
	$permissions: /bali/permissions/Public/v3
    $previous: ` + previous + `
)`
	return c.CertificateFromSource(source)
}

func (c *certificateClass_) CertificateFromSource(
	source string,
) CertificateLike {
	var component = doc.ParseComponent(source)
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

func (v *certificate_) AsSource() string {
	return doc.FormatComponent(v.ComponentLike) + "\n"
}

// Attribute Methods

func (v *certificate_) GetAlgorithm() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$algorithm"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *certificate_) GetKey() doc.BinaryLike {
	var composite = v.GetSubcomponent(doc.Symbol("$key"))
	return doc.Binary(doc.FormatComponent(composite))
}

// Parameterized Methods

func (v *certificate_) GetType() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *certificate_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *certificate_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *certificate_) GetPermissions() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Name(doc.FormatComponent(component))
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

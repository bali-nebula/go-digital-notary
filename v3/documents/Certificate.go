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
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	algorithm fra.QuoteLike,
	publicKey fra.BinaryLike,
	tag fra.TagLike,
	version fra.VersionLike,
	previous any,
) CertificateLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(publicKey) {
		panic("The \"publicKey\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(previous) {
		panic("The \"previous\" attribute is required by this class.")
	}

	var citation string
	switch actual := previous.(type) {
	case fra.PatternLike:
		citation = actual.AsString()
		if citation != "none" {
			var message = fmt.Sprintf(
				"An invalid previous pattern was passed: %v",
				citation,
			)
			panic(message)
		}
	case fra.ResourceLike:
		citation = actual.AsString()
	case CitationLike:
		citation = actual.AsResource().AsString()
	default:
		var message = fmt.Sprintf(
			"An invalid previous citation type was passed: %T",
			actual,
		)
		panic(message)
	}

	var component = doc.ParseSource(`[
    $algorithm: ` + algorithm.AsString() + `
    $publicKey: ` + publicKey.AsString() + `
](
    $type: <bali:/nebula/types/Certificate:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $permissions: <bali:/nebula/permissions/public:v3>
    $previous: ` + citation + `
)`,
	)

	var instance = &certificate_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}

	return instance
}

func (c *certificateClass_) CertificateFromString(
	source string,
) CertificateLike {
	var component = doc.ParseSource(source)
	var instance = &certificate_{
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

func (v *certificate_) GetClass() CertificateClassLike {
	return certificateClass()
}

func (v *certificate_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *certificate_) GetAlgorithm() fra.QuoteLike {
	var object = v.GetObject(fra.Symbol("algorithm"))
	var quote = fra.QuoteFromString(doc.FormatComponent(object))
	return quote
}

func (v *certificate_) GetPublicKey() fra.BinaryLike {
	var object = v.GetObject(fra.Symbol("publicKey"))
	var binary = fra.BinaryFromString(doc.FormatComponent(object))
	return binary
}

// Attribute Methods

// Parameterized Methods

func (v *certificate_) GetType() fra.ResourceLike {
	var document = v.GetParameter(fra.Symbol("type"))
	return fra.ResourceFromString(doc.FormatComponent(document))
}

func (v *certificate_) GetTag() fra.TagLike {
	var document = v.GetParameter(fra.Symbol("tag"))
	return fra.TagFromString(doc.FormatComponent(document))
}

func (v *certificate_) GetVersion() fra.VersionLike {
	var document = v.GetParameter(fra.Symbol("version"))
	return fra.VersionFromString(doc.FormatComponent(document))
}

func (v *certificate_) GetPermissions() fra.ResourceLike {
	var document = v.GetParameter(fra.Symbol("permissions"))
	return fra.ResourceFromString(doc.FormatComponent(document))
}

func (v *certificate_) GetPrevious() any {
	var constraint = v.GetParameter(fra.Symbol("previous"))
	var source = doc.FormatComponent(constraint)
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

type certificate_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
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

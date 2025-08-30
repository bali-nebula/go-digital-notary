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

func DigestClass() DigestClassLike {
	return digestClass()
}

// Constructor Methods

func (c *digestClass_) Digest(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
) DigestLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(base64) {
		panic("The \"base64\" attribute is required by this class.")
	}

	var component = doc.ParseSource(`[
    $algorithm: ` + algorithm.AsString() + `
    $base64: ` + base64.AsString() + `
]($type: <bali:/nebula/types/Digest:v3>)`,
	)

	var instance = &digest_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *digestClass_) DigestFromString(
	source string,
) DigestLike {
	var component = doc.ParseSource(source)
	var instance = &digest_{
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

func (v *digest_) GetClass() DigestClassLike {
	return digestClass()
}

func (v *digest_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *digest_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *digest_) GetAlgorithm() fra.QuoteLike {
	var object = v.GetObject(fra.Symbol("algorithm"))
	return fra.QuoteFromString(doc.FormatComponent(object))
}

func (v *digest_) GetBase64() fra.BinaryLike {
	var object = v.GetObject(fra.Symbol("base64"))
	return fra.BinaryFromString(doc.FormatComponent(object))
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type digest_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type digestClass_ struct {
	// Declare the class constants.
}

// Class Reference

func digestClass() *digestClass_ {
	return digestClassReference_
}

var digestClassReference_ = &digestClass_{
	// Initialize the class constants.
}

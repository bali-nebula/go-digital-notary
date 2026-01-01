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

func SealClass() SealClassLike {
	return sealClass()
}

// Constructor Methods

func (c *sealClass_) Seal(
	algorithm doc.QuoteLike,
	signature doc.BinaryLike,
) SealLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(signature) {
		panic("The \"signature\" attribute is required by this class.")
	}

	var source = `[
    $algorithm: ` + algorithm.AsSource() + `
    $signature: ` + signature.AsSource() + `
]($type: /bali/types/notary/Seal/v3)`
	return c.SealFromSource(source)
}

func (c *sealClass_) SealFromSource(
	source string,
) SealLike {
	var component = doc.ParseComponent(source)
	var instance = &seal_{
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

func (v *seal_) GetClass() SealClassLike {
	return sealClass()
}

func (v *seal_) AsIntrinsic() doc.Composite {
	return v.Composite
}

func (v *seal_) AsSource() string {
	return doc.FormatComponent(v.Composite) + "\n"
}

// Attribute Methods

func (v *seal_) GetAlgorithm() doc.QuoteLike {
	var component = v.GetSubcomponent(doc.Symbol("$algorithm"))
	return doc.Quote(doc.FormatComponent(component))
}

func (v *seal_) GetSignature() doc.BinaryLike {
	var component = v.GetSubcomponent(doc.Symbol("$signature"))
	return doc.Binary(doc.FormatComponent(component))
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type seal_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Composite
}

// Class Structure

type sealClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sealClass() *sealClass_ {
	return sealClassReference_
}

var sealClassReference_ = &sealClass_{
	// Initialize the class constants.
}

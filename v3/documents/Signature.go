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

func SignatureClass() SignatureClassLike {
	return signatureClass()
}

// Constructor Methods

func (c *signatureClass_) Signature(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
) SignatureLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(base64) {
		panic("The \"base64\" attribute is required by this class.")
	}

	var component = doc.ParseSource(`[
    $algorithm: ` + algorithm.AsString() + `
    $base64: ` + base64.AsString() + `
]($type: <bali:/types/notary/Signature:v3>)`,
	)

	var instance = &signature_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *signatureClass_) SignatureFromString(
	source string,
) SignatureLike {
	var component = doc.ParseSource(source)
	var instance = &signature_{
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

func (v *signature_) GetClass() SignatureClassLike {
	return signatureClass()
}

func (v *signature_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *signature_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *signature_) GetAlgorithm() fra.QuoteLike {
	var object = v.GetObject(fra.Symbol("algorithm"))
	return fra.QuoteFromString(doc.FormatComponent(object))
}

func (v *signature_) GetBase64() fra.BinaryLike {
	var object = v.GetObject(fra.Symbol("base64"))
	return fra.BinaryFromString(doc.FormatComponent(object))
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type signature_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
}

// Class Structure

type signatureClass_ struct {
	// Declare the class constants.
}

// Class Reference

func signatureClass() *signatureClass_ {
	return signatureClassReference_
}

var signatureClassReference_ = &signatureClass_{
	// Initialize the class constants.
}

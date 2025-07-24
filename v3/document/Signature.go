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

package document

import (
	fmt "fmt"
	not "github.com/bali-nebula/go-document-notation/v3"
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
	var instance = &signature_{
		// Initialize the instance attributes.
		algorithm_: algorithm,
		base64_:    base64,
	}
	return instance
}

func (c *signatureClass_) SignatureFromString(
	source string,
) SignatureLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid signature was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var document = not.ParseSource(source)
	var algorithm = c.extractAlgorithm(document)
	var base64 = c.extractBase64(document)
	return c.Signature(algorithm, base64)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *signature_) GetClass() SignatureClassLike {
	return signatureClass()
}

func (v *signature_) AsString() string {
	var string_ = `[
`
	string_ += `    $algorithm: ` + v.GetAlgorithm().AsString()
	string_ += `    $base64: ` + v.GetBase64().AsString()
	string_ += `]
`
	var signature = not.ParseSource(string_)
	string_ = not.FormatDocument(signature)
	return string_
}

// Attribute Methods

func (v *signature_) GetAlgorithm() fra.QuoteLike {
	return v.algorithm_
}

func (v *signature_) GetBase64() fra.BinaryLike {
	return v.base64_
}

// PROTECTED INTERFACE

// Private Methods

func (c *signatureClass_) extractAlgorithm(
	document not.DocumentLike,
) fra.QuoteLike {
	var attribute = c.extractAttribute("$algorithm", document)
	var algorithm = fra.QuoteFromString(attribute)
	return algorithm
}

func (c *signatureClass_) extractAttribute(
	name string,
	document not.DocumentLike,
) string {
	var attribute string
	var key = not.Primitive(not.Element(name))
	document = not.GetAttribute(document, key)
	if uti.IsDefined(document) {
		attribute = not.FormatDocument(document)
		attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
	}
	return attribute
}

func (c *signatureClass_) extractBase64(
	document not.DocumentLike,
) fra.BinaryLike {
	var attribute = c.extractAttribute("$base64", document)
	var base64 = fra.BinaryFromString(attribute)
	return base64
}

// Instance Structure

type signature_ struct {
	// Declare the instance attributes.
	algorithm_ fra.QuoteLike
	base64_    fra.BinaryLike
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

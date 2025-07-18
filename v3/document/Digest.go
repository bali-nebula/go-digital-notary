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
	var instance = &digest_{
		// Initialize the instance attributes.
		algorithm_: algorithm,
		base64_:    base64,
	}
	return instance
}

func (c *digestClass_) DigestFromString(
	source string,
) DigestLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid digest was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var document = not.ParseSource(source)
	var algorithm = c.extractAlgorithm(document)
	var base64 = c.extractBase64(document)
	return c.Digest(algorithm, base64)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *digest_) GetClass() DigestClassLike {
	return digestClass()
}

func (v *digest_) AsString() string {
	var string_ = `[
`
	string_ += `    $algorithm: ` + v.GetAlgorithm().AsString()
	string_ += `    $base64: ` + v.GetBase64().AsString()
	string_ += `]
`
	var digest = not.ParseSource(string_)
	string_ = not.FormatDocument(digest)
	return string_
}

// Attribute Methods

func (v *digest_) GetAlgorithm() fra.QuoteLike {
	return v.algorithm_
}

func (v *digest_) GetBase64() fra.BinaryLike {
	return v.base64_
}

// PROTECTED INTERFACE

// Private Methods

func (c *digestClass_) extractAlgorithm(
	document not.DocumentLike,
) fra.QuoteLike {
	var attribute = c.extractAttribute("$algorithm", document)
	var algorithm = fra.QuoteFromString(attribute)
	return algorithm
}

func (c *digestClass_) extractAttribute(
	name string,
	document not.DocumentLike,
) string {
	var attribute string
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(not.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			attribute = not.FormatDocument(association.GetDocument())
			attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
			break
		}
	}
	return attribute
}

func (c *digestClass_) extractBase64(
	document not.DocumentLike,
) fra.BinaryLike {
	var attribute = c.extractAttribute("$base64", document)
	var base64 = fra.BinaryFromString(attribute)
	return base64
}

// Instance Structure

type digest_ struct {
	// Declare the instance attributes.
	algorithm_ fra.QuoteLike
	base64_    fra.BinaryLike
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

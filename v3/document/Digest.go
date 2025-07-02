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
	bal "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func DigestClass() DigestClassLike {
	return digestClass()
}

// Constructor Methods

func (c *digestClass_) Digest(
	algorithm string,
	base64 string,
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
	var document = bal.ParseSource(source)
	var algorithm = DocumentClass().ExtractAlgorithm("$algorithm", document)
	var base64 = DocumentClass().ExtractAttribute("$base64", document)
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
	string_ += `    $algorithm: "` + v.GetAlgorithm() + `"`
	string_ += `    $base64: ` + v.GetBase64()
	string_ += `]
`
	var digest = bal.ParseSource(string_)
	string_ = bal.FormatDocument(digest)
	return string_
}

// Attribute Methods

func (v *digest_) GetAlgorithm() string {
	return v.algorithm_
}

func (v *digest_) GetBase64() string {
	return v.base64_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type digest_ struct {
	// Declare the instance attributes.
	algorithm_ string
	base64_    string
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

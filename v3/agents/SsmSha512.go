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

package agents

import (
	dig "crypto/sha512"
	fmt "fmt"
)

// CLASS INTERFACE

// Access Function

func SsmSha512Class() SsmSha512ClassLike {
	return ssmSha512Class()
}

// Constructor Methods

func (c *ssmSha512Class_) SsmSha512() SsmSha512Like {
	var instance = &ssmSha512_{
		// Initialize the instance attributes.
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *ssmSha512_) GetClass() SsmSha512ClassLike {
	return ssmSha512Class()
}

// Attribute Methods

// Trusted Methods

func (v *ssmSha512_) GetDigestAlgorithm() string {
	return ssmSha512Class().algorithm_
}

func (v *ssmSha512_) DigestBytes(
	bytes []byte,
) []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to digest bytes",
	)

	var array = dig.Sum512(bytes)
	var digest = array[:] // Convert the [64]byte array to a slice.
	return digest
}

// PROTECTED INTERFACE

// Private Methods

func (v *ssmSha512_) errorCheck(
	message string,
) {
	if e := recover(); e != nil {
		message = fmt.Sprintf(
			"SsmSha512: %s:\n        %v",
			message,
			e,
		)
		panic(message)
	}
}

// Instance Structure

type ssmSha512_ struct {
	// Declare the instance attributes.
}

// Class Structure

type ssmSha512Class_ struct {
	// Declare the class constants.
	algorithm_ string
}

// Class Reference

func ssmSha512Class() *ssmSha512Class_ {
	return ssmSha512ClassReference_
}

var ssmSha512ClassReference_ = &ssmSha512Class_{
	// Initialize the class constants.
	algorithm_: "SHA512",
}

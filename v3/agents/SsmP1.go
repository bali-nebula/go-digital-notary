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

package agents

import (
	dig "crypto/sha512"
	fmt "fmt"
)

// CLASS INTERFACE

// Access Function

func SsmP1Class() SsmP1ClassLike {
	return ssmP1Class()
}

// Constructor Methods

func (c *ssmP1Class_) SsmP1() SsmP1Like {
	var instance = &ssmP1_{
		// Initialize the instance attributes.
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *ssmP1_) GetClass() SsmP1ClassLike {
	return ssmP1Class()
}

// Attribute Methods

// Trusted Methods

func (v *ssmP1_) GetDigestAlgorithm() string {
	return ssmP1Class().algorithm_
}

func (v *ssmP1_) DigestBytes(
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

func (v *ssmP1_) errorCheck(
	message string,
) {
	if e := recover(); e != nil {
		message = fmt.Sprintf(
			"SsmP1: %s:\n        %v",
			message,
			e,
		)
		panic(message)
	}
}

// Instance Structure

type ssmP1_ struct {
	// Declare the instance attributes.
}

// Class Structure

type ssmP1Class_ struct {
	// Declare the class constants.
	algorithm_ string
}

// Class Reference

func ssmP1Class() *ssmP1Class_ {
	return ssmP1ClassReference_
}

var ssmP1ClassReference_ = &ssmP1Class_{
	// Initialize the class constants.
	algorithm_: "SHA512",
}

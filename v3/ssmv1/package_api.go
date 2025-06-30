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

/*
Package "ssmv1" provides implementations of a software security module (SSM)
that are compliant with version 1 of the Bali Nebula security protocol.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-digital-notary/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ssmv1

import ()

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
SsmV1ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
software-security-module-like class.
*/
type SsmV1ClassLike interface {
	// Constructor Methods
	SsmV1(
		directory string,
	) SsmV1Like
}

// INSTANCE DECLARATIONS

/*
SsmV1Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete software-security-module-like class.
*/
type SsmV1Like interface {
	// Principal Methods
	GetClass() SsmV1ClassLike

	// Aspect Interfaces
	V1Secure
}

// ASPECT DECLARATIONS

/*
V1Secure declares the set of method signatures that must be supported by all
version 1 compatible security modules.
*/
type V1Secure interface {
	GetProtocolVersion() string
	GetDigestAlgorithm() string
	GetSignatureAlgorithm() string
	DigestBytes(
		bytes []byte,
	) []byte
	IsValid(
		key []byte,
		signature []byte,
		bytes []byte,
	) bool
	GetTag() string
	GenerateKeys() []byte
	SignBytes(
		bytes []byte,
	) []byte
	RotateKeys() []byte
	EraseKeys()
}

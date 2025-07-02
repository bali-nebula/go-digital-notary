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
Package "ssm" provides implementations of a software security module (SSM)
that are compliant with Bali Nebula security protocol.

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
package ssm

import (
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
SsmClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
software-security-module-like class.

A software security module (SSM) may be used in place of a hardware security
module (HSM) for testing purposes, or in a physically secure environment like
the cloud.
*/
type SsmClassLike interface {
	// Constructor Methods
	Ssm() SsmLike
}

// INSTANCE DECLARATIONS

/*
SsmLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete software-security-module-like class.
*/
type SsmLike interface {
	// Principal Methods
	GetClass() SsmClassLike

	// Aspect Interfaces
	not.Hardened
	not.Trusted
}

// ASPECT DECLARATIONS

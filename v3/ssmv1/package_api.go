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

import (
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
These constants define the possible states for the state machine.
*/
const (
	invalid fra.State = iota
	keyless
	loneKey
	twoKeys
)

/*
These constants define the possible events for the state machine.
*/
const (
	none fra.Event = iota
	generateKeys
	signBytes
	rotateKeys
)

/*
This list defines the event headings for the state machine.
*/
var events = []fra.Event{generateKeys, signBytes, rotateKeys}

/*
This table defines the allowed transitions for the state machine.
*/
var transitions = map[fra.State]fra.Transitions{
	keyless: fra.Transitions{loneKey, invalid, invalid},
	loneKey: fra.Transitions{invalid, loneKey, twoKeys},
	twoKeys: fra.Transitions{invalid, loneKey, invalid},
}

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
	not.V1Secure
}

// ASPECT DECLARATIONS

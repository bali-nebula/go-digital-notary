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
Package "agents" provides an implementation of a digital notary that can be used
to digitally notarize Bali documents.  The digital notary delegates the actual
security sensitive operations to one or more versions of a security module.

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
package agents

import (
	not "github.com/bali-nebula/go-digital-notary/v3/documents"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
DigitalNotaryClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete digital-notary-like class.

A digital notary may be used to digitally notarize digital documents using a
hardware security module (HSM). It may also be used to validate the seal on a
document that was notarized using this or any other digital notary.
*/
type DigitalNotaryClassLike interface {
	// Constructor Methods
	DigitalNotary(
		directory string,
		ssm Trusted,
		hsm Hardened,
	) DigitalNotaryLike
}

/*
SsmP1ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
software-security-module-p1-like class.
*/
type SsmP1ClassLike interface {
	// Constructor Methods
	SsmP1() SsmP1Like
}

/*
HsmP1ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
hardward-security-module-p1-like class.
*/
type HsmP1ClassLike interface {
	// Constructor Methods
	HsmP1(
		device string,
	) HsmP1Like
}

/*
TsmP1ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
test-security-module-p1-like class.

A test security module (TSM) should only be used in place of an actual hardware
security module (HSM) for testing purposes only, or in a physically secure
environment like the cloud.
*/
type TsmP1ClassLike interface {
	// Constructor Methods
	TsmP1(
		directory string,
	) TsmP1Like
}

// INSTANCE DECLARATIONS

/*
DigitalNotaryLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete digital-notary-like class.
*/
type DigitalNotaryLike interface {
	// Principal Methods
	GetClass() DigitalNotaryClassLike
	CiteDocument(
		document not.DocumentLike,
	) not.CitationLike
	CitationMatches(
		citation not.CitationLike,
		document not.DocumentLike,
	) bool
	GenerateKey() not.DocumentLike
	RefreshKey() not.DocumentLike
	ForgetKey()
	GenerateCredential(
		context any,
	) not.DocumentLike
	RefreshCredential(
		credential not.DocumentLike,
		context any,
	) not.DocumentLike
	NotarizeDocument(
		document not.DocumentLike,
	)
	SealMatches(
		document not.DocumentLike,
		certificate not.DocumentLike,
	) bool
}

/*
SsmP1Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete software-security-module-p1-like class.
*/
type SsmP1Like interface {
	// Principal Methods
	GetClass() SsmP1ClassLike

	// Aspect Interfaces
	Trusted
}

/*
HsmP1Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete hardware-security-module-p1-like class.
*/
type HsmP1Like interface {
	// Principal Methods
	GetClass() HsmP1ClassLike

	// Aspect Interfaces
	Hardened
}

/*
TsmP1Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete test-security-module-p1-like class.
*/
type TsmP1Like interface {
	// Principal Methods
	GetClass() TsmP1ClassLike

	// Aspect Interfaces
	Hardened
}

// ASPECT DECLARATIONS

/*
Trusted declares the set of method signatures that must be supported by all
trusted security modules.  No private key is needed by this interface.
*/
type Trusted interface {
	GetDigestAlgorithm() string
	DigestBytes(
		bytes []byte,
	) []byte
}

/*
Hardened declares the set of method signatures that must be supported by all
hardened security modules.  This interface requires a private key.
*/
type Hardened interface {
	GetTag() string
	GetSignatureAlgorithm() string
	GenerateKeys() []byte
	SignBytes(
		bytes []byte,
	) []byte
	IsValid(
		key []byte,
		bytes []byte,
		signature []byte,
	) bool
	RotateKeys() []byte
	EraseKeys()
}

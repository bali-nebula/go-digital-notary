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
	doc "github.com/bali-nebula/go-bali-documents/v3"
	com "github.com/bali-nebula/go-digital-notary/v3/components"
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
		ssm Trusted,
		hsm Hardened,
	) DigitalNotaryLike
	DigitalNotaryWithCertificate(
		ssm Trusted,
		hsm Hardened,
		certificate com.DocumentLike,
	) DigitalNotaryLike
}

/*
SsmSha512ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
software-security-module-sha512-like class.
*/
type SsmSha512ClassLike interface {
	// Constructor Methods
	SsmSha512() SsmSha512Like
}

/*
HsmEd25519ClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
hardward-security-module-ed25519-like class.
*/
type HsmEd25519ClassLike interface {
	// Constructor Methods
	HsmEd25519(
		device string,
	) HsmEd25519Like
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
		document com.DocumentLike,
	) com.CitationLike
	CitationMatches(
		citation com.CitationLike,
		document com.DocumentLike,
	) bool
	GenerateKey(
		attributes doc.Composite,
	) com.DocumentLike
	RefreshKey() com.DocumentLike
	ForgetKey()
	GenerateCredential(
		context any,
	) com.DocumentLike
	RefreshCredential(
		context any,
		document com.DocumentLike,
	) com.DocumentLike
	NotarizeDocument(
		document com.DocumentLike,
	)
	SealMatches(
		document com.DocumentLike,
		certificate com.DocumentLike,
	) bool
}

/*
SsmSha512Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete software-security-module-sha512-like class.
*/
type SsmSha512Like interface {
	// Principal Methods
	GetClass() SsmSha512ClassLike

	// Aspect Interfaces
	Trusted
}

/*
HsmEd25519Like is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete hardware-security-module-ed25519-like class.
*/
type HsmEd25519Like interface {
	// Principal Methods
	GetClass() HsmEd25519ClassLike

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
	GetPublicKey() []byte
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

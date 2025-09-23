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
Package "notary" provides an implementation of a digital notary that can be used
to digitally notarize Bali documents.

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
package notary

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
	not "github.com/bali-nebula/go-digital-notary/v3/documents"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
DigitalNotaryClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete digital-notary-like class.

A digital notary may be used to digitally notarize digital draft documents using
a hardware security module (HSM). It may also be used to validate the seal on a
contract that was notarized using this or any other digital notary.
*/
type DigitalNotaryClassLike interface {
	// Constructor Methods
	DigitalNotary(
		directory string,
		ssm Trusted,
		hsm Hardened,
	) DigitalNotaryLike
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
	GenerateKey() not.ContractLike
	RefreshKey() not.ContractLike
	ForgetKey()
	GenerateCredential(
		tag doc.TagLike,
		version doc.VersionLike,
	) not.ContractLike
	NotarizeDocument(
		draft not.Parameterized,
	) not.ContractLike
	SealMatches(
		document not.ContractLike,
		certificate not.CertificateLike,
	) bool
	CiteDocument(
		draft not.Parameterized,
	) doc.ResourceLike
	CitationMatches(
		citation doc.ResourceLike,
		draft not.Parameterized,
	) bool
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
	IsValid(
		key []byte,
		seal []byte,
		bytes []byte,
	) bool
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
	RotateKeys() []byte
	EraseKeys()
}

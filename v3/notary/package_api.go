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
to digitally sign (notarize) digital documents.

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
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
)

// TYPE DECLARATIONS

/*
Salt is a constrained type representing the random salt added to a hash
generation.
*/
type Salt []byte

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
NotaryClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
notary-like class.
*/
type NotaryClassLike interface {
	// Constructor Methods
	Notary(
		directory string,
		hsm V1Secure,
	) NotaryLike
}

// INSTANCE DECLARATIONS

/*
NotaryLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete notary-like class.
*/
type NotaryLike interface {
	// Principal Methods
	GetClass() NotaryClassLike
	GenerateKey() doc.ContractLike
	GetCitation() doc.CitationLike
	RefreshKey() doc.ContractLike
	ForgetKey()
	GenerateCredential(
		salt Salt,
	) doc.ContractLike
	NotarizeDocument(
		document doc.DocumentLike,
	) doc.ContractLike
	SignatureMatches(
		contract doc.ContractLike,
		certificate doc.CertificateLike,
	) bool
	CiteDocument(
		document doc.DocumentLike,
	) doc.CitationLike
	CitationMatches(
		citation doc.CitationLike,
		document doc.DocumentLike,
	) bool
}

// ASPECT DECLARATIONS

/*
V1Secure declares the set of method signatures that must be supported by all
version 1 compatible security modules.
*/
type V1Secure interface {
	GetProtocol() string
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

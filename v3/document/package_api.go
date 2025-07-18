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
Package "document" provides an implementation of wrappers for various types of
Bali Document Notation™ documents that are required by digital notarization.

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
package document

import (
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
CertificateClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete certificate-like class.
*/
type CertificateClassLike interface {
	// Constructor Methods
	Certificate(
		algorithm fra.QuoteLike,
		publicKey fra.BinaryLike,
		tag fra.TagLike,
		version fra.VersionLike,
		optionalPrevious CitationLike,
	) CertificateLike
	CertificateFromString(
		source string,
	) CertificateLike
}

/*
CitationClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete citation-like class.
*/
type CitationClassLike interface {
	// Constructor Methods
	Citation(
		tag fra.TagLike,
		version fra.VersionLike,
		digest DigestLike,
	) CitationLike
	CitationFromString(
		source string,
	) CitationLike
}

/*
ContractClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete contract-like class.
*/
type ContractClassLike interface {
	// Constructor Methods
	Contract(
		draft DraftLike,
		account fra.TagLike,
		certificate CitationLike,
	) ContractLike
	ContractFromString(
		source string,
	) ContractLike
}

/*
DigestClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete digest-like class.
*/
type DigestClassLike interface {
	// Constructor Methods
	Digest(
		algorithm fra.QuoteLike,
		base64 fra.BinaryLike,
	) DigestLike
	DigestFromString(
		source string,
	) DigestLike
}

/*
DraftClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete draft-like class.
*/
type DraftClassLike interface {
	// Constructor Methods
	Draft(
		component not.ComponentLike,
		type_ fra.ResourceLike,
		tag fra.TagLike,
		version fra.VersionLike,
		permissions fra.ResourceLike,
		previous CitationLike,
	) DraftLike
	DraftFromString(
		source string,
	) DraftLike
}

/*
SignatureClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete signature-like class.
*/
type SignatureClassLike interface {
	// Constructor Methods
	Signature(
		algorithm fra.QuoteLike,
		base64 fra.BinaryLike,
	) SignatureLike
	SignatureFromString(
		source string,
	) SignatureLike
}

// INSTANCE DECLARATIONS

/*
CertificateLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete certificate-like class.
*/
type CertificateLike interface {
	// Principal Methods
	GetClass() CertificateClassLike
	AsString() string

	// Attribute Methods
	GetAlgorithm() fra.QuoteLike
	GetPublicKey() fra.BinaryLike

	// Aspect Interfaces
	Parameterized
}

/*
CitationLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete citation-like class.
*/
type CitationLike interface {
	// Principal Methods
	GetClass() CitationClassLike
	AsString() string

	// Attribute Methods
	GetTag() fra.TagLike
	GetVersion() fra.VersionLike
	GetDigest() DigestLike
}

/*
ContractLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete contract-like class.
*/
type ContractLike interface {
	// Principal Methods
	GetClass() ContractClassLike
	AsString() string

	// Attribute Methods
	GetDraft() DraftLike
	GetAccount() fra.TagLike
	GetCertificate() CitationLike
	GetSignature() SignatureLike
	SetSignature(
		signature SignatureLike,
	)
}

/*
DigestLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete digest-like class.
*/
type DigestLike interface {
	// Principal Methods
	GetClass() DigestClassLike
	AsString() string

	// Attribute Methods
	GetAlgorithm() fra.QuoteLike
	GetBase64() fra.BinaryLike
}

/*
DraftLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete draft-like class.
*/
type DraftLike interface {
	// Principal Methods
	GetClass() DraftClassLike
	AsString() string

	// Attribute Methods
	GetComponent() not.ComponentLike

	// Aspect Interfaces
	Parameterized
}

/*
SignatureLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete signature-like class.
*/
type SignatureLike interface {
	// Principal Methods
	GetClass() SignatureClassLike
	AsString() string

	// Attribute Methods
	GetAlgorithm() fra.QuoteLike
	GetBase64() fra.BinaryLike
}

// ASPECT DECLARATIONS

/*
Parameterized declares the set of method signatures that must be supported by
all parameterized documents.
*/
type Parameterized interface {
	GetType() fra.ResourceLike
	GetTag() fra.TagLike
	GetVersion() fra.VersionLike
	GetPermissions() fra.ResourceLike
	GetOptionalPrevious() CitationLike
}

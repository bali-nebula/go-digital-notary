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
Package "documents" provides an implementation of wrappers for various types of
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
package documents

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
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
		key KeyLike,
		account fra.TagLike,
		signatory fra.ResourceLike,
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
		isNotarized fra.BooleanLike,
		tag fra.TagLike,
		version fra.VersionLike,
		digest DigestLike,
	) CitationLike
	CitationFromResource(
		resource fra.ResourceLike,
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
		draft Parameterized,
		account fra.TagLike,
		signatory fra.ResourceLike,
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
		entity any,
		type_ fra.ResourceLike,
		tag fra.TagLike,
		version fra.VersionLike,
		permissions fra.ResourceLike,
		optionalPrevious fra.ResourceLike,
	) DraftLike
	DraftFromString(
		source string,
	) DraftLike
}

/*
KeyClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete key-like class.
*/
type KeyClassLike interface {
	// Constructor Methods
	Key(
		algorithm fra.QuoteLike,
		base64 fra.BinaryLike,
		tag fra.TagLike,
		version fra.VersionLike,
	) KeyLike
	KeyFromString(
		source string,
	) KeyLike
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
	AsIntrinsic() doc.ComponentLike
	AsString() string
	GetKey() KeyLike

	// Aspect Interfaces
	Notarized
}

/*
CitationLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete citation-like class.
*/
type CitationLike interface {
	// Principal Methods
	GetClass() CitationClassLike
	AsIntrinsic() doc.ComponentLike
	AsResource() fra.ResourceLike
	AsString() string
	IsNotarized() fra.BooleanLike
	GetTag() fra.TagLike
	GetVersion() fra.VersionLike
	GetDigest() DigestLike

	// Aspect Interfaces
	doc.Declarative
}

/*
ContractLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete contract-like class.
*/
type ContractLike interface {
	// Principal Methods
	GetClass() ContractClassLike
	AsIntrinsic() doc.ComponentLike
	AsString() string
	GetDraft() Parameterized

	// Aspect Interfaces
	Notarized
}

/*
DigestLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete digest-like class.
*/
type DigestLike interface {
	// Principal Methods
	GetClass() DigestClassLike
	AsIntrinsic() doc.ComponentLike
	AsString() string
	GetAlgorithm() fra.QuoteLike
	GetBase64() fra.BinaryLike

	// Aspect Interfaces
	doc.Declarative
}

/*
DraftLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete draft-like class.
*/
type DraftLike interface {
	// Principal Methods
	GetClass() DraftClassLike
	AsIntrinsic() doc.ComponentLike
	AsString() string

	// Aspect Interfaces
	Parameterized
}

/*
KeyLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete key-like class.
*/
type KeyLike interface {
	// Principal Methods
	GetClass() KeyClassLike
	AsIntrinsic() doc.ComponentLike
	AsString() string
	GetCreated() fra.MomentLike
	GetAlgorithm() fra.QuoteLike
	GetBase64() fra.BinaryLike

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
	AsIntrinsic() doc.ComponentLike
	AsString() string
	GetAlgorithm() fra.QuoteLike
	GetBase64() fra.BinaryLike

	// Aspect Interfaces
	doc.Declarative
}

// ASPECT DECLARATIONS

/*
Notarized declares the set of method signatures that must be supported by
all notarized documents.
*/
type Notarized interface {
	doc.Declarative
	AsString() string
	GetContent() Parameterized
	GetAccount() fra.TagLike
	GetSignatory() fra.ResourceLike
	GetSignature() SignatureLike
	SetSignature(
		signature SignatureLike,
	)
	RemoveSignature()
}

/*
Parameterized declares the set of method signatures that must be supported by
all parameterized documents.
*/
type Parameterized interface {
	doc.Declarative
	AsString() string
	GetEntity() any
	GetType() fra.ResourceLike
	GetTag() fra.TagLike
	GetVersion() fra.VersionLike
	GetPermissions() fra.ResourceLike
	GetOptionalPrevious() fra.ResourceLike
}

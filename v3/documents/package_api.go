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
		tag doc.TagLike,
		version doc.VersionLike,
		algorithm doc.QuoteLike,
		key doc.BinaryLike,
		optionalPrevious doc.ResourceLike,
	) CertificateLike
	CertificateFromSource(
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
		tag doc.TagLike,
		version doc.VersionLike,
		algorithm doc.QuoteLike,
		digest doc.BinaryLike,
	) CitationLike
	CitationFromResource(
		resource doc.ResourceLike,
	) CitationLike
	CitationFromSource(
		source string,
	) CitationLike
}

/*
CredentialClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete credential-like class.
*/
type CredentialClassLike interface {
	// Constructor Methods
	Credential(
		context any,
		tag doc.TagLike,
		version doc.VersionLike,
		optionalPrevious doc.ResourceLike,
	) CredentialLike
	CredentialFromSource(
		source string,
	) CredentialLike
}

/*
DocumentClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		content Parameterized,
	) DocumentLike
	DocumentFromSource(
		source string,
	) DocumentLike
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
		type_ doc.NameLike,
		tag doc.TagLike,
		version doc.VersionLike,
		permissions doc.NameLike,
		optionalPrevious doc.ResourceLike,
	) DraftLike
	DraftFromSource(
		source string,
	) DraftLike
}

/*
SealClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete seal-like class.
*/
type SealClassLike interface {
	// Constructor Methods
	Seal(
		algorithm doc.QuoteLike,
		signature doc.BinaryLike,
	) SealLike
	SealFromSource(
		source string,
	) SealLike
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
	GetAlgorithm() doc.QuoteLike
	GetKey() doc.BinaryLike

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
	AsIntrinsic() doc.ComponentLike
	AsSource() string
	AsResource() doc.ResourceLike
	GetTag() doc.TagLike
	GetVersion() doc.VersionLike
	GetAlgorithm() doc.QuoteLike
	GetDigest() doc.BinaryLike
}

/*
CredentialLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete credential-like class.
*/
type CredentialLike interface {
	// Principal Methods
	GetClass() CredentialClassLike
	AsIntrinsic() doc.ComponentLike
	GetContext() any

	// Aspect Interfaces
	Parameterized
}

/*
DocumentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike
	AsIntrinsic() doc.ComponentLike
	AsSource() string
	GetContent() Parameterized
	GetAccount() doc.TagLike
	GetTimestamp() doc.MomentLike
	SetNotary(
		account doc.TagLike,
		notary CitationLike,
	)
	GetNotary() CitationLike
	SetSeal(
		seal SealLike,
	)
	HasSeal() bool
	RemoveSeal() SealLike
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

	// Aspect Interfaces
	Parameterized
}

/*
SealLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete seal-like class.
*/
type SealLike interface {
	// Principal Methods
	GetClass() SealClassLike
	AsIntrinsic() doc.ComponentLike
	AsSource() string
	GetAlgorithm() doc.QuoteLike
	GetSignature() doc.BinaryLike
}

// ASPECT DECLARATIONS

/*
Parameterized declares the set of method signatures that must be supported by
all parameterized documents.
*/
type Parameterized interface {
	AsSource() string
	GetEntity() any
	GetType() doc.NameLike
	GetTag() doc.TagLike
	GetVersion() doc.VersionLike
	GetPermissions() doc.NameLike
	GetOptionalPrevious() doc.ResourceLike
}

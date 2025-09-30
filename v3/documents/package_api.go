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
		account doc.TagLike,
		tag doc.TagLike,
		version doc.VersionLike,
		algorithm doc.QuoteLike,
		key doc.BinaryLike,
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
		tag doc.TagLike,
		version doc.VersionLike,
		algorithm doc.QuoteLike,
		digest doc.BinaryLike,
	) CitationLike
	CitationFromResource(
		resource doc.ResourceLike,
	) CitationLike
	CitationFromString(
		source string,
	) CitationLike
}

/*
ContentClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete content-like class.
*/
type ContentClassLike interface {
	// Constructor Methods
	Content(
		entity any,
		type_ doc.ResourceLike,
		tag doc.TagLike,
		version doc.VersionLike,
		optionalPrevious doc.ResourceLike,
		permissions doc.ResourceLike,
		account doc.TagLike,
	) ContentLike
	ContentFromString(
		source string,
	) ContentLike
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
		account doc.TagLike,
		tag doc.TagLike,
		version doc.VersionLike,
	) CredentialLike
	CredentialFromString(
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
	DocumentFromString(
		source string,
	) DocumentLike
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
	SealFromString(
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
	AsString() string
	AsResource() doc.ResourceLike
	GetTag() doc.TagLike
	GetVersion() doc.VersionLike
	GetAlgorithm() doc.QuoteLike
	GetDigest() doc.BinaryLike
}

/*
ContentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete certificate-like class.
*/
type ContentLike interface {
	// Principal Methods
	GetClass() ContentClassLike
	AsIntrinsic() doc.ComponentLike

	// Aspect Interfaces
	Parameterized
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
	AsString() string
	GetContent() Parameterized
	GetTimestamp() doc.MomentLike
	GetNotary() CitationLike
	SetNotary(
		notary CitationLike,
	)
	HasSeal() bool
	SetSeal(
		seal SealLike,
	)
	RemoveSeal() SealLike
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
	AsString() string
	GetAlgorithm() doc.QuoteLike
	GetSignature() doc.BinaryLike
}

// ASPECT DECLARATIONS

/*
Parameterized declares the set of method signatures that must be supported by
all parameterized documents.
*/
type Parameterized interface {
	AsString() string
	GetEntity() any
	GetType() doc.ResourceLike
	GetTag() doc.TagLike
	GetVersion() doc.VersionLike
	GetOptionalPrevious() doc.ResourceLike
	GetPermissions() doc.ResourceLike
	GetAccount() doc.TagLike
}

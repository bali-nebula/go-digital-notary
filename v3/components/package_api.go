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
Package "components" provides an implementation of wrappers for various types of
Bali Document Notation™ components that are required by digital notarization.

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
package components

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
		algorithm doc.QuoteLike,
		key doc.BinaryLike,
		tag doc.TagLike,
		version doc.VersionLike,
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
ContentClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete content-like class.
*/
type ContentClassLike interface {
	// Constructor Methods
	Content(
		entity any,
		type_ doc.NameLike,
		tag doc.TagLike,
		version doc.VersionLike,
		permissions doc.NameLike,
		optionalPrevious doc.ResourceLike,
	) ContentLike
	ContentFromSource(
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
IdentityClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete identity-like class.
*/
type IdentityClassLike interface {
	// Constructor Methods
	Identity(
		name doc.QuoteLike,
		surname doc.QuoteLike,
		birthdate doc.MomentLike,
		birthplace doc.QuoteLike,
		sex doc.SymbolLike,
		nationality doc.QuoteLike,
		address doc.NarrativeLike,
		mobile doc.QuoteLike,
		email doc.QuoteLike,
		face doc.BinaryLike,
		tag doc.TagLike,
		version doc.VersionLike,
		optionalPrevious doc.ResourceLike,
	) IdentityLike
	IdentityFromSource(
		source string,
	) IdentityLike
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
ContentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete content-like class.
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
	AsSource() string
	GetContent() Parameterized
	GetOwner() doc.TagLike
	GetTimestamp() doc.MomentLike
	SetNotary(
		owner doc.TagLike,
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
IdentityLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete identity-like class.
*/
type IdentityLike interface {
	// Principal Methods
	GetClass() IdentityClassLike
	AsIntrinsic() doc.ComponentLike
	GetName() doc.QuoteLike
	GetSurname() doc.QuoteLike
	GetBirthdate() doc.MomentLike
	GetBirthplace() doc.QuoteLike
	GetSex() doc.SymbolLike
	GetNationality() doc.QuoteLike
	GetAddress() doc.NarrativeLike
	GetMobile() doc.QuoteLike
	GetEmail() doc.QuoteLike
	GetFace() doc.BinaryLike

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

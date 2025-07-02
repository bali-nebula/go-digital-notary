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
Package "document" provides an implementation of wrappers for Bali Document
Notation™ documents that are required by digital notarization.

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
	bal "github.com/bali-nebula/go-bali-documents/v3"
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
		algorithm string,
		publicKey string,
		tag string,
		version string,
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
		tag string,
		version string,
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
		document DocumentLike,
		account string,
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
		algorithm string,
		base64 string,
	) DigestLike
	DigestFromString(
		source string,
	) DigestLike
}

/*
DocumentClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		component bal.ComponentLike,
		type_ string,
		tag string,
		version string,
		permissions string,
		previous CitationLike,
	) DocumentLike
	DocumentFromString(
		source string,
	) DocumentLike

	// Function Methods
	ExtractAlgorithm(
		name string,
		document bal.DocumentLike,
	) string
	ExtractAttribute(
		name string,
		document bal.DocumentLike,
	) string
	ExtractCertificate(
		name string,
		document bal.DocumentLike,
	) CertificateLike
	ExtractCitation(
		name string,
		document bal.DocumentLike,
	) CitationLike
	ExtractDigest(
		name string,
		document bal.DocumentLike,
	) DigestLike
	ExtractDocument(
		name string,
		document bal.DocumentLike,
	) DocumentLike
	ExtractParameter(
		name string,
		document bal.DocumentLike,
	) string
	ExtractPrevious(
		name string,
		document bal.DocumentLike,
	) CitationLike
	ExtractSignature(
		name string,
		document bal.DocumentLike,
	) SignatureLike
}

/*
SignatureClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete signature-like class.
*/
type SignatureClassLike interface {
	// Constructor Methods
	Signature(
		algorithm string,
		base64 string,
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
	GetAlgorithm() string
	GetPublicKey() string

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
	GetTag() string
	GetVersion() string
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
	GetDocument() DocumentLike
	GetAccount() string
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
	GetAlgorithm() string
	GetBase64() string
}

/*
DocumentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike
	AsString() string

	// Attribute Methods
	GetComponent() bal.ComponentLike

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
	GetAlgorithm() string
	GetBase64() string
}

// ASPECT DECLARATIONS

/*
Parameterized declares the set of method signatures that must be supported by
all parameterized documents.
*/
type Parameterized interface {
	GetType() string
	GetTag() string
	GetVersion() string
	GetPermissions() string
	GetOptionalPrevious() CitationLike
}

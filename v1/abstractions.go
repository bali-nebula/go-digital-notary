/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package notary

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all notarized components.
type Notarized interface {
	GetDocument() DocumentLike
	GetTimestamp() MomentLike
	GetAccount() TagLike
	GetProtocol() VersionLike
	GetCertificate() CitationLike
	AddSignature(signature BinaryLike)
	RemoveSignature() BinaryLike
}

// This interface defines the methods supported by all published components.
type Published interface {
	GetAlgorithms() CatalogLike
	GetKey() BinaryLike
}

// This interface defines the methods supported by all referential components.
type Referential interface {
	GetTag() TagLike
	GetVersion() VersionLike
	GetProtocol() VersionLike
	GetDigest() BinaryLike
}

// This interface defines the methods supported by all restricted components.
type Restricted interface {
	GetPermissions() MonikerLike
}

// This interface defines the methods supported by all salted components.
type Salted interface {
	GetSalt() BinaryLike
}

// This interface defines the methods supported by all typed components.
type Typed interface {
	GetType() MonikerLike
}

// This interface defines the methods supported by all versioned components.
type Versioned interface {
	GetTag() TagLike
	GetVersion() VersionLike
	GetPrevious() CitationLike
}

// This interface defines the methods supported by all prudent notary
// agents.
type Prudent interface {
	GenerateKey() ContractLike
	GetCitation() CitationLike
	RefreshKey() ContractLike
	ForgetKey()
}

// This interface defines the methods supported by all certified notary
// agents.
type Certified interface {
	GenerateCredentials(salt BinaryLike) ContractLike
	NotarizeDocument(document DocumentLike) ContractLike
	SignatureMatches(contract ContractLike, certificate CertificateLike) bool
	CiteDocument(document DocumentLike) CitationLike
	CitationMatches(citation CitationLike, document DocumentLike) bool
}

// CONSOLIDATED INTERFACES

type TypeLike interface {
	Encapsulated
	Typed
}

type CitationLike interface {
	Encapsulated
	Referential
	Typed
}

type ContractLike interface {
	Encapsulated
	Notarized
	Typed
}

type DocumentLike interface {
	Encapsulated
	Typed
	Restricted
	Versioned
}

type CertificateLike interface {
	Encapsulated
	Published
	Typed
	Restricted
	Versioned
}

type CredentialsLike interface {
	Encapsulated
	Salted
	Typed
	Restricted
	Versioned
}

type NotaryLike interface {
	Prudent
	Certified
}

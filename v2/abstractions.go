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

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all notarized components.
type Notarized interface {
	GetDocument() DocumentLike
	GetAccount() abs.TagLike
	GetProtocol() abs.VersionLike
	GetCertificate() CitationLike
	AddSignature(signature abs.BinaryLike)
	RemoveSignature() abs.BinaryLike
}

// This interface defines the methods supported by all published components.
type Published interface {
	GetAlgorithms() abs.CatalogLike
	GetKey() abs.BinaryLike
}

// This interface defines the methods supported by all referential components.
type Referential interface {
	GetTag() abs.TagLike
	GetVersion() abs.VersionLike
	GetProtocol() abs.VersionLike
	GetDigest() abs.BinaryLike
}

// This interface defines the methods supported by all restricted components.
type Restricted interface {
	GetPermissions() abs.MonikerLike
}

// This interface defines the methods supported by all salted components.
type Seasoned interface {
	GetSalt() abs.BinaryLike
}

// This interface defines the methods supported by all typed components.
type Typed interface {
	GetType() abs.MonikerLike
}

// This interface defines the methods supported by all versioned components.
type Versioned interface {
	GetTag() abs.TagLike
	GetVersion() abs.VersionLike
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
	GenerateCredentials(salt abs.BinaryLike) ContractLike
	NotarizeDocument(document DocumentLike) ContractLike
	SignatureMatches(contract ContractLike, certificate CertificateLike) bool
	CiteDocument(document DocumentLike) CitationLike
	CitationMatches(citation CitationLike, document DocumentLike) bool
}

// This interface defines the methods supported by all trusted security
// modules.
type Trusted interface {
	GetProtocol() string
	DigestBytes(bytes []byte) []byte
	IsValid(key []byte, signature []byte, bytes []byte) bool
}

// This interface defines the methods supported by all hardened security
// modules.
type Hardened interface {
	GetTag() string
	GenerateKeys() []byte
	SignBytes(bytes []byte) []byte
	RotateKeys() []byte
	EraseKeys()
}

// CONSOLIDATED INTERFACES

type TypeLike interface {
	abs.Encapsulated
	Typed
}

type CitationLike interface {
	abs.Encapsulated
	Referential
	Typed
}

type ContractLike interface {
	abs.Encapsulated
	Notarized
	Typed
}

type DocumentLike interface {
	abs.Encapsulated
	Typed
	Restricted
	Versioned
}

type CertificateLike interface {
	abs.Encapsulated
	Published
	Typed
	Restricted
	Versioned
}

type CredentialsLike interface {
	abs.Encapsulated
	Seasoned
	Typed
	Restricted
	Versioned
}

type NotaryLike interface {
	Prudent
	Certified
}

// This interface consolidates all the interfaces supported by
// security-module-like devices.
type SecurityModuleLike interface {
	Trusted
	Hardened
}

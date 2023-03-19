/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
)

// CONSTANT DEFINITIONS

// These constants define the attribute names for the standard attribues.
var (
	TagAttribute         = bal.Symbol("$tag")
	AccountAttribute     = bal.Symbol("$account")
	AlgorithmsAttribute  = bal.Symbol("$algorithms")
	CertificateAttribute = bal.Symbol("$certificate")
	DigestAttribute      = bal.Symbol("$digest")
	DocumentAttribute    = bal.Symbol("$document")
	KeyAttribute         = bal.Symbol("$key")
	PermissionsAttribute = bal.Symbol("$permissions")
	PreviousAttribute    = bal.Symbol("$previous")
	ProtocolAttribute    = bal.Symbol("$protocol")
	SaltAttribute        = bal.Symbol("$salt")
	SignatureAttribute   = bal.Symbol("$signature")
	TimestampAttribute   = bal.Symbol("$timestamp")
	TypeAttribute        = bal.Symbol("$type")
	VersionAttribute     = bal.Symbol("$version")
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
	GetType() TypeLike
}

// This interface defines the methods supported by all named components.
type Named interface {
	GetName() abs.MonikerLike
}

// This interface defines the methods supported by all versioned components.
type Versioned interface {
	GetTag() abs.TagLike
	GetVersion() abs.VersionLike
	GetPrevious() CitationLike
}

// CONSOLIDATED INTERFACES

type CertificateLike interface {
	abs.Encapsulated
	Published
	Restricted
	Typed
	Versioned
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

type CredentialLike interface {
	abs.Encapsulated
	Restricted
	Seasoned
	Typed
	Versioned
}

type DocumentLike interface {
	abs.Encapsulated
	Restricted
	Typed
	Versioned
}

type TypeLike interface {
	abs.Encapsulated
	Named
}

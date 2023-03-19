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
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all prudent notary agents.
type Prudent interface {
	GenerateKey() ContractLike
	GetCitation() CitationLike
	RefreshKey() ContractLike
	ForgetKey()
}

// This interface defines the methods supported by all certified notary agents.
type Certified interface {
	GenerateCredential(salt abs.BinaryLike) ContractLike
	NotarizeRecord(record RecordLike) ContractLike
	SignatureMatches(contract ContractLike, certificate CertificateLike) bool
	CiteRecord(record RecordLike) CitationLike
	CitationMatches(citation CitationLike, record RecordLike) bool
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

// This interface consolidates all the interfaces supported by
// notary-like devices.
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

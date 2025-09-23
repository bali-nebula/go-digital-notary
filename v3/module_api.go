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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-digital-notary/wiki
*/
package module

import (
	bal "github.com/bali-nebula/go-bali-documents/v3"
	doc "github.com/bali-nebula/go-digital-notary/v3/documents"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssm"
)

// TYPE ALIASES

// Documents

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContractClassLike    = doc.ContractClassLike
	CredentialClassLike  = doc.CredentialClassLike
	DraftClassLike       = doc.DraftClassLike
	SealClassLike        = doc.SealClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContractLike    = doc.ContractLike
	CredentialLike  = doc.CredentialLike
	DraftLike       = doc.DraftLike
	SealLike        = doc.SealLike
)

type (
	Parameterized = doc.Parameterized
)

// Notary

type (
	DigitalNotaryClassLike = not.DigitalNotaryClassLike
)

type (
	DigitalNotaryLike = not.DigitalNotaryLike
)

type (
	Trusted  = not.Trusted
	Hardened = not.Hardened
)

// Ssm

type (
	SsmClassLike = ssm.SsmClassLike
)

type (
	SsmLike = ssm.SsmLike
)

// CLASS ACCESSORS

// Documents

func CertificateClass() CertificateClassLike {
	return doc.CertificateClass()
}

func Certificate(
	arguments ...any,
) CertificateLike {
	if len(arguments) == 1 {
		var source = arguments[0].(string)
		return CertificateClass().CertificateFromString(source)
	}
	var algorithm = arguments[0].(bal.QuoteLike)
	var key = arguments[1].(bal.BinaryLike)
	var tag = arguments[2].(bal.TagLike)
	var version = arguments[3].(bal.VersionLike)
	return CertificateClass().Certificate(
		algorithm,
		key,
		tag,
		version,
	)
}

func CitationClass() CitationClassLike {
	return doc.CitationClass()
}

func Citation(
	arguments ...any,
) CitationLike {
	if len(arguments) == 1 {
		switch actual := arguments[0].(type) {
		case string:
			return CitationClass().CitationFromString(actual)
		default:
			var resource = arguments[0].(bal.ResourceLike)
			return CitationClass().CitationFromResource(resource)
		}
	}
	var algorithm = arguments[0].(bal.QuoteLike)
	var digest = arguments[1].(bal.BinaryLike)
	var tag = arguments[2].(bal.TagLike)
	var version = arguments[3].(bal.VersionLike)
	return CitationClass().Citation(
		algorithm,
		digest,
		tag,
		version,
	)
}

func ContractClass() ContractClassLike {
	return doc.ContractClass()
}

func Contract(
	arguments ...any,
) ContractLike {
	if len(arguments) == 1 {
		var source = arguments[0].(string)
		return ContractClass().ContractFromString(source)
	}
	var content = arguments[0].(doc.Parameterized)
	var account = arguments[1].(bal.TagLike)
	var notary = arguments[2].(bal.ResourceLike)
	return ContractClass().Contract(
		content,
		account,
		notary,
	)
}

func CredentialClass() CredentialClassLike {
	return doc.CredentialClass()
}

func Credential(
	arguments ...any,
) CredentialLike {
	if len(arguments) == 1 {
		var source = arguments[0].(string)
		return CredentialClass().CredentialFromString(source)
	}
	var tag = arguments[0].(bal.TagLike)
	var version = arguments[1].(bal.VersionLike)
	return CredentialClass().Credential(
		tag,
		version,
	)
}

func DraftClass() DraftClassLike {
	return doc.DraftClass()
}

func Draft(
	arguments ...any,
) DraftLike {
	if len(arguments) == 1 {
		var source = arguments[0].(string)
		return DraftClass().DraftFromString(source)
	}
	var entity = arguments[0]
	var type_ = arguments[1].(bal.ResourceLike)
	var tag = arguments[2].(bal.TagLike)
	var version = arguments[3].(bal.VersionLike)
	var permissions = arguments[4].(bal.ResourceLike)
	var optionalPrevious bal.ResourceLike
	if len(arguments) == 6 && arguments[5] != nil {
		optionalPrevious = arguments[5].(bal.ResourceLike)
	}
	return DraftClass().Draft(
		entity,
		type_,
		tag,
		version,
		permissions,
		optionalPrevious,
	)
}

func SealClass() SealClassLike {
	return doc.SealClass()
}

func Seal(
	arguments ...any,
) SealLike {
	if len(arguments) == 1 {
		var source = arguments[0].(string)
		return SealClass().SealFromString(source)
	}
	var algorithm = arguments[0].(bal.QuoteLike)
	var signature = arguments[1].(bal.BinaryLike)
	return SealClass().Seal(
		algorithm,
		signature,
	)
}

// Notary

func DigitalNotaryClass() DigitalNotaryClassLike {
	return not.DigitalNotaryClass()
}

func DigitalNotary(
	directory string,
	ssm not.Trusted,
	hsm not.Hardened,
) DigitalNotaryLike {
	return DigitalNotaryClass().DigitalNotary(
		directory,
		ssm,
		hsm,
	)
}

// Ssm

func SsmClass() SsmClassLike {
	return ssm.SsmClass()
}

func Ssm(
	directory string,
) SsmLike {
	return SsmClass().Ssm(
		directory,
	)
}

// GLOBAL FUNCTIONS

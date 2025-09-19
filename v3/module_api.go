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
	doc "github.com/bali-nebula/go-digital-notary/v3/documents"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssm"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE ALIASES

// Documents

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContractClassLike    = doc.ContractClassLike
	CredentialClassLike  = doc.CredentialClassLike
	DigestClassLike      = doc.DigestClassLike
	DraftClassLike       = doc.DraftClassLike
	KeyClassLike         = doc.KeyClassLike
	SealClassLike        = doc.SealClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContractLike    = doc.ContractLike
	CredentialLike  = doc.CredentialLike
	DigestLike      = doc.DigestLike
	DraftLike       = doc.DraftLike
	KeyLike         = doc.KeyLike
	SealLike        = doc.SealLike
)

type (
	Notarized     = doc.Notarized
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
	key KeyLike,
	account fra.TagLike,
	notary fra.ResourceLike,
) CertificateLike {
	return CertificateClass().Certificate(
		key,
		account,
		notary,
	)
}

func CertificateFromString(
	source string,
) CertificateLike {
	return CertificateClass().CertificateFromString(
		source,
	)
}

func CitationClass() CitationClassLike {
	return doc.CitationClass()
}

func Citation(
	tag fra.TagLike,
	version fra.VersionLike,
	digest doc.DigestLike,
) CitationLike {
	return CitationClass().Citation(
		tag,
		version,
		digest,
	)
}

func CitationFromResource(
	resource fra.ResourceLike,
) CitationLike {
	return CitationClass().CitationFromResource(
		resource,
	)
}

func CitationFromString(
	source string,
) CitationLike {
	return CitationClass().CitationFromString(
		source,
	)
}

func ContractClass() ContractClassLike {
	return doc.ContractClass()
}

func Contract(
	entity doc.Parameterized,
	account fra.TagLike,
	certificate fra.ResourceLike,
) ContractLike {
	return ContractClass().Contract(
		entity,
		account,
		certificate,
	)
}

func ContractFromString(
	source string,
) ContractLike {
	return ContractClass().ContractFromString(
		source,
	)
}

func CredentialClass() CredentialClassLike {
	return doc.CredentialClass()
}

func Credential(
	account fra.TagLike,
	notary fra.ResourceLike,
) CredentialLike {
	return CredentialClass().Credential(
		account,
		notary,
	)
}

func CredentialFromString(
	source string,
) CredentialLike {
	return CredentialClass().CredentialFromString(
		source,
	)
}

func DigestClass() DigestClassLike {
	return doc.DigestClass()
}

func Digest(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
) DigestLike {
	return DigestClass().Digest(
		algorithm,
		base64,
	)
}

func DigestFromString(
	source string,
) DigestLike {
	return DigestClass().DigestFromString(
		source,
	)
}

func DraftClass() DraftClassLike {
	return doc.DraftClass()
}

func Draft(
	entity any,
	type_ fra.ResourceLike,
	tag fra.TagLike,
	version fra.VersionLike,
	permissions fra.ResourceLike,
	optionalPrevious fra.ResourceLike,
) DraftLike {
	return DraftClass().Draft(
		entity,
		type_,
		tag,
		version,
		permissions,
		optionalPrevious,
	)
}

func DraftFromString(
	source string,
) DraftLike {
	return DraftClass().DraftFromString(
		source,
	)
}

func KeyClass() KeyClassLike {
	return doc.KeyClass()
}

func Key(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
	tag fra.TagLike,
	version fra.VersionLike,
) KeyLike {
	return KeyClass().Key(
		algorithm,
		base64,
		tag,
		version,
	)
}

func KeyFromString(
	source string,
) KeyLike {
	return KeyClass().KeyFromString(
		source,
	)
}

func SealClass() SealClassLike {
	return doc.SealClass()
}

func Seal(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
) SealLike {
	return SealClass().Seal(
		algorithm,
		base64,
	)
}

func SealFromString(
	source string,
) SealLike {
	return SealClass().SealFromString(
		source,
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

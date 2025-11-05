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
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	age "github.com/bali-nebula/go-digital-notary/v3/agents"
	com "github.com/bali-nebula/go-digital-notary/v3/components"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// TYPE ALIASES

// Documents

type (
	CertificateClassLike = com.CertificateClassLike
	CitationClassLike    = com.CitationClassLike
	ContentClassLike     = com.ContentClassLike
	CredentialClassLike  = com.CredentialClassLike
	DocumentClassLike    = com.DocumentClassLike
	IdentityClassLike    = com.IdentityClassLike
	SealClassLike        = com.SealClassLike
)

type (
	CertificateLike = com.CertificateLike
	CitationLike    = com.CitationLike
	ContentLike     = com.ContentLike
	CredentialLike  = com.CredentialLike
	DocumentLike    = com.DocumentLike
	IdentityLike    = com.IdentityLike
	SealLike        = com.SealLike
)

type (
	Parameterized = com.Parameterized
)

// Agents

type (
	DigitalNotaryClassLike = age.DigitalNotaryClassLike
)

type (
	DigitalNotaryLike = age.DigitalNotaryLike
)

type (
	Trusted  = age.Trusted
	Hardened = age.Hardened
)

type (
	HsmEd25519ClassLike = age.HsmEd25519ClassLike
)

type (
	HsmEd25519Like = age.HsmEd25519Like
)

type (
	SsmSha512ClassLike = age.SsmSha512ClassLike
)

type (
	SsmSha512Like = age.SsmSha512Like
)

// CLASS ACCESSORS

// Documents

func CertificateClass() CertificateClassLike {
	return com.CertificateClass()
}

func CitationClass() CitationClassLike {
	return com.CitationClass()
}

func ContentClass() ContentClassLike {
	return com.ContentClass()
}

func CredentialClass() CredentialClassLike {
	return com.CredentialClass()
}

func DocumentClass() DocumentClassLike {
	return com.DocumentClass()
}

func IdentityClass() IdentityClassLike {
	return com.IdentityClass()
}

func SealClass() SealClassLike {
	return com.SealClass()
}

// Agents

func DigitalNotaryClass() DigitalNotaryClassLike {
	return age.DigitalNotaryClass()
}

func DigitalNotary(
	authority com.DocumentLike,
	ssm age.Trusted,
	hsm age.Hardened,
) DigitalNotaryLike {
	return DigitalNotaryClass().DigitalNotary(
		authority,
		ssm,
		hsm,
	)
}

func HsmEd25519Class() HsmEd25519ClassLike {
	return age.HsmEd25519Class()
}

func HsmEd25519(
	device string,
) HsmEd25519Like {
	return HsmEd25519Class().HsmEd25519(
		device,
	)
}

func SsmSha512Class() SsmSha512ClassLike {
	return age.SsmSha512Class()
}

func SsmSha512() SsmSha512Like {
	return SsmSha512Class().SsmSha512()
}

// GLOBAL FUNCTIONS

// Documents

func Certificate(
	value ...any,
) CertificateLike {
	if len(value) == 1 {
		var source string
		switch actual := value[0].(type) {
		case string:
			source = actual
		case com.Parameterized:
			source = actual.AsSource()
		}
		return com.CertificateClass().CertificateFromSource(source)
	}
	var algorithm = value[0].(doc.QuoteLike)
	var key = value[1].(doc.BinaryLike)
	var tag = value[2].(doc.TagLike)
	var version = value[3].(doc.VersionLike)
	var previous = value[4].(doc.ResourceLike)
	return CertificateClass().Certificate(algorithm, key, tag, version, previous)
}

func Citation(
	value ...any,
) CitationLike {
	if len(value) == 1 {
		switch actual := value[0].(type) {
		case string:
			return com.CitationClass().CitationFromSource(actual)
		case doc.ResourceLike:
			return com.CitationClass().CitationFromResource(actual)
		default:
			var message = fmt.Sprintf(
				"An invalid argument type was passed into the Citation contructor: %v(%T)",
				actual,
				actual,
			)
			panic(message)
		}
	}
	var tag = value[0].(doc.TagLike)
	var version = value[1].(doc.VersionLike)
	var algorithm = value[2].(doc.QuoteLike)
	var digest = value[3].(doc.BinaryLike)
	return CitationClass().Citation(tag, version, algorithm, digest)
}

func Content(
	value ...any,
) ContentLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return com.ContentClass().ContentFromSource(source)
	}
	var entity = value[0]
	var type_ = value[1].(doc.NameLike)
	var tag = value[2].(doc.TagLike)
	var version = value[3].(doc.VersionLike)
	var permissions = value[4].(doc.NameLike)
	var optionalPrevious doc.ResourceLike
	if uti.IsDefined(value[5]) {
		optionalPrevious = value[5].(doc.ResourceLike)
	}
	return ContentClass().Content(
		entity,
		type_,
		tag,
		version,
		permissions,
		optionalPrevious,
	)
}

func Credential(
	value ...any,
) CredentialLike {
	if len(value) == 1 {
		var source string
		switch actual := value[0].(type) {
		case string:
			source = actual
		case com.Parameterized:
			source = actual.AsSource()
		}
		return com.CredentialClass().CredentialFromSource(source)
	}
	var context = value[0]
	var tag = value[1].(doc.TagLike)
	var version = value[2].(doc.VersionLike)
	var previous = value[3].(doc.ResourceLike)
	return CredentialClass().Credential(
		context,
		tag,
		version,
		previous,
	)
}

func Document(
	value any,
) DocumentLike {
	switch actual := value.(type) {
	case string:
		return DocumentClass().DocumentFromSource(actual)
	case com.Parameterized:
		return DocumentClass().Document(actual)
	default:
		var message = fmt.Sprintf(
			"An invalid argument type was passed into the Document contructor: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func Identity(
	value ...any,
) IdentityLike {
	if len(value) == 1 {
		var source string
		switch actual := value[0].(type) {
		case string:
			source = actual
		case com.Parameterized:
			source = actual.AsSource()
		}
		return com.IdentityClass().IdentityFromSource(source)
	}
	var surname = value[0].(doc.QuoteLike)
	var birthname = value[1].(doc.QuoteLike)
	var birthdate = value[2].(doc.MomentLike)
	var birthplace = value[3].(doc.QuoteLike)
	var birthsex = value[4].(doc.SymbolLike)
	var nationality = value[5].(doc.QuoteLike)
	var address = value[6].(doc.NarrativeLike)
	var mobile = value[7].(doc.QuoteLike)
	var email = value[8].(doc.QuoteLike)
	var mugshot = value[9].(doc.BinaryLike)
	var tag = value[10].(doc.TagLike)
	var version = value[11].(doc.VersionLike)
	var previous = value[12].(doc.ResourceLike)
	return IdentityClass().Identity(
		surname,
		birthname,
		birthdate,
		birthplace,
		birthsex,
		nationality,
		address,
		mobile,
		email,
		mugshot,
		tag,
		version,
		previous,
	)
}

func Seal(
	value ...any,
) SealLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return com.SealClass().SealFromSource(source)
	}
	var algorithm = value[0].(doc.QuoteLike)
	var signature = value[1].(doc.BinaryLike)
	return SealClass().Seal(algorithm, signature)
}

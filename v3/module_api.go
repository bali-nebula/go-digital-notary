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
	bal "github.com/bali-nebula/go-bali-documents/v3"
	age "github.com/bali-nebula/go-digital-notary/v3/agents"
	doc "github.com/bali-nebula/go-digital-notary/v3/documents"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// TYPE ALIASES

// Documents

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	CredentialClassLike  = doc.CredentialClassLike
	DocumentClassLike    = doc.DocumentClassLike
	DraftClassLike       = doc.DraftClassLike
	SealClassLike        = doc.SealClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	CredentialLike  = doc.CredentialLike
	DocumentLike    = doc.DocumentLike
	DraftLike       = doc.DraftLike
	SealLike        = doc.SealLike
)

type (
	Parameterized = doc.Parameterized
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
	SsmP1ClassLike = age.SsmP1ClassLike
)

type (
	SsmP1Like = age.SsmP1Like
)

type (
	TsmP1ClassLike = age.TsmP1ClassLike
)

type (
	TsmP1Like = age.TsmP1Like
)

type (
	HsmP1ClassLike = age.HsmP1ClassLike
)

type (
	HsmP1Like = age.HsmP1Like
)

// CLASS ACCESSORS

// Documents

func CertificateClass() CertificateClassLike {
	return doc.CertificateClass()
}

func CitationClass() CitationClassLike {
	return doc.CitationClass()
}

func CredentialClass() CredentialClassLike {
	return doc.CredentialClass()
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func DraftClass() DraftClassLike {
	return doc.DraftClass()
}

func SealClass() SealClassLike {
	return doc.SealClass()
}

// Agents

func DigitalNotaryClass() DigitalNotaryClassLike {
	return age.DigitalNotaryClass()
}

func DigitalNotary(
	directory string,
	ssm age.Trusted,
	hsm age.Hardened,
) DigitalNotaryLike {
	return DigitalNotaryClass().DigitalNotary(
		directory,
		ssm,
		hsm,
	)
}

func SsmP1Class() SsmP1ClassLike {
	return age.SsmP1Class()
}

func SsmP1() SsmP1Like {
	return SsmP1Class().SsmP1()
}

func TsmP1Class() TsmP1ClassLike {
	return age.TsmP1Class()
}

func TsmP1(
	directory string,
) TsmP1Like {
	return TsmP1Class().TsmP1(
		directory,
	)
}

func HsmP1Class() HsmP1ClassLike {
	return age.HsmP1Class()
}

func HsmP1(
	device string,
) HsmP1Like {
	return HsmP1Class().HsmP1(
		device,
	)
}

// GLOBAL FUNCTIONS

// Documents

func Certificate(
	value ...any,
) CertificateLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.CertificateClass().CertificateFromSource(source)
	}
	var tag = value[0].(bal.TagLike)
	var version = value[1].(bal.VersionLike)
	var algorithm = value[2].(bal.QuoteLike)
	var key = value[3].(bal.BinaryLike)
	var previous = value[4].(bal.ResourceLike)
	return CertificateClass().Certificate(tag, version, algorithm, key, previous)
}

func Citation(
	value ...any,
) CitationLike {
	if len(value) == 1 {
		switch actual := value[0].(type) {
		case string:
			return doc.CitationClass().CitationFromSource(actual)
		case bal.ResourceLike:
			return doc.CitationClass().CitationFromResource(actual)
		default:
			var message = fmt.Sprintf(
				"An invalid argument type was passed into the Citation contructor: %v(%T)",
				actual,
				actual,
			)
			panic(message)
		}
	}
	var tag = value[0].(bal.TagLike)
	var version = value[1].(bal.VersionLike)
	var algorithm = value[2].(bal.QuoteLike)
	var digest = value[3].(bal.BinaryLike)
	return CitationClass().Citation(tag, version, algorithm, digest)
}

func Credential(
	value ...any,
) CredentialLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.CredentialClass().CredentialFromSource(source)
	}
	var context = value[0]
	var tag = value[1].(bal.TagLike)
	var version = value[2].(bal.VersionLike)
	var previous = value[3].(bal.ResourceLike)
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
	case doc.Parameterized:
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

func Draft(
	value ...any,
) DraftLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.DraftClass().DraftFromSource(source)
	}
	var entity = value[0]
	var type_ = value[1].(bal.NameLike)
	var tag = value[2].(bal.TagLike)
	var version = value[3].(bal.VersionLike)
	var permissions = value[4].(bal.NameLike)
	var optionalPrevious bal.ResourceLike
	if uti.IsDefined(value[5]) {
		optionalPrevious = value[5].(bal.ResourceLike)
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

func Seal(
	value ...any,
) SealLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.SealClass().SealFromSource(source)
	}
	var algorithm = value[0].(bal.QuoteLike)
	var signature = value[1].(bal.BinaryLike)
	return SealClass().Seal(algorithm, signature)
}

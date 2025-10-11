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
	doc "github.com/bali-nebula/go-digital-notary/v3/documents"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssm"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// TYPE ALIASES

// Documents

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContentClassLike     = doc.ContentClassLike
	CredentialClassLike  = doc.CredentialClassLike
	DocumentClassLike    = doc.DocumentClassLike
	SealClassLike        = doc.SealClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContentLike     = doc.ContentLike
	CredentialLike  = doc.CredentialLike
	DocumentLike    = doc.DocumentLike
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

func CitationClass() CitationClassLike {
	return doc.CitationClass()
}

func ContentClass() ContentClassLike {
	return doc.ContentClass()
}

func CredentialClass() CredentialClassLike {
	return doc.CredentialClass()
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func SealClass() SealClassLike {
	return doc.SealClass()
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

// Documents

func Certificate(
	value ...any,
) CertificateLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.CertificateClass().CertificateFromString(source)
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
			return doc.CitationClass().CitationFromString(actual)
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

func Content(
	value ...any,
) ContentLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.ContentClass().ContentFromString(source)
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
		var source = value[0].(string)
		return doc.CredentialClass().CredentialFromString(source)
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
		return DocumentClass().DocumentFromString(actual)
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

func Seal(
	value ...any,
) SealLike {
	if len(value) == 1 {
		var source = value[0].(string)
		return doc.SealClass().SealFromString(source)
	}
	var algorithm = value[0].(bal.QuoteLike)
	var signature = value[1].(bal.BinaryLike)
	return SealClass().Seal(algorithm, signature)
}

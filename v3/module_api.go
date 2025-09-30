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

func Certificate(
	tag bal.TagLike,
	version bal.VersionLike,
	algorithm bal.QuoteLike,
	key bal.BinaryLike,
) CertificateLike {
	return CertificateClass().Certificate(
		tag,
		version,
		algorithm,
		key,
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
	tag bal.TagLike,
	version bal.VersionLike,
	algorithm bal.QuoteLike,
	digest bal.BinaryLike,
) CitationLike {
	return CitationClass().Citation(
		tag,
		version,
		algorithm,
		digest,
	)
}

func CitationFromResource(
	resource bal.ResourceLike,
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

func ContentClass() ContentClassLike {
	return doc.ContentClass()
}

func Content(
	entity any,
	type_ bal.ResourceLike,
	tag bal.TagLike,
	version bal.VersionLike,
	optionalPrevious bal.ResourceLike,
	permissions bal.ResourceLike,
) ContentLike {
	return ContentClass().Content(
		entity,
		type_,
		tag,
		version,
		optionalPrevious,
		permissions,
	)
}

func ContentFromString(
	source string,
) ContentLike {
	return ContentClass().ContentFromString(
		source,
	)
}

func CredentialClass() CredentialClassLike {
	return doc.CredentialClass()
}

func Credential(
	context any,
	tag bal.TagLike,
	version bal.VersionLike,
) CredentialLike {
	return CredentialClass().Credential(
		context,
		tag,
		version,
	)
}

func CredentialFromString(
	source string,
) CredentialLike {
	return CredentialClass().CredentialFromString(
		source,
	)
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func Document(
	content doc.Parameterized,
) DocumentLike {
	return DocumentClass().Document(
		content,
	)
}

func DocumentFromString(
	source string,
) DocumentLike {
	return DocumentClass().DocumentFromString(
		source,
	)
}

func SealClass() SealClassLike {
	return doc.SealClass()
}

func Seal(
	algorithm bal.QuoteLike,
	signature bal.BinaryLike,
) SealLike {
	return SealClass().Seal(
		algorithm,
		signature,
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

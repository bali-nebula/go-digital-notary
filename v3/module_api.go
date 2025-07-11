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
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssm"
	bal "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE ALIASES

// Document

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContractClassLike    = doc.ContractClassLike
	DigestClassLike      = doc.DigestClassLike
	DraftClassLike       = doc.DraftClassLike
	SignatureClassLike   = doc.SignatureClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContractLike    = doc.ContractLike
	DigestLike      = doc.DigestLike
	DraftLike       = doc.DraftLike
	SignatureLike   = doc.SignatureLike
)

type (
	Parameterized = doc.Parameterized
)

// Notary

type (
	NotaryClassLike = not.NotaryClassLike
)

type (
	NotaryLike = not.NotaryLike
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

// Document

func CertificateClass() CertificateClassLike {
	return doc.CertificateClass()
}

func Certificate(
	algorithm fra.QuoteLike,
	publicKey fra.BinaryLike,
	tag fra.TagLike,
	version fra.VersionLike,
	optionalPrevious doc.CitationLike,
) CertificateLike {
	return CertificateClass().Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		optionalPrevious,
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
	draft doc.DraftLike,
	account fra.TagLike,
	certificate doc.CitationLike,
) ContractLike {
	return ContractClass().Contract(
		draft,
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
	component bal.ComponentLike,
	type_ fra.ResourceLike,
	tag fra.TagLike,
	version fra.VersionLike,
	permissions fra.ResourceLike,
	previous doc.CitationLike,
) DraftLike {
	return DraftClass().Draft(
		component,
		type_,
		tag,
		version,
		permissions,
		previous,
	)
}

func DraftFromString(
	source string,
) DraftLike {
	return DraftClass().DraftFromString(
		source,
	)
}

func SignatureClass() SignatureClassLike {
	return doc.SignatureClass()
}

func Signature(
	algorithm fra.QuoteLike,
	base64 fra.BinaryLike,
) SignatureLike {
	return SignatureClass().Signature(
		algorithm,
		base64,
	)
}

func SignatureFromString(
	source string,
) SignatureLike {
	return SignatureClass().SignatureFromString(
		source,
	)
}

// Notary

func NotaryClass() NotaryClassLike {
	return not.NotaryClass()
}

func Notary(
	ssm not.Trusted,
	hsm not.Hardened,
) NotaryLike {
	return NotaryClass().Notary(
		ssm,
		hsm,
	)
}

// Ssm

func SsmClass() SsmClassLike {
	return ssm.SsmClass()
}

func Ssm() SsmLike {
	return SsmClass().Ssm()
}

// GLOBAL FUNCTIONS

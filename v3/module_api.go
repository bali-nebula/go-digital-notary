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
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
	not "github.com/bali-nebula/go-digital-notary/v3/notary"
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssm"
)

// TYPE ALIASES

// Document

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContractClassLike    = doc.ContractClassLike
	DigestClassLike      = doc.DigestClassLike
	DocumentClassLike    = doc.DocumentClassLike
	SignatureClassLike   = doc.SignatureClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContractLike    = doc.ContractLike
	DigestLike      = doc.DigestLike
	DocumentLike    = doc.DocumentLike
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
	algorithm string,
	publicKey string,
	tag string,
	version string,
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
	tag string,
	version string,
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
	document doc.DocumentLike,
	account string,
	certificate doc.CitationLike,
) ContractLike {
	return ContractClass().Contract(
		document,
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
	algorithm string,
	base64 string,
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

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func Document(
	component bal.ComponentLike,
	type_ string,
	tag string,
	version string,
	permissions string,
	previous doc.CitationLike,
) DocumentLike {
	return DocumentClass().Document(
		component,
		type_,
		tag,
		version,
		permissions,
		previous,
	)
}

func DocumentFromString(
	source string,
) DocumentLike {
	return DocumentClass().DocumentFromString(
		source,
	)
}

func SignatureClass() SignatureClassLike {
	return doc.SignatureClass()
}

func Signature(
	algorithm string,
	base64 string,
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

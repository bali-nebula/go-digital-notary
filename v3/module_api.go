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
│             This "module_api.go" file was automatically generated.           │
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
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssmv1"
)

// TYPE ALIASES

// Document

type (
	CertificateClassLike = doc.CertificateClassLike
	CitationClassLike    = doc.CitationClassLike
	ContractClassLike    = doc.ContractClassLike
	DocumentClassLike    = doc.DocumentClassLike
)

type (
	CertificateLike = doc.CertificateLike
	CitationLike    = doc.CitationLike
	ContractLike    = doc.ContractLike
	DocumentLike    = doc.DocumentLike
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

// Ssmv1

type (
	SsmV1ClassLike = ssm.SsmV1ClassLike
)

type (
	SsmV1Like = ssm.SsmV1Like
)

type (
	V1Secure = ssm.V1Secure
)

// CLASS ACCESSORS

// Document

func CertificateClass() CertificateClassLike {
	return doc.CertificateClass()
}

func Certificate(
	digest string,
	signature string,
	key string,
	tag string,
	version string,
	optionalPrevious doc.CitationLike,
) CertificateLike {
	return CertificateClass().Certificate(
		digest,
		signature,
		key,
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
	protocol string,
	digest string,
) CitationLike {
	return CitationClass().Citation(
		tag,
		version,
		protocol,
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
	protocol string,
	certificate doc.CitationLike,
) ContractLike {
	return ContractClass().Contract(
		document,
		account,
		protocol,
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

// Notary

func NotaryClass() NotaryClassLike {
	return not.NotaryClass()
}

func Notary(
	optionalDirectory string,
	hsm ssm.V1Secure,
) NotaryLike {
	return NotaryClass().Notary(
		optionalDirectory,
		hsm,
	)
}

// Ssmv1

func SsmV1Class() SsmV1ClassLike {
	return ssm.SsmV1Class()
}

func SsmV1(
	directory string,
) SsmV1Like {
	return SsmV1Class().SsmV1(
		directory,
	)
}

// GLOBAL FUNCTIONS

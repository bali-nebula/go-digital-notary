/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package documents

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	ab2 "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CERTIFICATE INTERFACE

// This constructor creates a new certificate.
func Certificate(
	key abs.BinaryLike,
	algorithms abs.CatalogLike,
	tag abs.TagLike,
	version abs.VersionLike,
	previous ab2.CitationLike,
) ab2.CertificateLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(ab2.KeyAttribute, bal.Component(key))
	attributes.SetValue(ab2.AlgorithmsAttribute, bal.Component(algorithms))

	// Create a new record.
	var type_ = Type(bal.Moniker("/bali/types/documents/Certificate/v1"), nil)
	var permissions = bal.Moniker("/bali/permissions/public/v1")
	var record = Record(attributes, type_, tag, version, permissions, previous)

	// Create a new certificate.
	return &certificate{record}
}

// CERTIFICATE IMPLEMENTATION

type certificate struct {
	ab2.RecordLike
}

// PUBLISHED INTERFACE

func (v *certificate) GetAlgorithms() abs.CatalogLike {
	return v.ExtractCatalog().GetValue(ab2.AlgorithmsAttribute).ExtractCatalog()
}

func (v *certificate) GetKey() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.KeyAttribute).ExtractBinary()
}

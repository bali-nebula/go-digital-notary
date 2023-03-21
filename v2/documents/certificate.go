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
	gcf "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CERTIFICATE INTERFACE

// This constructor creates a new certificate.
func Certificate(
	key gcf.BinaryLike,
	algorithms gcf.CatalogLike,
	tag gcf.TagLike,
	version gcf.VersionLike,
	previous abs.CitationLike,
) abs.CertificateLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(abs.KeyAttribute, bal.Component(key))
	attributes.SetValue(abs.AlgorithmsAttribute, bal.Component(algorithms))

	// Create a new record.
	var type_ = Type(bal.Moniker("/bali/types/documents/Certificate/v1"), nil)
	var permissions = bal.Moniker("/bali/permissions/public/v1")
	var record = Record(attributes, type_, tag, version, permissions, previous)

	// Create a new certificate.
	return &certificate{record}
}

// CERTIFICATE IMPLEMENTATION

type certificate struct {
	abs.RecordLike
}

// PUBLISHED INTERFACE

func (v *certificate) GetAlgorithms() gcf.CatalogLike {
	return v.ExtractCatalog().GetValue(abs.AlgorithmsAttribute).ExtractCatalog()
}

func (v *certificate) GetKey() gcf.BinaryLike {
	return v.ExtractCatalog().GetValue(abs.KeyAttribute).ExtractBinary()
}

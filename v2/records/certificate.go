/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package records

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
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

	// Create a new context.
	var context = com.Context()
	context.SetValue(ab2.TypeAttribute, bal.ParseComponent("/bali/types/documents/Certificate/v1"))
	context.SetValue(ab2.TagAttribute, bal.Component(tag))
	context.SetValue(ab2.VersionAttribute, bal.Component(version))
	context.SetValue(ab2.PermissionsAttribute, bal.ParseComponent("/bali/permissions/public/v1"))
	if previous != nil {
		context.SetValue(ab2.PreviousAttribute, bal.Component(previous))
	}

	// Create a new certificate.
	return &certificate{bal.ComponentWithContext(attributes, context)}
}

// CERTIFICATE IMPLEMENTATION

type certificate struct {
	abs.Encapsulated
}

func (v *certificate) GetAlgorithms() abs.CatalogLike {
	return v.ExtractCatalog().GetValue(ab2.AlgorithmsAttribute).ExtractCatalog()
}

func (v *certificate) GetKey() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.KeyAttribute).ExtractBinary()
}

func (v *certificate) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(ab2.PermissionsAttribute).ExtractMoniker()
}

func (v *certificate) GetPrevious() ab2.CitationLike {
	return v.GetContext().GetValue(ab2.PreviousAttribute).ExtractCatalog().(ab2.CitationLike)
}

func (v *certificate) GetTag() abs.TagLike {
	return v.GetContext().GetValue(ab2.TagAttribute).ExtractTag()
}

func (v *certificate) GetType() ab2.TypeLike {
	return v.GetContext().GetValue(ab2.TypeAttribute).(ab2.TypeLike)
}

func (v *certificate) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(ab2.VersionAttribute).ExtractVersion()
}

/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package notary

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
)

// CERTIFICATE INTERFACE

// This constructor creates a new certificate.
func Certificate(
	key abs.BinaryLike,
	algorithms abs.CatalogLike,
	tag abs.TagLike,
	version abs.VersionLike,
	previous CitationLike,
) CertificateLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(keyAttribute, bal.Component(key))
	attributes.SetValue(algorithmsAttribute, bal.Component(algorithms))

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/bali/types/documents/Certificate/v1"))
	context.SetValue(tagAttribute, bal.Component(tag))
	context.SetValue(versionAttribute, bal.Component(version))
	context.SetValue(permissionsAttribute, bal.ParseComponent("/bali/permissions/public/v1"))
	if previous != nil {
		context.SetValue(previousAttribute, bal.Component(previous))
	}

	// Create a new certificate.
	return &certificate{bal.ComponentWithContext(attributes, context)}
}

// CERTIFICATE IMPLEMENTATION

type certificate struct {
	abs.Encapsulated
}

func (v *certificate) GetAlgorithms() abs.CatalogLike {
	return v.ExtractCatalog().GetValue(algorithmsAttribute).ExtractCatalog()
}

func (v *certificate) GetKey() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(keyAttribute).ExtractBinary()
}

func (v *certificate) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *certificate) GetPrevious() CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(CitationLike)
}

func (v *certificate) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *certificate) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

func (v *certificate) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}

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
	com "github.com/bali-nebula/go-component-framework/v2/components"
	ab2 "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// DOCUMENT INTERFACE

// This constructor creates a new document.
func Document(
	attributes abs.CatalogLike,
	type_ ab2.TypeLike,
	tag abs.TagLike,
	version abs.VersionLike,
	permissions abs.MonikerLike,
	previous ab2.CitationLike,
) ab2.DocumentLike {

	// Create a new context.
	var context = com.Context()
	context.SetValue(ab2.TypeAttribute, type_)
	context.SetValue(ab2.TagAttribute, bal.Component(tag))
	context.SetValue(ab2.VersionAttribute, bal.Component(version))
	context.SetValue(ab2.PermissionsAttribute, bal.Component(permissions))
	if previous != nil {
		context.SetValue(ab2.PreviousAttribute, bal.Component(previous))
	}

	// Create a new document.
	return &document{bal.ComponentWithContext(attributes, context)}
}

// DOCUMENT IMPLEMENTATION

type document struct {
	abs.Encapsulated
}

func (v *document) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(ab2.PermissionsAttribute).ExtractMoniker()
}

func (v *document) GetPrevious() ab2.CitationLike {
	return v.GetContext().GetValue(ab2.PreviousAttribute).ExtractCatalog().(ab2.CitationLike)
}

func (v *document) GetTag() abs.TagLike {
	return v.GetContext().GetValue(ab2.TagAttribute).ExtractTag()
}

func (v *document) GetType() ab2.TypeLike {
	return v.GetContext().GetValue(ab2.TypeAttribute).(ab2.TypeLike)
}

func (v *document) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(ab2.VersionAttribute).ExtractVersion()
}

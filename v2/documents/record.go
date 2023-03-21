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
	com "github.com/bali-nebula/go-component-framework/v2/components"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// RECORD INTERFACE

// This constructor creates a new record.
func Record(
	attributes gcf.CatalogLike,
	type_ abs.TypeLike,
	tag gcf.TagLike,
	version gcf.VersionLike,
	permissions gcf.MonikerLike,
	previous abs.CitationLike,
) abs.RecordLike {

	// Create a new context.
	var context = com.Context()
	context.SetValue(abs.TypeAttribute, type_)
	context.SetValue(abs.TagAttribute, bal.Component(tag))
	context.SetValue(abs.VersionAttribute, bal.Component(version))
	context.SetValue(abs.PermissionsAttribute, bal.Component(permissions))
	if previous != nil {
		context.SetValue(abs.PreviousAttribute, bal.Component(previous))
	}

	// Create a new record.
	return &record{bal.ComponentWithContext(attributes, context)}
}

// RECORD IMPLEMENTATION

type record struct {
	gcf.Encapsulated
}

// RESTRICTED INTERFACE

func (v *record) GetPermissions() gcf.MonikerLike {
	return v.GetContext().GetValue(abs.PermissionsAttribute).ExtractMoniker()
}

// VERSIONED INTERFACE

func (v *record) GetPrevious() abs.CitationLike {
	return v.GetContext().GetValue(abs.PreviousAttribute).ExtractCatalog().(abs.CitationLike)
}

func (v *record) GetTag() gcf.TagLike {
	return v.GetContext().GetValue(abs.TagAttribute).ExtractTag()
}

func (v *record) GetVersion() gcf.VersionLike {
	return v.GetContext().GetValue(abs.VersionAttribute).ExtractVersion()
}

// TYPED INTERFACE

func (v *record) GetType() abs.TypeLike {
	return v.GetContext().GetValue(abs.TypeAttribute).(abs.TypeLike)
}

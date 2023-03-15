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
	com "github.com/bali-nebula/go-component-framework/v2/components"
)

// DOCUMENT INTERFACE

// This constructor creates a new document.
func Document(
	attributes abs.CatalogLike,
	type_ abs.ComponentLike,
	tag abs.TagLike,
	version abs.VersionLike,
	permissions abs.MonikerLike,
	previous CitationLike,
) DocumentLike {

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, type_)
	context.SetValue(tagAttribute, bal.Component(tag))
	context.SetValue(versionAttribute, bal.Component(version))
	context.SetValue(permissionsAttribute, bal.Component(permissions))
	if previous != nil {
		context.SetValue(previousAttribute, bal.Component(previous))
	}

	// Create a new document.
	return &document{bal.ComponentWithContext(attributes, context)}
}

// DOCUMENT IMPLEMENTATION

type document struct {
	abs.Encapsulated
}

func (v *document) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *document) GetPrevious() CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(CitationLike)
}

func (v *document) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *document) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

func (v *document) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}

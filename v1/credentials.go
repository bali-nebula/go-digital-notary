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
	abs "github.com/bali-nebula/go-component-framework/v1/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v1/bali"
	col "github.com/bali-nebula/go-component-framework/v1/collections"
	com "github.com/bali-nebula/go-component-framework/v1/components"
)

// CREDENTIALS INTERFACE

// This constructor creates a new credentials.
func Credentials(
	salt abs.BinaryLike,
) abs.CredentialsLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(saltAttribute, bal.Component(salt))

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/bali/types/documents/Credentials/v1"))
	context.SetValue(tagAttribute, bal.Component(bal.NewTag()))
	context.SetValue(versionAttribute, bal.Component(v1))
	context.SetValue(permissionsAttribute, bal.ParseComponent("/bali/permissions/private/v1"))

	// Create a new credentials.
	return &credentials{bal.ComponentWithContext(attributes, context)}
}

// CREDENTIALS IMPLEMENTATION

type credentials struct {
	abs.Encapsulated
}

func (v *credentials) GetSalt() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(saltAttribute).ExtractBinary()
}

func (v *credentials) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *credentials) GetPrevious() abs.CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(abs.CitationLike)
}

func (v *credentials) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *credentials) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

func (v *credentials) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}

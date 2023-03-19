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

// CREDENTIALS INTERFACE

// This constructor creates a new credentials.
func Credentials(
	salt abs.BinaryLike,
) ab2.CredentialsLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(ab2.SaltAttribute, bal.Component(salt))

	// Create a new context.
	var context = com.Context()
	context.SetValue(ab2.TypeAttribute, bal.ParseComponent("/bali/types/documents/Credentials/v1"))
	context.SetValue(ab2.TagAttribute, bal.Component(bal.NewTag()))
	context.SetValue(ab2.VersionAttribute, bal.Component("v1"))
	context.SetValue(ab2.PermissionsAttribute, bal.ParseComponent("/bali/permissions/private/v1"))

	// Create a new credentials.
	return &credentials{bal.ComponentWithContext(attributes, context)}
}

// CREDENTIALS IMPLEMENTATION

type credentials struct {
	abs.Encapsulated
}

// SEASONED INTERFACE

func (v *credentials) GetSalt() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.SaltAttribute).ExtractBinary()
}

// TYPED INTERFACE

func (v *credentials) GetType() ab2.TypeLike {
	return v.GetContext().GetValue(ab2.TypeAttribute).(ab2.TypeLike)
}

// RESTRICTED INTERFACE

func (v *credentials) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(ab2.PermissionsAttribute).ExtractMoniker()
}

// VERSIONED INTERFACE

func (v *credentials) GetTag() abs.TagLike {
	return v.GetContext().GetValue(ab2.TagAttribute).ExtractTag()
}

func (v *credentials) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(ab2.VersionAttribute).ExtractVersion()
}

func (v *credentials) GetPrevious() ab2.CitationLike {
	return v.GetContext().GetValue(ab2.PreviousAttribute).ExtractCatalog().(ab2.CitationLike)
}

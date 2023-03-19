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

// CREDENTIALS INTERFACE

// This constructor creates a new credential.
func Credential(
	salt abs.BinaryLike,
) ab2.CredentialLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(ab2.SaltAttribute, bal.Component(salt))

	// Create a new record.
	var type_ = Type(bal.Moniker("/bali/types/documents/Credential/v1"), nil)
	var tag = bal.NewTag()
	var version = bal.Version("v1")
	var permissions = bal.Moniker("/bali/permissions/private/v1")
	var previous ab2.CitationLike
	var record = Record(attributes, type_, tag, version, permissions, previous)

	// Create a new credential.
	return &credential{record}
}

// CREDENTIALS IMPLEMENTATION

type credential struct {
	ab2.RecordLike
}

// SEASONED INTERFACE

func (v *credential) GetSalt() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.SaltAttribute).ExtractBinary()
}

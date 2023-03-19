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

	// Create a new document.
	var type_ = Type(bal.Moniker("/bali/types/documents/Credentials/v1"), nil)
	var tag = bal.NewTag()
	var version = bal.Version("v1")
	var permissions = bal.Moniker("/bali/permissions/private/v1")
	var previous ab2.CitationLike
	var document = Document(attributes, type_, tag, version, permissions, previous)

	// Create a new credentials.
	return &credentials{document}
}

// CREDENTIALS IMPLEMENTATION

type credentials struct {
	ab2.DocumentLike
}

// SEASONED INTERFACE

func (v *credentials) GetSalt() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.SaltAttribute).ExtractBinary()
}

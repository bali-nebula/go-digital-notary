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

// CITATION INTERFACE

// This constructor creates a new citation.
func Citation(
	tag abs.TagLike,
	version abs.VersionLike,
	protocol abs.VersionLike,
	digest abs.BinaryLike,
) ab2.CitationLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(ab2.TagAttribute, bal.Component(tag))
	attributes.SetValue(ab2.VersionAttribute, bal.Component(version))
	attributes.SetValue(ab2.ProtocolAttribute, bal.Component(protocol))
	attributes.SetValue(ab2.DigestAttribute, bal.Component(digest))

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(ab2.TypeAttribute, bal.ParseComponent("/bali/types/documents/Citation/v1"))

	// Create a new citation.
	return &citation{bal.ComponentWithContext(attributes, context)}
}

// CITATION IMPLEMENTATION

type citation struct {
	abs.Encapsulated
}

func (v *citation) GetDigest() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(ab2.DigestAttribute).ExtractBinary()
}

func (v *citation) GetProtocol() abs.VersionLike {
	return v.ExtractCatalog().GetValue(ab2.ProtocolAttribute).ExtractVersion()
}

func (v *citation) GetTag() abs.TagLike {
	return v.ExtractCatalog().GetValue(ab2.TagAttribute).ExtractTag()
}

func (v *citation) GetType() ab2.TypeLike {
	return v.GetContext().GetValue(ab2.TypeAttribute).(ab2.TypeLike)
}

func (v *citation) GetVersion() abs.VersionLike {
	return v.ExtractCatalog().GetValue(ab2.VersionAttribute).ExtractVersion()
}

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
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CITATION INTERFACE

// This constructor creates a new citation.
func Citation(
	tag gcf.TagLike,
	version gcf.VersionLike,
	protocol gcf.VersionLike,
	digest gcf.BinaryLike,
) abs.CitationLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(abs.TagAttribute, bal.Component(tag))
	attributes.SetValue(abs.VersionAttribute, bal.Component(version))
	attributes.SetValue(abs.ProtocolAttribute, bal.Component(protocol))
	attributes.SetValue(abs.DigestAttribute, bal.Component(digest))

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(abs.TypeAttribute, bal.ParseComponent("/bali/types/documents/Citation/v1"))

	// Create a new citation.
	return &citation{bal.ComponentWithContext(attributes, context)}
}

// CITATION IMPLEMENTATION

type citation struct {
	gcf.Encapsulated
}

// REFERENTIAL INTERFACE

func (v *citation) GetDigest() gcf.BinaryLike {
	return v.ExtractCatalog().GetValue(abs.DigestAttribute).ExtractBinary()
}

func (v *citation) GetProtocol() gcf.VersionLike {
	return v.ExtractCatalog().GetValue(abs.ProtocolAttribute).ExtractVersion()
}

func (v *citation) GetTag() gcf.TagLike {
	return v.ExtractCatalog().GetValue(abs.TagAttribute).ExtractTag()
}

func (v *citation) GetVersion() gcf.VersionLike {
	return v.ExtractCatalog().GetValue(abs.VersionAttribute).ExtractVersion()
}

// TYPED INTERFACE

func (v *citation) GetType() abs.TypeLike {
	return v.GetContext().GetValue(abs.TypeAttribute).(abs.TypeLike)
}

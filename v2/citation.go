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

// CITATION INTERFACE

// This constructor creates a new citation.
func Citation(
	tag abs.TagLike,
	version abs.VersionLike,
	protocol abs.VersionLike,
	digest abs.BinaryLike,
) CitationLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(tagAttribute, bal.Component(tag))
	attributes.SetValue(versionAttribute, bal.Component(version))
	attributes.SetValue(protocolAttribute, bal.Component(protocol))
	attributes.SetValue(digestAttribute, bal.Component(digest))

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/bali/types/documents/Citation/v1"))

	// Create a new citation.
	return &citation{bal.ComponentWithContext(attributes, context)}
}

// CITATION IMPLEMENTATION

type citation struct {
	abs.Encapsulated
}

func (v *citation) GetDigest() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(digestAttribute).ExtractBinary()
}

func (v *citation) GetProtocol() abs.VersionLike {
	return v.ExtractCatalog().GetValue(protocolAttribute).ExtractVersion()
}

func (v *citation) GetTag() abs.TagLike {
	return v.ExtractCatalog().GetValue(tagAttribute).ExtractTag()
}

func (v *citation) GetType() TypeLike {
	return v.GetContext().GetValue(typeAttribute).(TypeLike)
}

func (v *citation) GetVersion() abs.VersionLike {
	return v.ExtractCatalog().GetValue(versionAttribute).ExtractVersion()
}

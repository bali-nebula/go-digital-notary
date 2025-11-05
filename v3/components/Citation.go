/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package components

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-essential-utilities/v8"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func CitationClass() CitationClassLike {
	return citationClass()
}

// Constructor Methods

func (c *citationClass_) Citation(
	tag doc.TagLike,
	version doc.VersionLike,
	algorithm doc.QuoteLike,
	digest doc.BinaryLike,
) CitationLike {
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}

	var source = `[
    $tag: ` + tag.AsSource() + `
    $version: ` + version.AsSource() + `
    $algorithm: ` + algorithm.AsSource() + `
    $digest: ` + digest.AsSource() + `
]($type: /bali/types/notary/Citation/v3)`
	return c.CitationFromSource(source)
}

func (c *citationClass_) CitationFromSource(
	source string,
) CitationLike {
	var component = doc.ParseComponent(source)
	var instance = &citation_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Compound: component,
	}
	return instance
}

func (c *citationClass_) CitationFromResource(
	resource doc.ResourceLike,
) CitationLike {
	// Parse parts of the path.
	var path = resource.GetPath()
	var parts = sts.Split(path, "/")
	parts = sts.Split(parts[1], ":")
	var tag = doc.Tag("#" + parts[0])
	var version = doc.Version(parts[1])

	// Parse parts of the query string.
	var query = resource.GetQuery()
	parts = sts.Split(query, "=")
	var algorithm = parts[0]
	var digest = parts[1]
	digest = sts.ReplaceAll(digest, "-", "+")
	digest = sts.ReplaceAll(digest, "_", "/")
	digest = "'>\n    " + digest[:60] + "\n    " + digest[60:] + "\n<'"

	// Construct the citation.
	var instance = c.Citation(
		doc.Tag(tag),
		doc.Version(version),
		doc.Quote(algorithm),
		doc.Binary(digest),
	)

	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *citation_) GetClass() CitationClassLike {
	return citationClass()
}

func (v *citation_) AsIntrinsic() doc.Compound {
	return v.Compound
}

func (v *citation_) AsSource() string {
	return doc.FormatComponent(v.Compound) + "\n"
}

func (v *citation_) AsResource() doc.ResourceLike {
	var algorithm = v.GetAlgorithm().AsSource()
	algorithm = algorithm[1 : len(algorithm)-1] // Remove the double quotes.
	var digest = v.GetDigest().AsSource()
	digest = digest[2 : len(digest)-2]
	digest = sts.ReplaceAll(digest, " ", "")
	digest = sts.ReplaceAll(digest, "\n", "")
	digest = sts.ReplaceAll(digest, "+", "-")
	digest = sts.ReplaceAll(digest, "/", "_")
	var tag = v.GetTag().AsSource()[1:]
	var version = v.GetVersion().AsSource()
	var source = "<nebula:/" + tag + ":" + version + "?" +
		algorithm + "=" + digest + ">"
	var resource = doc.Resource(source)
	return resource
}

// Attribute Methods

func (v *citation_) GetTag() doc.TagLike {
	var component = v.GetSubcomponent(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *citation_) GetVersion() doc.VersionLike {
	var component = v.GetSubcomponent(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *citation_) GetAlgorithm() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$algorithm"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *citation_) GetDigest() doc.BinaryLike {
	var composite = v.GetSubcomponent(doc.Symbol("$digest"))
	var source = doc.FormatComponent(composite)
	return doc.Binary(source)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type citation_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Compound
}

// Class Structure

type citationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func citationClass() *citationClass_ {
	return citationClassReference_
}

var citationClassReference_ = &citationClass_{
	// Initialize the class constants.
}

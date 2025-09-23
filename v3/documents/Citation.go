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

package documents

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func CitationClass() CitationClassLike {
	return citationClass()
}

// Constructor Methods

func (c *citationClass_) Citation(
	algorithm doc.QuoteLike,
	digest doc.BinaryLike,
	tag doc.TagLike,
	version doc.VersionLike,
) CitationLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var previous = "none"
	var current = version.AsIntrinsic()[0]
	if current > 1 {
		previous = "<nebula:/" + tag.AsString()[1:] +
			":" + doc.Version([]uint{current - 1}).AsString() + ">"
	}
	var component = doc.ParseSource(`[
    $algorithm: ` + algorithm.AsString() + `
    $digest: ` + digest.AsString() + `
](
    $type: <bali:/types/notary/Citation:v3>
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
	$permissions: <bali:/permissions/Public:v3>
    $previous: ` + previous + `
)`,
	)

	var instance = &citation_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

func (c *citationClass_) CitationFromString(
	source string,
) CitationLike {
	var component = doc.ParseSource(source)
	var instance = &citation_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Declarative: component,
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
	digest = "'>\n        " + digest[:60] + "\n        " + digest[60:] + "\n<'"

	// Construct the citation.
	var instance = c.Citation(
		doc.Quote(algorithm),
		doc.Binary(digest),
		doc.Tag(tag),
		doc.Version(version),
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

func (v *citation_) AsIntrinsic() doc.ComponentLike {
	return v.Declarative.(doc.ComponentLike)
}

func (v *citation_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *citation_) AsResource() doc.ResourceLike {
	var algorithm = v.GetAlgorithm().AsString()
	algorithm = algorithm[1 : len(algorithm)-1] // Remove the double quotes.
	var digest = v.GetDigest().AsString()
	digest = digest[2 : len(digest)-2]
	digest = sts.ReplaceAll(digest, " ", "")
	digest = sts.ReplaceAll(digest, "\n", "")
	digest = sts.ReplaceAll(digest, "+", "-")
	digest = sts.ReplaceAll(digest, "/", "_")
	var tag = v.GetTag().AsString()[1:]
	var version = v.GetVersion().AsString()
	var resource = doc.Resource(
		"<nebula:/" + tag + ":" + version + "?" + algorithm + "=" + digest + ">",
	)
	return resource
}

// Attribute Methods

func (v *citation_) GetAlgorithm() doc.QuoteLike {
	var object = v.GetObject(doc.Symbol("algorithm"))
	return doc.Quote(doc.FormatComponent(object))
}

func (v *citation_) GetDigest() doc.BinaryLike {
	var object = v.GetObject(doc.Symbol("digest"))
	return doc.Binary(doc.FormatComponent(object))
}

// Parameterized Methods

func (v *citation_) GetEntity() any {
	return v.Declarative.(doc.ComponentLike).GetEntity()
}

func (v *citation_) GetType() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("type"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *citation_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *citation_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *citation_) GetPermissions() doc.ResourceLike {
	var component = v.GetParameter(doc.Symbol("permissions"))
	return doc.Resource(doc.FormatComponent(component))
}

func (v *citation_) GetOptionalPrevious() doc.ResourceLike {
	var previous doc.ResourceLike
	var component = v.GetParameter(doc.Symbol("previous"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		if source != "none" {
			previous = doc.Resource(source)
		}
	}
	return previous
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type citation_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Declarative
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

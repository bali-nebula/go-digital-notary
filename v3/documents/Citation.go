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
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	fra "github.com/craterdog/go-component-framework/v7"
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
	isNotarized fra.BooleanLike,
	tag fra.TagLike,
	version fra.VersionLike,
	digest DigestLike,
) CitationLike {
	if uti.IsUndefined(isNotarized) {
		panic("The \"isNotarized\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}

	var component = doc.ParseSource(`[
    $isNotarized: ` + isNotarized.AsString() + `
    $tag: ` + tag.AsString() + `
    $version: ` + version.AsString() + `
    $digest: ` + digest.AsString() + `
]($type: <bali:/nebula/types/Citation:v3>)`,
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
	resource fra.ResourceLike,
) CitationLike {
	// Parse parts of the path.
	var path = resource.GetPath()
	var parts = sts.Split(path, "/")
	var documents = parts[2]
	var isNotarized fra.BooleanLike
	switch documents {
	case "contracts":
		isNotarized = fra.Boolean(true)
	case "drafts":
		isNotarized = fra.Boolean(false)
	default:
		var message = fmt.Sprintf(
			"The resource has an invalid document type: %v",
			documents,
		)
		panic(message)
	}
	parts = sts.Split(parts[3], ":")
	var tag = fra.TagFromString("#" + parts[0])
	var version = fra.VersionFromString(parts[1])

	// Parse parts of the query string.
	var query = resource.GetQuery()
	parts = sts.Split(query, "=")
	var algorithm = parts[0]
	var base64 = parts[1]
	base64 = sts.ReplaceAll(base64, "-", "+")
	base64 = sts.ReplaceAll(base64, "_", "/")
	base64 = "\n        " + base64[:60] + "\n        " + base64[60:] + "\n"

	// Construct the digest.
	var digest = DigestClass().DigestFromString(`[
    $algorithm: "` + algorithm + `"
    $base64: '>` + base64 + `<'
]($type: <bali:/nebula/types/Digest:v3>)`,
	)

	return c.Citation(isNotarized, tag, version, digest)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *citation_) GetClass() CitationClassLike {
	return citationClass()
}

func (v *citation_) AsString() string {
	return doc.FormatDocument(v.Declarative.(doc.ComponentLike))
}

func (v *citation_) AsResource() fra.ResourceLike {
	var tag = v.GetTag().AsString()[1:]
	var version = v.GetVersion().AsString()
	var documents string
	switch v.IsNotarized().AsString() {
	case "true":
		documents = "contracts"
	case "false":
		documents = "drafts"
	}
	var digest = v.GetDigest()
	var algorithm = digest.GetAlgorithm().AsString()
	algorithm = algorithm[1 : len(algorithm)-1] // Remove the double quotes.
	var base64 = digest.GetBase64().AsString()
	base64 = base64[2 : len(base64)-2]
	base64 = sts.ReplaceAll(base64, " ", "")
	base64 = sts.ReplaceAll(base64, "\n", "")
	base64 = sts.ReplaceAll(base64, "+", "-")
	base64 = sts.ReplaceAll(base64, "/", "_")
	var resource = fra.ResourceFromString(
		"<bali:/nebula/" + documents + "/" + tag + ":" + version +
			"?" + algorithm + "=" + base64 + ">",
	)
	return resource
}

func (v *citation_) IsNotarized() fra.BooleanLike {
	var object = v.GetObject(fra.Symbol("isNotarized"))
	return fra.BooleanFromString(doc.FormatComponent(object))
}

func (v *citation_) GetTag() fra.TagLike {
	var object = v.GetObject(fra.Symbol("tag"))
	return fra.TagFromString(doc.FormatComponent(object))
}

func (v *citation_) GetVersion() fra.VersionLike {
	var object = v.GetObject(fra.Symbol("version"))
	return fra.VersionFromString(doc.FormatComponent(object))
}

func (v *citation_) GetDigest() DigestLike {
	var object = v.GetObject(fra.Symbol("digest"))
	return DigestClass().DigestFromString(doc.FormatComponent(object))
}

// Attribute Methods

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

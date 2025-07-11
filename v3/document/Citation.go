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

package document

import (
	fmt "fmt"
	bal "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func CitationClass() CitationClassLike {
	return citationClass()
}

// Constructor Methods

func (c *citationClass_) Citation(
	tag fra.TagLike,
	version fra.VersionLike,
	digest DigestLike,
) CitationLike {
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}
	var instance = &citation_{
		// Initialize the instance attributes.
		tag_:     tag,
		version_: version,
		digest_:  digest,
	}
	return instance
}

func (c *citationClass_) CitationFromString(
	source string,
) CitationLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid citation was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var document = bal.ParseSource(source)
	var tag = fra.TagFromString(
		DraftClass().ExtractAttribute("$tag", document),
	)
	var version = fra.VersionFromString(
		DraftClass().ExtractAttribute("$version", document),
	)
	var digest = DraftClass().ExtractDigest(document)
	return c.Citation(tag, version, digest)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *citation_) GetClass() CitationClassLike {
	return citationClass()
}

func (v *citation_) AsString() string {
	var string_ = `[
`
	string_ += `    $tag: ` + v.GetTag().AsString()
	string_ += `    $version: ` + v.GetVersion().AsString()
	string_ += `    $digest: ` + v.GetDigest().AsString()
	string_ += `]($type: <bali:/types/documents/Citation:v3>)
`
	var citation = bal.ParseSource(string_)
	string_ = bal.FormatDocument(citation)
	return string_
}

// Attribute Methods

func (v *citation_) GetTag() fra.TagLike {
	return v.tag_
}

func (v *citation_) GetVersion() fra.VersionLike {
	return v.version_
}

func (v *citation_) GetDigest() DigestLike {
	return v.digest_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type citation_ struct {
	// Declare the instance attributes.
	tag_     fra.TagLike
	version_ fra.VersionLike
	digest_  DigestLike
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

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
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func CitationClass() CitationClassLike {
	return citationClass()
}

// Constructor Methods

func (c *citationClass_) Citation(
	tag string,
	version string,
	protocol string,
	digest string,
) CitationLike {
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(protocol) {
		panic("The \"protocol\" attribute is required by this class.")
	}
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}
	var instance = &citation_{
		// Initialize the instance attributes.
		tag_:      tag,
		version_:  version,
		protocol_: protocol,
		digest_:   digest,
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
	var component = doc.ParseSource(source).GetComponent()
	var collection = component.GetAny().(doc.CollectionLike)
	var attributes = collection.GetAny().(doc.AttributesLike)
	var associations = attributes.GetAssociations()

	var association = associations.GetValue(1)
	var element = association.GetPrimitive().GetAny().(doc.ElementLike)
	var symbol = element.GetAny().(string)
	if symbol != "$tag" {
		panic("Missing the $tag attribute.")
	}
	var tag = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(2)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$version" {
		panic("Missing the $version attribute.")
	}
	var version = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(3)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$protocol" {
		panic("Missing the $protocol attribute.")
	}
	var protocol = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(4)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$digest" {
		panic("Missing the $digest attribute.")
	}
	var digest = doc.FormatDocument(association.GetDocument())

	return c.Citation(tag, version, protocol, digest)
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
	string_ += `    $tag: ` + v.GetTag()
	string_ += `    $version: ` + v.GetVersion()
	string_ += `    $protocol: ` + v.GetProtocol()
	string_ += `    $digest: ` + v.GetDigest()
	string_ += `]($type: <bali:/types/documents/Citation@v1>)
`
	var citation = doc.ParseSource(string_)
	string_ = doc.FormatDocument(citation)
	return string_
}

// Attribute Methods

func (v *citation_) GetTag() string {
	return v.tag_
}

func (v *citation_) GetVersion() string {
	return v.version_
}

func (v *citation_) GetProtocol() string {
	return v.protocol_
}

func (v *citation_) GetDigest() string {
	return v.digest_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type citation_ struct {
	// Declare the instance attributes.
	tag_      string
	version_  string
	protocol_ string
	digest_   string
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

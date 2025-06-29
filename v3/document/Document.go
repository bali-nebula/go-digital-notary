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

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	component doc.ComponentLike,
	type_ string,
	tag string,
	version string,
	permissions string,
	previous CitationLike,
) DocumentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	if uti.IsUndefined(permissions) {
		panic("The \"permissions\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		component_:   component,
		type_:        type_,
		tag_:         tag,
		version_:     version,
		permissions_: permissions,
		previous_:    previous,
	}
	return instance
}

func (c *documentClass_) DocumentFromString(
	source string,
) DocumentLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid document was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var document = doc.ParseSource(source)
	var component = document.GetComponent()
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()

	var association = associations.GetValue(1)
	var element = association.GetPrimitive().GetAny().(doc.ElementLike)
	var symbol = element.GetAny().(string)
	if symbol != "$type" {
		panic("Missing the $type attribute.")
	}
	var type_ = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(2)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$tag" {
		panic("Missing the $tag attribute.")
	}
	var tag = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(3)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$version" {
		panic("Missing the $version attribute.")
	}
	var version = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(4)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$permissions" {
		panic("Missing the $permissions attribute.")
	}
	var permissions = doc.FormatDocument(association.GetDocument())

	var previous CitationLike
	if associations.GetSize() > 4 {
		association = associations.GetValue(5)
		element = association.GetPrimitive().GetAny().(doc.ElementLike)
		symbol = element.GetAny().(string)
		if symbol != "$previous" {
			panic("Missing the $previous attribute.")
		}
		previous = citationClass().CitationFromString(
			doc.FormatDocument(association.GetDocument()),
		)
	}

	return c.Document(component, type_, tag, version, permissions, previous)
}

// Constant Methods

// Function Methods

func (c *documentClass_) ExtractAttribute(
	name string,
	document doc.DocumentLike,
) string {
	var attribute string
	var component = document.GetComponent()
	var collection = component.GetAny().(doc.CollectionLike)
	var attributes = collection.GetAny().(doc.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(doc.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			attribute = doc.FormatDocument(association.GetDocument())
			attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
			break
		}
	}
	return attribute
}

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

func (v *document_) AsString() string {
	var document = doc.Document(v.GetComponent(), nil, "")
	var string_ = doc.FormatDocument(document)
	string_ = string_[:len(string_)-1] // Remove the trailing newline.
	string_ += `(
`
	string_ += `    $type: ` + v.GetType()
	string_ += `    $tag: ` + v.GetTag()
	string_ += `    $version: ` + v.GetVersion()
	string_ += `    $permissions: ` + v.GetPermissions()
	var previous = v.GetOptionalPrevious()
	if uti.IsDefined(previous) {
		string_ += `    $previous: ` + previous.AsString()
	}
	string_ += `)
`
	string_ = doc.FormatDocument(doc.ParseSource(string_))
	return string_
}

// Attribute Methods

func (v *document_) GetComponent() doc.ComponentLike {
	return v.component_
}

// Parameterized Methods

func (v *document_) GetType() string {
	return v.type_
}

func (v *document_) GetTag() string {
	return v.tag_
}

func (v *document_) GetVersion() string {
	return v.version_
}

func (v *document_) GetPermissions() string {
	return v.permissions_
}

func (v *document_) GetOptionalPrevious() CitationLike {
	return v.previous_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	component_   doc.ComponentLike
	type_        string
	tag_         string
	version_     string
	permissions_ string
	previous_    CitationLike
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentClass() *documentClass_ {
	return documentClassReference_
}

var documentClassReference_ = &documentClass_{
	// Initialize the class constants.
}

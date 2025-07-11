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

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	component bal.ComponentLike,
	type_ fra.ResourceLike,
	tag fra.TagLike,
	version fra.VersionLike,
	permissions fra.ResourceLike,
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
	var document = bal.ParseSource(source)
	var component = document.GetComponent()
	var type_ = c.ExtractType(document)
	var tag = c.ExtractTag(document)
	var version = c.ExtractVersion(document)
	var permissions = c.ExtractPermissions(document)
	var previous = c.ExtractPrevious(document)
	return c.Document(
		component,
		type_,
		tag,
		version,
		permissions,
		previous,
	)
}

// Constant Methods

// Function Methods

func (c *documentClass_) ExtractAlgorithm(
	document bal.DocumentLike,
) fra.QuoteLike {
	var attribute fra.QuoteLike
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$algorithm" {
			attribute = fra.QuoteFromString(
				bal.FormatDocument(association.GetDocument()),
			)
			break
		}
	}
	return attribute
}

func (c *documentClass_) ExtractAttribute(
	name string,
	document bal.DocumentLike,
) string {
	var attribute string
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			attribute = bal.FormatDocument(association.GetDocument())
			attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
			break
		}
	}
	return attribute
}

func (c *documentClass_) ExtractCertificate(
	document bal.DocumentLike,
) CitationLike {
	var certificate CitationLike
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$certificate" {
			var source = bal.FormatDocument(association.GetDocument())
			certificate = CitationClass().CitationFromString(source)
			break
		}
	}
	return certificate
}

func (c *documentClass_) ExtractDigest(
	document bal.DocumentLike,
) DigestLike {
	var digest DigestLike
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$digest" {
			var string_ = bal.FormatDocument(association.GetDocument())
			digest = DigestClass().DigestFromString(string_)
			break
		}
	}
	return digest
}

func (c *documentClass_) ExtractDocument(
	document bal.DocumentLike,
) DocumentLike {
	var result DocumentLike
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$document" {
			var string_ = bal.FormatDocument(association.GetDocument())
			result = c.DocumentFromString(string_)
			break
		}
	}
	return result
}

func (c *documentClass_) ExtractPermissions(
	document bal.DocumentLike,
) fra.ResourceLike {
	var permissions fra.ResourceLike
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$permissions" {
			var source = bal.FormatDocument(association.GetDocument())
			permissions = fra.ResourceFromString(source)
			break
		}
	}
	return permissions
}

func (c *documentClass_) ExtractPrevious(
	document bal.DocumentLike,
) CitationLike {
	var previous CitationLike
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$previous" {
			var string_ = bal.FormatDocument(association.GetDocument())
			previous = CitationClass().CitationFromString(string_)
			break
		}
	}
	return previous
}

func (c *documentClass_) ExtractSignature(
	document bal.DocumentLike,
) SignatureLike {
	var signature SignatureLike
	var component = document.GetComponent()
	var collection = component.GetAny().(bal.CollectionLike)
	var attributes = collection.GetAny().(bal.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$signature" {
			var string_ = bal.FormatDocument(association.GetDocument())
			signature = SignatureClass().SignatureFromString(string_)
			break
		}
	}
	return signature
}

func (c *documentClass_) ExtractTag(
	document bal.DocumentLike,
) fra.TagLike {
	var tag fra.TagLike
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$tag" {
			var source = bal.FormatDocument(association.GetDocument())
			tag = fra.TagFromString(source)
			break
		}
	}
	return tag
}

func (c *documentClass_) ExtractType(
	document bal.DocumentLike,
) fra.ResourceLike {
	var type_ fra.ResourceLike
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$type" {
			var source = bal.FormatDocument(association.GetDocument())
			type_ = fra.ResourceFromString(source)
			break
		}
	}
	return type_
}

func (c *documentClass_) ExtractVersion(
	document bal.DocumentLike,
) fra.VersionLike {
	var version fra.VersionLike
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(bal.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == "$version" {
			var source = bal.FormatDocument(association.GetDocument())
			version = fra.VersionFromString(source)
			break
		}
	}
	return version
}

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

func (v *document_) AsString() string {
	var document = bal.Document(v.GetComponent(), nil, "")
	var string_ = bal.FormatDocument(document)
	string_ = string_[:len(string_)-1] // Remove the trailing newline.
	string_ += `(
`
	string_ += `    $type: ` + v.GetType().AsString()
	string_ += `    $tag: ` + v.GetTag().AsString()
	string_ += `    $version: ` + v.GetVersion().AsString()
	string_ += `    $permissions: ` + v.GetPermissions().AsString()
	var previous = v.GetOptionalPrevious()
	if uti.IsDefined(previous) {
		string_ += `    $previous: ` + previous.AsString()
	}
	string_ += `)
`
	string_ = bal.FormatDocument(bal.ParseSource(string_))
	return string_
}

// Attribute Methods

func (v *document_) GetComponent() bal.ComponentLike {
	return v.component_
}

// Parameterized Methods

func (v *document_) GetType() fra.ResourceLike {
	return v.type_
}

func (v *document_) GetTag() fra.TagLike {
	return v.tag_
}

func (v *document_) GetVersion() fra.VersionLike {
	return v.version_
}

func (v *document_) GetPermissions() fra.ResourceLike {
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
	component_   bal.ComponentLike
	type_        fra.ResourceLike
	tag_         fra.TagLike
	version_     fra.VersionLike
	permissions_ fra.ResourceLike
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

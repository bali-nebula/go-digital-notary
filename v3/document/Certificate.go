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
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	algorithm fra.QuoteLike,
	publicKey fra.BinaryLike,
	tag fra.TagLike,
	version fra.VersionLike,
	optionalPrevious CitationLike,
) CertificateLike {
	if uti.IsUndefined(algorithm) {
		panic("The \"algorithm\" attribute is required by this class.")
	}
	if uti.IsUndefined(publicKey) {
		panic("The \"publicKey\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}
	var type_ = fra.ResourceFromString("<bali:/types/documents/Certificate:v3>")
	var permissions = fra.ResourceFromString("<bali:/permissions/Public:v3>")
	var instance = &certificate_{
		// Initialize the instance attributes.
		algorithm_:   algorithm,
		publicKey_:   publicKey,
		type_:        type_,
		tag_:         tag,
		version_:     version,
		permissions_: permissions,
		previous_:    optionalPrevious,
	}
	return instance
}

func (c *certificateClass_) CertificateFromString(
	source string,
) CertificateLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid certificate was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var document = not.ParseSource(source)
	var algorithm = c.extractAlgorithm(document)
	var publicKey = c.extractPublicKey(document)
	var tag = c.extractTag(document)
	var version = c.extractVersion(document)
	var previous = c.extractPrevious(document)
	return c.Certificate(
		algorithm,
		publicKey,
		tag,
		version,
		previous)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *certificate_) GetClass() CertificateClassLike {
	return certificateClass()
}

func (v *certificate_) AsString() string {
	var string_ = `[
`
	string_ += `    $algorithm: ` + v.GetAlgorithm().AsString()
	string_ += `    $publicKey: ` + v.GetPublicKey().AsString()
	string_ += `
](
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
	string_ = not.FormatDocument(not.ParseSource(string_))
	return string_
}

// Attribute Methods

func (v *certificate_) GetAlgorithm() fra.QuoteLike {
	return v.algorithm_
}

func (v *certificate_) GetPublicKey() fra.BinaryLike {
	return v.publicKey_
}

// Parameterized Methods

func (v *certificate_) GetType() fra.ResourceLike {
	return v.type_
}

func (v *certificate_) GetTag() fra.TagLike {
	return v.tag_
}

func (v *certificate_) GetVersion() fra.VersionLike {
	return v.version_
}

func (v *certificate_) GetPermissions() fra.ResourceLike {
	return v.permissions_
}

func (v *certificate_) GetOptionalPrevious() CitationLike {
	return v.previous_
}

// PROTECTED INTERFACE

// Private Methods

func (c *certificateClass_) extractAlgorithm(
	document not.DocumentLike,
) fra.QuoteLike {
	var attribute = c.extractAttribute("$algorithm", document)
	var algorithm = fra.QuoteFromString(attribute)
	return algorithm
}

func (c *certificateClass_) extractAttribute(
	name string,
	document not.DocumentLike,
) string {
	var attribute string
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(not.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			attribute = not.FormatDocument(association.GetDocument())
			attribute = attribute[:len(attribute)-1] // Remove the trailing newline.
			break
		}
	}
	return attribute
}

func (c *certificateClass_) extractParameter(
	name string,
	document not.DocumentLike,
) string {
	var parameter string
	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var element = association.GetPrimitive().GetAny().(not.ElementLike)
		var symbol = element.GetAny().(string)
		if symbol == name {
			parameter = not.FormatDocument(association.GetDocument())
			break
		}
	}
	return parameter
}

func (c *certificateClass_) extractPublicKey(
	document not.DocumentLike,
) fra.BinaryLike {
	var attribute = c.extractAttribute("$publicKey", document)
	var publicKey = fra.BinaryFromString(attribute)
	return publicKey
}

func (c *certificateClass_) extractPrevious(
	document not.DocumentLike,
) CitationLike {
	var previous CitationLike
	var parameter = c.extractParameter("$previous", document)
	if uti.IsDefined(parameter) {
		previous = CitationClass().CitationFromString(parameter)
	}
	return previous
}

func (c *certificateClass_) extractTag(
	document not.DocumentLike,
) fra.TagLike {
	var parameter = c.extractParameter("$tag", document)
	var tag = fra.TagFromString(parameter)
	return tag
}

func (c *certificateClass_) extractVersion(
	document not.DocumentLike,
) fra.VersionLike {
	var parameter = c.extractParameter("$version", document)
	var version = fra.VersionFromString(parameter)
	return version
}

// Instance Structure

type certificate_ struct {
	// Declare the instance attributes.
	algorithm_   fra.QuoteLike
	publicKey_   fra.BinaryLike
	type_        fra.ResourceLike
	tag_         fra.TagLike
	version_     fra.VersionLike
	permissions_ fra.ResourceLike
	previous_    CitationLike
}

// Class Structure

type certificateClass_ struct {
	// Declare the class constants.
}

// Class Reference

func certificateClass() *certificateClass_ {
	return certificateClassReference_
}

var certificateClassReference_ = &certificateClass_{
	// Initialize the class constants.
}

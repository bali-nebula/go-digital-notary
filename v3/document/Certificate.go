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

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	digest string,
	signature string,
	key string,
	type_ string,
	tag string,
	version string,
	permissions string,
	previous CitationLike,
) CertificateLike {
	if uti.IsUndefined(digest) {
		panic("The \"digest\" attribute is required by this class.")
	}
	if uti.IsUndefined(signature) {
		panic("The \"signature\" attribute is required by this class.")
	}
	if uti.IsUndefined(key) {
		panic("The \"key\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type_\" attribute is required by this class.")
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
	var instance = &certificate_{
		// Initialize the instance attributes.
		digest_:      digest,
		signature_:   signature,
		key_:         key,
		type_:        type_,
		tag_:         tag,
		version_:     version,
		permissions_: permissions,
		previous_:    previous,
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
	var document = doc.ParseSource(source)
	var component = document.GetComponent()
	var collection = component.GetAny().(doc.CollectionLike)
	var attributes = collection.GetAny().(doc.AttributesLike)
	var associations = attributes.GetAssociations()

	var association = associations.GetValue(1)
	var element = association.GetPrimitive().GetAny().(doc.ElementLike)
	var symbol = element.GetAny().(string)
	if symbol != "$digest" {
		panic("Missing the $digest attribute.")
	}
	var digest = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(2)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$signature" {
		panic("Missing the $signature attribute.")
	}
	var signature = doc.FormatDocument(association.GetDocument())

	association = associations.GetValue(3)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$key" {
		panic("Missing the $key attribute.")
	}
	var key = doc.FormatDocument(association.GetDocument())

	var parameters = document.GetOptionalParameters() // Not optional here.
	associations = parameters.GetAssociations()

	association = associations.GetValue(1)
	element = association.GetPrimitive().GetAny().(doc.ElementLike)
	symbol = element.GetAny().(string)
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

	return c.Certificate(
		digest,
		signature,
		key,
		type_,
		tag,
		version,
		permissions,
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
	string_ += `    $digest: ` + v.GetDigest()
	string_ += `    $signature: ` + v.GetSignature()
	string_ += `    $key: ` + v.GetKey()
	string_ += `
](
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

func (v *certificate_) GetDigest() string {
	return v.digest_
}

func (v *certificate_) GetSignature() string {
	return v.signature_
}

func (v *certificate_) GetKey() string {
	return v.key_
}

// Parameterized Methods

func (v *certificate_) GetType() string {
	return v.type_
}

func (v *certificate_) GetTag() string {
	return v.tag_
}

func (v *certificate_) GetVersion() string {
	return v.version_
}

func (v *certificate_) GetPermissions() string {
	return v.permissions_
}

func (v *certificate_) GetOptionalPrevious() CitationLike {
	return v.previous_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type certificate_ struct {
	// Declare the instance attributes.
	digest_      string
	signature_   string
	key_         string
	type_        string
	tag_         string
	version_     string
	permissions_ string
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

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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func CertificateClass() CertificateClassLike {
	return certificateClass()
}

// Constructor Methods

func (c *certificateClass_) Certificate(
	algorithm string,
	publicKey string,
	tag string,
	version string,
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
	var type_ = "<bali:/types/documents/Certificate:v3>"
	var permissions = "<bali:/permissions/Public:v3>"
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
	var document = bal.ParseSource(source)
	var algorithm = DocumentClass().ExtractAlgorithm("$algorithm", document)
	var publicKey = DocumentClass().ExtractAttribute("$publicKey", document)

	var parameters = document.GetOptionalParameters() // Not optional here.
	var associations = parameters.GetAssociations()
	var association = associations.GetValue(2)
	var element = association.GetPrimitive().GetAny().(bal.ElementLike)
	var symbol = element.GetAny().(string)
	if symbol != "$tag" {
		panic("Missing the $tag attribute.")
	}
	var tag = bal.FormatDocument(association.GetDocument())

	association = associations.GetValue(3)
	element = association.GetPrimitive().GetAny().(bal.ElementLike)
	symbol = element.GetAny().(string)
	if symbol != "$version" {
		panic("Missing the $version attribute.")
	}
	var version = bal.FormatDocument(association.GetDocument())

	var previous CitationLike
	if associations.GetSize() > 4 {
		association = associations.GetValue(5)
		element = association.GetPrimitive().GetAny().(bal.ElementLike)
		symbol = element.GetAny().(string)
		if symbol != "$previous" {
			panic("Missing the $previous attribute.")
		}
		previous = citationClass().CitationFromString(
			bal.FormatDocument(association.GetDocument()),
		)
	}

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
	string_ += `    $algorithm: "` + v.GetAlgorithm() + `"`
	string_ += `    $publicKey: ` + v.GetPublicKey()
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
	string_ = bal.FormatDocument(bal.ParseSource(string_))
	return string_
}

// Attribute Methods

func (v *certificate_) GetAlgorithm() string {
	return v.algorithm_
}

func (v *certificate_) GetPublicKey() string {
	return v.publicKey_
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
	algorithm_   string
	publicKey_   string
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

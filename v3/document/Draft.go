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

func DraftClass() DraftClassLike {
	return draftClass()
}

// Constructor Methods

func (c *draftClass_) Draft(
	component not.ComponentLike,
	type_ fra.ResourceLike,
	tag fra.TagLike,
	version fra.VersionLike,
	permissions fra.ResourceLike,
	previous CitationLike,
) DraftLike {
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
	var instance = &draft_{
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

func (c *draftClass_) DraftFromString(
	source string,
) DraftLike {
	defer func() {
		if e := recover(); e != nil {
			var message = fmt.Sprintf(
				"The following invalid draft document was passed: %s\n%v",
				source,
				e,
			)
			panic(message)
		}
	}()
	var draft = not.ParseSource(source)
	var component = draft.GetComponent()
	var type_ = c.extractType(draft)
	var tag = c.extractTag(draft)
	var version = c.extractVersion(draft)
	var permissions = c.extractPermissions(draft)
	var previous = c.extractPrevious(draft)
	return c.Draft(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *draft_) GetClass() DraftClassLike {
	return draftClass()
}

func (v *draft_) AsString() string {
	var draft = not.Document(v.GetComponent(), nil, "")
	var string_ = not.FormatDocument(draft)
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
	string_ = not.FormatDocument(not.ParseSource(string_))
	return string_
}

// Attribute Methods

func (v *draft_) GetComponent() not.ComponentLike {
	return v.component_
}

// Parameterized Methods

func (v *draft_) GetType() fra.ResourceLike {
	return v.type_
}

func (v *draft_) GetTag() fra.TagLike {
	return v.tag_
}

func (v *draft_) GetVersion() fra.VersionLike {
	return v.version_
}

func (v *draft_) GetPermissions() fra.ResourceLike {
	return v.permissions_
}

func (v *draft_) GetOptionalPrevious() CitationLike {
	return v.previous_
}

// PROTECTED INTERFACE

// Private Methods

func (c *draftClass_) extractParameter(
	name string,
	document not.DocumentLike,
) string {
	var parameter string
	document = not.GetParameter(document, name)
	if uti.IsDefined(document) {
		parameter = not.FormatDocument(document)
		parameter = parameter[:len(parameter)-1] // Remove the trailing newline.
	}
	return parameter
}

func (c *draftClass_) extractPermissions(
	document not.DocumentLike,
) fra.ResourceLike {
	var parameter = c.extractParameter("$permissions", document)
	var permissions = fra.ResourceFromString(parameter)
	return permissions
}

func (c *draftClass_) extractPrevious(
	document not.DocumentLike,
) CitationLike {
	var previous CitationLike
	var parameter = c.extractParameter("$previous", document)
	if uti.IsDefined(parameter) {
		previous = CitationClass().CitationFromString(parameter)
	}
	return previous
}

func (c *draftClass_) extractTag(
	document not.DocumentLike,
) fra.TagLike {
	var parameter = c.extractParameter("$tag", document)
	var tag = fra.TagFromString(parameter)
	return tag
}

func (c *draftClass_) extractType(
	document not.DocumentLike,
) fra.ResourceLike {
	var parameter = c.extractParameter("$type", document)
	var type_ = fra.ResourceFromString(parameter)
	return type_
}

func (c *draftClass_) extractVersion(
	document not.DocumentLike,
) fra.VersionLike {
	var parameter = c.extractParameter("$version", document)
	var version = fra.VersionFromString(parameter)
	return version
}

// Instance Structure

type draft_ struct {
	// Declare the instance attributes.
	component_   not.ComponentLike
	type_        fra.ResourceLike
	tag_         fra.TagLike
	version_     fra.VersionLike
	permissions_ fra.ResourceLike
	previous_    CitationLike
}

// Class Structure

type draftClass_ struct {
	// Declare the class constants.
}

// Class Reference

func draftClass() *draftClass_ {
	return draftClassReference_
}

var draftClassReference_ = &draftClass_{
	// Initialize the class constants.
}

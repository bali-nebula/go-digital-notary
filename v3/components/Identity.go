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

package components

import (
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func IdentityClass() IdentityClassLike {
	return identityClass()
}

// Constructor Methods

func (c *identityClass_) Identity(
	surname doc.QuoteLike,
	birthname doc.QuoteLike,
	birthdate doc.MomentLike,
	birthplace doc.QuoteLike,
	birthsex doc.SymbolLike,
	nationality doc.QuoteLike,
	address doc.NarrativeLike,
	mobile doc.QuoteLike,
	email doc.QuoteLike,
	mugshot doc.BinaryLike,
	tag doc.TagLike,
	version doc.VersionLike,
	optionalPrevious doc.ResourceLike,
) IdentityLike {
	if uti.IsUndefined(surname) {
		panic("The \"surname\" attribute is required by this class.")
	}
	if uti.IsUndefined(birthname) {
		panic("The \"birthname\" attribute is required by this class.")
	}
	if uti.IsUndefined(birthdate) {
		panic("The \"birthdate\" attribute is required by this class.")
	}
	if uti.IsUndefined(birthplace) {
		panic("The \"birthplace\" attribute is required by this class.")
	}
	if uti.IsUndefined(birthsex) {
		panic("The \"birthsex\" attribute is required by this class.")
	}
	if uti.IsUndefined(nationality) {
		panic("The \"nationality\" attribute is required by this class.")
	}
	if uti.IsUndefined(address) {
		panic("The \"address\" attribute is required by this class.")
	}
	if uti.IsUndefined(mobile) {
		panic("The \"mobile\" attribute is required by this class.")
	}
	if uti.IsUndefined(email) {
		panic("The \"email\" attribute is required by this class.")
	}
	if uti.IsUndefined(mugshot) {
		panic("The \"mugshot\" attribute is required by this class.")
	}
	if uti.IsUndefined(tag) {
		panic("The \"tag\" attribute is required by this class.")
	}
	if uti.IsUndefined(version) {
		panic("The \"version\" attribute is required by this class.")
	}

	var previous = "none"
	if uti.IsDefined(optionalPrevious) {
		previous = optionalPrevious.AsSource()
	}
	var source = `[
    $surname: ` + surname.AsSource() + `
    $birthname: ` + birthname.AsSource() + `
    $birthdate: ` + birthdate.AsSource() + `
    $birthplace: ` + birthplace.AsSource() + `
    $birthsex: ` + birthsex.AsSource() + `
    $nationality: ` + nationality.AsSource() + `
    $address: ` + address.AsSource() + `
    $mobile: ` + mobile.AsSource() + `
    $email: ` + email.AsSource() + `
    $mugshot: ` + mugshot.AsSource() + `
](
	$type: /bali/types/notary/Identity/v3
    $tag: ` + tag.AsSource() + `
    $version: ` + version.AsSource() + `
	$permissions: /bali/permissions/Public/v3
    $previous: ` + previous + `
)`
	return c.IdentityFromSource(source)
}

func (c *identityClass_) IdentityFromSource(
	source string,
) IdentityLike {
	var component = doc.ParseComponent(source)
	var instance = &identity_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Compound: component,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *identity_) GetClass() IdentityClassLike {
	return identityClass()
}

func (v *identity_) AsIntrinsic() doc.Compound {
	return v.Compound
}

func (v *identity_) AsSource() string {
	return doc.FormatComponent(v.Compound) + "\n"
}

// Attribute Methods

func (v *identity_) GetSurname() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$surname"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetBirthname() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$birthname"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetBirthdate() doc.MomentLike {
	var composite = v.GetSubcomponent(doc.Symbol("$birthdate"))
	return doc.Moment(doc.FormatComponent(composite))
}

func (v *identity_) GetBirthplace() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$birthplace"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetBirthsex() doc.SymbolLike {
	var composite = v.GetSubcomponent(doc.Symbol("$birthsex"))
	return doc.Symbol(doc.FormatComponent(composite))
}

func (v *identity_) GetNationality() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$nationality"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetAddress() doc.NarrativeLike {
	var composite = v.GetSubcomponent(doc.Symbol("$address"))
	return doc.Narrative(doc.FormatComponent(composite))
}

func (v *identity_) GetMobile() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$mobile"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetEmail() doc.QuoteLike {
	var composite = v.GetSubcomponent(doc.Symbol("$email"))
	return doc.Quote(doc.FormatComponent(composite))
}

func (v *identity_) GetMugshot() doc.BinaryLike {
	var composite = v.GetSubcomponent(doc.Symbol("$mugshot"))
	return doc.Binary(doc.FormatComponent(composite))
}

func (v *identity_) SetCertificate(
	certificate doc.ResourceLike,
) {
	v.SetSubcomponent(doc.Symbol("$certificate"), certificate)
}

func (v *identity_) GetCertificate() doc.ResourceLike {
	var composite = v.GetSubcomponent(doc.Symbol("$certificate"))
	var certificate doc.ResourceLike
	if uti.IsDefined(composite) && doc.FormatComponent(composite) != "none" {
		certificate = doc.Resource(doc.FormatComponent(composite))
	}
	return certificate
}

// Parameterized Methods

func (v *identity_) GetType() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$type"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *identity_) GetTag() doc.TagLike {
	var component = v.GetParameter(doc.Symbol("$tag"))
	return doc.Tag(doc.FormatComponent(component))
}

func (v *identity_) GetVersion() doc.VersionLike {
	var component = v.GetParameter(doc.Symbol("$version"))
	return doc.Version(doc.FormatComponent(component))
}

func (v *identity_) GetPermissions() doc.NameLike {
	var component = v.GetParameter(doc.Symbol("$permissions"))
	return doc.Name(doc.FormatComponent(component))
}

func (v *identity_) GetOptionalPrevious() doc.ResourceLike {
	var previous doc.ResourceLike
	var component = v.GetParameter(doc.Symbol("$previous"))
	if uti.IsDefined(component) {
		var source = doc.FormatComponent(component)
		if source != "none" {
			previous = doc.Resource(source)
		}
	}
	return previous
}

// Private Methods

// Instance Structure

type identity_ struct {
	// Declare the instance attributes.

	// Declare the inherited aspects.
	doc.Compound
}

// Class Structure

type identityClass_ struct {
	// Declare the class constants.
}

// Class Reference

func identityClass() *identityClass_ {
	return identityClassReference_
}

var identityClassReference_ = &identityClass_{
	// Initialize the class constants.
}

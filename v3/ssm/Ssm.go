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

package ssm

import (
	sig "crypto/ed25519"
	dig "crypto/sha512"
	fmt "fmt"
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func SsmClass() SsmClassLike {
	return ssmClass()
}

// Constructor Methods

func (c *ssmClass_) Ssm(
	directory string,
) SsmLike {
	if uti.IsUndefined(directory) {
		panic("The \"directory\" attribute is required by this class.")
	}
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += "ssm/"
	uti.MakeDirectory(directory)
	var controller = fra.Controller(c.events_, c.transitions_, c.keyless_)
	var instance = &ssm_{
		// Initialize the instance attributes.
		directory_:  directory,
		filename_:   "Configuration.bali",
		controller_: controller,
	}
	var filename = directory + instance.filename_
	if uti.PathExists(filename) {
		instance.readConfiguration()
	} else {
		instance.createConfiguration()
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *ssm_) GetClass() SsmClassLike {
	return ssmClass()
}

// Attribute Methods

// Hardened Methods

func (v *ssm_) GetTag() string {
	return v.tag_
}

func (v *ssm_) GetSignatureAlgorithm() string {
	return "ED25519"
}

func (v *ssm_) GenerateKeys() []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate new keys",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to generate keys.")
	var err error
	v.controller_.ProcessEvent(ssmClass().generateKeys_)
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	v.updateConfiguration()
	return v.publicKey_
}

func (v *ssm_) SignBytes(
	bytes []byte,
) []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to sign bytes",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to sign bytes.")
	v.controller_.ProcessEvent(ssmClass().signBytes_)
	var privateKey = v.privateKey_
	if v.previousKey_ != nil {
		// Use the old key one more time to sign the new one.
		privateKey = v.previousKey_
		v.previousKey_ = nil
	}
	var signature = sig.Sign(privateKey, bytes)
	v.updateConfiguration()
	return signature
}

func (v *ssm_) RotateKeys() []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to rotate keys",
	)

	var err error
	fmt.Println("WARNING: Using a SOFTWARE security module to rotate keys.")
	v.controller_.ProcessEvent(ssmClass().rotateKeys_)
	v.previousKey_ = v.privateKey_
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	v.updateConfiguration()
	return v.publicKey_
}

func (v *ssm_) EraseKeys() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to erase the keys",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to erase keys.")
	v.createConfiguration()
}

// Trusted Methods

func (v *ssm_) GetDigestAlgorithm() string {
	return "SHA512"
}

func (v *ssm_) DigestBytes(
	bytes []byte,
) []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to digest bytes",
	)

	var array = dig.Sum512(bytes)
	var digest = array[:] // Convert the [64]byte array to a slice.
	return digest
}

func (v *ssm_) IsValid(
	key []byte,
	signature []byte,
	bytes []byte,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify bytes signature",
	)

	return sig.Verify(sig.PublicKey(key), bytes, signature)
}

// PROTECTED INTERFACE

// Private Methods

func (c *ssmClass_) extractAttribute(
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

func (c *ssmClass_) extractKey(
	name string,
	document not.DocumentLike,
) []byte {
	var key = c.extractAttribute(name, document)
	if key == "none" {
		return nil
	}
	return fra.BinaryFromString(key).AsIntrinsic()
}

func (c *ssmClass_) extractState(
	document not.DocumentLike,
) fra.State {
	var state fra.State
	var attribute = c.extractAttribute("$state", document)
	switch attribute {
	case "$Keyless":
		state = c.keyless_
	case "$LoneKey":
		state = c.loneKey_
	case "$TwoKeys":
		state = c.twoKeys_
	case "$Invalid":
		state = c.invalid_
	}
	return state
}

func (c *ssmClass_) extractTag(
	document not.DocumentLike,
) string {
	return c.extractAttribute("$tag", document)
}

func (v *ssm_) createConfiguration() {
	v.tag_ = fra.TagWithSize(20).AsString() // Results in a 32 character tag.
	var document = not.ParseSource(`[
    $tag: ` + v.tag_ + `
    $state: $Keyless
    $publicKey: none
    $privateKey: none
    $previousKey: none
]`)
	v.extractAttributes(document)
	var source = not.FormatDocument(document)
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

func (v *ssm_) errorCheck(
	message string,
) {
	if e := recover(); e != nil {
		message = fmt.Sprintf(
			"Ssm: %s:\n        %v",
			message,
			e,
		)
		panic(message)
	}
}

func (v *ssm_) extractAttributes(
	document not.DocumentLike,
) {
	v.tag_ = ssmClass().extractTag(document)
	v.publicKey_ = ssmClass().extractKey("$publicKey", document)
	v.privateKey_ = ssmClass().extractKey("$privateKey", document)
	v.previousKey_ = ssmClass().extractKey("$previousKey", document)
	var state = ssmClass().extractState(document)
	v.controller_.SetState(state)
}

func (v *ssm_) extractDraft() not.DocumentLike {
	var tag = v.tag_
	var state = v.getState()
	var publicKey = fra.Binary(v.publicKey_).AsString()
	var privateKey = fra.Binary(v.privateKey_).AsString()
	var previousKey string
	if v.previousKey_ != nil {
		previousKey = fra.Binary(v.previousKey_).AsString()
	} else {
		previousKey = "none"
	}
	var document = not.ParseSource(`[
    $tag: ` + tag + `
    $state: ` + state + `
    $publicKey: ` + publicKey + `
    $privateKey: ` + privateKey + `
    $previousKey: ` + previousKey + `
]`)
	return document
}

func (v *ssm_) getState() string {
	switch v.controller_.GetState() {
	case ssmClass().keyless_:
		return "$Keyless"
	case ssmClass().loneKey_:
		return "$LoneKey"
	case ssmClass().twoKeys_:
		return "$TwoKeys"
	default:
		return "$Invalid"
	}
}

func (v *ssm_) readConfiguration() {
	var filename = v.directory_ + v.filename_
	var source = uti.ReadFile(filename)
	var document = not.ParseSource(source)
	v.extractAttributes(document)
}

func (v *ssm_) updateConfiguration() {
	if v.controller_.GetState() == "$Invalid" {
		panic("Invalid State")
	}
	var draft = v.extractDraft()
	var source = not.FormatDocument(draft)
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

// Instance Structure

type ssm_ struct {
	// Declare the instance attributes.
	tag_         string
	publicKey_   []byte
	privateKey_  []byte
	previousKey_ []byte
	directory_   string
	filename_    string
	controller_  fra.ControllerLike
}

// Class Structure

type ssmClass_ struct {
	// Declare the class constants.
	invalid_      fra.State
	keyless_      fra.State
	loneKey_      fra.State
	twoKeys_      fra.State
	generateKeys_ fra.Event
	signBytes_    fra.Event
	rotateKeys_   fra.Event
	events_       []fra.Event
	transitions_  map[fra.State]fra.Transitions
}

// Class Reference

func ssmClass() *ssmClass_ {
	return ssmClassReference_
}

var ssmClassReference_ = &ssmClass_{
	// Initialize the class constants.
	keyless_:      "$Keyless",
	loneKey_:      "$LoneKey",
	twoKeys_:      "$TwoKeys",
	generateKeys_: "$GenerateKeys",
	signBytes_:    "$SignBytes",
	rotateKeys_:   "$RotateKeys",
	events_:       []fra.Event{"$GenerateKeys", "$SignBytes", "$RotateKeys"},
	transitions_: map[fra.State]fra.Transitions{
		"$Keyless": fra.Transitions{"$LoneKey", "$Invalid", "$Invalid"},
		"$LoneKey": fra.Transitions{"$Invalid", "$LoneKey", "$TwoKeys"},
		"$TwoKeys": fra.Transitions{"$Invalid", "$LoneKey", "$Invalid"},
	},
}

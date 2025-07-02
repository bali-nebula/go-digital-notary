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
	bal "github.com/bali-nebula/go-bali-documents/v3"
	doc "github.com/bali-nebula/go-digital-notary/v3/document"
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

func (c *ssmClass_) Ssm() SsmLike {
	var directory = uti.HomeDirectory()
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += ".bali/ssm/"
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
	fmt.Println("WARNING: Using a SOFTWARE security module to generate keys.")
	var err error
	v.controller_.ProcessEvent(ssmClass().generateKeys_)
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		var message = fmt.Sprintf(
			"Could not generate a new public-private keypair: %v.",
			err,
		)
		panic(message)
	}
	v.updateConfiguration()
	return v.publicKey_
}

func (v *ssm_) SignBytes(
	bytes []byte,
) []byte {
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
	var err error
	fmt.Println("WARNING: Using a SOFTWARE security module to rotate keys.")
	v.controller_.ProcessEvent(ssmClass().rotateKeys_)
	v.previousKey_ = v.privateKey_
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		var message = fmt.Sprintf(
			"Could not rotate the public-private keypair: %v.",
			err,
		)
		panic(message)
	}
	v.updateConfiguration()
	return v.publicKey_
}

func (v *ssm_) EraseKeys() {
	fmt.Println("WARNING: Using a SOFTWARE security module to erase keys.")
	v.createConfiguration()
}

// Trusted Methods

func (v *ssm_) GetProtocolVersion() string {
	return "v1"
}

func (v *ssm_) GetDigestAlgorithm() string {
	return "SHA512"
}

func (v *ssm_) DigestBytes(
	bytes []byte,
) []byte {
	var array = dig.Sum512(bytes)
	var digest = array[:] // Convert the [64]byte array to a slice.
	return digest
}

func (v *ssm_) IsValid(
	key []byte,
	signature []byte,
	bytes []byte,
) bool {
	return sig.Verify(sig.PublicKey(key), bytes, signature)
}

// PROTECTED INTERFACE

// Private Methods

func (v *ssm_) extractAttributes(
	document bal.DocumentLike,
) {
	v.tag_ = doc.DocumentClass().ExtractAttribute("$tag", document)
	v.publicKey_ = v.extractKey("$publicKey", document)
	v.privateKey_ = v.extractKey("$privateKey", document)
	v.previousKey_ = v.extractKey("$previousKey", document)
	var state = v.extractState(document)
	v.controller_.SetState(state)
}

func (v *ssm_) extractDocument() bal.DocumentLike {
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
	var document = bal.ParseSource(`[
    $tag: ` + tag + `
    $state: ` + state + `
    $publicKey: ` + publicKey + `
    $privateKey: ` + privateKey + `
    $previousKey: ` + previousKey + `
]`)
	return document
}

func (v *ssm_) extractKey(
	name string,
	document bal.DocumentLike,
) []byte {
	var documentClass = doc.DocumentClass()
	var key = documentClass.ExtractAttribute(name, document)
	if key == "none" {
		return nil
	}
	return fra.BinaryFromString(key).AsIntrinsic()
}

func (v *ssm_) extractState(
	document bal.DocumentLike,
) fra.State {
	var documentClass = doc.DocumentClass()
	var state fra.State
	var attribute = documentClass.ExtractAttribute("$state", document)
	switch attribute {
	case "$Keyless":
		state = ssmClass().keyless_
	case "$LoneKey":
		state = ssmClass().loneKey_
	case "$TwoKeys":
		state = ssmClass().twoKeys_
	case "$Invalid":
		state = ssmClass().invalid_
	}
	return state
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

func (v *ssm_) createConfiguration() {
	v.tag_ = fra.TagWithSize(20).AsString() // Results in a 32 character tag.
	var document = bal.ParseSource(`[
    $tag: ` + v.tag_ + `
    $state: $Keyless
    $publicKey: none
    $privateKey: none
    $previousKey: none
]`)
	v.extractAttributes(document)
	var source = bal.FormatDocument(document)
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

func (v *ssm_) readConfiguration() {
	var filename = v.directory_ + v.filename_
	var source = uti.ReadFile(filename)
	var document = bal.ParseSource(source)
	v.extractAttributes(document)
}

func (v *ssm_) updateConfiguration() {
	if v.controller_.GetState() == "$Invalid" {
		panic("Invalid State")
	}
	var document = v.extractDocument()
	var source = bal.FormatDocument(document)
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

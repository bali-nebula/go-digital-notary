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

package agents

import (
	sig "crypto/ed25519"
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-essential-utilities/v8"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func TsmP1Class() TsmP1ClassLike {
	return tsmP1Class()
}

// Constructor Methods

func (c *tsmP1Class_) TsmP1(
	directory string,
) TsmP1Like {
	if uti.IsUndefined(directory) {
		panic("The \"directory\" attribute is required by this class.")
	}
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += "tsmP1/"
	uti.MakeDirectory(directory)
	var controller = uti.Controller(c.events_, c.transitions_, c.keyless_)
	var instance = &tsmP1_{
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

func (v *tsmP1_) GetClass() TsmP1ClassLike {
	return tsmP1Class()
}

// Attribute Methods

// Hardened Methods

func (v *tsmP1_) GetTag() string {
	fmt.Println("WARNING: Using a SOFTWARE security module to retrieve the tag.")
	return v.tag_
}

func (v *tsmP1_) GetSignatureAlgorithm() string {
	fmt.Println("WARNING: Using a SOFTWARE security module to retrieve the signature algorithm.")
	return tsmP1Class().algorithm_
}

func (v *tsmP1_) GenerateKeys() []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to generate new keys",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to generate keys.")
	var err error
	v.controller_.ProcessEvent(tsmP1Class().generateKeys_)
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	v.writeConfiguration()
	return v.publicKey_
}

func (v *tsmP1_) SignBytes(
	bytes []byte,
) []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to sign bytes",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to sign bytes.")
	v.controller_.ProcessEvent(tsmP1Class().signBytes_)
	var privateKey = v.privateKey_
	if v.previousKey_ != nil {
		// Use the old key one more time to sign the new one.
		privateKey = v.previousKey_
		v.previousKey_ = nil
	}
	var signature = sig.Sign(privateKey, bytes)
	v.writeConfiguration()
	return signature
}

func (v *tsmP1_) IsValid(
	key []byte,
	bytes []byte,
	signature []byte,
) bool {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to verify bytes signature",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to verify signature.")
	return sig.Verify(sig.PublicKey(key), bytes, signature)
}

func (v *tsmP1_) RotateKeys() []byte {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to rotate keys",
	)

	var err error
	fmt.Println("WARNING: Using a SOFTWARE security module to rotate keys.")
	v.controller_.ProcessEvent(tsmP1Class().rotateKeys_)
	v.previousKey_ = v.privateKey_
	v.publicKey_, v.privateKey_, err = sig.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	v.writeConfiguration()
	return v.publicKey_
}

func (v *tsmP1_) EraseKeys() {
	// Check for any errors at the end.
	defer v.errorCheck(
		"An error occurred while attempting to erase the keys",
	)

	fmt.Println("WARNING: Using a SOFTWARE security module to erase keys.")
	v.createConfiguration()
}

// PROTECTED INTERFACE

// Private Methods

func (v *tsmP1_) createConfiguration() {
	v.tag_ = doc.Tag().AsSource() // Results in a 32 character tag.
	v.publicKey_ = nil
	v.privateKey_ = nil
	v.previousKey_ = nil
	v.controller_ = uti.Controller(
		tsmP1Class().events_,
		tsmP1Class().transitions_,
		tsmP1Class().keyless_,
	)
	v.writeConfiguration()
}

func (v *tsmP1_) errorCheck(
	message string,
) {
	if e := recover(); e != nil {
		message = fmt.Sprintf(
			"TsmP1: %s:\n        %v",
			message,
			e,
		)
		panic(message)
	}
}

func (v *tsmP1_) readConfiguration() {
	var filename = v.directory_ + v.filename_
	var source = uti.ReadFile(filename)
	var component = doc.ParseComponent(source)
	fmt.Println(filename)

	v.tag_ = doc.FormatComponent(
		component.GetSubcomponent(doc.Symbol("$tag")),
	)

	var publicKey = doc.FormatComponent(
		component.GetSubcomponent(doc.Symbol("$publicKey")),
	)
	if publicKey != "none" {
		v.publicKey_ = doc.Binary(publicKey).AsIntrinsic()
	}

	var privateKey = doc.FormatComponent(
		component.GetSubcomponent(doc.Symbol("$privateKey")),
	)
	if privateKey != "none" {
		v.privateKey_ = doc.Binary(privateKey).AsIntrinsic()
	}

	var previousKey = doc.FormatComponent(
		component.GetSubcomponent(doc.Symbol("$previousKey")),
	)
	if previousKey != "none" {
		v.previousKey_ = doc.Binary(previousKey).AsIntrinsic()
	}

	var state = doc.FormatComponent(
		component.GetSubcomponent(doc.Symbol("$state")),
	)
	switch state {
	case "$Keyless":
		v.controller_.SetState(tsmP1Class().keyless_)
	case "$LoneKey":
		v.controller_.SetState(tsmP1Class().loneKey_)
	case "$TwoKeys":
		v.controller_.SetState(tsmP1Class().twoKeys_)
	default:
		panic("Invalid State")
	}
}

func (v *tsmP1_) writeConfiguration() {
	var tag = v.tag_

	var state string
	switch v.controller_.GetState() {
	case tsmP1Class().keyless_:
		state = "$Keyless"
	case tsmP1Class().loneKey_:
		state = "$LoneKey"
	case tsmP1Class().twoKeys_:
		state = "$TwoKeys"
	default:
		panic("Invalid State")
	}

	var publicKey = "none"
	if uti.IsDefined(v.publicKey_) {
		publicKey = doc.Binary(v.publicKey_).AsSource()
	}

	var privateKey = "none"
	if uti.IsDefined(v.privateKey_) {
		privateKey = doc.Binary(v.privateKey_).AsSource()
	}

	var previousKey = "none"
	if uti.IsDefined(v.previousKey_) {
		previousKey = doc.Binary(v.previousKey_).AsSource()
	}

	var source = `[
    $tag: ` + tag + `
    $state: ` + state + `
    $publicKey: ` + publicKey + `
    $privateKey: ` + privateKey + `
    $previousKey: ` + previousKey + `
]($type: <bali:/types/notary/TsmP1@v3>)
`
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

// Instance Structure

type tsmP1_ struct {
	// Declare the instance attributes.
	tag_         string
	publicKey_   []byte
	privateKey_  []byte
	previousKey_ []byte
	directory_   string
	filename_    string
	controller_  uti.Stateful
}

// Class Structure

type tsmP1Class_ struct {
	// Declare the class constants.
	algorithm_    string
	keyless_      uti.State
	loneKey_      uti.State
	twoKeys_      uti.State
	generateKeys_ uti.Event
	signBytes_    uti.Event
	rotateKeys_   uti.Event
	events_       []uti.Event
	transitions_  map[uti.State]uti.Transitions
}

// Class Reference

func tsmP1Class() *tsmP1Class_ {
	return tsmP1ClassReference_
}

var tsmP1ClassReference_ = &tsmP1Class_{
	// Initialize the class constants.
	algorithm_:    "ED25519",
	keyless_:      "$Keyless",
	loneKey_:      "$LoneKey",
	twoKeys_:      "$TwoKeys",
	generateKeys_: "$GenerateKeys",
	signBytes_:    "$SignBytes",
	rotateKeys_:   "$RotateKeys",
	events_:       []uti.Event{"$GenerateKeys", "$SignBytes", "$RotateKeys"},
	transitions_: map[uti.State]uti.Transitions{
		"$Keyless": uti.Transitions{"$LoneKey", "$Invalid", "$Invalid"},
		"$LoneKey": uti.Transitions{"$Invalid", "$LoneKey", "$TwoKeys"},
		"$TwoKeys": uti.Transitions{"$Invalid", "$LoneKey", "$Invalid"},
	},
}

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
	doc "github.com/bali-nebula/go-bali-documents/v3"
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
	var controller = doc.Controller(c.events_, c.transitions_, c.keyless_)
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
	v.writeConfiguration()
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
	v.writeConfiguration()
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
	v.writeConfiguration()
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

func (v *ssm_) createConfiguration() {
	v.tag_ = doc.Tag().AsSource() // Results in a 32 character tag.
	v.publicKey_ = nil
	v.privateKey_ = nil
	v.previousKey_ = nil
	v.controller_ = doc.Controller(
		ssmClass().events_,
		ssmClass().transitions_,
		ssmClass().keyless_,
	)
	v.writeConfiguration()
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

func (v *ssm_) readConfiguration() {
	var filename = v.directory_ + v.filename_
	var source = uti.ReadFile(filename)
	var component = doc.ParseComponent(source)
	fmt.Println(filename)

	v.tag_ = doc.FormatComponent(
		component.GetObject(doc.Symbol("$tag")),
	)

	var publicKey = doc.FormatComponent(
		component.GetObject(doc.Symbol("$publicKey")),
	)
	if publicKey != "none" {
		v.publicKey_ = doc.Binary(publicKey).AsIntrinsic()
	}

	var privateKey = doc.FormatComponent(
		component.GetObject(doc.Symbol("$privateKey")),
	)
	if privateKey != "none" {
		v.privateKey_ = doc.Binary(privateKey).AsIntrinsic()
	}

	var previousKey = doc.FormatComponent(
		component.GetObject(doc.Symbol("$previousKey")),
	)
	if previousKey != "none" {
		v.previousKey_ = doc.Binary(previousKey).AsIntrinsic()
	}

	var state = doc.FormatComponent(
		component.GetObject(doc.Symbol("$state")),
	)
	switch state {
	case "$Keyless":
		v.controller_.SetState(ssmClass().keyless_)
	case "$LoneKey":
		v.controller_.SetState(ssmClass().loneKey_)
	case "$TwoKeys":
		v.controller_.SetState(ssmClass().twoKeys_)
	default:
		panic("Invalid State")
	}
}

func (v *ssm_) writeConfiguration() {
	var tag = v.tag_

	var state string
	switch v.controller_.GetState() {
	case ssmClass().keyless_:
		state = "$Keyless"
	case ssmClass().loneKey_:
		state = "$LoneKey"
	case ssmClass().twoKeys_:
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
]($type: <bali:/types/notary/Ssm@v3>)`
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
	controller_  doc.ControllerLike
}

// Class Structure

type ssmClass_ struct {
	// Declare the class constants.
	keyless_      doc.State
	loneKey_      doc.State
	twoKeys_      doc.State
	generateKeys_ doc.Event
	signBytes_    doc.Event
	rotateKeys_   doc.Event
	events_       []doc.Event
	transitions_  map[doc.State]doc.Transitions
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
	events_:       []doc.Event{"$GenerateKeys", "$SignBytes", "$RotateKeys"},
	transitions_: map[doc.State]doc.Transitions{
		"$Keyless": doc.Transitions{"$LoneKey", "$Invalid", "$Invalid"},
		"$LoneKey": doc.Transitions{"$Invalid", "$LoneKey", "$TwoKeys"},
		"$TwoKeys": doc.Transitions{"$Invalid", "$LoneKey", "$Invalid"},
	},
}

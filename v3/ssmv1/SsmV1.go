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

package ssmv1

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

func SsmV1Class() SsmV1ClassLike {
	return ssmV1Class()
}

// Constructor Methods

func (c *ssmV1Class_) SsmV1(
	directory string,
) SsmV1Like {
	fmt.Println("WARNING: Using a SOFTWARE security module instead of a HARDWARE security module.")
	if uti.IsUndefined(directory) {
		panic("The \"directory\" attribute is required by this class.")
	}
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	uti.MakeDirectory(directory)
	var controller = fra.Controller(events, transitions)
	var instance = &ssmV1_{
		// Initialize the instance attributes.
		directory_:  directory,
		filename_:   "SsmV1.bali",
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

func (v *ssmV1_) GetClass() SsmV1ClassLike {
	return ssmV1Class()
}

// Attribute Methods

// V1Secure Methods

func (v *ssmV1_) GetProtocolVersion() string {
	return "v1"
}

func (v *ssmV1_) GetDigestAlgorithm() string {
	return "SHA512"
}

func (v *ssmV1_) GetSignatureAlgorithm() string {
	return "ED25519"
}

func (v *ssmV1_) DigestBytes(
	bytes []byte,
) []byte {
	var array = dig.Sum512(bytes)
	var digest = array[:] // Convert the [64]byte array to a slice.
	return digest
}

func (v *ssmV1_) IsValid(
	key []byte,
	signature []byte,
	bytes []byte,
) bool {
	return sig.Verify(sig.PublicKey(key), bytes, signature)
}

func (v *ssmV1_) GetTag() string {
	return v.tag_
}

func (v *ssmV1_) GenerateKeys() []byte {
	var err error
	v.controller_.ProcessEvent(generateKeys)
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

func (v *ssmV1_) SignBytes(
	bytes []byte,
) []byte {
	v.controller_.ProcessEvent(signBytes)
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

func (v *ssmV1_) RotateKeys() []byte {
	var err error
	v.controller_.ProcessEvent(rotateKeys)
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

func (v *ssmV1_) EraseKeys() {
	v.controller_.SetState(keyless)
	v.deleteConfiguration()
}

// PROTECTED INTERFACE

// Private Methods

func (v *ssmV1_) extractAttributes(
	document bal.DocumentLike,
) {
	v.tag_ = doc.DocumentClass().ExtractAttribute("$tag", document)
	v.publicKey_ = v.extractKey("$publicKey", document)
	v.privateKey_ = v.extractKey("$privateKey", document)
	v.previousKey_ = v.extractKey("$previousKey", document)
	var state = v.extractState(document)
	v.controller_.SetState(state)
}

func (v *ssmV1_) extractDocument() bal.DocumentLike {
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

func (v *ssmV1_) extractKey(
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

func (v *ssmV1_) extractState(
	document bal.DocumentLike,
) fra.State {
	var documentClass = doc.DocumentClass()
	var state fra.State
	switch documentClass.ExtractAttribute("$state", document) {
	case "$keyless":
		state = keyless
	case "$loneKey":
		state = loneKey
	case "$twoKeys":
		state = twoKeys
	case "$invalid":
		state = invalid
	}
	return state
}

func (v *ssmV1_) getState() string {
	switch v.controller_.GetState() {
	case keyless:
		return "$keyless"
	case loneKey:
		return "$loneKey"
	case twoKeys:
		return "$twoKeys"
	default:
		return "$invalid"
	}
}

func (v *ssmV1_) createConfiguration() {
	v.tag_ = fra.TagWithSize(20).AsString() // Results in a 32 character tag.
	var document = bal.ParseSource(`[
    $tag: ` + v.tag_ + `
    $state: $keyless
    $publicKey: none
    $privateKey: none
    $previousKey: none
]`)
	v.extractAttributes(document)
	var source = bal.FormatDocument(document)
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

func (v *ssmV1_) readConfiguration() {
	var filename = v.directory_ + v.filename_
	var source = uti.ReadFile(filename)
	var document = bal.ParseSource(source)
	v.extractAttributes(document)
}

func (v *ssmV1_) updateConfiguration() {
	var document = v.extractDocument()
	var source = bal.FormatDocument(document)
	var filename = v.directory_ + v.filename_
	uti.WriteFile(filename, source)
}

func (v *ssmV1_) deleteConfiguration() {
	v.tag_ = ""
	v.publicKey_ = nil
	v.privateKey_ = nil
	v.previousKey_ = nil
	uti.RemovePath(v.directory_ + v.filename_)
}

// Instance Structure

type ssmV1_ struct {
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

type ssmV1Class_ struct {
	// Declare the class constants.
}

// Class Reference

func ssmV1Class() *ssmV1Class_ {
	return ssmV1ClassReference_
}

var ssmV1ClassReference_ = &ssmV1Class_{
	// Initialize the class constants.
}

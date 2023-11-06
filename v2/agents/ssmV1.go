/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents

/////////////////////////////////////////////////////////////////////////////////
// This module should only be used for LOCAL TESTING or on a PHYSICALLY SECURE //
// device.  It CANNOT guarantee the protection of the private keys from people //
// and other processes that have access to the RAM and storage devices for the //
// device.                                                                     //
//                           YOU HAVE BEEN WARNED!!!                           //
/////////////////////////////////////////////////////////////////////////////////

import (
	sig "crypto/ed25519"
	dig "crypto/sha512"
	fmt "fmt"
	gcf "github.com/bali-nebula/go-component-framework/v2/abstractions"
	age "github.com/bali-nebula/go-component-framework/v2/agents"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// SOFTWARE SECURITY MODULE (SSM) INTERFACE

// This constructor creates a new software security module. It emulates a
// hardware security module and may be used for testing or, in a trusted
// cloud environment where it cannot be tampered with.
func SSMv1(directory string) abs.SecurityModuleLike {
	var configurator gcf.ConfiguratorLike
	var controller gcf.ControllerLike
	var v = &ssmV1{} // Assume this is only a trusted security module.
	if len(directory) > 0 {
		// Nope, this is pretending to be a hardened security module!
		fmt.Println("WARNING: Using a SOFTWARE security module instead of a HARDWARE security module.")
		var filename = "SSMv1.bali"
		controller = age.Controller(states)
		configurator = age.Configurator(directory, filename)
		v = &ssmV1{controller: controller, configurator: configurator}
		if configurator.Exists() {
			v.readConfiguration()
		} else {
			v.createConfiguration()
		}
	}
	return v
}

// SOFTWARE SECURITY MODULE (SSM) IMPLEMENTATION

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssmV1 struct {
	tag          string
	publicKey    []byte
	privateKey   []byte
	previousKey  []byte
	controller   gcf.ControllerLike
	configurator gcf.ConfiguratorLike
}

// These constants define the possible states for the state machine.
const (
	invalid int = iota
	keyless
	loneKey
	twoKeys
)

// These constants define the possible events for the state machine.
const (
	events int = iota
	generateKeys
	signBytes
	rotateKeys
)

// This table defines the allowed transitions for the state machine.
var states = [][]int{
	{events, generateKeys, signBytes, rotateKeys},
	{keyless, loneKey, invalid, invalid},
	{loneKey, invalid, loneKey, twoKeys},
	{twoKeys, invalid, loneKey, invalid},
}

// These constants define the possible attribute names for the configuration.
var (
	tagAttribute         = bal.Symbol("$tag")
	stateAttribute       = bal.Symbol("$state")
	publicKeyAttribute   = bal.Symbol("$publicKey")
	privateKeyAttribute  = bal.Symbol("$privateKey")
	previousKeyAttribute = bal.Symbol("$previousKey")
)

// TRUSTED INTERFACE

// This method retrieves the protocol version for this security module.
func (v *ssmV1) GetProtocol() string {
	return "v1"
}

// This method generates a digital digest of the specified bytes and returns
// the digital digest.
func (v *ssmV1) DigestBytes(bytes []byte) []byte {
	var digest = dig.Sum512(bytes)
	return digest[:] // Convert to a slice.
}

// This method determines whether or not the specified digital signature is
// valid for the specified bytes using the specified public key.
func (v *ssmV1) IsValid(key []byte, signature []byte, bytes []byte) bool {
	var isValid = sig.Verify(sig.PublicKey(key), bytes, signature)
	return isValid
}

// HARDENED INTERFACE

// This method retrieves the unique tag for this security module.
func (v *ssmV1) GetTag() string {
	return v.tag
}

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssmV1) GenerateKeys() []byte {
	var err error
	if v.configurator == nil {
		panic("Attempted to generate keys on a non-hardened security module")
	}
	v.controller.TransitionState(generateKeys)
	v.publicKey, v.privateKey, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not generate a new public-private keypair: %v.", err)
		panic(message)
	}
	v.updateConfiguration()
	return v.publicKey
}

// This method digitally signs the specified bytes using the private key and
// returns the digital signature.
func (v *ssmV1) SignBytes(bytes []byte) []byte {
	if v.configurator == nil {
		panic("Attempted to sign bytes on a non-hardened security module")
	}
	v.controller.TransitionState(signBytes)
	var privateKey = v.previousKey
	if privateKey != nil {
		v.previousKey = nil
	} else {
		privateKey = v.privateKey
	}
	var signature = sig.Sign(privateKey, bytes)
	v.updateConfiguration()
	return signature
}

// This method replaces the existing public-key pair with a new one and returns
// the new public key.
func (v *ssmV1) RotateKeys() []byte {
	var err error
	if v.configurator == nil {
		panic("Attempted to rotated keys on a non-hardened security module")
	}
	v.controller.TransitionState(rotateKeys)
	v.previousKey = v.privateKey
	v.publicKey, v.privateKey, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not rotate the public-private keypair: %v.", err)
		panic(message)
	}
	v.updateConfiguration()
	return v.publicKey
}

// This method erases the existing public-key pair.
func (v *ssmV1) EraseKeys() {
	if v.configurator == nil {
		panic("Attempted to erase keys on a non-hardened security module")
	}
	v.controller.SetState(keyless)
	v.deleteConfiguration()
}

// PRIVATE METHODS

func (v *ssmV1) getState() string {
	switch v.controller.GetState() {
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

func (v *ssmV1) setState(state gcf.ComponentLike) {
	switch bal.FormatComponent(state) {
	case "$keyless":
		v.controller.SetState(keyless)
	case "$loneKey":
		v.controller.SetState(loneKey)
	case "$twoKeys":
		v.controller.SetState(twoKeys)
	default:
		v.controller.SetState(invalid)
	}
}

func (v *ssmV1) createConfiguration() {
	v.tag = bal.FormatEntity(bal.NewTag())
	var configuration = bal.Catalog(`[
    $tag: ` + v.tag + `
    $state: ` + v.getState() + `
]`)
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssmV1) readConfiguration() {
	var document = v.configurator.Load()
	var component = bal.ParseDocument(document)
	var configuration = component.ExtractCatalog()
	v.tag = bal.FormatComponent(configuration.GetValue(tagAttribute))
	var state = configuration.GetValue(stateAttribute)
	v.setState(state)
	component = configuration.GetValue(publicKeyAttribute)
	if component != nil {
		v.publicKey = component.ExtractBinary().AsArray()
	}
	component = configuration.GetValue(privateKeyAttribute)
	if component != nil {
		v.privateKey = component.ExtractBinary().AsArray()
	}
	component = configuration.GetValue(previousKeyAttribute)
	if component != nil {
		v.previousKey = component.ExtractBinary().AsArray()
	}
}

func (v *ssmV1) updateConfiguration() {
	var configuration = bal.Catalog("[:]")
	var component = bal.Component(v.tag)
	configuration.SetValue(tagAttribute, component)
	component = bal.Component(v.getState())
	configuration.SetValue(stateAttribute, component)
	if v.publicKey != nil {
		component = bal.Component(v.publicKey)
		configuration.SetValue(publicKeyAttribute, component)
	}
	if v.privateKey != nil {
		component = bal.Component(v.privateKey)
		configuration.SetValue(privateKeyAttribute, component)
	}
	if v.previousKey != nil {
		component = bal.Component(v.previousKey)
		configuration.SetValue(previousKeyAttribute, component)
	}
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssmV1) deleteConfiguration() {
	v.tag = ""
	v.publicKey = nil
	v.privateKey = nil
	v.previousKey = nil
	v.configurator.Delete()
}

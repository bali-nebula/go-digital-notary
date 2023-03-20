/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package documents

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	ab2 "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CREDENTIALS INTERFACE

// This constructor creates a new credential.
func Credential(
	salt abs.BinaryLike,
	account abs.TagLike,
	protocol abs.VersionLike,
	certificate ab2.CitationLike,
) ab2.CredentialLike {

	// Create a new contract.
	var component = bal.Component(salt)
	var contract = Contract(component, account, protocol, certificate)

	// Change the type of the contract.
	var type_ = Type(bal.Moniker("/bali/types/documents/Credential/v1"), nil)
	var context = contract.GetContext()
	context.SetValue(ab2.TypeAttribute, bal.Component(type_))

	// Create a new credential.
	return &credential{contract}
}

// CREDENTIALS IMPLEMENTATION

type credential struct {
	ab2.ContractLike
}

// SEASONED INTERFACE

func (v *credential) GetSalt() abs.BinaryLike {
	return v.ExtractBinary()
}

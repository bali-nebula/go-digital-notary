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
	gcf "github.com/bali-nebula/go-component-framework/v2/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v2/bali"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CREDENTIALS INTERFACE

// This constructor creates a new credential.
func Credential(
	salt gcf.BinaryLike,
	account gcf.TagLike,
	protocol gcf.VersionLike,
	certificate abs.CitationLike,
) abs.CredentialLike {

	// Create a new contract.
	var component = bal.Component(salt)
	var contract = Contract(component, account, protocol, certificate)

	// Change the type of the contract.
	var type_ = Type(bal.Moniker("/bali/types/documents/Credential/v1"), nil)
	var context = contract.GetContext()
	context.SetValue(abs.TypeAttribute, bal.Component(type_))

	// Create a new credential.
	return &credential{contract}
}

// CREDENTIALS IMPLEMENTATION

type credential struct {
	abs.ContractLike
}

// SEASONED INTERFACE

func (v *credential) GetSalt() gcf.BinaryLike {
	return v.ExtractBinary()
}

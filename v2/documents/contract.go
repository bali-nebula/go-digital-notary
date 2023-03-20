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
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
	ab2 "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CONTRACT INTERFACE

// This constructor creates a new contract.
func Contract(
	component abs.ComponentLike,
	account abs.TagLike,
	protocol abs.VersionLike,
	certificate ab2.CitationLike,
) ab2.ContractLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(ab2.ComponentAttribute, component)
	attributes.SetValue(ab2.AccountAttribute, bal.Component(account))
	attributes.SetValue(ab2.ProtocolAttribute, bal.Component(protocol))
	attributes.SetValue(ab2.CertificateAttribute, certificate)

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(ab2.TypeAttribute, bal.Component("/bali/types/documents/Contract/v1"))

	// Create a new contract.
	return &contract{bal.ComponentWithContext(attributes, context)}
}

// CONTRACT IMPLEMENTATION

type contract struct {
	abs.Encapsulated
}

// NOTARIZED INTERFACE

func (v *contract) GetAccount() abs.TagLike {
	return v.ExtractCatalog().GetValue(ab2.AccountAttribute).ExtractTag()
}

func (v *contract) GetCertificate() ab2.CitationLike {
	return v.ExtractCatalog().GetValue(ab2.CertificateAttribute).(ab2.CitationLike)
}

func (v *contract) GetComponent() abs.ComponentLike {
	return v.ExtractCatalog().GetValue(ab2.ComponentAttribute).(abs.ComponentLike)
}

func (v *contract) GetProtocol() abs.VersionLike {
	return v.ExtractCatalog().GetValue(ab2.ProtocolAttribute).ExtractVersion()
}

func (v *contract) AddSignature(signature abs.BinaryLike) {
	v.ExtractCatalog().SetValue(ab2.SignatureAttribute, bal.Component(signature))
}

func (v *contract) RemoveSignature() abs.BinaryLike {
	return v.ExtractCatalog().RemoveValue(ab2.SignatureAttribute).ExtractBinary()
}

// TYPED INTERFACE

func (v *contract) GetType() ab2.TypeLike {
	return v.GetContext().GetValue(ab2.TypeAttribute).(ab2.TypeLike)
}

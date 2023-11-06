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
	col "github.com/bali-nebula/go-component-framework/v2/collections"
	com "github.com/bali-nebula/go-component-framework/v2/components"
	abs "github.com/bali-nebula/go-digital-notary/v2/abstractions"
)

// CONTRACT INTERFACE

// This constructor creates a new contract.
func Contract(
	component gcf.ComponentLike,
	account gcf.TagLike,
	protocol gcf.VersionLike,
	certificate abs.CitationLike,
) abs.ContractLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(abs.ComponentAttribute, component)
	attributes.SetValue(abs.AccountAttribute, bal.Component(account))
	attributes.SetValue(abs.ProtocolAttribute, bal.Component(protocol))
	attributes.SetValue(abs.CertificateAttribute, certificate)

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(abs.TypeAttribute, bal.Component("/bali/types/documents/Contract/v1"))

	// Create a new contract.
	return &contract{bal.ComponentWithContext(attributes, context)}
}

// CONTRACT IMPLEMENTATION

type contract struct {
	gcf.Encapsulated
}

// NOTARIZED INTERFACE

func (v *contract) GetAccount() gcf.TagLike {
	return v.ExtractCatalog().GetValue(abs.AccountAttribute).ExtractTag()
}

func (v *contract) GetCertificate() abs.CitationLike {
	return v.ExtractCatalog().GetValue(abs.CertificateAttribute).(abs.CitationLike)
}

func (v *contract) GetComponent() gcf.ComponentLike {
	return v.ExtractCatalog().GetValue(abs.ComponentAttribute)
}

func (v *contract) GetProtocol() gcf.VersionLike {
	return v.ExtractCatalog().GetValue(abs.ProtocolAttribute).ExtractVersion()
}

func (v *contract) AddSignature(signature gcf.BinaryLike) {
	v.ExtractCatalog().SetValue(abs.SignatureAttribute, bal.Component(signature))
}

func (v *contract) RemoveSignature() gcf.BinaryLike {
	return v.ExtractCatalog().RemoveValue(abs.SignatureAttribute).ExtractBinary()
}

// TYPED INTERFACE

func (v *contract) GetType() abs.TypeLike {
	return v.GetContext().GetValue(abs.TypeAttribute).(abs.TypeLike)
}

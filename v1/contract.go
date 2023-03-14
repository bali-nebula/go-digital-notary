/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package notary

import (
	abs "github.com/bali-nebula/go-component-framework/v1/abstractions"
	bal "github.com/bali-nebula/go-component-framework/v1/bali"
	col "github.com/bali-nebula/go-component-framework/v1/collections"
	com "github.com/bali-nebula/go-component-framework/v1/components"
)

// CONTRACT INTERFACE

// This constructor creates a new contract.
func Contract(
	document abs.DocumentLike,
	account abs.TagLike,
	protocol abs.VersionLike,
	certificate abs.CitationLike,
) abs.ContractLike {
	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(documentAttribute, document)
	attributes.SetValue(accountAttribute, bal.Component(account))
	attributes.SetValue(protocolAttribute, bal.Component(protocol))
	attributes.SetValue(certificateAttribute, certificate)

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/bali/types/documents/Contract/v1"))

	// Create a new contract.
	return &contract{bal.ComponentWithContext(attributes, context)}
}

// CONTRACT IMPLEMENTATION

type contract struct {
	abs.Encapsulated
}

func (v *contract) GetAccount() abs.TagLike {
	return v.ExtractCatalog().GetValue(accountAttribute).ExtractTag()
}

func (v *contract) GetCertificate() abs.CitationLike {
	return v.ExtractCatalog().GetValue(versionAttribute).(abs.CitationLike)
}

func (v *contract) GetDocument() abs.DocumentLike {
	return v.ExtractCatalog().GetValue(documentAttribute).(abs.DocumentLike)
}

func (v *contract) GetProtocol() abs.VersionLike {
	return v.ExtractCatalog().GetValue(protocolAttribute).ExtractVersion()
}

func (v *contract) GetTimestamp() abs.MomentLike {
	return v.ExtractCatalog().GetValue(timestampAttribute).ExtractMoment()
}

func (v *contract) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

func (v *contract) AddSignature(signature abs.BinaryLike) {
	v.ExtractCatalog().SetValue(signatureAttribute, bal.Component(signature))
}

func (v *contract) RemoveSignature() abs.BinaryLike {
	return v.ExtractCatalog().RemoveValue(signatureAttribute).ExtractBinary()
}

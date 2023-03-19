/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	age "github.com/bali-nebula/go-digital-notary/v2/agents"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	var module = age.SSMv1("./")
	ass.Equal(t, "v1", module.GetProtocol())
	ass.Equal(t, 64, len(module.DigestBytes(bytes)))

	var publicKey = module.GenerateKeys()

	var signature = module.SignBytes(bytes)
	ass.True(t, module.IsValid(publicKey, signature, bytes))

	var newPublicKey = module.RotateKeys()
	signature = module.SignBytes(newPublicKey)
	ass.True(t, module.IsValid(publicKey, signature, newPublicKey))

	module.EraseKeys()
}

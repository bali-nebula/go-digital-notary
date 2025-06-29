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

package ssmv1_test

import (
	ssm "github.com/bali-nebula/go-digital-notary/v3/ssmv1"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSSM(t *tes.T) {
	var bytes = []byte{0x0, 0x1, 0x2, 0x3, 0x4}
	var module = ssm.SsmV1Class().SsmV1("../test/")
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

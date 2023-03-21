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

// TYPE INTERFACE

// This constructor creates a new type component.
func Type(
	name gcf.MonikerLike,
	context gcf.ContextLike,
) abs.TypeLike {
	return &type_{bal.ComponentWithContext(name, context)}
}

// TYPE IMPLEMENTATION

type type_ struct {
	gcf.Encapsulated
}

func (v *type_) GetName() gcf.MonikerLike {
	return v.ExtractMoniker()
}

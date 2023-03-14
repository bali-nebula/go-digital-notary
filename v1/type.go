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
)

// TYPE INTERFACE

// This constructor creates a new type component.
func Type(
	moniker abs.MonikerLike,
	context abs.ContextLike,
) abs.TypeLike {
	return &type_{bal.ComponentWithContext(moniker, context)}
}

// TYPE IMPLEMENTATION

type type_ struct {
	abs.Encapsulated
}

func (v *type_) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

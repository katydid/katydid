//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package expr

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/types"
)

func NewNestedFunction(name string, params ...*Expr) *Expr {
	return &Expr{Function: &Function{
		Name:       name,
		OpenParen:  newOpenParen(),
		Params:     params,
		CloseParen: newCloseParen(),
	}}
}

func NewVar(typ types.Type) *Expr {
	return &Expr{
		Terminal: &Terminal{
			Variable: &Variable{
				Type: typ,
			},
		},
	}
}

func NewDoubleVar() *Expr {
	return NewVar(types.SINGLE_DOUBLE)
}

func NewIntVar() *Expr {
	return NewVar(types.SINGLE_INT)
}

func NewUintVar() *Expr {
	return NewVar(types.SINGLE_UINT)
}

func NewBoolVar() *Expr {
	return NewVar(types.SINGLE_BOOL)
}

func NewStringVar() *Expr {
	return NewVar(types.SINGLE_STRING)
}

func NewBytesVar() *Expr {
	return NewVar(types.SINGLE_BYTES)
}

func NewDoubleConst(d float64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			DoubleValue: proto.Float64(d),
		},
	}
}

func NewIntConst(i int64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			IntValue: proto.Int64(i),
		},
	}
}

func NewUintConst(i uint64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			UintValue: proto.Uint64(i),
		},
	}
}

func NewTrue() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(true),
	}
}

func NewFalse() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(false),
	}
}

func NewStringConst(s string) *Expr {
	return &Expr{
		Terminal: &Terminal{
			StringValue: proto.String(s),
		},
	}
}

func NewBytesConst(buf []byte) *Expr {
	return &Expr{
		Terminal: &Terminal{
			BytesValue: buf,
		},
	}
}

func NewList(typ types.Type, elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  typ,
			Elems: elems,
		},
	}
}

func NewDoubleList(elems ...*Expr) *Expr {
	return NewList(types.LIST_DOUBLE, elems...)
}

func NewIntList(elems ...*Expr) *Expr {
	return NewList(types.LIST_INT, elems...)
}

func NewUintList(elems ...*Expr) *Expr {
	return NewList(types.LIST_UINT, elems...)
}

func NewBoolList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BOOL, elems...)
}

func NewStringList(elems ...*Expr) *Expr {
	return NewList(types.LIST_STRING, elems...)
}

func NewBytesList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BYTES, elems...)
}

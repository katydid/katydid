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
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/types"
)

func newUnderscore() *Keyword {
	return &Keyword{Value: "_"}
}

func newExclamation() *Keyword {
	return &Keyword{Value: "!"}
}

func newPipe() *Keyword {
	return &Keyword{Value: "|"}
}

func NewSDTName(space *Space, term *Terminal) *NameExpr {
	name := &NameExpr{
		Name: &Name{
			Before: space,
		},
	}
	if term.DoubleValue != nil {
		name.Name.DoubleValue = term.DoubleValue
	} else if term.IntValue != nil {
		name.Name.IntValue = term.IntValue
	} else if term.UintValue != nil {
		name.Name.UintValue = term.UintValue
	} else if term.BoolValue != nil {
		name.Name.BoolValue = term.BoolValue
	} else if term.StringValue != nil {
		name.Name.StringValue = term.StringValue
	} else if term.BytesValue != nil {
		name.Name.BytesValue = term.BytesValue
	} else {
		panic(fmt.Sprintf("unreachable name type %#v", term))
	}
	return name
}

func NewDoubleName(name float64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			DoubleValue: &name,
		},
	}
}

func NewIntName(name int64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			IntValue: &name,
		},
	}
}

func NewUintName(name uint64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			UintValue: &name,
		},
	}
}

func NewBoolName(name bool) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BoolValue: &name,
		},
	}
}

func NewStringName(name string) *NameExpr {
	return &NameExpr{
		Name: &Name{
			StringValue: &name,
		},
	}
}

func NewBytesName(name []byte) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BytesValue: name,
		},
	}
}

func NewAnyName() *NameExpr {
	return &NameExpr{
		AnyName: &AnyName{Underscore: newUnderscore()},
	}
}

func NewAnyNameExcept(name *NameExpr) *NameExpr {
	return &NameExpr{
		AnyNameExcept: &AnyNameExcept{
			Exclamation: newExclamation(),
			OpenParen:   newOpenParen(),
			Except:      name,
			CloseParen:  newCloseParen(),
		},
	}
}

func NewNameChoice(names ...*NameExpr) *NameExpr {
	if len(names) == 0 {
		return nil
	}
	if len(names) == 1 {
		return names[0]
	}
	return &NameExpr{
		NameChoice: &NameChoice{
			OpenParen:  newOpenParen(),
			Left:       names[0],
			Pipe:       newPipe(),
			Right:      newNameChoice(names[1:]),
			CloseParen: newCloseParen(),
		},
	}
}

func newNameChoice(names []*NameExpr) *NameExpr {
	if len(names) == 1 {
		return names[0]
	}
	return &NameExpr{
		NameChoice: &NameChoice{
			Left:  names[0],
			Pipe:  newPipe(),
			Right: newNameChoice(names[1:]),
		},
	}
}

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

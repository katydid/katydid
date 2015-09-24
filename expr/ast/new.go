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
	"github.com/katydid/katydid/types"
)

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

func NewNestedFunction(name string, params ...*Expr) *Expr {
	return &Expr{Function: &Function{
		Name:       name,
		OpenParen:  newOpenParen(),
		Params:     params,
		CloseParen: newCloseParen(),
	}}
}

func NewStringVar() *Expr {
	return NewVar(types.SINGLE_STRING)
}

func NewBytesVar() *Expr {
	return NewVar(types.SINGLE_BYTES)
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

func NewStringConst(s string) *Expr {
	return &Expr{
		Terminal: &Terminal{
			StringValue: proto.String(s),
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

func NewBytesConst(buf []byte) *Expr {
	return &Expr{
		Terminal: &Terminal{
			BytesValue: buf,
		},
	}
}

func NewBoolList(elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  types.LIST_BOOL,
			Elems: elems,
		},
	}
}

func NewStringList(elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  types.LIST_STRING,
			Elems: elems,
		},
	}
}

func NewIntList(elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  types.LIST_INT,
			Elems: elems,
		},
	}
}

func NewEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newDoubleEqual(),
			Expr:   e,
		},
	}
}

func NewLessThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessThan(),
			Expr:   e,
		},
	}
}

func NewGreaterThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterThan(),
			Expr:   e,
		},
	}
}

func NewLessEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessEqual(),
			Expr:   e,
		},
	}
}

func NewGreaterEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterEqual(),
			Expr:   e,
		},
	}
}

func NewRegex(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newTildeEqual(),
			Expr:   e,
		},
	}
}

func NewContains(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newStarEqual(),
			Expr:   e,
		},
	}
}

func NewHasPrefix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newCaretEqual(),
			Expr:   e,
		},
	}
}

func NewHasSuffix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newDollarEqual(),
			Expr:   e,
		},
	}
}

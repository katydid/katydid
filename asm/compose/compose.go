//  Copyright 2013 Walter Schulze
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

package compose

import (
	"fmt"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/funcs"
	"github.com/awalterschulze/katydid/types"
)

type errExpected struct {
	typ  string
	expr string
}

func (this *errExpected) Error() string {
	return fmt.Sprintf("expr %v is not type %v", this.expr, this.typ)
}

type errUnknownType struct {
	expr *ast.Expr
}

func (this *errUnknownType) Error() string {
	return "expr type is unknown: " + this.expr.String()
}

func prep(expr *ast.Expr, expType types.Type) (uniq string, err error) {
	if expr.Function != nil {
		fnc := expr.GetFunction()
		uniq, err = WhichFunc(fnc)
		if err != nil {
			return "", err
		}
		typ, err := funcs.Out(uniq)
		if err != nil {
			return "", err
		}
		if typ != expType {
			return "", &errExpected{expType.String(), expr.String()}
		}
		return uniq, err
	}
	if expr.List != nil {
		if !types.IsList(expType) {
			return "", &errExpected{expType.String(), expr.String()}
		}
	}
	typ, err := Which(expr)
	if err != nil {
		return "", err
	}
	if typ != expType {
		return "", &errExpected{expType.String(), expr.String()}
	}
	return "", nil
}

func newValues(params []*ast.Expr) ([]interface{}, error) {
	values := make([]interface{}, 0, len(params))
	for _, p := range params {
		v, err := newValue(p)
		if err != nil {
			return nil, err
		}
		if v != nil {
			values = append(values, v)
		}
	}
	return values, nil
}

func newValue(p *ast.Expr) (interface{}, error) {
	typ, err := Which(p)
	if err != nil {
		return nil, err
	}
	switch typ {
	case 0:
		return nil, nil
	case types.SINGLE_DOUBLE:
		return ComposeFloat64(p)
	case types.SINGLE_FLOAT:
		return ComposeFloat32(p)
	case types.SINGLE_INT64:
		return ComposeInt64(p)
	case types.SINGLE_UINT64:
		return ComposeUint64(p)
	case types.SINGLE_INT32:
		return ComposeInt32(p)
	case types.SINGLE_BOOL:
		return ComposeBool(p)
	case types.SINGLE_STRING:
		return ComposeString(p)
	case types.SINGLE_BYTES:
		return ComposeBytes(p)
	case types.SINGLE_UINT32:
		return ComposeUint32(p)
	case types.LIST_DOUBLE:
		return ComposeFloat64s(p)
	case types.LIST_FLOAT:
		return ComposeFloat32s(p)
	case types.LIST_INT64:
		return ComposeInt64s(p)
	case types.LIST_UINT64:
		return ComposeUint64s(p)
	case types.LIST_INT32:
		return ComposeInt32s(p)
	case types.LIST_BOOL:
		return ComposeBools(p)
	case types.LIST_STRING:
		return ComposeStrings(p)
	case types.LIST_BYTES:
		return ComposeListOfBytes(p)
	case types.LIST_UINT32:
		return ComposeUint32s(p)
	}
	panic("not implemented")
}

func WhichFunc(fnc *ast.Function) (string, error) {
	types := make([]types.Type, 0, len(fnc.GetParams()))
	for _, p := range fnc.GetParams() {
		typ, err := Which(p)
		if err != nil {
			return "", err
		}
		if typ > 0 {
			types = append(types, typ)
		}
	}
	return funcs.Which(fnc.GetName(), types...)
}

func Which(expr *ast.Expr) (types.Type, error) {
	if expr.Terminal != nil {
		term := expr.GetTerminal()
		if term.DoubleValue != nil {
			return types.SINGLE_DOUBLE, nil
		}
		if term.FloatValue != nil {
			return types.SINGLE_FLOAT, nil
		}
		if term.Int64Value != nil {
			return types.SINGLE_INT64, nil
		}
		if term.Uint64Value != nil {
			return types.SINGLE_UINT64, nil
		}
		if term.Int32Value != nil {
			return types.SINGLE_INT32, nil
		}
		if term.BoolValue != nil {
			return types.SINGLE_BOOL, nil
		}
		if term.StringValue != nil {
			return types.SINGLE_STRING, nil
		}
		if term.BytesValue != nil {
			return types.SINGLE_BYTES, nil
		}
		if term.Uint32Value != nil {
			return types.SINGLE_UINT32, nil
		}
		if term.Variable != nil {
			return 0, nil
		}
	}
	if expr.List != nil {
		if expr.GetList().Type != nil {
			return expr.GetList().GetType(), nil
		}
		elems := expr.GetList().GetElems()
		if len(elems) == 0 {
			return 0, &errUnknownType{expr}
		}
		typ1, err := Which(elems[0])
		if err != nil {
			return 0, err
		}
		for i := range elems {
			typ, err := Which(elems[i])
			if err != nil {
				return 0, err
			}
			if typ != typ1 {
				return 0, &errExpected{typ.String(), typ1.String()}
			}
		}
		return types.SingleToList(typ1), nil
	}
	if expr.Function != nil {
		fnc := expr.GetFunction()
		uniq, err := WhichFunc(fnc)
		if err != nil {
			return 0, err
		}
		return funcs.Out(uniq)
	}
	return 0, &errUnknownType{expr}
}

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
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/types"
)

type errExpected struct {
	typ  string
	expr string
}

func (this *errExpected) Error() string {
	return fmt.Sprintf("expr %v is not type %v", this.expr, this.typ)
}

type errUnknownType struct {
	expr *expr.Expr
}

func (this *errUnknownType) Error() string {
	return "expr type is unknown: " + this.expr.String()
}

func prep(expr *expr.Expr, expType types.Type) (uniq string, err error) {
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

func newValues(params []*expr.Expr) ([]interface{}, error) {
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

func composeVariable(v *expr.Variable) funcs.Variable {
	switch v.Type {
	case types.SINGLE_DOUBLE:
		return funcs.NewFloat64Variable()
	case types.SINGLE_FLOAT:
		return funcs.NewFloat32Variable()
	case types.SINGLE_INT64:
		return funcs.NewInt64Variable()
	case types.SINGLE_UINT64:
		return funcs.NewUint64Variable()
	case types.SINGLE_INT32:
		return funcs.NewInt32Variable()
	case types.SINGLE_BOOL:
		return funcs.NewBoolVariable()
	case types.SINGLE_STRING:
		return funcs.NewStringVariable()
	case types.SINGLE_BYTES:
		return funcs.NewBytesVariable()
	case types.SINGLE_UINT32:
		return funcs.NewUint32Variable()
	}
	panic("unreachable")
}

func newValue(p *expr.Expr) (interface{}, error) {
	if p.Terminal != nil && p.GetTerminal().Variable != nil {
		return composeVariable(p.GetTerminal().Variable), nil
	}
	typ, err := Which(p)
	if err != nil {
		return nil, err
	}
	switch typ {
	case types.SINGLE_DOUBLE:
		return composeFloat64(p)
	case types.SINGLE_FLOAT:
		return composeFloat32(p)
	case types.SINGLE_INT64:
		return composeInt64(p)
	case types.SINGLE_UINT64:
		return composeUint64(p)
	case types.SINGLE_INT32:
		return composeInt32(p)
	case types.SINGLE_BOOL:
		return composeBool(p)
	case types.SINGLE_STRING:
		return composeString(p)
	case types.SINGLE_BYTES:
		return composeBytes(p)
	case types.SINGLE_UINT32:
		return composeUint32(p)
	case types.LIST_DOUBLE:
		return composeFloat64s(p)
	case types.LIST_FLOAT:
		return composeFloat32s(p)
	case types.LIST_INT64:
		return composeInt64s(p)
	case types.LIST_UINT64:
		return composeUint64s(p)
	case types.LIST_INT32:
		return composeInt32s(p)
	case types.LIST_BOOL:
		return composeBools(p)
	case types.LIST_STRING:
		return composeStrings(p)
	case types.LIST_BYTES:
		return composeListOfBytes(p)
	case types.LIST_UINT32:
		return composeUint32s(p)
	}
	panic("not implemented")
}

func WhichFunc(fnc *expr.Function) (string, error) {
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

func Which(expr *expr.Expr) (types.Type, error) {
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
			return term.Variable.Type, nil
		}
	}
	if expr.List != nil {
		return expr.GetList().GetType(), nil
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

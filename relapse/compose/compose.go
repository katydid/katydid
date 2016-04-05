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

//Package compose compiles a parsed expression for evaluation.
package compose

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/types"
)

//Which returns the type that the expression will return.
func Which(expr *ast.Expr) (types.Type, error) {
	if expr.Terminal != nil {
		term := expr.GetTerminal()
		if term.DoubleValue != nil {
			return types.SINGLE_DOUBLE, nil
		}
		if term.IntValue != nil {
			return types.SINGLE_INT, nil
		}
		if term.UintValue != nil {
			return types.SINGLE_UINT, nil
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

//WhichFunc returns the unique name of the function, given the types of the parameters.
//For example, a function named eq could have a unique name of intEq, doubleEq, etc.
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

func composeVariable(v *ast.Variable) funcs.Variable {
	switch v.Type {
	case types.SINGLE_DOUBLE:
		return funcs.DoubleVar()
	case types.SINGLE_INT:
		return funcs.IntVar()
	case types.SINGLE_UINT:
		return funcs.UintVar()
	case types.SINGLE_BOOL:
		return funcs.BoolVar()
	case types.SINGLE_STRING:
		return funcs.StringVar()
	case types.SINGLE_BYTES:
		return funcs.BytesVar()
	}
	panic("unreachable")
}

func newValue(p *ast.Expr) (interface{}, error) {
	if p.Terminal != nil && p.GetTerminal().Variable != nil {
		return composeVariable(p.GetTerminal().Variable), nil
	}
	typ, err := Which(p)
	if err != nil {
		return nil, err
	}
	switch typ {
	case types.SINGLE_DOUBLE:
		return composeDouble(p)
	case types.SINGLE_INT:
		return composeInt(p)
	case types.SINGLE_UINT:
		return composeUint(p)
	case types.SINGLE_BOOL:
		return composeBool(p)
	case types.SINGLE_STRING:
		return composeString(p)
	case types.SINGLE_BYTES:
		return composeBytes(p)
	case types.LIST_DOUBLE:
		return composeDoubles(p)
	case types.LIST_INT:
		return composeInts(p)
	case types.LIST_UINT:
		return composeUints(p)
	case types.LIST_BOOL:
		return composeBools(p)
	case types.LIST_STRING:
		return composeStrings(p)
	case types.LIST_BYTES:
		return composeListOfBytes(p)
	}
	panic("not implemented")
}

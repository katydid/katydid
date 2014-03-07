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
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"fmt"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"github.com/awalterschulze/katydid/exp/funcs"
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

func prep(expr *ast.Expr, expType descriptor.FieldDescriptorProto_Type) (uniq string, err error) {
	if expr.Terminal != nil {
		typ, err := Which(expr)
		if err != nil {
			return "", err
		}
		if typ != expType {
			return "", &errExpected{expType.String(), expr.String()}
		}
		return "", nil
	}
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

func ComposeBool(expr *ast.Expr) (funcs.Bool, error) {
	uniq, err := prep(expr, descriptor.FieldDescriptorProto_TYPE_BOOL)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		return funcs.NewBool(expr.GetTerminal().GetBoolValue()), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewBoolFunc(uniq, values...)
}

func ComposeInt64(expr *ast.Expr) (funcs.Int64, error) {
	uniq, err := prep(expr, descriptor.FieldDescriptorProto_TYPE_INT64)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		return funcs.NewInt64(expr.GetTerminal().GetInt64Value()), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewInt64Func(uniq, values...)
}

func ComposeString(expr *ast.Expr) (funcs.String, error) {
	uniq, err := prep(expr, descriptor.FieldDescriptorProto_TYPE_STRING)
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		return funcs.NewString(expr.GetTerminal().GetStringValue()), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.NewStringFunc(uniq, values...)
}

func newValues(params []*ast.Expr) ([]interface{}, error) {
	values := make([]interface{}, 0, len(params))
	for _, p := range params {
		typ, err := Which(p)
		if err != nil {
			return nil, err
		}
		switch typ {
		case 0:
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			f, err := ComposeBool(p)
			if err != nil {
				return nil, err
			}
			values = append(values, f)
		case descriptor.FieldDescriptorProto_TYPE_INT64:
			f, err := ComposeInt64(p)
			if err != nil {
				return nil, err
			}
			values = append(values, f)
		case descriptor.FieldDescriptorProto_TYPE_STRING:
			f, err := ComposeString(p)
			if err != nil {
				return nil, err
			}
			values = append(values, f)
		default:
			panic("not implemented")
		}
	}
	return values, nil
}

func WhichFunc(fnc *ast.Function) (string, error) {
	types := make([]descriptor.FieldDescriptorProto_Type, 0, len(fnc.GetParams()))
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

func Which(expr *ast.Expr) (descriptor.FieldDescriptorProto_Type, error) {
	if expr.Terminal != nil {
		term := expr.GetTerminal()
		if term.BoolValue != nil {
			return descriptor.FieldDescriptorProto_TYPE_BOOL, nil
		}
		if term.Int64Value != nil {
			return descriptor.FieldDescriptorProto_TYPE_INT64, nil
		}
		if term.Uint64Value != nil {
			return descriptor.FieldDescriptorProto_TYPE_UINT64, nil
		}
		if term.StringValue != nil {
			return descriptor.FieldDescriptorProto_TYPE_STRING, nil
		}
		if term.Variable != nil {
			return 0, nil
		}
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

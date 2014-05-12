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
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/funcs"
	"reflect"
)

type Bool interface {
	Eval([]byte) bool
}

type Variable interface {
	SetVariable([]byte)
}

type composedBool struct {
	Vars []Variable
	Func funcs.Bool
}

func NewBool(expr *ast.Expr) (Bool, error) {
	e, err := composeBool(expr)
	if err != nil {
		return nil, err
	}
	e = TrimBool(e)
	typ := reflect.TypeOf((*Variable)(nil)).Elem()
	impls := FuncImplements(e, typ)
	vars := make([]Variable, len(impls))
	for i := range impls {
		vars[i] = impls[i].(Variable)
	}
	return &composedBool{vars, e}, nil
}

func (this *composedBool) Eval(buf []byte) bool {
	for _, v := range this.Vars {
		v.SetVariable(buf)
	}
	return this.Func.Eval()
}

func FuncImplements(i interface{}, typ reflect.Type) []interface{} {
	e := reflect.ValueOf(i).Elem()
	var is []interface{}
	for i := 0; i < e.NumField(); i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		is = append(is, FuncImplements(e.Field(i).Interface(), typ)...)
	}
	if reflect.ValueOf(i).Type().Implements(typ) {
		is = append(is, i)
	}
	return is
}

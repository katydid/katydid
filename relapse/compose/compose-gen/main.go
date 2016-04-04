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

//compose-gen is a command that generates some of the code in the compose package.
package main

import (
	"github.com/katydid/katydid/gen"
)

const composeStr = `
func compose{{.SingleName}}(expr *relapse.Expr) (funcs.{{.SingleName}}, error) {
	uniq, err := prep(expr, types.{{.SingleType}})
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.{{.SingleName}}Var(), nil
		} else {
			return funcs.{{.SingleName}}Const(expr.GetTerminal().Get{{.SingleValue}}Value()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.New{{.SingleName}}Func(uniq, values...)
}

func compose{{.ListName}}(expr *relapse.Expr) (funcs.{{.ListName}}, error) {
	uniq, err := prep(expr, types.{{.ListType}})
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.{{.SingleName}}, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.{{.SingleName}})
			if !ok {
				return nil, &errExpected{types.{{.SingleType}}.String(), expr.String()}
			}
		}
		return funcs.NewListOf{{.SingleName}}(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return funcs.New{{.ListName}}Func(uniq, values...)
}
`

type composer struct {
	SingleName  string
	SingleType  string
	SingleValue string
	ListName    string
	ListType    string
}

func main() {
	gen := gen.NewPackage("compose")
	gen(composeStr, "compose.gen.go", []interface{}{
		&composer{"Double", "SINGLE_DOUBLE", "Double", "Doubles", "LIST_DOUBLE"},
		&composer{"Int", "SINGLE_INT", "Int", "Ints", "LIST_INT"},
		&composer{"Uint", "SINGLE_UINT", "Uint", "Uints", "LIST_UINT"},
		&composer{"Bool", "SINGLE_BOOL", "Bool", "Bools", "LIST_BOOL"},
		&composer{"String", "SINGLE_STRING", "String", "Strings", "LIST_STRING"},
		&composer{"Bytes", "SINGLE_BYTES", "Bytes", "ListOfBytes", "LIST_BYTES"},
	},
		`"github.com/katydid/katydid/relapse"`,
		`"github.com/katydid/katydid/relapse/funcs"`,
		`"github.com/katydid/katydid/relapse/types"`)
}

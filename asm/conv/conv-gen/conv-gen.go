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

package main

import (
	"github.com/awalterschulze/katydid/gen"
)

const convStr = `
func {{.CNames}}FromGo(vs []{{.GoType}}) string {
	vss := make([]string, len(vs))
	for i, v := range vs {
		vss[i] = fmt.Sprintf("{{.Sprintf}}", v)
	}
	return "[]{{.Type}}{" + strings.Join(vss, ",") + "}"
}

func {{.CNames}}ToGo(s string) ([]{{.GoType}}, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		return nil, &ErrWrongType{"[]{{.Type}}", "not a list"}
	}
	list := expr.GetList()
	elems := list.GetElems()
	if list.GetType() != types.LIST_{{.ASMTYPE}} {
		return nil, &ErrWrongType{"[]{{.Type}}", list.GetType().String()}
	}
	return {{.Names}}ToGo(elems)
}

func {{.Names}}ToGo(elems []*ast.Expr) ([]{{.GoType}}, error) {
	vs := make([]{{.GoType}}, len(elems))
	for i := range elems {
		if elems[i].GetTerminal().{{.CValue}}Value == nil {
			return nil, &ErrWrongType{"{{.Type}}", elems[i].GetTerminal().String()}
		}
		vs[i] = elems[i].GetTerminal().Get{{.CValue}}Value()
	}
	return vs, nil
}

func {{.CName}}FromGo(v {{.GoType}}) string {
	return fmt.Sprintf("{{.Sprintf}}", v)
}

func {{.CName}}ToGo(s string) ({{.GoType}}, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return {{.Nil}}, err
	}
	if expr.Terminal == nil {
		return {{.Nil}}, &ErrWrongType{"{{.Type}}", "not a terminal"}
	}
	return {{.Name}}ToGo(expr.GetTerminal())
}

func {{.Name}}ToGo(term *ast.Terminal) ({{.GoType}}, error) {
	if term.{{.CValue}}Value == nil {
		return {{.Nil}}, &ErrWrongType{"{{.Type}}", term.String()}
	}
	return term.Get{{.CValue}}Value(), nil
}

`

type conver struct {
	GoType  string
	Type    string
	ASMTYPE string
	Sprintf string
	Nil     string
}

func (this *conver) Name() string {
	if this.Type == "[]byte" {
		return "bytes"
	}
	return this.Type
}

func (this *conver) CName() string {
	return gen.CapFirst(this.Name())
}

func (this *conver) Names() string {
	if this.Type == "[]byte" {
		return "listOfBytes"
	}
	return this.Type + "s"
}

func (this *conver) CNames() string {
	return gen.CapFirst(this.Names())
}

func (this *conver) CValue() string {
	if this.Type == "[]byte" {
		return "Bytes"
	}
	return gen.CapFirst(this.Type)
}

func main() {
	gen := gen.NewFunc("conv")
	gen(convStr, "conv.gen.go", []interface{}{
		&conver{"float64", "double", "DOUBLE", "double(%f)", "0"},
		&conver{"float32", "float", "FLOAT", "float(%f)", "0"},
		&conver{"int64", "int64", "INT64", "int64(%d)", "0"},
		&conver{"uint64", "uint64", "UINT64", "uint64(%d)", "0"},
		&conver{"int32", "int32", "INT32", "int32(%d)", "0"},
		&conver{"bool", "bool", "BOOL", "%v", "false"},
		&conver{"string", "string", "STRING", "`%s`", `""`},
		&conver{"[]byte", "[]byte", "BYTES", "%#v", "nil"},
		&conver{"uint32", "uint32", "UINT32", "uint32(%d)", "0"},
	},
		`"fmt"`,
		`"github.com/awalterschulze/katydid/asm/lexer"`,
		`"github.com/awalterschulze/katydid/asm/parser"`,
		`"github.com/awalterschulze/katydid/asm/ast"`,
		`"github.com/awalterschulze/katydid/types"`,
		`"strings"`)
}

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

//debug-gen is a command that generates some of the code in the debug package.
package main

import (
	"github.com/katydid/katydid/gen"
	"strings"
)

type valuer struct {
	CName   string
	GoType  string
	Default string
}

func (v *valuer) Name() string {
	return strings.ToLower(v.CName)
}

const valueStr = `
func (*errValue) {{.CName}}() ({{.GoType}}, error) {
	return {{.Default}}, parser.ErrNot{{.CName}}
}

type {{.Name}}Value struct {
	*errValue
	v {{.GoType}}
}

//New{{.CName}}Value wraps a native go type into a parser.Value.
func New{{.CName}}Value(v {{.GoType}}) parser.Value {
	return &{{.Name}}Value{&errValue{}, v}
}

func (v *{{.Name}}Value) {{.CName}}() ({{.GoType}}, error) {
	return v.v, nil
}

`

func main() {
	gen := gen.NewPackage("debug")
	gen(valueStr, "value.gen.go", []interface{}{
		&valuer{"Double", "float64", "0"},
		&valuer{"Int", "int64", "0"},
		&valuer{"Uint", "uint64", "0"},
		&valuer{"Bool", "bool", "false"},
		&valuer{"String", "string", `""`},
		&valuer{"Bytes", "[]byte", "nil"},
	}, `"github.com/katydid/katydid/parser"`)
}

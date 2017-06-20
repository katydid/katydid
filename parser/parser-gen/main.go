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

//Command parser-gen generates some of the code in the parser package.
package main

import (
	"strings"

	"github.com/katydid/katydid/gen"
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
//ErrNot{{.CName}} is an error that represents a type error.
var ErrNot{{.CName}} = fmt.Errorf("value is not a {{.Name}}")
`

func main() {
	gen := gen.NewPackage("parser")
	gen(valueStr, "err.gen.go", []interface{}{
		&valuer{"Double", "float64", "0"},
		&valuer{"Int", "int64", "0"},
		&valuer{"Uint", "uint64", "0"},
		&valuer{"Bool", "bool", "false"},
		&valuer{"String", "string", `""`},
		&valuer{"Bytes", "[]byte", "nil"},
	}, `"fmt"`)
}

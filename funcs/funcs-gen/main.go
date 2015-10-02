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
	"github.com/katydid/katydid/gen"
	"strings"
)

const compareStr = `
type {{.Type}}{{.CName}} struct {
	V1 {{.CType}}
	V2 {{.CType}}
}

func (this *{{.Type}}{{.CName}}) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, err
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, err
	}
	{{if .Eval}}{{.Eval}}{{else}}return v1 {{.Operator}} v2, nil{{end}}
}

func init() {
	Register("{{.Name}}", new({{.Type}}{{.CName}}))
}

func {{.CType}}{{.CName}}(a, b {{.CType}}) Bool {
	return &{{.Type}}{{.CName}}{V1: a, V2: b}
}

func {{.CType}}Var{{.CName}}(a {{.CType}}) Bool {
	return &{{.Type}}{{.CName}}{V1: a, V2: {{.CType}}Var()}
}
`

type compare struct {
	Name     string
	Operator string
	Type     string
	Eval     string
	CType    string
}

func (this *compare) CName() string {
	if this.Name == "ge" || this.Name == "le" {
		return strings.ToUpper(this.Name)
	}
	return gen.CapFirst(this.Name)
}

const newFuncStr = `
func New{{.}}Func(uniq string, values ...interface{}) ({{.}}, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.({{.}}), nil
}
`

const constStr = `
type Const{{.CType}} interface {
	{{.CType}}
}

var typConst{{.CType}} reflect.Type = reflect.TypeOf((*Const{{.CType}})(nil)).Elem()

type const{{.CType}} struct {
	v {{.GoType}}
}

func {{.CType}}Const(v {{.GoType}}) Const{{.CType}} {
	return &const{{.CType}}{v}
}

func (this *const{{.CType}}) IsConst() {}

func (this *const{{.CType}}) Eval() ({{.GoType}}, error) {
	return this.v, nil
}

func (this *const{{.CType}}) String() string {
	{{if .ListType}}ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("{{.String}}", this.v[i])
	}
	return "[]{{.ListType}}{" + strings.Join(ss, ",") + "}"{{else}}return fmt.Sprintf("{{.String}}", this.v){{end}}
}
`

type conster struct {
	CType    string
	GoType   string
	String   string
	ListType string
}

const listStr = `
type listOf{{.FuncType}} struct {
	List []{{.FuncType}}
}

func NewListOf{{.FuncType}}(v []{{.FuncType}}) {{.CType}} {
	return &listOf{{.FuncType}}{v}
}

func (this *listOf{{.FuncType}}) Eval() ([]{{.GoType}}, error) {
	res := make([]{{.GoType}}, len(this.List))
	var err error
	for i, e := range this.List {
		res[i], err = e.Eval()
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (this *listOf{{.FuncType}}) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]{{.GoType}}{" + strings.Join(ss, ",") + "}"
}

func (this *listOf{{.FuncType}}) IsListOf() {}
`

type list struct {
	CType    string
	FuncType string
	GoType   string
}

const printStr = `
type print{{.Name}} struct {
	E {{.Name}}
}

func (this *print{{.Name}}) Eval() ({{.GoType}}, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *print{{.Name}}) IsVariable() {}

func init() {
	Register("print", new(print{{.Name}}))
}

func Print{{.Name}}(e {{.Name}}) {{.Name}} {
	return &print{{.Name}}{E: e}
}
`

type printer struct {
	Name   string
	GoType string
}

const lengthStr = `
type len{{.}} struct {
	E {{.}}
}

func (this *len{{.}}) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func init() {
	Register("length", new(len{{.}}))
}

func Len{{.}}(e {{.}}) Int {
	return &len{{.}}{E: e}
}
`

const elemStr = `
type elem{{.ListType}} struct {
	List {{.ListType}}
	Index Int
}

func (this *elem{{.ListType}}) Eval() ({{.ReturnType}}, error) {
	list, err := this.List.Eval()
	if err != nil {
		return {{.Default}}, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return {{.Default}}, err
	}
	index := int(index64)
	if len(list) == 0 {
		return {{.Default}}, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)	
	}
	if len(list) <= index {
		return {{.Default}}, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func init() {
	Register("elem", new(elem{{.ListType}}))
}

func Elem{{.ListType}}(list {{.ListType}}, index Int) {{.ThrowType}} {
	return &elem{{.ListType}}{
		List: list,
		Index: index,
	}
}
`

type elemer struct {
	ListType   string
	ReturnType string
	ThrowType  string
	Default    string
}

const rangeStr = `
type range{{.ListType}} struct {
	List {{.ListType}}
	First Int
	Last Int
}

func (this *range{{.ListType}}) Eval() ({{.ReturnType}}, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)	
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)	
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func init() {
	Register("range", new(range{{.ListType}}))
}

func Range{{.ListType}}(list {{.ListType}}, from, to Int) {{.ListType}} {
	return &range{{.ListType}}{
		List: list,
		First: from,
		Last: to,
	}
}
`

type ranger struct {
	ListType   string
	ReturnType string
}

const variableStr = `
type var{{.Name}} struct {
	Dec serialize.Decoder
}

var _ Decoder = &var{{.Name}}{}
var _ Variable = &var{{.Name}}{}

func (this *var{{.Name}}) Eval() ({{.GoType}}, error) {
	v, err := this.Dec.{{.Name}}()
	if err != nil {
		return {{.Default}}, err
	}
	return v, nil
}

func (this *var{{.Name}}) IsVariable() {}

func (this *var{{.Name}}) SetDecoder(dec serialize.Decoder) {
	this.Dec = dec
}

func (this *var{{.Name}}) String() string {
	return "${{.Decode}}"
}

func {{.Name}}Var() *var{{.Name}} {
	return &var{{.Name}}{}
}

`

type varer struct {
	Name    string
	Decode  string
	GoType  string
	Default string
}

const typStr = `
type typ{{.Name}} struct {
	E {{.Name}}
}

func (this *typ{{.Name}}) Eval() (bool, error) {
	_, err := this.E.Eval()
	return (err == nil), nil
}

func init() {
	Register("type", new(typ{{.Name}}))
}

func Type{{.Name}}() Bool {
	return &typ{{.Name}}{E: {{.Name}}Var()}
}
`

type typer struct {
	Name string
}

func main() {
	gen := gen.NewFunc("funcs")
	gen(compareStr, "compare.gen.go", []interface{}{
		&compare{"ge", ">=", "double", "", "Double"},
		&compare{"ge", ">=", "int", "", "Int"},
		&compare{"ge", ">=", "uint", "", "Uint"},
		&compare{"ge", "", "bytes", "return bytes.Compare(v1, v2) >= 0, nil", "Bytes"},
		&compare{"gt", ">", "double", "", "Double"},
		&compare{"gt", ">", "int", "", "Int"},
		&compare{"gt", ">", "uint", "", "Uint"},
		&compare{"gt", "", "bytes", "return bytes.Compare(v1, v2) > 0, nil", "Bytes"},
		&compare{"le", "<=", "double", "", "Double"},
		&compare{"le", "<=", "int", "", "Int"},
		&compare{"le", "<=", "uint", "", "Uint"},
		&compare{"le", "", "bytes", "return bytes.Compare(v1, v2) <= 0, nil", "Bytes"},
		&compare{"lt", "<", "double", "", "Double"},
		&compare{"lt", "<", "int", "", "Int"},
		&compare{"lt", "<", "uint", "", "Uint"},
		&compare{"lt", "", "bytes", "return bytes.Compare(v1, v2) < 0, nil", "Bytes"},
		&compare{"eq", "==", "double", "", "Double"},
		&compare{"eq", "==", "int", "", "Int"},
		&compare{"eq", "==", "uint", "", "Uint"},
		&compare{"eq", "==", "bool", "", "Bool"},
		&compare{"eq", "==", "string", "", "String"},
		&compare{"eq", "", "bytes", "return bytes.Equal(v1, v2), nil", "Bytes"},
		&compare{"ne", "!=", "double", "", "Double"},
		&compare{"ne", "!=", "int", "", "Int"},
		&compare{"ne", "!=", "uint", "", "Uint"},
		&compare{"ne", "!=", "bool", "", "Bool"},
		&compare{"ne", "!=", "string", "", "String"},
		&compare{"ne", "", "bytes", "return !bytes.Equal(v1, v2), nil", "Bytes"},
	}, `"bytes"`)
	gen(newFuncStr, "newfunc.gen.go", []interface{}{
		"Double",
		"Int",
		"Uint",
		"Bool",
		"String",
		"Bytes",
		"Doubles",
		"Ints",
		"Uints",
		"Bools",
		"Strings",
		"ListOfBytes",
	})
	gen(constStr, "const.gen.go", []interface{}{
		&conster{"Double", "float64", "double(%f)", ""},
		&conster{"Int", "int64", "int(%d)", ""},
		&conster{"Uint", "uint64", "uint(%d)", ""},
		&conster{"Bool", "bool", "%v", ""},
		&conster{"String", "string", "`%s`", ""},
		&conster{"Bytes", "[]byte", "%#v", ""},
		&conster{"Doubles", "[]float64", "double(%f)", "double"},
		&conster{"Ints", "[]int64", "int(%d)", "int"},
		&conster{"Uints", "[]uint64", "uint(%d)", "uint"},
		&conster{"Bools", "[]bool", "%v", "bool"},
		&conster{"Strings", "[]string", "`%s`", "string"},
		&conster{"ListOfBytes", "[][]byte", "%#v", "[]byte"},
	}, `"fmt"`, `"strings"`, `"reflect"`)
	gen(listStr, "list.gen.go", []interface{}{
		&list{"Doubles", "Double", "float64"},
		&list{"Ints", "Int", "int64"},
		&list{"Uints", "Uint", "uint64"},
		&list{"Bools", "Bool", "bool"},
		&list{"Strings", "String", "string"},
		&list{"ListOfBytes", "Bytes", "[]byte"},
	}, `"strings"`)
	gen(printStr, "print.gen.go", []interface{}{
		&printer{"Double", "float64"},
		&printer{"Int", "int64"},
		&printer{"Uint", "uint64"},
		&printer{"Bool", "bool"},
		&printer{"String", "string"},
		&printer{"Bytes", "[]byte"},
		&printer{"Doubles", "[]float64"},
		&printer{"Ints", "[]int64"},
		&printer{"Uints", "[]uint64"},
		&printer{"Bools", "[]bool"},
		&printer{"Strings", "[]string"},
		&printer{"ListOfBytes", "[][]byte"},
	}, `"fmt"`)
	gen(lengthStr, "length.gen.go", []interface{}{
		"Doubles",
		"Ints",
		"Uints",
		"Bools",
		"Strings",
		"ListOfBytes",
		"String",
		"Bytes",
	})
	gen(elemStr, "elem.gen.go", []interface{}{
		&elemer{"Doubles", "float64", "Double", "0"},
		&elemer{"Ints", "int64", "Int", "0"},
		&elemer{"Uints", "uint64", "Uint", "0"},
		&elemer{"Bools", "bool", "Bool", "false"},
		&elemer{"Strings", "string", "String", `""`},
		&elemer{"ListOfBytes", "[]byte", "Bytes", "nil"},
	})
	gen(rangeStr, "range.gen.go", []interface{}{
		&ranger{"Doubles", "[]float64"},
		&ranger{"Ints", "[]int64"},
		&ranger{"Uints", "[]uint64"},
		&ranger{"Bools", "[]bool"},
		&ranger{"Strings", "[]string"},
		&ranger{"ListOfBytes", "[][]byte"},
	})
	gen(variableStr, "variable.gen.go", []interface{}{
		&varer{"Double", "double", "float64", "0"},
		&varer{"Int", "int", "int64", "0"},
		&varer{"Uint", "uint", "uint64", "0"},
		&varer{"Bool", "bool", "bool", "false"},
		&varer{"String", "string", "string", `""`},
		&varer{"Bytes", "[]byte", "[]byte", "nil"},
	}, `"github.com/katydid/katydid/serialize"`)
	gen(typStr, "type.gen.go", []interface{}{
		&typer{"Double"},
		&typer{"Int"},
		&typer{"Uint"},
		&typer{"Bool"},
		&typer{"String"},
		&typer{"Bytes"},
	})
}

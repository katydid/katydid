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

const compareStr = `
type {{.Type}}{{.CName}} struct {
	V1 {{.CType}}
	V2 {{.CType}}
}

func (this *{{.Type}}{{.CName}}) Eval() bool {
	{{if .Eval}}{{.Eval}}{{else}}return this.V1.Eval() {{.Operator}} this.V2.Eval(){{end}}
}

func init() {
	Register("{{.Name}}", new({{.Type}}{{.CName}}))
}
`

type compare struct {
	Name     string
	Operator string
	Type     string
	Eval     string
}

func (this *compare) CName() string {
	return gen.CapFirst(this.Name)
}

func (this *compare) CType() string {
	return gen.CapFirst(this.Type)
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

func NewConst{{.CType}}(v {{.GoType}}) Const{{.CType}} {
	return &const{{.CType}}{v}
}

func (this *const{{.CType}}) IsConst() {}

func (this *const{{.CType}}) Eval() {{.GoType}} {
	return this.v
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

func (this *listOf{{.FuncType}}) Eval() []{{.GoType}} {
	res := make([]{{.GoType}}, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
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

func (this *print{{.Name}}) Eval() {{.GoType}} {
	v := this.E.Eval()
	fmt.Printf("%#v\n", v)
	return v
}

func (this *print{{.Name}}) SetVariable([]byte) {}

func init() {
	Register("print", new(print{{.Name}}))
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

func (this *len{{.}}) Eval() int64 {
	return int64(len(this.E.Eval()))
}

func init() {
	Register("length", new(len{{.}}))
}
`

const elemStr = `
type elem{{.ListType}} struct {
	List {{.ListType}}
	Index Int64
	Thrower
}

func (this *elem{{.ListType}}) Eval() {{.ReturnType}} {
	list := this.List.Eval()
	index := int(this.Index.Eval())
	if len(list) == 0 {
		return this.Throw{{.ThrowType}}(NewRangeCheckErr(index, len(list)))
	}
	if index < 0 {
		index = index % len(list)	
	}
	if len(list) <= index {
		return this.Throw{{.ThrowType}}(NewRangeCheckErr(index, len(list)))
	}
	return list[index]
}

func init() {
	Register("elem", new(elem{{.ListType}}))
}
`

type elemer struct {
	ListType   string
	ReturnType string
	ThrowType  string
}

const rangeStr = `
type range{{.ListType}} struct {
	List {{.ListType}}
	First Int64
	Last Int64
	Thrower
}

func (this *range{{.ListType}}) Eval() {{.ReturnType}} {
	list := this.List.Eval()
	first := int(this.First.Eval())
	if len(list) == 0 {
		return this.Throw{{.ListType}}(NewRangeCheckErr(first, len(list)))
	}
	if first < 0 {
		first = first % len(list)	
	}
	if first > len(list) {
		return this.Throw{{.ListType}}(NewRangeCheckErr(first, len(list)))
	}
	last := int(this.Last.Eval())
	if last < 0 {
		last = last % len(list)	
	}
	if last > len(list) {
		return this.Throw{{.ListType}}(NewRangeCheckErr(last, len(list)))
	}
	if first > last {
		return this.Throw{{.ListType}}(NewRangeErr(first, last))	
	}
	return list[first:last]
}

func init() {
	Register("range", new(range{{.ListType}}))
}
`

type ranger struct {
	ListType   string
	ReturnType string
}

func main() {
	gen := gen.NewFunc("funcs")
	gen(compareStr, "compare.gen.go", []interface{}{
		&compare{"ge", ">=", "float64", ""},
		&compare{"ge", ">=", "float32", ""},
		&compare{"ge", ">=", "int64", ""},
		&compare{"ge", ">=", "uint64", ""},
		&compare{"ge", ">=", "int32", ""},
		&compare{"ge", ">=", "uint32", ""},
		&compare{"ge", "", "bytes", "return bytes.Compare(this.V1.Eval(), this.V2.Eval()) >= 0"},
		&compare{"gt", ">", "float64", ""},
		&compare{"gt", ">", "float32", ""},
		&compare{"gt", ">", "int64", ""},
		&compare{"gt", ">", "uint64", ""},
		&compare{"gt", ">", "int32", ""},
		&compare{"gt", ">", "uint32", ""},
		&compare{"gt", "", "bytes", "return bytes.Compare(this.V1.Eval(), this.V2.Eval()) > 0"},
		&compare{"le", "<=", "float64", ""},
		&compare{"le", "<=", "float32", ""},
		&compare{"le", "<=", "int64", ""},
		&compare{"le", "<=", "uint64", ""},
		&compare{"le", "<=", "int32", ""},
		&compare{"le", "<=", "uint32", ""},
		&compare{"le", "", "bytes", "return bytes.Compare(this.V1.Eval(), this.V2.Eval()) <= 0"},
		&compare{"lt", "<", "float64", ""},
		&compare{"lt", "<", "float32", ""},
		&compare{"lt", "<", "int64", ""},
		&compare{"lt", "<", "uint64", ""},
		&compare{"lt", "<", "int32", ""},
		&compare{"lt", "<", "uint32", ""},
		&compare{"lt", "", "bytes", "return bytes.Compare(this.V1.Eval(), this.V2.Eval()) < 0"},
		&compare{"eq", "==", "float64", ""},
		&compare{"eq", "==", "float32", ""},
		&compare{"eq", "==", "int64", ""},
		&compare{"eq", "==", "uint64", ""},
		&compare{"eq", "==", "int32", ""},
		&compare{"eq", "==", "uint32", ""},
		&compare{"eq", "==", "bool", ""},
		&compare{"eq", "==", "string", ""},
		&compare{"eq", "", "bytes", "return bytes.Equal(this.V1.Eval(), this.V2.Eval())"},
	}, `"bytes"`)
	gen(newFuncStr, "newfunc.gen.go", []interface{}{
		"Float64",
		"Float32",
		"Int64",
		"Uint64",
		"Int32",
		"Uint32",
		"Bool",
		"String",
		"Bytes",
		"Float64s",
		"Float32s",
		"Int64s",
		"Uint64s",
		"Int32s",
		"Uint32s",
		"Bools",
		"Strings",
		"ListOfBytes",
	})
	gen(constStr, "const.gen.go", []interface{}{
		&conster{"Float64", "float64", "double(%f)", ""},
		&conster{"Float32", "float32", "float(%f)", ""},
		&conster{"Int64", "int64", "int64(%d)", ""},
		&conster{"Uint64", "uint64", "uint64(%d)", ""},
		&conster{"Int32", "int32", "int32(%d)", ""},
		&conster{"Uint32", "uint32", "uint32(%d)", ""},
		&conster{"Bool", "bool", "%v", ""},
		&conster{"String", "string", "`%s`", ""},
		&conster{"Bytes", "[]byte", "%#v", ""},
		&conster{"Float64s", "[]float64", "double(%f)", "double"},
		&conster{"Float32s", "[]float32", "float(%f)", "float"},
		&conster{"Int64s", "[]int64", "int64(%d)", "int64"},
		&conster{"Uint64s", "[]uint64", "uint64(%d)", "uint64"},
		&conster{"Int32s", "[]int32", "int32(%d)", "int32"},
		&conster{"Uint32s", "[]uint32", "uint32(%d)", "uint32"},
		&conster{"Bools", "[]bool", "%v", "bool"},
		&conster{"Strings", "[]string", "`%s`", "string"},
		&conster{"ListOfBytes", "[][]byte", "%#v", "[]byte"},
	}, `"fmt"`, `"strings"`, `"reflect"`)
	gen(listStr, "list.gen.go", []interface{}{
		&list{"Float64s", "Float64", "float64"},
		&list{"Float32s", "Float32", "float32"},
		&list{"Int64s", "Int64", "int64"},
		&list{"Uint64s", "Uint64", "uint64"},
		&list{"Int32s", "Int32", "int32"},
		&list{"Bools", "Bool", "bool"},
		&list{"Strings", "String", "string"},
		&list{"ListOfBytes", "Bytes", "[]byte"},
		&list{"Uint32s", "Uint32", "uint32"},
	}, `"strings"`)
	gen(printStr, "print.gen.go", []interface{}{
		&printer{"Float64", "float64"},
		&printer{"Float32", "float32"},
		&printer{"Int64", "int64"},
		&printer{"Uint64", "uint64"},
		&printer{"Int32", "int32"},
		&printer{"Uint32", "uint32"},
		&printer{"Bool", "bool"},
		&printer{"String", "string"},
		&printer{"Bytes", "[]byte"},
		&printer{"Float64s", "[]float64"},
		&printer{"Float32s", "[]float32"},
		&printer{"Int64s", "[]int64"},
		&printer{"Uint64s", "[]uint64"},
		&printer{"Int32s", "[]int32"},
		&printer{"Uint32s", "[]uint32"},
		&printer{"Bools", "[]bool"},
		&printer{"Strings", "[]string"},
		&printer{"ListOfBytes", "[][]byte"},
	}, `"fmt"`)
	gen(lengthStr, "length.gen.go", []interface{}{
		"Float64s",
		"Float32s",
		"Int64s",
		"Uint64s",
		"Int32s",
		"Uint32s",
		"Bools",
		"Strings",
		"ListOfBytes",
		"String",
		"Bytes",
	})
	gen(elemStr, "elem.gen.go", []interface{}{
		&elemer{"Float64s", "float64", "Float64"},
		&elemer{"Float32s", "float32", "Float32"},
		&elemer{"Int64s", "int64", "Int64"},
		&elemer{"Uint64s", "uint64", "Uint64"},
		&elemer{"Int32s", "int32", "Int32"},
		&elemer{"Uint32s", "uint32", "Uint32"},
		&elemer{"Bools", "bool", "Bool"},
		&elemer{"Strings", "string", "String"},
		&elemer{"ListOfBytes", "[]byte", "Bytes"},
	})
	gen(rangeStr, "range.gen.go", []interface{}{
		&ranger{"Float64s", "[]float64"},
		&ranger{"Float32s", "[]float32"},
		&ranger{"Int64s", "[]int64"},
		&ranger{"Uint64s", "[]uint64"},
		&ranger{"Int32s", "[]int32"},
		&ranger{"Uint32s", "[]uint32"},
		&ranger{"Bools", "[]bool"},
		&ranger{"Strings", "[]string"},
		&ranger{"ListOfBytes", "[][]byte"},
	})
}

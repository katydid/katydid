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
	"github.com/katydid/katydid/expr/funcs"
	"github.com/katydid/katydid/serialize"
	"reflect"
)

type Bool interface {
	Eval(serialize.Decoder) (bool, error)
}

type Decoder interface {
	SetDecoder(serialize.Decoder)
}

type composedBool struct {
	Decoders []Decoder
	Func     funcs.Bool
}

type errInit struct {
	i   interface{}
	err error
}

func (this *errInit) Error() string {
	return fmt.Sprintf("%#v err: %s", this.i, this.err)
}

var (
	varTyp     = reflect.TypeOf((*funcs.Variable)(nil)).Elem()
	decoderTyp = reflect.TypeOf((*funcs.Decoder)(nil)).Elem()
	constTyp   = reflect.TypeOf((*funcs.Const)(nil)).Elem()
	initTyp    = reflect.TypeOf((*funcs.Init)(nil)).Elem()
	listOfTyp  = reflect.TypeOf((*funcs.ListOf)(nil)).Elem()
)

var (
	debug = true
)

func NewBool(expr *expr.Expr) (funcs.Bool, error) {
	expr2, err := rewriteBuiltIn(expr)
	if err != nil {
		return nil, err
	}
	return composeBool(expr2)
}

func NewBoolFunc(f funcs.Bool) (*composedBool, error) {
	e, err := TrimBool(f)
	if err != nil {
		return nil, err
	}
	initImpls := FuncImplements(e, initTyp)
	for _, i := range initImpls {
		err := i.(funcs.Init).Init()
		if err != nil {
			return nil, &errInit{i, err}
		}
	}
	impls := FuncImplements(e, decoderTyp)
	decs := make([]Decoder, len(impls))
	for i := range impls {
		decs[i] = impls[i].(Decoder)
	}
	return &composedBool{decs, e}, nil
}

func (this *composedBool) Eval(dec serialize.Decoder) (bool, error) {
	for _, v := range this.Decoders {
		v.SetDecoder(dec)
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
		if e.Field(i).Elem().Type().Implements(listOfTyp) {
			is = append(is, ListImplements(e.Field(i).Interface(), typ)...)
		} else {
			is = append(is, FuncImplements(e.Field(i).Interface(), typ)...)
		}
	}
	if reflect.ValueOf(i).Type().Implements(typ) {
		is = append(is, i)
	}
	return is
}

func ListImplements(i interface{}, typ reflect.Type) []interface{} {
	e := reflect.ValueOf(i).Elem()
	var is []interface{}
	list := e.Field(0)
	for i := 0; i < list.Len(); i++ {
		is = append(is, FuncImplements(list.Index(i).Interface(), typ)...)
	}
	return is
}

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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/funcs"
	"reflect"
)

//Bool is an interface that represents a function, that given a value for a variable returns a boolean or an error.
type Bool interface {
	Eval(parser.Value) (bool, error)
}

//Setter is an interface that represents a variable in a function of which the value can be set.
type Setter interface {
	SetValue(parser.Value)
}

type composedBool struct {
	Setters []Setter
	Func    funcs.Bool
}

type errInit struct {
	i   interface{}
	err error
}

func (this *errInit) Error() string {
	return fmt.Sprintf("%#v err: %s", this.i, this.err)
}

var (
	varTyp        = reflect.TypeOf((*funcs.Variable)(nil)).Elem()
	setterTyp     = reflect.TypeOf((*funcs.Setter)(nil)).Elem()
	constTyp      = reflect.TypeOf((*funcs.Const)(nil)).Elem()
	initTyp       = reflect.TypeOf((*funcs.Init)(nil)).Elem()
	listOfTyp     = reflect.TypeOf((*funcs.ListOf)(nil)).Elem()
	setContextTyp = reflect.TypeOf((*funcs.SetContext)(nil)).Elem()
)

//NewBool constructs a boolean function from a parsed expression.
func NewBool(expr *ast.Expr) (funcs.Bool, error) {
	expr2, err := ConvertBuiltInIntoFunction(expr)
	if err != nil {
		return nil, err
	}
	return composeBool(expr2)
}

func SetContext(f funcs.Bool, context *funcs.Context) {
	setContextImpls := FuncImplements(f, setContextTyp)
	for _, i := range setContextImpls {
		i.(funcs.SetContext).SetContext(context)
	}
}

//NewBoolFunc returns the same function that it was given, but that has been trimmed and which is ready for variable values and evaluation.
func NewBoolFunc(f funcs.Bool) (Bool, error) {
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
	impls := FuncImplements(e, setterTyp)
	setters := make([]Setter, len(impls))
	for i := range impls {
		setters[i] = impls[i].(Setter)
	}
	return &composedBool{setters, e}, nil
}

//Eval evaluates the function given a value.
func (this *composedBool) Eval(val parser.Value) (bool, error) {
	for _, s := range this.Setters {
		s.SetValue(val)
	}
	return this.Func.Eval()
}

//FuncImplements returns all the functions in the function tree that implements the provided type.
func FuncImplements(i interface{}, typ reflect.Type) []interface{} {
	e := reflect.ValueOf(i).Elem()
	var is []interface{}
	for i := 0; i < e.NumField(); i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if e.Field(i).Elem().Type().Implements(listOfTyp) {
			is = append(is, listImplements(e.Field(i).Interface(), typ)...)
		} else {
			is = append(is, FuncImplements(e.Field(i).Interface(), typ)...)
		}
	}
	if reflect.ValueOf(i).Type().Implements(typ) {
		is = append(is, i)
	}
	return is
}

func listImplements(i interface{}, typ reflect.Type) []interface{} {
	e := reflect.ValueOf(i).Elem()
	var is []interface{}
	list := e.Field(0)
	for i := 0; i < list.Len(); i++ {
		is = append(is, FuncImplements(list.Index(i).Interface(), typ)...)
	}
	return is
}

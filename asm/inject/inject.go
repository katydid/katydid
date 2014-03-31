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

package inject

import (
	"github.com/awalterschulze/katydid/asm/compiler"
	"github.com/awalterschulze/katydid/asm/exec"
	"github.com/awalterschulze/katydid/asm/ifexpr"
	"reflect"
)

//Returns all the function structs instances that implements the given interface
func Implements(exec *exec.Exec, typ reflect.Type) []interface{} {
	ifs := exec.Link.(compiler.Link).GetIfs()
	var is []interface{}
	for _, iffy := range ifs {
		is = append(is, stateExprImplements(iffy, typ)...)
	}
	return is
}

func stateExprImplements(stateExpr ifexpr.StateExpr, typ reflect.Type) []interface{} {
	var is []interface{}
	ifExpr, ok := stateExpr.(*ifexpr.IfExpr)
	if !ok {
		return nil
	}
	is = append(is, funcImplements(ifExpr.Cond, typ)...)
	is = append(is, stateExprImplements(ifExpr.Succ, typ)...)
	is = append(is, stateExprImplements(ifExpr.Fail, typ)...)
	return is
}

func funcImplements(i interface{}, typ reflect.Type) []interface{} {
	e := reflect.ValueOf(i).Elem()
	var is []interface{}
	for i := 0; i < e.NumField(); i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		is = append(is, funcImplements(e.Field(i).Interface(), typ)...)
	}
	if reflect.ValueOf(i).Type().Implements(typ) {
		is = append(is, i)
	}
	return is
}

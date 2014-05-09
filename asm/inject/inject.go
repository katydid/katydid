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
	"github.com/awalterschulze/katydid/asm/compose"
	"github.com/awalterschulze/katydid/asm/exec"
	"github.com/awalterschulze/katydid/asm/ifexpr"
	"reflect"
)

//Returns all the function structs instances that implements the given interface
func Implements(exec *exec.Exec, typ reflect.Type) []interface{} {
	ifs := exec.Link.(compiler.Link).GetIfs()
	var is []interface{}
	for _, iffy := range ifs {
		is = append(is, StateExprImplements(iffy, typ)...)
	}
	return is
}

func StateExprImplements(stateExpr ifexpr.StateExpr, typ reflect.Type) []interface{} {
	var is []interface{}
	ifExpr, ok := stateExpr.(*ifexpr.IfExpr)
	if !ok {
		return nil
	}
	is = append(is, compose.FuncImplements(ifExpr.Cond, typ)...)
	is = append(is, StateExprImplements(ifExpr.Succ, typ)...)
	is = append(is, StateExprImplements(ifExpr.Fail, typ)...)
	return is
}

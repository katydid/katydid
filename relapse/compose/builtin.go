//  Copyright 2015 Walter Schulze
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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/types"
)

//ConvertBuiltInIntoFunction converts a BuiltIn Expr into a Function Expr.
func ConvertBuiltInIntoFunction(e *ast.Expr) (*ast.Expr, error) {
	if e.BuiltIn == nil {
		return e, nil
	}
	s := e.GetBuiltIn().GetSymbol().GetValue()
	right := e.GetBuiltIn().GetExpr()
	typ, err := Which(right)
	if err != nil {
		return nil, err
	}
	if types.IsList(typ) {
		typ = types.ListToSingle(typ)
	}
	left := ast.NewVar(typ)
	funcName := ast.BuiltInFunctionName(s)
	e2 := ast.NewNestedFunction(funcName, left, right)
	if funcName == "regex" {
		e2 = ast.NewNestedFunction(funcName, right, ast.NewVar(types.SINGLE_STRING))
	} else if funcName == "type" {
		e2 = ast.NewNestedFunction(funcName, right)
	}
	return e2, nil
}

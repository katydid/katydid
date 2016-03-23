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
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/types"
)

func symbolToFunc(symbol string) string {
	switch symbol {
	case "==":
		return "eq"
	case "!=":
		return "ne"
	case "<":
		return "lt"
	case ">":
		return "gt"
	case "<=":
		return "le"
	case ">=":
		return "ge"
	case "~=":
		return "regex"
	case "*=":
		return "contains"
	case "^=":
		return "hasPrefix"
	case "$=":
		return "hasSuffix"
	case "::":
		return "type"
	}
	panic("unknown builtin symbol `" + symbol + "`")
}

func rewriteBuiltIn(e *expr.Expr) (*expr.Expr, error) {
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
	left := expr.NewVar(typ)
	funcName := symbolToFunc(s)
	e2 := expr.NewNestedFunction(funcName, left, right)
	if funcName == "regex" {
		e2 = expr.NewNestedFunction(funcName, right, expr.NewVar(types.SINGLE_STRING))
	} else if funcName == "type" {
		e2 = expr.NewNestedFunction(funcName, right)
	}
	return e2, nil
}

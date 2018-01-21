//  Copyright 2018 Walter Schulze
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

package combinator

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/types"
)

//Value represents a field value.
func Value(expr *ast.Expr) *ast.Pattern {
	return ast.NewLeafNode(expr)
}

func Eq(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("eq", left, right)
}

func StringVar() *ast.Expr {
	return ast.NewVar(types.SINGLE_STRING)
}

func StringConst(s string) *ast.Expr {
	return ast.NewStringConst(s)
}

func And(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("and", left, right)
}

func Contains(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("contains", left, right)
}

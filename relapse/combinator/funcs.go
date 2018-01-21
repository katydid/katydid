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
	v := ast.NewVar(types.SINGLE_STRING)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
}

func StringConst(s string) *ast.Expr {
	return ast.NewStringConst(s)
}

func BoolVar() *ast.Expr {
	v := ast.NewVar(types.SINGLE_BOOL)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
}

func BoolConst(b bool) *ast.Expr {
	if b {
		return ast.NewTrue()
	}
	return ast.NewFalse()
}

func IntVar() *ast.Expr {
	v := ast.NewVar(types.SINGLE_INT)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
}

func IntConst(i int64) *ast.Expr {
	return ast.NewIntConst(i)
}

func UintVar() *ast.Expr {
	v := ast.NewVar(types.SINGLE_UINT)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
}

func UintConst(u uint64) *ast.Expr {
	return ast.NewUintConst(u)
}

func DoubleVar() *ast.Expr {
	v := ast.NewVar(types.SINGLE_DOUBLE)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
}

func DoubleConst(d float64) *ast.Expr {
	return ast.NewDoubleConst(d)
}

func And(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("and", left, right)
}

func Contains(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("contains", left, right)
}

func Length(e *ast.Expr) *ast.Expr {
	return ast.NewFunction("length", e)
}

func Not(e *ast.Expr) *ast.Expr {
	return ast.NewFunction("not", e)
}

func Or(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("or", left, right)
}

func LT(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("lt", left, right)
}

func LE(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("le", left, right)
}

func GT(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("gt", left, right)
}

func GE(left, right *ast.Expr) *ast.Expr {
	return ast.NewFunction("ge", left, right)
}

func Regex(expr, input *ast.Expr) *ast.Expr {
	return ast.NewFunction("regex", expr, input)
}

func Type(expr *ast.Expr) *ast.Expr {
	return ast.NewFunction("type", expr)
}

func StringsConst(ss []string) *ast.Expr {
	es := make([]*ast.Expr, len(ss))
	for i := range ss {
		es[i] = StringConst(ss[i])
	}
	return ast.NewStringList(es...)
}

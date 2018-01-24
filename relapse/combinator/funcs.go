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
	"github.com/katydid/katydid/relapse/parser"
	"github.com/katydid/katydid/relapse/types"
)

//Value represents a field value.
func Value(expr *ast.Expr) *ast.Pattern {
	if expr.Function != nil && len(expr.Function.Params) == 2 {
		constructor := ast.FunctionNameToBuiltIn(expr.Function.Name)
		if constructor != nil {
			p1 := expr.Function.Params[0]
			p2 := expr.Function.Params[1]
			if p1.Terminal != nil && p1.Terminal.Variable != nil &&
				p2.Terminal != nil && p2.Terminal.Variable == nil {
				p2.Comma = nil
				expr = constructor(p2)
			} else if p2.Terminal != nil && p2.Terminal.Variable != nil &&
				p1.Terminal != nil && p1.Terminal.Variable == nil {
				p1.Comma = nil
				expr = constructor(p1)
			}
		}
	}
	e, err := parser.ParseExpr(expr.String())
	if err != nil {
		panic(err)
	}
	return ast.NewLeafNode(e)
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
	c := ast.NewStringConst(s)
	c.RightArrow = &ast.Keyword{Value: "->"}
	return c
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

func BytesConst(b []byte) *ast.Expr {
	return ast.NewBytesConst(b)
}

func BytesVar() *ast.Expr {
	v := ast.NewVar(types.SINGLE_BYTES)
	v.RightArrow = &ast.Keyword{Value: "->"}
	return v
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

func StringsConst(is []string) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = StringConst(is[i])
	}
	return ast.NewStringList(es...)
}

func IntsConst(is []int64) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = IntConst(is[i])
	}
	return ast.NewIntList(es...)
}

func UintsConst(is []uint64) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = UintConst(is[i])
	}
	return ast.NewUintList(es...)
}

func DoublesConst(is []float64) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = DoubleConst(is[i])
	}
	return ast.NewDoubleList(es...)
}

func BoolsConst(is []bool) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = BoolConst(is[i])
	}
	return ast.NewBoolList(es...)
}

func ListOfBytesConst(is [][]byte) *ast.Expr {
	es := make([]*ast.Expr, len(is))
	for i := range is {
		es[i] = BytesConst(is[i])
	}
	return ast.NewBytesList(es...)
}

func EqualFold(s, t *ast.Expr) *ast.Expr {
	return ast.NewFunction("eqFold", s, t)
}

func HasPrefix(a, b *ast.Expr) *ast.Expr {
	return ast.NewFunction("hasPrefix", a, b)
}

func HasSuffix(a, b *ast.Expr) *ast.Expr {
	return ast.NewFunction("hasSuffix", a, b)
}

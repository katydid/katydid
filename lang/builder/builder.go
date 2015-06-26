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

package builder

import (
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/lang/ast"
)

func Any() *lang.Pattern {
	return lang.NewNot(lang.NewEmptySet())
}

func concat(p *lang.Pattern, ps ...*lang.Pattern) *lang.Pattern {
	if len(ps) == 0 {
		return p
	}
	return lang.NewConcat(p, concat(ps[0], ps[1:]...))
}

func MatchIn(nameStr string, child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	name, err := exprparser.NewParser().ParseExpr(nameStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(name, concat(child, children...))
}

func None() *lang.Pattern {
	return lang.NewEmptySet()
}

func Eval(name string) *lang.Pattern {
	return lang.NewReference(name)
}

func MatchTree(child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	return concat(child, children...)
}

func MatchField(nameStr string, exprStr string) *lang.Pattern {
	name, err := exprparser.NewParser().ParseExpr(nameStr)
	if err != nil {
		panic(err)
	}
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(name, lang.NewLeafNode(expr))
}

func And(left, right *lang.Pattern) *lang.Pattern {
	return lang.NewAnd(left, right)
}

func Or(left, right *lang.Pattern) *lang.Pattern {
	return lang.NewOr(left, right)
}

func ZeroOrMore(p *lang.Pattern) *lang.Pattern {
	return lang.NewZeroOrMore(p)
}

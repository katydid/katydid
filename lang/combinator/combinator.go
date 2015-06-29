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

type G map[string]*lang.Pattern

func (g G) Grammar() *lang.Grammar {
	return lang.NewGrammar(g)
}

func Any() *lang.Pattern {
	return lang.NewNot(lang.NewEmptySet())
}

func concat(p *lang.Pattern, ps ...*lang.Pattern) *lang.Pattern {
	if len(ps) == 0 {
		return p
	}
	return lang.NewConcat(p, concat(ps[0], ps[1:]...))
}

func MatchIn(name string, child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	return lang.NewTreeNode(lang.NewName(name), concat(child, children...))
}

func MatchInAny(child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	return lang.NewTreeNode(lang.NewAnyName(), concat(child, children...))
}

func MatchInAnyExcept(name string, child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	return lang.NewTreeNode(lang.NewAnyNameExcept(lang.NewName(name)), concat(child, children...))
}

func nameChoice(p string, ps ...string) *lang.NameExpr {
	if len(ps) == 0 {
		return lang.NewName(p)
	}
	return lang.NewNameChoice(lang.NewName(p), nameChoice(ps[0], ps[1:]...))
}

func MatchInAnyOf(names []string, child *lang.Pattern, children ...*lang.Pattern) *lang.Pattern {
	if len(names) < 2 {
		panic("less than two names is not really a choice, is it?")
	}
	return lang.NewTreeNode(nameChoice(names[0], names[1:]...), concat(child, children...))
}

func MatchField(name string, exprStr string) *lang.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(lang.NewName(name), lang.NewLeafNode(expr))
}

func MatchAnyField(exprStr string) *lang.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(lang.NewAnyName(), lang.NewLeafNode(expr))
}

func MatchAnyFieldExcept(name string, exprStr string) *lang.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(lang.NewAnyNameExcept(lang.NewName(name)), lang.NewLeafNode(expr))
}

func MatchAnyFields(names []string, exprStr string) *lang.Pattern {
	if len(names) < 2 {
		panic("less than two fieldnames is not really a choice, is it?")
	}
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(nameChoice(names[0], names[1:]...), lang.NewLeafNode(expr))
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

func And(left, right *lang.Pattern) *lang.Pattern {
	return lang.NewAnd(left, right)
}

func Or(left, right *lang.Pattern) *lang.Pattern {
	return lang.NewOr(left, right)
}

func Many(p *lang.Pattern) *lang.Pattern {
	return lang.NewZeroOrMore(p)
}

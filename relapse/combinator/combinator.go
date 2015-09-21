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

package combinator

import (
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/relapse/ast"
)

type G map[string]*relapse.Pattern

func (g G) Grammar() *relapse.Grammar {
	return relapse.NewGrammar(g)
}

func Any() *relapse.Pattern {
	return relapse.NewNot(relapse.NewEmptySet())
}

func concat(p *relapse.Pattern, ps ...*relapse.Pattern) *relapse.Pattern {
	if len(ps) == 0 {
		return p
	}
	pss := append([]*relapse.Pattern{p}, ps...)
	return relapse.NewConcat(pss...)
}

func MatchIn(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewName(name), concat(child, children...))
}

func MatchInAny(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewAnyName(), concat(child, children...))
}

func MatchInAnyExcept(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewAnyNameExcept(relapse.NewName(name)), concat(child, children...))
}

func nameChoice(p string, ps ...string) *relapse.NameExpr {
	if len(ps) == 0 {
		return relapse.NewName(p)
	}
	return relapse.NewNameChoice(relapse.NewName(p), nameChoice(ps[0], ps[1:]...))
}

func MatchInAnyOf(names []string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	if len(names) < 2 {
		panic("less than two names is not really a choice, is it?")
	}
	return relapse.NewTreeNode(nameChoice(names[0], names[1:]...), concat(child, children...))
}

func MatchField(name string, exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewTreeNode(relapse.NewName(name), relapse.NewLeafNode(expr))
}

func MatchAnyField(exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewTreeNode(relapse.NewAnyName(), relapse.NewLeafNode(expr))
}

func MatchAnyFieldExcept(name string, exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewTreeNode(relapse.NewAnyNameExcept(relapse.NewName(name)), relapse.NewLeafNode(expr))
}

func MatchAnyFields(names []string, exprStr string) *relapse.Pattern {
	if len(names) < 2 {
		panic("less than two fieldnames is not really a choice, is it?")
	}
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewTreeNode(nameChoice(names[0], names[1:]...), relapse.NewLeafNode(expr))
}

func None() *relapse.Pattern {
	return relapse.NewEmptySet()
}

func Eval(name string) *relapse.Pattern {
	return relapse.NewReference(name)
}

func MatchInOrder(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return concat(child, children...)
}

func AllOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewAnd(patterns...)
}

func AnyOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewOr(patterns...)
}

func OppositeOf(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewNot(p)
}

func Many(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewZeroOrMore(p)
}

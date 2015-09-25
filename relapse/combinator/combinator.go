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
	return relapse.NewZAny()
}

func concat(p *relapse.Pattern, ps ...*relapse.Pattern) *relapse.Pattern {
	if len(ps) == 0 {
		return p
	}
	pss := append([]*relapse.Pattern{p}, ps...)
	return relapse.NewConcat(pss...)
}

func InPath(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewWithSomeTreeNode(relapse.NewName(name), concat(child, children...))
}

func InFieldPath(name string, exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewWithSomeLeafNode(relapse.NewName(name), expr)
}

func In(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewName(name), concat(child, children...))
}

func InAny(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewAnyName(), concat(child, children...))
}

func InAnyExcept(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewAnyNameExcept(relapse.NewName(name)), concat(child, children...))
}

func nameChoice(p string, ps ...string) *relapse.NameExpr {
	if len(ps) == 0 {
		return relapse.NewName(p)
	}
	return relapse.NewNameChoice(relapse.NewName(p), nameChoice(ps[0], ps[1:]...))
}

func InAnyOf(names []string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	if len(names) < 2 {
		panic("less than two names is not really a choice, is it?")
	}
	return relapse.NewTreeNode(nameChoice(names[0], names[1:]...), concat(child, children...))
}

func Field(name string, exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewLeafNode(relapse.NewName(name), expr)
}

func AnyField(exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewLeafNode(relapse.NewAnyName(), expr)
}

func AnyFieldExcept(name string, exprStr string) *relapse.Pattern {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewLeafNode(relapse.NewAnyNameExcept(relapse.NewName(name)), expr)
}

func AnyFields(names []string, exprStr string) *relapse.Pattern {
	if len(names) < 2 {
		panic("less than two fieldnames is not really a choice, is it?")
	}
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewLeafNode(nameChoice(names[0], names[1:]...), expr)
}

func None() *relapse.Pattern {
	return relapse.NewEmptySet()
}

func Eval(name string) *relapse.Pattern {
	return relapse.NewReference(name)
}

func InOrder(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return concat(child, children...)
}

func ExactAllOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewAnd(patterns...)
}

func MatchAllOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewWithSomeAnd(patterns...)
}

func ExactAnyOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewOr(patterns...)
}

func MatchAnyOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewWithSomeOr(patterns...)
}

func OppositeOf(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewNot(p)
}

func Many(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewZeroOrMore(p)
}

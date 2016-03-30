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

//Package combinator provides a user friendly way of constructing a relapse abstract syntax tree.
package combinator

import (
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/funcs"
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/relapse"
)

//G represents the relapse Grammar.
//This consists of a "main" map key with the main pattern value and any other references.
type G map[string]*relapse.Pattern

//Grammar returns G as a proper relapse.Grammar
func (g G) Grammar() *relapse.Grammar {
	return relapse.NewGrammar(g)
}

//Any represents a zero or more of anything pattern.
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

//InPath represents an ordered list of patterns in a field path.
func InPath(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewContains(relapse.NewTreeNode(expr.NewStringName(name), concat(child, children...)))
}

//InAnyPath represents an ordered list of patterns in any path.
func InAnyPath(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewContains(relapse.NewTreeNode(expr.NewAnyName(), concat(child, children...)))
}

//In represents an ordered list of patterns in a field.
func In(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(expr.NewStringName(name), concat(child, children...))
}

//Elem repesents an ordered list of patterns in an specific array element.
func Elem(index int, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(expr.NewIntName(int64(index)), concat(child, children...))
}

//InAny represents an ordered list of patterns in any field or index.
func InAny(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(expr.NewAnyName(), concat(child, children...))
}

//InAnyExcept represents an ordered list of patterns in any field except the specified one.
func InAnyExcept(name string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewTreeNode(expr.NewAnyNameExcept(expr.NewStringName(name)), concat(child, children...))
}

func nameChoice(p string, ps ...string) *expr.NameExpr {
	if len(ps) == 0 {
		return expr.NewStringName(p)
	}
	return expr.NewNameChoice(expr.NewStringName(p), nameChoice(ps[0], ps[1:]...))
}

//InAnyOf represents an ordered list of patterns in any of the specified fields.
func InAnyOf(names []string, child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	if len(names) < 2 {
		panic("less than two names is not really a choice, is it?")
	}
	return relapse.NewTreeNode(nameChoice(names[0], names[1:]...), concat(child, children...))
}

//Value represents a field value.
func Value(f funcs.Bool) *relapse.Pattern {
	exprStr := funcs.Sprint(f)
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return relapse.NewLeafNode(expr)
}

//None represents no possible match.
func None() *relapse.Pattern {
	return relapse.NewNot(relapse.NewZAny())
}

//Eval represents the evaluation of a reference name.
func Eval(name string) *relapse.Pattern {
	return relapse.NewReference(name)
}

//InOrder represents an ordered list of patterns.
func InOrder(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return concat(child, children...)
}

//AllOf represents an intersection of patterns.
func AllOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewAnd(patterns...)
}

//AnyOf represents a union of patterns.
func AnyOf(patterns ...*relapse.Pattern) *relapse.Pattern {
	return relapse.NewOr(patterns...)
}

//OppositeOf represents a compliment of a pattern.
func OppositeOf(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewNot(p)
}

//Many represents zero or more of the input pattern.
func Many(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewZeroOrMore(p)
}

//Maybe represents an optional pattern.
func Maybe(p *relapse.Pattern) *relapse.Pattern {
	return relapse.NewOptional(p)
}

func interleave(p *relapse.Pattern, ps ...*relapse.Pattern) *relapse.Pattern {
	if len(ps) == 0 {
		return p
	}
	pss := append([]*relapse.Pattern{p}, ps...)
	return relapse.NewInterleave(pss...)
}

//InAnyOrder represents interleaved patterns.
func InAnyOrder(child *relapse.Pattern, children ...*relapse.Pattern) *relapse.Pattern {
	return interleave(child, children...)
}

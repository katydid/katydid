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
//
//TODO: Review this API of this package, especially the Array methods.
package combinator

import (
	"github.com/katydid/katydid/relapse/ast"
)

//G represents the relapse Grammar.
//This consists of a "main" map key with the main pattern value and any other references.
type G map[string]*ast.Pattern

//Grammar returns G as a proper relapse.Grammar
func (g G) Grammar() *ast.Grammar {
	return ast.NewGrammar(g)
}

//Any represents a zero or more of anything pattern.
func Any() *ast.Pattern {
	return ast.NewZAny()
}

func concat(p *ast.Pattern, ps ...*ast.Pattern) *ast.Pattern {
	if len(ps) == 0 {
		return p
	}
	pss := append([]*ast.Pattern{p}, ps...)
	return ast.NewConcat(pss...)
}

//InPath represents an ordered list of patterns in a field path.
func InPath(name string, child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewContains(ast.NewTreeNode(ast.NewStringName(name), concat(child, children...)))
}

//InAnyPath represents an ordered list of patterns in any path.
func InAnyPath(child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewContains(ast.NewTreeNode(ast.NewAnyName(), concat(child, children...)))
}

//In represents an ordered list of patterns in a field.
func In(name string, child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewTreeNode(ast.NewStringName(name), concat(child, children...))
}

//Elem repesents an ordered list of patterns in an specific array element.
func Elem(index int, child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewTreeNode(ast.NewIntName(int64(index)), concat(child, children...))
}

//InAny represents an ordered list of patterns in any field or index.
func InAny(child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewTreeNode(ast.NewAnyName(), concat(child, children...))
}

//InAnyExcept represents an ordered list of patterns in any field except the specified one.
func InAnyExcept(name string, child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return ast.NewTreeNode(ast.NewAnyNameExcept(ast.NewStringName(name)), concat(child, children...))
}

func nameChoice(p string, ps ...string) *ast.NameExpr {
	if len(ps) == 0 {
		return ast.NewStringName(p)
	}
	return ast.NewNameChoice(ast.NewStringName(p), nameChoice(ps[0], ps[1:]...))
}

//InAnyOf represents an ordered list of patterns in any of the specified fields.
func InAnyOf(names []string, child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	if len(names) < 2 {
		panic("less than two names is not really a choice, is it?")
	}
	return ast.NewTreeNode(nameChoice(names[0], names[1:]...), concat(child, children...))
}

//None represents no possible match.
func None() *ast.Pattern {
	return ast.NewNot(ast.NewZAny())
}

//Eval represents the evaluation of a reference name.
func Eval(name string) *ast.Pattern {
	return ast.NewReference(name)
}

//InOrder represents an ordered list of patterns.
func InOrder(child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return concat(child, children...)
}

//AllOf represents an intersection of patterns.
func AllOf(patterns ...*ast.Pattern) *ast.Pattern {
	return ast.NewAnd(patterns...)
}

//AnyOf represents a union of patterns.
func AnyOf(patterns ...*ast.Pattern) *ast.Pattern {
	return ast.NewOr(patterns...)
}

//OppositeOf represents a compliment of a pattern.
func OppositeOf(p *ast.Pattern) *ast.Pattern {
	return ast.NewNot(p)
}

//Many represents zero or more of the input pattern.
func Many(p *ast.Pattern) *ast.Pattern {
	return ast.NewZeroOrMore(p)
}

//Maybe represents an optional pattern.
func Maybe(p *ast.Pattern) *ast.Pattern {
	return ast.NewOptional(p)
}

func interleave(p *ast.Pattern, ps ...*ast.Pattern) *ast.Pattern {
	if len(ps) == 0 {
		return p
	}
	pss := append([]*ast.Pattern{p}, ps...)
	return ast.NewInterleave(pss...)
}

//InAnyOrder represents interleaved patterns.
func InAnyOrder(child *ast.Pattern, children ...*ast.Pattern) *ast.Pattern {
	return interleave(child, children...)
}

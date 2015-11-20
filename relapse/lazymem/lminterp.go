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

package lazymem

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/serialize"
	"io"
)

func Interpret(g *relapse.Grammar, tree serialize.Parser) bool {
	fmt.Printf("BEFORE: %s\n", g)
	p := ConvertGrammar(g)
	res := p
	res = Simplify(res)
	err := tree.Next()
	for err == nil {
		res = deriv(res, tree)
		res = Simplify(res)
		err = tree.Next()
	}
	if err != io.EOF {
		panic(err)
	}
	fmt.Printf("CONVERTED: %s\n", ConvertPattern(res))
	return Nullable(res)
}

func derivNode(p *Node, tree serialize.Parser) *Pattern {
	matched, err := p.F.Eval(tree)
	if err != nil {
		return NewNot(NewZAny())
	}
	if !matched {
		return NewNot(NewZAny())
	}
	res := p.Pattern
	if !tree.IsLeaf() {
		tree.Down()
		err := tree.Next()
		for err == nil {
			res = deriv(res, tree)
			res = Simplify(res)
			err = tree.Next()
		}
		if err != io.EOF {
			panic(err)
		}
		tree.Up()
	}
	if !Nullable(res) {
		return NewNot(NewZAny())
	}
	return NewEmpty()
}

func deriv(p *Pattern, tree serialize.Parser) *Pattern {
	head := p.Head().GetValue()
	switch v := head.(type) {
	case *Empty:
		return NewNot(NewZAny())
	case *Node:
		return derivNode(v, tree)
	case *Concat:
		leftDeriv := NewConcat(deriv(v.Left, tree), v.Right)
		if Nullable(v.Left) {
			return NewOr(
				leftDeriv,
				deriv(v.Right, tree.Copy()),
			)
		} else {
			return leftDeriv
		}
	case *Or:
		treeCopy := tree.Copy()
		return NewOr(
			deriv(v.Left, tree),
			deriv(v.Right, treeCopy),
		)
	case *And:
		treeCopy := tree.Copy()
		return NewAnd(
			deriv(v.Left, tree),
			deriv(v.Right, treeCopy),
		)
	case *ZeroOrMore:
		return NewConcat(
			deriv(v.Pattern, tree),
			NewZeroOrMore(v.Pattern),
		)
	case *Not:
		return NewNot(deriv(v.Pattern, tree))
	case *ZAny:
		return p
	case *Contains:
		return deriv(NewConcat(NewZAny(), NewConcat(v.Pattern, NewZAny())), tree)
	case *Optional:
		return deriv(NewOr(NewEmpty(), v.Pattern), tree)
	case *Interleave:
		treeCopy := tree.Copy()
		return NewOr(
			NewInterleave(deriv(v.Left, tree), v.Right),
			NewInterleave(v.Left, deriv(v.Right, treeCopy)),
		)
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", head))
	}
}

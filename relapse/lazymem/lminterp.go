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
	"strconv"
	"sync"
)

var mem Memoize = nil

type Memoize map[*Pattern]map[serialize.Parser]*Pattern

func (this Memoize) Get(p *Pattern, tree serialize.Parser) *Pattern {
	t := this[p]
	if t == nil {
		return nil
	}
	return t[tree]
}

func (this Memoize) Add(p *Pattern, tree serialize.Parser, d *Pattern) {
	t := this[p]
	if t == nil {
		this[p] = make(map[serialize.Parser]*Pattern)
	}
	this[p][tree] = d
}

func Interpret(g *relapse.Grammar, tree serialize.Parser) bool {
	mem = make(Memoize)
	fmt.Printf("BEFORE: %s\n", g)
	p := ConvertGrammar(g)
	res := p
	res = Simplify(res)
	err := tree.Next()
	for err == nil {
		res = deriv(res, tree.Copy())
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
	tree = tree.Copy()
	res := p.Pattern
	if !tree.IsLeaf() {
		tree.Down()
		err := tree.Next()
		for err == nil {
			res = deriv(res, tree.Copy())
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

var i int
var s sync.Mutex

func newName() string {
	s.Lock()
	defer s.Unlock()
	i++
	return "int" + strconv.Itoa(i)
}

func lderiv(p *Pattern, tree serialize.Parser) *Pattern {
	m := mem.Get(p, tree)
	if m != nil {
		return m
	}
	if p.thunk == nil {
		fmt.Printf("deriv %s\n", p.String())
		res := deriv(p, tree)
		mem.Add(p, tree, res)
		return res
	}
	fmt.Printf("lazy %s\n", p.String())
	res := &Pattern{
		name: newName(),
		thunk: func() *Pattern {
			return deriv(p, tree)
		},
	}
	mem.Add(p, tree, res)
	return res
}

func deriv(p *Pattern, tree serialize.Parser) *Pattern {
	head := p.Head().GetValue()
	switch v := head.(type) {
	case *Empty:
		return NewNot(NewZAny())
	case *Node:
		return derivNode(v, tree)
	case *Concat:
		leftDeriv := NewConcat(lderiv(v.Left, tree), v.Right)
		if Nullable(v.Left) {
			return NewOr(
				leftDeriv,
				lderiv(v.Right, tree),
			)
		} else {
			return leftDeriv
		}
	case *Or:
		left := lderiv(v.Left, tree)
		right := lderiv(v.Right, tree)
		return NewOr(left, right)
	case *And:
		left := lderiv(v.Left, tree)
		right := lderiv(v.Right, tree)
		return NewAnd(left, right)
	case *ZeroOrMore:
		return NewConcat(
			lderiv(v.Pattern, tree),
			NewZeroOrMore(v.Pattern),
		)
	case *Not:
		return NewNot(lderiv(v.Pattern, tree))
	case *ZAny:
		return p
	case *Contains:
		return lderiv(NewConcat(NewZAny(), NewConcat(v.Pattern, NewZAny())), tree)
	case *Optional:
		return lderiv(NewOr(NewEmpty(), v.Pattern), tree)
	case *Interleave:
		left := lderiv(v.Left, tree)
		right := lderiv(v.Right, tree)
		return NewOr(
			NewInterleave(left, v.Right),
			NewInterleave(v.Left, right),
		)
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", head))
	}
}

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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"io"
	"strconv"
)

type memoize map[*Pattern]map[parser.Interface]*Pattern

func (this memoize) Get(p *Pattern, tree parser.Interface) *Pattern {
	t := this[p]
	if t == nil {
		return nil
	}
	return t[tree]
}

func (this memoize) Add(p *Pattern, tree parser.Interface, d *Pattern) {
	t := this[p]
	if t == nil {
		this[p] = make(map[parser.Interface]*Pattern)
	}
	this[p][tree] = d
}

type Interpreter struct {
	p *Pattern
}

func (this *Interpreter) Interpret(tree parser.Interface) bool {
	interpret := &interpreter{make(memoize), 0}
	res := this.p
	err := tree.Next()
	for err == nil {
		res = interpret.deriv(res, tree /*TODO remove .Copy()*/)
		res = Simplify(res)
		head := res.Head()
		if head.ZAny != nil {
			return true
		}
		if head.Not != nil && head.Not.Pattern.Head().ZAny != nil {
			return false
		}
		err = tree.Next()
	}
	if err != io.EOF {
		panic(err)
	}
	return Nullable(res)
}

func NewInterpreter(g *ast.Grammar) *Interpreter {
	p := ConvertGrammar(g)
	return &Interpreter{Simplify(p)}
}

func Interpret(g *ast.Grammar, tree parser.Interface) bool {
	return NewInterpreter(g).Interpret(tree)
}

type interpreter struct {
	mem    memoize
	refNum int
}

func (this *interpreter) newName() string {
	this.refNum++
	return "ref" + strconv.Itoa(this.refNum)
}

func (this *interpreter) derivNode(p *Node, tree parser.Interface) *Pattern {
	matched, err := p.F.Eval(tree)
	if err != nil {
		return NewNot(NewZAny())
	}
	if !matched {
		return NewNot(NewZAny())
	}
	tree = tree /* TODO remove .Copy()*/
	res := p.Pattern
	head := res.Head()
	if head.ZAny != nil {
		return NewEmpty()
	}
	if head.Not != nil && head.Not.Pattern.Head().ZAny != nil {
		return NewNot(NewZAny())
	}
	if !tree.IsLeaf() {
		tree.Down()
		err := tree.Next()
		for err == nil {
			res = this.deriv(res, tree /* TODO remove .Copy()*/)
			res = Simplify(res)
			head := res.Head()
			if head.ZAny != nil {
				break
			}
			if head.Not != nil && head.Not.Pattern.Head().ZAny != nil {
				break
			}
			err = tree.Next()
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		tree.Up()
	}
	if !Nullable(res) {
		return NewNot(NewZAny())
	}
	return NewEmpty()
}

func (this *interpreter) lderiv(p *Pattern, tree parser.Interface) *Pattern {
	m := this.mem.Get(p, tree)
	if m != nil {
		return m
	}
	if p.thunk == nil {
		//fmt.Printf("deriv %s\n", p.String())
		res := this.deriv(p, tree)
		this.mem.Add(p, tree, res)
		return res
	}
	//fmt.Printf("lazy %s\n", p.String())
	res := &Pattern{
		name: this.newName(),
		thunk: func() *Pattern {
			//fmt.Printf("-> evaluating lazy pattern %v\n", p)
			pp := this.deriv(p, tree)
			//fmt.Printf("<- evaluated lazy pattern %v\n", p)
			return pp
		},
	}
	this.mem.Add(p, tree, res)
	return res
}

func (this *interpreter) deriv(p *Pattern, tree parser.Interface) *Pattern {
	head := p.Head().GetValue()
	switch v := head.(type) {
	case *Empty:
		return NewNot(NewZAny())
	case *Node:
		return this.derivNode(v, tree)
	case *Concat:
		leftDeriv := NewConcat(this.lderiv(v.Left, tree), v.Right)
		if Nullable(v.Left) {
			return NewOr(
				leftDeriv,
				this.lderiv(v.Right, tree),
			)
		} else {
			return leftDeriv
		}
	case *Or:
		left := this.lderiv(v.Left, tree)
		right := this.lderiv(v.Right, tree)
		return NewOr(left, right)
	case *And:
		left := this.lderiv(v.Left, tree)
		right := this.lderiv(v.Right, tree)
		return NewAnd(left, right)
	case *ZeroOrMore:
		return NewConcat(
			this.lderiv(v.Pattern, tree),
			NewZeroOrMore(v.Pattern),
		)
	case *Not:
		return NewNot(this.lderiv(v.Pattern, tree))
	case *ZAny:
		return p
	case *Contains:
		return this.lderiv(NewConcat(NewZAny(), NewConcat(v.Pattern, NewZAny())), tree)
	case *Optional:
		return this.lderiv(NewOr(NewEmpty(), v.Pattern), tree)
	case *Interleave:
		left := this.lderiv(v.Left, tree)
		right := this.lderiv(v.Right, tree)
		return NewOr(
			NewInterleave(left, v.Right),
			NewInterleave(v.Left, right),
		)
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", head))
	}
}

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
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/nameexpr"
)

type Pattern struct {
	thunk    func() *PatternHead
	head     *PatternHead
	nullable *bool
}

func (this *Pattern) HasHead() bool {
	return this.head != nil
}

func (this *Pattern) Head() *PatternHead {
	if this.head == nil {
		this.head = this.thunk()
	}
	return this.head
}

func (this *Pattern) HasNullable() bool {
	return this.nullable != nil
}

func (this *Pattern) SetNullable(b bool) {
	this.nullable = &b
}

func (this *Pattern) GetNullable() bool {
	if !this.HasNullable() {
		return false
	}
	return *this.nullable
}

func NewLazyPattern(p *Pattern) *Pattern {
	return &Pattern{
		thunk: func() *PatternHead {
			return p.Head()
		},
	}
}

type PatternHead struct {
	Empty      *Empty
	Node       *Node
	Concat     *Concat
	Or         *Or
	And        *And
	ZeroOrMore *ZeroOrMore
	Not        *Not
	ZAny       *ZAny
	Contains   *Contains
	Optional   *Optional
	Interleave *Interleave
}

func (this *PatternHead) GetValue() interface{} {
	if this.Empty != nil {
		return this.Empty
	}
	if this.Node != nil {
		return this.Node
	}
	if this.Concat != nil {
		return this.Concat
	}
	if this.Or != nil {
		return this.Or
	}
	if this.And != nil {
		return this.And
	}
	if this.ZeroOrMore != nil {
		return this.ZeroOrMore
	}
	if this.Not != nil {
		return this.Not
	}
	if this.ZAny != nil {
		return this.ZAny
	}
	if this.Contains != nil {
		return this.Contains
	}
	if this.Optional != nil {
		return this.Optional
	}
	if this.Interleave != nil {
		return this.Interleave
	}
	panic("no pattern set")
}

type Empty struct{}

func NewEmpty() *Pattern {
	return &Pattern{
		head: &PatternHead{
			Empty: &Empty{},
		},
	}
}

type Node struct {
	F       compose.Bool
	Name    *relapse.NameExpr
	Expr    *expr.Expr
	Pattern *Pattern
}

func NewNode(f compose.Bool, n *relapse.NameExpr, e *expr.Expr, p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Node: &Node{
				F:       f,
				Name:    n,
				Expr:    e,
				Pattern: p,
			},
		},
	}
}

func NewTreeNode(name *relapse.NameExpr, p *Pattern) *Pattern {
	nameFunc := nameexpr.NameToFunc(name)
	f, err := compose.NewBoolFunc(nameFunc)
	if err != nil {
		panic(err)
	}
	return NewNode(f, name, nil, p)
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	f, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	return NewNode(f, nil, expr, NewEmpty())
}

type Concat struct {
	Left  *Pattern
	Right *Pattern
}

func NewConcat(l, r *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Concat: &Concat{
				Left:  l,
				Right: r,
			},
		},
	}
}

type Or struct {
	Left  *Pattern
	Right *Pattern
}

func NewOr(l, r *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Or: &Or{
				Left:  l,
				Right: r,
			},
		},
	}
}

type And struct {
	Left  *Pattern
	Right *Pattern
}

func NewAnd(l, r *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			And: &And{
				Left:  l,
				Right: r,
			},
		},
	}
}

type ZeroOrMore struct {
	Pattern *Pattern
}

func NewZeroOrMore(p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			ZeroOrMore: &ZeroOrMore{
				Pattern: p,
			},
		},
	}
}

type Not struct {
	Pattern *Pattern
}

func NewNot(p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Not: &Not{
				Pattern: p,
			},
		},
	}
}

type ZAny struct{}

func NewZAny() *Pattern {
	return &Pattern{
		head: &PatternHead{
			ZAny: &ZAny{},
		},
	}
}

type Contains struct {
	Pattern *Pattern
}

func NewContains(p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Contains: &Contains{
				Pattern: p,
			},
		},
	}
}

type Optional struct {
	Pattern *Pattern
}

func NewOptional(p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Optional: &Optional{
				Pattern: p,
			},
		},
	}
}

type Interleave struct {
	Left  *Pattern
	Right *Pattern
}

func NewInterleave(l, r *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			Interleave: &Interleave{
				Left:  l,
				Right: r,
			},
		},
	}
}

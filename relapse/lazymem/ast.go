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
	"github.com/katydid/katydid/relapse/ast"
)

type Pattern struct {
	thunk    func() *PatternHead
	head     *PatternHead
	nullable *bool
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

type PatternHead struct {
	Empty      *Empty
	EmptySet   *EmptySet
	TreeNode   *TreeNode
	LeafNode   *LeafNode
	Concat     *Concat
	Or         *Or
	And        *And
	ZeroOrMore *ZeroOrMore
	Not        *Not
}

func (this *PatternHead) GetValue() interface{} {
	if this.Empty != nil {
		return this.Empty
	}
	if this.EmptySet != nil {
		return this.EmptySet
	}
	if this.TreeNode != nil {
		return this.TreeNode
	}
	if this.LeafNode != nil {
		return this.LeafNode
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

type EmptySet struct{}

func NewEmptySet() *Pattern {
	return &Pattern{
		head: &PatternHead{
			EmptySet: &EmptySet{},
		},
	}
}

type TreeNode struct {
	Name    *relapse.NameExpr
	Pattern *Pattern
}

func NewTreeNode(name *relapse.NameExpr, p *Pattern) *Pattern {
	return &Pattern{
		head: &PatternHead{
			TreeNode: &TreeNode{
				Name:    name,
				Pattern: p,
			},
		},
	}
}

type LeafNode struct {
	Expr *expr.Expr
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	return &Pattern{
		head: &PatternHead{
			LeafNode: &LeafNode{
				Expr: expr,
			},
		},
	}
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

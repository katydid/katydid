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
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/compose"
	nameexpr "github.com/katydid/katydid/expr/name"
)

type Pattern struct {
	name     string
	thunk    func() *Pattern
	head     *PatternHead
	nullable *bool
}

func (this *Pattern) HasHead() bool {
	return this.head != nil
}

func (this *Pattern) Head() *PatternHead {
	if this.head == nil {
		this.head = this.thunk().Head()
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

func (this *Pattern) GoString() string {
	if this.thunk != nil {
		return fmt.Sprintf("&Pattern{thunk: %v}", this.name)
	}
	return fmt.Sprintf("&Pattern{head: %#v}", this.head)
}

func (this *Pattern) String() string {
	if this.thunk != nil {
		return fmt.Sprintf("@%v", this.name)
	}
	return this.head.String()
}

func NewLazyPattern(name string, p *Pattern) *Pattern {
	return &Pattern{
		name: name,
		thunk: func() *Pattern {
			return &Pattern{
				head: p.Head(),
			}
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

func (this *PatternHead) GoString() string {
	return fmt.Sprintf("%#v", this.GetValue())
}

func (this *PatternHead) String() string {
	return fmt.Sprintf("%v", this.GetValue())
}

type Empty struct{}

func NewEmpty() *Pattern {
	return &Pattern{
		head: &PatternHead{
			Empty: &Empty{},
		},
	}
}

func (this *Empty) GoString() string {
	return "&Empty{}"
}

func (this *Empty) String() string {
	return "<empty>"
}

type Node struct {
	F       compose.Bool
	Name    *expr.NameExpr
	Expr    *expr.Expr
	Pattern *Pattern
}

func NewNode(f compose.Bool, n *expr.NameExpr, e *expr.Expr, p *Pattern) *Pattern {
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

func (this *Node) GoString() string {
	if this.Name != nil {
		return fmt.Sprintf("&Node{Name: %#v, Pattern: %#v}", this.Name, this.Pattern)
	}
	return fmt.Sprintf("&Node{Expr: %#v, Pattern: %#v}", this.Expr, this.Pattern)
}

func (this *Node) String() string {
	if this.Name != nil {
		return this.Name.String() + ": " + this.Pattern.String()
	}
	return this.Expr.String()
}

func NewTreeNode(name *expr.NameExpr, p *Pattern) *Pattern {
	nameFunc := nameexpr.NameToFunc(name)
	f, err := compose.NewBoolFunc(nameFunc)
	if err != nil {
		panic(err)
	}
	return NewNode(f, name, nil, p)
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	b, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := compose.NewBoolFunc(b)
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

func (this *Concat) GoString() string {
	return fmt.Sprintf("&Concat{Left: %#v, Right: %#v}", this.Left, this.Right)
}

func (this *Concat) String() string {
	return fmt.Sprintf("[%v, %v]", this.Left, this.Right)
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

func (this *Or) GoString() string {
	return fmt.Sprintf("&Or{Left: %#v, Right: %#v}", this.Left, this.Right)
}

func (this *Or) String() string {
	return fmt.Sprintf("( %v | %v )", this.Left, this.Right)
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

func (this *And) GoString() string {
	return fmt.Sprintf("&And{Left: %#v, Right: %#v}", this.Left, this.Right)
}

func (this *And) String() string {
	return fmt.Sprintf("( %v & %v )", this.Left, this.Right)
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

func (this *ZeroOrMore) GoString() string {
	return fmt.Sprintf("&ZeroOrMore{Pattern: %#v}", this.Pattern)
}

func (this *ZeroOrMore) String() string {
	return fmt.Sprintf("(%v)*", this.Pattern)
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

func (this *Not) GoString() string {
	return fmt.Sprintf("&Not{Pattern: %#v}", this.Pattern)
}

func (this *Not) String() string {
	return fmt.Sprintf("!(%v)", this.Pattern)
}

type ZAny struct{}

func NewZAny() *Pattern {
	return &Pattern{
		head: &PatternHead{
			ZAny: &ZAny{},
		},
	}
}

func (this *ZAny) GoString() string {
	return "&ZAny{}"
}

func (this *ZAny) String() string {
	return "*"
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

func (this *Contains) GoString() string {
	return fmt.Sprintf("&Contains{Pattern: %#v}", this.Pattern)
}

func (this *Contains) String() string {
	return fmt.Sprintf(".%v", this.Pattern)
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

func (this *Optional) GoString() string {
	return fmt.Sprintf("&Optional{Pattern: %#v}", this.Pattern)
}

func (this *Optional) String() string {
	return fmt.Sprintf("(%v)?", this.Pattern)
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

func (this *Interleave) GoString() string {
	return fmt.Sprintf("&Interleave{Left: %#v, Right: %#v}", this.Left, this.Right)
}

func (this *Interleave) String() string {
	return fmt.Sprintf("{ %#v ; %#v}", this.Left, this.Right)
}

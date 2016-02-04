//  Copyright 2016 Walter Schulze
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

package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/nameexpr"
)

type patternType int

var (
	EmptyType      patternType = 1
	TreeNodeType   patternType = 3
	LeafNodeType   patternType = 4
	ConcatType     patternType = 5
	OrType         patternType = 6
	AndType        patternType = 7
	ZeroOrMoreType patternType = 8
	ReferenceType  patternType = 9
	NotType        patternType = 10
	ZAnyType       patternType = 11
	// ContainsType   patternType = 12
	// OptionalType   patternType = 13
	InterleaveType patternType = 14
)

type pattern struct {
	typ  patternType
	expr funcs.Bool
	sub0 *pattern
	sub1 *pattern
	name string
}

func (this pattern) child() pattern {
	return *this.sub0
}

func (this pattern) left() pattern {
	return *this.sub0
}

func (this pattern) right() pattern {
	return *this.sub1
}

func newNot(p pattern) pattern {
	return pattern{
		typ:  NotType,
		sub0: &p,
	}
}

func newZany() pattern {
	return pattern{
		expr: funcs.BoolConst(true),
		typ:  ZAnyType,
	}
}

func newEmpty() pattern {
	return pattern{
		typ:  EmptyType,
		expr: funcs.BoolConst(true),
	}
}

func newConcat(p1, p2 pattern) pattern {
	return pattern{
		typ:  ConcatType,
		sub0: &p1,
		sub1: &p2,
	}
}

func newOr(p1, p2 pattern) pattern {
	return pattern{
		typ:  OrType,
		sub0: &p1,
		sub1: &p2,
	}
}

func newAnd(p1, p2 pattern) pattern {
	return pattern{
		typ:  AndType,
		sub0: &p1,
		sub1: &p2,
	}
}

func newInterleave(p1, p2 pattern) pattern {
	return pattern{
		typ:  InterleaveType,
		sub0: &p1,
		sub1: &p2,
	}
}

func (p pattern) String() string {
	return ToPattern(p).String()
}

func ToPattern(p pattern) *relapse.Pattern {
	switch p.typ {
	case EmptyType:
		return relapse.NewEmpty()
	case TreeNodeType:
		return relapse.NewTreeNode(nameexpr.FuncToName(p.expr), ToPattern(p.child()))
	case LeafNodeType:
		return relapse.NewLeafNode(funcToExpr(p.expr))
	case ConcatType:
		return relapse.NewConcat(
			ToPattern(p.left()),
			ToPattern(p.right()),
		)
	case OrType:
		return relapse.NewOr(
			ToPattern(p.left()),
			ToPattern(p.right()),
		)
	case AndType:
		return relapse.NewAnd(
			ToPattern(p.left()),
			ToPattern(p.right()),
		)
	case ZeroOrMoreType:
		return relapse.NewZeroOrMore(ToPattern(p.child()))
	case ReferenceType:
		return relapse.NewReference(p.name)
	case NotType:
		return relapse.NewNot(ToPattern(p.child()))
	case ZAnyType:
		return relapse.NewZAny()
	// case ContainsType:
	// 	return relapse.NewContains(ToPattern(p.child()))
	// case OptionalType:
	// 	return relapse.NewOptional(ToPattern(p.child()))
	case InterleaveType:
		return relapse.NewInterleave(
			ToPattern(p.left()),
			ToPattern(p.right()),
		)
	}
	panic(fmt.Sprintf("unknown pattern typ %v", p))
}

func funcToExpr(f funcs.Bool) *expr.Expr {
	exprStr := funcs.Sprint(f)
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return expr
}

func FromPattern(p *relapse.Pattern) pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return newEmpty()
	case *relapse.TreeNode:
		f := nameToFunc(v.GetName())
		child := FromPattern(v.GetPattern())
		return pattern{typ: TreeNodeType, expr: f, sub0: &child}
	case *relapse.LeafNode:
		f := exprToFunc(v.GetExpr())
		return pattern{typ: LeafNodeType, expr: f}
	case *relapse.Concat:
		left := FromPattern(v.GetLeftPattern())
		right := FromPattern(v.GetRightPattern())
		return newConcat(left, right)
	case *relapse.Or:
		left := FromPattern(v.GetLeftPattern())
		right := FromPattern(v.GetRightPattern())
		return newOr(left, right)
	case *relapse.And:
		left := FromPattern(v.GetLeftPattern())
		right := FromPattern(v.GetRightPattern())
		return newAnd(left, right)
	case *relapse.ZeroOrMore:
		child := FromPattern(v.GetPattern())
		return pattern{typ: ZeroOrMoreType, sub0: &child}
	case *relapse.Reference:
		return pattern{typ: ReferenceType, name: v.GetName()}
	case *relapse.Not:
		child := FromPattern(v.GetPattern())
		return newNot(child)
	case *relapse.ZAny:
		return newZany()
	case *relapse.Contains:
		return FromPattern(relapse.NewConcat(relapse.NewZAny(),
			relapse.NewConcat(v.GetPattern(), relapse.NewZAny())))
		// child := FromPattern(v.GetPattern())
		// return pattern{typ: ContainsType, sub0: &child}
	case *relapse.Optional:
		return FromPattern(relapse.NewOr(v.GetPattern(), relapse.NewEmpty()))
		// child := FromPattern(v.GetPattern())
		// return pattern{typ: OptionalType, sub0: &child}
	case *relapse.Interleave:
		left := FromPattern(v.GetLeftPattern())
		right := FromPattern(v.GetRightPattern())
		return newInterleave(left, right)
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

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

package interp

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/nameexpr"
	"github.com/katydid/katydid/serialize"
	"io"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

//This is a naive implementation and it does not handle left recursion
func Interpret(g *relapse.Grammar, tree serialize.Parser) bool {
	refs := relapse.NewRefsLookup(g)
	res := refs["main"]
	res = refs["main"]
	err := tree.Next()
	for err == nil {
		// name, nameErr := tree.String()
		// if nameErr != nil {
		// 	panic(nameErr)
		// }
		// log.Printf("Interpret = %s given input %s", res, name)
		res = sderiv(refs, res, tree)
		err = tree.Next()
	}
	if err != io.EOF {
		panic(err)
	}
	log.Printf("Interpret Final = %s", res)
	return Nullable(refs, res)
}

//TODO improve nullable for left recursion using fix points
// https://github.com/kennknowles/go-yid/blob/master/src/yid/nullable.go
//This is a naive implementation and it does not handle left recursion
func Nullable(refs relapse.RefLookup, p *relapse.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return true
	case *relapse.TreeNode:
		return false
	case *relapse.LeafNode:
		return false
	case *relapse.Concat:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *relapse.Or:
		return Nullable(refs, v.GetLeftPattern()) || Nullable(refs, v.GetRightPattern())
	case *relapse.And:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return true
	case *relapse.Reference:
		return Nullable(refs, refs[v.GetName()])
	case *relapse.Not:
		return !(Nullable(refs, v.GetPattern()))
	case *relapse.ZAny:
		return true
	case *relapse.Contains:
		return Nullable(refs, v.GetPattern())
	case *relapse.Optional:
		return true
	case *relapse.Interleave:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func derivTreeNode(refs relapse.RefLookup, p *relapse.TreeNode, tree serialize.Parser) *relapse.Pattern {
	matched := nameexpr.EvalName(p.GetName(), tree)
	log.Printf("name %s -> %v", p, matched)
	if !matched {
		return relapse.NewNot(relapse.NewZAny())
	}
	tree.Down()
	res := p.GetPattern()
	err := tree.Next()
	for err == nil {
		res = sderiv(refs, res, tree)
		err = tree.Next()
	}
	if err != io.EOF {
		panic(err)
	}
	log.Printf("derivTreeNode Final = %s", res)
	tree.Up()
	if !Nullable(refs, res) {
		return relapse.NewNot(relapse.NewZAny())
	}
	return relapse.NewEmpty()
}

func sderiv(refs relapse.RefLookup, p *relapse.Pattern, tree serialize.Parser) *relapse.Pattern {
	d := deriv(refs, p, tree)
	s := Simplify(refs, d)
	log.Printf("sderiv %s -> %s -> %s", p, d, s)
	return s
}

func deriv(refs relapse.RefLookup, p *relapse.Pattern, tree serialize.Parser) *relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return relapse.NewNot(relapse.NewZAny())
	case *relapse.TreeNode:
		if tree.IsLeaf() {
			return relapse.NewNot(relapse.NewZAny())
		}
		return derivTreeNode(refs, v, tree)
	case *relapse.LeafNode:
		f, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		if !tree.IsLeaf() {
			return relapse.NewNot(relapse.NewZAny())
		}
		res, err := f.Eval(tree)
		if err != nil {
			return relapse.NewNot(relapse.NewZAny())
		}
		if res {
			return relapse.NewEmpty()
		}
		return relapse.NewNot(relapse.NewZAny())
	case *relapse.Concat:
		leftDeriv := relapse.NewConcat(sderiv(refs, v.GetLeftPattern(), tree), v.GetRightPattern())
		if Nullable(refs, v.GetLeftPattern()) {
			return relapse.NewOr(
				leftDeriv,
				sderiv(refs, v.GetRightPattern(), tree.Copy()),
			)
		} else {
			return leftDeriv
		}
	case *relapse.Or:
		treeCopy := tree.Copy()
		return relapse.NewOr(
			sderiv(refs, v.GetLeftPattern(), tree),
			sderiv(refs, v.GetRightPattern(), treeCopy),
		)
	case *relapse.And:
		treeCopy := tree.Copy()
		return relapse.NewAnd(
			sderiv(refs, v.GetLeftPattern(), tree),
			sderiv(refs, v.GetRightPattern(), treeCopy),
		)
	case *relapse.ZeroOrMore:
		return relapse.NewConcat(sderiv(refs, v.Pattern, tree), relapse.NewZeroOrMore(v.Pattern))
	case *relapse.Reference:
		return sderiv(refs, refs[v.GetName()], tree)
	case *relapse.Not:
		return relapse.NewNot(sderiv(refs, v.GetPattern(), tree))
	case *relapse.ZAny:
		return p
	case *relapse.Contains:
		newp := relapse.NewConcat(relapse.NewZAny(), v.GetPattern(), relapse.NewZAny())
		return deriv(refs, newp, tree)
	case *relapse.Optional:
		newp := relapse.NewOr(relapse.NewEmpty(), v.GetPattern())
		return deriv(refs, newp, tree)
	case *relapse.Interleave:
		treeCopy := tree.Copy()
		return relapse.NewOr(
			Simplify(refs, relapse.NewInterleave(sderiv(refs, v.GetLeftPattern(), tree), v.GetRightPattern())),
			Simplify(refs, relapse.NewInterleave(v.GetLeftPattern(), sderiv(refs, v.GetRightPattern(), treeCopy))),
		)
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

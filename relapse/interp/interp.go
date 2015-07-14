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
	"github.com/katydid/katydid/serialize"
	"io"
	"log"
)

func Interpret(g *relapse.Grammar, tree serialize.Scanner) bool {
	refs := newRefsLookup(g)
	res := refs["main"]
	res = refs["main"]
	for {
		err := tree.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		log.Printf("Interpret = %s given input %s", res, tree.Name())
		res = sderiv(refs, res, tree)
	}
	log.Printf("Interpret Final = %s", res)
	return nullable(refs, res)
}

type RefLookup map[string]*relapse.Pattern

func newRefsLookup(g *relapse.Grammar) RefLookup {
	decls := g.GetPatternDecls()
	refs := make(RefLookup, len(decls))
	for _, d := range decls {
		refs[d.Name] = d.Pattern
	}
	return refs
}

func nullable(refs RefLookup, p *relapse.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return true
	case *relapse.EmptySet:
		return false
	case *relapse.TreeNode:
		return false
	case *relapse.LeafNode:
		return false
	case *relapse.Concat:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *relapse.Or:
		return nullable(refs, v.GetLeftPattern()) || nullable(refs, v.GetRightPattern())
	case *relapse.And:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return true
	case *relapse.Reference:
		return nullable(refs, refs[v.GetName()])
	case *relapse.Not:
		return !(nullable(refs, v.GetPattern()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func evalName(nameExpr *relapse.NameExpr, name string) bool {
	typ := nameExpr.GetValue()
	switch v := typ.(type) {
	case *relapse.Name:
		return name == v.GetName()
	case *relapse.AnyName:
		return true
	case *relapse.AnyNameExcept:
		return !evalName(v.GetExcept(), name)
	case *relapse.NameChoice:
		return evalName(v.GetLeft(), name) || evalName(v.GetRight(), name)
	}
	panic(fmt.Sprintf("unknown nameExpr typ %T", typ))
}

func derivTreeNode(refs RefLookup, p *relapse.TreeNode, tree serialize.Scanner) *relapse.Pattern {
	matched := evalName(p.GetName(), tree.Name())
	if !matched {
		tree.Next()
		return relapse.NewEmptySet()
	}
	if tree.IsLeaf() {
		res := sderiv(refs, p.GetPattern(), tree)
		if !nullable(refs, res) {
			return relapse.NewEmptySet()
		}
		return relapse.NewEmpty()
	}
	tree.Down()
	res := p.GetPattern()
	for {
		err := tree.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		log.Printf("derivTreeNode = %s given input %s", res, tree.Name())
		res = sderiv(refs, res, tree)
	}
	log.Printf("derivTreeNode Final = %s", res)
	tree.Up()
	if !nullable(refs, res) {
		return relapse.NewEmptySet()
	}
	return relapse.NewEmpty()
}

func sderiv(refs RefLookup, p *relapse.Pattern, tree serialize.Scanner) *relapse.Pattern {
	return Simplify(refs, deriv(refs, p, tree))
}

func deriv(refs RefLookup, p *relapse.Pattern, tree serialize.Scanner) *relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return relapse.NewEmptySet()
	case *relapse.EmptySet:
		return relapse.NewEmptySet()
	case *relapse.TreeNode:
		return derivTreeNode(refs, v, tree)
	case *relapse.LeafNode:
		f, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		if !tree.IsLeaf() {
			return relapse.NewEmptySet()
		}
		res, err := f.Eval(tree)
		if err != nil {
			panic(err)
		}
		if res {
			return relapse.NewEmpty()
		}
		return relapse.NewEmptySet()
	case *relapse.Concat:
		leftDeriv := relapse.NewConcat(sderiv(refs, v.GetLeftPattern(), tree), v.GetRightPattern())
		if nullable(refs, v.GetLeftPattern()) {
			return relapse.NewOr(
				leftDeriv,
				sderiv(refs, v.GetRightPattern(), tree.Copy()),
			)
		} else {
			return leftDeriv
		}
	case *relapse.Or:
		return relapse.NewOr(
			sderiv(refs, v.GetLeftPattern(), tree),
			sderiv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *relapse.And:
		return relapse.NewAnd(
			sderiv(refs, v.GetLeftPattern(), tree),
			sderiv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *relapse.ZeroOrMore:
		return relapse.NewConcat(sderiv(refs, v.Pattern, tree), relapse.NewZeroOrMore(v.Pattern))
	case *relapse.Reference:
		return sderiv(refs, refs[v.GetName()], tree)
	case *relapse.Not:
		return relapse.NewNot(sderiv(refs, v.GetPattern(), tree))
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

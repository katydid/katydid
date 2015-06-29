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
	"github.com/katydid/katydid/lang/ast"
	"github.com/katydid/katydid/serialize"
	"io"
	"log"
)

func Interpret(g *lang.Grammar, tree serialize.Scanner) bool {
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
		res = deriv(refs, res, tree)
	}
	log.Printf("Interpret Final = %s", res)
	return nullable(refs, res)
}

type RefLookup map[string]*lang.Pattern

func newRefsLookup(g *lang.Grammar) RefLookup {
	decls := g.GetPatternDecls()
	refs := make(RefLookup, len(decls))
	for _, d := range decls {
		refs[d.Name] = d.Pattern
	}
	return refs
}

func nullable(refs RefLookup, p *lang.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return true
	case *lang.EmptySet:
		return false
	case *lang.TreeNode:
		return false
	case *lang.LeafNode:
		return false
	case *lang.Concat:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *lang.Or:
		return nullable(refs, v.GetLeftPattern()) || nullable(refs, v.GetRightPattern())
	case *lang.And:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *lang.ZeroOrMore:
		return true
	case *lang.Reference:
		return nullable(refs, refs[v.GetName()])
	case *lang.Not:
		return !(nullable(refs, v.GetPattern()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func evalName(nameExpr *lang.NameExpr, name string) bool {
	typ := nameExpr.GetValue()
	switch v := typ.(type) {
	case *lang.Name:
		return name == v.GetName()
	case *lang.AnyName:
		return true
	case *lang.AnyNameExcept:
		return !evalName(v.GetExcept(), name)
	case *lang.NameChoice:
		return evalName(v.GetLeft(), name) || evalName(v.GetRight(), name)
	}
	panic(fmt.Sprintf("unknown nameExpr typ %T", typ))
}

func derivTreeNode(refs RefLookup, p *lang.TreeNode, tree serialize.Scanner) *lang.Pattern {
	matched := evalName(p.GetName(), tree.Name())
	if !matched {
		tree.Next()
		return lang.NewEmptySet()
	}
	if tree.IsLeaf() {
		res := deriv(refs, p.GetPattern(), tree)
		if !nullable(refs, res) {
			return lang.NewEmptySet()
		}
		return lang.NewEmpty()
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
		res = deriv(refs, res, tree)
	}
	log.Printf("derivTreeNode Final = %s", res)
	tree.Up()
	if !nullable(refs, res) {
		return lang.NewEmptySet()
	}
	return lang.NewEmpty()
}

func deriv(refs RefLookup, p *lang.Pattern, tree serialize.Scanner) *lang.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return lang.NewEmptySet()
	case *lang.EmptySet:
		return lang.NewEmptySet()
	case *lang.TreeNode:
		return derivTreeNode(refs, v, tree)
	case *lang.LeafNode:
		f, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		if !tree.IsLeaf() {
			return lang.NewEmptySet()
		}
		res, err := f.Eval(tree)
		if err != nil {
			panic(err)
		}
		if res {
			return lang.NewEmpty()
		}
		return lang.NewEmptySet()
	case *lang.Concat:
		leftDeriv := lang.NewConcat(deriv(refs, v.GetLeftPattern(), tree), v.GetRightPattern())
		if nullable(refs, v.GetLeftPattern()) {
			return lang.NewOr(
				leftDeriv,
				deriv(refs, v.GetRightPattern(), tree.Copy()),
			)
		} else {
			return leftDeriv
		}
	case *lang.Or:
		return lang.NewOr(
			deriv(refs, v.GetLeftPattern(), tree),
			deriv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *lang.And:
		return lang.NewAnd(
			deriv(refs, v.GetLeftPattern(), tree),
			deriv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *lang.ZeroOrMore:
		return lang.NewConcat(deriv(refs, v.Pattern, tree), lang.NewZeroOrMore(v.Pattern))
	case *lang.Reference:
		return deriv(refs, refs[v.GetName()], tree)
	case *lang.Not:
		return lang.NewNot(deriv(refs, v.GetPattern(), tree))
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	nameexpr "github.com/katydid/katydid/relapse/name"
)

func checkRef(refs ast.RefLookup, p *ast.Pattern) *ast.Pattern {
	for name, rpat := range refs {
		if rpat.Equal(p) {
			return ast.NewReference(name)
		}
	}
	return p
}

func SimplifyGrammar(g *ast.Grammar) *ast.Grammar {
	refs := ast.NewRefLookup(g)
	p := Simplify(refs, refs["main"])
	refs["main"] = p
	return ast.NewGrammar(refs)
}

func Simplify(refs ast.RefLookup, p *ast.Pattern) *ast.Pattern {
	return simplify(refs, p, true)
}

func simplify(refs ast.RefLookup, p *ast.Pattern, top bool) *ast.Pattern {
	cRef := func(cp *ast.Pattern) *ast.Pattern {
		if top {
			return cp
		}
		return checkRef(refs, cp)
	}
	simp := func(sp *ast.Pattern) *ast.Pattern {
		return simplify(refs, sp, false)
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return p
	case *ast.TreeNode:
		child := simp(v.GetPattern())
		if isNotZany(child) {
			return ast.NewNot(ast.NewZAny())
		}
		name := v.GetName()
		b := nameexpr.NameToFunc(v.GetName())
		if funcs.IsFalse(b) {
			return ast.NewNot(ast.NewZAny())
		}
		if funcs.IsTrue(b) {
			name = ast.NewAnyName()
		}
		return cRef(ast.NewTreeNode(name, child))
	case *ast.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		if funcs.IsFalse(b) {
			return ast.NewNot(ast.NewZAny())
		}
		return p
	case *ast.Concat:
		return cRef(simplifyConcat(
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *ast.Or:
		return cRef(simplifyOr(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *ast.And:
		return cRef(simplifyAnd(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *ast.ZeroOrMore:
		return cRef(simplifyZeroOrMore(simp(v.GetPattern())))
	case *ast.Reference:
		return p
	case *ast.Not:
		return cRef(simplifyNot(simp(v.GetPattern())))
	case *ast.ZAny:
		return p
	case *ast.Contains:
		return cRef(simplifyContains(simp(v.GetPattern())))
	case *ast.Optional:
		return simplifyOptional(simp(v.GetPattern()))
	case *ast.Interleave:
		return cRef(simplifyInterleave(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func isNotZany(p *ast.Pattern) bool {
	return p.Not != nil && p.GetNot().GetPattern().ZAny != nil
}

func isEmpty(p *ast.Pattern) bool {
	return p.Empty != nil
}

func isZany(p *ast.Pattern) bool {
	return p.ZAny != nil
}

func simplifyConcat(p1, p2 *ast.Pattern) *ast.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return ast.NewNot(ast.NewZAny())
	}
	if p1.Concat != nil {
		return simplifyConcat(
			p1.Concat.GetLeftPattern(),
			ast.NewConcat(p1.Concat.GetRightPattern(), p2),
		)
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && p2.Concat != nil {
		if l := p2.Concat.GetLeftPattern(); isZany(p2.Concat.GetRightPattern()) {
			return ast.NewContains(l)
		}
	}
	return ast.NewConcat(p1, p2)
}

func getOrs(p *ast.Pattern) []*ast.Pattern {
	if p.Or != nil {
		return append(getOrs(p.Or.GetLeftPattern()), getOrs(p.Or.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func simplifyOr(refs ast.RefLookup, p1, p2 *ast.Pattern) *ast.Pattern {
	if isNotZany(p1) {
		return p2
	}
	if isNotZany(p2) {
		return p1
	}
	if isZany(p1) || isZany(p2) {
		return ast.NewZAny()
	}
	if isEmpty(p1) && Nullable(refs, p2) {
		return p2
	}
	if isEmpty(p2) && Nullable(refs, p1) {
		return p1
	}
	left := getOrs(p1)
	right := getOrs(p2)
	list := append(left, right...)
	list = ast.Set(list)
	list = simplifyChildren(list, func(left, right *ast.Pattern) *ast.Pattern {
		return simplifyOr(refs, left, right)
	})
	ast.Sort(list)
	var p *ast.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = ast.NewOr(p, list[i])
	}
	return p
}

func getAnds(p *ast.Pattern) []*ast.Pattern {
	if p.And != nil {
		return append(getAnds(p.And.GetLeftPattern()), getAnds(p.And.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func simplifyChildren(children []*ast.Pattern, op func(left, right *ast.Pattern) *ast.Pattern) []*ast.Pattern {
	if len(children) == 0 || len(children) == 1 {
		return children
	}

	t0 := children[0].TreeNode
	t1 := children[1].TreeNode
	if t0 != nil && t1 != nil {
		if t0.GetName().Equal(t1.GetName()) {
			newchild := ast.NewTreeNode(t0.GetName(), op(t0.GetPattern(), t1.GetPattern()))
			children[1] = newchild
			return simplifyChildren(children[1:], op)
		}
	}
	return append([]*ast.Pattern{children[0]}, simplifyChildren(children[1:], op)...)
}

func simplifyAnd(refs ast.RefLookup, p1, p2 *ast.Pattern) *ast.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return ast.NewNot(ast.NewZAny())
	}
	if isZany(p1) {
		return p2
	}
	if isZany(p2) {
		return p1
	}
	if isEmpty(p1) {
		if Nullable(refs, p2) {
			return ast.NewEmpty()
		} else {
			return ast.NewNot(ast.NewZAny())
		}
	}
	if isEmpty(p2) {
		if Nullable(refs, p1) {
			return ast.NewEmpty()
		} else {
			return ast.NewNot(ast.NewZAny())
		}
	}
	left := getAnds(p1)
	right := getAnds(p2)
	list := append(left, right...)
	list = ast.Set(list)
	list = simplifyChildren(list, func(left, right *ast.Pattern) *ast.Pattern {
		return simplifyAnd(refs, left, right)
	})
	ast.Sort(list)
	var p *ast.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = ast.NewAnd(p, list[i])
	}
	return p
}

func simplifyZeroOrMore(p *ast.Pattern) *ast.Pattern {
	if p.ZeroOrMore != nil {
		return p
	}
	return ast.NewZeroOrMore(p)
}

func simplifyNot(p *ast.Pattern) *ast.Pattern {
	if p.Not != nil {
		return p.Not.GetPattern()
	}
	return ast.NewNot(p)
}

func simplifyOptional(p *ast.Pattern) *ast.Pattern {
	if isEmpty(p) {
		return ast.NewEmpty()
	}
	return ast.NewOptional(p)
}

func getInterleaves(p *ast.Pattern) []*ast.Pattern {
	if p.Interleave != nil {
		return append(getInterleaves(p.Interleave.GetLeftPattern()), getInterleaves(p.Interleave.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func simplifyInterleave(refs ast.RefLookup, p1, p2 *ast.Pattern) *ast.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return ast.NewNot(ast.NewZAny())
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && isZany(p2) {
		return p1
	}
	left := getInterleaves(p1)
	right := getInterleaves(p2)
	list := append(left, right...)
	list = ast.Set(list)
	ast.Sort(list)
	var p *ast.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = ast.NewInterleave(p, list[i])
	}
	return p
}

func simplifyContains(p *ast.Pattern) *ast.Pattern {
	if isEmpty(p) || isZany(p) {
		return ast.NewZAny()
	}
	if isNotZany(p) {
		return p
	}
	return ast.NewContains(p)
}

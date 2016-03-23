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
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/expr/funcs"
	nameexpr "github.com/katydid/katydid/expr/name"
	"github.com/katydid/katydid/relapse/ast"
)

func checkRef(refs relapse.RefLookup, p *relapse.Pattern) *relapse.Pattern {
	for name, rpat := range refs {
		if rpat.Equal(p) {
			return relapse.NewReference(name)
		}
	}
	return p
}

func SimplifyGrammar(g *relapse.Grammar) *relapse.Grammar {
	refs := relapse.NewRefsLookup(g)
	p := Simplify(refs, refs["main"])
	refs["main"] = p
	return relapse.NewGrammar(refs)
}

func Simplify(refs relapse.RefLookup, p *relapse.Pattern) *relapse.Pattern {
	return simplify(refs, p, true)
}

func simplify(refs relapse.RefLookup, p *relapse.Pattern, top bool) *relapse.Pattern {
	cRef := func(cp *relapse.Pattern) *relapse.Pattern {
		if top {
			return cp
		}
		return checkRef(refs, cp)
	}
	simp := func(sp *relapse.Pattern) *relapse.Pattern {
		return simplify(refs, sp, false)
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return p
	case *relapse.TreeNode:
		child := simp(v.GetPattern())
		if isNotZany(child) {
			return relapse.NewNot(relapse.NewZAny())
		}
		name := v.GetName()
		b := nameexpr.NameToFunc(v.GetName())
		if funcs.IsFalse(b) {
			return relapse.NewNot(relapse.NewZAny())
		}
		if funcs.IsTrue(b) {
			name = expr.NewAnyName()
		}
		return cRef(relapse.NewTreeNode(name, child))
	case *relapse.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		if funcs.IsFalse(b) {
			return relapse.NewNot(relapse.NewZAny())
		}
		return p
	case *relapse.Concat:
		return cRef(simplifyConcat(
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *relapse.Or:
		return cRef(simplifyOr(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *relapse.And:
		return cRef(simplifyAnd(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *relapse.ZeroOrMore:
		return cRef(simplifyZeroOrMore(simp(v.GetPattern())))
	case *relapse.Reference:
		return p
	case *relapse.Not:
		return cRef(simplifyNot(simp(v.GetPattern())))
	case *relapse.ZAny:
		return p
	case *relapse.Contains:
		return cRef(simplifyContains(simp(v.GetPattern())))
	case *relapse.Optional:
		return simplifyOptional(simp(v.GetPattern()))
	case *relapse.Interleave:
		return cRef(simplifyInterleave(refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func isNotZany(p *relapse.Pattern) bool {
	return p.Not != nil && p.GetNot().GetPattern().ZAny != nil
}

func isEmpty(p *relapse.Pattern) bool {
	return p.Empty != nil
}

func isZany(p *relapse.Pattern) bool {
	return p.ZAny != nil
}

func simplifyConcat(p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return relapse.NewNot(relapse.NewZAny())
	}
	if p1.Concat != nil {
		return simplifyConcat(
			p1.Concat.GetLeftPattern(),
			relapse.NewConcat(p1.Concat.GetRightPattern(), p2),
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
			return relapse.NewContains(l)
		}
	}
	return relapse.NewConcat(p1, p2)
}

func getOrs(p *relapse.Pattern) []*relapse.Pattern {
	if p.Or != nil {
		return append(getOrs(p.Or.GetLeftPattern()), getOrs(p.Or.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyOr(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) {
		return p2
	}
	if isNotZany(p2) {
		return p1
	}
	if isZany(p1) || isZany(p2) {
		return relapse.NewZAny()
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
	list = relapse.Set(list)
	list = simplifyChildren(list, func(left, right *relapse.Pattern) *relapse.Pattern {
		return simplifyOr(refs, left, right)
	})
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewOr(p, list[i])
	}
	return p
}

func getAnds(p *relapse.Pattern) []*relapse.Pattern {
	if p.And != nil {
		return append(getAnds(p.And.GetLeftPattern()), getAnds(p.And.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyChildren(children []*relapse.Pattern, op func(left, right *relapse.Pattern) *relapse.Pattern) []*relapse.Pattern {
	if len(children) == 0 || len(children) == 1 {
		return children
	}

	t0 := children[0].TreeNode
	t1 := children[1].TreeNode
	if t0 != nil && t1 != nil {
		if t0.GetName().Equal(t1.GetName()) {
			newchild := relapse.NewTreeNode(t0.GetName(), op(t0.GetPattern(), t1.GetPattern()))
			children[1] = newchild
			return simplifyChildren(children[1:], op)
		}
	}
	return append([]*relapse.Pattern{children[0]}, simplifyChildren(children[1:], op)...)
}

func simplifyAnd(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return relapse.NewNot(relapse.NewZAny())
	}
	if isZany(p1) {
		return p2
	}
	if isZany(p2) {
		return p1
	}
	if isEmpty(p1) {
		if Nullable(refs, p2) {
			return relapse.NewEmpty()
		} else {
			return relapse.NewNot(relapse.NewZAny())
		}
	}
	if isEmpty(p2) {
		if Nullable(refs, p1) {
			return relapse.NewEmpty()
		} else {
			return relapse.NewNot(relapse.NewZAny())
		}
	}
	left := getAnds(p1)
	right := getAnds(p2)
	list := append(left, right...)
	list = relapse.Set(list)
	list = simplifyChildren(list, func(left, right *relapse.Pattern) *relapse.Pattern {
		return simplifyAnd(refs, left, right)
	})
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewAnd(p, list[i])
	}
	return p
}

func simplifyZeroOrMore(p *relapse.Pattern) *relapse.Pattern {
	if p.ZeroOrMore != nil {
		return p
	}
	return relapse.NewZeroOrMore(p)
}

func simplifyNot(p *relapse.Pattern) *relapse.Pattern {
	if p.Not != nil {
		return p.Not.GetPattern()
	}
	return relapse.NewNot(p)
}

func simplifyOptional(p *relapse.Pattern) *relapse.Pattern {
	if isEmpty(p) {
		return relapse.NewEmpty()
	}
	return relapse.NewOptional(p)
}

func getInterleaves(p *relapse.Pattern) []*relapse.Pattern {
	if p.Interleave != nil {
		return append(getInterleaves(p.Interleave.GetLeftPattern()), getInterleaves(p.Interleave.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyInterleave(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return relapse.NewNot(relapse.NewZAny())
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
	list = relapse.Set(list)
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewInterleave(p, list[i])
	}
	return p
}

func simplifyContains(p *relapse.Pattern) *relapse.Pattern {
	if isEmpty(p) || isZany(p) {
		return relapse.NewZAny()
	}
	if isNotZany(p) {
		return p
	}
	return relapse.NewContains(p)
}

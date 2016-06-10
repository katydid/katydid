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

//Simplifier simplifies the patterns of a given grammar.
type Simplifier interface {
	//Simplify returns a pattern that has been simplified, transformed to an equivalent, but smaller or equal pattern.
	Simplify(p *ast.Pattern) *ast.Pattern
	//Grammar returns a grammar that has been simplified, transformed to an equivalent, but smaller or equal grammar.
	Grammar() *ast.Grammar
	//OptimizeForRecord optimizes the simplifier to also takes into account the fact the structure being validated is a record.
	//A record can be json, protobuf, reflected go structures or any structure that have unique field names for each structure.
	//XML would be an example of a structure for which this simplification is NOT appropriate.
	OptimizeForRecord() Simplifier
}

type simplifier struct {
	refs   ast.RefLookup
	cache  map[*ast.Pattern]struct{}
	record bool
}

//NewSimplifier returns a new Simplifier that is used to simplify a grammar and its patterns.
func NewSimplifier(g *ast.Grammar) Simplifier {
	return &simplifier{
		refs:  ast.NewRefLookup(g),
		cache: make(map[*ast.Pattern]struct{}),
	}
}

func (this *simplifier) Simplify(p *ast.Pattern) *ast.Pattern {
	if _, ok := this.cache[p]; ok {
		return p
	}
	s := this.simplify(p, true)
	this.cache[s] = struct{}{}
	return s
}

func (this *simplifier) Grammar() *ast.Grammar {
	for name, p := range this.refs {
		this.refs[name] = this.Simplify(p)
	}
	return ast.NewGrammar(this.refs)
}

func (this *simplifier) OptimizeForRecord() Simplifier {
	this.record = true
	return this
}

func checkRef(refs ast.RefLookup, p *ast.Pattern) *ast.Pattern {
	for name, rpat := range refs {
		if rpat.Equal(p) {
			return ast.NewReference(name)
		}
	}
	return p
}

var emptyset = ast.NewNot(ast.NewZAny())

func (this *simplifier) simplify(p *ast.Pattern, top bool) *ast.Pattern {
	cRef := func(cp *ast.Pattern) *ast.Pattern {
		if top {
			return cp
		}
		return checkRef(this.refs, cp)
	}
	cachesimp := func(sp *ast.Pattern) *ast.Pattern {
		if _, ok := this.cache[sp]; ok {
			return sp
		}
		s := this.simplify(sp, false)
		this.cache[s] = struct{}{}
		return s
	}
	simp := func(sp *ast.Pattern) *ast.Pattern {
		return this.simplify(sp, false)
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return p
	case *ast.TreeNode:
		child := cachesimp(v.GetPattern())
		if isNotZany(child) {
			return emptyset
		}
		name := v.GetName()
		b := nameexpr.NameToFunc(v.GetName())
		if funcs.IsFalse(b) {
			return emptyset
		}
		if funcs.IsTrue(b) {
			name = ast.NewAnyName()
		}
		return cRef(ast.NewTreeNode(name, child))
	case *ast.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			//Don't simplify if there is an error to keep this function signature simple.
			return p
		}
		if funcs.IsFalse(b) {
			return emptyset
		}
		return p
	case *ast.Concat:
		return cRef(simplifyConcat(
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
		))
	case *ast.Or:
		return cRef(simplifyOr(this.refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
			this.record,
		))
	case *ast.And:
		return cRef(simplifyAnd(this.refs,
			simp(v.GetLeftPattern()),
			simp(v.GetRightPattern()),
			this.record,
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
		return cRef(simplifyInterleave(this.refs,
			cachesimp(v.GetLeftPattern()),
			cachesimp(v.GetRightPattern()),
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
		return emptyset
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

func simplifyOr(refs ast.RefLookup, p1, p2 *ast.Pattern, record bool) *ast.Pattern {
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
		return simplifyOr(refs, left, right, record)
	}, record)
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

func simplifyChildren(children []*ast.Pattern, op func(left, right *ast.Pattern) *ast.Pattern, record bool) []*ast.Pattern {
	if len(children) == 0 || len(children) == 1 {
		return children
	}

	if record {
		c0 := children[0].GetContains().GetPattern().GetTreeNode()
		c1 := children[1].GetContains().GetPattern().GetTreeNode()
		if c0 != nil && c1 != nil {
			if c0.GetName().Equal(c1.GetName()) {
				newchild := ast.NewContains(ast.NewTreeNode(c0.GetName(), op(c0.GetPattern(), c1.GetPattern())))
				children[1] = newchild
				return simplifyChildren(children[1:], op, record)
			}
		}
	}

	t0 := children[0].TreeNode
	t1 := children[1].TreeNode
	if t0 != nil && t1 != nil {
		if t0.GetName().Equal(t1.GetName()) {
			newchild := ast.NewTreeNode(t0.GetName(), op(t0.GetPattern(), t1.GetPattern()))
			children[1] = newchild
			return simplifyChildren(children[1:], op, record)
		}
	}

	return append([]*ast.Pattern{children[0]}, simplifyChildren(children[1:], op, record)...)
}

func simplifyAnd(refs ast.RefLookup, p1, p2 *ast.Pattern, record bool) *ast.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return emptyset
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
			return emptyset
		}
	}
	if isEmpty(p2) {
		if Nullable(refs, p1) {
			return ast.NewEmpty()
		} else {
			return emptyset
		}
	}
	left := getAnds(p1)
	right := getAnds(p2)
	list := append(left, right...)
	list = ast.Set(list)
	list = simplifyChildren(list, func(left, right *ast.Pattern) *ast.Pattern {
		return simplifyAnd(refs, left, right, record)
	}, record)
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
		return emptyset
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

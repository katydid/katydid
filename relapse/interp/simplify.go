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
)

func Simplify(refs relapse.RefLookup, p *relapse.Pattern) *relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return p
	case *relapse.EmptySet:
		return p
	case *relapse.TreeNode:
		return relapse.NewTreeNode(v.GetName(), Simplify(refs, v.GetPattern()))
	case *relapse.LeafNode:
		return p
	case *relapse.Concat:
		return simplifyConcat(
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		)
	case *relapse.Or:
		return simplifyOr(refs,
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		)
	case *relapse.And:
		return simplifyAnd(refs,
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		)
	case *relapse.ZeroOrMore:
		return relapse.NewZeroOrMore(Simplify(refs, v.GetPattern()))
	case *relapse.Reference:
		return p
	case *relapse.Not:
		return simplifyNot(v.GetPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func isEmptySet(p *relapse.Pattern) bool {
	return p.EmptySet != nil
}

func isEmpty(p *relapse.Pattern) bool {
	return p.Empty != nil
}

func isNotEmptySet(p *relapse.Pattern) bool {
	if p.Not != nil {
		return isEmptySet(p.Not.GetPattern())
	}
	return false
}

func simplifyConcat(p1, p2 *relapse.Pattern) *relapse.Pattern {
	if p1.Concat != nil {
		return simplifyConcat(
			p1.Concat.GetLeftPattern(),
			relapse.NewConcat(p1.Concat.GetRightPattern(), p2),
		)
	}
	if isEmptySet(p1) || isEmptySet(p2) {
		return relapse.NewEmptySet()
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	return relapse.NewConcat(p1, p2)
}

func simplifyOr(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isEmptySet(p1) {
		return p2
	}
	if isEmptySet(p2) {
		return p1
	}
	if isNotEmptySet(p1) || isNotEmptySet(p2) {
		return relapse.NewNot(relapse.NewEmptySet())
	}
	if isEmpty(p1) && Nullable(refs, p2) {
		return p2
	}
	if isEmpty(p2) && Nullable(refs, p1) {
		return p1
	}
	if p1.Equal(p2) {
		return p1
	}
	if p2.Or != nil {
		if p1.Equal(p2.Or.GetLeftPattern()) || p1.Equal(p2.Or.GetRightPattern()) {
			return simplifyOr(refs, p2.Or.GetLeftPattern(), p2.Or.GetRightPattern())
		}
	}
	return relapse.NewOr(p1, p2)
}

func simplifyAnd(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isEmptySet(p1) || isEmptySet(p2) {
		return relapse.NewEmptySet()
	}
	if isNotEmptySet(p1) {
		return p2
	}
	if isNotEmptySet(p2) {
		return p1
	}
	if isEmpty(p1) && Nullable(refs, p2) ||
		isEmpty(p2) && Nullable(refs, p1) {
		return relapse.NewEmpty()
	} else {
		return relapse.NewEmptySet()
	}
	if p1.Equal(p2) {
		return p1
	}
	return relapse.NewAnd(p1, p2)
}

func simplifyNot(p *relapse.Pattern) *relapse.Pattern {
	if p.Not != nil {
		return p.Not.GetPattern()
	}
	return relapse.NewNot(p)
}

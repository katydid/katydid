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
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/relapse"
)

// Experimental
func ChildrenOf(g *relapse.Grammar, path []*expr.NameExpr) []*relapse.Pattern {
	refs := relapse.NewRefsLookup(g)
	p := refs["main"]
	return childrenOf(refs, p, path)
}

func childrenOf(refs relapse.RefLookup, p *relapse.Pattern, path []*expr.NameExpr) []*relapse.Pattern {
	if len(path) == 0 {
		return []*relapse.Pattern{p}
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return nil
	case *relapse.TreeNode:
		if !v.GetName().Equal(path[0]) {
			return nil
		}
		return childrenOf(refs, v.GetPattern(), path[1:])
	case *relapse.LeafNode:
		return nil
	case *relapse.Concat:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.Or:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.And:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.ZeroOrMore:
		return childrenOf(refs, v.GetPattern(), path)
	case *relapse.Reference:
		return childrenOf(refs, refs[v.GetName()], path)
	case *relapse.Not:
		return nil
	case *relapse.ZAny:
		return nil
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

//Experimental
func CompareVarExpr(a, b *expr.Expr) bool {
	if !a.HasVar() && !b.HasVar() {
		return true
	}
	if !a.HasVar() || !b.HasVar() {
		return false
	}
	//both have a var somewhere, so lets continue comparing
	if a.Terminal != nil && b.Terminal != nil {
		if a.Terminal.GetVariable().GetType() != b.Terminal.GetVariable().GetType() {
			return false
		}
		return true
	}
	if a.List != nil && b.List != nil {
		return compareVarExprs(a.GetList().GetElems(), b.GetList().GetElems())
	}
	if a.Function != nil && b.Function != nil {
		if a.GetFunction().GetName() != b.GetFunction().GetName() {
			return false
		}
		as := a.GetFunction().GetParams()
		bs := b.GetFunction().GetParams()
		return compareVarExprs(as, bs)
	}
	if a.BuiltIn != nil && b.BuiltIn != nil {
		if a.GetBuiltIn().GetSymbol().GetValue() != b.GetBuiltIn().GetSymbol().GetValue() {
			return false
		}
		return CompareVarExpr(a.GetBuiltIn().GetExpr(), b.GetBuiltIn().GetExpr())
	}
	//types are different
	return false
}

func compareVarExprs(as, bs []*expr.Expr) bool {
	if len(as) != len(bs) {
		return false
	}
	for i, a := range as {
		if !CompareVarExpr(a, bs[i]) {
			return false
		}
	}
	return true
}

//Experimental
func LeafNodesOf(g *relapse.Grammar, ps ...*relapse.Pattern) []*relapse.Pattern {
	return ofs(relapse.NewRefsLookup(g), ps, func(p1 *relapse.Pattern) bool {
		return p1.LeafNode != nil
	})
}

//Experimental
func TreeNodesOf(g *relapse.Grammar, ps ...*relapse.Pattern) []*relapse.Pattern {
	return ofs(relapse.NewRefsLookup(g), ps, func(p1 *relapse.Pattern) bool {
		return p1.TreeNode != nil
	})
}

func ofs(refs relapse.RefLookup, ps []*relapse.Pattern, f func(p *relapse.Pattern) bool) []*relapse.Pattern {
	fs := []*relapse.Pattern{}
	for _, p := range ps {
		fs = append(fs, of(refs, p, f)...)
	}
	return fs
}

func of(refs relapse.RefLookup, p *relapse.Pattern, f func(p *relapse.Pattern) bool) []*relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return nil
	case *relapse.TreeNode:
		if f(p) {
			return []*relapse.Pattern{p}
		}
		return nil
	case *relapse.LeafNode:
		if f(p) {
			return []*relapse.Pattern{p}
		}
		return nil
	case *relapse.Concat:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.Or:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.And:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*relapse.Pattern{}, left...), right...)
	case *relapse.ZeroOrMore:
		return of(refs, v.GetPattern(), f)
	case *relapse.Reference:
		return of(refs, refs[v.GetName()], f)
	case *relapse.Not:
		return nil
	case *relapse.ZAny:
		return nil
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

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

// Experimental
func ChildrenOf(g *ast.Grammar, path []*ast.NameExpr) []*ast.Pattern {
	refs := ast.NewRefsLookup(g)
	p := refs["main"]
	return childrenOf(refs, p, path)
}

func childrenOf(refs ast.RefLookup, p *ast.Pattern, path []*ast.NameExpr) []*ast.Pattern {
	if len(path) == 0 {
		return []*ast.Pattern{p}
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return nil
	case *ast.TreeNode:
		if !v.GetName().Equal(path[0]) {
			return nil
		}
		return childrenOf(refs, v.GetPattern(), path[1:])
	case *ast.LeafNode:
		return nil
	case *ast.Concat:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.Or:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.And:
		left := childrenOf(refs, v.GetLeftPattern(), path)
		right := childrenOf(refs, v.GetRightPattern(), path)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.ZeroOrMore:
		return childrenOf(refs, v.GetPattern(), path)
	case *ast.Reference:
		return childrenOf(refs, refs[v.GetName()], path)
	case *ast.Not:
		return nil
	case *ast.ZAny:
		return nil
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

//Experimental
func CompareVarExpr(a, b *ast.Expr) bool {
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

func compareVarExprs(as, bs []*ast.Expr) bool {
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
func LeafNodesOf(g *ast.Grammar, ps ...*ast.Pattern) []*ast.Pattern {
	return ofs(ast.NewRefsLookup(g), ps, func(p1 *ast.Pattern) bool {
		return p1.LeafNode != nil
	})
}

//Experimental
func TreeNodesOf(g *ast.Grammar, ps ...*ast.Pattern) []*ast.Pattern {
	return ofs(ast.NewRefsLookup(g), ps, func(p1 *ast.Pattern) bool {
		return p1.TreeNode != nil
	})
}

func ofs(refs ast.RefLookup, ps []*ast.Pattern, f func(p *ast.Pattern) bool) []*ast.Pattern {
	fs := []*ast.Pattern{}
	for _, p := range ps {
		fs = append(fs, of(refs, p, f)...)
	}
	return fs
}

func of(refs ast.RefLookup, p *ast.Pattern, f func(p *ast.Pattern) bool) []*ast.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return nil
	case *ast.TreeNode:
		if f(p) {
			return []*ast.Pattern{p}
		}
		return nil
	case *ast.LeafNode:
		if f(p) {
			return []*ast.Pattern{p}
		}
		return nil
	case *ast.Concat:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.Or:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.And:
		left := of(refs, v.GetLeftPattern(), f)
		right := of(refs, v.GetRightPattern(), f)
		return append(append([]*ast.Pattern{}, left...), right...)
	case *ast.ZeroOrMore:
		return of(refs, v.GetPattern(), f)
	case *ast.Reference:
		return of(refs, refs[v.GetName()], f)
	case *ast.Not:
		return nil
	case *ast.ZAny:
		return nil
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

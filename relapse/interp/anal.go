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
	"github.com/katydid/katydid/relapse/ast"
)

// Experimental
func GetLeafs(g *relapse.Grammar, path ...string) []*expr.Expr {
	refs := relapse.NewRefsLookup(g)
	p := refs["main"]
	return getLeafs(refs, p, path)
}

func getLeafs(refs relapse.RefLookup, p *relapse.Pattern, path []string) []*expr.Expr {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return nil
	case *relapse.EmptySet:
		return nil
	case *relapse.TreeNode:
		if len(path) == 0 {
			return nil
		}
		n, ok := v.GetName().GetValue().(*relapse.Name)
		if !ok {
			return nil
		}
		if n.GetName() != path[0] {
			return nil
		}
		return getLeafs(refs, v.GetPattern(), path[1:])
	case *relapse.LeafNode:
		if len(path) > 0 {
			return nil
		}
		return []*expr.Expr{v.GetExpr()}
	case *relapse.Concat:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *relapse.Or:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *relapse.And:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *relapse.ZeroOrMore:
		return getLeafs(refs, v.GetPattern(), path)
	case *relapse.Reference:
		return getLeafs(refs, refs[v.GetName()], path)
	case *relapse.Not:
		return getLeafs(refs, v.GetPattern(), path)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

//Experimental
func CompareVarExpr(a, b *expr.Expr) bool {
	if !hasVar(a) && !hasVar(b) {
		return true
	}
	if !hasVar(a) || !hasVar(b) {
		return false
	}
	//both have a var somewhere, so lets continue comparing
	if a.Terminal != nil && b.Terminal != nil {
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

func hasVar(e *expr.Expr) bool {
	v, _ := e.GetTerminalVariable()
	return v != nil
}

//Experimental
func Replace(g *relapse.Grammar, current *expr.Expr, replacement *expr.Expr, path ...string) *relapse.Grammar {
	g2 := g.Clone()
	refs := relapse.NewRefsLookup(g2)
	p := refs["main"]
	replace(refs, p, path, current, replacement)
	return g2
}

func replace(refs relapse.RefLookup, p *relapse.Pattern, path []string, current *expr.Expr, replacement *expr.Expr) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return
	case *relapse.EmptySet:
		return
	case *relapse.TreeNode:
		if len(path) == 0 {
			return
		}
		n, ok := v.GetName().GetValue().(*relapse.Name)
		if !ok {
			return
		}
		if n.GetName() != path[0] {
			return
		}
		replace(refs, v.GetPattern(), path[1:], current, replacement)
		return
	case *relapse.LeafNode:
		if len(path) > 0 {
			return
		}
		if v.Expr.String() != current.String() {
			return
		}
		v.Expr = replacement
		return
	case *relapse.Concat:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *relapse.Or:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *relapse.And:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *relapse.ZeroOrMore:
		replace(refs, v.GetPattern(), path, current, replacement)
		return
	case *relapse.Reference:
		replace(refs, refs[v.GetName()], path, current, replacement)
		return
	case *relapse.Not:
		replace(refs, v.GetPattern(), path, current, replacement)
		return
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

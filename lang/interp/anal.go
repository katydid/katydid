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
	"github.com/katydid/katydid/lang/ast"
)

// Experimental
func GetLeafs(g *lang.Grammar, path ...string) []*expr.Expr {
	refs := newRefsLookup(g)
	p := refs["main"]
	return getLeafs(refs, p, path)
}

func getLeafs(refs RefLookup, p *lang.Pattern, path []string) []*expr.Expr {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return nil
	case *lang.EmptySet:
		return nil
	case *lang.TreeNode:
		if len(path) == 0 {
			return nil
		}
		n, ok := v.GetName().GetValue().(*lang.Name)
		if !ok {
			return nil
		}
		if n.GetName() != path[0] {
			return nil
		}
		return getLeafs(refs, v.GetPattern(), path[1:])
	case *lang.LeafNode:
		if len(path) > 0 {
			return nil
		}
		return []*expr.Expr{v.GetExpr()}
	case *lang.Concat:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *lang.Or:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *lang.And:
		left := getLeafs(refs, v.GetLeftPattern(), path)
		right := getLeafs(refs, v.GetRightPattern(), path)
		return append(append([]*expr.Expr{}, left...), right...)
	case *lang.ZeroOrMore:
		return getLeafs(refs, v.GetPattern(), path)
	case *lang.Reference:
		return getLeafs(refs, refs[v.GetName()], path)
	case *lang.Not:
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
func Replace(g *lang.Grammar, current *expr.Expr, replacement *expr.Expr, path ...string) *lang.Grammar {
	g2 := g.Clone()
	refs := newRefsLookup(g2)
	p := refs["main"]
	replace(refs, p, path, current, replacement)
	return g2
}

func replace(refs RefLookup, p *lang.Pattern, path []string, current *expr.Expr, replacement *expr.Expr) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return
	case *lang.EmptySet:
		return
	case *lang.TreeNode:
		if len(path) == 0 {
			return
		}
		n, ok := v.GetName().GetValue().(*lang.Name)
		if !ok {
			return
		}
		if n.GetName() != path[0] {
			return
		}
		replace(refs, v.GetPattern(), path[1:], current, replacement)
		return
	case *lang.LeafNode:
		if len(path) > 0 {
			return
		}
		if v.Expr.String() != current.String() {
			return
		}
		v.Expr = replacement
		return
	case *lang.Concat:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *lang.Or:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *lang.And:
		replace(refs, v.GetLeftPattern(), path, current, replacement)
		replace(refs, v.GetRightPattern(), path, current, replacement)
		return
	case *lang.ZeroOrMore:
		replace(refs, v.GetPattern(), path, current, replacement)
		return
	case *lang.Reference:
		replace(refs, refs[v.GetName()], path, current, replacement)
		return
	case *lang.Not:
		replace(refs, v.GetPattern(), path, current, replacement)
		return
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

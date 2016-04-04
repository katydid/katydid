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

package lazymem

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
)

func ConvertGrammar(g *ast.Grammar) *Pattern {
	refs := ast.NewRefsLookup(g)
	lazyRefs := make(map[string]*Pattern)
	for name, _ := range refs {
		lazyRefs[name] = &Pattern{}
	}
	for name, _ := range refs {
		newp := newPattern(lazyRefs, refs[name])
		lazyRefs[name].head = newp.head
		lazyRefs[name].thunk = newp.thunk
		lazyRefs[name].name = newp.name
	}
	return lazyRefs["main"]
}

func newPattern(refs map[string]*Pattern, p *ast.Pattern) *Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return NewEmpty()
	case *ast.TreeNode:
		return NewTreeNode(v.GetName(), newPattern(refs, v.GetPattern()))
	case *ast.LeafNode:
		return NewLeafNode(v.GetExpr())
	case *ast.Concat:
		return NewConcat(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *ast.Or:
		return NewOr(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *ast.And:
		return NewAnd(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *ast.ZeroOrMore:
		return NewZeroOrMore(newPattern(refs, v.GetPattern()))
	case *ast.Not:
		return NewNot(newPattern(refs, v.GetPattern()))
	case *ast.Reference:
		return NewLazyPattern(v.GetName(), refs[v.GetName()])
	case *ast.ZAny:
		return NewZAny()
	case *ast.Contains:
		return NewContains(newPattern(refs, v.GetPattern()))
	case *ast.Optional:
		return NewOptional(newPattern(refs, v.GetPattern()))
	case *ast.Interleave:
		return NewInterleave(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	}
	panic(fmt.Sprintf("unknown Pattern typ %T", typ))
}

func ConvertPattern(p *Pattern) *ast.Grammar {
	refs := make(map[string]*ast.Pattern)
	visited := make(map[*Pattern]string)
	visited[p] = "main"
	changed := true
	for changed {
		changed = false
		for p, name := range visited {
			if _, ok := refs[name]; !ok {
				headp := p
				if p.thunk != nil {
					headp = p.thunk()
				}
				visited[headp] = name
				newp := convertPattern(visited, headp)
				refs[name] = newp
				changed = true
			}
		}
	}
	for name, pat := range refs {
		if name != "main" {
			refs[name] = interp.Simplify(refs, pat)
		}
	}
	return ast.NewGrammar(refs)
}

func convertPattern(visited map[*Pattern]string, p *Pattern) *ast.Pattern {
	if p.thunk != nil {
		if _, ok := visited[p]; !ok {
			visited[p] = p.name
		}
		return ast.NewReference(visited[p])
	}
	typ := p.head.GetValue()
	switch v := typ.(type) {
	case *Empty:
		return ast.NewEmpty()
	case *Node:
		if v.Name != nil {
			c := convertPattern(visited, v.Pattern)
			return ast.NewTreeNode(v.Name, c)
		} else {
			return ast.NewLeafNode(v.Expr)
		}
	case *Concat:
		left := convertPattern(visited, v.Left)
		right := convertPattern(visited, v.Right)
		return ast.NewConcat(left, right)
	case *Or:
		left := convertPattern(visited, v.Left)
		right := convertPattern(visited, v.Right)
		return ast.NewOr(left, right)
	case *And:
		left := convertPattern(visited, v.Left)
		right := convertPattern(visited, v.Right)
		return ast.NewAnd(left, right)
	case *ZeroOrMore:
		c := convertPattern(visited, v.Pattern)
		return ast.NewZeroOrMore(c)
	case *Not:
		c := convertPattern(visited, v.Pattern)
		return ast.NewNot(c)
	case *ZAny:
		return ast.NewZAny()
	case *Optional:
		c := convertPattern(visited, v.Pattern)
		return ast.NewOptional(c)
	case *Contains:
		c := convertPattern(visited, v.Pattern)
		return ast.NewContains(c)
	case *Interleave:
		left := convertPattern(visited, v.Left)
		right := convertPattern(visited, v.Right)
		return ast.NewInterleave(left, right)
	}
	panic(fmt.Sprintf("unknown Pattern typ %T", typ))
}

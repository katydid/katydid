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
	"strconv"
)

func ConvertGrammar(g *relapse.Grammar) *Pattern {
	refs := relapse.NewRefsLookup(g)
	lazyRefs := make(map[string]*Pattern)
	for name, _ := range refs {
		lazyRefs[name] = &Pattern{}
	}
	for name, _ := range refs {
		newp := newPattern(lazyRefs, refs[name])
		lazyRefs[name].head = newp.head
		lazyRefs[name].thunk = newp.thunk
	}
	return lazyRefs["main"]
}

func newPattern(refs map[string]*Pattern, p *relapse.Pattern) *Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return NewEmpty()
	case *relapse.TreeNode:
		return NewTreeNode(v.GetName(), newPattern(refs, v.GetPattern()))
	case *relapse.LeafNode:
		return NewLeafNode(v.GetExpr())
	case *relapse.Concat:
		return NewConcat(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.Or:
		return NewOr(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.And:
		return NewAnd(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.ZeroOrMore:
		return NewZeroOrMore(newPattern(refs, v.GetPattern()))
	case *relapse.Not:
		return NewNot(newPattern(refs, v.GetPattern()))
	case *relapse.Reference:
		return NewLazyPattern(refs[v.GetName()])
	case *relapse.ZAny:
		return NewZAny()
	case *relapse.Contains:
		return NewContains(newPattern(refs, v.GetPattern()))
	case *relapse.Optional:
		return NewOptional(newPattern(refs, v.GetPattern()))
	case *relapse.Interleave:
		return NewInterleave(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	}
	panic(fmt.Sprintf("unknown Pattern typ %T", typ))
}

func ConvertPattern(p *Pattern) *relapse.Grammar {
	refs := make(map[string]*relapse.Pattern)
	visited := make(map[*Pattern]string)
	visited[p] = "main"
	changed := true
	s := &smallStr{}
	for changed {
		changed = false
		for p, name := range visited {
			if _, ok := refs[name]; !ok {
				headp := p
				if p.head == nil {
					headp = &Pattern{head: p.thunk()}
				}
				visited[headp] = name
				newp := convertPattern(visited, headp, s.newStr)
				refs[name] = newp
				changed = true
			}
		}
	}
	removeEqualRefs(refs)
	return relapse.NewGrammar(refs)
}

type smallStr struct {
	i int
}

func (this *smallStr) newStr() string {
	this.i++
	return "ref" + strconv.Itoa(this.i)
}

func convertPattern(visited map[*Pattern]string, p *Pattern, newStr func() string) *relapse.Pattern {
	if p.head == nil {
		if _, ok := visited[p]; !ok {
			name := newStr()
			visited[p] = name
		}
		return relapse.NewReference(visited[p])
	}
	typ := p.head.GetValue()
	switch v := typ.(type) {
	case *Empty:
		return relapse.NewEmpty()
	case *Node:
		if v.Name != nil {
			c := convertPattern(visited, v.Pattern, newStr)
			return relapse.NewTreeNode(v.Name, c)
		} else {
			return relapse.NewLeafNode(v.Expr)
		}
	case *Concat:
		left := convertPattern(visited, v.Left, newStr)
		right := convertPattern(visited, v.Right, newStr)
		return relapse.NewConcat(left, right)
	case *Or:
		left := convertPattern(visited, v.Left, newStr)
		right := convertPattern(visited, v.Right, newStr)
		return relapse.NewOr(left, right)
	case *And:
		left := convertPattern(visited, v.Left, newStr)
		right := convertPattern(visited, v.Right, newStr)
		return relapse.NewAnd(left, right)
	case *ZeroOrMore:
		c := convertPattern(visited, v.Pattern, newStr)
		return relapse.NewZeroOrMore(c)
	case *Not:
		c := convertPattern(visited, v.Pattern, newStr)
		return relapse.NewNot(c)
	case *ZAny:
		return relapse.NewZAny()
	case *Optional:
		c := convertPattern(visited, v.Pattern, newStr)
		return relapse.NewOptional(c)
	case *Contains:
		c := convertPattern(visited, v.Pattern, newStr)
		return relapse.NewContains(c)
	case *Interleave:
		left := convertPattern(visited, v.Left, newStr)
		right := convertPattern(visited, v.Right, newStr)
		return relapse.NewInterleave(left, right)
	}
	panic(fmt.Sprintf("unknown Pattern typ %T", typ))
}

func findEqual(refs relapse.RefLookup) (string, string, bool) {
	for name1 := range refs {
		for name2 := range refs {
			if name1 == name2 {
				continue
			}
			if refs[name1].Equal(refs[name2]) {
				if name2 == "main" {
					//always return main first
					return name2, name1, true
				}
				return name1, name2, true
			}
		}
	}
	return "", "", false
}

func replaceRef(p *relapse.Pattern, prev, next string) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty, *relapse.LeafNode, *relapse.ZAny:
		return
	case *relapse.TreeNode:
		replaceRef(v.GetPattern(), prev, next)
	case *relapse.Concat:
		replaceRef(v.GetLeftPattern(), prev, next)
		replaceRef(v.GetRightPattern(), prev, next)
	case *relapse.Or:
		replaceRef(v.GetLeftPattern(), prev, next)
		replaceRef(v.GetRightPattern(), prev, next)
	case *relapse.And:
		replaceRef(v.GetLeftPattern(), prev, next)
		replaceRef(v.GetRightPattern(), prev, next)
	case *relapse.ZeroOrMore:
		replaceRef(v.GetPattern(), prev, next)
	case *relapse.Reference:
		if v.Name == prev {
			v.Name = next
		}
	case *relapse.Not:
		replaceRef(v.GetPattern(), prev, next)
	case *relapse.Contains:
		replaceRef(v.GetPattern(), prev, next)
	case *relapse.Optional:
		replaceRef(v.GetPattern(), prev, next)
	case *relapse.Interleave:
		replaceRef(v.GetLeftPattern(), prev, next)
		replaceRef(v.GetRightPattern(), prev, next)
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", typ))
	}
}

func replaceRefs(refs relapse.RefLookup, prev, next string) {
	for name := range refs {
		replaceRef(refs[name], prev, next)
	}
}

func removeEqualRefs(refs relapse.RefLookup) {
	name1, name2, ok := findEqual(refs)
	for ok {
		delete(refs, name2)
		replaceRefs(refs, name2, name1)
		name1, name2, ok = findEqual(refs)
	}
}

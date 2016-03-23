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
	"github.com/katydid/katydid/relapse"
)

//TODO what about derived left recursion
//TODO what about right recursion when left is nullable
func HasLeftRecursion(g *relapse.Grammar) bool {
	refs := relapse.NewRefsLookup(g)
	visited := make(map[*relapse.Pattern]bool)
	return hasLeftRecursion(visited, refs, refs["main"])
}

func hasLeftRecursion(visited map[*relapse.Pattern]bool, refs relapse.RefLookup, p *relapse.Pattern) bool {
	if _, ok := visited[p]; ok {
		return true
	}
	visited[p] = true
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return false
	case *relapse.TreeNode:
		return false
	case *relapse.LeafNode:
		return false
	case *relapse.Concat:
		return hasLeftRecursion(visited, refs, v.GetLeftPattern())
	case *relapse.Or:
		return hasLeftRecursion(visited, refs, v.GetLeftPattern())
	case *relapse.And:
		return hasLeftRecursion(visited, refs, v.GetLeftPattern())
	case *relapse.ZeroOrMore:
		return hasLeftRecursion(visited, refs, v.GetPattern())
	case *relapse.Reference:
		return hasLeftRecursion(visited, refs, refs[v.GetName()])
	case *relapse.Not:
		return hasLeftRecursion(visited, refs, v.GetPattern())
	case *relapse.ZAny:
		return false
	case *relapse.Contains:
		return hasLeftRecursion(visited, refs, v.GetPattern())
	case *relapse.Optional:
		return false //TODO
	case *relapse.Interleave:
		return false //TODO
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func HasRecursion(g *relapse.Grammar) bool {
	refs := relapse.NewRefsLookup(g)
	for name, _ := range refs {
		visited := make(map[*relapse.Pattern]bool)
		if hasRecursion(visited, refs, refs[name]) {
			return true
		}
	}
	return false
}

func hasRecursion(visited map[*relapse.Pattern]bool, refs relapse.RefLookup, p *relapse.Pattern) bool {
	if _, ok := visited[p]; ok {
		return true
	}
	visited[p] = true
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return false
	case *relapse.TreeNode:
		return false
	case *relapse.LeafNode:
		return false
	case *relapse.Concat:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *relapse.Or:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *relapse.And:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return hasRecursion(visited, refs, v.GetPattern())
	case *relapse.Reference:
		return hasRecursion(visited, refs, refs[v.GetName()])
	case *relapse.Not:
		return hasRecursion(visited, refs, v.GetPattern())
	case *relapse.ZAny:
		return false
	case *relapse.Contains:
		return hasRecursion(visited, refs, v.GetPattern())
	case *relapse.Optional:
		return hasRecursion(visited, refs, v.GetPattern())
	case *relapse.Interleave:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

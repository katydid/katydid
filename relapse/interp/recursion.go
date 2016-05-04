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

//HasRecursion returns whether the grammar contains any references that does not have a TreeNode pattern in between.
//For example:
//
//  #main = @main
//  #main = (A:* | @main)
//
//Recursion can still be used when a TreeNode pattern is placed between references, for example:
//
//  #main = (A:@main | <empty>)
func HasRecursion(g *ast.Grammar) bool {
	refs := ast.NewRefLookup(g)
	for name, _ := range refs {
		visited := make(map[*ast.Pattern]bool)
		if hasRecursion(visited, refs, refs[name]) {
			return true
		}
	}
	return false
}

func hasRecursion(visited map[*ast.Pattern]bool, refs ast.RefLookup, p *ast.Pattern) bool {
	if _, ok := visited[p]; ok {
		return true
	}
	visited[p] = true
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return false
	case *ast.TreeNode:
		return false
	case *ast.LeafNode:
		return false
	case *ast.Concat:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *ast.Or:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *ast.And:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	case *ast.ZeroOrMore:
		return hasRecursion(visited, refs, v.GetPattern())
	case *ast.Reference:
		return hasRecursion(visited, refs, refs[v.GetName()])
	case *ast.Not:
		return hasRecursion(visited, refs, v.GetPattern())
	case *ast.ZAny:
		return false
	case *ast.Contains:
		return hasRecursion(visited, refs, v.GetPattern())
	case *ast.Optional:
		return hasRecursion(visited, refs, v.GetPattern())
	case *ast.Interleave:
		return hasRecursion(visited, refs, v.GetLeftPattern()) || hasRecursion(visited, refs, v.GetRightPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

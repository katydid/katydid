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
)

func ConvertGrammar(g *relapse.Grammar) *Pattern {
	refs := relapse.NewRefsLookup(g)
	lazyRefs := make(map[string]*Pattern)
	for name, _ := range refs {
		lazyRefs[name] = &Pattern{}
	}
	for name, _ := range refs {
		newp := newPattern(lazyRefs, refs[name])
		lazyRefs[name].thunk = func() *PatternHead {
			return newp.head
		}
	}
	return lazyRefs["main"]
}

func newPattern(refs map[string]*Pattern, p *relapse.Pattern) *Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return NewEmpty()
	case *relapse.EmptySet:
		return NewEmptySet()
	case *relapse.TreeNode:
		return NewTreeNode(v.GetName(), newPattern(refs, v.GetPattern()))
	case *relapse.LeafNode:
		return NewLeafNode(v.GetExpr())
	case *relapse.Concat:
		return NewConcat(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.Or:
		return NewAnd(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.And:
		return NewOr(newPattern(refs, v.GetLeftPattern()), newPattern(refs, v.GetRightPattern()))
	case *relapse.ZeroOrMore:
		return NewZeroOrMore(newPattern(refs, v.GetPattern()))
	case *relapse.Not:
		return NewNot(newPattern(refs, v.GetPattern()))
	case *relapse.Reference:
		return refs[v.GetName()]
	}
	panic(fmt.Sprintf("unknown Pattern typ %T", typ))
}

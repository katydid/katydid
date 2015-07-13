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

// Returns true if there exists a tree that can satisfy the Grammar.
// This function may return some false positives.
func Satisfiable(g *relapse.Grammar) bool {
	refs := newRefsLookup(g)
	p := refs["main"]
	return sat(refs, p)
}

func satName(nameExpr *relapse.NameExpr) bool {
	typ := nameExpr.GetValue()
	switch v := typ.(type) {
	case *relapse.Name:
		return true
	case *relapse.AnyName:
		return true
	case *relapse.AnyNameExcept:
		return true // TODO for example:
		// - AnyNameExcept(Any) == false
		// - AnyNameExcept(NameChoice(Any, Name("a"))) == false
		// - AnyNameExcept(AnyNameExcept(Any)) == true
	case *relapse.NameChoice:
		return satName(v.GetLeft()) || satName(v.GetRight())
	}
	panic(fmt.Sprintf("unknown nameExpr typ %T", typ))
}

// This function returns mostly false positives, given the nature of user defined functions.
func satLeaf(expr *expr.Expr) bool {
	if expr.Terminal != nil && expr.GetTerminal().BoolValue != nil {
		return expr.GetTerminal().GetBoolValue()
	}
	//TODO all functions without any variables should be evaluatable.
	return true //TODO some standard functions should be evaluatable.
}

func sat(refs RefLookup, p *relapse.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return true
	case *relapse.EmptySet:
		return false
	case *relapse.TreeNode:
		return satName(v.GetName()) && sat(refs, v.GetPattern())
	case *relapse.LeafNode:
		return satLeaf(v.GetExpr())
	case *relapse.Concat:
		return sat(refs, v.GetLeftPattern()) && sat(refs, v.GetRightPattern())
	case *relapse.Or:
		return sat(refs, v.GetLeftPattern()) || sat(refs, v.GetRightPattern())
	case *relapse.And:
		return sat(refs, v.GetLeftPattern()) && sat(refs, v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return true
	case *relapse.Reference:
		return sat(refs, refs[v.GetName()])
	case *relapse.Not:
		return true // TODO
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

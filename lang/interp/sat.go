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

// Returns true if there exists a tree that can satisfy the Grammar.
// This function may return some false positives.
func Satisfiable(g *lang.Grammar) bool {
	refs := newRefsLookup(g)
	p := refs["main"]
	return sat(refs, p)
}

func satName(nameExpr *lang.NameExpr) bool {
	typ := nameExpr.GetValue()
	switch v := typ.(type) {
	case *lang.Name:
		return true
	case *lang.AnyName:
		return true
	case *lang.AnyNameExcept:
		return true // TODO for example:
		// - AnyNameExcept(Any) == false
		// - AnyNameExcept(NameChoice(Any, Name("a"))) == false
		// - AnyNameExcept(AnyNameExcept(Any)) == true
	case *lang.NameChoice:
		return satName(v.GetLeft()) || satName(v.GetRight())
	}
	panic(fmt.Sprintf("unknown nameExpr typ %T", typ))
}

// This function returns mostly false positives, given the nature of user defined functions.
func satLeaf(expr *expr.Expr) bool {
	return true //TODO some standard functions should be evaluatable.
}

func sat(refs RefLookup, p *lang.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return true
	case *lang.EmptySet:
		return false
	case *lang.TreeNode:
		return satName(v.GetName()) && sat(refs, v.GetPattern())
	case *lang.LeafNode:
		return satLeaf(v.GetExpr())
	case *lang.Concat:
		return sat(refs, v.GetLeftPattern()) && sat(refs, v.GetRightPattern())
	case *lang.Or:
		return sat(refs, v.GetLeftPattern()) || sat(refs, v.GetRightPattern())
	case *lang.And:
		return sat(refs, v.GetLeftPattern()) && sat(refs, v.GetRightPattern())
	case *lang.ZeroOrMore:
		return true
	case *lang.Reference:
		return sat(refs, refs[v.GetName()])
	case *lang.Not:
		return true // TODO
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

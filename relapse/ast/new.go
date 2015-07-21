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

package relapse

import (
	"github.com/katydid/katydid/expr/ast"
)

type RefLookup map[string]*Pattern

func NewRefsLookup(g *Grammar) RefLookup {
	decls := g.GetPatternDecls()
	refs := make(RefLookup, len(decls))
	for _, d := range decls {
		refs[d.Name] = d.Pattern
	}
	return refs
}

func NewGrammar(m map[string]*Pattern) *Grammar {
	ps := make([]*PatternDecl, 0, len(m))
	first := true
	for name, p := range m {
		before := &expr.Space{Space: []string{"\n"}}
		if first {
			before = nil
			first = false
		}
		ps = append(ps, &PatternDecl{
			Before:  before,
			Name:    name,
			Eq:      newEqual(),
			Pattern: p,
		})
	}
	return &Grammar{PatternDecls: ps}
}

func NewEmpty() *Pattern {
	return &Pattern{
		Empty: &Empty{},
	}
}

func NewEmptySet() *Pattern {
	return &Pattern{
		EmptySet: &EmptySet{},
	}
}

func newOpenParen() *expr.Keyword {
	return &expr.Keyword{Value: "("}
}

func newCloseParen() *expr.Keyword {
	return &expr.Keyword{Value: ")"}
}

func newComma() *expr.Keyword {
	return &expr.Keyword{Value: ","}
}

func newEqual() *expr.Keyword {
	return &expr.Keyword{Value: "="}
}

func NewName(name string) *NameExpr {
	return &NameExpr{
		Name: &Name{
			OpenParen:  newOpenParen(),
			Name:       name,
			CloseParen: newCloseParen(),
		},
	}
}

func NewAnyName() *NameExpr {
	return &NameExpr{
		AnyName: &AnyName{},
	}
}

func NewAnyNameExcept(name *NameExpr) *NameExpr {
	return &NameExpr{
		AnyNameExcept: &AnyNameExcept{
			OpenParen:  newOpenParen(),
			Except:     name,
			CloseParen: newCloseParen(),
		},
	}
}

func NewNameChoice(left, right *NameExpr) *NameExpr {
	return &NameExpr{
		NameChoice: &NameChoice{
			OpenParen:  newOpenParen(),
			Left:       left,
			Comma:      newComma(),
			Right:      right,
			CloseParen: newCloseParen(),
		},
	}
}

func NewTreeNode(name *NameExpr, pattern *Pattern) *Pattern {
	return &Pattern{
		TreeNode: &TreeNode{
			OpenParen:  newOpenParen(),
			Name:       name,
			Comma:      newComma(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	return &Pattern{
		LeafNode: &LeafNode{
			OpenParen:  newOpenParen(),
			Expr:       expr,
			CloseParen: newCloseParen(),
		},
	}
}

func NewConcat(left, right *Pattern) *Pattern {
	return &Pattern{
		Concat: &Concat{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewOr(left, right *Pattern) *Pattern {
	return &Pattern{
		Or: &Or{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewAnd(left, right *Pattern) *Pattern {
	return &Pattern{
		And: &And{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewZeroOrMore(pattern *Pattern) *Pattern {
	return &Pattern{
		ZeroOrMore: &ZeroOrMore{
			OpenParen:  newOpenParen(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

func NewReference(name string) *Pattern {
	return &Pattern{
		Reference: &Reference{
			OpenParen:  newOpenParen(),
			Name:       name,
			CloseParen: newCloseParen(),
		},
	}
}

func NewNot(pattern *Pattern) *Pattern {
	return &Pattern{
		Not: &Not{
			OpenParen:  newOpenParen(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

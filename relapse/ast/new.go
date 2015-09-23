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
	if g.TopPattern != nil {
		refs["main"] = g.TopPattern
	}
	return refs
}

func NewGrammar(m map[string]*Pattern) *Grammar {
	ps := make([]*PatternDecl, 0, len(m))
	first := true
	g := &Grammar{}
	for name, _ := range m {
		if name == "main" {
			g.TopPattern = m[name]
		}
		before := &expr.Space{Space: []string{"\n"}}
		if first {
			before = nil
			first = false
		}
		ps = append(ps, &PatternDecl{
			At:      newAt(),
			Before:  before,
			Name:    name,
			Eq:      newEqual(),
			Pattern: m[name],
		})
	}
	g.PatternDecls = ps
	return g
}

func NewEmpty() *Pattern {
	return &Pattern{
		Empty: &Empty{&expr.Keyword{Value: "<empty>"}},
	}
}

func NewEmptySet() *Pattern {
	return &Pattern{
		EmptySet: &EmptySet{&expr.Keyword{Value: "<emptyset>"}},
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

func newHashTag() *expr.Keyword {
	return &expr.Keyword{Value: "#"}
}

func newAmpersand() *expr.Keyword {
	return &expr.Keyword{Value: "&"}
}

func newPipe() *expr.Keyword {
	return &expr.Keyword{Value: "|"}
}

func newOpenBracket() *expr.Keyword {
	return &expr.Keyword{Value: "["}
}

func newCloseBracket() *expr.Keyword {
	return &expr.Keyword{Value: "]"}
}

func newColon() *expr.Keyword {
	return &expr.Keyword{Value: ":"}
}

func newExclamation() *expr.Keyword {
	return &expr.Keyword{Value: "!"}
}

func newOpenCurly() *expr.Keyword {
	return &expr.Keyword{Value: "{"}
}

func newCloseCurly() *expr.Keyword {
	return &expr.Keyword{Value: "}"}
}

func newStar() *expr.Keyword {
	return &expr.Keyword{Value: "*"}
}

func newUnderscore() *expr.Keyword {
	return &expr.Keyword{Value: "_"}
}

func newTilde() *expr.Keyword {
	return &expr.Keyword{Value: "~"}
}

func newDot() *expr.Keyword {
	return &expr.Keyword{Value: "."}
}

func newAt() *expr.Keyword {
	return &expr.Keyword{Value: "@"}
}

func newRightArrow() *expr.Keyword {
	return &expr.Keyword{Value: "->"}
}

func NewName(name string) *NameExpr {
	return &NameExpr{
		Name: &Name{
			Name: name,
		},
	}
}

func NewAnyName() *NameExpr {
	return &NameExpr{
		AnyName: &AnyName{Underscore: newUnderscore()},
	}
}

func NewAnyNameExcept(name *NameExpr) *NameExpr {
	return &NameExpr{
		AnyNameExcept: &AnyNameExcept{
			Exclamation: newExclamation(),
			OpenParen:   newOpenParen(),
			Except:      name,
			CloseParen:  newCloseParen(),
		},
	}
}

func NewNameChoice(names ...*NameExpr) *NameExpr {
	if len(names) == 0 {
		return nil
	}
	if len(names) == 1 {
		return names[0]
	}
	return &NameExpr{
		NameChoice: &NameChoice{
			OpenParen:  newOpenParen(),
			Left:       names[0],
			Pipe:       newPipe(),
			Right:      newNameChoice(names[1:]),
			CloseParen: newCloseParen(),
		},
	}
}

func newNameChoice(names []*NameExpr) *NameExpr {
	if len(names) == 1 {
		return names[0]
	}
	return &NameExpr{
		NameChoice: &NameChoice{
			Left:  names[0],
			Pipe:  newPipe(),
			Right: newNameChoice(names[1:]),
		},
	}
}

func NewTreeNode(name *NameExpr, pattern *Pattern) *Pattern {
	switch pattern.GetValue().(type) {
	case *Concat, *WithSomeOr, *WithSomeAnd:
		return &Pattern{
			TreeNode: &TreeNode{
				Name:    name,
				Pattern: pattern,
			},
		}
	}
	return &Pattern{
		TreeNode: &TreeNode{
			Name:    name,
			Colon:   newColon(),
			Pattern: pattern,
		},
	}
}

func NewWithSomeTreeNode(name *NameExpr, pattern *Pattern) *Pattern {
	return &Pattern{
		WithSomeTreeNode: &WithSomeTreeNode{
			Dot:     newDot(),
			Pattern: NewTreeNode(name, pattern),
		},
	}
}

func NewLeafNode(name *NameExpr, expr *expr.Expr) *Pattern {
	return &Pattern{
		TreeNode: &TreeNode{
			Name: name,
			Pattern: &Pattern{
				LeafNode: &LeafNode{
					RightArrow: newRightArrow(),
					Expr:       expr,
				},
			},
		},
	}
}

func NewConcat(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Concat: &Concat{
			OpenBracket:  newOpenBracket(),
			LeftPattern:  patterns[0],
			Comma:        newComma(),
			RightPattern: newConcat(patterns[1:]),
			CloseBracket: newCloseBracket(),
		},
	}
}

func newConcat(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Concat: &Concat{
			LeftPattern:  patterns[0],
			Comma:        newComma(),
			RightPattern: newConcat(patterns[1:]),
		},
	}
}

func NewOr(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Or: &Or{
			OpenParen:    newOpenParen(),
			LeftPattern:  patterns[0],
			Pipe:         newPipe(),
			RightPattern: newOr(patterns[1:]),
			CloseParen:   newCloseParen(),
		},
	}
}

func newOr(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Or: &Or{
			LeftPattern:  patterns[0],
			Pipe:         newPipe(),
			RightPattern: newOr(patterns[1:]),
		},
	}
}

func NewWithSomeOr(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		WithSomeOr: &WithSomeOr{
			OpenCurly:    newOpenCurly(),
			LeftPattern:  patterns[0],
			Pipe:         newPipe(),
			RightPattern: newWithSomeOr(patterns[1:]),
			CloseCurly:   newCloseCurly(),
		},
	}
}

func newWithSomeOr(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		WithSomeOr: &WithSomeOr{
			LeftPattern:  patterns[0],
			Pipe:         newPipe(),
			RightPattern: newWithSomeOr(patterns[1:]),
		},
	}
}

func NewAnd(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		And: &And{
			OpenParen:    newOpenParen(),
			LeftPattern:  patterns[0],
			Ampersand:    newAmpersand(),
			RightPattern: newAnd(patterns[1:]),
			CloseParen:   newCloseParen(),
		},
	}
}

func newAnd(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		And: &And{
			LeftPattern:  patterns[0],
			Ampersand:    newAmpersand(),
			RightPattern: newAnd(patterns[1:]),
		},
	}
}

func NewWithSomeAnd(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		WithSomeAnd: &WithSomeAnd{
			OpenCurly:    newOpenCurly(),
			LeftPattern:  patterns[0],
			Ampersand:    newAmpersand(),
			RightPattern: newWithSomeAnd(patterns[1:]),
			CloseCurly:   newCloseCurly(),
		},
	}
}

func newWithSomeAnd(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		WithSomeAnd: &WithSomeAnd{
			LeftPattern:  patterns[0],
			Ampersand:    newAmpersand(),
			RightPattern: newWithSomeAnd(patterns[1:]),
		},
	}
}

func NewZeroOrMore(pattern *Pattern) *Pattern {
	return &Pattern{
		ZeroOrMore: &ZeroOrMore{
			OpenParen:  newOpenParen(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
			Star:       newStar(),
		},
	}
}

func NewReference(name string) *Pattern {
	return &Pattern{
		Reference: &Reference{
			HashTag: newHashTag(),
			Name:    name,
		},
	}
}

func NewNot(pattern *Pattern) *Pattern {
	return &Pattern{
		Not: &Not{
			Exclamation: newExclamation(),
			OpenParen:   newOpenParen(),
			Pattern:     pattern,
			CloseParen:  newCloseParen(),
		},
	}
}

func NewZAny() *Pattern {
	return &Pattern{
		ZAny: &ZAny{
			Star: newStar(),
		},
	}
}

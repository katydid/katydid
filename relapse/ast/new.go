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
		} else {
			before := &expr.Space{Space: []string{"\n"}}
			if first {
				before = nil
				first = false
			}
			ps = append(ps, &PatternDecl{
				Hash:    newHash(),
				Before:  before,
				Name:    name,
				Eq:      newEqual(),
				Pattern: m[name],
			})
		}
	}
	g.PatternDecls = ps
	return g
}

func NewPatternDecl(name string, p *Pattern) *PatternDecl {
	return &PatternDecl{
		Hash:    newHash(),
		Name:    name,
		Eq:      newEqual(),
		Pattern: p,
	}
}

func NewEmpty() *Pattern {
	return &Pattern{
		Empty: &Empty{&expr.Keyword{Value: "<empty>"}},
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

func newHash() *expr.Keyword {
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

func newQuestionMark() *expr.Keyword {
	return &expr.Keyword{Value: "?"}
}

func newSemiColon() *expr.Keyword {
	return &expr.Keyword{Value: ";"}
}

func NewTreeNode(name *expr.NameExpr, pattern *Pattern) *Pattern {
	switch pattern.GetValue().(type) {
	case *Concat:
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

func NewContains(pattern *Pattern) *Pattern {
	return &Pattern{
		Contains: &Contains{
			Dot:     newDot(),
			Pattern: pattern,
		},
	}
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	if expr.BuiltIn != nil {
		return &Pattern{
			LeafNode: &LeafNode{
				Expr: expr,
			},
		}
	}
	return &Pattern{
		LeafNode: &LeafNode{
			RightArrow: newRightArrow(),
			Expr:       expr,
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
			At:   newAt(),
			Name: name,
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

func NewOptional(pattern *Pattern) *Pattern {
	return &Pattern{
		Optional: &Optional{
			OpenParen:    newOpenParen(),
			Pattern:      pattern,
			CloseParen:   newCloseParen(),
			QuestionMark: newQuestionMark(),
		},
	}
}

func NewInterleave(patterns ...*Pattern) *Pattern {
	if len(patterns) == 0 {
		return nil
	}
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Interleave: &Interleave{
			OpenCurly:    newOpenCurly(),
			LeftPattern:  patterns[0],
			SemiColon:    newSemiColon(),
			RightPattern: newInterleave(patterns[1:]),
			CloseCurly:   newCloseCurly(),
		},
	}
}

func newInterleave(patterns []*Pattern) *Pattern {
	if len(patterns) == 1 {
		return patterns[0]
	}
	return &Pattern{
		Interleave: &Interleave{
			LeftPattern:  patterns[0],
			SemiColon:    newSemiColon(),
			RightPattern: newInterleave(patterns[1:]),
		},
	}
}

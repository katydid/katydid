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

package ast

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/relapse/types"
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
			before := &Space{Space: []string{"\n"}}
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
		Empty: &Empty{&Keyword{Value: "<empty>"}},
	}
}

func NewTreeNode(name *NameExpr, pattern *Pattern) *Pattern {
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

func NewLeafNode(e *Expr) *Pattern {
	if e.BuiltIn != nil {
		return &Pattern{
			LeafNode: &LeafNode{
				Expr: e,
			},
		}
	}
	return &Pattern{
		LeafNode: &LeafNode{
			Expr: e,
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

//NewNestedFunction returns a function expression given a name and a list of parameters.
func NewNestedFunction(name string, params ...*Expr) *Expr {
	return &Expr{Function: &Function{
		Name:       name,
		OpenParen:  newOpenParen(),
		Params:     params,
		CloseParen: newCloseParen(),
	}}
}

//NewVar return a variable expression given a type.
func NewVar(typ types.Type) *Expr {
	return &Expr{
		Terminal: &Terminal{
			Variable: &Variable{
				Type: typ,
			},
		},
	}
}

//NewDoubleVar returns a new variable expression of type double
func NewDoubleVar() *Expr {
	return NewVar(types.SINGLE_DOUBLE)
}

//NewIntVar returns a new variable expression of type int
func NewIntVar() *Expr {
	return NewVar(types.SINGLE_INT)
}

//NewUintVar returns a new variable expression of type uint
func NewUintVar() *Expr {
	return NewVar(types.SINGLE_UINT)
}

//NewBoolVar returns a new variable expression of type bool
func NewBoolVar() *Expr {
	return NewVar(types.SINGLE_BOOL)
}

//NewStringVar returns a new variable expression of type string
func NewStringVar() *Expr {
	return NewVar(types.SINGLE_STRING)
}

//NewBytesVar returns a new variable expression of type []byte
func NewBytesVar() *Expr {
	return NewVar(types.SINGLE_BYTES)
}

//NewDoubleConst returns a new terminal expression containing the given double value.
func NewDoubleConst(d float64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			DoubleValue: proto.Float64(d),
		},
	}
}

//NewIntConst returns a new terminal expression containing the given int value.
func NewIntConst(i int64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			IntValue: proto.Int64(i),
		},
	}
}

//NewUintConst returns a new terminal expression containing the given uint value.
func NewUintConst(i uint64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			UintValue: proto.Uint64(i),
		},
	}
}

//NewTrue returns a new terminal expression containing a true value.
func NewTrue() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(true),
	}
}

//NewFalse returns a new terminal expression containing a false value.
func NewFalse() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(false),
	}
}

//NewStringConst returns a new terminal expression containing the given string value.
func NewStringConst(s string) *Expr {
	return &Expr{
		Terminal: &Terminal{
			StringValue: proto.String(s),
		},
	}
}

//NewBytesConst returns a new terminal expression containing the given bytes value.
func NewBytesConst(buf []byte) *Expr {
	return &Expr{
		Terminal: &Terminal{
			BytesValue: buf,
		},
	}
}

//NewList returns a typed list of expressions.
func NewList(typ types.Type, elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  typ,
			Elems: elems,
		},
	}
}

//NewDoubleList returns a list of expressions, each of type double.
func NewDoubleList(elems ...*Expr) *Expr {
	return NewList(types.LIST_DOUBLE, elems...)
}

//NewIntList returns a list of expressions, each of type int.
func NewIntList(elems ...*Expr) *Expr {
	return NewList(types.LIST_INT, elems...)
}

//NewUintList returns a list of expressions, each of type uint.
func NewUintList(elems ...*Expr) *Expr {
	return NewList(types.LIST_UINT, elems...)
}

//NewBoolList returns a list of expressions, each of type bool.
func NewBoolList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BOOL, elems...)
}

//NewStringList returns a list of expressions, each of type string.
func NewStringList(elems ...*Expr) *Expr {
	return NewList(types.LIST_STRING, elems...)
}

//NewBytesList returns a list of expressions, each of type []byte.
func NewBytesList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BYTES, elems...)
}

//NewDoubleName returns a name expression containing the given double value.
func NewDoubleName(name float64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			DoubleValue: &name,
		},
	}
}

//NewIntName returns a name expression containing the given int value.
func NewIntName(name int64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			IntValue: &name,
		},
	}
}

//NewUintName returns a name expression containing the given uint value.
func NewUintName(name uint64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			UintValue: &name,
		},
	}
}

//NewBoolName returns a name expression containing the given bool value.
func NewBoolName(name bool) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BoolValue: &name,
		},
	}
}

//NewStringName returns a name expression containing the given string value.
func NewStringName(name string) *NameExpr {
	return &NameExpr{
		Name: &Name{
			StringValue: &name,
		},
	}
}

//NewBytesName returns a name expression containing the given []byte value.
func NewBytesName(name []byte) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BytesValue: name,
		},
	}
}

//NewAnyName returns a name expression that represents the any name expression.
func NewAnyName() *NameExpr {
	return &NameExpr{
		AnyName: &AnyName{Underscore: newUnderscore()},
	}
}

//NewAnyNameExcept returns a name expression that represents any name except the given name expression.
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

//NewNameChoice returns a name expression which represents of choice of the list of given name expressions.
//The function can also handle zero or even one name expression.
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

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

//RefLookup represents a relapse grammar as a map of references.
type RefLookup map[string]*Pattern

//NewRefLookup converts a grammar into a reference lookup map.
func NewRefLookup(g *Grammar) RefLookup {
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

//NewGrammar converts a refenence lookup map into a Grammar.
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

//NewPatternDecl creates a new pattern declaration.
//  #name = pattern
func NewPatternDecl(name string, pattern *Pattern) *PatternDecl {
	return &PatternDecl{
		Hash:    newHash(),
		Name:    name,
		Eq:      newEqual(),
		Pattern: pattern,
	}
}

//NewEmpty returns a new empty pattern.
//  <empty>
func NewEmpty() *Pattern {
	return &Pattern{
		Empty: &Empty{&Keyword{Value: "<empty>"}},
	}
}

//NewTreeNode returns a new TreeNode pattern.
//Depending on what is appropriate the relapse string could be one of the following:
//  name: pattern
//  name pattern
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

//NewContains returns a new Contains pattern.
//  .pattern
func NewContains(pattern *Pattern) *Pattern {
	return &Pattern{
		Contains: &Contains{
			Dot:     newDot(),
			Pattern: pattern,
		},
	}
}

//NewLeafNode returns a new LeafNode pattern.
func NewLeafNode(e *Expr) *Pattern {
	return &Pattern{
		LeafNode: &LeafNode{
			Expr: e,
		},
	}
}

//NewConcat returns a new Concat pattern.
//If the number of patterns provided is:
//
//0, nil is returned;
//
//1, the input pattern is returned;
//
//2, the concatenated pattern is returned;
//  [pattern[0], pattern[1]]
//> 2, the left curried concatenated pattern is returned.
//  [pattern[0], [pattern[1], [...]]]
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

//NewOr returns a new Or pattern.
//If the number of patterns provided is:
//
//0, nil is returned;
//
//1, the input pattern is returned;
//
//2, the ored pattern is returned;
//  (pattern[0] | pattern[1])
//> 2, the left curried ored pattern is returned.
//  (pattern[0] | (pattern[1] | (...)))
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

//NewAnd returns a new And pattern.
//If the number of patterns provided is:
//
//0, nil is returned;
//
//1, the input pattern is returned;
//
//2, the anded pattern is returned;
//  (pattern[0] & pattern[1])
//> 2, the left curried anded pattern is returned.
//  (pattern[0] & (pattern[1] & (...)))
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

//NewZeroOrMore returns a new ZeroOrMore pattern.
//  (pattern)*
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

//NewReference returns a new Reference pattern.
//  @name
func NewReference(name string) *Pattern {
	return &Pattern{
		Reference: &Reference{
			At:   newAt(),
			Name: name,
		},
	}
}

//NewNot returns a new Not pattern.
//  !(pattern)
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

//NewZAny returns a new ZAny pattern.
//  *
func NewZAny() *Pattern {
	return &Pattern{
		ZAny: &ZAny{
			Star: newStar(),
		},
	}
}

//NewOptional returns a new Optional pattern.
//  (pattern)?
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

//NewInterleave returns a new Interleave pattern.
//If the number of patterns provided is:
//
//0, nil is returned;
//
//1, the input pattern is returned;
//
//2, the interleaved pattern is returned;
//  {pattern[0]; pattern[1]}
//> 2, the left curried interleaved pattern is returned.
//  {pattern[0]; {pattern[1]; {...}}}
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

//NewFunction returns a function expression given a name and a list of parameters.
//  ->name(param1, param2, ...)
//This function should be the top level function and not a nested function.
//If the parameters don't have populated Comma fields, this function will add them.
//If a parameter has a populared RightArrow field, the contents of the RightArrow field is thrown away.
func NewFunction(name string, params ...*Expr) *Expr {
	for i, p := range params {
		if p.RightArrow != nil {
			p.RightArrow = nil
		}
		if i == 0 {
			continue
		}
		if p.Comma == nil {
			p.Comma = newComma()
		}
	}
	return &Expr{
		RightArrow: newRightArrow(),
		Function: &Function{
			Name:       name,
			OpenParen:  newOpenParen(),
			Params:     params,
			CloseParen: newCloseParen(),
		},
	}
}

//NewNestedFunction returns a function expression given a name and a list of parameters.
//  name(param1, param2, ...)
//If the parameters don't have populated Comma fields, this function will add them.
//If a parameter has a populared RightArrow field, the contents of the RightArrow field is thrown away.
func NewNestedFunction(name string, params ...*Expr) *Expr {
	for i, p := range params {
		if p.RightArrow != nil {
			p.RightArrow = nil
		}
		if i == 0 {
			continue
		}
		if p.Comma == nil {
			p.Comma = newComma()
		}
	}
	return &Expr{Function: &Function{
		Name:       name,
		OpenParen:  newOpenParen(),
		Params:     params,
		CloseParen: newCloseParen(),
	}}
}

//NewVar return a variable expression given a type.
//  $typ
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
//  $double
func NewDoubleVar() *Expr {
	return NewVar(types.SINGLE_DOUBLE)
}

//NewIntVar returns a new variable expression of type int
//  $int
func NewIntVar() *Expr {
	return NewVar(types.SINGLE_INT)
}

//NewUintVar returns a new variable expression of type uint
//  $uint
func NewUintVar() *Expr {
	return NewVar(types.SINGLE_UINT)
}

//NewBoolVar returns a new variable expression of type bool
//  $bool
func NewBoolVar() *Expr {
	return NewVar(types.SINGLE_BOOL)
}

//NewStringVar returns a new variable expression of type string
//  $string
func NewStringVar() *Expr {
	return NewVar(types.SINGLE_STRING)
}

//NewBytesVar returns a new variable expression of type []byte
//  $[]byte
func NewBytesVar() *Expr {
	return NewVar(types.SINGLE_BYTES)
}

//NewDoubleConst returns a new terminal expression containing the given double value.
//  double(d)
func NewDoubleConst(d float64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			DoubleValue: proto.Float64(d),
		},
	}
}

//NewIntConst returns a new terminal expression containing the given int value.
//  int(i)
func NewIntConst(i int64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			IntValue: proto.Int64(i),
		},
	}
}

//NewUintConst returns a new terminal expression containing the given uint value.
//  uint(i)
func NewUintConst(i uint64) *Expr {
	return &Expr{
		Terminal: &Terminal{
			UintValue: proto.Uint64(i),
		},
	}
}

//NewTrue returns a new terminal expression containing a true value.
//  true
func NewTrue() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(true),
	}
}

//NewFalse returns a new terminal expression containing a false value.
//  false
func NewFalse() *Expr {
	return &Expr{
		Terminal: NewBoolTerminal(false),
	}
}

//NewStringConst returns a new terminal expression containing the given string value.
//  "s"
//  `s`
func NewStringConst(s string) *Expr {
	return &Expr{
		Terminal: &Terminal{
			StringValue: proto.String(s),
		},
	}
}

//NewBytesConst returns a new terminal expression containing the given bytes value.
//  []byte(fmt.Sprintf("%#v", buf))
func NewBytesConst(buf []byte) *Expr {
	return &Expr{
		Terminal: &Terminal{
			BytesValue: buf,
		},
	}
}

//NewList returns a typed list of expressions.
//  []typ{elems}
func NewList(typ types.Type, elems ...*Expr) *Expr {
	return &Expr{
		List: &List{
			Type:  typ,
			Elems: elems,
		},
	}
}

//NewDoubleList returns a list of expressions, each of type double.
//  []double{elems}
func NewDoubleList(elems ...*Expr) *Expr {
	return NewList(types.LIST_DOUBLE, elems...)
}

//NewIntList returns a list of expressions, each of type int.
//  []int{elems}
func NewIntList(elems ...*Expr) *Expr {
	return NewList(types.LIST_INT, elems...)
}

//NewUintList returns a list of expressions, each of type uint.
//  []uint{elems}
func NewUintList(elems ...*Expr) *Expr {
	return NewList(types.LIST_UINT, elems...)
}

//NewBoolList returns a list of expressions, each of type bool.
//  []bool{elems}
func NewBoolList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BOOL, elems...)
}

//NewStringList returns a list of expressions, each of type string.
//  []string{elems}
func NewStringList(elems ...*Expr) *Expr {
	return NewList(types.LIST_STRING, elems...)
}

//NewBytesList returns a list of expressions, each of type []byte.
//  [][]byte{elems}
func NewBytesList(elems ...*Expr) *Expr {
	return NewList(types.LIST_BYTES, elems...)
}

//NewDoubleName returns a name expression containing the given double value.
//  double(name)
func NewDoubleName(name float64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			DoubleValue: &name,
		},
	}
}

//NewIntName returns a name expression containing the given int value.
//  int(name)
func NewIntName(name int64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			IntValue: &name,
		},
	}
}

//NewUintName returns a name expression containing the given uint value.
//  uint(name)
func NewUintName(name uint64) *NameExpr {
	return &NameExpr{
		Name: &Name{
			UintValue: &name,
		},
	}
}

//NewBoolName returns a name expression containing the given bool value.
//  bool(name)
func NewBoolName(name bool) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BoolValue: &name,
		},
	}
}

//NewStringName returns a name expression containing the given string value.
//  string(name)
func NewStringName(name string) *NameExpr {
	return &NameExpr{
		Name: &Name{
			StringValue: &name,
		},
	}
}

//NewBytesName returns a name expression containing the given []byte value.
//  []byte(name)
func NewBytesName(name []byte) *NameExpr {
	return &NameExpr{
		Name: &Name{
			BytesValue: name,
		},
	}
}

//NewAnyName returns a name expression that represents the any name expression.
//  _
func NewAnyName() *NameExpr {
	return &NameExpr{
		AnyName: &AnyName{Underscore: newUnderscore()},
	}
}

//NewAnyNameExcept returns a name expression that represents any name except the given name expression.
//  !(name)
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
//If the number of names provided is:
//
//0, nil is returned;
//
//1, the input name is returned;
//
//2, the ored names is returned;
//  (names[0] | names[1])
//> 2, the left curried ored names is returned.
//  (names[0] | (names[1] | (...)))
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

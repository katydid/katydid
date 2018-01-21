//  Copyright 2017 Walter Schulze
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

package intern

import (
	"fmt"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/funcs"
	nameexpr "github.com/katydid/katydid/relapse/name"
	"github.com/katydid/katydid/relapse/parser"
)

type PatternType int

const Empty = PatternType(1)
const Node = PatternType(2)
const Concat = PatternType(4)
const Or = PatternType(5)
const And = PatternType(6)
const ZeroOrMore = PatternType(7)
const Reference = PatternType(8)
const Not = PatternType(9)
const ZAny = PatternType(10)
const Contains = PatternType(11)
const Optional = PatternType(12)
const Interleave = PatternType(13)

type Pattern struct {
	Type     PatternType
	Func     funcs.Bool
	Patterns []*Pattern
	Ref      string
	hash     uint64
	nullable bool
}

func (p *Pattern) GoString() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf("&%#v", *p)
}

func (p *Pattern) NewAst() *ast.Pattern {
	pp, err := newASTPattern(p)
	if err != nil {
		panic(err)
	}
	return pp
}

func newASTPattern(p *Pattern) (*ast.Pattern, error) {
	switch p.Type {
	case Empty:
		return ast.NewEmpty(), nil
	case Node:
		if isEmpty(p.Patterns[0]) {
			exprStr := funcs.Sprint(p.Func)
			expr, err := parser.NewParser().ParseExpr(exprStr)
			if err != nil {
				return nil, err
			}
			return ast.NewLeafNode(expr), nil
		} else {
			name, err := nameexpr.FuncToName(p.Func)
			if err != nil {
				return nil, err
			}
			p, err := newASTPattern(p.Patterns[0])
			if err != nil {
				return nil, err
			}
			return ast.NewTreeNode(name, p), nil
		}
	case Concat:
		ps, err := deriveTraverseAst(newASTPattern, p.Patterns)
		if err != nil {
			return nil, err
		}
		return ast.NewConcat(ps...), nil
	case Or:
		ps, err := deriveTraverseAst(newASTPattern, p.Patterns)
		if err != nil {
			return nil, err
		}
		return ast.NewOr(ps...), nil
	case And:
		ps, err := deriveTraverseAst(newASTPattern, p.Patterns)
		if err != nil {
			return nil, err
		}
		return ast.NewAnd(ps...), nil
	case ZeroOrMore:
		pp, err := newASTPattern(p.Patterns[0])
		if err != nil {
			return nil, err
		}
		return ast.NewZeroOrMore(pp), nil
	case Reference:
		return ast.NewReference(p.Ref), nil
	case Not:
		pp, err := newASTPattern(p.Patterns[0])
		if err != nil {
			return nil, err
		}
		return ast.NewNot(pp), nil
	case ZAny:
		return ast.NewZAny(), nil
	case Contains:
		pp, err := newASTPattern(p.Patterns[0])
		if err != nil {
			return nil, err
		}
		return ast.NewContains(pp), nil
	case Optional:
		pp, err := newASTPattern(p.Patterns[0])
		if err != nil {
			return nil, err
		}
		return ast.NewOptional(pp), nil
	case Interleave:
		ps, err := deriveTraverseAst(newASTPattern, p.Patterns)
		if err != nil {
			return nil, err
		}
		return ast.NewInterleave(ps...), nil
	}
	panic(fmt.Sprintf("unknown pattern: %d", p.Type))
}

func (p *Pattern) String() string {
	if p == nil {
		return ""
	}
	st, err := newASTPattern(p)
	if err != nil {
		return fmt.Sprintf("could not get ast of: %v", *p)
	}
	return st.String()
}

func (p *Pattern) Equal(pp *Pattern) bool {
	if p.hash != pp.hash {
		return false
	}
	if p.Type != pp.Type {
		return false
	}
	if p.Func == nil && pp.Func != nil {
		return false
	}
	if p.Func != nil && pp.Func == nil {
		return false
	}
	if p.Func != nil && pp.Func != nil {
		if !funcs.Equal(p.Func, pp.Func) {
			return false
		}
	}
	if len(p.Patterns) != len(pp.Patterns) {
		return false
	}
	for i := range p.Patterns {
		if !p.Patterns[i].Equal(pp.Patterns[i]) {
			return false
		}
	}
	if p.Ref != pp.Ref {
		return false
	}
	return true
}

func (p *Pattern) SetHash() {
	h := uint64(17)
	h = 31*h + uint64(p.Type)
	if p.Func != nil {
		h = 31*h + p.Func.Hash()
	}
	for _, pattern := range p.Patterns {
		h = 31*h + pattern.hash
	}
	if len(p.Ref) > 0 {
		h = 31*h + deriveHashString(p.Ref)
	}
	if h == 0 {
		h = 1
	}
	p.hash = h
}

func (p *Pattern) Nullable() bool {
	return p.nullable
}

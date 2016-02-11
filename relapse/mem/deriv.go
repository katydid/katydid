//  Copyright 2016 Walter Schulze
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

package mem

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/relapse/nameexpr"
	"github.com/katydid/katydid/serialize"
	"io"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

//This is a naive implementation and it does not handle left recursion
func Interpret(g *relapse.Grammar, parser serialize.Parser) bool {
	refs := relapse.NewRefsLookup(g)
	mem := newMem(refs)
	start := mem.add([]*relapse.Pattern{refs["main"]})
	final := deriv(mem, start, parser)
	return mem.accept(final)
}

type Compiled interface {
	Interpret(serialize.Parser) bool
}

func (mem *mem) Interpret(parser serialize.Parser) bool {
	final := deriv(mem, mem.start, parser)
	return mem.accept(final)
}

func Compile(g *relapse.Grammar) Compiled {
	refs := relapse.NewRefsLookup(g)
	mem := newMem(refs)
	return mem
}

func escapable(patterns []*relapse.Pattern) bool {
	for _, pattern := range patterns {
		if pattern.ZAny != nil {
			continue
		}
		if pattern.Not != nil {
			if pattern.GetNot().GetPattern().ZAny != nil {
				continue
			}
		}
		return true
	}
	return false
}

func deriv(mem *mem, current state, tree serialize.Parser) state {
	for {
		if !mem.escapable(current) {
			return current
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callTree := mem.getCallTree(current)
		childState, zi := callTree.eval(tree)
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			childState = deriv(mem, childState, tree)
			tree.Up()
		}
		current = mem.getReturn(current, zi, childState)
	}
	return current
}

func nullables(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []bool {
	nulls := make([]bool, len(patterns))
	for i, p := range patterns {
		nulls[i] = interp.Nullable(refs, p)
	}
	return nulls
}

func simps(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []*relapse.Pattern {
	for i := range patterns {
		patterns[i] = interp.Simplify(refs, patterns[i])
	}
	return patterns
}

func derivCalls(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []*callable {
	res := []*callable{}
	for _, pattern := range patterns {
		cs := derivCall(refs, pattern)
		res = append(res, cs...)
	}
	return res
}

func derivCall(refs map[string]*relapse.Pattern, p *relapse.Pattern) []*callable {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return []*callable{&callable{nil, relapse.NewNot(relapse.NewZAny()), nil}}
	case *relapse.ZAny:
		return []*callable{&callable{nil, relapse.NewZAny(), nil}}
	case *relapse.TreeNode:
		b := nameexpr.NameToFunc(v.GetName())
		return []*callable{&callable{b, v.GetPattern(), relapse.NewNot(relapse.NewZAny())}}
	case *relapse.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		return []*callable{&callable{b, relapse.NewEmpty(), relapse.NewNot(relapse.NewZAny())}}
	case *relapse.Concat:
		l := derivCall(refs, v.GetLeftPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return l
		}
		r := derivCall(refs, v.GetRightPattern())
		return append(l, r...)
	case *relapse.Or:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern())
	case *relapse.And:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern())
	case *relapse.Interleave:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return derivCall(refs, v.GetPattern())
	case *relapse.Reference:
		return derivCall(refs, refs[v.GetName()])
	case *relapse.Not:
		return derivCall(refs, v.GetPattern())
	case *relapse.Contains:
		return derivCall(refs, relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())))
	case *relapse.Optional:
		return derivCall(refs, relapse.NewOr(v.GetPattern(), relapse.NewEmpty()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func derivCall2(refs map[string]*relapse.Pattern, left, right *relapse.Pattern) []*callable {
	l := derivCall(refs, left)
	r := derivCall(refs, right)
	return append(l, r...)
}

func derivReturns(refs map[string]*relapse.Pattern, originals []*relapse.Pattern, nullable []bool) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(originals))
	rest := nullable
	for i, original := range originals {
		res[i], rest = derivReturn(refs, original, rest)
	}
	return res
}

func derivReturn(refs map[string]*relapse.Pattern, p *relapse.Pattern, nullable []bool) (*relapse.Pattern, []bool) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return relapse.NewNot(relapse.NewZAny()), nullable[1:]
	case *relapse.ZAny:
		return relapse.NewZAny(), nullable[1:]
	case *relapse.TreeNode:
		if nullable[0] {
			return relapse.NewEmpty(), nullable[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), nullable[1:]
	case *relapse.LeafNode:
		if nullable[0] {
			return relapse.NewEmpty(), nullable[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), nullable[1:]
	case *relapse.Concat:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		leftConcat := relapse.NewConcat(l, v.GetRightPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return leftConcat, leftRest
		}
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(leftConcat, r), rightRest
	case *relapse.Or:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(l, r), rightRest
	case *relapse.And:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewAnd(l, r), rightRest
	case *relapse.Interleave:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(relapse.NewInterleave(l, v.GetRightPattern()), relapse.NewInterleave(r, v.GetLeftPattern())), rightRest
	case *relapse.ZeroOrMore:
		c, rest := derivReturn(refs, v.GetPattern(), nullable)
		return relapse.NewConcat(c, p), rest
	case *relapse.Reference:
		return derivReturn(refs, refs[v.GetName()], nullable)
	case *relapse.Not:
		c, rest := derivReturn(refs, v.GetPattern(), nullable)
		return relapse.NewNot(c), rest
	case *relapse.Contains:
		return derivReturn(refs, relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())), nullable)
	case *relapse.Optional:
		return derivReturn(refs, relapse.NewOr(v.GetPattern(), relapse.NewEmpty()), nullable)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

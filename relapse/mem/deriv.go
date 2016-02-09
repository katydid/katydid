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
	"github.com/katydid/katydid/funcs"
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
	finals := deriv(refs, []*relapse.Pattern{refs["main"]}, parser)
	return interp.Nullable(refs, finals[0])
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

func deriv(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern, tree serialize.Parser) []*relapse.Pattern {
	var resPatterns []*relapse.Pattern = patterns
	for {
		if !escapable(resPatterns) {
			return resPatterns
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callables := derivCalls(refs, resPatterns)
		calls := combos(callables)
		childPatterns := evalCalls(calls, tree)
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			zchild, zi := zip(childPatterns)
			zchild = deriv(refs, zchild, tree)
			childPatterns = unzip(zchild, zi)
			tree.Up()
		}
		resPatterns = derivReturns(refs, resPatterns, childPatterns)
		resPatterns = simps(refs, resPatterns)
	}
	return resPatterns
}

func evalCalls(calls []*call, label serialize.Decoder) []*relapse.Pattern {
	for _, call := range calls {
		f, err := compose.NewBoolFunc(call.cond)
		if err != nil {
			panic(err)
		}
		res, err := f.Eval(label)
		if err != nil {
			panic(err)
		}
		if res {
			return call.then
		}
	}
	panic("wtf")
}

func simps(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []*relapse.Pattern {
	for i := range patterns {
		patterns[i] = interp.Simplify(refs, patterns[i])
	}
	return patterns
}

func zip(patterns []*relapse.Pattern) ([]*relapse.Pattern, []int) {
	zipped := relapse.Set(patterns)
	relapse.Sort(zipped)
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		indexes[i] = relapse.Index(zipped, pattern)
	}
	return zipped, indexes
}

func unzip(patterns []*relapse.Pattern, indexes []int) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(indexes))
	for i, index := range indexes {
		res[i] = patterns[index]
	}
	return res
}

func derivCalls(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []*callable {
	res := []*callable{}
	for _, pattern := range patterns {
		cs := derivCall(refs, pattern)
		res = append(res, cs...)
	}
	return res
}

type call struct {
	cond funcs.Bool
	then []*relapse.Pattern
}

func combos(callables []*callable) []*call {
	cs := []*call{}
	for _, callable := range callables {
		c := combo(callable)
		cs = combine(cs, c)
	}
	return cs
}

func combine(left, right []*call) []*call {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}
	cs := make([]*call, 0, len(left)*len(right))
	for _, l := range left {
		for _, r := range right {
			cs = append(cs, &call{
				funcs.Simplify(funcs.And(l.cond, r.cond)),
				append([]*relapse.Pattern{}, append(l.then, r.then...)...),
			})
		}
	}
	return cs
}

func combo(c *callable) []*call {
	if c.els == nil {
		return []*call{&call{funcs.BoolConst(true), []*relapse.Pattern{c.then}}}
	}
	return []*call{
		&call{c.cond, []*relapse.Pattern{c.then}},
		&call{funcs.Not(c.cond), []*relapse.Pattern{c.els}},
	}
}

func evals(calls []*callable, label serialize.Decoder) []*relapse.Pattern {
	patterns := make([]*relapse.Pattern, len(calls))
	for i, call := range calls {
		patterns[i] = call.eval(label)
	}
	return patterns
}

type callable struct {
	cond funcs.Bool
	then *relapse.Pattern
	els  *relapse.Pattern
}

func (this *callable) eval(label serialize.Decoder) *relapse.Pattern {
	if this.els == nil {
		return this.then
	}
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		panic(err)
	}
	cond, err := f.Eval(label)
	if err != nil {
		panic(err)
	}
	if cond {
		return this.then
	}
	return this.els
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

func derivReturns(refs map[string]*relapse.Pattern, originals []*relapse.Pattern, evaluated []*relapse.Pattern) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(originals))
	rest := evaluated
	for i, original := range originals {
		res[i], rest = derivReturn(refs, original, rest)
	}
	return res
}

func derivReturn(refs map[string]*relapse.Pattern, p *relapse.Pattern, patterns []*relapse.Pattern) (*relapse.Pattern, []*relapse.Pattern) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return patterns[0], patterns[1:]
	case *relapse.ZAny:
		return patterns[0], patterns[1:]
	case *relapse.TreeNode:
		if interp.Nullable(refs, patterns[0]) {
			return relapse.NewEmpty(), patterns[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), patterns[1:]
	case *relapse.LeafNode:
		if interp.Nullable(refs, patterns[0]) {
			return relapse.NewEmpty(), patterns[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), patterns[1:]
	case *relapse.Concat:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		leftConcat := relapse.NewConcat(l, v.GetRightPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return leftConcat, leftRest
		}
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(leftConcat, r), rightRest
	case *relapse.Or:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(l, r), rightRest
	case *relapse.And:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewAnd(l, r), rightRest
	case *relapse.Interleave:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(relapse.NewInterleave(l, v.GetRightPattern()), relapse.NewInterleave(r, v.GetLeftPattern())), rightRest
	case *relapse.ZeroOrMore:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return relapse.NewConcat(c, p), rest
	case *relapse.Reference:
		return derivReturn(refs, refs[v.GetName()], patterns)
	case *relapse.Not:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return relapse.NewNot(c), rest
	case *relapse.Contains:
		return derivReturn(refs, relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())), patterns)
	case *relapse.Optional:
		return derivReturn(refs, relapse.NewOr(v.GetPattern(), relapse.NewEmpty()), patterns)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

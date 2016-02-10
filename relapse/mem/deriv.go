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
		callTree := newCallTree(callables)
		childPatterns := callTree.eval(tree)
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			zchild, zi, zignore := zip(childPatterns)
			zchild = deriv(refs, zchild, tree)
			childPatterns = unzip(zchild, zi, zignore)
			tree.Up()
		}
		nullable := nullables(refs, childPatterns)
		resPatterns = derivReturns(refs, resPatterns, nullable)
		resPatterns = simps(refs, resPatterns)
	}
	return resPatterns
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

func zip(patterns []*relapse.Pattern) ([]*relapse.Pattern, []int, []*relapse.Pattern) {
	zipped := relapse.Set(patterns)
	relapse.Sort(zipped)
	zignore := []*relapse.Pattern{
		relapse.NewZAny(),
		relapse.NewNot(relapse.NewZAny()),
	}
	if index := relapse.Index(zipped, relapse.NewZAny()); index != -1 {
		zipped = relapse.Remove(zipped, index)
	}
	if index := relapse.Index(zipped, relapse.NewNot(relapse.NewZAny())); index != -1 {
		zipped = relapse.Remove(zipped, index)
	}
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		index := relapse.Index(zipped, pattern)
		if index == -1 {
			index = relapse.Index(zignore, pattern)
			index *= -1
			index -= 1
		}
		indexes[i] = index
	}
	return zipped, indexes, zignore
}

func unzip(patterns []*relapse.Pattern, indexes []int, ignored []*relapse.Pattern) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(indexes))
	for i, index := range indexes {
		if index >= 0 {
			res[i] = patterns[index]
		} else {
			res[i] = ignored[(index+1)*-1]
		}
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

type callNode struct {
	cond funcs.Bool
	then *callNode
	els  *callNode
	term []*relapse.Pattern
}

func (this *callNode) eval(label serialize.Decoder) []*relapse.Pattern {
	if this.term != nil {
		return this.term
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
		return this.then.eval(label)
	}
	return this.els.eval(label)
}

func newCallTree(callables []*callable) *callNode {
	top := newCallNode(callables[0])
	for _, callable := range callables[1:] {
		callnode := newCallNode(callable)
		top = appendCallNode(funcs.BoolConst(true), top, callnode)
	}
	return top
}

func newCallNode(call *callable) *callNode {
	c := &callNode{}
	if call.els == nil {
		c.term = []*relapse.Pattern{call.then}
	} else {
		c.cond = call.cond
		c.then = &callNode{term: []*relapse.Pattern{call.then}}
		c.els = &callNode{term: []*relapse.Pattern{call.els}}
	}
	return c
}

func appendCallNode(cond funcs.Bool, top, bot *callNode) *callNode {
	if top.term != nil {
		return prependTerm(cond, top.term, bot)
	}
	thencond := funcs.Simplify(funcs.And(cond, top.cond))
	elscond := funcs.Simplify(funcs.And(cond, funcs.Not(top.cond)))
	then := appendCallNode(thencond, top.then, bot)
	els := appendCallNode(elscond, top.els, bot)
	return &callNode{
		cond: top.cond,
		then: then,
		els:  els,
	}
}

func prependTerm(cond funcs.Bool, patterns []*relapse.Pattern, bot *callNode) *callNode {
	if bot.term != nil {
		return &callNode{term: append([]*relapse.Pattern{}, append(patterns, bot.term...)...)}
	}
	thencond := funcs.Simplify(funcs.And(cond, bot.cond))
	if funcs.Sprint(thencond) == "false" {
		return prependTerm(cond, patterns, bot.els)
	}
	elscond := funcs.Simplify(funcs.And(cond, funcs.Not(bot.cond)))
	if funcs.Sprint(elscond) == "false" {
		return prependTerm(cond, patterns, bot.then)
	}
	then := prependTerm(thencond, patterns, bot.then)
	els := prependTerm(elscond, patterns, bot.els)
	return &callNode{
		cond: bot.cond,
		then: then,
		els:  els,
	}
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

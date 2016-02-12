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
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/serialize"
)

type composedBool interface {
	Eval(serialize.Decoder) (bool, error)
}

type memCallNode struct {
	cond       composedBool
	then       *memCallNode
	els        *memCallNode
	child      int
	stackIndex int
}

func newMemCallTree(parent int, stackElms *PairIndexedSet, patterns *PatternsIndexedSet, zis *IntsIndexedSet, node *callNode) *memCallNode {
	if node.term != nil {
		ps := node.term
		zps, zi := zip(ps)
		stackIndex := stackElms.add(stackElm{state: parent, zindex: zis.add(zi)})
		return &memCallNode{child: patterns.add(zps), stackIndex: stackIndex}
	}
	then := newMemCallTree(parent, stackElms, patterns, zis, node.then)
	els := newMemCallTree(parent, stackElms, patterns, zis, node.els)
	f, err := compose.NewBoolFunc(node.cond)
	if err != nil {
		panic(err)
	}
	return &memCallNode{
		cond: f,
		then: then,
		els:  els,
	}
}

func (this *memCallNode) eval(label serialize.Decoder) (int, int) {
	if this.cond == nil {
		return this.child, this.stackIndex
	}
	cond, err := this.cond.Eval(label)
	if err != nil {
		panic(err)
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
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

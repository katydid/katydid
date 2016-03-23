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
	"github.com/katydid/katydid/expr/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/serialize"
	"reflect"
	"strings"
)

type CallNode struct {
	cond       compose.Bool
	then       *CallNode
	els        *CallNode
	child      int
	stackIndex int
}

func Implements(c *CallNode, typ reflect.Type) []interface{} {
	var is []interface{}
	fs := getAllConditions(c)
	for _, f := range fs {
		is = append(is, compose.FuncImplements(f, typ)...)
	}
	return is
}

func newMemCallTree(parent int, stackElms *PairIndexedSet, patterns *PatternsIndexedSet, zis *IntsIndexedSet, node *callNode) *CallNode {
	if node.term != nil {
		ps := node.term
		zps, zi := zip(ps)
		stackIndex := stackElms.add(stackElm{State: parent, Zindex: zis.add(zi)})
		return &CallNode{child: patterns.add(zps), stackIndex: stackIndex}
	}
	then := newMemCallTree(parent, stackElms, patterns, zis, node.then)
	els := newMemCallTree(parent, stackElms, patterns, zis, node.els)
	f, err := compose.NewBoolFunc(node.cond)
	if err != nil {
		panic(err)
	}
	return &CallNode{
		cond: f,
		then: then,
		els:  els,
	}
}

func getAllConditions(c *CallNode) []compose.Bool {
	if c.cond == nil {
		return []compose.Bool{}
	}
	cs := []compose.Bool{c.cond}
	if c.then != nil {
		cs = append(cs, getAllConditions(c.then)...)
	}
	if c.els != nil {
		cs = append(cs, getAllConditions(c.then)...)
	}
	return cs
}

func (this *CallNode) Eval(label serialize.Decoder) (int, int) {
	if this.cond == nil {
		return this.child, this.stackIndex
	}
	cond, err := this.cond.Eval(label)
	if err != nil {
		panic(err)
	}
	if cond {
		return this.then.Eval(label)
	}
	return this.els.Eval(label)
}

type callNode struct {
	cond funcs.Bool
	then *callNode
	els  *callNode
	term []*relapse.Pattern
}

func addtab(s string) string {
	ss := strings.Split(s, "\n")
	for i := range ss {
		ss[i] = "\t" + ss[i]
	}
	return strings.Join(ss, "\n")
}

func (this *callNode) String() string {
	if this.term != nil {
		ss := make([]string, len(this.term))
		for i := range this.term {
			ss[i] = this.term[i].String()
		}
		return strings.Join(ss, ", ")
	}
	sthen := addtab(this.then.String())
	sels := addtab(this.els.String())
	sfunc := funcs.Sprint(this.cond)
	return "{\n" + sfunc + "\nThen:\n" + sthen + "\nElse:\n" + sels + "}"
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
	if len(callables) == 0 {
		return &callNode{
			term: []*relapse.Pattern{},
		}
	}
	top := newCallNode(callables[0])
	for _, callable := range callables[1:] {
		if callable.cond == nil {
			top = appendCallTerm(top, callable.then)
		} else {
			top = appendCallNode(top, callable.cond, callable.then, callable.els)
		}
	}
	return top
}

func newCallNode(call *callable) *callNode {
	c := &callNode{}
	if call.els == nil {
		c.term = []*relapse.Pattern{call.then}
	} else if call.then.Equal(call.els) {
		c.term = []*relapse.Pattern{call.then}
	} else {
		c.cond = call.cond
		c.then = &callNode{term: []*relapse.Pattern{call.then}}
		c.els = &callNode{term: []*relapse.Pattern{call.els}}
	}
	return c
}

func appendCallTerm(top *callNode, term *relapse.Pattern) *callNode {
	if top.term != nil {
		return &callNode{term: append([]*relapse.Pattern{}, append(top.term, term)...)}
	}
	then := appendCallTerm(top.then, term)
	els := appendCallTerm(top.els, term)
	return &callNode{
		cond: top.cond,
		then: then,
		els:  els,
	}
}

func appendCallNode(top *callNode, cond funcs.Bool, then, els *relapse.Pattern) *callNode {
	if top.term != nil {
		thens := append([]*relapse.Pattern{}, append(top.term, then)...)
		elss := append([]*relapse.Pattern{}, append(top.term, els)...)
		return &callNode{
			cond: cond,
			then: &callNode{term: thens},
			els:  &callNode{term: elss},
		}
	}
	if funcs.Equal(top.cond, cond) {
		thennode := appendCallTerm(top.then, then)
		elsnode := appendCallTerm(top.els, els)
		return &callNode{
			cond: top.cond,
			then: thennode,
			els:  elsnode,
		}
	}
	if funcs.Sprint(funcs.Simplify(funcs.And(top.cond, cond))) == "false" {
		thennode := appendCallTerm(top.then, els)
		elsnode := appendCallNode(top.els, cond, then, els)
		return &callNode{
			cond: top.cond,
			then: thennode,
			els:  elsnode,
		}
	}
	if funcs.Sprint(funcs.Simplify(funcs.And(top.cond, funcs.Not(cond)))) == "false" {
		thennode := appendCallNode(top.then, cond, then, els)
		elsnode := appendCallTerm(top.els, then)
		return &callNode{
			cond: top.cond,
			then: thennode,
			els:  elsnode,
		}
	}
	thennode := appendCallNode(top.then, cond, then, els)
	elsnode := appendCallNode(top.els, cond, then, els)
	return &callNode{
		cond: top.cond,
		then: thennode,
		els:  elsnode,
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

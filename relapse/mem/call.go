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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"strings"
)

//CallNode represents a node in the call tree.
//The call tree is basically a nested if then else tree, which results in child states and stack elements.
type CallNode struct {
	cond       compose.Bool
	then       *CallNode
	els        *CallNode
	child      int
	stackIndex int
}

func newMemCallTree(parent int, stackElms *pairSet, patterns *patternsSet, zis *intsSet, node *callNode) (*CallNode, error) {
	if node.term != nil {
		ps := node.term
		zps, zi := zip(ps)
		stackIndex := stackElms.add(stackElm{State: parent, Zindex: zis.add(zi)})
		return &CallNode{child: patterns.add(zps), stackIndex: stackIndex}, nil
	}
	then, err := newMemCallTree(parent, stackElms, patterns, zis, node.then)
	if err != nil {
		return nil, err
	}
	els, err := newMemCallTree(parent, stackElms, patterns, zis, node.els)
	if err != nil {
		return nil, err
	}
	f, err := compose.NewBoolFunc(node.cond)
	if err != nil {
		return nil, err
	}
	return &CallNode{
		cond: f,
		then: then,
		els:  els,
	}, nil
}

//Eval evaluates the call tree returning the child state and stack element given the label value of the parser element.
func (this *CallNode) Eval(label parser.Value) (int, int, error) {
	if this.cond == nil {
		return this.child, this.stackIndex, nil
	}
	cond, err := this.cond.Eval(label)
	if err != nil {
		return 0, 0, err
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
	term []*ast.Pattern
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

func (this *callNode) eval(label parser.Value) ([]*ast.Pattern, error) {
	if this.term != nil {
		return this.term, nil
	}
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
}

func newCallTree(callables []*callable) *callNode {
	if len(callables) == 0 {
		return &callNode{
			term: []*ast.Pattern{},
		}
	}
	top := newCallNode(callables[0])
	for _, callable := range callables[1:] {
		if callable.cond == nil {
			top.appendTerm(callable.then)
		} else {
			top.appendNode(callable.cond, callable.then, callable.els)
		}
	}
	return top
}

func newCallNode(call *callable) *callNode {
	c := &callNode{}
	if call.els == nil || call.then.Equal(call.els) {
		c.term = []*ast.Pattern{call.then}
	} else {
		c.cond = call.cond
		c.then = &callNode{term: []*ast.Pattern{call.then}}
		c.els = &callNode{term: []*ast.Pattern{call.els}}
	}
	return c
}

func (top *callNode) appendTerm(term *ast.Pattern) {
	if top.term != nil {
		top.term = append(top.term, term)
		return
	}
	top.then.appendTerm(term)
	top.els.appendTerm(term)
	return
}

func (top *callNode) appendNode(cond funcs.Bool, then, els *ast.Pattern) {
	if top.term != nil {
		top.cond = cond
		thenterms := make([]*ast.Pattern, len(top.term)+1)
		copy(thenterms, top.term)
		thenterms[len(thenterms)-1] = then
		top.then = &callNode{term: thenterms}
		top.els = &callNode{term: append(top.term, els)}
		top.term = nil
		return
	}
	if funcs.Equal(top.cond, cond) {
		top.then.appendTerm(then)
		top.els.appendTerm(els)
		return
	}
	if funcs.IsFalse(funcs.Simplify(funcs.And(top.cond, cond))) {
		top.then.appendTerm(els)
		top.els.appendNode(cond, then, els)
		return
	}
	if funcs.IsFalse(funcs.Simplify(funcs.And(top.cond, funcs.Not(cond)))) {
		top.then.appendNode(cond, then, els)
		top.els.appendTerm(then)
		return
	}
	top.then.appendNode(cond, then, els)
	top.els.appendNode(cond, then, els)
	return
}

type callable struct {
	cond funcs.Bool
	then *ast.Pattern
	els  *ast.Pattern
}

func (this *callable) eval(label parser.Value) (*ast.Pattern, error) {
	if this.els == nil {
		return this.then, nil
	}
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then, nil
	}
	return this.els, nil
}

//  Copyright 2013 Walter Schulze
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

package compiler

import (
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/asm/ifexpr"
	"github.com/awalterschulze/katydid/asm/protomap"
	"github.com/awalterschulze/katydid/asm/table"
)

type Link interface {
	ProtoToStart(protoState int) (startState int, exists bool)
	IfEval(protoState int, buf []byte) (state int, exists bool)
	GetIfs() []ifexpr.StateExpr
}

type link struct {
	protoToStart []int
	protoToIf    []ifexpr.StateExpr
}

func NewLink(rules *ast.Rules, pmap protomap.ProtoMap, tab table.Table) (Link, error) {
	l := pmap.LenStates()
	link := &link{
		protoToStart: make([]int, l),
		protoToIf:    make([]ifexpr.StateExpr, l),
	}
	for _, init := range rules.GetInit() {
		pstate, err := pmap.State(init.GetPackage(), init.GetMessage(), "")
		if err != nil {
			return nil, err
		}
		tstate, err := tab.NameToState(init.GetState())
		if err != nil {
			return nil, err
		}
		link.protoToStart[pstate] = tstate
	}
	for _, ifExpr := range rules.GetIfExpr() {
		var1, err := ifexpr.GetVariable(ifExpr)
		if err != nil {
			return nil, err
		}
		pstate, err := pmap.State(var1.GetPackage(), var1.GetMessage(), var1.GetField())
		if err != nil {
			return nil, err
		}
		stateExpr, err := ifexpr.Compile(ifExpr, tab)
		if err != nil {
			return nil, err
		}
		link.protoToIf[pstate] = stateExpr
	}
	return link, nil
}

func (this *link) ProtoToStart(protoState int) (startState int, exists bool) {
	if protoState < 0 || protoState >= len(this.protoToStart) {
		return 0, false
	}
	s := this.protoToStart[protoState]
	return s, (s != 0)
}

func (this *link) IfEval(protoState int, buf []byte) (state int, exists bool) {
	if protoState < 0 || protoState >= len(this.protoToIf) {
		return 0, false
	}
	ifExpr := this.protoToIf[protoState]
	if ifExpr == nil {
		return 0, false
	}
	return ifExpr.Eval(buf), true
}

func (this *link) GetIfs() []ifexpr.StateExpr {
	return this.protoToIf
}

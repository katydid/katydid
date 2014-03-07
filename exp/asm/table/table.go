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

package table

import (
	"fmt"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"strings"
)

type errUnknownStateName struct {
	name string
}

func (this *errUnknownStateName) Error() string {
	return "unknown state " + this.name
}

type errUnknownState struct {
	s int
}

func (this *errUnknownState) Error() string {
	return fmt.Sprintf("unknown state %d", this.s)
}

type errUnknownInput struct {
	src int
	i   int
}

func (this *errUnknownInput) Error() string {
	return fmt.Sprintf("unknown input %d from %d and no default available", this.i, this.src)
}

type table struct {
	nameToState map[string]int
	stateToName []string
	trans       []map[int]int
}

func (this *table) allocState(name string) int {
	if s, ok := this.nameToState[name]; ok {
		return s
	}
	s := len(this.trans)
	this.nameToState[name] = s
	this.trans = append(this.trans, make(map[int]int))
	this.stateToName = append(this.stateToName, name)
	return s
}

func (this *table) allocTrans(src string, input string, dst string) {
	s := this.allocState(src)
	var i int
	if input == "_" {
		i = 0 //default input
	} else {
		i = this.allocState(input)
	}
	d := this.allocState(dst)
	this.trans[s][i] = d
}

func (this *table) allocIfStates(expr *ast.StateExpr) {
	if expr.State != nil {
		this.allocState(expr.GetState())
		return
	}
	this.allocIfStates(expr.GetIfExpr().GetThen())
	this.allocIfStates(expr.GetIfExpr().GetElse())
}

func New(transitions []*ast.Transition, ifs []*ast.IfExpr) Table {
	tab := &table{make(map[string]int), make([]string, 0), make([]map[int]int, 0)}
	tab.allocState("_") // making the default input "_" = 0
	for _, trans := range transitions {
		tab.allocTrans(trans.GetSrc(), trans.GetInput(), trans.GetDst())
	}
	for _, ifExpr := range ifs {
		tab.allocIfStates(ifExpr.GetThen())
		tab.allocIfStates(ifExpr.GetElse())
	}
	return tab
}

type Table interface {
	NameToState(name string) (int, error)
	Trans(src int, input int) (int, error)
	Dot() string
}

func (this *table) NameToState(name string) (int, error) {
	s, ok := this.nameToState[name]
	if !ok {
		return 0, &errUnknownStateName{name}
	}
	return s, nil
}

func (this *table) Trans(src int, input int) (int, error) {
	if src >= len(this.trans) {
		return 0, &errUnknownState{src}
	}
	d, ok := this.trans[src][input]
	if !ok {
		def, ok := this.trans[src][0]
		if !ok {
			return 0, &errUnknownInput{src, input}
		}
		return def, nil
	}
	return d, nil
}

func (this *table) Dot() string {
	lines := make([]string, 0, len(this.stateToName)*4)
	for s, name := range this.stateToName {
		lines = append(lines, fmt.Sprintf("\tnode%d [label=\"%v(%d)\"]\n", s, name, s))
	}
	for s, trans := range this.trans {
		for i, d := range trans {
			lines = append(lines, fmt.Sprintf("\tnode%d -> node%d [label=\"%s(%d)\"]\n", s, d, this.stateToName[i], i))
		}
	}
	return "digraph map {\n" + strings.Join(lines, "") + "}"
}

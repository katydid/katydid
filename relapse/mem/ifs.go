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

package mem

import (
	"fmt"
	"strings"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/intern"
)

type ifExprs struct {
	*ifNode
}

func newIfExprs(ifs []*intern.IfExpr) *ifExprs {
	return &ifExprs{&ifNode{prev: funcs.BoolConst(true), ifs: ifs}}
}

type ifNode struct {
	prev funcs.Bool
	f    funcs.Bool
	ifs  []*intern.IfExpr

	cond compose.Bool
	thn  *ifNode
	els  *ifNode

	ps  []*intern.Pattern
	res int
}

func (this *ifNode) String() string {
	if this == nil {
		return "...\n"
	}
	if len(this.ifs) == 0 {
		ss := make([]string, len(this.ps))
		for i := range this.ps {
			ss[i] = this.ps[i].String()
		}
		return fmt.Sprintf("%s\n", strings.Join(ss, ","))
	}
	if this.f == nil {
		return "...\n"
	}
	return "if (" + funcs.Sprint(this.f) + ") then {\n" + this.thn.String() + "} else {\n" + this.els.String() + "}"
}

func (node *ifNode) eval(set *intern.SetOfPatterns, label parser.Value) (int, error) {
	if len(node.ifs) == 0 {
		if node.ps != nil {
			node.res = set.Add(node.ps)
			node.ps = nil
		}
		return node.res, nil
	}
	if node.f == nil {
		node.f = node.ifs[0].Cond
		node.f = funcs.Simplify(node.f)
		if funcs.Equal(node.prev, node.f) {
			node.f = funcs.BoolConst(true)
		}
		if funcs.IsFalse(funcs.Simplify(funcs.And(node.prev, node.f))) {
			node.f = funcs.BoolConst(false)
		}
		if funcs.IsFalse(funcs.Simplify(funcs.And(node.prev, funcs.Not(node.f)))) {
			node.f = funcs.BoolConst(true)
		}
		if funcs.IsTrue(node.f) {
			node.ps = append(node.ps, node.ifs[0].Thn)
			node.ifs = node.ifs[1:]
			node.f = nil
			return node.eval(set, label)
		}
		if funcs.IsFalse(node.f) {
			node.ps = append(node.ps, node.ifs[0].Els)
			node.ifs = node.ifs[1:]
			node.f = nil
			return node.eval(set, label)
		}
		c, err := compose.NewBoolFunc(node.f)
		if err != nil {
			return 0, err
		}
		node.cond = c
	}
	cond, err := node.cond.Eval(label)
	if err != nil {
		return 0, err
	}
	if cond {
		if node.thn == nil {
			node.thn = &ifNode{}
			node.thn.ps = make([]*intern.Pattern, 0, len(node.ps)+1)
			node.thn.ps = append(node.thn.ps, node.ps...)
			node.thn.ps = append(node.thn.ps, node.ifs[0].Thn)
			node.thn.prev = funcs.Simplify(funcs.And(node.prev, node.f))
			node.thn.ifs = node.ifs[1:]
		}
		return node.thn.eval(set, label)
	}
	if node.els == nil {
		node.els = &ifNode{}
		node.els.ps = make([]*intern.Pattern, 0, len(node.ps)+1)
		node.els.ps = append(node.els.ps, node.ps...)
		node.els.ps = append(node.els.ps, node.ifs[0].Els)
		node.els.prev = funcs.Simplify(funcs.And(node.prev, funcs.Not(node.f)))
		node.els.ifs = node.ifs[1:]
	}
	return node.els.eval(set, label)
}

func (this *Mem) calcNode(node *ifNode, label parser.Value) (int, int, error) {
	state, err := node.eval(this.states, label)
	if err != nil {
		return 0, 0, err
	}
	zippedPatternIndex, stackIndex := this.zipStackAndPatterns(state)
	return zippedPatternIndex, stackIndex, nil
}

func (this *Mem) zipStackAndPatterns(state int) (int, int) {
	p := this.states.Get(state)
	zipperIndex := p.IndexOfZippedIndexes
	zippedPatternIndex := p.IndexOfZippedPatterns
	return zippedPatternIndex, zipperIndex
}

func (this *Mem) eval(ifs *ifExprs, label parser.Value) (int, int, error) {
	return this.calcNode(ifs.ifNode, label)
}

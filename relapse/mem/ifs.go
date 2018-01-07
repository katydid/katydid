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
	"github.com/katydid/katydid/relapse/sets"
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
	zip *intern.ZippedPatterns
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

func (node *ifNode) eval(label parser.Value) (*intern.ZippedPatterns, error) {
	if len(node.ifs) == 0 {
		if node.zip == nil {
			node.zip = intern.Zip(node.ps)
		}
		return node.zip, nil
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
			return node.eval(label)
		}
		if funcs.IsFalse(node.f) {
			node.ps = append(node.ps, node.ifs[0].Els)
			node.ifs = node.ifs[1:]
			node.f = nil
			return node.eval(label)
		}
		c, err := compose.NewBoolFunc(node.f)
		if err != nil {
			return nil, err
		}
		node.cond = c
	}
	cond, err := node.cond.Eval(label)
	if err != nil {
		return nil, err
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
		return node.thn.eval(label)
	}
	if node.els == nil {
		node.els = &ifNode{}
		node.els.ps = make([]*intern.Pattern, 0, len(node.ps)+1)
		node.els.ps = append(node.els.ps, node.ps...)
		node.els.ps = append(node.els.ps, node.ifs[0].Els)
		node.els.prev = funcs.Simplify(funcs.And(node.prev, funcs.Not(node.f)))
		node.els.ifs = node.ifs[1:]
	}
	return node.els.eval(label)
}

func (this *Mem) calcNode(node *ifNode, parentPatterns int, label parser.Value) (int, int, error) {
	zipped, err := node.eval(label)
	if err != nil {
		return 0, 0, err
	}
	zippedPatternIndex, stackIndex := this.zipStackAndPatterns(parentPatterns, zipped)
	return zippedPatternIndex, stackIndex, nil
}

func (this *Mem) zipStackAndPatterns(parentPatterns int, zipped *intern.ZippedPatterns) (int, int) {
	zipperIndex := this.zis.Add(zipped.Indexes)
	stackElement := sets.Pair{
		First:  parentPatterns,
		Second: zipperIndex,
	}
	stackIndex := this.stackElms.Add(stackElement)
	zippedPatternIndex := this.states.Add(zipped.Patterns)
	return zippedPatternIndex, stackIndex
}

func (this *Mem) eval(parentPatterns int, ifs *ifExprs, label parser.Value) (int, int, error) {
	return this.calcNode(ifs.ifNode, parentPatterns, label)
}

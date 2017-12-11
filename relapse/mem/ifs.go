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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/intern"
	"github.com/katydid/katydid/relapse/sets"
)

type ifExprs struct {
	parentPatterns int
	ifs            []*ifExpr
	node           *ifNode
}

func newIfExprs(parentPatterns int, ifs []*ifExpr) *ifExprs {
	return &ifExprs{parentPatterns, ifs, nil}
}

type ifNode struct {
	prev               funcs.Bool
	f                  funcs.Bool
	ifs                []*ifExpr
	cond               compose.Bool
	thn                *ifNode
	els                *ifNode
	ps                 []*intern.Pattern
	ret                bool
	zippedPatternIndex int
	stackIndex         int
}

func (this *ifNode) String() string {
	if this == nil {
		return "...\n"
	}
	if this.ret {
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

func (this *Mem) calcNode(node *ifNode, parentPatterns int, label parser.Value) (int, int, error) {
	if len(node.ifs) == 0 {
		if !node.ret {
			node.zippedPatternIndex, node.stackIndex = this.zipStackAndPatterns(parentPatterns, node.ps)
			node.ret = true
		}
		return node.zippedPatternIndex, node.stackIndex, nil
	}
	if node.f == nil {
		node.f = node.ifs[0].cond
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
			node.ps = append(node.ps, node.ifs[0].thn)
			node.ifs = node.ifs[1:]
			node.f = nil
			return this.calcNode(node, parentPatterns, label)
		}
		if funcs.IsFalse(node.f) {
			node.ps = append(node.ps, node.ifs[0].els)
			node.ifs = node.ifs[1:]
			node.f = nil
			return this.calcNode(node, parentPatterns, label)
		}
		c, err := compose.NewBoolFunc(node.f)
		if err != nil {
			return 0, 0, err
		}
		node.cond = c
	}
	cond, err := node.cond.Eval(label)
	if err != nil {
		return 0, 0, err
	}
	if cond {
		if node.thn == nil {
			node.thn = &ifNode{}
			node.thn.ps = make([]*intern.Pattern, 0, len(node.ps)+1)
			node.thn.ps = append(node.thn.ps, node.ps...)
			node.thn.ps = append(node.thn.ps, node.ifs[0].thn)
			node.thn.prev = funcs.Simplify(funcs.And(node.prev, node.f))
			node.thn.ifs = node.ifs[1:]
		}
		return this.calcNode(node.thn, parentPatterns, label)
	}
	if node.els == nil {
		node.els = &ifNode{}
		node.els.ps = make([]*intern.Pattern, 0, len(node.ps)+1)
		node.els.ps = append(node.els.ps, node.ps...)
		node.els.ps = append(node.els.ps, node.ifs[0].els)
		node.els.prev = funcs.Simplify(funcs.And(node.prev, funcs.Not(node.f)))
		node.els.ifs = node.ifs[1:]
	}
	return this.calcNode(node.els, parentPatterns, label)
}

func newasts(ps []*intern.Pattern) []*ast.Pattern {
	pps := make([]*ast.Pattern, len(ps))
	for i := range pps {
		pps[i] = ps[i].NewAst()
	}
	return pps
}

func (this *Mem) zipStackAndPatterns(parentPatterns int, ps []*intern.Pattern) (int, int) {
	fmt.Printf("ps = %v\n", ps)
	zippedPatterns, zipper := intern.Zip(ps)
	fmt.Printf("zipped %v, %v\n", zippedPatterns, zipper)
	astPatterns := newasts(ps)
	astzippedPatterns, astzipper := sets.Zip(astPatterns)
	fmt.Printf("astzipped %v, %v\n", astzippedPatterns, astzipper)
	zipperIndex := this.zis.Add(zipper)
	stackElement := sets.Pair{
		First:  parentPatterns,
		Second: zipperIndex,
	}
	fmt.Printf("stackElement: %v\n", stackElement)
	stackIndex := this.stackElms.Add(stackElement)
	zippedPatternIndex := this.patterns.Add(zippedPatterns)
	return zippedPatternIndex, stackIndex
}

func (this *Mem) eval(ifs *ifExprs, label parser.Value) (int, int, error) {
	if ifs.node == nil {
		ifs.node = &ifNode{prev: funcs.BoolConst(true), ifs: ifs.ifs}
	}
	return this.calcNode(ifs.node, ifs.parentPatterns, label)
}

type ifExpr struct {
	cond funcs.Bool
	thn  *intern.Pattern
	els  *intern.Pattern
}

func (this *ifExpr) eval(label parser.Value) (*intern.Pattern, error) {
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.thn, nil
	}
	return this.els, nil
}

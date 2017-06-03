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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
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
	cond               compose.Bool
	thn                *ifNode
	els                *ifNode
	ps                 []*ast.Pattern
	ret                bool
	zippedPatternIndex int
	stackIndex         int
}

func (this *Mem) calc(node *ifNode, ifs []*ifExpr, parentPatterns int, label parser.Value) (int, int, error) {
	if len(ifs) == 0 {
		if !node.ret {
			node.zippedPatternIndex, node.stackIndex = this.zipStackAndPatterns(parentPatterns, node.ps)
			node.ret = true
		}
		return node.zippedPatternIndex, node.stackIndex, nil
	}
	if node.cond == nil {
		f, err := compose.NewBoolFunc(ifs[0].cond)
		if err != nil {
			return 0, 0, err
		}
		node.cond = f
	}
	cond, err := node.cond.Eval(label)
	if err != nil {
		return 0, 0, err
	}
	if cond {
		if node.thn == nil {
			node.thn = &ifNode{}
			node.thn.ps = make([]*ast.Pattern, len(node.ps)+1)
			copy(node.thn.ps, node.ps)
			node.thn.ps[len(node.thn.ps)-1] = ifs[0].thn
		}
		return this.calc(node.thn, ifs[1:], parentPatterns, label)
	}
	if node.els == nil {
		node.els = &ifNode{}
		node.els.ps = make([]*ast.Pattern, len(node.ps)+1)
		copy(node.els.ps, node.ps)
		node.els.ps[len(node.els.ps)-1] = ifs[0].els
	}
	return this.calc(node.els, ifs[1:], parentPatterns, label)
}

func (this *Mem) zipStackAndPatterns(parentPatterns int, ps []*ast.Pattern) (int, int) {
	zippedPatterns, zipper := sets.Zip(ps)
	zipperIndex := this.zis.Add(zipper)
	stackElement := sets.Pair{
		First:  parentPatterns,
		Second: zipperIndex,
	}
	stackIndex := this.stackElms.Add(stackElement)
	zippedPatternIndex := this.patterns.Add(zippedPatterns)
	return zippedPatternIndex, stackIndex
}

func (this *Mem) eval(ifs *ifExprs, label parser.Value) (int, int, error) {
	if ifs.node == nil {
		ifs.node = &ifNode{}
	}
	return this.calc(ifs.node, ifs.ifs, ifs.parentPatterns, label)
}

type ifExpr struct {
	cond funcs.Bool
	thn  *ast.Pattern
	els  *ast.Pattern
}

func (this *ifExpr) eval(label parser.Value) (*ast.Pattern, error) {
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

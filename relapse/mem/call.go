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
	"github.com/katydid/katydid/relapse/compose"
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

func newMemCallTree(parent int, stackElms *pairSet, patterns *patternsSet, zis *intsSet, node *ifExprs) (*CallNode, error) {
	if node.ret != nil {
		ps := node.ret
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

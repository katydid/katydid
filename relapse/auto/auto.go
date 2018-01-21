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

//Package auto compiles a parsed relapse grammar into a visual pushdown automaton and executes it.
//Compilation into a VPA may result in an exponential explosion, since fully converting a grammar to VPA is O(2^n^2).
//Rather use the mem package.  It gives comparable speed and has no exponential behaviour.
//This package is just here to provide a benchmark against the mem package.
package auto

import (
	"io"

	"github.com/katydid/katydid/parser"
)

//Auto is the structure that represents the automaton.
type Auto struct {
	calls           []*callNode
	returns         []map[int]int
	escapables      []bool
	start           int
	stateToNullable []int
	accept          []bool
}

//Validate executes an automaton with the given parser and returns whether the parser is valid given the automaton's original grammar.
func (auto *Auto) Validate(p parser.Interface) (bool, error) {
	final, err := deriv(auto, auto.start, p)
	if err != nil {
		return false, err
	}
	return auto.accept[final], nil
}

func deriv(auto *Auto, current int, tree parser.Interface) (int, error) {
	for {
		if !auto.escapables[current] {
			return current, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		callTree := auto.calls[current]
		childState, stackElm, err := callTree.eval(tree)
		if err != nil {
			return 0, err
		}
		if !tree.IsLeaf() {
			tree.Down()
			childState, err = deriv(auto, childState, tree)
			if err != nil {
				return 0, err
			}
			tree.Up()
		}
		nullIndex := auto.stateToNullable[childState]
		current = auto.returns[stackElm][nullIndex]
	}
	return current, nil
}

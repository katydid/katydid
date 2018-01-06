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
	"io"

	"github.com/katydid/katydid/parser"
)

func deriv(mem *Mem, patterns int, tree parser.Interface) (int, error) {
	for {
		if !mem.states.Get(patterns).Escapable {
			return patterns, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		callTree, err := mem.getCallTree(patterns)
		if err != nil {
			return 0, err
		}
		childPatterns, stackElm, err := mem.eval(callTree, tree)
		if err != nil {
			return 0, err
		}
		if !tree.IsLeaf() {
			tree.Down()
			childPatterns, err = deriv(mem, childPatterns, tree)
			if err != nil {
				return 0, err
			}
			tree.Up()
		}
		nullIndex := mem.getNullable(childPatterns)
		patterns, err = mem.getReturn(stackElm, nullIndex)
		if err != nil {
			return 0, err
		}
	}
	return patterns, nil
}

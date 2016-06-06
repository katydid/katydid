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
	"github.com/katydid/katydid/relapse/ast"
)

//Compile memoizes the full state space, all possible things that can be memoized.
func Compile(g *ast.Grammar) (*Mem, error) {
	mem, err := New(g)
	if err != nil {
		return nil, err
	}
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		for state := range mem.patterns {
			if visited[state] {
				continue
			}
			visited[state] = true
			if err := compile(mem, state); err != nil {
				return nil, err
			}
			changed = true
		}
	}
	return mem, nil
}

func compile(mem *Mem, current int) error {
	mem.getNullable(current)
	mem.accept(current)
	mem.escapable(current)

	callTree, err := mem.getCallTree(current)
	if err != nil {
		return err
	}
	leafs := getLeafs(callTree)
	for _, leaf := range leafs {
		childlen := len(mem.patterns[leaf.child])
		combos := uint64(1)
		if childlen > 0 {
			// pow(2, childlen)
			combos = combos << uint(childlen)
		}
		for combo := uint64(0); combo < combos; combo++ {
			nullIndex := mem.nullables.add(bitset{
				val:  combo,
				size: childlen,
			})
			mem.getReturnn(leaf.stackIndex, nullIndex)
		}
	}
	return nil
}

func getLeafs(callTree *CallNode) []*CallNode {
	if callTree.cond == nil {
		return []*CallNode{callTree}
	}
	then := getLeafs(callTree.then)
	els := getLeafs(callTree.els)
	return append(then, els...)
}

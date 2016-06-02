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
func Compile(g *ast.Grammar) *Mem {
	mem := New(g)
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		for state := range mem.patterns {
			if visited[state] {
				continue
			}
			visited[state] = true
			compile(mem, state)
			changed = true
		}
	}
	return mem
}

func compile(mem *Mem, current int) {
	mem.getNullable(current)
	mem.accept(current)
	mem.escapable(current)

	callTree := mem.getCallTree(current)
	leafs := getLeafs(callTree)
	for _, leaf := range leafs {
		childlen := len(mem.patterns[leaf.child])
		nullablecombos := getBitsetCombos(childlen)
		for _, nullablecombo := range nullablecombos {
			nullIndex := mem.nullables.add(nullablecombo)
			mem.getReturnn(leaf.stackIndex, nullIndex)
		}
	}
}

func getLeafs(callTree *CallNode) []*CallNode {
	if callTree.cond == nil {
		return []*CallNode{callTree}
	}
	then := getLeafs(callTree.then)
	els := getLeafs(callTree.els)
	return append(then, els...)
}

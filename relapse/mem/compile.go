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
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
)

func Compile(g *relapse.Grammar) Mem {
	refs := relapse.NewRefsLookup(g)
	mem := newMem(refs)
	changed := true
	for changed {
		changed = false
		for state := range mem.patternsMap {
			if !mem.escapable(state) {
				continue
			}
			if _, ok := mem.callTrees[state]; !ok {
				compile(mem, state)
				if _, ok := mem.callTrees[state]; !ok {
					panic("wtf")
				}
				changed = true
			}
		}
	}
	return mem
}

func prints(prefix string, state int, patterns []*relapse.Pattern) {
	fmt.Printf(prefix+":%d:", state)
	for _, p := range patterns {
		fmt.Printf("%v, ", p.String())
	}
	fmt.Printf("\n")
}

func compile(mem *mem, current int) {
	mem.getNullable(current)
	callTree := mem.getCallTree(current)
	leafs := getLeafs(callTree)
	for _, leaf := range leafs {
		childlen := len(mem.patternsMap[leaf.child])
		nullablecombos := mcombos(childlen)
		for _, nullablecombo := range nullablecombos {
			nullIndex := mem.nullables.add(nullablecombo)
			mem.getReturnn(leaf.stackIndex, nullIndex)
		}
	}
}

func getLeafs(callTree *memCallNode) []*memCallNode {
	if callTree.cond == nil {
		return []*memCallNode{callTree}
	}
	then := getLeafs(callTree.then)
	els := getLeafs(callTree.els)
	return append(then, els...)
}

var m map[int][][]bool = make(map[int][][]bool)

func mcombos(n int) [][]bool {
	_, ok := m[n]
	if !ok {
		m[n] = combos(n)
	}
	return m[n]
}

func combos(n int) [][]bool {
	if n == 0 {
		return [][]bool{}
	}
	if n == 1 {
		return [][]bool{[]bool{false}, []bool{true}}
	}
	cs := mcombos(n - 1)
	res := [][]bool{}
	for _, c := range cs {
		f := append([]bool{}, append(c, false)...)
		t := append([]bool{}, append(c, true)...)
		res = append(append(res, f), t)
	}
	return res
}

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
	"github.com/katydid/katydid/relapse/interp"
)

//Compile memoizes the full state space (all possible things that can be memoized) before even executing on a parser.
func Compile(g *ast.Grammar) *Mem {
	refs := ast.NewRefLookup(g)
	for name, p := range refs {
		refs[name] = interp.Simplify(refs, p)
	}
	mem := newMem(refs)
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		for state := range mem.PatternsMap {
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

func prints(prefix string, state int, patterns []*ast.Pattern) {
	fmt.Printf(prefix+":%d:", state)
	for _, p := range patterns {
		fmt.Printf("%v, ", p.String())
	}
	fmt.Printf("\n")
}

func compile(mem *Mem, current int) {
	mem.getNullable(current)
	mem.accept(current)
	mem.escapable(current)

	callTree := mem.getCallTree(current)
	leafs := getLeafs(callTree)
	for _, leaf := range leafs {
		childlen := len(mem.PatternsMap[leaf.child])
		nullablecombos := mcombos(childlen)
		for _, nullablecombo := range nullablecombos {
			nullIndex := mem.Nullables.add(nullablecombo)
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
		return [][]bool{[]bool{}}
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

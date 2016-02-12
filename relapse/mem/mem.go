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
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/serialize"
)

type Mem interface {
	Interpret(serialize.Parser) bool
}

//This is a naive implementation and it does not handle left recursion
func (mem *mem) Interpret(parser serialize.Parser) bool {
	final := deriv(mem, mem.start, parser)
	return mem.accept(final)
}

func New(g *relapse.Grammar) Mem {
	refs := relapse.NewRefsLookup(g)
	mem := newMem(refs)
	return mem
}

type mem struct {
	refs        map[string]*relapse.Pattern
	patternsMap PatternsIndexedSet
	callTrees   map[int]*memCallNode
	returns     map[int]map[int]int
	escapables  map[int]bool
	zis         IntsIndexedSet
	start       int

	stackElms       PairIndexedSet
	stateToNullable map[int]int
	nullables       BoolsIndexedSet
}

func newMem(refs map[string]*relapse.Pattern) *mem {
	m := &mem{
		refs:        refs,
		patternsMap: newPatternsIndexedSet(),
		callTrees:   make(map[int]*memCallNode),
		returns:     make(map[int]map[int]int),
		escapables:  make(map[int]bool),
		zis:         newIntsIndexedSet(),

		stackElms:       newPairIndexedSet(),
		stateToNullable: make(map[int]int),
		nullables:       newBoolsIndexedSet(),
	}
	start := m.patternsMap.add([]*relapse.Pattern{refs["main"]})
	m.start = start
	return m
}

func (this *mem) escapable(s int) bool {
	if e, ok := this.escapables[s]; ok {
		return e
	}
	patterns := this.patternsMap[s]
	e := escapable(patterns)
	this.escapables[s] = e
	return e
}

func (this *mem) accept(s int) bool {
	patterns := this.patternsMap[s]
	if len(patterns) != 1 {
		return false
	}
	return interp.Nullable(this.refs, patterns[0])
}

func (this *mem) getCallTree(s int) *memCallNode {
	ct, ok := this.callTrees[s]
	if ok {
		return ct
	}
	callables := derivCalls(this.refs, this.patternsMap[s])
	callTree := newCallTree(callables)
	memCallTree := newMemCallTree(s, &this.stackElms, &this.patternsMap, &this.zis, callTree)
	this.callTrees[s] = memCallTree
	return memCallTree
}

func (this *mem) getNullable(child int) int {
	nullIndex, ok := this.stateToNullable[child]
	if ok {
		return nullIndex
	}
	childPatterns := this.patternsMap[child]
	nullable := nullables(this.refs, childPatterns)
	nullIndex = this.nullables.add(nullable)
	this.stateToNullable[child] = nullIndex
	return nullIndex
}

func (this *mem) getReturn(stackIndex int, child int) int {
	children, ok := this.returns[stackIndex]
	if ok {
		nullIndex := this.getNullable(child)
		ret, ok := children[nullIndex]
		if ok {
			return ret
		}
	} else {
		this.returns[stackIndex] = make(map[int]int)
	}
	stackElm := this.stackElms[stackIndex]
	nullIndex := this.getNullable(child)
	zullable := this.nullables[nullIndex]
	zindex := stackElm.zindex
	nullable := unzipb(zullable, this.zis[zindex])
	current := stackElm.state
	currentPatterns := this.patternsMap[current]
	currentPatterns = derivReturns(this.refs, currentPatterns, nullable)
	simplePatterns := simps(this.refs, currentPatterns)
	res := this.patternsMap.add(simplePatterns)
	this.returns[stackIndex][nullIndex] = res
	return res
}

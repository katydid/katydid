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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
)

//This is a naive implementation and it does not handle left recursion
func (mem *Mem) Interpret(p parser.Interface) bool {
	final := deriv(mem, mem.Start, p)
	return mem.accept(final)
}

func New(g *ast.Grammar) *Mem {
	refs := ast.NewRefLookup(g)
	mem := newMem(refs)
	return mem
}

type Mem struct {
	Refs        map[string]*ast.Pattern
	PatternsMap PatternsIndexedSet
	Calls       map[int]*CallNode
	Returns     map[int]map[int]int
	Escapables  map[int]bool
	Zis         IntsIndexedSet
	Start       int

	StackElms       PairIndexedSet
	StateToNullable map[int]int
	Nullables       BoolsIndexedSet

	Accept map[int]bool
}

func newMem(refs map[string]*ast.Pattern) *Mem {
	m := &Mem{
		Refs:        refs,
		PatternsMap: newPatternsIndexedSet(),
		Calls:       make(map[int]*CallNode),
		Returns:     make(map[int]map[int]int),
		Escapables:  make(map[int]bool),
		Zis:         newIntsIndexedSet(),

		StackElms:       newPairIndexedSet(),
		StateToNullable: make(map[int]int),
		Nullables:       newBoolsIndexedSet(),
		Accept:          make(map[int]bool),
	}
	start := m.PatternsMap.add([]*ast.Pattern{refs["main"]})
	m.Start = start
	return m
}

func (this *Mem) escapable(s int) bool {
	if e, ok := this.Escapables[s]; ok {
		return e
	}
	patterns := this.PatternsMap[s]
	e := escapable(patterns)
	this.Escapables[s] = e
	return e
}

func (this *Mem) accept(s int) bool {
	if a, ok := this.Accept[s]; ok {
		return a
	}
	patterns := this.PatternsMap[s]
	if len(patterns) != 1 {
		this.Accept[s] = false
		return false
	}
	accept := interp.Nullable(this.Refs, patterns[0])
	this.Accept[s] = accept
	return accept
}

func (this *Mem) getCallTree(s int) *CallNode {
	ct, ok := this.Calls[s]
	if ok {
		return ct
	}
	callables := derivCalls(this.Refs, this.PatternsMap[s])
	callTree := newCallTree(callables)
	memCallTree := newMemCallTree(s, &this.StackElms, &this.PatternsMap, &this.Zis, callTree)
	this.Calls[s] = memCallTree
	return memCallTree
}

func (this *Mem) getNullable(state int) int {
	nullIndex, ok := this.StateToNullable[state]
	if ok {
		return nullIndex
	}
	childPatterns := this.PatternsMap[state]
	nullable := nullables(this.Refs, childPatterns)
	nullIndex = this.Nullables.add(nullable)
	this.StateToNullable[state] = nullIndex
	return nullIndex
}

func (this *Mem) getReturnn(stackIndex int, nullIndex int) int {
	children, ok := this.Returns[stackIndex]
	if ok {
		ret, ok := children[nullIndex]
		if ok {
			return ret
		}
	} else {
		this.Returns[stackIndex] = make(map[int]int)
	}
	stackElm := this.StackElms[stackIndex]
	zullable := this.Nullables[nullIndex]
	zindex := stackElm.Zindex
	nullable := unzipb(zullable, this.Zis[zindex])
	current := stackElm.State
	currentPatterns := this.PatternsMap[current]
	currentPatterns = derivReturns(this.Refs, currentPatterns, nullable)
	simplePatterns := simps(this.Refs, currentPatterns)
	res := this.PatternsMap.add(simplePatterns)
	this.Returns[stackIndex][nullIndex] = res
	return res
}

func (this *Mem) getReturn(stackIndex int, child int) int {
	nullIndex := this.getNullable(child)
	return this.getReturnn(stackIndex, nullIndex)
}

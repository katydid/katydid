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

//Package mem contains functions to interpret and memoize the execution of the grammar.
//
//TODO: cleanup
package mem

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
)

//New creates a new memoizable grammar.
func New(g *ast.Grammar) *Mem {
	refs := ast.NewRefLookup(g)
	for name, p := range refs {
		refs[name] = interp.Simplify(refs, p)
	}
	m := &Mem{
		refs:       refs,
		patterns:   newPatternsSet(),
		Calls:      []*CallNode{},
		Returns:    []map[int]int{},
		Escapables: []bool{},
		zis:        newIntsSet(),

		stackElms:       newPairSet(),
		StateToNullable: []int{},
		nullables:       newBitsetSet(),
		Accept:          []bool{},
	}
	start := m.patterns.add([]*ast.Pattern{refs["main"]})
	m.Start = start
	return m
}

//Interpret interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//The intermediate results are memoized to help with the speed of future executions.
//
//NOTE: This is a naive implementation and it does not handle left recursion.
func (mem *Mem) Interpret(p parser.Interface) bool {
	final := deriv(mem, mem.Start, p)
	return mem.accept(final)
}

//Mem is the structure containing the memoized grammar.
//TODO make more private fields.
type Mem struct {
	refs            map[string]*ast.Pattern
	patterns        patternsSet
	Calls           []*CallNode
	Returns         []map[int]int
	Escapables      []bool
	zis             intsSet
	Start           int
	stackElms       pairSet
	StateToNullable []int
	nullables       bitsetSet
	Accept          []bool
}

func escapable(patterns []*ast.Pattern) bool {
	for _, pattern := range patterns {
		if pattern.ZAny != nil {
			continue
		}
		if pattern.Not != nil {
			if pattern.GetNot().GetPattern().ZAny != nil {
				continue
			}
		}
		return true
	}
	return false
}

func (this *Mem) escapable(s int) bool {
	if len(this.Escapables) <= s {
		for i := len(this.Escapables); i <= s; i++ {
			patterns := this.patterns[i]
			this.Escapables = append(this.Escapables, escapable(patterns))
		}
	}
	return this.Escapables[s]
}

func (this *Mem) accept(s int) bool {
	if len(this.Accept) <= s {
		for i := len(this.Accept); i <= s; i++ {
			patterns := this.patterns[i]
			if len(patterns) != 1 {
				this.Accept = append(this.Accept, false)
			} else {
				this.Accept = append(this.Accept, interp.Nullable(this.refs, patterns[0]))
			}
		}
	}
	return this.Accept[s]
}

func (this *Mem) getCallTree(s int) *CallNode {
	if len(this.Calls) <= s {
		for i := len(this.Calls); i <= s; i++ {
			callables := derivCalls(this.refs, this.patterns[i])
			callTree := newCallTree(callables)
			memCallTree := newMemCallTree(s, &this.stackElms, &this.patterns, &this.zis, callTree)
			this.Calls = append(this.Calls, memCallTree)
		}
	}
	return this.Calls[s]
}

func nullables(refs map[string]*ast.Pattern, patterns []*ast.Pattern) bitset {
	nulls := newBitSet(len(patterns))
	for i, p := range patterns {
		nulls.set(i, interp.Nullable(refs, p))
	}
	return nulls
}

func (this *Mem) getNullable(s int) int {
	if len(this.StateToNullable) <= s {
		for i := len(this.StateToNullable); i <= s; i++ {
			childPatterns := this.patterns[s]
			nullable := nullables(this.refs, childPatterns)
			nullIndex := this.nullables.add(nullable)
			this.StateToNullable = append(this.StateToNullable, nullIndex)
		}
	}
	return this.StateToNullable[s]
}

func simps(refs map[string]*ast.Pattern, patterns []*ast.Pattern) []*ast.Pattern {
	for i := range patterns {
		patterns[i] = interp.Simplify(refs, patterns[i])
	}
	return patterns
}

func (this *Mem) getReturnn(stackIndex int, nullIndex int) int {
	if len(this.Returns) <= stackIndex {
		for i := len(this.Returns); i <= stackIndex; i++ {
			this.Returns = append(this.Returns, make(map[int]int))
		}
	}
	ret, ok := this.Returns[stackIndex][nullIndex]
	if ok {
		return ret
	}
	stackElm := this.stackElms[stackIndex]
	zullable := this.nullables[nullIndex]
	zindex := stackElm.Zindex
	nullable := unzipb(zullable, this.zis[zindex])
	current := stackElm.State
	currentPatterns := this.patterns[current]
	currentPatterns = derivReturns(this.refs, currentPatterns, nullable)
	simplePatterns := simps(this.refs, currentPatterns)
	res := this.patterns.add(simplePatterns)
	this.Returns[stackIndex][nullIndex] = res
	return res
}

func (this *Mem) getReturn(stackIndex int, child int) int {
	nullIndex := this.getNullable(child)
	return this.getReturnn(stackIndex, nullIndex)
}

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
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/interp"
)

//New creates a new memoizable grammar.
func New(g *ast.Grammar) (*Mem, error) {
	return new(g, false)
}

//New creates a new memoizable grammar which is optimized for records.
//A record can be json, protobuf, reflected go structures or any structure that have unique field names for each structure.
//XML would be an example of a structure for which this simplification is NOT appropriate.
func NewRecord(g *ast.Grammar) (*Mem, error) {
	return new(g, true)
}

func new(g *ast.Grammar, record bool) (*Mem, error) {
	simp := interp.NewSimplifier(g)
	if record {
		simp = simp.OptimizeForRecord()
	}
	g = simp.Grammar()
	refs := ast.NewRefLookup(g)
	m := &Mem{
		refs:       refs,
		simplifier: simp,

		patterns:  newPatternsSet(),
		zis:       newIntsSet(),
		stackElms: newPairSet(),
		nullables: newBitsetSet(),

		Calls:           []*CallNode{},
		Returns:         []map[int]int{},
		Escapables:      []bool{},
		StateToNullable: []int{},
		Accept:          []bool{},
	}
	start := m.patterns.add([]*ast.Pattern{refs["main"]})
	e := &exprToFunc{m: make(map[*ast.Expr]funcs.Bool)}
	for _, p := range refs {
		p.Walk(e)
		if e.err != nil {
			return nil, e.err
		}
	}
	m.funcs = e.m
	m.Start = start
	return m, nil
}

//Validate interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//The intermediate results are memoized to help with the speed of future executions.
//
//NOTE: This is a naive implementation and it does not handle left recursion.
func (mem *Mem) Validate(p parser.Interface) (bool, error) {
	final, err := deriv(mem, mem.Start, p)
	if err != nil {
		return false, err
	}
	return mem.accept(final), nil
}

func (mem *Mem) SetContext(context *funcs.Context) {
	mem.context = context
	for _, f := range mem.funcs {
		compose.SetContext(f, mem.context)
	}
}

//Mem is the structure containing the memoized grammar.
//TODO make more private fields.
type Mem struct {
	refs       map[string]*ast.Pattern
	funcs      map[*ast.Expr]funcs.Bool
	simplifier interp.Simplifier
	context    *funcs.Context

	patterns  patternsSet
	zis       intsSet
	stackElms pairSet
	nullables bitsetSet

	Start           int
	Calls           []*CallNode
	Returns         []map[int]int
	Escapables      []bool
	StateToNullable []int
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

func (this *Mem) calcEscapables(upto int) {
	for i := len(this.Escapables); i <= upto; i++ {
		patterns := this.patterns[i]
		this.Escapables = append(this.Escapables, escapable(patterns))
	}
}

func (this *Mem) escapable(s int) bool {
	this.calcEscapables(s)
	return this.Escapables[s]
}

func (this *Mem) calcAccepts(upto int) {
	for i := len(this.Accept); i <= upto; i++ {
		patterns := this.patterns[i]
		if len(patterns) != 1 {
			this.Accept = append(this.Accept, false)
		} else {
			this.Accept = append(this.Accept, interp.Nullable(this.refs, patterns[0]))
		}
	}
}

func (this *Mem) accept(s int) bool {
	this.calcAccepts(s)
	return this.Accept[s]
}

func (this *Mem) getFunc(expr *ast.Expr) funcs.Bool {
	if f, ok := this.funcs[expr]; ok {
		return f
	}
	f, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	compose.SetContext(f, this.context)
	this.funcs[expr] = f
	return f
}

func (this *Mem) calcCallTrees(upto int) error {
	for i := len(this.Calls); i <= upto; i++ {
		callables := derivCalls(this.refs, this.getFunc, this.patterns[i])
		callTree := newIfExprs(callables)
		memCallTree, err := newMemCallTree(i, &this.stackElms, &this.patterns, &this.zis, callTree)
		if err != nil {
			return err
		}
		this.Calls = append(this.Calls, memCallTree)
	}
	return nil
}

func (this *Mem) getCallTree(s int) (*CallNode, error) {
	if err := this.calcCallTrees(s); err != nil {
		return nil, err
	}
	return this.Calls[s], nil
}

func nullables(refs map[string]*ast.Pattern, patterns []*ast.Pattern) bitset {
	nulls := newBitSet(len(patterns))
	for i, p := range patterns {
		nulls.set(i, interp.Nullable(refs, p))
	}
	return nulls
}

func (this *Mem) calcNullables(upto int) {
	for i := len(this.StateToNullable); i <= upto; i++ {
		childPatterns := this.patterns[i]
		nullable := nullables(this.refs, childPatterns)
		nullIndex := this.nullables.add(nullable)
		this.StateToNullable = append(this.StateToNullable, nullIndex)
	}
}

func (this *Mem) getNullable(s int) int {
	this.calcNullables(s)
	return this.StateToNullable[s]
}

func (this *Mem) simps(patterns []*ast.Pattern) []*ast.Pattern {
	for i := range patterns {
		patterns[i] = this.simplifier.Simplify(patterns[i])
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
	simplePatterns := this.simps(currentPatterns)
	res := this.patterns.add(simplePatterns)
	this.Returns[stackIndex][nullIndex] = res
	return res
}

func (this *Mem) getReturn(stackIndex int, child int) int {
	nullIndex := this.getNullable(child)
	return this.getReturnn(stackIndex, nullIndex)
}

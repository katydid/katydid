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
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/intern"
	"github.com/katydid/katydid/relapse/sets"
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
	c := intern.NewConstructor()
	if record {
		c = intern.NewConstructorOptimizedForRecords()
	}
	main, err := c.AddGrammar(g)
	if err != nil {
		return nil, err
	}
	m := &Mem{
		construct: c,

		states:    intern.NewSetOfPatterns(),
		stackElms: sets.NewPairs(),

		calls:   []*ifExprs{},
		returns: []map[int]int{},
	}
	start := m.states.Add([]*intern.Pattern{main})

	m.start = start
	return m, nil
}

//Validate interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//The intermediate results are memoized to help with the speed of future executions.
//
//NOTE: This is a naive implementation and it does not handle left recursion.
func (mem *Mem) Validate(p parser.Interface) (bool, error) {
	final, err := deriv(mem, mem.start, p)
	if err != nil {
		return false, err
	}
	return mem.states.Get(final).Accept, nil
}

func (mem *Mem) SetContext(context *funcs.Context) {
	mem.construct.SetContext(context)
}

//Mem is the structure containing the memoized grammar.
type Mem struct {
	construct intern.Construct

	states    *intern.SetOfPatterns
	stackElms sets.Pairs

	start   int
	calls   []*ifExprs    // state -> (ifExprs : state -> label -> state)
	returns []map[int]int // stackIndex -> nullIndex -> state
}

func (this *Mem) getCall(state int) (*ifExprs, error) {
	for i := len(this.calls); i <= state; i++ {
		listOfIfExpr := intern.DeriveCalls(this.construct, this.states.Get(i).Patterns)
		ifs := newIfExprs(listOfIfExpr)
		this.calls = append(this.calls, ifs)
	}
	return this.calls[state], nil
}

func (this *Mem) getReturn(stackIndex int, nullIndex int) (int, error) {
	if len(this.returns) <= stackIndex {
		for i := len(this.returns); i <= stackIndex; i++ {
			this.returns = append(this.returns, make(map[int]int))
		}
	}
	if ret, ok := this.returns[stackIndex][nullIndex]; ok {
		return ret, nil
	}
	stackElm := this.stackElms[stackIndex]
	zullable := this.states.SetOfBits[nullIndex]
	childrenZipper := stackElm.Second
	nullable := sets.UnzipBits(zullable, this.states.SetOfZipIndexes[childrenZipper])
	parentPatterns := stackElm.First
	currentPatterns := this.states.Get(parentPatterns).Patterns
	retPatterns, err := intern.DeriveReturns(this.construct, currentPatterns, nullable)
	if err != nil {
		return 0, err
	}
	res := this.states.Add(retPatterns)
	this.returns[stackIndex][nullIndex] = res
	return res, nil
}

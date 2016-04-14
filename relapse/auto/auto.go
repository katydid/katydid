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

//Package auto represents the compilation and execution of a visual pushdown automaton from a parsed relapse grammar.
//
//TODO: cleanup
//
//TODO: optimize
package auto

import (
	"fmt"
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
	"io"
	"reflect"
)

//Auto is the structure that represents the automaton.
type Auto struct {
	maps   *auto
	slices *autos
}

type auto struct {
	calls           map[int]*mem.CallNode
	returns         map[int]map[int]int
	escapables      map[int]bool
	start           int
	stateToNullable map[int]int
	accept          map[int]bool
}

type autos struct {
	calls           []*mem.CallNode
	returns         [][]int
	escapables      []bool
	start           int
	stateToNullable []int
	accept          []bool
}

//Compile compiles a parsed relapse grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar) *Auto {
	m := mem.Compile(g)
	auto := &auto{
		calls:           m.Calls,
		returns:         m.Returns,
		escapables:      m.Escapables,
		start:           m.Start,
		stateToNullable: m.StateToNullable,
		accept:          m.Accept,
	}
	if size(auto) > 1000 {
		return &Auto{
			maps: auto,
		}
	}
	maxState := maxState(auto)
	calls := make([]*mem.CallNode, maxState+1)
	for k := range auto.calls {
		calls[k] = auto.calls[k]
	}
	returns := make([][]int, maxReturns(auto)+1)
	for k := range auto.returns {
		returns[k] = make([]int, maxReturn2(auto.returns[k])+1)
		for rk := range auto.returns[k] {
			returns[k][rk] = auto.returns[k][rk]
		}
	}
	escs := make([]bool, maxState+1)
	for k := range auto.escapables {
		escs[k] = auto.escapables[k]
	}
	nullables := make([]int, maxState+1)
	for k := range auto.stateToNullable {
		nullables[k] = auto.stateToNullable[k]
	}
	accepts := make([]bool, maxState+1)
	for k := range auto.accept {
		accepts[k] = auto.accept[k]
	}
	return &Auto{
		slices: &autos{
			calls:           calls,
			returns:         returns,
			escapables:      escs,
			start:           auto.start,
			stateToNullable: nullables,
			accept:          accepts,
		},
	}
}

//Execute executes an automaton with the given parser and returns whether the parser is valid given the automaton's original grammar.
func Execute(auto *Auto, p parser.Interface) bool {
	if auto.maps != nil {
		final := deriv(auto.maps, auto.maps.start, p)
		return auto.maps.accept[final]
	}
	final := derivs(auto.slices, auto.slices.start, p)
	return auto.slices.accept[final]
}

//Implements returns all funcs in the compiled automaton that implements a specific interface.
func Implements(auto *Auto, typ reflect.Type) []interface{} {
	if auto.maps != nil {
		allis := []interface{}{}
		for _, call := range auto.maps.calls {
			is := mem.Implements(call, typ)
			allis = append(allis, is...)
		}
		return allis
	}
	allis := []interface{}{}
	for _, call := range auto.slices.calls {
		is := mem.Implements(call, typ)
		allis = append(allis, is...)
	}
	return allis
}

func maxState(auto *auto) int {
	maxCalls := 0
	for i := range auto.calls {
		if i > maxCalls {
			maxCalls = i
		}
	}
	maxEscapes := 0
	for i := range auto.escapables {
		if i > maxEscapes {
			maxEscapes = i
		}
	}
	maxStateToNullable := 0
	for i := range auto.stateToNullable {
		if i > maxStateToNullable {
			maxStateToNullable = i
		}
	}
	maxAccept := 0
	for i := range auto.accept {
		if i > maxAccept {
			maxAccept = i
		}
	}
	if maxCalls != maxEscapes || maxCalls != maxStateToNullable || maxCalls != maxAccept {
		panic(fmt.Sprintf("states: maxCalls(%d) != maxEscapes(%d) != maxStateToNullabe(%d) != maxAccept(%d)", maxCalls, maxEscapes, maxStateToNullable, maxAccept))
	}
	return maxCalls
}

func maxReturns(auto *auto) int {
	maxReturns := 0
	for i := range auto.returns {
		if i > maxReturns {
			maxReturns = i
		}
	}
	return maxReturns
}

func maxReturn2(returns map[int]int) int {
	maxReturns2 := 0
	for j := range returns {
		if j > maxReturns2 {
			maxReturns2 = j
		}
	}
	return maxReturns2
}

func size(auto *auto) int {
	maxReturns2 := 0
	for i := range auto.returns {
		j := maxReturn2(auto.returns[i])
		if j > maxReturns2 {
			maxReturns2 = j
		}
	}
	states := maxState(auto)
	maxReturns := maxReturns(auto)
	size := 4*states + maxReturns*maxReturns2
	//fmt.Printf("size = %d, states = %d, maxReturns = %d, maxReturns2 = %d\n", size, states, maxReturns, maxReturns2)
	return size
}

func derivs(auto *autos, current int, tree parser.Interface) int {
	for {
		if !auto.escapables[current] {
			return current
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callTree := auto.calls[current]
		childState, stackElm := callTree.Eval(tree)
		if !tree.IsLeaf() {
			tree.Down()
			childState = derivs(auto, childState, tree)
			tree.Up()
		}
		nullIndex := auto.stateToNullable[childState]
		current = auto.returns[stackElm][nullIndex]
	}
	return current
}

func deriv(auto *auto, current int, tree parser.Interface) int {
	for {
		if !auto.escapables[current] {
			return current
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callTree := auto.calls[current]
		childState, stackElm := callTree.Eval(tree)
		if !tree.IsLeaf() {
			tree.Down()
			childState = deriv(auto, childState, tree)
			tree.Up()
		}
		nullIndex := auto.stateToNullable[childState]
		current = auto.returns[stackElm][nullIndex]
	}
	return current
}

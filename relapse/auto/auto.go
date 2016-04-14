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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
	"io"
	"reflect"
)

//Compile compiles a parsed relapse grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar) *Auto {
	mem := mem.Compile(g)
	return &Auto{
		Refs:        mem.Refs,
		PatternsMap: mem.PatternsMap,
		Calls:       mem.Calls,
		Returns:     mem.Returns,
		Escapables:  mem.Escapables,
		Zis:         mem.Zis,
		Start:       mem.Start,

		StackElms:       mem.StackElms,
		StateToNullable: mem.StateToNullable,
		Nullables:       mem.Nullables,
		Accept:          mem.Accept,
	}
}

//Execute executes an automaton with the given parser and returns whether the parser is valid given the automaton's original grammar.
func Execute(auto *Auto, p parser.Interface) bool {
	final := deriv(auto, auto.Start, p)
	return auto.Accept[final]
}

//Implements returns all funcs in the compiled automaton that implements a specific interface.
func Implements(auto *Auto, typ reflect.Type) []interface{} {
	allis := []interface{}{}
	for _, call := range auto.Calls {
		is := mem.Implements(call, typ)
		allis = append(allis, is...)
	}
	return allis
}

//Auto is the structure that represents the automaton.
//TODO make more private fields.
type Auto struct {
	Refs        map[string]*ast.Pattern
	PatternsMap mem.PatternsIndexedSet
	Calls       map[int]*mem.CallNode
	Returns     map[int]map[int]int
	Escapables  map[int]bool
	Zis         mem.IntsIndexedSet
	Start       int

	StackElms       mem.PairIndexedSet
	StateToNullable map[int]int
	Nullables       mem.BoolsIndexedSet
	Accept          map[int]bool
}

func deriv(auto *Auto, current int, tree parser.Interface) int {
	for {
		if !auto.Escapables[current] {
			return current
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callTree := auto.Calls[current]
		childState, stackElm := callTree.Eval(tree)
		if !tree.IsLeaf() {
			tree.Down()
			childState = deriv(auto, childState, tree)
			tree.Up()
		}
		nullIndex := auto.StateToNullable[childState]
		current = auto.Returns[stackElm][nullIndex]
	}
	return current
}

func checkderiv(auto *Auto, current int, tree parser.Interface) int {
	for {
		if esc, ok := auto.Escapables[current]; ok {
			if !esc {
				// mem.Prints("!Escapable", current, auto.PatternsMap[current])
				return current
			}
		} else {
			panic("wtf")
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		callTree, cok := auto.Calls[current]
		if !cok {
			panic("wtf")
		}
		childState, stackElm := callTree.Eval(tree)
		//callpair := auto.StackElms[stackElm]
		// mem.Prints("call", callpair.State, auto.PatternsMap[callpair.State])
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			childState = deriv(auto, childState, tree)
			tree.Up()
		}
		nullIndex, nok := auto.StateToNullable[childState]
		if !nok {
			panic("wtf")
		}
		//retpair := auto.StackElms[stackElm]
		//mem.Prints("return", retpair.State, auto.PatternsMap[retpair.State])
		children, rok := auto.Returns[stackElm]
		if !rok {
			panic("wtf")
		}
		current, rok = children[nullIndex]
		if !rok {
			panic("wtf")
		}
	}
	return current
}

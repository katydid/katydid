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

package convert

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
	"github.com/katydid/katydid/serialize"
	"io"
)

func Compile(g *relapse.Grammar) *Auto {
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

type Auto struct {
	Refs        map[string]*relapse.Pattern
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

func Interpret(auto *Auto, parser serialize.Parser) bool {
	final := deriv(auto, auto.Start, parser)
	return auto.Accept[final]
}

func deriv(auto *Auto, current int, tree serialize.Parser) int {
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

func checkderiv(auto *Auto, current int, tree serialize.Parser) int {
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

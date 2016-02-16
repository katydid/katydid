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
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
	"github.com/katydid/katydid/serialize"
	"io"
)

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
		if esc, ok := auto.Escapables[current]; ok {
			if !esc {
				mem.Prints("!Escapable", current, auto.PatternsMap[current])
				return current
			}
		} else {
			panic("wtf")
		}
		fmt.Printf("Next\n")
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
		callpair := auto.StackElms[stackElm]
		mem.Prints("call", callpair.State, auto.PatternsMap[callpair.State])
		if tree.IsLeaf() {
			//do nothing
		} else {
			fmt.Printf("Down\n")
			tree.Down()
			childState = deriv(auto, childState, tree)
			tree.Up()
			fmt.Printf("Up\n")
		}
		nullIndex, nok := auto.StateToNullable[childState]
		if !nok {
			panic("wtf")
		}
		retpair := auto.StackElms[stackElm]
		mem.Prints("return", retpair.State, auto.PatternsMap[retpair.State])
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

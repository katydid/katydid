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

//Package auto compiles a parsed relapse grammar into a visual pushdown automaton and executes it.
package auto

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/mem"
	"io"
	"reflect"
)

//Auto is the structure that represents the automaton.
type Auto struct {
	calls           []*mem.CallNode
	returns         []intmap
	escapables      []bool
	start           int
	stateToNullable []int
	accept          []bool
}

//Execute executes an automaton with the given parser and returns whether the parser is valid given the automaton's original grammar.
func Execute(auto *Auto, p parser.Interface) bool {
	final := deriv(auto, auto.start, p)
	return auto.accept[final]
}

//Implements returns all funcs in the compiled automaton that implements a specific interface.
func Implements(auto *Auto, typ reflect.Type) []interface{} {
	allis := []interface{}{}
	for _, call := range auto.calls {
		is := mem.Implements(call, typ)
		allis = append(allis, is...)
	}
	return allis
}

func deriv(auto *Auto, current int, tree parser.Interface) int {
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
		current = auto.returns[stackElm].lookup(nullIndex)
	}
	return current
}

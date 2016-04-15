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

package auto

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
)

//Compile compiles a parsed relapse grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar) *Auto {
	m := mem.Compile(g)
	a := &Auto{
		calls:           newCalls(m.Calls),
		returns:         newReturns(m.Returns),
		escapables:      newEscapes(m.Escapables),
		start:           m.Start,
		stateToNullable: newNullables(m.StateToNullable),
		accept:          newAccepts(m.Accept),
	}
	numStates := len(a.calls)
	if numStates != len(a.escapables) || numStates != len(a.stateToNullable) || numStates != len(a.accept) {
		panic(fmt.Sprintf("states: calls(%d) != escapes(%d) != nullables(%d) != accepts(%d)", len(a.calls), len(a.escapables), len(a.stateToNullable), len(a.accept)))
	}
	return a
}

func newCalls(callsmap map[int]*mem.CallNode) []*mem.CallNode {
	max := 0
	for i := range callsmap {
		if i > max {
			max = i
		}
	}
	callslice := make([]*mem.CallNode, max+1)
	for k := range callsmap {
		callslice[k] = callsmap[k]
	}
	return callslice
}

func newEscapes(escs map[int]bool) []bool {
	return newBoolSlice(escs)
}

func newBoolSlice(m map[int]bool) []bool {
	max := 0
	for i := range m {
		if i > max {
			max = i
		}
	}
	s := make([]bool, max+1)
	for k := range m {
		s[k] = m[k]
	}
	return s
}

func newAccepts(accept map[int]bool) []bool {
	return newBoolSlice(accept)
}

func newNullables(m map[int]int) []int {
	max := 0
	for i := range m {
		if i > max {
			max = i
		}
	}
	s := make([]int, max+1)
	for k := range m {
		s[k] = m[k]
	}
	return s
}

func newReturns(m map[int]map[int]int) []intmap {
	max := 0
	for i := range m {
		if i > max {
			max = i
		}
	}
	returns := make([]intmap, max+1)
	for k := range m {
		returns[k] = newIntMap(m[k])
	}
	return returns
}

//intmap is a little bit faster than a map, since lots of cases have a very small key space,
//but it still covers the case where we have a sparse key space.
type intmap struct {
	slice []int
	m     map[int]int
}

func newIntMap(m map[int]int) intmap {
	max := 10
	i := intmap{
		slice: make([]int, max),
		m:     make(map[int]int),
	}
	for k, v := range m {
		if k < max {
			i.slice[k] = v
		} else {
			i.m[k] = v
		}
	}
	return i
}

func (this intmap) lookup(i int) int {
	if i < len(this.slice) {
		return this.slice[i]
	}
	return this.m[i]
}

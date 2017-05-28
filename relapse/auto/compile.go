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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
)

//Compile compiles a parsed relapse grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar) (*Auto, error) {
	return compile(g, false)
}

//CompileRecord compiles a parsed relapse grammar and optimizes it for the case where the input structures are records.
func CompileRecord(g *ast.Grammar) (*Auto, error) {
	return compile(g, true)
}

func compile(g *ast.Grammar, record bool) (*Auto, error) {
	var m *mem.Mem
	var err error
	if record {
		m, err = mem.NewRecord(g)
	} else {
		m, err = mem.New(g)
	}
	if err != nil {
		return nil, err
	}
	if err := m.Compile(); err != nil {
		return nil, err
	}
	a := &Auto{
		calls:           m.Calls,
		returns:         newReturns(m.Returns),
		escapables:      m.Escapables,
		start:           m.Start,
		stateToNullable: m.StateToNullable,
		accept:          m.Accept,
	}
	return a, nil
}

func newReturns(m []map[int]int) []intmap {
	returns := make([]intmap, len(m))
	for k := range m {
		returns[k] = newIntMap(m[k])
	}
	return returns
}

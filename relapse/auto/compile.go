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
	"errors"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/sets"
)

//TODO document
var ErrTooManyStates = errors.New("a state explosion has occured")

//Compile compiles a parsed relapse grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar) (*Auto, error) {
	return compileAuto(g, false)
}

//CompileRecord compiles a parsed relapse grammar and optimizes it for the case where the input structures are records.
func CompileRecord(g *ast.Grammar) (*Auto, error) {
	return compileAuto(g, true)
}

func compileAuto(g *ast.Grammar, record bool) (*Auto, error) {
	c, err := newCompiler(g, record)
	if err != nil {
		return nil, err
	}
	if err := c.compile(); err != nil {
		return nil, err
	}
	a := &Auto{
		calls:           c.calls,
		returns:         c.returns,
		escapables:      c.escapables,
		start:           c.start,
		stateToNullable: c.stateToNullable,
		accept:          c.accept,
	}
	return a, nil
}

//Compile memoizes the full state space, all possible things that can be memoized.
func (c *compiler) compile() error {
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		for state := range c.patterns {
			if visited[state] {
				continue
			}
			visited[state] = true
			if err := compile(c, state); err != nil {
				return err
			}
			changed = true
		}
	}
	return nil
}

func compile(c *compiler, patterns int) error {
	c.getNullable(patterns)
	c.getAccept(patterns)
	c.escapable(patterns)

	callTree, err := c.getCallTree(patterns)
	if err != nil {
		return err
	}
	allPossibleCalls := getLeafs(callTree)
	for _, call := range allPossibleCalls {

		numOfChildPatterns := len(c.patterns[call.child])
		if numOfChildPatterns > 64 {
			return ErrTooManyStates
		}

		maxPossibleNullables := sets.NewBits(numOfChildPatterns)
		for i := 0; i < numOfChildPatterns; i++ {
			maxPossibleNullables.Set(i, true)
		}

		possibleNullables := sets.NewBits(numOfChildPatterns)
		for {
			nullIndex := c.nullables.Add(possibleNullables)
			c.getReturn(call.stackIndex, nullIndex)
			if possibleNullables.Equal(maxPossibleNullables) {
				break
			}
			possibleNullables = possibleNullables.Inc()
		}
	}
	return nil
}

func getLeafs(callTree *callNode) []*callNode {
	if callTree.cond == nil {
		return []*callNode{callTree}
	}
	then := getLeafs(callTree.then)
	els := getLeafs(callTree.els)
	return append(then, els...)
}

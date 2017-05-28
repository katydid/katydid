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

package mem

import (
	"errors"

	"github.com/katydid/katydid/relapse/sets"
)

//TODO document
var ErrTooManyStates = errors.New("a state explosion has occured")

//Compile memoizes the full state space, all possible things that can be memoized.
func (mem *Mem) Compile() error {
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		for state := range mem.patterns {
			if visited[state] {
				continue
			}
			visited[state] = true
			if err := compile(mem, state); err != nil {
				return err
			}
			changed = true
		}
	}
	return nil
}

func compile(mem *Mem, patterns int) error {
	mem.getNullable(patterns)
	mem.accept(patterns)
	mem.escapable(patterns)

	callTree, err := mem.getCallTree(patterns)
	if err != nil {
		return err
	}
	allPossibleCalls := getLeafs(callTree)
	for _, call := range allPossibleCalls {

		numOfChildPatterns := len(mem.patterns[call.child])
		if numOfChildPatterns > 64 {
			return ErrTooManyStates
		}

		maxPossibleNullables := sets.NewBits(numOfChildPatterns)
		for i := 0; i < numOfChildPatterns; i++ {
			maxPossibleNullables.Set(i, true)
		}

		possibleNullables := sets.NewBits(numOfChildPatterns)
		for {
			nullIndex := mem.nullables.Add(possibleNullables)
			mem.getReturn(call.stackIndex, nullIndex)
			if possibleNullables.Equal(maxPossibleNullables) {
				break
			}
			possibleNullables = possibleNullables.Inc()
		}
	}
	return nil
}

func getLeafs(callTree *CallNode) []*CallNode {
	if callTree.cond == nil {
		return []*CallNode{callTree}
	}
	then := getLeafs(callTree.then)
	els := getLeafs(callTree.els)
	return append(then, els...)
}

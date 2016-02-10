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
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
)

type state int

type mem struct {
	refs        map[string]*relapse.Pattern
	patternsMap [][]*relapse.Pattern
	callTrees   map[state]*memCallNode
	returns     map[state]map[string]map[state]state
	escapables  map[state]bool
}

func newMem(refs map[string]*relapse.Pattern) *mem {
	return &mem{
		refs:        refs,
		patternsMap: [][]*relapse.Pattern{},
		callTrees:   make(map[state]*memCallNode),
		returns:     make(map[state]map[string]map[state]state),
		escapables:  make(map[state]bool),
	}
}

func (this *mem) escapable(s state) bool {
	if e, ok := this.escapables[s]; ok {
		return e
	}
	patterns := this.patternsMap[s]
	e := escapable(patterns)
	this.escapables[s] = e
	return e
}

func (this *mem) accept(s state) bool {
	patterns := this.patternsMap[s]
	if len(patterns) != 1 {
		return false
	}
	return interp.Nullable(this.refs, patterns[0])
}

func (this *mem) lookup(patterns []*relapse.Pattern) state {
	for i, ps := range this.patternsMap {
		if relapse.Equals(ps, patterns) {
			return state(i)
		}
	}
	return state(-1)
}

func (this *mem) add(patterns []*relapse.Pattern) state {
	index := this.lookup(patterns)
	if index != -1 {
		return index
	}
	this.patternsMap = append(this.patternsMap, patterns)
	return state(len(this.patternsMap) - 1)
}

func (this *mem) getCallTree(s state) *memCallNode {
	ct, ok := this.callTrees[s]
	if ok {
		return ct
	}
	callables := derivCalls(this.refs, this.patternsMap[s])
	callTree := newCallTree(callables)
	memCallTree := newMemCallTree(this, callTree)
	this.callTrees[s] = memCallTree
	return memCallTree
}

func (this *mem) getReturn(current state, zi []int, child state) state {
	zistr := fmt.Sprintf("%v", zi)
	zis, ok := this.returns[current]
	if ok {
		children, ok := zis[zistr]
		if ok {
			ret, ok := children[child]
			if ok {
				return ret
			}
		} else {
			this.returns[current][zistr] = make(map[state]state)
		}
	} else {
		this.returns[current] = make(map[string]map[state]state)
		this.returns[current][zistr] = make(map[state]state)
	}
	zchild := this.patternsMap[child]
	childPatterns := unzip(zchild, zi)
	nullable := nullables(this.refs, childPatterns)
	currentPatterns := this.patternsMap[current]
	currentPatterns = derivReturns(this.refs, currentPatterns, nullable)
	simplePatterns := simps(this.refs, currentPatterns)
	res := this.add(simplePatterns)
	this.returns[current][zistr][child] = res
	return res
}

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
	"github.com/katydid/katydid/relapse/ast"
)

type IntsIndexedSet [][]int

func newIntsIndexedSet() IntsIndexedSet {
	return IntsIndexedSet([][]int{})
}

func (this IntsIndexedSet) index(is []int) int {
	for i, iis := range this {
		if len(is) == len(iis) {
			eq := true
			for j := range iis {
				if iis[j] != is[j] {
					eq = false
					break
				}
			}
			if eq {
				return i
			}
		}
	}
	return -1
}

func (this *IntsIndexedSet) add(zi []int) int {
	index := this.index(zi)
	if index != -1 {
		return index
	}
	*this = append(*this, zi)
	return len(*this) - 1
}

type PatternsIndexedSet [][]*relapse.Pattern

func newPatternsIndexedSet() PatternsIndexedSet {
	return PatternsIndexedSet([][]*relapse.Pattern{})
}

func (this PatternsIndexedSet) index(patterns []*relapse.Pattern) int {
	for i, ps := range this {
		if relapse.Equals(ps, patterns) {
			return i
		}
	}
	return -1
}

func (this *PatternsIndexedSet) add(patterns []*relapse.Pattern) int {
	index := this.index(patterns)
	if index != -1 {
		return index
	}
	*this = append(*this, patterns)
	return len(*this) - 1
}

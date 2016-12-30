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

//intsSet represents an indexed list of list of integers.
//It reverse maps a list of ints into a single int.
type intsSet [][]int

func newIntsSet() intsSet {
	return intsSet([][]int{})
}

func (this intsSet) index(is []int) int {
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

func (this *intsSet) add(is []int) int {
	index := this.index(is)
	if index != -1 {
		return index
	}
	*this = append(*this, is)
	return len(*this) - 1
}

//patternsSet represents an indexed list of list of Patterns.
//It reverse maps a list of Patterns into a single int.
type patternsSet [][]*ast.Pattern

func newPatternsSet() patternsSet {
	return patternsSet([][]*ast.Pattern{})
}

func (this patternsSet) index(patterns []*ast.Pattern) int {
	for i, ps := range this {
		if ast.Equals(ps, patterns) {
			return i
		}
	}
	return -1
}

func (this *patternsSet) add(patterns []*ast.Pattern) int {
	index := this.index(patterns)
	if index != -1 {
		return index
	}
	*this = append(*this, patterns)
	return len(*this) - 1
}

//bitsetSet represents an indexed list of bitsets.
//It reverse maps a bitset into a single int.
type bitsetSet []bitset

func newBitsetSet() bitsetSet {
	return bitsetSet([]bitset{})
}

func (this bitsetSet) index(bs bitset) int {
	for i, ibs := range this {
		if ibs.equal(bs) {
			return i
		}
	}
	return -1
}

func (this *bitsetSet) add(bs bitset) int {
	index := this.index(bs)
	if index != -1 {
		return index
	}
	*this = append(*this, bs)
	return len(*this) - 1
}

type stackElm struct {
	parentPatterns int
	childrenZipper int
}

//pairSet represents an indexed list of stackElm pairs.
//It reverse maps a list of stackElms into a single int.
type pairSet []stackElm

func newPairSet() pairSet {
	return pairSet([]stackElm{})
}

func (this pairSet) index(se stackElm) int {
	for i, ise := range this {
		if ise == se {
			return i
		}
	}
	return -1
}

func (this *pairSet) add(se stackElm) int {
	index := this.index(se)
	if index != -1 {
		return index
	}
	*this = append(*this, se)
	return len(*this) - 1
}

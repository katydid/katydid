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

package sets

import (
	"github.com/katydid/katydid/relapse/ast"
)

//Ints represents an indexed list of list of integers.
//It reverse maps a list of ints into a single int.
type Ints [][]int

func NewInts() Ints {
	return Ints([][]int{})
}

func (this Ints) Index(is []int) int {
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

func (this *Ints) Add(is []int) int {
	index := this.Index(is)
	if index != -1 {
		return index
	}
	*this = append(*this, is)
	return len(*this) - 1
}

//Patterns represents an indexed list of list of Patterns.
//It reverse maps a list of Patterns into a single int.
type Patterns [][]*ast.Pattern

func NewPatterns() Patterns {
	return Patterns([][]*ast.Pattern{})
}

func (this Patterns) Index(patterns []*ast.Pattern) int {
	for i, ps := range this {
		// TODO maybe we should rather use hashes to do this more efficiently?
		if deriveEquals(ps, patterns) {
			return i
		}
	}
	return -1
}

func (this *Patterns) Add(patterns []*ast.Pattern) int {
	index := this.Index(patterns)
	if index != -1 {
		return index
	}
	*this = append(*this, patterns)
	return len(*this) - 1
}

//BitsSet represents an indexed list of Bits.
//It reverse maps a Bits into a single int.
type BitsSet []Bits

func NewBitsSet() BitsSet {
	return BitsSet([]Bits{})
}

func (this BitsSet) Index(bs Bits) int {
	for i, ibs := range this {
		if ibs.Equal(bs) {
			return i
		}
	}
	return -1
}

func (this *BitsSet) Add(bs Bits) int {
	index := this.Index(bs)
	if index != -1 {
		return index
	}
	*this = append(*this, bs)
	return len(*this) - 1
}

type Pair struct {
	First  int
	Second int
}

//Pairs represents an indexed list of Pair pairs.
//It reverse maps a list of Pairs into a single int.
type Pairs []Pair

func NewPairs() Pairs {
	return Pairs([]Pair{})
}

func (this Pairs) Index(se Pair) int {
	for i, ise := range this {
		if ise == se {
			return i
		}
	}
	return -1
}

func (this *Pairs) Add(se Pair) int {
	index := this.Index(se)
	if index != -1 {
		return index
	}
	*this = append(*this, se)
	return len(*this) - 1
}

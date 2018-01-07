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

package intern

import (
	"github.com/katydid/katydid/relapse/sets"
)

//SetOfPatterns represents an indexed list of list of Patterns.
//It reverse maps a list of Patterns into a single int.
type SetOfPatterns struct {
	List      []*Patterns
	Hashes    map[uint64][]int
	SetOfBits sets.BitsSet
}

func NewSetOfPatterns() *SetOfPatterns {
	return &SetOfPatterns{
		List:      []*Patterns{},
		Hashes:    make(map[uint64][]int),
		SetOfBits: sets.NewBitsSet(),
	}
}

func (this *SetOfPatterns) Get(i int) *Patterns {
	return this.List[i]
}

func (this *SetOfPatterns) Index(patterns []*Pattern) int {
	h := hashes(patterns)
	pss := this.Hashes[h]
	for _, index := range pss {
		ps := this.List[index]
		if deriveEquals(ps.Patterns, patterns) {
			return index
		}
	}
	return -1
}

func (this *SetOfPatterns) Add(ps []*Pattern) int {
	index := this.Index(ps)
	if index != -1 {
		return index
	}
	h := hashes(ps)
	index = len(this.List)
	this.Hashes[h] = append(this.Hashes[h], index)
	patterns := NewPatterns(ps)
	this.List = append(this.List, patterns)
	patterns.NullIndex = this.SetOfBits.Add(patterns.Nullables)
	return index
}

type Patterns struct {
	Patterns  []*Pattern
	Escapable bool
	Nullables sets.Bits
	Accept    bool
	NullIndex int
	Zipped    *ZippedPatterns
}

func NewPatterns(ps []*Pattern) *Patterns {
	return &Patterns{
		Patterns:  ps,
		Escapable: escapable(ps),
		Nullables: newNullableSet(ps),
		Accept:    len(ps) == 1 && ps[0].nullable,
		Zipped:    Zip(ps),
	}
}

func hashes(patterns []*Pattern) uint64 {
	h := uint64(17)
	for _, pattern := range patterns {
		h = 31*h + pattern.hash
	}
	return h
}

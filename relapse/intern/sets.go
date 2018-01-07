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
	List            []*Patterns
	Hashes          map[uint64][]int
	SetOfBits       sets.BitsSet
	SetOfZipIndexes sets.Ints
}

func NewSetOfPatterns() *SetOfPatterns {
	return &SetOfPatterns{
		List:            []*Patterns{},
		Hashes:          make(map[uint64][]int),
		SetOfBits:       sets.NewBitsSet(),
		SetOfZipIndexes: sets.NewInts(),
	}
}

func (this *SetOfPatterns) Get(i int) *Patterns {
	return this.List[i]
}

func (this *SetOfPatterns) indexOf(hash uint64, patterns []*Pattern) int {
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
	h := hashes(ps)
	index := this.indexOf(h, ps)
	if index != -1 {
		return index
	}
	index = len(this.List)
	patterns := &Patterns{}
	this.List = append(this.List, patterns)
	patterns.Index = index

	this.Hashes[h] = append(this.Hashes[h], index)

	nulls := newNullableSet(ps)
	patterns.Patterns = ps
	patterns.Escapable = escapable(ps)
	patterns.Accept = len(ps) == 1 && ps[0].nullable
	zipped := Zip(ps)
	patterns.NullIndex = this.SetOfBits.Add(nulls)
	patterns.IndexOfZippedPatterns = this.Add(zipped.Patterns)
	patterns.IndexOfZippedIndexes = this.SetOfZipIndexes.Add(zipped.Indexes)

	return index
}

type Patterns struct {
	Patterns []*Pattern
	Index    int

	Escapable bool
	Accept    bool
	NullIndex int

	IndexOfZippedPatterns int
	IndexOfZippedIndexes  int
}

func hashes(patterns []*Pattern) uint64 {
	h := uint64(17)
	for _, pattern := range patterns {
		h = 31*h + pattern.hash
	}
	return h
}

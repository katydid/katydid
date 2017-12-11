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

//Patterns represents an indexed list of list of Patterns.
//It reverse maps a list of Patterns into a single int.
type Patterns struct {
	List   [][]*Pattern
	Hashes map[uint64][]int
}

func NewPatterns() *Patterns {
	return &Patterns{
		List:   [][]*Pattern{},
		Hashes: make(map[uint64][]int),
	}
}

func hashes(patterns []*Pattern) uint64 {
	h := uint64(17)
	for _, pattern := range patterns {
		h = 31*h + pattern.hash
	}
	return h
}

func (this *Patterns) Get(i int) []*Pattern {
	return this.List[i]
}

func (this *Patterns) Index(patterns []*Pattern) int {
	h := hashes(patterns)
	pss := this.Hashes[h]
	for i, index := range pss {
		ps := this.List[index]
		if deriveEquals(ps, patterns) {
			return i
		}
	}
	return -1
}

func (this *Patterns) Add(patterns []*Pattern) int {
	index := this.Index(patterns)
	if index != -1 {
		return index
	}
	h := hashes(patterns)
	index = len(this.List)
	this.List = append(this.List, patterns)
	this.Hashes[h] = append(this.Hashes[h], index)
	return index
}

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

package ast

import (
	"bytes"
	"github.com/gogo/protobuf/proto"
	"sort"
)

//Compare compares two patterns the result will be 0 if p1 == p2, -1 if p1 < p2 and +1 is p1 > p2.
//What smaller and bigger means may change over time, but this function is still useful for deterministic ordering.
//TODO this is really slow, make it faster, generate a method or simply write one.
func Compare(p1, p2 *Pattern) int {
	d1, err := proto.Marshal(p1)
	if err != nil {
		panic(err)
	}
	d2, err := proto.Marshal(p2)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(d1, d2)
}

//Less returns whether one pattern is smaller than another.
//What smaller means may change over time, but this function is still useful for deterministic ordering.
func (p1 *Pattern) Less(p2 *Pattern) bool {
	return Compare(p1, p2) < 0
}

//Index returns the index of the Pattern p in the list of Patterns ps.  If p is not found -1 is returned.
func Index(ps []*Pattern, p *Pattern) int {
	for i, pp := range ps {
		if Compare(pp, p) == 0 {
			return i
		}
	}
	return -1
}

//Remove removes the i'th index from the list of patterns.
func Remove(ps []*Pattern, i int) []*Pattern {
	return append(append([]*Pattern{}, ps[:i]...), ps[i+1:]...)
}

//Equals returns whether two lists of patterns are equal.
func Equals(this, that []*Pattern) bool {
	if len(this) != len(that) {
		return false
	}
	for i := range this {
		if !this[i].Equal(that[i]) {
			return false
		}
	}
	return true
}

//Has returns whether the Pattern p is contained in the list of Patterns ps.
func Has(ps []*Pattern, p *Pattern) bool {
	return Index(ps, p) > -1
}

//Set returns the list of Patterns as a set, where all duplicates have been removed.
func Set(ps []*Pattern) []*Pattern {
	set := make([]*Pattern, 0, len(ps))
	for i := range ps {
		if !Has(set, ps[i]) {
			set = append(set, ps[i])
		}
	}
	return set
}

//Sort sorts a list of Patterns.
func Sort(ps []*Pattern) {
	sort.Sort(Sortable(ps))
}

//Sortable attaches the sort.Interface methods to []*Pattern.
type Sortable []*Pattern

func (s Sortable) Less(i, j int) bool {
	return s[i].Less(s[j])
}

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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

package relapse

import (
	"bytes"
	"github.com/gogo/protobuf/proto"
	"sort"
)

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

func (p1 *Pattern) Less(p2 *Pattern) bool {
	return Compare(p1, p2) < 0
}

func Index(ps []*Pattern, p *Pattern) int {
	for i, pp := range ps {
		if Compare(pp, p) == 0 {
			return i
		}
	}
	return -1
}

func Remove(ps []*Pattern, index int) []*Pattern {
	return append(append([]*Pattern{}, ps[:index]...), ps[index+1:]...)
}

func Equals(this, that []*Pattern) bool {
	if len(this) != len(that) {
		return false
	}
	for _, l := range this {
		for _, r := range that {
			if !l.Equal(r) {
				return false
			}
		}
	}
	return true
}

func Has(ps []*Pattern, p *Pattern) bool {
	return Index(ps, p) > -1
}

func Set(ps []*Pattern) []*Pattern {
	set := make([]*Pattern, 0, len(ps))
	for i := range ps {
		if !Has(set, ps[i]) {
			set = append(set, ps[i])
		}
	}
	return set
}

func Sort(ps []*Pattern) {
	sort.Sort(Sortable(ps))
}

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

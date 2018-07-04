//  Copyright 2017 Walter Schulze
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
	"sort"
)

type sortable []*Pattern

func (s sortable) Len() int { return len(s) }

func (s sortable) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less says Node patterns are the largest and then falls back to comparing hashes.
func (s sortable) Less(i, j int) bool {
	if s[i].Type != Node || s[j].Type != Node {
		if s[i].Type == Node {
			// Nodes are larger
			return false
		}
		if s[j].Type == Node {
			// Nodes are larger
			return true
		}
		return s[i].hash < s[j].hash
	}
	// if both are nodes
	c := s[i].Func.Compare(s[j].Func)
	if c == 0 {
		return s[i].hash < s[j].hash
	}
	return c < 0
}

// orderedSet creates a stable order, while taking into account treenodes and removing duplicates.
func orderedSet(ps []*Pattern) []*Pattern {

	if len(ps) == 0 {
		return nil
	}

	sort.Sort(sortable(ps))

	// remove duplicates from a sorted list
	u := 0
	for i := 1; i < len(ps); i++ {
		if ps[i].Equal(ps[u]) {
			continue
		}
		u++
		if u != i {
			ps[u] = ps[i]
		}
	}
	ps = ps[:u+1]

	return ps

}

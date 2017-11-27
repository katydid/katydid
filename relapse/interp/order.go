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

package interp

import (
	"sort"

	"github.com/katydid/katydid/relapse/ast"
)

type sortable struct {
	hashes []uint64
	list   []*ast.Pattern
}

func (ps *sortable) Len() int {
	return len(ps.list)
}

func (ps *sortable) Swap(i, j int) {
	ps.list[i], ps.list[j] = ps.list[j], ps.list[i]
	ps.hashes[i], ps.hashes[j] = ps.hashes[j], ps.hashes[i]
}

// Less says TreeNode patterns are the largest and then fallsback to comparing hashes.
func (ps *sortable) Less(i, j int) bool {
	if ps.list[i].GetTreeNode() == nil && ps.list[j].GetTreeNode() == nil {
		return ps.hashes[i] < ps.hashes[j]
	}
	if ps.list[i].GetTreeNode() != nil && ps.list[j].GetTreeNode() != nil {
		c := ps.list[i].GetTreeNode().GetName().Compare(ps.list[j].GetTreeNode().GetName())
		if c == 0 {
			return ps.hashes[i] < ps.hashes[j]
		}
		return c < 0
	}
	return ps.list[i].GetTreeNode() == nil
}

// create a stable order, taking into account treenodes and remove duplicates.
func orderedSet(ps []*ast.Pattern) []*ast.Pattern {
	if len(ps) == 0 || len(ps) == 1 {
		return ps
	}

	hashes := make([]uint64, len(ps))
	for i, p := range ps {
		if p.GetTreeNode() == nil {
			hashes[i] = deriveHash(p)
		} else {
			hashes[i] = deriveHash(p.GetTreeNode().GetPattern())
		}
	}

	sortable := &sortable{hashes, ps}
	sort.Sort(sortable)

	u := 0
	for i := 1; i < len(ps); i++ {
		if hashes[i] == hashes[u] &&
			ps[i].Equal(ps[u]) {
			continue
		}
		u++
		if u != i {
			ps[u] = ps[i]
			hashes[u] = hashes[i]
		}
	}
	return ps[:u+1]

}

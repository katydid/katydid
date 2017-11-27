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

package sets

import (
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

func (ps *sortable) Less(i, j int) bool {
	return ps.hashes[i] < ps.hashes[j]
}

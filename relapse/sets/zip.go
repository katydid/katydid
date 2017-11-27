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
	"sort"

	"github.com/katydid/katydid/relapse/ast"
)

var (
	zipIgnoreSet = []*ast.Pattern{
		ast.NewZAny(),
		ast.NewNot(ast.NewZAny()),
	}
	zignoreb = []bool{
		true,
		false,
	}
)

func Zip(patterns []*ast.Pattern) ([]*ast.Pattern, []int) {

	ps := make([]*ast.Pattern, len(patterns))
	orighashes := make([]uint64, len(patterns))
	hashes := make([]uint64, len(patterns))
	for i, p := range patterns {
		ps[i] = patterns[i]
		h := p.Hash()
		hashes[i] = h
		orighashes[i] = h
	}

	// sort
	sortable := &sortable{hashes, ps}
	sort.Sort(sortable)

	// remove zany and not zany
	u := 0
	for i := 0; i < len(ps); i++ {
		if ps[i].ZAny != nil {
			continue
		}
		if ps[i].Not != nil && ps[i].Not.Pattern.ZAny != nil {
			continue
		}
		ps[u] = ps[i]
		hashes[u] = hashes[i]
		u++
	}
	ps = ps[:u]
	hashes = hashes[:u]

	// remove duplicates
	if len(ps) > 0 {
		u = 0
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
		ps = ps[:u+1]
		hashes = hashes[:u+1]
	}

	// calculate indexes by doing a reverse lookup using the original hashes and moved hashes.
	revhashes := make(map[uint64][]int)
	for i, h := range hashes {
		revhashes[h] = append(revhashes[h], i)
	}
	indexes := make([]int, len(patterns))
	for i := range patterns {
		if patterns[i].ZAny != nil {
			indexes[i] = -1
		} else if patterns[i].Not != nil && patterns[i].Not.Pattern.ZAny != nil {
			indexes[i] = -2
		} else {
			hashindexes, ok := revhashes[orighashes[i]]
			if !ok {
				panic("wtf")
			}
			if len(hashindexes) == 1 {
				indexes[i] = hashindexes[0]
			} else {
				for _, index := range hashindexes {
					if ps[index].Equal(patterns[i]) {
						indexes[i] = index
					}
				}
			}
		}
	}

	return ps, indexes
}

func Unzip(patterns []*ast.Pattern, indexes []int) []*ast.Pattern {
	res := make([]*ast.Pattern, len(indexes))
	for i, index := range indexes {
		if index < 0 {
			index += 1
			index = index * -1
			res[i] = zipIgnoreSet[index]
		} else {
			res[i] = patterns[index]
		}
	}
	return res
}

func UnzipBits(bools Bits, indexes []int) []bool {
	res := make([]bool, len(indexes))
	for i, index := range indexes {
		if index >= 0 {
			res[i] = bools.Get(index)
		} else {
			res[i] = zignoreb[(index+1)*-1]
		}
	}
	return res
}

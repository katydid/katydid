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
	set := ast.Set(patterns)
	ast.Sort(set)

	if index := ast.Index(set, zipIgnoreSet[0]); index != -1 {
		set = ast.Remove(set, index)
	}
	if index := ast.Index(set, zipIgnoreSet[1]); index != -1 {
		set = ast.Remove(set, index)
	}
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		index := ast.Index(set, pattern)
		if index == -1 {
			index = ast.Index(zipIgnoreSet, pattern)
			index *= -1
			index -= 1
		}
		indexes[i] = index
	}
	return set, indexes
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

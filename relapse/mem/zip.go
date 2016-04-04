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

package mem

import (
	"github.com/katydid/katydid/relapse/ast"
)

var (
	zignore = []*ast.Pattern{
		ast.NewZAny(),
		ast.NewNot(ast.NewZAny()),
	}
	zignoreb = []bool{
		true,
		false,
	}
)

func zip(patterns []*ast.Pattern) ([]*ast.Pattern, []int) {
	zipped := ast.Set(patterns)
	ast.Sort(zipped)

	if index := ast.Index(zipped, zignore[0]); index != -1 {
		zipped = ast.Remove(zipped, index)
	}
	if index := ast.Index(zipped, zignore[1]); index != -1 {
		zipped = ast.Remove(zipped, index)
	}
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		index := ast.Index(zipped, pattern)
		if index == -1 {
			index = ast.Index(zignore, pattern)
			index *= -1
			index -= 1
		}
		indexes[i] = index
	}
	return zipped, indexes
}

func unzip(patterns []*ast.Pattern, indexes []int) []*ast.Pattern {
	res := make([]*ast.Pattern, len(indexes))
	for i, index := range indexes {
		if index >= 0 {
			res[i] = patterns[index]
		} else {
			res[i] = zignore[(index+1)*-1]
		}
	}
	return res
}

func unzipb(bools []bool, indexes []int) []bool {
	res := make([]bool, len(indexes))
	for i, index := range indexes {
		if index >= 0 {
			res[i] = bools[index]
		} else {
			res[i] = zignoreb[(index+1)*-1]
		}
	}
	return res
}

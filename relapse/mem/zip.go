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
	zignore = []*relapse.Pattern{
		relapse.NewZAny(),
		relapse.NewNot(relapse.NewZAny()),
	}
	zignoreb = []bool{
		true,
		false,
	}
)

func zip(patterns []*relapse.Pattern) ([]*relapse.Pattern, []int) {
	zipped := relapse.Set(patterns)
	relapse.Sort(zipped)

	if index := relapse.Index(zipped, zignore[0]); index != -1 {
		zipped = relapse.Remove(zipped, index)
	}
	if index := relapse.Index(zipped, zignore[1]); index != -1 {
		zipped = relapse.Remove(zipped, index)
	}
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		index := relapse.Index(zipped, pattern)
		if index == -1 {
			index = relapse.Index(zignore, pattern)
			index *= -1
			index -= 1
		}
		indexes[i] = index
	}
	return zipped, indexes
}

func unzip(patterns []*relapse.Pattern, indexes []int) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(indexes))
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

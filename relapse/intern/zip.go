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

var (
	zipIgnoreSet = []*Pattern{
		zany,
		notzany,
	}
	zignoreb = []bool{
		true,
		false,
	}
)

func Zip(patterns []*Pattern) ([]*Pattern, []int) {

	ps := make([]*Pattern, len(patterns))
	for i := range patterns {
		ps[i] = patterns[i]
	}
	ps = orderedSet(ps)

	// remove zany and not zany
	ps = removeZAny(ps)
	ps = removeNotZAny(ps)

	// calculate indexes by doing a reverse lookup using the original hashes and moved hashes.
	revhashes := make(map[uint64][]int)
	for i, p := range ps {
		revhashes[p.hash] = append(revhashes[p.hash], i)
	}
	indexes := make([]int, len(patterns))
	for i := range patterns {
		if patterns[i] == zany {
			indexes[i] = -1
		} else if patterns[i] == notzany {
			indexes[i] = -2
		} else {
			hashindexes, ok := revhashes[patterns[i].hash]
			if !ok {
				panic("unreachable: unknown hash")
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

func Unzip(patterns []*Pattern, indexes []int) []*Pattern {
	res := make([]*Pattern, len(indexes))
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

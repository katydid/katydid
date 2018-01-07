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
		newZAny(),
		newNotZAny(),
	}
	zignoreb = []bool{
		true,
		false,
	}
)

func removeAllNotZAny(ps []*Pattern) []*Pattern {
	return deriveFilter(notEmptySet, ps)
}

func removeAllZAny(ps []*Pattern) []*Pattern {
	return deriveFilter(func(p *Pattern) bool {
		return p.Type != ZAny
	}, ps)
}

type ZippedPatterns struct {
	Patterns []*Pattern
	Indexes  []int
}

func Zip(patterns []*Pattern) *ZippedPatterns {
	zips := make([]*Pattern, len(patterns))
	for i := range patterns {
		zips[i] = patterns[i]
	}
	zips = orderedSet(zips)

	// remove zany and not zany
	zips = removeAllZAny(zips)
	zips = removeAllNotZAny(zips)

	// calculate indexes by doing a reverse lookup using the original hashes and moved hashes.
	revhashes := make(map[uint64][]int)
	for i, p := range zips {
		revhashes[p.hash] = append(revhashes[p.hash], i)
	}
	indexes := make([]int, len(patterns))
	for i := range indexes {
		if isZAny(patterns[i]) {
			indexes[i] = -1
			continue
		}
		if isNotZAny(patterns[i]) {
			indexes[i] = -2
			continue
		}
		hashindexes := revhashes[patterns[i].hash]
		if len(hashindexes) == 0 {
			panic("unreachable: unknown hash")
		}
		if len(hashindexes) == 1 {
			indexes[i] = hashindexes[0]
			continue
		}
		found := false
		for _, index := range hashindexes {
			if zips[index].Equal(patterns[i]) {
				indexes[i] = index
				found = true
				break
			}
		}
		if !found {
			panic("wtf")
		}
	}

	return &ZippedPatterns{zips, indexes}
}

func (z *ZippedPatterns) Unzip() []*Pattern {
	res := make([]*Pattern, len(z.Indexes))
	for i, index := range z.Indexes {
		if index < 0 {
			index += 1
			index = index * -1
			res[i] = zipIgnoreSet[index]
		} else {
			res[i] = z.Patterns[index]
		}
	}
	return res
}

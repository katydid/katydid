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

package auto

//intmap is a little bit faster than a map, since lots of cases have a very small key space,
//but it still covers the case where we have a sparse key space.
type intmap struct {
	slice []int
	m     map[int]int
}

func newIntMap(m map[int]int) intmap {
	max := 10
	i := intmap{
		slice: make([]int, max),
		m:     make(map[int]int),
	}
	for k, v := range m {
		if k < max {
			i.slice[k] = v
		} else {
			i.m[k] = v
		}
	}
	return i
}

func (this intmap) lookup(i int) int {
	if i < len(this.slice) {
		return this.slice[i]
	}
	return this.m[i]
}

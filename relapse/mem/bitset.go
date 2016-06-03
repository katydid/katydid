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

type bitset struct {
	val  uint64
	size int
}

func newBitSet(size int) bitset {
	return bitset{val: 0, size: size}
}

func (bs bitset) get(i int) bool {
	if i < 0 || i >= bs.size {
		panic("range check error")
	}
	return ((bs.val >> uint(i)) & 0x1) == 1
}

func (bs bitset) list() []bool {
	list := make([]bool, bs.size)
	for i := 0; i < bs.size; i++ {
		list[i] = bs.get(i)
	}
	return list
}

func (bs *bitset) set(i int, b bool) {
	if i < 0 || i >= bs.size {
		panic("range check error")
	}
	if (((bs.val >> uint(i)) & 0x1) == 1) == b {
		return
	}
	bs.val ^= 1 << uint(i)
}

func (bs *bitset) append(b bool) {
	if b {
		bs.val ^= 1 << uint(bs.size)
	}
	bs.size++
}

func (bs bitset) copy() bitset {
	return bitset{
		val:  bs.val,
		size: bs.size,
	}
}

func (this bitset) equal(that bitset) bool {
	return this == that
}


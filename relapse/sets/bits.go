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
	"fmt"
	"math"
)

type Bits struct {
	val0 uint64
	vals []uint64
	size int
}

func NewBits(size int) Bits {
	if size < 64 {
		return Bits{size: size}
	}
	return Bits{size: size, vals: make([]uint64, (size / 64))}
}

func (bs Bits) Get(i int) bool {
	if i < 0 || i >= bs.size {
		panic(fmt.Sprintf("range check error i:%d, size:%d", i, bs.size))
	}
	b := bs.val0
	if i >= 64 {
		index := (i / 64) - 1
		b = bs.vals[index]
		i = i % 64
	}
	return ((b >> uint(i)) & 0x1) == 1
}

func (bs Bits) Inc() Bits {
	if bs.size < 64 {
		return Bits{
			val0: bs.val0 + 1,
			size: bs.size,
		}
	}
	if bs.val0 < math.MaxUint64 {
		return Bits{
			val0: bs.val0 + 1,
			vals: bs.vals,
			size: bs.size,
		}
	}
	maxi := bs.size / 64
	newvals := make([]uint64, maxi)
	inced := false
	for i := 0; i < maxi; i++ {
		if !inced {
			if bs.vals[i] < math.MaxUint64 {
				newvals[i] = bs.vals[i] + 1
				inced = true
			}
		} else {
			newvals[i] = bs.vals[i]
		}
	}
	if !inced {
		panic("maximum reached")
	}
	return Bits{
		vals: newvals,
		size: bs.size,
	}
}

func (bs Bits) list() []bool {
	list := make([]bool, bs.size)
	for i := 0; i < bs.size; i++ {
		list[i] = bs.Get(i)
	}
	return list
}

func (bs *Bits) Set(i int, b bool) {
	if i < 0 || i >= bs.size {
		panic(fmt.Sprintf("range check error i:%d, size:%d", i, bs.size))
	}
	if i < 64 {
		if (((bs.val0 >> uint(i)) & 0x1) == 1) == b {
			return
		}
		bs.val0 ^= 1 << uint(i)
		return
	}
	index := (i / 64) - 1
	i = i % 64
	if (((bs.vals[index] >> uint(i)) & 0x1) == 1) == b {
		return
	}
	bs.vals[index] ^= 1 << uint(i)
}

func (this Bits) Equal(that Bits) bool {
	if this.size != that.size {
		return false
	}
	if this.val0 != that.val0 {
		return false
	}
	for i, v := range this.vals {
		if v != that.vals[i] {
			return false
		}
	}
	return true
}

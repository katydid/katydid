//  Copyright 2013 Walter Schulze
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

package maps

type IntToInt interface {
	Lookup(int) (int, bool)
	ToMap() map[int]int
}

type Uint64ToInt interface {
	Lookup(uint64) (int, bool)
	ToMap() map[uint64]int
}

const (
	zero    = 0
	one     = 1
	listMax = 14
)

func NewIntToInt(m map[int]int) IntToInt {
	if len(m) == zero {
		return &zeroIntToInt{}
	}
	if len(m) == one {
		for k, v := range m {
			return &singleIntToInt{k, v}
		}
	}
	if len(m) <= listMax {
		ls := &listIntToInt{
			ks: make([]int, len(m)),
			vs: make([]int, len(m)),
		}
		i := 0
		for k, v := range m {
			ls.ks[i] = k
			ls.vs[i] = v
			i++
		}
		return ls
	}
	mm := make(mapIntToInt)
	for k, v := range m {
		mm[k] = v
	}
	return mm
}

func NewUint64ToInt(m map[uint64]int) Uint64ToInt {
	if len(m) == zero {
		return &zeroUint64ToInt{}
	}
	if len(m) == one {
		for k, v := range m {
			return &singleUint64ToInt{k, v}
		}
	}
	if len(m) <= listMax {
		ls := &listUint64ToInt{
			ks: make([]uint64, len(m)),
			vs: make([]int, len(m)),
		}
		i := 0
		for k, v := range m {
			ls.ks[i] = k
			ls.vs[i] = v
			i++
		}
		return ls
	}
	mm := make(mapUint64ToInt)
	for k, v := range m {
		mm[k] = v
	}
	return mm
}

type mapIntToInt map[int]int

func (this mapIntToInt) Lookup(k int) (int, bool) {
	v, ok := this[k]
	return v, ok
}

func (this mapIntToInt) ToMap() map[int]int {
	return this
}

type mapUint64ToInt map[uint64]int

func (this mapUint64ToInt) Lookup(k uint64) (int, bool) {
	v, ok := this[k]
	return v, ok
}

func (this mapUint64ToInt) ToMap() map[uint64]int {
	return this
}

type zeroIntToInt struct {
}

func (this *zeroIntToInt) Lookup(k int) (int, bool) {
	return 0, false
}

func (this *zeroIntToInt) ToMap() map[int]int {
	return make(map[int]int)
}

type zeroUint64ToInt struct {
}

func (this *zeroUint64ToInt) Lookup(k uint64) (int, bool) {
	return 0, false
}

func (this *zeroUint64ToInt) ToMap() map[uint64]int {
	return make(map[uint64]int)
}

type singleIntToInt struct {
	k int
	v int
}

func (this *singleIntToInt) Lookup(k int) (int, bool) {
	if this.k == k {
		return this.v, true
	}
	return 0, false
}

func (this *singleIntToInt) ToMap() map[int]int {
	return map[int]int{this.k: this.v}
}

type singleUint64ToInt struct {
	k uint64
	v int
}

func (this *singleUint64ToInt) Lookup(k uint64) (int, bool) {
	if this.k == k {
		return this.v, true
	}
	return 0, false
}

func (this *singleUint64ToInt) ToMap() map[uint64]int {
	return map[uint64]int{this.k: this.v}
}

type listIntToInt struct {
	ks []int
	vs []int
}

func (this *listIntToInt) Lookup(key int) (int, bool) {
	for i, k := range this.ks {
		if k == key {
			return this.vs[i], true
		}
	}
	return 0, false
}

func (this *listIntToInt) ToMap() map[int]int {
	m := make(map[int]int)
	for i, k := range this.ks {
		m[k] = this.vs[i]
	}
	return m
}

type listUint64ToInt struct {
	ks []uint64
	vs []int
}

func (this *listUint64ToInt) Lookup(key uint64) (int, bool) {
	for i, k := range this.ks {
		if k == key {
			return this.vs[i], true
		}
	}
	return 0, false
}

func (this *listUint64ToInt) ToMap() map[uint64]int {
	m := make(map[uint64]int)
	for i, k := range this.ks {
		m[k] = this.vs[i]
	}
	return m
}

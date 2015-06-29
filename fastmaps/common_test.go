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

package fastmaps

import (
	"math/rand"
	"time"
)

var (
	seed = time.Now().UnixNano()
	r    = rand.New(rand.NewSource(seed))
)

const (
	size   = 100
	double = size * 2
)

func reset() {
	r = rand.New(rand.NewSource(seed))
}

type MapInt interface {
	benchkeys() [double]int
	Lookup(int) (int, bool)
}

type MapUint64 interface {
	benchkeys() [double]uint64
	Lookup(uint64) (int, bool)
}

func newMapInt() MapInt {
	reset()
	m := make(mapIntToInt)
	for i := 0; i < size; i++ {
		m[r.Int()] = r.Int()
	}
	return m
}

func (this mapIntToInt) benchkeys() [double]int {
	var ks [double]int
	i := 0
	for k := range this {
		ks[i] = k
		i++
		ks[i] = r.Int()
		i++
	}
	return ks
}

func newMapUint64() MapUint64 {
	reset()
	m := make(mapUint64ToInt)
	for i := 0; i < size; i++ {
		m[uint64(r.Uint32())] = r.Int()
	}
	return m
}

func (this mapUint64ToInt) benchkeys() [double]uint64 {
	var ks [double]uint64
	i := 0
	for k := range this {
		ks[i] = k
		i++
		ks[i] = uint64(r.Uint32())
		i++
	}
	return ks
}

func newSingleInt() MapInt {
	reset()
	return &singleIntToInt{r.Int(), r.Int()}
}

func (this *singleIntToInt) benchkeys() [double]int {
	var ks [double]int
	ks[0] = this.k
	for i := 1; i < double; i++ {
		ks[i] = r.Int()
	}
	return ks
}

type singleRefInt struct {
	k int
	v int
}

func newSingleRefInt() MapInt {
	reset()
	return singleRefInt{r.Int(), r.Int()}
}

func (this singleRefInt) benchkeys() [double]int {
	var ks [double]int
	ks[0] = this.k
	for i := 1; i < double; i++ {
		ks[i] = r.Int()
	}
	return ks
}

func (this singleRefInt) Lookup(k int) (int, bool) {
	if this.k == k {
		return this.v, true
	}
	return 0, false
}

func newUint64Value() MapUint64 {
	reset()
	return &singleUint64ToInt{uint64(r.Uint32()), r.Int()}
}

func (this *singleUint64ToInt) benchkeys() [double]uint64 {
	var ks [double]uint64
	ks[0] = this.k
	for i := 1; i < double; i++ {
		ks[i] = uint64(r.Uint32())
	}
	return ks
}

func newListInt(n int) MapInt {
	reset()
	l := &listIntToInt{make([]int, n), make([]int, n)}
	for i := 0; i < n; i++ {
		l.ks[i] = r.Int()
		l.vs[i] = r.Int()
	}
	return l
}

func (this *listIntToInt) benchkeys() [double]int {
	var ks [double]int
	for i, k := range this.ks {
		ks[i] = k
	}
	for i := len(this.ks); i < double; i++ {
		ks[i] = r.Int()
	}
	return ks
}

type arrayInt struct {
	ks [100]int
	vs [100]int
	n  int
}

func newArrayInt(n int) MapInt {
	reset()
	arr := &arrayInt{}
	for i := 0; i < n; i++ {
		arr.ks[i] = r.Int()
		arr.vs[i] = r.Int()
	}
	arr.n = n
	return arr
}

func (this *arrayInt) benchkeys() [double]int {
	var ks [double]int
	for i := 0; i < this.n; i++ {
		ks[i] = this.ks[i]
	}
	for i := this.n; i < double; i++ {
		ks[i] = r.Int()
	}
	return ks
}

func (this *arrayInt) Lookup(key int) (int, bool) {
	for i := 0; i < this.n; i++ {
		if this.ks[i] == key {
			return this.vs[i], true
		}
	}
	return 0, false
}

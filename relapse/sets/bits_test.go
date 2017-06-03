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
	"math/rand"
	"testing"
	"time"
)

func check(t *testing.T, this, that []bool) {
	if len(this) != len(that) {
		t.Fatalf("expected %#v, but got %#v", this, that)
	}
	for i := range this {
		if this[i] != that[i] {
			t.Fatalf("expected %#v, but got %#v", this, that)
		}
	}
	return
}

func TestBits1(t *testing.T) {
	bs := NewBits(1)
	bs.Set(0, true)
	check(t, bs.list(), []bool{true})
	if !bs.Get(0) {
		t.Fatalf("expected true")
	}
	bs.Set(0, false)
	check(t, bs.list(), []bool{false})
	if bs.Get(0) {
		t.Fatalf("expected false")
	}
}

func TestBits10(t *testing.T) {
	bs := NewBits(10)
	bs.Set(5, true)
	bs.Set(6, false)
	bs.Set(7, true)
	bs.Set(8, false)
	bs.Set(9, true)
	check(t, bs.list(), []bool{false, false, false, false, false, true, false, true, false, true})
	if !bs.Get(5) {
		t.Fatalf("expected true")
	}
	bs.Set(3, true)
	check(t, bs.list(), []bool{false, false, false, true, false, true, false, true, false, true})
	if !bs.Get(3) {
		t.Fatalf("expected true")
	}
}

func TestBits100(t *testing.T) {
	bs := NewBits(100)
	bs.Set(5, true)
	bs.Set(6, false)
	bs.Set(7, true)
	bs.Set(8, false)
	bs.Set(9, true)
	if !bs.Get(5) {
		t.Fatalf("expected true")
	}
	bs.Set(3, true)
	if !bs.Get(3) {
		t.Fatalf("expected true")
	}
	bs.Set(99, true)
	if bs.Get(98) {
		t.Fatalf("expected false")
	}
	if !bs.Get(99) {
		t.Fatalf("expected true")
	}
	bs.Set(64, true)
	if bs.Get(63) {
		t.Fatalf("expected false")
	}
	if !bs.Get(64) {
		t.Fatalf("expected true")
	}
	bs.Set(63, true)
	if bs.Get(62) {
		t.Fatalf("expected false")
	}
	if !bs.Get(63) {
		t.Fatalf("expected true")
	}
}

func TestBits1000(t *testing.T) {
	bs := NewBits(1000)
	bs.Set(3, true)
	if !bs.Get(3) {
		t.Fatalf("expected true")
	}
	bs.Set(99, true)
	if bs.Get(98) {
		t.Fatalf("expected false")
	}
	if !bs.Get(99) {
		t.Fatalf("expected true")
	}
	bs.Set(999, true)
	if bs.Get(998) {
		t.Fatalf("expected false")
	}
	if !bs.Get(999) {
		t.Fatalf("expected true")
	}
}

func TestBitsRandom(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := make([]int, r.Intn(20)+10)
	bools := make([]bool, len(ints))
	for runs := 0; runs < 10; runs++ {
		size := r.Intn(1000) + 100
		bs := NewBits(size)
		set := make(map[int]struct{}, len(ints))
		for i := range ints {
			index := r.Intn(size)
			_, ok := set[index]
			for ok {
				index = r.Intn(size)
				_, ok = set[index]
			}
			set[index] = struct{}{}
			ints[i] = index
			bools[i] = r.Intn(2) == 1
			bs.Set(ints[i], bools[i])
		}
		for i := range ints {
			if bools[i] != bs.Get(ints[i]) {
				t.Errorf("size := %d", size)
				t.Errorf("ints := %#v", ints)
				t.Errorf("bools := %#v", bools)
				t.Errorf("index := %d", ints[i])
				t.Fatalf("expected %v, but got %v", bools[i], bs.Get(ints[i]))
			}
		}
	}
}

func TestBitsInc10(t *testing.T) {
	size := 10
	max := NewBits(size)
	for i := 0; i < size; i++ {
		max.Set(i, true)
	}
	current := NewBits(size)
	for {
		current = current.Inc()
		if current.Equal(max) {
			break
		}
	}
}

func TestBitSetInc65(t *testing.T) {
	size := 65
	max := NewBits(size)
	max.Set(size-1, true)
	max.Set(0, true)
	current := NewBits(size)
	for i := 0; i < 64; i++ {
		current.Set(i, true)
	}
	for {
		current = current.Inc()
		if current.Equal(max) {
			break
		}
	}
}

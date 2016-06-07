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

func TestBitSet1(t *testing.T) {
	bs := newBitSet(1)
	bs.set(0, true)
	check(t, bs.list(), []bool{true})
	if !bs.get(0) {
		t.Fatalf("expected true")
	}
	bs.set(0, false)
	check(t, bs.list(), []bool{false})
	if bs.get(0) {
		t.Fatalf("expected false")
	}
}

func TestBitSet10(t *testing.T) {
	bs := newBitSet(10)
	bs.set(5, true)
	bs.set(6, false)
	bs.set(7, true)
	bs.set(8, false)
	bs.set(9, true)
	check(t, bs.list(), []bool{false, false, false, false, false, true, false, true, false, true})
	if !bs.get(5) {
		t.Fatalf("expected true")
	}
	bs.set(3, true)
	check(t, bs.list(), []bool{false, false, false, true, false, true, false, true, false, true})
	if !bs.get(3) {
		t.Fatalf("expected true")
	}
}

func TestBitSet100(t *testing.T) {
	bs := newBitSet(100)
	bs.set(5, true)
	bs.set(6, false)
	bs.set(7, true)
	bs.set(8, false)
	bs.set(9, true)
	if !bs.get(5) {
		t.Fatalf("expected true")
	}
	bs.set(3, true)
	if !bs.get(3) {
		t.Fatalf("expected true")
	}
	bs.set(99, true)
	if bs.get(98) {
		t.Fatalf("expected false")
	}
	if !bs.get(99) {
		t.Fatalf("expected true")
	}
	bs.set(64, true)
	if bs.get(63) {
		t.Fatalf("expected false")
	}
	if !bs.get(64) {
		t.Fatalf("expected true")
	}
	bs.set(63, true)
	if bs.get(62) {
		t.Fatalf("expected false")
	}
	if !bs.get(63) {
		t.Fatalf("expected true")
	}
}

func TestBitSet1000(t *testing.T) {
	bs := newBitSet(1000)
	bs.set(3, true)
	if !bs.get(3) {
		t.Fatalf("expected true")
	}
	bs.set(99, true)
	if bs.get(98) {
		t.Fatalf("expected false")
	}
	if !bs.get(99) {
		t.Fatalf("expected true")
	}
	bs.set(999, true)
	if bs.get(998) {
		t.Fatalf("expected false")
	}
	if !bs.get(999) {
		t.Fatalf("expected true")
	}
}

func TestBitSetRandom(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := make([]int, r.Intn(20)+10)
	bools := make([]bool, len(ints))
	for runs := 0; runs < 10; runs++ {
		size := r.Intn(1000) + 100
		bs := newBitSet(size)
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
			bs.set(ints[i], bools[i])
		}
		for i := range ints {
			if bools[i] != bs.get(ints[i]) {
				t.Errorf("size := %d", size)
				t.Errorf("ints := %#v", ints)
				t.Errorf("bools := %#v", bools)
				t.Errorf("index := %d", ints[i])
				t.Fatalf("expected %v, but got %v", bools[i], bs.get(ints[i]))
			}
		}
	}
}

func TestBitSetInc10(t *testing.T) {
	size := 10
	max := newBitSet(size)
	for i := 0; i < size; i++ {
		max.set(i, true)
	}
	current := newBitSet(size)
	for {
		current = current.inc()
		if current.equal(max) {
			break
		}
	}
}

func TestBitSetInc65(t *testing.T) {
	size := 65
	max := newBitSet(size)
	max.set(size-1, true)
	max.set(0, true)
	current := newBitSet(size)
	for i := 0; i < 64; i++ {
		current.set(i, true)
	}
	for {
		current = current.inc()
		if current.equal(max) {
			break
		}
	}
}

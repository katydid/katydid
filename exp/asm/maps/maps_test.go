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

import (
	"reflect"
	"testing"
)

func testIntToInt(t *testing.T, num int) {
	reset()
	m := make(map[int]int)
	for i := 0; i < num; i++ {
		m[r.Int()] = r.Int()
	}
	mapper := NewIntToInt(m)
	for k, v := range m {
		vv, ok := mapper.Lookup(k)
		if !ok {
			t.Fatalf("map should contain key = %v", k)
		}
		if vv != v {
			t.Fatalf("value should be %v got %v", v, vv)
		}
	}
	if !reflect.DeepEqual(m, mapper.ToMap()) {
		t.Fatalf("maps not equal")
	}
}

func TestZeroIntToInt(t *testing.T) {
	testIntToInt(t, 0)
}

func TestOneIntToInt(t *testing.T) {
	testIntToInt(t, 1)
}

func TestFiveIntToInt(t *testing.T) {
	testIntToInt(t, 5)
}

func TestTenIntToInt(t *testing.T) {
	testIntToInt(t, 10)
}

func TestTwentyIntToInt(t *testing.T) {
	testIntToInt(t, 20)
}

func TestFiftyIntToInt(t *testing.T) {
	testIntToInt(t, 50)
}

func TestHundredIntToInt(t *testing.T) {
	testIntToInt(t, 100)
}

func TestThousandIntToInt(t *testing.T) {
	testIntToInt(t, 1000)
}

func testUint64ToInt(t *testing.T, num int) {
	reset()
	m := make(map[uint64]int)
	for i := 0; i < num; i++ {
		m[uint64(r.Uint32())] = r.Int()
	}
	mapper := NewUint64ToInt(m)
	for k, v := range m {
		vv, ok := mapper.Lookup(k)
		if !ok {
			t.Fatalf("map should contain key = %v", k)
		}
		if vv != v {
			t.Fatalf("value should be %v got %v", v, vv)
		}
	}
	if !reflect.DeepEqual(m, mapper.ToMap()) {
		t.Fatalf("maps not equal")
	}
}

func TestZeroUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 0)
}

func TestOneUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 1)
}

func TestFiveUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 5)
}

func TestTenUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 10)
}

func TestTwentyUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 20)
}

func TestFiftyUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 50)
}

func TestHundredUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 100)
}

func TestThousandUint64ToInt(t *testing.T) {
	testUint64ToInt(t, 1000)
}

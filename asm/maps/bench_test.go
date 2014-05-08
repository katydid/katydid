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
	"testing"
)

func benchInt(b *testing.B, m MapInt) {
	keys := m.benchkeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Lookup(keys[i%double])
	}
}

func benchUint64(b *testing.B, m MapUint64) {
	keys := m.benchkeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Lookup(keys[i%double])
	}
}

func BenchmarkMapInt(b *testing.B) {
	benchInt(b, newMapInt())
}

func BenchmarkSingleInt(b *testing.B) {
	benchInt(b, newSingleInt())
}

func BenchmarkSingleRefInt(b *testing.B) {
	benchInt(b, newSingleRefInt())
}

func BenchmarkListOneInt(b *testing.B) {
	benchInt(b, newListInt(1))
}

func BenchmarkListTwoInt(b *testing.B) {
	benchInt(b, newListInt(2))
}

func BenchmarkListTenInt(b *testing.B) {
	benchInt(b, newListInt(10))
}

func BenchmarkListForteenInt(b *testing.B) {
	benchInt(b, newListInt(14))
}

func BenchmarkListFifteenInt(b *testing.B) {
	benchInt(b, newListInt(15))
}

func BenchmarkListTwentyInt(b *testing.B) {
	benchInt(b, newListInt(20))
}

func BenchmarkArrayOneInt(b *testing.B) {
	benchInt(b, newArrayInt(1))
}

func BenchmarkArrayTwoInt(b *testing.B) {
	benchInt(b, newArrayInt(2))
}

func BenchmarkArrayTenInt(b *testing.B) {
	benchInt(b, newArrayInt(10))
}

func BenchmarkArrayTwentyInt(b *testing.B) {
	benchInt(b, newArrayInt(20))
}

func BenchmarkMapUint64(b *testing.B) {
	benchUint64(b, newMapUint64())
}

func BenchmarkUint64Value(b *testing.B) {
	benchUint64(b, newUint64Value())
}

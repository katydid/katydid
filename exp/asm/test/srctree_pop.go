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

package main

func NewPopulatedSrcTree(r randySrctree, easy bool) *SrcTree {
	this := &SrcTree{}
	if r.Intn(10) != 0 {
		v1 := randStringSrctree(r)
		this.PackageName = &v1
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Imports = make([]*SrcTree, v2)
		for i := 0; i < v2; i++ {
			this.Imports[i] = NewPopulatedSrcTree(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSrctree(r, 3)
	}
	return this
}

type randySrctree interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSrctree(r randySrctree) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringSrctree(r randySrctree) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneSrctree(r)
	}
	return string(tmps)
}
func randUnrecognizedSrctree(r randySrctree, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldSrctree(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldSrctree(data []byte, r randySrctree, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateSrctree(data, uint64(key))
		data = encodeVarintPopulateSrctree(data, uint64(r.Int63()))
	case 1:
		data = encodeVarintPopulateSrctree(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateSrctree(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateSrctree(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateSrctree(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateSrctree(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}

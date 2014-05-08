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

package conv

import (
	"math"
	"reflect"
	"testing"
)

func TestConv(t *testing.T) {
	vs := []interface{}{
		[]float64{1.1, 0},
		[]float64{},
		[]float64{0},
		[]float64{1.455, 1.222, math.MaxFloat64, 10000, 98234098.1},
		[]float32{1.1, 0},
		[]float32{},
		[]float32{0},
		[]float32{1.455, 1.222, 0.4, math.MaxFloat32, 98234098.1},
		[]int64{},
		[]int64{0},
		[]int64{1, 2},
		[]int64{math.MaxInt64, math.MinInt64, -1, 10000},
		[]uint64{},
		[]uint64{0},
		[]uint64{1, 2},
		[]uint64{math.MaxUint64, 0, 9123098, 10000},
		[]int32{},
		[]int32{0},
		[]int32{1, 2},
		[]int32{math.MaxInt32, math.MinInt32, -1, 10000},
		[]bool{true, false},
		[]bool{},
		[]bool{false},
		[]bool{true, false, true, false, false},
		[]string{`alsdfk`},
		[]string{""},
		[]string{},
		[]string{`"adsfadf"`, "asdfasdf", "'a'", "a;ksldfhasl;dfhalkjhfdalkjsdfhalkjdfh"},
		[][]byte{[]byte(`alsdfk`)},
		[][]byte{{}},
		[][]byte{},
		[][]byte{[]byte(`"adsfadf"`), []byte("asdfasdf"), []byte("'a'"), []byte("a;ksldfhasl;dfhalkjhfdalkjsdfhalkjdfh")},
		[]uint32{},
		[]uint32{0},
		[]uint32{1, 2},
		[]uint32{math.MaxUint32, 0, 9123098, 10000},
		float64(0.0),
		float64(-1.111),
		float64(math.MaxFloat64),
		float32(0.0),
		float32(-1.111),
		float32(math.MaxFloat32),
		int64(1),
		int64(math.MaxInt64),
		int64(math.MinInt64),
		uint64(1),
		uint64(math.MaxUint64),
		uint64(0),
		int32(1),
		int32(math.MaxInt32),
		int32(math.MinInt32),
		bool(false),
		bool(true),
		string("adflkjasdf"),
		string(`"adsfadf"`),
		string("'a'"),
		string(""),
		[]byte("01289309"),
		[]byte{},
		[]byte("983-49583-589340-589345-098345-0934850-34985-43098543-095834"),
		uint32(1),
		uint32(math.MaxUint32),
		uint32(0),
	}
	for _, ins := range vs {
		sins, err := FromGo(ins)
		if err != nil {
			panic(err)
		}
		t.Logf("%s", sins)
		outs, err := ToGo(sins)
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(ins, outs) {
			t.Fatalf("expected %#v got %#v", ins, outs)
		}
	}
}

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

package test

import (
	"code.google.com/p/gogoprotobuf/proto"
	"math/rand"
	"testing"
	"time"
)

func (this test) bench(b *testing.B, newPop func(r randyTest, easy bool) proto.Message) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := 1000
	datas := make([][]byte, num)
	for i := 0; i < num; i++ {
		m := newPop(r, true)
		data, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := this.exec.Eval(datas[i%num]); err != nil {
			panic(err)
		}
	}
}

func BenchmarkContextPerson(b *testing.B) {
	newTest("`BenchContextPerson", b, contextPerson).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkRecursiveSrcTree(b *testing.B) {
	pops := []*SrcTree{ioUtil, path, runtime, syscall}
	newTest("TestRecursiveSrcTree", b, recursiveSrcTree).bench(b, func(r randyTest, easy bool) proto.Message {
		return pops[r.Intn(4)]
	})
}

func BenchmarkListIndexAddress(b *testing.B) {
	newTest("`BenchmarkListIndexAddress", b, listIndexAddress).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkNilName(b *testing.B) {
	newTest("`BenchmarkNilName", b, nilName).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkLenName(b *testing.B) {
	newTest("`BenchmarkLenName", b, lenName).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkEmptyOrNil(b *testing.B) {
	newTest("`BenchmarkEmptyOrNil", b, emptyOrNil).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkIncorrectNotName(b *testing.B) {
	newTest("`BenchmarkIncorrectNotName", b, incorrentNotName).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkCorrectNotName(b *testing.B) {
	newTest("`BenchmarkIncorrectNotName", b, correctNotName).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkAndNameTelephone(b *testing.B) {
	newTest("`BenchmarkIncorrectNotName", b, andNameTelephone).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

func BenchmarkOrNameTelephone(b *testing.B) {
	newTest("`BenchmarkIncorrectNotName", b, orNameTelephone).bench(b, func(r randyTest, easy bool) proto.Message {
		return NewPopulatedPerson(r, easy)
	})
}

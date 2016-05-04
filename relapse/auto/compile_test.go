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

package auto

import (
	"github.com/katydid/katydid/relapse/parser"
	"github.com/katydid/katydid/relapse/protonum"
	"github.com/katydid/katydid/relapse/tests"
	"testing"
)

func benchCompile(b *testing.B, str string) {
	st, err := parser.ParseGrammar(str)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compile(st)
	}
}

func BenchmarkCompileSmall(b *testing.B) {
	benchCompile(b, "*")
}

func BenchmarkCompileMedium(b *testing.B) {
	benchCompile(b, ".A.B.C: (A:* & B:* & C:* & D:*)")
}

func BenchmarkCompileLarge(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:*})")
}

func BenchmarkCompileLarger(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:* ; F.D:* ; G:* ; H:*; I:*})")
}

var typewriterQueryStr = `(.WineMessenger:* & .ShoelaceBeer:* & 
		.PocketRoses: ( 
			.ScarBusStop == "a" & 
			.BadgeShopping > 1 &
			.DaisySled < 5 &
			.SubmarineSaw == 0 &
			.SmileLetter :: $bool &
			.MenuPaperclip._ *= "A" &
			.BeetlePoker._ $= "b" &
			.WigPride._ ^= "c"
		)
	)`

func BenchmarkCompileProtoNum(b *testing.B) {
	st, err := parser.ParseGrammar(typewriterQueryStr)
	if err != nil {
		b.Fatal(err)
	}
	numst, err := protonum.FieldNamesToNumbers("tests", "TypewriterPrison", tests.TypewriterprisonDescription(), st)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compile(numst)
	}
}

func BenchmarkCompileProto(b *testing.B) {
	st, err := parser.ParseGrammar(typewriterQueryStr)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compile(st)
	}
}

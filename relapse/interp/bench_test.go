//  Copyright 2015 Walter Schulze
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

package interp_test

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/interp"
	pparser "github.com/katydid/katydid/serialize/proto/parser"
	"github.com/katydid/katydid/tests"
	"testing"
)

func bench(b *testing.B, patternDecls G, pkg string, msg string, gen func() proto.Message) {
	num := 1000
	datas := make([][]byte, num)
	for i := 0; i < num; i++ {
		m := gen()
		data, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	tmp := gen()
	desc := tmp.(tests.ProtoMessage).Description()
	g := patternDecls.Grammar()
	s := pparser.NewProtoParser(pkg, msg, desc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := s.Init(datas[i%num])
		if err != nil {
			b.Fatal(err)
		}
		interp.Interpret(g, s)
	}
}

func benchTest(t *testing.T, patternDecls G, expected bool, pkg string, msg string, m tests.ProtoMessage) {
	g := patternDecls.Grammar()
	desc := fmt.Sprintf("%#v", m)
	parser := tests.NewProtoParser(pkg, msg, m)
	match := interp.Interpret(g, parser)
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}

func BenchmarkContextPerson(b *testing.B) {
	bench(b, tests.ContextPerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkRecursiveSrcTree(b *testing.B) {
	bench(b, tests.RecursiveSrcTree, "tests", "SrcTree", tests.RandomSrcTree)
}

func BenchmarkListIndexAddress(b *testing.B) {
	bench(b, tests.ListIndexAddressPerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkNilName(b *testing.B) {
	bench(b, tests.NilNamePerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkLenName(b *testing.B) {
	bench(b, tests.LenNamePerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkEmptyOrNil(b *testing.B) {
	bench(b, tests.EmptyOrNilPerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkIncorrectNotName(b *testing.B) {
	bench(b, tests.NaiveNotNamePerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkCorrectNotName(b *testing.B) {
	bench(b, tests.ProperNotNamePerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkAndNameTelephone(b *testing.B) {
	bench(b, tests.AndNameTelephonePerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkOrNameTelephone(b *testing.B) {
	bench(b, tests.OrNameTelephonePerson, "tests", "Person", tests.RandomPerson)
}

var scarBusStop = G{
	"main": InOrder(
		Any(),
		In("PocketRoses",
			Any(),
			Field("ScarBusStop", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonScarBusStop(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{ScarBusStop: proto.String("a")}}
	benchTest(t, scarBusStop, true, "tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonScarBusStop(b *testing.B) {
	bench(b, scarBusStop, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var daisySled = G{
	"main": InOrder(
		Any(),
		In("PocketRoses",
			Any(),
			Field("DaisySled", Sprint(IntEq(IntVar(), IntConst(1)))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonDaisySled(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{DaisySled: proto.Int64(1)}}
	benchTest(t, daisySled, true, "tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonDaisySled(b *testing.B) {
	bench(b, daisySled, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var smileLetter = G{
	"main": InOrder(
		Any(),
		In("PocketRoses",
			Any(),
			Field("SmileLetter", Sprint(BoolVar())),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonSmileLetter(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{SmileLetter: proto.Bool(true)}}
	benchTest(t, smileLetter, true, "tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonSmileLetter(b *testing.B) {
	bench(b, smileLetter, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var menuPaperClip = G{
	"main": InOrder(
		Any(),
		In("PocketRoses",
			Any(),
			Field("MenuPaperclip", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonMenuPaperclip(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{MenuPaperclip: []string{"a"}}}
	benchTest(t, menuPaperClip, true, "tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonMenuPaperclip(b *testing.B) {
	bench(b, menuPaperClip, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var mapShark = G{
	"main": InOrder(
		Any(),
		In("PocketRoses",
			Any(),
			Field("MapShark", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonMapShark(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{MapShark: proto.String("a")}}
	benchTest(t, mapShark, true, "tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonMapShark(b *testing.B) {
	bench(b, mapShark, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var bridgePepper = G{
	"main": InOrder(
		Any(),
		In("FinanceJudo",
			Any(),
			In("SaladWorry",
				Any(),
				In("SpyCarpenter",
					Any(),
					Field("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
					Any(),
				),
				Any(),
			),
			Any(),
		),
		Any(),
	),
}

func TestBridgePepper(t *testing.T) {
	m := &tests.PuddingMilkshake{FinanceJudo: &tests.FinanceJudo{SaladWorry: &tests.SaladWorry{SpyCarpenter: &tests.SpyCarpenter{BridgePepper: []string{"a"}}}}}
	benchTest(t, bridgePepper, true, "tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepper(b *testing.B) {
	bench(b, bridgePepper, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

var bridgePepperAndFountainTarget = G{
	"main": InOrder(
		Any(),
		In("FinanceJudo",
			Any(),
			In("SaladWorry",
				Any(),
				In("SpyCarpenter",
					MatchAllOf(
						Field("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
						Field("FountainTarget", Sprint(Contains(StringVar(), StringConst("a")))),
					),
				),
				Any(),
			),
			Any(),
		),
		Any(),
	),
}

func TestBridgePepperAndFountainTarget(t *testing.T) {
	m := &tests.PuddingMilkshake{FinanceJudo: &tests.FinanceJudo{SaladWorry: &tests.SaladWorry{SpyCarpenter: &tests.SpyCarpenter{BridgePepper: []string{"a"}, FountainTarget: []string{"a"}}}}}
	benchTest(t, bridgePepperAndFountainTarget, true, "tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepperAndFountainTarget(b *testing.B) {
	bench(b, bridgePepperAndFountainTarget, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

var bridgePepperOrFountainTarget = G{
	"main": InOrder(
		Any(),
		In("FinanceJudo",
			Any(),
			In("SaladWorry",
				Any(),
				In("SpyCarpenter",
					MatchAnyOf(
						Field("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
						Field("FountainTarget", Sprint(Contains(StringVar(), StringConst("a")))),
					),
				),
				Any(),
			),
			Any(),
		),
		Any(),
	),
}

func TestBridgePepperOrFountainTarget(t *testing.T) {
	m := &tests.PuddingMilkshake{FinanceJudo: &tests.FinanceJudo{SaladWorry: &tests.SaladWorry{SpyCarpenter: &tests.SpyCarpenter{BridgePepper: []string{"b"}, FountainTarget: []string{"a"}}}}}
	benchTest(t, bridgePepperOrFountainTarget, true, "tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepperOrFountainTarget(b *testing.B) {
	bench(b, bridgePepperOrFountainTarget, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

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
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/interp"
	pscanner "github.com/katydid/katydid/serialize/proto/scanner"
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
	desc := tmp.(protoMessage).Description()
	g := patternDecls.Grammar()
	s := pscanner.NewProtoScanner(pkg, msg, desc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := s.Init(datas[i%num])
		if err != nil {
			b.Fatal(err)
		}
		interp.Interpret(g, s)
	}
}

func BenchmarkContextPerson(b *testing.B) {
	bench(b, contextPerson, "tests", "Person", tests.RandomPerson)
}

func BenchmarkRecursiveSrcTree(b *testing.B) {
	bench(b, recursiveSrcTree, "tests", "SrcTree", tests.RandomSrcTree)
}

func BenchmarkListIndexAddress(b *testing.B) {
	bench(b, listIndexAddress, "tests", "Person", tests.RandomPerson)
}

func BenchmarkNilName(b *testing.B) {
	bench(b, nilName, "tests", "Person", tests.RandomPerson)
}

func BenchmarkLenName(b *testing.B) {
	bench(b, lenName, "tests", "Person", tests.RandomPerson)
}

func BenchmarkEmptyOrNil(b *testing.B) {
	bench(b, emptyOrNil, "tests", "Person", tests.RandomPerson)
}

func BenchmarkIncorrectNotName(b *testing.B) {
	bench(b, incorrentNotName, "tests", "Person", tests.RandomPerson)
}

func BenchmarkCorrectNotName(b *testing.B) {
	bench(b, correctNotName, "tests", "Person", tests.RandomPerson)
}

func BenchmarkAndNameTelephone(b *testing.B) {
	bench(b, andNameTelephone, "tests", "Person", tests.RandomPerson)
}

func BenchmarkOrNameTelephone(b *testing.B) {
	bench(b, orNameTelephone, "tests", "Person", tests.RandomPerson)
}

var scarBusStop = G{
	"main": MatchTree(
		Any(),
		MatchIn("PocketRoses",
			Any(),
			MatchField("ScarBusStop", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonScarBusStop(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{ScarBusStop: proto.String("a")}}
	newTester(t, scarBusStop, true).proto("tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonScarBusStop(b *testing.B) {
	bench(b, scarBusStop, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var daisySled = G{
	"main": MatchTree(
		Any(),
		MatchIn("PocketRoses",
			Any(),
			MatchField("DaisySled", Sprint(IntEq(IntVar(), IntConst(1)))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonDaisySled(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{DaisySled: proto.Int64(1)}}
	newTester(t, daisySled, true).proto("tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonDaisySled(b *testing.B) {
	bench(b, daisySled, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var smileLetter = G{
	"main": MatchTree(
		Any(),
		MatchIn("PocketRoses",
			Any(),
			MatchField("SmileLetter", Sprint(BoolVar())),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonSmileLetter(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{SmileLetter: proto.Bool(true)}}
	newTester(t, smileLetter, true).proto("tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonSmileLetter(b *testing.B) {
	bench(b, smileLetter, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var menuPaperClip = G{
	"main": MatchTree(
		Any(),
		MatchIn("PocketRoses",
			Any(),
			MatchField("MenuPaperclip", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonMenuPaperclip(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{MenuPaperclip: []string{"a"}}}
	newTester(t, menuPaperClip, true).proto("tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonMenuPaperclip(b *testing.B) {
	bench(b, menuPaperClip, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var mapShark = G{
	"main": MatchTree(
		Any(),
		MatchIn("PocketRoses",
			Any(),
			MatchField("MapShark", Sprint(Contains(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func TestTypewriterPrisonMapShark(t *testing.T) {
	m := &tests.TypewriterPrison{PocketRoses: &tests.PocketRoses{MapShark: proto.String("a")}}
	newTester(t, mapShark, true).proto("tests", "TypewriterPrison", m)
}

func BenchmarkTypewriterPrisonMapShark(b *testing.B) {
	bench(b, mapShark, "tests", "TypewriterPrison", tests.RandomTypewriterPrison)
}

var bridgePepper = G{
	"main": MatchTree(
		Any(),
		MatchIn("FinanceJudo",
			Any(),
			MatchIn("SaladWorry",
				Any(),
				MatchIn("SpyCarpenter",
					Any(),
					MatchField("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
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
	newTester(t, bridgePepper, true).proto("tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepper(b *testing.B) {
	bench(b, bridgePepper, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

var bridgePepperAndFountainTarget = G{
	"main": MatchTree(
		Any(),
		MatchIn("FinanceJudo",
			Any(),
			MatchIn("SaladWorry",
				Any(),
				MatchIn("SpyCarpenter",
					AndMatch(
						MatchTree(
							Any(),
							MatchField("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
							Any(),
						),
						MatchTree(
							Any(),
							MatchField("FountainTarget", Sprint(Contains(StringVar(), StringConst("a")))),
							Any(),
						),
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
	newTester(t, bridgePepperAndFountainTarget, true).proto("tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepperAndFountainTarget(b *testing.B) {
	bench(b, bridgePepperAndFountainTarget, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

var bridgePepperOrFountainTarget = G{
	"main": MatchTree(
		Any(),
		MatchIn("FinanceJudo",
			Any(),
			MatchIn("SaladWorry",
				Any(),
				MatchIn("SpyCarpenter",
					OrMatch(
						MatchTree(
							Any(),
							MatchField("BridgePepper", Sprint(Contains(StringVar(), StringConst("a")))),
							Any(),
						),
						MatchTree(
							Any(),
							MatchField("FountainTarget", Sprint(Contains(StringVar(), StringConst("a")))),
							Any(),
						),
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
	newTester(t, bridgePepperOrFountainTarget, true).proto("tests", "PuddingMilkshake", m)
}

func BenchmarkBridgePepperOrFountainTarget(b *testing.B) {
	bench(b, bridgePepperOrFountainTarget, "tests", "PuddingMilkshake", tests.RandomPuddingMilkshake)
}

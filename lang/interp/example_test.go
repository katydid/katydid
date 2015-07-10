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
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/lang/combinator"
	"github.com/katydid/katydid/tests"
	"testing"
)

func example(t *testing.T, patternDecls G, m interface{}, expected bool) {
	newTester(t, patternDecls, expected).reflect(m)
}

//Has this person ever lived at 456 TheStreet
var contextPerson = G{"main": MatchTree(
	Any(),
	MatchIn("Addresses",
		MatchField("Number", Sprint(IntEq(IntVar(), IntConst(456)))),
		MatchField("Street", Sprint(StringEq(StringVar(), StringConst("TheStreet")))),
	),
	Any(),
)}

func TestContextDavid(t *testing.T) {
	example(t, contextPerson, tests.DavidPerson, false)
}

func TestContextRobert(t *testing.T) {
	example(t, contextPerson, tests.RobertPerson, true)
}

//Does this SrcTree depend on io or is its package name io
var recursiveSrcTree = G{
	"main": MatchTree(
		Any(),
		OrMatch(
			MatchField("PackageName", Sprint(StringEq(StringVar(), StringConst("io")))),
			MatchIn("Imports", Eval("main")),
		),
		Any(),
	),
}

func TestRecursiveSrcTreeIoUtil(t *testing.T) {
	example(t, recursiveSrcTree, tests.IoUtilSrcTree, true)
}

func TestRecursiveSrcTreePath(t *testing.T) {
	example(t, recursiveSrcTree, tests.PathSrcTree, true)
}

func TestRecursiveSrcTreeRuntime(t *testing.T) {
	example(t, recursiveSrcTree, tests.RuntimeSrcTree, false)
}

func TestRecursiveSrcTreeSyscall(t *testing.T) {
	example(t, recursiveSrcTree, tests.SyscallSrcTree, false)
}

//Is this Person's newest street number 1 and second newest street number 2.
//Assume that addresses are appended to the list, so the last address is the newest address.
//find main.Person where { main.Person { Addresses[-2].Number == 2 && Addresses[-1].Number == 1 } }
var listIndexAddress = G{
	"main": MatchTree(
		Any(),
		MatchIn("Addresses",
			Any(),
			MatchField("Number", Sprint(IntEq(IntVar(), IntConst(2)))),
			Any(),
		),
		MatchIn("Addresses",
			Any(),
			MatchField("Number", Sprint(IntEq(IntVar(), IntConst(1)))),
			Any(),
		),
		Many(MatchInAnyExcept("Addresses", Any())),
	),
}

func TestListIndexAddressMover(t *testing.T) {
	example(t, listIndexAddress, tests.MoverPerson, false)
}

func TestListIndexAddressShaker(t *testing.T) {
	example(t, listIndexAddress, tests.ShakerPerson, true)
}

func TestListIndexAddressRoutine(t *testing.T) {
	example(t, listIndexAddress, tests.RoutinePerson, false)
}

//Is this person's name missing
var nilName = G{"main": NotMatch(MatchTree(
	Any(),
	MatchIn("Name", Any()),
	Any(),
))}

func TestNilNameNoName(t *testing.T) {
	example(t, nilName, tests.NonamePerson, true)
}

func TestNilNameJohn(t *testing.T) {
	example(t, nilName, tests.JohnPerson, false)
}

func TestNilNameSmith(t *testing.T) {
	example(t, nilName, tests.SmithPerson, false)
}

//Is this person's name an empty string
var lenName = G{"main": MatchTree(
	Any(),
	MatchField("Name", Sprint(IntEq(LenString(StringVar()), IntConst(0)))),
	Any(),
)}

func TestLenNameNoName(t *testing.T) {
	example(t, lenName, tests.NonamePerson, false)
}

func TestLenNameJohn(t *testing.T) {
	example(t, lenName, tests.JohnPerson, false)
}

func TestLenNameSmith(t *testing.T) {
	example(t, lenName, tests.SmithPerson, true)
}

//Is this person's name empty or an empty string
var emptyOrNil = G{
	"main":  OrMatch(Eval("empty"), Eval("nil")),
	"empty": lenName["main"],
	"nil":   nilName["main"],
}

func TestEmptyOrNilNoName(t *testing.T) {
	example(t, emptyOrNil, tests.NonamePerson, true)
}

func TestEmptyOrNilJohn(t *testing.T) {
	example(t, emptyOrNil, tests.JohnPerson, false)
}

func TestEmptyOrNilSmith(t *testing.T) {
	example(t, emptyOrNil, tests.SmithPerson, true)
}

//Is the person's name not David
var incorrentNotName = G{
	"main": MatchTree(
		Any(),
		MatchField("Name", Sprint(Not(StringEq(StringVar(), StringConst("David"))))),
		Any(),
	),
}

func TestIncorrectNotNameNoName(t *testing.T) {
	example(t, incorrentNotName, tests.NonamePerson, false)
}

func TestIncorrectNotNameRobert(t *testing.T) {
	example(t, incorrentNotName, tests.RobertPerson, true)
}

func TestIncorrectNotNameSmith(t *testing.T) {
	example(t, incorrentNotName, tests.SmithPerson, true)
}

func TestIncorrectNotNameDavid(t *testing.T) {
	example(t, incorrentNotName, tests.DavidPerson, false)
}

//Is the person's name not David and make sure that the case where the name does not exist is also covered
var correctNotName = G{
	"main": OrMatch(Eval("name"), Eval("nil")),
	"nil":  nilName["main"],
	"name": MatchTree(
		Any(),
		MatchField("Name", Sprint(Not(StringEq(StringVar(), StringConst("David"))))),
		Any(),
	),
}

func TestCorrectNotNameNoName(t *testing.T) {
	example(t, correctNotName, tests.NonamePerson, true)
}

func TestCorrectNotNameRobert(t *testing.T) {
	example(t, correctNotName, tests.RobertPerson, true)
}

func TestCorrectNotNameSmith(t *testing.T) {
	example(t, correctNotName, tests.SmithPerson, true)
}

func TestCorrectNotNameDavid(t *testing.T) {
	example(t, correctNotName, tests.DavidPerson, false)
}

//Is this person's name David and telephone number 0123456789
var andNameTelephone = G{
	"main": MatchTree(
		AndMatch(
			MatchTree(
				Any(),
				MatchField("Name", Sprint(StringEq(StringVar(), StringConst("David")))),
				Any(),
			),
			MatchTree(
				Any(),
				MatchField("Telephone", Sprint(StringEq(StringVar(), StringConst("0123456789")))),
				Any(),
			),
		),
	),
}

func TestAndNameTelephoneDavid(t *testing.T) {
	example(t, andNameTelephone, tests.DavidPerson, true)
}

func TestAndNameTelephoneJohn(t *testing.T) {
	example(t, andNameTelephone, tests.JohnPerson, false)
}

func TestAndNameTelephoneMover(t *testing.T) {
	example(t, andNameTelephone, tests.MoverPerson, false)
}

func TestAndNameTelephoneSmith(t *testing.T) {
	example(t, andNameTelephone, tests.SmithPerson, false)
}

//Is this person's name David or telephone number 0123456789
var orNameTelephone = G{
	"main": MatchTree(
		OrMatch(
			MatchTree(
				Any(),
				MatchField("Name", Sprint(StringEq(StringVar(), StringConst("David")))),
				Any(),
			),
			MatchTree(
				Any(),
				MatchField("Telephone", Sprint(StringEq(StringVar(), StringConst("0123456789")))),
				Any(),
			),
		),
	),
}

func TestOrNameTelephoneDavid(t *testing.T) {
	example(t, orNameTelephone, tests.DavidPerson, true)
}

func TestOrNameTelephoneJohn(t *testing.T) {
	example(t, orNameTelephone, tests.JohnPerson, true)
}

func TestOrNameTelephoneMover(t *testing.T) {
	example(t, orNameTelephone, tests.MoverPerson, true)
}

func TestOrNameTelephoneSmith(t *testing.T) {
	example(t, orNameTelephone, tests.SmithPerson, false)
}

//Is this person's telephone number 0123456789 or 0127897897
var listOfTelephones = G{
	"main": MatchTree(
		Any(),
		MatchField("Telephone", Sprint(Or(StringEq(StringVar(), StringConst("0123456789")), StringEq(StringVar(), StringConst("0127897897"))))),
		Any(),
	),
}

func TestListOfTelephonesDavid(t *testing.T) {
	example(t, listOfTelephones, tests.DavidPerson, true)
}

func TestListOfTelephonesShaker(t *testing.T) {
	example(t, listOfTelephones, tests.ShakerPerson, true)
}

func TestListOfTelephonesRoutine(t *testing.T) {
	example(t, listOfTelephones, tests.RoutinePerson, false)
}

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

package tests

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
)

var DavidPerson = &Person{
	Name: proto.String("David"),
	Addresses: []*Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var RobertPerson = &Person{
	Name: proto.String("Robert"),
	Addresses: []*Address{
		{
			Number: proto.Int64(456),
			Street: proto.String("TheStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

//Has this person ever lived at 456 TheStreet
var ContextPerson = G{"main": InPath("Addresses", InAny(InOrder(
	In("Number", Value(IntVarEq(IntConst(456)))),
	In("Street", Value(StringVarEq(StringConst("TheStreet")))),
)))}

func init() {
	Validate("ContextDavid", ContextPerson, AllCodecs(DavidPerson), false)
	Validate("ContextRobert", ContextPerson, AllCodecs(RobertPerson), true)
}

var XmlContextPerson = G{"main": In("Person", InPath("Addresses",
	In("Number", Value(IntVarEq(IntConst(456)))),
	In("Street", Value(StringVarEq(StringConst("TheStreet")))),
	Any(),
))}

func init() {
	Validate("XmlContextDavid", XmlContextPerson, XML(DavidPerson), false)
	Validate("XmlContextRobert", XmlContextPerson, XML(RobertPerson), true)
}

var MoverPerson = &Person{
	Name: proto.String("Mover"),
	Addresses: []*Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var ShakerPerson = &Person{
	Name: proto.String("Shaker"),
	Addresses: []*Address{
		{
			Number: proto.Int64(55),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var RoutinePerson = &Person{
	Name: proto.String("Routine"),
	Addresses: []*Address{
		{
			Number: proto.Int64(3),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0124444444"),
}

var NonamePerson = &Person{
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var JohnPerson = &Person{
	Name: proto.String("John"),
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var SmithPerson = &Person{
	Name: proto.String(""),
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

//Is this Person's newest street number 1 and second newest street number 2.
//Assume that addresses are appended to the list, so the last address is the newest address.
//find tests.Person where { tests.Person { Addresses[-2].Number == 2 && Addresses[-1].Number == 1 } }
var ListIndexAddressPerson = G{
	"main": InOrder(
		Any(),
		In("Addresses",
			Any(),
			InAny(InPath("Number", Value(IntVarEq(IntConst(2))))),
			InAny(InPath("Number", Value(IntVarEq(IntConst(1))))),
		),
		Any(),
	),
}

func init() {
	Validate("ListIndexAddressMover", ListIndexAddressPerson, AllCodecs(MoverPerson), false)
	Validate("ListIndexAddressShaker", ListIndexAddressPerson, AllCodecs(ShakerPerson), true)
	Validate("ListIndexAddressRoutine", ListIndexAddressPerson, AllCodecs(RoutinePerson), false)
}

//Is this person's name missing
var NilNamePerson = G{"main": OppositeOf(
	InPath("Name", Any()),
)}

func init() {
	Validate("NilNameNoname", NilNamePerson, AllCodecs(NonamePerson), true)
	Validate("NilNameJohn", NilNamePerson, AllCodecs(JohnPerson), false)
	Validate("NilNameSmith", NilNamePerson, AllCodecs(SmithPerson), false)
}

//Is this person's name an empty string
var LenNamePerson = G{"main": InOrder(
	Any(),
	In("Name", Value(IntEq(LenString(StringVar()), IntConst(0)))),
	Any(),
)}

func init() {
	Validate("LenNameNoname", LenNamePerson, AllCodecs(NonamePerson), false)
	Validate("LenNameJohn", LenNamePerson, AllCodecs(JohnPerson), false)
	Validate("LenNameSmith", LenNamePerson, AllCodecs(SmithPerson), true)
}

//Is this person's name empty or an empty string
var EmptyOrNilPerson = G{
	"main":  AnyOf(Eval("empty"), Eval("nil")),
	"empty": LenNamePerson["main"],
	"nil":   NilNamePerson["main"],
}

func init() {
	Validate("EmptyOrNilNoname", EmptyOrNilPerson, AllCodecs(NonamePerson), true)
	Validate("EmptyOrNilJohn", EmptyOrNilPerson, AllCodecs(JohnPerson), false)
	Validate("EmptyOrNilSmith", EmptyOrNilPerson, AllCodecs(SmithPerson), true)
}

//Is the person's name not David
var NaiveNotNamePerson = G{
	"main": InOrder(
		Any(),
		In("Name", Value(Not(StringEq(StringVar(), StringConst("David"))))),
		Any(),
	),
}

func init() {
	Validate("NaiveNotNameNoname", NaiveNotNamePerson, AllCodecs(NonamePerson), false)
	Validate("NaiveNotNameRobert", NaiveNotNamePerson, AllCodecs(RobertPerson), true)
	Validate("NaiveNotNameSmith", NaiveNotNamePerson, AllCodecs(SmithPerson), true)
	Validate("NaiveNotNameDavid", NaiveNotNamePerson, AllCodecs(DavidPerson), false)
}

//Is the person's name not David and make sure that the case where the name does not exist is also covered
var ProperNotNamePerson = G{
	"main": AnyOf(Eval("name"), Eval("nil")),
	"nil":  NilNamePerson["main"],
	"name": InOrder(
		Any(),
		In("Name", Value(Not(StringEq(StringVar(), StringConst("David"))))),
		Any(),
	),
}

func init() {
	Validate("ProperNotNamePersonNoname", ProperNotNamePerson, AllCodecs(NonamePerson), true)
	Validate("ProperNotNamePersonRobert", ProperNotNamePerson, AllCodecs(RobertPerson), true)
	Validate("ProperNotNamePersonSmith", ProperNotNamePerson, AllCodecs(SmithPerson), true)
	Validate("ProperNotNamePersonRobert", ProperNotNamePerson, AllCodecs(DavidPerson), false)
}

//Is this person's name David and telephone number 0123456789
var AndNameTelephonePerson = G{
	"main": InOrder(
		AllOf(
			InOrder(
				Any(),
				In("Name", Value(StringEq(StringVar(), StringConst("David")))),
				Any(),
			),
			InOrder(
				Any(),
				In("Telephone", Value(StringEq(StringVar(), StringConst("0123456789")))),
				Any(),
			),
		),
	),
}

func init() {
	Validate("AndNameTelephoneDavid", AndNameTelephonePerson, AllCodecs(DavidPerson), true)
	Validate("AndNameTelephoneJohn", AndNameTelephonePerson, AllCodecs(JohnPerson), false)
	Validate("AndNameTelephoneMover", AndNameTelephonePerson, AllCodecs(MoverPerson), false)
	Validate("AndNameTelephoneSmith", AndNameTelephonePerson, AllCodecs(SmithPerson), false)
}

//Is this person's name David or telephone number 0123456789
var OrNameTelephonePerson = G{
	"main": AnyOf(
		InPath("Name", Value(StringEq(StringVar(), StringConst("David")))),
		InPath("Telephone", Value(StringEq(StringVar(), StringConst("0123456789")))),
	),
}

func init() {
	Validate("OrNameTelephoneDavid", OrNameTelephonePerson, AllCodecs(DavidPerson), true)
	Validate("OrNameTelephoneJohn", OrNameTelephonePerson, AllCodecs(JohnPerson), true)
	Validate("OrNameTelephoneMover", OrNameTelephonePerson, AllCodecs(MoverPerson), true)
	Validate("OrNameTelephoneSmith", OrNameTelephonePerson, AllCodecs(SmithPerson), false)
}

//Is this person's telephone number 0123456789 or 0127897897
var ListOfTelephonesPerson = G{
	"main": InOrder(
		Any(),
		In("Telephone", Value(Or(StringEq(StringVar(), StringConst("0123456789")), StringEq(StringVar(), StringConst("0127897897"))))),
		Any(),
	),
}

func init() {
	Validate("ListOfTelephonesDavid", ListOfTelephonesPerson, AllCodecs(DavidPerson), true)
	Validate("ListOfTelephonesShaker", ListOfTelephonesPerson, AllCodecs(ShakerPerson), true)
	Validate("ListOfTelephonesRoutine", ListOfTelephonesPerson, AllCodecs(RoutinePerson), false)
}

var LeftRecursion = G{
	"main": InOrder(
		AnyOf(
			Eval("main"),
			InOrder(
				Any(),
				In("Telephone", Value(StringEq(StringVar(), StringConst("0123456789")))),
			),
		),
	),
}

func init() {
	Validate("LeftRecursionDavid", LeftRecursion, AllCodecs(DavidPerson), true)
	Validate("LeftRecursionRobert", LeftRecursion, AllCodecs(RobertPerson), false)
}

var HiddenLeftRecursion = G{
	"main": InOrder(
		AnyOf(
			Eval("hidden"),
			InOrder(
				Any(),
				In("Telephone", Value(StringEq(StringVar(), StringConst("0123456789")))),
			),
		),
	),
	"hidden": Eval("main"),
}

func init() {
	Validate("HiddenLeftRecursionDavid", HiddenLeftRecursion, AllCodecs(DavidPerson), true)
	Validate("HiddenLeftRecursionRobert", HiddenLeftRecursion, AllCodecs(RobertPerson), false)
}

var NegativePerson = &Person{
	Addresses: []*Address{
		{
			Number: proto.Int64(-1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var PositiveNumber = G{
	"main": InOrder(
		In("Addresses",
			In("Number", Value(UintGE(UintVar(), UintConst(0)))),
			Any(),
		),
		Any(),
	),
}

func init() {
	Validate("DontErrorGivenWrongType", PositiveNumber, AllCodecs(NegativePerson), false)
}

var CorrectTypePerson = G{"main": InOrder(
	In("Name", Value(TypeString(StringVar()))),
	Any(),
)}

var WrongTypePerson = G{"main": InOrder(
	In("Name", Value(TypeInt(IntVar()))),
	Any(),
)}

func init() {
	Validate("CorrectTypeRobert", CorrectTypePerson, AllCodecs(RobertPerson), true)
	Validate("WrongTypeRobert", WrongTypePerson, AllCodecs(RobertPerson), false)
}

var InSetPerson = G{"main": InPath("Name", Value(ContainsString(StringVar(),
	StringsConst([]string{"The", "Robert", "Smith"}),
)))}

func init() {
	Validate("InSetPersonRobert", InSetPerson, AllCodecs(RobertPerson), true)
	Validate("InSetPersonDavid", InSetPerson, AllCodecs(DavidPerson), false)
}

var OptionalName = G{"main": InOrder(
	Maybe(In("Name", Any())),
	In("Addresses", Any()),
	In("Telephone", Value(StringEq(StringVar(), StringConst("0127897897")))),
)}

func init() {
	Validate("OptionalNameShakerPerson", OptionalName, AllCodecs(ShakerPerson), true)
	Validate("OptionalNameNonamePerson", OptionalName, AllCodecs(NonamePerson), true)
	Validate("OptionalNameSmithPerson", OptionalName, AllCodecs(SmithPerson), true)
	Validate("OptionalNameRoutinePerson", OptionalName, AllCodecs(RoutinePerson), false)
	Validate("OptionalNameJohnPerson", OptionalName, AllCodecs(JohnPerson), false)
}

var OptionalAddress = G{"main": InPath("Addresses",
	Maybe(InAny(In("Number", Any()), In("Street", Any()))),
	InAny(In("Number", Value(IntEq(IntVar(), IntConst(456)))), In("Street", Any())),
)}

func init() {
	Validate("OptionalAddressRobertPerson", OptionalAddress, AllCodecs(RobertPerson), true)
	Validate("OptionalAddressDavidPerson", OptionalAddress, AllCodecs(DavidPerson), true)
	Validate("OptionalAddressMoverPerson", OptionalAddress, AllCodecs(MoverPerson), false)
}

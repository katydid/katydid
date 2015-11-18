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
	ValidateProtoKeyEtc("ContextDavid", ContextPerson, DavidPerson, false)
	ValidateProtoKeyEtc("ContextRobert", ContextPerson, RobertPerson, true)
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
	ValidateProtoKeyEtc("ListIndexAddressMover", ListIndexAddressPerson, MoverPerson, false)
	ValidateProtoKeyEtc("ListIndexAddressShaker", ListIndexAddressPerson, ShakerPerson, true)
	ValidateProtoKeyEtc("ListIndexAddressRoutine", ListIndexAddressPerson, RoutinePerson, false)
}

//Is this person's name missing
var NilNamePerson = G{"main": OppositeOf(
	InPath("Name", Any()),
)}

func init() {
	ValidateProtoKeyEtc("NilNameNoname", NilNamePerson, NonamePerson, true)
	ValidateProtoKeyEtc("NilNameJohn", NilNamePerson, JohnPerson, false)
	ValidateProtoKeyEtc("NilNameSmith", NilNamePerson, SmithPerson, false)
}

//Is this person's name an empty string
var LenNamePerson = G{"main": InOrder(
	Any(),
	In("Name", Value(IntEq(LenString(StringVar()), IntConst(0)))),
	Any(),
)}

func init() {
	ValidateProtoKeyEtc("LenNameNoname", LenNamePerson, NonamePerson, false)
	ValidateProtoKeyEtc("LenNameJohn", LenNamePerson, JohnPerson, false)
	ValidateProtoKeyEtc("LenNameSmith", LenNamePerson, SmithPerson, true)
}

//Is this person's name empty or an empty string
var EmptyOrNilPerson = G{
	"main":  AnyOf(Eval("empty"), Eval("nil")),
	"empty": LenNamePerson["main"],
	"nil":   NilNamePerson["main"],
}

func init() {
	ValidateProtoKeyEtc("EmptyOrNilNoname", EmptyOrNilPerson, NonamePerson, true)
	ValidateProtoKeyEtc("EmptyOrNilJohn", EmptyOrNilPerson, JohnPerson, false)
	ValidateProtoKeyEtc("EmptyOrNilSmith", EmptyOrNilPerson, SmithPerson, true)
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
	ValidateProtoKeyEtc("NaiveNotNameNoname", NaiveNotNamePerson, NonamePerson, false)
	ValidateProtoKeyEtc("NaiveNotNameRobert", NaiveNotNamePerson, RobertPerson, true)
	ValidateProtoKeyEtc("NaiveNotNameSmith", NaiveNotNamePerson, SmithPerson, true)
	ValidateProtoKeyEtc("NaiveNotNameDavid", NaiveNotNamePerson, DavidPerson, false)
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
	ValidateProtoKeyEtc("ProperNotNamePersonNoname", ProperNotNamePerson, NonamePerson, true)
	ValidateProtoKeyEtc("ProperNotNamePersonRobert", ProperNotNamePerson, RobertPerson, true)
	ValidateProtoKeyEtc("ProperNotNamePersonSmith", ProperNotNamePerson, SmithPerson, true)
	ValidateProtoKeyEtc("ProperNotNamePersonRobert", ProperNotNamePerson, DavidPerson, false)
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
	ValidateProtoKeyEtc("AndNameTelephoneDavid", AndNameTelephonePerson, DavidPerson, true)
	ValidateProtoKeyEtc("AndNameTelephoneJohn", AndNameTelephonePerson, JohnPerson, false)
	ValidateProtoKeyEtc("AndNameTelephoneMover", AndNameTelephonePerson, MoverPerson, false)
	ValidateProtoKeyEtc("AndNameTelephoneSmith", AndNameTelephonePerson, SmithPerson, false)
}

//Is this person's name David or telephone number 0123456789
var OrNameTelephonePerson = G{
	"main": AnyOf(
		InPath("Name", Value(StringEq(StringVar(), StringConst("David")))),
		InPath("Telephone", Value(StringEq(StringVar(), StringConst("0123456789")))),
	),
}

func init() {
	ValidateProtoKeyEtc("OrNameTelephoneDavid", OrNameTelephonePerson, DavidPerson, true)
	ValidateProtoKeyEtc("OrNameTelephoneJohn", OrNameTelephonePerson, JohnPerson, true)
	ValidateProtoKeyEtc("OrNameTelephoneMover", OrNameTelephonePerson, MoverPerson, true)
	ValidateProtoKeyEtc("OrNameTelephoneSmith", OrNameTelephonePerson, SmithPerson, false)
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
	ValidateProtoKeyEtc("ListOfTelephonesDavid", ListOfTelephonesPerson, DavidPerson, true)
	ValidateProtoKeyEtc("ListOfTelephonesShaker", ListOfTelephonesPerson, ShakerPerson, true)
	ValidateProtoKeyEtc("ListOfTelephonesRoutine", ListOfTelephonesPerson, RoutinePerson, false)
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
	ValidateProtoKeyEtc("LeftRecursionDavid", LeftRecursion, DavidPerson, true)
	ValidateProtoKeyEtc("LeftRecursionRobert", LeftRecursion, RobertPerson, false)
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
	ValidateProtoKeyEtc("HiddenLeftRecursionDavid", HiddenLeftRecursion, DavidPerson, true)
	ValidateProtoKeyEtc("HiddenLeftRecursionRobert", HiddenLeftRecursion, RobertPerson, false)
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
			InAny(
				In("Number", Value(UintGE(UintVar(), UintConst(0)))),
				Any(),
			),
		),
		Any(),
	),
}

func init() {
	ValidateProtoKeyEtc("DontErrorGivenWrongType", PositiveNumber, NegativePerson, false)
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
	ValidateProtoKeyEtc("CorrectTypeRobert", CorrectTypePerson, RobertPerson, true)
	ValidateProtoKeyEtc("WrongTypeRobert", WrongTypePerson, RobertPerson, false)
}

var InSetPerson = G{"main": InPath("Name", Value(ContainsString(StringVar(),
	StringsConst([]string{"The", "Robert", "Smith"}),
)))}

func init() {
	ValidateProtoKeyEtc("InSetPersonRobert", InSetPerson, RobertPerson, true)
	ValidateProtoKeyEtc("InSetPersonDavid", InSetPerson, DavidPerson, false)
}

var OptionalName = G{"main": InOrder(
	Maybe(In("Name", Any())),
	In("Addresses", Any()),
	In("Telephone", Value(StringEq(StringVar(), StringConst("0127897897")))),
)}

func init() {
	ValidateProtoKeyEtc("OptionalNameShakerPerson", OptionalName, ShakerPerson, true)
	ValidateProtoKeyEtc("OptionalNameNonamePerson", OptionalName, NonamePerson, true)
	ValidateProtoKeyEtc("OptionalNameSmithPerson", OptionalName, SmithPerson, true)
	ValidateProtoKeyEtc("OptionalNameRoutinePerson", OptionalName, RoutinePerson, false)
	ValidateProtoKeyEtc("OptionalNameJohnPerson", OptionalName, JohnPerson, false)
}

var OptionalAddress = G{"main": InPath("Addresses",
	Maybe(InAny(In("Number", Any()), In("Street", Any()))),
	InAny(In("Number", Value(IntEq(IntVar(), IntConst(456)))), In("Street", Any())),
)}

func init() {
	ValidateProtoKeyEtc("OptionalAddressRobertPerson", OptionalAddress, RobertPerson, true)
	ValidateProtoKeyEtc("OptionalAddressDavidPerson", OptionalAddress, DavidPerson, true)
	ValidateProtoKeyEtc("OptionalAddressMoverPerson", OptionalAddress, MoverPerson, false)
}

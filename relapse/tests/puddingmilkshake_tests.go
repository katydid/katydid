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
	. "github.com/katydid/katydid/expr/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
)

var Ab21FinanceJudo = &FinanceJudo{
	SaladWorry: &SaladWorry{
		MagazineFrame: []string{"a", "b"},
		XrayPilot: &XrayPilot{
			AnkleCoat: proto.Int64(2),
		},
	},
	RumourSpirit: proto.Int64(1),
}

var AnyFinanceJudo = G{"main": Any()}

func init() {
	ValidateProtoNumEtc("Ab21Any", AnyFinanceJudo, Ab21FinanceJudo, true)
}

var NoneFinanceJudo = G{"main": None()}

func init() {
	ValidateProtoNumEtc("Ab21None", NoneFinanceJudo, Ab21FinanceJudo, false)
}

var HasSpirit1FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": In("RumourSpirit", Value(IntEq(IntVar(), IntConst(1)))),
}

func init() {
	ValidateProtoNumEtc("Ab21Spirit1", HasSpirit1FinanceJudo, Ab21FinanceJudo, true)
}

var HasSpirit2FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": In("RumourSpirit", Value(IntEq(IntVar(), IntConst(2)))),
}

func init() {
	ValidateProtoNumEtc("Ab21Spirit2", HasSpirit2FinanceJudo, Ab21FinanceJudo, false)
}

var MagazineFrameAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			In("MagazineFrame", InOrder(
				InAny(Value(StringEq(StringVar(), StringConst("a")))),
				Any()),
			),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoNumEtc("Ab21MagazineFrameA", MagazineFrameAFinanceJudo, Ab21FinanceJudo, true)
}

var MagazineFrameSingleAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			In("MagazineFrame", Value(StringEq(StringVar(), StringConst("a")))),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoNumEtc("Ab21MagazineFrameSingleA", MagazineFrameSingleAFinanceJudo, Ab21FinanceJudo, false)
}

var InAnyExceptNotAFieldNameFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("NotAFieldName",
			In("MagazineFrame",
				InOrder(
					InAny(Value(StringEq(StringVar(), StringConst("a")))),
					Any()),
			),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21NotAFieldName", InAnyExceptNotAFieldNameFinanceJudo, Ab21FinanceJudo, true)
}

var InAnyExceptSaladWorryFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("SaladWorry",
			In("MagazineFrame", Value(StringEq(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21InAnyExceptSaladWorry", InAnyExceptSaladWorryFinanceJudo, Ab21FinanceJudo, false)
}

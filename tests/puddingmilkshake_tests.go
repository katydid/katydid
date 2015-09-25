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
	Validate("Ab21Any", AnyFinanceJudo, AllCodecs(Ab21FinanceJudo), true)
}

var NoneFinanceJudo = G{"main": None()}

func init() {
	Validate("Ab21None", NoneFinanceJudo, AllCodecs(Ab21FinanceJudo), false)
}

var HasSpirit1FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": Field("RumourSpirit", Sprint(IntEq(IntVar(), IntConst(1)))),
}

func init() {
	Validate("Ab21Spirit1", HasSpirit1FinanceJudo, AllCodecs(Ab21FinanceJudo), true)
}

var HasSpirit2FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": Field("RumourSpirit", Sprint(IntEq(IntVar(), IntConst(2)))),
}

func init() {
	Validate("Ab21Spirit2", HasSpirit2FinanceJudo, AllCodecs(Ab21FinanceJudo), false)
}

var MagazineFrameAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			Field("MagazineFrame", Sprint(StringEq(StringVar(), StringConst("a")))),
			Any(),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	Validate("Ab21MagazineFrameA", MagazineFrameAFinanceJudo, AllCodecs(Ab21FinanceJudo), true)
}

var MagazineFrameSingleAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			Field("MagazineFrame", Sprint(StringEq(StringVar(), StringConst("a")))),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	Validate("Ab21MagazineFrameSingleA", MagazineFrameSingleAFinanceJudo, AllCodecs(Ab21FinanceJudo), false)
}

var InAnyExceptNotAFieldNameFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("NotAFieldName",
			Field("MagazineFrame", Sprint(StringEq(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func init() {
	Validate("Ab21NotAFieldName", InAnyExceptNotAFieldNameFinanceJudo, AllCodecs(Ab21FinanceJudo), true)
}

var InAnyExceptSaladWorryFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("SaladWorry",
			Field("MagazineFrame", Sprint(StringEq(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func init() {
	Validate("Ab21InAnyExceptSaladWorry", InAnyExceptSaladWorryFinanceJudo, AllCodecs(Ab21FinanceJudo), false)
}

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
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
)

var AbcPocketRoses = &PocketRoses{
	MenuPaperclip: []string{"a", "b", "c"},
}

var AStringArray = G{"main": In("MenuPaperclip", Elem(0, Value(StringEq(StringVar(), StringConst("a")))), Any())}
var FinalCStringArray = G{"main": In("MenuPaperclip", Any(), Elem(2, Value(StringEq(StringVar(), StringConst("c")))))}
var AbcStringArray = G{"main": In("MenuPaperclip",
	Elem(0, Value(StringEq(StringVar(), StringConst("a")))),
	Elem(1, Value(StringEq(StringVar(), StringConst("b")))),
	Elem(2, Value(StringEq(StringVar(), StringConst("c")))),
)}
var NextAStringArray = G{"main": In("MenuPaperclip",
	Elem(1, Value(StringEq(StringVar(), StringConst("a")))),
	Elem(0, Value(StringEq(StringVar(), StringConst("b")))),
	Elem(2, Value(StringEq(StringVar(), StringConst("c")))),
)}
var DStringArray = G{"main": In("MenuPaperclip", Elem(0, Value(StringEq(StringVar(), StringConst("d")))), Any())}

func init() {
	ValidateProtoKeyEtc("APocketRoses", AStringArray, AbcPocketRoses, true)
	ValidateProtoKeyEtc("FinalCPocketRoses", FinalCStringArray, AbcPocketRoses, true)
	ValidateProtoKeyEtc("AbcPocketRoses", AbcStringArray, AbcPocketRoses, true)
	ValidateProtoKeyEtc("NextAPocketRoses", NextAStringArray, AbcPocketRoses, false)
	ValidateProtoKeyEtc("DPocketRoses", DStringArray, AbcPocketRoses, false)
}

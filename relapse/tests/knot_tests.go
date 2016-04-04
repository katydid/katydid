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
	. "github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/funcs"
)

var KnotElbow = G{"main": InAny(InAny(In("Elbow", Value(BoolEq(BoolVar(), BoolConst(true))))))}

var KnotKnotElbow = &Knot{
	BitterEnd: &Knot{
		BitterEnd: &Knot{
			Elbow: proto.Bool(true),
		},
	},
}

var BightTurn = &Knot{
	Bight: []*BightKnot{
		&BightKnot{
			Turn: proto.Bool(true),
		},
	},
}

func init() {
	ValidateProtoEtc("KnotKnotElbow", KnotElbow, KnotKnotElbow, true)
	ValidateProtoEtc("BightTurn", KnotElbow, BightTurn, false)
}

var RecursiveTurn = G{"main": AnyOf(
	InAny(Eval("main")),
	In("Turn", Value(BoolEq(BoolVar(), BoolConst(true)))),
)}

func init() {
	ValidateProtoEtc("RecursiveTurnKnotKnotElbow", RecursiveTurn, KnotKnotElbow, false)
	ValidateProtoEtc("RecursiveTurnBightTurn", RecursiveTurn, BightTurn, true)
}

var RecursiveElbow = G{"main": AnyOf(
	InAny(Eval("main")),
	In("Elbow", Value(BoolEq(BoolVar(), BoolConst(true)))),
)}

func init() {
	ValidateProtoEtc("RecursiveElbowKnotKnotElbow", RecursiveElbow, KnotKnotElbow, true)
	ValidateProtoEtc("RecursiveElbowBightTurn", RecursiveElbow, BightTurn, false)
}

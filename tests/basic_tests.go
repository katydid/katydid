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

package tests

import (
	. "github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	. "github.com/katydid/katydid/relapse/combinator"
)

func init() {
	Validate(
		"BasicNone_A",
		G{"main": relapse.NewNot(relapse.NewZAny())},
		XMLString("<A/>"),
		false,
	)

	Validate(
		"BasicA_A",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())},
		XMLString("<A/>"),
		true,
	)
	Validate(
		"BasicA_B",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())},
		XMLString("<B/>"),
		false,
	)

	Validate(
		"BasicNotA_A",
		G{"main": relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))},
		XMLString("<A/>"),
		false,
	)
	Validate(
		"BasicNotA_B",
		G{"main": relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))},
		XMLString("<B/>"),
		true,
	)

	Validate(
		"BasicAB_AB",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))},
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicAB_BB",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))},
		XMLString("<B><B/></B>"),
		false,
	)

	Validate(
		"BasicALeafB_AB",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), Value(StringEq(StringConst("B"), StringVar())))},
		XMLString("<A>B</A>"),
		true,
	)
	Validate(
		"BasicALeafB_BB",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), Value(StringEq(StringConst("B"), StringVar())))},
		XMLString("<B>B</B>"),
		false,
	)

	Validate(
		"BasicConcatBC_BC",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		))},
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatBC_BB",
		G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		))},
		XMLString("<A><B/><B/></A>"),
		false,
	)
}

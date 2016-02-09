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
	basicNone := G{"main": relapse.NewNot(relapse.NewZAny())}
	Validate(
		"BasicNone_A",
		basicNone,
		XMLString("<A/>"),
		false,
	)

	basicA := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())}
	Validate(
		"BasicA_A",
		basicA,
		XMLString("<A/>"),
		true,
	)
	Validate(
		"BasicA_B",
		basicA,
		XMLString("<B/>"),
		false,
	)

	basicNotA := G{"main": relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))}
	Validate(
		"BasicNotA_A",
		basicNotA,
		XMLString("<A/>"),
		false,
	)
	Validate(
		"BasicNotA_B",
		basicNotA,
		XMLString("<B/>"),
		true,
	)

	basicAB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))}
	Validate(
		"BasicAB_AB",
		basicAB,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicAB_BB",
		basicAB,
		XMLString("<B><B/></B>"),
		false,
	)

	basicALeafB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), Value(StringEq(StringConst("B"), StringVar())))}
	Validate(
		"BasicALeafB_AB",
		basicALeafB,
		XMLString("<A>B</A>"),
		true,
	)
	Validate(
		"BasicALeafB_BB",
		basicALeafB,
		XMLString("<B>B</B>"),
		false,
	)

	basicConcatBC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicConcatBC_BC",
		basicConcatBC,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatBC_BB",
		basicConcatBC,
		XMLString("<A><B/><B/></A>"),
		false,
	)

	basicNotConcatBC := G{"main": relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	)))}
	Validate(
		"BasicNotConcatBC_BC",
		basicNotConcatBC,
		XMLString("<A><B/><C/></A>"),
		false,
	)
	Validate(
		"BasicNotConcatBC_BB",
		basicNotConcatBC,
		XMLString("<A><B/><B/></A>"),
		true,
	)

	basicAorB := G{"main": relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	)}
	Validate(
		"BasicAorB_A",
		basicAorB,
		XMLString("<A/>"),
		true,
	)

	Validate(
		"BasicAorB_C",
		basicAorB,
		XMLString("<C/>"),
		false,
	)

	basicTreeAAorBB := G{"main": relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicTreeAAorBB_AA",
		basicTreeAAorBB,
		XMLString("<A><A/></A>"),
		true,
	)
	Validate(
		"BasicTreeAAorBB_AB",
		basicTreeAAorBB,
		XMLString("<A><B/></A>"),
		false,
	)

	basicTreeBAorBB := G{"main": relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicTreeBAorBB_BA",
		basicTreeBAorBB,
		XMLString("<B><A/></B>"),
		true,
	)
	Validate(
		"BasicTreeBAorBB_AA",
		basicTreeBAorBB,
		XMLString("<A><A/></A>"),
		false,
	)

	basicTreeAOrOrC := G{"main": relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"),
			relapse.NewOr(
				relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
				relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			),
		),
		relapse.NewTreeNode(relapse.NewStringName("C"),
			relapse.NewOr(
				relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
				relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			),
		),
	)}
	Validate(
		"BasicTreeAOrOrC_AB",
		basicTreeAOrOrC,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicTreeAOrOrC_CA",
		basicTreeAOrOrC,
		XMLString("<C><A/></C>"),
		false,
	)

	basicConcatZAnyC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicConcatZAnyC_AC",
		basicConcatZAnyC,
		XMLString("<A><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatZAnyC_ABC",
		basicConcatZAnyC,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatZAnyC_BBBC",
		basicConcatZAnyC,
		XMLString("<A><B/><B/><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatZAnyC_BCCC",
		basicConcatZAnyC,
		XMLString("<A><B/><C/><C/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatZAnyC_BCBC",
		basicConcatZAnyC,
		XMLString("<A><B/><C/><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatZAnyC_AB",
		basicConcatZAnyC,
		XMLString("<A><B/></A>"),
		false,
	)
	//typical fundemental flaw
	Validate(
		"BasicConcatZAnyC_ACchildB_TypicalFundementalFlaw",
		basicConcatZAnyC,
		XMLString("<A><C>B</C></A>"),
		false,
	)

	basicZeroOrMoreB0 := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicZeroOrMoreB_0",
		basicZeroOrMoreB0,
		XMLString("<A></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreB_1",
		basicZeroOrMoreB0,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreB_3",
		basicZeroOrMoreB0,
		XMLString("<A><B/><B/><B/></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreB_C",
		basicZeroOrMoreB0,
		XMLString("<A><C/></A>"),
		false,
	)
	Validate(
		"BasicZeroOrMoreB_BC",
		basicZeroOrMoreB0,
		XMLString("<A><B/><C/><B/></A>"),
		false,
	)

	basicZeroOrMoreEmpty := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewEmpty(),
	))}
	Validate(
		"BasicZeroOrMoreEmpty_Empty",
		basicZeroOrMoreEmpty,
		XMLString("<A></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreEmpty_B",
		basicZeroOrMoreEmpty,
		XMLString("<A><B/></A>"),
		false,
	)

	basicZeroOrMoreZeroOrMoreB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewZeroOrMore(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))}
	Validate(
		"BasicZeroOrMoreZeroOrMoreB_BB",
		basicZeroOrMoreZeroOrMoreB,
		XMLString("<A><B/><B/></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreZeroOrMoreB_C",
		basicZeroOrMoreZeroOrMoreB,
		XMLString("<A><C/></A>"),
		false,
	)

	basicConcatOrEmpty := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewOr(
			relapse.NewEmpty(),
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicConcatOrEmpty_BC",
		basicConcatOrEmpty,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatOrEmpty_C",
		basicConcatOrEmpty,
		XMLString("<A><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatOrEmpty_BD",
		basicConcatOrEmpty,
		XMLString("<A><B/><D/></A>"),
		false,
	)

	basicZeroOrMoreBOrEmpty := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewOr(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewEmpty(),
		),
	))}
	Validate(
		"BasicZeroOrMoreBOrEmpty_BB",
		basicZeroOrMoreBOrEmpty,
		XMLString("<A><B/><B/></A>"),
		true,
	)
	Validate(
		"BasicZeroOrMoreBOrEmpty_BC",
		basicZeroOrMoreBOrEmpty,
		XMLString("<A><B/><C/></A>"),
		false,
	)

	basicConcatCStar := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))}
	Validate(
		"BasicConcatCStar_0",
		basicConcatCStar,
		XMLString("<A></A>"),
		false,
	)
	Validate(
		"BasicConcatCStar_1",
		basicConcatCStar,
		XMLString("<A><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatCStar_2",
		basicConcatCStar,
		XMLString("<A><C/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatCStar_3",
		basicConcatCStar,
		XMLString("<A><C/><C/><C/></A>"),
		true,
	)
	Validate(
		"BasicConcatCStar_B",
		basicConcatCStar,
		XMLString("<A><B/></A>"),
		false,
	)
	Validate(
		"BasicConcatCStar_CB",
		basicConcatCStar,
		XMLString("<A><C/><B/></A>"),
		false,
	)

}

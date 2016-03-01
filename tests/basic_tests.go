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

	basicTreeAandA := G{"main": relapse.NewAnd(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicTreeAandA_A",
		basicTreeAandA,
		XMLString("<A><A/></A>"),
		true,
	)
	Validate(
		"BasicTreeAandA_B",
		basicTreeAandA,
		XMLString("<A><B/></A>"),
		false,
	)

	basicTreeAandB := G{"main": relapse.NewAnd(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicTreeAandB_B",
		basicTreeAandB,
		XMLString("<A><B/></A>"),
		false,
	)

	basicAndBAnyC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))}
	Validate(
		"BasicAndBAnyC_BC",
		basicAndBAnyC,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicAndBAnyC_CB",
		basicAndBAnyC,
		XMLString("<A><C/><B/></A>"),
		false,
	)
	Validate(
		"BasicAndBAnyC_B",
		basicAndBAnyC,
		XMLString("<A><B/></A>"),
		false,
	)
	Validate(
		"BasicAndBAnyC_C",
		basicAndBAnyC,
		XMLString("<A><C/></A>"),
		false,
	)
	Validate(
		"BasicAndBAnyC_BXXXC",
		basicAndBAnyC,
		XMLString("<A><B/><X/><X/><X/><C/></A>"),
		true,
	)

	basicTreeAndBAnyC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("A"),
				relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("A"),
				relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
			),
		),
	))}
	Validate(
		"BasicTreeAndBAnyC_BC",
		basicTreeAndBAnyC,
		XMLString("<A><A><B/></A><A><C/></A></A>"),
		true,
	)
	Validate(
		"BasicTreeAndBAnyC_CB",
		basicTreeAndBAnyC,
		XMLString("<A><A><C/></A><A><B/></A></A>"),
		false,
	)
	Validate(
		"BasicTreeAndBAnyC_BXXXC",
		basicTreeAndBAnyC,
		XMLString("<A><A><B/></A><X/><X/><X/><A><C/></A></A>"),
		true,
	)
	Validate(
		"BasicTreeAndBAnyC_CBC",
		basicTreeAndBAnyC,
		XMLString("<A><A><C/></A><A><B/></A><A><C/></A></A>"),
		false,
	)

	basicAContainsB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewContains(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicAContainsB_B",
		basicAContainsB,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicAContainsB_CBC",
		basicAContainsB,
		XMLString("<A><C/><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicAContainsB_CC",
		basicAContainsB,
		XMLString("<A><C/><C/></A>"),
		false,
	)
	Validate(
		"BasicAContainsB_0",
		basicAContainsB,
		XMLString("<A></A>"),
		false,
	)

	basicOptionalB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOptional(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicOptionalB_Empty",
		basicOptionalB,
		XMLString("<A></A>"),
		true,
	)
	Validate(
		"BasicOptionalB_B",
		basicOptionalB,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicOptionalB_BB",
		basicOptionalB,
		XMLString("<A><B/><B/></A>"),
		false,
	)
	Validate(
		"BasicOptionalB_C",
		basicOptionalB,
		XMLString("<A><C/></A>"),
		false,
	)

	basicInterleaveBC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))}
	Validate(
		"BasicInterleaveBC_BC",
		basicInterleaveBC,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBC_CB",
		basicInterleaveBC,
		XMLString("<A><C/><B/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBC_C",
		basicInterleaveBC,
		XMLString("<A><C/></A>"),
		false,
	)

	basicInterleaveBAnyC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))}
	Validate(
		"BasicInterleaveBAnyC_BC",
		basicInterleaveBAnyC,
		XMLString("<A><B/><C/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBAnyC_BAC",
		basicInterleaveBAnyC,
		XMLString("<A><B/><A/><C/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBAnyC_ABACA",
		basicInterleaveBAnyC,
		XMLString("<A><A/><B/><A/><C/><A/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBAnyC_ABAA",
		basicInterleaveBAnyC,
		XMLString("<A><A/><B/><A/><A/></A>"),
		false,
	)
	Validate(
		"BasicInterleaveBAnyC_ACCBA",
		basicInterleaveBAnyC,
		XMLString("<A><A/><C/><C/><B/><A/></A>"),
		true,
	)
	Validate(
		"BasicInterleaveBAnyC_ACCCA",
		basicInterleaveBAnyC,
		XMLString("<A><A/><C/><C/><C/><A/></A>"),
		false,
	)

	basicRefLoop := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOr(
		relapse.NewEmpty(),
		relapse.NewReference("main"),
	))}
	Validate(
		"BasicRefLoop_A",
		basicRefLoop,
		XMLString("<A/>"),
		true,
	)
	Validate(
		"BasicRefLoop_AA",
		basicRefLoop,
		XMLString("<A><A/></A>"),
		true,
	)
	Validate(
		"BasicRefLoop_AB",
		basicRefLoop,
		XMLString("<A><B/></A>"),
		false,
	)

	basicConcatBOptionalD := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewOptional(relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty())),
		),
	)}
	Validate(
		"BasicConcatBOptionalD_BD",
		basicConcatBOptionalD,
		XMLString("<A><B/><D/></A>"),
		true,
	)
	Validate(
		"BasicConcatBOptionalD_B",
		basicConcatBOptionalD,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicConcatBOptionalD_D",
		basicConcatBOptionalD,
		XMLString("<A><D/></A>"),
		false,
	)

	//!(B:<empty>) can be <empty> and * can be B:<empty> which means B:<empty> can be accepted
	basicAnyNotB := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))}
	Validate(
		"BasicAnyNotB_B",
		basicAnyNotB,
		XMLString("<A><B/></A>"),
		true,
	)
	Validate(
		"BasicAnyNotB_C",
		basicAnyNotB,
		XMLString("<A><C/></A>"),
		true,
	)

	basicNotAndBStarC := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewNot(
		relapse.NewAnd(
			relapse.NewConcat(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()), relapse.NewZAny()),
			relapse.NewConcat(relapse.NewZAny(), relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty())),
		),
	))}
	Validate(
		"BasicNotAndBStarC_BC",
		basicNotAndBStarC,
		XMLString("<A><B/><C/></A>"),
		false,
	)
	Validate(
		"BasicNotAndBStarC_CB",
		basicNotAndBStarC,
		XMLString("<A><C/><B/></A>"),
		true,
	)

	basicAndNotAB := G{"main": relapse.NewAnd(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicAndNotAB_A",
		basicAndNotAB,
		XMLString("<A/>"),
		false,
	)
	Validate(
		"BasicAndNotAB_B",
		basicAndNotAB,
		XMLString("<B/>"),
		false,
	)
	Validate(
		"BasicAndNotAB_C",
		basicAndNotAB,
		XMLString("<C/>"),
		true,
	)

	basicOrNotAB := G{"main": relapse.NewOr(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)}
	Validate(
		"BasicOrNotAB_A",
		basicOrNotAB,
		XMLString("<A/>"),
		true,
	)
	Validate(
		"BasicOrNotAB_C",
		basicOrNotAB,
		XMLString("<C/>"),
		true,
	)

	//deeper fundamental flaw
	basicAEndsWithBContainsAnyD := G{"main": relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("B"),
			relapse.NewContains(relapse.NewTreeNode(relapse.NewAnyName(), relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty()))),
		)),
	)}
	Validate(
		"BasicAEndsWithBContainsAnyD_BCD_DeeperFundementalFlaw",
		basicAEndsWithBContainsAnyD,
		XMLString("<A><B><C><D/></C></B></A>"),
		true,
	)
	Validate(
		"BasicAEndsWithBContainsAnyD_BCA_DeeperFundementalFlaw",
		basicAEndsWithBContainsAnyD,
		XMLString("<A><B><C><A/></C></B></A>"),
		false,
	)

	basicAndContainsTree := G{"main": In("A", AllOf(
		InPath("B", In("C", relapse.NewEmpty())),
		InPath("B", In("D", relapse.NewEmpty())),
	))}
	Validate(
		"BasicAndContainsTree_BCBD",
		basicAndContainsTree,
		XMLString("<A><B><C/></B><B><D/></B></A>"),
		true,
	)
	Validate(
		"BasicAndContainsTree_BC",
		basicAndContainsTree,
		XMLString("<A><B><C/></B></A>"),
		false,
	)
}

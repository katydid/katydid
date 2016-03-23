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
	. "github.com/katydid/katydid/expr/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
)

// Foundations of XML Processing: The Tree Automata Approach - Example 8.1.1
// Without simplification rules the state space for the respective automata can be become very large
var Example811 = G{
	"main": AnyOf(Eval("q1"), Eval("q2")),
	"q1": AnyOf(
		In("A",
			In("Left", Eval("q1")),
			In("Right", Eval("q2")),
		),
		In("A",
			In("Left", Eval("q1")),
			In("Right", Eval("q1")),
		),
	),
	"q2": AnyOf(
		In("A",
			In("Left", Eval("q2")),
			In("Right", Eval("q2")),
		),
		In("A",
			In("Value", Value(StringEq(StringVar(), StringConst("#")))),
		),
	),
}

type M map[string]interface{}

var TwoHashes811 = M{
	"A": M{
		"Left": M{
			"A": M{
				"Value": "#",
			},
		},
		"Right": M{
			"A": M{
				"Value": "#",
			},
		},
	},
}

func init() {
	Validate("TwoHashes811", Example811, Json(TwoHashes811), true)
}

var OneLeftAndTwoRightHashes811 = M{
	"A": M{
		"Left": M{
			"A": M{
				"Value": "#",
			},
		},
		"Right": M{
			"A": M{
				"Left": M{
					"A": M{
						"Value": "#",
					},
				},
				"Right": M{
					"A": M{
						"Value": "#",
					},
				},
			},
		},
	},
}

func init() {
	Validate("OneLeftAndTwoRightHashes811", Example811, Json(OneLeftAndTwoRightHashes811), true)
}

var TwoLeftAndOneRightHashes811 = M{
	"A": M{
		"Left": M{
			"A": M{
				"Left": M{
					"A": M{
						"Value": "#",
					},
				},
				"Right": M{
					"A": M{
						"Value": "#",
					},
				},
			},
		},
		"Right": M{
			"A": M{
				"Value": "#",
			},
		},
	},
}

func init() {
	Validate("TwoLeftAndOneRightHashes811", Example811, Json(TwoLeftAndOneRightHashes811), true)
}

var DeepLeft811 = M{
	"A": M{
		"Left": M{
			"A": M{
				"Left": M{
					"A": M{
						"Left": M{
							"A": M{
								"Value": "#",
							},
						},
						"Right": M{
							"A": M{
								"Value": "#",
							},
						},
					},
				},
				"Right": M{
					"A": M{
						"Value": "#",
					},
				},
			},
		},
		"Right": M{
			"A": M{
				"Value": "#",
			},
		},
	},
}

func init() {
	Validate("DeepLeft811", Example811, Json(DeepLeft811), true)
}

var OneHash811 = M{
	"A": M{
		"Right": M{
			"A": M{
				"Value": "#",
			},
		},
	},
}

func init() {
	Validate("OneHash811", Example811, Json(OneHash811), false)
}

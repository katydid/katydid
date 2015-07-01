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

package interp_test

import (
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/lang/combinator"
	"github.com/katydid/katydid/lang/interp"
	"github.com/katydid/katydid/serialize"
	jscanner "github.com/katydid/katydid/serialize/json"
	pscanner "github.com/katydid/katydid/serialize/proto/scanner"
	rscanner "github.com/katydid/katydid/serialize/reflect"
	"github.com/katydid/katydid/tests"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type tree struct {
	scannerName string
	scanner     serialize.Scanner
	description string
}

func newJsonScanner(m interface{}) serialize.Scanner {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := jscanner.NewJsonScanner()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	return s
}

func newJsonTree(m interface{}) tree {
	return tree{
		scannerName: "json",
		scanner:     newJsonScanner(m),
		description: fmt.Sprintf("%#v", m),
	}
}

func newReflectScanner(m interface{}) serialize.Scanner {
	s := rscanner.NewReflectScanner()
	s.Init(reflect.ValueOf(m))
	return s
}

func newReflectTree(m interface{}) tree {
	return tree{
		scannerName: "reflect",
		scanner:     newReflectScanner(m),
		description: fmt.Sprintf("%#v", m),
	}
}

func newProtoScanner(pkg, msg string, m protoMessage) serialize.Scanner {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := pscanner.NewProtoScanner(pkg, msg, m.Description())
	if err := s.Init(data); err != nil {
		panic(err)
	}
	return s
}

type protoMessage interface {
	Description() (desc *descriptor.FileDescriptorSet)
	proto.Message
}

func newProtoTree(pkg, msg string, m protoMessage) tree {
	return tree{
		scannerName: "proto",
		scanner:     newProtoScanner(pkg, msg, m),
		description: fmt.Sprintf("%#v", m),
	}
}

type tester struct {
	t            *testing.T
	patternDecls G
	expected     bool
}

func newTester(t *testing.T, patternDecls G, expected bool) tester {
	return tester{t, patternDecls, expected}
}

func (t tester) test(tree tree) tester {
	g := t.patternDecls.Grammar()
	match := interp.Interpret(g, tree.scanner)
	if match != t.expected {
		t.t.Fatalf("Expected a %v match given %s scanner from \n%v \non \n%s", t.expected, tree.scannerName, g.String(), tree.description)
	}
	return t
}

func (t tester) json(m interface{}) tester {
	return t.test(newJsonTree(m))
}

func (t tester) reflect(m interface{}) tester {
	return t.test(newReflectTree(m))
}

func (t tester) proto(pkg, msg string, m protoMessage) tester {
	return t.test(newProtoTree(pkg, msg, m))
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestInterp(t *testing.T) {
	m := &tests.FinanceJudo{
		SaladWorry: &tests.SaladWorry{
			MagazineFrame: []string{"a", "b"},
			XrayPilot: &tests.XrayPilot{
				AnkleCoat: proto.Int64(2),
			},
		},
		RumourSpirit: proto.Int64(1),
	}
	eq1 := Sprint(IntEq(IntVar(), IntConst(1)))
	t.Logf("%v", eq1)
	eq2 := Sprint(IntEq(IntVar(), IntConst(2)))
	t.Logf("%v", eq2)
	eqa := Sprint(StringEq(StringVar(), StringConst("a")))
	t.Logf("%v", eqa)
	testers := []tester{
		newTester(t, G{"main": Any()}, true),
		newTester(t, G{"main": None()}, false),
		newTester(t, G{
			"main":   MatchTree(Any(), Eval("spirit"), Any()),
			"spirit": MatchField("RumourSpirit", eq1),
		}, true),
		newTester(t, G{
			"main":   MatchTree(Any(), Eval("spirit"), Any()),
			"spirit": MatchField("RumourSpirit", eq2),
		}, false),
		newTester(t, G{
			"main": MatchTree(
				MatchIn("SaladWorry",
					MatchField("MagazineFrame", eqa),
					Any(),
					MatchIn("XrayPilot", Any()),
					Any(),
				),
				Any(),
			),
		}, true),
		newTester(t, G{
			"main": MatchTree(
				MatchIn("SaladWorry",
					MatchField("MagazineFrame", eqa),
					MatchIn("XrayPilot", Any()),
					Any(),
				),
				Any(),
			),
		}, false),
		newTester(t, G{
			"main": MatchTree(
				MatchInAnyExcept("NotAFieldName",
					MatchField("MagazineFrame", eqa),
					Any(),
				),
				Any(),
			),
		}, true),
		newTester(t, G{
			"main": MatchTree(
				MatchInAnyExcept("SaladWorry",
					MatchField("MagazineFrame", eqa),
					Any(),
				),
				Any(),
			),
		}, false),
	}
	for _, testy := range testers {
		testy.json(m).reflect(m).proto("tests", "FinanceJudo", m)
	}
}

type M map[string]interface{}

func Test811(t *testing.T) {
	// Foundations of XML Processing: The Tree Automata Approach - Example 8.1.1
	// Without simplification rules the state space for the respective automata can be become very large
	eqhash := Sprint(StringEq(StringVar(), StringConst("#")))
	ex811 := G{
		"main": OrMatch(Eval("q1"), Eval("q2")),
		"q1": OrMatch(
			MatchIn("A",
				MatchIn("Left", Eval("q1")),
				MatchIn("Right", Eval("q2")),
			),
			MatchIn("A",
				MatchIn("Left", Eval("q1")),
				MatchIn("Right", Eval("q1")),
			),
		),
		"q2": OrMatch(
			MatchIn("A",
				MatchIn("Left", Eval("q2")),
				MatchIn("Right", Eval("q2")),
			),
			MatchIn("A",
				MatchField("Value", eqhash),
			),
		),
	}
	newTester(t, ex811, true).json(M{
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
	})
	newTester(t, ex811, true).json(M{
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
	})
	newTester(t, ex811, true).json(M{
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
	})
	newTester(t, ex811, true).json(M{
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
	})
	newTester(t, ex811, false).json(M{
		"A": M{
			"Right": M{
				"A": M{
					"Value": "#",
				},
			},
		},
	})
}

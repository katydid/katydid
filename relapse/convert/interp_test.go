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

package convert

import (
	//"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/serialize"
	xparser "github.com/katydid/katydid/serialize/xml"
	"testing"
)

func TestInterpA1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A/>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<B/>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotA(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A/>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotB(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<B/>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpBB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<B><B/></B>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpALeafB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), combinator.Value(funcs.StringEq(funcs.StringConst("B"), funcs.StringVar())))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A>B</A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpBLeafB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), combinator.Value(funcs.StringEq(funcs.StringConst("B"), funcs.StringVar())))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<B>B</B>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatBB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><B/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotConcatBC(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	)))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotConcatBB(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	)))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A/>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpCorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<C/>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeAorA(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><A/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeAorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeBorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<B><A/></B>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeBorA(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><A/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeAororC(t *testing.T) {
	p := relapse.NewOr(
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
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeCororA(t *testing.T) {
	p := relapse.NewOr(
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
	)
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<C><A/></C>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatZAnyC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBBBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><B/><B/><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBCCC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/><C/><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/><B/><C/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

//current union cross product is wrong
func TestInterpConcatZAnyCChild(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><C>B</C></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreB0(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreB3(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><B/><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><C/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreBCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><C/><B/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewEmpty(),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreNotEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewEmpty(),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreZeroOrMoreB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewZeroOrMore(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><B/><B/></A>")
	if !Interp(auto, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreZeroOrMoreC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewZeroOrMore(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	auto := Convert(make(map[string]*relapse.Pattern), p)
	tree := newXMLStringParser("<A><C/></A>")
	if Interp(auto, tree) {
		t.Fatalf("unexpected match")
	}
}

// func TestInterpConcatOrEmptyBC(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
// 		relapse.NewOr(
// 			relapse.NewEmpty(),
// 			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
// 		),
// 		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
// 	))
// 	auto := Convert(make(map[string]*relapse.Pattern), p)
// 	tree := newXMLStringParser("<A><B/><C/></A>")
// 	if !Interp(auto, tree) {
// 		t.Fatalf("expected match")
// 	}
// }

// func TestInterpConcatOrEmptyC(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
// 		relapse.NewOr(
// 			relapse.NewEmpty(),
// 			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
// 		),
// 		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
// 	))
// 	auto := Convert(make(map[string]*relapse.Pattern), p)
// 	tree := newXMLStringParser("<A><C/></A>")
// 	if !Interp(auto, tree) {
// 		t.Fatalf("expected match")
// 	}
// }

// func TestInterpConcatOrEmptyBD(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
// 		relapse.NewOr(
// 			relapse.NewEmpty(),
// 			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
// 		),
// 		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
// 	))
// 	auto := Convert(make(map[string]*relapse.Pattern), p)
// 	tree := newXMLStringParser("<A><B/><D/></A>")
// 	if Interp(auto, tree) {
// 		t.Fatalf("unexpected match")
// 	}
// }

// func TestInterpZeroOrMoreBOrEmpty(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
// 		relapse.NewOr(
// 			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
// 			relapse.NewEmpty(),
// 		),
// 	))
// 	auto := Convert(make(map[string]*relapse.Pattern), p)
// 	tree := newXMLStringParser("<A><B/><B/></A>")
// 	if !Interp(auto, tree) {
// 		t.Fatalf("expected match")
// 	}
// }

// func TestInterpZeroOrMoreCOrEmpty(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
// 		relapse.NewOr(
// 			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
// 			relapse.NewEmpty(),
// 		),
// 	))
// 	auto := Convert(make(map[string]*relapse.Pattern), p)
// 	tree := newXMLStringParser("<A><B/><C/></A>")
// 	if Interp(auto, tree) {
// 		t.Fatalf("unexpected match")
// 	}
// }

// func TestInterpConcatCStarC(t *testing.T) {
// 	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(

// 	))
// }

func newXMLStringParser(s string) serialize.Parser {
	p := xparser.NewXMLParser()
	err := p.Init([]byte(s))
	if err != nil {
		panic(err)
	}
	return p
}

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

func newXMLStringParser(s string) serialize.Parser {
	p := xparser.NewXMLParser()
	err := p.Init([]byte(s))
	if err != nil {
		panic(err)
	}
	return p
}

func interpTest(t *testing.T, refs map[string]*relapse.Pattern, tree string) bool {
	parser := newXMLStringParser(tree)
	// auto := Convert(refs, refs["main"])
	// return Interp(auto, parser)
	return Eval(refs, parser)
}

func TestInterpNone(t *testing.T) {
	p := relapse.NewNot(relapse.NewZAny())
	tree := "<A/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpA1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	tree := "<A/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	tree := "<B/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotA(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))
	tree := "<A/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotB(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()))
	tree := "<B/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpBB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()))
	tree := "<B><B/></B>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpALeafB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), combinator.Value(funcs.StringEq(funcs.StringConst("B"), funcs.StringVar())))
	tree := "<A>B</A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpBLeafB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), combinator.Value(funcs.StringEq(funcs.StringConst("B"), funcs.StringVar())))
	tree := "<B>B</B>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatBB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotConcatBC(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	)))
	tree := "<A><B/><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotConcatBB(t *testing.T) {
	p := relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	)))
	tree := "<A><B/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	)
	tree := "<A/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpCorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	)
	tree := "<C/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeAorA(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A><A/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeAorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeBorB(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<B><A/></B>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeBorA(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A><A/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
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
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
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
	tree := "<C><A/></C>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatZAnyC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBBBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><B/><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBCCC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/><C/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyBCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatZAnyB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

//typical fundemental flaw
func TestInterpConcatZAnyCChild(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><C>B</C></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreB0(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreB3(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/><B/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreBCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewEmpty(),
	))
	tree := "<A></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreNotEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewEmpty(),
	))
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreZeroOrMoreB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewZeroOrMore(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	tree := "<A><B/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreZeroOrMoreC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewZeroOrMore(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	tree := "<A><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatOrEmptyBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewOr(
			relapse.NewEmpty(),
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatOrEmptyC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewOr(
			relapse.NewEmpty(),
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatOrEmptyBD(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewOr(
			relapse.NewEmpty(),
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><D/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpZeroOrMoreBOrEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewOr(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewEmpty(),
		),
	))
	tree := "<A><B/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpZeroOrMoreCOrEmpty(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewZeroOrMore(
		relapse.NewOr(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewEmpty(),
		),
	))
	tree := "<A><B/><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatCStarC0(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatCStarC1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatCStarC2(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatCStarC3(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/><C/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatCStarCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatCStarCCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		relapse.NewZeroOrMore(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeAAndAA(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
	)
	tree := "<A><A/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpTreeAAndAB(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
	)
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpTreeAAndBB(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndBAnyCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAndBAnyCCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndBAnyCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndBAnyCC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndBAnyCBXXXC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewZAny(),
		),
		relapse.NewConcat(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/><X/><X/><X/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAAndBAnyCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
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
	))
	tree := "<A><A><B/></A><A><C/></A></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAAndBAnyCCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
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
	))
	tree := "<A><A><C/></A><A><B/></A></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAAndBAnyCBXXXC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
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
	))
	tree := "<A><A><B/></A><X/><X/><X/><A><C/></A></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAAndBAnyCCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewAnd(
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
	))
	tree := "<A><A><C/></A><A><B/></A><A><C/></A></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAContainsB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewContains(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAContainsBCBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewContains(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><C/><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAContainsBCC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewContains(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><C/><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAContainsB0(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewContains(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpOptionalB0(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOptional(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpOptionalB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOptional(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpOptionalB2(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOptional(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><B/><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpOptionalC1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOptional(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
	))
	tree := "<A><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpInterleaveBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveCB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><C/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
	))
	tree := "<A><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpInterleaveBCAnyBC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveBCAnyBAC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><B/><A/><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveBCAnyABACA(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><A/><B/><A/><C/><A/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveBCAnyABAA(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><A/><B/><A/><A/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpInterleaveBCAnyACCBA(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><A/><C/><C/><B/><A/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpInterleaveBCAnyACCCA(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewInterleave(
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
		relapse.NewInterleave(
			relapse.NewZAny(),
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty()),
		),
	))
	tree := "<A><A/><C/><C/><C/><A/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpRefA1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	tree := "<A/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpRefB1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())
	tree := "<B/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpRefLoopA1(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOr(
		relapse.NewEmpty(),
		relapse.NewReference("main"),
	))
	tree := "<A/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpRefLoopA2(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOr(
		relapse.NewEmpty(),
		relapse.NewReference("main"),
	))
	tree := "<A><A/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpRefLoopAB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewOr(
		relapse.NewEmpty(),
		relapse.NewReference("main"),
	))
	tree := "<A><B/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpConcatOptionalBD_BD(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewOptional(relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty())),
		),
	)
	tree := "<A><B/><D/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatOptionalBD_B(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewOptional(relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty())),
		),
	)
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpConcatOptionalBD_D(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewConcat(
			relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()),
			relapse.NewOptional(relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty())),
		),
	)
	tree := "<A><D/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

//!(B:<empty>) can be <empty> and * can be B:<empty> which means B:<empty> can be accepted
func TestInterpAnyButNotB_B(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	tree := "<A><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAnyButNotB_C(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	))
	tree := "<A><C/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpNotAndBStarC_BC(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewNot(
		relapse.NewAnd(
			relapse.NewConcat(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()), relapse.NewZAny()),
			relapse.NewConcat(relapse.NewZAny(), relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty())),
		),
	))
	tree := "<A><B/><C/></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpNotAndBStarC_CB(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewNot(
		relapse.NewAnd(
			relapse.NewConcat(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty()), relapse.NewZAny()),
			relapse.NewConcat(relapse.NewZAny(), relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewEmpty())),
		),
	))
	tree := "<A><C/><B/></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAndNotAB_A(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndNotAB_B(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<B/>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

func TestInterpAndNotAB_C(t *testing.T) {
	p := relapse.NewAnd(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<C/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpOrNotAB_A(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<A/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpOrNotAB_C(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewEmpty())),
		relapse.NewNot(relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewEmpty())),
	)
	tree := "<C/>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAEndsWithBContainsAnyD_ABCD(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("B"),
			relapse.NewContains(relapse.NewTreeNode(relapse.NewAnyName(), relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty()))),
		)),
	)
	tree := "<A><B><C><D/></C></B></A>"
	if !interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("expected match")
	}
}

func TestInterpAEndsWithBContainsAnyD_ABCA(t *testing.T) {
	p := relapse.NewTreeNode(relapse.NewStringName("A"), relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(relapse.NewStringName("B"),
			relapse.NewContains(relapse.NewTreeNode(relapse.NewAnyName(), relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewEmpty()))),
		)),
	)
	tree := "<A><B><C><A/></C></B></A>"
	if interpTest(t, map[string]*relapse.Pattern{"main": p}, tree) {
		t.Fatalf("unexpected match")
	}
}

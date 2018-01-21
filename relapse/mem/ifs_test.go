//  Copyright 2018 Walter Schulze
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

package mem

import (
	"testing"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/parser/debug"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/intern"
)

var (
	c       = intern.NewConstructor()
	zany    = c.NewZAny()
	notzany = c.NewNotZAny()
	empty   = c.NewEmpty()
	set     = intern.NewSetOfPatterns()
)

func eval(ifs *ifExprs, value parser.Value) ([]*intern.Pattern, error) {
	state, err := ifs.eval(set, value)
	if err != nil {
		return nil, err
	}
	return set.Get(state).Patterns, nil
}

func TestIfsOneTrue(t *testing.T) {
	allTrue := intern.NewIfExpr(funcs.BoolConst(true), zany, notzany)
	ifs := newIfExprs([]*intern.IfExpr{allTrue})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{zany}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfsOneFalse(t *testing.T) {
	allFalse := intern.NewIfExpr(funcs.BoolConst(false), zany, notzany)
	ifs := newIfExprs([]*intern.IfExpr{allFalse})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{notzany}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfsTwoTrues(t *testing.T) {
	allTrue := intern.NewIfExpr(funcs.BoolConst(true), zany, notzany)
	ifs := newIfExprs([]*intern.IfExpr{allTrue, allTrue})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{zany, zany}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfs10Trues(t *testing.T) {
	allTrue := intern.NewIfExpr(funcs.BoolConst(true), zany, notzany)
	ifs := newIfExprs([]*intern.IfExpr{allTrue, allTrue, allTrue, allTrue, allTrue, allTrue, allTrue, allTrue, allTrue, allTrue})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{zany, zany, zany, zany, zany, zany, zany, zany, zany, zany}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfs10Mixed(t *testing.T) {
	allTrue := intern.NewIfExpr(funcs.BoolConst(true), zany, notzany)
	allFalse := intern.NewIfExpr(funcs.BoolConst(false), zany, notzany)
	ifs := newIfExprs([]*intern.IfExpr{allTrue, allFalse, allTrue, allFalse, allTrue, allFalse, allTrue, allFalse, allFalse, allTrue})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{zany, notzany, zany, notzany, zany, notzany, zany, notzany, notzany, zany}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfsOneStringVar(t *testing.T) {
	eqa := funcs.StringEq(funcs.StringConst("a"), funcs.StringVar())
	emptyIfA := intern.NewIfExpr(eqa, empty, notzany)
	ifs := newIfExprs([]*intern.IfExpr{emptyIfA})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{empty}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfsABC(t *testing.T) {
	eqa := funcs.StringEq(funcs.StringConst("a"), funcs.StringVar())
	eqb := funcs.StringEq(funcs.StringConst("b"), funcs.StringVar())
	eqc := funcs.StringEq(funcs.StringConst("c"), funcs.StringVar())
	emptyIfA := intern.NewIfExpr(eqa, empty, notzany)
	notzanyIfB := intern.NewIfExpr(eqb, notzany, zany)
	zanyIfC := intern.NewIfExpr(eqc, zany, empty)
	ifs := newIfExprs([]*intern.IfExpr{emptyIfA, zanyIfC, zanyIfC, notzanyIfB, emptyIfA, notzanyIfB, zanyIfC})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{empty, empty, empty, zany, empty, zany, empty}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

func TestIfsContainsABC(t *testing.T) {
	containsA := funcs.Contains(funcs.StringVar(), funcs.StringConst("a"))
	containsB := funcs.Contains(funcs.StringVar(), funcs.StringConst("b"))
	containsC := funcs.Contains(funcs.StringVar(), funcs.StringConst("c"))
	emptyIfA := intern.NewIfExpr(containsA, empty, notzany)
	notzanyIfB := intern.NewIfExpr(containsB, notzany, zany)
	zanyIfC := intern.NewIfExpr(containsC, zany, empty)
	ifs := newIfExprs([]*intern.IfExpr{emptyIfA, zanyIfC, zanyIfC, notzanyIfB, emptyIfA, notzanyIfB, zanyIfC})
	gots, err := eval(ifs, debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	wants := []*intern.Pattern{empty, empty, empty, zany, empty, zany, empty}
	if !deriveEquals(gots, wants) {
		t.Fatalf("want %#v got %#v", wants, gots)
	}
}

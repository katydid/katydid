//  Copyright 2017 Walter Schulze
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

package intern

import (
	"testing"

	"github.com/katydid/katydid/relapse/ast"
)

func TestZip0(t *testing.T) {
	want := []*Pattern{}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipZAny(t *testing.T) {
	want := []*Pattern{zany, zany}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipNotZAny(t *testing.T) {
	want := []*Pattern{notzany, notzany, notzany}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipNotAndZAny(t *testing.T) {
	want := []*Pattern{zany, notzany, zany, zany, zany, zany, notzany, notzany, notzany}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipA(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	c := NewConstructor()
	ap, err := c.AddPatternDecl("main", a)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap, ap, ap, ap}
	zips, zipi := Zip(want)
	if len(zips) != 1 {
		t.Errorf("wanted one in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipAB(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	b := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("b")))
	c := NewConstructor()
	ap, err := c.AddPatternDecl("a", a)
	if err != nil {
		t.Fatal(err)
	}
	bp, err := c.AddPatternDecl("b", b)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap, bp, ap, bp, bp}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABNotAndZAny(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	b := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("b")))
	c := NewConstructor()
	ap, err := c.AddPatternDecl("a", a)
	if err != nil {
		t.Fatal(err)
	}
	bp, err := c.AddPatternDecl("b", b)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap, bp, ap, zany, bp, zany, bp, bp, notzany}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipANoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	c := NewConstructor()
	ap, err := c.AddPatternDecl("a", a)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap}
	zips, zipi := Zip(want)
	if len(zips) != 1 {
		t.Errorf("wanted one in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABNoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	b := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("b")))
	c := NewConstructor()
	ap, err := c.AddPatternDecl("a", a)
	if err != nil {
		t.Fatal(err)
	}
	bp, err := c.AddPatternDecl("b", b)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap, bp}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABCNoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("a")))
	b := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("b")))
	c := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("c")))
	cons := NewConstructor()
	ap, err := cons.AddPatternDecl("a", a)
	if err != nil {
		t.Fatal(err)
	}
	bp, err := cons.AddPatternDecl("b", b)
	if err != nil {
		t.Fatal(err)
	}
	cp, err := cons.AddPatternDecl("c", c)
	if err != nil {
		t.Fatal(err)
	}
	want := []*Pattern{ap, bp, cp}
	zips, zipi := Zip(want)
	if len(zips) != 3 {
		t.Errorf("wanted three in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !deriveEquals(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

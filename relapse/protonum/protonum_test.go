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

package protonum

import (
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/relapse"
	"github.com/katydid/katydid/serialize/debug"
	"testing"
)

func check(t *testing.T, g *relapse.Grammar) {
	refs := relapse.NewRefsLookup(g)
	for name := range refs {
		if !onlyUintNames(refs[name]) {
			t.Fatalf("expected only uint names")
		}
		if anyStringNames(refs[name]) {
			t.Fatalf("expected only uint names")
		}
	}

}

func TestKeyField(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny())
	t.Logf("%v", p)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
	if gkey.GetTopPattern().GetTreeNode().GetName().GetName().GetUintValue() != 1 {
		t.Fatalf("expected field 1, but got %v", gkey)
	}
}

func TestKeyOr(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()),
		relapse.NewTreeNode(expr.NewStringName("B"), relapse.NewZAny()),
	)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
	if gkey.GetTopPattern().GetOr().GetLeftPattern().GetTreeNode().GetName().GetName().GetUintValue() != 1 {
		t.Fatalf("expected field 1, but got %v", gkey)
	}
	if gkey.GetTopPattern().GetOr().GetRightPattern().GetTreeNode().GetName().GetName().GetUintValue() != 2 {
		t.Fatalf("expected field 2, but got %v", gkey)
	}
}

func TestKeyTree(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewStringName("C"), relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()))
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	if gkey.GetTopPattern().GetTreeNode().GetName().GetName().GetUintValue() != 3 {
		t.Fatalf("expected field 3, but got %v", gkey)
	}
	if gkey.GetTopPattern().GetTreeNode().GetPattern().GetTreeNode().GetName().GetName().GetUintValue() != 1 {
		t.Fatalf("expected field 1, but got %v", gkey)
	}
	check(t, gkey)
}

func TestKeyAnyName(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(expr.NewNameChoice(expr.NewAnyName(), expr.NewStringName("C")), relapse.NewZAny()),
		relapse.NewTreeNode(expr.NewStringName("B"), relapse.NewZAny()),
	)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err == nil {
		t.Fatalf("Expected: Any Field Not Supported: Name: _, but got %v", gkey)
	}
}

func TestKeyRecursive(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewTreeNode(expr.NewStringName("C"), relapse.NewReference("main")),
		relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()),
	)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
}

func TestKeyLeftRecursive(t *testing.T) {
	p := relapse.NewOr(
		relapse.NewReference("a"),
		relapse.NewTreeNode(expr.NewStringName("C"), relapse.NewReference("main")),
		relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()),
	)
	g := p.Grammar().AddRef("a", relapse.NewReference("main"))
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	if gkey.GetTopPattern().GetOr().GetRightPattern().GetOr().GetLeftPattern().GetTreeNode().GetName().GetName().GetUintValue() != 3 {
		t.Fatalf("expected field 3, but got %v", gkey)
	}
	check(t, gkey)
}

func TestKeyLeaf(t *testing.T) {
	p := relapse.NewLeafNode(expr.NewStringVar())
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
}

func TestKeyAnyArrayIndex(t *testing.T) {
	p := relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(expr.NewStringName("E"),
			relapse.NewTreeNode(expr.NewAnyName(),
				relapse.NewConcat(
					relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()),
					relapse.NewTreeNode(expr.NewStringName("B"), relapse.NewZAny()),
				),
			),
		),
		relapse.NewZAny(),
	)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
}

func TestRepeatedMessageWithNoFieldsOfTypeMessage(t *testing.T) {
	p := relapse.NewConcat(
		relapse.NewZAny(),
		relapse.NewTreeNode(expr.NewStringName("KeyValue"),
			relapse.NewTreeNode(expr.NewAnyName(),
				relapse.NewConcat(
					relapse.NewTreeNode(expr.NewStringName("Key"), relapse.NewZAny()),
					relapse.NewTreeNode(expr.NewStringName("Value"), relapse.NewZAny()),
				),
			),
		),
		relapse.NewZAny(),
	)
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("protonum", "ProtoNum", ProtonumDescription(), g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
}

func TestUnreachable(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewStringName("NotC"), relapse.NewTreeNode(expr.NewStringName("C"), relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny())))
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err == nil {
		t.Fatal("Expected: Unknown Field Error: Name: NotC, Msg: Debug, but got %v", gkey)
	}
}

func TestNotUnreachable(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewAnyNameExcept(expr.NewStringName("NotC")), relapse.NewTreeNode(expr.NewStringName("C"), relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny())))
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err == nil {
		t.Fatalf("Expected: AnyNameExcept Not Supported Error: Name: !(NotC), but got %v", gkey)
	}
}

func TestNotUnreachableArray(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewAnyNameExcept(expr.NewStringName("NotC")), relapse.NewTreeNode(expr.NewStringName("F"),
		relapse.NewConcat(relapse.NewZAny(), relapse.NewTreeNode(expr.NewAnyName(),
			relapse.NewZAny(),
		))))
	g := p.Grammar()
	gkey, err := FieldNamesToNumbers("debug", "Debug", debug.DebugDescription(), g)
	if err == nil {
		t.Fatalf("Expected: AnyNameExcept Not Supported Error: Name: !(NotC), but got %v", gkey)
	}
}

func TestTopsyTurvy(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewAnyName(), relapse.NewTreeNode(expr.NewStringName("A"), relapse.NewZAny()))
	gkey, err := FieldNamesToNumbers("protonum", "TopsyTurvy", ProtonumDescription(), p.Grammar())
	if err == nil {
		t.Fatalf("Expected: Any Field Not Supported: Name: _, but got %v", gkey)
	}
}

func TestKnot(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewAnyName(), relapse.NewTreeNode(expr.NewAnyName(), relapse.NewTreeNode(expr.NewStringName("Elbow"), relapse.NewZAny())))
	gkey, err := FieldNamesToNumbers("protonum", "Knot", ProtonumDescription(), p.Grammar())
	if err == nil {
		t.Fatal("Expected: Any Field Not Supported: Name: _, but got %v", gkey)
	}
}

func TestRecursiveKnotTurn(t *testing.T) {
	p := relapse.NewOr(relapse.NewTreeNode(expr.NewAnyName(), relapse.NewReference("main")), relapse.NewTreeNode(expr.NewStringName("Turn"), relapse.NewZAny()))
	gkey, err := FieldNamesToNumbers("protonum", "Knot", ProtonumDescription(), p.Grammar())
	if err == nil {
		t.Fatal("Expected: Any Field Not Supported: Name: _, but got %v", gkey)
	}
}

func TestRecursiveKnotElbow(t *testing.T) {
	p := relapse.NewOr(relapse.NewTreeNode(expr.NewAnyName(), relapse.NewReference("main")), relapse.NewTreeNode(expr.NewStringName("Elbow"), relapse.NewZAny()))
	gkey, err := FieldNamesToNumbers("protonum", "Knot", ProtonumDescription(), p.Grammar())
	if err == nil {
		t.Fatal("Expected: Any Field Not Supported: Name: _, but got %v", gkey)
	}
}

func TestAnyIndex(t *testing.T) {
	p := relapse.NewTreeNode(expr.NewStringName("KeyValue"), relapse.NewTreeNode(expr.NewAnyName(), relapse.NewConcat(
		relapse.NewTreeNode(expr.NewStringName("Key"), relapse.NewZAny()),
		relapse.NewTreeNode(expr.NewStringName("Value"), relapse.NewZAny()),
	)))
	gkey, err := FieldNamesToNumbers("protonum", "ProtoNum", ProtonumDescription(), p.Grammar())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", gkey)
	check(t, gkey)
}

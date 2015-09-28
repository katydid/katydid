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

package eval

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/serialize"
	rparser "github.com/katydid/katydid/serialize/reflect"
	"github.com/katydid/katydid/tests"
	"github.com/katydid/katydid/viper/ast"
	"reflect"
	"testing"
)

func newReflectParser(m interface{}) serialize.Parser {
	s := rparser.NewReflectParser()
	s.Init(reflect.ValueOf(m))
	return s
}

func test(t *testing.T, rules *viper.Rules, m interface{}, expected bool) {
	got := Eval(rules, newReflectParser(m))
	if got != expected {
		t.Fatalf("expected %v got %v", expected, got)
	}
}

func newExpr(e interface{}) *expr.Expr {
	es := funcs.Sprint(e)
	ee, err := parser.NewParser().ParseExpr(es)
	if err != nil {
		panic(err)
	}
	return ee
}

func name(name string) *expr.Expr {
	return newExpr(funcs.StringEq(funcs.StringVar(), funcs.StringConst(name)))
}

func anyNameExcept(name string) *expr.Expr {
	return newExpr(funcs.Not(funcs.StringEq(funcs.StringVar(), funcs.StringConst(name))))
}

func anyName() *expr.Expr {
	return newExpr(funcs.BoolConst(true))
}

func any() *expr.Expr {
	return anyName()
}

func TestAny(t *testing.T) {
	test(t, viper.NewRules(
		viper.NewStart("start"),
		viper.NewFinal("start"),
		viper.NewCall("start", anyName(), "start", "start"),
		viper.NewReturn("start", "start", anyName(), "start"),
		viper.NewInternal("start", any(), "start"),
	), tests.RandomPerson(), true)
}

func TestNone(t *testing.T) {
	test(t, viper.NewRules(
		viper.NewStart("start"),
		viper.NewCall("start", anyName(), "start", "start"),
		viper.NewReturn("start", "start", anyName(), "start"),
		viper.NewInternal("start", any(), "start"),
	), tests.RandomPerson(), false)
}

func TestSaladWorry(t *testing.T) {
	m := &tests.FinanceJudo{
		SaladWorry: &tests.SaladWorry{
			MagazineFrame: []string{"a", "b"},
			XrayPilot: &tests.XrayPilot{
				AnkleCoat: proto.Int64(2),
			},
		},
		RumourSpirit: proto.Int64(1),
	}
	rules := viper.NewRules(
		viper.NewStart("start"),
		viper.NewFinal("final"),

		viper.NewCall("start", anyNameExcept("SaladWorry"), "start", "start"),
		viper.NewReturn("start", "start", anyName(), "start"),
		viper.NewInternal("start", any(), "start"),

		viper.NewCall("start", name("SaladWorry"), "start", "saladChild"),
		viper.NewReturn("start", "final", name("SaladWorry"), "final"),
		viper.NewReturn("start", "saladChild", name("SaladWorry"), "start"),
		viper.NewCall("saladChild", name("MagazineFrame"), "saladChild", "magazineFrame"),
		viper.NewCall("saladChild", anyNameExcept("MagazineFrame"), "saladChild", "start"),
		viper.NewReturn("saladChild", "start", anyNameExcept("MagazineFrame"), "saladChild"),
		viper.NewReturn("saladChild", "final", name("MagazineFrame"), "final"),
		viper.NewReturn("saladChild", "magazineFrame", name("MagazineFrame"), "saladChild"),

		viper.NewCall("final", anyName(), "final", "final"),
		viper.NewReturn("final", "final", anyName(), "final"),
		viper.NewInternal("final", any(), "final"),

		viper.NewInternal("magazineFrame", newExpr(funcs.Contains(funcs.StringVar(), funcs.StringConst("a"))), "final"),
		viper.NewInternal("magazineFrame", newExpr(funcs.Contains(funcs.StringVar(), funcs.StringConst("b"))), "final"),
		viper.NewInternal("magazineFrame", newExpr(funcs.Not(funcs.Contains(funcs.StringVar(), funcs.StringConst("a")))), "magazineFrame"),
	)
	test(t, rules, m, true)
	m.SaladWorry.MagazineFrame = []string{"c", "b"}
	test(t, rules, m, true)
	m.SaladWorry.MagazineFrame = []string{"c", "d"}
	test(t, rules, m, false)
}

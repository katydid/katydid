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
	"github.com/katydid/katydid/lang/ast"
	. "github.com/katydid/katydid/lang/builder"
	"github.com/katydid/katydid/lang/interp"
	"github.com/katydid/katydid/serialize"
	jscanner "github.com/katydid/katydid/serialize/json/scanner"
	rscanner "github.com/katydid/katydid/serialize/reflect/scanner"
	"github.com/katydid/katydid/tests"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type G map[string]*lang.Pattern

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

func newReflectScanner(m interface{}) serialize.Scanner {
	s := rscanner.NewReflectScanner()
	s.Init(reflect.ValueOf(m))
	return s
}

func test(t *testing.T, m interface{}, ps G, positive bool) {
	g := lang.NewGrammar(ps)
	scanners := map[string]func(m interface{}) serialize.Scanner{
		"json":    newJsonScanner,
		"reflect": newReflectScanner,
	}
	for scannerName, newScanner := range scanners {
		s := newScanner(m)
		match := interp.Interpret(g, s)
		if match != positive {
			t.Errorf("Expected a %v match given %s scanner from \n%v \non \n%#v", positive, scannerName, g.String(), m)
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestInterp(t *testing.T) {
	var v int64 = 1
	m := &tests.FinanceJudo{
		RumourSpirit: &v,
	}
	test(t, m, G{"main": Any()}, true)
	test(t, m, G{"main": None()}, false)
	someSpirit := MatchTree(Any(), Eval("spirit"), Any())
	test(t, m, G{
		"main":   someSpirit,
		"spirit": MatchField(`"RumourSpirit"`, "eq($int, int(1))"),
	}, true)
	test(t, m, G{
		"main":   someSpirit,
		"spirit": MatchField(`"RumourSpirit"`, "eq($int, int(2))"),
	}, false)
}

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

package auto_test

import (
	"strings"
	"testing"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/auto"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/relapse/testsuite"
)

func TestSuite(t *testing.T) {
	if !testsuite.TestSuiteExists() {
		t.Skip("testsuite not avaliable")
	}
	tests, err := testsuite.ReadTestSuite()
	if err != nil {
		t.Fatal(err)
	}
	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			test(t, testCase.Name, testCase.Grammar, testCase.Parser, testCase.Expected, "", testCase.Record)
		})
	}
}

func test(t *testing.T, name string, g *ast.Grammar, p parser.Interface, expected bool, desc string, record bool) {
	if interp.HasRecursion(g) {
		t.Skipf("convert was not designed to handle recursion")
	}
	if strings.HasPrefix(name, "GoBigOr") {
		t.Skipf("too big to fail: the number of Ors creates a state space explosion")
	}
	var a *auto.Auto
	var err error
	if record {
		a, err = auto.CompileRecord(g)
	} else {
		a, err = auto.Compile(g)
	}
	if err != nil {
		t.Fatal(err)
	}
	match, err := a.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}

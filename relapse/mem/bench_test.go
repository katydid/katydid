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

package mem_test

import (
	"flag"
	"testing"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/mem"
	"github.com/katydid/katydid/relapse/testsuite"
)

var bN = flag.Int("b.N", 0, "the number of times the benchmark function's target code must run")

func init() {
	flag.Parse()
}

func BenchmarkSuite(b *testing.B) {
	if !testsuite.BenchSuiteExists() {
		b.Skip("benchsuite not available")
	}
	benches, err := testsuite.ReadBenchmarkSuite()
	if err != nil {
		b.Fatal(err)
	}
	for _, benchCase := range benches {
		b.Run(benchCase.Name, func(b *testing.B) {
			bench(b, benchCase.Grammar, benchCase.Parsers, benchCase.Record)
		})
	}
}

func bench(b *testing.B, grammar *ast.Grammar, parsers []testsuite.ResetParser, record bool) {
	num := len(parsers)
	if *bN != 0 {
		b.N = *bN
	}
	var m *mem.Mem
	var err error
	if record {
		m, err = mem.NewRecord(grammar)
	} else {
		m, err = mem.New(grammar)
	}
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := parsers[i%num].Reset(); err != nil {
			b.Fatal(err)
		}
		if _, err := m.Validate(parsers[i%num]); err != nil {
			b.Fatal(err)
		}
	}
}

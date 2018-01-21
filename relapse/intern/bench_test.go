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

package intern_test

import (
	"testing"

	"github.com/katydid/katydid/relapse/intern"
	"github.com/katydid/katydid/relapse/testsuite"
)

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
			num := len(benchCase.Parsers)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := benchCase.Parsers[i%num].Reset(); err != nil {
					b.Fatal(err)
				}
				if _, err := intern.Interpret(benchCase.Grammar, benchCase.Record, benchCase.Parsers[i%num]); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

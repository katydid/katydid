//  Copyright 2013 Walter Schulze
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

package main

import (
	"github.com/katydid/katydid/gen"
	"github.com/katydid/katydid/tests"
)

const testStr = `
func Test{{.Name}}{{capFirst .CodecName}}(t *testing.T) {
	v := tests.Validators["{{.Name}}"]["{{.CodecName}}"]
	test(t, v.Grammar, v.Parser(), v.Expected, v.Description)
}
`

const benchStr = `
func Benchmark{{.Name}}{{capFirst .CodecName}}(b *testing.B) {
	v := tests.BenchValidators["{{.Name}}"]["{{.CodecName}}"]
	bench(b, v.Grammar, tests.Random{{.MessageName}}{{capFirst .CodecName}}Parser)
}
`

func main() {
	gen := gen.NewFunc("convert_test")
	gen(testStr, "convert.gen_test.go", tests.ValidatorList(), `"testing"`, `"github.com/katydid/katydid/tests"`)
	gen(benchStr, "convert.gen_bench_test.go", tests.BenchValidatorList(), `"testing"`, `"github.com/katydid/katydid/tests"`)
}

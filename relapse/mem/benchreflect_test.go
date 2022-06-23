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
	"reflect"
	"testing"

	"github.com/katydid/katydid/relapse/testsuite"

	parser "github.com/katydid/katydid/parser/reflect"
	"github.com/katydid/katydid/relapse"
)

type TestStruct struct {
	String string `validate:"required" json:"StringVal"`
}

// 219 ns/op
// func BenchmarkStructLevelValidationSuccess(b *testing.B) {

// 	validate := New()
// 	validate.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

// 	tst := TestStruct{
// 		String: "good value",
// 	}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		validate.Struct(tst)
// 	}
// }

// 132 ns/op
func BenchmarkGoPlaygroundStructLevelValidationSuccess(b *testing.B) {
	tst := TestStruct{
		String: "good value",
	}
	s := ".String:*"
	p := parser.NewReflectParser().Init(reflect.ValueOf(tst))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if !valid {
		b.Fatalf("expected valid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

// 263 ns/op
// func BenchmarkStructSimpleSuccess(b *testing.B) {

// 	validate := New()

// 	type Foo struct {
// 		StringValue string `validate:"min=5,max=10"`
// 		IntValue    int    `validate:"min=5,max=10"`
// 	}

// 	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		validate.Struct(validFoo)
// 	}
// }

// 155 ns/op
func BenchmarkGoPlaygroundStructSimpleSuccess(b *testing.B) {
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}
	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	s := `{
		(StringValue->and(ge(length($string), 5), le(length($string), 10)))? ;
		(IntValue->and(ge($int, 5), le($int, 10)))?
	}`
	p := parser.NewReflectParser().Init(reflect.ValueOf(validFoo))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if !valid {
		b.Fatalf("expected valid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

type TestString struct {
	BlankTag  string `validate:""`
	Required  string `validate:"required"`
	Len       string `validate:"len=10"`
	Min       string `validate:"min=1"`
	Max       string `validate:"max=10"`
	MinMax    string `validate:"min=1,max=10"`
	Lt        string `validate:"lt=10"`
	Lte       string `validate:"lte=10"`
	Gt        string `validate:"gt=10"`
	Gte       string `validate:"gte=10"`
	OmitEmpty string `validate:"omitempty,min=1,max=10"`
	Sub       *SubTest
	SubIgnore *SubTest `validate:"-"`
}

type SubTest struct {
	Test string `validate:"required"`
}

// 1504 ns/op
// func BenchmarkStructComplexSuccess(b *testing.B) {

// 	validate := New()

// 	tSuccess := &TestString{
// 		Required:  "Required",
// 		Len:       "length==10",
// 		Min:       "min=1",
// 		Max:       "1234567890",
// 		MinMax:    "12345",
// 		Lt:        "012345678",
// 		Lte:       "0123456789",
// 		Gt:        "01234567890",
// 		Gte:       "0123456789",
// 		OmitEmpty: "",
// 		Sub: &SubTest{
// 			Test: "1",
// 		},
// 		SubIgnore: &SubTest{
// 			Test: "",
// 		},
// 		Anonymous: struct {
// 			A string `validate:"required"`
// 		}{
// 			A: "1",
// 		},
// 		Iface: &Impl{
// 			F: "123",
// 		},
// 	}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		validate.Struct(tSuccess)
// 	}
// }

// 444 ns/op
func BenchmarkGoPlaygroundStructComplexSuccess(b *testing.B) {
	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
		Lt:        "012345678",
		Lte:       "0123456789",
		Gt:        "01234567890",
		Gte:       "0123456789",
		OmitEmpty: "",
		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
	}
	s := `#main = {
		BlankTag->eq(length($string), 0) ;
		Required->gt(length($string), 0) ;
		(Len->eq(length($string), 10))? ;
		(Min->ge(length($string), 1))? ;
		(Max->le(length($string), 10))? ;
		(MinMax->and(ge(length($string), 1), le(length($string), 10)))? ;
		(Lt->lt(length($string), 10))? ;
		(Lte->le(length($string), 10))? ;
		(Gt->gt(length($string), 10))? ;
		(Gte->ge(length($string), 10))? ;
		(OmitEmpty->or(eq(length($string), 0), and(ge(length($string), 1), le(length($string), 10))))? ;
		(Sub: @subtest)? ;
		(SubIgnore: *)? ;
	}

	#subtest = Test->gt(length($string), 0)
	`
	p := parser.NewReflectParser().Init(reflect.ValueOf(tSuccess))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if !valid {
		b.Fatalf("expected valid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

// 7585 ns/op
// func BenchmarkStructComplexFailure(b *testing.B) {

// 	validate := New()

// 	tFail := &TestString{
// 		Required:  "",
// 		Len:       "",
// 		Min:       "",
// 		Max:       "12345678901",
// 		MinMax:    "",
// 		Lt:        "0123456789",
// 		Lte:       "01234567890",
// 		Gt:        "1",
// 		Gte:       "1",
// 		OmitEmpty: "12345678901",
// 		Sub: &SubTest{
// 			Test: "",
// 		},
// 		Anonymous: struct {
// 			A string `validate:"required"`
// 		}{
// 			A: "",
// 		},
// 		Iface: &Impl{
// 			F: "12",
// 		},
// 	}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		validate.Struct(tFail)
// 	}
// }

// 414 ns/op
func BenchmarkGoPlaygroundStructComplexFailure(b *testing.B) {
	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
		Lt:        "0123456789",
		Lte:       "01234567890",
		Gt:        "1",
		Gte:       "1",
		OmitEmpty: "12345678901",
		Sub: &SubTest{
			Test: "",
		},
	}
	s := `#main = {
		BlankTag->eq(length($string), 0) ;
		Required->gt(length($string), 0) ;
		(Len->eq(length($string), 10))? ;
		(Min->ge(length($string), 1))? ;
		(Max->le(length($string), 10))? ;
		(MinMax->and(ge(length($string), 1), le(length($string), 10)))? ;
		(Lt->lt(length($string), 10))? ;
		(Lte->le(length($string), 10))? ;
		(Gt->gt(length($string), 10))? ;
		(Gte->ge(length($string), 10))? ;
		(OmitEmpty->or(eq(length($string), 0), and(ge(length($string), 1), le(length($string), 10))))? ;
		(Sub: @subtest)? ;
		(SubIgnore: *)? ;
	}

	#subtest = Test->gt(length($string), 0)
	`
	p := parser.NewReflectParser().Init(reflect.ValueOf(tFail))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if valid {
		b.Fatalf("expected invalid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

// 132 ns/op
func BenchmarkGoPlaygroundStructComplexSuccess2(b *testing.B) {
	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
		Lt:        "012345678",
		Lte:       "0123456789",
		Gt:        "01234567890",
		Gte:       "0123456789",
		OmitEmpty: "",
		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
	}
	s := `#main = [
		BlankTag->eq(length($string), 0) ,
		Required->gt(length($string), 0) ,
		Len->eq(length($string), 10) ,
		Min->ge(length($string), 1) ,
		Max->le(length($string), 10) ,
		MinMax->and(ge(length($string), 1), le(length($string), 10)) ,
		Lt->lt(length($string), 10) ,
		Lte->le(length($string), 10) ,
		Gt->gt(length($string), 10) ,
		Gte->ge(length($string), 10) ,
		OmitEmpty->or(eq(length($string), 0), and(ge(length($string), 1), le(length($string), 10))) ,
		Sub: @subtest ,
		SubIgnore: *
	]

	#subtest = Test->gt(length($string), 0)
	`
	p := parser.NewReflectParser().Init(reflect.ValueOf(tSuccess))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if !valid {
		b.Fatalf("expected valid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

// 129 ns/op
func BenchmarkGoPlaygroundStructComplexFailure2(b *testing.B) {
	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
		Lt:        "0123456789",
		Lte:       "01234567890",
		Gt:        "1",
		Gte:       "1",
		OmitEmpty: "12345678901",
		Sub: &SubTest{
			Test: "",
		},
	}
	s := `#main = [
		BlankTag->eq(length($string), 0) ,
		Required->gt(length($string), 0) ,
		Len->eq(length($string), 10) ,
		Min->ge(length($string), 1) ,
		Max->le(length($string), 10) ,
		MinMax->and(ge(length($string), 1), le(length($string), 10)) ,
		Lt->lt(length($string), 10) ,
		Lte->le(length($string), 10) ,
		Gt->gt(length($string), 10) ,
		Gte->ge(length($string), 10) ,
		OmitEmpty->or(eq(length($string), 0), and(ge(length($string), 1), le(length($string), 10))) ,
		Sub: @subtest ,
		SubIgnore: *
	]

	#subtest = Test->gt(length($string), 0)
	`
	p := parser.NewReflectParser().Init(reflect.ValueOf(tFail))
	g, err := relapse.Parse(s)
	if err != nil {
		b.Fatal(err)
	}
	m, err := relapse.Prepare(g)
	if err != nil {
		b.Fatal(err)
	}
	valid, err := relapse.Validate(m, p)
	if err != nil {
		b.Fatal(err)
	}
	if valid {
		b.Fatalf("expected invalid")
	}
	bench(b, g, []testsuite.ResetParser{p}, true)
}

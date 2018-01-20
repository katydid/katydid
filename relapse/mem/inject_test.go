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
	"strings"
	"testing"

	reflectparser "github.com/katydid/katydid/parser/reflect"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/mem"
)

func NewInjectable() *injectableInt {
	return &injectableInt{}
}

type injectableInt struct {
	context *funcs.Context
}

func (this *injectableInt) Eval() (int64, error) {
	return this.context.Value.(int64), nil
}

func (this *injectableInt) Compare(that funcs.Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if _, ok := that.(*injectableInt); ok {
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *injectableInt) String() string {
	return "inject()"
}

func (this *injectableInt) Hash() uint64 {
	return 17
}

func (this *injectableInt) SetContext(context *funcs.Context) {
	this.context = context
}

func (this *injectableInt) HasVariable() bool {
	return true
}

func init() {
	funcs.Register("inject", NewInjectable)

	injectNumber = G{
		"main": InPath("Num", Value(funcs.IntEq(funcs.IntVar(), NewInjectable()))),
	}
}

var injectNumber = G{}

type Number struct {
	Num int64
}

func testInject(t *testing.T, m *mem.Mem) bool {
	parser := reflectparser.NewReflectParser()
	parser.Init(reflect.ValueOf(&Number{Num: 456}))
	res, err := m.Validate(parser)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestInject(t *testing.T) {
	grammar := injectNumber.Grammar()
	m, err := mem.New(grammar)
	if err != nil {
		t.Fatal(err)
	}
	c := &funcs.Context{Value: int64(0)}
	m.SetContext(c)
	c.Value = int64(456)
	if !testInject(t, m) {
		t.Fatalf("expected match")
	}
	c.Value = int64(123)
	if testInject(t, m) {
		t.Fatalf("expected non match")
	}
}

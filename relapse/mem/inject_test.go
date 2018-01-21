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
	"fmt"
	"reflect"
	"strings"
	"testing"

	reflectparser "github.com/katydid/katydid/parser/reflect"
	"github.com/katydid/katydid/relapse/ast"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/mem"
	relapseparser "github.com/katydid/katydid/relapse/parser"
)

func NewInjectable() *injectableInt {
	fmt.Printf("NewInjectable\n")
	return &injectableInt{}
}

type injectableInt struct {
	context *funcs.Context
}

func (this *injectableInt) Eval() (int64, error) {
	v := this.context.Value.(int64)
	fmt.Printf("eval = %d\n", v)
	return v, nil
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
	fmt.Printf("context set\n")
}

func (this *injectableInt) HasVariable() bool {
	return true
}

func init() {
	funcs.Register("inject", NewInjectable)

	parsedGrammar, err := relapseparser.ParseGrammar("Num:->eq($int, inject())")
	if err != nil {
		panic(err)
	}
	injectNumber = G(ast.NewRefLookup(parsedGrammar))
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
	fmt.Printf("parsed Grammar: %s\n", grammar)
	m, err := mem.New(grammar)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("trying to set context...\n")
	c := &funcs.Context{Value: int64(0)}
	m.SetContext(c)
	fmt.Printf("hopefully context was set\n")
	c.Value = int64(456)
	if !testInject(t, m) {
		t.Fatalf("expected match")
	}
	c.Value = int64(123)
	if testInject(t, m) {
		t.Fatalf("expected non match")
	}
}

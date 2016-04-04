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

package compose_test

import (
	"github.com/katydid/katydid/parser/debug"
	"github.com/katydid/katydid/relapse/compose"
	. "github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func TestNoEqualError(t *testing.T) {
	exprStr := Sprint(StringEq(ElemStrings(NewListOfString([]String{StringVar()}), IntConst(3)), StringConst("0123456789")))
	t.Logf(exprStr)
	e, err := parser.NewParser().ParseExpr(exprStr)
	if err != nil {
		t.Fatal(err)
	}
	b, err := compose.NewBool(e)
	if err != nil {
		t.Fatal(err)
	}
	f, err := compose.NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	if v, err := f.Eval(debug.NewStringValue("0")); err != nil {
		t.Fatal(err)
	} else if v {
		t.Fatalf("expected false")
	}
}

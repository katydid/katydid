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

package funcs

import (
	"testing"
)

func TestString1(t *testing.T) {
	out := Sprint(BoolEq(BoolConst(true), BoolVar()))
	exp := "== true"
	if out != exp {
		t.Fatalf("expected %s, but got %s", exp, out)
	}
}

func TestString2(t *testing.T) {
	out := Sprint(IntGE(ElemInts(IntsConst([]int64{1, 2}), IntVar()), IntVar()))
	exp := "->ge(elem([]int{int(1),int(2)},$int),$int)"
	if out != exp {
		t.Fatalf("expected %s, but got %s", exp, out)
	}
}

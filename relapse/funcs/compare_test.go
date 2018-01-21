//  Copyright 2018 Walter Schulze
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

import "testing"

func TestCompare(t *testing.T) {
	eq1 := IntEq(IntConst(1), IntVar())
	eq2 := IntEq(IntConst(1), IntVar())
	if eq1.Compare(eq2) != 0 {
		t.Fatalf("Compare is not equal")
	}
	eq2 = IntEq(IntConst(2), IntVar())
	if eq1.Compare(eq2) != (-1 * eq2.Compare(eq1)) {
		t.Fatalf("Compare is not symmetrical")
	}
	eq1.(*intEq).hash = eq2.(*intEq).hash
	if eq1.Compare(eq2) != (-1 * eq2.Compare(eq1)) {
		t.Fatalf("hash has too much power")
	}
}

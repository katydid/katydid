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

package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	positives := []string{
		`"String"`,
		"false",
		`== "bla"`,
		`eq($string, "bla")`,
		`1`,
		`1.0`,
	}
	negatives := []string{
		"a",
		`= "bla"`,
	}
	for _, in := range positives {
		_, err := NewParser().ParseExpr(in)
		if err != nil {
			t.Errorf("%s results in error: %s", in, err)
		}
	}
	for _, in := range negatives {
		_, err := NewParser().ParseExpr(in)
		if err == nil {
			t.Errorf("%s results in no error", in)
		}
	}
}

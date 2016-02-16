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
		`return parent child eq($int, int(1)) next;`,
		"call current eq($int, int(1)) nextparent nextchild;",
		"internal current eq($int, int(1)) next ;",
		`return parent child eq($int, int(1)) next;
		internal current eq($int, int(1)) next ;`,
	}
	negatives := []string{
		"external current eq($int, int(1)) next ;",
	}
	for _, in := range positives {
		_, err := NewParser().ParseViper(in)
		if err != nil {
			t.Errorf("%s results in error: %s", in, err)
		}
	}
	for _, in := range negatives {
		_, err := NewParser().ParseViper(in)
		if err == nil {
			t.Errorf("%s results in no error", in)
		}
	}
}

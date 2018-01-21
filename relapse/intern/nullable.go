//  Copyright 2017 Walter Schulze
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

package intern

import "github.com/katydid/katydid/relapse/sets"

func anyNullable(ps []*Pattern) bool {
	for _, any := range ps {
		if any.nullable {
			return true
		}
	}
	return false
}

func allNullable(ps []*Pattern) bool {
	for _, all := range ps {
		if !all.nullable {
			return false
		}
	}
	return true
}

func newNullableSet(patterns []*Pattern) sets.Bits {
	nulls := sets.NewBits(len(patterns))
	for i, p := range patterns {
		nulls.Set(i, p.nullable)
	}
	return nulls
}

func nullables(ps []*Pattern) []bool {
	nulls := make([]bool, len(ps))
	for i, p := range ps {
		nulls[i] = p.nullable
	}
	return nulls
}

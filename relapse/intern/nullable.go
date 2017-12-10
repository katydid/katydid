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

import (
	"fmt"
)

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

func Nullable(refs map[string]*Pattern, p *Pattern) bool {
	switch p.Type {
	case Empty:
		return true
	case Node:
		return false
	case Concat:
		return allNullable(p.Patterns)
	case Or:
		return anyNullable(p.Patterns)
	case And:
		return allNullable(p.Patterns)
	case ZeroOrMore:
		return true
	case Reference:
		return refs[p.Ref].nullable
	case Not:
		return !p.Patterns[0].nullable
	case Contains:
		return p.Patterns[0].nullable
	case Optional:
		return true
	case Interleave:
		return allNullable(p.Patterns)
	}
	panic(fmt.Sprintf("unknown pattern type %v", p))
}

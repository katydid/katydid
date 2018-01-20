//  Copyright 2013 Walter Schulze
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
	"strings"
)

// sjoin calls the String method of all parameters and joins them with a comma as a separator.
func sjoin(s ...Stringer) string {
	ss := make([]string, len(s))
	for i := range s {
		ss[i] = s[i].String()
	}
	return strings.Join(ss, ",")
}

//Sprint returns the string printout of the function.
func Sprint(s Stringer) string {
	if shorthand, ok := s.(interface {
		Shorthand() (string, bool)
	}); ok {
		if short, sok := shorthand.Shorthand(); sok {
			return short
		}
	}
	return "->" + s.String()
}

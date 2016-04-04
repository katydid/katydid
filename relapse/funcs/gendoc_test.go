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
	"fmt"
	"github.com/katydid/katydid/gen"
	"os"
	"sort"
	"strings"
	"testing"
)

func toString(s string) string {
	return strings.ToLower(strings.Replace(strings.Replace(strings.Replace(s, "_BYTES", "_[]byte", 1), "SINGLE_", "", 1), "LIST_", "[]", 1))
}

//TestGenFunList generates the current list of functions.
func TestGenFuncList(t *testing.T) {
	funcs := []string{}
	for name, us := range funcsMap.nameToUniq {
		for _, u := range us {
			f := funcsMap.uniqToFunc[u]
			ins := make([]string, len(f.In))
			for i, in := range f.In {
				if f.InConst[i] {
					ins[i] = gen.LowerFirst(f.InNames[i]) + " const " + toString(in.String())
				} else {
					ins[i] = gen.LowerFirst(f.InNames[i]) + " " + toString(in.String())
				}
			}
			funcs = append(funcs, fmt.Sprintf("func %v(%v) %v", name, strings.Join(ins, ","), toString(f.Out.String())))
		}
	}
	sort.Strings(funcs)
	fmt.Fprintf(os.Stderr, "%v\n", strings.Join(funcs, "\n"))
}

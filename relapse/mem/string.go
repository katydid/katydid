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

package mem

import (
	"strings"

	"github.com/katydid/katydid/relapse/funcs"
)

func addtab(s string) string {
	ss := strings.Split(s, "\n")
	for i := range ss {
		ss[i] = "\t" + ss[i]
	}
	return strings.Join(ss, "\n")
}

func (this *ifExprs) String() string {
	if this.ret != nil {
		ss := make([]string, len(this.ret))
		for i := range this.ret {
			ss[i] = this.ret[i].String()
		}
		return strings.Join(ss, ", ")
	}
	sthen := addtab(this.then.String())
	sels := addtab(this.els.String())
	sfunc := funcs.Sprint(this.cond)
	return "{\n" + sfunc + "\nThen:\n" + sthen + "\nElse:\n" + sels + "}"
}

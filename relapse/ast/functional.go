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

package ast

//Filter returns a filtered list of patterns as specified by the predicate function.
func Filter(f func(p *Pattern) bool, ps []*Pattern) []*Pattern {
	fs := make([]*Pattern, 0, len(ps))
	for i, p := range ps {
		if f(p) {
			fs = append(fs, ps[i])
		}
	}
	return fs
}

//Map applies the provided function to each pattern in the list and returns the modified list.
func Map(f func(p *Pattern) *Pattern, ps []*Pattern) []*Pattern {
	fs := make([]*Pattern, len(ps))
	for i, p := range ps {
		fs[i] = f(p)
	}
	return fs
}

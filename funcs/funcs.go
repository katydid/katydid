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
	"code.google.com/p/go.text/unicode/norm"
	"strings"
)

type contains struct {
	V1 String
	V2 String
}

func (this *contains) Eval(buf []byte) bool {
	v1 := this.V1.Eval(buf)
	v2 := this.V2.Eval(buf)
	return strings.Contains(v1, v2)
}

func init() {
	Register("contains", new(contains))
}

type nfkc struct {
	V1 String
}

func (this *nfkc) Eval(buf []byte) string {
	return norm.NFKC.String(this.V1.Eval(buf))
}

func init() {
	Register("nfkc", new(nfkc))
}

type not struct {
	V1 Bool
}

func (this *not) Eval(buf []byte) bool {
	return !this.V1.Eval(buf)
}

func init() {
	Register("not", new(not))
}

type exists struct {
}

func (this *exists) Eval(buf []byte) bool {
	return true
}

func init() {
	Register("exists", new(exists))
}

type lenString struct {
	V1 String
}

func (this *lenString) Eval(buf []byte) int64 {
	return int64(len(this.V1.Eval(buf)))
}

func init() {
	Register("length", new(lenString))
}

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

type equalFold struct {
	V1 String
	V2 String
}

func (this *equalFold) Eval(buf []byte) bool {
	v1 := this.V1.Eval(buf)
	v2 := this.V2.Eval(buf)
	return strings.EqualFold(v1, v2)
}

func init() {
	Register("eqFold", new(equalFold))
}

type prefix struct {
	V1 String
	V2 String
}

func (this *prefix) Eval(buf []byte) bool {
	v1 := this.V1.Eval(buf)
	v2 := this.V2.Eval(buf)
	return strings.HasPrefix(v1, v2)
}

func init() {
	Register("prefix", new(prefix))
}

type suffix struct {
	V1 String
	V2 String
}

func (this *suffix) Eval(buf []byte) bool {
	v1 := this.V1.Eval(buf)
	v2 := this.V2.Eval(buf)
	return strings.HasSuffix(v1, v2)
}

func init() {
	Register("suffix", new(suffix))
}

type nfc struct {
	V1 String
}

func (this *nfc) Eval(buf []byte) string {
	return norm.NFC.String(this.V1.Eval(buf))
}

func init() {
	Register("nfc", new(nfc))
}

type nfd struct {
	V1 String
}

func (this *nfd) Eval(buf []byte) string {
	return norm.NFD.String(this.V1.Eval(buf))
}

func init() {
	Register("nfd", new(nfd))
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

type nfkd struct {
	V1 String
}

func (this *nfkd) Eval(buf []byte) string {
	return norm.NFKD.String(this.V1.Eval(buf))
}

func init() {
	Register("nfkd", new(nfkd))
}

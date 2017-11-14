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

package ast

import (
	"bytes"
	"fmt"
	"testing"
)

func testParseStrErr(t *testing.T, s string) {
	_, err := parseBytes(s)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func testParseStr(t *testing.T, in []byte, s string) {
	out, err := parseBytes(s)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(in, out) {
		t.Fatalf("expected %#v got %#v from %s", in, out, s)
	}
}

func testParse(t *testing.T, in []byte) {
	s := fmt.Sprintf("%#v", in)
	testParseStr(t, in, s)
}

func TestParseBytes(t *testing.T) {
	testParse(t, []byte{})
	testParse(t, []byte{1})
	testParse(t, []byte{1, 2})
	testParse(t, []byte{255})
	testParseStr(t, []byte{255, 1, 2}, `[]byte{255,1,2}`)
	testParseStr(t, []byte{255, 1, 2, 0, 0, 0}, `[]byte{0xff,0x01,2, 0x0, 0, 0x00}`)
	testParseStr(t, []byte{255, 'a', 'b', 0, 0, '0'}, `[]byte{0xff,'a','b', 0x0, 0, '0'}`)
	testParseStr(t, []byte{}, `[]byte { }`)
	testParseStrErr(t, `[]byte{0922}`)
	testParseStrErr(t, `[]byte{0x255}`)
	testParseStrErr(t, `[]byte{256}`)
}

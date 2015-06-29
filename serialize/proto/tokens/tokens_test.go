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

package tokens

import (
	"github.com/katydid/katydid/graphviz"
	"github.com/katydid/katydid/protoparser"
	"io/ioutil"
	"testing"
)

func output(t *testing.T, dot string, filename string) {
	t.Logf("%v\n", dot)
	if graphviz.Installed() {
		svg, err := graphviz.DrawSVG(dot)
		if err != nil {
			t.Fatal(err)
		}
		err = ioutil.WriteFile(filename, []byte(svg), 0666)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTokensPerson(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("test.proto", ".")
	if err != nil {
		t.Fatal(err)
	}
	m, err := New("test", "Person", fileDescriptorSet)
	if err != nil {
		t.Fatal(err)
	}
	output(t, m.Dot(), "person.svg")
}

func TestTokensSrcTree(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("test.proto", ".")
	if err != nil {
		t.Fatal(err)
	}
	m, err := New("test", "SrcTree", fileDescriptorSet)
	if err != nil {
		t.Fatal(err)
	}
	output(t, m.Dot(), "srctree.svg")
}

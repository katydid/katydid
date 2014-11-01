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

package main_test

import (
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	"testing"
)

func TestAttachedComments(t *testing.T) {
	var q = `//Unattached

root = a.b
//attached
a.b = start
/*attached*/
start c = accept
/*attached*/start _ = start
/*unattached*/

start d = accept`
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(q)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	count := 0
	for _, rule := range rules.Rules {
		c := rule.GetAttachedComment().GetContent()
		if len(c) > 0 {
			if c != "attached" {
				t.Fatalf("received attached comment = %s", c)
			} else {
				count++
			}
		}
	}
	if count != 3 {
		t.Fatalf("did not find enough attached comments")
	}
}

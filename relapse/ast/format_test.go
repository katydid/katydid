//  Copyright 2015 Walter Schulze
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

package ast_test

import (
	"strings"
	"testing"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/parser"
)

func format(t *testing.T, input string, expected string) {
	g, err := parser.ParseGrammar(input)
	if err != nil {
		t.Errorf("parse error: %v given input: <%s>", err, input)
		return
	}
	g.Format()
	output := g.String()
	if expected != output {
		expectedStr := strings.Replace(strings.Replace(expected, " ", "_", -1), "\t", "<tab>", -1)
		outputStr := strings.Replace(strings.Replace(output, " ", "_", -1), "\t", "<tab>", -1)
		t.Errorf("format failure: expected \n<%s>, but got \n<%s>", expectedStr, outputStr)
	}
	input = expected
	g2, err := parser.ParseGrammar(input)
	if err != nil {
		t.Errorf("parse error2: %v given input: <%s>", err, input)
		return
	}
	g2.Format()
	output = g.String()
	if expected != output {
		expectedStr := strings.Replace(strings.Replace(expected, " ", "_", -1), "\t", "<tab>", -1)
		outputStr := strings.Replace(strings.Replace(output, " ", "_", -1), "\t", "<tab>", -1)
		t.Errorf("format failure2: expected \n<%s>, but got \n<%s>", expectedStr, outputStr)
	}
}

func TestFormatGrammar(t *testing.T) {
	format(t, " #main =  *", "#main = *")
	format(t,
		`//attachedcomment
	  *`,
		`//attachedcomment
*`)
	format(t,
		`//unattachedcomment

		*`,
		`//unattachedcomment
*`)
	//2 pattern declarations
	format(t,
		`#main = *
			#other = a->any()`,
		`#main = *
#other = a->any()`)
	//2 pattern declarations without a new line
	format(t,
		`#main = *	#other = a->any()`,
		`#main = * #other = a->any()`)
	//3 pattern declarations
	format(t,
		`*
			#other = a->any()

			#more = (*)*`,
		`*
#other = a->any()
#more = (*)*`)
	//treenode
	format(t,
		`#main =   "a":*`,
		`#main = a:*`)
	format(t,
		ast.NewGrammar(map[string]*ast.Pattern{"main": ast.NewReference("ref1"), "ref1": ast.NewReference("main")}).String(),
		`@ref1
#ref1=@main`)
}

func formatSpace(t *testing.T, input string, expected string) {
	format(t, "*"+input, "*"+expected)
}

func TestFormatSpace(t *testing.T) {
	formatSpace(t, "//linecomment\n", "//linecomment\n")
	formatSpace(t, "/*blockcomment*/", "/*blockcomment*/")
	formatSpace(t, " \t\t\t //linecomment \t \n", "//linecomment\n")
	formatSpace(t, "//linecomment\n\n\n", "//linecomment\n")
	formatSpace(t, "//linecomment\n", "//linecomment\n")
	formatSpace(t, "/*blockcomment*/\n\n", "/*blockcomment*/")
	formatSpace(t, "\t\t/*blockcomment*/  ", "/*blockcomment*/")
	formatSpace(t,
		`//linecomment

/*blockcomment*/

//linecomment2
`, `//linecomment

/*blockcomment*/

//linecomment2
`)
	formatSpace(t,
		`  //linecomment



	/*blockcomment*/  



  //linecomment2  
`, `//linecomment

/*blockcomment*/

//linecomment2
`)
	formatSpace(t,
		`/*blockcomment
*/`, `/*blockcomment
*/`)
	formatSpace(t,
		`/*blockcomment

*/`, `/*blockcomment

*/`)
	formatSpace(t,
		`	/*blockcomment
		more text
*/	
		`, `/*blockcomment
		more text
*/`)
	formatSpace(t,
		`	/*blockcomment
		more text
*/	

		/*blockcomment2*/





	/*  
		block3
		block2
		block1
	*/
		`, `/*blockcomment
		more text
*/

/*blockcomment2*/

/*  
		block3
		block2
		block1
	*/`)
	formatSpace(t,
		`	/*blockcomment
		more text
*/	

		//linecomment





	/*  
		block3
		block2
		block1
	*/
		`, `/*blockcomment
		more text
*/

//linecomment

/*  
		block3
		block2
		block1
	*/`)
}

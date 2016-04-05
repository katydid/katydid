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

//Package relapse contains the relapse validation language and the functions necessary for running it.
//See katydid.github.io for the language documentation.
package relapse

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/auto"
	"github.com/katydid/katydid/relapse/interp"
	relapseParser "github.com/katydid/katydid/relapse/parser"
)

//Parse parses the relapse string into an ast (abstract syntax tree)
func Parse(relapse string) (*ast.Grammar, error) {
	return relapseParser.ParseGrammar(relapse)
}

//Interpret validates the parser with the given grammar.
func Interpret(g *ast.Grammar, p parser.Interface) bool {
	return interp.Interpret(g, p)
}

//Compile compiles the grammar into an efficient automaton.
func Compile(g *ast.Grammar) *auto.Auto {
	return auto.Compile(g)
}

//Execute validates the parser with the given automaton.
func Execute(a *auto.Auto, p parser.Interface) bool {
	return auto.Interpret(a, p)
}

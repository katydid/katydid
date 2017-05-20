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

package parser

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/lexer"
)

type errWrongType struct {
	typ string
	res interface{}
}

func (this *errWrongType) Error() string {
	return fmt.Sprintf("relapse/parser: expected %s, but got %#v", this.typ, this.res)
}

//ParseGrammar parses a relapse grammar and returns an abstract syntax tree.
func (this *Parser) ParseGrammar(s string) (*ast.Grammar, error) {
	scanner := lexer.NewLexer([]byte(s))
	g, err := this.Parse(scanner)
	if err != nil {
		return nil, err
	}
	gram, ok := g.(*ast.Grammar)
	if !ok {
		return nil, &errWrongType{"*relapse.Grammar", g}
	}
	return gram, nil
}

//ParseGrammar parses a relapse grammar and returns an abstract syntax tree.
func ParseGrammar(s string) (*ast.Grammar, error) {
	p := NewParser()
	return p.ParseGrammar(s)
}

//ParseExpr returns a parsed expression or error, given a string.
func (this *Parser) ParseExpr(s string) (res *ast.Expr, err error) {
	scanner := lexer.NewLexer([]byte(s))
	g, err := this.Parse(scanner)
	if err != nil {
		return nil, err
	}
	gram, ok := g.(*ast.Grammar)
	if !ok {
		return nil, &errWrongType{"*relapse.Grammar", g}
	}
	if len(gram.GetPatternDecls()) != 0 {
		return nil, &errWrongType{"found pattern declarations where none were expected", g}
	}
	if gram.GetTopPattern().LeafNode == nil {
		return nil, &errWrongType{"LeafNode == nil", g}
	}
	return gram.GetTopPattern().GetLeafNode().GetExpr(), nil
}

//ParseExpr returns a parsed expression or error, given a string.
func ParseExpr(s string) (*ast.Expr, error) {
	p := NewParser()
	return p.ParseExpr(s)
}

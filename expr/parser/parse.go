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

package parser

import (
	"fmt"
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/lexer"
)

//ErrWrongType contains the value that failed to assert to *expr.Expr
type ErrWrongType struct {
	typ string
	res interface{}
}

func (this *ErrWrongType) Error() string {
	return fmt.Sprintf("expected %s, but got %#v", this.typ, this.res)
}

//ParseExpr returns a parsed expression or error, given a string.
func (this *Parser) ParseExpr(s string) (res *expr.Expr, err error) {
	scanner := lexer.NewLexer([]byte(s))
	e, err := this.Parse(scanner)
	if err != nil {
		return nil, err
	}
	expr, ok := e.(*expr.Expr)
	if !ok {
		return nil, &ErrWrongType{"*expr.Expr", e}
	}
	return expr, nil
}

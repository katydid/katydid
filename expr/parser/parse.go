package parser

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
)

type ErrWrongType struct {
	typ string
	res interface{}
}

func (this *ErrWrongType) Error() string {
	return fmt.Sprintf("expected %s, but got %#v", this.typ, this.res)
}

func (this *Parser) ParseExpr(scanner Scanner) (res *expr.Expr, err error) {
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

package lang

import (
	"github.com/katydid/katydid/expr/ast"
)

func NewEmpty() *Pattern {
	return &Pattern{
		Empty: &Empty{},
	}
}

func NewEmptySet() *Pattern {
	return &Pattern{
		EmptySet: &EmptySet{},
	}
}

func newOpenParen() *expr.Keyword {
	return &expr.Keyword{Value: "("}
}

func newCloseParen() *expr.Keyword {
	return &expr.Keyword{Value: ")"}
}

func newComma() *expr.Keyword {
	return &expr.Keyword{Value: ","}
}

func NewTreeNode(name *expr.Expr, pattern *Pattern) *Pattern {
	return &Pattern{
		TreeNode: &TreeNode{
			OpenParen:  newOpenParen(),
			Name:       name,
			Comma:      newComma(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

func NewLeafNode(expr *expr.Expr) *Pattern {
	return &Pattern{
		LeafNode: &LeafNode{
			OpenParen:  newOpenParen(),
			Expr:       expr,
			CloseParen: newCloseParen(),
		},
	}
}

func NewConcat(left, right *Pattern) *Pattern {
	return &Pattern{
		Concat: &Concat{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewOr(left, right *Pattern) *Pattern {
	return &Pattern{
		Or: &Or{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewAnd(left, right *Pattern) *Pattern {
	return &Pattern{
		And: &And{
			OpenParen:    newOpenParen(),
			LeftPattern:  left,
			Comma:        newComma(),
			RightPattern: right,
			CloseParen:   newCloseParen(),
		},
	}
}

func NewZeroOrMore(pattern *Pattern) *Pattern {
	return &Pattern{
		ZeroOrMore: &ZeroOrMore{
			OpenParen:  newOpenParen(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

func NewReference(name string) *Pattern {
	return &Pattern{
		Reference: &Reference{
			OpenParen:  newOpenParen(),
			Name:       name,
			CloseParen: newCloseParen(),
		},
	}
}

func NewNot(pattern *Pattern) *Pattern {
	return &Pattern{
		Not: &Not{
			OpenParen:  newOpenParen(),
			Pattern:    pattern,
			CloseParen: newCloseParen(),
		},
	}
}

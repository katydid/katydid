package parser

import (
	"github.com/katydid/katydid/expr/ast"
	. "github.com/katydid/katydid/lang/ast"
	"github.com/katydid/katydid/lang/token"
	"github.com/katydid/katydid/types"
)

func newString(v interface{}) string {
	t := v.(*token.Token)
	return string(t.Lit)
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Grammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Grammar : PatternDecls	<< &Grammar{X[0].([]*PatternDecl), nil}, nil >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].([]*PatternDecl), nil}, nil
		},
	},
	ProdTabEntry{
		String: `Grammar : PatternDecls Space	<< &Grammar{X[0].([]*PatternDecl), X[1].(*expr.Space)}, nil >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].([]*PatternDecl), X[1].(*expr.Space)}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecls : PatternDecl	<< []*PatternDecl{X[0].(*PatternDecl)}, nil >>`,
		Id:         "PatternDecls",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*PatternDecl{X[0].(*PatternDecl)}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecls : PatternDecls PatternDecl	<< append(X[0].([]*PatternDecl), X[1].(*PatternDecl)), nil >>`,
		Id:         "PatternDecls",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*PatternDecl), X[1].(*PatternDecl)), nil
		},
	},
	ProdTabEntry{
		String: `PatternDecl : Space id Equal Pattern	<< &PatternDecl{
      Before: X[0].(*expr.Space),
      Name: newString(X[1]),
      Equal: X[2].(*expr.Keyword),
      Pattern: X[3].(*Pattern),
    }, nil >>`,
		Id:         "PatternDecl",
		NTType:     3,
		Index:      5,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &PatternDecl{
				Before:  X[0].(*expr.Space),
				Name:    newString(X[1]),
				Equal:   X[2].(*expr.Keyword),
				Pattern: X[3].(*Pattern),
			}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecl : id Equal Pattern	<< &PatternDecl{
      Name: newString(X[0]),
      Equal: X[1].(*expr.Keyword),
      Pattern: X[2].(*Pattern),
    }, nil >>`,
		Id:         "PatternDecl",
		NTType:     3,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &PatternDecl{
				Name:    newString(X[0]),
				Equal:   X[1].(*expr.Keyword),
				Pattern: X[2].(*Pattern),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Empty"	<< &Pattern{Empty: &Empty{X[0].(*expr.Space)}}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Empty: &Empty{X[0].(*expr.Space)}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Empty"	<< &Pattern{Empty: &Empty{}}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Empty: &Empty{}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "EmptySet"	<< &Pattern{EmptySet: &EmptySet{X[0].(*expr.Space)}}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{EmptySet: &EmptySet{X[0].(*expr.Space)}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "EmptySet"	<< &Pattern{EmptySet: &EmptySet{}}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{EmptySet: &EmptySet{}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "TreeNode" OpenParen Expr Comma Pattern CloseParen	<< &Pattern{TreeNode: &TreeNode{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      Name: X[3].(*expr.Expr),
      Comma: X[4].(*expr.Keyword),
      Pattern: X[5].(*Pattern),
      CloseParen: X[6].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      11,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{TreeNode: &TreeNode{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				Name:       X[3].(*expr.Expr),
				Comma:      X[4].(*expr.Keyword),
				Pattern:    X[5].(*Pattern),
				CloseParen: X[6].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "TreeNode" OpenParen Expr Comma Pattern CloseParen	<< &Pattern{TreeNode: &TreeNode{
      OpenParen: X[1].(*expr.Keyword),
      Name: X[2].(*expr.Expr),
      Comma: X[3].(*expr.Keyword),
      Pattern: X[4].(*Pattern),
      CloseParen: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      12,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{TreeNode: &TreeNode{
				OpenParen:  X[1].(*expr.Keyword),
				Name:       X[2].(*expr.Expr),
				Comma:      X[3].(*expr.Keyword),
				Pattern:    X[4].(*Pattern),
				CloseParen: X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "LeafNode" OpenParen Expr CloseParen	<< &Pattern{LeafNode: &LeafNode{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      Expr: X[3].(*expr.Expr),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      13,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{LeafNode: &LeafNode{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				Expr:       X[3].(*expr.Expr),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "LeafNode" OpenParen Expr CloseParen	<< &Pattern{LeafNode: &LeafNode{
      OpenParen: X[1].(*expr.Keyword),
      Expr: X[2].(*expr.Expr),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{LeafNode: &LeafNode{
				OpenParen:  X[1].(*expr.Keyword),
				Expr:       X[2].(*expr.Expr),
				CloseParen: X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Concat" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{Concat: &Concat{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      LeftPattern: X[3].(*Pattern),
      Comma: X[4].(*expr.Keyword),
      RightPattern: X[5].(*Pattern),
      CloseParen: X[6].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      15,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Concat: &Concat{
				Before:       X[0].(*expr.Space),
				OpenParen:    X[2].(*expr.Keyword),
				LeftPattern:  X[3].(*Pattern),
				Comma:        X[4].(*expr.Keyword),
				RightPattern: X[5].(*Pattern),
				CloseParen:   X[6].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Concat" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{Concat: &Concat{
      OpenParen: X[1].(*expr.Keyword),
      LeftPattern: X[2].(*Pattern),
      Comma: X[3].(*expr.Keyword),
      RightPattern: X[4].(*Pattern),
      CloseParen: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      16,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Concat: &Concat{
				OpenParen:    X[1].(*expr.Keyword),
				LeftPattern:  X[2].(*Pattern),
				Comma:        X[3].(*expr.Keyword),
				RightPattern: X[4].(*Pattern),
				CloseParen:   X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Or" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{Or: &Or{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      LeftPattern: X[3].(*Pattern),
      Comma: X[4].(*expr.Keyword),
      RightPattern: X[5].(*Pattern),
      CloseParen: X[6].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      17,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Or: &Or{
				Before:       X[0].(*expr.Space),
				OpenParen:    X[2].(*expr.Keyword),
				LeftPattern:  X[3].(*Pattern),
				Comma:        X[4].(*expr.Keyword),
				RightPattern: X[5].(*Pattern),
				CloseParen:   X[6].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Or" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{Or: &Or{
      OpenParen: X[1].(*expr.Keyword),
      LeftPattern: X[2].(*Pattern),
      Comma: X[3].(*expr.Keyword),
      RightPattern: X[4].(*Pattern),
      CloseParen: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      18,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Or: &Or{
				OpenParen:    X[1].(*expr.Keyword),
				LeftPattern:  X[2].(*Pattern),
				Comma:        X[3].(*expr.Keyword),
				RightPattern: X[4].(*Pattern),
				CloseParen:   X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "And" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{And: &And{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      LeftPattern: X[3].(*Pattern),
      Comma: X[4].(*expr.Keyword),
      RightPattern: X[5].(*Pattern),
      CloseParen: X[6].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      19,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{And: &And{
				Before:       X[0].(*expr.Space),
				OpenParen:    X[2].(*expr.Keyword),
				LeftPattern:  X[3].(*Pattern),
				Comma:        X[4].(*expr.Keyword),
				RightPattern: X[5].(*Pattern),
				CloseParen:   X[6].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "And" OpenParen Pattern Comma Pattern CloseParen	<< &Pattern{And: &And{
      OpenParen: X[1].(*expr.Keyword),
      LeftPattern: X[2].(*Pattern),
      Comma: X[3].(*expr.Keyword),
      RightPattern: X[4].(*Pattern),
      CloseParen: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      20,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{And: &And{
				OpenParen:    X[1].(*expr.Keyword),
				LeftPattern:  X[2].(*Pattern),
				Comma:        X[3].(*expr.Keyword),
				RightPattern: X[4].(*Pattern),
				CloseParen:   X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "ZeroOrMore" OpenParen Pattern CloseParen	<< &Pattern{ZeroOrMore: &ZeroOrMore{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      Pattern: X[3].(*Pattern),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      21,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{ZeroOrMore: &ZeroOrMore{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				Pattern:    X[3].(*Pattern),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "ZeroOrMore" OpenParen Pattern CloseParen	<< &Pattern{ZeroOrMore: &ZeroOrMore{
      OpenParen: X[1].(*expr.Keyword),
      Pattern: X[2].(*Pattern),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      22,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{ZeroOrMore: &ZeroOrMore{
				OpenParen:  X[1].(*expr.Keyword),
				Pattern:    X[2].(*Pattern),
				CloseParen: X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Reference" OpenParen Space id CloseParen	<< &Pattern{Reference: &Reference{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      BeforeName: X[3].(*expr.Space),
      Name: newString(X[4]),
      CloseParen: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      23,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Reference: &Reference{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				BeforeName: X[3].(*expr.Space),
				Name:       newString(X[4]),
				CloseParen: X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Reference" OpenParen Space id CloseParen	<< &Pattern{Reference: &Reference{
      OpenParen: X[1].(*expr.Keyword),
      BeforeName: X[2].(*expr.Space),
      Name: newString(X[3]),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      24,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Reference: &Reference{
				OpenParen:  X[1].(*expr.Keyword),
				BeforeName: X[2].(*expr.Space),
				Name:       newString(X[3]),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Reference" OpenParen id CloseParen	<< &Pattern{Reference: &Reference{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      Name: newString(X[3]),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      25,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Reference: &Reference{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				Name:       newString(X[3]),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Reference" OpenParen id CloseParen	<< &Pattern{Reference: &Reference{
      OpenParen: X[1].(*expr.Keyword),
      Name: newString(X[2]),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      26,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Reference: &Reference{
				OpenParen:  X[1].(*expr.Keyword),
				Name:       newString(X[2]),
				CloseParen: X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Space "Not" OpenParen Pattern CloseParen	<< &Pattern{Not: &Not{
      Before: X[0].(*expr.Space),
      OpenParen: X[2].(*expr.Keyword),
      Pattern: X[3].(*Pattern),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      27,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Not: &Not{
				Before:     X[0].(*expr.Space),
				OpenParen:  X[2].(*expr.Keyword),
				Pattern:    X[3].(*Pattern),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : "Not" OpenParen Pattern CloseParen	<< &Pattern{Not: &Not{
      OpenParen: X[1].(*expr.Keyword),
      Pattern: X[2].(*Pattern),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     4,
		Index:      28,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Not: &Not{
				OpenParen:  X[1].(*expr.Keyword),
				Pattern:    X[2].(*Pattern),
				CloseParen: X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     5,
		Index:      29,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     5,
		Index:      30,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     5,
		Index:      31,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     5,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     6,
		Index:      33,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     6,
		Index:      34,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     6,
		Index:      35,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     6,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*expr.Expr{X[0].(*expr.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     7,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*expr.Expr{X[0].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     7,
		Index:      38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     8,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &expr.Expr{Function: X[0].(*expr.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     8,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Function: X[0].(*expr.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &expr.Expr{List: X[0].(*expr.List)}, nil >>`,
		Id:         "Expr",
		NTType:     8,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{List: X[0].(*expr.List)}, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int"	<< types.LIST_INT, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint"	<< types.LIST_UINT, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     9,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     10,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< expr.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     10,
		Index:      49,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Terminal : Bool	<< expr.NewBoolTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBoolTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : int_lit	<< expr.NewIntTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewIntTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : uint_lit	<< expr.NewUintTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUintTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : double_lit	<< expr.NewDoubleTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewDoubleTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : string_lit	<< expr.NewStringTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewStringTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_lit	<< expr.NewBytesTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBytesTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< expr.NewVariableTerminal(types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int_var	<< expr.NewVariableTerminal(types.SINGLE_INT) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint_var	<< expr.NewVariableTerminal(types.SINGLE_UINT) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< expr.NewVariableTerminal(types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< expr.NewVariableTerminal(types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< expr.NewVariableTerminal(types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     11,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     12,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     12,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< &expr.Keyword{Value: "="}, nil >>`,
		Id:         "Equal",
		NTType:     13,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "="}, nil
		},
	},
	ProdTabEntry{
		String: `Equal : Space "="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "=",
    }, nil >>`,
		Id:         "Equal",
		NTType:     13,
		Index:      65,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : "("	<< &expr.Keyword{Value: "("}, nil >>`,
		Id:         "OpenParen",
		NTType:     14,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "("}, nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : Space "("	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "(",
    }, nil >>`,
		Id:         "OpenParen",
		NTType:     14,
		Index:      67,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "(",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : ")"	<< &expr.Keyword{Value: ")"}, nil >>`,
		Id:         "CloseParen",
		NTType:     15,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ")"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : Space ")"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ")",
    }, nil >>`,
		Id:         "CloseParen",
		NTType:     15,
		Index:      69,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ")",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : "{"	<< &expr.Keyword{Value: "{"}, nil >>`,
		Id:         "OpenCurly",
		NTType:     16,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "{"}, nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : Space "{"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "{",
    }, nil >>`,
		Id:         "OpenCurly",
		NTType:     16,
		Index:      71,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "{",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : "}"	<< &expr.Keyword{Value: "}"}, nil >>`,
		Id:         "CloseCurly",
		NTType:     17,
		Index:      72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "}"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : Space "}"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "}",
    }, nil >>`,
		Id:         "CloseCurly",
		NTType:     17,
		Index:      73,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "}",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Comma : ","	<< &expr.Keyword{Value: ","}, nil >>`,
		Id:         "Comma",
		NTType:     18,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ","}, nil
		},
	},
	ProdTabEntry{
		String: `Comma : Space ","	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ",",
    }, nil >>`,
		Id:         "Comma",
		NTType:     18,
		Index:      75,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ",",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Space : Space space	<< expr.AppendSpace(X[0], newString(X[1])), nil >>`,
		Id:         "Space",
		NTType:     19,
		Index:      76,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.AppendSpace(X[0], newString(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< &expr.Space{Space: []string{newString(X[0])}}, nil >>`,
		Id:         "Space",
		NTType:     19,
		Index:      77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Space{Space: []string{newString(X[0])}}, nil
		},
	},
}

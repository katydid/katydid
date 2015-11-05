package parser

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/ast"
	. "github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/token"
	"github.com/katydid/katydid/types"
	"strconv"
)

func newString(v interface{}) string {
	t := v.(*token.Token)
	return string(t.Lit)
}

func unquote(s1 string) string {
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
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
		String: `S' : AllGrammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllGrammar : Grammar	<<  >>`,
		Id:         "AllGrammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllGrammar : GrammarWithEndingSpace	<<  >>`,
		Id:         "AllGrammar",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Grammar : Pattern	<< &Grammar{X[0].(*Pattern), nil, nil}, nil >>`,
		Id:         "Grammar",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].(*Pattern), nil, nil}, nil
		},
	},
	ProdTabEntry{
		String: `Grammar : Pattern PatternDecls	<< &Grammar{X[0].(*Pattern), X[1].([]*PatternDecl), nil}, nil >>`,
		Id:         "Grammar",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].(*Pattern), X[1].([]*PatternDecl), nil}, nil
		},
	},
	ProdTabEntry{
		String: `Grammar : PatternDecls	<< &Grammar{nil, X[0].([]*PatternDecl), nil}, nil >>`,
		Id:         "Grammar",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{nil, X[0].([]*PatternDecl), nil}, nil
		},
	},
	ProdTabEntry{
		String: `GrammarWithEndingSpace : Pattern Space	<< &Grammar{X[0].(*Pattern), nil, X[1].(*expr.Space)}, nil >>`,
		Id:         "GrammarWithEndingSpace",
		NTType:     3,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].(*Pattern), nil, X[1].(*expr.Space)}, nil
		},
	},
	ProdTabEntry{
		String: `GrammarWithEndingSpace : Pattern PatternDecls Space	<< &Grammar{X[0].(*Pattern), X[1].([]*PatternDecl), X[2].(*expr.Space)}, nil >>`,
		Id:         "GrammarWithEndingSpace",
		NTType:     3,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{X[0].(*Pattern), X[1].([]*PatternDecl), X[2].(*expr.Space)}, nil
		},
	},
	ProdTabEntry{
		String: `GrammarWithEndingSpace : PatternDecls Space	<< &Grammar{nil, X[0].([]*PatternDecl), X[1].(*expr.Space)}, nil >>`,
		Id:         "GrammarWithEndingSpace",
		NTType:     3,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Grammar{nil, X[0].([]*PatternDecl), X[1].(*expr.Space)}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecls : PatternDecl	<< []*PatternDecl{X[0].(*PatternDecl)}, nil >>`,
		Id:         "PatternDecls",
		NTType:     4,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*PatternDecl{X[0].(*PatternDecl)}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecls : PatternDecls PatternDecl	<< append(X[0].([]*PatternDecl), X[1].(*PatternDecl)), nil >>`,
		Id:         "PatternDecls",
		NTType:     4,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*PatternDecl), X[1].(*PatternDecl)), nil
		},
	},
	ProdTabEntry{
		String: `PatternDecl : Hash Space id Equal Pattern	<< &PatternDecl{
      Hash: X[0].(*expr.Keyword),
      Before: X[1].(*expr.Space),
      Name: newString(X[2]),
      Eq: X[3].(*expr.Keyword),
      Pattern: X[4].(*Pattern),
    }, nil >>`,
		Id:         "PatternDecl",
		NTType:     5,
		Index:      11,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &PatternDecl{
				Hash:    X[0].(*expr.Keyword),
				Before:  X[1].(*expr.Space),
				Name:    newString(X[2]),
				Eq:      X[3].(*expr.Keyword),
				Pattern: X[4].(*Pattern),
			}, nil
		},
	},
	ProdTabEntry{
		String: `PatternDecl : Hash id Equal Pattern	<< &PatternDecl{
      Hash: X[0].(*expr.Keyword),
      Name: newString(X[1]),
      Eq: X[2].(*expr.Keyword),
      Pattern: X[3].(*Pattern),
    }, nil >>`,
		Id:         "PatternDecl",
		NTType:     5,
		Index:      12,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &PatternDecl{
				Hash:    X[0].(*expr.Keyword),
				Name:    newString(X[1]),
				Eq:      X[2].(*expr.Keyword),
				Pattern: X[3].(*Pattern),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Name : Space Literal	<< NewSDTName(X[0].(*expr.Space), X[1].(*expr.Terminal)), nil >>`,
		Id:         "Name",
		NTType:     6,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewSDTName(X[0].(*expr.Space), X[1].(*expr.Terminal)), nil
		},
	},
	ProdTabEntry{
		String: `Name : Literal	<< NewSDTName(nil, X[0].(*expr.Terminal)), nil >>`,
		Id:         "Name",
		NTType:     6,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewSDTName(nil, X[0].(*expr.Terminal)), nil
		},
	},
	ProdTabEntry{
		String: `Name : Space id	<< NewSDTName(X[0].(*expr.Space), &expr.Terminal{StringValue: proto.String(newString(X[1]))}), nil >>`,
		Id:         "Name",
		NTType:     6,
		Index:      15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewSDTName(X[0].(*expr.Space), &expr.Terminal{StringValue: proto.String(newString(X[1]))}), nil
		},
	},
	ProdTabEntry{
		String: `Name : id	<< NewSDTName(nil, &expr.Terminal{StringValue: proto.String(newString(X[0]))}), nil >>`,
		Id:         "Name",
		NTType:     6,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewSDTName(nil, &expr.Terminal{StringValue: proto.String(newString(X[0]))}), nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Underscore	<< &NameExpr{AnyName: &AnyName{X[0].(*expr.Keyword)}}, nil >>`,
		Id:         "NameExpr",
		NTType:     7,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &NameExpr{AnyName: &AnyName{X[0].(*expr.Keyword)}}, nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Name	<<  >>`,
		Id:         "NameExpr",
		NTType:     7,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Exclamation OpenParen NameExpr CloseParen	<< &NameExpr{AnyNameExcept: &AnyNameExcept{
      Exclamation: X[0].(*expr.Keyword),
      OpenParen: X[1].(*expr.Keyword),
      Except: X[2].(*NameExpr),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "NameExpr",
		NTType:     7,
		Index:      19,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &NameExpr{AnyNameExcept: &AnyNameExcept{
				Exclamation: X[0].(*expr.Keyword),
				OpenParen:   X[1].(*expr.Keyword),
				Except:      X[2].(*NameExpr),
				CloseParen:  X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : StartNameChoice	<<  >>`,
		Id:         "NameExpr",
		NTType:     7,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StartNameChoice : OpenParen NameExpr Pipe ContinueNameChoice CloseParen	<< &NameExpr{NameChoice: &NameChoice{
      OpenParen: X[0].(*expr.Keyword),
      Left: X[1].(*NameExpr),
      Pipe: X[2].(*expr.Keyword),
      Right: X[3].(*NameExpr),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartNameChoice",
		NTType:     8,
		Index:      21,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &NameExpr{NameChoice: &NameChoice{
				OpenParen:  X[0].(*expr.Keyword),
				Left:       X[1].(*NameExpr),
				Pipe:       X[2].(*expr.Keyword),
				Right:      X[3].(*NameExpr),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueNameChoice : NameExpr	<<  >>`,
		Id:         "ContinueNameChoice",
		NTType:     9,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueNameChoice : ContinueNameChoice Pipe NameExpr	<< &NameExpr{NameChoice: &NameChoice{
      Left: X[0].(*NameExpr),
      Pipe: X[1].(*expr.Keyword),
      Right: X[2].(*NameExpr),
    }}, nil >>`,
		Id:         "ContinueNameChoice",
		NTType:     9,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &NameExpr{NameChoice: &NameChoice{
				Left:  X[0].(*NameExpr),
				Pipe:  X[1].(*expr.Keyword),
				Right: X[2].(*NameExpr),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Empty : "<empty>"	<< &expr.Keyword{Value: "<empty>"}, nil >>`,
		Id:         "Empty",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<empty>"}, nil
		},
	},
	ProdTabEntry{
		String: `Empty : Space "<empty>"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<empty>",
    }, nil >>`,
		Id:         "Empty",
		NTType:     10,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<empty>",
			}, nil
		},
	},
	ProdTabEntry{
		String: `EmptySet : "<emptyset>"	<< &expr.Keyword{Value: "<emptyset>"}, nil >>`,
		Id:         "EmptySet",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<emptyset>"}, nil
		},
	},
	ProdTabEntry{
		String: `EmptySet : Space "<emptyset>"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<emptyset>",
    }, nil >>`,
		Id:         "EmptySet",
		NTType:     11,
		Index:      27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<emptyset>",
			}, nil
		},
	},
	ProdTabEntry{
		String: `DepthPattern : StartConcat	<<  >>`,
		Id:         "DepthPattern",
		NTType:     12,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DepthPattern : StartInterleave	<<  >>`,
		Id:         "DepthPattern",
		NTType:     12,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DepthPattern : Dot TreeNode	<< &Pattern{WithSomeTreeNode: &WithSomeTreeNode{
      Dot: X[0].(*expr.Keyword),
      Pattern: X[1].(*Pattern),
    }}, nil >>`,
		Id:         "DepthPattern",
		NTType:     12,
		Index:      30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{WithSomeTreeNode: &WithSomeTreeNode{
				Dot:     X[0].(*expr.Keyword),
				Pattern: X[1].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `DepthPattern : RightArrow Function	<< &Pattern{LeafNode: &LeafNode{
        RightArrow: X[0].(*expr.Keyword),
        Expr: &expr.Expr{Function: X[1].(*expr.Function)},
    }}, nil >>`,
		Id:         "DepthPattern",
		NTType:     12,
		Index:      31,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{LeafNode: &LeafNode{
				RightArrow: X[0].(*expr.Keyword),
				Expr:       &expr.Expr{Function: X[1].(*expr.Function)},
			}}, nil
		},
	},
	ProdTabEntry{
		String: `DepthPattern : BuiltIn	<< &Pattern{LeafNode: &LeafNode{
        Expr: &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)},
    }}, nil >>`,
		Id:         "DepthPattern",
		NTType:     12,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{LeafNode: &LeafNode{
				Expr: &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)},
			}}, nil
		},
	},
	ProdTabEntry{
		String: `TreeNode : NameExpr Colon Pattern	<< &Pattern{TreeNode: &TreeNode{
      Name: X[0].(*NameExpr),
      Colon: X[1].(*expr.Keyword),
      Pattern: X[2].(*Pattern),
    }}, nil >>`,
		Id:         "TreeNode",
		NTType:     13,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{TreeNode: &TreeNode{
				Name:    X[0].(*NameExpr),
				Colon:   X[1].(*expr.Keyword),
				Pattern: X[2].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `TreeNode : NameExpr DepthPattern	<< &Pattern{TreeNode: &TreeNode{
      Name: X[0].(*NameExpr),
      Pattern: X[1].(*Pattern),
    }}, nil >>`,
		Id:         "TreeNode",
		NTType:     13,
		Index:      34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{TreeNode: &TreeNode{
				Name:    X[0].(*NameExpr),
				Pattern: X[1].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Empty	<< &Pattern{Empty: &Empty{X[0].(*expr.Keyword)}}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Empty: &Empty{X[0].(*expr.Keyword)}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : EmptySet	<< &Pattern{EmptySet: &EmptySet{X[0].(*expr.Keyword)}}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{EmptySet: &EmptySet{X[0].(*expr.Keyword)}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Star	<< &Pattern{ZAny: &ZAny{Star: X[0].(*expr.Keyword)}}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{ZAny: &ZAny{Star: X[0].(*expr.Keyword)}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : TreeNode	<<  >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Pattern : DepthPattern	<<  >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Pattern : StartOr	<<  >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Pattern : StartAnd	<<  >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Pattern : OpenParen Pattern CloseParen Star	<< &Pattern{ZeroOrMore: &ZeroOrMore{
      OpenParen: X[0].(*expr.Keyword),
      Pattern: X[1].(*Pattern),
      CloseParen: X[2].(*expr.Keyword),
      Star: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      42,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{ZeroOrMore: &ZeroOrMore{
				OpenParen:  X[0].(*expr.Keyword),
				Pattern:    X[1].(*Pattern),
				CloseParen: X[2].(*expr.Keyword),
				Star:       X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : OpenParen Pattern CloseParen QuestionMark	<< &Pattern{Optional: &Optional{
      OpenParen: X[0].(*expr.Keyword),
      Pattern: X[1].(*Pattern),
      CloseParen: X[2].(*expr.Keyword),
      QuestionMark: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      43,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Optional: &Optional{
				OpenParen:    X[0].(*expr.Keyword),
				Pattern:      X[1].(*Pattern),
				CloseParen:   X[2].(*expr.Keyword),
				QuestionMark: X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : At id	<< &Pattern{Reference: &Reference{
      At: X[0].(*expr.Keyword),
      Name: newString(X[1]),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      44,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Reference: &Reference{
				At:   X[0].(*expr.Keyword),
				Name: newString(X[1]),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Pattern : Exclamation OpenParen Pattern CloseParen	<< &Pattern{Not: &Not{
      Exclamation: X[0].(*expr.Keyword),
      OpenParen: X[1].(*expr.Keyword),
      Pattern: X[2].(*Pattern),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "Pattern",
		NTType:     14,
		Index:      45,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Not: &Not{
				Exclamation: X[0].(*expr.Keyword),
				OpenParen:   X[1].(*expr.Keyword),
				Pattern:     X[2].(*Pattern),
				CloseParen:  X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartConcat : OpenBracket Pattern Comma ContinueConcat CloseBracket	<< &Pattern{Concat: &Concat{
      OpenBracket: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      Comma: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      CloseBracket: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartConcat",
		NTType:     15,
		Index:      46,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Concat: &Concat{
				OpenBracket:  X[0].(*expr.Keyword),
				LeftPattern:  X[1].(*Pattern),
				Comma:        X[2].(*expr.Keyword),
				RightPattern: X[3].(*Pattern),
				CloseBracket: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartConcat : OpenBracket Pattern Comma ContinueConcat Comma CloseBracket	<< &Pattern{Concat: &Concat{
      OpenBracket: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      Comma: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      ExtraComma: X[4].(*expr.Keyword),
      CloseBracket: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartConcat",
		NTType:     15,
		Index:      47,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Concat: &Concat{
				OpenBracket:  X[0].(*expr.Keyword),
				LeftPattern:  X[1].(*Pattern),
				Comma:        X[2].(*expr.Keyword),
				RightPattern: X[3].(*Pattern),
				ExtraComma:   X[4].(*expr.Keyword),
				CloseBracket: X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueConcat : Pattern	<<  >>`,
		Id:         "ContinueConcat",
		NTType:     16,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueConcat : ContinueConcat Comma Pattern	<< &Pattern{Concat: &Concat{
      LeftPattern: X[0].(*Pattern),
      Comma: X[1].(*expr.Keyword),
      RightPattern: X[2].(*Pattern),
    }}, nil >>`,
		Id:         "ContinueConcat",
		NTType:     16,
		Index:      49,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Concat: &Concat{
				LeftPattern:  X[0].(*Pattern),
				Comma:        X[1].(*expr.Keyword),
				RightPattern: X[2].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartOr : OpenParen Pattern Pipe ContinueOr CloseParen	<< &Pattern{Or: &Or{
      OpenParen: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      Pipe: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartOr",
		NTType:     17,
		Index:      50,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Or: &Or{
				OpenParen:    X[0].(*expr.Keyword),
				LeftPattern:  X[1].(*Pattern),
				Pipe:         X[2].(*expr.Keyword),
				RightPattern: X[3].(*Pattern),
				CloseParen:   X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueOr : Pattern	<<  >>`,
		Id:         "ContinueOr",
		NTType:     18,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueOr : ContinueOr Pipe Pattern	<< &Pattern{Or: &Or{
      LeftPattern: X[0].(*Pattern),
      Pipe: X[1].(*expr.Keyword),
      RightPattern: X[2].(*Pattern),
    }}, nil >>`,
		Id:         "ContinueOr",
		NTType:     18,
		Index:      52,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Or: &Or{
				LeftPattern:  X[0].(*Pattern),
				Pipe:         X[1].(*expr.Keyword),
				RightPattern: X[2].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartAnd : OpenParen Pattern Ampersand ContinueAnd CloseParen	<< &Pattern{And: &And{
      OpenParen: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      Ampersand: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartAnd",
		NTType:     19,
		Index:      53,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{And: &And{
				OpenParen:    X[0].(*expr.Keyword),
				LeftPattern:  X[1].(*Pattern),
				Ampersand:    X[2].(*expr.Keyword),
				RightPattern: X[3].(*Pattern),
				CloseParen:   X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueAnd : Pattern	<<  >>`,
		Id:         "ContinueAnd",
		NTType:     20,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueAnd : ContinueAnd Ampersand Pattern	<< &Pattern{And: &And{
      LeftPattern: X[0].(*Pattern),
      Ampersand: X[1].(*expr.Keyword),
      RightPattern: X[2].(*Pattern),
    }}, nil >>`,
		Id:         "ContinueAnd",
		NTType:     20,
		Index:      55,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{And: &And{
				LeftPattern:  X[0].(*Pattern),
				Ampersand:    X[1].(*expr.Keyword),
				RightPattern: X[2].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartInterleave : OpenCurly Pattern SemiColon ContinueInterleave CloseCurly	<< &Pattern{Interleave: &Interleave{
      OpenCurly: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      SemiColon: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      CloseCurly: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartInterleave",
		NTType:     21,
		Index:      56,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Interleave: &Interleave{
				OpenCurly:    X[0].(*expr.Keyword),
				LeftPattern:  X[1].(*Pattern),
				SemiColon:    X[2].(*expr.Keyword),
				RightPattern: X[3].(*Pattern),
				CloseCurly:   X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `StartInterleave : OpenCurly Pattern SemiColon ContinueInterleave SemiColon CloseCurly	<< &Pattern{Interleave: &Interleave{
      OpenCurly: X[0].(*expr.Keyword),
      LeftPattern: X[1].(*Pattern),
      SemiColon: X[2].(*expr.Keyword),
      RightPattern: X[3].(*Pattern),
      ExtraSemiColon: X[4].(*expr.Keyword),
      CloseCurly: X[5].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartInterleave",
		NTType:     21,
		Index:      57,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Interleave: &Interleave{
				OpenCurly:      X[0].(*expr.Keyword),
				LeftPattern:    X[1].(*Pattern),
				SemiColon:      X[2].(*expr.Keyword),
				RightPattern:   X[3].(*Pattern),
				ExtraSemiColon: X[4].(*expr.Keyword),
				CloseCurly:     X[5].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueInterleave : Pattern	<<  >>`,
		Id:         "ContinueInterleave",
		NTType:     22,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueInterleave : ContinueInterleave SemiColon Pattern	<< &Pattern{Interleave: &Interleave{
      LeftPattern: X[0].(*Pattern),
      SemiColon: X[1].(*expr.Keyword),
      RightPattern: X[2].(*Pattern),
    }}, nil >>`,
		Id:         "ContinueInterleave",
		NTType:     22,
		Index:      59,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Pattern{Interleave: &Interleave{
				LeftPattern:  X[0].(*Pattern),
				SemiColon:    X[1].(*expr.Keyword),
				RightPattern: X[2].(*Pattern),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : Expr Space	<<  >>`,
		Id:         "AllExpr",
		NTType:     23,
		Index:      60,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : Expr	<<  >>`,
		Id:         "AllExpr",
		NTType:     23,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : BuiltIn Space	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "AllExpr",
		NTType:     23,
		Index:      62,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : BuiltIn	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "AllExpr",
		NTType:     23,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     24,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &expr.Expr{Function: X[0].(*expr.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     24,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Function: X[0].(*expr.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &expr.Expr{List: X[0].(*expr.List)}, nil >>`,
		Id:         "Expr",
		NTType:     24,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{List: X[0].(*expr.List)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     25,
		Index:      67,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     25,
		Index:      68,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     25,
		Index:      69,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     25,
		Index:      70,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : EqualEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      71,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ExclamationEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      72,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      73,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      74,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      75,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      76,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : TildeEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      77,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : StarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      78,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : CaretEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      79,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : DollarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      80,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ColonColon Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     26,
		Index:      81,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     27,
		Index:      82,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     27,
		Index:      83,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     27,
		Index:      84,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     27,
		Index:      85,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*expr.Expr{X[0].(*expr.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     28,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*expr.Expr{X[0].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     28,
		Index:      87,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      88,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int"	<< types.LIST_INT, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint"	<< types.LIST_UINT, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     29,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     30,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< expr.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     30,
		Index:      95,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : Bool	<< expr.NewBoolTerminal(X[0]), nil >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBoolTerminal(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : int_lit	<< expr.NewIntTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewIntTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : uint_lit	<< expr.NewUintTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUintTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : double_lit	<< expr.NewDoubleTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewDoubleTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : string_lit	<< expr.NewStringTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewStringTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : bytes_lit	<< expr.NewBytesTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     31,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBytesTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : Literal	<<  >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< expr.NewVariableTerminal(types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      103,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int_var	<< expr.NewVariableTerminal(types.SINGLE_INT) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      104,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint_var	<< expr.NewVariableTerminal(types.SINGLE_UINT) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< expr.NewVariableTerminal(types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      106,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< expr.NewVariableTerminal(types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< expr.NewVariableTerminal(types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     32,
		Index:      108,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     33,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     33,
		Index:      110,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< &expr.Keyword{Value: "="}, nil >>`,
		Id:         "Equal",
		NTType:     34,
		Index:      111,
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
		NTType:     34,
		Index:      112,
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
		NTType:     35,
		Index:      113,
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
		NTType:     35,
		Index:      114,
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
		NTType:     36,
		Index:      115,
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
		NTType:     36,
		Index:      116,
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
		NTType:     37,
		Index:      117,
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
		NTType:     37,
		Index:      118,
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
		NTType:     38,
		Index:      119,
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
		NTType:     38,
		Index:      120,
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
		NTType:     39,
		Index:      121,
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
		NTType:     39,
		Index:      122,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ",",
			}, nil
		},
	},
	ProdTabEntry{
		String: `SemiColon : ";"	<< &expr.Keyword{Value: ";"}, nil >>`,
		Id:         "SemiColon",
		NTType:     40,
		Index:      123,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ";"}, nil
		},
	},
	ProdTabEntry{
		String: `SemiColon : Space ";"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ";",
    }, nil >>`,
		Id:         "SemiColon",
		NTType:     40,
		Index:      124,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ";",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Hash : "#"	<< &expr.Keyword{Value: "#"}, nil >>`,
		Id:         "Hash",
		NTType:     41,
		Index:      125,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "#"}, nil
		},
	},
	ProdTabEntry{
		String: `Hash : Space "#"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "#",
    }, nil >>`,
		Id:         "Hash",
		NTType:     41,
		Index:      126,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "#",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Ampersand : "&"	<< &expr.Keyword{Value: "&"}, nil >>`,
		Id:         "Ampersand",
		NTType:     42,
		Index:      127,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "&"}, nil
		},
	},
	ProdTabEntry{
		String: `Ampersand : Space "&"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "&",
    }, nil >>`,
		Id:         "Ampersand",
		NTType:     42,
		Index:      128,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "&",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Pipe : "|"	<< &expr.Keyword{Value: "|"}, nil >>`,
		Id:         "Pipe",
		NTType:     43,
		Index:      129,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "|"}, nil
		},
	},
	ProdTabEntry{
		String: `Pipe : Space "|"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "|",
    }, nil >>`,
		Id:         "Pipe",
		NTType:     43,
		Index:      130,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "|",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenBracket : "["	<< &expr.Keyword{Value: "["}, nil >>`,
		Id:         "OpenBracket",
		NTType:     44,
		Index:      131,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "["}, nil
		},
	},
	ProdTabEntry{
		String: `OpenBracket : Space "["	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "[",
    }, nil >>`,
		Id:         "OpenBracket",
		NTType:     44,
		Index:      132,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "[",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseBracket : "]"	<< &expr.Keyword{Value: "]"}, nil >>`,
		Id:         "CloseBracket",
		NTType:     45,
		Index:      133,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "]"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseBracket : Space "]"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "]",
    }, nil >>`,
		Id:         "CloseBracket",
		NTType:     45,
		Index:      134,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "]",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Colon : ":"	<< &expr.Keyword{Value: ":"}, nil >>`,
		Id:         "Colon",
		NTType:     46,
		Index:      135,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ":"}, nil
		},
	},
	ProdTabEntry{
		String: `Colon : Space ":"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ":",
    }, nil >>`,
		Id:         "Colon",
		NTType:     46,
		Index:      136,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ":",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Exclamation : "!"	<< &expr.Keyword{Value: "!"}, nil >>`,
		Id:         "Exclamation",
		NTType:     47,
		Index:      137,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "!"}, nil
		},
	},
	ProdTabEntry{
		String: `Exclamation : Space "!"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "!",
    }, nil >>`,
		Id:         "Exclamation",
		NTType:     47,
		Index:      138,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "!",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Star : "*"	<< &expr.Keyword{Value: "*"}, nil >>`,
		Id:         "Star",
		NTType:     48,
		Index:      139,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "*"}, nil
		},
	},
	ProdTabEntry{
		String: `Star : Space "*"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "*",
    }, nil >>`,
		Id:         "Star",
		NTType:     48,
		Index:      140,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "*",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Underscore : "_"	<< &expr.Keyword{Value: "_"}, nil >>`,
		Id:         "Underscore",
		NTType:     49,
		Index:      141,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "_"}, nil
		},
	},
	ProdTabEntry{
		String: `Underscore : Space "_"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "_",
    }, nil >>`,
		Id:         "Underscore",
		NTType:     49,
		Index:      142,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "_",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Tilde : "~"	<< &expr.Keyword{Value: "~"}, nil >>`,
		Id:         "Tilde",
		NTType:     50,
		Index:      143,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "~"}, nil
		},
	},
	ProdTabEntry{
		String: `Tilde : Space "~"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "~",
    }, nil >>`,
		Id:         "Tilde",
		NTType:     50,
		Index:      144,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "~",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Dot : "."	<< &expr.Keyword{Value: "."}, nil >>`,
		Id:         "Dot",
		NTType:     51,
		Index:      145,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "."}, nil
		},
	},
	ProdTabEntry{
		String: `Dot : Space "."	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ".",
    }, nil >>`,
		Id:         "Dot",
		NTType:     51,
		Index:      146,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ".",
			}, nil
		},
	},
	ProdTabEntry{
		String: `At : "@"	<< &expr.Keyword{Value: "@"}, nil >>`,
		Id:         "At",
		NTType:     52,
		Index:      147,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "@"}, nil
		},
	},
	ProdTabEntry{
		String: `At : Space "@"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "@",
    }, nil >>`,
		Id:         "At",
		NTType:     52,
		Index:      148,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "@",
			}, nil
		},
	},
	ProdTabEntry{
		String: `RightArrow : "->"	<< &expr.Keyword{Value: "->"}, nil >>`,
		Id:         "RightArrow",
		NTType:     53,
		Index:      149,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "->"}, nil
		},
	},
	ProdTabEntry{
		String: `RightArrow : Space "->"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "->",
    }, nil >>`,
		Id:         "RightArrow",
		NTType:     53,
		Index:      150,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "->",
			}, nil
		},
	},
	ProdTabEntry{
		String: `EqualEqual : "=="	<< &expr.Keyword{Value: "=="}, nil >>`,
		Id:         "EqualEqual",
		NTType:     54,
		Index:      151,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "=="}, nil
		},
	},
	ProdTabEntry{
		String: `EqualEqual : Space "=="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "==",
    }, nil >>`,
		Id:         "EqualEqual",
		NTType:     54,
		Index:      152,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "==",
			}, nil
		},
	},
	ProdTabEntry{
		String: `ExclamationEqual : "!="	<< &expr.Keyword{Value: "!="}, nil >>`,
		Id:         "ExclamationEqual",
		NTType:     55,
		Index:      153,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "!="}, nil
		},
	},
	ProdTabEntry{
		String: `ExclamationEqual : Space "!="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "!=",
    }, nil >>`,
		Id:         "ExclamationEqual",
		NTType:     55,
		Index:      154,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "!=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `LessThan : "<"	<< &expr.Keyword{Value: "<"}, nil >>`,
		Id:         "LessThan",
		NTType:     56,
		Index:      155,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<"}, nil
		},
	},
	ProdTabEntry{
		String: `LessThan : Space "<"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<",
    }, nil >>`,
		Id:         "LessThan",
		NTType:     56,
		Index:      156,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<",
			}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterThan : ">"	<< &expr.Keyword{Value: ">"}, nil >>`,
		Id:         "GreaterThan",
		NTType:     57,
		Index:      157,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ">"}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterThan : Space ">"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ">",
    }, nil >>`,
		Id:         "GreaterThan",
		NTType:     57,
		Index:      158,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ">",
			}, nil
		},
	},
	ProdTabEntry{
		String: `LessEqual : "<="	<< &expr.Keyword{Value: "<="}, nil >>`,
		Id:         "LessEqual",
		NTType:     58,
		Index:      159,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<="}, nil
		},
	},
	ProdTabEntry{
		String: `LessEqual : Space "<="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<=",
    }, nil >>`,
		Id:         "LessEqual",
		NTType:     58,
		Index:      160,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterEqual : ">="	<< &expr.Keyword{Value: ">="}, nil >>`,
		Id:         "GreaterEqual",
		NTType:     59,
		Index:      161,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ">="}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterEqual : Space ">="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ">=",
    }, nil >>`,
		Id:         "GreaterEqual",
		NTType:     59,
		Index:      162,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ">=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `TildeEqual : "~="	<< &expr.Keyword{Value: "~="}, nil >>`,
		Id:         "TildeEqual",
		NTType:     60,
		Index:      163,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "~="}, nil
		},
	},
	ProdTabEntry{
		String: `TildeEqual : Space "~="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "~=",
    }, nil >>`,
		Id:         "TildeEqual",
		NTType:     60,
		Index:      164,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "~=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `StarEqual : "*="	<< &expr.Keyword{Value: "*="}, nil >>`,
		Id:         "StarEqual",
		NTType:     61,
		Index:      165,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "*="}, nil
		},
	},
	ProdTabEntry{
		String: `StarEqual : Space "*="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "*=",
    }, nil >>`,
		Id:         "StarEqual",
		NTType:     61,
		Index:      166,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "*=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CaretEqual : "^="	<< &expr.Keyword{Value: "^="}, nil >>`,
		Id:         "CaretEqual",
		NTType:     62,
		Index:      167,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "^="}, nil
		},
	},
	ProdTabEntry{
		String: `CaretEqual : Space "^="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "^=",
    }, nil >>`,
		Id:         "CaretEqual",
		NTType:     62,
		Index:      168,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "^=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `DollarEqual : "$="	<< &expr.Keyword{Value: "$="}, nil >>`,
		Id:         "DollarEqual",
		NTType:     63,
		Index:      169,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "$="}, nil
		},
	},
	ProdTabEntry{
		String: `DollarEqual : Space "$="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "$=",
    }, nil >>`,
		Id:         "DollarEqual",
		NTType:     63,
		Index:      170,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "$=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `ColonColon : "::"	<< &expr.Keyword{Value: "::"}, nil >>`,
		Id:         "ColonColon",
		NTType:     64,
		Index:      171,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "::"}, nil
		},
	},
	ProdTabEntry{
		String: `ColonColon : Space "::"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "::",
    }, nil >>`,
		Id:         "ColonColon",
		NTType:     64,
		Index:      172,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "::",
			}, nil
		},
	},
	ProdTabEntry{
		String: `QuestionMark : "?"	<< &expr.Keyword{Value: "?"}, nil >>`,
		Id:         "QuestionMark",
		NTType:     65,
		Index:      173,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "?"}, nil
		},
	},
	ProdTabEntry{
		String: `QuestionMark : Space "?"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "?",
    }, nil >>`,
		Id:         "QuestionMark",
		NTType:     65,
		Index:      174,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "?",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Space : Space space	<< expr.AppendSpace(X[0], newString(X[1])), nil >>`,
		Id:         "Space",
		NTType:     66,
		Index:      175,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.AppendSpace(X[0], newString(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< &expr.Space{Space: []string{newString(X[0])}}, nil >>`,
		Id:         "Space",
		NTType:     66,
		Index:      176,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Space{Space: []string{newString(X[0])}}, nil
		},
	},
}

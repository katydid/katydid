package parser

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/token"
	"github.com/katydid/katydid/expr/types"
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
		String: `S' : AllExpr	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : Expr Space	<<  >>`,
		Id:         "AllExpr",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : Expr	<<  >>`,
		Id:         "AllExpr",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : BuiltIn Space	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "AllExpr",
		NTType:     1,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `AllExpr : BuiltIn	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "AllExpr",
		NTType:     1,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &expr.Expr{Function: X[0].(*expr.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Function: X[0].(*expr.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &expr.Expr{List: X[0].(*expr.List)}, nil >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{List: X[0].(*expr.List)}, nil
		},
	},
	ProdTabEntry{
		String: `Name : Space Literal	<< expr.NewSDTName(X[0].(*expr.Space), X[1].(*expr.Terminal)), nil >>`,
		Id:         "Name",
		NTType:     3,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewSDTName(X[0].(*expr.Space), X[1].(*expr.Terminal)), nil
		},
	},
	ProdTabEntry{
		String: `Name : Literal	<< expr.NewSDTName(nil, X[0].(*expr.Terminal)), nil >>`,
		Id:         "Name",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewSDTName(nil, X[0].(*expr.Terminal)), nil
		},
	},
	ProdTabEntry{
		String: `Name : Space id	<< expr.NewSDTName(X[0].(*expr.Space), &expr.Terminal{StringValue: proto.String(newString(X[1]))}), nil >>`,
		Id:         "Name",
		NTType:     3,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewSDTName(X[0].(*expr.Space), &expr.Terminal{StringValue: proto.String(newString(X[1]))}), nil
		},
	},
	ProdTabEntry{
		String: `Name : id	<< expr.NewSDTName(nil, &expr.Terminal{StringValue: proto.String(newString(X[0]))}), nil >>`,
		Id:         "Name",
		NTType:     3,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewSDTName(nil, &expr.Terminal{StringValue: proto.String(newString(X[0]))}), nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Underscore	<< &expr.NameExpr{AnyName: &expr.AnyName{X[0].(*expr.Keyword)}}, nil >>`,
		Id:         "NameExpr",
		NTType:     4,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.NameExpr{AnyName: &expr.AnyName{X[0].(*expr.Keyword)}}, nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Name	<<  >>`,
		Id:         "NameExpr",
		NTType:     4,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : Exclamation OpenParen NameExpr CloseParen	<< &expr.NameExpr{AnyNameExcept: &expr.AnyNameExcept{
      Exclamation: X[0].(*expr.Keyword),
      OpenParen: X[1].(*expr.Keyword),
      Except: X[2].(*expr.NameExpr),
      CloseParen: X[3].(*expr.Keyword),
    }}, nil >>`,
		Id:         "NameExpr",
		NTType:     4,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.NameExpr{AnyNameExcept: &expr.AnyNameExcept{
				Exclamation: X[0].(*expr.Keyword),
				OpenParen:   X[1].(*expr.Keyword),
				Except:      X[2].(*expr.NameExpr),
				CloseParen:  X[3].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `NameExpr : StartNameChoice	<<  >>`,
		Id:         "NameExpr",
		NTType:     4,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StartNameChoice : OpenParen NameExpr Pipe ContinueNameChoice CloseParen	<< &expr.NameExpr{NameChoice: &expr.NameChoice{
      OpenParen: X[0].(*expr.Keyword),
      Left: X[1].(*expr.NameExpr),
      Pipe: X[2].(*expr.Keyword),
      Right: X[3].(*expr.NameExpr),
      CloseParen: X[4].(*expr.Keyword),
    }}, nil >>`,
		Id:         "StartNameChoice",
		NTType:     5,
		Index:      16,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.NameExpr{NameChoice: &expr.NameChoice{
				OpenParen:  X[0].(*expr.Keyword),
				Left:       X[1].(*expr.NameExpr),
				Pipe:       X[2].(*expr.Keyword),
				Right:      X[3].(*expr.NameExpr),
				CloseParen: X[4].(*expr.Keyword),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueNameChoice : NameExpr	<<  >>`,
		Id:         "ContinueNameChoice",
		NTType:     6,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContinueNameChoice : ContinueNameChoice Pipe NameExpr	<< &expr.NameExpr{NameChoice: &expr.NameChoice{
      Left: X[0].(*expr.NameExpr),
      Pipe: X[1].(*expr.Keyword),
      Right: X[2].(*expr.NameExpr),
    }}, nil >>`,
		Id:         "ContinueNameChoice",
		NTType:     6,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.NameExpr{NameChoice: &expr.NameChoice{
				Left:  X[0].(*expr.NameExpr),
				Pipe:  X[1].(*expr.Keyword),
				Right: X[2].(*expr.NameExpr),
			}}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     7,
		Index:      19,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     7,
		Index:      20,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     7,
		Index:      21,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     7,
		Index:      22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : EqualEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ExclamationEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : TildeEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      29,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : StarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : CaretEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      31,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : DollarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ColonColon Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     8,
		Index:      33,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     9,
		Index:      34,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     9,
		Index:      35,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     9,
		Index:      36,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     9,
		Index:      37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*expr.Expr{X[0].(*expr.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     10,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*expr.Expr{X[0].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     10,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int"	<< types.LIST_INT, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint"	<< types.LIST_UINT, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     11,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     12,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< expr.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     12,
		Index:      47,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : Bool	<< expr.NewBoolTerminal(X[0]), nil >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBoolTerminal(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : int_lit	<< expr.NewIntTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewIntTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : uint_lit	<< expr.NewUintTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUintTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : double_lit	<< expr.NewDoubleTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewDoubleTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : string_lit	<< expr.NewStringTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewStringTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : bytes_lit	<< expr.NewBytesTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     13,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBytesTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : Literal	<<  >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< expr.NewVariableTerminal(types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int_var	<< expr.NewVariableTerminal(types.SINGLE_INT) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint_var	<< expr.NewVariableTerminal(types.SINGLE_UINT) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< expr.NewVariableTerminal(types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< expr.NewVariableTerminal(types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< expr.NewVariableTerminal(types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     14,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     15,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     15,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< &expr.Keyword{Value: "="}, nil >>`,
		Id:         "Equal",
		NTType:     16,
		Index:      63,
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
		NTType:     16,
		Index:      64,
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
		NTType:     17,
		Index:      65,
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
		NTType:     17,
		Index:      66,
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
		NTType:     18,
		Index:      67,
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
		NTType:     18,
		Index:      68,
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
		NTType:     19,
		Index:      69,
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
		NTType:     19,
		Index:      70,
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
		NTType:     20,
		Index:      71,
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
		NTType:     20,
		Index:      72,
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
		NTType:     21,
		Index:      73,
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
		NTType:     21,
		Index:      74,
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
		NTType:     22,
		Index:      75,
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
		NTType:     22,
		Index:      76,
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
		NTType:     23,
		Index:      77,
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
		NTType:     23,
		Index:      78,
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
		NTType:     24,
		Index:      79,
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
		NTType:     24,
		Index:      80,
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
		NTType:     25,
		Index:      81,
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
		NTType:     25,
		Index:      82,
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
		NTType:     26,
		Index:      83,
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
		NTType:     26,
		Index:      84,
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
		NTType:     27,
		Index:      85,
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
		NTType:     27,
		Index:      86,
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
		NTType:     28,
		Index:      87,
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
		NTType:     28,
		Index:      88,
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
		NTType:     29,
		Index:      89,
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
		NTType:     29,
		Index:      90,
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
		NTType:     30,
		Index:      91,
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
		NTType:     30,
		Index:      92,
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
		NTType:     31,
		Index:      93,
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
		NTType:     31,
		Index:      94,
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
		NTType:     32,
		Index:      95,
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
		NTType:     32,
		Index:      96,
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
		NTType:     33,
		Index:      97,
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
		NTType:     33,
		Index:      98,
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
		NTType:     34,
		Index:      99,
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
		NTType:     34,
		Index:      100,
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
		NTType:     35,
		Index:      101,
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
		NTType:     35,
		Index:      102,
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
		NTType:     36,
		Index:      103,
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
		NTType:     36,
		Index:      104,
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
		NTType:     37,
		Index:      105,
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
		NTType:     37,
		Index:      106,
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
		NTType:     38,
		Index:      107,
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
		NTType:     38,
		Index:      108,
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
		NTType:     39,
		Index:      109,
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
		NTType:     39,
		Index:      110,
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
		NTType:     40,
		Index:      111,
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
		NTType:     40,
		Index:      112,
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
		NTType:     41,
		Index:      113,
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
		NTType:     41,
		Index:      114,
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
		NTType:     42,
		Index:      115,
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
		NTType:     42,
		Index:      116,
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
		NTType:     43,
		Index:      117,
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
		NTType:     43,
		Index:      118,
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
		NTType:     44,
		Index:      119,
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
		NTType:     44,
		Index:      120,
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
		NTType:     45,
		Index:      121,
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
		NTType:     45,
		Index:      122,
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
		NTType:     46,
		Index:      123,
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
		NTType:     46,
		Index:      124,
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
		NTType:     47,
		Index:      125,
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
		NTType:     47,
		Index:      126,
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
		NTType:     48,
		Index:      127,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.AppendSpace(X[0], newString(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< &expr.Space{Space: []string{newString(X[0])}}, nil >>`,
		Id:         "Space",
		NTType:     48,
		Index:      128,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Space{Space: []string{newString(X[0])}}, nil
		},
	},
}

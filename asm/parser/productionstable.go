package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/types"
)

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
		String: `S' : AllRules	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AllRules : Rules Space	<< X[0].(*ast.Rules).SetSpace(X[1].(*ast.Space)), nil >>`,
		Id:         "AllRules",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(*ast.Rules).SetSpace(X[1].(*ast.Space)), nil
		},
	},
	ProdTabEntry{
		String: `AllRules : Rules	<< X[0], nil >>`,
		Id:         "AllRules",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rules Rule	<< ast.AppendRule(X[0], X[1]) >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendRule(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< ast.NewRule(X[0]) >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[0])
		},
	},
	ProdTabEntry{
		String: `Rules : Expr	<<  >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Root	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Init	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Transition	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : IfExpr	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Root : Space "root" Equal Space id "." id	<< &ast.Root{Before: X[0].(*ast.Space), Equal: X[2].(*ast.Keyword), BeforeQualId: X[3].(*ast.Space), Package: ast.NewString(X[4]), Message: ast.NewString(X[6]), State: "root"}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      10,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Root{Before: X[0].(*ast.Space), Equal: X[2].(*ast.Keyword), BeforeQualId: X[3].(*ast.Space), Package: ast.NewString(X[4]), Message: ast.NewString(X[6]), State: "root"}, nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" Equal Space id "." id	<< &ast.Root{Equal: X[1].(*ast.Keyword), BeforeQualId: X[2].(*ast.Space), Package: ast.NewString(X[3]), Message: ast.NewString(X[5]), State: "root"}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      11,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Root{Equal: X[1].(*ast.Keyword), BeforeQualId: X[2].(*ast.Space), Package: ast.NewString(X[3]), Message: ast.NewString(X[5]), State: "root"}, nil
		},
	},
	ProdTabEntry{
		String: `Root : Space "root" Equal id "." id	<< &ast.Root{Before: X[0].(*ast.Space), Equal: X[2].(*ast.Keyword), Package: ast.NewString(X[3]), Message: ast.NewString(X[5]), State: "root"}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      12,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Root{Before: X[0].(*ast.Space), Equal: X[2].(*ast.Keyword), Package: ast.NewString(X[3]), Message: ast.NewString(X[5]), State: "root"}, nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" Equal id "." id	<< &ast.Root{Equal: X[1].(*ast.Keyword), Package: ast.NewString(X[2]), Message: ast.NewString(X[4]), State: "root"}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      13,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Root{Equal: X[1].(*ast.Keyword), Package: ast.NewString(X[2]), Message: ast.NewString(X[4]), State: "root"}, nil
		},
	},
	ProdTabEntry{
		String: `Init : Space id "." id Equal Space id	<< &ast.Init{
		Before: X[0].(*ast.Space),
		Package: ast.NewString(X[1]),
		Message: ast.NewString(X[3]),
		Equal: X[4].(*ast.Keyword),
		BeforeState: X[5].(*ast.Space),
		State: ast.NewString(X[6])}, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      14,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{
				Before:      X[0].(*ast.Space),
				Package:     ast.NewString(X[1]),
				Message:     ast.NewString(X[3]),
				Equal:       X[4].(*ast.Keyword),
				BeforeState: X[5].(*ast.Space),
				State:       ast.NewString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `Init : id "." id Equal Space id	<< &ast.Init{
		Package: ast.NewString(X[0]),
		Message: ast.NewString(X[2]),
		Equal: X[3].(*ast.Keyword),
		BeforeState: X[4].(*ast.Space),
		State: ast.NewString(X[5])}, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      15,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{
				Package:     ast.NewString(X[0]),
				Message:     ast.NewString(X[2]),
				Equal:       X[3].(*ast.Keyword),
				BeforeState: X[4].(*ast.Space),
				State:       ast.NewString(X[5])}, nil
		},
	},
	ProdTabEntry{
		String: `Init : Space id "." id Equal id	<< &ast.Init{
		Before: X[0].(*ast.Space),
		Package: ast.NewString(X[1]),
		Message: ast.NewString(X[3]),
		Equal: X[4].(*ast.Keyword),
		State: ast.NewString(X[5])}, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      16,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{
				Before:  X[0].(*ast.Space),
				Package: ast.NewString(X[1]),
				Message: ast.NewString(X[3]),
				Equal:   X[4].(*ast.Keyword),
				State:   ast.NewString(X[5])}, nil
		},
	},
	ProdTabEntry{
		String: `Init : id "." id Equal id	<< &ast.Init{
		Package: ast.NewString(X[0]),
		Message: ast.NewString(X[2]),
		Equal: X[3].(*ast.Keyword),
		State: ast.NewString(X[4])}, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      17,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{
				Package: ast.NewString(X[0]),
				Message: ast.NewString(X[2]),
				Equal:   X[3].(*ast.Keyword),
				State:   ast.NewString(X[4])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : Space id Space id Equal Space id	<< &ast.Transition{Before: X[0].(*ast.Space), Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), BeforeDst: X[5].(*ast.Space), Dst: ast.NewString(X[6])}, nil >>`,
		Id:         "Transition",
		NTType:     6,
		Index:      18,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Before: X[0].(*ast.Space), Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), BeforeDst: X[5].(*ast.Space), Dst: ast.NewString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id Space id Equal Space id	<< &ast.Transition{Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), BeforeDst: X[5].(*ast.Space), Dst: ast.NewString(X[6])}, nil >>`,
		Id:         "Transition",
		NTType:     6,
		Index:      19,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), BeforeDst: X[5].(*ast.Space), Dst: ast.NewString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : Space id Space id Equal id	<< &ast.Transition{Before: X[0].(*ast.Space), Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), Dst: ast.NewString(X[6])}, nil >>`,
		Id:         "Transition",
		NTType:     6,
		Index:      20,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Before: X[0].(*ast.Space), Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), Dst: ast.NewString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id Space id Equal id	<< &ast.Transition{Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), Dst: ast.NewString(X[6])}, nil >>`,
		Id:         "Transition",
		NTType:     6,
		Index:      21,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Src: ast.NewString(X[1]), BeforeInput: X[2].(*ast.Space), Input: ast.NewString(X[3]), Equal: X[4].(*ast.Keyword), Dst: ast.NewString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `IfExpr : Space "if" Expr Then StateExpr Else StateExpr	<< &ast.IfExpr{Before: X[0].(*ast.Space), Condition: X[2].(*ast.Expr), ThenWord: X[3].(*ast.Keyword), ThenClause: X[4].(*ast.StateExpr), ElseWord: X[5].(*ast.Keyword), ElseClause: X[6].(*ast.StateExpr)}, nil >>`,
		Id:         "IfExpr",
		NTType:     7,
		Index:      22,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.IfExpr{Before: X[0].(*ast.Space), Condition: X[2].(*ast.Expr), ThenWord: X[3].(*ast.Keyword), ThenClause: X[4].(*ast.StateExpr), ElseWord: X[5].(*ast.Keyword), ElseClause: X[6].(*ast.StateExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `IfExpr : "if" Expr Then StateExpr Else StateExpr	<< &ast.IfExpr{Condition: X[2].(*ast.Expr), ThenWord: X[3].(*ast.Keyword), ThenClause: X[4].(*ast.StateExpr), ElseWord: X[5].(*ast.Keyword), ElseClause: X[6].(*ast.StateExpr)}, nil >>`,
		Id:         "IfExpr",
		NTType:     7,
		Index:      23,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.IfExpr{Condition: X[2].(*ast.Expr), ThenWord: X[3].(*ast.Keyword), ThenClause: X[4].(*ast.StateExpr), ElseWord: X[5].(*ast.Keyword), ElseClause: X[6].(*ast.StateExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : Space "{" IfExpr CloseCurly	<< &ast.StateExpr{Before: X[0].(*ast.Space), IfExpr: X[2].(*ast.IfExpr), CloseCurly: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "StateExpr",
		NTType:     8,
		Index:      24,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{Before: X[0].(*ast.Space), IfExpr: X[2].(*ast.IfExpr), CloseCurly: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : Space id	<< &ast.StateExpr{Before: X[0].(*ast.Space), State: proto.String(ast.NewString(X[1]))}, nil >>`,
		Id:         "StateExpr",
		NTType:     8,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{Before: X[0].(*ast.Space), State: proto.String(ast.NewString(X[1]))}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : "{" IfExpr CloseCurly	<< &ast.StateExpr{IfExpr: X[2].(*ast.IfExpr), CloseCurly: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "StateExpr",
		NTType:     8,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{IfExpr: X[2].(*ast.IfExpr), CloseCurly: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : id	<< &ast.StateExpr{State: proto.String(ast.NewString(X[1]))}, nil >>`,
		Id:         "StateExpr",
		NTType:     8,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{State: proto.String(ast.NewString(X[1]))}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &ast.Function{Before: X[0].(*ast.Space), Name: ast.NewString(X[1]), OpenParen: X[2].(*ast.Keyword), Params: X[3].([]*ast.Expr), CloseParen: X[4].(*ast.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     9,
		Index:      28,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Before: X[0].(*ast.Space), Name: ast.NewString(X[1]), OpenParen: X[2].(*ast.Keyword), Params: X[3].([]*ast.Expr), CloseParen: X[4].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &ast.Function{Before: X[0].(*ast.Space), Name: ast.NewString(X[1]), OpenParen: X[2].(*ast.Keyword), CloseParen: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     9,
		Index:      29,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Before: X[0].(*ast.Space), Name: ast.NewString(X[1]), OpenParen: X[2].(*ast.Keyword), CloseParen: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &ast.Function{Name: ast.NewString(X[0]), OpenParen: X[1].(*ast.Keyword), Params: X[2].([]*ast.Expr), CloseParen: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     9,
		Index:      30,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.NewString(X[0]), OpenParen: X[1].(*ast.Keyword), Params: X[2].([]*ast.Expr), CloseParen: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &ast.Function{Name: ast.NewString(X[0]), OpenParen: X[1].(*ast.Keyword), CloseParen: X[2].(*ast.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     9,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.NewString(X[0]), OpenParen: X[1].(*ast.Keyword), CloseParen: X[2].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &ast.List{Before: X[0].(*ast.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*ast.Keyword), Elems: X[3].([]*ast.Expr), CloseCurly: X[4].(*ast.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     10,
		Index:      32,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Before: X[0].(*ast.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*ast.Keyword), Elems: X[3].([]*ast.Expr), CloseCurly: X[4].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &ast.List{Type: X[0].(types.Type), OpenCurly: X[1].(*ast.Keyword), Elems: X[2].([]*ast.Expr), CloseCurly: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     10,
		Index:      33,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Type: X[0].(types.Type), OpenCurly: X[1].(*ast.Keyword), Elems: X[2].([]*ast.Expr), CloseCurly: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &ast.List{Before: X[0].(*ast.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*ast.Keyword), CloseCurly: X[3].(*ast.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     10,
		Index:      34,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Before: X[0].(*ast.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*ast.Keyword), CloseCurly: X[3].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &ast.List{Type: X[0].(types.Type), OpenCurly: X[1].(*ast.Keyword), CloseCurly: X[2].(*ast.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     10,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Type: X[0].(types.Type), OpenCurly: X[1].(*ast.Keyword), CloseCurly: X[2].(*ast.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*ast.Expr{X[0].(*ast.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     11,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*ast.Expr{X[0].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*ast.Expr), ast.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     11,
		Index:      37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*ast.Expr), ast.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     12,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &ast.Expr{Function: X[0].(*ast.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     12,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Function: X[0].(*ast.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &ast.Expr{List: X[0].(*ast.List)}, nil >>`,
		Id:         "Expr",
		NTType:     12,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{List: X[0].(*ast.List)}, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int64"	<< types.LIST_INT64, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT64, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int32"	<< types.LIST_INT32, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT32, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint64"	<< types.LIST_UINT64, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT64, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint32"	<< types.LIST_UINT32, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT32, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]float"	<< types.LIST_FLOAT, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_FLOAT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     13,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     14,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< ast.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     14,
		Index:      51,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Terminal : Bool	<< ast.NewBoolTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : int64_lit	<< ast.NewInt64Terminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInt64Terminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : int32_lit	<< ast.NewInt32Terminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInt32Terminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : uint64_lit	<< ast.NewUint64Terminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUint64Terminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : uint32_lit	<< ast.NewUint32Terminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUint32Terminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : double_lit	<< ast.NewDoubleTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDoubleTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : float_lit	<< ast.NewFloatTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFloatTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : string_lit	<< ast.NewStringTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_lit	<< ast.NewBytesTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBytesTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int64_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_INT64) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_INT64)
		},
	},
	ProdTabEntry{
		String: `Terminal : int32_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_INT32) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_INT32)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint64_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_UINT64) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_UINT64)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint32_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_UINT32) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_UINT32)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : float_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_FLOAT) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_FLOAT)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< ast.NewVariableTerminal(X[0], types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableTerminal(X[0], types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     16,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     16,
		Index:      71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "Equal",
		NTType:     17,
		Index:      72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Equal : Space "="	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "Equal",
		NTType:     17,
		Index:      73,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Then : "then"	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "Then",
		NTType:     18,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Then : Space "then"	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "Then",
		NTType:     18,
		Index:      75,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Else : "else"	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "Else",
		NTType:     19,
		Index:      76,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Else : Space "else"	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "Else",
		NTType:     19,
		Index:      77,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : "("	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "OpenParen",
		NTType:     20,
		Index:      78,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : Space "("	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "OpenParen",
		NTType:     20,
		Index:      79,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : ")"	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "CloseParen",
		NTType:     21,
		Index:      80,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : Space ")"	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "CloseParen",
		NTType:     21,
		Index:      81,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : "{"	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "OpenCurly",
		NTType:     22,
		Index:      82,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : Space "{"	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "OpenCurly",
		NTType:     22,
		Index:      83,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : "}"	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "CloseCurly",
		NTType:     23,
		Index:      84,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : Space "}"	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "CloseCurly",
		NTType:     23,
		Index:      85,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Comma : ","	<< ast.NewKeyword(nil, X[0]), nil >>`,
		Id:         "Comma",
		NTType:     24,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(nil, X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Comma : Space ","	<< ast.NewKeyword(X[0], X[1]), nil >>`,
		Id:         "Comma",
		NTType:     24,
		Index:      87,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyword(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Space : Space space	<< ast.AppendSpace(X[0], X[1]), nil >>`,
		Id:         "Space",
		NTType:     25,
		Index:      88,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendSpace(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< ast.NewSpace(X[0]), nil >>`,
		Id:         "Space",
		NTType:     25,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSpace(X[0]), nil
		},
	},
}

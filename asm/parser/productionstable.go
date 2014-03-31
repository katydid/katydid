package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/asm/ast"
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
		String: `S' : Rules	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rules Rule	<< ast.AppendRule(X[0], X[1]) >>`,
		Id:         "Rules",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendRule(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< ast.NewRule(X[0]) >>`,
		Id:         "Rules",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[0])
		},
	},
	ProdTabEntry{
		String: `Rule : Root	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Init	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Transition	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : IfExpr	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" "=" id "." id	<< &ast.Init{Package: ast.IdToString(X[2]), Message: ast.IdToString(X[4]), State: proto.String("root")}, nil >>`,
		Id:         "Root",
		NTType:     3,
		Index:      7,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{Package: ast.IdToString(X[2]), Message: ast.IdToString(X[4]), State: proto.String("root")}, nil
		},
	},
	ProdTabEntry{
		String: `Init : id "." id "=" id	<< &ast.Init{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), State: ast.IdToString(X[4])}, nil >>`,
		Id:         "Init",
		NTType:     4,
		Index:      8,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), State: ast.IdToString(X[4])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id id "=" id	<< &ast.Transition{Src: ast.IdToString(X[0]), Input: ast.IdToString(X[1]), Dst: ast.IdToString(X[3])}, nil >>`,
		Id:         "Transition",
		NTType:     5,
		Index:      9,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Src: ast.IdToString(X[0]), Input: ast.IdToString(X[1]), Dst: ast.IdToString(X[3])}, nil
		},
	},
	ProdTabEntry{
		String: `IfExpr : "if" Expr "then" StateExpr "else" StateExpr	<< &ast.IfExpr{Condition: X[1].(*ast.Expr), Then: X[3].(*ast.StateExpr), Else: X[5].(*ast.StateExpr)}, nil >>`,
		Id:         "IfExpr",
		NTType:     6,
		Index:      10,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.IfExpr{Condition: X[1].(*ast.Expr), Then: X[3].(*ast.StateExpr), Else: X[5].(*ast.StateExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : "{" IfExpr "}"	<< &ast.StateExpr{IfExpr: X[1].(*ast.IfExpr)}, nil >>`,
		Id:         "StateExpr",
		NTType:     7,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{IfExpr: X[1].(*ast.IfExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : id	<< &ast.StateExpr{State: ast.IdToString(X[0])}, nil >>`,
		Id:         "StateExpr",
		NTType:     7,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{State: ast.IdToString(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id "(" Params ")"	<< &ast.Function{Name: ast.IdToString(X[0]), Params: X[2].([]*ast.Expr)}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      13,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.IdToString(X[0]), Params: X[2].([]*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id "(" ")"	<< &ast.Function{Name: ast.IdToString(X[0]), Params: nil}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.IdToString(X[0]), Params: nil}, nil
		},
	},
	ProdTabEntry{
		String: `Function : "(" Expr Comparator Expr ")"	<< &ast.Function{Name: proto.String(X[2].(string)), Params: []*ast.Expr{X[1].(*ast.Expr), X[3].(*ast.Expr)}}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: proto.String(X[2].(string)), Params: []*ast.Expr{X[1].(*ast.Expr), X[3].(*ast.Expr)}}, nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "=="	<< "eq", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "eq", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "<"	<< "lt", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "lt", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "<="	<< "le", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "le", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ">"	<< "gt", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "gt", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ">="	<< "ge", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "ge", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "&&"	<< "and", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "and", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "||"	<< "or", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "or", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "or"	<<  >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "and"	<<  >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Params : Params "," Expr	<< append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil >>`,
		Id:         "Params",
		NTType:     10,
		Index:      25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil
		},
	},
	ProdTabEntry{
		String: `Params : Expr	<< []*ast.Expr{X[0].(*ast.Expr)}, nil >>`,
		Id:         "Params",
		NTType:     10,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*ast.Expr{X[0].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Terminal	<< &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     11,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &ast.Expr{Function: X[0].(*ast.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     11,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Function: X[0].(*ast.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : Bool	<< &ast.Terminal{BoolValue: proto.Bool(X[0].(bool))}, nil >>`,
		Id:         "Terminal",
		NTType:     12,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{BoolValue: proto.Bool(X[0].(bool))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : "int64" "(" int_lit ")"	<< &ast.Terminal{Int64Value: ast.ToInt64(X[2])}, nil >>`,
		Id:         "Terminal",
		NTType:     12,
		Index:      30,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Int64Value: ast.ToInt64(X[2])}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : "uint64" "(" int_lit ")"	<< &ast.Terminal{Uint64Value: ast.ToUint64(X[2])}, nil >>`,
		Id:         "Terminal",
		NTType:     12,
		Index:      31,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Uint64Value: ast.ToUint64(X[2])}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : string_lit	<< &ast.Terminal{StringValue: ast.ToString(X[0])}, nil >>`,
		Id:         "Terminal",
		NTType:     12,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{StringValue: ast.ToString(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : id "." id "." id	<< &ast.Terminal{Variable: &ast.Variable{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), Field: ast.IdToString(X[4])}}, nil >>`,
		Id:         "Terminal",
		NTType:     12,
		Index:      33,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Variable: &ast.Variable{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), Field: ast.IdToString(X[4])}}, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     13,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "True"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     13,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     13,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "False"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     13,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
}

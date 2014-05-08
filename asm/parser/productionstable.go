package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/types"
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
		String: `Rules : Expr	<<  >>`,
		Id:         "Rules",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Root	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Init	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Transition	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : IfExpr	<<  >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" "=" id "." id	<< &ast.Init{Package: ast.IdToString(X[2]), Message: ast.IdToString(X[4]), State: "root"}, nil >>`,
		Id:         "Root",
		NTType:     3,
		Index:      8,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{Package: ast.IdToString(X[2]), Message: ast.IdToString(X[4]), State: "root"}, nil
		},
	},
	ProdTabEntry{
		String: `Init : id "." id "=" id	<< &ast.Init{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), State: ast.IdToString(X[4])}, nil >>`,
		Id:         "Init",
		NTType:     4,
		Index:      9,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Init{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), State: ast.IdToString(X[4])}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id id "=" id	<< &ast.Transition{Src: ast.IdToString(X[0]), Input: ast.IdToString(X[1]), Dst: ast.IdToString(X[3])}, nil >>`,
		Id:         "Transition",
		NTType:     5,
		Index:      10,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Transition{Src: ast.IdToString(X[0]), Input: ast.IdToString(X[1]), Dst: ast.IdToString(X[3])}, nil
		},
	},
	ProdTabEntry{
		String: `IfExpr : "if" Expr "then" StateExpr "else" StateExpr	<< &ast.IfExpr{Condition: X[1].(*ast.Expr), Then: X[3].(*ast.StateExpr), Else: X[5].(*ast.StateExpr)}, nil >>`,
		Id:         "IfExpr",
		NTType:     6,
		Index:      11,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.IfExpr{Condition: X[1].(*ast.Expr), Then: X[3].(*ast.StateExpr), Else: X[5].(*ast.StateExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : "{" IfExpr "}"	<< &ast.StateExpr{IfExpr: X[1].(*ast.IfExpr)}, nil >>`,
		Id:         "StateExpr",
		NTType:     7,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{IfExpr: X[1].(*ast.IfExpr)}, nil
		},
	},
	ProdTabEntry{
		String: `StateExpr : id	<< &ast.StateExpr{State: proto.String(ast.IdToString(X[0]))}, nil >>`,
		Id:         "StateExpr",
		NTType:     7,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.StateExpr{State: proto.String(ast.IdToString(X[0]))}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id "(" Params ")"	<< &ast.Function{Name: ast.IdToString(X[0]), Params: X[2].([]*ast.Expr)}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.IdToString(X[0]), Params: X[2].([]*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id "(" ")"	<< &ast.Function{Name: ast.IdToString(X[0]), Params: nil}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: ast.IdToString(X[0]), Params: nil}, nil
		},
	},
	ProdTabEntry{
		String: `Function : "(" Expr Comparator Expr ")"	<< &ast.Function{Name: X[2].(string), Params: []*ast.Expr{X[1].(*ast.Expr), X[3].(*ast.Expr)}}, nil >>`,
		Id:         "Function",
		NTType:     8,
		Index:      16,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Function{Name: X[2].(string), Params: []*ast.Expr{X[1].(*ast.Expr), X[3].(*ast.Expr)}}, nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "=="	<< "eq", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "eq", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "<"	<< "lt", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "lt", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "<="	<< "le", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "le", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ">"	<< "gt", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "gt", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : ">="	<< "ge", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "ge", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "&&"	<< "and", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "and", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "||"	<< "or", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "or", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "or"	<< "or", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "or", nil
		},
	},
	ProdTabEntry{
		String: `Comparator : "and"	<< "and", nil >>`,
		Id:         "Comparator",
		NTType:     9,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "and", nil
		},
	},
	ProdTabEntry{
		String: `Params : Params "," Expr	<< append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil >>`,
		Id:         "Params",
		NTType:     10,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil
		},
	},
	ProdTabEntry{
		String: `Params : Expr	<< []*ast.Expr{X[0].(*ast.Expr)}, nil >>`,
		Id:         "Params",
		NTType:     10,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*ast.Expr{X[0].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Terminal	<< &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     11,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Terminal: X[0].(*ast.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &ast.Expr{Function: X[0].(*ast.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     11,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{Function: X[0].(*ast.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &ast.Expr{List: X[0].(*ast.List)}, nil >>`,
		Id:         "Expr",
		NTType:     11,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{List: X[0].(*ast.List)}, nil
		},
	},
	ProdTabEntry{
		String: `List : "{" Exprs "}"	<< &ast.List{Elems: X[1].([]*ast.Expr)}, nil >>`,
		Id:         "List",
		NTType:     12,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Elems: X[1].([]*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType "{" Exprs "}"	<< &ast.List{Type: X[0].(*types.Type), Elems: X[2].([]*ast.Expr)}, nil >>`,
		Id:         "List",
		NTType:     12,
		Index:      32,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Type: X[0].(*types.Type), Elems: X[2].([]*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType "{" "}"	<< &ast.List{Type: X[0].(*types.Type)}, nil >>`,
		Id:         "List",
		NTType:     12,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.List{Type: X[0].(*types.Type)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*ast.Expr{X[0].(*ast.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     13,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*ast.Expr{X[0].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs "," Expr	<< append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil >>`,
		Id:         "Exprs",
		NTType:     13,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*ast.Expr), X[2].(*ast.Expr)), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int64"	<< types.LIST_INT64.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT64.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int32"	<< types.LIST_INT32.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT32.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint64"	<< types.LIST_UINT64.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT64.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint32"	<< types.LIST_UINT32.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT32.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]float"	<< types.LIST_FLOAT.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_FLOAT.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES.Enum(), nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES.Enum(), nil
		},
	},
	ProdTabEntry{
		String: `Terminal : Bool	<< &ast.Terminal{BoolValue: proto.Bool(X[0].(bool))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{BoolValue: proto.Bool(X[0].(bool))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : int64_lit	<< &ast.Terminal{Int64Value: ast.ToInt64(ast.Strip(X[0], "int64"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Int64Value: ast.ToInt64(ast.Strip(X[0], "int64"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : int32_lit	<< &ast.Terminal{Int32Value: ast.ToInt32(ast.Strip(X[0], "int32"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Int32Value: ast.ToInt32(ast.Strip(X[0], "int32"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : uint64_lit	<< &ast.Terminal{Uint64Value: ast.ToUint64(ast.Strip(X[0], "uint64"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Uint64Value: ast.ToUint64(ast.Strip(X[0], "uint64"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : uint32_lit	<< &ast.Terminal{Uint32Value: ast.ToUint32(ast.Strip(X[0], "uint32"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Uint32Value: ast.ToUint32(ast.Strip(X[0], "uint32"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : double_lit	<< &ast.Terminal{DoubleValue: ast.ToFloat64(ast.Strip(X[0], "double"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{DoubleValue: ast.ToFloat64(ast.Strip(X[0], "double"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : float_lit	<< &ast.Terminal{FloatValue: ast.ToFloat32(ast.Strip(X[0], "float"))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{FloatValue: ast.ToFloat32(ast.Strip(X[0], "float"))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : string_lit	<< &ast.Terminal{StringValue: proto.String(ast.ToString(X[0]))}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{StringValue: proto.String(ast.ToString(X[0]))}, nil
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_lit	<< ast.NewBytesTerminal(X[0]) >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBytesTerminal(X[0])
		},
	},
	ProdTabEntry{
		String: `Terminal : id "." id "." id	<< &ast.Terminal{Variable: &ast.Variable{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), Field: ast.IdToString(X[4])}}, nil >>`,
		Id:         "Terminal",
		NTType:     15,
		Index:      54,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Terminal{Variable: &ast.Variable{Package: ast.IdToString(X[0]), Message: ast.IdToString(X[2]), Field: ast.IdToString(X[4])}}, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     16,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     16,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
}

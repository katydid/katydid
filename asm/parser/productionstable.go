package parser

import (
	. "github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/token"
	"github.com/katydid/katydid/expr/ast"
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
		String: `AllRules : Rules Space	<< X[0].(*Rules).SetSpace(X[1].(*expr.Space)), nil >>`,
		Id:         "AllRules",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(*Rules).SetSpace(X[1].(*expr.Space)), nil
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
		String: `Rules : Rules Rule	<< AppendRule(X[0], X[1]) >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return AppendRule(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< NewRule(X[0]) >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewRule(X[0])
		},
	},
	ProdTabEntry{
		String: `Rule : Root	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Init	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rule : Final	<<  >>`,
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
		String: `Rule : FunctionDecl	<<  >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Root : Space "root" Equal Space id "." id	<< &Root{Before: X[0].(*expr.Space), Equal: X[2].(*expr.Keyword), BeforeQualId: X[3].(*expr.Space), Package: newString(X[4]), Message: newString(X[6])}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      10,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Root{Before: X[0].(*expr.Space), Equal: X[2].(*expr.Keyword), BeforeQualId: X[3].(*expr.Space), Package: newString(X[4]), Message: newString(X[6])}, nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" Equal Space id "." id	<< &Root{Equal: X[1].(*expr.Keyword), BeforeQualId: X[2].(*expr.Space), Package: newString(X[3]), Message: newString(X[5])}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      11,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Root{Equal: X[1].(*expr.Keyword), BeforeQualId: X[2].(*expr.Space), Package: newString(X[3]), Message: newString(X[5])}, nil
		},
	},
	ProdTabEntry{
		String: `Root : Space "root" Equal id "." id	<< &Root{Before: X[0].(*expr.Space), Equal: X[2].(*expr.Keyword), Package: newString(X[3]), Message: newString(X[5])}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      12,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Root{Before: X[0].(*expr.Space), Equal: X[2].(*expr.Keyword), Package: newString(X[3]), Message: newString(X[5])}, nil
		},
	},
	ProdTabEntry{
		String: `Root : "root" Equal id "." id	<< &Root{Equal: X[1].(*expr.Keyword), Package: newString(X[2]), Message: newString(X[4])}, nil >>`,
		Id:         "Root",
		NTType:     4,
		Index:      13,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Root{Equal: X[1].(*expr.Keyword), Package: newString(X[2]), Message: newString(X[4])}, nil
		},
	},
	ProdTabEntry{
		String: `Init : Space "init" Equal Space id	<< &Init{
		Before: X[0].(*expr.Space), 
		Equal: X[2].(*expr.Keyword),
		BeforeState: X[3].(*expr.Space),
		State: newString(X[4]),
	  }, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      14,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Init{
				Before:      X[0].(*expr.Space),
				Equal:       X[2].(*expr.Keyword),
				BeforeState: X[3].(*expr.Space),
				State:       newString(X[4]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Init : "init" Equal Space id	<< &Init{
		Equal: X[1].(*expr.Keyword),
		BeforeState: X[2].(*expr.Space),
		State: newString(X[3]),
	  }, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      15,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Init{
				Equal:       X[1].(*expr.Keyword),
				BeforeState: X[2].(*expr.Space),
				State:       newString(X[3]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Init : Space "init" Equal id	<< &Init{
		Before: X[0].(*expr.Space), 
		Equal: X[2].(*expr.Keyword),
		State: newString(X[3]),
	  }, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      16,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Init{
				Before: X[0].(*expr.Space),
				Equal:  X[2].(*expr.Keyword),
				State:  newString(X[3]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Init : "init" Equal id	<< &Init{
		Equal: X[1].(*expr.Keyword),
		State: newString(X[2]),
	  }, nil >>`,
		Id:         "Init",
		NTType:     5,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Init{
				Equal: X[1].(*expr.Keyword),
				State: newString(X[2]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : Space "final" Equal Space id	<< &Final{
		Before: X[0].(*expr.Space), 
		Equal: X[2].(*expr.Keyword),
		BeforeState: X[3].(*expr.Space),
		State: newString(X[4]),
	  }, nil >>`,
		Id:         "Final",
		NTType:     6,
		Index:      18,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Before:      X[0].(*expr.Space),
				Equal:       X[2].(*expr.Keyword),
				BeforeState: X[3].(*expr.Space),
				State:       newString(X[4]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : "final" Equal Space id	<< &Final{
		Equal: X[1].(*expr.Keyword),
		BeforeState: X[2].(*expr.Space),
		State: newString(X[3]),
	  }, nil >>`,
		Id:         "Final",
		NTType:     6,
		Index:      19,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Equal:       X[1].(*expr.Keyword),
				BeforeState: X[2].(*expr.Space),
				State:       newString(X[3]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : Space "final" Equal id	<< &Final{
		Before: X[0].(*expr.Space), 
		Equal: X[2].(*expr.Keyword),
		State: newString(X[3]),
	  }, nil >>`,
		Id:         "Final",
		NTType:     6,
		Index:      20,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Before: X[0].(*expr.Space),
				Equal:  X[2].(*expr.Keyword),
				State:  newString(X[3]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : "final" Equal id	<< &Final{
		Equal: X[1].(*expr.Keyword),
		State: newString(X[2]),
	  }, nil >>`,
		Id:         "Final",
		NTType:     6,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Equal: X[1].(*expr.Keyword),
				State: newString(X[2]),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : Space id Space id Equal Destination	<< &Transition{
		Before: X[0].(*expr.Space),
		Src: newString(X[1]),
		BeforeInput: X[2].(*expr.Space),
		Input: newString(X[3]),
		Equal: X[4].(*expr.Keyword),
		Dst: X[5].(*Destination),
	}, nil >>`,
		Id:         "Transition",
		NTType:     7,
		Index:      22,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Transition{
				Before:      X[0].(*expr.Space),
				Src:         newString(X[1]),
				BeforeInput: X[2].(*expr.Space),
				Input:       newString(X[3]),
				Equal:       X[4].(*expr.Keyword),
				Dst:         X[5].(*Destination),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id Space id Equal Destination	<< &Transition{
		Src: newString(X[0]),
		BeforeInput: X[1].(*expr.Space),
		Input: newString(X[2]),
		Equal: X[3].(*expr.Keyword),
		Dst: X[4].(*Destination),
	}, nil >>`,
		Id:         "Transition",
		NTType:     7,
		Index:      23,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Transition{
				Src:         newString(X[0]),
				BeforeInput: X[1].(*expr.Space),
				Input:       newString(X[2]),
				Equal:       X[3].(*expr.Keyword),
				Dst:         X[4].(*Destination),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : Space id id Equal Destination	<< &Transition{
		Before: X[0].(*expr.Space),
		Src: newString(X[1]),
		Input: newString(X[2]),
		Equal: X[3].(*expr.Keyword),
		Dst: X[4].(*Destination),
	}, nil >>`,
		Id:         "Transition",
		NTType:     7,
		Index:      24,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Transition{
				Before: X[0].(*expr.Space),
				Src:    newString(X[1]),
				Input:  newString(X[2]),
				Equal:  X[3].(*expr.Keyword),
				Dst:    X[4].(*Destination),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Transition : id id Equal Destination	<< &Transition{
		Src: newString(X[0]),
		Input: newString(X[1]),
		Equal: X[2].(*expr.Keyword),
		Dst: X[3].(*Destination),
	}, nil >>`,
		Id:         "Transition",
		NTType:     7,
		Index:      25,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Transition{
				Src:   newString(X[0]),
				Input: newString(X[1]),
				Equal: X[2].(*expr.Keyword),
				Dst:   X[3].(*Destination),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen Space id Comma Space id Comma Space id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		BeforeChild: X[1].(*expr.Space),
		Child: newString(X[2]),
		CommaOne: X[3].(*expr.Keyword),
		BeforeSuccess: X[4].(*expr.Space),
		Success: newString(X[5]),
		CommaTwo: X[6].(*expr.Keyword),
		BeforeFailure: X[7].(*expr.Space),
		Failure: newString(X[8]),
		CloseParen: X[9].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      26,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				BeforeChild:   X[1].(*expr.Space),
				Child:         newString(X[2]),
				CommaOne:      X[3].(*expr.Keyword),
				BeforeSuccess: X[4].(*expr.Space),
				Success:       newString(X[5]),
				CommaTwo:      X[6].(*expr.Keyword),
				BeforeFailure: X[7].(*expr.Space),
				Failure:       newString(X[8]),
				CloseParen:    X[9].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen id Comma Space id Comma Space id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		Child: newString(X[1]),
		CommaOne: X[2].(*expr.Keyword),
		BeforeSuccess: X[3].(*expr.Space),
		Success: newString(X[4]),
		CommaTwo: X[5].(*expr.Keyword),
		BeforeFailure: X[6].(*expr.Space),
		Failure: newString(X[7]),
		CloseParen: X[8].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      27,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				Child:         newString(X[1]),
				CommaOne:      X[2].(*expr.Keyword),
				BeforeSuccess: X[3].(*expr.Space),
				Success:       newString(X[4]),
				CommaTwo:      X[5].(*expr.Keyword),
				BeforeFailure: X[6].(*expr.Space),
				Failure:       newString(X[7]),
				CloseParen:    X[8].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen Space id Comma id Comma Space id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		BeforeChild: X[1].(*expr.Space),
		Child: newString(X[2]),
		CommaOne: X[3].(*expr.Keyword),
		Success: newString(X[4]),
		CommaTwo: X[5].(*expr.Keyword),
		BeforeFailure: X[6].(*expr.Space),
		Failure: newString(X[7]),
		CloseParen: X[8].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      28,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				BeforeChild:   X[1].(*expr.Space),
				Child:         newString(X[2]),
				CommaOne:      X[3].(*expr.Keyword),
				Success:       newString(X[4]),
				CommaTwo:      X[5].(*expr.Keyword),
				BeforeFailure: X[6].(*expr.Space),
				Failure:       newString(X[7]),
				CloseParen:    X[8].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen Space id Comma Space id Comma id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		BeforeChild: X[1].(*expr.Space),
		Child: newString(X[2]),
		CommaOne: X[3].(*expr.Keyword),
		BeforeSuccess: X[4].(*expr.Space),
		Success: newString(X[5]),
		CommaTwo: X[6].(*expr.Keyword),
		Failure: newString(X[7]),
		CloseParen: X[8].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      29,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				BeforeChild:   X[1].(*expr.Space),
				Child:         newString(X[2]),
				CommaOne:      X[3].(*expr.Keyword),
				BeforeSuccess: X[4].(*expr.Space),
				Success:       newString(X[5]),
				CommaTwo:      X[6].(*expr.Keyword),
				Failure:       newString(X[7]),
				CloseParen:    X[8].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen id Comma id Comma Space id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		Child: newString(X[1]),
		CommaOne: X[2].(*expr.Keyword),
		Success: newString(X[3]),
		CommaTwo: X[4].(*expr.Keyword),
		BeforeFailure: X[5].(*expr.Space),
		Failure: newString(X[6]),
		CloseParen: X[7].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      30,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				Child:         newString(X[1]),
				CommaOne:      X[2].(*expr.Keyword),
				Success:       newString(X[3]),
				CommaTwo:      X[4].(*expr.Keyword),
				BeforeFailure: X[5].(*expr.Space),
				Failure:       newString(X[6]),
				CloseParen:    X[7].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen id Comma Space id Comma id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		Child: newString(X[1]),
		CommaOne: X[2].(*expr.Keyword),
		BeforeSuccess: X[3].(*expr.Space),
		Success: newString(X[4]),
		CommaTwo: X[5].(*expr.Keyword),
		Failure: newString(X[6]),
		CloseParen: X[7].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      31,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				Child:         newString(X[1]),
				CommaOne:      X[2].(*expr.Keyword),
				BeforeSuccess: X[3].(*expr.Space),
				Success:       newString(X[4]),
				CommaTwo:      X[5].(*expr.Keyword),
				Failure:       newString(X[6]),
				CloseParen:    X[7].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen id Comma id Comma id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		Child: newString(X[1]),
		CommaOne: X[2].(*expr.Keyword),
		Success: newString(X[3]),
		CommaTwo: X[4].(*expr.Keyword),
		Failure: newString(X[5]),
		CloseParen: X[6].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      32,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:  X[0].(*expr.Keyword),
				Child:      newString(X[1]),
				CommaOne:   X[2].(*expr.Keyword),
				Success:    newString(X[3]),
				CommaTwo:   X[4].(*expr.Keyword),
				Failure:    newString(X[5]),
				CloseParen: X[6].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen id Comma id Comma Space id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		Child: newString(X[1]),
		CommaOne: X[2].(*expr.Keyword),
		Success: newString(X[3]),
		CommaTwo: X[4].(*expr.Keyword),
		BeforeFailure: X[5].(*expr.Space),
		Failure: newString(X[6]),
		CloseParen: X[7].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      33,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:     X[0].(*expr.Keyword),
				Child:         newString(X[1]),
				CommaOne:      X[2].(*expr.Keyword),
				Success:       newString(X[3]),
				CommaTwo:      X[4].(*expr.Keyword),
				BeforeFailure: X[5].(*expr.Space),
				Failure:       newString(X[6]),
				CloseParen:    X[7].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Destination : OpenParen Space id Comma id Comma id CloseParen	<< &Destination{
		OpenParen: X[0].(*expr.Keyword),
		BeforeChild: X[1].(*expr.Space),
		Child: newString(X[2]),
		CommaOne: X[3].(*expr.Keyword),
		Success: newString(X[4]),
		CommaTwo: X[5].(*expr.Keyword),
		Failure: newString(X[6]),
		CloseParen: X[7].(*expr.Keyword),
	}, nil >>`,
		Id:         "Destination",
		NTType:     8,
		Index:      34,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Destination{
				OpenParen:   X[0].(*expr.Keyword),
				BeforeChild: X[1].(*expr.Space),
				Child:       newString(X[2]),
				CommaOne:    X[3].(*expr.Keyword),
				Success:     newString(X[4]),
				CommaTwo:    X[5].(*expr.Keyword),
				Failure:     newString(X[6]),
				CloseParen:  X[7].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `FunctionDecl : Space "func" Space id Equal Function	<< &FunctionDecl{
		Before: X[0].(*expr.Space),
		BeforeName: X[2].(*expr.Space),
		Name: newString(X[3]),
		Equal: X[4].(*expr.Keyword),
		Function: X[5].(*expr.Function),
	}, nil >>`,
		Id:         "FunctionDecl",
		NTType:     9,
		Index:      35,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &FunctionDecl{
				Before:     X[0].(*expr.Space),
				BeforeName: X[2].(*expr.Space),
				Name:       newString(X[3]),
				Equal:      X[4].(*expr.Keyword),
				Function:   X[5].(*expr.Function),
			}, nil
		},
	},
	ProdTabEntry{
		String: `FunctionDecl : "func" Space id Equal Function	<< &FunctionDecl{
		BeforeName: X[1].(*expr.Space),
		Name: newString(X[2]),
		Equal: X[3].(*expr.Keyword),
		Function: X[4].(*expr.Function),
	}, nil >>`,
		Id:         "FunctionDecl",
		NTType:     9,
		Index:      36,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &FunctionDecl{
				BeforeName: X[1].(*expr.Space),
				Name:       newString(X[2]),
				Equal:      X[3].(*expr.Keyword),
				Function:   X[4].(*expr.Function),
			}, nil
		},
	},
	ProdTabEntry{
		String: `FunctionDecl : Space "func" id Equal Function	<< &FunctionDecl{
		Before: X[0].(*expr.Space),
		Name: newString(X[2]),
		Equal: X[3].(*expr.Keyword),
		Function: X[4].(*expr.Function),
	}, nil >>`,
		Id:         "FunctionDecl",
		NTType:     9,
		Index:      37,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &FunctionDecl{
				Before:   X[0].(*expr.Space),
				Name:     newString(X[2]),
				Equal:    X[3].(*expr.Keyword),
				Function: X[4].(*expr.Function),
			}, nil
		},
	},
	ProdTabEntry{
		String: `FunctionDecl : "func" id Equal Function	<< &FunctionDecl{
		Name: newString(X[1]),
		Equal: X[2].(*expr.Keyword),
		Function: X[3].(*expr.Function),
	}, nil >>`,
		Id:         "FunctionDecl",
		NTType:     9,
		Index:      38,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &FunctionDecl{
				Name:     newString(X[1]),
				Equal:    X[2].(*expr.Keyword),
				Function: X[3].(*expr.Function),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     10,
		Index:      39,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     10,
		Index:      40,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     10,
		Index:      41,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     10,
		Index:      42,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     11,
		Index:      43,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     11,
		Index:      44,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     11,
		Index:      45,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     11,
		Index:      46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*expr.Expr{X[0].(*expr.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     12,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*expr.Expr{X[0].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     12,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     13,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &expr.Expr{Function: X[0].(*expr.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     13,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Function: X[0].(*expr.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &expr.Expr{List: X[0].(*expr.List)}, nil >>`,
		Id:         "Expr",
		NTType:     13,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{List: X[0].(*expr.List)}, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int64"	<< types.LIST_INT64, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT64, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int32"	<< types.LIST_INT32, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT32, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint64"	<< types.LIST_UINT64, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT64, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint32"	<< types.LIST_UINT32, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT32, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]float"	<< types.LIST_FLOAT, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_FLOAT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     14,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     15,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< expr.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     15,
		Index:      62,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Terminal : Bool	<< expr.NewBoolTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBoolTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : int64_lit	<< expr.NewInt64Terminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewInt64Terminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : int32_lit	<< expr.NewInt32Terminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewInt32Terminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : uint64_lit	<< expr.NewUint64Terminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUint64Terminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : uint32_lit	<< expr.NewUint32Terminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUint32Terminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : double_lit	<< expr.NewDoubleTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewDoubleTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : float_lit	<< expr.NewFloatTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewFloatTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : string_lit	<< expr.NewStringTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewStringTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_lit	<< expr.NewBytesTerminal(newString(X[0])) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBytesTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< expr.NewVariableTerminal(types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int64_var	<< expr.NewVariableTerminal(types.SINGLE_INT64) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      73,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT64)
		},
	},
	ProdTabEntry{
		String: `Terminal : int32_var	<< expr.NewVariableTerminal(types.SINGLE_INT32) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT32)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint64_var	<< expr.NewVariableTerminal(types.SINGLE_UINT64) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      75,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT64)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint32_var	<< expr.NewVariableTerminal(types.SINGLE_UINT32) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      76,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT32)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< expr.NewVariableTerminal(types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : float_var	<< expr.NewVariableTerminal(types.SINGLE_FLOAT) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      78,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_FLOAT)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< expr.NewVariableTerminal(types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< expr.NewVariableTerminal(types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     16,
		Index:      80,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     17,
		Index:      81,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     17,
		Index:      82,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< &expr.Keyword{Value: "="}, nil >>`,
		Id:         "Equal",
		NTType:     18,
		Index:      83,
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
		NTType:     18,
		Index:      84,
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
		NTType:     19,
		Index:      85,
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
		NTType:     19,
		Index:      86,
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
		NTType:     20,
		Index:      87,
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
		NTType:     20,
		Index:      88,
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
		NTType:     21,
		Index:      89,
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
		NTType:     21,
		Index:      90,
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
		NTType:     22,
		Index:      91,
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
		NTType:     22,
		Index:      92,
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
		NTType:     23,
		Index:      93,
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
		NTType:     23,
		Index:      94,
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
		NTType:     24,
		Index:      95,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.AppendSpace(X[0], newString(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< &expr.Space{Space: []string{newString(X[0])}}, nil >>`,
		Id:         "Space",
		NTType:     24,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Space{Space: []string{newString(X[0])}}, nil
		},
	},
}

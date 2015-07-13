/*
 */
package parser

const numNTSymbols = 21

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // Grammar
		2,  // PatternDecls
		4,  // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		3,  // Space

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		8,  // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		7,  // Space

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S5

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		13, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		12, // Space

	},
	gotoRow{ // S6

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S7

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S8

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S9

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S10

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		17, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		12, // Space

	},
	gotoRow{ // S11

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S12

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S13

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		21, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		20, // Space

	},
	gotoRow{ // S14

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S15

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S16

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S17

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		33, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		20, // Space

	},
	gotoRow{ // S18

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S19

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S20

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S21

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S22

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S23

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S24

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		46, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		45, // Space

	},
	gotoRow{ // S25

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		50, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		49, // Space

	},
	gotoRow{ // S26

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		53, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S27

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		55, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S28

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		56, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S29

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		57, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S30

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		59, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		58, // Space

	},
	gotoRow{ // S31

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		61, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S32

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S33

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S34

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S35

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S36

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		62, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		45, // Space

	},
	gotoRow{ // S37

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		63, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		49, // Space

	},
	gotoRow{ // S38

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		64, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S39

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		65, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S40

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		66, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S41

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		67, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S42

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		68, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		58, // Space

	},
	gotoRow{ // S43

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		69, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		52, // Space

	},
	gotoRow{ // S44

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S45

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S46

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		73, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		72, // Space

	},
	gotoRow{ // S47

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S48

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S49

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S50

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		83, // Expr
		85, // Function
		86, // List
		-1, // Exprs
		87, // ListType
		84, // SpaceTerminal
		94, // Terminal
		95, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		80, // Space

	},
	gotoRow{ // S51

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S52

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S53

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		111, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S54

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S55

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		122, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S56

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		123, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S57

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		125, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S58

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S59

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		137, // Space

	},
	gotoRow{ // S60

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S61

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		139, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S62

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		140, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S63

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		141, // Expr
		85,  // Function
		86,  // List
		-1,  // Exprs
		87,  // ListType
		84,  // SpaceTerminal
		94,  // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		80,  // Space

	},
	gotoRow{ // S64

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		142, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S65

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		143, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S66

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		144, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S67

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		145, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S68

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		146, // Space

	},
	gotoRow{ // S69

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		148, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S70

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S71

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S72

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S73

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		155, // Comma
		154, // Space

	},
	gotoRow{ // S74

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S75

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		159, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		158, // Space

	},
	gotoRow{ // S76

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		161, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S77

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		162, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S78

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S79

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S80

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		164, // ListType
		-1,  // SpaceTerminal
		165, // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S81

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		168, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S82

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S83

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		171, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S84

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S85

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S86

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S87

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		175, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S88

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S89

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S90

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S91

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S92

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S93

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S94

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S95

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S96

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S97

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S98

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S99

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S100

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S101

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S102

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S103

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S104

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S105

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S106

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S107

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S108

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S109

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S110

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S111

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		188, // Comma
		154, // Space

	},
	gotoRow{ // S112

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S113

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S114

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		189, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S115

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		190, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		49,  // Space

	},
	gotoRow{ // S116

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		191, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S117

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		192, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S118

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		193, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S119

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		194, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S120

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		195, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		58,  // Space

	},
	gotoRow{ // S121

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		196, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S122

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		197, // Comma
		154, // Space

	},
	gotoRow{ // S123

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		198, // Comma
		154, // Space

	},
	gotoRow{ // S124

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S125

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		209, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S126

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S127

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S128

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		210, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S129

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		211, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		49,  // Space

	},
	gotoRow{ // S130

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		212, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S131

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		213, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S132

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		214, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S133

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		215, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S134

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		216, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		58,  // Space

	},
	gotoRow{ // S135

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		217, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S136

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S137

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S138

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		219, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S139

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		220, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S140

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		221, // Comma
		154, // Space

	},
	gotoRow{ // S141

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		222, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S142

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		223, // Comma
		154, // Space

	},
	gotoRow{ // S143

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		224, // Comma
		154, // Space

	},
	gotoRow{ // S144

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		225, // Comma
		154, // Space

	},
	gotoRow{ // S145

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		226, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S146

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S147

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		228, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S148

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		229, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S149

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S150

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		230, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		158, // Space

	},
	gotoRow{ // S151

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		231, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S152

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		232, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S153

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S154

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S155

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		235, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S156

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S157

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S158

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S159

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		237, // Space

	},
	gotoRow{ // S160

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S161

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		241, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S162

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		246, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S163

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		247, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S164

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		248, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S165

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S166

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S167

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S168

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		258, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		253, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		250, // Space

	},
	gotoRow{ // S169

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S170

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S171

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S172

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S173

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S174

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S175

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		287, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		289, // CloseCurly
		-1,  // Comma
		280, // Space

	},
	gotoRow{ // S176

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S177

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S178

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S179

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S180

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		306, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S181

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		307, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		49,  // Space

	},
	gotoRow{ // S182

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		308, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S183

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		309, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S184

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		310, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S185

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		311, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S186

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		312, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		58,  // Space

	},
	gotoRow{ // S187

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		313, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S188

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		314, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S189

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		315, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S190

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		316, // Expr
		85,  // Function
		86,  // List
		-1,  // Exprs
		87,  // ListType
		84,  // SpaceTerminal
		94,  // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		80,  // Space

	},
	gotoRow{ // S191

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		317, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S192

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		318, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S193

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		319, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S194

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		320, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S195

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		321, // Space

	},
	gotoRow{ // S196

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		323, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S197

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		324, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S198

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		325, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S199

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S200

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S201

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		326, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S202

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		327, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		49,  // Space

	},
	gotoRow{ // S203

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		328, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S204

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		329, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S205

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		330, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S206

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		331, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S207

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		332, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		58,  // Space

	},
	gotoRow{ // S208

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		333, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		52,  // Space

	},
	gotoRow{ // S209

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S210

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		334, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S211

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		335, // Expr
		85,  // Function
		86,  // List
		-1,  // Exprs
		87,  // ListType
		84,  // SpaceTerminal
		94,  // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		80,  // Space

	},
	gotoRow{ // S212

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		336, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S213

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		337, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S214

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		338, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S215

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		339, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S216

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		340, // Space

	},
	gotoRow{ // S217

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		342, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S218

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		343, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S219

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S220

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S221

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		344, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S222

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S223

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		345, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S224

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		346, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S225

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		347, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S226

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S227

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		348, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S228

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S229

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S230

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		349, // Space

	},
	gotoRow{ // S231

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		351, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S232

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		352, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S233

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S234

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S235

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		353, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S236

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S237

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S238

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		357, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S239

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S240

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S241

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		363, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S242

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S243

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		364, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		158, // Space

	},
	gotoRow{ // S244

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		365, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S245

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		366, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S246

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		368, // Comma
		367, // Space

	},
	gotoRow{ // S247

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		371, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		370, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		250, // Space

	},
	gotoRow{ // S248

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		372, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		373, // CloseCurly
		-1,  // Comma
		280, // Space

	},
	gotoRow{ // S249

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S250

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		375, // ListType
		-1,  // SpaceTerminal
		376, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S251

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		379, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S252

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S253

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S254

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S255

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S256

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S257

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S258

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		381, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		380, // Space

	},
	gotoRow{ // S259

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		385, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S260

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S261

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S262

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S263

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S264

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S265

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S266

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S267

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S268

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S269

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S270

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S271

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S272

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S273

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S274

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S275

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S276

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S277

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S278

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S279

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S280

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		387, // ListType
		-1,  // SpaceTerminal
		388, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S281

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		391, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S282

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S283

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S284

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S285

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S286

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S287

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		394, // CloseCurly
		393, // Comma
		392, // Space

	},
	gotoRow{ // S288

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		396, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S289

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S290

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S291

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S292

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S293

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S294

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S295

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S296

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S297

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S298

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S299

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S300

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S301

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S302

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S303

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S304

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S305

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S306

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		397, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S307

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		398, // Expr
		85,  // Function
		86,  // List
		-1,  // Exprs
		87,  // ListType
		84,  // SpaceTerminal
		94,  // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		80,  // Space

	},
	gotoRow{ // S308

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		399, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S309

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		400, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S310

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		401, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S311

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		402, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S312

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		403, // Space

	},
	gotoRow{ // S313

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		405, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S314

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		406, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S315

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		407, // Comma
		154, // Space

	},
	gotoRow{ // S316

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		408, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S317

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		409, // Comma
		154, // Space

	},
	gotoRow{ // S318

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		410, // Comma
		154, // Space

	},
	gotoRow{ // S319

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		411, // Comma
		154, // Space

	},
	gotoRow{ // S320

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		412, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S321

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S322

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		414, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S323

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		415, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S324

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		416, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S325

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		417, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S326

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		418, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S327

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		419, // Expr
		85,  // Function
		86,  // List
		-1,  // Exprs
		87,  // ListType
		84,  // SpaceTerminal
		94,  // Terminal
		95,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		80,  // Space

	},
	gotoRow{ // S328

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		420, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S329

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		421, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S330

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		422, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		110, // Space

	},
	gotoRow{ // S331

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		423, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S332

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		424, // Space

	},
	gotoRow{ // S333

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		426, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S334

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		427, // Comma
		154, // Space

	},
	gotoRow{ // S335

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		429, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S336

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		430, // Comma
		154, // Space

	},
	gotoRow{ // S337

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		431, // Comma
		154, // Space

	},
	gotoRow{ // S338

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		432, // Comma
		154, // Space

	},
	gotoRow{ // S339

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		433, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S340

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S341

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		435, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S342

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		436, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S343

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S344

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		437, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S345

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		438, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S346

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		439, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S347

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		440, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		170, // Space

	},
	gotoRow{ // S348

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S349

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S350

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		442, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S351

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		443, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S352

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		444, // Comma
		367, // Space

	},
	gotoRow{ // S353

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S354

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		445, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S355

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S356

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S357

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S358

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S359

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S360

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		447, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		158, // Space

	},
	gotoRow{ // S361

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		448, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S362

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		449, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S363

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S364

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		450, // Space

	},
	gotoRow{ // S365

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		452, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S366

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		453, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S367

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S368

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		455, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S369

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S370

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S371

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		456, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		380, // Space

	},
	gotoRow{ // S372

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		457, // CloseCurly
		393, // Comma
		392, // Space

	},
	gotoRow{ // S373

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S374

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		458, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S375

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		459, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S376

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S377

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S378

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S379

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		462, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		461, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		460, // Space

	},
	gotoRow{ // S380

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S381

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S382

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		467, // Expr
		256, // Function
		257, // List
		-1,  // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		466, // Space

	},
	gotoRow{ // S383

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S384

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S385

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		469, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		470, // CloseCurly
		-1,  // Comma
		468, // Space

	},
	gotoRow{ // S386

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		472, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		167, // Space

	},
	gotoRow{ // S387

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		473, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		174, // Space

	},
	gotoRow{ // S388

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S389

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S390

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S391

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		476, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		475, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		474, // Space

	},
	gotoRow{ // S392

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S393

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		480, // Expr
		285, // Function
		286, // List
		-1,  // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		479, // Space

	},
	gotoRow{ // S394

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S395

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S396

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		482, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		483, // CloseCurly
		-1,  // Comma
		481, // Space

	},
	gotoRow{ // S397

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		485, // Comma
		154, // Space

	},
	gotoRow{ // S398

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		486, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S399

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		487, // Comma
		154, // Space

	},
	gotoRow{ // S400

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		488, // Comma
		154, // Space

	},
	gotoRow{ // S401

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		489, // Comma
		154, // Space

	},
	gotoRow{ // S402

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		490, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S403

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S404

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		492, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S405

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		493, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S406

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S407

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		494, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S408

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S409

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		495, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S410

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		496, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S411

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		497, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S412

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S413

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		498, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S414

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S415

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S416

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S417

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S418

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		499, // Comma
		154, // Space

	},
	gotoRow{ // S419

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		500, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S420

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		501, // Comma
		154, // Space

	},
	gotoRow{ // S421

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		502, // Comma
		154, // Space

	},
	gotoRow{ // S422

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		503, // Comma
		154, // Space

	},
	gotoRow{ // S423

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		504, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S424

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S425

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		506, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S426

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		507, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S427

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		508, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S428

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S429

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S430

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		509, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S431

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		510, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S432

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		511, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S433

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S434

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		512, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S435

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S436

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S437

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S438

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S439

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S440

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S441

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		513, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S442

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S443

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S444

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		514, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S445

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S446

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S447

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		515, // Space

	},
	gotoRow{ // S448

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		517, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S449

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		518, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		72,  // Space

	},
	gotoRow{ // S450

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S451

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		520, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S452

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		521, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S453

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		522, // Comma
		367, // Space

	},
	gotoRow{ // S454

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S455

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		523, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S456

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S457

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S458

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		525, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		524, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		460, // Space

	},
	gotoRow{ // S459

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		526, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		527, // CloseCurly
		-1,  // Comma
		468, // Space

	},
	gotoRow{ // S460

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		375, // ListType
		-1,  // SpaceTerminal
		376, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S461

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S462

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		530, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		529, // Space

	},
	gotoRow{ // S463

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S464

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S465

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S466

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		375, // ListType
		-1,  // SpaceTerminal
		376, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S467

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S468

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		387, // ListType
		-1,  // SpaceTerminal
		388, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S469

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		533, // CloseCurly
		393, // Comma
		532, // Space

	},
	gotoRow{ // S470

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S471

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S472

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		254, // Expr
		256, // Function
		257, // List
		535, // Exprs
		259, // ListType
		255, // SpaceTerminal
		260, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		534, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		474, // Space

	},
	gotoRow{ // S473

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		283, // Expr
		285, // Function
		286, // List
		536, // Exprs
		288, // ListType
		284, // SpaceTerminal
		290, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		537, // CloseCurly
		-1,  // Comma
		481, // Space

	},
	gotoRow{ // S474

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		375, // ListType
		-1,  // SpaceTerminal
		376, // Terminal
		261, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S475

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S476

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		540, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		539, // Space

	},
	gotoRow{ // S477

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S478

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S479

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		387, // ListType
		-1,  // SpaceTerminal
		388, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S480

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S481

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		387, // ListType
		-1,  // SpaceTerminal
		388, // Terminal
		291, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S482

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		543, // CloseCurly
		393, // Comma
		542, // Space

	},
	gotoRow{ // S483

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S484

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S485

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		544, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S486

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S487

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		545, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S488

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		546, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S489

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		547, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S490

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S491

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		548, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S492

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S493

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S494

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		549, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S495

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		550, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S496

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		551, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S497

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		552, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S498

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S499

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		553, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S500

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S501

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		554, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S502

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		555, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S503

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		556, // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		124, // Space

	},
	gotoRow{ // S504

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S505

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		557, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S506

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S507

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S508

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		558, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S509

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		559, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S510

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		560, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S511

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		561, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S512

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S513

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S514

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		562, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S515

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S516

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		564, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S517

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		565, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S518

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		566, // Comma
		367, // Space

	},
	gotoRow{ // S519

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		567, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S520

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S521

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S522

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		568, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S523

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S524

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S525

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		569, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		529, // Space

	},
	gotoRow{ // S526

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		570, // CloseCurly
		393, // Comma
		532, // Space

	},
	gotoRow{ // S527

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S528

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S529

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S530

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S531

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S532

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S533

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S534

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S535

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		571, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		382, // Comma
		539, // Space

	},
	gotoRow{ // S536

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		572, // CloseCurly
		393, // Comma
		542, // Space

	},
	gotoRow{ // S537

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S538

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S539

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S540

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S541

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S542

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S543

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S544

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		573, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S545

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		574, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S546

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		575, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S547

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		576, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		356, // Space

	},
	gotoRow{ // S548

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S549

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S550

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S551

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S552

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S553

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		577, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S554

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		578, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S555

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		579, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S556

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		580, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S557

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S558

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S559

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S560

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S561

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S562

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S563

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		581, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S564

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S565

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S566

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		582, // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		240, // Space

	},
	gotoRow{ // S567

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S568

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		583, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S569

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S570

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S571

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S572

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S573

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S574

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S575

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S576

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S577

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S578

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S579

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S580

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S581

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S582

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // NameExpr
		-1,  // Pattern
		-1,  // Expr
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		584, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S583

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S584

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // NameExpr
		-1, // Pattern
		-1, // Expr
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
}

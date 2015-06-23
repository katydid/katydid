/*
 */
package parser

const numNTSymbols = 20

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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		21, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		33, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		49, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		45, // Space

	},
	gotoRow{ // S26

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		51, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		50, // Space

	},
	gotoRow{ // S27

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S28

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		54, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		50, // Space

	},
	gotoRow{ // S29

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S30

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		56, // Space

	},
	gotoRow{ // S31

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S32

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		60, // OpenParen
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		45, // Space

	},
	gotoRow{ // S38

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S39

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S40

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S41

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S42

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		56, // Space

	},
	gotoRow{ // S43

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
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
		50, // Space

	},
	gotoRow{ // S44

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		73, // Function
		74, // List
		-1, // Exprs
		72, // Expr
		75, // ListType
		76, // SpaceTerminal
		86, // Terminal
		87, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		70, // Space

	},
	gotoRow{ // S47

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		110, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S50

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S51

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		138, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S52

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		149, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S54

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		150, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S55

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		152, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S56

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S57

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		164, // Space

	},
	gotoRow{ // S58

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		166, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S60

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		73,  // Function
		74,  // List
		-1,  // Exprs
		167, // Expr
		75,  // ListType
		76,  // SpaceTerminal
		86,  // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		70,  // Space

	},
	gotoRow{ // S61

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		168, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S62

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		169, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S63

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		170, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S64

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		171, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S65

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		172, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S66

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		173, // Space

	},
	gotoRow{ // S67

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		175, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S68

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S69

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S70

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		177, // ListType
		-1,  // SpaceTerminal
		178, // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S71

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		181, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S72

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		184, // Comma
		183, // Space

	},
	gotoRow{ // S73

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S74

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		188, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S76

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S77

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S78

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S81

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S82

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S84

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S88

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		192, // ListType
		-1,  // SpaceTerminal
		193, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S109

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		180, // Space

	},
	gotoRow{ // S110

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		196, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S111

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S112

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		199, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S114

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S115

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S116

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S117

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S118

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S119

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S120

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S121

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S122

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S123

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S124

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S126

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S129

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S130

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S131

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S132

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S133

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S134

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S135

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S136

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		210, // Comma
		183, // Space

	},
	gotoRow{ // S139

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S140

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S141

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		45,  // Space

	},
	gotoRow{ // S142

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		45,  // Space

	},
	gotoRow{ // S143

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		50,  // Space

	},
	gotoRow{ // S144

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		50,  // Space

	},
	gotoRow{ // S145

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		50,  // Space

	},
	gotoRow{ // S146

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		50,  // Space

	},
	gotoRow{ // S147

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		56,  // Space

	},
	gotoRow{ // S148

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		218, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S149

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		219, // Comma
		183, // Space

	},
	gotoRow{ // S150

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		220, // Comma
		183, // Space

	},
	gotoRow{ // S151

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S152

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		231, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S153

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S156

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		233, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S157

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		234, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S158

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		235, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S159

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		236, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S160

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		237, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S161

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		238, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		56,  // Space

	},
	gotoRow{ // S162

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		239, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S163

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S164

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S165

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		241, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S166

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		242, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S167

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		243, // Comma
		183, // Space

	},
	gotoRow{ // S168

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		244, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S169

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		245, // Comma
		183, // Space

	},
	gotoRow{ // S170

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		246, // Comma
		183, // Space

	},
	gotoRow{ // S171

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		247, // Comma
		183, // Space

	},
	gotoRow{ // S172

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		248, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S173

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		250, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S175

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		251, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S176

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		252, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S177

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		253, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S178

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S181

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		260, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		258, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		255, // Space

	},
	gotoRow{ // S182

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S183

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S184

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		289, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S185

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S186

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S187

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S188

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		296, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		299, // CloseCurly
		-1,  // Comma
		292, // Space

	},
	gotoRow{ // S189

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S190

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S191

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		324, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S192

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		325, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S193

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S194

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		328, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		327, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		326, // Space

	},
	gotoRow{ // S195

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S196

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S197

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S198

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S199

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		333, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		334, // CloseCurly
		-1,  // Comma
		332, // Space

	},
	gotoRow{ // S200

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S202

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		336, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S203

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		337, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S204

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		338, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S205

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		339, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S206

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		340, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S207

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		341, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S208

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		342, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		56,  // Space

	},
	gotoRow{ // S209

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		343, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S210

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		344, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S211

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		73,  // Function
		74,  // List
		-1,  // Exprs
		345, // Expr
		75,  // ListType
		76,  // SpaceTerminal
		86,  // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		70,  // Space

	},
	gotoRow{ // S212

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		346, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S213

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		347, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S214

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		348, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S215

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		349, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S216

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		350, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S217

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		351, // Space

	},
	gotoRow{ // S218

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		353, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S219

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		354, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S220

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		355, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S221

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S222

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		356, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S224

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		357, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		45,  // Space

	},
	gotoRow{ // S225

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		358, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S226

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		359, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S227

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		360, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S228

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		361, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S229

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		362, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		56,  // Space

	},
	gotoRow{ // S230

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		363, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		50,  // Space

	},
	gotoRow{ // S231

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S232

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		73,  // Function
		74,  // List
		-1,  // Exprs
		364, // Expr
		75,  // ListType
		76,  // SpaceTerminal
		86,  // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		70,  // Space

	},
	gotoRow{ // S233

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		365, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S234

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		366, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S235

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		367, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S236

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		368, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S237

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		369, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S238

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		370, // Space

	},
	gotoRow{ // S239

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		372, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S240

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		373, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S241

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S242

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		374, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S244

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S245

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		375, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S246

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		376, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S247

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		377, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S248

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S249

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		378, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S250

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S251

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S252

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		380, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		379, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		255, // Space

	},
	gotoRow{ // S253

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		381, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		382, // CloseCurly
		-1,  // Comma
		292, // Space

	},
	gotoRow{ // S254

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		384, // ListType
		-1,  // SpaceTerminal
		385, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S256

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		388, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S257

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S259

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S260

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		391, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		389, // Space

	},
	gotoRow{ // S261

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		394, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S263

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S281

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S282

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S288

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S289

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		395, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S290

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		397, // ListType
		-1,  // SpaceTerminal
		398, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S293

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		401, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S294

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		404, // CloseCurly
		403, // Comma
		402, // Space

	},
	gotoRow{ // S297

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		406, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S299

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S307

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S308

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S309

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S310

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S311

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S312

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S313

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S314

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S315

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S316

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S317

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S318

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S319

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S320

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S321

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S323

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S324

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		408, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		407, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		326, // Space

	},
	gotoRow{ // S325

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		409, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		410, // CloseCurly
		-1,  // Comma
		332, // Space

	},
	gotoRow{ // S326

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		384, // ListType
		-1,  // SpaceTerminal
		385, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S327

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S328

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		413, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		412, // Space

	},
	gotoRow{ // S329

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S330

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S331

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S332

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		397, // ListType
		-1,  // SpaceTerminal
		398, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S333

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		416, // CloseCurly
		403, // Comma
		415, // Space

	},
	gotoRow{ // S334

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S335

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S336

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		73,  // Function
		74,  // List
		-1,  // Exprs
		417, // Expr
		75,  // ListType
		76,  // SpaceTerminal
		86,  // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		70,  // Space

	},
	gotoRow{ // S337

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		418, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S338

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		419, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S339

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		420, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S340

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		421, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S341

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		422, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S342

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		423, // Space

	},
	gotoRow{ // S343

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		425, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S344

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		426, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S345

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		183, // Space

	},
	gotoRow{ // S346

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S347

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		183, // Space

	},
	gotoRow{ // S348

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		183, // Space

	},
	gotoRow{ // S349

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		183, // Space

	},
	gotoRow{ // S350

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S351

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S352

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S353

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S354

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		195, // Space

	},
	gotoRow{ // S355

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		195, // Space

	},
	gotoRow{ // S356

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		73,  // Function
		74,  // List
		-1,  // Exprs
		439, // Expr
		75,  // ListType
		76,  // SpaceTerminal
		86,  // Terminal
		87,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		70,  // Space

	},
	gotoRow{ // S357

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		111, // Function
		112, // List
		-1,  // Exprs
		440, // Expr
		113, // ListType
		114, // SpaceTerminal
		115, // Terminal
		116, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		108, // Space

	},
	gotoRow{ // S358

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		441, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S359

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		442, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S360

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		443, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S361

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		444, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S362

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		445, // Space

	},
	gotoRow{ // S363

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		447, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S364

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		448, // Comma
		183, // Space

	},
	gotoRow{ // S365

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		450, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S366

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		451, // Comma
		183, // Space

	},
	gotoRow{ // S367

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		452, // Comma
		183, // Space

	},
	gotoRow{ // S368

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		453, // Comma
		183, // Space

	},
	gotoRow{ // S369

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		454, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S370

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		456, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S372

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		457, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S373

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		458, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S375

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		459, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S376

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		460, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S377

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		461, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		195, // Space

	},
	gotoRow{ // S378

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S380

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		462, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		389, // Space

	},
	gotoRow{ // S381

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		463, // CloseCurly
		403, // Comma
		402, // Space

	},
	gotoRow{ // S382

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S383

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		464, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S384

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		465, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S385

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S386

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S387

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S388

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		468, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		467, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		466, // Space

	},
	gotoRow{ // S389

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		-1,  // Exprs
		473, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		472, // Space

	},
	gotoRow{ // S391

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S392

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S394

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		475, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		476, // CloseCurly
		-1,  // Comma
		474, // Space

	},
	gotoRow{ // S395

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		478, // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		180, // Space

	},
	gotoRow{ // S397

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		479, // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		187, // Space

	},
	gotoRow{ // S398

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S399

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S400

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S401

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		482, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		481, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		480, // Space

	},
	gotoRow{ // S402

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S403

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		-1,  // Exprs
		486, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		485, // Space

	},
	gotoRow{ // S404

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S405

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S406

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		488, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		489, // CloseCurly
		-1,  // Comma
		487, // Space

	},
	gotoRow{ // S407

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S408

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		491, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		412, // Space

	},
	gotoRow{ // S409

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		492, // CloseCurly
		403, // Comma
		415, // Space

	},
	gotoRow{ // S410

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S411

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S412

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S414

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		493, // Comma
		183, // Space

	},
	gotoRow{ // S418

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		494, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S419

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		495, // Comma
		183, // Space

	},
	gotoRow{ // S420

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		496, // Comma
		183, // Space

	},
	gotoRow{ // S421

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		497, // Comma
		183, // Space

	},
	gotoRow{ // S422

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		428, // Space

	},
	gotoRow{ // S423

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S424

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S425

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		501, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S426

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S427

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		502, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S428

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		503, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S431

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		504, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S432

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		505, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S433

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S435

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		507, // Comma
		183, // Space

	},
	gotoRow{ // S440

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		508, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S441

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		509, // Comma
		183, // Space

	},
	gotoRow{ // S442

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		510, // Comma
		183, // Space

	},
	gotoRow{ // S443

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		511, // Comma
		183, // Space

	},
	gotoRow{ // S444

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		449, // Space

	},
	gotoRow{ // S445

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		514, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S447

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		515, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S448

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		516, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S449

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S450

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		517, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S452

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		518, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S453

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		519, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S454

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		449, // Space

	},
	gotoRow{ // S456

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S459

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S460

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S461

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S463

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		522, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		521, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		466, // Space

	},
	gotoRow{ // S465

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		523, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		524, // CloseCurly
		-1,  // Comma
		474, // Space

	},
	gotoRow{ // S466

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		384, // ListType
		-1,  // SpaceTerminal
		385, // Terminal
		265, // Bool
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		527, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		526, // Space

	},
	gotoRow{ // S469

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S470

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		384, // ListType
		-1,  // SpaceTerminal
		385, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S473

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S474

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		397, // ListType
		-1,  // SpaceTerminal
		398, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S475

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		530, // CloseCurly
		403, // Comma
		529, // Space

	},
	gotoRow{ // S476

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S477

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		259, // Function
		261, // List
		532, // Exprs
		257, // Expr
		262, // ListType
		263, // SpaceTerminal
		264, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		531, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		480, // Space

	},
	gotoRow{ // S479

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		295, // Function
		297, // List
		533, // Exprs
		294, // Expr
		298, // ListType
		300, // SpaceTerminal
		301, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		534, // CloseCurly
		-1,  // Comma
		487, // Space

	},
	gotoRow{ // S480

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		384, // ListType
		-1,  // SpaceTerminal
		385, // Terminal
		265, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S481

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S482

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		537, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		536, // Space

	},
	gotoRow{ // S483

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		397, // ListType
		-1,  // SpaceTerminal
		398, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S486

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		397, // ListType
		-1,  // SpaceTerminal
		398, // Terminal
		302, // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		-1,  // Space

	},
	gotoRow{ // S488

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		540, // CloseCurly
		403, // Comma
		539, // Space

	},
	gotoRow{ // S489

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S490

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S492

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		541, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S494

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S495

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		542, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S496

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		543, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S497

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		544, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S498

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		545, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S500

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S502

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		546, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S503

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		547, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S504

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		428, // Space

	},
	gotoRow{ // S505

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		428, // Space

	},
	gotoRow{ // S506

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		550, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S508

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S509

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		551, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S510

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		552, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S511

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		553, // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		151, // Space

	},
	gotoRow{ // S512

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		554, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S514

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S515

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		555, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S517

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		556, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S518

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		449, // Space

	},
	gotoRow{ // S519

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		449, // Space

	},
	gotoRow{ // S520

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		559, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		526, // Space

	},
	gotoRow{ // S523

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		560, // CloseCurly
		403, // Comma
		529, // Space

	},
	gotoRow{ // S524

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S526

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S527

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		561, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		390, // Comma
		536, // Space

	},
	gotoRow{ // S533

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		-1,  // CloseParen
		-1,  // OpenCurly
		562, // CloseCurly
		403, // Comma
		539, // Space

	},
	gotoRow{ // S534

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S536

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S537

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		563, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S542

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S543

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
	gotoRow{ // S544

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		566, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		428, // Space

	},
	gotoRow{ // S545

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S546

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S547

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S548

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
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
		449, // Space

	},
	gotoRow{ // S551

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		568, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S552

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		569, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S553

		-1,  // S'
		-1,  // Grammar
		-1,  // PatternDecls
		-1,  // PatternDecl
		-1,  // Pattern
		-1,  // Function
		-1,  // List
		-1,  // Exprs
		-1,  // Expr
		-1,  // ListType
		-1,  // SpaceTerminal
		-1,  // Terminal
		-1,  // Bool
		-1,  // Equal
		-1,  // OpenParen
		570, // CloseParen
		-1,  // OpenCurly
		-1,  // CloseCurly
		-1,  // Comma
		449, // Space

	},
	gotoRow{ // S554

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S555

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S556

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S557

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S564

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S567

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
		-1, // Bool
		-1, // Equal
		-1, // OpenParen
		-1, // CloseParen
		-1, // OpenCurly
		-1, // CloseCurly
		-1, // Comma
		-1, // Space

	},
	gotoRow{ // S569

		-1, // S'
		-1, // Grammar
		-1, // PatternDecls
		-1, // PatternDecl
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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
		-1, // Pattern
		-1, // Function
		-1, // List
		-1, // Exprs
		-1, // Expr
		-1, // ListType
		-1, // SpaceTerminal
		-1, // Terminal
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

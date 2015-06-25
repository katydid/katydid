package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(3), /* id */
			nil,      /* []bool */
			nil,      /* []int */
			nil,      /* []uint */
			nil,      /* []double */
			nil,      /* []string */
			nil,      /* [][]byte */
			nil,      /* int_lit */
			nil,      /* uint_lit */
			nil,      /* double_lit */
			nil,      /* string_lit */
			nil,      /* bytes_lit */
			nil,      /* bool_var */
			nil,      /* int_var */
			nil,      /* uint_var */
			nil,      /* double_var */
			nil,      /* string_var */
			nil,      /* bytes_var */
			nil,      /* true */
			nil,      /* false */
			nil,      /* = */
			nil,      /* ( */
			nil,      /* ) */
			nil,      /* { */
			nil,      /* } */
			nil,      /* , */
			shift(4), /* space */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* id */
			nil,          /* []bool */
			nil,          /* []int */
			nil,          /* []uint */
			nil,          /* []double */
			nil,          /* []string */
			nil,          /* [][]byte */
			nil,          /* int_lit */
			nil,          /* uint_lit */
			nil,          /* double_lit */
			nil,          /* string_lit */
			nil,          /* bytes_lit */
			nil,          /* bool_var */
			nil,          /* int_var */
			nil,          /* uint_var */
			nil,          /* double_var */
			nil,          /* string_var */
			nil,          /* bytes_var */
			nil,          /* true */
			nil,          /* false */
			nil,          /* = */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* { */
			nil,          /* } */
			nil,          /* , */
			nil,          /* space */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(5), /* id */
			nil,      /* []bool */
			nil,      /* []int */
			nil,      /* []uint */
			nil,      /* []double */
			nil,      /* []string */
			nil,      /* [][]byte */
			nil,      /* int_lit */
			nil,      /* uint_lit */
			nil,      /* double_lit */
			nil,      /* string_lit */
			nil,      /* bytes_lit */
			nil,      /* bool_var */
			nil,      /* int_var */
			nil,      /* uint_var */
			nil,      /* double_var */
			nil,      /* string_var */
			nil,      /* bytes_var */
			nil,      /* true */
			nil,      /* false */
			nil,      /* = */
			nil,      /* ( */
			nil,      /* ) */
			nil,      /* { */
			nil,      /* } */
			nil,      /* , */
			shift(6), /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(49), /* id, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(48), /* id, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(12), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(13), /* space */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(44), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(45), /* space */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(38), /* id, reduce: OpenParen */
			reduce(38), /* []bool, reduce: OpenParen */
			reduce(38), /* []int, reduce: OpenParen */
			reduce(38), /* []uint, reduce: OpenParen */
			reduce(38), /* []double, reduce: OpenParen */
			reduce(38), /* []string, reduce: OpenParen */
			reduce(38), /* [][]byte, reduce: OpenParen */
			reduce(38), /* int_lit, reduce: OpenParen */
			reduce(38), /* uint_lit, reduce: OpenParen */
			reduce(38), /* double_lit, reduce: OpenParen */
			reduce(38), /* string_lit, reduce: OpenParen */
			reduce(38), /* bytes_lit, reduce: OpenParen */
			reduce(38), /* bool_var, reduce: OpenParen */
			reduce(38), /* int_var, reduce: OpenParen */
			reduce(38), /* uint_var, reduce: OpenParen */
			reduce(38), /* double_var, reduce: OpenParen */
			reduce(38), /* string_var, reduce: OpenParen */
			reduce(38), /* bytes_var, reduce: OpenParen */
			reduce(38), /* true, reduce: OpenParen */
			reduce(38), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(38), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(38), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(49), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(44), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(45), /* space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(39), /* id, reduce: OpenParen */
			reduce(39), /* []bool, reduce: OpenParen */
			reduce(39), /* []int, reduce: OpenParen */
			reduce(39), /* []uint, reduce: OpenParen */
			reduce(39), /* []double, reduce: OpenParen */
			reduce(39), /* []string, reduce: OpenParen */
			reduce(39), /* [][]byte, reduce: OpenParen */
			reduce(39), /* int_lit, reduce: OpenParen */
			reduce(39), /* uint_lit, reduce: OpenParen */
			reduce(39), /* double_lit, reduce: OpenParen */
			reduce(39), /* string_lit, reduce: OpenParen */
			reduce(39), /* bytes_lit, reduce: OpenParen */
			reduce(39), /* bool_var, reduce: OpenParen */
			reduce(39), /* int_var, reduce: OpenParen */
			reduce(39), /* uint_var, reduce: OpenParen */
			reduce(39), /* double_var, reduce: OpenParen */
			reduce(39), /* string_var, reduce: OpenParen */
			reduce(39), /* bytes_var, reduce: OpenParen */
			reduce(39), /* true, reduce: OpenParen */
			reduce(39), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(39), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(48), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(12), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(12), /* ,, reduce: Expr */
			reduce(12), /* space, reduce: Expr */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(48), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(51), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(44), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(57), /* , */
			shift(58), /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Function */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(13), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(13), /* ,, reduce: Expr */
			reduce(13), /* space, reduce: Expr */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			shift(61), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(62), /* space */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(9), /* ), reduce: Exprs */
			nil,       /* { */
			nil,       /* } */
			reduce(9), /* ,, reduce: Exprs */
			reduce(9), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(11), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(11), /* ,, reduce: Expr */
			reduce(11), /* space, reduce: Expr */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(14), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(14), /* space, reduce: ListType */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(15), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(15), /* space, reduce: ListType */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(16), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(16), /* space, reduce: ListType */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(17), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(17), /* space, reduce: ListType */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(18), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(18), /* space, reduce: ListType */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(19), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(19), /* space, reduce: ListType */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(20), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(20), /* ,, reduce: SpaceTerminal */
			reduce(20), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(22), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(22), /* ,, reduce: Terminal */
			reduce(22), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(23), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(23), /* ,, reduce: Terminal */
			reduce(23), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(24), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(24), /* ,, reduce: Terminal */
			reduce(24), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(25), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(25), /* ,, reduce: Terminal */
			reduce(25), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(26), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(26), /* ,, reduce: Terminal */
			reduce(26), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(27), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(27), /* ,, reduce: Terminal */
			reduce(27), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(28), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(28), /* ,, reduce: Terminal */
			reduce(28), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(29), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(29), /* ,, reduce: Terminal */
			reduce(29), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(30), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(30), /* ,, reduce: Terminal */
			reduce(30), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(31), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(31), /* ,, reduce: Terminal */
			reduce(31), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(32), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(32), /* ,, reduce: Terminal */
			reduce(32), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(33), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(33), /* ,, reduce: Terminal */
			reduce(33), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(34), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(34), /* ,, reduce: Bool */
			reduce(34), /* space, reduce: Bool */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(35), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(35), /* ,, reduce: Bool */
			reduce(35), /* space, reduce: Bool */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(40), /* $, reduce: CloseParen */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(49), /* id, reduce: Space */
			reduce(49), /* []bool, reduce: Space */
			reduce(49), /* []int, reduce: Space */
			reduce(49), /* []uint, reduce: Space */
			reduce(49), /* []double, reduce: Space */
			reduce(49), /* []string, reduce: Space */
			reduce(49), /* [][]byte, reduce: Space */
			reduce(49), /* int_lit, reduce: Space */
			reduce(49), /* uint_lit, reduce: Space */
			reduce(49), /* double_lit, reduce: Space */
			reduce(49), /* string_lit, reduce: Space */
			reduce(49), /* bytes_lit, reduce: Space */
			reduce(49), /* bool_var, reduce: Space */
			reduce(49), /* int_var, reduce: Space */
			reduce(49), /* uint_var, reduce: Space */
			reduce(49), /* double_var, reduce: Space */
			reduce(49), /* string_var, reduce: Space */
			reduce(49), /* bytes_var, reduce: Space */
			reduce(49), /* true, reduce: Space */
			reduce(49), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(44), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(57), /* , */
			shift(58), /* space */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Function */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			shift(61), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(62), /* space */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(21), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(21), /* ,, reduce: SpaceTerminal */
			reduce(21), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(41), /* $, reduce: CloseParen */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(48), /* id, reduce: Space */
			reduce(48), /* []bool, reduce: Space */
			reduce(48), /* []int, reduce: Space */
			reduce(48), /* []uint, reduce: Space */
			reduce(48), /* []double, reduce: Space */
			reduce(48), /* []string, reduce: Space */
			reduce(48), /* [][]byte, reduce: Space */
			reduce(48), /* int_lit, reduce: Space */
			reduce(48), /* uint_lit, reduce: Space */
			reduce(48), /* double_lit, reduce: Space */
			reduce(48), /* string_lit, reduce: Space */
			reduce(48), /* bytes_lit, reduce: Space */
			reduce(48), /* bool_var, reduce: Space */
			reduce(48), /* int_var, reduce: Space */
			reduce(48), /* uint_var, reduce: Space */
			reduce(48), /* double_var, reduce: Space */
			reduce(48), /* string_var, reduce: Space */
			reduce(48), /* bytes_var, reduce: Space */
			reduce(48), /* true, reduce: Space */
			reduce(48), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(48), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(69), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(45), /* space */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(51), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(70), /* , */
			shift(71), /* space */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Function */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(74), /* space */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(46), /* id, reduce: Comma */
			reduce(46), /* []bool, reduce: Comma */
			reduce(46), /* []int, reduce: Comma */
			reduce(46), /* []uint, reduce: Comma */
			reduce(46), /* []double, reduce: Comma */
			reduce(46), /* []string, reduce: Comma */
			reduce(46), /* [][]byte, reduce: Comma */
			reduce(46), /* int_lit, reduce: Comma */
			reduce(46), /* uint_lit, reduce: Comma */
			reduce(46), /* double_lit, reduce: Comma */
			reduce(46), /* string_lit, reduce: Comma */
			reduce(46), /* bytes_lit, reduce: Comma */
			reduce(46), /* bool_var, reduce: Comma */
			reduce(46), /* int_var, reduce: Comma */
			reduce(46), /* uint_var, reduce: Comma */
			reduce(46), /* double_var, reduce: Comma */
			reduce(46), /* string_var, reduce: Comma */
			reduce(46), /* bytes_var, reduce: Comma */
			reduce(46), /* true, reduce: Comma */
			reduce(46), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(46), /* space, reduce: Comma */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(49), /* ,, reduce: Space */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			shift(75), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(76), /* space */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(101), /* } */
			nil,        /* , */
			shift(102), /* space */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(42), /* id, reduce: OpenCurly */
			reduce(42), /* []bool, reduce: OpenCurly */
			reduce(42), /* []int, reduce: OpenCurly */
			reduce(42), /* []uint, reduce: OpenCurly */
			reduce(42), /* []double, reduce: OpenCurly */
			reduce(42), /* []string, reduce: OpenCurly */
			reduce(42), /* [][]byte, reduce: OpenCurly */
			reduce(42), /* int_lit, reduce: OpenCurly */
			reduce(42), /* uint_lit, reduce: OpenCurly */
			reduce(42), /* double_lit, reduce: OpenCurly */
			reduce(42), /* string_lit, reduce: OpenCurly */
			reduce(42), /* bytes_lit, reduce: OpenCurly */
			reduce(42), /* bool_var, reduce: OpenCurly */
			reduce(42), /* int_var, reduce: OpenCurly */
			reduce(42), /* uint_var, reduce: OpenCurly */
			reduce(42), /* double_var, reduce: OpenCurly */
			reduce(42), /* string_var, reduce: OpenCurly */
			reduce(42), /* bytes_var, reduce: OpenCurly */
			reduce(42), /* true, reduce: OpenCurly */
			reduce(42), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(42), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(42), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(49), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Function */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int */
			shift(25), /* []uint */
			shift(26), /* []double */
			shift(27), /* []string */
			shift(28), /* [][]byte */
			shift(31), /* int_lit */
			shift(32), /* uint_lit */
			shift(33), /* double_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* bool_var */
			shift(37), /* int_var */
			shift(38), /* uint_var */
			shift(39), /* double_var */
			shift(40), /* string_var */
			shift(41), /* bytes_var */
			shift(42), /* true */
			shift(43), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(69), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(45), /* space */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(101), /* } */
			nil,        /* , */
			shift(102), /* space */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(48),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(31),  /* int_lit */
			shift(32),  /* uint_lit */
			shift(33),  /* double_lit */
			shift(34),  /* string_lit */
			shift(35),  /* bytes_lit */
			shift(36),  /* bool_var */
			shift(37),  /* int_var */
			shift(38),  /* uint_var */
			shift(39),  /* double_var */
			shift(40),  /* string_var */
			shift(41),  /* bytes_var */
			shift(42),  /* true */
			shift(43),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(107), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(52),  /* space */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(69), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(57), /* , */
			shift(58), /* space */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(4), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(4), /* ,, reduce: Function */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: CloseParen */
			reduce(40), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(47), /* id, reduce: Comma */
			reduce(47), /* []bool, reduce: Comma */
			reduce(47), /* []int, reduce: Comma */
			reduce(47), /* []uint, reduce: Comma */
			reduce(47), /* []double, reduce: Comma */
			reduce(47), /* []string, reduce: Comma */
			reduce(47), /* [][]byte, reduce: Comma */
			reduce(47), /* int_lit, reduce: Comma */
			reduce(47), /* uint_lit, reduce: Comma */
			reduce(47), /* double_lit, reduce: Comma */
			reduce(47), /* string_lit, reduce: Comma */
			reduce(47), /* bytes_lit, reduce: Comma */
			reduce(47), /* bool_var, reduce: Comma */
			reduce(47), /* int_var, reduce: Comma */
			reduce(47), /* uint_var, reduce: Comma */
			reduce(47), /* double_var, reduce: Comma */
			reduce(47), /* string_var, reduce: Comma */
			reduce(47), /* bytes_var, reduce: Comma */
			reduce(47), /* true, reduce: Comma */
			reduce(47), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(47), /* space, reduce: Comma */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(48), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(48), /* ,, reduce: Space */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(48),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(31),  /* int_lit */
			shift(32),  /* uint_lit */
			shift(33),  /* double_lit */
			shift(34),  /* string_lit */
			shift(35),  /* bytes_lit */
			shift(36),  /* bool_var */
			shift(37),  /* int_var */
			shift(38),  /* uint_var */
			shift(39),  /* double_var */
			shift(40),  /* string_var */
			shift(41),  /* bytes_var */
			shift(42),  /* true */
			shift(43),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(110), /* space */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(10), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(10), /* ,, reduce: Exprs */
			reduce(10), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(49), /* id, reduce: Space */
			reduce(49), /* []bool, reduce: Space */
			reduce(49), /* []int, reduce: Space */
			reduce(49), /* []uint, reduce: Space */
			reduce(49), /* []double, reduce: Space */
			reduce(49), /* []string, reduce: Space */
			reduce(49), /* [][]byte, reduce: Space */
			reduce(49), /* int_lit, reduce: Space */
			reduce(49), /* uint_lit, reduce: Space */
			reduce(49), /* double_lit, reduce: Space */
			reduce(49), /* string_lit, reduce: Space */
			reduce(49), /* bytes_lit, reduce: Space */
			reduce(49), /* bool_var, reduce: Space */
			reduce(49), /* int_var, reduce: Space */
			reduce(49), /* uint_var, reduce: Space */
			reduce(49), /* double_var, reduce: Space */
			reduce(49), /* string_var, reduce: Space */
			reduce(49), /* bytes_var, reduce: Space */
			reduce(49), /* true, reduce: Space */
			reduce(49), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(43), /* id, reduce: OpenCurly */
			reduce(43), /* []bool, reduce: OpenCurly */
			reduce(43), /* []int, reduce: OpenCurly */
			reduce(43), /* []uint, reduce: OpenCurly */
			reduce(43), /* []double, reduce: OpenCurly */
			reduce(43), /* []string, reduce: OpenCurly */
			reduce(43), /* [][]byte, reduce: OpenCurly */
			reduce(43), /* int_lit, reduce: OpenCurly */
			reduce(43), /* uint_lit, reduce: OpenCurly */
			reduce(43), /* double_lit, reduce: OpenCurly */
			reduce(43), /* string_lit, reduce: OpenCurly */
			reduce(43), /* bytes_lit, reduce: OpenCurly */
			reduce(43), /* bool_var, reduce: OpenCurly */
			reduce(43), /* int_var, reduce: OpenCurly */
			reduce(43), /* uint_var, reduce: OpenCurly */
			reduce(43), /* double_var, reduce: OpenCurly */
			reduce(43), /* string_var, reduce: OpenCurly */
			reduce(43), /* bytes_var, reduce: OpenCurly */
			reduce(43), /* true, reduce: OpenCurly */
			reduce(43), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(43), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(43), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(48), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(12), /* }, reduce: Expr */
			reduce(12), /* ,, reduce: Expr */
			reduce(12), /* space, reduce: Expr */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(111), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(114), /* } */
			nil,        /* , */
			shift(115), /* space */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(101), /* } */
			shift(57),  /* , */
			shift(120), /* space */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(13), /* }, reduce: Expr */
			reduce(13), /* ,, reduce: Expr */
			reduce(13), /* space, reduce: Expr */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			shift(61), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(62), /* space */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(8), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(8), /* ,, reduce: List */
			reduce(8), /* space, reduce: List */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(9), /* }, reduce: Exprs */
			reduce(9), /* ,, reduce: Exprs */
			reduce(9), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(11), /* }, reduce: Expr */
			reduce(11), /* ,, reduce: Expr */
			reduce(11), /* space, reduce: Expr */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(20), /* }, reduce: SpaceTerminal */
			reduce(20), /* ,, reduce: SpaceTerminal */
			reduce(20), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(22), /* }, reduce: Terminal */
			reduce(22), /* ,, reduce: Terminal */
			reduce(22), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(23), /* }, reduce: Terminal */
			reduce(23), /* ,, reduce: Terminal */
			reduce(23), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(24), /* }, reduce: Terminal */
			reduce(24), /* ,, reduce: Terminal */
			reduce(24), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(25), /* }, reduce: Terminal */
			reduce(25), /* ,, reduce: Terminal */
			reduce(25), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(26), /* }, reduce: Terminal */
			reduce(26), /* ,, reduce: Terminal */
			reduce(26), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(27), /* }, reduce: Terminal */
			reduce(27), /* ,, reduce: Terminal */
			reduce(27), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(28), /* }, reduce: Terminal */
			reduce(28), /* ,, reduce: Terminal */
			reduce(28), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(29), /* }, reduce: Terminal */
			reduce(29), /* ,, reduce: Terminal */
			reduce(29), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(30), /* }, reduce: Terminal */
			reduce(30), /* ,, reduce: Terminal */
			reduce(30), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(31), /* }, reduce: Terminal */
			reduce(31), /* ,, reduce: Terminal */
			reduce(31), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(32), /* }, reduce: Terminal */
			reduce(32), /* ,, reduce: Terminal */
			reduce(32), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(33), /* }, reduce: Terminal */
			reduce(33), /* ,, reduce: Terminal */
			reduce(33), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(34), /* }, reduce: Bool */
			reduce(34), /* ,, reduce: Bool */
			reduce(34), /* space, reduce: Bool */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(35), /* }, reduce: Bool */
			reduce(35), /* ,, reduce: Bool */
			reduce(35), /* space, reduce: Bool */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(44), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(44), /* ,, reduce: CloseCurly */
			reduce(44), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(49), /* id, reduce: Space */
			reduce(49), /* []bool, reduce: Space */
			reduce(49), /* []int, reduce: Space */
			reduce(49), /* []uint, reduce: Space */
			reduce(49), /* []double, reduce: Space */
			reduce(49), /* []string, reduce: Space */
			reduce(49), /* [][]byte, reduce: Space */
			reduce(49), /* int_lit, reduce: Space */
			reduce(49), /* uint_lit, reduce: Space */
			reduce(49), /* double_lit, reduce: Space */
			reduce(49), /* string_lit, reduce: Space */
			reduce(49), /* bytes_lit, reduce: Space */
			reduce(49), /* bool_var, reduce: Space */
			reduce(49), /* int_var, reduce: Space */
			reduce(49), /* uint_var, reduce: Space */
			reduce(49), /* double_var, reduce: Space */
			reduce(49), /* string_var, reduce: Space */
			reduce(49), /* bytes_var, reduce: Space */
			reduce(49), /* true, reduce: Space */
			reduce(49), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(49), /* }, reduce: Space */
			nil,        /* , */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(69), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(57), /* , */
			shift(58), /* space */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(2), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(2), /* ,, reduce: Function */
			reduce(2), /* space, reduce: Function */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(101), /* } */
			shift(57),  /* , */
			shift(120), /* space */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(7), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(7), /* ,, reduce: List */
			reduce(7), /* space, reduce: List */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(41), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(41), /* ,, reduce: CloseParen */
			reduce(41), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(107), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(70),  /* , */
			shift(71),  /* space */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(3), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(3), /* ,, reduce: Function */
			reduce(3), /* space, reduce: Function */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(48), /* id, reduce: Space */
			reduce(48), /* []bool, reduce: Space */
			reduce(48), /* []int, reduce: Space */
			reduce(48), /* []uint, reduce: Space */
			reduce(48), /* []double, reduce: Space */
			reduce(48), /* []string, reduce: Space */
			reduce(48), /* [][]byte, reduce: Space */
			reduce(48), /* int_lit, reduce: Space */
			reduce(48), /* uint_lit, reduce: Space */
			reduce(48), /* double_lit, reduce: Space */
			reduce(48), /* string_lit, reduce: Space */
			reduce(48), /* bytes_lit, reduce: Space */
			reduce(48), /* bool_var, reduce: Space */
			reduce(48), /* int_var, reduce: Space */
			reduce(48), /* uint_var, reduce: Space */
			reduce(48), /* double_var, reduce: Space */
			reduce(48), /* string_var, reduce: Space */
			reduce(48), /* bytes_var, reduce: Space */
			reduce(48), /* true, reduce: Space */
			reduce(48), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			shift(61), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(62), /* space */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(21), /* }, reduce: SpaceTerminal */
			reduce(21), /* ,, reduce: SpaceTerminal */
			reduce(21), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(45), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(45), /* ,, reduce: CloseCurly */
			reduce(45), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(48), /* id, reduce: Space */
			reduce(48), /* []bool, reduce: Space */
			reduce(48), /* []int, reduce: Space */
			reduce(48), /* []uint, reduce: Space */
			reduce(48), /* []double, reduce: Space */
			reduce(48), /* []string, reduce: Space */
			reduce(48), /* [][]byte, reduce: Space */
			reduce(48), /* int_lit, reduce: Space */
			reduce(48), /* uint_lit, reduce: Space */
			reduce(48), /* double_lit, reduce: Space */
			reduce(48), /* string_lit, reduce: Space */
			reduce(48), /* bytes_lit, reduce: Space */
			reduce(48), /* bool_var, reduce: Space */
			reduce(48), /* int_var, reduce: Space */
			reduce(48), /* uint_var, reduce: Space */
			reduce(48), /* double_var, reduce: Space */
			reduce(48), /* string_var, reduce: Space */
			reduce(48), /* bytes_var, reduce: Space */
			reduce(48), /* true, reduce: Space */
			reduce(48), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(48), /* }, reduce: Space */
			nil,        /* , */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(16),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(31),  /* int_lit */
			shift(32),  /* uint_lit */
			shift(33),  /* double_lit */
			shift(34),  /* string_lit */
			shift(35),  /* bytes_lit */
			shift(36),  /* bool_var */
			shift(37),  /* int_var */
			shift(38),  /* uint_var */
			shift(39),  /* double_var */
			shift(40),  /* string_var */
			shift(41),  /* bytes_var */
			shift(42),  /* true */
			shift(43),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(129), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(45),  /* space */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(114), /* } */
			shift(70),  /* , */
			shift(130), /* space */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(6), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(6), /* ,, reduce: List */
			reduce(6), /* space, reduce: List */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(74),  /* space */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(49), /* }, reduce: Space */
			reduce(49), /* ,, reduce: Space */
			reduce(49), /* space, reduce: Space */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(136), /* } */
			nil,        /* , */
			shift(102), /* space */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(1), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(1), /* ,, reduce: Function */
			reduce(1), /* space, reduce: Function */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			reduce(5), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(5), /* ,, reduce: List */
			reduce(5), /* space, reduce: List */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(16),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(31),  /* int_lit */
			shift(32),  /* uint_lit */
			shift(33),  /* double_lit */
			shift(34),  /* string_lit */
			shift(35),  /* bytes_lit */
			shift(36),  /* bool_var */
			shift(37),  /* int_var */
			shift(38),  /* uint_var */
			shift(39),  /* double_var */
			shift(40),  /* string_var */
			shift(41),  /* bytes_var */
			shift(42),  /* true */
			shift(43),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(129), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(45),  /* space */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(136), /* } */
			nil,        /* , */
			shift(102), /* space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(48),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(31),  /* int_lit */
			shift(32),  /* uint_lit */
			shift(33),  /* double_lit */
			shift(34),  /* string_lit */
			shift(35),  /* bytes_lit */
			shift(36),  /* bool_var */
			shift(37),  /* int_var */
			shift(38),  /* uint_var */
			shift(39),  /* double_var */
			shift(40),  /* string_var */
			shift(41),  /* bytes_var */
			shift(42),  /* true */
			shift(43),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(141), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(52),  /* space */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(129), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(57),  /* , */
			shift(58),  /* space */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(4), /* }, reduce: Function */
			reduce(4), /* ,, reduce: Function */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(40), /* }, reduce: CloseParen */
			reduce(40), /* ,, reduce: CloseParen */
			reduce(40), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(48), /* }, reduce: Space */
			reduce(48), /* ,, reduce: Space */
			reduce(48), /* space, reduce: Space */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(111), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(110), /* space */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(10), /* }, reduce: Exprs */
			reduce(10), /* ,, reduce: Exprs */
			reduce(10), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(111), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int */
			shift(25),  /* []uint */
			shift(26),  /* []double */
			shift(27),  /* []string */
			shift(28),  /* [][]byte */
			shift(88),  /* int_lit */
			shift(89),  /* uint_lit */
			shift(90),  /* double_lit */
			shift(91),  /* string_lit */
			shift(92),  /* bytes_lit */
			shift(93),  /* bool_var */
			shift(94),  /* int_var */
			shift(95),  /* uint_var */
			shift(96),  /* double_var */
			shift(97),  /* string_var */
			shift(98),  /* bytes_var */
			shift(99),  /* true */
			shift(100), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(144), /* } */
			nil,        /* , */
			shift(115), /* space */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(136), /* } */
			shift(57),  /* , */
			shift(120), /* space */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(8), /* }, reduce: List */
			reduce(8), /* ,, reduce: List */
			reduce(8), /* space, reduce: List */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(44), /* }, reduce: CloseCurly */
			reduce(44), /* ,, reduce: CloseCurly */
			reduce(44), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(129), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(57),  /* , */
			shift(58),  /* space */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(2), /* }, reduce: Function */
			reduce(2), /* ,, reduce: Function */
			reduce(2), /* space, reduce: Function */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(136), /* } */
			shift(57),  /* , */
			shift(120), /* space */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(7), /* }, reduce: List */
			reduce(7), /* ,, reduce: List */
			reduce(7), /* space, reduce: List */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(41), /* }, reduce: CloseParen */
			reduce(41), /* ,, reduce: CloseParen */
			reduce(41), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(141), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(70),  /* , */
			shift(71),  /* space */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(3), /* }, reduce: Function */
			reduce(3), /* ,, reduce: Function */
			reduce(3), /* space, reduce: Function */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(45), /* }, reduce: CloseCurly */
			reduce(45), /* ,, reduce: CloseCurly */
			reduce(45), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(144), /* } */
			shift(70),  /* , */
			shift(130), /* space */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(6), /* }, reduce: List */
			reduce(6), /* ,, reduce: List */
			reduce(6), /* space, reduce: List */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(1), /* }, reduce: Function */
			reduce(1), /* ,, reduce: Function */
			reduce(1), /* space, reduce: Function */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(5), /* }, reduce: List */
			reduce(5), /* ,, reduce: List */
			reduce(5), /* space, reduce: List */

		},
	},
}

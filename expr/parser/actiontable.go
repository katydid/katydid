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
			nil,       /* INVALID */
			nil,       /* $ */
			shift(6),  /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(16), /* int_lit */
			shift(17), /* uint_lit */
			shift(18), /* double_lit */
			shift(19), /* string_lit */
			shift(20), /* bytes_lit */
			shift(21), /* bool_var */
			shift(22), /* int_var */
			shift(23), /* uint_var */
			shift(24), /* double_var */
			shift(25), /* string_var */
			shift(26), /* bytes_var */
			shift(27), /* true */
			shift(28), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(29), /* space */

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
			nil,          /* ; */
			nil,          /* # */
			nil,          /* & */
			nil,          /* | */
			nil,          /* [ */
			nil,          /* ] */
			nil,          /* : */
			nil,          /* ! */
			nil,          /* space */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Expr */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Expr */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Expr */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(30), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(16), /* int_lit */
			shift(17), /* uint_lit */
			shift(18), /* double_lit */
			shift(19), /* string_lit */
			shift(20), /* bytes_lit */
			shift(21), /* bool_var */
			shift(22), /* int_var */
			shift(23), /* uint_var */
			shift(24), /* double_var */
			shift(25), /* string_var */
			shift(26), /* bytes_var */
			shift(27), /* true */
			shift(28), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(33), /* space */

		},
	},
	actionRow{ // S6
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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

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
			nil,       /* ( */
			nil,       /* ) */
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S8
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(14), /* space, reduce: ListType */

		},
	},
	actionRow{ // S9
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(15), /* space, reduce: ListType */

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
			nil,        /* ( */
			nil,        /* ) */
			reduce(16), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(16), /* space, reduce: ListType */

		},
	},
	actionRow{ // S11
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(17), /* space, reduce: ListType */

		},
	},
	actionRow{ // S12
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(18), /* space, reduce: ListType */

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
			nil,        /* ( */
			nil,        /* ) */
			reduce(19), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(19), /* space, reduce: ListType */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(20), /* $, reduce: SpaceTerminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(22), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(23), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(24), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(25), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(26), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(27), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(28), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(29), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(30), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(31), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(32), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(33), /* $, reduce: Terminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(34), /* $, reduce: Bool */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(35), /* $, reduce: Bool */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(65), /* id, reduce: Space */
			reduce(65), /* []bool, reduce: Space */
			reduce(65), /* []int, reduce: Space */
			reduce(65), /* []uint, reduce: Space */
			reduce(65), /* []double, reduce: Space */
			reduce(65), /* []string, reduce: Space */
			reduce(65), /* [][]byte, reduce: Space */
			reduce(65), /* int_lit, reduce: Space */
			reduce(65), /* uint_lit, reduce: Space */
			reduce(65), /* double_lit, reduce: Space */
			reduce(65), /* string_lit, reduce: Space */
			reduce(65), /* bytes_lit, reduce: Space */
			reduce(65), /* bool_var, reduce: Space */
			reduce(65), /* int_var, reduce: Space */
			reduce(65), /* uint_var, reduce: Space */
			reduce(65), /* double_var, reduce: Space */
			reduce(65), /* string_var, reduce: Space */
			reduce(65), /* bytes_var, reduce: Space */
			reduce(65), /* true, reduce: Space */
			reduce(65), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S30
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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

		},
	},
	actionRow{ // S31
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
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(21), /* $, reduce: SpaceTerminal */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(64), /* id, reduce: Space */
			reduce(64), /* []bool, reduce: Space */
			reduce(64), /* []int, reduce: Space */
			reduce(64), /* []uint, reduce: Space */
			reduce(64), /* []double, reduce: Space */
			reduce(64), /* []string, reduce: Space */
			reduce(64), /* [][]byte, reduce: Space */
			reduce(64), /* int_lit, reduce: Space */
			reduce(64), /* uint_lit, reduce: Space */
			reduce(64), /* double_lit, reduce: Space */
			reduce(64), /* string_lit, reduce: Space */
			reduce(64), /* bytes_lit, reduce: Space */
			reduce(64), /* bool_var, reduce: Space */
			reduce(64), /* int_var, reduce: Space */
			reduce(64), /* uint_var, reduce: Space */
			reduce(64), /* double_var, reduce: Space */
			reduce(64), /* string_var, reduce: Space */
			reduce(64), /* bytes_var, reduce: Space */
			reduce(64), /* true, reduce: Space */
			reduce(64), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S34
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
			shift(44), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(45), /* space */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(51), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(57), /* int_lit */
			shift(58), /* uint_lit */
			shift(59), /* double_lit */
			shift(60), /* string_lit */
			shift(61), /* bytes_lit */
			shift(62), /* bool_var */
			shift(63), /* int_var */
			shift(64), /* uint_var */
			shift(65), /* double_var */
			shift(66), /* string_var */
			shift(67), /* bytes_var */
			shift(68), /* true */
			shift(69), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(70), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(71), /* space */

		},
	},
	actionRow{ // S36
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: OpenParen */

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
			reduce(65), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S38
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
			shift(72), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(73), /* space */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(79), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(85), /* int_lit */
			shift(86), /* uint_lit */
			shift(87), /* double_lit */
			shift(88), /* string_lit */
			shift(89), /* bytes_lit */
			shift(90), /* bool_var */
			shift(91), /* int_var */
			shift(92), /* uint_var */
			shift(93), /* double_var */
			shift(94), /* string_var */
			shift(95), /* bytes_var */
			shift(96), /* true */
			shift(97), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			shift(98), /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(99), /* space */

		},
	},
	actionRow{ // S40
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(42), /* space, reduce: OpenCurly */

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
			nil,        /* ) */
			reduce(65), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(51), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(57), /* int_lit */
			shift(58), /* uint_lit */
			shift(59), /* double_lit */
			shift(60), /* string_lit */
			shift(61), /* bytes_lit */
			shift(62), /* bool_var */
			shift(63), /* int_var */
			shift(64), /* uint_var */
			shift(65), /* double_var */
			shift(66), /* string_var */
			shift(67), /* bytes_var */
			shift(68), /* true */
			shift(69), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(70), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(71), /* space */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(79), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(85), /* int_lit */
			shift(86), /* uint_lit */
			shift(87), /* double_lit */
			shift(88), /* string_lit */
			shift(89), /* bytes_lit */
			shift(90), /* bool_var */
			shift(91), /* int_var */
			shift(92), /* uint_var */
			shift(93), /* double_var */
			shift(94), /* string_var */
			shift(95), /* bytes_var */
			shift(96), /* true */
			shift(97), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			shift(98), /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(99), /* space */

		},
	},
	actionRow{ // S44
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S45
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
			reduce(64), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S46
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
			reduce(12), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(12), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(12), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S47
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
			reduce(1), /* ), reduce: Expr */
			nil,       /* { */
			nil,       /* } */
			reduce(1), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(1), /* space, reduce: Expr */

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
			nil,       /* ( */
			reduce(2), /* ), reduce: Expr */
			nil,       /* { */
			nil,       /* } */
			reduce(2), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(2), /* space, reduce: Expr */

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
			reduce(3), /* ), reduce: Expr */
			nil,       /* { */
			nil,       /* } */
			reduce(3), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(3), /* space, reduce: Expr */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(104), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(107), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(108), /* space */

		},
	},
	actionRow{ // S51
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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

		},
	},
	actionRow{ // S52
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
			shift(70),  /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: Function */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

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
			nil,       /* ) */
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S55
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(20), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S56
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(22), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S57
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(23), /* space, reduce: Terminal */

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
			reduce(24), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(24), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(24), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S59
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(25), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S60
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(26), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S61
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(27), /* space, reduce: Terminal */

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
			reduce(28), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(28), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(28), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S63
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(29), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S64
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S65
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S66
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S67
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S68
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Bool */

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
			reduce(35), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(35), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Bool */

		},
	},
	actionRow{ // S70
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(65), /* id, reduce: Space */
			reduce(65), /* []bool, reduce: Space */
			reduce(65), /* []int, reduce: Space */
			reduce(65), /* []uint, reduce: Space */
			reduce(65), /* []double, reduce: Space */
			reduce(65), /* []string, reduce: Space */
			reduce(65), /* [][]byte, reduce: Space */
			reduce(65), /* int_lit, reduce: Space */
			reduce(65), /* uint_lit, reduce: Space */
			reduce(65), /* double_lit, reduce: Space */
			reduce(65), /* string_lit, reduce: Space */
			reduce(65), /* bytes_lit, reduce: Space */
			reduce(65), /* bool_var, reduce: Space */
			reduce(65), /* int_var, reduce: Space */
			reduce(65), /* uint_var, reduce: Space */
			reduce(65), /* double_var, reduce: Space */
			reduce(65), /* string_var, reduce: Space */
			reduce(65), /* bytes_var, reduce: Space */
			reduce(65), /* true, reduce: Space */
			reduce(65), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(65), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S72
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(43), /* space, reduce: OpenCurly */

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
			nil,        /* ) */
			reduce(64), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S74
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
			reduce(12), /* }, reduce: Exprs */
			reduce(12), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(12), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S75
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
			reduce(1), /* }, reduce: Expr */
			reduce(1), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(1), /* space, reduce: Expr */

		},
	},
	actionRow{ // S76
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
			reduce(2), /* }, reduce: Expr */
			reduce(2), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(2), /* space, reduce: Expr */

		},
	},
	actionRow{ // S77
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
			reduce(3), /* }, reduce: Expr */
			reduce(3), /* ,, reduce: Expr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(3), /* space, reduce: Expr */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(116), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(119), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(120), /* space */

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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

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
			shift(98),  /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

		},
	},
	actionRow{ // S81
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
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: List */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S83
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(20), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S84
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(22), /* space, reduce: Terminal */

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
			reduce(23), /* }, reduce: Terminal */
			reduce(23), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(23), /* space, reduce: Terminal */

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
			reduce(24), /* }, reduce: Terminal */
			reduce(24), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(24), /* space, reduce: Terminal */

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
			reduce(25), /* }, reduce: Terminal */
			reduce(25), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(25), /* space, reduce: Terminal */

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
			reduce(26), /* }, reduce: Terminal */
			reduce(26), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(26), /* space, reduce: Terminal */

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
			reduce(27), /* }, reduce: Terminal */
			reduce(27), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(27), /* space, reduce: Terminal */

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
			reduce(28), /* }, reduce: Terminal */
			reduce(28), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(28), /* space, reduce: Terminal */

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
			reduce(29), /* }, reduce: Terminal */
			reduce(29), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(29), /* space, reduce: Terminal */

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
			reduce(30), /* }, reduce: Terminal */
			reduce(30), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Terminal */

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
			reduce(31), /* }, reduce: Terminal */
			reduce(31), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Terminal */

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
			reduce(32), /* }, reduce: Terminal */
			reduce(32), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Terminal */

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
			reduce(33), /* }, reduce: Terminal */
			reduce(33), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Terminal */

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
			reduce(34), /* }, reduce: Bool */
			reduce(34), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Bool */

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
			reduce(35), /* }, reduce: Bool */
			reduce(35), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Bool */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(44), /* $, reduce: CloseCurly */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(65), /* id, reduce: Space */
			reduce(65), /* []bool, reduce: Space */
			reduce(65), /* []int, reduce: Space */
			reduce(65), /* []uint, reduce: Space */
			reduce(65), /* []double, reduce: Space */
			reduce(65), /* []string, reduce: Space */
			reduce(65), /* [][]byte, reduce: Space */
			reduce(65), /* int_lit, reduce: Space */
			reduce(65), /* uint_lit, reduce: Space */
			reduce(65), /* double_lit, reduce: Space */
			reduce(65), /* string_lit, reduce: Space */
			reduce(65), /* bytes_lit, reduce: Space */
			reduce(65), /* bool_var, reduce: Space */
			reduce(65), /* int_var, reduce: Space */
			reduce(65), /* uint_var, reduce: Space */
			reduce(65), /* double_var, reduce: Space */
			reduce(65), /* string_var, reduce: Space */
			reduce(65), /* bytes_var, reduce: Space */
			reduce(65), /* true, reduce: Space */
			reduce(65), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(65), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

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
			shift(70),  /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Function */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S102
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
			shift(98),  /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: List */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

		},
	},
	actionRow{ // S105
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
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S106
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(21), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S107
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(64), /* id, reduce: Space */
			reduce(64), /* []bool, reduce: Space */
			reduce(64), /* []int, reduce: Space */
			reduce(64), /* []uint, reduce: Space */
			reduce(64), /* []double, reduce: Space */
			reduce(64), /* []string, reduce: Space */
			reduce(64), /* [][]byte, reduce: Space */
			reduce(64), /* int_lit, reduce: Space */
			reduce(64), /* uint_lit, reduce: Space */
			reduce(64), /* double_lit, reduce: Space */
			reduce(64), /* string_lit, reduce: Space */
			reduce(64), /* bytes_lit, reduce: Space */
			reduce(64), /* bool_var, reduce: Space */
			reduce(64), /* int_var, reduce: Space */
			reduce(64), /* uint_var, reduce: Space */
			reduce(64), /* double_var, reduce: Space */
			reduce(64), /* string_var, reduce: Space */
			reduce(64), /* bytes_var, reduce: Space */
			reduce(64), /* true, reduce: Space */
			reduce(64), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(64), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(51),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(134), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(71),  /* space */

		},
	},
	actionRow{ // S110
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
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(136), /* space */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Function */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(51), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(57), /* int_lit */
			shift(58), /* uint_lit */
			shift(59), /* double_lit */
			shift(60), /* string_lit */
			shift(61), /* bytes_lit */
			shift(62), /* bool_var */
			shift(63), /* int_var */
			shift(64), /* uint_var */
			shift(65), /* double_var */
			shift(66), /* string_var */
			shift(67), /* bytes_var */
			shift(68), /* true */
			shift(69), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(29), /* space */

		},
	},
	actionRow{ // S113
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(46), /* space, reduce: Comma */

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
			reduce(65), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(65), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(142), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(99),  /* space */

		},
	},
	actionRow{ // S116
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
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(37), /* space */

		},
	},
	actionRow{ // S117
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
			shift(40), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(41), /* space */

		},
	},
	actionRow{ // S118
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(21), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(45), /* $, reduce: CloseCurly */
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* space */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(64), /* id, reduce: Space */
			reduce(64), /* []bool, reduce: Space */
			reduce(64), /* []int, reduce: Space */
			reduce(64), /* []uint, reduce: Space */
			reduce(64), /* []double, reduce: Space */
			reduce(64), /* []string, reduce: Space */
			reduce(64), /* [][]byte, reduce: Space */
			reduce(64), /* int_lit, reduce: Space */
			reduce(64), /* uint_lit, reduce: Space */
			reduce(64), /* double_lit, reduce: Space */
			reduce(64), /* string_lit, reduce: Space */
			reduce(64), /* bytes_lit, reduce: Space */
			reduce(64), /* bool_var, reduce: Space */
			reduce(64), /* int_var, reduce: Space */
			reduce(64), /* uint_var, reduce: Space */
			reduce(64), /* double_var, reduce: Space */
			reduce(64), /* string_var, reduce: Space */
			reduce(64), /* bytes_var, reduce: Space */
			reduce(64), /* true, reduce: Space */
			reduce(64), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(64), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(51),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(71),  /* space */

		},
	},
	actionRow{ // S122
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
			shift(119), /* } */
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(149), /* space */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: List */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(79), /* id */
			shift(8),  /* []bool */
			shift(9),  /* []int */
			shift(10), /* []uint */
			shift(11), /* []double */
			shift(12), /* []string */
			shift(13), /* [][]byte */
			shift(85), /* int_lit */
			shift(86), /* uint_lit */
			shift(87), /* double_lit */
			shift(88), /* string_lit */
			shift(89), /* bytes_lit */
			shift(90), /* bool_var */
			shift(91), /* int_var */
			shift(92), /* uint_var */
			shift(93), /* double_var */
			shift(94), /* string_var */
			shift(95), /* bytes_var */
			shift(96), /* true */
			shift(97), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			shift(29), /* space */

		},
	},
	actionRow{ // S125
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
			reduce(65), /* }, reduce: Space */
			reduce(65), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(65), /* space, reduce: Space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(155), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(99),  /* space */

		},
	},
	actionRow{ // S127
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: List */
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* space */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(51),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(134), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(71),  /* space */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(142), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(99),  /* space */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(104), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(160), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(108), /* space */

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
			shift(134), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

		},
	},
	actionRow{ // S133
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
			reduce(7), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(7), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(7), /* space, reduce: Function */

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
			reduce(40), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S135
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(47), /* space, reduce: Comma */

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
			reduce(64), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(64), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

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
			reduce(13), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(13), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(13), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(104), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(33),  /* space */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(116), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(163), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(120), /* space */

		},
	},
	actionRow{ // S140
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
			shift(142), /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

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
			reduce(11), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(11), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(11), /* space, reduce: List */

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
			reduce(44), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(44), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(44), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(51),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(71),  /* space */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(79),  /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(155), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(99),  /* space */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(104), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(57),  /* int_lit */
			shift(58),  /* uint_lit */
			shift(59),  /* double_lit */
			shift(60),  /* string_lit */
			shift(61),  /* bytes_lit */
			shift(62),  /* bool_var */
			shift(63),  /* int_var */
			shift(64),  /* uint_var */
			shift(65),  /* double_var */
			shift(66),  /* string_var */
			shift(67),  /* bytes_var */
			shift(68),  /* true */
			shift(69),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(170), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(108), /* space */

		},
	},
	actionRow{ // S146
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
			shift(148), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

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
			reduce(7), /* }, reduce: Function */
			reduce(7), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(7), /* space, reduce: Function */

		},
	},
	actionRow{ // S148
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S149
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
			reduce(64), /* }, reduce: Space */
			reduce(64), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Space */

		},
	},
	actionRow{ // S150
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
			reduce(13), /* }, reduce: Exprs */
			reduce(13), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(13), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(116), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(33),  /* space */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(116), /* id */
			shift(8),   /* []bool */
			shift(9),   /* []int */
			shift(10),  /* []uint */
			shift(11),  /* []double */
			shift(12),  /* []string */
			shift(13),  /* [][]byte */
			shift(85),  /* int_lit */
			shift(86),  /* uint_lit */
			shift(87),  /* double_lit */
			shift(88),  /* string_lit */
			shift(89),  /* bytes_lit */
			shift(90),  /* bool_var */
			shift(91),  /* int_var */
			shift(92),  /* uint_var */
			shift(93),  /* double_var */
			shift(94),  /* string_var */
			shift(95),  /* bytes_var */
			shift(96),  /* true */
			shift(97),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(173), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(120), /* space */

		},
	},
	actionRow{ // S153
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
			shift(155), /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

		},
	},
	actionRow{ // S154
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
			reduce(11), /* }, reduce: List */
			reduce(11), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(11), /* space, reduce: List */

		},
	},
	actionRow{ // S155
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(44), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S156
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
			shift(134), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

		},
	},
	actionRow{ // S157
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
			reduce(5), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(5), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(5), /* space, reduce: Function */

		},
	},
	actionRow{ // S158
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
			shift(142), /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

		},
	},
	actionRow{ // S159
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
			reduce(10), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(10), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(10), /* space, reduce: List */

		},
	},
	actionRow{ // S160
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(41), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S161
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
			shift(160), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(136), /* space */

		},
	},
	actionRow{ // S162
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
			reduce(6), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(6), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(6), /* space, reduce: Function */

		},
	},
	actionRow{ // S163
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(45), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S164
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
			shift(163), /* } */
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(149), /* space */

		},
	},
	actionRow{ // S165
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
			reduce(9), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(9), /* ,, reduce: List */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(9), /* space, reduce: List */

		},
	},
	actionRow{ // S166
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
			shift(148), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(114), /* space */

		},
	},
	actionRow{ // S167
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
			reduce(5), /* }, reduce: Function */
			reduce(5), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(5), /* space, reduce: Function */

		},
	},
	actionRow{ // S168
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
			shift(155), /* } */
			shift(113), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(125), /* space */

		},
	},
	actionRow{ // S169
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
			reduce(10), /* }, reduce: List */
			reduce(10), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(10), /* space, reduce: List */

		},
	},
	actionRow{ // S170
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(41), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S171
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
			shift(170), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(136), /* space */

		},
	},
	actionRow{ // S172
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
			reduce(6), /* }, reduce: Function */
			reduce(6), /* ,, reduce: Function */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(6), /* space, reduce: Function */

		},
	},
	actionRow{ // S173
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
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(45), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S174
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
			shift(173), /* } */
			shift(135), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(149), /* space */

		},
	},
	actionRow{ // S175
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
			reduce(9), /* }, reduce: List */
			reduce(9), /* ,, reduce: List */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(9), /* space, reduce: List */

		},
	},
	actionRow{ // S176
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S177
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(8), /* space, reduce: List */

		},
	},
	actionRow{ // S178
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S179
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
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(8), /* space, reduce: List */

		},
	},
}

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
			shift(5), /* id */
			nil,      /* string_lit */
			nil,      /* []bool */
			nil,      /* []int */
			nil,      /* []uint */
			nil,      /* []double */
			nil,      /* []string */
			nil,      /* [][]byte */
			nil,      /* int_lit */
			nil,      /* uint_lit */
			nil,      /* double_lit */
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
			nil,      /* ; */
			nil,      /* # */
			nil,      /* & */
			nil,      /* | */
			nil,      /* [ */
			nil,      /* ] */
			nil,      /* : */
			nil,      /* ! */
			nil,      /* * */
			nil,      /* _ */
			nil,      /* ~ */
			nil,      /* . */
			shift(6), /* space */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* id */
			nil,          /* string_lit */
			nil,          /* []bool */
			nil,          /* []int */
			nil,          /* []uint */
			nil,          /* []double */
			nil,          /* []string */
			nil,          /* [][]byte */
			nil,          /* int_lit */
			nil,          /* uint_lit */
			nil,          /* double_lit */
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
			nil,          /* * */
			nil,          /* _ */
			nil,          /* ~ */
			nil,          /* . */
			nil,          /* space */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Grammar */
			shift(5),  /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(9),  /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(10), /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(11), /* space */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: PatternDecls */
			reduce(3), /* id, reduce: PatternDecls */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(3), /* space, reduce: PatternDecls */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(14), /* = */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(15), /* space */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(108), /* id, reduce: Space */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Grammar */
			shift(10), /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(16), /* space */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: PatternDecls */
			reduce(4), /* id, reduce: PatternDecls */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(4), /* space, reduce: PatternDecls */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(108), /* $, reduce: Space */
			reduce(108), /* id, reduce: Space */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(14), /* = */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(15), /* space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(107), /* id, reduce: Space */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
			nil,       /* bytes_lit */
			nil,       /* bool_var */
			nil,       /* int_var */
			nil,       /* uint_var */
			nil,       /* double_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(18), /* = */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(19), /* space */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(24), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(42), /* * */
			shift(43), /* _ */
			shift(44), /* ~ */
			shift(45), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(71), /* string_lit, reduce: Equal */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(71), /* (, reduce: Equal */
			nil,        /* ) */
			reduce(71), /* {, reduce: Equal */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(71), /* #, reduce: Equal */
			nil,        /* & */
			nil,        /* | */
			reduce(71), /* [, reduce: Equal */
			nil,        /* ] */
			nil,        /* : */
			reduce(71), /* !, reduce: Equal */
			reduce(71), /* *, reduce: Equal */
			reduce(71), /* _, reduce: Equal */
			reduce(71), /* ~, reduce: Equal */
			reduce(71), /* ., reduce: Equal */
			reduce(71), /* space, reduce: Equal */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			reduce(108), /* =, reduce: Space */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(107), /* $, reduce: Space */
			reduce(107), /* id, reduce: Space */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(24), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(42), /* * */
			shift(43), /* _ */
			shift(44), /* ~ */
			shift(45), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(72), /* string_lit, reduce: Equal */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(72), /* (, reduce: Equal */
			nil,        /* ) */
			reduce(72), /* {, reduce: Equal */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(72), /* #, reduce: Equal */
			nil,        /* & */
			nil,        /* | */
			reduce(72), /* [, reduce: Equal */
			nil,        /* ] */
			nil,        /* : */
			reduce(72), /* !, reduce: Equal */
			reduce(72), /* *, reduce: Equal */
			reduce(72), /* _, reduce: Equal */
			reduce(72), /* ~, reduce: Equal */
			reduce(72), /* ., reduce: Equal */
			reduce(72), /* space, reduce: Equal */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			reduce(107), /* =, reduce: Space */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(48), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(49), /* ( */
			nil,       /* ) */
			shift(50), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(51), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(52), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(53), /* ! */
			shift(54), /* * */
			shift(55), /* _ */
			shift(56), /* ~ */
			shift(57), /* . */
			shift(58), /* space */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: PatternDecl */
			reduce(6), /* id, reduce: PatternDecl */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(6), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(15), /* $, reduce: Pattern */
			reduce(15), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(16), /* $, reduce: Pattern */
			reduce(16), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(17), /* $, reduce: Pattern */
			reduce(17), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(20), /* $, reduce: Pattern */
			reduce(20), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(21), /* $, reduce: Pattern */
			reduce(21), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(22), /* $, reduce: Pattern */
			reduce(22), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(116), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(73), /* string_lit, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(73), /* (, reduce: OpenParen */
			nil,        /* ) */
			reduce(73), /* {, reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(73), /* #, reduce: OpenParen */
			nil,        /* & */
			nil,        /* | */
			reduce(73), /* [, reduce: OpenParen */
			nil,        /* ] */
			nil,        /* : */
			reduce(73), /* !, reduce: OpenParen */
			reduce(73), /* *, reduce: OpenParen */
			reduce(73), /* _, reduce: OpenParen */
			reduce(73), /* ~, reduce: OpenParen */
			reduce(73), /* ., reduce: OpenParen */
			reduce(73), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(77), /* id, reduce: OpenCurly */
			reduce(77), /* string_lit, reduce: OpenCurly */
			reduce(77), /* []bool, reduce: OpenCurly */
			reduce(77), /* []int, reduce: OpenCurly */
			reduce(77), /* []uint, reduce: OpenCurly */
			reduce(77), /* []double, reduce: OpenCurly */
			reduce(77), /* []string, reduce: OpenCurly */
			reduce(77), /* [][]byte, reduce: OpenCurly */
			reduce(77), /* int_lit, reduce: OpenCurly */
			reduce(77), /* uint_lit, reduce: OpenCurly */
			reduce(77), /* double_lit, reduce: OpenCurly */
			reduce(77), /* bytes_lit, reduce: OpenCurly */
			reduce(77), /* bool_var, reduce: OpenCurly */
			reduce(77), /* int_var, reduce: OpenCurly */
			reduce(77), /* uint_var, reduce: OpenCurly */
			reduce(77), /* double_var, reduce: OpenCurly */
			reduce(77), /* string_var, reduce: OpenCurly */
			reduce(77), /* bytes_var, reduce: OpenCurly */
			reduce(77), /* true, reduce: OpenCurly */
			reduce(77), /* false, reduce: OpenCurly */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(77), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(85), /* id, reduce: HashTag */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(91), /* string_lit, reduce: OpenBracket */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(91), /* (, reduce: OpenBracket */
			nil,        /* ) */
			reduce(91), /* {, reduce: OpenBracket */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(91), /* #, reduce: OpenBracket */
			nil,        /* & */
			nil,        /* | */
			reduce(91), /* [, reduce: OpenBracket */
			nil,        /* ] */
			nil,        /* : */
			reduce(91), /* !, reduce: OpenBracket */
			reduce(91), /* *, reduce: OpenBracket */
			reduce(91), /* _, reduce: OpenBracket */
			reduce(91), /* ~, reduce: OpenBracket */
			reduce(91), /* ., reduce: OpenBracket */
			reduce(91), /* space, reduce: OpenBracket */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(97), /* (, reduce: Exclamation */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(97), /* space, reduce: Exclamation */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(99), /* $, reduce: Star */
			reduce(99), /* id, reduce: Star */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(101), /* $, reduce: Underscore */
			reduce(101), /* id, reduce: Underscore */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(103), /* $, reduce: Tilde */
			reduce(103), /* id, reduce: Tilde */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(105), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			reduce(108), /* string_lit, reduce: Space */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(108), /* (, reduce: Space */
			nil,         /* ) */
			reduce(108), /* {, reduce: Space */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			reduce(108), /* #, reduce: Space */
			nil,         /* & */
			nil,         /* | */
			reduce(108), /* [, reduce: Space */
			nil,         /* ] */
			nil,         /* : */
			reduce(108), /* !, reduce: Space */
			reduce(108), /* *, reduce: Space */
			reduce(108), /* _, reduce: Space */
			reduce(108), /* ~, reduce: Space */
			reduce(108), /* ., reduce: Space */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: PatternDecl */
			reduce(5), /* id, reduce: PatternDecl */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(5), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(74), /* string_lit, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(74), /* (, reduce: OpenParen */
			nil,        /* ) */
			reduce(74), /* {, reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(74), /* #, reduce: OpenParen */
			nil,        /* & */
			nil,        /* | */
			reduce(74), /* [, reduce: OpenParen */
			nil,        /* ] */
			nil,        /* : */
			reduce(74), /* !, reduce: OpenParen */
			reduce(74), /* *, reduce: OpenParen */
			reduce(74), /* _, reduce: OpenParen */
			reduce(74), /* ~, reduce: OpenParen */
			reduce(74), /* ., reduce: OpenParen */
			reduce(74), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(78), /* id, reduce: OpenCurly */
			reduce(78), /* string_lit, reduce: OpenCurly */
			reduce(78), /* []bool, reduce: OpenCurly */
			reduce(78), /* []int, reduce: OpenCurly */
			reduce(78), /* []uint, reduce: OpenCurly */
			reduce(78), /* []double, reduce: OpenCurly */
			reduce(78), /* []string, reduce: OpenCurly */
			reduce(78), /* [][]byte, reduce: OpenCurly */
			reduce(78), /* int_lit, reduce: OpenCurly */
			reduce(78), /* uint_lit, reduce: OpenCurly */
			reduce(78), /* double_lit, reduce: OpenCurly */
			reduce(78), /* bytes_lit, reduce: OpenCurly */
			reduce(78), /* bool_var, reduce: OpenCurly */
			reduce(78), /* int_var, reduce: OpenCurly */
			reduce(78), /* uint_var, reduce: OpenCurly */
			reduce(78), /* double_var, reduce: OpenCurly */
			reduce(78), /* string_var, reduce: OpenCurly */
			reduce(78), /* bytes_var, reduce: OpenCurly */
			reduce(78), /* true, reduce: OpenCurly */
			reduce(78), /* false, reduce: OpenCurly */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(78), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(86), /* id, reduce: HashTag */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(92), /* string_lit, reduce: OpenBracket */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(92), /* (, reduce: OpenBracket */
			nil,        /* ) */
			reduce(92), /* {, reduce: OpenBracket */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(92), /* #, reduce: OpenBracket */
			nil,        /* & */
			nil,        /* | */
			reduce(92), /* [, reduce: OpenBracket */
			nil,        /* ] */
			nil,        /* : */
			reduce(92), /* !, reduce: OpenBracket */
			reduce(92), /* *, reduce: OpenBracket */
			reduce(92), /* _, reduce: OpenBracket */
			reduce(92), /* ~, reduce: OpenBracket */
			reduce(92), /* ., reduce: OpenBracket */
			reduce(92), /* space, reduce: OpenBracket */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(98), /* (, reduce: Exclamation */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(98), /* space, reduce: Exclamation */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(100), /* $, reduce: Star */
			reduce(100), /* id, reduce: Star */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(102), /* $, reduce: Underscore */
			reduce(102), /* id, reduce: Underscore */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			reduce(104), /* $, reduce: Tilde */
			reduce(104), /* id, reduce: Tilde */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(106), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			reduce(107), /* string_lit, reduce: Space */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(107), /* (, reduce: Space */
			nil,         /* ) */
			reduce(107), /* {, reduce: Space */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			reduce(107), /* #, reduce: Space */
			nil,         /* & */
			nil,         /* | */
			reduce(107), /* [, reduce: Space */
			nil,         /* ] */
			nil,         /* : */
			reduce(107), /* !, reduce: Space */
			reduce(107), /* *, reduce: Space */
			reduce(107), /* _, reduce: Space */
			reduce(107), /* ~, reduce: Space */
			reduce(107), /* ., reduce: Space */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(134), /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(135), /* space */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(24), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(42), /* * */
			shift(43), /* _ */
			shift(44), /* ~ */
			shift(45), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(95), /* string_lit, reduce: Colon */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(95), /* (, reduce: Colon */
			nil,        /* ) */
			reduce(95), /* {, reduce: Colon */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(95), /* #, reduce: Colon */
			nil,        /* & */
			nil,        /* | */
			reduce(95), /* [, reduce: Colon */
			nil,        /* ] */
			nil,        /* : */
			reduce(95), /* !, reduce: Colon */
			reduce(95), /* *, reduce: Colon */
			reduce(95), /* _, reduce: Colon */
			reduce(95), /* ~, reduce: Colon */
			reduce(95), /* ., reduce: Colon */
			reduce(95), /* space, reduce: Colon */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(108), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(137), /* space */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(108), /* (, reduce: Space */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(159), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(160), /* * */
			shift(161), /* _ */
			shift(162), /* ~ */
			shift(163), /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(176), /* space */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			reduce(7), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			reduce(9), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			reduce(11), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(15), /* &, reduce: Pattern */
			reduce(15), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(16), /* &, reduce: Pattern */
			reduce(16), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(17), /* &, reduce: Pattern */
			reduce(17), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(20), /* &, reduce: Pattern */
			reduce(20), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(21), /* &, reduce: Pattern */
			reduce(21), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(22), /* &, reduce: Pattern */
			reduce(22), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(181), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ), reduce: Star */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(99), /* &, reduce: Star */
			reduce(99), /* |, reduce: Star */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(101), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(101), /* &, reduce: Underscore */
			reduce(101), /* |, reduce: Underscore */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(103), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(103), /* &, reduce: Tilde */
			reduce(103), /* |, reduce: Tilde */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(105), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			reduce(105), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(183), /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(186), /* space */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(61), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(192), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(36), /* }, reduce: Expr */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(36), /* space, reduce: Expr */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(37), /* }, reduce: Expr */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(37), /* space, reduce: Expr */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(38), /* }, reduce: Expr */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(49), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(49), /* space, reduce: ListType */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(50), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(50), /* space, reduce: ListType */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(51), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(51), /* space, reduce: ListType */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(52), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(52), /* space, reduce: ListType */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(53), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(53), /* space, reduce: ListType */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(54), /* {, reduce: ListType */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(54), /* space, reduce: ListType */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(55), /* }, reduce: SpaceTerminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(55), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(57), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(58), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(59), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(60), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(62), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(63), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(64), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(65), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(66), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(67), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(68), /* }, reduce: Terminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(69), /* }, reduce: Bool */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(69), /* space, reduce: Bool */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(70), /* }, reduce: Bool */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(70), /* space, reduce: Bool */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(108), /* id, reduce: Space */
			reduce(108), /* string_lit, reduce: Space */
			reduce(108), /* []bool, reduce: Space */
			reduce(108), /* []int, reduce: Space */
			reduce(108), /* []uint, reduce: Space */
			reduce(108), /* []double, reduce: Space */
			reduce(108), /* []string, reduce: Space */
			reduce(108), /* [][]byte, reduce: Space */
			reduce(108), /* int_lit, reduce: Space */
			reduce(108), /* uint_lit, reduce: Space */
			reduce(108), /* double_lit, reduce: Space */
			reduce(108), /* bytes_lit, reduce: Space */
			reduce(108), /* bool_var, reduce: Space */
			reduce(108), /* int_var, reduce: Space */
			reduce(108), /* uint_var, reduce: Space */
			reduce(108), /* double_var, reduce: Space */
			reduce(108), /* string_var, reduce: Space */
			reduce(108), /* bytes_var, reduce: Space */
			reduce(108), /* true, reduce: Space */
			reduce(108), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(25), /* $, reduce: Pattern */
			reduce(25), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(198), /* * */
			shift(199), /* _ */
			shift(200), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(209), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ,, reduce: Star */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(101), /* ,, reduce: Underscore */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(103), /* ,, reduce: Tilde */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(96), /* string_lit, reduce: Colon */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(96), /* (, reduce: Colon */
			nil,        /* ) */
			reduce(96), /* {, reduce: Colon */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(96), /* #, reduce: Colon */
			nil,        /* & */
			nil,        /* | */
			reduce(96), /* [, reduce: Colon */
			nil,        /* ] */
			nil,        /* : */
			reduce(96), /* !, reduce: Colon */
			reduce(96), /* *, reduce: Colon */
			reduce(96), /* _, reduce: Colon */
			reduce(96), /* ~, reduce: Colon */
			reduce(96), /* ., reduce: Colon */
			reduce(96), /* space, reduce: Colon */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(107), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(18), /* $, reduce: Pattern */
			reduce(18), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(107), /* (, reduce: Space */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(211), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(212), /* * */
			shift(213), /* _ */
			shift(214), /* ~ */
			shift(215), /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(218), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(223), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(224), /* space */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			reduce(7), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			reduce(9), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* ), reduce: NameExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			reduce(11), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S150
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(229), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S154
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ), reduce: Star */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(101), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S157
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(103), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S158
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(105), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(105), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			reduce(8), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(100), /* ), reduce: Star */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(100), /* &, reduce: Star */
			reduce(100), /* |, reduce: Star */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(102), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(102), /* &, reduce: Underscore */
			reduce(102), /* |, reduce: Underscore */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(104), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(104), /* &, reduce: Tilde */
			reduce(104), /* |, reduce: Tilde */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(106), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			reduce(106), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(231), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(234), /* space */

		},
	},
	actionRow{ // S165
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(42),  /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S166
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S167
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S168
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* *, reduce: CloseParen */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S169
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(87), /* string_lit, reduce: Ampersand */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(87), /* (, reduce: Ampersand */
			nil,        /* ) */
			reduce(87), /* {, reduce: Ampersand */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(87), /* #, reduce: Ampersand */
			nil,        /* & */
			nil,        /* | */
			reduce(87), /* [, reduce: Ampersand */
			nil,        /* ] */
			nil,        /* : */
			reduce(87), /* !, reduce: Ampersand */
			reduce(87), /* *, reduce: Ampersand */
			reduce(87), /* _, reduce: Ampersand */
			reduce(87), /* ~, reduce: Ampersand */
			reduce(87), /* ., reduce: Ampersand */
			reduce(87), /* space, reduce: Ampersand */

		},
	},
	actionRow{ // S170
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(89), /* string_lit, reduce: Pipe */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(89), /* (, reduce: Pipe */
			nil,        /* ) */
			reduce(89), /* {, reduce: Pipe */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* #, reduce: Pipe */
			nil,        /* & */
			nil,        /* | */
			reduce(89), /* [, reduce: Pipe */
			nil,        /* ] */
			nil,        /* : */
			reduce(89), /* !, reduce: Pipe */
			reduce(89), /* *, reduce: Pipe */
			reduce(89), /* _, reduce: Pipe */
			reduce(89), /* ~, reduce: Pipe */
			reduce(89), /* ., reduce: Pipe */
			reduce(89), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S171
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(108), /* &, reduce: Space */
			reduce(108), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S172
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(134), /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(275), /* space */

		},
	},
	actionRow{ // S173
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S174
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(24), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(45), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S175
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(89), /* string_lit, reduce: Pipe */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(89), /* (, reduce: Pipe */
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
			reduce(89), /* !, reduce: Pipe */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			reduce(89), /* ., reduce: Pipe */
			reduce(89), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S176
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(108), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			reduce(108), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S177
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S178
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S179
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(176), /* space */

		},
	},
	actionRow{ // S180
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(300), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S181
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(25), /* &, reduce: Pattern */
			reduce(25), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S182
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S183
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S184
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S185
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(56), /* }, reduce: SpaceTerminal */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(56), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S186
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(107), /* id, reduce: Space */
			reduce(107), /* string_lit, reduce: Space */
			reduce(107), /* []bool, reduce: Space */
			reduce(107), /* []int, reduce: Space */
			reduce(107), /* []uint, reduce: Space */
			reduce(107), /* []double, reduce: Space */
			reduce(107), /* []string, reduce: Space */
			reduce(107), /* [][]byte, reduce: Space */
			reduce(107), /* int_lit, reduce: Space */
			reduce(107), /* uint_lit, reduce: Space */
			reduce(107), /* double_lit, reduce: Space */
			reduce(107), /* bytes_lit, reduce: Space */
			reduce(107), /* bool_var, reduce: Space */
			reduce(107), /* int_var, reduce: Space */
			reduce(107), /* uint_var, reduce: Space */
			reduce(107), /* double_var, reduce: Space */
			reduce(107), /* string_var, reduce: Space */
			reduce(107), /* bytes_var, reduce: Space */
			reduce(107), /* true, reduce: Space */
			reduce(107), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S187
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(304), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(137), /* space */

		},
	},
	actionRow{ // S188
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(329), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S189
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(73), /* id, reduce: OpenParen */
			reduce(73), /* string_lit, reduce: OpenParen */
			reduce(73), /* []bool, reduce: OpenParen */
			reduce(73), /* []int, reduce: OpenParen */
			reduce(73), /* []uint, reduce: OpenParen */
			reduce(73), /* []double, reduce: OpenParen */
			reduce(73), /* []string, reduce: OpenParen */
			reduce(73), /* [][]byte, reduce: OpenParen */
			reduce(73), /* int_lit, reduce: OpenParen */
			reduce(73), /* uint_lit, reduce: OpenParen */
			reduce(73), /* double_lit, reduce: OpenParen */
			reduce(73), /* bytes_lit, reduce: OpenParen */
			reduce(73), /* bool_var, reduce: OpenParen */
			reduce(73), /* int_var, reduce: OpenParen */
			reduce(73), /* uint_var, reduce: OpenParen */
			reduce(73), /* double_var, reduce: OpenParen */
			reduce(73), /* string_var, reduce: OpenParen */
			reduce(73), /* bytes_var, reduce: OpenParen */
			reduce(73), /* true, reduce: OpenParen */
			reduce(73), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(73), /* ), reduce: OpenParen */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(73), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S190
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(331), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S191
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(19), /* $, reduce: Pattern */
			reduce(19), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S192
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(79), /* $, reduce: CloseCurly */
			reduce(79), /* id, reduce: CloseCurly */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S193
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(108), /* }, reduce: Space */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S194
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(333), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(334), /* space */

		},
	},
	actionRow{ // S195
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(359), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S196
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(77), /* id, reduce: OpenCurly */
			reduce(77), /* string_lit, reduce: OpenCurly */
			reduce(77), /* []bool, reduce: OpenCurly */
			reduce(77), /* []int, reduce: OpenCurly */
			reduce(77), /* []uint, reduce: OpenCurly */
			reduce(77), /* []double, reduce: OpenCurly */
			reduce(77), /* []string, reduce: OpenCurly */
			reduce(77), /* [][]byte, reduce: OpenCurly */
			reduce(77), /* int_lit, reduce: OpenCurly */
			reduce(77), /* uint_lit, reduce: OpenCurly */
			reduce(77), /* double_lit, reduce: OpenCurly */
			reduce(77), /* bytes_lit, reduce: OpenCurly */
			reduce(77), /* bool_var, reduce: OpenCurly */
			reduce(77), /* int_var, reduce: OpenCurly */
			reduce(77), /* uint_var, reduce: OpenCurly */
			reduce(77), /* double_var, reduce: OpenCurly */
			reduce(77), /* string_var, reduce: OpenCurly */
			reduce(77), /* bytes_var, reduce: OpenCurly */
			reduce(77), /* true, reduce: OpenCurly */
			reduce(77), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(77), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(77), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S197
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			reduce(108), /* {, reduce: Space */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S198
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(100), /* ,, reduce: Star */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S199
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(102), /* ,, reduce: Underscore */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S200
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(104), /* ,, reduce: Tilde */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S201
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(362), /* space */

		},
	},
	actionRow{ // S202
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S203
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(81), /* string_lit, reduce: Comma */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(81), /* (, reduce: Comma */
			nil,        /* ) */
			reduce(81), /* {, reduce: Comma */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(81), /* #, reduce: Comma */
			nil,        /* & */
			nil,        /* | */
			reduce(81), /* [, reduce: Comma */
			nil,        /* ] */
			nil,        /* : */
			reduce(81), /* !, reduce: Comma */
			reduce(81), /* *, reduce: Comma */
			reduce(81), /* _, reduce: Comma */
			reduce(81), /* ~, reduce: Comma */
			reduce(81), /* ., reduce: Comma */
			reduce(81), /* space, reduce: Comma */

		},
	},
	actionRow{ // S204
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(108), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S205
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S206
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S207
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S208
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(388), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S209
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S210
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S211
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			reduce(8), /* :, reduce: NameExpr */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S212
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(100), /* ), reduce: Star */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S213
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(102), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S214
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(104), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S215
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(106), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(106), /* :, reduce: Dot */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S216
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(390), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S217
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(26), /* $, reduce: Pattern */
			reduce(26), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S218
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(75), /* $, reduce: CloseParen */
			reduce(75), /* id, reduce: CloseParen */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S219
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S220
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(392), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(134), /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(393), /* space */

		},
	},
	actionRow{ // S221
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S222
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S223
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S224
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(108), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S225
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S226
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S227
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(176), /* space */

		},
	},
	actionRow{ // S228
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(407), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S229
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S230
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S231
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* *, reduce: CloseParen */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S232
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(88), /* string_lit, reduce: Ampersand */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(88), /* (, reduce: Ampersand */
			nil,        /* ) */
			reduce(88), /* {, reduce: Ampersand */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* #, reduce: Ampersand */
			nil,        /* & */
			nil,        /* | */
			reduce(88), /* [, reduce: Ampersand */
			nil,        /* ] */
			nil,        /* : */
			reduce(88), /* !, reduce: Ampersand */
			reduce(88), /* *, reduce: Ampersand */
			reduce(88), /* _, reduce: Ampersand */
			reduce(88), /* ~, reduce: Ampersand */
			reduce(88), /* ., reduce: Ampersand */
			reduce(88), /* space, reduce: Ampersand */

		},
	},
	actionRow{ // S233
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(90), /* string_lit, reduce: Pipe */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(90), /* (, reduce: Pipe */
			nil,        /* ) */
			reduce(90), /* {, reduce: Pipe */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(90), /* #, reduce: Pipe */
			nil,        /* & */
			nil,        /* | */
			reduce(90), /* [, reduce: Pipe */
			nil,        /* ] */
			nil,        /* : */
			reduce(90), /* !, reduce: Pipe */
			reduce(90), /* *, reduce: Pipe */
			reduce(90), /* _, reduce: Pipe */
			reduce(90), /* ~, reduce: Pipe */
			reduce(90), /* ., reduce: Pipe */
			reduce(90), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S234
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(107), /* &, reduce: Space */
			reduce(107), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S235
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(54),  /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S236
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(23), /* $, reduce: Pattern */
			reduce(23), /* id, reduce: Pattern */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S237
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			reduce(108), /* *, reduce: Space */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S238
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(410), /* * */
			shift(411), /* _ */
			shift(412), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S239
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(31), /* ), reduce: ContinueOr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(31), /* |, reduce: ContinueOr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(31), /* space, reduce: ContinueOr */

		},
	},
	actionRow{ // S240
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S241
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S242
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S243
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(15), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S244
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(16), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S245
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(17), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S246
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S247
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(20), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S248
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(21), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S249
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(22), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S250
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(417), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S251
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S252
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(218), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S253
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ), reduce: Star */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(99), /* |, reduce: Star */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S254
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(101), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(101), /* |, reduce: Underscore */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S255
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(103), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(103), /* |, reduce: Tilde */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S256
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(423), /* * */
			shift(424), /* _ */
			shift(425), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S257
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(34), /* ), reduce: ContinueAnd */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(34), /* &, reduce: ContinueAnd */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(34), /* space, reduce: ContinueAnd */

		},
	},
	actionRow{ // S258
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S259
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S260
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S261
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(15), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S262
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(16), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S263
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(17), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S264
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S265
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(20), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S266
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(21), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S267
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(22), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S268
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(430), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S269
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S270
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(218), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S271
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ), reduce: Star */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(99), /* &, reduce: Star */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S272
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(101), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(101), /* &, reduce: Underscore */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S273
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(103), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(103), /* &, reduce: Tilde */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S274
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(90), /* string_lit, reduce: Pipe */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(90), /* (, reduce: Pipe */
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
			reduce(90), /* !, reduce: Pipe */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			reduce(90), /* ., reduce: Pipe */
			reduce(90), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S275
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(107), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			reduce(107), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S276
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(436), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(437), /* ( */
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
			shift(53),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(438), /* . */
			shift(439), /* space */

		},
	},
	actionRow{ // S277
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(13), /* ), reduce: ContinueNameChoice */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(13), /* |, reduce: ContinueNameChoice */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(13), /* space, reduce: ContinueNameChoice */

		},
	},
	actionRow{ // S278
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			reduce(7), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S279
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			reduce(9), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S280
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S281
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(445), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(449), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S282
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* ), reduce: NameExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(11), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S283
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(223), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S284
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(73), /* string_lit, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(73), /* (, reduce: OpenParen */
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
			reduce(73), /* !, reduce: OpenParen */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			reduce(73), /* ., reduce: OpenParen */
			reduce(73), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S285
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(105), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(105), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S286
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			reduce(108), /* string_lit, reduce: Space */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(108), /* (, reduce: Space */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			reduce(108), /* !, reduce: Space */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			reduce(108), /* ., reduce: Space */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S287
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(160), /* * */
			shift(161), /* _ */
			shift(162), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S288
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(18), /* &, reduce: Pattern */
			reduce(18), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S289
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S290
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S291
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S292
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(456), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S293
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(459), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(224), /* space */

		},
	},
	actionRow{ // S294
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(83),  /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S295
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S296
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S297
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S298
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(465), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S299
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(19), /* &, reduce: Pattern */
			reduce(19), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S300
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(79), /* &, reduce: CloseCurly */
			reduce(79), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S301
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S302
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(329), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S303
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(359), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S304
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(74), /* id, reduce: OpenParen */
			reduce(74), /* string_lit, reduce: OpenParen */
			reduce(74), /* []bool, reduce: OpenParen */
			reduce(74), /* []int, reduce: OpenParen */
			reduce(74), /* []uint, reduce: OpenParen */
			reduce(74), /* []double, reduce: OpenParen */
			reduce(74), /* []string, reduce: OpenParen */
			reduce(74), /* [][]byte, reduce: OpenParen */
			reduce(74), /* int_lit, reduce: OpenParen */
			reduce(74), /* uint_lit, reduce: OpenParen */
			reduce(74), /* double_lit, reduce: OpenParen */
			reduce(74), /* bytes_lit, reduce: OpenParen */
			reduce(74), /* bool_var, reduce: OpenParen */
			reduce(74), /* int_var, reduce: OpenParen */
			reduce(74), /* uint_var, reduce: OpenParen */
			reduce(74), /* double_var, reduce: OpenParen */
			reduce(74), /* string_var, reduce: OpenParen */
			reduce(74), /* bytes_var, reduce: OpenParen */
			reduce(74), /* true, reduce: OpenParen */
			reduce(74), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(74), /* ), reduce: OpenParen */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(74), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S305
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(471), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(474), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(475), /* space */

		},
	},
	actionRow{ // S306
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S307
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(61), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(61), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S308
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(42), /* }, reduce: Function */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(42), /* space, reduce: Function */

		},
	},
	actionRow{ // S309
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(47), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(47), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(47), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S310
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(36), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(36), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(36), /* space, reduce: Expr */

		},
	},
	actionRow{ // S311
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(37), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(37), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(37), /* space, reduce: Expr */

		},
	},
	actionRow{ // S312
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(38), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(38), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S313
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(329), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S314
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S315
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(55), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(55), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(55), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S316
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(57), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(57), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S317
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(58), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(58), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S318
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(59), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(59), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S319
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(60), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(60), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S320
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(62), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(62), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S321
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(63), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(63), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S322
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(64), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(64), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S323
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(65), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(65), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S324
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(66), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(66), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S325
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(67), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(67), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S326
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(68), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(68), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S327
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(69), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(69), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(69), /* space, reduce: Bool */

		},
	},
	actionRow{ // S328
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(70), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(70), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(70), /* space, reduce: Bool */

		},
	},
	actionRow{ // S329
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* }, reduce: CloseParen */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S330
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(108), /* id, reduce: Space */
			reduce(108), /* string_lit, reduce: Space */
			reduce(108), /* []bool, reduce: Space */
			reduce(108), /* []int, reduce: Space */
			reduce(108), /* []uint, reduce: Space */
			reduce(108), /* []double, reduce: Space */
			reduce(108), /* []string, reduce: Space */
			reduce(108), /* [][]byte, reduce: Space */
			reduce(108), /* int_lit, reduce: Space */
			reduce(108), /* uint_lit, reduce: Space */
			reduce(108), /* double_lit, reduce: Space */
			reduce(108), /* bytes_lit, reduce: Space */
			reduce(108), /* bool_var, reduce: Space */
			reduce(108), /* int_var, reduce: Space */
			reduce(108), /* uint_var, reduce: Space */
			reduce(108), /* double_var, reduce: Space */
			reduce(108), /* string_var, reduce: Space */
			reduce(108), /* bytes_var, reduce: Space */
			reduce(108), /* true, reduce: Space */
			reduce(108), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S331
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(80), /* $, reduce: CloseCurly */
			reduce(80), /* id, reduce: CloseCurly */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S332
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(107), /* }, reduce: Space */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S333
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(78), /* id, reduce: OpenCurly */
			reduce(78), /* string_lit, reduce: OpenCurly */
			reduce(78), /* []bool, reduce: OpenCurly */
			reduce(78), /* []int, reduce: OpenCurly */
			reduce(78), /* []uint, reduce: OpenCurly */
			reduce(78), /* []double, reduce: OpenCurly */
			reduce(78), /* []string, reduce: OpenCurly */
			reduce(78), /* [][]byte, reduce: OpenCurly */
			reduce(78), /* int_lit, reduce: OpenCurly */
			reduce(78), /* uint_lit, reduce: OpenCurly */
			reduce(78), /* double_lit, reduce: OpenCurly */
			reduce(78), /* bytes_lit, reduce: OpenCurly */
			reduce(78), /* bool_var, reduce: OpenCurly */
			reduce(78), /* int_var, reduce: OpenCurly */
			reduce(78), /* uint_var, reduce: OpenCurly */
			reduce(78), /* double_var, reduce: OpenCurly */
			reduce(78), /* string_var, reduce: OpenCurly */
			reduce(78), /* bytes_var, reduce: OpenCurly */
			reduce(78), /* true, reduce: OpenCurly */
			reduce(78), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(78), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(78), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S334
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			reduce(107), /* {, reduce: Space */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S335
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(483), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(486), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(487), /* space */

		},
	},
	actionRow{ // S336
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S337
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(61), /* }, reduce: Terminal */
			reduce(61), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S338
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(47), /* }, reduce: Exprs */
			reduce(47), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(47), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S339
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(46), /* }, reduce: List */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(46), /* space, reduce: List */

		},
	},
	actionRow{ // S340
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(36), /* }, reduce: Expr */
			reduce(36), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(36), /* space, reduce: Expr */

		},
	},
	actionRow{ // S341
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(37), /* }, reduce: Expr */
			reduce(37), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(37), /* space, reduce: Expr */

		},
	},
	actionRow{ // S342
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(38), /* }, reduce: Expr */
			reduce(38), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S343
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(359), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S344
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S345
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(55), /* }, reduce: SpaceTerminal */
			reduce(55), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(55), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S346
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(57), /* }, reduce: Terminal */
			reduce(57), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S347
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(58), /* }, reduce: Terminal */
			reduce(58), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S348
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(59), /* }, reduce: Terminal */
			reduce(59), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S349
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(60), /* }, reduce: Terminal */
			reduce(60), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S350
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(62), /* }, reduce: Terminal */
			reduce(62), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S351
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(63), /* }, reduce: Terminal */
			reduce(63), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S352
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(64), /* }, reduce: Terminal */
			reduce(64), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S353
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(65), /* }, reduce: Terminal */
			reduce(65), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S354
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(66), /* }, reduce: Terminal */
			reduce(66), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S355
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(67), /* }, reduce: Terminal */
			reduce(67), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S356
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(68), /* }, reduce: Terminal */
			reduce(68), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S357
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(69), /* }, reduce: Bool */
			reduce(69), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(69), /* space, reduce: Bool */

		},
	},
	actionRow{ // S358
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(70), /* }, reduce: Bool */
			reduce(70), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(70), /* space, reduce: Bool */

		},
	},
	actionRow{ // S359
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* }, reduce: CloseCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S360
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(108), /* id, reduce: Space */
			reduce(108), /* string_lit, reduce: Space */
			reduce(108), /* []bool, reduce: Space */
			reduce(108), /* []int, reduce: Space */
			reduce(108), /* []uint, reduce: Space */
			reduce(108), /* []double, reduce: Space */
			reduce(108), /* []string, reduce: Space */
			reduce(108), /* [][]byte, reduce: Space */
			reduce(108), /* int_lit, reduce: Space */
			reduce(108), /* uint_lit, reduce: Space */
			reduce(108), /* double_lit, reduce: Space */
			reduce(108), /* bytes_lit, reduce: Space */
			reduce(108), /* bool_var, reduce: Space */
			reduce(108), /* int_var, reduce: Space */
			reduce(108), /* uint_var, reduce: Space */
			reduce(108), /* double_var, reduce: Space */
			reduce(108), /* string_var, reduce: Space */
			reduce(108), /* bytes_var, reduce: Space */
			reduce(108), /* true, reduce: Space */
			reduce(108), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(108), /* }, reduce: Space */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S361
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(82), /* string_lit, reduce: Comma */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(82), /* (, reduce: Comma */
			nil,        /* ) */
			reduce(82), /* {, reduce: Comma */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(82), /* #, reduce: Comma */
			nil,        /* & */
			nil,        /* | */
			reduce(82), /* [, reduce: Comma */
			nil,        /* ] */
			nil,        /* : */
			reduce(82), /* !, reduce: Comma */
			reduce(82), /* *, reduce: Comma */
			reduce(82), /* _, reduce: Comma */
			reduce(82), /* ~, reduce: Comma */
			reduce(82), /* ., reduce: Comma */
			reduce(82), /* space, reduce: Comma */

		},
	},
	actionRow{ // S362
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(107), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S363
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(494), /* * */
			shift(495), /* _ */
			shift(496), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S364
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(28), /* ,, reduce: ContinueConcat */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(28), /* ], reduce: ContinueConcat */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(28), /* space, reduce: ContinueConcat */

		},
	},
	actionRow{ // S365
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S366
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S367
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S368
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(15), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(15), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(15), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S369
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(16), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(16), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(16), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S370
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(17), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(17), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S371
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(89),  /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(103), /* int_lit */
			shift(104), /* uint_lit */
			shift(105), /* double_lit */
			shift(106), /* bytes_lit */
			shift(107), /* bool_var */
			shift(108), /* int_var */
			shift(109), /* uint_var */
			shift(110), /* double_var */
			shift(111), /* string_var */
			shift(112), /* bytes_var */
			shift(113), /* true */
			shift(114), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S372
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(20), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(20), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S373
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(21), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(21), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S374
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(22), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(22), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S375
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(501), /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			nil,        /* space */

		},
	},
	actionRow{ // S376
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(131), /* * */
			shift(132), /* _ */
			shift(133), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S377
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(506), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S378
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(99), /* ,, reduce: Star */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(99), /* ], reduce: Star */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(99), /* space, reduce: Star */

		},
	},
	actionRow{ // S379
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(101), /* ,, reduce: Underscore */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(101), /* ], reduce: Underscore */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(101), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S380
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(103), /* ,, reduce: Tilde */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(103), /* ], reduce: Tilde */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(103), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S381
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S382
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(510), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S383
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(131), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S384
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S385
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S386
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(515), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S387
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S388
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S389
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S390
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(76), /* $, reduce: CloseParen */
			reduce(76), /* id, reduce: CloseParen */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S391
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S392
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S393
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			reduce(107), /* :, reduce: Space */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S394
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(48),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(49),  /* ( */
			nil,        /* ) */
			shift(50),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(51),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(52),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(53),  /* ! */
			shift(212), /* * */
			shift(213), /* _ */
			shift(214), /* ~ */
			shift(57),  /* . */
			shift(58),  /* space */

		},
	},
	actionRow{ // S395
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S396
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(61), /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(62), /* space */

		},
	},
	actionRow{ // S397
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			shift(65), /* space */

		},
	},
	actionRow{ // S398
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(70), /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			shift(37), /* ( */
			nil,       /* ) */
			shift(38), /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(39), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(40), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(41), /* ! */
			shift(83), /* * */
			shift(84), /* _ */
			shift(85), /* ~ */
			shift(86), /* . */
			shift(46), /* space */

		},
	},
	actionRow{ // S399
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(520), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S400
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(523), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(61),  /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(224), /* space */

		},
	},
	actionRow{ // S401
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(155), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S402
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S403
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S404
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S405
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(529), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S406
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S407
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ), reduce: CloseCurly */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S408
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S409
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			reduce(107), /* *, reduce: Space */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S410
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(100), /* ), reduce: Star */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(100), /* |, reduce: Star */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S411
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(102), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(102), /* |, reduce: Underscore */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S412
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(104), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(104), /* |, reduce: Tilde */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S413
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S414
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S415
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S416
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(538), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S417
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(25), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S418
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S419
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(390), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S420
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(30), /* $, reduce: StartOr */
			reduce(30), /* id, reduce: StartOr */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S421
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S422
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(108), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S423
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(100), /* ), reduce: Star */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(100), /* &, reduce: Star */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S424
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(102), /* ), reduce: Underscore */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(102), /* &, reduce: Underscore */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S425
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(104), /* ), reduce: Tilde */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(104), /* &, reduce: Tilde */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S426
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S427
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S428
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S429
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(549), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S430
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(25), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S431
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S432
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(390), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S433
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(33), /* $, reduce: StartAnd */
			reduce(33), /* id, reduce: StartAnd */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S434
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S435
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(108), /* &, reduce: Space */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S436
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* ), reduce: NameExpr */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			reduce(8), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S437
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(74), /* string_lit, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(74), /* (, reduce: OpenParen */
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
			reduce(74), /* !, reduce: OpenParen */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			reduce(74), /* ., reduce: OpenParen */
			reduce(74), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S438
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(106), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(106), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S439
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			reduce(107), /* string_lit, reduce: Space */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			reduce(107), /* (, reduce: Space */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			reduce(107), /* !, reduce: Space */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			reduce(107), /* ., reduce: Space */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S440
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(437), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(137), /* space */

		},
	},
	actionRow{ // S441
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(556), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(560), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S442
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(561), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(437), /* ( */
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
			shift(53),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(562), /* . */
			shift(439), /* space */

		},
	},
	actionRow{ // S443
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(565), /* space */

		},
	},
	actionRow{ // S444
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S445
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S446
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S447
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(445), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(449), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S448
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S449
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(105), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S450
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(392), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S451
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* :, reduce: StartNameChoice */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
	actionRow{ // S452
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S453
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S454
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(569), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S455
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(26), /* &, reduce: Pattern */
			reduce(26), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S456
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(75), /* &, reduce: CloseParen */
			reduce(75), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S457
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(570), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(134), /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(393), /* space */

		},
	},
	actionRow{ // S458
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			reduce(10), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S459
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			reduce(75), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S460
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(160), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S461
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(23), /* &, reduce: Pattern */
			reduce(23), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S462
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(456), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S463
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(456), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S464
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(459), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S465
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(80), /* &, reduce: CloseCurly */
			reduce(80), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S466
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(579), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S467
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(40), /* }, reduce: Function */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(40), /* space, reduce: Function */

		},
	},
	actionRow{ // S468
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(329), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S469
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(45), /* }, reduce: List */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(45), /* space, reduce: List */

		},
	},
	actionRow{ // S470
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(359), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S471
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S472
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S473
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(56), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(56), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(56), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S474
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* }, reduce: CloseParen */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S475
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(107), /* id, reduce: Space */
			reduce(107), /* string_lit, reduce: Space */
			reduce(107), /* []bool, reduce: Space */
			reduce(107), /* []int, reduce: Space */
			reduce(107), /* []uint, reduce: Space */
			reduce(107), /* []double, reduce: Space */
			reduce(107), /* []string, reduce: Space */
			reduce(107), /* [][]byte, reduce: Space */
			reduce(107), /* int_lit, reduce: Space */
			reduce(107), /* uint_lit, reduce: Space */
			reduce(107), /* double_lit, reduce: Space */
			reduce(107), /* bytes_lit, reduce: Space */
			reduce(107), /* bool_var, reduce: Space */
			reduce(107), /* int_var, reduce: Space */
			reduce(107), /* uint_var, reduce: Space */
			reduce(107), /* double_var, reduce: Space */
			reduce(107), /* string_var, reduce: Space */
			reduce(107), /* bytes_var, reduce: Space */
			reduce(107), /* true, reduce: Space */
			reduce(107), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S476
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(587), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S477
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(474), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(589), /* space */

		},
	},
	actionRow{ // S478
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(41), /* }, reduce: Function */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(41), /* space, reduce: Function */

		},
	},
	actionRow{ // S479
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S480
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(81), /* id, reduce: Comma */
			reduce(81), /* string_lit, reduce: Comma */
			reduce(81), /* []bool, reduce: Comma */
			reduce(81), /* []int, reduce: Comma */
			reduce(81), /* []uint, reduce: Comma */
			reduce(81), /* []double, reduce: Comma */
			reduce(81), /* []string, reduce: Comma */
			reduce(81), /* [][]byte, reduce: Comma */
			reduce(81), /* int_lit, reduce: Comma */
			reduce(81), /* uint_lit, reduce: Comma */
			reduce(81), /* double_lit, reduce: Comma */
			reduce(81), /* bytes_lit, reduce: Comma */
			reduce(81), /* bool_var, reduce: Comma */
			reduce(81), /* int_var, reduce: Comma */
			reduce(81), /* uint_var, reduce: Comma */
			reduce(81), /* double_var, reduce: Comma */
			reduce(81), /* string_var, reduce: Comma */
			reduce(81), /* bytes_var, reduce: Comma */
			reduce(81), /* true, reduce: Comma */
			reduce(81), /* false, reduce: Comma */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(81), /* space, reduce: Comma */

		},
	},
	actionRow{ // S481
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(108), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			reduce(108), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S482
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(595), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S483
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(189), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S484
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(196), /* { */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(197), /* space */

		},
	},
	actionRow{ // S485
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(56), /* }, reduce: SpaceTerminal */
			reduce(56), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(56), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S486
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* }, reduce: CloseCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S487
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			reduce(107), /* id, reduce: Space */
			reduce(107), /* string_lit, reduce: Space */
			reduce(107), /* []bool, reduce: Space */
			reduce(107), /* []int, reduce: Space */
			reduce(107), /* []uint, reduce: Space */
			reduce(107), /* []double, reduce: Space */
			reduce(107), /* []string, reduce: Space */
			reduce(107), /* [][]byte, reduce: Space */
			reduce(107), /* int_lit, reduce: Space */
			reduce(107), /* uint_lit, reduce: Space */
			reduce(107), /* double_lit, reduce: Space */
			reduce(107), /* bytes_lit, reduce: Space */
			reduce(107), /* bool_var, reduce: Space */
			reduce(107), /* int_var, reduce: Space */
			reduce(107), /* uint_var, reduce: Space */
			reduce(107), /* double_var, reduce: Space */
			reduce(107), /* string_var, reduce: Space */
			reduce(107), /* bytes_var, reduce: Space */
			reduce(107), /* true, reduce: Space */
			reduce(107), /* false, reduce: Space */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(107), /* }, reduce: Space */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S488
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(601), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S489
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(486), /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(602), /* space */

		},
	},
	actionRow{ // S490
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(44), /* }, reduce: List */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(44), /* space, reduce: List */

		},
	},
	actionRow{ // S491
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(115), /* space */

		},
	},
	actionRow{ // S492
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(108), /* }, reduce: Space */
			reduce(108), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S493
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(608), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S494
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(100), /* ,, reduce: Star */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(100), /* ], reduce: Star */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(100), /* space, reduce: Star */

		},
	},
	actionRow{ // S495
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(102), /* ,, reduce: Underscore */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(102), /* ], reduce: Underscore */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(102), /* space, reduce: Underscore */

		},
	},
	actionRow{ // S496
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(104), /* ,, reduce: Tilde */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(104), /* ], reduce: Tilde */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(104), /* space, reduce: Tilde */

		},
	},
	actionRow{ // S497
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S498
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S499
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(168), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(171), /* space */

		},
	},
	actionRow{ // S500
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(616), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(193), /* space */

		},
	},
	actionRow{ // S501
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(25), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(25), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S502
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(204), /* space */

		},
	},
	actionRow{ // S503
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(618), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S504
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S505
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(27), /* $, reduce: StartConcat */
			reduce(27), /* id, reduce: StartConcat */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S506
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(93), /* $, reduce: CloseBracket */
			reduce(93), /* id, reduce: CloseBracket */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S507
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(108), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(108), /* ], reduce: Space */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S508
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(621), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S509
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S510
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S511
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(198), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S512
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S513
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(510), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S514
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(510), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S515
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S516
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(628), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S517
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(142), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(155), /* * */
			shift(156), /* _ */
			shift(157), /* ~ */
			shift(158), /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S518
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(629), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S519
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S520
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S521
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(630), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			shift(134), /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(393), /* space */

		},
	},
	actionRow{ // S522
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* ), reduce: NameExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			reduce(10), /* :, reduce: NameExpr */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S523
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			reduce(75), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S524
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(212), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S525
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ), reduce: Pattern */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S526
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(520), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S527
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(520), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S528
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(523), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S529
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ), reduce: CloseCurly */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S530
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(639), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S531
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(18), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S532
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(642), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S533
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(253), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S534
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S535
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S536
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(647), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S537
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(19), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S538
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(79), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S539
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S540
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(107), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S541
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(32), /* ), reduce: ContinueOr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(32), /* |, reduce: ContinueOr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(32), /* space, reduce: ContinueOr */

		},
	},
	actionRow{ // S542
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(18), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S543
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(651), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S544
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(271), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S545
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S546
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S547
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(656), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S548
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(19), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S549
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(79), /* &, reduce: CloseCurly */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S550
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S551
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			reduce(107), /* &, reduce: Space */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S552
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(35), /* ), reduce: ContinueAnd */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(35), /* &, reduce: ContinueAnd */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(35), /* space, reduce: ContinueAnd */

		},
	},
	actionRow{ // S553
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(658), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(437), /* ( */
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
			shift(53),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(659), /* . */
			shift(439), /* space */

		},
	},
	actionRow{ // S554
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(642), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S555
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(7), /* ), reduce: NameExpr */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S556
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(9), /* ), reduce: NameExpr */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S557
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(65),  /* space */

		},
	},
	actionRow{ // S558
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(445), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(449), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S559
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(11), /* ), reduce: NameExpr */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S560
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(105), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(105), /* space, reduce: Dot */

		},
	},
	actionRow{ // S561
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* |, reduce: NameExpr */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S562
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(106), /* |, reduce: Dot */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S563
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(663), /* space */

		},
	},
	actionRow{ // S564
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S565
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(108), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(108), /* space, reduce: Space */

		},
	},
	actionRow{ // S566
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(556), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(560), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S567
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(565), /* space */

		},
	},
	actionRow{ // S568
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(14), /* ), reduce: ContinueNameChoice */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(14), /* |, reduce: ContinueNameChoice */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(14), /* space, reduce: ContinueNameChoice */

		},
	},
	actionRow{ // S569
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(76), /* &, reduce: CloseParen */
			reduce(76), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S570
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			reduce(76), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S571
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(569), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S572
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ), reduce: StartOr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(30), /* &, reduce: StartOr */
			reduce(30), /* |, reduce: StartOr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S573
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(569), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S574
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ), reduce: StartAnd */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(33), /* &, reduce: StartAnd */
			reduce(33), /* |, reduce: StartAnd */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S575
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(570), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S576
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* |, reduce: StartNameChoice */
			nil,        /* [ */
			nil,        /* ] */
			reduce(12), /* :, reduce: StartNameChoice */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
	actionRow{ // S577
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(667), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S578
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ), reduce: StartConcat */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(27), /* &, reduce: StartConcat */
			reduce(27), /* |, reduce: StartConcat */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S579
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(93), /* &, reduce: CloseBracket */
			reduce(93), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S580
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(39), /* }, reduce: Function */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(39), /* space, reduce: Function */

		},
	},
	actionRow{ // S581
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(43), /* }, reduce: List */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(43), /* space, reduce: List */

		},
	},
	actionRow{ // S582
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(587), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S583
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(595), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S584
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(471), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(672), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(475), /* space */

		},
	},
	actionRow{ // S585
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(42), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(42), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(42), /* space, reduce: Function */

		},
	},
	actionRow{ // S586
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(587), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S587
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(75), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S588
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(82), /* id, reduce: Comma */
			reduce(82), /* string_lit, reduce: Comma */
			reduce(82), /* []bool, reduce: Comma */
			reduce(82), /* []int, reduce: Comma */
			reduce(82), /* []uint, reduce: Comma */
			reduce(82), /* []double, reduce: Comma */
			reduce(82), /* []string, reduce: Comma */
			reduce(82), /* [][]byte, reduce: Comma */
			reduce(82), /* int_lit, reduce: Comma */
			reduce(82), /* uint_lit, reduce: Comma */
			reduce(82), /* double_lit, reduce: Comma */
			reduce(82), /* bytes_lit, reduce: Comma */
			reduce(82), /* bool_var, reduce: Comma */
			reduce(82), /* int_var, reduce: Comma */
			reduce(82), /* uint_var, reduce: Comma */
			reduce(82), /* double_var, reduce: Comma */
			reduce(82), /* string_var, reduce: Comma */
			reduce(82), /* bytes_var, reduce: Comma */
			reduce(82), /* true, reduce: Comma */
			reduce(82), /* false, reduce: Comma */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(82), /* space, reduce: Comma */

		},
	},
	actionRow{ // S589
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(107), /* ), reduce: Space */
			nil,         /* { */
			nil,         /* } */
			reduce(107), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S590
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(471), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(186), /* space */

		},
	},
	actionRow{ // S591
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(48), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(48), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(48), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S592
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(483), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(675), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(487), /* space */

		},
	},
	actionRow{ // S593
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(46), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(46), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(46), /* space, reduce: List */

		},
	},
	actionRow{ // S594
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(595), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S595
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(79), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S596
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(306), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(601), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(330), /* space */

		},
	},
	actionRow{ // S597
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(336), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(608), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(360), /* space */

		},
	},
	actionRow{ // S598
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(471), /* id */
			shift(307), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(317), /* int_lit */
			shift(318), /* uint_lit */
			shift(319), /* double_lit */
			shift(320), /* bytes_lit */
			shift(321), /* bool_var */
			shift(322), /* int_var */
			shift(323), /* uint_var */
			shift(324), /* double_var */
			shift(325), /* string_var */
			shift(326), /* bytes_var */
			shift(327), /* true */
			shift(328), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(682), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(475), /* space */

		},
	},
	actionRow{ // S599
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(42), /* }, reduce: Function */
			reduce(42), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(42), /* space, reduce: Function */

		},
	},
	actionRow{ // S600
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(601), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S601
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* }, reduce: CloseParen */
			reduce(75), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S602
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			reduce(107), /* }, reduce: Space */
			reduce(107), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S603
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(483), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(186), /* space */

		},
	},
	actionRow{ // S604
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(48), /* }, reduce: Exprs */
			reduce(48), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(48), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S605
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(483), /* id */
			shift(337), /* string_lit */
			shift(95),  /* []bool */
			shift(96),  /* []int */
			shift(97),  /* []uint */
			shift(98),  /* []double */
			shift(99),  /* []string */
			shift(100), /* [][]byte */
			shift(347), /* int_lit */
			shift(348), /* uint_lit */
			shift(349), /* double_lit */
			shift(350), /* bytes_lit */
			shift(351), /* bool_var */
			shift(352), /* int_var */
			shift(353), /* uint_var */
			shift(354), /* double_var */
			shift(355), /* string_var */
			shift(356), /* bytes_var */
			shift(357), /* true */
			shift(358), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(685), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(487), /* space */

		},
	},
	actionRow{ // S606
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(46), /* }, reduce: List */
			reduce(46), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(46), /* space, reduce: List */

		},
	},
	actionRow{ // S607
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(608), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S608
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* }, reduce: CloseCurly */
			reduce(79), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S609
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(18), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(18), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S610
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(690), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S611
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(378), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(237), /* space */

		},
	},
	actionRow{ // S612
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(253), /* * */
			shift(254), /* _ */
			shift(255), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S613
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(271), /* * */
			shift(272), /* _ */
			shift(273), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S614
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(695), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(332), /* space */

		},
	},
	actionRow{ // S615
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(19), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(19), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S616
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(79), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(79), /* ], reduce: CloseCurly */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(79), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S617
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(24),  /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(37),  /* ( */
			nil,        /* ) */
			shift(38),  /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(39),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(40),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(41),  /* ! */
			shift(378), /* * */
			shift(379), /* _ */
			shift(380), /* ~ */
			shift(45),  /* . */
			shift(46),  /* space */

		},
	},
	actionRow{ // S618
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(94), /* $, reduce: CloseBracket */
			reduce(94), /* id, reduce: CloseBracket */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S619
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			reduce(107), /* ,, reduce: Space */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			reduce(107), /* ], reduce: Space */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S620
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(29), /* ,, reduce: ContinueConcat */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(29), /* ], reduce: ContinueConcat */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(29), /* space, reduce: ContinueConcat */

		},
	},
	actionRow{ // S621
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S622
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(621), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S623
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ,, reduce: StartOr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S624
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(621), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S625
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ,, reduce: StartAnd */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S626
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(697), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S627
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ,, reduce: StartConcat */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S628
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S629
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S630
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			reduce(76), /* :, reduce: CloseParen */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S631
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(629), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S632
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ), reduce: StartOr */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S633
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(629), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S634
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ), reduce: StartAnd */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S635
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(630), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S636
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* ), reduce: StartNameChoice */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			reduce(12), /* :, reduce: StartNameChoice */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
	actionRow{ // S637
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(698), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S638
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ), reduce: StartConcat */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S639
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ), reduce: CloseBracket */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S640
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(699), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S641
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(26), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S642
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(75), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S643
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(410), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S644
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(23), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S645
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(642), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S646
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(642), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S647
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(80), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S648
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(706), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S649
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(707), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S650
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(26), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S651
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(75), /* &, reduce: CloseParen */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S652
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(423), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S653
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(23), /* &, reduce: Pattern */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S654
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(651), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S655
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(651), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S656
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(80), /* &, reduce: CloseCurly */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S657
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(714), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S658
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* string_lit */
			nil,       /* []bool */
			nil,       /* []int */
			nil,       /* []uint */
			nil,       /* []double */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int_lit */
			nil,       /* uint_lit */
			nil,       /* double_lit */
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
			reduce(8), /* ), reduce: NameExpr */
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
			nil,       /* * */
			nil,       /* _ */
			nil,       /* ~ */
			nil,       /* . */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S659
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			reduce(106), /* ), reduce: Dot */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			nil,         /* | */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(106), /* space, reduce: Dot */

		},
	},
	actionRow{ // S660
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* ), reduce: NameExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(10), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S661
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(556), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(560), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S662
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(565), /* space */

		},
	},
	actionRow{ // S663
		canRecover: false,
		actions: [numSymbols]action{
			nil,         /* INVALID */
			nil,         /* $ */
			nil,         /* id */
			nil,         /* string_lit */
			nil,         /* []bool */
			nil,         /* []int */
			nil,         /* []uint */
			nil,         /* []double */
			nil,         /* []string */
			nil,         /* [][]byte */
			nil,         /* int_lit */
			nil,         /* uint_lit */
			nil,         /* double_lit */
			nil,         /* bytes_lit */
			nil,         /* bool_var */
			nil,         /* int_var */
			nil,         /* uint_var */
			nil,         /* double_var */
			nil,         /* string_var */
			nil,         /* bytes_var */
			nil,         /* true */
			nil,         /* false */
			nil,         /* = */
			nil,         /* ( */
			nil,         /* ) */
			nil,         /* { */
			nil,         /* } */
			nil,         /* , */
			nil,         /* ; */
			nil,         /* # */
			nil,         /* & */
			reduce(107), /* |, reduce: Space */
			nil,         /* [ */
			nil,         /* ] */
			nil,         /* : */
			nil,         /* ! */
			nil,         /* * */
			nil,         /* _ */
			nil,         /* ~ */
			nil,         /* . */
			reduce(107), /* space, reduce: Space */

		},
	},
	actionRow{ // S664
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(642), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S665
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(721), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S666
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S667
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(94), /* &, reduce: CloseBracket */
			reduce(94), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S668
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(40), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(40), /* space, reduce: Function */

		},
	},
	actionRow{ // S669
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(587), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S670
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(45), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(45), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(45), /* space, reduce: List */

		},
	},
	actionRow{ // S671
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(595), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S672
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(76), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S673
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(672), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(589), /* space */

		},
	},
	actionRow{ // S674
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(41), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(41), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(41), /* space, reduce: Function */

		},
	},
	actionRow{ // S675
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(80), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S676
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(675), /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(602), /* space */

		},
	},
	actionRow{ // S677
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(44), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(44), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(44), /* space, reduce: List */

		},
	},
	actionRow{ // S678
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(40), /* }, reduce: Function */
			reduce(40), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(40), /* space, reduce: Function */

		},
	},
	actionRow{ // S679
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(601), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(481), /* space */

		},
	},
	actionRow{ // S680
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(45), /* }, reduce: List */
			reduce(45), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(45), /* space, reduce: List */

		},
	},
	actionRow{ // S681
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(608), /* } */
			shift(480), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(492), /* space */

		},
	},
	actionRow{ // S682
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* }, reduce: CloseParen */
			reduce(76), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S683
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(682), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(589), /* space */

		},
	},
	actionRow{ // S684
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(41), /* }, reduce: Function */
			reduce(41), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(41), /* space, reduce: Function */

		},
	},
	actionRow{ // S685
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* }, reduce: CloseCurly */
			reduce(80), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S686
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(685), /* } */
			shift(588), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(602), /* space */

		},
	},
	actionRow{ // S687
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(44), /* }, reduce: List */
			reduce(44), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(44), /* space, reduce: List */

		},
	},
	actionRow{ // S688
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(727), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S689
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(26), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(26), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S690
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(75), /* ], reduce: CloseParen */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S691
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(494), /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(409), /* space */

		},
	},
	actionRow{ // S692
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(23), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(23), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S693
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(690), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(170), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S694
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(690), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(169), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(435), /* space */

		},
	},
	actionRow{ // S695
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(80), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(80), /* ], reduce: CloseCurly */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(80), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S696
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(203), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(734), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(507), /* space */

		},
	},
	actionRow{ // S697
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S698
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ), reduce: CloseBracket */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S699
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(76), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S700
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(699), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S701
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ), reduce: StartOr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(30), /* |, reduce: StartOr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S702
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(699), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S703
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ), reduce: StartAnd */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(33), /* |, reduce: StartAnd */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S704
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(735), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S705
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ), reduce: StartConcat */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(27), /* |, reduce: StartConcat */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S706
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(93), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S707
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(76), /* &, reduce: CloseParen */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S708
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(707), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S709
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ), reduce: StartOr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(30), /* &, reduce: StartOr */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S710
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(707), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S711
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ), reduce: StartAnd */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(33), /* &, reduce: StartAnd */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S712
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(736), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S713
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ), reduce: StartConcat */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(27), /* &, reduce: StartConcat */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S714
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(93), /* &, reduce: CloseBracket */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S715
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(520), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(219), /* space */

		},
	},
	actionRow{ // S716
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(279), /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(284), /* ( */
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
			shift(41),  /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			shift(285), /* . */
			shift(286), /* space */

		},
	},
	actionRow{ // S717
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(699), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S718
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* ), reduce: StartNameChoice */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(12), /* |, reduce: StartNameChoice */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
	actionRow{ // S719
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(739), /* ) */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(391), /* space */

		},
	},
	actionRow{ // S720
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* |, reduce: NameExpr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S721
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(75), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(75), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S722
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(721), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S723
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(39), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(39), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(39), /* space, reduce: Function */

		},
	},
	actionRow{ // S724
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(43), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(43), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(43), /* space, reduce: List */

		},
	},
	actionRow{ // S725
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(39), /* }, reduce: Function */
			reduce(39), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(39), /* space, reduce: Function */

		},
	},
	actionRow{ // S726
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(43), /* }, reduce: List */
			reduce(43), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(43), /* space, reduce: List */

		},
	},
	actionRow{ // S727
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(76), /* ], reduce: CloseParen */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S728
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(727), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(233), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S729
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(30), /* ,, reduce: StartOr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(30), /* ], reduce: StartOr */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(30), /* space, reduce: StartOr */

		},
	},
	actionRow{ // S730
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(727), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			shift(232), /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(551), /* space */

		},
	},
	actionRow{ // S731
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(33), /* ,, reduce: StartAnd */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(33), /* ], reduce: StartAnd */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(33), /* space, reduce: StartAnd */

		},
	},
	actionRow{ // S732
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(361), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			shift(742), /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(619), /* space */

		},
	},
	actionRow{ // S733
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(27), /* ,, reduce: StartConcat */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(27), /* ], reduce: StartConcat */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(27), /* space, reduce: StartConcat */

		},
	},
	actionRow{ // S734
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(93), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(93), /* ], reduce: CloseBracket */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(93), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S735
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			reduce(94), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S736
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ), reduce: CloseBracket */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			reduce(94), /* &, reduce: CloseBracket */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S737
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(10), /* ), reduce: NameExpr */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S738
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(520), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(175), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(422), /* space */

		},
	},
	actionRow{ // S739
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(76), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(76), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S740
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(739), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S741
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* |, reduce: StartNameChoice */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
	actionRow{ // S742
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(94), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			reduce(94), /* ], reduce: CloseBracket */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(94), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S743
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			shift(629), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			shift(274), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			shift(540), /* space */

		},
	},
	actionRow{ // S744
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* string_lit */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
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
			reduce(12), /* ), reduce: StartNameChoice */
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
			nil,        /* * */
			nil,        /* _ */
			nil,        /* ~ */
			nil,        /* . */
			reduce(12), /* space, reduce: StartNameChoice */

		},
	},
}

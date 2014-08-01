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
			shift(10), /* root */
			shift(11), /* id */
			nil,       /* . */
			shift(12), /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(28), /* int64_lit */
			shift(29), /* int32_lit */
			shift(30), /* uint64_lit */
			shift(31), /* uint32_lit */
			shift(32), /* double_lit */
			shift(33), /* float_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* true */
			shift(37), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(38), /* space */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* root */
			nil,          /* id */
			nil,          /* . */
			nil,          /* if */
			nil,          /* { */
			nil,          /* []bool */
			nil,          /* []int64 */
			nil,          /* []int32 */
			nil,          /* []uint64 */
			nil,          /* []uint32 */
			nil,          /* []double */
			nil,          /* []float */
			nil,          /* []string */
			nil,          /* [][]byte */
			nil,          /* int64_lit */
			nil,          /* int32_lit */
			nil,          /* uint64_lit */
			nil,          /* uint32_lit */
			nil,          /* double_lit */
			nil,          /* float_lit */
			nil,          /* string_lit */
			nil,          /* bytes_lit */
			nil,          /* true */
			nil,          /* false */
			nil,          /* = */
			nil,          /* then */
			nil,          /* else */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* } */
			nil,          /* , */
			nil,          /* space */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: AllRules */
			shift(10), /* root */
			shift(41), /* id */
			nil,       /* . */
			shift(12), /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(42), /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(43), /* root */
			shift(44), /* id */
			nil,       /* . */
			shift(45), /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(28), /* int64_lit */
			shift(29), /* int32_lit */
			shift(30), /* uint64_lit */
			shift(31), /* uint32_lit */
			shift(32), /* double_lit */
			shift(33), /* float_lit */
			shift(34), /* string_lit */
			shift(35), /* bytes_lit */
			shift(36), /* true */
			shift(37), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(48), /* space */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Rules */
			reduce(4), /* root, reduce: Rules */
			reduce(4), /* id, reduce: Rules */
			nil,       /* . */
			reduce(4), /* if, reduce: Rules */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(4), /* space, reduce: Rules */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Rules */
			reduce(5), /* root, reduce: Rules */
			reduce(5), /* id, reduce: Rules */
			nil,       /* . */
			reduce(5), /* if, reduce: Rules */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(5), /* space, reduce: Rules */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Rule */
			reduce(6), /* root, reduce: Rule */
			reduce(6), /* id, reduce: Rule */
			nil,       /* . */
			reduce(6), /* if, reduce: Rule */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(6), /* space, reduce: Rule */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: Rule */
			reduce(7), /* root, reduce: Rule */
			reduce(7), /* id, reduce: Rule */
			nil,       /* . */
			reduce(7), /* if, reduce: Rule */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(7), /* space, reduce: Rule */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: Rule */
			reduce(8), /* root, reduce: Rule */
			reduce(8), /* id, reduce: Rule */
			nil,       /* . */
			reduce(8), /* if, reduce: Rule */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(8), /* space, reduce: Rule */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: Rule */
			reduce(9), /* root, reduce: Rule */
			reduce(9), /* id, reduce: Rule */
			nil,       /* . */
			reduce(9), /* if, reduce: Rule */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(9), /* space, reduce: Rule */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			shift(54), /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			shift(56), /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(57), /* space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(60), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(67), /* int64_lit */
			shift(68), /* int32_lit */
			shift(69), /* uint64_lit */
			shift(70), /* uint32_lit */
			shift(71), /* double_lit */
			shift(72), /* float_lit */
			shift(73), /* string_lit */
			shift(74), /* bytes_lit */
			shift(75), /* true */
			shift(76), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(77), /* space */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(39), /* $, reduce: Expr */
			reduce(39), /* root, reduce: Expr */
			reduce(39), /* id, reduce: Expr */
			nil,        /* . */
			reduce(39), /* if, reduce: Expr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(39), /* space, reduce: Expr */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(40), /* $, reduce: Expr */
			reduce(40), /* root, reduce: Expr */
			reduce(40), /* id, reduce: Expr */
			nil,        /* . */
			reduce(40), /* if, reduce: Expr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(40), /* space, reduce: Expr */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(38), /* $, reduce: Expr */
			reduce(38), /* root, reduce: Expr */
			reduce(38), /* id, reduce: Expr */
			nil,        /* . */
			reduce(38), /* if, reduce: Expr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(41), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(41), /* space, reduce: ListType */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(42), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(42), /* space, reduce: ListType */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(43), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(43), /* space, reduce: ListType */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(44), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(44), /* space, reduce: ListType */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(45), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(45), /* space, reduce: ListType */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(46), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(46), /* space, reduce: ListType */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(47), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(47), /* space, reduce: ListType */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(48), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: ListType */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(49), /* {, reduce: ListType */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(49), /* space, reduce: ListType */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(50), /* $, reduce: SpaceTerminal */
			reduce(50), /* root, reduce: SpaceTerminal */
			reduce(50), /* id, reduce: SpaceTerminal */
			nil,        /* . */
			reduce(50), /* if, reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(52), /* $, reduce: Terminal */
			reduce(52), /* root, reduce: Terminal */
			reduce(52), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(52), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(53), /* $, reduce: Terminal */
			reduce(53), /* root, reduce: Terminal */
			reduce(53), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(53), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(54), /* $, reduce: Terminal */
			reduce(54), /* root, reduce: Terminal */
			reduce(54), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(54), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(55), /* $, reduce: Terminal */
			reduce(55), /* root, reduce: Terminal */
			reduce(55), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(55), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(56), /* $, reduce: Terminal */
			reduce(56), /* root, reduce: Terminal */
			reduce(56), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(56), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(57), /* $, reduce: Terminal */
			reduce(57), /* root, reduce: Terminal */
			reduce(57), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(57), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(58), /* $, reduce: Terminal */
			reduce(58), /* root, reduce: Terminal */
			reduce(58), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(58), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(59), /* $, reduce: Terminal */
			reduce(59), /* root, reduce: Terminal */
			reduce(59), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(59), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(60), /* $, reduce: Terminal */
			reduce(60), /* root, reduce: Terminal */
			reduce(60), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(60), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(62), /* $, reduce: Bool */
			reduce(62), /* root, reduce: Bool */
			reduce(62), /* id, reduce: Bool */
			nil,        /* . */
			reduce(62), /* if, reduce: Bool */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(62), /* space, reduce: Bool */

		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(63), /* $, reduce: Bool */
			reduce(63), /* root, reduce: Bool */
			reduce(63), /* id, reduce: Bool */
			nil,        /* . */
			reduce(63), /* if, reduce: Bool */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(81), /* root, reduce: Space */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			reduce(81), /* if, reduce: Space */
			nil,        /* { */
			reduce(81), /* []bool, reduce: Space */
			reduce(81), /* []int64, reduce: Space */
			reduce(81), /* []int32, reduce: Space */
			reduce(81), /* []uint64, reduce: Space */
			reduce(81), /* []uint32, reduce: Space */
			reduce(81), /* []double, reduce: Space */
			reduce(81), /* []float, reduce: Space */
			reduce(81), /* []string, reduce: Space */
			reduce(81), /* [][]byte, reduce: Space */
			reduce(81), /* int64_lit, reduce: Space */
			reduce(81), /* int32_lit, reduce: Space */
			reduce(81), /* uint64_lit, reduce: Space */
			reduce(81), /* uint32_lit, reduce: Space */
			reduce(81), /* double_lit, reduce: Space */
			reduce(81), /* float_lit, reduce: Space */
			reduce(81), /* string_lit, reduce: Space */
			reduce(81), /* bytes_lit, reduce: Space */
			reduce(81), /* true, reduce: Space */
			reduce(81), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: AllRules */
			shift(43), /* root */
			shift(82), /* id */
			nil,       /* . */
			shift(45), /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(83), /* space */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Rules */
			reduce(3), /* root, reduce: Rules */
			reduce(3), /* id, reduce: Rules */
			nil,       /* . */
			reduce(3), /* if, reduce: Rules */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			reduce(3), /* space, reduce: Rules */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			shift(85), /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(86), /* space */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(81), /* $, reduce: Space */
			reduce(81), /* root, reduce: Space */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			reduce(81), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			shift(89), /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			shift(56), /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(57), /* space */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(60), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(67), /* int64_lit */
			shift(68), /* int32_lit */
			shift(69), /* uint64_lit */
			shift(70), /* uint32_lit */
			shift(71), /* double_lit */
			shift(72), /* float_lit */
			shift(73), /* string_lit */
			shift(74), /* bytes_lit */
			shift(75), /* true */
			shift(76), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(77), /* space */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(51), /* $, reduce: SpaceTerminal */
			reduce(51), /* root, reduce: SpaceTerminal */
			reduce(51), /* id, reduce: SpaceTerminal */
			nil,        /* . */
			reduce(51), /* if, reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(51), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(80), /* root, reduce: Space */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			reduce(80), /* if, reduce: Space */
			nil,        /* { */
			reduce(80), /* []bool, reduce: Space */
			reduce(80), /* []int64, reduce: Space */
			reduce(80), /* []int32, reduce: Space */
			reduce(80), /* []uint64, reduce: Space */
			reduce(80), /* []uint32, reduce: Space */
			reduce(80), /* []double, reduce: Space */
			reduce(80), /* []float, reduce: Space */
			reduce(80), /* []string, reduce: Space */
			reduce(80), /* [][]byte, reduce: Space */
			reduce(80), /* int64_lit, reduce: Space */
			reduce(80), /* int32_lit, reduce: Space */
			reduce(80), /* uint64_lit, reduce: Space */
			reduce(80), /* uint32_lit, reduce: Space */
			reduce(80), /* double_lit, reduce: Space */
			reduce(80), /* float_lit, reduce: Space */
			reduce(80), /* string_lit, reduce: Space */
			reduce(80), /* bytes_lit, reduce: Space */
			reduce(80), /* true, reduce: Space */
			reduce(80), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(93), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(94), /* space */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(96), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(86), /* space */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(64), /* id, reduce: Equal */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(64), /* space, reduce: Equal */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			reduce(81), /* =, reduce: Space */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(97), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			shift(98), /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(100), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(70), /* id, reduce: OpenParen */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(70), /* []bool, reduce: OpenParen */
			reduce(70), /* []int64, reduce: OpenParen */
			reduce(70), /* []int32, reduce: OpenParen */
			reduce(70), /* []uint64, reduce: OpenParen */
			reduce(70), /* []uint32, reduce: OpenParen */
			reduce(70), /* []double, reduce: OpenParen */
			reduce(70), /* []float, reduce: OpenParen */
			reduce(70), /* []string, reduce: OpenParen */
			reduce(70), /* [][]byte, reduce: OpenParen */
			reduce(70), /* int64_lit, reduce: OpenParen */
			reduce(70), /* int32_lit, reduce: OpenParen */
			reduce(70), /* uint64_lit, reduce: OpenParen */
			reduce(70), /* uint32_lit, reduce: OpenParen */
			reduce(70), /* double_lit, reduce: OpenParen */
			reduce(70), /* float_lit, reduce: OpenParen */
			reduce(70), /* string_lit, reduce: OpenParen */
			reduce(70), /* bytes_lit, reduce: OpenParen */
			reduce(70), /* true, reduce: OpenParen */
			reduce(70), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(70), /* ), reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			reduce(70), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(81), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(124), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(67),  /* int64_lit */
			shift(68),  /* int32_lit */
			shift(69),  /* uint64_lit */
			shift(70),  /* uint32_lit */
			shift(71),  /* double_lit */
			shift(72),  /* float_lit */
			shift(73),  /* string_lit */
			shift(74),  /* bytes_lit */
			shift(75),  /* true */
			shift(76),  /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(127), /* space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(130), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(131), /* space */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(133), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(39), /* then, reduce: Expr */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(39), /* space, reduce: Expr */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(40), /* then, reduce: Expr */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(40), /* space, reduce: Expr */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(38), /* then, reduce: Expr */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(50), /* then, reduce: SpaceTerminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(52), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(53), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(54), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(55), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(56), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(57), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(58), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(59), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(60), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(62), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(62), /* space, reduce: Bool */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(63), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(81), /* []bool, reduce: Space */
			reduce(81), /* []int64, reduce: Space */
			reduce(81), /* []int32, reduce: Space */
			reduce(81), /* []uint64, reduce: Space */
			reduce(81), /* []uint32, reduce: Space */
			reduce(81), /* []double, reduce: Space */
			reduce(81), /* []float, reduce: Space */
			reduce(81), /* []string, reduce: Space */
			reduce(81), /* [][]byte, reduce: Space */
			reduce(81), /* int64_lit, reduce: Space */
			reduce(81), /* int32_lit, reduce: Space */
			reduce(81), /* uint64_lit, reduce: Space */
			reduce(81), /* uint32_lit, reduce: Space */
			reduce(81), /* double_lit, reduce: Space */
			reduce(81), /* float_lit, reduce: Space */
			reduce(81), /* string_lit, reduce: Space */
			reduce(81), /* bytes_lit, reduce: Space */
			reduce(81), /* true, reduce: Space */
			reduce(81), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			shift(137), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(138), /* space */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(74), /* id, reduce: OpenCurly */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(74), /* []bool, reduce: OpenCurly */
			reduce(74), /* []int64, reduce: OpenCurly */
			reduce(74), /* []int32, reduce: OpenCurly */
			reduce(74), /* []uint64, reduce: OpenCurly */
			reduce(74), /* []uint32, reduce: OpenCurly */
			reduce(74), /* []double, reduce: OpenCurly */
			reduce(74), /* []float, reduce: OpenCurly */
			reduce(74), /* []string, reduce: OpenCurly */
			reduce(74), /* [][]byte, reduce: OpenCurly */
			reduce(74), /* int64_lit, reduce: OpenCurly */
			reduce(74), /* int32_lit, reduce: OpenCurly */
			reduce(74), /* uint64_lit, reduce: OpenCurly */
			reduce(74), /* uint32_lit, reduce: OpenCurly */
			reduce(74), /* double_lit, reduce: OpenCurly */
			reduce(74), /* float_lit, reduce: OpenCurly */
			reduce(74), /* string_lit, reduce: OpenCurly */
			reduce(74), /* bytes_lit, reduce: OpenCurly */
			reduce(74), /* true, reduce: OpenCurly */
			reduce(74), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(74), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(74), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(81), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(163), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(80), /* $, reduce: Space */
			reduce(80), /* root, reduce: Space */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			reduce(80), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(97),  /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(165), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(167), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(168), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(98),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(99),  /* space */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(169), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(130), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(131), /* space */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(65), /* id, reduce: Equal */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(65), /* space, reduce: Equal */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			reduce(80), /* =, reduce: Space */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(175), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(176), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(71), /* id, reduce: OpenParen */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(71), /* []bool, reduce: OpenParen */
			reduce(71), /* []int64, reduce: OpenParen */
			reduce(71), /* []int32, reduce: OpenParen */
			reduce(71), /* []uint64, reduce: OpenParen */
			reduce(71), /* []uint32, reduce: OpenParen */
			reduce(71), /* []double, reduce: OpenParen */
			reduce(71), /* []float, reduce: OpenParen */
			reduce(71), /* []string, reduce: OpenParen */
			reduce(71), /* [][]byte, reduce: OpenParen */
			reduce(71), /* int64_lit, reduce: OpenParen */
			reduce(71), /* int32_lit, reduce: OpenParen */
			reduce(71), /* uint64_lit, reduce: OpenParen */
			reduce(71), /* uint32_lit, reduce: OpenParen */
			reduce(71), /* double_lit, reduce: OpenParen */
			reduce(71), /* float_lit, reduce: OpenParen */
			reduce(71), /* string_lit, reduce: OpenParen */
			reduce(71), /* bytes_lit, reduce: OpenParen */
			reduce(71), /* true, reduce: OpenParen */
			reduce(71), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(71), /* ), reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			reduce(71), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(80), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(179), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			shift(51),  /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(52),  /* space */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(180), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(183), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(184), /* space */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(36), /* ), reduce: Exprs */
			nil,        /* } */
			reduce(36), /* ,, reduce: Exprs */
			reduce(36), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(185), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(39), /* ), reduce: Expr */
			nil,        /* } */
			reduce(39), /* ,, reduce: Expr */
			reduce(39), /* space, reduce: Expr */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(31), /* $, reduce: Function */
			reduce(31), /* root, reduce: Function */
			reduce(31), /* id, reduce: Function */
			nil,        /* . */
			reduce(31), /* if, reduce: Function */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(31), /* space, reduce: Function */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(40), /* ), reduce: Expr */
			nil,        /* } */
			reduce(40), /* ,, reduce: Expr */
			reduce(40), /* space, reduce: Expr */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(38), /* ), reduce: Expr */
			nil,        /* } */
			reduce(38), /* ,, reduce: Expr */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(50), /* ), reduce: SpaceTerminal */
			nil,        /* } */
			reduce(50), /* ,, reduce: SpaceTerminal */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(52), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(52), /* ,, reduce: Terminal */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(53), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(53), /* ,, reduce: Terminal */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(54), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(54), /* ,, reduce: Terminal */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(55), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(55), /* ,, reduce: Terminal */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(56), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(56), /* ,, reduce: Terminal */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(57), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(57), /* ,, reduce: Terminal */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(58), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(58), /* ,, reduce: Terminal */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(59), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(59), /* ,, reduce: Terminal */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(60), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(60), /* ,, reduce: Terminal */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(62), /* ), reduce: Bool */
			nil,        /* } */
			reduce(62), /* ,, reduce: Bool */
			reduce(62), /* space, reduce: Bool */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(63), /* ), reduce: Bool */
			nil,        /* } */
			reduce(63), /* ,, reduce: Bool */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(72), /* $, reduce: CloseParen */
			reduce(72), /* root, reduce: CloseParen */
			reduce(72), /* id, reduce: CloseParen */
			nil,        /* . */
			reduce(72), /* if, reduce: CloseParen */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(72), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(81), /* []bool, reduce: Space */
			reduce(81), /* []int64, reduce: Space */
			reduce(81), /* []int32, reduce: Space */
			reduce(81), /* []uint64, reduce: Space */
			reduce(81), /* []uint32, reduce: Space */
			reduce(81), /* []double, reduce: Space */
			reduce(81), /* []float, reduce: Space */
			reduce(81), /* []string, reduce: Space */
			reduce(81), /* [][]byte, reduce: Space */
			reduce(81), /* int64_lit, reduce: Space */
			reduce(81), /* int32_lit, reduce: Space */
			reduce(81), /* uint64_lit, reduce: Space */
			reduce(81), /* uint32_lit, reduce: Space */
			reduce(81), /* double_lit, reduce: Space */
			reduce(81), /* float_lit, reduce: Space */
			reduce(81), /* string_lit, reduce: Space */
			reduce(81), /* bytes_lit, reduce: Space */
			reduce(81), /* true, reduce: Space */
			reduce(81), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(81), /* ), reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(133), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(51), /* then, reduce: SpaceTerminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(51), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(80), /* []bool, reduce: Space */
			reduce(80), /* []int64, reduce: Space */
			reduce(80), /* []int32, reduce: Space */
			reduce(80), /* []uint64, reduce: Space */
			reduce(80), /* []uint32, reduce: Space */
			reduce(80), /* []double, reduce: Space */
			reduce(80), /* []float, reduce: Space */
			reduce(80), /* []string, reduce: Space */
			reduce(80), /* [][]byte, reduce: Space */
			reduce(80), /* int64_lit, reduce: Space */
			reduce(80), /* int32_lit, reduce: Space */
			reduce(80), /* uint64_lit, reduce: Space */
			reduce(80), /* uint32_lit, reduce: Space */
			reduce(80), /* double_lit, reduce: Space */
			reduce(80), /* float_lit, reduce: Space */
			reduce(80), /* string_lit, reduce: Space */
			reduce(80), /* bytes_lit, reduce: Space */
			reduce(80), /* true, reduce: Space */
			reduce(80), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(195), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(196), /* space */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(198), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(200), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(66), /* id, reduce: Then */
			nil,        /* . */
			nil,        /* if */
			reduce(66), /* {, reduce: Then */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(66), /* space, reduce: Then */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(81), /* then, reduce: Space */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(98),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(202), /* space */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(203), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(207), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(81), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(211), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(75), /* id, reduce: OpenCurly */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(75), /* []bool, reduce: OpenCurly */
			reduce(75), /* []int64, reduce: OpenCurly */
			reduce(75), /* []int32, reduce: OpenCurly */
			reduce(75), /* []uint64, reduce: OpenCurly */
			reduce(75), /* []uint32, reduce: OpenCurly */
			reduce(75), /* []double, reduce: OpenCurly */
			reduce(75), /* []float, reduce: OpenCurly */
			reduce(75), /* []string, reduce: OpenCurly */
			reduce(75), /* [][]byte, reduce: OpenCurly */
			reduce(75), /* int64_lit, reduce: OpenCurly */
			reduce(75), /* int32_lit, reduce: OpenCurly */
			reduce(75), /* uint64_lit, reduce: OpenCurly */
			reduce(75), /* uint32_lit, reduce: OpenCurly */
			reduce(75), /* double_lit, reduce: OpenCurly */
			reduce(75), /* float_lit, reduce: OpenCurly */
			reduce(75), /* string_lit, reduce: OpenCurly */
			reduce(75), /* bytes_lit, reduce: OpenCurly */
			reduce(75), /* true, reduce: OpenCurly */
			reduce(75), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(75), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(75), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(80), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(212), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(215), /* } */
			nil,        /* , */
			shift(216), /* space */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(36), /* }, reduce: Exprs */
			reduce(36), /* ,, reduce: Exprs */
			reduce(36), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(217), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(35), /* $, reduce: List */
			reduce(35), /* root, reduce: List */
			reduce(35), /* id, reduce: List */
			nil,        /* . */
			reduce(35), /* if, reduce: List */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(35), /* space, reduce: List */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(39), /* }, reduce: Expr */
			reduce(39), /* ,, reduce: Expr */
			reduce(39), /* space, reduce: Expr */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(40), /* }, reduce: Expr */
			reduce(40), /* ,, reduce: Expr */
			reduce(40), /* space, reduce: Expr */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(38), /* }, reduce: Expr */
			reduce(38), /* ,, reduce: Expr */
			reduce(38), /* space, reduce: Expr */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(50), /* }, reduce: SpaceTerminal */
			reduce(50), /* ,, reduce: SpaceTerminal */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(52), /* }, reduce: Terminal */
			reduce(52), /* ,, reduce: Terminal */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S150
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(53), /* }, reduce: Terminal */
			reduce(53), /* ,, reduce: Terminal */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(54), /* }, reduce: Terminal */
			reduce(54), /* ,, reduce: Terminal */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(55), /* }, reduce: Terminal */
			reduce(55), /* ,, reduce: Terminal */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(56), /* }, reduce: Terminal */
			reduce(56), /* ,, reduce: Terminal */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S154
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(57), /* }, reduce: Terminal */
			reduce(57), /* ,, reduce: Terminal */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(58), /* }, reduce: Terminal */
			reduce(58), /* ,, reduce: Terminal */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(59), /* }, reduce: Terminal */
			reduce(59), /* ,, reduce: Terminal */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S157
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(60), /* }, reduce: Terminal */
			reduce(60), /* ,, reduce: Terminal */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S158
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(62), /* }, reduce: Bool */
			reduce(62), /* ,, reduce: Bool */
			reduce(62), /* space, reduce: Bool */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(63), /* }, reduce: Bool */
			reduce(63), /* ,, reduce: Bool */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(76), /* $, reduce: CloseCurly */
			reduce(76), /* root, reduce: CloseCurly */
			reduce(76), /* id, reduce: CloseCurly */
			nil,        /* . */
			reduce(76), /* if, reduce: CloseCurly */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(81), /* []bool, reduce: Space */
			reduce(81), /* []int64, reduce: Space */
			reduce(81), /* []int32, reduce: Space */
			reduce(81), /* []uint64, reduce: Space */
			reduce(81), /* []uint32, reduce: Space */
			reduce(81), /* []double, reduce: Space */
			reduce(81), /* []float, reduce: Space */
			reduce(81), /* []string, reduce: Space */
			reduce(81), /* [][]byte, reduce: Space */
			reduce(81), /* int64_lit, reduce: Space */
			reduce(81), /* int32_lit, reduce: Space */
			reduce(81), /* uint64_lit, reduce: Space */
			reduce(81), /* uint32_lit, reduce: Space */
			reduce(81), /* double_lit, reduce: Space */
			reduce(81), /* float_lit, reduce: Space */
			reduce(81), /* string_lit, reduce: Space */
			reduce(81), /* bytes_lit, reduce: Space */
			reduce(81), /* true, reduce: Space */
			reduce(81), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(81), /* }, reduce: Space */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(168), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(224), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S165
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S166
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(225), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S167
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(226), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S168
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S169
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(179), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			shift(51),  /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(52),  /* space */

		},
	},
	actionRow{ // S170
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S171
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(29), /* $, reduce: Function */
			reduce(29), /* root, reduce: Function */
			reduce(29), /* id, reduce: Function */
			nil,        /* . */
			reduce(29), /* if, reduce: Function */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(29), /* space, reduce: Function */

		},
	},
	actionRow{ // S172
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(198), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(200), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S173
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(34), /* $, reduce: List */
			reduce(34), /* root, reduce: List */
			reduce(34), /* id, reduce: List */
			nil,        /* . */
			reduce(34), /* if, reduce: List */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(34), /* space, reduce: List */

		},
	},
	actionRow{ // S174
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S175
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(232), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S176
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(233), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S177
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(235), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S178
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(237), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S179
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(238), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S180
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(185), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S181
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S182
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(51), /* ), reduce: SpaceTerminal */
			nil,        /* } */
			reduce(51), /* ,, reduce: SpaceTerminal */
			reduce(51), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S183
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(73), /* $, reduce: CloseParen */
			reduce(73), /* root, reduce: CloseParen */
			reduce(73), /* id, reduce: CloseParen */
			nil,        /* . */
			reduce(73), /* if, reduce: CloseParen */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(73), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S184
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(80), /* []bool, reduce: Space */
			reduce(80), /* []int64, reduce: Space */
			reduce(80), /* []int32, reduce: Space */
			reduce(80), /* []uint64, reduce: Space */
			reduce(80), /* []uint32, reduce: Space */
			reduce(80), /* []double, reduce: Space */
			reduce(80), /* []float, reduce: Space */
			reduce(80), /* []string, reduce: Space */
			reduce(80), /* [][]byte, reduce: Space */
			reduce(80), /* int64_lit, reduce: Space */
			reduce(80), /* int32_lit, reduce: Space */
			reduce(80), /* uint64_lit, reduce: Space */
			reduce(80), /* uint32_lit, reduce: Space */
			reduce(80), /* double_lit, reduce: Space */
			reduce(80), /* float_lit, reduce: Space */
			reduce(80), /* string_lit, reduce: Space */
			reduce(80), /* bytes_lit, reduce: Space */
			reduce(80), /* true, reduce: Space */
			reduce(80), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(80), /* ), reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S185
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S186
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(245), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S187
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(183), /* ) */
			nil,        /* } */
			shift(246), /* , */
			shift(247), /* space */

		},
	},
	actionRow{ // S188
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(30), /* $, reduce: Function */
			reduce(30), /* root, reduce: Function */
			reduce(30), /* id, reduce: Function */
			nil,        /* . */
			reduce(30), /* if, reduce: Function */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(30), /* space, reduce: Function */

		},
	},
	actionRow{ // S189
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(77),  /* space */

		},
	},
	actionRow{ // S190
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(78), /* id, reduce: Comma */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(78), /* []bool, reduce: Comma */
			reduce(78), /* []int64, reduce: Comma */
			reduce(78), /* []int32, reduce: Comma */
			reduce(78), /* []uint64, reduce: Comma */
			reduce(78), /* []uint32, reduce: Comma */
			reduce(78), /* []double, reduce: Comma */
			reduce(78), /* []float, reduce: Comma */
			reduce(78), /* []string, reduce: Comma */
			reduce(78), /* [][]byte, reduce: Comma */
			reduce(78), /* int64_lit, reduce: Comma */
			reduce(78), /* int32_lit, reduce: Comma */
			reduce(78), /* uint64_lit, reduce: Comma */
			reduce(78), /* uint32_lit, reduce: Comma */
			reduce(78), /* double_lit, reduce: Comma */
			reduce(78), /* float_lit, reduce: Comma */
			reduce(78), /* string_lit, reduce: Comma */
			reduce(78), /* bytes_lit, reduce: Comma */
			reduce(78), /* true, reduce: Comma */
			reduce(78), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(78), /* space, reduce: Comma */

		},
	},
	actionRow{ // S191
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(81), /* ), reduce: Space */
			nil,        /* } */
			reduce(81), /* ,, reduce: Space */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S192
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(253), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S193
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(207), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S194
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(211), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S195
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(67), /* id, reduce: Then */
			nil,        /* . */
			nil,        /* if */
			reduce(67), /* {, reduce: Then */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(67), /* space, reduce: Then */

		},
	},
	actionRow{ // S196
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(80), /* then, reduce: Space */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S197
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(258), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(259), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(260), /* space */

		},
	},
	actionRow{ // S198
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(27), /* else, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(27), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S199
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(263), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(264), /* space */

		},
	},
	actionRow{ // S200
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S201
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(81), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			reduce(81), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S202
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(80), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S203
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(269), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S204
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(180), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(184), /* space */

		},
	},
	actionRow{ // S205
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(207), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S206
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(31), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(31), /* space, reduce: Function */

		},
	},
	actionRow{ // S207
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(72), /* then, reduce: CloseParen */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(72), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S208
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(212), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(273), /* } */
			nil,        /* , */
			shift(216), /* space */

		},
	},
	actionRow{ // S209
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(35), /* then, reduce: List */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(35), /* space, reduce: List */

		},
	},
	actionRow{ // S210
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(211), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S211
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(76), /* then, reduce: CloseCurly */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S212
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(217), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(56),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S213
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(79), /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(81), /* space */

		},
	},
	actionRow{ // S214
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(51), /* }, reduce: SpaceTerminal */
			reduce(51), /* ,, reduce: SpaceTerminal */
			reduce(51), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S215
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(77), /* $, reduce: CloseCurly */
			reduce(77), /* root, reduce: CloseCurly */
			reduce(77), /* id, reduce: CloseCurly */
			nil,        /* . */
			reduce(77), /* if, reduce: CloseCurly */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S216
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(80), /* []bool, reduce: Space */
			reduce(80), /* []int64, reduce: Space */
			reduce(80), /* []int32, reduce: Space */
			reduce(80), /* []uint64, reduce: Space */
			reduce(80), /* []uint32, reduce: Space */
			reduce(80), /* []double, reduce: Space */
			reduce(80), /* []float, reduce: Space */
			reduce(80), /* []string, reduce: Space */
			reduce(80), /* [][]byte, reduce: Space */
			reduce(80), /* int64_lit, reduce: Space */
			reduce(80), /* int32_lit, reduce: Space */
			reduce(80), /* uint64_lit, reduce: Space */
			reduce(80), /* uint32_lit, reduce: Space */
			reduce(80), /* double_lit, reduce: Space */
			reduce(80), /* float_lit, reduce: Space */
			reduce(80), /* string_lit, reduce: Space */
			reduce(80), /* bytes_lit, reduce: Space */
			reduce(80), /* true, reduce: Space */
			reduce(80), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(80), /* }, reduce: Space */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S217
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(278), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S218
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(282), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S219
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(215), /* } */
			shift(246), /* , */
			shift(283), /* space */

		},
	},
	actionRow{ // S220
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(33), /* $, reduce: List */
			reduce(33), /* root, reduce: List */
			reduce(33), /* id, reduce: List */
			nil,        /* . */
			reduce(33), /* if, reduce: List */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(33), /* space, reduce: List */

		},
	},
	actionRow{ // S221
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(77),  /* space */

		},
	},
	actionRow{ // S222
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(81), /* }, reduce: Space */
			reduce(81), /* ,, reduce: Space */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S223
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(289), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S224
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			nil,       /* []bool */
			nil,       /* []int64 */
			nil,       /* []int32 */
			nil,       /* []uint64 */
			nil,       /* []uint32 */
			nil,       /* []double */
			nil,       /* []float */
			nil,       /* []string */
			nil,       /* [][]byte */
			nil,       /* int64_lit */
			nil,       /* int32_lit */
			nil,       /* uint64_lit */
			nil,       /* uint32_lit */
			nil,       /* double_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* bytes_lit */
			nil,       /* true */
			nil,       /* false */
			shift(51), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(52), /* space */

		},
	},
	actionRow{ // S225
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(290), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S226
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(291), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S227
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(293), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S228
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(295), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(86),  /* space */

		},
	},
	actionRow{ // S229
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(28), /* $, reduce: Function */
			reduce(28), /* root, reduce: Function */
			reduce(28), /* id, reduce: Function */
			nil,        /* . */
			reduce(28), /* if, reduce: Function */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(28), /* space, reduce: Function */

		},
	},
	actionRow{ // S230
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(263), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(264), /* space */

		},
	},
	actionRow{ // S231
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(32), /* $, reduce: List */
			reduce(32), /* root, reduce: List */
			reduce(32), /* id, reduce: List */
			nil,        /* . */
			reduce(32), /* if, reduce: List */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(32), /* space, reduce: List */

		},
	},
	actionRow{ // S232
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(297), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S233
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: Root */
			reduce(13), /* root, reduce: Root */
			reduce(13), /* id, reduce: Root */
			nil,        /* . */
			reduce(13), /* if, reduce: Root */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(13), /* space, reduce: Root */

		},
	},
	actionRow{ // S234
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(298), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S235
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(21), /* $, reduce: Transition */
			reduce(21), /* root, reduce: Transition */
			reduce(21), /* id, reduce: Transition */
			nil,        /* . */
			reduce(21), /* if, reduce: Transition */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(21), /* space, reduce: Transition */

		},
	},
	actionRow{ // S236
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(299), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S237
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(17), /* $, reduce: Init */
			reduce(17), /* root, reduce: Init */
			reduce(17), /* id, reduce: Init */
			nil,        /* . */
			reduce(17), /* if, reduce: Init */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(17), /* space, reduce: Init */

		},
	},
	actionRow{ // S238
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(61), /* $, reduce: Terminal */
			reduce(61), /* root, reduce: Terminal */
			reduce(61), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(61), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S239
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(245), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S240
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(253), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S241
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(304), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S242
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(180), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(184), /* space */

		},
	},
	actionRow{ // S243
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(245), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S244
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(31), /* ), reduce: Function */
			nil,        /* } */
			reduce(31), /* ,, reduce: Function */
			reduce(31), /* space, reduce: Function */

		},
	},
	actionRow{ // S245
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(72), /* ), reduce: CloseParen */
			nil,        /* } */
			reduce(72), /* ,, reduce: CloseParen */
			reduce(72), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S246
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(79), /* id, reduce: Comma */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(79), /* []bool, reduce: Comma */
			reduce(79), /* []int64, reduce: Comma */
			reduce(79), /* []int32, reduce: Comma */
			reduce(79), /* []uint64, reduce: Comma */
			reduce(79), /* []uint32, reduce: Comma */
			reduce(79), /* []double, reduce: Comma */
			reduce(79), /* []float, reduce: Comma */
			reduce(79), /* []string, reduce: Comma */
			reduce(79), /* [][]byte, reduce: Comma */
			reduce(79), /* int64_lit, reduce: Comma */
			reduce(79), /* int32_lit, reduce: Comma */
			reduce(79), /* uint64_lit, reduce: Comma */
			reduce(79), /* uint32_lit, reduce: Comma */
			reduce(79), /* double_lit, reduce: Comma */
			reduce(79), /* float_lit, reduce: Comma */
			reduce(79), /* string_lit, reduce: Comma */
			reduce(79), /* bytes_lit, reduce: Comma */
			reduce(79), /* true, reduce: Comma */
			reduce(79), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(79), /* space, reduce: Comma */

		},
	},
	actionRow{ // S247
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(80), /* ), reduce: Space */
			nil,        /* } */
			reduce(80), /* ,, reduce: Space */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S248
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(180), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(127), /* space */

		},
	},
	actionRow{ // S249
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(37), /* ), reduce: Exprs */
			nil,        /* } */
			reduce(37), /* ,, reduce: Exprs */
			reduce(37), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S250
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(212), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(308), /* } */
			nil,        /* , */
			shift(216), /* space */

		},
	},
	actionRow{ // S251
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(35), /* ), reduce: List */
			nil,        /* } */
			reduce(35), /* ,, reduce: List */
			reduce(35), /* space, reduce: List */

		},
	},
	actionRow{ // S252
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(253), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S253
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(76), /* ), reduce: CloseCurly */
			nil,        /* } */
			reduce(76), /* ,, reduce: CloseCurly */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S254
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(207), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S255
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(29), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(29), /* space, reduce: Function */

		},
	},
	actionRow{ // S256
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(34), /* then, reduce: List */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(34), /* space, reduce: List */

		},
	},
	actionRow{ // S257
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(211), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S258
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(25), /* else, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(25), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S259
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S260
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(80), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			reduce(80), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S261
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(314), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(315), /* space */

		},
	},
	actionRow{ // S262
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(317), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(319), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S263
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(68), /* id, reduce: Else */
			nil,        /* . */
			nil,        /* if */
			reduce(68), /* {, reduce: Else */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(68), /* space, reduce: Else */

		},
	},
	actionRow{ // S264
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(81), /* else, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S265
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(320), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(321), /* space */

		},
	},
	actionRow{ // S266
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(324), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S267
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(60), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(67), /* int64_lit */
			shift(68), /* int32_lit */
			shift(69), /* uint64_lit */
			shift(70), /* uint32_lit */
			shift(71), /* double_lit */
			shift(72), /* float_lit */
			shift(73), /* string_lit */
			shift(74), /* bytes_lit */
			shift(75), /* true */
			shift(76), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(77), /* space */

		},
	},
	actionRow{ // S268
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			reduce(81), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S269
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(327), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S270
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(73), /* then, reduce: CloseParen */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(73), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S271
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			shift(246), /* , */
			shift(247), /* space */

		},
	},
	actionRow{ // S272
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(30), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(30), /* space, reduce: Function */

		},
	},
	actionRow{ // S273
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(77), /* then, reduce: CloseCurly */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S274
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(273), /* } */
			shift(246), /* , */
			shift(283), /* space */

		},
	},
	actionRow{ // S275
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(33), /* then, reduce: List */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(33), /* space, reduce: List */

		},
	},
	actionRow{ // S276
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(282), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(123), /* space */

		},
	},
	actionRow{ // S277
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(141), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(289), /* } */
			nil,        /* , */
			shift(161), /* space */

		},
	},
	actionRow{ // S278
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(332), /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S279
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(180), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(112), /* int64_lit */
			shift(113), /* int32_lit */
			shift(114), /* uint64_lit */
			shift(115), /* uint32_lit */
			shift(116), /* double_lit */
			shift(117), /* float_lit */
			shift(118), /* string_lit */
			shift(119), /* bytes_lit */
			shift(120), /* true */
			shift(121), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(333), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(184), /* space */

		},
	},
	actionRow{ // S280
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(282), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S281
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(31), /* }, reduce: Function */
			reduce(31), /* ,, reduce: Function */
			reduce(31), /* space, reduce: Function */

		},
	},
	actionRow{ // S282
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(72), /* }, reduce: CloseParen */
			reduce(72), /* ,, reduce: CloseParen */
			reduce(72), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S283
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(80), /* }, reduce: Space */
			reduce(80), /* ,, reduce: Space */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S284
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(212), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(127), /* space */

		},
	},
	actionRow{ // S285
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(37), /* }, reduce: Exprs */
			reduce(37), /* ,, reduce: Exprs */
			reduce(37), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S286
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(212), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			shift(17),  /* []bool */
			shift(18),  /* []int64 */
			shift(19),  /* []int32 */
			shift(20),  /* []uint64 */
			shift(21),  /* []uint32 */
			shift(22),  /* []double */
			shift(23),  /* []float */
			shift(24),  /* []string */
			shift(25),  /* [][]byte */
			shift(150), /* int64_lit */
			shift(151), /* int32_lit */
			shift(152), /* uint64_lit */
			shift(153), /* uint32_lit */
			shift(154), /* double_lit */
			shift(155), /* float_lit */
			shift(156), /* string_lit */
			shift(157), /* bytes_lit */
			shift(158), /* true */
			shift(159), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(336), /* } */
			nil,        /* , */
			shift(216), /* space */

		},
	},
	actionRow{ // S287
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(35), /* }, reduce: List */
			reduce(35), /* ,, reduce: List */
			reduce(35), /* space, reduce: List */

		},
	},
	actionRow{ // S288
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(289), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S289
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(76), /* }, reduce: CloseCurly */
			reduce(76), /* ,, reduce: CloseCurly */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S290
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(339), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S291
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: Root */
			reduce(12), /* root, reduce: Root */
			reduce(12), /* id, reduce: Root */
			nil,        /* . */
			reduce(12), /* if, reduce: Root */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(12), /* space, reduce: Root */

		},
	},
	actionRow{ // S292
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(340), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S293
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(20), /* $, reduce: Transition */
			reduce(20), /* root, reduce: Transition */
			reduce(20), /* id, reduce: Transition */
			nil,        /* . */
			reduce(20), /* if, reduce: Transition */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(20), /* space, reduce: Transition */

		},
	},
	actionRow{ // S294
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(341), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(164), /* space */

		},
	},
	actionRow{ // S295
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(16), /* $, reduce: Init */
			reduce(16), /* root, reduce: Init */
			reduce(16), /* id, reduce: Init */
			nil,        /* . */
			reduce(16), /* if, reduce: Init */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(16), /* space, reduce: Init */

		},
	},
	actionRow{ // S296
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(317), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(319), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S297
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: Root */
			reduce(11), /* root, reduce: Root */
			reduce(11), /* id, reduce: Root */
			nil,        /* . */
			reduce(11), /* if, reduce: Root */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(11), /* space, reduce: Root */

		},
	},
	actionRow{ // S298
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(19), /* $, reduce: Transition */
			reduce(19), /* root, reduce: Transition */
			reduce(19), /* id, reduce: Transition */
			nil,        /* . */
			reduce(19), /* if, reduce: Transition */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(19), /* space, reduce: Transition */

		},
	},
	actionRow{ // S299
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(15), /* $, reduce: Init */
			reduce(15), /* root, reduce: Init */
			reduce(15), /* id, reduce: Init */
			nil,        /* . */
			reduce(15), /* if, reduce: Init */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(15), /* space, reduce: Init */

		},
	},
	actionRow{ // S300
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(245), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S301
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(29), /* ), reduce: Function */
			nil,        /* } */
			reduce(29), /* ,, reduce: Function */
			reduce(29), /* space, reduce: Function */

		},
	},
	actionRow{ // S302
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(34), /* ), reduce: List */
			nil,        /* } */
			reduce(34), /* ,, reduce: List */
			reduce(34), /* space, reduce: List */

		},
	},
	actionRow{ // S303
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(253), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S304
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(345), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S305
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(73), /* ), reduce: CloseParen */
			nil,        /* } */
			reduce(73), /* ,, reduce: CloseParen */
			reduce(73), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S306
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			shift(246), /* , */
			shift(247), /* space */

		},
	},
	actionRow{ // S307
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(30), /* ), reduce: Function */
			nil,        /* } */
			reduce(30), /* ,, reduce: Function */
			reduce(30), /* space, reduce: Function */

		},
	},
	actionRow{ // S308
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(77), /* ), reduce: CloseCurly */
			nil,        /* } */
			reduce(77), /* ,, reduce: CloseCurly */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S309
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(308), /* } */
			shift(246), /* , */
			shift(283), /* space */

		},
	},
	actionRow{ // S310
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(33), /* ), reduce: List */
			nil,        /* } */
			reduce(33), /* ,, reduce: List */
			reduce(33), /* space, reduce: List */

		},
	},
	actionRow{ // S311
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(28), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(28), /* space, reduce: Function */

		},
	},
	actionRow{ // S312
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(32), /* then, reduce: List */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(32), /* space, reduce: List */

		},
	},
	actionRow{ // S313
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(324), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S314
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(69), /* id, reduce: Else */
			nil,        /* . */
			nil,        /* if */
			reduce(69), /* {, reduce: Else */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(69), /* space, reduce: Else */

		},
	},
	actionRow{ // S315
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(80), /* else, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S316
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(347), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(348), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(260), /* space */

		},
	},
	actionRow{ // S317
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(27), /* $, reduce: StateExpr */
			reduce(27), /* root, reduce: StateExpr */
			reduce(27), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(27), /* if, reduce: StateExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(27), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S318
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(23), /* $, reduce: IfExpr */
			reduce(23), /* root, reduce: IfExpr */
			reduce(23), /* id, reduce: IfExpr */
			nil,        /* . */
			reduce(23), /* if, reduce: IfExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(23), /* space, reduce: IfExpr */

		},
	},
	actionRow{ // S319
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S320
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(60), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* { */
			shift(17), /* []bool */
			shift(18), /* []int64 */
			shift(19), /* []int32 */
			shift(20), /* []uint64 */
			shift(21), /* []uint32 */
			shift(22), /* []double */
			shift(23), /* []float */
			shift(24), /* []string */
			shift(25), /* [][]byte */
			shift(67), /* int64_lit */
			shift(68), /* int32_lit */
			shift(69), /* uint64_lit */
			shift(70), /* uint32_lit */
			shift(71), /* double_lit */
			shift(72), /* float_lit */
			shift(73), /* string_lit */
			shift(74), /* bytes_lit */
			shift(75), /* true */
			shift(76), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(77), /* space */

		},
	},
	actionRow{ // S321
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			reduce(80), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S322
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(351), /* } */
			nil,        /* , */
			shift(352), /* space */

		},
	},
	actionRow{ // S323
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(26), /* else, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(26), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S324
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(76), /* else, reduce: CloseCurly */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S325
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(81), /* }, reduce: Space */
			nil,        /* , */
			reduce(81), /* space, reduce: Space */

		},
	},
	actionRow{ // S326
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(130), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(131), /* space */

		},
	},
	actionRow{ // S327
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(61), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S328
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(282), /* ) */
			nil,        /* } */
			shift(190), /* , */
			shift(191), /* space */

		},
	},
	actionRow{ // S329
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(29), /* }, reduce: Function */
			reduce(29), /* ,, reduce: Function */
			reduce(29), /* space, reduce: Function */

		},
	},
	actionRow{ // S330
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(34), /* }, reduce: List */
			reduce(34), /* ,, reduce: List */
			reduce(34), /* space, reduce: List */

		},
	},
	actionRow{ // S331
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(289), /* } */
			shift(190), /* , */
			shift(222), /* space */

		},
	},
	actionRow{ // S332
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(356), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			nil,        /* space */

		},
	},
	actionRow{ // S333
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(73), /* }, reduce: CloseParen */
			reduce(73), /* ,, reduce: CloseParen */
			reduce(73), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S334
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(333), /* ) */
			nil,        /* } */
			shift(246), /* , */
			shift(247), /* space */

		},
	},
	actionRow{ // S335
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(30), /* }, reduce: Function */
			reduce(30), /* ,, reduce: Function */
			reduce(30), /* space, reduce: Function */

		},
	},
	actionRow{ // S336
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(77), /* }, reduce: CloseCurly */
			reduce(77), /* ,, reduce: CloseCurly */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S337
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(336), /* } */
			shift(246), /* , */
			shift(283), /* space */

		},
	},
	actionRow{ // S338
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(33), /* }, reduce: List */
			reduce(33), /* ,, reduce: List */
			reduce(33), /* space, reduce: List */

		},
	},
	actionRow{ // S339
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: Root */
			reduce(10), /* root, reduce: Root */
			reduce(10), /* id, reduce: Root */
			nil,        /* . */
			reduce(10), /* if, reduce: Root */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(10), /* space, reduce: Root */

		},
	},
	actionRow{ // S340
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(18), /* $, reduce: Transition */
			reduce(18), /* root, reduce: Transition */
			reduce(18), /* id, reduce: Transition */
			nil,        /* . */
			reduce(18), /* if, reduce: Transition */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(18), /* space, reduce: Transition */

		},
	},
	actionRow{ // S341
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(14), /* $, reduce: Init */
			reduce(14), /* root, reduce: Init */
			reduce(14), /* id, reduce: Init */
			nil,        /* . */
			reduce(14), /* if, reduce: Init */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(14), /* space, reduce: Init */

		},
	},
	actionRow{ // S342
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(22), /* $, reduce: IfExpr */
			reduce(22), /* root, reduce: IfExpr */
			reduce(22), /* id, reduce: IfExpr */
			nil,        /* . */
			reduce(22), /* if, reduce: IfExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(22), /* space, reduce: IfExpr */

		},
	},
	actionRow{ // S343
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(28), /* ), reduce: Function */
			nil,        /* } */
			reduce(28), /* ,, reduce: Function */
			reduce(28), /* space, reduce: Function */

		},
	},
	actionRow{ // S344
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(32), /* ), reduce: List */
			nil,        /* } */
			reduce(32), /* ,, reduce: List */
			reduce(32), /* space, reduce: List */

		},
	},
	actionRow{ // S345
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(61), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(61), /* ,, reduce: Terminal */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S346
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(24), /* else, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(24), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S347
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(25), /* $, reduce: StateExpr */
			reduce(25), /* root, reduce: StateExpr */
			reduce(25), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(25), /* if, reduce: StateExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(25), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S348
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S349
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S350
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(130), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(131), /* space */

		},
	},
	actionRow{ // S351
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(77), /* else, reduce: CloseCurly */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S352
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(80), /* }, reduce: Space */
			nil,        /* , */
			reduce(80), /* space, reduce: Space */

		},
	},
	actionRow{ // S353
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(198), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(200), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S354
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(28), /* }, reduce: Function */
			reduce(28), /* ,, reduce: Function */
			reduce(28), /* space, reduce: Function */

		},
	},
	actionRow{ // S355
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(32), /* }, reduce: List */
			reduce(32), /* ,, reduce: List */
			reduce(32), /* space, reduce: List */

		},
	},
	actionRow{ // S356
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(61), /* }, reduce: Terminal */
			reduce(61), /* ,, reduce: Terminal */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S357
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(160), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S358
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(215), /* } */
			nil,        /* , */
			shift(352), /* space */

		},
	},
	actionRow{ // S359
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(26), /* $, reduce: StateExpr */
			reduce(26), /* root, reduce: StateExpr */
			reduce(26), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(26), /* if, reduce: StateExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(26), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S360
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(198), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(200), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S361
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(263), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(264), /* space */

		},
	},
	actionRow{ // S362
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(24), /* $, reduce: StateExpr */
			reduce(24), /* root, reduce: StateExpr */
			reduce(24), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(24), /* if, reduce: StateExpr */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(24), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S363
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(263), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(264), /* space */

		},
	},
	actionRow{ // S364
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(367), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(369), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S365
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(367), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(369), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(201), /* space */

		},
	},
	actionRow{ // S366
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(371), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(372), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(260), /* space */

		},
	},
	actionRow{ // S367
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(27), /* }, reduce: StateExpr */
			nil,        /* , */
			reduce(27), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S368
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(23), /* }, reduce: IfExpr */
			nil,        /* , */
			reduce(23), /* space, reduce: IfExpr */

		},
	},
	actionRow{ // S369
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S370
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(22), /* }, reduce: IfExpr */
			nil,        /* , */
			reduce(22), /* space, reduce: IfExpr */

		},
	},
	actionRow{ // S371
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(25), /* }, reduce: StateExpr */
			nil,        /* , */
			reduce(25), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S372
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(267), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(268), /* space */

		},
	},
	actionRow{ // S373
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(377), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S374
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(377), /* } */
			nil,        /* , */
			shift(325), /* space */

		},
	},
	actionRow{ // S375
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(379), /* } */
			nil,        /* , */
			shift(352), /* space */

		},
	},
	actionRow{ // S376
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(26), /* }, reduce: StateExpr */
			nil,        /* , */
			reduce(26), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S377
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(76), /* }, reduce: CloseCurly */
			nil,        /* , */
			reduce(76), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S378
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(24), /* }, reduce: StateExpr */
			nil,        /* , */
			reduce(24), /* space, reduce: StateExpr */

		},
	},
	actionRow{ // S379
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(77), /* }, reduce: CloseCurly */
			nil,        /* , */
			reduce(77), /* space, reduce: CloseCurly */

		},
	},
}

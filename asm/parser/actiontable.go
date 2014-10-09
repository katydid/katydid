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
			shift(36), /* bool_var */
			shift(37), /* int64_var */
			shift(38), /* int32_var */
			shift(39), /* uint64_var */
			shift(40), /* uint32_var */
			shift(41), /* double_var */
			shift(42), /* float_var */
			shift(43), /* string_var */
			shift(44), /* bytes_var */
			shift(45), /* true */
			shift(46), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(47), /* space */

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
			nil,          /* bool_var */
			nil,          /* int64_var */
			nil,          /* int32_var */
			nil,          /* uint64_var */
			nil,          /* uint32_var */
			nil,          /* double_var */
			nil,          /* float_var */
			nil,          /* string_var */
			nil,          /* bytes_var */
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
			shift(50), /* id */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(51), /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(52), /* root */
			shift(53), /* id */
			nil,       /* . */
			shift(54), /* if */
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
			shift(36), /* bool_var */
			shift(37), /* int64_var */
			shift(38), /* int32_var */
			shift(39), /* uint64_var */
			shift(40), /* uint32_var */
			shift(41), /* double_var */
			shift(42), /* float_var */
			shift(43), /* string_var */
			shift(44), /* bytes_var */
			shift(45), /* true */
			shift(46), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(57), /* space */

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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			shift(63), /* . */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			shift(65), /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(66), /* space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(69), /* id */
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
			shift(76), /* int64_lit */
			shift(77), /* int32_lit */
			shift(78), /* uint64_lit */
			shift(79), /* uint32_lit */
			shift(80), /* double_lit */
			shift(81), /* float_lit */
			shift(82), /* string_lit */
			shift(83), /* bytes_lit */
			shift(84), /* bool_var */
			shift(85), /* int64_var */
			shift(86), /* int32_var */
			shift(87), /* uint64_var */
			shift(88), /* uint32_var */
			shift(89), /* double_var */
			shift(90), /* float_var */
			shift(91), /* string_var */
			shift(92), /* bytes_var */
			shift(93), /* true */
			shift(94), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(95), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(62), /* $, reduce: Terminal */
			reduce(62), /* root, reduce: Terminal */
			reduce(62), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(62), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(63), /* $, reduce: Terminal */
			reduce(63), /* root, reduce: Terminal */
			reduce(63), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(63), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(64), /* $, reduce: Terminal */
			reduce(64), /* root, reduce: Terminal */
			reduce(64), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(64), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(65), /* $, reduce: Terminal */
			reduce(65), /* root, reduce: Terminal */
			reduce(65), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(65), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(66), /* $, reduce: Terminal */
			reduce(66), /* root, reduce: Terminal */
			reduce(66), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(66), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(67), /* $, reduce: Terminal */
			reduce(67), /* root, reduce: Terminal */
			reduce(67), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(67), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(68), /* $, reduce: Terminal */
			reduce(68), /* root, reduce: Terminal */
			reduce(68), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(68), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(69), /* $, reduce: Terminal */
			reduce(69), /* root, reduce: Terminal */
			reduce(69), /* id, reduce: Terminal */
			nil,        /* . */
			reduce(69), /* if, reduce: Terminal */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(70), /* $, reduce: Bool */
			reduce(70), /* root, reduce: Bool */
			reduce(70), /* id, reduce: Bool */
			nil,        /* . */
			reduce(70), /* if, reduce: Bool */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(70), /* space, reduce: Bool */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(71), /* $, reduce: Bool */
			reduce(71), /* root, reduce: Bool */
			reduce(71), /* id, reduce: Bool */
			nil,        /* . */
			reduce(71), /* if, reduce: Bool */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(71), /* space, reduce: Bool */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* root, reduce: Space */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			reduce(89), /* if, reduce: Space */
			nil,        /* { */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int64, reduce: Space */
			reduce(89), /* []int32, reduce: Space */
			reduce(89), /* []uint64, reduce: Space */
			reduce(89), /* []uint32, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []float, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int64_lit, reduce: Space */
			reduce(89), /* int32_lit, reduce: Space */
			reduce(89), /* uint64_lit, reduce: Space */
			reduce(89), /* uint32_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* float_lit, reduce: Space */
			reduce(89), /* string_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int64_var, reduce: Space */
			reduce(89), /* int32_var, reduce: Space */
			reduce(89), /* uint64_var, reduce: Space */
			reduce(89), /* uint32_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* float_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(1),  /* $, reduce: AllRules */
			shift(52),  /* root */
			shift(100), /* id */
			nil,        /* . */
			shift(54),  /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(101), /* space */

		},
	},
	actionRow{ // S49
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
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
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(63),  /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(89), /* $, reduce: Space */
			reduce(89), /* root, reduce: Space */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			reduce(89), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S52
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(106), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(66),  /* space */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(69), /* id */
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
			shift(76), /* int64_lit */
			shift(77), /* int32_lit */
			shift(78), /* uint64_lit */
			shift(79), /* uint32_lit */
			shift(80), /* double_lit */
			shift(81), /* float_lit */
			shift(82), /* string_lit */
			shift(83), /* bytes_lit */
			shift(84), /* bool_var */
			shift(85), /* int64_var */
			shift(86), /* int32_var */
			shift(87), /* uint64_var */
			shift(88), /* uint32_var */
			shift(89), /* double_var */
			shift(90), /* float_var */
			shift(91), /* string_var */
			shift(92), /* bytes_var */
			shift(93), /* true */
			shift(94), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(95), /* space */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

		},
	},
	actionRow{ // S56
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(88), /* root, reduce: Space */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			reduce(88), /* if, reduce: Space */
			nil,        /* { */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int64, reduce: Space */
			reduce(88), /* []int32, reduce: Space */
			reduce(88), /* []uint64, reduce: Space */
			reduce(88), /* []uint32, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []float, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int64_lit, reduce: Space */
			reduce(88), /* int32_lit, reduce: Space */
			reduce(88), /* uint64_lit, reduce: Space */
			reduce(88), /* uint32_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* float_lit, reduce: Space */
			reduce(88), /* string_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int64_var, reduce: Space */
			reduce(88), /* int32_var, reduce: Space */
			reduce(88), /* uint64_var, reduce: Space */
			reduce(88), /* uint32_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* float_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S58
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			shift(110), /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(111), /* space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(113), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(72), /* id, reduce: Equal */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(72), /* space, reduce: Equal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(89), /* =, reduce: Space */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(114), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(115), /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(116), /* space */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(117), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(78), /* id, reduce: OpenParen */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(78), /* []bool, reduce: OpenParen */
			reduce(78), /* []int64, reduce: OpenParen */
			reduce(78), /* []int32, reduce: OpenParen */
			reduce(78), /* []uint64, reduce: OpenParen */
			reduce(78), /* []uint32, reduce: OpenParen */
			reduce(78), /* []double, reduce: OpenParen */
			reduce(78), /* []float, reduce: OpenParen */
			reduce(78), /* []string, reduce: OpenParen */
			reduce(78), /* [][]byte, reduce: OpenParen */
			reduce(78), /* int64_lit, reduce: OpenParen */
			reduce(78), /* int32_lit, reduce: OpenParen */
			reduce(78), /* uint64_lit, reduce: OpenParen */
			reduce(78), /* uint32_lit, reduce: OpenParen */
			reduce(78), /* double_lit, reduce: OpenParen */
			reduce(78), /* float_lit, reduce: OpenParen */
			reduce(78), /* string_lit, reduce: OpenParen */
			reduce(78), /* bytes_lit, reduce: OpenParen */
			reduce(78), /* bool_var, reduce: OpenParen */
			reduce(78), /* int64_var, reduce: OpenParen */
			reduce(78), /* int32_var, reduce: OpenParen */
			reduce(78), /* uint64_var, reduce: OpenParen */
			reduce(78), /* uint32_var, reduce: OpenParen */
			reduce(78), /* double_var, reduce: OpenParen */
			reduce(78), /* float_var, reduce: OpenParen */
			reduce(78), /* string_var, reduce: OpenParen */
			reduce(78), /* bytes_var, reduce: OpenParen */
			reduce(78), /* true, reduce: OpenParen */
			reduce(78), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(78), /* ), reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			reduce(78), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(89), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(150), /* id */
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
			shift(76),  /* int64_lit */
			shift(77),  /* int32_lit */
			shift(78),  /* uint64_lit */
			shift(79),  /* uint32_lit */
			shift(80),  /* double_lit */
			shift(81),  /* float_lit */
			shift(82),  /* string_lit */
			shift(83),  /* bytes_lit */
			shift(84),  /* bool_var */
			shift(85),  /* int64_var */
			shift(86),  /* int32_var */
			shift(87),  /* uint64_var */
			shift(88),  /* uint32_var */
			shift(89),  /* double_var */
			shift(90),  /* float_var */
			shift(91),  /* string_var */
			shift(92),  /* bytes_var */
			shift(93),  /* true */
			shift(94),  /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(153), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(156), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(157), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S77
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S78
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S79
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S80
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S81
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S82
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S83
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S84
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S85
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(62), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S86
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(63), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S87
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(64), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S88
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(65), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S89
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(66), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S90
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(67), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(67), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(68), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S92
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(69), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S93
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(70), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(70), /* space, reduce: Bool */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(71), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(71), /* space, reduce: Bool */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int64, reduce: Space */
			reduce(89), /* []int32, reduce: Space */
			reduce(89), /* []uint64, reduce: Space */
			reduce(89), /* []uint32, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []float, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int64_lit, reduce: Space */
			reduce(89), /* int32_lit, reduce: Space */
			reduce(89), /* uint64_lit, reduce: Space */
			reduce(89), /* uint32_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* float_lit, reduce: Space */
			reduce(89), /* string_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int64_var, reduce: Space */
			reduce(89), /* int32_var, reduce: Space */
			reduce(89), /* uint64_var, reduce: Space */
			reduce(89), /* uint32_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* float_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			shift(162), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(163), /* space */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(82), /* id, reduce: OpenCurly */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(82), /* []bool, reduce: OpenCurly */
			reduce(82), /* []int64, reduce: OpenCurly */
			reduce(82), /* []int32, reduce: OpenCurly */
			reduce(82), /* []uint64, reduce: OpenCurly */
			reduce(82), /* []uint32, reduce: OpenCurly */
			reduce(82), /* []double, reduce: OpenCurly */
			reduce(82), /* []float, reduce: OpenCurly */
			reduce(82), /* []string, reduce: OpenCurly */
			reduce(82), /* [][]byte, reduce: OpenCurly */
			reduce(82), /* int64_lit, reduce: OpenCurly */
			reduce(82), /* int32_lit, reduce: OpenCurly */
			reduce(82), /* uint64_lit, reduce: OpenCurly */
			reduce(82), /* uint32_lit, reduce: OpenCurly */
			reduce(82), /* double_lit, reduce: OpenCurly */
			reduce(82), /* float_lit, reduce: OpenCurly */
			reduce(82), /* string_lit, reduce: OpenCurly */
			reduce(82), /* bytes_lit, reduce: OpenCurly */
			reduce(82), /* bool_var, reduce: OpenCurly */
			reduce(82), /* int64_var, reduce: OpenCurly */
			reduce(82), /* int32_var, reduce: OpenCurly */
			reduce(82), /* uint64_var, reduce: OpenCurly */
			reduce(82), /* uint32_var, reduce: OpenCurly */
			reduce(82), /* double_var, reduce: OpenCurly */
			reduce(82), /* float_var, reduce: OpenCurly */
			reduce(82), /* string_var, reduce: OpenCurly */
			reduce(82), /* bytes_var, reduce: OpenCurly */
			reduce(82), /* true, reduce: OpenCurly */
			reduce(82), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(82), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(82), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(89), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(106), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(88), /* $, reduce: Space */
			reduce(88), /* root, reduce: Space */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			reduce(88), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(114), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(199), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(200), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(115), /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(116), /* space */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(201), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S108
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(156), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(157), /* space */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(73), /* id, reduce: Equal */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(73), /* space, reduce: Equal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(88), /* =, reduce: Space */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(207), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(208), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S114
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(79), /* id, reduce: OpenParen */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(79), /* []bool, reduce: OpenParen */
			reduce(79), /* []int64, reduce: OpenParen */
			reduce(79), /* []int32, reduce: OpenParen */
			reduce(79), /* []uint64, reduce: OpenParen */
			reduce(79), /* []uint32, reduce: OpenParen */
			reduce(79), /* []double, reduce: OpenParen */
			reduce(79), /* []float, reduce: OpenParen */
			reduce(79), /* []string, reduce: OpenParen */
			reduce(79), /* [][]byte, reduce: OpenParen */
			reduce(79), /* int64_lit, reduce: OpenParen */
			reduce(79), /* int32_lit, reduce: OpenParen */
			reduce(79), /* uint64_lit, reduce: OpenParen */
			reduce(79), /* uint32_lit, reduce: OpenParen */
			reduce(79), /* double_lit, reduce: OpenParen */
			reduce(79), /* float_lit, reduce: OpenParen */
			reduce(79), /* string_lit, reduce: OpenParen */
			reduce(79), /* bytes_lit, reduce: OpenParen */
			reduce(79), /* bool_var, reduce: OpenParen */
			reduce(79), /* int64_var, reduce: OpenParen */
			reduce(79), /* int32_var, reduce: OpenParen */
			reduce(79), /* uint64_var, reduce: OpenParen */
			reduce(79), /* uint32_var, reduce: OpenParen */
			reduce(79), /* double_var, reduce: OpenParen */
			reduce(79), /* float_var, reduce: OpenParen */
			reduce(79), /* string_var, reduce: OpenParen */
			reduce(79), /* bytes_var, reduce: OpenParen */
			reduce(79), /* true, reduce: OpenParen */
			reduce(79), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(79), /* ), reduce: OpenParen */
			nil,        /* } */
			nil,        /* , */
			reduce(79), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(88), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S117
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(211), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(214), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(215), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S122
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

		},
	},
	actionRow{ // S123
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S124
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S127
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S129
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S130
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S133
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S134
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S136
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S137
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S138
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(62), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(62), /* ,, reduce: Terminal */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S139
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(63), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(63), /* ,, reduce: Terminal */
			reduce(63), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(64), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(64), /* ,, reduce: Terminal */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S141
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(65), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(65), /* ,, reduce: Terminal */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S142
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(66), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(66), /* ,, reduce: Terminal */
			reduce(66), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(67), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(67), /* ,, reduce: Terminal */
			reduce(67), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(68), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(68), /* ,, reduce: Terminal */
			reduce(68), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(69), /* ), reduce: Terminal */
			nil,        /* } */
			reduce(69), /* ,, reduce: Terminal */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S146
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(70), /* ), reduce: Bool */
			nil,        /* } */
			reduce(70), /* ,, reduce: Bool */
			reduce(70), /* space, reduce: Bool */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(71), /* ), reduce: Bool */
			nil,        /* } */
			reduce(71), /* ,, reduce: Bool */
			reduce(71), /* space, reduce: Bool */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(80), /* $, reduce: CloseParen */
			reduce(80), /* root, reduce: CloseParen */
			reduce(80), /* id, reduce: CloseParen */
			nil,        /* . */
			reduce(80), /* if, reduce: CloseParen */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int64, reduce: Space */
			reduce(89), /* []int32, reduce: Space */
			reduce(89), /* []uint64, reduce: Space */
			reduce(89), /* []uint32, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []float, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int64_lit, reduce: Space */
			reduce(89), /* int32_lit, reduce: Space */
			reduce(89), /* uint64_lit, reduce: Space */
			reduce(89), /* uint32_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* float_lit, reduce: Space */
			reduce(89), /* string_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int64_var, reduce: Space */
			reduce(89), /* int32_var, reduce: Space */
			reduce(89), /* uint64_var, reduce: Space */
			reduce(89), /* uint32_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* float_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(89), /* ), reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int64, reduce: Space */
			reduce(88), /* []int32, reduce: Space */
			reduce(88), /* []uint64, reduce: Space */
			reduce(88), /* []uint32, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []float, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int64_lit, reduce: Space */
			reduce(88), /* int32_lit, reduce: Space */
			reduce(88), /* uint64_lit, reduce: Space */
			reduce(88), /* uint32_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* float_lit, reduce: Space */
			reduce(88), /* string_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int64_var, reduce: Space */
			reduce(88), /* int32_var, reduce: Space */
			reduce(88), /* uint64_var, reduce: Space */
			reduce(88), /* uint32_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* float_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(225), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(226), /* space */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(228), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(230), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(74), /* id, reduce: Then */
			nil,        /* . */
			nil,        /* if */
			reduce(74), /* {, reduce: Then */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(74), /* space, reduce: Then */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(89), /* then, reduce: Space */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(115), /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(232), /* space */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(236), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S160
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(89), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(240), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(83), /* id, reduce: OpenCurly */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(83), /* []bool, reduce: OpenCurly */
			reduce(83), /* []int64, reduce: OpenCurly */
			reduce(83), /* []int32, reduce: OpenCurly */
			reduce(83), /* []uint64, reduce: OpenCurly */
			reduce(83), /* []uint32, reduce: OpenCurly */
			reduce(83), /* []double, reduce: OpenCurly */
			reduce(83), /* []float, reduce: OpenCurly */
			reduce(83), /* []string, reduce: OpenCurly */
			reduce(83), /* [][]byte, reduce: OpenCurly */
			reduce(83), /* int64_lit, reduce: OpenCurly */
			reduce(83), /* int32_lit, reduce: OpenCurly */
			reduce(83), /* uint64_lit, reduce: OpenCurly */
			reduce(83), /* uint32_lit, reduce: OpenCurly */
			reduce(83), /* double_lit, reduce: OpenCurly */
			reduce(83), /* float_lit, reduce: OpenCurly */
			reduce(83), /* string_lit, reduce: OpenCurly */
			reduce(83), /* bytes_lit, reduce: OpenCurly */
			reduce(83), /* bool_var, reduce: OpenCurly */
			reduce(83), /* int64_var, reduce: OpenCurly */
			reduce(83), /* int32_var, reduce: OpenCurly */
			reduce(83), /* uint64_var, reduce: OpenCurly */
			reduce(83), /* uint32_var, reduce: OpenCurly */
			reduce(83), /* double_var, reduce: OpenCurly */
			reduce(83), /* float_var, reduce: OpenCurly */
			reduce(83), /* string_var, reduce: OpenCurly */
			reduce(83), /* bytes_var, reduce: OpenCurly */
			reduce(83), /* true, reduce: OpenCurly */
			reduce(83), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(83), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(83), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(88), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(244), /* } */
			nil,        /* , */
			shift(245), /* space */

		},
	},
	actionRow{ // S165
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S166
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

		},
	},
	actionRow{ // S167
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S168
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S169
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			shift(220), /* , */
			shift(250), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S171
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

		},
	},
	actionRow{ // S172
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S173
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S175
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S176
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S177
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S178
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S179
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S180
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S181
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S183
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S184
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(62), /* }, reduce: Terminal */
			reduce(62), /* ,, reduce: Terminal */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S185
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(63), /* }, reduce: Terminal */
			reduce(63), /* ,, reduce: Terminal */
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S186
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(64), /* }, reduce: Terminal */
			reduce(64), /* ,, reduce: Terminal */
			reduce(64), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(65), /* }, reduce: Terminal */
			reduce(65), /* ,, reduce: Terminal */
			reduce(65), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S188
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(66), /* }, reduce: Terminal */
			reduce(66), /* ,, reduce: Terminal */
			reduce(66), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S189
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(67), /* }, reduce: Terminal */
			reduce(67), /* ,, reduce: Terminal */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S190
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(68), /* }, reduce: Terminal */
			reduce(68), /* ,, reduce: Terminal */
			reduce(68), /* space, reduce: Terminal */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(69), /* }, reduce: Terminal */
			reduce(69), /* ,, reduce: Terminal */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S192
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(70), /* }, reduce: Bool */
			reduce(70), /* ,, reduce: Bool */
			reduce(70), /* space, reduce: Bool */

		},
	},
	actionRow{ // S193
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(71), /* }, reduce: Bool */
			reduce(71), /* ,, reduce: Bool */
			reduce(71), /* space, reduce: Bool */

		},
	},
	actionRow{ // S194
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(84), /* $, reduce: CloseCurly */
			reduce(84), /* root, reduce: CloseCurly */
			reduce(84), /* id, reduce: CloseCurly */
			nil,        /* . */
			reduce(84), /* if, reduce: CloseCurly */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S195
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int64, reduce: Space */
			reduce(89), /* []int32, reduce: Space */
			reduce(89), /* []uint64, reduce: Space */
			reduce(89), /* []uint32, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []float, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int64_lit, reduce: Space */
			reduce(89), /* int32_lit, reduce: Space */
			reduce(89), /* uint64_lit, reduce: Space */
			reduce(89), /* uint32_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* float_lit, reduce: Space */
			reduce(89), /* string_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int64_var, reduce: Space */
			reduce(89), /* int32_var, reduce: Space */
			reduce(89), /* uint64_var, reduce: Space */
			reduce(89), /* uint32_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* float_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(89), /* }, reduce: Space */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S196
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(200), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S197
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S198
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(252), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S199
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(253), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S200
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S201
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			shift(60), /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

		},
	},
	actionRow{ // S203
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S204
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(228), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(230), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S205
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			shift(220), /* , */
			shift(250), /* space */

		},
	},
	actionRow{ // S207
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(259), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S208
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(260), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S209
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(262), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S210
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(264), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

		},
	},
	actionRow{ // S212
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

		},
	},
	actionRow{ // S213
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S214
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(81), /* $, reduce: CloseParen */
			reduce(81), /* root, reduce: CloseParen */
			reduce(81), /* id, reduce: CloseParen */
			nil,        /* . */
			reduce(81), /* if, reduce: CloseParen */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S215
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int64, reduce: Space */
			reduce(88), /* []int32, reduce: Space */
			reduce(88), /* []uint64, reduce: Space */
			reduce(88), /* []uint32, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []float, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int64_lit, reduce: Space */
			reduce(88), /* int32_lit, reduce: Space */
			reduce(88), /* uint64_lit, reduce: Space */
			reduce(88), /* uint32_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* float_lit, reduce: Space */
			reduce(88), /* string_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int64_var, reduce: Space */
			reduce(88), /* int32_var, reduce: Space */
			reduce(88), /* uint64_var, reduce: Space */
			reduce(88), /* uint32_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* float_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(88), /* ), reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S216
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S217
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(214), /* ) */
			nil,        /* } */
			shift(271), /* , */
			shift(272), /* space */

		},
	},
	actionRow{ // S218
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S219
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(95),  /* space */

		},
	},
	actionRow{ // S220
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(86), /* id, reduce: Comma */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(86), /* []bool, reduce: Comma */
			reduce(86), /* []int64, reduce: Comma */
			reduce(86), /* []int32, reduce: Comma */
			reduce(86), /* []uint64, reduce: Comma */
			reduce(86), /* []uint32, reduce: Comma */
			reduce(86), /* []double, reduce: Comma */
			reduce(86), /* []float, reduce: Comma */
			reduce(86), /* []string, reduce: Comma */
			reduce(86), /* [][]byte, reduce: Comma */
			reduce(86), /* int64_lit, reduce: Comma */
			reduce(86), /* int32_lit, reduce: Comma */
			reduce(86), /* uint64_lit, reduce: Comma */
			reduce(86), /* uint32_lit, reduce: Comma */
			reduce(86), /* double_lit, reduce: Comma */
			reduce(86), /* float_lit, reduce: Comma */
			reduce(86), /* string_lit, reduce: Comma */
			reduce(86), /* bytes_lit, reduce: Comma */
			reduce(86), /* bool_var, reduce: Comma */
			reduce(86), /* int64_var, reduce: Comma */
			reduce(86), /* int32_var, reduce: Comma */
			reduce(86), /* uint64_var, reduce: Comma */
			reduce(86), /* uint32_var, reduce: Comma */
			reduce(86), /* double_var, reduce: Comma */
			reduce(86), /* float_var, reduce: Comma */
			reduce(86), /* string_var, reduce: Comma */
			reduce(86), /* bytes_var, reduce: Comma */
			reduce(86), /* true, reduce: Comma */
			reduce(86), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(86), /* space, reduce: Comma */

		},
	},
	actionRow{ // S221
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(89), /* ), reduce: Space */
			nil,        /* } */
			reduce(89), /* ,, reduce: Space */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S222
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(278), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S223
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(236), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S224
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(240), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S225
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(75), /* id, reduce: Then */
			nil,        /* . */
			nil,        /* if */
			reduce(75), /* {, reduce: Then */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(75), /* space, reduce: Then */

		},
	},
	actionRow{ // S226
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(88), /* then, reduce: Space */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S227
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(283), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(284), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(285), /* space */

		},
	},
	actionRow{ // S228
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S229
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(288), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(289), /* space */

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
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

		},
	},
	actionRow{ // S231
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(89), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			reduce(89), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S232
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			reduce(88), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S233
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(211), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(294), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(215), /* space */

		},
	},
	actionRow{ // S234
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(236), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

		},
	},
	actionRow{ // S235
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S236
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(80), /* then, reduce: CloseParen */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(80), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S237
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(297), /* } */
			nil,        /* , */
			shift(245), /* space */

		},
	},
	actionRow{ // S238
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S239
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(240), /* } */
			shift(220), /* , */
			shift(250), /* space */

		},
	},
	actionRow{ // S240
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(84), /* then, reduce: CloseCurly */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S241
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			shift(65),  /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(160), /* space */

		},
	},
	actionRow{ // S242
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(97), /* { */
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
			nil,       /* bool_var */
			nil,       /* int64_var */
			nil,       /* int32_var */
			nil,       /* uint64_var */
			nil,       /* uint32_var */
			nil,       /* double_var */
			nil,       /* float_var */
			nil,       /* string_var */
			nil,       /* bytes_var */
			nil,       /* true */
			nil,       /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(99), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S244
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(85), /* $, reduce: CloseCurly */
			reduce(85), /* root, reduce: CloseCurly */
			reduce(85), /* id, reduce: CloseCurly */
			nil,        /* . */
			reduce(85), /* if, reduce: CloseCurly */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(85), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S245
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int64, reduce: Space */
			reduce(88), /* []int32, reduce: Space */
			reduce(88), /* []uint64, reduce: Space */
			reduce(88), /* []uint32, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []float, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int64_lit, reduce: Space */
			reduce(88), /* int32_lit, reduce: Space */
			reduce(88), /* uint64_lit, reduce: Space */
			reduce(88), /* uint32_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* float_lit, reduce: Space */
			reduce(88), /* string_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int64_var, reduce: Space */
			reduce(88), /* int32_var, reduce: Space */
			reduce(88), /* uint64_var, reduce: Space */
			reduce(88), /* uint32_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* float_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(88), /* }, reduce: Space */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S246
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(244), /* } */
			shift(271), /* , */
			shift(306), /* space */

		},
	},
	actionRow{ // S248
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S249
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(95),  /* space */

		},
	},
	actionRow{ // S250
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(89), /* }, reduce: Space */
			reduce(89), /* ,, reduce: Space */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S251
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(312), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S252
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			shift(313), /* . */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S253
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(314), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S254
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(316), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S255
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(318), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(103), /* space */

		},
	},
	actionRow{ // S256
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(288), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(289), /* space */

		},
	},
	actionRow{ // S258
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S259
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(320), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S260
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S261
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(321), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S262
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S263
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(322), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S264
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S265
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S266
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(278), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S267
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(211), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(327), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(215), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

		},
	},
	actionRow{ // S269
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(80), /* ), reduce: CloseParen */
			nil,        /* } */
			reduce(80), /* ,, reduce: CloseParen */
			reduce(80), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S271
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(87), /* id, reduce: Comma */
			nil,        /* . */
			nil,        /* if */
			nil,        /* { */
			reduce(87), /* []bool, reduce: Comma */
			reduce(87), /* []int64, reduce: Comma */
			reduce(87), /* []int32, reduce: Comma */
			reduce(87), /* []uint64, reduce: Comma */
			reduce(87), /* []uint32, reduce: Comma */
			reduce(87), /* []double, reduce: Comma */
			reduce(87), /* []float, reduce: Comma */
			reduce(87), /* []string, reduce: Comma */
			reduce(87), /* [][]byte, reduce: Comma */
			reduce(87), /* int64_lit, reduce: Comma */
			reduce(87), /* int32_lit, reduce: Comma */
			reduce(87), /* uint64_lit, reduce: Comma */
			reduce(87), /* uint32_lit, reduce: Comma */
			reduce(87), /* double_lit, reduce: Comma */
			reduce(87), /* float_lit, reduce: Comma */
			reduce(87), /* string_lit, reduce: Comma */
			reduce(87), /* bytes_lit, reduce: Comma */
			reduce(87), /* bool_var, reduce: Comma */
			reduce(87), /* int64_var, reduce: Comma */
			reduce(87), /* int32_var, reduce: Comma */
			reduce(87), /* uint64_var, reduce: Comma */
			reduce(87), /* uint32_var, reduce: Comma */
			reduce(87), /* double_var, reduce: Comma */
			reduce(87), /* float_var, reduce: Comma */
			reduce(87), /* string_var, reduce: Comma */
			reduce(87), /* bytes_var, reduce: Comma */
			reduce(87), /* true, reduce: Comma */
			reduce(87), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(87), /* space, reduce: Comma */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(88), /* ), reduce: Space */
			nil,        /* } */
			reduce(88), /* ,, reduce: Space */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S273
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(211), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(153), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S275
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(330), /* } */
			nil,        /* , */
			shift(245), /* space */

		},
	},
	actionRow{ // S276
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S277
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(278), /* } */
			shift(220), /* , */
			shift(250), /* space */

		},
	},
	actionRow{ // S278
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(84), /* ), reduce: CloseCurly */
			nil,        /* } */
			reduce(84), /* ,, reduce: CloseCurly */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S279
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(236), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(240), /* } */
			shift(220), /* , */
			shift(250), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S284
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

		},
	},
	actionRow{ // S285
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(88), /* id, reduce: Space */
			nil,        /* . */
			nil,        /* if */
			reduce(88), /* {, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S286
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(336), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(337), /* space */

		},
	},
	actionRow{ // S287
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(339), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(341), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S288
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(76), /* id, reduce: Else */
			nil,        /* . */
			nil,        /* if */
			reduce(76), /* {, reduce: Else */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(76), /* space, reduce: Else */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(89), /* else, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S290
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(342), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(343), /* space */

		},
	},
	actionRow{ // S291
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(346), /* } */
			nil,        /* , */
			shift(347), /* space */

		},
	},
	actionRow{ // S292
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(69), /* id */
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
			shift(76), /* int64_lit */
			shift(77), /* int32_lit */
			shift(78), /* uint64_lit */
			shift(79), /* uint32_lit */
			shift(80), /* double_lit */
			shift(81), /* float_lit */
			shift(82), /* string_lit */
			shift(83), /* bytes_lit */
			shift(84), /* bool_var */
			shift(85), /* int64_var */
			shift(86), /* int32_var */
			shift(87), /* uint64_var */
			shift(88), /* uint32_var */
			shift(89), /* double_var */
			shift(90), /* float_var */
			shift(91), /* string_var */
			shift(92), /* bytes_var */
			shift(93), /* true */
			shift(94), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(95), /* space */

		},
	},
	actionRow{ // S293
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			reduce(89), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S294
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(81), /* then, reduce: CloseParen */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(81), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S295
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(294), /* ) */
			nil,        /* } */
			shift(271), /* , */
			shift(272), /* space */

		},
	},
	actionRow{ // S296
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S297
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(85), /* then, reduce: CloseCurly */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(85), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S298
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(297), /* } */
			shift(271), /* , */
			shift(306), /* space */

		},
	},
	actionRow{ // S299
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S300
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(120), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(149), /* space */

		},
	},
	actionRow{ // S301
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(166), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(312), /* } */
			nil,        /* , */
			shift(195), /* space */

		},
	},
	actionRow{ // S302
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(211), /* id */
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
			shift(129), /* int64_lit */
			shift(130), /* int32_lit */
			shift(131), /* uint64_lit */
			shift(132), /* uint32_lit */
			shift(133), /* double_lit */
			shift(134), /* float_lit */
			shift(135), /* string_lit */
			shift(136), /* bytes_lit */
			shift(137), /* bool_var */
			shift(138), /* int64_var */
			shift(139), /* int32_var */
			shift(140), /* uint64_var */
			shift(141), /* uint32_var */
			shift(142), /* double_var */
			shift(143), /* float_var */
			shift(144), /* string_var */
			shift(145), /* bytes_var */
			shift(146), /* true */
			shift(147), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(353), /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(215), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

		},
	},
	actionRow{ // S304
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(80), /* }, reduce: CloseParen */
			reduce(80), /* ,, reduce: CloseParen */
			reduce(80), /* space, reduce: CloseParen */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(88), /* }, reduce: Space */
			reduce(88), /* ,, reduce: Space */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S307
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(153), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S309
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(241), /* id */
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
			shift(175), /* int64_lit */
			shift(176), /* int32_lit */
			shift(177), /* uint64_lit */
			shift(178), /* uint32_lit */
			shift(179), /* double_lit */
			shift(180), /* float_lit */
			shift(181), /* string_lit */
			shift(182), /* bytes_lit */
			shift(183), /* bool_var */
			shift(184), /* int64_var */
			shift(185), /* int32_var */
			shift(186), /* uint64_var */
			shift(187), /* uint32_var */
			shift(188), /* double_var */
			shift(189), /* float_var */
			shift(190), /* string_var */
			shift(191), /* bytes_var */
			shift(192), /* true */
			shift(193), /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(356), /* } */
			nil,        /* , */
			shift(245), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(312), /* } */
			shift(220), /* , */
			shift(250), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(84), /* }, reduce: CloseCurly */
			reduce(84), /* ,, reduce: CloseCurly */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S313
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(359), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S314
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S315
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(360), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S316
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S317
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(361), /* id */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(197), /* space */

		},
	},
	actionRow{ // S318
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S319
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(339), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(341), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S320
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S321
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S322
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(270), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(278), /* } */
			shift(220), /* , */
			shift(250), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(81), /* ), reduce: CloseParen */
			nil,        /* } */
			reduce(81), /* ,, reduce: CloseParen */
			reduce(81), /* space, reduce: CloseParen */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(327), /* ) */
			nil,        /* } */
			shift(271), /* , */
			shift(272), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			reduce(85), /* ), reduce: CloseCurly */
			nil,        /* } */
			reduce(85), /* ,, reduce: CloseCurly */
			reduce(85), /* space, reduce: CloseCurly */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(330), /* } */
			shift(271), /* , */
			shift(306), /* space */

		},
	},
	actionRow{ // S332
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(346), /* } */
			nil,        /* , */
			shift(347), /* space */

		},
	},
	actionRow{ // S336
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			reduce(77), /* id, reduce: Else */
			nil,        /* . */
			nil,        /* if */
			reduce(77), /* {, reduce: Else */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(77), /* space, reduce: Else */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(88), /* else, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S338
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(366), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(367), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(285), /* space */

		},
	},
	actionRow{ // S339
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S340
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S341
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

		},
	},
	actionRow{ // S342
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(69), /* id */
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
			shift(76), /* int64_lit */
			shift(77), /* int32_lit */
			shift(78), /* uint64_lit */
			shift(79), /* uint32_lit */
			shift(80), /* double_lit */
			shift(81), /* float_lit */
			shift(82), /* string_lit */
			shift(83), /* bytes_lit */
			shift(84), /* bool_var */
			shift(85), /* int64_var */
			shift(86), /* int32_var */
			shift(87), /* uint64_var */
			shift(88), /* uint32_var */
			shift(89), /* double_var */
			shift(90), /* float_var */
			shift(91), /* string_var */
			shift(92), /* bytes_var */
			shift(93), /* true */
			shift(94), /* false */
			nil,       /* = */
			nil,       /* then */
			nil,       /* else */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* } */
			nil,       /* , */
			shift(95), /* space */

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
			reduce(88), /* if, reduce: Space */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(370), /* } */
			nil,        /* , */
			shift(371), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(84), /* else, reduce: CloseCurly */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S347
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(89), /* }, reduce: Space */
			nil,        /* , */
			reduce(89), /* space, reduce: Space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(156), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(157), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(305), /* ) */
			nil,        /* } */
			shift(220), /* , */
			shift(221), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(312), /* } */
			shift(220), /* , */
			shift(250), /* space */

		},
	},
	actionRow{ // S353
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(81), /* }, reduce: CloseParen */
			reduce(81), /* ,, reduce: CloseParen */
			reduce(81), /* space, reduce: CloseParen */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			shift(353), /* ) */
			nil,        /* } */
			shift(271), /* , */
			shift(272), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(85), /* }, reduce: CloseCurly */
			reduce(85), /* ,, reduce: CloseCurly */
			reduce(85), /* space, reduce: CloseCurly */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(356), /* } */
			shift(271), /* , */
			shift(306), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S359
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S360
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S361
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S362
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S364
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S365
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S366
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S367
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			nil,        /* , */
			shift(347), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(156), /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(157), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			reduce(85), /* else, reduce: CloseCurly */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			reduce(85), /* space, reduce: CloseCurly */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(88), /* }, reduce: Space */
			nil,        /* , */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S372
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(228), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(230), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(194), /* } */
			nil,        /* , */
			shift(347), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(244), /* } */
			nil,        /* , */
			shift(371), /* space */

		},
	},
	actionRow{ // S377
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S378
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(228), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(230), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(288), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(289), /* space */

		},
	},
	actionRow{ // S380
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S381
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			shift(288), /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(289), /* space */

		},
	},
	actionRow{ // S382
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(385), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(387), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S383
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(385), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(387), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(231), /* space */

		},
	},
	actionRow{ // S384
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			shift(389), /* id */
			nil,        /* . */
			nil,        /* if */
			shift(390), /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(285), /* space */

		},
	},
	actionRow{ // S385
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S386
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S387
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

		},
	},
	actionRow{ // S388
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S389
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S390
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* id */
			nil,        /* . */
			shift(292), /* if */
			nil,        /* { */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* } */
			nil,        /* , */
			shift(293), /* space */

		},
	},
	actionRow{ // S391
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(395), /* } */
			nil,        /* , */
			shift(347), /* space */

		},
	},
	actionRow{ // S392
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(395), /* } */
			nil,        /* , */
			shift(347), /* space */

		},
	},
	actionRow{ // S393
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			shift(397), /* } */
			nil,        /* , */
			shift(371), /* space */

		},
	},
	actionRow{ // S394
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S395
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(84), /* }, reduce: CloseCurly */
			nil,        /* , */
			reduce(84), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S396
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
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
	actionRow{ // S397
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
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* then */
			nil,        /* else */
			nil,        /* ( */
			nil,        /* ) */
			reduce(85), /* }, reduce: CloseCurly */
			nil,        /* , */
			reduce(85), /* space, reduce: CloseCurly */

		},
	},
}

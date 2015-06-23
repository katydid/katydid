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
			nil,      /* []int64 */
			nil,      /* []int32 */
			nil,      /* []uint64 */
			nil,      /* []uint32 */
			nil,      /* []double */
			nil,      /* []float */
			nil,      /* []string */
			nil,      /* [][]byte */
			nil,      /* int64_lit */
			nil,      /* int32_lit */
			nil,      /* uint64_lit */
			nil,      /* uint32_lit */
			nil,      /* double_lit */
			nil,      /* float_lit */
			nil,      /* string_lit */
			nil,      /* bytes_lit */
			nil,      /* bool_var */
			nil,      /* int64_var */
			nil,      /* int32_var */
			nil,      /* uint64_var */
			nil,      /* uint32_var */
			nil,      /* double_var */
			nil,      /* float_var */
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
			nil,      /* []int64 */
			nil,      /* []int32 */
			nil,      /* []uint64 */
			nil,      /* []uint32 */
			nil,      /* []double */
			nil,      /* []float */
			nil,      /* []string */
			nil,      /* [][]byte */
			nil,      /* int64_lit */
			nil,      /* int32_lit */
			nil,      /* uint64_lit */
			nil,      /* uint32_lit */
			nil,      /* double_lit */
			nil,      /* float_lit */
			nil,      /* string_lit */
			nil,      /* bytes_lit */
			nil,      /* bool_var */
			nil,      /* int64_var */
			nil,      /* int32_var */
			nil,      /* uint64_var */
			nil,      /* uint32_var */
			nil,      /* double_var */
			nil,      /* float_var */
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
			reduce(58), /* id, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			reduce(57), /* id, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(53), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(54), /* space */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(47), /* id, reduce: OpenParen */
			reduce(47), /* []bool, reduce: OpenParen */
			reduce(47), /* []int64, reduce: OpenParen */
			reduce(47), /* []int32, reduce: OpenParen */
			reduce(47), /* []uint64, reduce: OpenParen */
			reduce(47), /* []uint32, reduce: OpenParen */
			reduce(47), /* []double, reduce: OpenParen */
			reduce(47), /* []float, reduce: OpenParen */
			reduce(47), /* []string, reduce: OpenParen */
			reduce(47), /* [][]byte, reduce: OpenParen */
			reduce(47), /* int64_lit, reduce: OpenParen */
			reduce(47), /* int32_lit, reduce: OpenParen */
			reduce(47), /* uint64_lit, reduce: OpenParen */
			reduce(47), /* uint32_lit, reduce: OpenParen */
			reduce(47), /* double_lit, reduce: OpenParen */
			reduce(47), /* float_lit, reduce: OpenParen */
			reduce(47), /* string_lit, reduce: OpenParen */
			reduce(47), /* bytes_lit, reduce: OpenParen */
			reduce(47), /* bool_var, reduce: OpenParen */
			reduce(47), /* int64_var, reduce: OpenParen */
			reduce(47), /* int32_var, reduce: OpenParen */
			reduce(47), /* uint64_var, reduce: OpenParen */
			reduce(47), /* uint32_var, reduce: OpenParen */
			reduce(47), /* double_var, reduce: OpenParen */
			reduce(47), /* float_var, reduce: OpenParen */
			reduce(47), /* string_var, reduce: OpenParen */
			reduce(47), /* bytes_var, reduce: OpenParen */
			reduce(47), /* true, reduce: OpenParen */
			reduce(47), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(47), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(47), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(58), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(53), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(54), /* space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(48), /* id, reduce: OpenParen */
			reduce(48), /* []bool, reduce: OpenParen */
			reduce(48), /* []int64, reduce: OpenParen */
			reduce(48), /* []int32, reduce: OpenParen */
			reduce(48), /* []uint64, reduce: OpenParen */
			reduce(48), /* []uint32, reduce: OpenParen */
			reduce(48), /* []double, reduce: OpenParen */
			reduce(48), /* []float, reduce: OpenParen */
			reduce(48), /* []string, reduce: OpenParen */
			reduce(48), /* [][]byte, reduce: OpenParen */
			reduce(48), /* int64_lit, reduce: OpenParen */
			reduce(48), /* int32_lit, reduce: OpenParen */
			reduce(48), /* uint64_lit, reduce: OpenParen */
			reduce(48), /* uint32_lit, reduce: OpenParen */
			reduce(48), /* double_lit, reduce: OpenParen */
			reduce(48), /* float_lit, reduce: OpenParen */
			reduce(48), /* string_lit, reduce: OpenParen */
			reduce(48), /* bytes_lit, reduce: OpenParen */
			reduce(48), /* bool_var, reduce: OpenParen */
			reduce(48), /* int64_var, reduce: OpenParen */
			reduce(48), /* int32_var, reduce: OpenParen */
			reduce(48), /* uint64_var, reduce: OpenParen */
			reduce(48), /* uint32_var, reduce: OpenParen */
			reduce(48), /* double_var, reduce: OpenParen */
			reduce(48), /* float_var, reduce: OpenParen */
			reduce(48), /* string_var, reduce: OpenParen */
			reduce(48), /* bytes_var, reduce: OpenParen */
			reduce(48), /* true, reduce: OpenParen */
			reduce(48), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(48), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(48), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(57), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			shift(57), /* id */
			shift(23), /* []bool */
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(60), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(61), /* space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			shift(53), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(66), /* , */
			shift(67), /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Function */
			nil,       /* id */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,       /* ( */
			nil,       /* ) */
			shift(70), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(71), /* space */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(20), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(20), /* space, reduce: ListType */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(21), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(21), /* space, reduce: ListType */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(22), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			reduce(22), /* space, reduce: ListType */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(23), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(23), /* ,, reduce: SpaceTerminal */
			reduce(23), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(34), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(34), /* ,, reduce: Terminal */
			reduce(34), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(35), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(35), /* ,, reduce: Terminal */
			reduce(35), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(36), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(36), /* ,, reduce: Terminal */
			reduce(36), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(37), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(37), /* ,, reduce: Terminal */
			reduce(37), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(38), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(38), /* ,, reduce: Terminal */
			reduce(38), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(39), /* ,, reduce: Terminal */
			reduce(39), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: Terminal */
			reduce(40), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(41), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(41), /* ,, reduce: Terminal */
			reduce(41), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(42), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(42), /* ,, reduce: Terminal */
			reduce(42), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(43), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(43), /* ,, reduce: Bool */
			reduce(43), /* space, reduce: Bool */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(44), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(44), /* ,, reduce: Bool */
			reduce(44), /* space, reduce: Bool */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(49), /* $, reduce: CloseParen */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(58), /* id, reduce: Space */
			reduce(58), /* []bool, reduce: Space */
			reduce(58), /* []int64, reduce: Space */
			reduce(58), /* []int32, reduce: Space */
			reduce(58), /* []uint64, reduce: Space */
			reduce(58), /* []uint32, reduce: Space */
			reduce(58), /* []double, reduce: Space */
			reduce(58), /* []float, reduce: Space */
			reduce(58), /* []string, reduce: Space */
			reduce(58), /* [][]byte, reduce: Space */
			reduce(58), /* int64_lit, reduce: Space */
			reduce(58), /* int32_lit, reduce: Space */
			reduce(58), /* uint64_lit, reduce: Space */
			reduce(58), /* uint32_lit, reduce: Space */
			reduce(58), /* double_lit, reduce: Space */
			reduce(58), /* float_lit, reduce: Space */
			reduce(58), /* string_lit, reduce: Space */
			reduce(58), /* bytes_lit, reduce: Space */
			reduce(58), /* bool_var, reduce: Space */
			reduce(58), /* int64_var, reduce: Space */
			reduce(58), /* int32_var, reduce: Space */
			reduce(58), /* uint64_var, reduce: Space */
			reduce(58), /* uint32_var, reduce: Space */
			reduce(58), /* double_var, reduce: Space */
			reduce(58), /* float_var, reduce: Space */
			reduce(58), /* string_var, reduce: Space */
			reduce(58), /* bytes_var, reduce: Space */
			reduce(58), /* true, reduce: Space */
			reduce(58), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(58), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			shift(53), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(66), /* , */
			shift(67), /* space */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Function */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			shift(70), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(71), /* space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(24), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(24), /* ,, reduce: SpaceTerminal */
			reduce(24), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(50), /* $, reduce: CloseParen */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(57), /* id, reduce: Space */
			reduce(57), /* []bool, reduce: Space */
			reduce(57), /* []int64, reduce: Space */
			reduce(57), /* []int32, reduce: Space */
			reduce(57), /* []uint64, reduce: Space */
			reduce(57), /* []uint32, reduce: Space */
			reduce(57), /* []double, reduce: Space */
			reduce(57), /* []float, reduce: Space */
			reduce(57), /* []string, reduce: Space */
			reduce(57), /* [][]byte, reduce: Space */
			reduce(57), /* int64_lit, reduce: Space */
			reduce(57), /* int32_lit, reduce: Space */
			reduce(57), /* uint64_lit, reduce: Space */
			reduce(57), /* uint32_lit, reduce: Space */
			reduce(57), /* double_lit, reduce: Space */
			reduce(57), /* float_lit, reduce: Space */
			reduce(57), /* string_lit, reduce: Space */
			reduce(57), /* bytes_lit, reduce: Space */
			reduce(57), /* bool_var, reduce: Space */
			reduce(57), /* int64_var, reduce: Space */
			reduce(57), /* int32_var, reduce: Space */
			reduce(57), /* uint64_var, reduce: Space */
			reduce(57), /* uint32_var, reduce: Space */
			reduce(57), /* double_var, reduce: Space */
			reduce(57), /* float_var, reduce: Space */
			reduce(57), /* string_var, reduce: Space */
			reduce(57), /* bytes_var, reduce: Space */
			reduce(57), /* true, reduce: Space */
			reduce(57), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(57), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(78), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(54), /* space */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			shift(60), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(79), /* , */
			shift(80), /* space */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Function */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(83), /* space */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(55), /* id, reduce: Comma */
			reduce(55), /* []bool, reduce: Comma */
			reduce(55), /* []int64, reduce: Comma */
			reduce(55), /* []int32, reduce: Comma */
			reduce(55), /* []uint64, reduce: Comma */
			reduce(55), /* []uint32, reduce: Comma */
			reduce(55), /* []double, reduce: Comma */
			reduce(55), /* []float, reduce: Comma */
			reduce(55), /* []string, reduce: Comma */
			reduce(55), /* [][]byte, reduce: Comma */
			reduce(55), /* int64_lit, reduce: Comma */
			reduce(55), /* int32_lit, reduce: Comma */
			reduce(55), /* uint64_lit, reduce: Comma */
			reduce(55), /* uint32_lit, reduce: Comma */
			reduce(55), /* double_lit, reduce: Comma */
			reduce(55), /* float_lit, reduce: Comma */
			reduce(55), /* string_lit, reduce: Comma */
			reduce(55), /* bytes_lit, reduce: Comma */
			reduce(55), /* bool_var, reduce: Comma */
			reduce(55), /* int64_var, reduce: Comma */
			reduce(55), /* int32_var, reduce: Comma */
			reduce(55), /* uint64_var, reduce: Comma */
			reduce(55), /* uint32_var, reduce: Comma */
			reduce(55), /* double_var, reduce: Comma */
			reduce(55), /* float_var, reduce: Comma */
			reduce(55), /* string_var, reduce: Comma */
			reduce(55), /* bytes_var, reduce: Comma */
			reduce(55), /* true, reduce: Comma */
			reduce(55), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(55), /* space, reduce: Comma */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(58), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(58), /* ,, reduce: Space */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			shift(84), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(85), /* space */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(116), /* } */
			nil,        /* , */
			shift(117), /* space */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(51), /* id, reduce: OpenCurly */
			reduce(51), /* []bool, reduce: OpenCurly */
			reduce(51), /* []int64, reduce: OpenCurly */
			reduce(51), /* []int32, reduce: OpenCurly */
			reduce(51), /* []uint64, reduce: OpenCurly */
			reduce(51), /* []uint32, reduce: OpenCurly */
			reduce(51), /* []double, reduce: OpenCurly */
			reduce(51), /* []float, reduce: OpenCurly */
			reduce(51), /* []string, reduce: OpenCurly */
			reduce(51), /* [][]byte, reduce: OpenCurly */
			reduce(51), /* int64_lit, reduce: OpenCurly */
			reduce(51), /* int32_lit, reduce: OpenCurly */
			reduce(51), /* uint64_lit, reduce: OpenCurly */
			reduce(51), /* uint32_lit, reduce: OpenCurly */
			reduce(51), /* double_lit, reduce: OpenCurly */
			reduce(51), /* float_lit, reduce: OpenCurly */
			reduce(51), /* string_lit, reduce: OpenCurly */
			reduce(51), /* bytes_lit, reduce: OpenCurly */
			reduce(51), /* bool_var, reduce: OpenCurly */
			reduce(51), /* int64_var, reduce: OpenCurly */
			reduce(51), /* int32_var, reduce: OpenCurly */
			reduce(51), /* uint64_var, reduce: OpenCurly */
			reduce(51), /* uint32_var, reduce: OpenCurly */
			reduce(51), /* double_var, reduce: OpenCurly */
			reduce(51), /* float_var, reduce: OpenCurly */
			reduce(51), /* string_var, reduce: OpenCurly */
			reduce(51), /* bytes_var, reduce: OpenCurly */
			reduce(51), /* true, reduce: OpenCurly */
			reduce(51), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(51), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(51), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(58), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Function */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* space */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* id */
			shift(23), /* []bool */
			shift(24), /* []int64 */
			shift(25), /* []int32 */
			shift(26), /* []uint64 */
			shift(27), /* []uint32 */
			shift(28), /* []double */
			shift(29), /* []float */
			shift(30), /* []string */
			shift(31), /* [][]byte */
			shift(34), /* int64_lit */
			shift(35), /* int32_lit */
			shift(36), /* uint64_lit */
			shift(37), /* uint32_lit */
			shift(38), /* double_lit */
			shift(39), /* float_lit */
			shift(40), /* string_lit */
			shift(41), /* bytes_lit */
			shift(42), /* bool_var */
			shift(43), /* int64_var */
			shift(44), /* int32_var */
			shift(45), /* uint64_var */
			shift(46), /* uint32_var */
			shift(47), /* double_var */
			shift(48), /* float_var */
			shift(49), /* string_var */
			shift(50), /* bytes_var */
			shift(51), /* true */
			shift(52), /* false */
			nil,       /* = */
			nil,       /* ( */
			shift(78), /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(54), /* space */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(116), /* } */
			nil,        /* , */
			shift(117), /* space */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(57),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(34),  /* int64_lit */
			shift(35),  /* int32_lit */
			shift(36),  /* uint64_lit */
			shift(37),  /* uint32_lit */
			shift(38),  /* double_lit */
			shift(39),  /* float_lit */
			shift(40),  /* string_lit */
			shift(41),  /* bytes_lit */
			shift(42),  /* bool_var */
			shift(43),  /* int64_var */
			shift(44),  /* int32_var */
			shift(45),  /* uint64_var */
			shift(46),  /* uint32_var */
			shift(47),  /* double_var */
			shift(48),  /* float_var */
			shift(49),  /* string_var */
			shift(50),  /* bytes_var */
			shift(51),  /* true */
			shift(52),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(61),  /* space */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			shift(78), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(66), /* , */
			shift(67), /* space */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(4), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(4), /* ,, reduce: Function */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(49), /* ,, reduce: CloseParen */
			reduce(49), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(56), /* id, reduce: Comma */
			reduce(56), /* []bool, reduce: Comma */
			reduce(56), /* []int64, reduce: Comma */
			reduce(56), /* []int32, reduce: Comma */
			reduce(56), /* []uint64, reduce: Comma */
			reduce(56), /* []uint32, reduce: Comma */
			reduce(56), /* []double, reduce: Comma */
			reduce(56), /* []float, reduce: Comma */
			reduce(56), /* []string, reduce: Comma */
			reduce(56), /* [][]byte, reduce: Comma */
			reduce(56), /* int64_lit, reduce: Comma */
			reduce(56), /* int32_lit, reduce: Comma */
			reduce(56), /* uint64_lit, reduce: Comma */
			reduce(56), /* uint32_lit, reduce: Comma */
			reduce(56), /* double_lit, reduce: Comma */
			reduce(56), /* float_lit, reduce: Comma */
			reduce(56), /* string_lit, reduce: Comma */
			reduce(56), /* bytes_lit, reduce: Comma */
			reduce(56), /* bool_var, reduce: Comma */
			reduce(56), /* int64_var, reduce: Comma */
			reduce(56), /* int32_var, reduce: Comma */
			reduce(56), /* uint64_var, reduce: Comma */
			reduce(56), /* uint32_var, reduce: Comma */
			reduce(56), /* double_var, reduce: Comma */
			reduce(56), /* float_var, reduce: Comma */
			reduce(56), /* string_var, reduce: Comma */
			reduce(56), /* bytes_var, reduce: Comma */
			reduce(56), /* true, reduce: Comma */
			reduce(56), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(56), /* space, reduce: Comma */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(57), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(57), /* ,, reduce: Space */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(57),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(34),  /* int64_lit */
			shift(35),  /* int32_lit */
			shift(36),  /* uint64_lit */
			shift(37),  /* uint32_lit */
			shift(38),  /* double_lit */
			shift(39),  /* float_lit */
			shift(40),  /* string_lit */
			shift(41),  /* bytes_lit */
			shift(42),  /* bool_var */
			shift(43),  /* int64_var */
			shift(44),  /* int32_var */
			shift(45),  /* uint64_var */
			shift(46),  /* uint32_var */
			shift(47),  /* double_var */
			shift(48),  /* float_var */
			shift(49),  /* string_var */
			shift(50),  /* bytes_var */
			shift(51),  /* true */
			shift(52),  /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(125), /* space */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(58), /* id, reduce: Space */
			reduce(58), /* []bool, reduce: Space */
			reduce(58), /* []int64, reduce: Space */
			reduce(58), /* []int32, reduce: Space */
			reduce(58), /* []uint64, reduce: Space */
			reduce(58), /* []uint32, reduce: Space */
			reduce(58), /* []double, reduce: Space */
			reduce(58), /* []float, reduce: Space */
			reduce(58), /* []string, reduce: Space */
			reduce(58), /* [][]byte, reduce: Space */
			reduce(58), /* int64_lit, reduce: Space */
			reduce(58), /* int32_lit, reduce: Space */
			reduce(58), /* uint64_lit, reduce: Space */
			reduce(58), /* uint32_lit, reduce: Space */
			reduce(58), /* double_lit, reduce: Space */
			reduce(58), /* float_lit, reduce: Space */
			reduce(58), /* string_lit, reduce: Space */
			reduce(58), /* bytes_lit, reduce: Space */
			reduce(58), /* bool_var, reduce: Space */
			reduce(58), /* int64_var, reduce: Space */
			reduce(58), /* int32_var, reduce: Space */
			reduce(58), /* uint64_var, reduce: Space */
			reduce(58), /* uint32_var, reduce: Space */
			reduce(58), /* double_var, reduce: Space */
			reduce(58), /* float_var, reduce: Space */
			reduce(58), /* string_var, reduce: Space */
			reduce(58), /* bytes_var, reduce: Space */
			reduce(58), /* true, reduce: Space */
			reduce(58), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(52), /* id, reduce: OpenCurly */
			reduce(52), /* []bool, reduce: OpenCurly */
			reduce(52), /* []int64, reduce: OpenCurly */
			reduce(52), /* []int32, reduce: OpenCurly */
			reduce(52), /* []uint64, reduce: OpenCurly */
			reduce(52), /* []uint32, reduce: OpenCurly */
			reduce(52), /* []double, reduce: OpenCurly */
			reduce(52), /* []float, reduce: OpenCurly */
			reduce(52), /* []string, reduce: OpenCurly */
			reduce(52), /* [][]byte, reduce: OpenCurly */
			reduce(52), /* int64_lit, reduce: OpenCurly */
			reduce(52), /* int32_lit, reduce: OpenCurly */
			reduce(52), /* uint64_lit, reduce: OpenCurly */
			reduce(52), /* uint32_lit, reduce: OpenCurly */
			reduce(52), /* double_lit, reduce: OpenCurly */
			reduce(52), /* float_lit, reduce: OpenCurly */
			reduce(52), /* string_lit, reduce: OpenCurly */
			reduce(52), /* bytes_lit, reduce: OpenCurly */
			reduce(52), /* bool_var, reduce: OpenCurly */
			reduce(52), /* int64_var, reduce: OpenCurly */
			reduce(52), /* int32_var, reduce: OpenCurly */
			reduce(52), /* uint64_var, reduce: OpenCurly */
			reduce(52), /* uint32_var, reduce: OpenCurly */
			reduce(52), /* double_var, reduce: OpenCurly */
			reduce(52), /* float_var, reduce: OpenCurly */
			reduce(52), /* string_var, reduce: OpenCurly */
			reduce(52), /* bytes_var, reduce: OpenCurly */
			reduce(52), /* true, reduce: OpenCurly */
			reduce(52), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(52), /* }, reduce: OpenCurly */
			nil,        /* , */
			reduce(52), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(57), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(126), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(129), /* } */
			nil,        /* , */
			shift(130), /* space */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(116), /* } */
			shift(66),  /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			shift(70), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(71), /* space */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(8), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(8), /* ,, reduce: List */
			reduce(8), /* space, reduce: List */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(9), /* }, reduce: Exprs */
			reduce(9), /* ,, reduce: Exprs */
			reduce(9), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(23), /* }, reduce: SpaceTerminal */
			reduce(23), /* ,, reduce: SpaceTerminal */
			reduce(23), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(34), /* }, reduce: Terminal */
			reduce(34), /* ,, reduce: Terminal */
			reduce(34), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(35), /* }, reduce: Terminal */
			reduce(35), /* ,, reduce: Terminal */
			reduce(35), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(36), /* }, reduce: Terminal */
			reduce(36), /* ,, reduce: Terminal */
			reduce(36), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(37), /* }, reduce: Terminal */
			reduce(37), /* ,, reduce: Terminal */
			reduce(37), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(38), /* }, reduce: Terminal */
			reduce(38), /* ,, reduce: Terminal */
			reduce(38), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(39), /* }, reduce: Terminal */
			reduce(39), /* ,, reduce: Terminal */
			reduce(39), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(40), /* }, reduce: Terminal */
			reduce(40), /* ,, reduce: Terminal */
			reduce(40), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(41), /* }, reduce: Terminal */
			reduce(41), /* ,, reduce: Terminal */
			reduce(41), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(42), /* }, reduce: Terminal */
			reduce(42), /* ,, reduce: Terminal */
			reduce(42), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(43), /* }, reduce: Bool */
			reduce(43), /* ,, reduce: Bool */
			reduce(43), /* space, reduce: Bool */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(44), /* }, reduce: Bool */
			reduce(44), /* ,, reduce: Bool */
			reduce(44), /* space, reduce: Bool */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(53), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(53), /* ,, reduce: CloseCurly */
			reduce(53), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(58), /* id, reduce: Space */
			reduce(58), /* []bool, reduce: Space */
			reduce(58), /* []int64, reduce: Space */
			reduce(58), /* []int32, reduce: Space */
			reduce(58), /* []uint64, reduce: Space */
			reduce(58), /* []uint32, reduce: Space */
			reduce(58), /* []double, reduce: Space */
			reduce(58), /* []float, reduce: Space */
			reduce(58), /* []string, reduce: Space */
			reduce(58), /* [][]byte, reduce: Space */
			reduce(58), /* int64_lit, reduce: Space */
			reduce(58), /* int32_lit, reduce: Space */
			reduce(58), /* uint64_lit, reduce: Space */
			reduce(58), /* uint32_lit, reduce: Space */
			reduce(58), /* double_lit, reduce: Space */
			reduce(58), /* float_lit, reduce: Space */
			reduce(58), /* string_lit, reduce: Space */
			reduce(58), /* bytes_lit, reduce: Space */
			reduce(58), /* bool_var, reduce: Space */
			reduce(58), /* int64_var, reduce: Space */
			reduce(58), /* int32_var, reduce: Space */
			reduce(58), /* uint64_var, reduce: Space */
			reduce(58), /* uint32_var, reduce: Space */
			reduce(58), /* double_var, reduce: Space */
			reduce(58), /* float_var, reduce: Space */
			reduce(58), /* string_var, reduce: Space */
			reduce(58), /* bytes_var, reduce: Space */
			reduce(58), /* true, reduce: Space */
			reduce(58), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(58), /* }, reduce: Space */
			nil,        /* , */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			shift(78), /* ) */
			nil,       /* { */
			nil,       /* } */
			shift(66), /* , */
			shift(67), /* space */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(2), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(2), /* ,, reduce: Function */
			reduce(2), /* space, reduce: Function */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(116), /* } */
			shift(66),  /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(7), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(7), /* ,, reduce: List */
			reduce(7), /* space, reduce: List */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(50), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(50), /* ,, reduce: CloseParen */
			reduce(50), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(122), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(79),  /* , */
			shift(80),  /* space */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(3), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(3), /* ,, reduce: Function */
			reduce(3), /* space, reduce: Function */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(57), /* id, reduce: Space */
			reduce(57), /* []bool, reduce: Space */
			reduce(57), /* []int64, reduce: Space */
			reduce(57), /* []int32, reduce: Space */
			reduce(57), /* []uint64, reduce: Space */
			reduce(57), /* []uint32, reduce: Space */
			reduce(57), /* []double, reduce: Space */
			reduce(57), /* []float, reduce: Space */
			reduce(57), /* []string, reduce: Space */
			reduce(57), /* [][]byte, reduce: Space */
			reduce(57), /* int64_lit, reduce: Space */
			reduce(57), /* int32_lit, reduce: Space */
			reduce(57), /* uint64_lit, reduce: Space */
			reduce(57), /* uint32_lit, reduce: Space */
			reduce(57), /* double_lit, reduce: Space */
			reduce(57), /* float_lit, reduce: Space */
			reduce(57), /* string_lit, reduce: Space */
			reduce(57), /* bytes_lit, reduce: Space */
			reduce(57), /* bool_var, reduce: Space */
			reduce(57), /* int64_var, reduce: Space */
			reduce(57), /* int32_var, reduce: Space */
			reduce(57), /* uint64_var, reduce: Space */
			reduce(57), /* uint32_var, reduce: Space */
			reduce(57), /* double_var, reduce: Space */
			reduce(57), /* float_var, reduce: Space */
			reduce(57), /* string_var, reduce: Space */
			reduce(57), /* bytes_var, reduce: Space */
			reduce(57), /* true, reduce: Space */
			reduce(57), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			shift(9),  /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			shift(10), /* space */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			shift(70), /* { */
			nil,       /* } */
			nil,       /* , */
			shift(71), /* space */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(24), /* }, reduce: SpaceTerminal */
			reduce(24), /* ,, reduce: SpaceTerminal */
			reduce(24), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(54), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(54), /* ,, reduce: CloseCurly */
			reduce(54), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(57), /* id, reduce: Space */
			reduce(57), /* []bool, reduce: Space */
			reduce(57), /* []int64, reduce: Space */
			reduce(57), /* []int32, reduce: Space */
			reduce(57), /* []uint64, reduce: Space */
			reduce(57), /* []uint32, reduce: Space */
			reduce(57), /* []double, reduce: Space */
			reduce(57), /* []float, reduce: Space */
			reduce(57), /* []string, reduce: Space */
			reduce(57), /* [][]byte, reduce: Space */
			reduce(57), /* int64_lit, reduce: Space */
			reduce(57), /* int32_lit, reduce: Space */
			reduce(57), /* uint64_lit, reduce: Space */
			reduce(57), /* uint32_lit, reduce: Space */
			reduce(57), /* double_lit, reduce: Space */
			reduce(57), /* float_lit, reduce: Space */
			reduce(57), /* string_lit, reduce: Space */
			reduce(57), /* bytes_lit, reduce: Space */
			reduce(57), /* bool_var, reduce: Space */
			reduce(57), /* int64_var, reduce: Space */
			reduce(57), /* int32_var, reduce: Space */
			reduce(57), /* uint64_var, reduce: Space */
			reduce(57), /* uint32_var, reduce: Space */
			reduce(57), /* double_var, reduce: Space */
			reduce(57), /* float_var, reduce: Space */
			reduce(57), /* string_var, reduce: Space */
			reduce(57), /* bytes_var, reduce: Space */
			reduce(57), /* true, reduce: Space */
			reduce(57), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(57), /* }, reduce: Space */
			nil,        /* , */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(16),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(34),  /* int64_lit */
			shift(35),  /* int32_lit */
			shift(36),  /* uint64_lit */
			shift(37),  /* uint32_lit */
			shift(38),  /* double_lit */
			shift(39),  /* float_lit */
			shift(40),  /* string_lit */
			shift(41),  /* bytes_lit */
			shift(42),  /* bool_var */
			shift(43),  /* int64_var */
			shift(44),  /* int32_var */
			shift(45),  /* uint64_var */
			shift(46),  /* uint32_var */
			shift(47),  /* double_var */
			shift(48),  /* float_var */
			shift(49),  /* string_var */
			shift(50),  /* bytes_var */
			shift(51),  /* true */
			shift(52),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(144), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(54),  /* space */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(129), /* } */
			shift(79),  /* , */
			shift(145), /* space */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(6), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(6), /* ,, reduce: List */
			reduce(6), /* space, reduce: List */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(83),  /* space */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(58), /* }, reduce: Space */
			reduce(58), /* ,, reduce: Space */
			reduce(58), /* space, reduce: Space */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(151), /* } */
			nil,        /* , */
			shift(117), /* space */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(1), /* ), reduce: Function */
			nil,       /* { */
			nil,       /* } */
			reduce(1), /* ,, reduce: Function */
			reduce(1), /* space, reduce: Function */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			reduce(5), /* ), reduce: List */
			nil,       /* { */
			nil,       /* } */
			reduce(5), /* ,, reduce: List */
			reduce(5), /* space, reduce: List */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(16),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(34),  /* int64_lit */
			shift(35),  /* int32_lit */
			shift(36),  /* uint64_lit */
			shift(37),  /* uint32_lit */
			shift(38),  /* double_lit */
			shift(39),  /* float_lit */
			shift(40),  /* string_lit */
			shift(41),  /* bytes_lit */
			shift(42),  /* bool_var */
			shift(43),  /* int64_var */
			shift(44),  /* int32_var */
			shift(45),  /* uint64_var */
			shift(46),  /* uint32_var */
			shift(47),  /* double_var */
			shift(48),  /* float_var */
			shift(49),  /* string_var */
			shift(50),  /* bytes_var */
			shift(51),  /* true */
			shift(52),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(144), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(54),  /* space */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(88),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(151), /* } */
			nil,        /* , */
			shift(117), /* space */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(57),  /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(34),  /* int64_lit */
			shift(35),  /* int32_lit */
			shift(36),  /* uint64_lit */
			shift(37),  /* uint32_lit */
			shift(38),  /* double_lit */
			shift(39),  /* float_lit */
			shift(40),  /* string_lit */
			shift(41),  /* bytes_lit */
			shift(42),  /* bool_var */
			shift(43),  /* int64_var */
			shift(44),  /* int32_var */
			shift(45),  /* uint64_var */
			shift(46),  /* uint32_var */
			shift(47),  /* double_var */
			shift(48),  /* float_var */
			shift(49),  /* string_var */
			shift(50),  /* bytes_var */
			shift(51),  /* true */
			shift(52),  /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(156), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(61),  /* space */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(144), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(66),  /* , */
			shift(67),  /* space */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(4), /* }, reduce: Function */
			reduce(4), /* ,, reduce: Function */
			reduce(4), /* space, reduce: Function */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(49), /* }, reduce: CloseParen */
			reduce(49), /* ,, reduce: CloseParen */
			reduce(49), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(57), /* }, reduce: Space */
			reduce(57), /* ,, reduce: Space */
			reduce(57), /* space, reduce: Space */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(126), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			shift(125), /* space */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
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
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(126), /* id */
			shift(23),  /* []bool */
			shift(24),  /* []int64 */
			shift(25),  /* []int32 */
			shift(26),  /* []uint64 */
			shift(27),  /* []uint32 */
			shift(28),  /* []double */
			shift(29),  /* []float */
			shift(30),  /* []string */
			shift(31),  /* [][]byte */
			shift(97),  /* int64_lit */
			shift(98),  /* int32_lit */
			shift(99),  /* uint64_lit */
			shift(100), /* uint32_lit */
			shift(101), /* double_lit */
			shift(102), /* float_lit */
			shift(103), /* string_lit */
			shift(104), /* bytes_lit */
			shift(105), /* bool_var */
			shift(106), /* int64_var */
			shift(107), /* int32_var */
			shift(108), /* uint64_var */
			shift(109), /* uint32_var */
			shift(110), /* double_var */
			shift(111), /* float_var */
			shift(112), /* string_var */
			shift(113), /* bytes_var */
			shift(114), /* true */
			shift(115), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(159), /* } */
			nil,        /* , */
			shift(130), /* space */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(151), /* } */
			shift(66),  /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S150
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(8), /* }, reduce: List */
			reduce(8), /* ,, reduce: List */
			reduce(8), /* space, reduce: List */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(53), /* }, reduce: CloseCurly */
			reduce(53), /* ,, reduce: CloseCurly */
			reduce(53), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(144), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(66),  /* , */
			shift(67),  /* space */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(2), /* }, reduce: Function */
			reduce(2), /* ,, reduce: Function */
			reduce(2), /* space, reduce: Function */

		},
	},
	actionRow{ // S154
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(151), /* } */
			shift(66),  /* , */
			shift(135), /* space */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(7), /* }, reduce: List */
			reduce(7), /* ,, reduce: List */
			reduce(7), /* space, reduce: List */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(50), /* }, reduce: CloseParen */
			reduce(50), /* ,, reduce: CloseParen */
			reduce(50), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S157
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(156), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(79),  /* , */
			shift(80),  /* space */

		},
	},
	actionRow{ // S158
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(3), /* }, reduce: Function */
			reduce(3), /* ,, reduce: Function */
			reduce(3), /* space, reduce: Function */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(54), /* }, reduce: CloseCurly */
			reduce(54), /* ,, reduce: CloseCurly */
			reduce(54), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* []bool */
			nil,        /* []int64 */
			nil,        /* []int32 */
			nil,        /* []uint64 */
			nil,        /* []uint32 */
			nil,        /* []double */
			nil,        /* []float */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int64_lit */
			nil,        /* int32_lit */
			nil,        /* uint64_lit */
			nil,        /* uint32_lit */
			nil,        /* double_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int64_var */
			nil,        /* int32_var */
			nil,        /* uint64_var */
			nil,        /* uint32_var */
			nil,        /* double_var */
			nil,        /* float_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(159), /* } */
			shift(79),  /* , */
			shift(145), /* space */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(6), /* }, reduce: List */
			reduce(6), /* ,, reduce: List */
			reduce(6), /* space, reduce: List */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(1), /* }, reduce: Function */
			reduce(1), /* ,, reduce: Function */
			reduce(1), /* space, reduce: Function */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
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
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* { */
			reduce(5), /* }, reduce: List */
			reduce(5), /* ,, reduce: List */
			reduce(5), /* space, reduce: List */

		},
	},
}

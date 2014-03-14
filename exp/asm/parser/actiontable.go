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
			shift(7), /* root */
			nil,      /* = */
			shift(8), /* id */
			nil,      /* . */
			shift(9), /* if */
			nil,      /* then */
			nil,      /* else */
			nil,      /* { */
			nil,      /* } */
			nil,      /* ( */
			nil,      /* ) */
			nil,      /* == */
			nil,      /* < */
			nil,      /* <= */
			nil,      /* > */
			nil,      /* >= */
			nil,      /* && */
			nil,      /* || */
			nil,      /* or */
			nil,      /* and */
			nil,      /* , */
			nil,      /* int64 */
			nil,      /* int_lit */
			nil,      /* uint64 */
			nil,      /* string_lit */
			nil,      /* true */
			nil,      /* True */
			nil,      /* false */
			nil,      /* False */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			shift(7),     /* root */
			nil,          /* = */
			shift(8),     /* id */
			nil,          /* . */
			shift(9),     /* if */
			nil,          /* then */
			nil,          /* else */
			nil,          /* { */
			nil,          /* } */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* == */
			nil,          /* < */
			nil,          /* <= */
			nil,          /* > */
			nil,          /* >= */
			nil,          /* && */
			nil,          /* || */
			nil,          /* or */
			nil,          /* and */
			nil,          /* , */
			nil,          /* int64 */
			nil,          /* int_lit */
			nil,          /* uint64 */
			nil,          /* string_lit */
			nil,          /* true */
			nil,          /* True */
			nil,          /* false */
			nil,          /* False */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Rules */
			reduce(2), /* root, reduce: Rules */
			nil,       /* = */
			reduce(2), /* id, reduce: Rules */
			nil,       /* . */
			reduce(2), /* if, reduce: Rules */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Rule */
			reduce(3), /* root, reduce: Rule */
			nil,       /* = */
			reduce(3), /* id, reduce: Rule */
			nil,       /* . */
			reduce(3), /* if, reduce: Rule */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Rule */
			reduce(4), /* root, reduce: Rule */
			nil,       /* = */
			reduce(4), /* id, reduce: Rule */
			nil,       /* . */
			reduce(4), /* if, reduce: Rule */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Rule */
			reduce(5), /* root, reduce: Rule */
			nil,       /* = */
			reduce(5), /* id, reduce: Rule */
			nil,       /* . */
			reduce(5), /* if, reduce: Rule */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Rule */
			reduce(6), /* root, reduce: Rule */
			nil,       /* = */
			reduce(6), /* id, reduce: Rule */
			nil,       /* . */
			reduce(6), /* if, reduce: Rule */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(11), /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(12), /* id */
			shift(13), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(14), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(17), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(20), /* int64 */
			nil,       /* int_lit */
			shift(21), /* uint64 */
			shift(22), /* string_lit */
			shift(23), /* true */
			shift(24), /* True */
			shift(25), /* false */
			shift(26), /* False */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Rules */
			reduce(1), /* root, reduce: Rules */
			nil,       /* = */
			reduce(1), /* id, reduce: Rules */
			nil,       /* . */
			reduce(1), /* if, reduce: Rules */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(27), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(28), /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(29), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			shift(30), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(31), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			shift(32), /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(28), /* then, reduce: Expr */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(33), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(39), /* int64 */
			nil,       /* int_lit */
			shift(40), /* uint64 */
			shift(41), /* string_lit */
			shift(42), /* true */
			shift(43), /* True */
			shift(44), /* false */
			shift(45), /* False */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(27), /* then, reduce: Expr */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(29), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(46), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(32), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(34), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(35), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(36), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(37), /* then, reduce: Bool */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			shift(48), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(49), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			shift(50), /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(51), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(52), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(55), /* ( */
			shift(57), /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(60), /* int64 */
			nil,       /* int_lit */
			shift(61), /* uint64 */
			shift(62), /* string_lit */
			shift(63), /* true */
			shift(64), /* True */
			shift(65), /* false */
			shift(66), /* False */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(67), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			shift(69), /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			shift(70), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(71), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			shift(73), /* == */
			shift(74), /* < */
			shift(75), /* <= */
			shift(76), /* > */
			shift(77), /* >= */
			shift(78), /* && */
			shift(79), /* || */
			shift(80), /* or */
			shift(81), /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(28), /* ==, reduce: Expr */
			reduce(28), /* <, reduce: Expr */
			reduce(28), /* <=, reduce: Expr */
			reduce(28), /* >, reduce: Expr */
			reduce(28), /* >=, reduce: Expr */
			reduce(28), /* &&, reduce: Expr */
			reduce(28), /* ||, reduce: Expr */
			reduce(28), /* or, reduce: Expr */
			reduce(28), /* and, reduce: Expr */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(33), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(39), /* int64 */
			nil,       /* int_lit */
			shift(40), /* uint64 */
			shift(41), /* string_lit */
			shift(42), /* true */
			shift(43), /* True */
			shift(44), /* false */
			shift(45), /* False */

		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(27), /* ==, reduce: Expr */
			reduce(27), /* <, reduce: Expr */
			reduce(27), /* <=, reduce: Expr */
			reduce(27), /* >, reduce: Expr */
			reduce(27), /* >=, reduce: Expr */
			reduce(27), /* &&, reduce: Expr */
			reduce(27), /* ||, reduce: Expr */
			reduce(27), /* or, reduce: Expr */
			reduce(27), /* and, reduce: Expr */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(29), /* ==, reduce: Terminal */
			reduce(29), /* <, reduce: Terminal */
			reduce(29), /* <=, reduce: Terminal */
			reduce(29), /* >, reduce: Terminal */
			reduce(29), /* >=, reduce: Terminal */
			reduce(29), /* &&, reduce: Terminal */
			reduce(29), /* ||, reduce: Terminal */
			reduce(29), /* or, reduce: Terminal */
			reduce(29), /* and, reduce: Terminal */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(83), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(84), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(32), /* ==, reduce: Terminal */
			reduce(32), /* <, reduce: Terminal */
			reduce(32), /* <=, reduce: Terminal */
			reduce(32), /* >, reduce: Terminal */
			reduce(32), /* >=, reduce: Terminal */
			reduce(32), /* &&, reduce: Terminal */
			reduce(32), /* ||, reduce: Terminal */
			reduce(32), /* or, reduce: Terminal */
			reduce(32), /* and, reduce: Terminal */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(34), /* ==, reduce: Bool */
			reduce(34), /* <, reduce: Bool */
			reduce(34), /* <=, reduce: Bool */
			reduce(34), /* >, reduce: Bool */
			reduce(34), /* >=, reduce: Bool */
			reduce(34), /* &&, reduce: Bool */
			reduce(34), /* ||, reduce: Bool */
			reduce(34), /* or, reduce: Bool */
			reduce(34), /* and, reduce: Bool */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(35), /* ==, reduce: Bool */
			reduce(35), /* <, reduce: Bool */
			reduce(35), /* <=, reduce: Bool */
			reduce(35), /* >, reduce: Bool */
			reduce(35), /* >=, reduce: Bool */
			reduce(35), /* &&, reduce: Bool */
			reduce(35), /* ||, reduce: Bool */
			reduce(35), /* or, reduce: Bool */
			reduce(35), /* and, reduce: Bool */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(36), /* ==, reduce: Bool */
			reduce(36), /* <, reduce: Bool */
			reduce(36), /* <=, reduce: Bool */
			reduce(36), /* >, reduce: Bool */
			reduce(36), /* >=, reduce: Bool */
			reduce(36), /* &&, reduce: Bool */
			reduce(36), /* ||, reduce: Bool */
			reduce(36), /* or, reduce: Bool */
			reduce(36), /* and, reduce: Bool */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(37), /* ==, reduce: Bool */
			reduce(37), /* <, reduce: Bool */
			reduce(37), /* <=, reduce: Bool */
			reduce(37), /* >, reduce: Bool */
			reduce(37), /* >=, reduce: Bool */
			reduce(37), /* &&, reduce: Bool */
			reduce(37), /* ||, reduce: Bool */
			reduce(37), /* or, reduce: Bool */
			reduce(37), /* and, reduce: Bool */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			shift(85), /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			shift(86), /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(87), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: Transition */
			reduce(9), /* root, reduce: Transition */
			nil,       /* = */
			reduce(9), /* id, reduce: Transition */
			nil,       /* . */
			reduce(9), /* if, reduce: Transition */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(88), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			shift(89), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			shift(90), /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(91), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(26), /* ), reduce: Params */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(26), /* ,, reduce: Params */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(28), /* ), reduce: Expr */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(28), /* ,, reduce: Expr */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(33), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(39), /* int64 */
			nil,       /* int_lit */
			shift(40), /* uint64 */
			shift(41), /* string_lit */
			shift(42), /* true */
			shift(43), /* True */
			shift(44), /* false */
			shift(45), /* False */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			shift(93), /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			shift(94), /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(14), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(27), /* ), reduce: Expr */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(27), /* ,, reduce: Expr */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(29), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(29), /* ,, reduce: Terminal */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(95), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(96), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(32), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(32), /* ,, reduce: Terminal */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(34), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(34), /* ,, reduce: Bool */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(35), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(35), /* ,, reduce: Bool */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(36), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(36), /* ,, reduce: Bool */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(37), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(37), /* ,, reduce: Bool */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			reduce(12), /* else, reduce: StateExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			shift(97), /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			shift(99), /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(100), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(52),  /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(55),  /* ( */
			shift(102), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(60),  /* int64 */
			nil,        /* int_lit */
			shift(61),  /* uint64 */
			shift(62),  /* string_lit */
			shift(63),  /* true */
			shift(64),  /* True */
			shift(65),  /* false */
			shift(66),  /* False */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(106), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(109), /* int64 */
			nil,        /* int_lit */
			shift(110), /* uint64 */
			shift(111), /* string_lit */
			shift(112), /* true */
			shift(113), /* True */
			shift(114), /* false */
			shift(115), /* False */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(16), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(16), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(16), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(16), /* uint64, reduce: Comparator */
			reduce(16), /* string_lit, reduce: Comparator */
			reduce(16), /* true, reduce: Comparator */
			reduce(16), /* True, reduce: Comparator */
			reduce(16), /* false, reduce: Comparator */
			reduce(16), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(17), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(17), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(17), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(17), /* uint64, reduce: Comparator */
			reduce(17), /* string_lit, reduce: Comparator */
			reduce(17), /* true, reduce: Comparator */
			reduce(17), /* True, reduce: Comparator */
			reduce(17), /* false, reduce: Comparator */
			reduce(17), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(18), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(18), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(18), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(18), /* uint64, reduce: Comparator */
			reduce(18), /* string_lit, reduce: Comparator */
			reduce(18), /* true, reduce: Comparator */
			reduce(18), /* True, reduce: Comparator */
			reduce(18), /* false, reduce: Comparator */
			reduce(18), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(19), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(19), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(19), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(19), /* uint64, reduce: Comparator */
			reduce(19), /* string_lit, reduce: Comparator */
			reduce(19), /* true, reduce: Comparator */
			reduce(19), /* True, reduce: Comparator */
			reduce(19), /* false, reduce: Comparator */
			reduce(19), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(20), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(20), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(20), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(20), /* uint64, reduce: Comparator */
			reduce(20), /* string_lit, reduce: Comparator */
			reduce(20), /* true, reduce: Comparator */
			reduce(20), /* True, reduce: Comparator */
			reduce(20), /* false, reduce: Comparator */
			reduce(20), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(21), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(21), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(21), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(21), /* uint64, reduce: Comparator */
			reduce(21), /* string_lit, reduce: Comparator */
			reduce(21), /* true, reduce: Comparator */
			reduce(21), /* True, reduce: Comparator */
			reduce(21), /* false, reduce: Comparator */
			reduce(21), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(22), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(22), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(22), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(22), /* uint64, reduce: Comparator */
			reduce(22), /* string_lit, reduce: Comparator */
			reduce(22), /* true, reduce: Comparator */
			reduce(22), /* True, reduce: Comparator */
			reduce(22), /* false, reduce: Comparator */
			reduce(22), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(23), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(23), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(23), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(23), /* uint64, reduce: Comparator */
			reduce(23), /* string_lit, reduce: Comparator */
			reduce(23), /* true, reduce: Comparator */
			reduce(23), /* True, reduce: Comparator */
			reduce(23), /* false, reduce: Comparator */
			reduce(23), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			reduce(24), /* id, reduce: Comparator */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			reduce(24), /* (, reduce: Comparator */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			reduce(24), /* int64, reduce: Comparator */
			nil,        /* int_lit */
			reduce(24), /* uint64, reduce: Comparator */
			reduce(24), /* string_lit, reduce: Comparator */
			reduce(24), /* true, reduce: Comparator */
			reduce(24), /* True, reduce: Comparator */
			reduce(24), /* false, reduce: Comparator */
			reduce(24), /* False, reduce: Comparator */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			shift(73), /* == */
			shift(74), /* < */
			shift(75), /* <= */
			shift(76), /* > */
			shift(77), /* >= */
			shift(78), /* && */
			shift(79), /* || */
			shift(80), /* or */
			shift(81), /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(117), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(118), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(119), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(120), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: Root */
			reduce(7), /* root, reduce: Root */
			nil,       /* = */
			reduce(7), /* id, reduce: Root */
			nil,       /* . */
			reduce(7), /* if, reduce: Root */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: Init */
			reduce(8), /* root, reduce: Init */
			nil,       /* = */
			reduce(8), /* id, reduce: Init */
			nil,       /* . */
			reduce(8), /* if, reduce: Init */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(121), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(122), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(52),  /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(55),  /* ( */
			shift(124), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(60),  /* int64 */
			nil,        /* int_lit */
			shift(61),  /* uint64 */
			shift(62),  /* string_lit */
			shift(63),  /* true */
			shift(64),  /* True */
			shift(65),  /* false */
			shift(66),  /* False */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			shift(73), /* == */
			shift(74), /* < */
			shift(75), /* <= */
			shift(76), /* > */
			shift(77), /* >= */
			shift(78), /* && */
			shift(79), /* || */
			shift(80), /* or */
			shift(81), /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(13), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(52), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(55), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(60), /* int64 */
			nil,       /* int_lit */
			shift(61), /* uint64 */
			shift(62), /* string_lit */
			shift(63), /* true */
			shift(64), /* True */
			shift(65), /* false */
			shift(66), /* False */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(127), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(128), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(129), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			shift(131), /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			shift(132), /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(14), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(17), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(20), /* int64 */
			nil,       /* int_lit */
			shift(21), /* uint64 */
			shift(22), /* string_lit */
			shift(23), /* true */
			shift(24), /* True */
			shift(25), /* false */
			shift(26), /* False */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			shift(134), /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(135), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			shift(94),  /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(14), /* ==, reduce: Function */
			reduce(14), /* <, reduce: Function */
			reduce(14), /* <=, reduce: Function */
			reduce(14), /* >, reduce: Function */
			reduce(14), /* >=, reduce: Function */
			reduce(14), /* &&, reduce: Function */
			reduce(14), /* ||, reduce: Function */
			reduce(14), /* or, reduce: Function */
			reduce(14), /* and, reduce: Function */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			shift(136), /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(137), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(138), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(28), /* ), reduce: Expr */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(33), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			shift(36), /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			shift(39), /* int64 */
			nil,       /* int_lit */
			shift(40), /* uint64 */
			shift(41), /* string_lit */
			shift(42), /* true */
			shift(43), /* True */
			shift(44), /* false */
			shift(45), /* False */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(27), /* ), reduce: Expr */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(29), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(140), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(141), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(32), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(34), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(35), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(36), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(37), /* ), reduce: Bool */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(106), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(109), /* int64 */
			nil,        /* int_lit */
			shift(110), /* uint64 */
			shift(111), /* string_lit */
			shift(112), /* true */
			shift(113), /* True */
			shift(114), /* false */
			shift(115), /* False */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(143), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(144), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(30), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(31), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(33), /* then, reduce: Terminal */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			shift(145), /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(146), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			shift(94),  /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(14), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(14), /* ,, reduce: Function */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(106), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(109), /* int64 */
			nil,        /* int_lit */
			shift(110), /* uint64 */
			shift(111), /* string_lit */
			shift(112), /* true */
			shift(113), /* True */
			shift(114), /* false */
			shift(115), /* False */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(25), /* ), reduce: Params */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(25), /* ,, reduce: Params */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(148), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(149), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: StateExpr */
			reduce(12), /* root, reduce: StateExpr */
			nil,        /* = */
			reduce(12), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(12), /* if, reduce: StateExpr */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: IfExpr */
			reduce(10), /* root, reduce: IfExpr */
			nil,        /* = */
			reduce(10), /* id, reduce: IfExpr */
			nil,        /* . */
			reduce(10), /* if, reduce: IfExpr */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			shift(99), /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			reduce(11), /* else, reduce: StateExpr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			shift(151), /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(152), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(13), /* ==, reduce: Function */
			reduce(13), /* <, reduce: Function */
			reduce(13), /* <=, reduce: Function */
			reduce(13), /* >, reduce: Function */
			reduce(13), /* >=, reduce: Function */
			reduce(13), /* &&, reduce: Function */
			reduce(13), /* ||, reduce: Function */
			reduce(13), /* or, reduce: Function */
			reduce(13), /* and, reduce: Function */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(153), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(52),  /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(55),  /* ( */
			shift(155), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(60),  /* int64 */
			nil,        /* int_lit */
			shift(61),  /* uint64 */
			shift(62),  /* string_lit */
			shift(63),  /* true */
			shift(64),  /* True */
			shift(65),  /* false */
			shift(66),  /* False */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			reduce(15), /* then, reduce: Function */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			shift(73), /* == */
			shift(74), /* < */
			shift(75), /* <= */
			shift(76), /* > */
			shift(77), /* >= */
			shift(78), /* && */
			shift(79), /* || */
			shift(80), /* or */
			shift(81), /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(157), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			shift(158), /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(159), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(30), /* ==, reduce: Terminal */
			reduce(30), /* <, reduce: Terminal */
			reduce(30), /* <=, reduce: Terminal */
			reduce(30), /* >, reduce: Terminal */
			reduce(30), /* >=, reduce: Terminal */
			reduce(30), /* &&, reduce: Terminal */
			reduce(30), /* ||, reduce: Terminal */
			reduce(30), /* or, reduce: Terminal */
			reduce(30), /* and, reduce: Terminal */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(31), /* ==, reduce: Terminal */
			reduce(31), /* <, reduce: Terminal */
			reduce(31), /* <=, reduce: Terminal */
			reduce(31), /* >, reduce: Terminal */
			reduce(31), /* >=, reduce: Terminal */
			reduce(31), /* &&, reduce: Terminal */
			reduce(31), /* ||, reduce: Terminal */
			reduce(31), /* or, reduce: Terminal */
			reduce(31), /* and, reduce: Terminal */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(160), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(13), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(13), /* ,, reduce: Function */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(161), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(30), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(30), /* ,, reduce: Terminal */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(31), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(31), /* ,, reduce: Terminal */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S150
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			shift(162), /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			shift(67), /* id */
			nil,       /* . */
			nil,       /* if */
			nil,       /* then */
			nil,       /* else */
			shift(69), /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(33), /* ==, reduce: Terminal */
			reduce(33), /* <, reduce: Terminal */
			reduce(33), /* <=, reduce: Terminal */
			reduce(33), /* >, reduce: Terminal */
			reduce(33), /* >=, reduce: Terminal */
			reduce(33), /* &&, reduce: Terminal */
			reduce(33), /* ||, reduce: Terminal */
			reduce(33), /* or, reduce: Terminal */
			reduce(33), /* and, reduce: Terminal */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			shift(164), /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S154
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(165), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			shift(94),  /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(14), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(103), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			shift(106), /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			shift(109), /* int64 */
			nil,        /* int_lit */
			shift(110), /* uint64 */
			shift(111), /* string_lit */
			shift(112), /* true */
			shift(113), /* True */
			shift(114), /* false */
			shift(115), /* False */

		},
	},
	actionRow{ // S157
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(167), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S158
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(168), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			reduce(15), /* ==, reduce: Function */
			reduce(15), /* <, reduce: Function */
			reduce(15), /* <=, reduce: Function */
			reduce(15), /* >, reduce: Function */
			reduce(15), /* >=, reduce: Function */
			reduce(15), /* &&, reduce: Function */
			reduce(15), /* ||, reduce: Function */
			reduce(15), /* or, reduce: Function */
			reduce(15), /* and, reduce: Function */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(33), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(33), /* ,, reduce: Terminal */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(15), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			reduce(15), /* ,, reduce: Function */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: StateExpr */
			reduce(11), /* root, reduce: StateExpr */
			nil,        /* = */
			reduce(11), /* id, reduce: StateExpr */
			nil,        /* . */
			reduce(11), /* if, reduce: StateExpr */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			shift(169), /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(170), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S165
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(13), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S166
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			shift(171), /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S167
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(30), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S168
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(31), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S169
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			shift(172), /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			shift(174), /* { */
			nil,        /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S170
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(33), /* ), reduce: Terminal */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S171
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			nil,        /* } */
			nil,        /* ( */
			reduce(15), /* ), reduce: Function */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S172
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			reduce(12), /* }, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S173
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			reduce(10), /* }, reduce: IfExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S174
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* root */
			nil,       /* = */
			nil,       /* id */
			nil,       /* . */
			shift(99), /* if */
			nil,       /* then */
			nil,       /* else */
			nil,       /* { */
			nil,       /* } */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* == */
			nil,       /* < */
			nil,       /* <= */
			nil,       /* > */
			nil,       /* >= */
			nil,       /* && */
			nil,       /* || */
			nil,       /* or */
			nil,       /* and */
			nil,       /* , */
			nil,       /* int64 */
			nil,       /* int_lit */
			nil,       /* uint64 */
			nil,       /* string_lit */
			nil,       /* true */
			nil,       /* True */
			nil,       /* false */
			nil,       /* False */

		},
	},
	actionRow{ // S175
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			shift(176), /* } */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
	actionRow{ // S176
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* root */
			nil,        /* = */
			nil,        /* id */
			nil,        /* . */
			nil,        /* if */
			nil,        /* then */
			nil,        /* else */
			nil,        /* { */
			reduce(11), /* }, reduce: StateExpr */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* == */
			nil,        /* < */
			nil,        /* <= */
			nil,        /* > */
			nil,        /* >= */
			nil,        /* && */
			nil,        /* || */
			nil,        /* or */
			nil,        /* and */
			nil,        /* , */
			nil,        /* int64 */
			nil,        /* int_lit */
			nil,        /* uint64 */
			nil,        /* string_lit */
			nil,        /* true */
			nil,        /* True */
			nil,        /* false */
			nil,        /* False */

		},
	},
}

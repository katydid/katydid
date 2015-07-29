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
			nil,      /* AnyName */
			nil,      /* Name */
			nil,      /* string_lit */
			nil,      /* AnyNameExcept */
			nil,      /* NameChoice */
			nil,      /* Empty */
			nil,      /* EmptySet */
			nil,      /* ZeroOrMore */
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
			shift(6), /* space */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* id */
			nil,          /* AnyName */
			nil,          /* Name */
			nil,          /* string_lit */
			nil,          /* AnyNameExcept */
			nil,          /* NameChoice */
			nil,          /* Empty */
			nil,          /* EmptySet */
			nil,          /* ZeroOrMore */
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
			nil,          /* space */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Grammar */
			shift(5),  /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(9),  /* space */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(10), /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(11), /* space */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: PatternDecls */
			reduce(3), /* id, reduce: PatternDecls */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(3), /* space, reduce: PatternDecls */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(15), /* space */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(94), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Grammar */
			shift(10), /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(16), /* space */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: PatternDecls */
			reduce(4), /* id, reduce: PatternDecls */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(4), /* space, reduce: PatternDecls */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(94), /* $, reduce: Space */
			reduce(94), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(15), /* space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(93), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(19), /* space */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(21), /* id */
			shift(24), /* AnyName */
			shift(25), /* Name */
			shift(27), /* string_lit */
			shift(28), /* AnyNameExcept */
			shift(29), /* NameChoice */
			shift(30), /* Empty */
			shift(31), /* EmptySet */
			shift(34), /* ZeroOrMore */
			shift(41), /* []bool */
			shift(42), /* []int */
			shift(43), /* []uint */
			shift(44), /* []double */
			shift(45), /* []string */
			shift(46), /* [][]byte */
			shift(49), /* int_lit */
			shift(50), /* uint_lit */
			shift(51), /* double_lit */
			shift(52), /* bytes_lit */
			shift(53), /* bool_var */
			shift(54), /* int_var */
			shift(55), /* uint_var */
			shift(56), /* double_var */
			shift(57), /* string_var */
			shift(58), /* bytes_var */
			shift(59), /* true */
			shift(60), /* false */
			nil,       /* = */
			shift(61), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(62), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(63), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(64), /* ! */
			shift(65), /* space */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(65), /* id, reduce: Equal */
			reduce(65), /* AnyName, reduce: Equal */
			reduce(65), /* Name, reduce: Equal */
			reduce(65), /* string_lit, reduce: Equal */
			reduce(65), /* AnyNameExcept, reduce: Equal */
			reduce(65), /* NameChoice, reduce: Equal */
			reduce(65), /* Empty, reduce: Equal */
			reduce(65), /* EmptySet, reduce: Equal */
			reduce(65), /* ZeroOrMore, reduce: Equal */
			reduce(65), /* []bool, reduce: Equal */
			reduce(65), /* []int, reduce: Equal */
			reduce(65), /* []uint, reduce: Equal */
			reduce(65), /* []double, reduce: Equal */
			reduce(65), /* []string, reduce: Equal */
			reduce(65), /* [][]byte, reduce: Equal */
			reduce(65), /* int_lit, reduce: Equal */
			reduce(65), /* uint_lit, reduce: Equal */
			reduce(65), /* double_lit, reduce: Equal */
			reduce(65), /* bytes_lit, reduce: Equal */
			reduce(65), /* bool_var, reduce: Equal */
			reduce(65), /* int_var, reduce: Equal */
			reduce(65), /* uint_var, reduce: Equal */
			reduce(65), /* double_var, reduce: Equal */
			reduce(65), /* string_var, reduce: Equal */
			reduce(65), /* bytes_var, reduce: Equal */
			reduce(65), /* true, reduce: Equal */
			reduce(65), /* false, reduce: Equal */
			nil,        /* = */
			reduce(65), /* (, reduce: Equal */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(65), /* #, reduce: Equal */
			nil,        /* & */
			nil,        /* | */
			reduce(65), /* [, reduce: Equal */
			nil,        /* ] */
			nil,        /* : */
			reduce(65), /* !, reduce: Equal */
			reduce(65), /* space, reduce: Equal */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(94), /* =, reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(93), /* $, reduce: Space */
			reduce(93), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(21), /* id */
			shift(24), /* AnyName */
			shift(25), /* Name */
			shift(27), /* string_lit */
			shift(28), /* AnyNameExcept */
			shift(29), /* NameChoice */
			shift(30), /* Empty */
			shift(31), /* EmptySet */
			shift(34), /* ZeroOrMore */
			shift(41), /* []bool */
			shift(42), /* []int */
			shift(43), /* []uint */
			shift(44), /* []double */
			shift(45), /* []string */
			shift(46), /* [][]byte */
			shift(49), /* int_lit */
			shift(50), /* uint_lit */
			shift(51), /* double_lit */
			shift(52), /* bytes_lit */
			shift(53), /* bool_var */
			shift(54), /* int_var */
			shift(55), /* uint_var */
			shift(56), /* double_var */
			shift(57), /* string_var */
			shift(58), /* bytes_var */
			shift(59), /* true */
			shift(60), /* false */
			nil,       /* = */
			shift(61), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(62), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(63), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(64), /* ! */
			shift(65), /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(66), /* id, reduce: Equal */
			reduce(66), /* AnyName, reduce: Equal */
			reduce(66), /* Name, reduce: Equal */
			reduce(66), /* string_lit, reduce: Equal */
			reduce(66), /* AnyNameExcept, reduce: Equal */
			reduce(66), /* NameChoice, reduce: Equal */
			reduce(66), /* Empty, reduce: Equal */
			reduce(66), /* EmptySet, reduce: Equal */
			reduce(66), /* ZeroOrMore, reduce: Equal */
			reduce(66), /* []bool, reduce: Equal */
			reduce(66), /* []int, reduce: Equal */
			reduce(66), /* []uint, reduce: Equal */
			reduce(66), /* []double, reduce: Equal */
			reduce(66), /* []string, reduce: Equal */
			reduce(66), /* [][]byte, reduce: Equal */
			reduce(66), /* int_lit, reduce: Equal */
			reduce(66), /* uint_lit, reduce: Equal */
			reduce(66), /* double_lit, reduce: Equal */
			reduce(66), /* bytes_lit, reduce: Equal */
			reduce(66), /* bool_var, reduce: Equal */
			reduce(66), /* int_var, reduce: Equal */
			reduce(66), /* uint_var, reduce: Equal */
			reduce(66), /* double_var, reduce: Equal */
			reduce(66), /* string_var, reduce: Equal */
			reduce(66), /* bytes_var, reduce: Equal */
			reduce(66), /* true, reduce: Equal */
			reduce(66), /* false, reduce: Equal */
			nil,        /* = */
			reduce(66), /* (, reduce: Equal */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(66), /* #, reduce: Equal */
			nil,        /* & */
			nil,        /* | */
			reduce(66), /* [, reduce: Equal */
			nil,        /* ] */
			nil,        /* : */
			reduce(66), /* !, reduce: Equal */
			reduce(66), /* space, reduce: Equal */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(93), /* =, reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(67), /* id */
			shift(68), /* AnyName */
			shift(69), /* Name */
			shift(27), /* string_lit */
			shift(70), /* AnyNameExcept */
			shift(71), /* NameChoice */
			shift(72), /* Empty */
			shift(73), /* EmptySet */
			shift(74), /* ZeroOrMore */
			shift(41), /* []bool */
			shift(42), /* []int */
			shift(43), /* []uint */
			shift(44), /* []double */
			shift(45), /* []string */
			shift(46), /* [][]byte */
			shift(49), /* int_lit */
			shift(50), /* uint_lit */
			shift(51), /* double_lit */
			shift(52), /* bytes_lit */
			shift(53), /* bool_var */
			shift(54), /* int_var */
			shift(55), /* uint_var */
			shift(56), /* double_var */
			shift(57), /* string_var */
			shift(58), /* bytes_var */
			shift(59), /* true */
			shift(60), /* false */
			nil,       /* = */
			shift(77), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(78), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(79), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(80), /* ! */
			shift(81), /* space */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: PatternDecl */
			reduce(6), /* id, reduce: PatternDecl */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(6), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(88), /* : */
			nil,       /* ! */
			shift(89), /* space */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(55), /* $, reduce: Terminal */
			reduce(55), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(18), /* $, reduce: Pattern */
			reduce(18), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(20), /* $, reduce: Pattern */
			reduce(20), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(22), /* $, reduce: Pattern */
			reduce(22), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(161), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(30), /* $, reduce: Expr */
			reduce(30), /* id, reduce: Expr */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(31), /* $, reduce: Expr */
			reduce(31), /* id, reduce: Expr */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(32), /* $, reduce: Expr */
			reduce(32), /* id, reduce: Expr */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(43), /* {, reduce: ListType */
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
			reduce(43), /* space, reduce: ListType */

		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(44), /* {, reduce: ListType */
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
			reduce(44), /* space, reduce: ListType */

		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(45), /* {, reduce: ListType */
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
			reduce(45), /* space, reduce: ListType */

		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(46), /* {, reduce: ListType */
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
			reduce(46), /* space, reduce: ListType */

		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(47), /* {, reduce: ListType */
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
			reduce(47), /* space, reduce: ListType */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(48), /* {, reduce: ListType */
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
			reduce(48), /* space, reduce: ListType */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(49), /* $, reduce: SpaceTerminal */
			reduce(49), /* id, reduce: SpaceTerminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(51), /* $, reduce: Terminal */
			reduce(51), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(52), /* $, reduce: Terminal */
			reduce(52), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(53), /* $, reduce: Terminal */
			reduce(53), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(54), /* $, reduce: Terminal */
			reduce(54), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(56), /* $, reduce: Terminal */
			reduce(56), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(57), /* $, reduce: Terminal */
			reduce(57), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(58), /* $, reduce: Terminal */
			reduce(58), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(59), /* $, reduce: Terminal */
			reduce(59), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(60), /* $, reduce: Terminal */
			reduce(60), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(61), /* $, reduce: Terminal */
			reduce(61), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(62), /* $, reduce: Terminal */
			reduce(62), /* id, reduce: Terminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(63), /* $, reduce: Bool */
			reduce(63), /* id, reduce: Bool */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(64), /* $, reduce: Bool */
			reduce(64), /* id, reduce: Bool */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(67), /* id, reduce: OpenParen */
			reduce(67), /* AnyName, reduce: OpenParen */
			reduce(67), /* Name, reduce: OpenParen */
			reduce(67), /* string_lit, reduce: OpenParen */
			reduce(67), /* AnyNameExcept, reduce: OpenParen */
			reduce(67), /* NameChoice, reduce: OpenParen */
			reduce(67), /* Empty, reduce: OpenParen */
			reduce(67), /* EmptySet, reduce: OpenParen */
			reduce(67), /* ZeroOrMore, reduce: OpenParen */
			reduce(67), /* []bool, reduce: OpenParen */
			reduce(67), /* []int, reduce: OpenParen */
			reduce(67), /* []uint, reduce: OpenParen */
			reduce(67), /* []double, reduce: OpenParen */
			reduce(67), /* []string, reduce: OpenParen */
			reduce(67), /* [][]byte, reduce: OpenParen */
			reduce(67), /* int_lit, reduce: OpenParen */
			reduce(67), /* uint_lit, reduce: OpenParen */
			reduce(67), /* double_lit, reduce: OpenParen */
			reduce(67), /* bytes_lit, reduce: OpenParen */
			reduce(67), /* bool_var, reduce: OpenParen */
			reduce(67), /* int_var, reduce: OpenParen */
			reduce(67), /* uint_var, reduce: OpenParen */
			reduce(67), /* double_var, reduce: OpenParen */
			reduce(67), /* string_var, reduce: OpenParen */
			reduce(67), /* bytes_var, reduce: OpenParen */
			reduce(67), /* true, reduce: OpenParen */
			reduce(67), /* false, reduce: OpenParen */
			nil,        /* = */
			reduce(67), /* (, reduce: OpenParen */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(67), /* #, reduce: OpenParen */
			nil,        /* & */
			nil,        /* | */
			reduce(67), /* [, reduce: OpenParen */
			nil,        /* ] */
			nil,        /* : */
			reduce(67), /* !, reduce: OpenParen */
			reduce(67), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(79), /* id, reduce: HashTag */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(85), /* id, reduce: OpenBracket */
			reduce(85), /* AnyName, reduce: OpenBracket */
			reduce(85), /* Name, reduce: OpenBracket */
			reduce(85), /* string_lit, reduce: OpenBracket */
			reduce(85), /* AnyNameExcept, reduce: OpenBracket */
			reduce(85), /* NameChoice, reduce: OpenBracket */
			reduce(85), /* Empty, reduce: OpenBracket */
			reduce(85), /* EmptySet, reduce: OpenBracket */
			reduce(85), /* ZeroOrMore, reduce: OpenBracket */
			reduce(85), /* []bool, reduce: OpenBracket */
			reduce(85), /* []int, reduce: OpenBracket */
			reduce(85), /* []uint, reduce: OpenBracket */
			reduce(85), /* []double, reduce: OpenBracket */
			reduce(85), /* []string, reduce: OpenBracket */
			reduce(85), /* [][]byte, reduce: OpenBracket */
			reduce(85), /* int_lit, reduce: OpenBracket */
			reduce(85), /* uint_lit, reduce: OpenBracket */
			reduce(85), /* double_lit, reduce: OpenBracket */
			reduce(85), /* bytes_lit, reduce: OpenBracket */
			reduce(85), /* bool_var, reduce: OpenBracket */
			reduce(85), /* int_var, reduce: OpenBracket */
			reduce(85), /* uint_var, reduce: OpenBracket */
			reduce(85), /* double_var, reduce: OpenBracket */
			reduce(85), /* string_var, reduce: OpenBracket */
			reduce(85), /* bytes_var, reduce: OpenBracket */
			reduce(85), /* true, reduce: OpenBracket */
			reduce(85), /* false, reduce: OpenBracket */
			nil,        /* = */
			reduce(85), /* (, reduce: OpenBracket */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(85), /* #, reduce: OpenBracket */
			nil,        /* & */
			nil,        /* | */
			reduce(85), /* [, reduce: OpenBracket */
			nil,        /* ] */
			nil,        /* : */
			reduce(85), /* !, reduce: OpenBracket */
			reduce(85), /* space, reduce: OpenBracket */

		},
	},
	actionRow{ // S64
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(91), /* (, reduce: Exclamation */
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
			reduce(91), /* space, reduce: Exclamation */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(94), /* id, reduce: Space */
			reduce(94), /* AnyName, reduce: Space */
			reduce(94), /* Name, reduce: Space */
			reduce(94), /* string_lit, reduce: Space */
			reduce(94), /* AnyNameExcept, reduce: Space */
			reduce(94), /* NameChoice, reduce: Space */
			reduce(94), /* Empty, reduce: Space */
			reduce(94), /* EmptySet, reduce: Space */
			reduce(94), /* ZeroOrMore, reduce: Space */
			reduce(94), /* []bool, reduce: Space */
			reduce(94), /* []int, reduce: Space */
			reduce(94), /* []uint, reduce: Space */
			reduce(94), /* []double, reduce: Space */
			reduce(94), /* []string, reduce: Space */
			reduce(94), /* [][]byte, reduce: Space */
			reduce(94), /* int_lit, reduce: Space */
			reduce(94), /* uint_lit, reduce: Space */
			reduce(94), /* double_lit, reduce: Space */
			reduce(94), /* bytes_lit, reduce: Space */
			reduce(94), /* bool_var, reduce: Space */
			reduce(94), /* int_var, reduce: Space */
			reduce(94), /* uint_var, reduce: Space */
			reduce(94), /* double_var, reduce: Space */
			reduce(94), /* string_var, reduce: Space */
			reduce(94), /* bytes_var, reduce: Space */
			reduce(94), /* true, reduce: Space */
			reduce(94), /* false, reduce: Space */
			nil,        /* = */
			reduce(94), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(94), /* #, reduce: Space */
			nil,        /* & */
			nil,        /* | */
			reduce(94), /* [, reduce: Space */
			nil,        /* ] */
			nil,        /* : */
			reduce(94), /* !, reduce: Space */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: PatternDecl */
			reduce(5), /* id, reduce: PatternDecl */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(5), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S71
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(17), /* $, reduce: Pattern */
			reduce(17), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(19), /* $, reduce: Pattern */
			reduce(19), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S74
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S75
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S76
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(50), /* $, reduce: SpaceTerminal */
			reduce(50), /* id, reduce: SpaceTerminal */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S77
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(68), /* id, reduce: OpenParen */
			reduce(68), /* AnyName, reduce: OpenParen */
			reduce(68), /* Name, reduce: OpenParen */
			reduce(68), /* string_lit, reduce: OpenParen */
			reduce(68), /* AnyNameExcept, reduce: OpenParen */
			reduce(68), /* NameChoice, reduce: OpenParen */
			reduce(68), /* Empty, reduce: OpenParen */
			reduce(68), /* EmptySet, reduce: OpenParen */
			reduce(68), /* ZeroOrMore, reduce: OpenParen */
			reduce(68), /* []bool, reduce: OpenParen */
			reduce(68), /* []int, reduce: OpenParen */
			reduce(68), /* []uint, reduce: OpenParen */
			reduce(68), /* []double, reduce: OpenParen */
			reduce(68), /* []string, reduce: OpenParen */
			reduce(68), /* [][]byte, reduce: OpenParen */
			reduce(68), /* int_lit, reduce: OpenParen */
			reduce(68), /* uint_lit, reduce: OpenParen */
			reduce(68), /* double_lit, reduce: OpenParen */
			reduce(68), /* bytes_lit, reduce: OpenParen */
			reduce(68), /* bool_var, reduce: OpenParen */
			reduce(68), /* int_var, reduce: OpenParen */
			reduce(68), /* uint_var, reduce: OpenParen */
			reduce(68), /* double_var, reduce: OpenParen */
			reduce(68), /* string_var, reduce: OpenParen */
			reduce(68), /* bytes_var, reduce: OpenParen */
			reduce(68), /* true, reduce: OpenParen */
			reduce(68), /* false, reduce: OpenParen */
			nil,        /* = */
			reduce(68), /* (, reduce: OpenParen */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(68), /* #, reduce: OpenParen */
			nil,        /* & */
			nil,        /* | */
			reduce(68), /* [, reduce: OpenParen */
			nil,        /* ] */
			nil,        /* : */
			reduce(68), /* !, reduce: OpenParen */
			reduce(68), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(80), /* id, reduce: HashTag */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(86), /* id, reduce: OpenBracket */
			reduce(86), /* AnyName, reduce: OpenBracket */
			reduce(86), /* Name, reduce: OpenBracket */
			reduce(86), /* string_lit, reduce: OpenBracket */
			reduce(86), /* AnyNameExcept, reduce: OpenBracket */
			reduce(86), /* NameChoice, reduce: OpenBracket */
			reduce(86), /* Empty, reduce: OpenBracket */
			reduce(86), /* EmptySet, reduce: OpenBracket */
			reduce(86), /* ZeroOrMore, reduce: OpenBracket */
			reduce(86), /* []bool, reduce: OpenBracket */
			reduce(86), /* []int, reduce: OpenBracket */
			reduce(86), /* []uint, reduce: OpenBracket */
			reduce(86), /* []double, reduce: OpenBracket */
			reduce(86), /* []string, reduce: OpenBracket */
			reduce(86), /* [][]byte, reduce: OpenBracket */
			reduce(86), /* int_lit, reduce: OpenBracket */
			reduce(86), /* uint_lit, reduce: OpenBracket */
			reduce(86), /* double_lit, reduce: OpenBracket */
			reduce(86), /* bytes_lit, reduce: OpenBracket */
			reduce(86), /* bool_var, reduce: OpenBracket */
			reduce(86), /* int_var, reduce: OpenBracket */
			reduce(86), /* uint_var, reduce: OpenBracket */
			reduce(86), /* double_var, reduce: OpenBracket */
			reduce(86), /* string_var, reduce: OpenBracket */
			reduce(86), /* bytes_var, reduce: OpenBracket */
			reduce(86), /* true, reduce: OpenBracket */
			reduce(86), /* false, reduce: OpenBracket */
			nil,        /* = */
			reduce(86), /* (, reduce: OpenBracket */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(86), /* #, reduce: OpenBracket */
			nil,        /* & */
			nil,        /* | */
			reduce(86), /* [, reduce: OpenBracket */
			nil,        /* ] */
			nil,        /* : */
			reduce(86), /* !, reduce: OpenBracket */
			reduce(86), /* space, reduce: OpenBracket */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(92), /* (, reduce: Exclamation */
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
			reduce(92), /* space, reduce: Exclamation */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(93), /* id, reduce: Space */
			reduce(93), /* AnyName, reduce: Space */
			reduce(93), /* Name, reduce: Space */
			reduce(93), /* string_lit, reduce: Space */
			reduce(93), /* AnyNameExcept, reduce: Space */
			reduce(93), /* NameChoice, reduce: Space */
			reduce(93), /* Empty, reduce: Space */
			reduce(93), /* EmptySet, reduce: Space */
			reduce(93), /* ZeroOrMore, reduce: Space */
			reduce(93), /* []bool, reduce: Space */
			reduce(93), /* []int, reduce: Space */
			reduce(93), /* []uint, reduce: Space */
			reduce(93), /* []double, reduce: Space */
			reduce(93), /* []string, reduce: Space */
			reduce(93), /* [][]byte, reduce: Space */
			reduce(93), /* int_lit, reduce: Space */
			reduce(93), /* uint_lit, reduce: Space */
			reduce(93), /* double_lit, reduce: Space */
			reduce(93), /* bytes_lit, reduce: Space */
			reduce(93), /* bool_var, reduce: Space */
			reduce(93), /* int_var, reduce: Space */
			reduce(93), /* uint_var, reduce: Space */
			reduce(93), /* double_var, reduce: Space */
			reduce(93), /* string_var, reduce: Space */
			reduce(93), /* bytes_var, reduce: Space */
			reduce(93), /* true, reduce: Space */
			reduce(93), /* false, reduce: Space */
			nil,        /* = */
			reduce(93), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(93), /* #, reduce: Space */
			nil,        /* & */
			nil,        /* | */
			reduce(93), /* [, reduce: Space */
			nil,        /* ] */
			nil,        /* : */
			reduce(93), /* !, reduce: Space */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S82
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(173), /* ( */
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
			shift(174), /* space */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(67), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(67), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(67), /* []bool, reduce: OpenParen */
			reduce(67), /* []int, reduce: OpenParen */
			reduce(67), /* []uint, reduce: OpenParen */
			reduce(67), /* []double, reduce: OpenParen */
			reduce(67), /* []string, reduce: OpenParen */
			reduce(67), /* [][]byte, reduce: OpenParen */
			reduce(67), /* int_lit, reduce: OpenParen */
			reduce(67), /* uint_lit, reduce: OpenParen */
			reduce(67), /* double_lit, reduce: OpenParen */
			reduce(67), /* bytes_lit, reduce: OpenParen */
			reduce(67), /* bool_var, reduce: OpenParen */
			reduce(67), /* int_var, reduce: OpenParen */
			reduce(67), /* uint_var, reduce: OpenParen */
			reduce(67), /* double_var, reduce: OpenParen */
			reduce(67), /* string_var, reduce: OpenParen */
			reduce(67), /* bytes_var, reduce: OpenParen */
			reduce(67), /* true, reduce: OpenParen */
			reduce(67), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(67), /* ), reduce: OpenParen */
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
			reduce(67), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S85
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(94), /* (, reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S86
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(201), /* : */
			nil,        /* ! */
			shift(202), /* space */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(21), /* id */
			shift(24), /* AnyName */
			shift(25), /* Name */
			shift(27), /* string_lit */
			shift(28), /* AnyNameExcept */
			shift(29), /* NameChoice */
			shift(30), /* Empty */
			shift(31), /* EmptySet */
			shift(34), /* ZeroOrMore */
			shift(41), /* []bool */
			shift(42), /* []int */
			shift(43), /* []uint */
			shift(44), /* []double */
			shift(45), /* []string */
			shift(46), /* [][]byte */
			shift(49), /* int_lit */
			shift(50), /* uint_lit */
			shift(51), /* double_lit */
			shift(52), /* bytes_lit */
			shift(53), /* bool_var */
			shift(54), /* int_var */
			shift(55), /* uint_var */
			shift(56), /* double_var */
			shift(57), /* string_var */
			shift(58), /* bytes_var */
			shift(59), /* true */
			shift(60), /* false */
			nil,       /* = */
			shift(61), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(62), /* # */
			nil,       /* & */
			nil,       /* | */
			shift(63), /* [ */
			nil,       /* ] */
			nil,       /* : */
			shift(64), /* ! */
			shift(65), /* space */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* id, reduce: Colon */
			reduce(89), /* AnyName, reduce: Colon */
			reduce(89), /* Name, reduce: Colon */
			reduce(89), /* string_lit, reduce: Colon */
			reduce(89), /* AnyNameExcept, reduce: Colon */
			reduce(89), /* NameChoice, reduce: Colon */
			reduce(89), /* Empty, reduce: Colon */
			reduce(89), /* EmptySet, reduce: Colon */
			reduce(89), /* ZeroOrMore, reduce: Colon */
			reduce(89), /* []bool, reduce: Colon */
			reduce(89), /* []int, reduce: Colon */
			reduce(89), /* []uint, reduce: Colon */
			reduce(89), /* []double, reduce: Colon */
			reduce(89), /* []string, reduce: Colon */
			reduce(89), /* [][]byte, reduce: Colon */
			reduce(89), /* int_lit, reduce: Colon */
			reduce(89), /* uint_lit, reduce: Colon */
			reduce(89), /* double_lit, reduce: Colon */
			reduce(89), /* bytes_lit, reduce: Colon */
			reduce(89), /* bool_var, reduce: Colon */
			reduce(89), /* int_var, reduce: Colon */
			reduce(89), /* uint_var, reduce: Colon */
			reduce(89), /* double_var, reduce: Colon */
			reduce(89), /* string_var, reduce: Colon */
			reduce(89), /* bytes_var, reduce: Colon */
			reduce(89), /* true, reduce: Colon */
			reduce(89), /* false, reduce: Colon */
			nil,        /* = */
			reduce(89), /* (, reduce: Colon */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* #, reduce: Colon */
			nil,        /* & */
			nil,        /* | */
			reduce(89), /* [, reduce: Colon */
			nil,        /* ] */
			nil,        /* : */
			reduce(89), /* !, reduce: Colon */
			reduce(89), /* space, reduce: Colon */

		},
	},
	actionRow{ // S89
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* :, reduce: Space */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S90
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(204), /* ( */
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
			shift(174), /* space */

		},
	},
	actionRow{ // S91
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(206), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S92
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(67), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(67), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(208), /* id */
			shift(68),  /* AnyName */
			shift(69),  /* Name */
			shift(98),  /* string_lit */
			shift(70),  /* AnyNameExcept */
			shift(71),  /* NameChoice */
			shift(209), /* Empty */
			shift(210), /* EmptySet */
			shift(211), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(77),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(79),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(80),  /* ! */
			shift(81),  /* space */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S95
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(218), /* & */
			shift(219), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(220), /* space */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(88), /* : */
			nil,       /* ! */
			shift(89), /* space */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S98
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(55), /* &, reduce: Terminal */
			reduce(55), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S99
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(18), /* &, reduce: Pattern */
			reduce(18), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S100
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(20), /* &, reduce: Pattern */
			reduce(20), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S101
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(22), /* &, reduce: Pattern */
			reduce(22), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(225), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S106
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(30), /* &, reduce: Expr */
			reduce(30), /* |, reduce: Expr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S107
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(31), /* &, reduce: Expr */
			reduce(31), /* |, reduce: Expr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(32), /* &, reduce: Expr */
			reduce(32), /* |, reduce: Expr */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S109
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S110
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(49), /* &, reduce: SpaceTerminal */
			reduce(49), /* |, reduce: SpaceTerminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S111
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(51), /* &, reduce: Terminal */
			reduce(51), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S112
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(52), /* &, reduce: Terminal */
			reduce(52), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S113
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(53), /* &, reduce: Terminal */
			reduce(53), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S114
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(54), /* &, reduce: Terminal */
			reduce(54), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S115
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(56), /* &, reduce: Terminal */
			reduce(56), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S116
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(57), /* &, reduce: Terminal */
			reduce(57), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S117
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(58), /* &, reduce: Terminal */
			reduce(58), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S118
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(59), /* &, reduce: Terminal */
			reduce(59), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S119
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* &, reduce: Terminal */
			reduce(60), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S120
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* &, reduce: Terminal */
			reduce(61), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S121
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* &, reduce: Terminal */
			reduce(62), /* |, reduce: Terminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S122
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(63), /* &, reduce: Bool */
			reduce(63), /* |, reduce: Bool */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S123
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(64), /* &, reduce: Bool */
			reduce(64), /* |, reduce: Bool */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S124
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(228), /* ( */
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
			shift(174), /* space */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(67), /* AnyName, reduce: OpenParen */
			reduce(67), /* Name, reduce: OpenParen */
			nil,        /* string_lit */
			reduce(67), /* AnyNameExcept, reduce: OpenParen */
			reduce(67), /* NameChoice, reduce: OpenParen */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(67), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S128
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(242), /* id */
			shift(68),  /* AnyName */
			shift(69),  /* Name */
			shift(133), /* string_lit */
			shift(70),  /* AnyNameExcept */
			shift(71),  /* NameChoice */
			shift(243), /* Empty */
			shift(244), /* EmptySet */
			shift(245), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(77),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(79),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(80),  /* ! */
			shift(81),  /* space */

		},
	},
	actionRow{ // S129
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S130
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(251), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S131
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(88), /* : */
			nil,       /* ! */
			shift(89), /* space */

		},
	},
	actionRow{ // S132
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S133
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(55), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S134
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S135
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(257), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S141
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(30), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S142
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(31), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S143
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(32), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S144
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S145
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(49), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(51), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S147
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(52), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S148
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(53), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S149
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(54), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S150
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(56), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S151
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(57), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S152
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(58), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(59), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S154
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S155
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S156
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S157
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(63), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S158
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(64), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S159
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(77),  /* ( */
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
			shift(174), /* space */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(28), /* $, reduce: Pattern */
			reduce(28), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S163
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(292), /* { */
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
			shift(293), /* space */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(318), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S165
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(71), /* id, reduce: OpenCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(71), /* string_lit, reduce: OpenCurly */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(71), /* []bool, reduce: OpenCurly */
			reduce(71), /* []int, reduce: OpenCurly */
			reduce(71), /* []uint, reduce: OpenCurly */
			reduce(71), /* []double, reduce: OpenCurly */
			reduce(71), /* []string, reduce: OpenCurly */
			reduce(71), /* [][]byte, reduce: OpenCurly */
			reduce(71), /* int_lit, reduce: OpenCurly */
			reduce(71), /* uint_lit, reduce: OpenCurly */
			reduce(71), /* double_lit, reduce: OpenCurly */
			reduce(71), /* bytes_lit, reduce: OpenCurly */
			reduce(71), /* bool_var, reduce: OpenCurly */
			reduce(71), /* int_var, reduce: OpenCurly */
			reduce(71), /* uint_var, reduce: OpenCurly */
			reduce(71), /* double_var, reduce: OpenCurly */
			reduce(71), /* string_var, reduce: OpenCurly */
			reduce(71), /* bytes_var, reduce: OpenCurly */
			reduce(71), /* true, reduce: OpenCurly */
			reduce(71), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(71), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(71), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S166
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(94), /* {, reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S167
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S168
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(323), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S169
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S170
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S171
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S172
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(318), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S173
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(68), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(68), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(68), /* []bool, reduce: OpenParen */
			reduce(68), /* []int, reduce: OpenParen */
			reduce(68), /* []uint, reduce: OpenParen */
			reduce(68), /* []double, reduce: OpenParen */
			reduce(68), /* []string, reduce: OpenParen */
			reduce(68), /* [][]byte, reduce: OpenParen */
			reduce(68), /* int_lit, reduce: OpenParen */
			reduce(68), /* uint_lit, reduce: OpenParen */
			reduce(68), /* double_lit, reduce: OpenParen */
			reduce(68), /* bytes_lit, reduce: OpenParen */
			reduce(68), /* bool_var, reduce: OpenParen */
			reduce(68), /* int_var, reduce: OpenParen */
			reduce(68), /* uint_var, reduce: OpenParen */
			reduce(68), /* double_var, reduce: OpenParen */
			reduce(68), /* string_var, reduce: OpenParen */
			reduce(68), /* bytes_var, reduce: OpenParen */
			reduce(68), /* true, reduce: OpenParen */
			reduce(68), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(68), /* ), reduce: OpenParen */
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
			reduce(68), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S174
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(93), /* (, reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S175
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(332), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S176
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S177
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(55), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(55), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S178
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(36), /* $, reduce: Function */
			reduce(36), /* id, reduce: Function */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S179
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(41), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(41), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(41), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S180
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(30), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(30), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S181
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(31), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(31), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S182
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(32), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(32), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S183
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S184
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S185
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(49), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S186
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(51), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(51), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S187
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(52), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(52), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S188
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(53), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(53), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S189
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(54), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(54), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S190
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(56), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(56), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S191
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S192
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S193
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S194
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S195
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S196
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S197
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(63), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(63), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S198
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(64), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(64), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S199
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(69), /* $, reduce: CloseParen */
			reduce(69), /* id, reduce: CloseParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S200
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(94), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(94), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(94), /* []bool, reduce: Space */
			reduce(94), /* []int, reduce: Space */
			reduce(94), /* []uint, reduce: Space */
			reduce(94), /* []double, reduce: Space */
			reduce(94), /* []string, reduce: Space */
			reduce(94), /* [][]byte, reduce: Space */
			reduce(94), /* int_lit, reduce: Space */
			reduce(94), /* uint_lit, reduce: Space */
			reduce(94), /* double_lit, reduce: Space */
			reduce(94), /* bytes_lit, reduce: Space */
			reduce(94), /* bool_var, reduce: Space */
			reduce(94), /* int_var, reduce: Space */
			reduce(94), /* uint_var, reduce: Space */
			reduce(94), /* double_var, reduce: Space */
			reduce(94), /* string_var, reduce: Space */
			reduce(94), /* bytes_var, reduce: Space */
			reduce(94), /* true, reduce: Space */
			reduce(94), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(94), /* ), reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S201
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(90), /* id, reduce: Colon */
			reduce(90), /* AnyName, reduce: Colon */
			reduce(90), /* Name, reduce: Colon */
			reduce(90), /* string_lit, reduce: Colon */
			reduce(90), /* AnyNameExcept, reduce: Colon */
			reduce(90), /* NameChoice, reduce: Colon */
			reduce(90), /* Empty, reduce: Colon */
			reduce(90), /* EmptySet, reduce: Colon */
			reduce(90), /* ZeroOrMore, reduce: Colon */
			reduce(90), /* []bool, reduce: Colon */
			reduce(90), /* []int, reduce: Colon */
			reduce(90), /* []uint, reduce: Colon */
			reduce(90), /* []double, reduce: Colon */
			reduce(90), /* []string, reduce: Colon */
			reduce(90), /* [][]byte, reduce: Colon */
			reduce(90), /* int_lit, reduce: Colon */
			reduce(90), /* uint_lit, reduce: Colon */
			reduce(90), /* double_lit, reduce: Colon */
			reduce(90), /* bytes_lit, reduce: Colon */
			reduce(90), /* bool_var, reduce: Colon */
			reduce(90), /* int_var, reduce: Colon */
			reduce(90), /* uint_var, reduce: Colon */
			reduce(90), /* double_var, reduce: Colon */
			reduce(90), /* string_var, reduce: Colon */
			reduce(90), /* bytes_var, reduce: Colon */
			reduce(90), /* true, reduce: Colon */
			reduce(90), /* false, reduce: Colon */
			nil,        /* = */
			reduce(90), /* (, reduce: Colon */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(90), /* #, reduce: Colon */
			nil,        /* & */
			nil,        /* | */
			reduce(90), /* [, reduce: Colon */
			nil,        /* ] */
			nil,        /* : */
			reduce(90), /* !, reduce: Colon */
			reduce(90), /* space, reduce: Colon */

		},
	},
	actionRow{ // S202
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* :, reduce: Space */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S203
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(21), /* $, reduce: Pattern */
			reduce(21), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S204
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(68), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(68), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S205
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(341), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S206
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S207
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(94), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S208
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S209
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(17), /* &, reduce: Pattern */
			reduce(17), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S210
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(19), /* &, reduce: Pattern */
			reduce(19), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S211
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S212
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S213
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(50), /* &, reduce: SpaceTerminal */
			reduce(50), /* |, reduce: SpaceTerminal */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S214
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S215
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(354), /* & */
			shift(355), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(356), /* space */

		},
	},
	actionRow{ // S216
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S217
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S218
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(81), /* id, reduce: Ampersand */
			reduce(81), /* AnyName, reduce: Ampersand */
			reduce(81), /* Name, reduce: Ampersand */
			reduce(81), /* string_lit, reduce: Ampersand */
			reduce(81), /* AnyNameExcept, reduce: Ampersand */
			reduce(81), /* NameChoice, reduce: Ampersand */
			reduce(81), /* Empty, reduce: Ampersand */
			reduce(81), /* EmptySet, reduce: Ampersand */
			reduce(81), /* ZeroOrMore, reduce: Ampersand */
			reduce(81), /* []bool, reduce: Ampersand */
			reduce(81), /* []int, reduce: Ampersand */
			reduce(81), /* []uint, reduce: Ampersand */
			reduce(81), /* []double, reduce: Ampersand */
			reduce(81), /* []string, reduce: Ampersand */
			reduce(81), /* [][]byte, reduce: Ampersand */
			reduce(81), /* int_lit, reduce: Ampersand */
			reduce(81), /* uint_lit, reduce: Ampersand */
			reduce(81), /* double_lit, reduce: Ampersand */
			reduce(81), /* bytes_lit, reduce: Ampersand */
			reduce(81), /* bool_var, reduce: Ampersand */
			reduce(81), /* int_var, reduce: Ampersand */
			reduce(81), /* uint_var, reduce: Ampersand */
			reduce(81), /* double_var, reduce: Ampersand */
			reduce(81), /* string_var, reduce: Ampersand */
			reduce(81), /* bytes_var, reduce: Ampersand */
			reduce(81), /* true, reduce: Ampersand */
			reduce(81), /* false, reduce: Ampersand */
			nil,        /* = */
			reduce(81), /* (, reduce: Ampersand */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(81), /* #, reduce: Ampersand */
			nil,        /* & */
			nil,        /* | */
			reduce(81), /* [, reduce: Ampersand */
			nil,        /* ] */
			nil,        /* : */
			reduce(81), /* !, reduce: Ampersand */
			reduce(81), /* space, reduce: Ampersand */

		},
	},
	actionRow{ // S219
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(83), /* id, reduce: Pipe */
			reduce(83), /* AnyName, reduce: Pipe */
			reduce(83), /* Name, reduce: Pipe */
			reduce(83), /* string_lit, reduce: Pipe */
			reduce(83), /* AnyNameExcept, reduce: Pipe */
			reduce(83), /* NameChoice, reduce: Pipe */
			reduce(83), /* Empty, reduce: Pipe */
			reduce(83), /* EmptySet, reduce: Pipe */
			reduce(83), /* ZeroOrMore, reduce: Pipe */
			reduce(83), /* []bool, reduce: Pipe */
			reduce(83), /* []int, reduce: Pipe */
			reduce(83), /* []uint, reduce: Pipe */
			reduce(83), /* []double, reduce: Pipe */
			reduce(83), /* []string, reduce: Pipe */
			reduce(83), /* [][]byte, reduce: Pipe */
			reduce(83), /* int_lit, reduce: Pipe */
			reduce(83), /* uint_lit, reduce: Pipe */
			reduce(83), /* double_lit, reduce: Pipe */
			reduce(83), /* bytes_lit, reduce: Pipe */
			reduce(83), /* bool_var, reduce: Pipe */
			reduce(83), /* int_var, reduce: Pipe */
			reduce(83), /* uint_var, reduce: Pipe */
			reduce(83), /* double_var, reduce: Pipe */
			reduce(83), /* string_var, reduce: Pipe */
			reduce(83), /* bytes_var, reduce: Pipe */
			reduce(83), /* true, reduce: Pipe */
			reduce(83), /* false, reduce: Pipe */
			nil,        /* = */
			reduce(83), /* (, reduce: Pipe */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(83), /* #, reduce: Pipe */
			nil,        /* & */
			nil,        /* | */
			reduce(83), /* [, reduce: Pipe */
			nil,        /* ] */
			nil,        /* : */
			reduce(83), /* !, reduce: Pipe */
			reduce(83), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S220
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* &, reduce: Space */
			reduce(94), /* |, reduce: Space */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S221
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S222
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(218), /* & */
			shift(219), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(220), /* space */

		},
	},
	actionRow{ // S223
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(251), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S224
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S225
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(28), /* &, reduce: Pattern */
			reduce(28), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S226
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S227
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(368), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S228
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(68), /* AnyName, reduce: OpenParen */
			reduce(68), /* Name, reduce: OpenParen */
			nil,        /* string_lit */
			reduce(68), /* AnyNameExcept, reduce: OpenParen */
			reduce(68), /* NameChoice, reduce: OpenParen */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(68), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S229
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(369), /* AnyName */
			shift(370), /* Name */
			nil,        /* string_lit */
			shift(371), /* AnyNameExcept */
			shift(372), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(373), /* space */

		},
	},
	actionRow{ // S230
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S231
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S232
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S233
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S234
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S235
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(94), /* AnyName, reduce: Space */
			reduce(94), /* Name, reduce: Space */
			nil,        /* string_lit */
			reduce(94), /* AnyNameExcept, reduce: Space */
			reduce(94), /* NameChoice, reduce: Space */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S236
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(378), /* AnyName */
			shift(379), /* Name */
			nil,        /* string_lit */
			shift(380), /* AnyNameExcept */
			shift(381), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(373), /* space */

		},
	},
	actionRow{ // S237
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S238
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(8), /* ,, reduce: NameExpr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(8), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S239
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S240
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S241
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S242
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S243
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S244
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S245
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S246
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S247
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(50), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S248
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S249
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(395), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(396), /* space */

		},
	},
	actionRow{ // S250
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S251
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(75), /* id, reduce: Comma */
			reduce(75), /* AnyName, reduce: Comma */
			reduce(75), /* Name, reduce: Comma */
			reduce(75), /* string_lit, reduce: Comma */
			reduce(75), /* AnyNameExcept, reduce: Comma */
			reduce(75), /* NameChoice, reduce: Comma */
			reduce(75), /* Empty, reduce: Comma */
			reduce(75), /* EmptySet, reduce: Comma */
			reduce(75), /* ZeroOrMore, reduce: Comma */
			reduce(75), /* []bool, reduce: Comma */
			reduce(75), /* []int, reduce: Comma */
			reduce(75), /* []uint, reduce: Comma */
			reduce(75), /* []double, reduce: Comma */
			reduce(75), /* []string, reduce: Comma */
			reduce(75), /* [][]byte, reduce: Comma */
			reduce(75), /* int_lit, reduce: Comma */
			reduce(75), /* uint_lit, reduce: Comma */
			reduce(75), /* double_lit, reduce: Comma */
			reduce(75), /* bytes_lit, reduce: Comma */
			reduce(75), /* bool_var, reduce: Comma */
			reduce(75), /* int_var, reduce: Comma */
			reduce(75), /* uint_var, reduce: Comma */
			reduce(75), /* double_var, reduce: Comma */
			reduce(75), /* string_var, reduce: Comma */
			reduce(75), /* bytes_var, reduce: Comma */
			reduce(75), /* true, reduce: Comma */
			reduce(75), /* false, reduce: Comma */
			nil,        /* = */
			reduce(75), /* (, reduce: Comma */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(75), /* #, reduce: Comma */
			nil,        /* & */
			nil,        /* | */
			reduce(75), /* [, reduce: Comma */
			nil,        /* ] */
			nil,        /* : */
			reduce(75), /* !, reduce: Comma */
			reduce(75), /* space, reduce: Comma */

		},
	},
	actionRow{ // S252
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S253
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S254
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(218), /* & */
			shift(219), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(220), /* space */

		},
	},
	actionRow{ // S255
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(251), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S256
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S257
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(28), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S258
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S259
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(437), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S260
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(438), /* id */
			shift(68),  /* AnyName */
			shift(69),  /* Name */
			shift(265), /* string_lit */
			shift(70),  /* AnyNameExcept */
			shift(71),  /* NameChoice */
			shift(439), /* Empty */
			shift(440), /* EmptySet */
			shift(441), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(77),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(79),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(80),  /* ! */
			shift(81),  /* space */

		},
	},
	actionRow{ // S261
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S262
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S263
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(88), /* : */
			nil,       /* ! */
			shift(89), /* space */

		},
	},
	actionRow{ // S264
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S265
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(55), /* ), reduce: Terminal */
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
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S266
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S267
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S268
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S269
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S270
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S271
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(451), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S272
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S273
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(30), /* ), reduce: Expr */
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
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S274
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(31), /* ), reduce: Expr */
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
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S275
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(32), /* ), reduce: Expr */
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
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S276
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S277
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: SpaceTerminal */
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
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S278
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(51), /* ), reduce: Terminal */
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
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S279
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(52), /* ), reduce: Terminal */
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
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S280
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(53), /* ), reduce: Terminal */
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
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S281
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(54), /* ), reduce: Terminal */
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
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S282
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(56), /* ), reduce: Terminal */
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
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S283
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S284
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S285
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S286
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S287
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S288
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S289
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(63), /* ), reduce: Bool */
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
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S290
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(64), /* ), reduce: Bool */
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
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S291
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S292
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(72), /* id, reduce: OpenCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(72), /* string_lit, reduce: OpenCurly */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(72), /* []bool, reduce: OpenCurly */
			reduce(72), /* []int, reduce: OpenCurly */
			reduce(72), /* []uint, reduce: OpenCurly */
			reduce(72), /* []double, reduce: OpenCurly */
			reduce(72), /* []string, reduce: OpenCurly */
			reduce(72), /* [][]byte, reduce: OpenCurly */
			reduce(72), /* int_lit, reduce: OpenCurly */
			reduce(72), /* uint_lit, reduce: OpenCurly */
			reduce(72), /* double_lit, reduce: OpenCurly */
			reduce(72), /* bytes_lit, reduce: OpenCurly */
			reduce(72), /* bool_var, reduce: OpenCurly */
			reduce(72), /* int_var, reduce: OpenCurly */
			reduce(72), /* uint_var, reduce: OpenCurly */
			reduce(72), /* double_var, reduce: OpenCurly */
			reduce(72), /* string_var, reduce: OpenCurly */
			reduce(72), /* bytes_var, reduce: OpenCurly */
			reduce(72), /* true, reduce: OpenCurly */
			reduce(72), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(72), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(72), /* space, reduce: OpenCurly */

		},
	},
	actionRow{ // S293
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(93), /* {, reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S294
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(458), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S295
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S296
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(55), /* }, reduce: Terminal */
			reduce(55), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S297
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(41), /* }, reduce: Exprs */
			reduce(41), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(41), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S298
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(30), /* }, reduce: Expr */
			reduce(30), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S299
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(31), /* }, reduce: Expr */
			reduce(31), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S300
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(32), /* }, reduce: Expr */
			reduce(32), /* ,, reduce: Expr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S301
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(318), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S302
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S303
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(40), /* $, reduce: List */
			reduce(40), /* id, reduce: List */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S304
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(49), /* }, reduce: SpaceTerminal */
			reduce(49), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S305
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(51), /* }, reduce: Terminal */
			reduce(51), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S306
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(52), /* }, reduce: Terminal */
			reduce(52), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S307
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(53), /* }, reduce: Terminal */
			reduce(53), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S308
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(54), /* }, reduce: Terminal */
			reduce(54), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S309
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(56), /* }, reduce: Terminal */
			reduce(56), /* ,, reduce: Terminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S310
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S311
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S312
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S313
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S314
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S315
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S316
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(63), /* }, reduce: Bool */
			reduce(63), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S317
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(64), /* }, reduce: Bool */
			reduce(64), /* ,, reduce: Bool */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S318
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(73), /* $, reduce: CloseCurly */
			reduce(73), /* id, reduce: CloseCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S319
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(94), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(94), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(94), /* []bool, reduce: Space */
			reduce(94), /* []int, reduce: Space */
			reduce(94), /* []uint, reduce: Space */
			reduce(94), /* []double, reduce: Space */
			reduce(94), /* []string, reduce: Space */
			reduce(94), /* [][]byte, reduce: Space */
			reduce(94), /* int_lit, reduce: Space */
			reduce(94), /* uint_lit, reduce: Space */
			reduce(94), /* double_lit, reduce: Space */
			reduce(94), /* bytes_lit, reduce: Space */
			reduce(94), /* bool_var, reduce: Space */
			reduce(94), /* int_var, reduce: Space */
			reduce(94), /* uint_var, reduce: Space */
			reduce(94), /* double_var, reduce: Space */
			reduce(94), /* string_var, reduce: Space */
			reduce(94), /* bytes_var, reduce: Space */
			reduce(94), /* true, reduce: Space */
			reduce(94), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(94), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S320
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(34), /* $, reduce: Function */
			reduce(34), /* id, reduce: Function */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S321
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S322
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(467), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S323
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S324
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S325
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S326
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S327
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(318), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S328
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(39), /* $, reduce: List */
			reduce(39), /* id, reduce: List */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S329
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S330
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S331
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(50), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(50), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S332
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(70), /* $, reduce: CloseParen */
			reduce(70), /* id, reduce: CloseParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S333
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(93), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(93), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(93), /* []bool, reduce: Space */
			reduce(93), /* []int, reduce: Space */
			reduce(93), /* []uint, reduce: Space */
			reduce(93), /* []double, reduce: Space */
			reduce(93), /* []string, reduce: Space */
			reduce(93), /* [][]byte, reduce: Space */
			reduce(93), /* int_lit, reduce: Space */
			reduce(93), /* uint_lit, reduce: Space */
			reduce(93), /* double_lit, reduce: Space */
			reduce(93), /* bytes_lit, reduce: Space */
			reduce(93), /* bool_var, reduce: Space */
			reduce(93), /* int_var, reduce: Space */
			reduce(93), /* uint_var, reduce: Space */
			reduce(93), /* double_var, reduce: Space */
			reduce(93), /* string_var, reduce: Space */
			reduce(93), /* bytes_var, reduce: Space */
			reduce(93), /* true, reduce: Space */
			reduce(93), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(93), /* ), reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S334
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(478), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S335
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(332), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S336
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(35), /* $, reduce: Function */
			reduce(35), /* id, reduce: Function */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S337
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
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
			shift(483), /* space */

		},
	},
	actionRow{ // S338
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(75), /* id, reduce: Comma */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(75), /* string_lit, reduce: Comma */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(75), /* []bool, reduce: Comma */
			reduce(75), /* []int, reduce: Comma */
			reduce(75), /* []uint, reduce: Comma */
			reduce(75), /* []double, reduce: Comma */
			reduce(75), /* []string, reduce: Comma */
			reduce(75), /* [][]byte, reduce: Comma */
			reduce(75), /* int_lit, reduce: Comma */
			reduce(75), /* uint_lit, reduce: Comma */
			reduce(75), /* double_lit, reduce: Comma */
			reduce(75), /* bytes_lit, reduce: Comma */
			reduce(75), /* bool_var, reduce: Comma */
			reduce(75), /* int_var, reduce: Comma */
			reduce(75), /* uint_var, reduce: Comma */
			reduce(75), /* double_var, reduce: Comma */
			reduce(75), /* string_var, reduce: Comma */
			reduce(75), /* bytes_var, reduce: Comma */
			reduce(75), /* true, reduce: Comma */
			reduce(75), /* false, reduce: Comma */
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
			reduce(75), /* space, reduce: Comma */

		},
	},
	actionRow{ // S339
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(94), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(94), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S340
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(487), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S341
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S342
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(93), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S343
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(489), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S344
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(12), /* :, reduce: NameExpr */
			nil,        /* ! */
			reduce(12), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S345
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(69), /* :, reduce: CloseParen */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S346
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(94), /* ), reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S347
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S348
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S349
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(368), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S350
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(496), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S351
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(36), /* &, reduce: Function */
			reduce(36), /* |, reduce: Function */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S352
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S353
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(69), /* &, reduce: CloseParen */
			reduce(69), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S354
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(82), /* id, reduce: Ampersand */
			reduce(82), /* AnyName, reduce: Ampersand */
			reduce(82), /* Name, reduce: Ampersand */
			reduce(82), /* string_lit, reduce: Ampersand */
			reduce(82), /* AnyNameExcept, reduce: Ampersand */
			reduce(82), /* NameChoice, reduce: Ampersand */
			reduce(82), /* Empty, reduce: Ampersand */
			reduce(82), /* EmptySet, reduce: Ampersand */
			reduce(82), /* ZeroOrMore, reduce: Ampersand */
			reduce(82), /* []bool, reduce: Ampersand */
			reduce(82), /* []int, reduce: Ampersand */
			reduce(82), /* []uint, reduce: Ampersand */
			reduce(82), /* []double, reduce: Ampersand */
			reduce(82), /* []string, reduce: Ampersand */
			reduce(82), /* [][]byte, reduce: Ampersand */
			reduce(82), /* int_lit, reduce: Ampersand */
			reduce(82), /* uint_lit, reduce: Ampersand */
			reduce(82), /* double_lit, reduce: Ampersand */
			reduce(82), /* bytes_lit, reduce: Ampersand */
			reduce(82), /* bool_var, reduce: Ampersand */
			reduce(82), /* int_var, reduce: Ampersand */
			reduce(82), /* uint_var, reduce: Ampersand */
			reduce(82), /* double_var, reduce: Ampersand */
			reduce(82), /* string_var, reduce: Ampersand */
			reduce(82), /* bytes_var, reduce: Ampersand */
			reduce(82), /* true, reduce: Ampersand */
			reduce(82), /* false, reduce: Ampersand */
			nil,        /* = */
			reduce(82), /* (, reduce: Ampersand */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(82), /* #, reduce: Ampersand */
			nil,        /* & */
			nil,        /* | */
			reduce(82), /* [, reduce: Ampersand */
			nil,        /* ] */
			nil,        /* : */
			reduce(82), /* !, reduce: Ampersand */
			reduce(82), /* space, reduce: Ampersand */

		},
	},
	actionRow{ // S355
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(84), /* id, reduce: Pipe */
			reduce(84), /* AnyName, reduce: Pipe */
			reduce(84), /* Name, reduce: Pipe */
			reduce(84), /* string_lit, reduce: Pipe */
			reduce(84), /* AnyNameExcept, reduce: Pipe */
			reduce(84), /* NameChoice, reduce: Pipe */
			reduce(84), /* Empty, reduce: Pipe */
			reduce(84), /* EmptySet, reduce: Pipe */
			reduce(84), /* ZeroOrMore, reduce: Pipe */
			reduce(84), /* []bool, reduce: Pipe */
			reduce(84), /* []int, reduce: Pipe */
			reduce(84), /* []uint, reduce: Pipe */
			reduce(84), /* []double, reduce: Pipe */
			reduce(84), /* []string, reduce: Pipe */
			reduce(84), /* [][]byte, reduce: Pipe */
			reduce(84), /* int_lit, reduce: Pipe */
			reduce(84), /* uint_lit, reduce: Pipe */
			reduce(84), /* double_lit, reduce: Pipe */
			reduce(84), /* bytes_lit, reduce: Pipe */
			reduce(84), /* bool_var, reduce: Pipe */
			reduce(84), /* int_var, reduce: Pipe */
			reduce(84), /* uint_var, reduce: Pipe */
			reduce(84), /* double_var, reduce: Pipe */
			reduce(84), /* string_var, reduce: Pipe */
			reduce(84), /* bytes_var, reduce: Pipe */
			reduce(84), /* true, reduce: Pipe */
			reduce(84), /* false, reduce: Pipe */
			nil,        /* = */
			reduce(84), /* (, reduce: Pipe */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(84), /* #, reduce: Pipe */
			nil,        /* & */
			nil,        /* | */
			reduce(84), /* [, reduce: Pipe */
			nil,        /* ] */
			nil,        /* : */
			reduce(84), /* !, reduce: Pipe */
			reduce(84), /* space, reduce: Pipe */

		},
	},
	actionRow{ // S356
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* &, reduce: Space */
			reduce(93), /* |, reduce: Space */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S357
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S358
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(199), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S359
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(21), /* &, reduce: Pattern */
			reduce(21), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S360
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S361
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S362
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S363
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S364
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S365
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(507), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S366
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(368), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S367
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(40), /* &, reduce: List */
			reduce(40), /* |, reduce: List */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S368
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(73), /* &, reduce: CloseCurly */
			reduce(73), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S369
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S370
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S371
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S372
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S373
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(93), /* AnyName, reduce: Space */
			reduce(93), /* Name, reduce: Space */
			nil,        /* string_lit */
			reduce(93), /* AnyNameExcept, reduce: Space */
			reduce(93), /* NameChoice, reduce: Space */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S374
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(14), /* :, reduce: NameExpr */
			nil,        /* ! */
			reduce(14), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S375
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(514), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S376
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S377
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S378
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(7), /* ,, reduce: NameExpr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S379
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(92), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S380
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S381
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(126), /* ( */
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
			shift(85),  /* space */

		},
	},
	actionRow{ // S382
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(520), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(396), /* space */

		},
	},
	actionRow{ // S383
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S384
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(75), /* AnyName, reduce: Comma */
			reduce(75), /* Name, reduce: Comma */
			nil,        /* string_lit */
			reduce(75), /* AnyNameExcept, reduce: Comma */
			reduce(75), /* NameChoice, reduce: Comma */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(75), /* space, reduce: Comma */

		},
	},
	actionRow{ // S385
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(523), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S386
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S387
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S388
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S389
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S390
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(437), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S391
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(531), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S392
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(36), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S393
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S394
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(69), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S395
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(76), /* id, reduce: Comma */
			reduce(76), /* AnyName, reduce: Comma */
			reduce(76), /* Name, reduce: Comma */
			reduce(76), /* string_lit, reduce: Comma */
			reduce(76), /* AnyNameExcept, reduce: Comma */
			reduce(76), /* NameChoice, reduce: Comma */
			reduce(76), /* Empty, reduce: Comma */
			reduce(76), /* EmptySet, reduce: Comma */
			reduce(76), /* ZeroOrMore, reduce: Comma */
			reduce(76), /* []bool, reduce: Comma */
			reduce(76), /* []int, reduce: Comma */
			reduce(76), /* []uint, reduce: Comma */
			reduce(76), /* []double, reduce: Comma */
			reduce(76), /* []string, reduce: Comma */
			reduce(76), /* [][]byte, reduce: Comma */
			reduce(76), /* int_lit, reduce: Comma */
			reduce(76), /* uint_lit, reduce: Comma */
			reduce(76), /* double_lit, reduce: Comma */
			reduce(76), /* bytes_lit, reduce: Comma */
			reduce(76), /* bool_var, reduce: Comma */
			reduce(76), /* int_var, reduce: Comma */
			reduce(76), /* uint_var, reduce: Comma */
			reduce(76), /* double_var, reduce: Comma */
			reduce(76), /* string_var, reduce: Comma */
			reduce(76), /* bytes_var, reduce: Comma */
			reduce(76), /* true, reduce: Comma */
			reduce(76), /* false, reduce: Comma */
			nil,        /* = */
			reduce(76), /* (, reduce: Comma */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* #, reduce: Comma */
			nil,        /* & */
			nil,        /* | */
			reduce(76), /* [, reduce: Comma */
			nil,        /* ] */
			nil,        /* : */
			reduce(76), /* !, reduce: Comma */
			reduce(76), /* space, reduce: Comma */

		},
	},
	actionRow{ // S396
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S397
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(534), /* id */
			shift(68),  /* AnyName */
			shift(69),  /* Name */
			shift(402), /* string_lit */
			shift(70),  /* AnyNameExcept */
			shift(71),  /* NameChoice */
			shift(535), /* Empty */
			shift(536), /* EmptySet */
			shift(537), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(77),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(79),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(80),  /* ! */
			shift(81),  /* space */

		},
	},
	actionRow{ // S398
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S399
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(543), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(544), /* space */

		},
	},
	actionRow{ // S400
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(88), /* : */
			nil,       /* ! */
			shift(89), /* space */

		},
	},
	actionRow{ // S401
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(94),  /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(98),  /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(99),  /* Empty */
			shift(100), /* EmptySet */
			shift(103), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(112), /* int_lit */
			shift(113), /* uint_lit */
			shift(114), /* double_lit */
			shift(115), /* bytes_lit */
			shift(116), /* bool_var */
			shift(117), /* int_var */
			shift(118), /* uint_var */
			shift(119), /* double_var */
			shift(120), /* string_var */
			shift(121), /* bytes_var */
			shift(122), /* true */
			shift(123), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S402
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(55), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(55), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S403
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(18), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S404
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(20), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S405
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(22), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S406
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(129), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(133), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(134), /* Empty */
			shift(135), /* EmptySet */
			shift(138), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(147), /* int_lit */
			shift(148), /* uint_lit */
			shift(149), /* double_lit */
			shift(150), /* bytes_lit */
			shift(151), /* bool_var */
			shift(152), /* int_var */
			shift(153), /* uint_var */
			shift(154), /* double_var */
			shift(155), /* string_var */
			shift(156), /* bytes_var */
			shift(157), /* true */
			shift(158), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S407
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S408
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(549), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
	actionRow{ // S409
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S410
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(30), /* ], reduce: Expr */
			nil,        /* : */
			nil,        /* ! */
			reduce(30), /* space, reduce: Expr */

		},
	},
	actionRow{ // S411
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(31), /* ], reduce: Expr */
			nil,        /* : */
			nil,        /* ! */
			reduce(31), /* space, reduce: Expr */

		},
	},
	actionRow{ // S412
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(32), /* ], reduce: Expr */
			nil,        /* : */
			nil,        /* ! */
			reduce(32), /* space, reduce: Expr */

		},
	},
	actionRow{ // S413
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S414
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(49), /* ], reduce: SpaceTerminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(49), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S415
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(51), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(51), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S416
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(52), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(52), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S417
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(53), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(53), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S418
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(54), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(54), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S419
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(56), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(56), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S420
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(57), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(57), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S421
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(58), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(58), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S422
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(59), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(59), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S423
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S424
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S425
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* ], reduce: Terminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S426
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(63), /* ], reduce: Bool */
			nil,        /* : */
			nil,        /* ! */
			reduce(63), /* space, reduce: Bool */

		},
	},
	actionRow{ // S427
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(64), /* ], reduce: Bool */
			nil,        /* : */
			nil,        /* ! */
			reduce(64), /* space, reduce: Bool */

		},
	},
	actionRow{ // S428
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S429
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S430
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S431
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S432
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S433
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S434
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(558), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S435
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(437), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S436
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(40), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S437
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(73), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S438
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S439
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S440
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S441
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S442
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S443
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(50), /* ), reduce: SpaceTerminal */
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
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S444
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S445
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(332), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S446
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(27), /* $, reduce: Pattern */
			reduce(27), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S447
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S448
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(218), /* & */
			shift(219), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(220), /* space */

		},
	},
	actionRow{ // S449
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(251), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S450
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S451
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(28), /* ), reduce: Pattern */
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
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S452
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S453
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(577), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S454
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(29), /* $, reduce: Pattern */
			reduce(29), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S455
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S456
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S457
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(50), /* }, reduce: SpaceTerminal */
			reduce(50), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S458
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(74), /* $, reduce: CloseCurly */
			reduce(74), /* id, reduce: CloseCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S459
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(93), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(93), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(93), /* []bool, reduce: Space */
			reduce(93), /* []int, reduce: Space */
			reduce(93), /* []uint, reduce: Space */
			reduce(93), /* []double, reduce: Space */
			reduce(93), /* []string, reduce: Space */
			reduce(93), /* [][]byte, reduce: Space */
			reduce(93), /* int_lit, reduce: Space */
			reduce(93), /* uint_lit, reduce: Space */
			reduce(93), /* double_lit, reduce: Space */
			reduce(93), /* bytes_lit, reduce: Space */
			reduce(93), /* bool_var, reduce: Space */
			reduce(93), /* int_var, reduce: Space */
			reduce(93), /* uint_var, reduce: Space */
			reduce(93), /* double_var, reduce: Space */
			reduce(93), /* string_var, reduce: Space */
			reduce(93), /* bytes_var, reduce: Space */
			reduce(93), /* true, reduce: Space */
			reduce(93), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(93), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S460
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(583), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S461
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(458), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S462
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
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
			shift(483), /* space */

		},
	},
	actionRow{ // S463
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(38), /* $, reduce: List */
			reduce(38), /* id, reduce: List */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S464
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(94), /* }, reduce: Space */
			reduce(94), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S465
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(590), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S466
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(33), /* $, reduce: Function */
			reduce(33), /* id, reduce: Function */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S467
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S468
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S469
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(13), /* :, reduce: NameExpr */
			nil,        /* ! */
			reduce(13), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S470
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S471
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(26), /* $, reduce: Pattern */
			reduce(26), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S472
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(37), /* $, reduce: List */
			reduce(37), /* id, reduce: List */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S473
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(478), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S474
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(487), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S475
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(597), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S476
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(36), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(36), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S477
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(478), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S478
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(69), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(69), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S479
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(76), /* id, reduce: Comma */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(76), /* string_lit, reduce: Comma */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(76), /* []bool, reduce: Comma */
			reduce(76), /* []int, reduce: Comma */
			reduce(76), /* []uint, reduce: Comma */
			reduce(76), /* []double, reduce: Comma */
			reduce(76), /* []string, reduce: Comma */
			reduce(76), /* [][]byte, reduce: Comma */
			reduce(76), /* int_lit, reduce: Comma */
			reduce(76), /* uint_lit, reduce: Comma */
			reduce(76), /* double_lit, reduce: Comma */
			reduce(76), /* bytes_lit, reduce: Comma */
			reduce(76), /* bool_var, reduce: Comma */
			reduce(76), /* int_var, reduce: Comma */
			reduce(76), /* uint_var, reduce: Comma */
			reduce(76), /* double_var, reduce: Comma */
			reduce(76), /* string_var, reduce: Comma */
			reduce(76), /* bytes_var, reduce: Comma */
			reduce(76), /* true, reduce: Comma */
			reduce(76), /* false, reduce: Comma */
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
			reduce(76), /* space, reduce: Comma */

		},
	},
	actionRow{ // S480
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(93), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(93), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S481
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
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
			shift(600), /* space */

		},
	},
	actionRow{ // S482
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(42), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(42), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(42), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S483
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(94), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(94), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(94), /* []bool, reduce: Space */
			reduce(94), /* []int, reduce: Space */
			reduce(94), /* []uint, reduce: Space */
			reduce(94), /* []double, reduce: Space */
			reduce(94), /* []string, reduce: Space */
			reduce(94), /* [][]byte, reduce: Space */
			reduce(94), /* int_lit, reduce: Space */
			reduce(94), /* uint_lit, reduce: Space */
			reduce(94), /* double_lit, reduce: Space */
			reduce(94), /* bytes_lit, reduce: Space */
			reduce(94), /* bool_var, reduce: Space */
			reduce(94), /* int_var, reduce: Space */
			reduce(94), /* uint_var, reduce: Space */
			reduce(94), /* double_var, reduce: Space */
			reduce(94), /* string_var, reduce: Space */
			reduce(94), /* bytes_var, reduce: Space */
			reduce(94), /* true, reduce: Space */
			reduce(94), /* false, reduce: Space */
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
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S484
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(601), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S485
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(487), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S486
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S487
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(73), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(73), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S488
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S489
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(70), /* :, reduce: CloseParen */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S490
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(93), /* ), reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S491
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(34), /* &, reduce: Function */
			reduce(34), /* |, reduce: Function */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S492
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S493
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S494
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(368), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S495
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(39), /* &, reduce: List */
			reduce(39), /* |, reduce: List */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S496
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(70), /* &, reduce: CloseParen */
			reduce(70), /* |, reduce: CloseParen */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S497
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(496), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S498
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(35), /* &, reduce: Function */
			reduce(35), /* |, reduce: Function */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S499
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(24), /* $, reduce: Pattern */
			reduce(24), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S500
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(25), /* $, reduce: Pattern */
			reduce(25), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S501
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S502
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(353), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S503
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(611), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(544), /* space */

		},
	},
	actionRow{ // S504
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(496), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S505
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(27), /* &, reduce: Pattern */
			reduce(27), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S506
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(29), /* &, reduce: Pattern */
			reduce(29), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S507
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(74), /* &, reduce: CloseCurly */
			reduce(74), /* |, reduce: CloseCurly */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S508
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(507), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S509
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(38), /* &, reduce: List */
			reduce(38), /* |, reduce: List */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S510
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(613), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S511
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S512
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S513
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(616), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S514
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S515
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S516
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S517
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(622), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(207), /* space */

		},
	},
	actionRow{ // S518
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S519
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(238), /* AnyName */
			shift(239), /* Name */
			nil,        /* string_lit */
			shift(240), /* AnyNameExcept */
			shift(241), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S520
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(76), /* AnyName, reduce: Comma */
			reduce(76), /* Name, reduce: Comma */
			nil,        /* string_lit */
			reduce(76), /* AnyNameExcept, reduce: Comma */
			reduce(76), /* NameChoice, reduce: Comma */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(76), /* space, reduce: Comma */

		},
	},
	actionRow{ // S521
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S522
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(626), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S523
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S524
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S525
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S526
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(34), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S527
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S528
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S529
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(437), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S530
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(39), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S531
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(70), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S532
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(531), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S533
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(35), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S534
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(84), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S535
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(17), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S536
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(19), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S537
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			shift(61), /* ( */
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
			shift(85), /* space */

		},
	},
	actionRow{ // S538
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(165), /* { */
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
			shift(166), /* space */

		},
	},
	actionRow{ // S539
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(50), /* ], reduce: SpaceTerminal */
			nil,        /* : */
			nil,        /* ! */
			reduce(50), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S540
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S541
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(640), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(641), /* space */

		},
	},
	actionRow{ // S542
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(23), /* $, reduce: Pattern */
			reduce(23), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S543
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(87), /* $, reduce: CloseBracket */
			reduce(87), /* id, reduce: CloseBracket */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(87), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S544
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(94), /* ], reduce: Space */
			nil,        /* : */
			nil,        /* ! */
			reduce(94), /* space, reduce: Space */

		},
	},
	actionRow{ // S545
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S546
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(218), /* & */
			shift(219), /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(220), /* space */

		},
	},
	actionRow{ // S547
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(251), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S548
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S549
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(28), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S550
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S551
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(651), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S552
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S553
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S554
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(656), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(544), /* space */

		},
	},
	actionRow{ // S555
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(531), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S556
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(27), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S557
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(29), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S558
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(74), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S559
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(558), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S560
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(38), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S561
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S562
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S563
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(577), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S564
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(662), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S565
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(36), /* ), reduce: Function */
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
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S566
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S567
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(69), /* ), reduce: CloseParen */
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
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S568
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S569
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S570
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S571
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S572
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S573
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S574
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(670), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S575
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(577), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S576
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: List */
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
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S577
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(73), /* ), reduce: CloseCurly */
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
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S578
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(583), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S579
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(590), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S580
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(677), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S581
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(36), /* }, reduce: Function */
			reduce(36), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S582
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(583), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S583
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(69), /* }, reduce: CloseParen */
			reduce(69), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S584
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(93), /* }, reduce: Space */
			reduce(93), /* ,, reduce: Space */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S585
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
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
			shift(600), /* space */

		},
	},
	actionRow{ // S586
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(42), /* }, reduce: Exprs */
			reduce(42), /* ,, reduce: Exprs */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(42), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S587
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(680), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S588
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(590), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S589
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(40), /* }, reduce: List */
			reduce(40), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S590
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(73), /* }, reduce: CloseCurly */
			reduce(73), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S591
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S592
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(345), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S593
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(34), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(34), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S594
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(478), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S595
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(487), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S596
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(39), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S597
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(70), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(70), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S598
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(597), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S599
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(35), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(35), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S600
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(93), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(93), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			reduce(93), /* []bool, reduce: Space */
			reduce(93), /* []int, reduce: Space */
			reduce(93), /* []uint, reduce: Space */
			reduce(93), /* []double, reduce: Space */
			reduce(93), /* []string, reduce: Space */
			reduce(93), /* [][]byte, reduce: Space */
			reduce(93), /* int_lit, reduce: Space */
			reduce(93), /* uint_lit, reduce: Space */
			reduce(93), /* double_lit, reduce: Space */
			reduce(93), /* bytes_lit, reduce: Space */
			reduce(93), /* bool_var, reduce: Space */
			reduce(93), /* int_var, reduce: Space */
			reduce(93), /* uint_var, reduce: Space */
			reduce(93), /* double_var, reduce: Space */
			reduce(93), /* string_var, reduce: Space */
			reduce(93), /* bytes_var, reduce: Space */
			reduce(93), /* true, reduce: Space */
			reduce(93), /* false, reduce: Space */
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
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S601
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(74), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(74), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S602
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(601), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S603
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(38), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(38), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S604
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(33), /* &, reduce: Function */
			reduce(33), /* |, reduce: Function */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S605
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(26), /* &, reduce: Pattern */
			reduce(26), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S606
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(37), /* &, reduce: List */
			reduce(37), /* |, reduce: List */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S607
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(24), /* &, reduce: Pattern */
			reduce(24), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S608
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(25), /* &, reduce: Pattern */
			reduce(25), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S609
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(686), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(641), /* space */

		},
	},
	actionRow{ // S610
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(23), /* &, reduce: Pattern */
			reduce(23), /* |, reduce: Pattern */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S611
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(87), /* &, reduce: CloseBracket */
			reduce(87), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(87), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S612
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(687), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S613
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S614
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S615
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S616
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S617
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(662), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S618
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(12), /* ), reduce: NameExpr */
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
			reduce(12), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S619
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(14), /* ), reduce: NameExpr */
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
			reduce(14), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S620
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S621
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(693), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(342), /* space */

		},
	},
	actionRow{ // S622
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S623
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S624
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(384), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(252), /* space */

		},
	},
	actionRow{ // S625
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(16), /* :, reduce: NameExpr */
			nil,        /* ! */
			reduce(16), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S626
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S627
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(12), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(12), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S628
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(14), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(14), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S629
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S630
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(33), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S631
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S632
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(37), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S633
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(176), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(200), /* space */

		},
	},
	actionRow{ // S634
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S635
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(295), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(651), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(319), /* space */

		},
	},
	actionRow{ // S636
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(329), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(177), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(187), /* int_lit */
			shift(188), /* uint_lit */
			shift(189), /* double_lit */
			shift(190), /* bytes_lit */
			shift(191), /* bool_var */
			shift(192), /* int_var */
			shift(193), /* uint_var */
			shift(194), /* double_var */
			shift(195), /* string_var */
			shift(196), /* bytes_var */
			shift(197), /* true */
			shift(198), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(704), /* ) */
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
			shift(333), /* space */

		},
	},
	actionRow{ // S637
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(36), /* ], reduce: Function */
			nil,        /* : */
			nil,        /* ! */
			reduce(36), /* space, reduce: Function */

		},
	},
	actionRow{ // S638
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S639
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(69), /* ], reduce: CloseParen */
			nil,        /* : */
			nil,        /* ! */
			reduce(69), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S640
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(88), /* $, reduce: CloseBracket */
			reduce(88), /* id, reduce: CloseBracket */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(88), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S641
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(93), /* ], reduce: Space */
			nil,        /* : */
			nil,        /* ! */
			reduce(93), /* space, reduce: Space */

		},
	},
	actionRow{ // S642
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(21), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S643
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S644
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(261), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(265), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(266), /* Empty */
			shift(267), /* EmptySet */
			shift(270), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(279), /* int_lit */
			shift(280), /* uint_lit */
			shift(281), /* double_lit */
			shift(282), /* bytes_lit */
			shift(283), /* bool_var */
			shift(284), /* int_var */
			shift(285), /* uint_var */
			shift(286), /* double_var */
			shift(287), /* string_var */
			shift(288), /* bytes_var */
			shift(289), /* true */
			shift(290), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S645
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(398), /* id */
			shift(24),  /* AnyName */
			shift(25),  /* Name */
			shift(402), /* string_lit */
			shift(28),  /* AnyNameExcept */
			shift(29),  /* NameChoice */
			shift(403), /* Empty */
			shift(404), /* EmptySet */
			shift(407), /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(416), /* int_lit */
			shift(417), /* uint_lit */
			shift(418), /* double_lit */
			shift(419), /* bytes_lit */
			shift(420), /* bool_var */
			shift(421), /* int_var */
			shift(422), /* uint_var */
			shift(423), /* double_var */
			shift(424), /* string_var */
			shift(425), /* bytes_var */
			shift(426), /* true */
			shift(427), /* false */
			nil,        /* = */
			shift(61),  /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(62),  /* # */
			nil,        /* & */
			nil,        /* | */
			shift(63),  /* [ */
			nil,        /* ] */
			nil,        /* : */
			shift(64),  /* ! */
			shift(65),  /* space */

		},
	},
	actionRow{ // S646
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S647
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S648
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(455), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(296), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			shift(41),  /* []bool */
			shift(42),  /* []int */
			shift(43),  /* []uint */
			shift(44),  /* []double */
			shift(45),  /* []string */
			shift(46),  /* [][]byte */
			shift(306), /* int_lit */
			shift(307), /* uint_lit */
			shift(308), /* double_lit */
			shift(309), /* bytes_lit */
			shift(310), /* bool_var */
			shift(311), /* int_var */
			shift(312), /* uint_var */
			shift(313), /* double_var */
			shift(314), /* string_var */
			shift(315), /* bytes_var */
			shift(316), /* true */
			shift(317), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(713), /* } */
			nil,        /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(459), /* space */

		},
	},
	actionRow{ // S649
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(651), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S650
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(40), /* ], reduce: List */
			nil,        /* : */
			nil,        /* ! */
			reduce(40), /* space, reduce: List */

		},
	},
	actionRow{ // S651
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(73), /* ], reduce: CloseCurly */
			nil,        /* : */
			nil,        /* ! */
			reduce(73), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S652
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(24), /* ,, reduce: Pattern */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S653
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S654
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(716), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(641), /* space */

		},
	},
	actionRow{ // S655
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S656
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(87), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(87), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S657
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(34), /* ), reduce: Function */
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
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S658
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S659
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S660
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(577), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S661
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: List */
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
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S662
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(70), /* ), reduce: CloseParen */
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
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S663
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(662), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S664
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(35), /* ), reduce: Function */
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
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S665
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S666
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S667
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(724), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(544), /* space */

		},
	},
	actionRow{ // S668
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(27), /* ), reduce: Pattern */
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
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S669
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(29), /* ), reduce: Pattern */
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
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S670
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(74), /* ), reduce: CloseCurly */
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
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S671
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(670), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S672
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(38), /* ), reduce: List */
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
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S673
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(34), /* }, reduce: Function */
			reduce(34), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S674
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(583), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S675
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(590), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S676
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(39), /* }, reduce: List */
			reduce(39), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S677
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(70), /* }, reduce: CloseParen */
			reduce(70), /* ,, reduce: CloseParen */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S678
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(677), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S679
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(35), /* }, reduce: Function */
			reduce(35), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S680
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(74), /* }, reduce: CloseCurly */
			reduce(74), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S681
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(680), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S682
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(38), /* }, reduce: List */
			reduce(38), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S683
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(15), /* :, reduce: NameExpr */
			nil,        /* ! */
			reduce(15), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S684
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(33), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(33), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S685
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(37), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(37), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S686
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(88), /* &, reduce: CloseBracket */
			reduce(88), /* |, reduce: CloseBracket */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(88), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S687
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S688
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S689
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(13), /* ), reduce: NameExpr */
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
			reduce(13), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S690
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S691
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S692
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S693
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S694
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(11), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(11), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S695
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(13), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(13), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S696
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(231), /* AnyName */
			shift(232), /* Name */
			nil,        /* string_lit */
			shift(233), /* AnyNameExcept */
			shift(234), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(235), /* space */

		},
	},
	actionRow{ // S697
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(10), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S698
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S699
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(34), /* ], reduce: Function */
			nil,        /* : */
			nil,        /* ! */
			reduce(34), /* space, reduce: Function */

		},
	},
	actionRow{ // S700
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(339), /* space */

		},
	},
	actionRow{ // S701
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S702
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(651), /* } */
			shift(338), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(464), /* space */

		},
	},
	actionRow{ // S703
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(39), /* ], reduce: List */
			nil,        /* : */
			nil,        /* ! */
			reduce(39), /* space, reduce: List */

		},
	},
	actionRow{ // S704
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(70), /* ], reduce: CloseParen */
			nil,        /* : */
			nil,        /* ! */
			reduce(70), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S705
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(704), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(480), /* space */

		},
	},
	actionRow{ // S706
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(35), /* ], reduce: Function */
			nil,        /* : */
			nil,        /* ! */
			reduce(35), /* space, reduce: Function */

		},
	},
	actionRow{ // S707
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S708
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(639), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S709
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(740), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(544), /* space */

		},
	},
	actionRow{ // S710
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(704), /* ) */
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
			shift(490), /* space */

		},
	},
	actionRow{ // S711
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(27), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S712
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(29), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S713
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(74), /* ], reduce: CloseCurly */
			nil,        /* : */
			nil,        /* ! */
			reduce(74), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S714
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(713), /* } */
			shift(479), /* , */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(584), /* space */

		},
	},
	actionRow{ // S715
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(38), /* ], reduce: List */
			nil,        /* : */
			nil,        /* ! */
			reduce(38), /* space, reduce: List */

		},
	},
	actionRow{ // S716
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(88), /* ,, reduce: CloseBracket */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(88), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S717
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(33), /* ), reduce: Function */
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
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S718
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S719
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(37), /* ), reduce: List */
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
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S720
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(24), /* ), reduce: Pattern */
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
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S721
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S722
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(741), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(641), /* space */

		},
	},
	actionRow{ // S723
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S724
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(87), /* ), reduce: CloseBracket */
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
			reduce(87), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S725
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(33), /* }, reduce: Function */
			reduce(33), /* ,, reduce: Function */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S726
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(37), /* }, reduce: List */
			reduce(37), /* ,, reduce: List */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S727
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S728
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(567), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S729
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(16), /* ), reduce: NameExpr */
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
			reduce(16), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S730
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			nil,       /* AnyName */
			nil,       /* Name */
			nil,       /* string_lit */
			nil,       /* AnyNameExcept */
			nil,       /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* ZeroOrMore */
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
			reduce(9), /* ,, reduce: NameExpr */
			nil,       /* ; */
			nil,       /* # */
			nil,       /* & */
			nil,       /* | */
			nil,       /* [ */
			nil,       /* ] */
			nil,       /* : */
			nil,       /* ! */
			reduce(9), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S731
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(394), /* ) */
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
			shift(346), /* space */

		},
	},
	actionRow{ // S732
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(16), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(16), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S733
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(33), /* ], reduce: Function */
			nil,        /* : */
			nil,        /* ! */
			reduce(33), /* space, reduce: Function */

		},
	},
	actionRow{ // S734
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(26), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S735
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(37), /* ], reduce: List */
			nil,        /* : */
			nil,        /* ! */
			reduce(37), /* space, reduce: List */

		},
	},
	actionRow{ // S736
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(24), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S737
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(25), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S738
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			shift(744), /* ] */
			nil,        /* : */
			nil,        /* ! */
			shift(641), /* space */

		},
	},
	actionRow{ // S739
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(23), /* ], reduce: Pattern */
			nil,        /* : */
			nil,        /* ! */
			reduce(23), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S740
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(87), /* ], reduce: CloseBracket */
			nil,        /* : */
			nil,        /* ! */
			reduce(87), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S741
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(88), /* ), reduce: CloseBracket */
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
			reduce(88), /* space, reduce: CloseBracket */

		},
	},
	actionRow{ // S742
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(15), /* ), reduce: NameExpr */
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
			reduce(15), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S743
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(15), /* ,, reduce: NameExpr */
			nil,        /* ; */
			nil,        /* # */
			nil,        /* & */
			nil,        /* | */
			nil,        /* [ */
			nil,        /* ] */
			nil,        /* : */
			nil,        /* ! */
			reduce(15), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S744
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* ZeroOrMore */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(88), /* ], reduce: CloseBracket */
			nil,        /* : */
			nil,        /* ! */
			reduce(88), /* space, reduce: CloseBracket */

		},
	},
}

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
			nil,      /* TreeNode */
			nil,      /* LeafNode */
			nil,      /* Concat */
			nil,      /* Or */
			nil,      /* And */
			nil,      /* ZeroOrMore */
			nil,      /* Reference */
			nil,      /* Not */
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
			nil,          /* TreeNode */
			nil,          /* LeafNode */
			nil,          /* Concat */
			nil,          /* Or */
			nil,          /* And */
			nil,          /* ZeroOrMore */
			nil,          /* Reference */
			nil,          /* Not */
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(15), /* space */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(4), /* space, reduce: PatternDecls */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(89), /* $, reduce: Space */
			reduce(89), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(15), /* space */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(88), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(19), /* space */

		},
	},
	actionRow{ // S13
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
			shift(22), /* Empty */
			shift(23), /* EmptySet */
			shift(24), /* TreeNode */
			shift(25), /* LeafNode */
			shift(26), /* Concat */
			shift(27), /* Or */
			shift(28), /* And */
			shift(29), /* ZeroOrMore */
			shift(30), /* Reference */
			shift(31), /* Not */
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
			shift(32), /* space */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			reduce(74), /* Empty, reduce: Equal */
			reduce(74), /* EmptySet, reduce: Equal */
			reduce(74), /* TreeNode, reduce: Equal */
			reduce(74), /* LeafNode, reduce: Equal */
			reduce(74), /* Concat, reduce: Equal */
			reduce(74), /* Or, reduce: Equal */
			reduce(74), /* And, reduce: Equal */
			reduce(74), /* ZeroOrMore, reduce: Equal */
			reduce(74), /* Reference, reduce: Equal */
			reduce(74), /* Not, reduce: Equal */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(74), /* space, reduce: Equal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(89), /* =, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(88), /* $, reduce: Space */
			reduce(88), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S17
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
			shift(22), /* Empty */
			shift(23), /* EmptySet */
			shift(24), /* TreeNode */
			shift(25), /* LeafNode */
			shift(26), /* Concat */
			shift(27), /* Or */
			shift(28), /* And */
			shift(29), /* ZeroOrMore */
			shift(30), /* Reference */
			shift(31), /* Not */
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
			shift(32), /* space */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			reduce(75), /* Empty, reduce: Equal */
			reduce(75), /* EmptySet, reduce: Equal */
			reduce(75), /* TreeNode, reduce: Equal */
			reduce(75), /* LeafNode, reduce: Equal */
			reduce(75), /* Concat, reduce: Equal */
			reduce(75), /* Or, reduce: Equal */
			reduce(75), /* And, reduce: Equal */
			reduce(75), /* ZeroOrMore, reduce: Equal */
			reduce(75), /* Reference, reduce: Equal */
			reduce(75), /* Not, reduce: Equal */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(75), /* space, reduce: Equal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			reduce(88), /* =, reduce: Space */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S20
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
			shift(34), /* Empty */
			shift(35), /* EmptySet */
			shift(36), /* TreeNode */
			shift(37), /* LeafNode */
			shift(38), /* Concat */
			shift(39), /* Or */
			shift(40), /* And */
			shift(41), /* ZeroOrMore */
			shift(42), /* Reference */
			shift(43), /* Not */
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
			shift(44), /* space */

		},
	},
	actionRow{ // S21
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(6), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S22
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S23
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(20), /* space, reduce: Pattern */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S26
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S27
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S28
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S29
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S30
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S31
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			reduce(89), /* Empty, reduce: Space */
			reduce(89), /* EmptySet, reduce: Space */
			reduce(89), /* TreeNode, reduce: Space */
			reduce(89), /* LeafNode, reduce: Space */
			reduce(89), /* Concat, reduce: Space */
			reduce(89), /* Or, reduce: Space */
			reduce(89), /* And, reduce: Space */
			reduce(89), /* ZeroOrMore, reduce: Space */
			reduce(89), /* Reference, reduce: Space */
			reduce(89), /* Not, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S33
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(5), /* space, reduce: PatternDecl */

		},
	},
	actionRow{ // S34
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S35
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(19), /* space, reduce: Pattern */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S37
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S38
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S39
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S40
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S41
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S42
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S43
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			reduce(88), /* Empty, reduce: Space */
			reduce(88), /* EmptySet, reduce: Space */
			reduce(88), /* TreeNode, reduce: Space */
			reduce(88), /* LeafNode, reduce: Space */
			reduce(88), /* Concat, reduce: Space */
			reduce(88), /* Or, reduce: Space */
			reduce(88), /* And, reduce: Space */
			reduce(88), /* ZeroOrMore, reduce: Space */
			reduce(88), /* Reference, reduce: Space */
			reduce(88), /* Not, reduce: Space */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S45
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(70), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(71), /* space */

		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(76), /* AnyName, reduce: OpenParen */
			reduce(76), /* Name, reduce: OpenParen */
			nil,        /* string_lit */
			reduce(76), /* AnyNameExcept, reduce: OpenParen */
			reduce(76), /* NameChoice, reduce: OpenParen */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(89), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S49
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(79), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(71), /* space */

		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(76), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(76), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(76), /* []bool, reduce: OpenParen */
			reduce(76), /* []int, reduce: OpenParen */
			reduce(76), /* []uint, reduce: OpenParen */
			reduce(76), /* []double, reduce: OpenParen */
			reduce(76), /* []string, reduce: OpenParen */
			reduce(76), /* [][]byte, reduce: OpenParen */
			reduce(76), /* int_lit, reduce: OpenParen */
			reduce(76), /* uint_lit, reduce: OpenParen */
			reduce(76), /* double_lit, reduce: OpenParen */
			reduce(76), /* bytes_lit, reduce: OpenParen */
			reduce(76), /* bool_var, reduce: OpenParen */
			reduce(76), /* int_var, reduce: OpenParen */
			reduce(76), /* uint_var, reduce: OpenParen */
			reduce(76), /* double_var, reduce: OpenParen */
			reduce(76), /* string_var, reduce: OpenParen */
			reduce(76), /* bytes_var, reduce: OpenParen */
			reduce(76), /* true, reduce: OpenParen */
			reduce(76), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(109), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(71),  /* space */

		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			reduce(76), /* Empty, reduce: OpenParen */
			reduce(76), /* EmptySet, reduce: OpenParen */
			reduce(76), /* TreeNode, reduce: OpenParen */
			reduce(76), /* LeafNode, reduce: OpenParen */
			reduce(76), /* Concat, reduce: OpenParen */
			reduce(76), /* Or, reduce: OpenParen */
			reduce(76), /* And, reduce: OpenParen */
			reduce(76), /* ZeroOrMore, reduce: OpenParen */
			reduce(76), /* Reference, reduce: OpenParen */
			reduce(76), /* Not, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(136), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(71),  /* space */

		},
	},
	actionRow{ // S59
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(138), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

		},
	},
	actionRow{ // S60
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(76), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S61
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S62
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S63
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S65
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S66
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S67
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S68
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(147), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

		},
	},
	actionRow{ // S69
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S70
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(77), /* AnyName, reduce: OpenParen */
			reduce(77), /* Name, reduce: OpenParen */
			nil,        /* string_lit */
			reduce(77), /* AnyNameExcept, reduce: OpenParen */
			reduce(77), /* NameChoice, reduce: OpenParen */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			reduce(88), /* (, reduce: Space */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S72
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(149), /* AnyName */
			shift(150), /* Name */
			nil,        /* string_lit */
			shift(151), /* AnyNameExcept */
			shift(152), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(153), /* space */

		},
	},
	actionRow{ // S73
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(8), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(160), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S76
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S77
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S78
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(89), /* AnyName, reduce: Space */
			reduce(89), /* Name, reduce: Space */
			nil,        /* string_lit */
			reduce(89), /* AnyNameExcept, reduce: Space */
			reduce(89), /* NameChoice, reduce: Space */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S79
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(77), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(77), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(77), /* []bool, reduce: OpenParen */
			reduce(77), /* []int, reduce: OpenParen */
			reduce(77), /* []uint, reduce: OpenParen */
			reduce(77), /* []double, reduce: OpenParen */
			reduce(77), /* []string, reduce: OpenParen */
			reduce(77), /* [][]byte, reduce: OpenParen */
			reduce(77), /* int_lit, reduce: OpenParen */
			reduce(77), /* uint_lit, reduce: OpenParen */
			reduce(77), /* double_lit, reduce: OpenParen */
			reduce(77), /* bytes_lit, reduce: OpenParen */
			reduce(77), /* bool_var, reduce: OpenParen */
			reduce(77), /* int_var, reduce: OpenParen */
			reduce(77), /* uint_var, reduce: OpenParen */
			reduce(77), /* double_var, reduce: OpenParen */
			reduce(77), /* string_var, reduce: OpenParen */
			reduce(77), /* bytes_var, reduce: OpenParen */
			reduce(77), /* true, reduce: OpenParen */
			reduce(77), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S80
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(163), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(166), /* space */

		},
	},
	actionRow{ // S81
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S83
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S84
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(39), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(40), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(41), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(41), /* space, reduce: Expr */

		},
	},
	actionRow{ // S87
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

		},
	},
	actionRow{ // S88
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
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
			reduce(52), /* space, reduce: ListType */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
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
			reduce(53), /* space, reduce: ListType */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
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
			reduce(54), /* space, reduce: ListType */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(55), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(55), /* space, reduce: ListType */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(56), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(56), /* space, reduce: ListType */

		},
	},
	actionRow{ // S93
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(57), /* {, reduce: ListType */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(57), /* space, reduce: ListType */

		},
	},
	actionRow{ // S94
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(58), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(58), /* space, reduce: SpaceTerminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S96
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S97
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(62), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(63), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(65), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(66), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(67), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S102
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S103
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(69), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S104
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(70), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(70), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S105
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(71), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(71), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(72), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(72), /* space, reduce: Bool */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(73), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(73), /* space, reduce: Bool */

		},
	},
	actionRow{ // S108
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(89), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int, reduce: Space */
			reduce(89), /* []uint, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int_lit, reduce: Space */
			reduce(89), /* uint_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int_var, reduce: Space */
			reduce(89), /* uint_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

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
			reduce(77), /* Empty, reduce: OpenParen */
			reduce(77), /* EmptySet, reduce: OpenParen */
			reduce(77), /* TreeNode, reduce: OpenParen */
			reduce(77), /* LeafNode, reduce: OpenParen */
			reduce(77), /* Concat, reduce: OpenParen */
			reduce(77), /* Or, reduce: OpenParen */
			reduce(77), /* And, reduce: OpenParen */
			reduce(77), /* ZeroOrMore, reduce: OpenParen */
			reduce(77), /* Reference, reduce: OpenParen */
			reduce(77), /* Not, reduce: OpenParen */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

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
			shift(178), /* Empty */
			shift(179), /* EmptySet */
			shift(180), /* TreeNode */
			shift(181), /* LeafNode */
			shift(182), /* Concat */
			shift(183), /* Or */
			shift(184), /* And */
			shift(185), /* ZeroOrMore */
			shift(186), /* Reference */
			shift(187), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(44),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(18), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S114
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S115
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S116
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S117
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S118
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S119
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S120
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S121
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			shift(199), /* Empty */
			shift(200), /* EmptySet */
			shift(201), /* TreeNode */
			shift(202), /* LeafNode */
			shift(203), /* Concat */
			shift(204), /* Or */
			shift(205), /* And */
			shift(206), /* ZeroOrMore */
			shift(207), /* Reference */
			shift(208), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(44),  /* space */

		},
	},
	actionRow{ // S125
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S126
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(18), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S127
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(20), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S128
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S130
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S132
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S133
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S134
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S135
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S136
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(77), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S137
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(218), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

		},
	},
	actionRow{ // S138
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S139
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S140
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S146
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(227), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S149
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(7), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(160), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S151
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S152
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S153
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(88), /* AnyName, reduce: Space */
			reduce(88), /* Name, reduce: Space */
			nil,        /* string_lit */
			reduce(88), /* AnyNameExcept, reduce: Space */
			reduce(88), /* NameChoice, reduce: Space */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(233), /* , */
			nil,        /* ; */
			shift(234), /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			reduce(84), /* Empty, reduce: Comma */
			reduce(84), /* EmptySet, reduce: Comma */
			reduce(84), /* TreeNode, reduce: Comma */
			reduce(84), /* LeafNode, reduce: Comma */
			reduce(84), /* Concat, reduce: Comma */
			reduce(84), /* Or, reduce: Comma */
			reduce(84), /* And, reduce: Comma */
			reduce(84), /* ZeroOrMore, reduce: Comma */
			reduce(84), /* Reference, reduce: Comma */
			reduce(84), /* Not, reduce: Comma */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(84), /* space, reduce: Comma */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(89), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(236), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(71),  /* space */

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
			shift(238), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(239), /* space */

		},
	},
	actionRow{ // S160
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(76), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S161
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

		},
	},
	actionRow{ // S162
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S164
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

		},
	},
	actionRow{ // S165
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(59), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(59), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S166
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(88), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(88), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int, reduce: Space */
			reduce(88), /* []uint, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int_lit, reduce: Space */
			reduce(88), /* uint_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int_var, reduce: Space */
			reduce(88), /* uint_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S167
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(249), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(71),  /* space */

		},
	},
	actionRow{ // S168
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

		},
	},
	actionRow{ // S169
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(76), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(76), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(76), /* []bool, reduce: OpenParen */
			reduce(76), /* []int, reduce: OpenParen */
			reduce(76), /* []uint, reduce: OpenParen */
			reduce(76), /* []double, reduce: OpenParen */
			reduce(76), /* []string, reduce: OpenParen */
			reduce(76), /* [][]byte, reduce: OpenParen */
			reduce(76), /* int_lit, reduce: OpenParen */
			reduce(76), /* uint_lit, reduce: OpenParen */
			reduce(76), /* double_lit, reduce: OpenParen */
			reduce(76), /* bytes_lit, reduce: OpenParen */
			reduce(76), /* bool_var, reduce: OpenParen */
			reduce(76), /* int_var, reduce: OpenParen */
			reduce(76), /* uint_var, reduce: OpenParen */
			reduce(76), /* double_var, reduce: OpenParen */
			reduce(76), /* string_var, reduce: OpenParen */
			reduce(76), /* bytes_var, reduce: OpenParen */
			reduce(76), /* true, reduce: OpenParen */
			reduce(76), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(76), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(76), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S170
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(276), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(277), /* space */

		},
	},
	actionRow{ // S171
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S172
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(78), /* $, reduce: CloseParen */
			reduce(78), /* id, reduce: CloseParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(78), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S173
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(89), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(278), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(279), /* space */

		},
	},
	actionRow{ // S175
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(304), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S176
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(80), /* id, reduce: OpenCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(80), /* string_lit, reduce: OpenCurly */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(80), /* []bool, reduce: OpenCurly */
			reduce(80), /* []int, reduce: OpenCurly */
			reduce(80), /* []uint, reduce: OpenCurly */
			reduce(80), /* []double, reduce: OpenCurly */
			reduce(80), /* []string, reduce: OpenCurly */
			reduce(80), /* [][]byte, reduce: OpenCurly */
			reduce(80), /* int_lit, reduce: OpenCurly */
			reduce(80), /* uint_lit, reduce: OpenCurly */
			reduce(80), /* double_lit, reduce: OpenCurly */
			reduce(80), /* bytes_lit, reduce: OpenCurly */
			reduce(80), /* bool_var, reduce: OpenCurly */
			reduce(80), /* int_var, reduce: OpenCurly */
			reduce(80), /* uint_var, reduce: OpenCurly */
			reduce(80), /* double_var, reduce: OpenCurly */
			reduce(80), /* string_var, reduce: OpenCurly */
			reduce(80), /* bytes_var, reduce: OpenCurly */
			reduce(80), /* true, reduce: OpenCurly */
			reduce(80), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(80), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			reduce(80), /* space, reduce: OpenCurly */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(89), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S178
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(17), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S180
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S181
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S182
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S183
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S184
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S185
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S186
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S187
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S189
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S190
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S195
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(322), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S199
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(17), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S200
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(19), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S201
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S202
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(51), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S203
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S204
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S205
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S206
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S207
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(60), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(54), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S209
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(32), /* $, reduce: Pattern */
			reduce(32), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(32), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S210
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S211
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S214
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S216
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(341), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

		},
	},
	actionRow{ // S217
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S218
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S219
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(36), /* $, reduce: Pattern */
			reduce(36), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(36), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S220
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(38), /* $, reduce: Pattern */
			reduce(38), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(38), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S221
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S222
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(23), /* space, reduce: Pattern */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S224
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S226
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(31), /* $, reduce: Pattern */
			reduce(31), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(31), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S227
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S228
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(35), /* $, reduce: Pattern */
			reduce(35), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(35), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S229
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(37), /* $, reduce: Pattern */
			reduce(37), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(37), /* space, reduce: Pattern */

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
			shift(350), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(239), /* space */

		},
	},
	actionRow{ // S231
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

		},
	},
	actionRow{ // S232
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

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
			reduce(85), /* Empty, reduce: Comma */
			reduce(85), /* EmptySet, reduce: Comma */
			reduce(85), /* TreeNode, reduce: Comma */
			reduce(85), /* LeafNode, reduce: Comma */
			reduce(85), /* Concat, reduce: Comma */
			reduce(85), /* Or, reduce: Comma */
			reduce(85), /* And, reduce: Comma */
			reduce(85), /* ZeroOrMore, reduce: Comma */
			reduce(85), /* Reference, reduce: Comma */
			reduce(85), /* Not, reduce: Comma */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(85), /* space, reduce: Comma */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(88), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S235
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S236
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(77), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

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
			shift(354), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(355), /* space */

		},
	},
	actionRow{ // S238
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S239
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(89), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S240
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(359), /* AnyName */
			shift(360), /* Name */
			nil,        /* string_lit */
			shift(361), /* AnyNameExcept */
			shift(362), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(153), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(8), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(160), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S244
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(369), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S247
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

		},
	},
	actionRow{ // S248
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(304), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S249
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(77), /* id, reduce: OpenParen */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(77), /* string_lit, reduce: OpenParen */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(77), /* []bool, reduce: OpenParen */
			reduce(77), /* []int, reduce: OpenParen */
			reduce(77), /* []uint, reduce: OpenParen */
			reduce(77), /* []double, reduce: OpenParen */
			reduce(77), /* []string, reduce: OpenParen */
			reduce(77), /* [][]byte, reduce: OpenParen */
			reduce(77), /* int_lit, reduce: OpenParen */
			reduce(77), /* uint_lit, reduce: OpenParen */
			reduce(77), /* double_lit, reduce: OpenParen */
			reduce(77), /* bytes_lit, reduce: OpenParen */
			reduce(77), /* bool_var, reduce: OpenParen */
			reduce(77), /* int_var, reduce: OpenParen */
			reduce(77), /* uint_var, reduce: OpenParen */
			reduce(77), /* double_var, reduce: OpenParen */
			reduce(77), /* string_var, reduce: OpenParen */
			reduce(77), /* bytes_var, reduce: OpenParen */
			reduce(77), /* true, reduce: OpenParen */
			reduce(77), /* false, reduce: OpenParen */
			nil,        /* = */
			nil,        /* ( */
			reduce(77), /* ), reduce: OpenParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(77), /* space, reduce: OpenParen */

		},
	},
	actionRow{ // S250
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(374), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(377), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(378), /* space */

		},
	},
	actionRow{ // S251
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(64), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S253
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(45), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(45), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(50), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(50), /* ,, reduce: Exprs */
			nil,        /* ; */
			reduce(50), /* space, reduce: Exprs */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(39), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(39), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(39), /* space, reduce: Expr */

		},
	},
	actionRow{ // S256
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(40), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(40), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(40), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(41), /* ), reduce: Expr */
			nil,        /* { */
			nil,        /* } */
			reduce(41), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(41), /* space, reduce: Expr */

		},
	},
	actionRow{ // S258
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

		},
	},
	actionRow{ // S259
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

		},
	},
	actionRow{ // S260
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(58), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(58), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			reduce(58), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S261
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(60), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(61), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S263
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S264
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(63), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(65), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(66), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(67), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(68), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S269
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(69), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(69), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(69), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S270
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(70), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(70), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(70), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S271
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(71), /* ), reduce: Terminal */
			nil,        /* { */
			nil,        /* } */
			reduce(71), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(71), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S272
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(72), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(72), /* ,, reduce: Bool */
			nil,        /* ; */
			reduce(72), /* space, reduce: Bool */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(73), /* ), reduce: Bool */
			nil,        /* { */
			nil,        /* } */
			reduce(73), /* ,, reduce: Bool */
			nil,        /* ; */
			reduce(73), /* space, reduce: Bool */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(78), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(78), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S275
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(89), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int, reduce: Space */
			reduce(89), /* []uint, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int_lit, reduce: Space */
			reduce(89), /* uint_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int_var, reduce: Space */
			reduce(89), /* uint_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(89), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S276
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(79), /* $, reduce: CloseParen */
			reduce(79), /* id, reduce: CloseParen */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(79), /* space, reduce: CloseParen */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(88), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S278
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(81), /* id, reduce: OpenCurly */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(81), /* string_lit, reduce: OpenCurly */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(81), /* []bool, reduce: OpenCurly */
			reduce(81), /* []int, reduce: OpenCurly */
			reduce(81), /* []uint, reduce: OpenCurly */
			reduce(81), /* []double, reduce: OpenCurly */
			reduce(81), /* []string, reduce: OpenCurly */
			reduce(81), /* [][]byte, reduce: OpenCurly */
			reduce(81), /* int_lit, reduce: OpenCurly */
			reduce(81), /* uint_lit, reduce: OpenCurly */
			reduce(81), /* double_lit, reduce: OpenCurly */
			reduce(81), /* bytes_lit, reduce: OpenCurly */
			reduce(81), /* bool_var, reduce: OpenCurly */
			reduce(81), /* int_var, reduce: OpenCurly */
			reduce(81), /* uint_var, reduce: OpenCurly */
			reduce(81), /* double_var, reduce: OpenCurly */
			reduce(81), /* string_var, reduce: OpenCurly */
			reduce(81), /* bytes_var, reduce: OpenCurly */
			reduce(81), /* true, reduce: OpenCurly */
			reduce(81), /* false, reduce: OpenCurly */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(81), /* }, reduce: OpenCurly */
			nil,        /* , */
			nil,        /* ; */
			reduce(81), /* space, reduce: OpenCurly */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			reduce(88), /* {, reduce: Space */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S280
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(386), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(389), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(390), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(64), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(50), /* }, reduce: Exprs */
			reduce(50), /* ,, reduce: Exprs */
			nil,        /* ; */
			reduce(50), /* space, reduce: Exprs */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(39), /* }, reduce: Expr */
			reduce(39), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(39), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(40), /* }, reduce: Expr */
			reduce(40), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(40), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(41), /* }, reduce: Expr */
			reduce(41), /* ,, reduce: Expr */
			nil,        /* ; */
			reduce(41), /* space, reduce: Expr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(304), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(49), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(58), /* }, reduce: SpaceTerminal */
			reduce(58), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			reduce(58), /* space, reduce: SpaceTerminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(60), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S292
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(61), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(62), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S294
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(63), /* space, reduce: Terminal */

		},
	},
	actionRow{ // S295
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(65), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(66), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(67), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(68), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(69), /* }, reduce: Terminal */
			reduce(69), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(69), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(70), /* }, reduce: Terminal */
			reduce(70), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(70), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(71), /* }, reduce: Terminal */
			reduce(71), /* ,, reduce: Terminal */
			nil,        /* ; */
			reduce(71), /* space, reduce: Terminal */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(72), /* }, reduce: Bool */
			reduce(72), /* ,, reduce: Bool */
			nil,        /* ; */
			reduce(72), /* space, reduce: Bool */

		},
	},
	actionRow{ // S303
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(73), /* }, reduce: Bool */
			reduce(73), /* ,, reduce: Bool */
			nil,        /* ; */
			reduce(73), /* space, reduce: Bool */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(82), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(82), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S305
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(89), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(89), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(89), /* []bool, reduce: Space */
			reduce(89), /* []int, reduce: Space */
			reduce(89), /* []uint, reduce: Space */
			reduce(89), /* []double, reduce: Space */
			reduce(89), /* []string, reduce: Space */
			reduce(89), /* [][]byte, reduce: Space */
			reduce(89), /* int_lit, reduce: Space */
			reduce(89), /* uint_lit, reduce: Space */
			reduce(89), /* double_lit, reduce: Space */
			reduce(89), /* bytes_lit, reduce: Space */
			reduce(89), /* bool_var, reduce: Space */
			reduce(89), /* int_var, reduce: Space */
			reduce(89), /* uint_var, reduce: Space */
			reduce(89), /* double_var, reduce: Space */
			reduce(89), /* string_var, reduce: Space */
			reduce(89), /* bytes_var, reduce: Space */
			reduce(89), /* true, reduce: Space */
			reduce(89), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(89), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S306
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S307
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S312
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(404), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S318
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S319
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S320
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S321
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(413), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S326
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S327
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(81),  /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(82),  /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(96),  /* int_lit */
			shift(97),  /* uint_lit */
			shift(98),  /* double_lit */
			shift(99),  /* bytes_lit */
			shift(100), /* bool_var */
			shift(101), /* int_var */
			shift(102), /* uint_var */
			shift(103), /* double_var */
			shift(104), /* string_var */
			shift(105), /* bytes_var */
			shift(106), /* true */
			shift(107), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

		},
	},
	actionRow{ // S328
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S329
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(112), /* Empty */
			shift(113), /* EmptySet */
			shift(114), /* TreeNode */
			shift(115), /* LeafNode */
			shift(116), /* Concat */
			shift(117), /* Or */
			shift(118), /* And */
			shift(119), /* ZeroOrMore */
			shift(120), /* Reference */
			shift(121), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S332
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(425), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(6),   /* space */

		},
	},
	actionRow{ // S333
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S334
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S336
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S337
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S338
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S340
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(434), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S343
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(34), /* $, reduce: Pattern */
			reduce(34), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(34), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S347
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(172), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S348
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(33), /* $, reduce: Pattern */
			reduce(33), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(33), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S349
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(441), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(355), /* space */

		},
	},
	actionRow{ // S350
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(369), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S353
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S354
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S355
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(88), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(446), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(277), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(12), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(78), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(78), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S359
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(7), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S360
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(160), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S361
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

		},
	},
	actionRow{ // S362
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(47), /* ( */
			nil,       /* ) */
			nil,       /* { */
			nil,       /* } */
			nil,       /* , */
			nil,       /* ; */
			shift(48), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(14), /* space, reduce: NameExpr */

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
			shift(451), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(239), /* space */

		},
	},
	actionRow{ // S365
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

		},
	},
	actionRow{ // S366
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(454), /* , */
			nil,        /* ; */
			shift(234), /* space */

		},
	},
	actionRow{ // S368
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

		},
	},
	actionRow{ // S369
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(84), /* AnyName, reduce: Comma */
			reduce(84), /* Name, reduce: Comma */
			nil,        /* string_lit */
			reduce(84), /* AnyNameExcept, reduce: Comma */
			reduce(84), /* NameChoice, reduce: Comma */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(84), /* space, reduce: Comma */

		},
	},
	actionRow{ // S370
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(43), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(43), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(304), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

		},
	},
	actionRow{ // S373
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(48), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(48), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

		},
	},
	actionRow{ // S376
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(59), /* ), reduce: SpaceTerminal */
			nil,        /* { */
			nil,        /* } */
			reduce(59), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			reduce(59), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S377
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(79), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(79), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S378
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(88), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(88), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int, reduce: Space */
			reduce(88), /* []uint, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int_lit, reduce: Space */
			reduce(88), /* uint_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int_var, reduce: Space */
			reduce(88), /* uint_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			reduce(88), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S379
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(463), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(377), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(465), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(44), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(44), /* space, reduce: Function */

		},
	},
	actionRow{ // S382
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

		},
	},
	actionRow{ // S383
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(84), /* id, reduce: Comma */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(84), /* string_lit, reduce: Comma */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(84), /* []bool, reduce: Comma */
			reduce(84), /* []int, reduce: Comma */
			reduce(84), /* []uint, reduce: Comma */
			reduce(84), /* []double, reduce: Comma */
			reduce(84), /* []string, reduce: Comma */
			reduce(84), /* [][]byte, reduce: Comma */
			reduce(84), /* int_lit, reduce: Comma */
			reduce(84), /* uint_lit, reduce: Comma */
			reduce(84), /* double_lit, reduce: Comma */
			reduce(84), /* bytes_lit, reduce: Comma */
			reduce(84), /* bool_var, reduce: Comma */
			reduce(84), /* int_var, reduce: Comma */
			reduce(84), /* uint_var, reduce: Comma */
			reduce(84), /* double_var, reduce: Comma */
			reduce(84), /* string_var, reduce: Comma */
			reduce(84), /* bytes_var, reduce: Comma */
			reduce(84), /* true, reduce: Comma */
			reduce(84), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(84), /* space, reduce: Comma */

		},
	},
	actionRow{ // S384
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(89), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(89), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S385
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(471), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S386
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			shift(169), /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(48),  /* space */

		},
	},
	actionRow{ // S387
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			shift(176), /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(177), /* space */

		},
	},
	actionRow{ // S388
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(59), /* }, reduce: SpaceTerminal */
			reduce(59), /* ,, reduce: SpaceTerminal */
			nil,        /* ; */
			reduce(59), /* space, reduce: SpaceTerminal */

		},
	},
	actionRow{ // S389
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(83), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(83), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S390
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(88), /* id, reduce: Space */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(88), /* string_lit, reduce: Space */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(88), /* []bool, reduce: Space */
			reduce(88), /* []int, reduce: Space */
			reduce(88), /* []uint, reduce: Space */
			reduce(88), /* []double, reduce: Space */
			reduce(88), /* []string, reduce: Space */
			reduce(88), /* [][]byte, reduce: Space */
			reduce(88), /* int_lit, reduce: Space */
			reduce(88), /* uint_lit, reduce: Space */
			reduce(88), /* double_lit, reduce: Space */
			reduce(88), /* bytes_lit, reduce: Space */
			reduce(88), /* bool_var, reduce: Space */
			reduce(88), /* int_var, reduce: Space */
			reduce(88), /* uint_var, reduce: Space */
			reduce(88), /* double_var, reduce: Space */
			reduce(88), /* string_var, reduce: Space */
			reduce(88), /* bytes_var, reduce: Space */
			reduce(88), /* true, reduce: Space */
			reduce(88), /* false, reduce: Space */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(88), /* }, reduce: Space */
			nil,        /* , */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S391
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(477), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(389), /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(478), /* space */

		},
	},
	actionRow{ // S393
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(108), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(47), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(47), /* space, reduce: List */

		},
	},
	actionRow{ // S395
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(89), /* }, reduce: Space */
			reduce(89), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(89), /* space, reduce: Space */

		},
	},
	actionRow{ // S396
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(484), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S397
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S398
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S400
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S401
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S403
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(491), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S406
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S407
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S408
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S409
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(32), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(32), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(36), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(36), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(38), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(38), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S416
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S417
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(30), /* $, reduce: Pattern */
			reduce(30), /* id, reduce: Pattern */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(30), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(156), /* , */
			nil,        /* ; */
			shift(157), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S424
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(505), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(11),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(377), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(277), /* space */

		},
	},
	actionRow{ // S429
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(24), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S430
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S431
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(32), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(32), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S434
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(36), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(36), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(38), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(38), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S437
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S438
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S439
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S440
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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S441
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(11), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(13), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S444
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(10), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S446
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(79), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(79), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S447
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(516), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(239), /* space */

		},
	},
	actionRow{ // S448
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

		},
	},
	actionRow{ // S449
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* id */
			shift(74), /* AnyName */
			shift(75), /* Name */
			nil,       /* string_lit */
			shift(76), /* AnyNameExcept */
			shift(77), /* NameChoice */
			nil,       /* Empty */
			nil,       /* EmptySet */
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			shift(78), /* space */

		},
	},
	actionRow{ // S450
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(519), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(355), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S452
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S453
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(369), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S454
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			reduce(85), /* AnyName, reduce: Comma */
			reduce(85), /* Name, reduce: Comma */
			nil,        /* string_lit */
			reduce(85), /* AnyNameExcept, reduce: Comma */
			reduce(85), /* NameChoice, reduce: Comma */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(85), /* space, reduce: Comma */

		},
	},
	actionRow{ // S455
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(42), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			nil,        /* , */
			nil,        /* ; */
			reduce(46), /* space, reduce: List */

		},
	},
	actionRow{ // S458
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(463), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

		},
	},
	actionRow{ // S459
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(471), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S460
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(374), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(528), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(378), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(45), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(45), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(45), /* space, reduce: Function */

		},
	},
	actionRow{ // S462
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(463), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

		},
	},
	actionRow{ // S463
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(78), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(78), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(78), /* space, reduce: CloseParen */

		},
	},
	actionRow{ // S464
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			reduce(85), /* id, reduce: Comma */
			nil,        /* AnyName */
			nil,        /* Name */
			reduce(85), /* string_lit, reduce: Comma */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			reduce(85), /* []bool, reduce: Comma */
			reduce(85), /* []int, reduce: Comma */
			reduce(85), /* []uint, reduce: Comma */
			reduce(85), /* []double, reduce: Comma */
			reduce(85), /* []string, reduce: Comma */
			reduce(85), /* [][]byte, reduce: Comma */
			reduce(85), /* int_lit, reduce: Comma */
			reduce(85), /* uint_lit, reduce: Comma */
			reduce(85), /* double_lit, reduce: Comma */
			reduce(85), /* bytes_lit, reduce: Comma */
			reduce(85), /* bool_var, reduce: Comma */
			reduce(85), /* int_var, reduce: Comma */
			reduce(85), /* uint_var, reduce: Comma */
			reduce(85), /* double_var, reduce: Comma */
			reduce(85), /* string_var, reduce: Comma */
			reduce(85), /* bytes_var, reduce: Comma */
			reduce(85), /* true, reduce: Comma */
			reduce(85), /* false, reduce: Comma */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(85), /* space, reduce: Comma */

		},
	},
	actionRow{ // S465
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(88), /* ), reduce: Space */
			nil,        /* { */
			nil,        /* } */
			reduce(88), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S466
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(374), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(166), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(51), /* ), reduce: Exprs */
			nil,        /* { */
			nil,        /* } */
			reduce(51), /* ,, reduce: Exprs */
			nil,        /* ; */
			reduce(51), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S468
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(386), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(531), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(390), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(471), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

		},
	},
	actionRow{ // S470
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(49), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(49), /* ,, reduce: List */
			nil,        /* ; */
			reduce(49), /* space, reduce: List */

		},
	},
	actionRow{ // S471
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(82), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(82), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			reduce(82), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S472
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(251), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(477), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(275), /* space */

		},
	},
	actionRow{ // S473
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(281), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(484), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(305), /* space */

		},
	},
	actionRow{ // S474
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(374), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(252), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(262), /* int_lit */
			shift(263), /* uint_lit */
			shift(264), /* double_lit */
			shift(265), /* bytes_lit */
			shift(266), /* bool_var */
			shift(267), /* int_var */
			shift(268), /* uint_var */
			shift(269), /* double_var */
			shift(270), /* string_var */
			shift(271), /* bytes_var */
			shift(272), /* true */
			shift(273), /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(538), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(378), /* space */

		},
	},
	actionRow{ // S475
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(45), /* }, reduce: Function */
			reduce(45), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(45), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(477), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(78), /* }, reduce: CloseParen */
			reduce(78), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(78), /* space, reduce: CloseParen */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(88), /* }, reduce: Space */
			reduce(88), /* ,, reduce: Space */
			nil,        /* ; */
			reduce(88), /* space, reduce: Space */

		},
	},
	actionRow{ // S479
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(386), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(166), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(51), /* }, reduce: Exprs */
			reduce(51), /* ,, reduce: Exprs */
			nil,        /* ; */
			reduce(51), /* space, reduce: Exprs */

		},
	},
	actionRow{ // S481
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			shift(386), /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			shift(282), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			shift(88),  /* []bool */
			shift(89),  /* []int */
			shift(90),  /* []uint */
			shift(91),  /* []double */
			shift(92),  /* []string */
			shift(93),  /* [][]byte */
			shift(292), /* int_lit */
			shift(293), /* uint_lit */
			shift(294), /* double_lit */
			shift(295), /* bytes_lit */
			shift(296), /* bool_var */
			shift(297), /* int_var */
			shift(298), /* uint_var */
			shift(299), /* double_var */
			shift(300), /* string_var */
			shift(301), /* bytes_var */
			shift(302), /* true */
			shift(303), /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(541), /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(390), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(484), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

		},
	},
	actionRow{ // S483
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(49), /* }, reduce: List */
			reduce(49), /* ,, reduce: List */
			nil,        /* ; */
			reduce(49), /* space, reduce: List */

		},
	},
	actionRow{ // S484
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(82), /* }, reduce: CloseCurly */
			reduce(82), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			reduce(82), /* space, reduce: CloseCurly */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(23), /* space, reduce: Pattern */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(31), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(31), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(35), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(35), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(37), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(37), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(34), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(34), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S499
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

		},
	},
	actionRow{ // S500
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(23), /* space, reduce: Pattern */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			shift(126), /* Empty */
			shift(127), /* EmptySet */
			shift(128), /* TreeNode */
			shift(129), /* LeafNode */
			shift(130), /* Concat */
			shift(131), /* Or */
			shift(132), /* And */
			shift(133), /* ZeroOrMore */
			shift(134), /* Reference */
			shift(135), /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(32),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(31), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(31), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(35), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(35), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(37), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(37), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S511
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S512
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(34), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(34), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S513
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(9), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			shift(563), /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(355), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S518
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(369), /* , */
			nil,        /* ; */
			shift(157), /* space */

		},
	},
	actionRow{ // S519
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S520
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(12), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(14), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S522
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(16), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(43), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(43), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(43), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(463), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(471), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(48), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(48), /* ,, reduce: List */
			nil,        /* ; */
			reduce(48), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(79), /* ), reduce: CloseParen */
			nil,        /* { */
			nil,        /* } */
			reduce(79), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(79), /* space, reduce: CloseParen */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(528), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(465), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(44), /* ), reduce: Function */
			nil,        /* { */
			nil,        /* } */
			reduce(44), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(44), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(83), /* ), reduce: CloseCurly */
			nil,        /* { */
			nil,        /* } */
			reduce(83), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			reduce(83), /* space, reduce: CloseCurly */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(531), /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(478), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(47), /* ), reduce: List */
			nil,        /* { */
			nil,        /* } */
			reduce(47), /* ,, reduce: List */
			nil,        /* ; */
			reduce(47), /* space, reduce: List */

		},
	},
	actionRow{ // S534
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(43), /* }, reduce: Function */
			reduce(43), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(43), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(477), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(384), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(484), /* } */
			shift(383), /* , */
			nil,        /* ; */
			shift(395), /* space */

		},
	},
	actionRow{ // S537
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(48), /* }, reduce: List */
			reduce(48), /* ,, reduce: List */
			nil,        /* ; */
			reduce(48), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(79), /* }, reduce: CloseParen */
			reduce(79), /* ,, reduce: CloseParen */
			nil,        /* ; */
			reduce(79), /* space, reduce: CloseParen */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(538), /* ) */
			nil,        /* { */
			nil,        /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(465), /* space */

		},
	},
	actionRow{ // S540
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(44), /* }, reduce: Function */
			reduce(44), /* ,, reduce: Function */
			nil,        /* ; */
			reduce(44), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(83), /* }, reduce: CloseCurly */
			reduce(83), /* ,, reduce: CloseCurly */
			nil,        /* ; */
			reduce(83), /* space, reduce: CloseCurly */

		},
	},
	actionRow{ // S542
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			shift(541), /* } */
			shift(464), /* , */
			nil,        /* ; */
			shift(478), /* space */

		},
	},
	actionRow{ // S543
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			reduce(47), /* }, reduce: List */
			reduce(47), /* ,, reduce: List */
			nil,        /* ; */
			reduce(47), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S545
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(358), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S548
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(33), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(33), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(22), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S550
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(26), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S551
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(28), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			reduce(30), /* ,, reduce: Pattern */
			nil,        /* ; */
			reduce(30), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(33), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(33), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(22), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(26), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(28), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S561
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			reduce(30), /* ), reduce: Pattern */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			reduce(30), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S562
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(15), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S563
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S564
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(11), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(13), /* space, reduce: NameExpr */

		},
	},
	actionRow{ // S566
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			shift(242), /* AnyName */
			shift(243), /* Name */
			nil,        /* string_lit */
			shift(244), /* AnyNameExcept */
			shift(245), /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(78),  /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(10), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

		},
	},
	actionRow{ // S569
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(42), /* space, reduce: Function */

		},
	},
	actionRow{ // S570
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(46), /* space, reduce: List */

		},
	},
	actionRow{ // S571
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(42), /* space, reduce: Function */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
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
			reduce(46), /* space, reduce: List */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S574
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(25), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(27), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
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
			reduce(29), /* space, reduce: Pattern */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(21), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S578
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(25), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S579
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(27), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S580
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* id */
			nil,        /* AnyName */
			nil,        /* Name */
			nil,        /* string_lit */
			nil,        /* AnyNameExcept */
			nil,        /* NameChoice */
			nil,        /* Empty */
			nil,        /* EmptySet */
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(29), /* space, reduce: Pattern */

		},
	},
	actionRow{ // S581
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
			nil,       /* TreeNode */
			nil,       /* LeafNode */
			nil,       /* Concat */
			nil,       /* Or */
			nil,       /* And */
			nil,       /* ZeroOrMore */
			nil,       /* Reference */
			nil,       /* Not */
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
			reduce(9), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
			nil,        /* int_var */
			nil,        /* uint_var */
			nil,        /* double_var */
			nil,        /* string_var */
			nil,        /* bytes_var */
			nil,        /* true */
			nil,        /* false */
			nil,        /* = */
			nil,        /* ( */
			shift(274), /* ) */
			nil,        /* { */
			nil,        /* } */
			nil,        /* , */
			nil,        /* ; */
			shift(173), /* space */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(16), /* space, reduce: NameExpr */

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
			nil,        /* TreeNode */
			nil,        /* LeafNode */
			nil,        /* Concat */
			nil,        /* Or */
			nil,        /* And */
			nil,        /* ZeroOrMore */
			nil,        /* Reference */
			nil,        /* Not */
			nil,        /* []bool */
			nil,        /* []int */
			nil,        /* []uint */
			nil,        /* []double */
			nil,        /* []string */
			nil,        /* [][]byte */
			nil,        /* int_lit */
			nil,        /* uint_lit */
			nil,        /* double_lit */
			nil,        /* bytes_lit */
			nil,        /* bool_var */
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
			reduce(15), /* space, reduce: NameExpr */

		},
	},
}

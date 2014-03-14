package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{

	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 38: // ['&','&']
			return 3
		case r == 40: // ['(','(']
			return 4
		case r == 41: // [')',')']
			return 5
		case r == 44: // [',',',']
			return 6
		case r == 45: // ['-','-']
			return 7
		case r == 46: // ['.','.']
			return 8
		case r == 47: // ['/','/']
			return 9
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 60: // ['<','<']
			return 11
		case r == 61: // ['=','=']
			return 12
		case r == 62: // ['>','>']
			return 13
		case 65 <= r && r <= 69: // ['A','E']
			return 14
		case r == 70: // ['F','F']
			return 15
		case 71 <= r && r <= 83: // ['G','S']
			return 14
		case r == 84: // ['T','T']
			return 16
		case 85 <= r && r <= 90: // ['U','Z']
			return 14
		case r == 95: // ['_','_']
			return 17
		case r == 96: // ['`','`']
			return 18
		case r == 97: // ['a','a']
			return 19
		case 98 <= r && r <= 100: // ['b','d']
			return 20
		case r == 101: // ['e','e']
			return 21
		case r == 102: // ['f','f']
			return 22
		case 103 <= r && r <= 104: // ['g','h']
			return 20
		case r == 105: // ['i','i']
			return 23
		case 106 <= r && r <= 110: // ['j','n']
			return 20
		case r == 111: // ['o','o']
			return 24
		case 112 <= r && r <= 113: // ['p','q']
			return 20
		case r == 114: // ['r','r']
			return 25
		case r == 115: // ['s','s']
			return 20
		case r == 116: // ['t','t']
			return 26
		case r == 117: // ['u','u']
			return 27
		case 118 <= r && r <= 122: // ['v','z']
			return 20
		case r == 123: // ['{','{']
			return 28
		case r == 124: // ['|','|']
			return 29
		case r == 125: // ['}','}']
			return 30

		}
		return NoState
	},

	// S1
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S2
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S3
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 34

		}
		return NoState
	},

	// S4
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S5
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S6
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S7
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10

		}
		return NoState
	},

	// S8
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S9
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 35
		case r == 47: // ['/','/']
			return 36

		}
		return NoState
	},

	// S10
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 10

		}
		return NoState
	},

	// S11
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 37

		}
		return NoState
	},

	// S12
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 38

		}
		return NoState
	},

	// S13
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 39

		}
		return NoState
	},

	// S14
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S15
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case r == 97: // ['a','a']
			return 44
		case 98 <= r && r <= 122: // ['b','z']
			return 43

		}
		return NoState
	},

	// S16
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 113: // ['a','q']
			return 43
		case r == 114: // ['r','r']
			return 45
		case 115 <= r && r <= 122: // ['s','z']
			return 43

		}
		return NoState
	},

	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S18
	func(r rune) int {
		switch {
		case r == 96: // ['`','`']
			return 46

		default:
			return 18

		}
		return NoState
	},

	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 109: // ['a','m']
			return 43
		case r == 110: // ['n','n']
			return 47
		case 111 <= r && r <= 122: // ['o','z']
			return 43

		}
		return NoState
	},

	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S21
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 107: // ['a','k']
			return 43
		case r == 108: // ['l','l']
			return 48
		case 109 <= r && r <= 122: // ['m','z']
			return 43

		}
		return NoState
	},

	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case r == 97: // ['a','a']
			return 49
		case 98 <= r && r <= 122: // ['b','z']
			return 43

		}
		return NoState
	},

	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 101: // ['a','e']
			return 43
		case r == 102: // ['f','f']
			return 50
		case 103 <= r && r <= 109: // ['g','m']
			return 43
		case r == 110: // ['n','n']
			return 51
		case 111 <= r && r <= 122: // ['o','z']
			return 43

		}
		return NoState
	},

	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 113: // ['a','q']
			return 43
		case r == 114: // ['r','r']
			return 52
		case 115 <= r && r <= 122: // ['s','z']
			return 43

		}
		return NoState
	},

	// S25
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 110: // ['a','n']
			return 43
		case r == 111: // ['o','o']
			return 53
		case 112 <= r && r <= 122: // ['p','z']
			return 43

		}
		return NoState
	},

	// S26
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 103: // ['a','g']
			return 43
		case r == 104: // ['h','h']
			return 54
		case 105 <= r && r <= 113: // ['i','q']
			return 43
		case r == 114: // ['r','r']
			return 55
		case 115 <= r && r <= 122: // ['s','z']
			return 43

		}
		return NoState
	},

	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 104: // ['a','h']
			return 43
		case r == 105: // ['i','i']
			return 56
		case 106 <= r && r <= 122: // ['j','z']
			return 43

		}
		return NoState
	},

	// S28
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S29
	func(r rune) int {
		switch {
		case r == 124: // ['|','|']
			return 57

		}
		return NoState
	},

	// S30
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S31
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S32
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 58
		case r == 39: // [''',''']
			return 58
		case 48 <= r && r <= 55: // ['0','7']
			return 59
		case r == 85: // ['U','U']
			return 60
		case r == 92: // ['\','\']
			return 58
		case r == 97: // ['a','a']
			return 58
		case r == 98: // ['b','b']
			return 58
		case r == 102: // ['f','f']
			return 58
		case r == 110: // ['n','n']
			return 58
		case r == 114: // ['r','r']
			return 58
		case r == 116: // ['t','t']
			return 58
		case r == 117: // ['u','u']
			return 61
		case r == 118: // ['v','v']
			return 58
		case r == 120: // ['x','x']
			return 62

		}
		return NoState
	},

	// S33
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S34
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S35
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 63

		default:
			return 35

		}
		return NoState
	},

	// S36
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 64

		default:
			return 36

		}
		return NoState
	},

	// S37
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S38
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S39
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S40
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S41
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S42
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S43
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S44
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 107: // ['a','k']
			return 43
		case r == 108: // ['l','l']
			return 65
		case 109 <= r && r <= 122: // ['m','z']
			return 43

		}
		return NoState
	},

	// S45
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 116: // ['a','t']
			return 43
		case r == 117: // ['u','u']
			return 66
		case 118 <= r && r <= 122: // ['v','z']
			return 43

		}
		return NoState
	},

	// S46
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 99: // ['a','c']
			return 43
		case r == 100: // ['d','d']
			return 67
		case 101 <= r && r <= 122: // ['e','z']
			return 43

		}
		return NoState
	},

	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 114: // ['a','r']
			return 43
		case r == 115: // ['s','s']
			return 68
		case 116 <= r && r <= 122: // ['t','z']
			return 43

		}
		return NoState
	},

	// S49
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 107: // ['a','k']
			return 43
		case r == 108: // ['l','l']
			return 69
		case 109 <= r && r <= 122: // ['m','z']
			return 43

		}
		return NoState
	},

	// S50
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S51
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 115: // ['a','s']
			return 43
		case r == 116: // ['t','t']
			return 70
		case 117 <= r && r <= 122: // ['u','z']
			return 43

		}
		return NoState
	},

	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S53
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 110: // ['a','n']
			return 43
		case r == 111: // ['o','o']
			return 71
		case 112 <= r && r <= 122: // ['p','z']
			return 43

		}
		return NoState
	},

	// S54
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 72
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S55
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 116: // ['a','t']
			return 43
		case r == 117: // ['u','u']
			return 73
		case 118 <= r && r <= 122: // ['v','z']
			return 43

		}
		return NoState
	},

	// S56
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 109: // ['a','m']
			return 43
		case r == 110: // ['n','n']
			return 74
		case 111 <= r && r <= 122: // ['o','z']
			return 43

		}
		return NoState
	},

	// S57
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S58
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S59
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 75

		}
		return NoState
	},

	// S60
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		case 65 <= r && r <= 70: // ['A','F']
			return 76
		case 97 <= r && r <= 102: // ['a','f']
			return 76

		}
		return NoState
	},

	// S61
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 77
		case 65 <= r && r <= 70: // ['A','F']
			return 77
		case 97 <= r && r <= 102: // ['a','f']
			return 77

		}
		return NoState
	},

	// S62
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 78
		case 65 <= r && r <= 70: // ['A','F']
			return 78
		case 97 <= r && r <= 102: // ['a','f']
			return 78

		}
		return NoState
	},

	// S63
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 63
		case r == 47: // ['/','/']
			return 79

		default:
			return 35

		}
		return NoState
	},

	// S64
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S65
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 114: // ['a','r']
			return 43
		case r == 115: // ['s','s']
			return 80
		case 116 <= r && r <= 122: // ['t','z']
			return 43

		}
		return NoState
	},

	// S66
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 81
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S67
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S68
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 82
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S69
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 114: // ['a','r']
			return 43
		case r == 115: // ['s','s']
			return 83
		case 116 <= r && r <= 122: // ['t','z']
			return 43

		}
		return NoState
	},

	// S70
	func(r rune) int {
		switch {
		case 48 <= r && r <= 53: // ['0','5']
			return 40
		case r == 54: // ['6','6']
			return 84
		case 55 <= r && r <= 57: // ['7','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S71
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 115: // ['a','s']
			return 43
		case r == 116: // ['t','t']
			return 85
		case 117 <= r && r <= 122: // ['u','z']
			return 43

		}
		return NoState
	},

	// S72
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 109: // ['a','m']
			return 43
		case r == 110: // ['n','n']
			return 86
		case 111 <= r && r <= 122: // ['o','z']
			return 43

		}
		return NoState
	},

	// S73
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 87
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S74
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 115: // ['a','s']
			return 43
		case r == 116: // ['t','t']
			return 88
		case 117 <= r && r <= 122: // ['u','z']
			return 43

		}
		return NoState
	},

	// S75
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 89

		}
		return NoState
	},

	// S76
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 90
		case 65 <= r && r <= 70: // ['A','F']
			return 90
		case 97 <= r && r <= 102: // ['a','f']
			return 90

		}
		return NoState
	},

	// S77
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 91
		case 65 <= r && r <= 70: // ['A','F']
			return 91
		case 97 <= r && r <= 102: // ['a','f']
			return 91

		}
		return NoState
	},

	// S78
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 92
		case 65 <= r && r <= 70: // ['A','F']
			return 92
		case 97 <= r && r <= 102: // ['a','f']
			return 92

		}
		return NoState
	},

	// S79
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S80
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 93
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S81
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S82
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S83
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 100: // ['a','d']
			return 43
		case r == 101: // ['e','e']
			return 94
		case 102 <= r && r <= 122: // ['f','z']
			return 43

		}
		return NoState
	},

	// S84
	func(r rune) int {
		switch {
		case 48 <= r && r <= 51: // ['0','3']
			return 40
		case r == 52: // ['4','4']
			return 95
		case 53 <= r && r <= 57: // ['5','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S85
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S86
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S87
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S88
	func(r rune) int {
		switch {
		case 48 <= r && r <= 53: // ['0','5']
			return 40
		case r == 54: // ['6','6']
			return 96
		case 55 <= r && r <= 57: // ['7','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S89
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S90
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 97
		case 65 <= r && r <= 70: // ['A','F']
			return 97
		case 97 <= r && r <= 102: // ['a','f']
			return 97

		}
		return NoState
	},

	// S91
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 98
		case 65 <= r && r <= 70: // ['A','F']
			return 98
		case 97 <= r && r <= 102: // ['a','f']
			return 98

		}
		return NoState
	},

	// S92
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S93
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S94
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S95
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S96
	func(r rune) int {
		switch {
		case 48 <= r && r <= 51: // ['0','3']
			return 40
		case r == 52: // ['4','4']
			return 99
		case 53 <= r && r <= 57: // ['5','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S97
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 100
		case 65 <= r && r <= 70: // ['A','F']
			return 100
		case 97 <= r && r <= 102: // ['a','f']
			return 100

		}
		return NoState
	},

	// S98
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 101
		case 65 <= r && r <= 70: // ['A','F']
			return 101
		case 97 <= r && r <= 102: // ['a','f']
			return 101

		}
		return NoState
	},

	// S99
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 41
		case r == 95: // ['_','_']
			return 42
		case 97 <= r && r <= 122: // ['a','z']
			return 43

		}
		return NoState
	},

	// S100
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 102
		case 65 <= r && r <= 70: // ['A','F']
			return 102
		case 97 <= r && r <= 102: // ['a','f']
			return 102

		}
		return NoState
	},

	// S101
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},

	// S102
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 103
		case 65 <= r && r <= 70: // ['A','F']
			return 103
		case 97 <= r && r <= 102: // ['a','f']
			return 103

		}
		return NoState
	},

	// S103
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 104
		case 65 <= r && r <= 70: // ['A','F']
			return 104
		case 97 <= r && r <= 102: // ['a','f']
			return 104

		}
		return NoState
	},

	// S104
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 105
		case 65 <= r && r <= 70: // ['A','F']
			return 105
		case 97 <= r && r <= 102: // ['a','f']
			return 105

		}
		return NoState
	},

	// S105
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32

		default:
			return 33

		}
		return NoState
	},
}

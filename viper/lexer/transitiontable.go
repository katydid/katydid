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
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 35: // ['#','#']
			return 4
		case r == 36: // ['$','$']
			return 5
		case r == 38: // ['&','&']
			return 6
		case r == 40: // ['(','(']
			return 7
		case r == 41: // [')',')']
			return 8
		case r == 42: // ['*','*']
			return 9
		case r == 44: // [',',',']
			return 10
		case r == 46: // ['.','.']
			return 11
		case r == 47: // ['/','/']
			return 12
		case r == 58: // [':',':']
			return 13
		case r == 59: // [';',';']
			return 14
		case r == 61: // ['=','=']
			return 15
		case r == 64: // ['@','@']
			return 16
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 91: // ['[','[']
			return 18
		case r == 93: // [']',']']
			return 19
		case r == 95: // ['_','_']
			return 20
		case r == 96: // ['`','`']
			return 21
		case 97 <= r && r <= 98: // ['a','b']
			return 22
		case r == 99: // ['c','c']
			return 23
		case r == 100: // ['d','d']
			return 24
		case r == 101: // ['e','e']
			return 22
		case r == 102: // ['f','f']
			return 25
		case 103 <= r && r <= 104: // ['g','h']
			return 22
		case r == 105: // ['i','i']
			return 26
		case 106 <= r && r <= 113: // ['j','q']
			return 22
		case r == 114: // ['r','r']
			return 27
		case r == 115: // ['s','s']
			return 28
		case r == 116: // ['t','t']
			return 29
		case r == 117: // ['u','u']
			return 30
		case 118 <= r && r <= 122: // ['v','z']
			return 22
		case r == 123: // ['{','{']
			return 31
		case r == 124: // ['|','|']
			return 32
		case r == 125: // ['}','}']
			return 33
		case r == 126: // ['~','~']
			return 34

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

		}
		return NoState

	},

	// S3
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

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
		case r == 91: // ['[','[']
			return 38
		case r == 98: // ['b','b']
			return 39
		case r == 100: // ['d','d']
			return 40
		case r == 105: // ['i','i']
			return 41
		case r == 115: // ['s','s']
			return 42
		case r == 117: // ['u','u']
			return 43

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

		}
		return NoState

	},

	// S10
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S11
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S12
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 44
		case r == 47: // ['/','/']
			return 45

		}
		return NoState

	},

	// S13
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S14
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S15
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S16
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S18
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 50

		}
		return NoState

	},

	// S19
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S21
	func(r rune) int {
		switch {
		case r == 96: // ['`','`']
			return 51

		default:
			return 21
		}

	},

	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 52
		case 98 <= r && r <= 122: // ['b','z']
			return 49

		}
		return NoState

	},

	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 110: // ['a','n']
			return 49
		case r == 111: // ['o','o']
			return 53
		case 112 <= r && r <= 122: // ['p','z']
			return 49

		}
		return NoState

	},

	// S25
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 54
		case 98 <= r && r <= 104: // ['b','h']
			return 49
		case r == 105: // ['i','i']
			return 55
		case 106 <= r && r <= 122: // ['j','z']
			return 49

		}
		return NoState

	},

	// S26
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 109: // ['a','m']
			return 49
		case r == 110: // ['n','n']
			return 56
		case 111 <= r && r <= 122: // ['o','z']
			return 49

		}
		return NoState

	},

	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 100: // ['a','d']
			return 49
		case r == 101: // ['e','e']
			return 57
		case 102 <= r && r <= 122: // ['f','z']
			return 49

		}
		return NoState

	},

	// S28
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 115: // ['a','s']
			return 49
		case r == 116: // ['t','t']
			return 58
		case 117 <= r && r <= 122: // ['u','z']
			return 49

		}
		return NoState

	},

	// S29
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 113: // ['a','q']
			return 49
		case r == 114: // ['r','r']
			return 59
		case 115 <= r && r <= 122: // ['s','z']
			return 49

		}
		return NoState

	},

	// S30
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 104: // ['a','h']
			return 49
		case r == 105: // ['i','i']
			return 60
		case 106 <= r && r <= 122: // ['j','z']
			return 49

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

		}
		return NoState

	},

	// S33
	func(r rune) int {
		switch {

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

		}
		return NoState

	},

	// S36
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 61
		case r == 39: // [''',''']
			return 61
		case 48 <= r && r <= 55: // ['0','7']
			return 62
		case r == 85: // ['U','U']
			return 63
		case r == 92: // ['\','\']
			return 61
		case r == 97: // ['a','a']
			return 61
		case r == 98: // ['b','b']
			return 61
		case r == 102: // ['f','f']
			return 61
		case r == 110: // ['n','n']
			return 61
		case r == 114: // ['r','r']
			return 61
		case r == 116: // ['t','t']
			return 61
		case r == 117: // ['u','u']
			return 64
		case r == 118: // ['v','v']
			return 61
		case r == 120: // ['x','x']
			return 65

		}
		return NoState

	},

	// S37
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S38
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 66

		}
		return NoState

	},

	// S39
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 67

		}
		return NoState

	},

	// S40
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 68

		}
		return NoState

	},

	// S41
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 69

		}
		return NoState

	},

	// S42
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 70

		}
		return NoState

	},

	// S43
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 71

		}
		return NoState

	},

	// S44
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 72

		default:
			return 44
		}

	},

	// S45
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 73

		default:
			return 45
		}

	},

	// S46
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S49
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S50
	func(r rune) int {
		switch {
		case r == 91: // ['[','[']
			return 74
		case r == 98: // ['b','b']
			return 75
		case r == 100: // ['d','d']
			return 76
		case r == 105: // ['i','i']
			return 77
		case r == 115: // ['s','s']
			return 78
		case r == 117: // ['u','u']
			return 79

		}
		return NoState

	},

	// S51
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 80
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S53
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 116: // ['a','t']
			return 49
		case r == 117: // ['u','u']
			return 81
		case 118 <= r && r <= 122: // ['v','z']
			return 49

		}
		return NoState

	},

	// S54
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 82
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S55
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 109: // ['a','m']
			return 49
		case r == 110: // ['n','n']
			return 83
		case 111 <= r && r <= 122: // ['o','z']
			return 49

		}
		return NoState

	},

	// S56
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 115: // ['a','s']
			return 49
		case r == 116: // ['t','t']
			return 84
		case 117 <= r && r <= 122: // ['u','z']
			return 49

		}
		return NoState

	},

	// S57
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 115: // ['a','s']
			return 49
		case r == 116: // ['t','t']
			return 85
		case 117 <= r && r <= 122: // ['u','z']
			return 49

		}
		return NoState

	},

	// S58
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 86
		case 98 <= r && r <= 122: // ['b','z']
			return 49

		}
		return NoState

	},

	// S59
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 116: // ['a','t']
			return 49
		case r == 117: // ['u','u']
			return 87
		case 118 <= r && r <= 122: // ['v','z']
			return 49

		}
		return NoState

	},

	// S60
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 109: // ['a','m']
			return 49
		case r == 110: // ['n','n']
			return 88
		case 111 <= r && r <= 122: // ['o','z']
			return 49

		}
		return NoState

	},

	// S61
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S62
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 89

		}
		return NoState

	},

	// S63
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

	// S64
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

	// S65
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

	// S66
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 93

		}
		return NoState

	},

	// S67
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 94

		}
		return NoState

	},

	// S68
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 95

		}
		return NoState

	},

	// S69
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 96

		}
		return NoState

	},

	// S70
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 97

		}
		return NoState

	},

	// S71
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 98

		}
		return NoState

	},

	// S72
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 72
		case r == 47: // ['/','/']
			return 99

		default:
			return 44
		}

	},

	// S73
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S74
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 100

		}
		return NoState

	},

	// S75
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 101
		case r == 121: // ['y','y']
			return 102

		}
		return NoState

	},

	// S76
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 103

		}
		return NoState

	},

	// S77
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 104

		}
		return NoState

	},

	// S78
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 105

		}
		return NoState

	},

	// S79
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 106

		}
		return NoState

	},

	// S80
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 107
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S81
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 49
		case r == 98: // ['b','b']
			return 108
		case 99 <= r && r <= 122: // ['c','z']
			return 49

		}
		return NoState

	},

	// S82
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 114: // ['a','r']
			return 49
		case r == 115: // ['s','s']
			return 109
		case 116 <= r && r <= 122: // ['t','z']
			return 49

		}
		return NoState

	},

	// S83
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 110
		case 98 <= r && r <= 122: // ['b','z']
			return 49

		}
		return NoState

	},

	// S84
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 111
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 100: // ['a','d']
			return 49
		case r == 101: // ['e','e']
			return 112
		case 102 <= r && r <= 122: // ['f','z']
			return 49

		}
		return NoState

	},

	// S85
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 116: // ['a','t']
			return 49
		case r == 117: // ['u','u']
			return 113
		case 118 <= r && r <= 122: // ['v','z']
			return 49

		}
		return NoState

	},

	// S86
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 113: // ['a','q']
			return 49
		case r == 114: // ['r','r']
			return 114
		case 115 <= r && r <= 122: // ['s','z']
			return 49

		}
		return NoState

	},

	// S87
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 100: // ['a','d']
			return 49
		case r == 101: // ['e','e']
			return 115
		case 102 <= r && r <= 122: // ['f','z']
			return 49

		}
		return NoState

	},

	// S88
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 115: // ['a','s']
			return 49
		case r == 116: // ['t','t']
			return 116
		case 117 <= r && r <= 122: // ['u','z']
			return 49

		}
		return NoState

	},

	// S89
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 117

		}
		return NoState

	},

	// S90
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 118
		case 65 <= r && r <= 70: // ['A','F']
			return 118
		case 97 <= r && r <= 102: // ['a','f']
			return 118

		}
		return NoState

	},

	// S91
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 119
		case 65 <= r && r <= 70: // ['A','F']
			return 119
		case 97 <= r && r <= 102: // ['a','f']
			return 119

		}
		return NoState

	},

	// S92
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 120
		case 65 <= r && r <= 70: // ['A','F']
			return 120
		case 97 <= r && r <= 102: // ['a','f']
			return 120

		}
		return NoState

	},

	// S93
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 121

		}
		return NoState

	},

	// S94
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 122

		}
		return NoState

	},

	// S95
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 123

		}
		return NoState

	},

	// S96
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S97
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 124

		}
		return NoState

	},

	// S98
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 125

		}
		return NoState

	},

	// S99
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S100
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 126

		}
		return NoState

	},

	// S101
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 127

		}
		return NoState

	},

	// S102
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 128

		}
		return NoState

	},

	// S103
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 129

		}
		return NoState

	},

	// S104
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 130

		}
		return NoState

	},

	// S105
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 131

		}
		return NoState

	},

	// S106
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 132

		}
		return NoState

	},

	// S107
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S108
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 133
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S109
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 100: // ['a','d']
			return 49
		case r == 101: // ['e','e']
			return 134
		case 102 <= r && r <= 122: // ['f','z']
			return 49

		}
		return NoState

	},

	// S110
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 135
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S111
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 136
		case r == 48: // ['0','0']
			return 137
		case 49 <= r && r <= 57: // ['1','9']
			return 138

		}
		return NoState

	},

	// S112
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 113: // ['a','q']
			return 49
		case r == 114: // ['r','r']
			return 139
		case 115 <= r && r <= 122: // ['s','z']
			return 49

		}
		return NoState

	},

	// S113
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 113: // ['a','q']
			return 49
		case r == 114: // ['r','r']
			return 140
		case 115 <= r && r <= 122: // ['s','z']
			return 49

		}
		return NoState

	},

	// S114
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 115: // ['a','s']
			return 49
		case r == 116: // ['t','t']
			return 141
		case 117 <= r && r <= 122: // ['u','z']
			return 49

		}
		return NoState

	},

	// S115
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S116
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 142
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S117
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S118
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 143
		case 65 <= r && r <= 70: // ['A','F']
			return 143
		case 97 <= r && r <= 102: // ['a','f']
			return 143

		}
		return NoState

	},

	// S119
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 144
		case 65 <= r && r <= 70: // ['A','F']
			return 144
		case 97 <= r && r <= 102: // ['a','f']
			return 144

		}
		return NoState

	},

	// S120
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S121
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 145

		}
		return NoState

	},

	// S122
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S123
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 146

		}
		return NoState

	},

	// S124
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 147

		}
		return NoState

	},

	// S125
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S126
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 148

		}
		return NoState

	},

	// S127
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 149

		}
		return NoState

	},

	// S128
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 150

		}
		return NoState

	},

	// S129
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 151

		}
		return NoState

	},

	// S130
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S131
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 152

		}
		return NoState

	},

	// S132
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 153

		}
		return NoState

	},

	// S133
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 100: // ['a','d']
			return 49
		case r == 101: // ['e','e']
			return 154
		case 102 <= r && r <= 122: // ['f','z']
			return 49

		}
		return NoState

	},

	// S134
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S135
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S136
	func(r rune) int {
		switch {
		case r == 48: // ['0','0']
			return 137
		case 49 <= r && r <= 57: // ['1','9']
			return 138

		}
		return NoState

	},

	// S137
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 155
		case 48 <= r && r <= 55: // ['0','7']
			return 156
		case r == 88: // ['X','X']
			return 157
		case r == 120: // ['x','x']
			return 157

		}
		return NoState

	},

	// S138
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 155
		case 48 <= r && r <= 57: // ['0','9']
			return 158

		}
		return NoState

	},

	// S139
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 109: // ['a','m']
			return 49
		case r == 110: // ['n','n']
			return 159
		case 111 <= r && r <= 122: // ['o','z']
			return 49

		}
		return NoState

	},

	// S140
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 109: // ['a','m']
			return 49
		case r == 110: // ['n','n']
			return 160
		case 111 <= r && r <= 122: // ['o','z']
			return 49

		}
		return NoState

	},

	// S141
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S142
	func(r rune) int {
		switch {
		case r == 48: // ['0','0']
			return 161
		case 49 <= r && r <= 57: // ['1','9']
			return 162

		}
		return NoState

	},

	// S143
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 163
		case 65 <= r && r <= 70: // ['A','F']
			return 163
		case 97 <= r && r <= 102: // ['a','f']
			return 163

		}
		return NoState

	},

	// S144
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 164
		case 65 <= r && r <= 70: // ['A','F']
			return 164
		case 97 <= r && r <= 102: // ['a','f']
			return 164

		}
		return NoState

	},

	// S145
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 165

		}
		return NoState

	},

	// S146
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 166

		}
		return NoState

	},

	// S147
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 167

		}
		return NoState

	},

	// S148
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 168

		}
		return NoState

	},

	// S149
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S150
	func(r rune) int {
		switch {
		case r == 123: // ['{','{']
			return 169

		}
		return NoState

	},

	// S151
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 170

		}
		return NoState

	},

	// S152
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 171

		}
		return NoState

	},

	// S153
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S154
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 172
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S155
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S156
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 155
		case 48 <= r && r <= 55: // ['0','7']
			return 156

		}
		return NoState

	},

	// S157
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 173
		case 65 <= r && r <= 70: // ['A','F']
			return 173
		case 97 <= r && r <= 102: // ['a','f']
			return 173

		}
		return NoState

	},

	// S158
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 155
		case 48 <= r && r <= 57: // ['0','9']
			return 158

		}
		return NoState

	},

	// S159
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case r == 97: // ['a','a']
			return 174
		case 98 <= r && r <= 122: // ['b','z']
			return 49

		}
		return NoState

	},

	// S160
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S161
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 175
		case 48 <= r && r <= 55: // ['0','7']
			return 176
		case r == 88: // ['X','X']
			return 177
		case r == 120: // ['x','x']
			return 177

		}
		return NoState

	},

	// S162
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 175
		case 48 <= r && r <= 57: // ['0','9']
			return 178

		}
		return NoState

	},

	// S163
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 179
		case 65 <= r && r <= 70: // ['A','F']
			return 179
		case 97 <= r && r <= 102: // ['a','f']
			return 179

		}
		return NoState

	},

	// S164
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S165
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S166
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S167
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S168
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 180

		}
		return NoState

	},

	// S169
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 181
		case r == 10: // ['\n','\n']
			return 181
		case r == 13: // ['\r','\r']
			return 181
		case r == 32: // [' ',' ']
			return 181
		case r == 39: // [''',''']
			return 182
		case r == 48: // ['0','0']
			return 183
		case 49 <= r && r <= 57: // ['1','9']
			return 184
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S170
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 186

		}
		return NoState

	},

	// S171
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 187

		}
		return NoState

	},

	// S172
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 188
		case r == 46: // ['.','.']
			return 189
		case r == 48: // ['0','0']
			return 190
		case 49 <= r && r <= 57: // ['1','9']
			return 191

		}
		return NoState

	},

	// S173
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 155
		case 48 <= r && r <= 57: // ['0','9']
			return 173
		case 65 <= r && r <= 70: // ['A','F']
			return 173
		case 97 <= r && r <= 102: // ['a','f']
			return 173

		}
		return NoState

	},

	// S174
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 107: // ['a','k']
			return 49
		case r == 108: // ['l','l']
			return 192
		case 109 <= r && r <= 122: // ['m','z']
			return 49

		}
		return NoState

	},

	// S175
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S176
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 175
		case 48 <= r && r <= 55: // ['0','7']
			return 176

		}
		return NoState

	},

	// S177
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 193
		case 65 <= r && r <= 70: // ['A','F']
			return 193
		case 97 <= r && r <= 102: // ['a','f']
			return 193

		}
		return NoState

	},

	// S178
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 175
		case 48 <= r && r <= 57: // ['0','9']
			return 178

		}
		return NoState

	},

	// S179
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 194
		case 65 <= r && r <= 70: // ['A','F']
			return 194
		case 97 <= r && r <= 102: // ['a','f']
			return 194

		}
		return NoState

	},

	// S180
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S181
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 181
		case r == 10: // ['\n','\n']
			return 181
		case r == 13: // ['\r','\r']
			return 181
		case r == 32: // [' ',' ']
			return 181
		case r == 39: // [''',''']
			return 182
		case r == 48: // ['0','0']
			return 183
		case 49 <= r && r <= 57: // ['1','9']
			return 184
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S182
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 195

		default:
			return 196
		}

	},

	// S183
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 55: // ['0','7']
			return 199
		case r == 88: // ['X','X']
			return 200
		case r == 120: // ['x','x']
			return 200
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S184
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 201
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S185
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S186
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S187
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S188
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 189
		case r == 48: // ['0','0']
			return 190
		case 49 <= r && r <= 57: // ['1','9']
			return 191

		}
		return NoState

	},

	// S189
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 202

		}
		return NoState

	},

	// S190
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case r == 46: // ['.','.']
			return 204
		case 48 <= r && r <= 55: // ['0','7']
			return 205
		case 56 <= r && r <= 57: // ['8','9']
			return 206
		case r == 69: // ['E','E']
			return 207
		case r == 88: // ['X','X']
			return 208
		case r == 101: // ['e','e']
			return 207
		case r == 120: // ['x','x']
			return 208

		}
		return NoState

	},

	// S191
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case r == 46: // ['.','.']
			return 204
		case 48 <= r && r <= 57: // ['0','9']
			return 191
		case r == 69: // ['E','E']
			return 207
		case r == 101: // ['e','e']
			return 207

		}
		return NoState

	},

	// S192
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 46
		case 65 <= r && r <= 90: // ['A','Z']
			return 47
		case r == 95: // ['_','_']
			return 48
		case 97 <= r && r <= 122: // ['a','z']
			return 49

		}
		return NoState

	},

	// S193
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 175
		case 48 <= r && r <= 57: // ['0','9']
			return 193
		case 65 <= r && r <= 70: // ['A','F']
			return 193
		case 97 <= r && r <= 102: // ['a','f']
			return 193

		}
		return NoState

	},

	// S194
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 209
		case 65 <= r && r <= 70: // ['A','F']
			return 209
		case 97 <= r && r <= 102: // ['a','f']
			return 209

		}
		return NoState

	},

	// S195
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 210
		case r == 39: // [''',''']
			return 210
		case 48 <= r && r <= 55: // ['0','7']
			return 211
		case r == 85: // ['U','U']
			return 212
		case r == 92: // ['\','\']
			return 210
		case r == 97: // ['a','a']
			return 210
		case r == 98: // ['b','b']
			return 210
		case r == 102: // ['f','f']
			return 210
		case r == 110: // ['n','n']
			return 210
		case r == 114: // ['r','r']
			return 210
		case r == 116: // ['t','t']
			return 210
		case r == 117: // ['u','u']
			return 213
		case r == 118: // ['v','v']
			return 210
		case r == 120: // ['x','x']
			return 214

		}
		return NoState

	},

	// S196
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S197
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S198
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 216
		case r == 10: // ['\n','\n']
			return 216
		case r == 13: // ['\r','\r']
			return 216
		case r == 32: // [' ',' ']
			return 216
		case r == 39: // [''',''']
			return 217
		case r == 48: // ['0','0']
			return 218
		case 49 <= r && r <= 57: // ['1','9']
			return 219

		}
		return NoState

	},

	// S199
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 55: // ['0','7']
			return 199
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S200
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 220
		case 65 <= r && r <= 70: // ['A','F']
			return 220
		case 97 <= r && r <= 102: // ['a','f']
			return 220

		}
		return NoState

	},

	// S201
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 201
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S202
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 202
		case r == 69: // ['E','E']
			return 221
		case r == 101: // ['e','e']
			return 221

		}
		return NoState

	},

	// S203
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S204
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 222
		case r == 69: // ['E','E']
			return 223
		case r == 101: // ['e','e']
			return 223

		}
		return NoState

	},

	// S205
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case r == 46: // ['.','.']
			return 204
		case 48 <= r && r <= 55: // ['0','7']
			return 205
		case 56 <= r && r <= 57: // ['8','9']
			return 206
		case r == 69: // ['E','E']
			return 207
		case r == 101: // ['e','e']
			return 207

		}
		return NoState

	},

	// S206
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 204
		case 48 <= r && r <= 57: // ['0','9']
			return 206
		case r == 69: // ['E','E']
			return 207
		case r == 101: // ['e','e']
			return 207

		}
		return NoState

	},

	// S207
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 224
		case r == 45: // ['-','-']
			return 224
		case 48 <= r && r <= 57: // ['0','9']
			return 225

		}
		return NoState

	},

	// S208
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 226
		case 65 <= r && r <= 70: // ['A','F']
			return 226
		case 97 <= r && r <= 102: // ['a','f']
			return 226

		}
		return NoState

	},

	// S209
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 227
		case 65 <= r && r <= 70: // ['A','F']
			return 227
		case 97 <= r && r <= 102: // ['a','f']
			return 227

		}
		return NoState

	},

	// S210
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S211
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 228

		}
		return NoState

	},

	// S212
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 229
		case 65 <= r && r <= 70: // ['A','F']
			return 229
		case 97 <= r && r <= 102: // ['a','f']
			return 229

		}
		return NoState

	},

	// S213
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 230
		case 65 <= r && r <= 70: // ['A','F']
			return 230
		case 97 <= r && r <= 102: // ['a','f']
			return 230

		}
		return NoState

	},

	// S214
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 231
		case 65 <= r && r <= 70: // ['A','F']
			return 231
		case 97 <= r && r <= 102: // ['a','f']
			return 231

		}
		return NoState

	},

	// S215
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S216
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 216
		case r == 10: // ['\n','\n']
			return 216
		case r == 13: // ['\r','\r']
			return 216
		case r == 32: // [' ',' ']
			return 216
		case r == 39: // [''',''']
			return 217
		case r == 48: // ['0','0']
			return 218
		case 49 <= r && r <= 57: // ['1','9']
			return 219

		}
		return NoState

	},

	// S217
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 232

		default:
			return 233
		}

	},

	// S218
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 55: // ['0','7']
			return 234
		case r == 88: // ['X','X']
			return 235
		case r == 120: // ['x','x']
			return 235
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S219
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 236
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S220
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 220
		case 65 <= r && r <= 70: // ['A','F']
			return 220
		case 97 <= r && r <= 102: // ['a','f']
			return 220
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S221
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 237
		case r == 45: // ['-','-']
			return 237
		case 48 <= r && r <= 57: // ['0','9']
			return 238

		}
		return NoState

	},

	// S222
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 222
		case r == 69: // ['E','E']
			return 239
		case r == 101: // ['e','e']
			return 239

		}
		return NoState

	},

	// S223
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 240
		case r == 45: // ['-','-']
			return 240
		case 48 <= r && r <= 57: // ['0','9']
			return 241

		}
		return NoState

	},

	// S224
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 225

		}
		return NoState

	},

	// S225
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 225

		}
		return NoState

	},

	// S226
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 226
		case 65 <= r && r <= 70: // ['A','F']
			return 226
		case 97 <= r && r <= 102: // ['a','f']
			return 226

		}
		return NoState

	},

	// S227
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 92: // ['\','\']
			return 36

		default:
			return 37
		}

	},

	// S228
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 242

		}
		return NoState

	},

	// S229
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 243
		case 65 <= r && r <= 70: // ['A','F']
			return 243
		case 97 <= r && r <= 102: // ['a','f']
			return 243

		}
		return NoState

	},

	// S230
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 244
		case 65 <= r && r <= 70: // ['A','F']
			return 244
		case 97 <= r && r <= 102: // ['a','f']
			return 244

		}
		return NoState

	},

	// S231
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 245
		case 65 <= r && r <= 70: // ['A','F']
			return 245
		case 97 <= r && r <= 102: // ['a','f']
			return 245

		}
		return NoState

	},

	// S232
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 246
		case r == 39: // [''',''']
			return 246
		case 48 <= r && r <= 55: // ['0','7']
			return 247
		case r == 85: // ['U','U']
			return 248
		case r == 92: // ['\','\']
			return 246
		case r == 97: // ['a','a']
			return 246
		case r == 98: // ['b','b']
			return 246
		case r == 102: // ['f','f']
			return 246
		case r == 110: // ['n','n']
			return 246
		case r == 114: // ['r','r']
			return 246
		case r == 116: // ['t','t']
			return 246
		case r == 117: // ['u','u']
			return 249
		case r == 118: // ['v','v']
			return 246
		case r == 120: // ['x','x']
			return 250

		}
		return NoState

	},

	// S233
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S234
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 55: // ['0','7']
			return 234
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S235
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 251
		case 65 <= r && r <= 70: // ['A','F']
			return 251
		case 97 <= r && r <= 102: // ['a','f']
			return 251

		}
		return NoState

	},

	// S236
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 236
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S237
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 238

		}
		return NoState

	},

	// S238
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 238

		}
		return NoState

	},

	// S239
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 252
		case r == 45: // ['-','-']
			return 252
		case 48 <= r && r <= 57: // ['0','9']
			return 253

		}
		return NoState

	},

	// S240
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 241

		}
		return NoState

	},

	// S241
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 241

		}
		return NoState

	},

	// S242
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S243
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 254
		case 65 <= r && r <= 70: // ['A','F']
			return 254
		case 97 <= r && r <= 102: // ['a','f']
			return 254

		}
		return NoState

	},

	// S244
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 255
		case 65 <= r && r <= 70: // ['A','F']
			return 255
		case 97 <= r && r <= 102: // ['a','f']
			return 255

		}
		return NoState

	},

	// S245
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S246
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S247
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 256

		}
		return NoState

	},

	// S248
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 257
		case 65 <= r && r <= 70: // ['A','F']
			return 257
		case 97 <= r && r <= 102: // ['a','f']
			return 257

		}
		return NoState

	},

	// S249
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 258
		case 65 <= r && r <= 70: // ['A','F']
			return 258
		case 97 <= r && r <= 102: // ['a','f']
			return 258

		}
		return NoState

	},

	// S250
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 259
		case 65 <= r && r <= 70: // ['A','F']
			return 259
		case 97 <= r && r <= 102: // ['a','f']
			return 259

		}
		return NoState

	},

	// S251
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 197
		case r == 10: // ['\n','\n']
			return 197
		case r == 13: // ['\r','\r']
			return 197
		case r == 32: // [' ',' ']
			return 197
		case r == 44: // [',',',']
			return 198
		case 48 <= r && r <= 57: // ['0','9']
			return 251
		case 65 <= r && r <= 70: // ['A','F']
			return 251
		case 97 <= r && r <= 102: // ['a','f']
			return 251
		case r == 125: // ['}','}']
			return 185

		}
		return NoState

	},

	// S252
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 253

		}
		return NoState

	},

	// S253
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 203
		case 48 <= r && r <= 57: // ['0','9']
			return 253

		}
		return NoState

	},

	// S254
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 260
		case 65 <= r && r <= 70: // ['A','F']
			return 260
		case 97 <= r && r <= 102: // ['a','f']
			return 260

		}
		return NoState

	},

	// S255
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 261
		case 65 <= r && r <= 70: // ['A','F']
			return 261
		case 97 <= r && r <= 102: // ['a','f']
			return 261

		}
		return NoState

	},

	// S256
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 262

		}
		return NoState

	},

	// S257
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 263
		case 65 <= r && r <= 70: // ['A','F']
			return 263
		case 97 <= r && r <= 102: // ['a','f']
			return 263

		}
		return NoState

	},

	// S258
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 264
		case 65 <= r && r <= 70: // ['A','F']
			return 264
		case 97 <= r && r <= 102: // ['a','f']
			return 264

		}
		return NoState

	},

	// S259
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 265
		case 65 <= r && r <= 70: // ['A','F']
			return 265
		case 97 <= r && r <= 102: // ['a','f']
			return 265

		}
		return NoState

	},

	// S260
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 266
		case 65 <= r && r <= 70: // ['A','F']
			return 266
		case 97 <= r && r <= 102: // ['a','f']
			return 266

		}
		return NoState

	},

	// S261
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S262
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S263
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 267
		case 65 <= r && r <= 70: // ['A','F']
			return 267
		case 97 <= r && r <= 102: // ['a','f']
			return 267

		}
		return NoState

	},

	// S264
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 268
		case 65 <= r && r <= 70: // ['A','F']
			return 268
		case 97 <= r && r <= 102: // ['a','f']
			return 268

		}
		return NoState

	},

	// S265
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S266
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 269
		case 65 <= r && r <= 70: // ['A','F']
			return 269
		case 97 <= r && r <= 102: // ['a','f']
			return 269

		}
		return NoState

	},

	// S267
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 270
		case 65 <= r && r <= 70: // ['A','F']
			return 270
		case 97 <= r && r <= 102: // ['a','f']
			return 270

		}
		return NoState

	},

	// S268
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 271
		case 65 <= r && r <= 70: // ['A','F']
			return 271
		case 97 <= r && r <= 102: // ['a','f']
			return 271

		}
		return NoState

	},

	// S269
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 272
		case 65 <= r && r <= 70: // ['A','F']
			return 272
		case 97 <= r && r <= 102: // ['a','f']
			return 272

		}
		return NoState

	},

	// S270
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 273
		case 65 <= r && r <= 70: // ['A','F']
			return 273
		case 97 <= r && r <= 102: // ['a','f']
			return 273

		}
		return NoState

	},

	// S271
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S272
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 274
		case 65 <= r && r <= 70: // ['A','F']
			return 274
		case 97 <= r && r <= 102: // ['a','f']
			return 274

		}
		return NoState

	},

	// S273
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 275
		case 65 <= r && r <= 70: // ['A','F']
			return 275
		case 97 <= r && r <= 102: // ['a','f']
			return 275

		}
		return NoState

	},

	// S274
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},

	// S275
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 276
		case 65 <= r && r <= 70: // ['A','F']
			return 276
		case 97 <= r && r <= 102: // ['a','f']
			return 276

		}
		return NoState

	},

	// S276
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 277
		case 65 <= r && r <= 70: // ['A','F']
			return 277
		case 97 <= r && r <= 102: // ['a','f']
			return 277

		}
		return NoState

	},

	// S277
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 215

		}
		return NoState

	},
}

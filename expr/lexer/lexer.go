package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/expr/util"

	"github.com/katydid/katydid/expr/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 289
	NumSymbols = 199
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {

	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)

	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {

		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)

		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '('
1: ')'
2: '0'
3: '-'
4: '('
5: ')'
6: '('
7: '-'
8: ')'
9: '$'
10: '$'
11: '$'
12: '$'
13: '$'
14: '$'
15: '{'
16: ','
17: '}'
18: '['
19: ']'
20: 'b'
21: 'o'
22: 'o'
23: 'l'
24: '['
25: ']'
26: 'i'
27: 'n'
28: 't'
29: '['
30: ']'
31: 'u'
32: 'i'
33: 'n'
34: 't'
35: '['
36: ']'
37: 'd'
38: 'o'
39: 'u'
40: 'b'
41: 'l'
42: 'e'
43: '['
44: ']'
45: 's'
46: 't'
47: 'r'
48: 'i'
49: 'n'
50: 'g'
51: '['
52: ']'
53: '['
54: ']'
55: 'b'
56: 'y'
57: 't'
58: 'e'
59: 't'
60: 'r'
61: 'u'
62: 'e'
63: 'f'
64: 'a'
65: 'l'
66: 's'
67: 'e'
68: '='
69: '('
70: ')'
71: '{'
72: '}'
73: ','
74: ';'
75: '#'
76: '&'
77: '|'
78: '['
79: ']'
80: ':'
81: '!'
82: '*'
83: '_'
84: '~'
85: '.'
86: '@'
87: '-'
88: '>'
89: '='
90: '='
91: '!'
92: '='
93: '<'
94: '>'
95: '<'
96: '='
97: '>'
98: '='
99: '~'
100: '='
101: '*'
102: '='
103: '^'
104: '='
105: '$'
106: '='
107: ':'
108: ':'
109: '?'
110: '/'
111: '/'
112: '\n'
113: '/'
114: '*'
115: '*'
116: '*'
117: '/'
118: ' '
119: '\t'
120: '\n'
121: '\r'
122: '0'
123: '0'
124: 'x'
125: 'X'
126: '-'
127: 'e'
128: 'E'
129: '+'
130: '-'
131: '.'
132: '.'
133: '.'
134: '_'
135: '_'
136: 'd'
137: 'o'
138: 'u'
139: 'b'
140: 'l'
141: 'e'
142: 'i'
143: 'n'
144: 't'
145: 'u'
146: 'i'
147: 'n'
148: 't'
149: '['
150: ']'
151: 'b'
152: 'y'
153: 't'
154: 'e'
155: 's'
156: 't'
157: 'r'
158: 'i'
159: 'n'
160: 'g'
161: 'b'
162: 'o'
163: 'o'
164: 'l'
165: '.'
166: '\'
167: 'U'
168: '\'
169: 'u'
170: '\'
171: 'x'
172: '\'
173: '`'
174: '`'
175: '\'
176: 'a'
177: 'b'
178: 'f'
179: 'n'
180: 'r'
181: 't'
182: 'v'
183: '\'
184: '''
185: '"'
186: '"'
187: '"'
188: '''
189: '''
190: '0'-'9'
191: '0'-'7'
192: '0'-'9'
193: 'A'-'F'
194: 'a'-'f'
195: '1'-'9'
196: 'A'-'Z'
197: 'a'-'z'
198: .

*/

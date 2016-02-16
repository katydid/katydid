package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/relapse/util"

	"github.com/katydid/katydid/relapse/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 295
	NumSymbols = 206
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
18: '<'
19: 'e'
20: 'm'
21: 'p'
22: 't'
23: 'y'
24: '>'
25: '['
26: ']'
27: 'b'
28: 'o'
29: 'o'
30: 'l'
31: '['
32: ']'
33: 'i'
34: 'n'
35: 't'
36: '['
37: ']'
38: 'u'
39: 'i'
40: 'n'
41: 't'
42: '['
43: ']'
44: 'd'
45: 'o'
46: 'u'
47: 'b'
48: 'l'
49: 'e'
50: '['
51: ']'
52: 's'
53: 't'
54: 'r'
55: 'i'
56: 'n'
57: 'g'
58: '['
59: ']'
60: '['
61: ']'
62: 'b'
63: 'y'
64: 't'
65: 'e'
66: 't'
67: 'r'
68: 'u'
69: 'e'
70: 'f'
71: 'a'
72: 'l'
73: 's'
74: 'e'
75: '='
76: '('
77: ')'
78: '{'
79: '}'
80: ','
81: ';'
82: '#'
83: '&'
84: '|'
85: '['
86: ']'
87: ':'
88: '!'
89: '*'
90: '_'
91: '~'
92: '.'
93: '@'
94: '-'
95: '>'
96: '='
97: '='
98: '!'
99: '='
100: '<'
101: '>'
102: '<'
103: '='
104: '>'
105: '='
106: '~'
107: '='
108: '*'
109: '='
110: '^'
111: '='
112: '$'
113: '='
114: ':'
115: ':'
116: '?'
117: '/'
118: '/'
119: '\n'
120: '/'
121: '*'
122: '*'
123: '*'
124: '/'
125: ' '
126: '\t'
127: '\n'
128: '\r'
129: '0'
130: '0'
131: 'x'
132: 'X'
133: '-'
134: 'e'
135: 'E'
136: '+'
137: '-'
138: '.'
139: '.'
140: '.'
141: '_'
142: '_'
143: 'd'
144: 'o'
145: 'u'
146: 'b'
147: 'l'
148: 'e'
149: 'i'
150: 'n'
151: 't'
152: 'u'
153: 'i'
154: 'n'
155: 't'
156: '['
157: ']'
158: 'b'
159: 'y'
160: 't'
161: 'e'
162: 's'
163: 't'
164: 'r'
165: 'i'
166: 'n'
167: 'g'
168: 'b'
169: 'o'
170: 'o'
171: 'l'
172: '.'
173: '\'
174: 'U'
175: '\'
176: 'u'
177: '\'
178: 'x'
179: '\'
180: '`'
181: '`'
182: '\'
183: 'a'
184: 'b'
185: 'f'
186: 'n'
187: 'r'
188: 't'
189: 'v'
190: '\'
191: '''
192: '"'
193: '"'
194: '"'
195: '''
196: '''
197: '0'-'9'
198: '0'-'7'
199: '0'-'9'
200: 'A'-'F'
201: 'a'-'f'
202: '1'-'9'
203: 'A'-'Z'
204: 'a'-'z'
205: .

*/

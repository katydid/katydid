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
	NumStates  = 299
	NumSymbols = 216
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
25: '<'
26: 'e'
27: 'm'
28: 'p'
29: 't'
30: 'y'
31: 's'
32: 'e'
33: 't'
34: '>'
35: '['
36: ']'
37: 'b'
38: 'o'
39: 'o'
40: 'l'
41: '['
42: ']'
43: 'i'
44: 'n'
45: 't'
46: '['
47: ']'
48: 'u'
49: 'i'
50: 'n'
51: 't'
52: '['
53: ']'
54: 'd'
55: 'o'
56: 'u'
57: 'b'
58: 'l'
59: 'e'
60: '['
61: ']'
62: 's'
63: 't'
64: 'r'
65: 'i'
66: 'n'
67: 'g'
68: '['
69: ']'
70: '['
71: ']'
72: 'b'
73: 'y'
74: 't'
75: 'e'
76: 't'
77: 'r'
78: 'u'
79: 'e'
80: 'f'
81: 'a'
82: 'l'
83: 's'
84: 'e'
85: '='
86: '('
87: ')'
88: '{'
89: '}'
90: ','
91: ';'
92: '#'
93: '&'
94: '|'
95: '['
96: ']'
97: ':'
98: '!'
99: '*'
100: '_'
101: '~'
102: '.'
103: '@'
104: '-'
105: '>'
106: '='
107: '='
108: '!'
109: '='
110: '<'
111: '>'
112: '<'
113: '='
114: '>'
115: '='
116: '~'
117: '='
118: '*'
119: '='
120: '^'
121: '='
122: '$'
123: '='
124: ':'
125: ':'
126: '?'
127: '/'
128: '/'
129: '\n'
130: '/'
131: '*'
132: '*'
133: '*'
134: '/'
135: ' '
136: '\t'
137: '\n'
138: '\r'
139: '0'
140: '0'
141: 'x'
142: 'X'
143: '-'
144: 'e'
145: 'E'
146: '+'
147: '-'
148: '.'
149: '.'
150: '.'
151: '_'
152: '_'
153: 'd'
154: 'o'
155: 'u'
156: 'b'
157: 'l'
158: 'e'
159: 'i'
160: 'n'
161: 't'
162: 'u'
163: 'i'
164: 'n'
165: 't'
166: '['
167: ']'
168: 'b'
169: 'y'
170: 't'
171: 'e'
172: 's'
173: 't'
174: 'r'
175: 'i'
176: 'n'
177: 'g'
178: 'b'
179: 'o'
180: 'o'
181: 'l'
182: '.'
183: '\'
184: 'U'
185: '\'
186: 'u'
187: '\'
188: 'x'
189: '\'
190: '`'
191: '`'
192: '\'
193: 'a'
194: 'b'
195: 'f'
196: 'n'
197: 'r'
198: 't'
199: 'v'
200: '\'
201: '''
202: '"'
203: '"'
204: '"'
205: '''
206: '''
207: '0'-'9'
208: '0'-'7'
209: '0'-'9'
210: 'A'-'F'
211: 'a'-'f'
212: '1'-'9'
213: 'A'-'Z'
214: 'a'-'z'
215: .

*/

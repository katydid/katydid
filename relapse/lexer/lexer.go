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
	NumStates  = 291
	NumSymbols = 226
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
2: '('
3: ')'
4: '('
5: '-'
6: ')'
7: '$'
8: '$'
9: '$'
10: '$'
11: '$'
12: '$'
13: '{'
14: ','
15: '}'
16: 'A'
17: 'n'
18: 'y'
19: 'N'
20: 'a'
21: 'm'
22: 'e'
23: 'N'
24: 'a'
25: 'm'
26: 'e'
27: 'A'
28: 'n'
29: 'y'
30: 'N'
31: 'a'
32: 'm'
33: 'e'
34: 'E'
35: 'x'
36: 'c'
37: 'e'
38: 'p'
39: 't'
40: 'N'
41: 'a'
42: 'm'
43: 'e'
44: 'C'
45: 'h'
46: 'o'
47: 'i'
48: 'c'
49: 'e'
50: 'E'
51: 'm'
52: 'p'
53: 't'
54: 'y'
55: 'E'
56: 'm'
57: 'p'
58: 't'
59: 'y'
60: 'S'
61: 'e'
62: 't'
63: 'Z'
64: 'e'
65: 'r'
66: 'o'
67: 'O'
68: 'r'
69: 'M'
70: 'o'
71: 'r'
72: 'e'
73: '['
74: ']'
75: 'b'
76: 'o'
77: 'o'
78: 'l'
79: '['
80: ']'
81: 'i'
82: 'n'
83: 't'
84: '['
85: ']'
86: 'u'
87: 'i'
88: 'n'
89: 't'
90: '['
91: ']'
92: 'd'
93: 'o'
94: 'u'
95: 'b'
96: 'l'
97: 'e'
98: '['
99: ']'
100: 's'
101: 't'
102: 'r'
103: 'i'
104: 'n'
105: 'g'
106: '['
107: ']'
108: '['
109: ']'
110: 'b'
111: 'y'
112: 't'
113: 'e'
114: 't'
115: 'r'
116: 'u'
117: 'e'
118: 'f'
119: 'a'
120: 'l'
121: 's'
122: 'e'
123: '='
124: '('
125: ')'
126: '{'
127: '}'
128: ','
129: ';'
130: '#'
131: '&'
132: '|'
133: '['
134: ']'
135: ':'
136: '!'
137: '/'
138: '/'
139: '\n'
140: '/'
141: '*'
142: '*'
143: '*'
144: '/'
145: ' '
146: '\t'
147: '\n'
148: '\r'
149: '0'
150: '0'
151: 'x'
152: 'X'
153: '-'
154: 'e'
155: 'E'
156: '+'
157: '-'
158: '.'
159: '.'
160: '.'
161: '_'
162: '_'
163: 'd'
164: 'o'
165: 'u'
166: 'b'
167: 'l'
168: 'e'
169: 'i'
170: 'n'
171: 't'
172: 'u'
173: 'i'
174: 'n'
175: 't'
176: '['
177: ']'
178: 'b'
179: 'y'
180: 't'
181: 'e'
182: 's'
183: 't'
184: 'r'
185: 'i'
186: 'n'
187: 'g'
188: 'b'
189: 'o'
190: 'o'
191: 'l'
192: '.'
193: '\'
194: 'U'
195: '\'
196: 'u'
197: '\'
198: 'x'
199: '\'
200: '`'
201: '`'
202: '\'
203: 'a'
204: 'b'
205: 'f'
206: 'n'
207: 'r'
208: 't'
209: 'v'
210: '\'
211: '''
212: '"'
213: '"'
214: '"'
215: '''
216: '''
217: '0'-'9'
218: '0'-'7'
219: '0'-'9'
220: 'A'-'F'
221: 'a'-'f'
222: '1'-'9'
223: 'A'-'Z'
224: 'a'-'z'
225: .

*/

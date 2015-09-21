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
	NumStates  = 282
	NumSymbols = 213
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
23: 'A'
24: 'n'
25: 'y'
26: 'N'
27: 'a'
28: 'm'
29: 'e'
30: 'E'
31: 'x'
32: 'c'
33: 'e'
34: 'p'
35: 't'
36: 'N'
37: 'a'
38: 'm'
39: 'e'
40: 'C'
41: 'h'
42: 'o'
43: 'i'
44: 'c'
45: 'e'
46: 'E'
47: 'm'
48: 'p'
49: 't'
50: 'y'
51: 'E'
52: 'm'
53: 'p'
54: 't'
55: 'y'
56: 'S'
57: 'e'
58: 't'
59: '['
60: ']'
61: 'b'
62: 'o'
63: 'o'
64: 'l'
65: '['
66: ']'
67: 'i'
68: 'n'
69: 't'
70: '['
71: ']'
72: 'u'
73: 'i'
74: 'n'
75: 't'
76: '['
77: ']'
78: 'd'
79: 'o'
80: 'u'
81: 'b'
82: 'l'
83: 'e'
84: '['
85: ']'
86: 's'
87: 't'
88: 'r'
89: 'i'
90: 'n'
91: 'g'
92: '['
93: ']'
94: '['
95: ']'
96: 'b'
97: 'y'
98: 't'
99: 'e'
100: 't'
101: 'r'
102: 'u'
103: 'e'
104: 'f'
105: 'a'
106: 'l'
107: 's'
108: 'e'
109: '='
110: '('
111: ')'
112: '{'
113: '}'
114: ','
115: ';'
116: '#'
117: '&'
118: '|'
119: '['
120: ']'
121: ':'
122: '!'
123: '*'
124: '/'
125: '/'
126: '\n'
127: '/'
128: '*'
129: '*'
130: '*'
131: '/'
132: ' '
133: '\t'
134: '\n'
135: '\r'
136: '0'
137: '0'
138: 'x'
139: 'X'
140: '-'
141: 'e'
142: 'E'
143: '+'
144: '-'
145: '.'
146: '.'
147: '.'
148: '_'
149: '_'
150: 'd'
151: 'o'
152: 'u'
153: 'b'
154: 'l'
155: 'e'
156: 'i'
157: 'n'
158: 't'
159: 'u'
160: 'i'
161: 'n'
162: 't'
163: '['
164: ']'
165: 'b'
166: 'y'
167: 't'
168: 'e'
169: 's'
170: 't'
171: 'r'
172: 'i'
173: 'n'
174: 'g'
175: 'b'
176: 'o'
177: 'o'
178: 'l'
179: '.'
180: '\'
181: 'U'
182: '\'
183: 'u'
184: '\'
185: 'x'
186: '\'
187: '`'
188: '`'
189: '\'
190: 'a'
191: 'b'
192: 'f'
193: 'n'
194: 'r'
195: 't'
196: 'v'
197: '\'
198: '''
199: '"'
200: '"'
201: '"'
202: '''
203: '''
204: '0'-'9'
205: '0'-'7'
206: '0'-'9'
207: 'A'-'F'
208: 'a'-'f'
209: '1'-'9'
210: 'A'-'Z'
211: 'a'-'z'
212: .

*/

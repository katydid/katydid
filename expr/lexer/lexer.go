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
	NumStates  = 288
	NumSymbols = 200
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
9: '-'
10: '$'
11: '$'
12: '$'
13: '$'
14: '$'
15: '$'
16: '{'
17: ','
18: '}'
19: '['
20: ']'
21: 'b'
22: 'o'
23: 'o'
24: 'l'
25: '['
26: ']'
27: 'i'
28: 'n'
29: 't'
30: '['
31: ']'
32: 'u'
33: 'i'
34: 'n'
35: 't'
36: '['
37: ']'
38: 'd'
39: 'o'
40: 'u'
41: 'b'
42: 'l'
43: 'e'
44: '['
45: ']'
46: 's'
47: 't'
48: 'r'
49: 'i'
50: 'n'
51: 'g'
52: '['
53: ']'
54: '['
55: ']'
56: 'b'
57: 'y'
58: 't'
59: 'e'
60: 't'
61: 'r'
62: 'u'
63: 'e'
64: 'f'
65: 'a'
66: 'l'
67: 's'
68: 'e'
69: '='
70: '('
71: ')'
72: '{'
73: '}'
74: ','
75: ';'
76: '#'
77: '&'
78: '|'
79: '['
80: ']'
81: ':'
82: '!'
83: '*'
84: '_'
85: '~'
86: '.'
87: '@'
88: '-'
89: '>'
90: '='
91: '='
92: '!'
93: '='
94: '<'
95: '>'
96: '<'
97: '='
98: '>'
99: '='
100: '~'
101: '='
102: '*'
103: '='
104: '^'
105: '='
106: '$'
107: '='
108: ':'
109: ':'
110: '?'
111: '/'
112: '/'
113: '\n'
114: '/'
115: '*'
116: '*'
117: '*'
118: '/'
119: ' '
120: '\t'
121: '\n'
122: '\r'
123: '0'
124: '0'
125: 'x'
126: 'X'
127: '-'
128: 'e'
129: 'E'
130: '+'
131: '-'
132: '.'
133: '.'
134: '.'
135: '_'
136: '_'
137: 'd'
138: 'o'
139: 'u'
140: 'b'
141: 'l'
142: 'e'
143: 'i'
144: 'n'
145: 't'
146: 'u'
147: 'i'
148: 'n'
149: 't'
150: '['
151: ']'
152: 'b'
153: 'y'
154: 't'
155: 'e'
156: 's'
157: 't'
158: 'r'
159: 'i'
160: 'n'
161: 'g'
162: 'b'
163: 'o'
164: 'o'
165: 'l'
166: '.'
167: '\'
168: 'U'
169: '\'
170: 'u'
171: '\'
172: 'x'
173: '\'
174: '`'
175: '`'
176: '\'
177: 'a'
178: 'b'
179: 'f'
180: 'n'
181: 'r'
182: 't'
183: 'v'
184: '\'
185: '''
186: '"'
187: '"'
188: '"'
189: '''
190: '''
191: '0'-'9'
192: '0'-'7'
193: '0'-'9'
194: 'A'-'F'
195: 'a'-'f'
196: '1'-'9'
197: 'A'-'Z'
198: 'a'-'z'
199: .

*/

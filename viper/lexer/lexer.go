package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/viper/util"

	"github.com/katydid/katydid/viper/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 279
	NumSymbols = 203
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
16: 's'
17: 't'
18: 'a'
19: 'r'
20: 't'
21: 'f'
22: 'i'
23: 'n'
24: 'a'
25: 'l'
26: 'i'
27: 'n'
28: 't'
29: 'e'
30: 'r'
31: 'n'
32: 'a'
33: 'l'
34: 'c'
35: 'a'
36: 'l'
37: 'l'
38: 'r'
39: 'e'
40: 't'
41: 'u'
42: 'r'
43: 'n'
44: '['
45: ']'
46: 'b'
47: 'o'
48: 'o'
49: 'l'
50: '['
51: ']'
52: 'i'
53: 'n'
54: 't'
55: '['
56: ']'
57: 'u'
58: 'i'
59: 'n'
60: 't'
61: '['
62: ']'
63: 'd'
64: 'o'
65: 'u'
66: 'b'
67: 'l'
68: 'e'
69: '['
70: ']'
71: 's'
72: 't'
73: 'r'
74: 'i'
75: 'n'
76: 'g'
77: '['
78: ']'
79: '['
80: ']'
81: 'b'
82: 'y'
83: 't'
84: 'e'
85: 't'
86: 'r'
87: 'u'
88: 'e'
89: 'f'
90: 'a'
91: 'l'
92: 's'
93: 'e'
94: '='
95: '('
96: ')'
97: '{'
98: '}'
99: ','
100: ';'
101: '#'
102: '&'
103: '|'
104: '['
105: ']'
106: ':'
107: '!'
108: '*'
109: '_'
110: '~'
111: '.'
112: '@'
113: '>'
114: '/'
115: '/'
116: '\n'
117: '/'
118: '*'
119: '*'
120: '*'
121: '/'
122: ' '
123: '\t'
124: '\n'
125: '\r'
126: '0'
127: '0'
128: 'x'
129: 'X'
130: '-'
131: 'e'
132: 'E'
133: '+'
134: '-'
135: '.'
136: '.'
137: '.'
138: '_'
139: '_'
140: 'd'
141: 'o'
142: 'u'
143: 'b'
144: 'l'
145: 'e'
146: 'i'
147: 'n'
148: 't'
149: 'u'
150: 'i'
151: 'n'
152: 't'
153: '['
154: ']'
155: 'b'
156: 'y'
157: 't'
158: 'e'
159: 's'
160: 't'
161: 'r'
162: 'i'
163: 'n'
164: 'g'
165: 'b'
166: 'o'
167: 'o'
168: 'l'
169: '.'
170: '\'
171: 'U'
172: '\'
173: 'u'
174: '\'
175: 'x'
176: '\'
177: '`'
178: '`'
179: '\'
180: 'a'
181: 'b'
182: 'f'
183: 'n'
184: 'r'
185: 't'
186: 'v'
187: '\'
188: '''
189: '"'
190: '"'
191: '"'
192: '''
193: '''
194: '0'-'9'
195: '0'-'7'
196: '0'-'9'
197: 'A'-'F'
198: 'a'-'f'
199: '1'-'9'
200: 'A'-'Z'
201: 'a'-'z'
202: .

*/

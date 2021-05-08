//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Package yaml contains the implementation of a YAML parser.
package yaml

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/katydid/katydid/parser"
)

//YamlParser is a parser for YAML
type YamlParser interface {
	parser.Interface
	//Init initialises the parser with a byte buffer containing YAML.
	Init(buf []byte) error
	Reset() error
}

//NewYamlParser returns a new YAML parser.
func NewYamlParser() YamlParser {
	return &yamlParser{
		state: state{
			firstObjectValue: true,
		},
		stack: make([]state, 0, 10),
	}
}

func (s *yamlParser) Init(buf []byte) error {
	s.state = state{
		firstObjectValue: true,
		buf:              buf,
	}
	// empty stack in case this is an initialization due to a reset
	s.stack = s.stack[:0]
	if err := s.skipEmptyLines(); err != nil {
		return err
	}
	line, err := s.peekNextLine()
	if err != nil {
		return err
	}

	if line.isArray {
		if err := s.scanValue(); err != nil {
			return err
		}
		s.inArray = true
		s.firstArrayValue = true
		s.buf = s.buf[s.startValueOffset:s.endValueOffset]
		s.offset = 0
	}
	return nil
}

func (s *yamlParser) scanValue() error {
	line, err := s.peekNextLine()
	if err != nil {
		return err
	}
	if line.isArray {
		return s.scanArray()
	}

	if line.hasValue { // tighten this condition
		value := []byte(strings.TrimSpace(string(line.value)))
		if len(value) > 0 {
			return s.scanString()
		}
	}

	return s.scanString()
	// return s.expected("objects not supported")
}

func (s *yamlParser) expected(expected string) error {
	return fmt.Errorf("katydid/yaml/parser: expected '%s' at offset %d, but got '%c'", expected, s.offset, s.buf[s.offset])
}

func scanString(buf []byte) (int, error) {
	for i, c := range buf[1:] {
		if c == '\n' {
			return i, nil
		}
	}
	return len(buf), nil
}

func (s *yamlParser) scanString() error {
	line, err := s.scanLine()
	if err != nil {
		return err
	}

	s.startValueOffset = line.valueOffset
	s.endValueOffset = s.offset

	return nil
}

func (s *yamlParser) Reset() error {
	if len(s.stack) > 0 {
		return s.Init(s.stack[0].buf)
	}
	return s.Init(s.buf)
}

func (s *yamlParser) Next() error {
	if s.isLeaf {
		if s.firstObjectValue {
			s.firstObjectValue = false
			return nil
		}
		return io.EOF
	}
	s.isValueObject = false
	s.isValueArray = false

	err := s.skipEmptyLines()
	if err != nil {
		return io.EOF
	}

	if s.inArray {
		if !s.firstArrayValue {
			s.arrayIndex++
		}
		return s.nextValueInArray()
	}
	return s.nextValueInObject()
}

func (s *yamlParser) nextValueInObject() error {
	if s.firstObjectValue {
		s.firstObjectValue = false
	}

	line, err := s.scanLine()
	if err != nil {
		return io.EOF
	}

	if !line.hasValue {
		s.name = string(line.name)
		s.startValueOffset = s.offset
		s.endValueOffset = s.offset
		return nil
	}

	// has inline value
	if len(strings.TrimSpace(string(line.value))) > 0 {
		s.name = string(line.name)
		s.startValueOffset = line.valueOffset
		s.endValueOffset = s.offset
		return nil
	}

	nextLine, err := s.peekNextLine()
	if err != nil { // no next line
		s.name = string(line.name)
		s.startValueOffset = s.offset
		s.endValueOffset = s.offset
		return nil
	}

	// is child
	if len(nextLine.indent) > len(line.indent) {
		s.name = string(line.name)
		if nextLine.isArray {
			return s.scanArray()
		}
		return s.scanObjectValue()
	}

	s.name = string(line.name)
	s.startValueOffset = line.valueOffset
	s.endValueOffset = s.offset
	return nil
}

func (s *yamlParser) nextValueInArray() error {
	if s.firstArrayValue {
		s.firstArrayValue = false
	}

	line, err := s.scanLine()
	if err != nil {
		return io.EOF
	}

	// check if this is last element of array
	// inadequate
	// if !line.isArray {
	// 	return io.EOF
	// }

	if !line.hasValue { // and its end of line
		s.startValueOffset = line.nameStart
		s.endValueOffset = s.offset
		return nil
	}

	// has inline value
	if len(strings.TrimSpace(string(line.value))) > 0 {
		s.startValueOffset = line.nameStart
		s.endValueOffset = s.offset
		s.isValueObject = true
		return nil
	}

	// has really no value but could
	// have objects

	nextLine, err := s.peekNextLine()
	if err != nil { // no next line
		s.startValueOffset = line.nameStart
		s.endValueOffset = s.offset
		return nil
	}

	// is child
	if len(nextLine.indent) > len(line.indent) {
		s.name = string(line.name)
		if nextLine.isArray {
			err := s.scanArray()
			if err != nil {
				return err
			}
			s.startValueOffset = line.nameStart
			s.isValueArray = false
			s.isValueObject = true
			return nil
		}
		err := s.scanObjectValue()
		if err != nil {
			return err
		}
		s.startValueOffset = line.nameStart
		s.isValueObject = true
		return nil
	}

	s.startValueOffset = line.nameStart
	s.endValueOffset = s.offset
	s.isValueObject = true
	return nil
}

func (s *yamlParser) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *yamlParser) Down() {
	if s.isValueObject {
		s.stack = append(s.stack, s.state)
		s.state = state{
			buf:              s.buf[s.startValueOffset:s.endValueOffset],
			firstObjectValue: true,
		}
	} else if s.isValueArray {
		s.stack = append(s.stack, s.state)
		s.state = state{
			buf:             s.buf[s.startValueOffset:s.endValueOffset],
			firstArrayValue: true,
			inArray:         true,
		}
	} else {
		s.stack = append(s.stack, s.state)
		s.state.isLeaf = true
		s.state.firstObjectValue = true
	}
}

type yamlParser struct {
	state
	stack []state
}

type state struct {
	buf              []byte
	offset           int
	name             string
	startValueOffset int
	endValueOffset   int
	inArray          bool
	firstObjectValue bool
	firstArrayValue  bool
	isValueObject    bool
	isValueArray     bool
	isLeaf           bool
	arrayIndex       int
}

func (s *yamlParser) Bool() (bool, error) {
	if s.isLeaf {
		v := string(s.Value())
		if v == "true" {
			return true, nil
		}
		if v == "false" {
			return false, nil
		}
	}
	return false, parser.ErrNotBool
}

func (s *yamlParser) Bytes() ([]byte, error) {
	return nil, parser.ErrNotBytes
}

func (s *yamlParser) Double() (float64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseFloat(v, 64)
		return i, err
	}
	return 0, parser.ErrNotDouble
}

func (s *yamlParser) Int() (int64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			f, ferr := strconv.ParseFloat(v, 64)
			if ferr != nil {
				return i, err
			}
			if float64(int64(f)) == f {
				return int64(f), nil
			}
		}
		return i, err
	}
	if s.inArray {
		return int64(s.arrayIndex), nil
	}
	return 0, parser.ErrNotInt
}

func (s *yamlParser) IsLeaf() bool {
	return s.isLeaf
}

func (s *yamlParser) String() (string, error) {
	if s.isLeaf {
		v := s.Value()
		s := strings.TrimSpace(string(v))
		if len(s) == 0 {
			// for nulls
			return "", parser.ErrNotString
		}
		return strings.TrimSpace(string(v)), nil
	}
	if s.inArray {
		return "", parser.ErrNotString
	}
	return s.name, nil
}

func (s *yamlParser) Uint() (uint64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseUint(v, 10, 64)
		return uint64(i), err
	}
	return 0, parser.ErrNotUint
}

func (s *yamlParser) Value() []byte {
	return s.buf[s.startValueOffset:s.endValueOffset]
}

func isSpace(c byte) bool {
	return (c == ' ') || (c == '\n') || (c == '\r') || (c == '\t')
}

func skipSpace(buf []byte) int {
	for i, c := range buf {
		if !isSpace(c) {
			return i
		}
	}
	return len(buf)
}

func skipEmptyLines(buf []byte) int {
	offset := 0
	for i, c := range buf {
		if i > 0 && buf[i-1] == '\n' {
			offset = i
		}

		if !isSpace(c) {
			return offset
		}
	}
	return len(buf)
}

func (s *yamlParser) skipEmptyLines() error {
	if s.offset >= len(s.buf) {
		return io.ErrShortBuffer
	}
	s.offset += skipEmptyLines(s.buf[s.offset:])
	if s.offset >= len(s.buf) {
		return io.ErrShortBuffer
	}
	return nil
}

func (s *yamlParser) nextLine() error {
	offset := s.offset
	err := s.skipEmptyLines()
	if err != nil {
		return err
	}
	if offset == s.offset {
		return io.EOF
	}
	return nil
}

type Line struct {
	indent      []byte
	isArray     bool
	isEmpty     bool
	name        []byte
	value       []byte
	hasValue    bool
	valueOffset int
	nameStart   int
}

func (s *yamlParser) peekNextLine() (*Line, error) {
	err := s.skipEmptyLines()
	if err != nil {
		return nil, err
	}

	if s.offset >= len(s.buf) {
		return nil, io.ErrShortBuffer
	}

	var indent []byte
	i := 0

	buf := s.buf[s.offset:]
	isArray := false

	// scan indent
	for ; i < len(buf); i++ {
		c := buf[i]
		if c == '\n' {
			return &Line{isEmpty: true}, nil
		}
		if !isSpace(c) {
			indent = buf[:i]
			if c == '-' {
				if i == len(buf)-1 || isSpace(buf[i+1]) {
					isArray = true
				}
			}
			break
		}
	}

	if isArray {
		// skip dash
		i++
		// skip spaces
		for ; i < len(buf); i++ {
			if buf[i] == '\n' {
				return &Line{
						isArray: true, indent: indent,
					},
					nil
			}
			if !isSpace(buf[i]) {
				break
			}
		}
	}

	var name []byte

	// scan name
	nameStart := i + s.offset
	j := i
	for ; j < len(buf); j++ {
		if buf[j] == '\n' {
			name = buf[i:j] // and remember to trim
			return &Line{
				name:      name,
				isArray:   isArray,
				indent:    indent,
				nameStart: nameStart,
			}, nil
		}
		if buf[j] == ':' {
			name = buf[i:j] // and remember to trim
			break
		}
		if j == len(buf)-1 {
			j++
			name = buf[i:j] // and remember to trim
			return &Line{
				name:      name,
				isArray:   isArray,
				indent:    indent,
				hasValue:  false,
				nameStart: nameStart,
			}, nil
		}
	}

	// scan value
	// skip colon
	j++
	k := j
	for ; k < len(buf); k++ {
		if buf[k] == '\n' {
			return &Line{
					name:      name,
					isArray:   isArray,
					value:     buf[j:k], // trim this
					indent:    indent,
					hasValue:  true,
					nameStart: nameStart,
				},
				nil
		}
	}

	return &Line{
			name:      name,
			isArray:   isArray,
			value:     buf[j:k], // trim this
			indent:    indent,
			hasValue:  true,
			nameStart: nameStart,
		},
		nil
}

func isNumber(c byte) bool {
	return (c == '-') || ((c >= '0') && (c <= '9'))
}

func (s *yamlParser) scanNumber() error {
	s.startValueOffset = s.offset
	if s.buf[s.offset] == '-' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
	}
	if s.buf[s.offset] == '0' {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
	} else if isDigit19(s.buf[s.offset]) {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.offset < len(s.buf) && s.buf[s.offset] == '.' {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.offset < len(s.buf) &&
		(s.buf[s.offset] == 'e' || s.buf[s.offset] == 'E') {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		if s.offset < len(s.buf) {
			if s.buf[s.offset] == '+' || s.buf[s.offset] == '-' {
				s.offset++
				if s.offset > len(s.buf) {
					return io.ErrShortBuffer
				}
			}
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	s.endValueOffset = s.offset
	return nil
}

func (s *yamlParser) scanLine() (*Line, error) {
	err := s.skipEmptyLines()
	if err != nil {
		return nil, err
	}

	if s.offset >= len(s.buf) {
		return nil, io.ErrShortBuffer
	}

	var indent []byte
	i := 0

	buf := s.buf[s.offset:]
	isArray := false

	// scan indent
	for ; i < len(buf); i++ {
		c := buf[i]
		if c == '\n' {
			s.offset += i
			return &Line{isEmpty: true}, nil
		}
		if !isSpace(c) {
			indent = buf[:i]
			if c == '-' {
				if i == len(buf)-1 || isSpace(buf[i+1]) {
					isArray = true
				}
			}
			break
		}
	}

	if isArray {
		// skip dash
		i++
		// skip spaces
		for ; i < len(buf); i++ {
			if buf[i] == '\n' {
				s.offset += i
				return &Line{
						isArray: true, indent: indent,
					},
					nil
			}
			if !isSpace(buf[i]) {
				break
			}
		}
	}

	var name []byte

	// scan name
	nameStart := i + s.offset
	j := i
	for ; j < len(buf); j++ {
		if buf[j] == '\n' {
			name = buf[i:j] // and remember to trim
			s.offset += j
			return &Line{
				name:      name,
				isArray:   isArray,
				indent:    indent,
				hasValue:  false,
				nameStart: nameStart,
			}, nil
		}
		if buf[j] == ':' {
			name = buf[i:j] // and remember to trim
			break
		}
		if j == len(buf)-1 {
			j++
			name = buf[i:j] // and remember to trim
			s.offset += j
			return &Line{
				name:      name,
				isArray:   isArray,
				indent:    indent,
				hasValue:  false,
				nameStart: nameStart,
			}, nil
		}
	}

	// scan value
	// skip colon
	j++
	valueOffset := s.offset + j
	k := j
	for ; k < len(buf); k++ {
		if buf[k] == '\n' {
			s.offset += k
			return &Line{
					name:        name,
					isArray:     isArray,
					value:       buf[j:k], // trim this
					indent:      indent,
					hasValue:    true,
					valueOffset: valueOffset,
					nameStart:   nameStart,
				},
				nil
		}
	}

	s.offset += k
	return &Line{
			name:        name,
			isArray:     isArray,
			value:       buf[j:k], // trim this
			indent:      indent,
			hasValue:    true,
			valueOffset: valueOffset,
			nameStart:   nameStart,
		},
		nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isDigit19(c byte) bool {
	return c >= '1' && c <= '9'
}

func (s *yamlParser) scanArray() error {
	s.startValueOffset = s.offset

	line, err := s.scanLine()
	if err != nil {
		return err
	}

	indent := line.indent

	for {
		// increment?
		s.endValueOffset = s.offset

		nextLine, err := s.peekNextLine()
		if err != nil {
			if err == io.ErrShortBuffer {
				s.isValueArray = true
				return nil
			}
			return err
		}
		if len(nextLine.indent) < len(indent) {
			s.isValueArray = true
			return nil
		}
		if len(nextLine.indent) == len(indent) && !nextLine.isArray {
			return errors.New("Cannot have non array object") // this is also an error
		}
		_, err = s.scanLine()
		if err != nil {
			return errors.New("Server error")
		}
	}
}

func (s *yamlParser) scanObjectValue() error {
	s.startValueOffset = s.offset

	line, err := s.scanLine()
	if err != nil {
		return err
	}

	indent := line.indent

	for {
		// increment?
		s.endValueOffset = s.offset

		nextLine, err := s.peekNextLine()
		if err != nil {
			if err == io.ErrShortBuffer {
				s.isValueObject = true
				return nil
			}
			return err
		}
		// next line is not child or sibling
		if len(nextLine.indent) < len(indent) {
			s.isValueObject = true
			return nil
		}
		if len(nextLine.indent) == len(indent) && nextLine.isArray {
			return errors.New("Not implemented same line array children") // this is also an error
		}
		_, err = s.scanLine()
		if err != nil {
			return errors.New("Server error")
		}
	}
}

//  Copyright 2015 Walter Schulze
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

//Package xml contains a parser for XML.
package xml

import (
	"bytes"
	"fmt"
	"github.com/katydid/katydid/parser"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type xmlParser struct {
	buf        []byte
	dec        *Decoder
	tok        Token
	attrs      []Attr
	attrIndex  int
	attrValue  bool
	attrFirst  bool
	attrPrefix string
	elemPrefix string
}

//XMLParser is an xml parser.
type XMLParser interface {
	parser.Interface
	//Init intialises the parser with a byte buffer containing xml.
	Init([]byte) error
}

//NewXMLParser returns a new xml parser.
func NewXMLParser(options ...Option) XMLParser {
	x := &xmlParser{}
	for _, option := range options {
		option(x)
	}
	return x
}

//Option is used set options when creating a new XMLParser
type Option func(x *xmlParser)

//WithAttrPrefix specifies the prefix which will be added to attributes returned by the parser.
func WithAttrPrefix(a string) func(x *xmlParser) {
	return func(x *xmlParser) {
		x.attrPrefix = a
	}
}

//WithElemPrefix specifies the prefix which will be added to elements returned by the parser.
func WithElemPrefix(e string) func(x *xmlParser) {
	return func(x *xmlParser) {
		x.elemPrefix = e
	}
}

var procInstPattern = regexp.MustCompile(`<\?.*\?>`)

func (p *xmlParser) Init(buf []byte) error {
	buf = procInstPattern.ReplaceAll(buf, []byte{})
	buf = bytes.TrimSpace(buf)
	p.buf = buf
	p.dec = NewDecoder(NewBuffer(buf))
	p.dec.Strict = false
	return nil
}

func hasContent(c CharData) bool {
	return len(string(c)) > 0
}

func (p *xmlParser) Next() (err error) {
	if p.attrValue {
		if p.attrFirst {
			p.attrFirst = false
			return nil
		} else {
			return io.EOF
		}
	}
	if p.tok == nil {
		if p.attrs != nil {
			p.attrIndex++
			if p.attrIndex < len(p.attrs) {
				return nil
			}
		}
	}
	if p.tok != nil {
		for {
			if _, ok := p.tok.(StartElement); ok {
				//fmt.Printf("Skipping %s\n", s.Name)
				if err := p.dec.Skip(); err != nil {
					//fmt.Printf("Skip err = %v\n", err)
					return err
				}
				break
			} else if _, ok := p.tok.(EndElement); ok {
				return io.EOF
			} else if c, ok := p.tok.(CharData); ok {
				if hasContent(c) {
					break
				}
			} else if _, ok := p.tok.(Comment); ok {
				p.tok, err = p.dec.Token()
				//fmt.Printf("Comment Next Token %#v, err = %v\n", p.tok, err)
				if err != nil {
					return err
				}
			} else {
				panic(fmt.Sprintf("unknown token %T", p.tok))
			}
		}
	}
	p.tok, err = p.dec.Token()
	//fmt.Printf("Next Token %#v, err %v\n", p.tok, err)
	for err == nil {
		if _, ok := p.tok.(StartElement); ok {
			break
		} else if c, ok := p.tok.(CharData); ok {
			if hasContent(c) {
				break
			}
		} else if _, ok := p.tok.(EndElement); ok {
			return io.EOF
		}
		p.tok, err = p.dec.Token()
		//fmt.Printf("Next Next Token %#v, err = %v\n", p.tok, err)
	}
	return err
}

func (p *xmlParser) IsLeaf() bool {
	if p.tok == nil {
		if p.attrValue {
			return true
		}
		return false
	}
	_, ok := p.tok.(CharData)
	//fmt.Printf("IsLeaf %#v\n", p.tok)
	return ok
}

func (p *xmlParser) getValue() string {
	if p.tok == nil && p.attrValue {
		return p.attrs[p.attrIndex].Value
	}
	if c, ok := p.tok.(CharData); ok {
		return string(c)
	}
	return ""
}

func (p *xmlParser) Double() (float64, error) {
	return strconv.ParseFloat(p.getValue(), 64)
}

func (p *xmlParser) Int() (int64, error) {
	i, err := strconv.ParseInt(p.getValue(), 10, 64)
	return int64(i), err
}

func (p *xmlParser) Uint() (uint64, error) {
	i, err := strconv.ParseUint(p.getValue(), 10, 64)
	return uint64(i), err
}

func (p *xmlParser) Bool() (bool, error) {
	return strconv.ParseBool(strings.TrimSpace(p.getValue()))
}

func (p *xmlParser) String() (string, error) {
	if p.tok == nil && p.attrIndex < len(p.attrs) {
		if p.attrValue {
			return p.attrs[p.attrIndex].Value, nil
		} else {
			return p.attrPrefix + p.attrs[p.attrIndex].Name.Local, nil
		}
	}
	if s, ok := p.tok.(StartElement); ok {
		return p.elemPrefix + s.Name.Local, nil
	}
	if c, ok := p.tok.(CharData); ok {
		return string(c), nil
	}
	return "", parser.ErrNotString
}

func (p *xmlParser) Bytes() ([]byte, error) {
	if c, ok := p.tok.(CharData); ok {
		return []byte(c), nil
	}
	return nil, parser.ErrNotBytes
}

func (p *xmlParser) Up() {
	if p.tok == nil {
		if p.attrValue {
			p.attrValue = false
			p.attrFirst = false
			return
		}
	}
	if _, ok := p.tok.(EndElement); ok {
		p.tok = nil
		p.attrs = nil
		p.attrIndex = 0
		return
	}
	if err := p.dec.Skip(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
}

func (p *xmlParser) Down() {
	if p.tok == nil {
		if p.attrIndex < len(p.attrs) {
			p.attrValue = true
			p.attrFirst = true
			return
		}
	}
	if s, ok := p.tok.(StartElement); ok {
		p.tok = nil
		p.attrs = s.Attr
		p.attrIndex = -1
		return
	}
	panic(fmt.Sprintf("not a start element %T", p.tok))
}

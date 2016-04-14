//  Copyright 2016 Walter Schulze
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

//Package xml encodes a parser.Interface into xml.
//This can be used for transcoding.
//
//TODO: currently only handles very naive cases, more tests would help.
package xml

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"github.com/katydid/katydid/parser"
	"io"
	"strconv"
)

//Encode encodes a parser.Interface into a byte slice containing valid xml.
func Encode(p parser.Interface) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	e := xml.NewEncoder(buf)
	if err := encode(e, p); err != nil {
		return nil, err
	}
	if err := e.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(e *xml.Encoder, p parser.Interface) error {
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if p.IsLeaf() {
			if err := e.EncodeToken(getCharData(p)); err != nil {
				return err
			}
		} else {
			name := getName(p)
			if err := e.EncodeToken(xml.StartElement{Name: name}); err != nil {
				return err
			}
			p.Down()
			if err := encode(e, p); err != nil {
				return err
			}
			p.Up()
			if err := e.EncodeToken(xml.EndElement{Name: name}); err != nil {
				return err
			}
		}
	}
	return nil
}

func getName(p parser.Interface) xml.Name {
	i, err := p.Int()
	if err == nil {
		return xml.Name{Local: strconv.FormatInt(i, 10)}
	}
	u, err := p.Uint()
	if err == nil {
		return xml.Name{Local: strconv.FormatUint(u, 10)}
	}
	d, err := p.Double()
	if err == nil {
		return xml.Name{Local: strconv.FormatFloat(d, 'e', -1, 64)}
	}
	b, err := p.Bool()
	if err == nil {
		return xml.Name{Local: strconv.FormatBool(b)}
	}
	s, err := p.String()
	if err == nil {
		return xml.Name{Local: s}
	}
	v, err := p.Bytes()
	if err == nil {
		return xml.Name{Local: base64.StdEncoding.EncodeToString(v)}
	}
	return xml.Name{}
}

func getCharData(p parser.Interface) xml.CharData {
	i, err := p.Int()
	if err == nil {
		return xml.CharData([]byte(strconv.FormatInt(i, 10)))
	}
	u, err := p.Uint()
	if err == nil {
		return xml.CharData([]byte(strconv.FormatUint(u, 10)))
	}
	d, err := p.Double()
	if err == nil {
		return xml.CharData([]byte(strconv.FormatFloat(d, 'e', -1, 64)))
	}
	b, err := p.Bool()
	if err == nil {
		return xml.CharData([]byte(strconv.FormatBool(b)))
	}
	s, err := p.String()
	if err == nil {
		return xml.CharData([]byte(s))
	}
	v, err := p.Bytes()
	if err == nil {
		return xml.CharData([]byte(base64.StdEncoding.EncodeToString(v)))
	}
	return nil
}

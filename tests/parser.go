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

package tests

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/serialize"
	jparser "github.com/katydid/katydid/serialize/json"
	pparser "github.com/katydid/katydid/serialize/proto"
	rparser "github.com/katydid/katydid/serialize/reflect"
	xparser "github.com/katydid/katydid/serialize/xml"
	"reflect"
)

type Codecs struct {
	Description string
	Parsers     map[string]NewParser
}

func getDesc(m interface{}) string {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func AllCodecs(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"reflect": Reflect(m).Parsers["reflect"],
			"json":    Json(m).Parsers["json"],
			"proto":   Proto(m).Parsers["proto"],
		},
	}
}

func Reflect(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"reflect": func() serialize.Parser {
				return NewReflectParser(m)
			},
		},
	}
}

func Json(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"json": func() serialize.Parser {
				return NewJsonParser(m)
			},
		},
	}
}

func Proto(m interface{}) Codecs {
	messageName := reflect.TypeOf(m).Elem().Name()
	packageName := "tests"
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"proto": func() serialize.Parser {
				return NewProtoParser(packageName, messageName, m.(ProtoMessage))
			},
		},
	}
}

func XML(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"xml": func() serialize.Parser {
				return NewXMLParser(m)
			},
		},
	}
}

type NewParser func() serialize.Parser

func NewReflectParser(m interface{}) serialize.Parser {
	s := rparser.NewReflectParser()
	s.Init(reflect.ValueOf(m))
	return s
}

func NewJsonParser(m interface{}) serialize.Parser {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := jparser.NewJsonParser()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	return s
}

type ProtoMessage interface {
	Description() (desc *descriptor.FileDescriptorSet)
	proto.Message
}

func NewProtoParser(pkg, msg string, m ProtoMessage) serialize.Parser {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := pparser.NewProtoParser(pkg, msg, m.Description())
	if err := s.Init(data); err != nil {
		panic(err)
	}
	return s
}

func NewXMLParser(m interface{}) serialize.Parser {
	data, err := xml.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := xparser.NewXMLParser()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	return s
}

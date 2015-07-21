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
	"fmt"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/serialize"
	jscanner "github.com/katydid/katydid/serialize/json"
	pscanner "github.com/katydid/katydid/serialize/proto/scanner"
	rscanner "github.com/katydid/katydid/serialize/reflect"
	"reflect"
)

type Codecs struct {
	Description string
	Scanners    map[string]NewScanner
}

func AllCodecs(m interface{}) Codecs {
	return Codecs{
		Description: fmt.Sprintf("%#v", m),
		Scanners: map[string]NewScanner{
			"reflect": Reflect(m).Scanners["reflect"],
			"json":    Json(m).Scanners["json"],
			"proto":   Proto(m).Scanners["proto"],
		},
	}
}

func Reflect(m interface{}) Codecs {
	return Codecs{
		Description: fmt.Sprintf("%#v", m),
		Scanners: map[string]NewScanner{
			"reflect": func() serialize.Scanner {
				return NewReflectScanner(m)
			},
		},
	}
}

func Json(m interface{}) Codecs {
	return Codecs{
		Description: fmt.Sprintf("%#v", m),
		Scanners: map[string]NewScanner{
			"json": func() serialize.Scanner {
				return NewJsonScanner(m)
			},
		},
	}
}

func Proto(m interface{}) Codecs {
	messageName := reflect.TypeOf(m).Elem().Name()
	packageName := "tests"
	return Codecs{
		Description: fmt.Sprintf("%#v", m),
		Scanners: map[string]NewScanner{
			"proto": func() serialize.Scanner {
				return NewProtoScanner(packageName, messageName, m.(ProtoMessage))
			},
		},
	}
}

type NewScanner func() serialize.Scanner

func NewReflectScanner(m interface{}) serialize.Scanner {
	s := rscanner.NewReflectScanner()
	s.Init(reflect.ValueOf(m))
	return s
}

func NewJsonScanner(m interface{}) serialize.Scanner {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := jscanner.NewJsonScanner()
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

func NewProtoScanner(pkg, msg string, m ProtoMessage) serialize.Scanner {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := pscanner.NewProtoScanner(pkg, msg, m.Description())
	if err := s.Init(data); err != nil {
		panic(err)
	}
	return s
}

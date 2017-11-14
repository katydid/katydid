//  Copyright 2017 Walter Schulze
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

package testsuite

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/gen"
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/parser/json"
	protoparser "github.com/katydid/katydid/parser/proto"
	"github.com/katydid/katydid/parser/xml"
	"github.com/katydid/katydid/relapse"
	"github.com/katydid/katydid/relapse/ast"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

var testpath string
var benchpath string

func init() {
	gopath := os.Getenv("GOPATH")
	testpath = filepath.Join(gopath, "src/github.com/katydid/testsuite/relapse/tests")
	benchpath = filepath.Join(gopath, "src/github.com/katydid/testsuite/relapse/benches")
}

func TestSuiteExists() bool {
	return exists(testpath)
}

func BenchSuiteExists() bool {
	return exists(benchpath)
}

func getFolders(path string) (map[string][]string, error) {
	folders := make(map[string][]string)
	codecFileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, codecFileInfo := range codecFileInfos {
		if !codecFileInfo.IsDir() {
			continue
		}
		codecFolderName := codecFileInfo.Name()
		codecPath := filepath.Join(path, codecFolderName)
		caseDirInfos, err := ioutil.ReadDir(codecPath)
		if err != nil {
			return nil, err
		}
		for _, caseDirInfo := range caseDirInfos {
			if !caseDirInfo.IsDir() {
				continue
			}
			casePath := filepath.Join(codecPath, caseDirInfo.Name())
			folders[codecFolderName] = append(folders[codecFolderName], casePath)
		}
	}
	return folders, nil
}

func ReadTestSuite() ([]Test, error) {
	tests := []Test{}
	codecs, err := getFolders(testpath)
	if err != nil {
		return nil, err
	}
	for codec, folders := range codecs {
		switch codec {
		case "pbname", "pbnum", "json", "xml":
		default:
			// codec not supported
			continue
		}
		for _, folder := range folders {
			test, err := readTestFolder(folder)
			if err != nil {
				return nil, err
			}
			tests = append(tests, *test)
		}
	}
	return tests, nil
}

func ReadBenchmarkSuite() ([]Bench, error) {
	benches := []Bench{}
	codecs, err := getFolders(benchpath)
	if err != nil {
		return nil, err
	}
	for codec, folders := range codecs {
		switch codec {
		case "pbname", "pbnum", "json", "xml":
		default:
			// codec not supported
			continue
		}
		for _, folder := range folders {
			bench, err := readBenchFolder(folder)
			if err != nil {
				return nil, err
			}
			benches = append(benches, *bench)
		}
	}
	return benches, nil
}

type Test struct {
	Name     string
	Grammar  *ast.Grammar
	Parser   parser.Interface
	Expected bool
	Record   bool
}

func readTestFolder(path string) (*Test, error) {
	name := filepath.Base(path)
	g, err := readGrammar(path)
	if err != nil {
		return nil, err
	}
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading folder <%s>", err, path)
	}
	var p parser.Interface
	var expected bool
	var codecName string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		filebase := fileInfo.Name()
		filename := filepath.Join(path, filebase)
		names := strings.Split(filebase, ".")
		valid := names[0] == "valid"
		invalid := names[0] == "invalid"
		if !valid && !invalid {
			continue
		}
		expected = valid
		codecName = names[len(names)-1]
		switch codecName {
		case "pbname":
			pkgName, msgName, desc, err := getProtoDesc(filename)
			if err != nil {
				return nil, err
			}
			p, err = newProtoNameParser(pkgName, msgName, desc, filename)
			if err != nil {
				return nil, err
			}
		case "pbnum":
			pkgName, msgName, desc, err := getProtoDesc(filename)
			if err != nil {
				return nil, err
			}
			p, err = newProtoNumParser(pkgName, msgName, desc, filename)
			if err != nil {
				return nil, err
			}
		case "json":
			p, err = newJsonParser(filename)
			if err != nil {
				return nil, err
			}
		case "xml":
			p, err = newXMLParser(filename)
			if err != nil {
				return nil, err
			}
		default:
			// unsupported codec
			continue
		}
	}
	if p == nil {
		return nil, fmt.Errorf("couldn't find valid.* or invalid.* filename inside <%s>", path)
	}
	name = name + gen.CapFirst(codecName)
	return &Test{
		Name:     name,
		Grammar:  g,
		Parser:   p,
		Expected: expected,
		Record:   codecName != "xml",
	}, nil
}

type Bench struct {
	Name    string
	Grammar *ast.Grammar
	Parsers []ResetParser
	Record  bool
}

type ResetParser interface {
	parser.Interface
	Reset() error
}

func readBenchFolder(path string) (*Bench, error) {
	name := filepath.Base(path)
	g, err := readGrammar(path)
	if err != nil {
		return nil, err
	}
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading folder <%s>", err, path)
	}
	var parsers []ResetParser
	var pkgName, msgName string
	var codecName string
	var desc *descriptor.FileDescriptorSet
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		filebase := fileInfo.Name()
		if _, err := strconv.Atoi(strings.Split(filebase, ".")[0]); err != nil {
			continue
		}
		filename := filepath.Join(path, filebase)
		codecName = filepath.Ext(filename)[1:]
		switch codecName {
		case "pbname":
			if desc == nil {
				pkgName, msgName, desc, err = getProtoDesc(filename)
				if err != nil {
					return nil, err
				}
			}
			p, err := newProtoNameParser(pkgName, msgName, desc, filename)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, p)
		case "pbnum":
			if desc == nil {
				pkgName, msgName, desc, err = getProtoDesc(filename)
				if err != nil {
					return nil, err
				}
			}
			p, err := newProtoNumParser(pkgName, msgName, desc, filename)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, p)
		case "json":
			p, err := newJsonParser(filename)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, p)
		case "xml":
			p, err := newXMLParser(filename)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, p)
		default:
			// unsupported codec
			continue
		}
	}
	return &Bench{
		Name:    name + capFirst(codecName),
		Grammar: g,
		Parsers: parsers,
		Record:  true,
	}, nil
}

func capFirst(s string) string {
	b := []byte(s)
	b[0] ^= ' '
	return string(b)
}

func readGrammar(path string) (*ast.Grammar, error) {
	relapseTxt := filepath.Join(path, "relapse.txt")
	relapseBytes, err := ioutil.ReadFile(relapseTxt)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading file <%s>", err, relapseTxt)
	}
	g, err := relapse.Parse(string(relapseBytes))
	if err != nil {
		return nil, fmt.Errorf("err <%v> parsing grammar from file <%s>", err, relapseTxt)
	}
	return g, nil
}

func newXMLParser(filename string) (ResetParser, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading file <%s>", err, filename)
	}
	x := xml.NewXMLParser()
	if err := x.Init(bytes); err != nil {
		return nil, fmt.Errorf("err <%v> parser.Init with bytes from filename <%s>", err, filename)
	}
	return x, nil
}

func newJsonParser(filename string) (ResetParser, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading file <%s>", err, filename)
	}
	j := json.NewJsonParser()
	if err := j.Init(bytes); err != nil {
		return nil, fmt.Errorf("err <%v> parser.Init with bytes from filename <%s>", err, filename)
	}
	return j, nil
}

func getProtoDesc(filename string) (pkgName, msgName string, desc *descriptor.FileDescriptorSet, err error) {
	names := strings.Split(filepath.Base(filename), ".")
	schemaName := strings.Join(names[1:len(names)-1], ".")
	schemaFilename := filepath.Clean(filepath.Join(filepath.Join(filepath.Dir(filename), ".."), schemaName))
	descData, err := ioutil.ReadFile(schemaFilename)
	if err != nil {
		return "", "", nil, fmt.Errorf("err <%v> reading filename <%s>", err, schemaFilename)
	}
	desc = &descriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(descData, desc); err != nil {
		return "", "", nil, fmt.Errorf("err <%v> unmarshaling descriptor from filename <%s>", err, schemaFilename)
	}
	pkgName = names[1]
	msgName = names[2]
	return pkgName, msgName, desc, nil
}

func newProtoNumParser(pkgName, msgName string, desc *descriptor.FileDescriptorSet, filename string) (ResetParser, error) {
	pp, err := protoparser.NewProtoNumParser(pkgName, msgName, desc)
	if err != nil {
		return nil, fmt.Errorf("err <%v> createing proto parser", err)
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading file <%s>", err, filename)
	}
	if err := pp.Init(bytes); err != nil {
		return nil, fmt.Errorf("err <%v> parser.Init with bytes from filename <%s>", err, filename)
	}
	return pp, nil
}

func newProtoNameParser(pkgName, msgName string, desc *descriptor.FileDescriptorSet, filename string) (ResetParser, error) {
	pp, err := protoparser.NewProtoNameParser(pkgName, msgName, desc)
	if err != nil {
		return nil, fmt.Errorf("err <%v> createing proto parser", err)
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("err <%v> reading file <%s>", err, filename)
	}
	if err := pp.Init(bytes); err != nil {
		return nil, fmt.Errorf("err <%v> parser.Init with bytes from filename <%s>", err, filename)
	}
	return pp, nil
}

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

package main_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"github.com/awalterschulze/katydid/exp/asm/lexer"
	"github.com/awalterschulze/katydid/exp/asm/parser"
	main "github.com/awalterschulze/katydid/exp/asm/test"

	"fmt"
	"path"
	"path/filepath"
	goruntime "runtime"
	"strings"

	"io/ioutil"
)

func newJsonFile(name string, m proto.Message) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	os.Remove(name + ".json")
	jsonFile, err := os.Create(name + ".json")
	if err != nil {
		panic(err)
	}
	if _, err := jsonFile.Write(data); err != nil {
		panic(err)
	}
	if err := jsonFile.Close(); err != nil {
		panic(err)
	}
}

func exampleName() string {
	pc, _, _, _ := goruntime.Caller(2)
	name := goruntime.FuncForPC(pc).Name()
	names := strings.Split(name, ".")
	name = names[len(names)-1]
	return strings.Replace(name, "Test", "", 1)
}

func exampleDir(name string) string {
	var dir string
	if path.IsAbs(os.Args[0]) {
		dir, _ = path.Split(path.Clean(os.Args[0]))
	} else {
		wd, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("Getwd failed: %s", err))
		}
		dir, _ = path.Split(path.Join(wd, os.Args[0]))
	}
	fullpath := path.Join(dir, "example", name)
	err := os.MkdirAll(fullpath, 0777)
	if err != nil {
		panic(err)
	}
	return fullpath
}

func example(t *testing.T, protoFilename string, m proto.Message, katydidStr string, positive bool) {
	name := exampleName()

	test(t, protoFilename, m, katydidStr, positive)

	dir := exampleDir(name)

	dotFilename := filepath.Join(dir, name+".dot")
	pdfFilename := filepath.Join(dir, name+".pdf")
	txtFilename := filepath.Join(dir, name+".txt")
	oneFilename := filepath.Join(dir, "one.box")
	twoFilename := filepath.Join(dir, "two.box")
	threeFilename := filepath.Join(dir, "three.box")

	if err := ioutil.WriteFile(txtFilename, []byte(katydidStr), 0666); err != nil {
		panic(err)
	}

	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer([]byte(katydidStr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	rules := r.(*ast.Rules)
	dotString := rules.Dot()
	if err := ioutil.WriteFile(dotFilename, []byte(dotString), 0666); err != nil {
		panic(err)
	}
	cmd := exec.Command("dot", dotFilename, "-Tpdf", "-o", pdfFilename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("dot error: " + err.Error() + ":" + string(output))
	}

	protoBytes, err := ioutil.ReadFile(protoFilename)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(oneFilename, protoBytes, 0666); err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err)
	}
	fulltyp := fmt.Sprintf("%T", m)
	typ := strings.Replace(fulltyp, "main.", "", -1)
	newtyp := strings.Replace(typ, "*", "&", -1)
	popStr := `package main

import (
	"encoding/json"
)

func Populate() (` + typ + `, error) {
	m := ` + newtyp + `{}

	jsonValue := ` + "`" + string(jsonBytes) + "`" + `

	if err := json.Unmarshal([]byte(jsonValue), m); err != nil {
		panic(err)
	}
	return m, nil
}

`
	if err := ioutil.WriteFile(twoFilename, []byte(popStr), 0666); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(threeFilename, []byte(katydidStr), 0666); err != nil {
		panic(err)
	}

}

var david = &main.Person{
	Name: proto.String("David"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var robert = &main.Person{
	Name: proto.String("Robert"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(456),
			Street: proto.String("TheStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var contextPerson = `//Has this person ever lived at 456 TheStreet

root = main.Person
main.Person = start
start numberAndStreet = accept
start _ = start

main.Address = address
address number = number
address street = street
address _ = address
number street = numberAndStreet
number _ = number
street number = numberAndStreet
street _ = street

if (decInt64(main.Address.Number) == int64(456)) then number else noNumber

if contains(nfkc(decString(main.Address.Street)), nfkc("TheStreet")) then street else noStreet
`

func TestContextDavid(t *testing.T) {
	example(t, "person.proto", david, contextPerson, false)
}

func TestContextRobert(t *testing.T) {
	example(t, "person.proto", robert, contextPerson, true)
}

var ioUtil = &main.SrcTree{
	PackageName: proto.String("io/ioutil"),
	Imports: []*main.SrcTree{
		{
			PackageName: proto.String("io"),
			Imports: []*main.SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("sync"),
				},
			},
		},
		{
			PackageName: proto.String("os"),
			Imports: []*main.SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("runtime"),
				},
			},
		},
	},
}

var pathSrcTree = &main.SrcTree{
	PackageName: proto.String("path"),
	Imports: []*main.SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("strings"),
			Imports: []*main.SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("uncode"),
				},
				{
					PackageName: proto.String("uncode/utf8"),
				},
			},
		},
		{
			PackageName: proto.String("unicode/utf8"),
		},
	},
}

var runtime = &main.SrcTree{
	PackageName: proto.String("runtime"),
	Imports: []*main.SrcTree{
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

var syscall = &main.SrcTree{
	PackageName: proto.String("syscall"),
	Imports: []*main.SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("runtime"),
		},
		{
			PackageName: proto.String("sync"),
			Imports: []*main.SrcTree{
				{
					PackageName: proto.String("sync/atomic"),
				},
				{
					PackageName: proto.String("unsafe"),
				},
			},
		},
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

var recursiveSrcTree = `//Does this SrcTree depend on io or is its package name io

root = main.SrcTree
main.SrcTree = start
start accept = accept
start _ = start
accept _ = accept

if (decString(main.SrcTree.PackageName) == "io") 
  then accept 
  else packageName

`

func TestRecursiveSrcTreeIoUtil(t *testing.T) {
	example(t, "srctree.proto", ioUtil, recursiveSrcTree, true)
}

func TestRecursiveSrcTreePath(t *testing.T) {
	example(t, "srctree.proto", pathSrcTree, recursiveSrcTree, true)
}

func TestRecursiveSrcTreeRuntime(t *testing.T) {
	example(t, "srctree.proto", runtime, recursiveSrcTree, false)
}

func TestRecursiveSrcTreeSyscall(t *testing.T) {
	example(t, "srctree.proto", syscall, recursiveSrcTree, false)
}

var mover = &main.Person{
	Name: proto.String("Mover"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var shaker = &main.Person{
	Name: proto.String("Shaker"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(55),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var routine = &main.Person{
	Name: proto.String("Routine"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(3),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0124444444"),
}

var listIndexAddress = `//Is this Person's newest street number 1 and second newest street number 2.
//Assume that addresses are appended to the list, so the last address is the newest address.
// find main.Person where { main.Person { Addresses[-2].Number == 2 && Addresses[-1].Number == 1 } }

root = main.Person
main.Person = start
start numberTwo = topNumberTwo
start _ = start
topNumberTwo numberOne = accept
topNumberTwo numberTwo = topNumberTwo
topNumberTwo _ = start
accept numberTwo = topNumberTwo
accept _ = start

main.Address = address
address numberTwo = numberTwo
address numberOne = numberOne
address _ = address
numberTwo numberTwo = numberTwo
numberTwo numberOne = numberOne
numberOne numberTwo = numberTwo
numberOne numberOne = numberOne

if (decInt64(main.Address.Number) == int64(1))
  then numberOne
  else {
    if (decInt64(main.Address.Number) == int64(2))
    then numberTwo
    else noNumber
  }
`

func TestListIndexAddressMover(t *testing.T) {
	example(t, "person.proto", mover, listIndexAddress, false)
}

func TestListIndexAddressShaker(t *testing.T) {
	example(t, "person.proto", shaker, listIndexAddress, true)
}

func TestListIndexAddressRoutine(t *testing.T) {
	example(t, "person.proto", routine, listIndexAddress, false)
}

var noname = &main.Person{
	Addresses: []*main.Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var john = &main.Person{
	Name: proto.String("John"),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var smith = &main.Person{
	Name: proto.String(""),
	Addresses: []*main.Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var nilName = `//Is this person's name missing

root = main.Person
main.Person = accept
accept name = reject
accept _ = accept

if exists(main.Person.Name)
  then name
  else noname
`

func TestNilNameNoName(t *testing.T) {
	example(t, "person.proto", noname, nilName, true)
}

func TestNilNameJohn(t *testing.T) {
	example(t, "person.proto", john, nilName, false)
}

func TestNilNameSmith(t *testing.T) {
	example(t, "person.proto", smith, nilName, false)
}

var lenName = `//Is this person's name an empty string

root = main.Person
main.Person = start
start name = reject
start noname = accept
start _ = start

if (length(decString(main.Person.Name)) == int64(0))
  then noname
  else name
`

func TestLenNameNoName(t *testing.T) {
	example(t, "person.proto", noname, lenName, false)
}

func TestLenNameJohn(t *testing.T) {
	example(t, "person.proto", john, lenName, false)
}

func TestLenNameSmith(t *testing.T) {
	example(t, "person.proto", smith, lenName, true)
}

var emptyOrNil = `//Is this person's name empty or an empty string

root = main.Person
main.Person = accept
accept name = reject
accept _ = accept

if (length(decString(main.Person.Name)) == int64(0))
  then noname
  else name
`

func TestEmptyOrNilNoName(t *testing.T) {
	example(t, "person.proto", noname, emptyOrNil, true)
}

func TestEmptyOrNilJohn(t *testing.T) {
	example(t, "person.proto", john, emptyOrNil, false)
}

func TestEmptyOrNilSmith(t *testing.T) {
	example(t, "person.proto", smith, emptyOrNil, true)
}

var incorrentNotName = `root = main.Person
main.Person = start
start notname = accept
start _ = start

if not((decString(main.Person.Name) == "David")) 
  then notname 
  else name
`

func TestIncorrectNotNameNoName(t *testing.T) {
	example(t, "person.proto", noname, incorrentNotName, false)
}

func TestIncorrectNotNameRobert(t *testing.T) {
	example(t, "person.proto", robert, incorrentNotName, true)
}

func TestIncorrectNotNameSmith(t *testing.T) {
	example(t, "person.proto", smith, incorrentNotName, true)
}

func TestIncorrectNotNameDavid(t *testing.T) {
	example(t, "person.proto", david, incorrentNotName, false)
}

var correctNotName = `root = main.Person
main.Person = accept
accept name = reject
reject _ = reject
accept _ = accept

if (decString(main.Person.Name) == "David") 
  then name 
  else noname
`

func TestCorrectNotNameNoName(t *testing.T) {
	example(t, "person.proto", noname, correctNotName, true)
}

func TestCorrectNotNameRobert(t *testing.T) {
	example(t, "person.proto", robert, correctNotName, true)
}

func TestCorrectNotNameSmith(t *testing.T) {
	example(t, "person.proto", smith, correctNotName, true)
}

func TestCorrectNotNameDavid(t *testing.T) {
	example(t, "person.proto", david, correctNotName, false)
}

var andNameTelephone = `//Is this person's name David and telephone number 0123456789

root = main.Person
main.Person = start
start name = name
start tel = tel
start _ = start
name tel = accept
name _ = name
tel name = accept
tel _ = tel

if (decString(main.Person.Name) == "David") 
  then name 
  else noname

if (decString(main.Person.Telephone) == "0123456789") 
  then tel 
  else notel
`

func TestAndNameTelephoneDavid(t *testing.T) {
	example(t, "person.proto", david, andNameTelephone, true)
}

func TestAndNameTelephoneJohn(t *testing.T) {
	example(t, "person.proto", john, andNameTelephone, false)
}

func TestAndNameTelephoneMover(t *testing.T) {
	example(t, "person.proto", mover, andNameTelephone, false)
}

func TestAndNameTelephoneSmith(t *testing.T) {
	example(t, "person.proto", smith, andNameTelephone, false)
}

var orNameTelephone = `//Is this person's name David or telephone number 0123456789

root = main.Person
main.Person = start
start name = accept
start tel = accept
start _ = start
accept _ = accept

if (decString(main.Person.Name) == "David") 
  then name 
  else noname

if (decString(main.Person.Telephone) == "0123456789") 
  then tel 
  else notel
`

func TestOrNameTelephoneDavid(t *testing.T) {
	example(t, "person.proto", david, orNameTelephone, true)
}

func TestOrNameTelephoneJohn(t *testing.T) {
	example(t, "person.proto", john, orNameTelephone, true)
}

func TestOrNameTelephoneMover(t *testing.T) {
	example(t, "person.proto", mover, orNameTelephone, true)
}

func TestOrNameTelephoneSmith(t *testing.T) {
	example(t, "person.proto", smith, orNameTelephone, false)
}

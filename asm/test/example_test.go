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

	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	main "github.com/katydid/katydid/asm/test"

	"fmt"
	"path"
	"path/filepath"
	goruntime "runtime"
	"strings"

	"io/ioutil"
)

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
	jsonFilename := filepath.Join(dir, name+".json")

	if err := ioutil.WriteFile(txtFilename, []byte(katydidStr), 0666); err != nil {
		panic(err)
	}

	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(katydidStr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
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

	jsonBytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(jsonFilename, jsonBytes, 0666); err != nil {
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
init = start
final = accept

accept _ = (accept, accept, accept)
reject _ = (reject, reject, reject)
ignore _ = (ignore, ignore, ignore)

start Addresses = (address, accept, start)
start _ = (ignore, start, start)
address Number = (number, hasNumber, address)
address Street = (street, hasStreet, address)
address _ = (ignore, address, address)
hasNumber Street = (street, accept, hasNumber)
hasNumber _ = (ignore, hasNumber, hasNumber)
hasStreet Number = (number, accept, hasStreet)
hasStreet _ = (ignore, hasStreet, hasStreet)

func number = eq($int64, int64(456))
func street = contains(toLower($string), toLower("TheStreet"))

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
init = start
final = accept

accept _ = (accept, accept, accept)

start Imports = (start, accept, start)
start PackageName = (packageName, accept, start)

func packageName = eq($string, "io")

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
init = start
final = accept

ignore _ = (ignore, ignore, ignore)

func numberOne = eq($int64, int64(1))
func numberTwo = eq($int64, int64(2))

start Addresses = (lookingForTwo, hasTwo, start)
start _ = (ignore, start, start)

lookingForTwo Number = (numberTwo, accept, ignore)
lookingForTwo _ = (ignore, lookingForTwo, lookingForTwo)

hasTwo Addresses = (lookingForOne, accept, start)
hasTwo _ = (ignore, hasTwo, hasTwo)

lookingForOne Number = (numberOne, accept, ignore)
lookingForOne _ = (ignore, lookingForOne, lookingForOne)

// main.Person = start
// start numberTwo = topNumberTwo
// start _ = start
// topNumberTwo numberOne = accept
// topNumberTwo numberTwo = topNumberTwo
// topNumberTwo _ = start
// accept numberTwo = topNumberTwo
// accept _ = start

// main.Address = address
// address numberTwo = numberTwo
// address numberOne = numberOne
// address _ = address
// numberTwo numberTwo = numberTwo
// numberTwo numberOne = numberOne
// numberOne numberTwo = numberTwo
// numberOne numberOne = numberOne

// if eq($int64(main.Address.Number), int64(1))
//   then numberOne
//   else {
//     if eq($int64(main.Address.Number), int64(2))
//     then numberTwo
//     else noNumber
//   }
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
init = accept
final = accept

accept Name = (accept, reject, reject)
reject _ = (reject, reject, reject)
accept _ = (accept, accept, accept)

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
init = start
final = accept
start Name = (zeroLength, accept, start)
start _ = (accept, start, start)
accept _ = (accept, accept, accept)

func zeroLength = eq(length($string), int64(0))
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
init = accept
final = accept

accept Name = (zeroLength, accept, reject)
accept _ = (accept, accept, accept)
reject _ = (reject, reject, reject)

func zeroLength = eq(length($string), int64(0))

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

var incorrentNotName = `

root = main.Person
init = start
final = accept

start Name = (notDavid, accept, start)
start _ = (accept, start, start)
accept _ = (accept, accept, accept)

func notDavid = not(eq($string, "David"))

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

var correctNotName = `

root = main.Person
init = accept
final = accept

accept Name = (david, reject, accept)
accept _ = (accept, accept, accept)
reject _ = (reject, reject, reject)

func david = eq($string, "David") 

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
init = start
final = accept

start Name = (name, hasName, start)
start Telephone = (tel, hasTel, start)
hasName Telephone = (tel, accept, hasName)
hasTel Name = (name, accept, hasTel)

start _ = (accept, start, start)
hasName _ = (accept, hasName, hasName)
hasTel _ = (accept, hasTel, hasTel)
accept _ = (accept, accept, accept)

func name = eq($string, "David") 

func tel = eq($string, "0123456789") 

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
init = start
final = accept

start Name = (number, accept, start)
start Telephone = (tel, accept, start)
start _ = (accept, start, start)
accept _ = (accept, accept, accept)

func name = eq($string, "David") 
func tel = eq($string, "0123456789") 
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

var listOfTelephones = `//Is this person's telephone number 0123456789 or 0127897897

root = main.Person
init = start
final = accept

start Telephone = (tel, accept, start)
start _ = (accept, start, start)
accept _ = (accept, accept, accept)

func tel = or(eq($string, "0123456789"), eq($string, "0127897897"))

`

func TestListOfTelephonesDavid(t *testing.T) {
	example(t, "person.proto", david, listOfTelephones, true)
}

func TestListOfTelephonesShaker(t *testing.T) {
	example(t, "person.proto", shaker, listOfTelephones, true)
}

func TestListOfTelephonesRoutine(t *testing.T) {
	example(t, "person.proto", routine, listOfTelephones, false)
}

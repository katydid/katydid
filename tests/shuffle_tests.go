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
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
)

//Foundations of XML Processing: The Tree Automata Approach - Chapter 15.2 Page 194
//interleave((a[], b[]); (c[], d[]))
var Page194 = G{"main": InAnyOrder(
	InOrder(
		In("A", Any()),
		In("B", Any()),
	),
	InOrder(
		In("C", Any()),
		In("D", Any()),
	),
)}

//a[], b[], c[], d[]
var Page194abcd = `{
	"A": "#",
	"B": "#",
	"C": "#",
	"D": "#",
}`

var Page194bacd = `{
	"B": "#",
	"A": "#",
	"C": "#",
	"D": "#",
}`

var Page194abdc = `{
	"A": "#",
	"B": "#",
	"D": "#",
	"C": "#",
}`

var Page194badc = `{
	"B": "#",
	"A": "#",
	"D": "#",
	"C": "#",
}`

func init() {
	Validate("Page194abcd", Page194, JsonString(Page194abcd), true)
	Validate("Page194bacd", Page194, JsonString(Page194bacd), false)
	Validate("Page194abdc", Page194, JsonString(Page194abdc), false)
	Validate("Page194badc", Page194, JsonString(Page194badc), false)
}

//a[], c[], b[], d[]
var Page194acbd = `{
	"A": "#",
	"C": "#",
	"B": "#",
	"D": "#",
}`

var Page194bcad = `{
	"B": "#",
	"C": "#",
	"A": "#",
	"D": "#",
}`

var Page194adbc = `{
	"A": "#",
	"D": "#",
	"B": "#",
	"C": "#",
}`

var Page194bdac = `{
	"B": "#",
	"D": "#",
	"A": "#",
	"C": "#",
}`

func init() {
	Validate("Page194acbd", Page194, JsonString(Page194acbd), true)
	Validate("Page194bcad", Page194, JsonString(Page194bcad), false)
	Validate("Page194adbc", Page194, JsonString(Page194adbc), false)
	Validate("Page194bdac", Page194, JsonString(Page194bdac), false)
}

//a[], c[], d[], b[]
var Page194acdb = `{
	"A": "#",
	"C": "#",
	"D": "#",
	"B": "#",
}`

var Page194bcda = `{
	"B": "#",
	"C": "#",
	"D": "#",
	"A": "#",
}`

var Page194adcb = `{
	"A": "#",
	"D": "#",
	"C": "#",
	"B": "#",
}`

var Page194bdca = `{
	"D": "#",
	"C": "#",
	"B": "#",
	"A": "#",
}`

func init() {
	Validate("Page194acdb", Page194, JsonString(Page194acdb), true)
	Validate("Page194bcda", Page194, JsonString(Page194bcda), false)
	Validate("Page194adcb", Page194, JsonString(Page194adcb), false)
	Validate("Page194bdca", Page194, JsonString(Page194bdca), false)
}

//c[], a[], b[], d[]
var Page194cabd = `{
	"C": "#",
	"A": "#",
	"B": "#",
	"D": "#",
}`

var Page194cbad = `{
	"C": "#",
	"B": "#",
	"A": "#",
	"D": "#",
}`

var Page194dabc = `{
	"D": "#",
	"A": "#",
	"B": "#",
	"C": "#",
}`

var Page194dbac = `{
	"D": "#",
	"B": "#",
	"A": "#",
	"C": "#",
}`

func init() {
	Validate("Page194cabd", Page194, JsonString(Page194cabd), true)
	Validate("Page194cbad", Page194, JsonString(Page194cbad), false)
	Validate("Page194dabc", Page194, JsonString(Page194dabc), false)
	Validate("Page194dbac", Page194, JsonString(Page194dbac), false)
}

//c[], a[], d[], b[]
var Page194cadb = `{
	"C": "#",
	"A": "#",
	"D": "#",
	"B": "#",
}`

var Page194cbda = `{
	"C": "#",
	"B": "#",
	"D": "#",
	"A": "#",
}`

var Page194dacb = `{
	"D": "#",
	"A": "#",
	"C": "#",
	"B": "#",
}`

var Page194dbca = `{
	"D": "#",
	"B": "#",
	"C": "#",
	"A": "#",
}`

func init() {
	Validate("Page194cadb", Page194, JsonString(Page194cadb), true)
	Validate("Page194cbda", Page194, JsonString(Page194cbda), false)
	Validate("Page194dacb", Page194, JsonString(Page194dacb), false)
	Validate("Page194dbca", Page194, JsonString(Page194dbca), false)
}

//c[], d[], a[], b[]
var Page194cdab = `{
	"C": "#",
	"D": "#",
	"A": "#",
	"B": "#",
}`

var Page194cdba = `{
	"C": "#",
	"D": "#",
	"B": "#",
	"A": "#",
}`

var Page194dcab = `{
	"D": "#",
	"C": "#",
	"A": "#",
	"B": "#",
}`

var Page194dcba = `{
	"D": "#",
	"C": "#",
	"B": "#",
	"A": "#",
}`

func init() {
	Validate("Page194cdab", Page194, JsonString(Page194cdab), true)
	Validate("Page194cdba", Page194, JsonString(Page194cdba), false)
	Validate("Page194dcab", Page194, JsonString(Page194dcab), false)
	Validate("Page194dcba", Page194, JsonString(Page194dcba), false)
}

var whitespace = Many(Value(Regex(StringConst("^(\\s)+$"), StringVar())))

//Name && Address? && Email*
var Page195NameAddrEmail = G{"main": In("Person", InAnyOrder(
	In("Name", Any()),
	Maybe(In("Addr", Any())),
	Many(In("Email", Any())),
	whitespace,
))}

//Email*, Name, Email*, Address?, Email*
var Page195E2NameE2AddrE2 = `<Person>
	<Email>an@e.mail</Email>
	<Email>an@e.mail2</Email>
	<Name>Shannon</Name>
	<Email>other@e.mail</Email>
	<Email>other@e.mail2</Email>
	<Addr>123 Some Street</Addr>
	<Email>another@e.mail</Email>
	<Email>another@e.mail2</Email>
</Person>`

var Page195E1NameE1AddrE1 = `<Person>
	<Email>an@e.mail</Email>
	<Name>Shannon</Name>
	<Email>other@e.mail</Email>
	<Addr>123 Some Street</Addr>
	<Email>another@e.mail</Email>
</Person>`

var Page195E0NameE0AddrE0 = `<Person>
	<Name>Shannon</Name>
	<Addr>123 Some Street</Addr>
</Person>`

func init() {
	Validate("Page195E2NameE2AddrE2", Page195NameAddrEmail, XMLString(Page195E2NameE2AddrE2), true)
	Validate("Page195E1NameE1AddrE1", Page195NameAddrEmail, XMLString(Page195E1NameE1AddrE1), true)
	Validate("Page195E0NameE0AddrE0", Page195NameAddrEmail, XMLString(Page195E0NameE0AddrE0), true)
}

//Email*, Address?, Email*, Name, Email*
var Page195E2AddrE2NameE2 = `<Person>
	<Email>an@e.mail</Email>
	<Email>an@e.mail2</Email>
	<Addr>123 Some Street</Addr>
	<Email>other@e.mail</Email>
	<Email>other@e.mail2</Email>
	<Name>Shannon</Name>
	<Email>another@e.mail</Email>
	<Email>another@e.mail2</Email>
</Person>`

var Page195E1AddrE1NameE1 = `<Person>
	<Email>an@e.mail</Email>
	<Addr>123 Some Street</Addr>
	<Email>other@e.mail</Email>
	<Name>Shannon</Name>
	<Email>another@e.mail</Email>
</Person>`

var Page195E0AddrE0NameE0 = `<Person>
	<Addr>123 Some Street</Addr>
	<Name>Shannon</Name>
</Person>`

func init() {
	Validate("Page195E2AddrE2NameE2", Page195NameAddrEmail, XMLString(Page195E2AddrE2NameE2), true)
	Validate("Page195E1AddrE1NameE1", Page195NameAddrEmail, XMLString(Page195E1AddrE1NameE1), true)
	Validate("Page195E0AddrE0NameE0", Page195NameAddrEmail, XMLString(Page195E0AddrE0NameE0), true)
}

//Email*, Name, Email*
var Page195E2NameE2 = `<Person>
	<Email>an@e.mail</Email>
	<Email>an@e.mail2</Email>
	<Name>Shannon</Name>
	<Email>other@e.mail</Email>
	<Email>other@e.mail2</Email>
</Person>`

var Page195E1NameE1 = `<Person>
	<Email>an@e.mail</Email>
	<Name>Shannon</Name>
	<Email>other@e.mail</Email>
</Person>`

var Page195E0NameE0 = `<Person>
	<Name>Shannon</Name>
</Person>`

func init() {
	Validate("Page195E2NameE2", Page195NameAddrEmail, XMLString(Page195E2NameE2), true)
	Validate("Page195E1NameE1", Page195NameAddrEmail, XMLString(Page195E1NameE1), true)
	Validate("Page195E0NameE0", Page195NameAddrEmail, XMLString(Page195E0NameE0), true)
}

//Email*, Address?, Email*
var Page195E2AddrE2 = `<Person>
	<Email>an@e.mail</Email>
	<Email>an@e.mail2</Email>
	<Addr>123 Some Street</Addr>
	<Email>another@e.mail</Email>
	<Email>another@e.mail2</Email>
</Person>`

var Page195E1AddrE1 = `<Person>
	<Email>an@e.mail</Email>
	<Addr>123 Some Street</Addr>
	<Email>another@e.mail</Email>
</Person>`

var Page195E0AddrE0 = `<Person>
	<Addr>123 Some Street</Addr>
</Person>`

func init() {
	Validate("Page195E2AddrE2", Page195NameAddrEmail, XMLString(Page195E2AddrE2), false)
	Validate("Page195E1AddrE1", Page195NameAddrEmail, XMLString(Page195E1AddrE1), false)
	Validate("Page195E0AddrE0", Page195NameAddrEmail, XMLString(Page195E0AddrE0), false)
}

//Email*
var Page195E2 = `<Person>
	<Email>an@e.mail</Email>
	<Email>an@e.mail2</Email>
</Person>`

var Page195E1 = `<Person>
	<Email>an@e.mail</Email>
</Person>`

var Page195E0 = `<Person>
</Person>`

func init() {
	Validate("Page195E2", Page195NameAddrEmail, XMLString(Page195E2), false)
	Validate("Page195E1", Page195NameAddrEmail, XMLString(Page195E1), false)
	Validate("Page195E0", Page195NameAddrEmail, XMLString(Page195E0), false)
}

//(Name,Tel?) && (Name,Email*)
var Page195NameTelNameEmail = G{"main": In("Person", InAnyOrder(
	InOrder(
		In("Name", Any()),
		Maybe(In("Tel", Any())),
	),
	InOrder(
		In("Name", Any()),
		Many(In("Email", Any())),
	),
	whitespace,
))}

var Page195E2NameE0TelE0NameE2 = `
<Person>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Shannon</Name>
	<Tel>123</Tel>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE2TelE0NameE2 = `
<Person>
	<Name>Shannon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Tel>123</Tel>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE0TelE2NameE2 = `
<Person>
	<Name>Shannon</Name>
	<Tel>123</Tel>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE2TelE2NameE2 = `
<Person>
	<Name>Shannon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Tel>123</Tel>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE0TelE0NameE2 = `
<Person>
	<Name>Shannon</Name>
	<Tel>123</Tel>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE0TelE2NameE0 = `
<Person>
	<Name>Shannon</Name>
	<Tel>123</Tel>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Sheldon</Name>
</Person>
`

var Page195E0NameE2TelE0NameE0 = `
<Person>
	<Name>Shannon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Tel>123</Tel>
	<Name>Sheldon</Name>
</Person>
`

var Page195E0NameE2NameE0 = `
<Person>
	<Name>Shannon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Sheldon</Name>
</Person>
`

var Page195E0NameE0NameE2 = `
<Person>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

var Page195E0NameE0NameE0 = `
<Person>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
</Person>
`

var Page195E0NameE2NameE0TelE0 = `
<Person>
	<Name>Shannon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Name>Sheldon</Name>
	<Tel>123</Tel>
</Person>
`

var Page195E0NameE0NameE2TelE0 = `
<Person>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
	<Tel>123</Tel>
</Person>
`

var Page195E0NameE0NameE0TelE0 = `
<Person>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
	<Tel>123</Tel>
</Person>
`

var Page195E0TelE0NameE0NameE0 = `
<Person>
	<Tel>123</Tel>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
</Person>
`

var Page195E0TelE0NameE0NameE2 = `
<Person>
	<Tel>123</Tel>
	<Name>Shannon</Name>
	<Name>Sheldon</Name>
	<Email>an@e.mail</Email>
	<Email>another@e.mail</Email>
</Person>
`

func init() {
	Validate("Page195E2NameE0TelE0NameE2", Page195NameTelNameEmail, XMLString(Page195E2NameE0TelE0NameE2), false)
	Validate("Page195E0NameE2TelE0NameE2", Page195NameTelNameEmail, XMLString(Page195E0NameE2TelE0NameE2), false)
	Validate("Page195E0NameE0TelE2NameE2", Page195NameTelNameEmail, XMLString(Page195E0NameE0TelE2NameE2), false)
	Validate("Page195E0NameE2TelE2NameE2", Page195NameTelNameEmail, XMLString(Page195E0NameE2TelE2NameE2), false)
	Validate("Page195E0NameE0TelE0NameE2", Page195NameTelNameEmail, XMLString(Page195E0NameE0TelE0NameE2), true)
	Validate("Page195E0NameE0TelE2NameE0", Page195NameTelNameEmail, XMLString(Page195E0NameE0TelE2NameE0), false)
	Validate("Page195E0NameE2TelE0NameE0", Page195NameTelNameEmail, XMLString(Page195E0NameE2TelE0NameE0), false)
	Validate("Page195E0NameE0NameE2", Page195NameTelNameEmail, XMLString(Page195E0NameE0NameE2), true)
	Validate("Page195E0NameE2NameE0", Page195NameTelNameEmail, XMLString(Page195E0NameE2NameE0), true)
	Validate("Page195E0NameE0NameE0", Page195NameTelNameEmail, XMLString(Page195E0NameE0NameE0), true)
	Validate("Page195E0NameE2NameE0TelE0", Page195NameTelNameEmail, XMLString(Page195E0NameE2NameE0TelE0), true)
	Validate("Page195E0NameE0NameE2TelE0", Page195NameTelNameEmail, XMLString(Page195E0NameE0NameE2TelE0), true)
	Validate("Page195E0NameE0NameE0TelE0", Page195NameTelNameEmail, XMLString(Page195E0NameE0NameE0TelE0), true)
	Validate("Page195E0TelE0NameE0NameE0", Page195NameTelNameEmail, XMLString(Page195E0TelE0NameE0NameE0), false)
	Validate("Page195E0TelE0NameE0NameE2", Page195NameTelNameEmail, XMLString(Page195E0TelE0NameE0NameE2), false)
}

//ABStar{A:a & B:b & *}
var ABStar = G{"main": In("ABStar",
	InAnyOrder(
		In("A", Value(StringEq(StringVar(), StringConst("a")))),
		In("B", Value(StringEq(StringVar(), StringConst("b")))),
		Any(),
	),
)}

var ABStarABC = `<ABStar>
	<A>a</A>
	<B>b</B>
	<C>c</C>
</ABStar>`

var ABStarCACBC = `<ABStar>
	<C>c</C>
	<A>a</A>
	<C>c</C>
	<B>b</B>
	<C>c</C>
</ABStar>`

var ABStarCBCAC = `<ABStar>
	<C>c</C>
	<B>b</B>
	<C>c</C>
	<A>a</A>
	<C>c</C>
</ABStar>`

var ABStarBBCAC = `<ABStar>
	<B>b</B>
	<B>b</B>
	<C>c</C>
	<A>a</A>
	<C>c</C>
</ABStar>`

var ABStarBBAAA = `<ABStar>
	<B>b</B>
	<B>b</B>
	<A>c</A>
	<A>a</A>
	<A>b</A>
</ABStar>`

var ABStarAAA = `<ABStar>
	<A>c</A>
	<A>a</A>
	<A>b</A>
</ABStar>`

var ABStarBB = `<ABStar>
	<B>b</B>
	<B>b</B>
</ABStar>`

var ABStarBBC = `<ABStar>
	<B>b</B>
	<B>b</B>
	<C>c</C>
</ABStar>`

func init() {
	Validate("ABStarABC", ABStar, XMLString(ABStarABC), true)
	Validate("ABStarCACBC", ABStar, XMLString(ABStarCACBC), true)
	Validate("ABStarCBCAC", ABStar, XMLString(ABStarCBCAC), true)
	Validate("ABStarBBCAC", ABStar, XMLString(ABStarBBCAC), true)
	Validate("ABStarBBAAA", ABStar, XMLString(ABStarBBAAA), true)
	Validate("ABStarAAA", ABStar, XMLString(ABStarAAA), false)
	Validate("ABStarBB", ABStar, XMLString(ABStarBB), false)
	Validate("ABStarBBC", ABStar, XMLString(ABStarBBC), false)
}

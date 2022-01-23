## Katydid

[Katydid](http://katydid.github.io) is a toolkit for trees.

[![GoDoc](https://godoc.org/github.com/katydid/katydid?status.svg)](https://godoc.org/github.com/katydid/katydid) [![Build Status](https://travis-ci.org/katydid/katydid.svg?branch=master)](https://travis-ci.org/katydid/katydid)

![Katydid Logo](https://cdn.rawgit.com/katydid/katydid.github.io/master/logo.png)

Currently there are three tools in the katydid toolkit:

  * Relapse: a regular expression type language for trees that matches up to 1000000s of records per second, 
  * A collection of parsers (protobuf, json, xml, reflected go structures, yaml*) which are easily extendable, and
  * A collection of encoders (protobuf, json, xml, reflected go structures) which are useful for dynamic transcoding.

[more...](http://katydid.github.io)

*The YAML parser implementation is in progress and has the following knwown limitations
- Array elements with more than 1 key are parsed incorrectly
- Comments are not supported
- The file delimeter "---" is not supported
- JSON syntax is not yet supported
- Quotes are currently included as part of strings
- Mixing tabs and spaces in indendation is not supported
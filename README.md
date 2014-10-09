katydid
=======

Katydid is a language based on Tree Grammars.
Currently it only has an assembler for a validator which uses bottom up hedge (unranked tree) automata and functions on the leaves.
In future a more intuitive language will be included that translates to this assembler.

Katydid supports multiple types of trees:
 - [Protocol Buffers](http://code.google.com/p/protobuf/)
 - [Json](http://json.org/)
 - [Reflected structures in Go](http://golang.org/pkg/reflect)

Currently they all need to have a protocol buffer specification.

Tree Automata have been found to be very applicable to XML processing.
Katydid tries to apply Tree Automata to different types of trees.

[![Build Status](https://drone.io/github.com/awalterschulze/katydid/status.png)](https://drone.io/github.com/awalterschulze/katydid/latest)

Katydid is ready for use, but backwards compatibility is not yet guaranteed.

Documentation is a work in progress http://awalterschulze.github.io/katydid

Ideals
------

 - Fits into the theoretical model of Tree Grammars
 - Solves practical use cases
 - Fast
 - Decidable
 - Expressive

Examples
--------

http://awalterschulze.github.io/katydid/examples/


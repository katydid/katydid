katydid
=======

Katydid is a language based on Tree Grammars.
Currently it only has an assembler for a recognizer which uses bottom up hedge (unranked tree) automata and functions on the leaves.
In future a more logical and intuitive language will be included that translates to this assembler.

Currently Katydid only supports a [protocol buffers](http://code.google.com/p/protobuf/) serialization format.
Protocol buffers encode data structures.
The encoded protocol buffers have a semi-unordered unranked tree structure.

Tree Automata have been found to be very applicable to XML processing.
Katydid tries to do the same by applying Tree Automata to Protocol Buffers.

[![Build Status](https://drone.io/github.com/awalterschulze/katydid/status.png)](https://drone.io/github.com/awalterschulze/katydid/latest)

Katydid is still in an experimental phase.

Documentation is a work in progress http://awalterschulze.github.io/katydid

Ideals
------

 - Fits into a theoretical model, probably tree grammars
 - Solves practical use cases
 - Fast
 - Decidable
 - Expressive

Examples
--------

http://awalterschulze.github.io/katydid/examples/


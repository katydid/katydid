katydid
=======

[protocol buffers](http://code.google.com/p/protobuf/) + tree grammars = katydid

[![Build Status](https://drone.io/github.com/awalterschulze/katydid/status.png)](https://drone.io/github.com/awalterschulze/katydid/latest)

Katydid is still in an experimental phase.

Name
----

https://www.youtube.com/watch?v=SvjSP2xYZm8

Ideals
------

 - Fits into a theoretical model, probably tree grammars
 - Solves practical use cases
 - Fast
 - Decidable
 - Expressive 

Tree Grammars
-------------

Protocol buffers encode data structures.
The encoded protocol buffers have a semi-unordered unranked tree structure.
Tree Automata have been found to be very applicable to XML processing.
Katydid tries to do the same by applying Tree Automata to Protocol Buffers.

First Experiment
----------------

Using a bottom up hedge (unranked tree) automata we can create pretty expressive queries.

http://arborist.katydid.ws

Language
--------

The current language reflects the proposed inner workings (assembler) of the final language.

https://github.com/awalterschulze/katydid/blob/master/exp/asm/asm.bnf

This will be useful for debugging in future.

Current Functions
-----------------

The idea is to make katydid easily extendable with other functions.
I am still proving the concept and I only needed these functions for my examples.

 - decString([]byte) string
 - decInt64([]byte) string
 - contains(string, string) bool
 - eq(int64, int64) bool
 - eq(string, string) bool
 - ge(int64, int64) bool
 - nfkc(string) string
 - not(bool) bool
 - length(string) int
 - exists([]byte) bool

Next steps
----------

 - Compilation and efficient execution
 - Create a usable language that translates to the current debugging/middle language
 - Boolean Logic Operators (Possibly using alternating tree automata)
 - Optimize for specific usecases like negative indices, without breaking the ideals
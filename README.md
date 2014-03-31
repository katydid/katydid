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

Examples
--------

Using a bottom up hedge (unranked tree) automata we can create pretty expressive queries.

http://arborist.katydid.ws

http://arborist.katydid.ws/example/ContextRobert

http://arborist.katydid.ws/example/ContextDavid

http://arborist.katydid.ws/example/AndNameTelephoneDavid

http://arborist.katydid.ws/example/OrNameTelephoneDavid

http://arborist.katydid.ws/example/EmptyOrNilNoName

http://arborist.katydid.ws/example/ListIndexAddressMover

http://arborist.katydid.ws/example/ListIndexAddressShaker

http://arborist.katydid.ws/example/RecursiveSrcTreePath

http://arborist.katydid.ws/example/RecursiveSrcTreeSyscall

Language
--------

The current language reflects the proposed inner workings (assembler) of the final language.

https://github.com/awalterschulze/katydid/blob/master/asm/asm.bnf

This will be useful for debugging in future.

Current Functions
-----------------

Please see the [list of functions](https://github.com/awalterschulze/katydid/blob/master/functions.txt).

This list can easily be extended with your own functions using the Register pattern.  See the funcs package for examples.

Streaming
---------

Katydid does some streaming processing which is incompatible with the merging feature of protocol buffers.
Merging allows a marshaled protocol buffer to contain more than one instance of the same optional field.
This assumes that the last instance will be unmarshaled as the actual value.
Katydid always assumes the first instance is the actual value.
This means Katydid will not always process protocol buffers, that have been marshaled with the merging feature, correctly.

For more information please see:

https://developers.google.com/protocol-buffers/docs/encoding#optional

NoMerge is a function in gogoprotobuf that checks whether a marshaled protocol buffer has used this feature.

http://godoc.org/code.google.com/p/gogoprotobuf/fieldpath#NoMerge


Next steps
----------

 - Create a usable language that translates to the current debugging/middle language
 - Boolean Logic Operators (Possibly using alternating tree automata)


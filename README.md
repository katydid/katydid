#katydid

Katydid is a tree validation language.

###What is a validation language?
Regular expressions are used as validators for strings.
RelaxNG, DTD and XSchema are validation languages for XML.
JsonSchema is a validation language for JSON.

###What are trees?
A tree is a structure, record, class that does not contain any loops.

Katydid supports multiple types of trees:
 - [Protocol Buffers](https://developers.google.com/protocol-buffers/)
 - [Json](http://json.org/)
 - [Reflected structures in Go](http://golang.org/pkg/reflect)
 - and even XML for the giant lizards
 - Easy to add your own. Simply implement a parser ...

##Status
Katydid is ready for use, but backwards compatibility is not yet guaranteed.
[![Build Status](https://drone.io/github.com/katydid/katydid/status.png)](https://drone.io/github.com/katydid/katydid/latest)

Documentation is a work in progress http://katydid.github.io

##Ideals

 - Fits into a theoretical model
 - Solves practical use cases
 - Fast
 - Decidable (excluding User Defined Functions)
 - Expressive

##Examples

http://katydid.github.io/examples/

##Theory

Tree Automata have been found to be very applicable to XML processing.
Katydid tries to apply Tree Automata to different types of trees.

##Tools

 - [Translate RelaxNG](https://github.com/katydid/relanxg) to Katydid
 - [Translate JsonSchema](https://github.com/katydid/jsonschema) to Katydid

##Possible Futures

 - Capturing (like Regular Expression Capturing, but for trees)
 - Dynamic Transcoding between different Tree Types
 - Optionally enable Javascript/Lua user defined functions.
 - More serialization formats: CapnProto?, Bson, Yaml, Toml ...
 - Remove all dependencies from core Katydid repo and move all other extra functionality into richer feature set repos that extend the core of Katydid.



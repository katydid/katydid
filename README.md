## Katydid

Katydid is a toolkit for trees.
Currently there are three tools:

  * A collection of parsers,
  * Relapse: a tree validation language, and
  * A collection of encoders.

### What is a validation language?
Regular expressions are used as validators for strings.
RelaxNG, DTD and XSchema are validation languages for XML.
JsonSchema is a validation language for JSON.

### What are trees?
A tree is a structure, record, class that does not contain any loops.

Katydid has implemented parsers for multiple types of trees:

  * [Protocol Buffers](https://developers.google.com/protocol-buffers/)
  * [Json](http://json.org/)
  * [Reflected structures in Go](http://golang.org/pkg/reflect)
  * and even XML for the dinosaurs and dragons.
  * Easy to add your own. Simply implement a [parser](http://katydid.github.io/dev/parsers)

Relapse can validate these trees, since they have implemented Parsers.

Encoders can take parsers and encode (or transcode) them into other types of trees.

### What is a Katydid?

A Katydid is the common name for a [leaf insect](https://avatars1.githubusercontent.com/u/9207606?v=3&s=200).
You might enjoy this [video](https://www.youtube.com/watch?v=SvjSP2xYZm8).

## Status
Katydid version 0.2 is in alpha.
[![Build Status](https://drone.io/github.com/katydid/katydid/status.png)](https://drone.io/github.com/katydid/katydid/latest)

## Ideals

  * Fits into a theoretical model
  * Solves practical use cases
  * Fast
  * Decidable (excluding User Defined Functions)
  * Expressive

## A Tour of Relapse

[Relapse Tour](http://katydid.github.io/tour)

## A Playground for Relapse

[Relapse Playground](http://katydid.github.io/play)

## GoDoc

[![GoDoc](https://godoc.org/github.com/katydid/katydid?status.svg)](https://godoc.org/github.com/katydid/katydid)

## Theory

Tree Automata and Visual Pushdown Automata have been found to be very applicable to XML processing.
Relapse tries to apply these Automata to different types of trees.

## External Tools

  * [Translate RelaxNG](https://github.com/katydid/relaxng) to Relapse
  * [Translate JsonSchema](https://github.com/katydid/jsonschema) to Relapse

## Possible Futures

  * Language Agnostic Test set
  * Capturing (like Regular Expression Capturing, but for trees)
  * Dynamic Transcoding between different Tree Types
  * Optionally enable Javascript/Lua user defined functions.
  * More serialization formats: CapnProto?, Bson, Yaml, Toml ...
  * Remove all dependencies from core Katydid repository and move all other extra functionality into richer feature set repositories that extend the core of Katydid.

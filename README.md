## Katydid

Katydid is a toolkit for trees.
Currently there are two tools:

  * A collection of parsers, and
  * Relapse: a tree validation language.

### What is a validation language
Regular expressions are used as validators for strings.
RelaxNG, DTD and XSchema are validation languages for XML.
JsonSchema is a validation language for JSON.

### What are trees
A tree is a structure, record, class that does not contain any loops.

Katydid has implemented parsers for multiple types of trees:

  * [Protocol Buffers](https://developers.google.com/protocol-buffers/)
  * [Json](http://json.org/)
  * [Reflected structures in Go](http://golang.org/pkg/reflect)
  * and even XML for the dinosaurs and dragons.
  * Easy to add your own. Simply implement a [parser](http://katydid.github.io/dev/parsers)

Relapse can validate these trees, since they have implemented Parsers.

## Status
Katydid version 0.2 is in beta.
[![Build Status](https://drone.io/github.com/katydid/katydid/status.png)](https://drone.io/github.com/katydid/katydid/latest)

## Ideals

  * Fits into a theoretical model
  * Solves practical use cases
  * Fast
  * Decidable (excluding User Defined Functions)
  * Expressive

## Playground

[Relapse Playgroun](http://katydid.github.io/play)

## Examples

[Relapse Examples](http://katydid.github.io/doc/examples)

## Theory

Tree Automata and Visual Pushdown Automata have been found to be very applicable to XML processing.
Relapse tries to apply these Automata to different types of trees.

## Tools

  * [Translate RelaxNG](https://github.com/katydid/relaxng) to Relapse
  * [Translate JsonSchema](https://github.com/katydid/jsonschema) to Relapse

## Possible Futures

  * Language Agnostic Testset
  * Capturing (like Regular Expression Capturing, but for trees)
  * Dynamic Transcoding between different Tree Types
  * Optionally enable Javascript/Lua user defined functions.
  * More serialization formats: CapnProto?, Bson, Yaml, Toml ...
  * Remove all dependencies from core Katydid repo and move all other extra functionality into richer feature set repos that extend the core of Katydid.


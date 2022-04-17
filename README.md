# reser
Polymorphic serialization util for golang.

## Features
* Abstracting serialization using interfaces: `Encoder, PolyEncoder, PolyMarshaler` and others
* Polymorphic serialization

It's very simplistic and does not implement any serialization format on it's own. Instead it builds tools on top of existing serialization formats like JSON and their existing implementations, like one in golang stdlib.

## Limitations
It does not support embedded polymorphic serialization.
You can serialize/deserialize types, which implement some interface, but you can't have these embedded in your structure, unless you manually handle 
polymorphic marshaling for them.

It also simplifies doing lots of amazing things like integrating ZSTD dictionary compression directly into serializer/deserializer function.

## Examples 
See docs and examples files (WIP)
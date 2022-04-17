# reser
Serialization util for golang.

It allows for things like:
* Abstracting serialization using interfaces
* Polymorphic serialization with tag
* Polymorphic serialization without tag

It's very simplistic and does not implement any serialization format on it's own. Instead it builds tools on top of existing serialization formats like JSON and their existing implementations, like one in golang stdlib.

It also simplifies doing lots of amazing things like integrating ZSTD dictionary compression directly into serializer/deserializer function.

## Examples 
See docs and examples files (WIP)
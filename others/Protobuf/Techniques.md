# [Techniques](https://developers.google.com/protocol-buffers/docs/techniques#self-description)

- [Streaming Multiple Messages](#)
- [Large Data Sets](#)
- [Self-describing Messages](#)

This page describes some commonly-used design patterns for dealing with Protocol Buffers. You can also send design and usage questions to the [Protocol Buffers discussion group](#).

## Streaming Multiple Messages

If you want to write multiple messages to a single file or stream, it is up to you to keep track of where one message ends and the next begins. The Protocol Buffer wire format is not self-delimiting, so protocol buffer parsers cannot determine where a message ends on their own. The easiest way to solve this problem is to write the size of each message before you write the message itself. When you read the messages back in, you read the size, then read the bytes into a separate buffer, then parse from that buffer. (If you want to avoid copying bytes to a separate buffer, check out the `CodedInputStream` class (in both C++ and Java) which can be told to limit reads to a certain number of bytes.)

## Large Data Sets

Protocol Buffers are not designed to handle large messages. As a general rule of thumb, if you are dealing in messages larger than a megabyte each, it may be time to consider an alternate strategy.

That said, Protocol Buffers are great for handling individual messages *within* a large data set. Usually, large data sets are really just a collection of small pieces, where each small piece may be a structured piece of data. Even though Protocol Buffers cannot handle the entire set at once, using Protocol Buffers to encode each piece greatly simplifies your problem: now all you need is to handle a set of byte strings rather than a set of structures.

Protocol Buffers do not include any built-in support for large data sets because different situations call for different solutions. Sometimes a simple list of records will do while other times you may want something more like a database. Each solution should be developed as a separate library, so that only those who need it need to pay the costs.

## Self-describing Messages

Protocol Buffers do not contain descriptions of their own types. Thus, given only a raw message without the corresponding `.proto` file defining its type, it is difficult to extract any useful data.

However, note that the contents of a .proto file can itself be represented using protocol buffers. The file `src/google/protobuf/descriptor.proto` in the source code package defines the message types involved. `protoc` can output a `FileDescriptorSet` – which represents a set of .proto files – using the `--descriptor_set_out` option. With this, you could define a self-describing protocol message like so:

```protobuf

syntax = "proto3";

import "google/protobuf/any.proto";
import "google/protobuf/descriptor.proto";

message SelfDescribingMessage {
  // Set of FileDescriptorProtos which describe the type and its dependencies.
  google.protobuf.FileDescriptorSet descriptor_set = 1;

  // The message and its type, encoded as an Any message.
  google.protobuf.Any message = 2;
}

```

By using classes like `DynamicMessage` (available in C++ and Java), you can then write tools which can manipulate `SelfDescribingMessage`s.

All that said, the reason that this functionality is not included in the Protocol Buffer library is because we have never had a use for it inside Google.

This technique requires support for dynamic messages using descriptors. Please check that your platforms support this feature before using self-describing messages.

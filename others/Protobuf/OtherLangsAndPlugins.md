# [Other Languages](https://developers.google.com/protocol-buffers/docs/reference/other)

While the current release just includes compilers and APIs for C++, Java, and Python, the compiler code is designed so that it's easy to add support for other languages. There are several ongoing projects to add new language implementations to Protocol Buffers, including C, C#, Haskell, Perl, Ruby, and more.

For a list of links to projects we know about, see the [third-party add-ons wiki page](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md).

## Compiler Plugins

As of version 2.3.0 (January 2010), protoc, the Protocol Buffers Compiler, can be extended to support new languages via plugins. A plugin is just a program which reads a `CodeGeneratorRequest` protocol buffer from standard input and then writes a `CodeGeneratorResponse` protocol buffer to standard output. These message types are defined in [`plugin.proto`](https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.compiler.plugin.pb). We recommend that all third-party code generators be written as plugins, as this allows all generators to provide a consistent interface and share a single parser implementation.

Additionally, plugins are able to insert code into the files generated by other code generators. See the comments about "insertion points" in [`plugin.proto`](https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.compiler.plugin.pb) for more on this. This could be used, for example, to write a plugin which generates RPC service code that is tailored for a particular RPC system. See the documentation for the generated code in each language to find out what insertion points they provide.
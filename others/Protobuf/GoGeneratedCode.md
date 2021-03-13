# [Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

This page describes exactly what Go code the protocol buffer compiler generates for any given protocol definition. Any differences between proto2 and proto3 generated code are highlighted - note that these differences are in the generated code as described in this document, not the base API, which are the same in both versions. You should read the [proto2 language guide](#) and/or the [proto3 language guide](#) before reading this document.

## Compiler Invocation

The protocol buffer compiler requires a plugin to generate Go code. Install it with:

```sh

go install google.golang.org/protobuf/cmd/protoc-gen-go

```

This will install a `protoc-gen-go` binary in `$GOBIN`. Set the `$GOBIN` environment variable to change the installation location. It must be in your `$PATH` for the protocol buffer compiler to find it. The protocol buffer compiler produces Go output when invoked with the `--go_out` flag. The parameter to the `--go_out` flag is the directory where you want the compiler to write your Go output. The compiler creates a single source file for each `.proto` file input. The name of the output file is created by replacing the `.proto` extension with `.pb.go`. The `.proto` file should contain a `go_package` option specifying the full import path of the Go package that contains the generated code.

```protobuf

option go_package = "example.com/foo/bar";

```

The subdirectory of the output directory the output file is placed in depends on the `go_package` option and the compiler flags:

- By default, the output file is placed in a directory named after the Go package's import path. For example, a file `protos/foo.proto` with the above `go_package` option results in a file named `example.com/foo/bar/foo.pb.go`.
- If the `--go_opt=module=$PREFIX` flag is given to `protoc`, the specified directory prefix is removed from the output filename. For example, a file `protos/foo.proto` with the above `go_package` option and the flag `--go_opt=module=example.com/foo` results in a file named `bar/foo.pb.go`.
- If the `--go_opt=paths=source_relative` flag is given to `protoc`, the output file is placed in the same relative directory as the input file. For example, the file `protos/foo.proto` results in a file named `protos/foo.pb.go`.

When you run the proto compiler like this:

```sh

protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/foo.proto src/bar/baz.proto

```

the compiler will read the files `src/foo.proto` and `src/bar/baz.proto`. It produces two output files: `build/gen/foo.pb.go` and `build/gen/bar/baz.pb.go`.

The compiler automatically creates the directory `build/gen/bar` if necessary, but it will *not* create `build` or `build/gen`; they must already exist.

## Packages

Source `.proto` files should contain a `go_package` option specifying the full Go import path for the package containing the file. If there is no `go_package` option, the compiler will try to guess at one. A future release of the compiler will make the `go_package` option a requirement.

The Go package name of generated code will be the last path component of the `go_package` option. For example:

```protobuf

// The Go package name is "timestamppb".
option go_package = "google.golang.org/protobuf/types/known/timestamppb";

```

The import path is used to determine which import statements must be generated when one `.proto` file imports another `.proto` file. For example, if `a.proto` imports `b.proto`, the generated `a.pb.go` file needs to import the Go package which contains the generated `b.pb.go` file (unless both files are in the same package).

The import path is also used to construct output filenames. See the "Compiler Invocation" section above for details.

The `go_package` option may also include an explicit package name separated from the import path by a semicolon. For example: "example.com/foo;package_name". This usage is discouraged, since it is almost always clearer for the package name to correspond to the import path (the default). As an alternative to the `go_package` option, the Go import path for a `.proto` file may be specified on the command line with the `--go_opt=M$FILENAME=$IMPORT_PATH` flag to `protoc`.

## Messages

Given a simple message declaration:

```protobuf

message Foo {}

```

the protocol buffer compiler generates a struct called `Foo`. A `*Foo` implements the [proto.Message](https://pkg.go.dev/google.golang.org/protobuf/proto#Message) interface.

The [`proto` package](https://pkg.go.dev/google.golang.org/protobuf/proto) provides functions which operate on messages, including conversion to and from binary format.

The `proto.Message` interface defines a `ProtoReflect` method. This method returns a [protoreflect.Message](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#Message) which provides a reflection-based view of the message.

The `optimize_for` option does not affect the output of the Go code generator.

### Nested Types

A message can be declared inside another message. For example:

```protobuf

message Foo {
  message Bar {
  }
}

```

In this case, the compiler generates two structs: ``Foo`` and `Foo_Bar`.

### Fields

The protocol buffer compiler generates a struct field for each field defined within a message. The exact nature of this field depends on its type and whether it is a singular, repeated, map, or oneof field.

Note that the generated Go field names always use camel-case naming, even if the field name in the `.proto` file uses lower-case with underscores ([as it should](#)). The case-conversion works as follows:

    1. The first letter is capitalized for export. If the first character is an underscore, it is removed and a capital X is prepended.
    2. If an interior underscore is followed by a lower-case letter, the underscore is removed, and the following letter is capitalized.

Thus, the proto field `foo_bar_baz` becomes `FooBarBaz` in Go, and `_my_field_name_2` becomes `XMyFieldName_2`.

#### Singular Scalar Fields (proto2)

For either of these field definitions:

```protobuf

optional int32 foo = 1;
required int32 foo = 1;

```

the compiler generates a struct with an `*int32` field named `Foo` and an accessor method `GetFoo()` which returns the `int32` value in `Foo` or the default value if the field is unset. If the default is not explicitly set, the [zero value](https://golang.org/ref/spec#The_zero_value) of that type is used instead (`0` for numbers, the empty string for strings).

For other scalar field types (including `bool`, `bytes`, and `string`), `*int32` is replaced with the corresponding Go type according to the [scalar value types table](#).

#### Singular Scalar Fields (proto3)

For this field definition:

```protobuf

int32 foo = 1;

```

The compiler will generate a struct with an `int32` field named `Foo` and an accessor method `GetFoo()` which returns the `int32` value in `Foo` or the [zero value](https://golang.org/ref/spec#The_zero_value) of that type if the field is unset (`0` for numbers, the empty string for strings).

For other scalar field types (including `bool`, `bytes`, and `string`), `int32` is replaced with the corresponding Go type according to the [scalar value types table](#). Unset values in the proto will be represented as the [zero value](https://golang.org/ref/spec#The_zero_value) of that type (`0` for numbers, the empty string for strings).

#### Singular Message Fields

Given the message type:

```protobuf

message Bar {}

```

For a message with a `Bar` field:

```protobuf

// proto2
message Baz {
  optional Bar foo = 1;
  // The generated code is the same result if required instead of optional.
}

// proto3
message Baz {
  Bar foo = 1;
}

```

The compiler will generate a Go struct

```protobuf

type Baz struct {
        Foo *Bar
}

```

Message fields can be set to `nil`, which means that the field is unset, effectively clearing the field. This is not equivalent to setting the value to an "empty" instance of the message struct.

The compiler also generates a `func (m *Baz) GetFoo() *Bar` helper function. This function returns a `nil` `*Bar` if `m` is nil or `foo` is unset. This makes it possible to chain get calls without intermediate `nil` checks.

#### Repeated Fields

Each repeated field generates a slice of `T` field in the struct in Go, where `T` is the field's element type. For this message with a repeated field:

```protobuf

message Baz {
  repeated Bar foo = 1;
}

```

the compiler generates the Go struct:

```protobuf

type Baz struct {
        Foo  []*Bar
}

```

Likewise, for the field definition `repeated bytes foo = 1;` the compiler will generate a Go struct with a `[][]byte` field named `Foo`. For a repeated [enumeration](#) `repeated MyEnum bar = 2;`, the compiler generates a struct with a `[]MyEnum` field called `Bar`.

The following example shows how to set the field:

```go

baz := &Baz{
  Foo: []*Bar{
    {}, // First element.
    {}, // Second element.
  },
}

```

To access the field, you can do the following:

```go

foo := baz.GetFoo() // foo type is []*Bar.
b1 := foo[0] // b1 type is *Bar, the first element in foo.

```

#### Map Fields

Each map field generates a field in the struct of type `map[TKey]TValue` where `TKey` is the field's key type and `TValue` is the field's value type. For this message with a map field:

```protobuf

message Bar {}

message Baz {
  map<string, Bar> foo = 1;
}

```

the compiler generates the Go struct:


type Baz struct {
        Foo map[string]*Bar
}

#### Oneof Fields

For a oneof field, the protobuf compiler generates a single field with an interface type `isMessageName_MyField`. It also generates a struct for each of the [singular fields](#) within the oneof. These all implement this `isMessageName_MyField` interface.

For this message with a oneof field:

```protobuf

package account;
message Profile {
  oneof avatar {
    string image_url = 1;
    bytes image_data = 2;
  }
}

```

the compiler generates the structs:

```protobuf

type Profile struct {
        // Types that are valid to be assigned to Avatar:
        //      *Profile_ImageUrl
        //      *Profile_ImageData
        Avatar isProfile_Avatar `protobuf_oneof:"avatar"`
}

type Profile_ImageUrl struct {
        ImageUrl string
}
type Profile_ImageData struct {
        ImageData []byte
}

```

Both `*Profile_ImageUrl` and `*Profile_ImageData` implement `isProfile_Avatar` by providing an empty `isProfile_Avatar()` method.

The following example shows how to set the field:

```go

p1 := &account.Profile{
  Avatar: &account.Profile_ImageUrl{"http://example.com/image.png"},
}

// imageData is []byte
imageData := getImageData()
p2 := &account.Profile{
  Avatar: &account.Profile_ImageData{imageData},
}

```

To access the field, you can use a type switch on the value to handle the different message types.

```go

switch x := m.Avatar.(type) {
case *account.Profile_ImageUrl:
        // Load profile image based on URL
        // using x.ImageUrl
case *account.Profile_ImageData:
        // Load profile image based on bytes
        // using x.ImageData
case nil:
        // The field is not set.
default:
        return fmt.Errorf("Profile.Avatar has unexpected type %T", x)
}

```

The compiler also generates get methods `func (m *Profile) GetImageUrl() string` and `func (m *Profile) GetImageData() []byte`. Each get function returns the value for that field or the zero value if it is not set.

## Enumerations

Given an enumeration like:

```protobuf

message SearchRequest {
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 1;
  ...
}

```

the protocol buffer compiler generates a type and a series of constants with that type.

For enums within a message (like the one above), the type name begins with the message name:

```go

type SearchRequest_Corpus int32

```

For a package-level enum:

```protobuf

enum Foo {
  DEFAULT_BAR = 0;
  BAR_BELLS = 1;
  BAR_B_CUE = 2;
}

```

the Go type name is unmodified from the proto enum name:

```go

type Foo int32

```

This type has a `String()` method that returns the name of a given value.

The `Enum()` method initializes freshly allocated memory with a given value and returns the corresponding pointer:

```go

func (Foo) Enum() *Foo

```

The protocol buffer compiler generates a constant for each value in the enum. For enums within a message, the constants begin with the enclosing message's name:

```go

const (
        SearchRequest_UNIVERSAL SearchRequest_Corpus = 0
        SearchRequest_WEB       SearchRequest_Corpus = 1
        SearchRequest_IMAGES    SearchRequest_Corpus = 2
        SearchRequest_LOCAL     SearchRequest_Corpus = 3
        SearchRequest_NEWS      SearchRequest_Corpus = 4
        SearchRequest_PRODUCTS  SearchRequest_Corpus = 5
        SearchRequest_VIDEO     SearchRequest_Corpus = 6
)

```

For a package-level enum, the constants begin with the enum name instead:

```go

const (
        Foo_DEFAULT_BAR Foo = 0
        Foo_BAR_BELLS   Foo = 1
        Foo_BAR_B_CUE   Foo = 2
)

```

The protobuf compiler also generates a map from integer values to the string names and a map from the names to the values:

```go

var Foo_name = map[int32]string{
        0: "DEFAULT_BAR",
        1: "BAR_BELLS",
        2: "BAR_B_CUE",
}
var Foo_value = map[string]int32{
        "DEFAULT_BAR": 0,
        "BAR_BELLS":   1,
        "BAR_B_CUE":   2,
}

```

Note that the `.proto` language allows multiple enum symbols to have the same numeric value. Symbols with the same numeric value are synonyms. These are represented in Go in exactly the same way, with multiple names corresponding to the same numeric value. The reverse mapping contains a single entry for the numeric value to the name which appears first in the `.proto` file.

## Extensions (proto2)

Given an extension definition:

```protobuf

extend Foo {
  optional int32 bar = 123;
}

```

The protocol buffer compiler will generate an [protoreflect.ExtensionType](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#ExtensionType) value named `E_Bar`. This value may be used with the [proto.GetExtension](https://pkg.go.dev/google.golang.org/protobuf/proto#GetExtension), [proto.SetExtension](https://pkg.go.dev/google.golang.org/protobuf/proto#SetExtension), [proto.HasExtension](https://pkg.go.dev/google.golang.org/protobuf/proto#HasExtension), and [proto.ClearExtension](https://pkg.go.dev/google.golang.org/protobuf/proto#ClearExtension) functions to access an extension in a message. The `GetExtension` function and `SetExtension` functions respectively accept and return an `interface{}` value containing the extension value type.

For singular scalar extension fields, the extension value type is the corresponding Go type from the [scalar value types table](#).

For singular embedded message extension fields, the extension value type is `*M`, where `M` is the field message type.

For repeated extension fields, the extension value type is a slice of the singular type.

For example, given the following definition:

```protobuf

extend Foo {
  optional int32 singular_int32 = 1;
  repeated bytes repeated_string = 2;
  optional Bar repeated_message = 3;
}

```

Extension values may be accessed as:

```go

m := &somepb.Foo{}
proto.SetExtension(m, extpb.E_SingularInt32, int32(1))
proto.SetExtension(m, extpb.E_RepeatedString, []string{"a", "b", "c"})
proto.SetExtension(m, extpb.E_RepeatedMessage, &extpb.Bar{})

v1 := proto.GetExtension(m, extpb.E_SingularInt32).(int32)
v2 := proto.GetExtension(m, extpb.E_RepeatedString).([][]byte)
v3 := proto.GetExtension(m, extpb.E_RepeatedMessage).(*extpb.Bar)

```

Extensions can be declared nested inside of another type. For example, a common pattern is to do something like this:

```protobuf

message Baz {
  extend Foo {
    optional Baz foo_ext = 124;
  }
}

```

In this case, the `ExtensionType` value is named `E_Baz_Foo`.

Services

The Go code generator does not produce output for services by default. If you enable the [gRPC](https://www.grpc.io) plugin (see the [gRPC Go Quickstart guide](https://github.com/grpc/grpc-go/tree/master/examples)) then code will be generated to support gRPC.

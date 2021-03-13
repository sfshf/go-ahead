# [Protocol Buffers Version 2 Language Specification](https://developers.google.com/protocol-buffers/docs/reference/proto2-spec)

This is a language specification reference for version 2 of the Protocol Buffers language (proto2). The syntax is specified using [Extended Backus-Naur Form (EBNF)](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_Form):

```txt

|   alternation
()  grouping
[]  option (zero or one time)
{}  repetition (any number of times)

```

For more information about using proto2, see the [language guide](#).

## Lexical elements

### Letters and digits

```txt

letter = "A" … "Z" | "a" … "z"
capitalLetter =  "A" … "Z"
decimalDigit = "0" … "9"
octalDigit   = "0" … "7"
hexDigit     = "0" … "9" | "A" … "F" | "a" … "f"

```

### Identifiers

```txt

ident = letter { letter | decimalDigit | "_" }
fullIdent = ident { "." ident }
messageName = ident
enumName = ident
fieldName = ident
oneofName = ident
mapName = ident
serviceName = ident
rpcName = ident
streamName = ident
messageType = [ "." ] { ident "." } messageName
enumType = [ "." ] { ident "." } enumName
groupName = capitalLetter { letter | decimalDigit | "_" }

```

### Integer literals

```txt

intLit     = decimalLit | octalLit | hexLit
decimalLit = ( "1" … "9" ) { decimalDigit }
octalLit   = "0" { octalDigit }
hexLit     = "0" ( "x" | "X" ) hexDigit { hexDigit }

```

### Floating-point literals

```txt

floatLit = ( decimals "." [ decimals ] [ exponent ] | decimals exponent | "."decimals [ exponent ] ) | "inf" | "nan"
decimals  = decimalDigit { decimalDigit }
exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals

```

### Boolean

```txt

boolLit = "true" | "false"

```

### String literals

```txt

strLit = ( "'" { charValue } "'" ) | ( '"' { charValue } '"' )
charValue = hexEscape | octEscape | charEscape | /[^\0\n\\]/
hexEscape = '\' ( "x" | "X" ) hexDigit hexDigit
octEscape = '\' octalDigit octalDigit octalDigit
charEscape = '\' ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | '\' | "'" | '"' )
quote = "'" | '"'

```

### EmptyStatement

```txt

emptyStatement = ";"

```

### Constant

```txt

constant = fullIdent | ( [ "-" | "+" ] intLit ) | ( [ "-" | "+" ] floatLit ) |
                strLit | boolLit

```

## Syntax

The syntax statement is used to define the protobuf version.

```txt

syntax = "syntax" "=" quote "proto2" quote ";"

```

## Import Statement

The import statement is used to import another .proto's definitions.

```txt

import = "import" [ "weak" | "public" ] strLit ";"

```

Example:

```protobuf

import public "other.proto";

```

## Package

The package specifier can be used to prevent name clashes between protocol message types.

```txt

package = "package" fullIdent ";"

```

Example:

```protobuf

package foo.bar;

```

## Option

Options can be used in proto files, messages, enums and services. An option can be a protobuf defined option or a custom option. For more information, see [Options](#) in the language guide.

```txt

option = "option" optionName  "=" constant ";"
optionName = ( ident | "(" fullIdent ")" ) { "." ident }

```

For examples:

```protobuf

option java_package = "com.example.foo";

```

## Fields

Fields are the basic elements of a protocol buffer message. Fields can be normal fields, group fields, oneof fields, or map fields. A field has a label, type and field number.

```txt

label = "required" | "optional" | "repeated"
type = "double" | "float" | "int32" | "int64" | "uint32" | "uint64"
      | "sint32" | "sint64" | "fixed32" | "fixed64" | "sfixed32" | "sfixed64"
      | "bool" | "string" | "bytes" | messageType | enumType
fieldNumber = intLit;

```

### Normal field

Each field has label, type, name and field number. It may have field options.

```txt

field = label type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
fieldOptions = fieldOption { ","  fieldOption }
fieldOption = optionName "=" constant

```

Examples:

```protobuf

optional foo.bar nested_message = 2;
repeated int32 samples = 4 [packed=true];

```

### Group field

**Note that this feature is deprecated and should not be used when creating new message types – use nested message types instead.**

Groups are one way to nest information in message definitions. The group name must begin with capital letter.

```txt

group = label "group" groupName "=" fieldNumber messageBody

```

Example:

```protobuf

repeated group Result = 1 {
    required string url = 2;
    optional string title = 3;
    repeated string snippets = 4;
}

```

### Oneof and oneof field

A oneof consists of oneof fields and a oneof name. Oneof fields do not have labels.

```txt

oneof = "oneof" oneofName "{" { option | oneofField | emptyStatement } "}"
oneofField = type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"

```

Example:

```protobuf

oneof foo {
    string name = 4;
    SubMessage sub_message = 9;
}

```

### Map field

A map field has a key type, value type, name, and field number. The key type can be any integral or string type. Note, the key type may not be an enum.

```txt

mapField = "map" "<" keyType "," type ">" mapName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
keyType = "int32" | "int64" | "uint32" | "uint64" | "sint32" | "sint64" |
          "fixed32" | "fixed64" | "sfixed32" | "sfixed64" | "bool" | "string"

```

Example:

```protobuf

map<string, Project> projects = 3;

```

## Extensions and Reserved

Extensions and reserved are message elements that declare a range of field numbers or field names.

### Extensions

Extensions declare that a range of field numbers in a message are available for third-party extensions. Other people can declare new fields for your message type with those numeric tags in their own .proto files without having to edit the original file.

```txt

extensions = "extensions" ranges ";"
ranges = range { "," range }
range =  intLit [ "to" ( intLit | "max" ) ]

```

Examples:

```protobuf

extensions 100 to 199;
extensions 4, 20 to max;

```

### Reserved

Reserved declares a range of field numbers or field names in a message that can not be used.

```txt

reserved = "reserved" ( ranges | fieldNames ) ";"
fieldNames = fieldName { "," fieldName }

```

Examples:

```protobuf

reserved 2, 15, 9 to 11;
reserved "foo", "bar";

```

## Top Level definitions

### Enum definition

The enum definition consists of a name and an enum body. The enum body can have options and enum fields.

```txt

enum = "enum" enumName enumBody
enumBody = "{" { option | enumField | emptyStatement } "}"
enumField = ident "=" [ "-" ] intLit [ "[" enumValueOption { ","  enumValueOption } "]" ]";"
enumValueOption = optionName "=" constant

```

Example:

```protobuf

enum EnumAllowingAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
}

```

### Message definition

A message consists of a message name and a message body. The message body can have fields, nested enum definitions, nested message definitions, extend statements, extensions, groups, options, oneofs, map fields, and reserved statements.

```txt

message = "message" messageName messageBody
messageBody = "{" { field | enum | message | extend | extensions | group |
option | oneof | mapField | reserved | emptyStatement } "}"

```

Example:

```protobuf

message Outer {
  option (my_option).a = true;
  message Inner {   // Level 2
    required int64 ival = 1;
  }
  map<int32, string> my_map = 2;
  extensions 20 to 30;
}

```

### Extend

If a message in the same or imported .proto file has reserved a range for extensions, the message can be extended.

```txt

extend = "extend" messageType "{" {field | group | emptyStatement} "}"

```

Example:

```protobuf

extend Foo {
  optional int32 bar = 126;
}

```

### Service definition

```txt

service = "service" serviceName "{" { option | rpc | stream | emptyStatement } "}"
rpc = "rpc" rpcName "(" [ "stream" ] messageType ")" "returns" "(" [ "stream" ]
messageType ")" (( "{" { option | emptyStatement } "}" ) | ";" )
stream = "stream" streamName "(" messageType "," messageType ")" (( "{"
{ option | emptyStatement } "}") | ";" )

```

Example:

```protobuf

service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}

```

## Proto file

```txt

proto = syntax { import | package | option | topLevelDef | emptyStatement }
topLevelDef = message | enum | extend | service

```

An example .proto file:

```protobuf

syntax = "proto2";
import public "other.proto";
option java_package = "com.example.foo";
enum EnumAllowingAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
}
message Outer {
  option (my_option).a = true;
  message Inner {   // Level 2
    required int64 ival = 1;
  }
  repeated Inner inner_message = 2;
  optional EnumAllowingAlias enum_field = 3;
  map<int32, string> my_map = 4;
  extensions 20 to 30;
}
message Foo {
  optional group GroupMessage {
    optional a = 1;
  }
}

```

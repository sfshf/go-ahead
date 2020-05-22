

# [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)

`mapstructure`是一个`Go库`，用于将通用`map`值解码为`结构`，反之亦可；同时还提供有用的`错误处理`。

当从某些您不是相当了解其底层数据结构，直到读取一部分才能了解的`数据流`（`JSON`，`Gob`等）中解码出值时，该库非常有用。因此，您可以读取`map[string]interface{}`，并使用此库将其解码为合适的底层的本地`Go语言结构`。


## But Why?!

`Go语言`提供了出色的标准库来解码`JSON`等格式。`标准方法是预先创建一个结构体，然后从编码格式的字节中填充该结构体`。很好，但问题是您的配置或编码会根据特定字段而略有变化。例如，考虑以下`JSON`：

```json

{
    "type": "person",
    "name": "Mitchell"
}

```

也许我们必须先从`JSON`中读取`"type"`字段，然后才能填充特定的结构；我们总是对`JSON`的解码进行两次传递（首先读取`"type"`，然后读取其余信息）。但是，将其解码为`map[string]interface{}`结构，读取`"type"`键，然后使用类似`mitchellh/mapstructure`库的东西将其解码为合适的结构，会简单很多。


## Usage & Example

`mapstructure`包，公开了将任意一种`Go语言类型`转换为另一种`Go语言类型`的功能，通常将`map[string]interface{}`转换为`本地Go语言结构`。

`Go语言结构`可以是任意复杂的，包含`切片`，其他`结构体`等，并且`解码器`将正确解码`嵌套的映射`等，成为`本地Go语言结构体`中的适当`结构`。请参阅示例以了解`解码器`的功能。

首先开始的最简单的函数是`Decode`。


### Field Tags（字段标记）

解码为`结构体`时，`默认情况下`，`mapstructure`将使用`字段名称`来执行映射。例如，如果`结构体`具有字段`Username`，那么`mapstructure`将在`username`的`源值`中查找`键`（`不区分大小写`）。

```go

type User struct {
    Username string
}

```

您可以使用`结构体标记`来更改`mapstructure`的行为。`mapstructure`查找的默认`结构体标记`是`"mapstructure"`，但是您可以使用`DecoderConfig`对其进行`自定义`。


### Renaming Fields（重命名字段）

要`重命名``mapstructure`查找的`键`，请使用`"mapstructure"标记`并`直接设置一个值`。例如，要将上面的`"username"`示例更改为`"user"`：

```go

type User struct {
    Username string `mapstructure:"user"`
}

```


### Embedded Structs and Squashing（内嵌结构体和挤压）

`内嵌结构体`，被视为是具有该名称的一个字段。`默认情况下`，使用`mapstructure`进行解码时，以下`两个结构体`是等效的：

```go

type Person struct {
    Name string
}

type Friend struct {
    Person
}

type Friend struct {
    Person Person
}

```

这将需要如下所示的输入：

```go

map[string]interface{}{
    "person": map[string]interface{}{"name": "alice"},
}

```

如果您的`"person"`值不是内嵌的，则可以将`",squash"`追加到`标记值`中，`mapstructure`会将该`内嵌结构体`视为`结构体`的一部分。示例：

```go

type Friend struct {
    Person `mapstructure:",squash"`
}

```

现在可以接收如下输入：

```go

map[string]interface{}{
    "name": "alice",
}

```

`DecoderConfig`具有一个字段，该字段可以将`mapstructure`的行为更改为`始终挤压内嵌结构`。


### Remainder Values（剩余的值）

如果`源值`中有任何未映射的`键`，则`默认情况下`，`mapstructure`将静默忽略它们。您可以通过在`DecoderConfig`中设置`ErrorUnused`来报错。如果您使用的是`元数据`，则您还可以维护一个`未使用的键`的`切片`。

您还可以在`标记`上使用`",remain"后缀`来收集映射中所有`未使用的值`。`带有此标记的字段必须为map类型，并且应该为"map[string]interface{}"或"map[interface{}]interface{}"`。请参见如下示例：

```go

type Friend struct {
    Name string
    Other map[string]interface{} `mapstructure:",remain"`
}

```

给定如下输入，将用`其他未使用的值`（`"name"`以外的所有值）来填充`Other`字段：

```go

map[string]interface{}{
    "name": "bob",
    "address": "123 Maple St.",
}

```


### Other Configuration（其他配置）

`mapstructure`是`高度可配置的`。有关其他被支持的功能和选项，请参见`DecoderConfig`结构体。


### 示例代码

- [解码为结构体](struct/)
- [解码为内嵌结构体](embeddedstruct/)
- [报错的处理](errors/)
- [元数据的使用](metadata/)
- [剩余数据的处理](remainingdata/)
- [标记的使用](tags/)
- [弱类型输入的使用](weaklytypedinput/)

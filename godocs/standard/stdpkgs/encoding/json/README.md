
# `encoding/json`包

`json`包实现了[RFC 7159](https://www.rfc-editor.org/info/rfc7159)中定义的`JSON编码和解码`。`JSON`和`Go值`之间的映射在`Marshal`和`Unmarshal`函数的文档中进行了描述。

有关此程序包的介绍，请参见`"JSON and Go"`：[https://golang.org/doc/articles/json_and_go.html](https://golang.org/doc/articles/json_and_go.html)。


## Marshal

```go

func Marshal(v interface{}) ([]byte, error)

```

`Marshal`返回`v`的`JSON编码`。

`Marshal`会`递归地`遍历值`v`。如果`要进行编码的值`实现了`Marshaler接口`并且`不是nil指针`，则`Marshal`会调用`该要进行编码的值`的`MarshalJSON方法`来生成JSON。`如果不存在MarshalJSON方法`，但该值实现`encoding.TextMarshaler`，则`Marshal`调用`其MarshalText方法`并将结果编码为JSON字符串。`nil指针异常`不是严格必需的，只是模仿了`UnmarshalJSON`行为（`UnmarshalJSON`中该异常是必需的）。

其他地，`Marshal`将使用以下`类型相关的默认编码`：

`布尔值`编码为`JSON布尔值`。

`浮点数`，`整数`和`数字值`编码为`JSON数字`。

`字符串值`编码为`强制转换`为`有效UTF-8的JSON字符串`，用`Unicode替换字符`替换`无效字节`。为了使`JSON`可以安全地嵌入`HTML <script>标记`中，`使用HTMLEscape对字符串进行编码`，会将`<`、`>`、`&`、`U+2028`和`U+2029`转义为`\u003c`、`\u003e`、`\u0026`、`\u2028`和`\u2029`。在使用`Encoder`时，可以通过调用`SetEscapeHTML(false)`来禁用该替换功能。

`数组和切片值`编码为`JSON数组`，但`[]byte`编码为`base64编码的字符串`，而`空切片`编码为`空JSON值`。

`结构值`编码为`JSON对象`。每个`导出的结构字段`都将成为对象的成员，使用`字段名`称作为`JSON对象关键字（key）`，除非出于之后要讲述的原因之一省略了该字段。

每个`结构字段的编码`可以通过存储在`结构字段的标签`中`json`关键字下的`格式字符串`进行自定义。`格式字符串`提供`字段在JSON里的名称`，可能后跟`英文逗号分隔的选项列表`。`名称可以为空`，以便在不覆盖默认字段名称的情况下指定选项。

`omitempty`选项，指定如果字段具有空值（定义为`false`，`0`，`空指针`，`空接口值`以及`任何空数组，切片，map或字符串`），则应`将该字段从编码中省略`。

`作为特殊情况`，如果`字段标签`为`-`，则会`始终省略该字段`。请注意，名称为`-`的字段仍可以使用标签`-`来生成。

`结构字段标签及其含义`的示例：

```go

// Field appears in JSON as key "myName"
Field int `json:"myName"`

// Field appears in JSON as key "myName" and the field is omitted from the object if its value is empty, as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`

```

`string`选项，表示`该字段`在`JSON编码的字符串`中存储为`JSON`。它仅适用于`字符串`，`浮点数`，`整数`或`布尔类型`的字段。这种额外的编码级别在与JavaScript程序进行通信时有时会使用：

```go

Int64String int64 `json:",string"`

```

如果`关键字的名称`是一个`非空字符串`，且该字符串是由`Unicode字母`，`数字`和`ASCII标点（引号，反斜杠和逗号除外）`组成，则将使用`该关键字的名称`。

`匿名结构的字段`通常就好像`它们内部的导出字段`是`外部结构的字段`一样地进行打包，且遵从`通常的Go可视规则`（如下一段文字中描述的进行修缮）。在`其JSON标签中具有名称的匿名结构字段`被视为具有该名称，而`不是匿名的`。`接口类型的匿名结构字段`将使用`该接口类型的名称`作为其字段名称，而`不是被视为匿名`。

在决定哪些字段要进行打包或解包时，为JSON修缮了`结构字段的Go可见性规则`。如果在同一层级上有多个字段，并且该层级是嵌套最少的（并且因此将是`通常的Go规则`选择的嵌套层级），则适用以下额外规则：

1）在这些字段中，如果有任何字段`带有JSON标签`，则即使存在多个`非标记字段`（这些`非标记字段`在没有`标记字段`时会发生冲突），也`仅考虑标记字段`。

2）如果`正好只有一个字段`（根据第一条规则标记或未标记），则将其选中。

3）否则，存在多个字段，所有字段`都将被忽略`；`不会有错误发生`。

`Go 1.1`中新增了处理`匿名结构字段`的功能。`在Go 1.1之前`，`匿名结构字段`会被忽略。要强制忽略`当前版本和早期版本中`的`匿名结构字段`，请为该字段指定`JSON标记"-"`。

`map值`会编码为`JSON对象`。`map的键类型`必须是`字符串`，`整数类型`或`实现encoding.TextMarshaler`。通过应用以下规则对`map键`进行`排序`并`用作JSON对象键`，还要遵守上述针对字符串值描述的`强制UTF-8`：

```txt

- keys of any string type are used directly
- encoding.TextMarshalers are marshaled
- integer keys are converted to strings

```

`指针值`会编码为指向的值。`空指针`编码为`空JSON值`。

`接口值`会编码为接口中包含的值。`空接口值`编码为`空JSON值`。

`通道`，`复数`和`函数`值`不能用JSON编码`。尝试对这样的值进行编码会导致`Marshal`返回`UnsupportedTypeError`。

JSON无法表示`循环的数据结构`，并且`Marshal`不会处理它们。将`循环结构`传递给`Marshal`将导致错误。


## UnMarshal

```go

func Unmarshal(data []byte, v interface{}) error

```

`Unmarshal`会解析`JSON编码的数据`，并将结果存储在`v`指向的值中。如果`v`为`nil或不是指针`，则`Unmarshal`返回`InvalidUnmarshalError`。

`Unmarshal`使用`Marshal`使用的编码的逆过程，根据需要和以下附加规则来分配`map`，`切片`和`指针`：

为了将`JSON`解组为一个`指针`，`Unmarshal`首先处理`JSON`为`JSON常量null`的情况。在这种情况下，`Unmarshal`会将`指针`设置为`nil`。否则，`Unmarshal`会将`JSON`解组为`指针所指向的值`。如果`指针`为`nil`，则`Unmarshal`为其`分配一个新值`以使其指向。

为了将`JSON`解组为`实现Unmarshaler接口`的值，`Unmarshal`调用该值的`UnmarshalJSON`方法，包括当`输入`为`JSON null`时。否则，如果该值实现`encoding.TextUnmarshaler`且`输入`是`带引号的JSON字符串`，则`Unmarshal`会使用`该字符串的无引号形式`来调用该值的`UnmarshalText`方法。

为了将`JSON`解组到`结构`中，`Unmarshal`将`传入的对象的键`与`Marshal使用的键`（结构字段名称或其标记）进行匹配，`最好使用精确匹配`，但还可以`接受不区分大小写的匹配`。默认情况下，`没有对应结构字段的对象键`将被忽略（有关替代方法，请参见`Decoder.DisallowUnknownFields`）。

要将`JSON`解组为`接口值`，`Unmarshal`将如下的其中之一存储在`接口值`中：

```txt

bool,                       for JSON booleans
float64,                    for JSON numbers
string,                     for JSON strings
[]interface{},              for JSON arrays
map[string]interface{},     for JSON objects
nil,                        for JSON null

```

要将一个`JSON数组`解组到`切片`中，`Unmarshal`将切片长度重置为零，然后将每个元素添加到切片中。`作为一种特殊情况`，要将一个`空的JSON数组`解组到一个切片中，`Unmarshal`会用一个新的空切片替换该切片。

为了将`JSON数组`解组为`Go数组`，`Unmarshal`将`JSON数组元素`解码为`对应的Go数组元素`。`如果Go数组小于JSON数组，则其他JSON数组元素将被丢弃`。`如果JSON数组小于Go数组，则将其他Go数组元素设置为零值`。

要将`JSON对象`解组到`map`中，`Unmarshal`首先要建立`要使用的map`。`如果map为空`，`Unmarshal`将分配一个`新map`。否则，`Unmarshal`会重用`现有的map，并保留现有条目`。然后，`Unmarshal`将来自`JSON对象的键值对`存储到`map`中。`map的键类型`必须是`字符串类型`，`整数`，`实现json.Unmarshaler`或`实现encoding.TextUnmarshaler`。

如果`JSON值不适用于给定的目标类型`，或者`JSON数字溢出目标类型`，则`Unmarshal`会跳过该字段并尽最大可能完成解组。如果没有遇到更多的严重错误，则`Unmarshal`返回`UnmarshalTypeError`来描述第一个的此类错误。无论如何，都不能保证有出现问题的字段之后的所有其余字段都将被解组到目标对象中。

通过将该Go值设置为`nil`，`JSON空值`可解组到`接口`，`map`，`指针`或`切片`中。由于在JSON中通常使用`null`来表示`不存在`，因此将`JSON null`解组到任何其他Go类型中不会影响该值，也不会产生任何错误。

解组`带引号的字符串`时，`无效的UTF-8`或`无效的UTF-16`的`替代对`不会被视为错误。而是将它们替换为`Unicode替换字符U+FFFD`。

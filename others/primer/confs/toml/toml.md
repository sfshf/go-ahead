

# [TOML -- Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)


## 目的

作为一种易读的，语义清晰的，极小的，配置文件格式；能清晰地映射到一个散列表；可以在很多编程语言的应用中简单地解析为数据结构。


## 说明

- 大小写敏感；
- 必须是有效的`UTF-8`编码的文件；
- `空白字符（whitespace）`是指`tab(0x09)`或`space(0x20)`；
- `换行符（newline）`是指`LF(0x0A)`或`CRLF(0x0D 0x0A)`；



## 文件扩展名 -- `.toml`


## MIME类型 -- `application/toml`


## 注释 -- hash symbol（"#"符号）

除了`tab(0x09)`之外的`控制字符（即：U+0000至U+0008，U+000A至U+001F，以及U+007F）`不能出现在注释中。


## 键值对格式 -- toml文件基石

就像砌房子的砖头，toml文件使用`键值对（key/value pair）`来构建，格式如下：

```

key = "value"

```

注意：有`键`必有`值`；一行不能有多个`键值对`。


## 键：

`键`可以是`不带引号的（bare）`、`带引号的（quoted）`或`带点号的（dotted）`：
- `不带引号的`：只能由`ascii字母`、`ascii数字`、`下划线（_）`和`破折号（-）`组成，且非空；
- `带引号的`：可以单双引号嵌套；可以是空字符串；
- `带点的`：`不带引号的`或`带引号的`键之间可以使用`.`来组成序列，用来将相似的属性组织在一起；

注意：多次定义一个`键`是不允许的；定义`键`时要避免歧义；只要一个`键`没有直接被定义，后续还是可以定义它或其内部的键的，举例如下：

```toml
# THIS IS VALID
fruit.apple.smooth = true
fruit.apple = 1
```
```toml
# THIS IS INVALID
fruit.apple = 1
fruit.apple.smooth = true
```

## 值：

`值`可以是下列类型之一：

- [字符串](#字符串)
- [整型](#整型)
- [浮点型](#浮点型)
- [布尔型](#布尔型)
- [偏移式日期时间](#偏移式日期时间)
- [本地日期时间](#本地日期时间)
- [本地日期](#本地日期)
- [本地时间](#本地时间)
- [数组](#数组)
- [内联表](#内联表)


## 字符串

有四种方式表达字符串：`基本量（basic）`、`多行基本量（multi-line basic）`、`字面量（literal）`、`多行字面量（multi-line literal）`；所有的字符串必须是由有效的`UTF-8`字符组成。

`基本量（basic）：`

字符串基本量是使用`双引号（quotation marks）`括起来的；可以包含除了`必须要转义的字符`之外的任何`Unicode`字符；`必须要转义的字符`有：`单引号（quotation mark）`、`反斜杠（backslash）`和`控制字符（除了tab之外）`。为了方便，一些常用的必须转义的字符有`压缩的转义序列`，如下：

```
\b         - backspace       (U+0008)
\t         - tab             (U+0009)
\n         - linefeed        (U+000A)
\f         - form feed       (U+000C)
\r         - carriage return (U+000D)
\"         - quote           (U+0022)
\\         - backslash       (U+005C)
\uXXXX     - unicode         (U+XXXX)
\UXXXXXXXX - unicode         (U+XXXXXXXX)

```

任何`Unicode`字符都能以`\uXXXX`或`\uXXXXXXXX`的形式进行转义；被转义的编码必须是有效的[`Unicode scalar values`](http://unicode.org/glossary/#unicode_scalar_value)。

上面没有列举出的转义序列都是保留的，如果使用保留的转义序列，TOML会产生错误。


`多行基本量（multi-line basic）：`

字符串多行基本量是使用`三个双引号`括起来的;字符串内容`不能出现连续三个的双引号`。

`字面量（literal）：`

字符串字面量里，转义字符将失去效果；字符串字面量是使用`单引号（quotation mark）`括起来的。

`多行字面量（multi-line literal）：`

多行字符串字面量是使用`三个单引号`括起来的；字符串内容`不能出现连续三个的单引号`。

更详细的描述，请查看[原文](https://github.com/toml-lang/toml#string)。


## 整型

`十进制整数`不能以`0`开头；正数前面可以加一个`+`号；负数前面加一个`-`号；比较长的数可以使用`_`来分隔。

非负整数可以使用`十六进制`、`八进制`或`二进制`来表示；前面不用加一个`+`号；可以使用`0`开头；`十六进制`数大小写不敏感；比较长的数可以使用`_`来分隔。

整型可以支持`64bit`。

更详细的描述，请查看[原文](https://github.com/toml-lang/toml#integer)。


## 浮点型

浮点型数字按`IEEE 754 binary64`来实现的。

浮点型支持`小数表示方式`、`指数表示方式`和`混合方式`；正数前面可以加一个`+`号；负数前面加一个`-`号。

更详细的描述，请查看[原文](https://github.com/toml-lang/toml#float)。


## 布尔型

布尔型使用小写方式 -- true/false。


## 偏移式日期时间

时间格式如下：

```toml
odt1 = 1979-05-27T07:32:00Z
odt2 = 1979-05-27T00:32:00-07:00
odt3 = 1979-05-27T00:32:00.999999-07:00

odt4 = 1979-05-27 07:32:00Z
```

注意：时间中`秒`的小数部分会精确到`毫秒`，如果超出`毫秒`，最终会被剪裁至`毫秒`，而不是四舍五入。

## 本地日期时间

时间格式如下：

```toml
ldt1 = 1979-05-27T07:32:00
ldt2 = 1979-05-27T00:32:00.999999
```

注意：时间中`秒`的小数部分会精确到`毫秒`，如果超出`毫秒`，最终会被剪裁至`毫秒`，而不是四舍五入。


## 本地日期

时间格式如下：

```toml
ldt1 = 1979-05-27
```


## 本地时间

时间格式如下：

```toml
lt1 = 07:32:00
lt2 = 00:32:00.999999
```

注意：时间中`秒`的小数部分会精确到`毫秒`，如果超出`毫秒`，最终会被剪裁至`毫秒`，而不是四舍五入。


## 数组

数组是使用`[]（square brackets）`括起来的；数组内可以包含任何类型的数据，数据的类型可以是混合的，数据之间用`,（comma）`分隔。


## 表

`表`，通常又叫做散列表或字典，是键值对的集合。TOML文档中，`表`的定义要独占一行，是使用`[]（square brackets）`括起来的，从下一行开始直到下一个表或EOF出现之前的所有键值对，都属于当前表的内容。

`表`的命令规则和`键`的命令规则一样；就像`键`的定义一样，`表`也不能重复定义。

空表也是允许的，即该表内没有键值对。


## 内联表

内联表是一个用来表示`表`的更压缩的语法，用来分组数据非常有用；内联表使用`{}（curly braces，大括号）`括起来，大括号里可以放零个或多个用`,`号分隔的键值对。

```toml
name = { first = "Tom", last = "Preston-Werner" }
# name = { first = "Tom", last = "Preston-Werner"， } # INVALID
point = { x = 1, y = 2 }
animal = { type.name = "pug" }
```
对应的标准表格式如下：

```toml
[name]
first = "Tom"
last = "Preston-Werner"

[point]
x = 1
y = 2

[animal]
type.name = "pug"
```

最后的键值对不能用`,`结尾；内联表正常使用在一行内；如果需要多行，请使用标准的[`表`](#表)格式。

内联表是完整独立的表；定义之后，不能在添加新的键值对或子表，例如：

```toml
[product]
type = { name = "Nail" }
# type.edible = false  # INVALID
```

反之，内联表也不能被用来作为新的添加对象，例如：

```toml
[product]
type.name = "Nail"
# type = { edible = false }  # INVALID
```


## 表的数组

表的数组是使用`两个中括号`括起来的；定义了一个表的数组，直到后续出现了另一个表或EOF，中间出现的使用同名数组的表都是同一数组里的元素。例如：

```toml
[[products]]
name = "Hammer"
sku = 738594937

[[products]]

[[products]]
name = "Nail"
sku = 284758393

color = "gray"
```
对应到`JSON`格式，就是如下结构：
```json
{
  "products": [
    { "name": "Hammer", "sku": 738594937 },
    { },
    { "name": "Nail", "sku": 284758393, "color": "gray" }
  ]
}
```

表的数组还可以嵌套表或表的数组，例如：

```toml
[[fruit]]
  name = "apple"

  [fruit.physical]  # subtable
    color = "red"
    shape = "round"

  [[fruit.variety]]  # nested array of tables
    name = "red delicious"

  [[fruit.variety]]
    name = "granny smith"

[[fruit]]
  name = "banana"

  [[fruit.variety]]
    name = "plantain"
```

相对应的`JSON`格式结构如下：

```json
{
  "fruit": [
    {
      "name": "apple",
      "physical": {
        "color": "red",
        "shape": "round"
      },
      "variety": [
        { "name": "red delicious" },
        { "name": "granny smith" }
      ]
    },
    {
      "name": "banana",
      "variety": [
        { "name": "plantain" }
      ]
    }
  ]
}
```

请注意，嵌套关系的使用必须要先定义父级表或表的数组，再定义子级表或表的数组；反之，则会出错。

尝试静态定义一个数组是要产生解析错误的，例如：

```toml
# INVALID TOML DOC
fruit = []

[[fruit]] # Not allowed
```

尝试定义与一个表的数组同名的表或与一个表同名的表的数组，都是要产生解析错误的，例如：
```toml
# INVALID TOML DOC
[[fruit]]
  name = "apple"

  [[fruit.variety]]
    name = "red delicious"

  # INVALID: This table conflicts with the previous array of tables
  [fruit.variety]
    name = "granny smith"

  [fruit.physical]
    color = "red"
    shape = "round"

  # INVALID: This array of tables conflicts with the previous table
  [[fruit.physical]]
    color = "green"
```

使用内联表的数组也是可行的，例如：
```toml
points = [ { x = 1, y = 2, z = 3 },
           { x = 7, y = 8, z = 9 },
           { x = 2, y = 4, z = 8 } ]
```


[Show documentation for package or symbol](https://golang.google.cn/cmd/go/#hdr-Show_documentation_for_package_or_symbol)


##### 一、用法

```

go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]

```

##### 二、doc命令说明

`doc命令`会打印与传入的参数所指明的项目相关联的文档注释（所传入的参数可以是一个包、一个常量、一个函数、一个类型、一个变量、一个方法或一个结构体字段）；打印出来的文档注释是所指明的项目下的每一个一级项目的一行概述（如果所指明的项目是包，则打印包级别的声明；如果所指明的项目是类型，则打印该类型的方法；等等以此类推；）。

`doc命令`接收零个、一个或两个参数。

**零个参数:**

零个参数，也就是如下所示运行：

```

go doc

```

将打印当前目录中包的包文档。如果该包是命令包（package main），则将从打印出的文稿中忽略包的导出标识符（exported symbol），除非提供了`-cmd`标志。

**一个参数:**

当使用一个参数运行时，该参数要采用被展示文档内容的项目的类Go语法（Go-syntax-like）表示形式。参数选择的内容取决于`GOROOT`和`GOPATH`中安装的内容，以及参数的形式；参数的形式示意如下：

```
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>

```

此列表中与参数匹配的第一项是打印其文档的项（可以参考下面的示例）。但是，如果参数以大写字母开头，则假定它标识当前目录中的标识符或方法。

对于包，扫描顺序是按词法上宽度优先顺序确定的。也就是说，呈现的包是要与搜索匹配，并且最接近根节点，还要在词法的层次结构上排第一。`GOROOT`树总是在`GOPATH`之前被完整地扫描。

如果没有指定或匹配的包，则选择当前目录中的包；因此执行`go doc Foo`会显示当前包中标识符`Foo`的文档。

包路径必须要么是限定的路径要么是一个路径的正确的后缀。`go`工具通常的包机制对一些包路径元素不适用，例如`.`和`...`；`go doc`工具没有实现对这些包路径元素的应用。

**两个参数:**

当使用两个参数运行时，第一个参数必须是完整的包路径（而不仅仅是后缀），第二个参数是一个标识符或者是带有方法或结构体字段的标识符。这与`godoc`接受的语法类似：

```

go doc <pkg> <sym>[.<methodField>]

```

在所有形式中，当匹配标识符时，参数中的小写字母能匹配不论大小写的字母，而大写字母只能精准匹配大写字母。这意味着，如果不同的标识符有不同的大小写，则包中的小写参数可能有多个匹配项。如果发生这种情况，将打印所有匹配项的文档。

示例：

| 命令 | 结果 |
|--|--|
| go doc | 显示当前包的文档。 |
| go doc Foo | 显示当前包中的关于标识符"Foo"的文档。（"Foo"是以大写字母开头，所以不会匹配到一个包路径。） |
| go doc encoding/json | 显示"encoding/json"包的文档。 |
| go doc json | 缩写；显示"encoding/json"包的文档。 |
| go doc json.Number (or go doc json.number) | 显示"json.Number"的文档和方法概述。 |
| go doc json.Number.Int64 (or go doc json.number.int64) | 显示"json.Number"的"Int64"方法的文档。 |
| go doc cmd/doc | 显示"doc"命令的包文档。 |
| go doc -cmd cmd/doc | 显示"doc"命令的包文档和导出标识符。 |
| go doc template.new | 显示"html/template"包里的"New"函数的文档。（词法上，"html/template"包在"text/template"包之前。） |
| go doc text/template.new # One argument | 显示"text/template"包里的"New"函数的文档。 |
| go doc text/template new # Two arguments | 显示"text/template"包里的"New"函数的文档。 |
|--|--|
| 下面的命令调用都会打印"json.Decoder"的"Decode"方法的文档: |--|
| go doc json.Decoder.Decode |--|
| go doc json.decoder.decode |--|
| go doc json.decode |--|
| cd go/src/encoding/json; go doc decode |--|


##### 三、doc命令常用标志说明

| 标志名 | 说明 |
|--|--|
| -all | 显示包的所有文档。 |
| -c | 匹配标识符时区分大小写。 |
| -cmd | 像处理常规包一样处理命令包（package main）。否则，在显示包的顶层文档时，main包的导出标识符会被隐藏。 |
| -src | 显示标识符的完整源代码。这将会显示显示完整的Go源码的声明和定义，例如一个函数的定义（包括函数主体），类型定义，或者内置的常量代码段。因此，输出可能包含非导出细节。 |
| -u | 显示非导出和导出的标识符、方法和字段的文档。 |

https://github.com/sfshf/Notes_IT_Basic

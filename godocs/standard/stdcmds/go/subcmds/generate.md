
[Generate Go files by processing source](https://golang.google.cn/cmd/go/#hdr-Generate_Go_files_by_processing_source)

##### 一、用法

```

go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

```

##### 二、generate命令说明

`generate命令`会按照现有文件中指示出的命令逐一运行。这些命令可以是任何进程，但其目的是创建或更新Go源文件。

`go generate`从来不会被`go build`、`go get`、`go test`等命令导致自动运行。它必须显式地运行。

`go generate`扫描文件得到命令指示，命令指示均为成行的如下格式：

```

//go:generate command argument...

```

（注意：在"//go"中没有前导空格且"//go"中间也没有空格）`command`是要被运行的生成器，对应一个可以在本地运行的可执行文件。它必须是shell路径(gofmt)、完全限定路径（/usr/you/bin/mytool）或命令别名，如下所述。

为了将生成的代码传达给人类和机器工具，生成的源代码应该有一行匹配以下正则表达式(使用Go语法):

```

^// Code generated .* DO NOT EDIT\.$

```

如上一行可以出现在文件的任何地方，但是它典型地方式在开头附近，因为这样容易找到。

注意：`go generate`不会解析文件，所以在注释或者多行字符串中的看起来像命令指示的文本行会被当做命令指示对待。

用于命令指示的参数（argument）是空格隔开的标志符号（token）或者用于在生成器运行时传递给生成器的独特参数是用双引号括起来的字符串。

带引号的字符串要使用Go语法，并且会在执行之前进行求值;一个带引号字符串被视为生成器的单个参数。

`go generate`在它运行生成器时有如下表所示的几个变量可供作`环境变量替换`使用:

| 环境变量 | 简单说明 |
|--|--|
| $GOARCH | 执行机的系统架构。（arm、amd64等） |
| $GOOS | 执行机的操作系统。（linux、windows等） |
| $GOFILE | 文件的基名。（即：不含扩展名的文件名） |
| $GOLINE | 在源文件中的命令指示的行数。 |
| $GOPACKAGE | 命令指示文件所在的包名。 |
| $DOLLAR | 一个`$`（dollar）符号。 |

除了变量替换和引用字符串求值之外，命令行上没有提供诸如`globbing`之类的特殊处理。

作为运行命令之前的最后一步，在整个命令行中，对任何调用的经过`环境变量替换`的环境变量(如$GOFILE或$HOME)都将被展开。变量展开的语法在所有操作系统上都是`$NAME`。由于求值的顺序，变量甚至在带引号的字符串中也会展开。如果环境变量`NAME`没有设置，`$NAME`将展开为空字符串。

命令指示还有如下所示的格式：

```

//go:generate -command xxx args...

```
该格式仅作用于当前命令指示源文件的剩余部分，其中字符串`xxx`表示由参数标识的命令。该格式可以用来创建别名或处理多单词的生成器。比如：

```

//go:generate -command foo go tool foo

```
该指示命令指明了用`foo`命令代表`go tool foo`生成器。

`generate命令`会按照命令行给出的顺序依次处理包。如果命令行中列出了来自同目录下的`.go文件`，则将这些`.go文件`视为同一个包。在一个包里，`generate命令`按文件名顺序依次处理源文件。在一个源文件里，`generate命令`按文件中出现的顺序依次运行生成器。`go generate`还设置了`build标志`--"generate"，这样文件可以被`go generate`检查，而在`build`过程中会被忽略。

如果任何生成器返回错误退出状态，`go generate`将跳过之后的对该包的所有处理。

生成器会在包的源目录中运行。

`generate命令`有一个独立的标志：

| 标志名 | 说明 |
|--|--|
| -run="" | 如果非空，指定一个正则表达式来选择指令，该指令的完整的原始的源文本（该源文本已去除了指示命令后的任何尾随空格及最后的换行）与正则表达式匹配。 |

`generate命令`还接收标准的`build标志`，包括`-v`、`-n`和`-x`。`-v`标志会打印被处理的包名和文件名。`-n`标志打印将要执行的命令（不运行）。`-x`标志会执行命令并打印出命令。

使用`go help build`查看更多关于`build标志`。

使用`go help packages`看看更多关于包的说明。

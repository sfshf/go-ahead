
[Module maintenance](https://golang.google.cn/cmd/go/#hdr-Module_maintenance)

`go mod`命令提供对模块操作的访问。

注意：对模块的支持内置在所有的`go`命令中，而不仅仅是`go mod`。例如，日常对依赖项的添加、删除、升级和降级应该使用`go get`来完成。有关模块功能的概述，查看`go help modules`。

##### 一、`go mod`命令的用法

```

go mod <command> [arguments]

```

##### 二、`go mod`命令的说明

`<command>`支持的命令有下表所列：

| 命令名 | 简单说明 |
|--|--|
| download | 下载模块到本地缓存。 |
| edit | 编辑来自工具或脚本的`go.mod`文件。 |
| graph | 打印模块引用图。 |
| init | 在当前目录下初始化新模块。 |
| tidy | 添加遗漏的和删除的没有使用的模块。 |
| vendor | 创建用于供应的（vendored）依赖项副本。 |
| verify | 验证依赖项是否具有预期的内容。 |
| why | 解释为什么这些包或模块被需要。 |

使用`go help mod <command>`查看关于某个命令更多信息。


[Download modules to local cache](https://golang.google.cn/cmd/go/#hdr-Download_modules_to_local_cache)

##### 一、`go mod download`命令的用法

```

go mod download [-json] [modules]

```

##### 二、`go mod download`命令的说明

`download`命令下载指名的模块，或用于选择主模块的依赖项的模块模式，还可以是格式为`path@version`的模块查询。在没有参数的情况下，`download`命令应用于主模块的所有依赖项。

`go`命令将在普通执行过程中根据需要自动下载模块。`go mod download`命令主要用于预先填充本地缓存或推断一个`go模块代理`（Go module proxy）的回应。

默认情况下，`download`命令将错误报告到标准错误，否则就是静默的。`-json`标志导致`download`命令将一系列JSON对象打印到标准输出，描述每个成功下载的模块(或失败的)，根据下面这个Go结构体:

```go

type Module struct {
    Path     string // module path
    Version  string // module version
    Error    string // error loading module
    Info     string // absolute path to cached .info file
    GoMod    string // absolute path to cached .mod file
    Zip      string // absolute path to cached .zip file
    Dir      string // absolute path to cached source root directory
    Sum      string // checksum for path, version (as in go.sum)
    GoModSum string // checksum for go.mod (as in go.sum)
}

```

使用`go help modules`查看更多关于模块查询的信息。


[Edit go.mod from tools or scripts](https://golang.google.cn/cmd/go/#hdr-Edit_go_mod_from_tools_or_scripts)

##### 一、`go mod edit`命令的用法

```

go mod edit [editing flags] [go.mod]

```

##### 二、`go mod edit`命令的说明

`edit`命令为编辑`go.mod`文件提供了一个命令行接口，主要用于工具或脚本。它只读取`go.mod`文件；它不会查找涉及到的模块的信息。默认情况下，`edit`命令读写主模块的`go.mod`文件，但一个不同的目标文件可以被指定在编辑标志（editing flag）之后。

编辑标志（editing flag）指定了编辑操作的一个序列。

`-fmt`标志将重新格式化`go.mod`文件，不做其他更改。使用或重写`go.mod`文件的任何其他修改也暗示了这种重新格式化。只有在没有指定其他标志的情况下才需要`-fmt`标志，比如在`go mod edit -fmt`中。

`-module`标志用于改变模块的路径(`go.mod`文件的模块行)。

`-require=path@version`和`-droprequire=path`标志在给定的模块路径和版本上添加和删除需求。注意，`-require`标志覆盖`path`上的所有现有需求。这些标志主要用于理解模块图的工具。用户应该优先使用`go get path@version`或`go get path@none`，因为它们会让其他`go.mod`文件的调整按需满足由其他模块强加的约束。

`-exclude=path@version`和`-dropexclude=path@version`标志为给定的模块路径和版本添加和删除一个排除项。请注意，`-exclude=path@version`是一个空操作（no-op），如果这个排除项已经存在的话。

`-replace=old[@v]=new[@v]`和`-dropreplace=old[@v]`标志用于添加和删除给定模块路径和版本对的替换。如果省略了`old[@v]`中的`@v`，则替换将适用于具有旧模块路径的所有版本。如果省略`new[@v]`中的`@v`，则新路径应该是本地模块根目录，而不是模块路径。注意，`-replace`会覆盖所有现有的对`old[@v]`的替换。

可以重复使用`-require`、`-droprequire`、`-exclude`、`-dropexclude`、`-replace`和`-dropreplace`这些编辑标志，并按给定的顺序进行更改。

`-go=version`标志用于设置想要的Go语言版本。

`-print`标志会以其自己的文本格式打印最终的`go.mod`内容，而不是将内容写回到`go.mod`文件。

`-json`标志会以JSON格式打印最终的`go.mod`内容，而不是将内容写回到`go.mod`文件。JSON输出是对照以下几个Go类型：

```go

type Module struct {
	Path string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path string
	Version string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

```

注意，`GoMod`结构体只描述了`go.mod`文件本身，而不是其他被简介引用的模块。查看对于一次构建可获得的所有模块的完整集合，可以使用`go list -m -json all`命令。

举例，一个工具可以通过解析`go mod edit -json`命令的输出，获得`go.mod`文件作为一个数据结构，还可以通过调用使用了`-require`、`-exclude`等标志的`go mod edit`命令来进行更改。


[Print module requirement graph](https://golang.google.cn/cmd/go/#hdr-Print_module_requirement_graph)

##### 一、`go mod graph`命令的用法

```

go mod graph

```

##### 二、`go mod graph`命令的说明

`graph`命令以文本形式打印模块需求图(包含应用的替换)。输出中的每一行都有两个字段（以空格分隔的）：模块及其需求之一。每个模块被标识为格式为`path@version`的字符串，主模块除外，因为主模块没有`@version`后缀。


[Initialize new module in current directory](https://golang.google.cn/cmd/go/#hdr-Initialize_new_module_in_current_directory)

##### 一、`go mod init`命令的用法

```

go mod init [module]

```

##### 二、`go mod init`命令的说明

`init`命令会初始化并编写一个新的`go.mod`文件到当前目录，实际上是在当前目录下创建一个新的模块；当前目录必须之前是不存在`go.mod`文件的。如果可能，`init命令`将根据导入注释（import comments，参见`go help importpath`）或版本控制配置猜测模块路径。要覆盖此猜测，请提供模块路径作为参数。


[Add missing and remove unused modules](https://golang.google.cn/cmd/go/#hdr-Add_missing_and_remove_unused_modules)

##### 一、`go mod tidy`命令的用法

```

go mod tidy [-v]

```

##### 二、`go mod tidy`命令的说明

`tidy`命令是用于保证`go.mod`文件匹配模块中的源代码。它添加用于构建当前模块的包和依赖项所需的任何缺少的模块，并删除不提供任何相关包的未使用的模块。它还会添加任何丢失的条目到`go.sum`文件，并把所有不必要的条目删除。

`-v`标志导致`tidy`命令打印被删除模块的信息到标准错误里。


[Make vendored copy of dependencies](https://golang.google.cn/cmd/go/#hdr-Make_vendored_copy_of_dependencies)

##### 一、`go mod vendor`命令的用法

```

go mod vendor [-v]

```

##### 二、`go mod vendor`命令的说明

`vendor`命令会重置主模块的`vendor目录`，来包含用于构建和测试所有主模块包所需的所有包。它不包括被供应包（vendored packages）的测试代码。

`-v`标志导致`vendor`命令将被供应模块和包的名字打印到标准错误里。


[Verify dependencies have expected content](https://golang.google.cn/cmd/go/#hdr-Verify_dependencies_have_expected_content)

##### 一、`go mod verify`命令的用法

```

go mod verify

```

##### 二、`go mod verify`命令的说明

`verify`命令检查当前模块的依赖项(存储在本地下载的源缓存中)是否在下载后未被修改。如果所有模块都没有修改，`verify`命令会打印`all modules verified.`。否则它会报告哪些模块被更改了，并导致`go mod`命令以非零状态退出。


[Explain why packages or modules are needed](https://golang.google.cn/cmd/go/#hdr-Explain_why_packages_or_modules_are_needed)

##### 一、`go mod why`命令的用法

```

go mod why [-m] [-venvor] packages

```

##### 二、`go mod why`命令的说明

`why`命令显示导入图中从主模块到列出的每个包的最短路径。如果给出了`-m`标志，那么`why`命令要将参数视为模块列表，并找出每个模块中的任何包的路径。

默认情况下，`why`命令会查询与`go list all`命令匹配的包的图，该图包含了可达包的测试。`-vendor`标志会导致`why`命令排除依赖项的测试。

输出是一个文段（stanza）序列，每个文段对应命令行上的每个包或模块名，文段之间用空行分隔。每个文段以注释行`# package`或`# module`开始，表示目标包或模块。随后的几行给出了导入图的路径，每行一个包。如果包或模块没有从主模块中引用，则文段将显示一个使用圆括号括起来提示来说明这一事实。

举例：

```

$ go mod why golang.org/x/text/language golang.org/x/text/encoding
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
$

```

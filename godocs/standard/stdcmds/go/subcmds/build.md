
[Compile packages and dependencies](https://golang.google.cn/cmd/go/#hdr-Compile_packages_and_dependencies)


##### 一、用法

```

go build [-o output] [-i] [build flags] [packages]

```


##### 二、build命令说明

 `build命令`编译由导入路径命名的包及其依赖项，但不安装结果。

如果 `build命令`的参数是来自单个目录下的一列`.go`文件，则 `build命令`会将它们视为指定为单个包的一列源文件。

编译包时， `build命令`将忽略以`_test.go`结尾的测试文件。

编译单个main包时， `build命令`会将生成的可执行文件写入以第一个源文件（例如，`go build ed.go rx.go`将输出`ed`文件或`ed.exe`文件）或源代码目录（`go build unix/sam`将输出`sam`文件或`sam.exe`文件）命名的输出文件内。写入`Windows`可执行文件时会添加`.exe`后缀。

编译多个包或单个非main包时， `build命令`编译包，但丢弃结果对象，仅用作检查是否可以生成包。

`-o`标志，强制将生成的可执行文件或对象写入指明的输出文件或目录中，而不是前两段中描述的默认行为。如果指明的输出是一个存在的目录，那么任何生成的可执行文件将被写入该目录。

`-i`标志，用于安装依赖包。


##### 三、build命令的flag说明

 `build命令`的标志（flag）是`build命令`、`clean命令`、`get命令`、`install命令`、`list命令`、`run命令`及`test命令`共用的。

| 标志名 | 说明 |
|--|--|
| -a | 强制重新编译已准备更新的包。 |
| -n | 打印命令但不运行它们。 |
| -p n | 指明程序的数量。比如，可以让build命令或test命令并行运行。该标志的默认值为可用的CPU数量。 |
| -race | 启用[数据竞争检测](https://link.springer.com/referenceworkentry/10.1007%2F978-0-387-09766-4_2248)。仅支持linux/amd64、freebsd/amd64、darwin/amd64和windows/amd64系统。 |
| -msan | 启用与内存清理程序的互操作。仅支持linux/amd64、linux/arm64和使用Clang/LLVM作为宿主C编译器。 |
| -v | 打印被编译的包名。 |
| -work | 打印临时工作目录，并在程序运行结束时不删除临时工作目录。 |
| -x | 打印命令。与`-n`标志对照。 |
| -asmflags '[pattern=]arg list' | 用于执行`go tool asm`命令时传参。 |
| -buildmode mode | 指明build命令所要用的模式。使用`go help buildmode`命令查询更多信息。 |
| -compiler name | 指明所要用的编译器的名字。例如`runtime.Compiler`（gccgo或者gc）。 |
| -gccgoflags '[pattern=]arg list' | 用于执行`gccgo compiler/linker`命令时传参。 |
| -gcflags '[pattern=]arg list' | 用于执行`go tool compile`命令时传参。 |
| -installsuffix suffix| 为了让输出与默认构建区分开，要在包安装目录的名称中使用的后缀。在使用`-race`标志时，安装的后缀会将自动设置为`race`；如果`-installsuffix`与`-race`标志同时启用，则在指明的后缀的后面追加`_race`；同样情况适用于`-msan`标志；使用需要非默认编译标志的`-buildmode`选项也有相似的效果。 |
| -ldflags '[pattern=]arg list' | 用于执行`go tool link`命令时传参。 |
| -linkshared | 链接到之前使用`-buildmode=shared`创建的共享库。 |
| -mod mode | 要用到的模块下载模式：`readonly`或`vendor`。使用`go help modules`命令查看更多信息。 |
| -pkgdir dir | 从`dir`而不是通常的位置安装和加载所有包。例如，当使用非标准配置构建时，使用`-pkgdir`将生成的包保存在单独的位置。 |
| -tags tag,list | 用于指明在构建期间认为用得到的，使用逗号分隔的构建标签列表。有关构建标签的详细信息，请参见在`go/build`标准包的文档中构建约束条款。（早期版本的Go使用了空格分隔的列表，该方式已被弃用，但仍被认可。） |
| -trimpath | 从生成的可执行文件中删除所有文件系统路径。替代了绝对的文件系统路径，被记录的文件名将要么以`go`（对于标准库）要么以一个模块`path@version`（使用模块时），再要么以朴素的`import path`（使用GOPATH时）开头。 |
| -toolexec 'cmd args' | 用于调用工具链程序（如`vet`和`asm`）的程序。例如，`go`命令将运行`cmd args /path/to/asm <arguments for asm>`，而不是运行`asm`。 |

`-asmflags`、`-gccgoflags`、`-gcflags`和`-ldflags`标志,在生成期间接受一个以空格分隔的参数列表，传递到一个底层工具。要在列表里的一个元素中嵌入空格，请用单引号或双引号将其括起来。参数列表前面可以有一个包模式（package pattern）和一个等号（equal sign），这将参数列表的使用限制为生成与该模式匹配的包（使用`go help packages`查看更多有关包模式的说明）。如果没有模式，参数列表只应用于命令行中指名的包。这些标志可以用不同的模式重复，以便为不同的包集指定不同的参数。如果包与多个标志中给定的模式匹配，则只有命令行上的最新的匹配将有效。例如，`go build -gcflags=-S fmt`只为包`fmt`打印反汇编，而`go build -gcflags=all=-S fmt`则为`fmt`及其所有依赖项打印反汇编。

使用`go help packages`查看更多有关包的详细说明。使用`go help gopath`查看更多有关包和二进制文件的安装位置的详细信息。使用`go help c`查看更多有关在Go语言和C/C++语言之间调用的更多信息。

`注意：build命令`遵循某些约定，如`go help gopath`可查信息中所描述的约定。然而，并不是所有的项目都能遵循这些约定。有自己的约定或使用独立的软件生成系统的装置可以选择使用较底层的调用，如`go tool compile`和`go tool link`，以避免构建工具的一些高开销和生成工具的设计决策。

更多：[`go install`](install.md)、 [`go get`](get.md)、 [`go clean`](clean.md)。


[Environment variables](https://golang.google.cn/cmd/go/#hdr-Environment_variables)


##### 说明：

`go`命令和它调用的工具会查阅环境变量来获取配置信息。如果环境变量未设置，`go`命令使用合理的默认设置。要查看变量`<NAME>`的有效设置，请运行`go env <NAME>`。要更改默认设置，请运行`go env -w <NAME>=<VALUE>`。使用`go env -w`更改的默认值被记录在每个用户配置目录里存储的go环境配置文件中，该文件路径可以使用`os.UserConfigDir`获取。通过设置环境变量`GOENV`可以更改配置文件的位置；`go env GOENV`打印有效位置，而使用`go env -w`不能更改默认位置。详情请参见`go help env`。

通用的环境变量有：

| 环境变量名 | 说明 |
|--|--|
| GCCGO | 用于运行`go build -compiler=gccgo`的`gccgo`命令。 |
| GOARCH | 用来编译代码的体系结构或处理器，例如：`amd64`、`386`、`arm`或`ppc64`。 |
| GOBIN | `go install`安装命令的目录。 |
| GOCACHE | go命令存放缓存信息的目录。缓存信息是为了在将来的构建中进行复用的。 |
| GODEBUG | 启用不同的调试工具。详见`go doc runtime`。 |
| GOENV | go环境配置文件的定位。不能使用`go env -w`设置。 |
| GOFLAGS | 一个空格分隔的`-flag=value`设置列表，默认应用于go命令，如果被给定的标志能被该go命令识别的话。每个条目必须是一个独立的标志。因为条目是用空格分隔的，所以标记值不能包含空格。命令行中列出的标志将应用于此列表之后，并因此覆盖它。 |
| GOOS | 用于编译代码的操作系统，例如：`linux`、`darwin`、`windows`或`netbsd`。 |
| GOPATH | 详见`go help gopath`。 |
| GOPROXY | Go模块代理的URL。更多查看`go help modules`。 |
| GOPRIVATE, GONOPROXY, GONOSUMDB | 模块路径前缀的通配符模式（glob pattern，可以查看Go的`path.Match`语法）的用逗号分隔的列表，模块路径前缀应该总能被直接获取，或者不应该与校验和数据库（checksum database）进行比较。参见`go help module-private`。 |
| GOROOT | go的sdk包树的根路径。 |
| GOSUMBD | 要使用的校验和数据库（checksum database）的名称，可能还有检验和数据库的公钥和URL。参见`go help module-auth`。 |
| GOTMPDIR | go命令将写入临时的源文件、包和二进制文件的目录。 |

用于`cgo`的环境变量有：

| 环境变量名 | 说明 |
|--|--|
| AR | 使用`gccgo`编译器构建时用于操作库存档的命令。默认是`ar`。 |
| CC | 用于编译C语言代码的命令。 |
| CGO_ENABLED | 标志是否支持`cgo`命令。值要么是`0`，要么是`1`。 |
| CGO_CFLAGS | 当编译C语言代码时`cgo`将传入编译器的标志。 |
| CGO_CFLAGS_ALLOW | 指明额外标志的正则表达式，额外标志是允许出现在`#cgo CFLAGS`源代码指令中的标志。不适用`CGO_CFLAGS`环境变量。 |
| CGO_CFLAGS_DISALLOW | 指明禁止出现在`#cgo CFLAGS`源代码指令中的标志的正则表达式。不适用`CGO_CFLAGS`环境变量。 |
| CGO_CPPFLAGS, CGO_CPPFLAGS_ALLOW, CGO_CPPFLAGS_DISALLOW | 类似`CGO_CFLAGS`、`CGO_CFLAGS_ALLOW`和`CGO_CFLAGS_DISALLOW`，但是应用于`C++预处理器`。 |
| CGO_CXXFLAGS, CGO_CXXFLAGS_ALLOW, CGO_CXXFLAGS_DISALLOW | 类似`CGO_CFLAGS`、`CGO_CFLAGS_ALLOW`和`CGO_CFLAGS_DISALLOW`，但是应用于`C++编译器`。 |
| CGO_FFLAGS, CGO_FFLAGS_ALLOW, CGO_FFLAGS_DISALLOW | 类似`CGO_CFLAGS`、`CGO_CFLAGS_ALLOW`和`CGO_CFLAGS_DISALLOW`，但是应用于`Fortran编译器`。 |
| CGO_LDFLAGS, CGO_LDFLAGS_ALLOW, CGO_LDFLAGS_DISALLOW | 类似`CGO_CFLAGS`、`CGO_CFLAGS_ALLOW`和`CGO_CFLAGS_DISALLOW`，但是应用于`链接器`。 |
| CXX | 用于编译`C++代码`的命令。 |
| FC | 用于编译`Fortran代码`的命令。 |
| PKG_CONFIG | `pkg-config`工具的路径。 |

特定体系结构的环境变量有：

| 环境变量名 | 说明 |
|--|--|
| GOARM | 代表`GOARCH=arm`，为`ARM`体系结构进行编译。有效值有`5`、`6`、`7`。 |
| GO386 | 代表`GOARCH=386`，浮点指令集。有效值有`387`、`sse2`。 |
| GOMIPS | 代表`GOARCH=mips{,le}`，是否使用浮点指令。有效值有`hardfloat`（默认，硬件浮点运算）、`softfloat`。 |
| GOMIPS64 | 代表`GOARCH=mips64{,le}`，是否使用浮点指令。有效值有`hardfloat`（默认，硬件浮点运算）、`softfloat`。 |
| GOWASM | 代表`GOARCH=wasm`，使用逗号分隔的实验性`WebAssembly`特性列表。有效值有`satconv`, `signext`。 |

特殊目的的环境变量有：

| 环境变量名 | 说明 |
|--|--|
| GCCGOTOOLDIR | 如果设置，指示在哪里可以找到`gccgo`工具，比如`cgo`。默认设置基于`gccgo`的配置。 |
| GOROOT_FINAL | 指示被安装的go的sdk包树的根路径，go的sdk包是被安装在一个地方，而不是被构建。堆栈跟踪中的文件名被从`GOROOT`重写到`GOROOT_FINAL`。 |
| GO_EXTLINK_ENABLED | 指示当使用存在`cgo`代码的`-linkmode=auto`时，链接器是否应该使用外部链接模式。设置为`0`可以禁用外部链接模式，设置为`1`可以启用外部链接模式。 |
| GIT_ALLOW_PROTOCOL | 被`Git`定义。允许与`git fetch/clone`一起使用的冒号分隔的模式列表。如果设置，任何没有明确提到的模式将被`go get`认为是不安全的。因为变量是由`Git`定义的，所以不能使用`go env -w`来设置默认值。 |

可从`go env`获取但无法从环境中读取的其他信息:

| 环境变量名 | 说明 |
|--|--|
| GOEXE | 可执行文件的名字后缀（即：扩展名）。（在`Windows`系统上是`.exe`，在其他系统上是`""`） |
| GOGCCFLAGS | 应用于`CC`命令的一个用空格隔开的参数列表。 |
| GOHOSTARCH | go工具链二进制文件的结构体系（即：`GOARCH`）。 |
| GOHOSTOS | go工具链二进制文件的操作系统（即：`GOOS`）。 |
| GOMOD | 主模块的`go.mod`文件的绝对路径，或者如果没有使用模块，则为空字符串。 |
| GOTOOLDIR | go命令工具（`compile`、`cover`、`doc`等等）被安装的目录。 |

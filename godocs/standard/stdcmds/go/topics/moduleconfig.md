
[Module configuration for non-public modules](https://golang.google.cn/cmd/go/#hdr-Module_configuration_for_non_public_modules)


##### 说明：

`go命令`默认从公共go模块镜像`proxy.golang.org`下载模块。它还默认针对`sum.golang.org`上的`公共Go校验和数据库`验证下载的模块（无论其来源）。这些默认设置适用于公开可用的源代码。

`GOPRIVATE`环境变量控制`go命令`认为哪些模块是私有的(不可公开使用)，因此不应该使用代理或校验和数据库。该变量是模块路径前缀的`glob`模式的逗号分隔列表（在Go的`path.Match`语法中）。例如：`GOPRIVATE=*.corp.example.com,rsc.io/private`，该值会导致`go命令`将任何路径前缀与`GOPRIVATE`变量中任一模式匹配的模块当作私有的，包括`git.corp.example.com/xyzzy`、`rsc.io/private`和`rsc.io/private/quux`。

`GOPRIVATE`环境变量也可以被其他工具用于标识非公共模块。例如，编辑器可以使用`GOPRIVATE`来决定是否将包导入链接到`godoc.org`页面。

为了对模块下载和验证进行细粒度（fine-grained）控制，`GONOPROXY`和`GONOSUMDB`环境变量接受相同类型的`glob`列表，并重写`GOPRIVATE`以分别决定是否使用`代理`及`校验和数据库`。

举例，如果一个公司管理一个模块代理提供私有模块的服务，用户可以如下配置`go环境`：

```

GOPRIVATE=*.corp.example.com
GOPROXY=proxy.example.com
GONOPROXY=none

```

如上配置，将告诉`go命令`和其他工具，以`corp.example.com`子域开头的模块是私有的；但是公司代理应该用于下载公共和私有模块，因为`GONOPROXY`被设置为与任何模块都不匹配的模式；以上配置重写了`GOPRIVATE`。

`go env -w`命令（请参阅`go help env`）可用于为以后的`go命令`调用设置这些变量。

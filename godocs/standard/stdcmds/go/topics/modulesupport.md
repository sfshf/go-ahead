
[Module support](https://golang.google.cn/cmd/go/#hdr-Module_support)


##### 说明：

`Go 1.13`包含对`Go模块`的支持。`模式感知模式（module-aware mode）`在默认情况下是激活的，无论`go.mod`文件是否可以在当前目录中被找到，或在当前目录的父目录中被找到。

利用模块支持的最快方法是签出存储库，创建一个`go.mod`文件(在下一节中描述)，并在该文件树中运行`go命令`。

对于更细粒度的（fine-grained）控制，`Go 1.13`继续使用一个临时环境变量`GO111MODULE`，该变量可以设置为三个字符串值之一：`off`、`on`或`auto`(默认值)。如果`GO111MODULE=on`，那么`go命令`需要使用module，而不是查询`GOPATH`。我们将此称为`模块感知的命令`或以`模块感知模式`运行的命令。如果`GO111MODULE=off`，则`go命令`从不使用模块支持。取而代之地是，它会在`vendor`目录和`GOPATH`中查找依赖项;我们现在称之为`GOPATH模式（GOPATH mode）`。如果`GO111MODULE=auto`或未设置，则`go命令`根据当前目录启用或禁用模块支持。当前目录包含`go.mod`文件，或者当前目录位于包含`go.mod`文件的目录下，才启用模块支持。

在`模块感知模式`下，`GOPATH`不再在构建期间定义导入的意义，但它仍然存储下载的依赖项(在`GOPATH/pkg/mod`中)和安装的命令(没有设置`GOBIN`的情况下在`GOPATH/bin`里)。

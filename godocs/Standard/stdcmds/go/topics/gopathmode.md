
[Legacy GOPATH go get](https://golang.google.cn/cmd/go/#hdr-Legacy_GOPATH_go_get)


##### 说明：

`go get`命令根据`go命令`是在`模块感知模式（module-aware mode）`下运行还是在`遗留GOPATH模式（legacy GOPATH mode）`下运行来更改行为。帮助文本（在`模块感知模式`下以`go help gopath-get`的形式访问）描述了`go get`在`遗留GOPATH模式`下的操作。

用法：`go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]`

`get命令`会下载以导入路径的方式指名的包，以及它们的依赖项；然后`get命令`会安装被指名的包，类似`go install`。

`-d`标志会指示`get命令`在下载包之后停止运行；也就是说，`-d`标志会指示`get命令`不去安装下载好的包。

`-f`标志，仅当`-u`标志被设置时有效，强制`get -u`不去验证每个包是不是从以导入路径方式指示的源代码控制存储库中检出的。如果源文件是原始文件的本地分支（local fork），那么`-f`标志将非常有用。

`-fix`标志指示`get命令`在解析依赖项或构建代码之前，对下载的包运行`fix`工具。

`-insecure`标志允许从存储库获取数据并使用`HTTP等不安全模式`解析自定义域。谨慎使用。

`-t`标志指示`get命令`还可以下载为指定包构建测试所需的包。

`-u`标志指示`get命令`使用网络更新被指名的包及其依赖项。默认情况下，`get命令`会使用网络来检出丢失的包，但不会查找现有包的更新。

`-v`标志启用进程和调试的详细输出。

`get命令`还接受构建标志来控制安装。参见`go help build`。

在检出一个新包时，`get命令`创建目标目录`GOPATH/src/<import-path>`。如果`GOPATH`包含多个条目，`get命令`将使用第一个条目。更多细节请看：`go help gopath`。

在检出或更新一个包时，`get命令`会查找与本地安装的Go语言版本匹配的分支（branch）或标记（tag）。最重要的规则是，如果本地安装正在运行版本`go1`，则`get命令`会搜索名为`go1`的分支或标记。如果不存在这样的版本，`get命令`将检索包的默认分支。

当`go get`检出或更新`Git`存储库时，它还更新存储库引用的任何`Git`子模块。

`get命令`绝不会检出和更新存放在`vendor`目录下的代码。

使用`go help packages`查看更多关于包的详细信息。

使用`go help importpath`查看更多关于`go get`如何寻找需要下载的源代码的。

本文描述了，使用`GOPATH`管理源代码和依赖项时`get命令`的行为。如果`go命令`以`模块感知模式（module-aware mode）`运行，`get命令`的标志和效果的细节就会改变，`go help get`也是如此。参见`go help modules`和`go help module-get`。

另见：[`go build`](../subcmds/build.md)、[`go install`](../subcmds/install.md)、[`go clean`](../subcmds/clean.md)。

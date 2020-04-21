
[Compile and install packages and dependencies](https://golang.google.cn/cmd/go/#hdr-Compile_and_install_packages_and_dependencies)

##### 一、用法

```

go install [-i] [build flags] [packages]

```

##### 二、install命令说明

`install命令`会编译并安装使用导入路径命名的包。

可执行文件会被安装在`GOBIN`环境变量命名的目录中，如果没有设置`GOPATH`环境变量，则默认为`$GOPATH/bin`或`$HOME/go/bin`。`$GOROOT`中的可执行文件安装在`$GOROOT/bin`或`$GOTOOLDIR`中，而不是`$GOBIN`。

当禁用模块感知模式（`module-aware mode`）时，将在目录`$GOPATH/pkg/$GOOS_$GOARCH`中安装其他包。当启用模块感知模式（`module-aware mode`）时，将构建和缓存其他包，但不安装它们。

`-i`标志也安装指名的包的依赖项。

使用`go help build`查看更多关于构建的标志。使用`go help packages`查看更多关于包的详细说明。

更多信息参考[`go build`](build.md)，[`go get`](get.md)，[`go clean`](clean.md)。

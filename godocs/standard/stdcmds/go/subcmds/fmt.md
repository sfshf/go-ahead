
[Gofmt (reformat) package sources](https://golang.google.cn/cmd/go/#hdr-Gofmt__reformat__package_sources)

##### 一、用法

```

go fmt [-n] [-x] [packages]

```

##### 二、fmt命令说明

`fmt命令`对以导入路径形式命名的包运行`gofmt -l -w`命令。它会打印出被修改的文件的名字。

使用`go doc cmd/gofmt`查看更多关于`gofmt`命令的信息。使用`go help packages`查看更多关于包的详细说明。

`-n`标志会打印可能被执行的命令（不运行）。`-x`标志打印被执行的命令（运行）。

想要使用特殊的选项运行`gofmt`命令，可以直接使用`gofmt`命令。

更多信息参考[`go fix`](fix.md)，[`go vet`](vet.md)。

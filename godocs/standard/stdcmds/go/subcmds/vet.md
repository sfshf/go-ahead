
[Report likely mistakes in packages](https://golang.google.cn/cmd/go/#hdr-Report_likely_mistakes_in_packages)

##### 一、`go vet`命令的用法

```

go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]

```

##### 二、`go vet`命令的说明

`go vet`用于检查由导入路径指名的包。

使用`go doc cmd/vet`查看关于`vet`命令和它的标志的更多信息。使用`go help packages`查看更多关于包的详细说明。使用`go tool vet help`查看检查器和它们的标志的列表；还可以查看一个指明的检查器的详细信息，例如：使用`go tool vet help printf`查看`printf`。

`-n`标志打印可能被执行的命令（不运行）。`-x`标志会打印被执行的命令（运行）。

`-vettool=prog`标志用于选择一个不同的分析工具，并进行备选或附加检查。例如：使用下面命令可以使用`shadow`分析器进行构建并运行：

```

go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
go vet -vettool=$(which shadow)

```

`go vet`支持的构建标志都是用于控制包的分析和执行的标志，例如`-n`、`-x`、`-v`、`-tags`和`-toolexec`。使用`go help build`查看更多关于这些标志的信息。

更多查看：[`go fmt`](fmt.md)、[`go fix`](fix.md)。


[Compile and run Go program](https://golang.google.cn/cmd/go/#hdr-Compile_and_run_Go_program)

##### 一、`go run`命令的用法

```

go run [build flags] [-exec xprog] package [arguments...]

```

##### 二、`go run`命令的说明

`run`编译并运行命名为`main`的Go包。通常，`package`是被指定为来自同一个目录的`.go`源文件列表，但它也可能是一个导入路径、文件系统路径或与单个已知包匹配的模式，如`go run .`或`go run my/cmd`。

默认情况下，`go run`直接运行编译后的二进制文件：`a.out arguments...`。如果给定了`-exec`标志，`go run`使用`xprog`调用二进制文件：

```

xprog a.out arguments...

```

如果未给定`-exec`标志，`GOOS`或`GOARCH`与系统默认值不同，并且当前搜索路径上没有找到名为`go_$GOOS_$GOARCH_exec`的程序，那么`go run`使用该程序调用二进制文件，例如`go-nacl-386-exec a.out arguments...`。这允许在模拟器或其他执行方法可用时执行交叉编译程序。

`run`命令的退出状态不是被编译出的二进制文件的退出状态。

有关`build`标志的详细信息，请使用`go help build`查看。有关指定包的详细信息，请使用`go help packages`查看。

另请参见：[`go build`](build.md)。

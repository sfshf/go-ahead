
[Run specified go tool](https://golang.google.cn/cmd/go/#hdr-Run_specified_go_tool)

##### 一、`go tool`命令的用法

```

go tool [-n] command [args...]

```

##### 二、`go tool`命令的说明

`go tool`用于运行被参数标识的go工具命令。没有参数的情况下，会打印已知工具命令的列表。

`-n`标志会导致`go tool`打印出可能被执行的命令，但不执行。

使用`go doc cmd/<command>`查看更多关于每个工具命令的信息。

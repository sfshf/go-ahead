
[Print Go version](https://golang.google.cn/cmd/go/#hdr-Print_Go_version)

##### 一、`go version`命令的用法

```

go version [-m] [-v] [file ...]

```

##### 二、`go version`命令的说明

`go version`用于打印Go可执行文件的构建信息。

`go version`会报告用于构建每个指名的可执行文件使用的Go版本。

如果命令行中没有指明文件，`go version`会打印本地安装的Go版本。

如果指定了一个目录，则`go version`递归地遍历该目录，查找可识别的Go二进制文件并报告它们的版本。默认情况下，`go version`不会报告在目录扫描过程中发现的无法识别的文件。`-v`标志使得`go version`报告无法识别的文件。

`-m`标志使`go version`打印每个可执行文件的嵌入式模块版本信息（如果可获得的话）。在输出中，模块信息由版本行之后的多行组成，每一行由一个前导制表符缩进。

使用`go doc runtime/debug.BuildInfo`查看更多。

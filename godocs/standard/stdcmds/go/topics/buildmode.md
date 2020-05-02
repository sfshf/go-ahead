
[Build modes](https://golang.google.cn/cmd/go/#hdr-Build_modes)


##### 说明：

`go build`和`go install`命令接收一个`-buildmode`参数，该参数指示要构建哪种对象文件。目前支持的值有：

| 标志值 | 说明 |
|--|--|
| -buildmode=archive | 将列出的非主包构建到`.a`文件中。名为`main`的包将被忽略。 |
| -buildmode=c-archive | 将列出的主包及其导入的所有包构建到一个`C语言`归档文件中。可调用的标识符（symbol）是那些使用了`cgo`的`//export`注释导出的函数。只需要列出一个主包。 |
| -buildmode=c-shared | 将列出的主包及其导入的所有包构建到`C语言`共享库中。唯一可调用的符号是那些使用了`cgo`的`//export`注释导出的函数。只需要列出一个主包。 |
| -buildmode=default | 列出的主包构建到可执行文件中，而列出的非主包构建到`.a`文件中(默认行为)。 |
| -buildmode=shared | 将列出的所有非主包合并到一个共享库中；在使用`-linkshared`选项构建时将使用共享库。名为`main`的包将被忽略。 |
| -buildmode=exe | 构建列出的主包以及它们导入到可执行文件中的所有内容。未命名为`main`的包将被忽略。 |
| -buildmode=pie | 构建列出的主包及其导入到位置独立的可执行文件（PIE，position independent executable）中的所有内容。未命名为`main`的包将被忽略。 |
| -buildmode=plugin | 将列出的主包以及它们导入的所有包构建到一个Go插件中。未命名为`main`的包将被忽略。 |

在`AIX`（advanced interactive executive）系统上，当链接一个使用了由`-buildmode=c-archive`构建的Go归档文件的C语言程序时，必须传递`-WI`和`-bnoobjreorder`两个标志到`C语言编译器`。

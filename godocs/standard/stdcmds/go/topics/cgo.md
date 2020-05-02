
[Calling between Go and C](https://golang.google.cn/cmd/go/#hdr-Calling_between_Go_and_C)


##### 说明：

在Go代码和C/C++代码之间调用有两种不同的方式。

第一种方式是`cgo`工具，它是Go供应版（Go distribution）的一部分。有关如何使用它的信息，请参阅`cgo`文档（使用`go doc cmd/cgo`查看）。

第二种方式是使用`SWIG`程序，它是一种通用的语言间接口工具。关于`SWIG`更多信息查看[`SWIG网站`](http://swig.org/)。当运行`go build`时，任何以`.swig`为扩展名的文件都将被传递给`SWIG`。任何以`.swigcxx`为扩展名的文件都将被使用`-c++`选项传递给`SWIG`。

当无论`cgo`还是`SWIG`被使用时，`go build`会将任何扩展名为`.c`、`.m`、`.s`或`.S`的文件传递给`C语言编译器`，并且将任何扩展名为`.cc`、`.cpp`或`.cxx`的文件传递给`C++语言编译器`。可以设置`CC`或`CXX`环境变量来分别确定要使用的`C`或`C++`编译器。


[Go Command Documentation](https://golang.google.cn/doc/cmd)

#### 一、简介

有一套用来构建和处理Go源码的程序。这套命令程序通常使用`go`命令调用的方式运行，从而避免直接运行。

这套命令程序通常是作为`go`命令的子命令来运行，比如`go fmt`；这样运行，被调用的命令会对Go源代码的完整包进行操作；`go`命令会使用适合于包级处理的参数调用底层二进制文件。

这些命令程序也可以使用`go tool`子命令，在不修改参数时，作为独立的二进制文件运行（如`go tool cgo`）。对于大部分命令而言，这种方式主要用于调试。某些命令（如`pprof`）只能通过`go tool`子命令访问。

最后，因为`fmt`和`godoc`命令经常被引用，所以作为名为`gofmt`和`godoc`的常规二进制文件安装。

单击下面链接以获取更多文档、调用方法和使用详细信息。

| 命令名 | 概要 |
|--|--|
| [addr2line](addr2line/README.md) | addr2line是对GNU的addr2line工具的一个最小模拟，仅足以支持`pprof`。  |
| [api](api/README.md) | api命令推断出一组Go包的导出的（开放的）API。 |
| [asm](asm/README.md) | asm命令通常使用`go tool asm`，它将源文件汇编到一个对象文件中，该对象文件的名称与被处理的源文件的文件名相同，后缀为`.o`。 |
| [buildid](buildid/README.md) | buildid命令显示或更新存储在Go包或二进制文件中的构建ID。 |
| [cgo](cgo/README.md) | cgo命令可以创建调用了C/C++代码的Go包。 |
| [compile](compile/README.md) | compile命令通常使用`go tool compile`，它将编译列在命令行上的属于同一个go包的所有源码文件。 |
| [cover](cover/README.md) | cover命令用于创建和分析由`go test -coverprofile=cover.out`生成的覆盖率配置文件。 |
| [dist](dist/README.md) | dist命令帮助引导、构建和测试Go发行版。 |
| [doc](doc/README.md) | doc命令，通常使用`go doc`，为Go包提取并生成说明文档。 |
| [fix](fix/README.md) | fix命令查找使用了过时的API的Go程序，并使用新的API重写程序。 |
| [fmt](fmt/README.md) | fmt命令格式化Go包。 |
| [go](go/README.md) | go命令管理Go源码以及运行其他被列出的命令。 |
| [link](link/README.md) | link命令通常使用`go tool link`，读取main包的go档案文件或对象文件，以及main包的依赖项，然后将它们组合成可执行二进制文件。 |
| [nm](nm/README.md) | nm命令列出对象文件、档案文件或可执行文件定义或使用的symbol。 |
| [objdump](objdump/README.md) | objdump命令反汇编可执行文件。 |
| [pack](pack/README.md) | pack命令是传统Unix系统上`ar`命令工具（ar - create, modify, and extract from archives）的简单版本。 |
| [pprof](pprof/README.md) | pprof命令用来解释和显示Go程序的配置文件。 |
| [test2json](test2json/README.md) | test2json命令将`go test`的输出转换为机器可读的JSON流。 |
| [trace](trace/README.md) | trace命令是查看跟踪文件的工具。 |
| [vet](vet/README.md) | vet命令检查Go源代码并报告可疑的构造，例如，Printf调用时参数与格式字符串不匹配的情况。 |

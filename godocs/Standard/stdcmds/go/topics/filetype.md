
[File types](https://golang.google.cn/cmd/go/#hdr-File_types)


##### 说明：

`go命令`会检查每个目录中受限制的一组文件的内容。它根据文件名的扩展名识别要检查的文件。这些扩展名有:

| 扩展名 | 说明 |
|--|--|
| .go | `Go语言`源码文件。 |
| .c, .h | `C语言`源码文件。如果包使用`cgo`或`SWIG`，这些文件将用操作系统原生编译器(OS-native compiler，通常是`gcc`)进行编译;否则它们将触发一个错误。 |
| .cc, .cpp, .cxx, .hh, .hpp, .hxx | `C++语言`源码文件。只对`cgo`或`SWIG`有用，并且总是用操作系统原生编译器进行编译。 |
| .m | `Objective-C源码文件`。只对`cgo`有用，并且总是用操作系统原生编译器进行编译。 |
| .s, .S | `Assembler`（汇编）源码文件。如果包使用了`cgo`或`SWIG`，这些文件将被操作系统原生汇编器（通常是`gcc (sic)`）进行汇编。否则，这些文件将会被Go语言汇编器进行汇编。 |
| .swig, .swigcxx | `SWIG`定义文件。 |
| .syso | 系统对象文件。 |

除了`.syso`之外，所有这些类型的文件都可能包含构建约束，但是`go命令`将停止扫描文件中第一个`不是空行或//样式的行注释`的项的构建约束。有关更多细节，请参阅`go/build`包文档。
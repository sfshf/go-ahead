
[Remove object files and cached files](https://golang.google.cn/cmd/go/#hdr-Remove_object_files_and_cached_files)


##### 一、用法

```

go clean [clean flags] [build flags] [packages]

```


##### 二、clean命令说明

`clean命令`从包源目录中删除对象文件。`go命令`在一个临时目录中构建大多数对象，所以`go clean`主要涉及其他工具留下的对象文件或手动调用`go build`时遗留的对象文件。

如果给定了一个包参数或设置了`-i`或`-r`标志，`clean命令`将根据导入的路径在每个源目录中删除下表中列出的各种文件：

|文件类型| 简单说明 |
|--|--|
| _obj/ | 编译生成文件时（Makefiles）遗留的旧的对象目录。 |
| _test/ | 编译生成文件时（Makefiles）遗留的旧的测试目录。 |
| _testmain.go | 编译生成文件时（Makefiles）遗留的旧的测试文件。 |
| test.out | 编译生成文件时（Makefiles）遗留的旧的测试日志。 |
| build.out | 编译生成文件时（Makefiles）遗留的旧的生成日志。 |
| *.[568ao] | 编译生成文件时（Makefiles）遗留的对象文件。 |
|||
| DIR(.exe) | `go build`遗留的旧文件。 |
| DIR.test(.exe) | `go test -C`遗留的旧文件。 |
| MAINFILE(.exe) | `go build MAINFILE.go`遗留的旧文件。 |
| *.so | [SWIG](http://www.swig.org/)遗留的旧文件。 |

在上表中，`DIR`表示目录的最后一个path元素，`MAINFILE`表示在构建包时没有涵盖的目录中的任何Go源文件的基名。

`-i`标志使`clean命令`删除相应的已安装存档文件或二进制文件（`go install`可能创建的文件）。

`-n`标志使`clean命令`打印它将执行的删除命令，但不运行。

`-r`标志使`clean命令`递归地应用于由导入路径指名的包的所有依赖项。

`-x`标志使`clean命令`在执行删除命令时打印它们。（与`-n`标志对照）。

`-cache`标志使`clean命令`删除所有由`go build`产生的缓存。

`-testcache`标志使`clean命令`使所有在`go build`产生的缓存里的测试结果过期。

`-modcache`标志使`clean命令`删除整个模块的下载缓存，包括已解压的设定了版本的依赖项的源代码。

使用`go help build`[查看更多构建的标志](build.md#三build命令的flag说明)。

使用`go help packages`查看更多关于包的说明。

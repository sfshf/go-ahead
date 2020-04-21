
[GOPATH environment variable](https://golang.google.cn/cmd/go/#hdr-GOPATH_environment_variable)


##### 说明：

Go路径用于解析导入语句。它是由`go/build`包实现并记录的。

`GOPATH`环境变量列出了搜索Go代码的地方。在`Unix系统`上，`GOPATH`的值是一个用逗号隔开的字符串；在`Windows系统`上，`GOPATH`的值是一个用分号隔开的字符串；在`Plan 9系统`上，`GOPATH`的值是一个列表。

如果环境变量未设置，`GOPATH`默认设置为用户主目录中名为`go`的子目录（在`Unix系统`上为`$HOME/go`， 在`Windows系统`上为`%USERPROFILE%\go`），除非该目录包含有Go分发版。运行`go env GOPATH`来查看当前系统的`GOPATH`。

`GOPATH`中列出的每个目录必须有一个规定的结构:

`src`目录里持有源代码。`src`目录下的路径决定了导入路径或可执行文件的名字。

`pkg`目录里持有被安装的包对象。与Go语言SDK包树一样，每个目标操作系统与体系结构的对都在`pkg`目录下有自己的子目录（即：`pkg/GOOS_GOARCH`）。

如果`DIR`是`GOPATH`中列出的目录，那么在`DIR/src/foo/bar`中的源码包可以被导入为`foo/bar`，并将其编译后的形式安装到`DIR/pkg/GOOS_GOARCH/foo/bar.a`中。

`bin`目录里持有已编译的命令。每个命令都是根据其源目录命名的，但是只使用了最后一个元素，而不是整个路径。也就是说，在`DIR/src/foo/quux`目录里命令源码会被安装到`DIR/bin/quux`中，而不是`DIR/bin/foo/quux`中。`foo/`前缀被剥去了，以便可以将`DIR/bin`添加到`PATH`环境变量中，以获得已安装的命令。如果设置了`GOBIN`环境变量，则将命令安装到`GOBIN`所指名的目录，而不是`DIR/bin`。`GOBIN`一定是一条绝对路径。

下面是一个示例目录：

```

GOPATH=/home/user/go

/home/user/go/
    src/
        foo/
            bar/               (go code in package bar)
                x.go
            quux/              (go code in package main)
                y.go
    bin/
        quux                   (installed command)
    pkg/
        linux_amd64/
            foo/
                bar.a          (installed package object)

```

`Go编译器`会搜索在`GOPATH`中列出的每个目录以找到源代码，但是新的包总是下载到列表中的第一个目录中。

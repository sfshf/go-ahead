
[Internal Directories](https://golang.google.cn/cmd/go/#hdr-Internal_Directories)


##### 说明：

`internal`目录里的或目录之下的代码只能由以`internal`目录的父目录为根目录的目录树中的代码导入。

下面是一个示例目录：

```

/home/user/go/
    src/
        crash/
            bang/              (go code in package bang)
                b.go
        foo/                   (go code in package foo)
            f.go
            bar/               (go code in package bar)
                x.go
            internal/
                baz/           (go code in package baz)
                    z.go
            quux/              (go code in package main)
                y.go

```

`z.go`中的代码可以被导入为`foo/internal/baz`，但该导入语句只能出现在以`foo`为根目录的子树中的源文件中。源文件`foo/f.go`，`foo/bar/x.go`和`foo/quux/y.go`都可以导入`foo/internal/baz`，但是源文件`crash/bang/b.go`不能。

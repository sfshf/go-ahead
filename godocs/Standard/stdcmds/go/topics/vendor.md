
[Vendor Directories](https://golang.google.cn/cmd/go/#hdr-Vendor_Directories)


##### 说明：

`Go 1.6`支持使用外部依赖项的本地副本来满足依赖项的导入，通常称为`供应（vendoring）`。

在名为`vendor`的目录下的代码只能被以`vendor`目录的父目录为根目录的目录树里的代码导入；只使用将前缀省略到`vendor`元素并包含该元素的导入路径。

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
            vendor/
                crash/
                    bang/      (go code in package bang)
                        b.go
                baz/           (go code in package baz)
                    z.go
            quux/              (go code in package main)
                y.go

```

类似应用于`internal`目录的可见性规则，但是`z.go`中的代码被导入为`baz`，而不是`foo/vendor/baz`。

源代码树中较深的`vendor`目录中的代码会隐蔽掉较高的目录中的代码。在以`foo`为根的子树中，`crash/bang`的导入会解析为`foo/vendor/crash/bang`，而不是顶级的`crash/bang`。

`vendor`目录中的代码不受导入路径检查的约束(参见`go help importpath`)。

当`go get`检验或更新`git存储库`时，它现在还更新子模块。

`vendor`目录不会影响`go get`首次检出的新存储库的位置：这些存储库总是放在主`GOPATH`中，而不是放在`vendor`目录子树中。

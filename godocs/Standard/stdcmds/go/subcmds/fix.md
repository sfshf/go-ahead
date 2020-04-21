
[Update packages to use new APIs](https://golang.google.cn/cmd/go/#hdr-Update_packages_to_use_new_APIs)

##### 一、用法

```

go fix [packages]

```

##### 二、fix命令说明

`fix命令`对以导入路径形式命名的包运行`go fix`命令。

使用`go doc cmd/fix`查看更多关于`fix命令`的信息。使用`go help packages`查看更多关于包的详细说明。

想要使用特殊的选项运行`fix命令`，可以使用`go tool fix`命令。

更多信息参考[`go fmt`](fmt.md)，[`go vet`](vet.md)。

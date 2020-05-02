
[Module and vendoring](https://golang.google.cn/cmd/go/#hdr-Modules_and_vendoring)


##### 说明：

当使用模块时，`go命令`完全忽略`vendor`目录。

默认情况下，`go命令`通过从模块的源下载模块并使用这些下载的副本（在验证之后，如[Module downloading and verification](moduledownload.md)所述）来满足依赖关系。为了可以与老版本的Go语言SDK互操作,或为了确保用于构建的所有文件一起存储在一个独立的文件树中,`go mod vendor`会在主模块的根目录里创建一个名为`vendor`的目录，并存储依赖模块的所有包，这些包是用来支持主模块的构建和测试的依赖包。

以使用主模块的顶级`vendor`目录来满足依赖关系的方式来构建项目(禁用通用的网络源和本地缓存)，请使用`go build -mod=vendor`。注意，只有主模块的顶级`vendor`目录起作用，其他位置的`vendor`目录仍然会被忽略。

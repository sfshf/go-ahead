
[Build and test caching](https://golang.google.cn/cmd/go/#hdr-Build_and_test_caching)


##### 说明：

`go`命令会缓存`build`的输出，以便在将来的`build`中重用。缓存数据的默认位置是当前操作系统的标准用户缓存目录中名为`go-build`的子目录。设置`GOCACHE`环境变量会覆盖这个默认值，运行`go env GOCACHE`会打印当前缓存目录。

`go`命令定期删除最近没有使用的缓存数据。运行`go clean -cache`删除所有缓存的数据。

构建缓存正确地考虑到源文件、编译器、编译器选项等的更改：在通常情况下不需要显式地清理缓存。但是，构建缓存不会检测到使用`cgo`导入的`C语言库`的更改。如果您对系统上的`C语言库`进行了更改，那么您需要显式地清理缓存，或者使用构建标志`-a`(参见`go help build`)强制重新构建依赖于更新的`C语言库`的包。

`go`命令还缓存成功的包测试结果。详见`go help test`。运行`go clean -testcache`将删除所有缓存的测试结果(但不删除缓存的构建结果)。

`GODEBUG`环境变量可以打印关于缓存状态的调试信息:

`GODEBUG=gocacheverify=1`可以让`go`命令避免任何缓存条目的使用，取而代之的是重新构建所有内容，然后检查结果是否与现有的缓存条目匹配。

`GODEBUG=gocachehash=1`可以让`go`命令打印用于构造缓存查找键的所有内容散列的输入。输出很大，但是对于调试缓存非常有用。

`GODEBUG=gocachetest=1`可以让`go`命令打印关于是否重用缓存的测试结果的决策细节。

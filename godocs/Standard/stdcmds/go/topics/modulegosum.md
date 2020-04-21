
[Module authentication using go.sum](https://golang.google.cn/cmd/go/#hdr-Module_authentication_using_go_sum)


##### 说明：

`go命令`尝试对每个下载的模块进行身份验证，检查今天下载的特定模块版本的位与昨天下载的位是否匹配。这确保了可重复的构建，并检测引入的意外更改，无论是否有恶意。

在每个模块的根目录中，与`go.mod`一起的，`go命令`还维护了一个名为`go.sum`的文件。包含模块依赖项的密码校验和。

`go.sum`文件的每一行有三个字段：`<module> <version>[/go.mod] <hash>`。

每个已知的模块版本在`go.sum`文件中都占有两行。第一行给出模块版本的文件树的散列。第二行追加了`/go.mod`到版本号之后，并给出用来管理整个的模块版本的（因为模块版本可能是合成的）唯一的`go.mod`文件的散列。`go.mod-only hash`允许下载和验证模块版本的`go.mod`文件（计算依赖关系图是需要这样做），而不需要下载所有模块的源代码。

散列以`h<N>:`形式的算法前缀开始。目前唯一定义的算法前缀是`h1:`，它使用`SHA-256`。

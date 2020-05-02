
[Module authentication failures](https://golang.google.cn/cmd/go/#hdr-Module_authentication_failures)


##### 说明：

`go命令`维护下载包的缓存，并在下载时计算和记录每个包的密码校验和。在正常操作中，`go命令`会根据这些在下载时预先计算的校验和检查主模块的`go.sum`文件，而不是在每次命令调用时重新计算它们。`go mod verify`命令可以检查模块下载的缓存副本是否仍然匹配`go.sum`文件中记录的校验和以及条目。

在日常开发中，给定模块版本的校验和不应该改变。每当一个给定的主模块使用一个依赖项时，`go命令`就会根据主模块的`go.sum`文件检查它的本地缓存副本(是否已下载)。如果校验和不匹配，`go命令`将不匹配报告为安全错误，并拒绝运行构建。当这种情况发生时，请谨慎从事:代码意外更改意味着今天的构建将与昨天的不匹配，而且意外更改可能没有好处。


如果`go命令`报告`go.sum`文件不匹配，则被报告的模块版本的下载代码与以前的主模块构建中使用的代码不匹配。在这一点上，找出正确的校验和是很重要的，这将决定是`go.sum`文件错误还是下载的代码错误。通常`go.sum`文件是正确的：前提是您希望使用与昨天相同的代码。

如果下载的模块还没有包含在`go.sum`文件中，且它是一个公共可用的模块，`go命令`会咨询`Go checksum database（Go语言校验和数据库）`来获取期望的`go.sum`行。如果下载的代码与这些`go.sum`行不匹配，`go命令`将报告不匹配并退出。注意，对于已经在`go.sum`中列出的模块版本，不需要查询数据库。

当一个`go.sum`不匹配的问题被报告时，调查为什么今天下载的代码与昨天下载的不同是很有必要的，是值得去调查的。

`GOSUMDB`环境变量用于指明被应用的校验和数据库的名称，以及可选择性地指明数据库的公钥和URL。举例：

```

GOSUMDB="sum.golang.org"
GOSUMDB="sum.golang.org+<publickey>"
GOSUMDB="sum.golang.org+<publickey> https://sum.golang.org"

```

`go命令`知道`sum.golang.org`的公钥，也知道`sum.golang.google.cn`（在中国大陆地区可用）连接到`sum.golang.org`校验和数据库；使用任何其他数据库都需要显式地提供公钥。URL默认为`https://`，后跟数据库名。

`GOSUMDB`默认值为`sum.golang.org`，一个由谷歌管理的`Go checksum database`。关于服务的隐私策略可查看`https://sum.golang.org/privacy`。

如果`GOSUMDB`被设置为`off`，或者`go get`命令使用了`-insecure`标志进行调用，那么将不访问校验和数据库，并接受所有未识别的模块，代价是放弃对所有模块的可重复下载验证的安全保证。绕过特定模块的校验和数据库的更好方法是使用`GOPRIVATE`或`GONOSUMDB`环境变量。详情请参见`go help module-private`。

`go env -w`命令(参见`go help env`)可用于为将来的`go命令`调用设置这些变量。

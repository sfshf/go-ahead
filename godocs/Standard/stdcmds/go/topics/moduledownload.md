
[Module downloading and verification](https://golang.google.cn/cmd/go/#hdr-Module_downloading_and_verification)


##### 说明：

根据`GOPROXY`环境变量的设置（参见`go help env`），`go命令`可以从代理中获取模块，也可以直接连接到源代码管理服务器。`GOPROXY`环境变量的默认值为`https://proxy.golang.org,direct`，这意味着尝试由谷歌运行的Go模块镜像，如果代理报告它没有该模块（HTTP错误`404`或`410`），则退回到直接连接。获取服务的隐私策略可查看[https://proxy.golang.org/privacy](https://proxy.golang.org/privacy)。如果`GOPROXY`被设置为字符串`direct`，下载将使用到源代码管理服务器的直接连接。设置`GOPROXY`为`off`不允许从任何来源下载模块。否则，`GOPROXY`将是一个逗号分隔的模块代理url列表，在这种情况下，`go命令`将从这些代理中获取模块。对于每个请求，`go命令`依次尝试每个代理，只有在当前代理返回`404`或`410`HTTP响应时才移动到下一个代理。字符串`direct`可能会出现在代理列表中，从而导致在搜索的那一点尝试直接连接。在`direct`之后列出的任何代理都不会被咨询。

`GOPRIVATE`和`GONOPROXY`环境变量允许绕过选定模块的代理。详情请参见`go help module-private`。

无论模块的来源是什么，`go命令`都会根据已知的`checksum（检验和）`检查下载，以检测任何特定模块版本的内容从第一天到第二天的意外变化。这个检查首先参考当前模块的`go.sum`文件，但是返回到`Go checksum database（go语言校验和数据库）`，该数据库可通过`GOSUMDB`和`GONOSUMDB`环境变量来控制。详情请参见`go help module-auth`。

有关代理协议和缓存下载包格式的详细信息，请参阅`go help goproxy`。

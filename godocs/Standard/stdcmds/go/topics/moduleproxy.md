
[Module proxy protocol](https://golang.google.cn/cmd/go/#hdr-Module_proxy_protocol)


##### 说明：

Go模块代理是任何可以响应特定格式的URL的`GET`请求的web服务器。请求没有查询参数，因此即使是来自固定文件系统的站点（包括`file:///`格式的URL）也可以是模块代理。

被传给Go模块代理的`GET`请求有：

`$GOPROXY/<module>/@v/list`的`GET`请求，返回被给定模块的所有已知版本的一个列表，逐行显示模块版本。

`$GOPROXY/<module>/@v/<version>.info`的`GET`请求，返回被给定模块版本的JSON格式的源信息。

`$GOPROXY/<module>/@v/<version>.mod`的`GET`请求，返回被给定模块版本的`go.mod`文件。

`$GOPROXY/<module>/@v/<version>.zip`的`GET`请求，返回被给定模块版本的`zip`档案文件。

为了避免从区分大小写的文件系统提供服务时出现问题，`<module>`和`<version>`两个元素是用大小写编码的（case-encoded），每个大写字母都用一个感叹号后面跟着相应的小写字母的格式来替代，例如：`github.com/Azure`编码为`github.com/!azure`。

给定模块的JSON格式元数据对应于下面这个Go数据结构，将来可能会被扩展:

```go

type Info struct {
    Version string    // version string
    Time    time.Time // commit time
}

```

给定模块的特定版本的`zip`归档文件是一个标准`zip`文件，其中包含与模块源代码和相关文件对应的文件树。`zip`存档使用斜杠分隔的路径，存档中的每个文件路径都必须以`<module>@<version>/`开始，其中的`<module>`和`<version>`会被直接地替换，而不是大小写编码地（case-encoded）。模块文件树的根对应于归档文件中的`<module>@<version>/`前缀。

即使当从版本控制系统直接下载时，`go命令`会合成显式的`info`、`mod`和`zip`文件，并将这些文件存储在本地缓存目录`$GOPATH/pkg/mod/cache/download`里；从一个代理站点直接下载这些文件时也是一样的。缓存布局和代理站点的URL空间是一样的，所以在`https://example.com/proxy`代理站点提供`$GOPATH/pkg/mod/cache/download`服务（或复制到代理站点），可以让其他用户通过设置`GOPROXY=https://example.com/proxy`访问被缓存的模块版本。


# `net`包

`net`包为`网络I/O`提供了一个`可移植的接口`，包括`TCP/IP`，`UDP`，`域名解析`和`Unix域套接字`。

尽管该包`提供了对低级网络原语的访问权限`，但大多数客户端仅需要`Dial`，`Listen`和`Accept`函数提供的基本接口以及关联的`Conn`和`Listener`接口。`crypto/tls`包使用相同的接口和相似的`Dial`and`Listen`功能。

`Dial`函数用来连接到服务器：

```go

conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...

```

`Listen`函数用来创建服务器：

```go

ln, err := net.Listen("tcp", ":8080")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
    }
    go handleConnection(conn)
}

```

# 域名解析

`域名解析方法`（无论是`间接`使用`Dial`函数还是`直接`使用`LookupHost`和`LookupAddr`等函数）因操作系统而异。

`在Unix系统上`，`解析器`有两个用于解析名称的选项。它可以使用`纯Go解析器`直接将`DNS请求`发送到`/etc/resolv.conf`中列出的服务器，也可以使用`基于cgo的解析器`来`调用C库例程`，例如`getaddrinfo`和`getnameinfo`。

`默认情况下，使用纯Go解析器`，因为`被阻止的DNS请求仅消耗一个goroutine`，而`被阻止的C调用消耗一个操作系统线程`。当`cgo`可用时，将在多种情况下使用`基于cgo的解析器`：在`不允许程序直接发出DNS请求（OS X）的系统`上，当存在`LOCALDOMAIN`环境变量（即使为空）时，`RES_OPTIONS`或`HOSTALIASES`环境变量为非空，当`ASR_CONFIG`环境变量为非空（仅适用于OpenBSD）时，当`/etc/resolv.conf`或`/etc/nsswitch.conf`指定使用`Go解析器`未实现的功能时，以及当要查找的名称以`.local`结尾或者为`mDNS`时。

可以通过将`GODEBUG`环境变量的`netdns`值（请参阅`runtime`包）设置为`go`或`cgo`来覆盖`解析程序`的决策，如下所示：

```sh

export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force cgo resolver

```

在`构建Go源代码树`时可以通过设置`netgo`或`netcgo build`标签来强制执行该决策。

将`netdns`设置为数字，如`GODEBUG=netdns=1`，会使`解析程序`打印有关其决策的调试信息。要`在打印调试信息的同时强制使用特定的解析器`，请使用`加号`将两个设置结合在一起，如`GODEBUG=netdns=go+1`。

`在Plan 9上`，`解析程序`始终访问`/net/cs`和`/net/dns`。

`在Windows上`，`解析器`始终使用`C库函数`，例如`GetAddrInfo`和`DnsQuery`。


# `net/http`包

`http`包提供了`HTTP客户端`和`HTTP服务器`的实现。

`Get`，`Head`，`Post`和`PostForm`用于创建`HTTP（或HTTPS）`请求：

```go

resp, err := http.Get("http://exmaple.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"key": {"Value"}, "id": {"123"}}})

```

`客户端a`必需在使用完`响应体`后将其关闭。

```go

resp, err := http.Get("http://example.com/")
if err != nil {
    // handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...

```

想要控制`HTTP客户端`的`头部`，`重定向策略`和其他设置，可以自定义创建`Client`：

```go

client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy`)
resp, err := client.Do(req)
// ...

```

想要控制`策略`，`TLS配置`，`长连接`，`压缩`和其他设置，可以创建一个`Transport`：

```go

tr := &http.Transport{
    MaxIdleConns:       10,
    IdleConnTimeout:    30*time.Second,
    DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://exmaple.com")

```

`Client`和`Transport`在多协程下并发安全，为了效能应该只要创建一次并重复使用。

`ListenAndServe`可以使用给定的`地址`和`处理函数`开启一个`HTTP服务器`。`处理函数`通常是`nil`，意味着使用`DefaultServeMux`。`Handle`和`HandleFunc`可以用于给`DefaultServeMux`添加`处理函数`：

```go

http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))

```

想要对`服务器的行为`进行更过的控制，可以通过一个自定义的`Server`来实现：

```go

s := &http.Server{
    Addr:           ":8080",
    Handler:        myHandler,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())

```

`从Go 1.6开始`，使用`HTTPS`时，`http包`对`HTTP/2协议`具有显式的支持。必须要`禁用HTTP/2`的程序，可以通过将`Transport.TLSNextProto（对于客户端）`或`Server.TLSNextProto（对于服务器）`设置为`非nil的空映射`来执行此操作。另一种方式是，使用目前支持的以下`GODEBUG`环境变量：

```go

GODEBUG=http2client=0 # disable HTTP/2 client support
GODEBUG=http2server=0 # disable HTTP/2 server support
GODEBUG=http2debug=1  # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2  # ... even more verbose, with frame dumps

```

`Go的API兼容性的保证`不涵盖`GODEBUG`变量。请在禁用HTTP/2支持之前，先报告所有问题：[https://golang.org/s/http2bug](https://golang.org/s/http2bug)。

`http包`的`Transport`和`Server`都自动允许`HTTP/2`进行简单配置。想要允许`HTTP/2`进行更复杂的配置，使用`较低级别的HTTP/2功能`或使用`Go的http2软件包的较新版本`，请直接导入`golang.org/x/net/http2`并使用其`ConfigureTransport`和/或`ConfigureServer`函数。通过`golang.org/x/net/http2包`手动配置的`HTTP/2`要优先于`net/http包的内置HTTP/2支持`。

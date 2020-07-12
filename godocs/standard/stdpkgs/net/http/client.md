
# `Client`结构体

`Client`是`HTTP客户端`。它的`零值`（`DefaultClient`）是使用`DefaultTransport`的可用客户端。

`Client`的`Transport（传输）`通常具有`内部状态（缓存的TCP连接）`，因此客户端应该要重复使用，而不是根据需要重新创建。`客户端可以安全地被多个goroutine并发使用`。

`Client`是比`RoundTripper`（例如`Transport`）更高层级的，并且还能处理`HTTP详细信息`，例如`cookie`和`重定向`。

后跟重定向时，`Client`将转发在初始`Request`时设置的所有`头部`，但以下情况除外：

- 将`敏感的头部`（如`Authorization`，`WWW-Authenticate`和`Cookie`）转发到`不受信任的目标`时。当重定向到`与子域不匹配`或`与初始域不完全匹配的域`时，将忽略这些`头部`。例如，从`foo.com`重定向到`foo.com`或`sub.foo.com`将会转发`敏感头部`，但重定向到`bar.com`则不会转发。

- 使用`非空cookie包（cookie Jar）`转发`Cookie标头`时。由于每个重定向都可能会更改`cookie包（cookie Jar）的状态`，因此重定向可能会更改`初始请求中设置的cookie`。当转发`Cookie头部`时，任何`更改的cookie`都将被省略，不要期望`cookie包（Jar）`将`更新的值`插入到那些`更改的cookie`（假设原点匹配）。如果`cookie包（Jar）`为`nil`，则将转发`原始cookie`，而不进行任何更改。

```go

type Client struct {

    /*
        `Transport`，指定单个HTTP请求采取的机制。

        如果为`nil`，则使用`DefaultTransport`。
    */
    Transport RoundTripper

    /*
        `CheckRedirect`，指定用于处理重定向的策略。
        如果`CheckRedirect`不为`nil`，则客户端将在进行HTTP重定向之前调用它。参数`req`和`via`是即将到来的请求和已发出的请求，先到先处理。如果`CheckRedirect`返回错误，则`Client`的`Get`方法将返回先前的`Response`（关闭其`Body`）和`CheckRedirect`的错误（包装在`url.Error`中），而不是发出请求`req`。

        作为一种特殊情况，如果`CheckRedirect`返回`ErrUseLastResponse`，则返回最近的响应（要求其主体是未关闭的），并返回`nil`错误。

        如果`CheckRedirect`为`nil`，则客户端使用其`默认策略`，该策略将在`连续10个请求`后停止。
    */
    CheckRedirect func(req *Request, via []*Request) error

    /*
        `Jar`，指定cookie包。

        `Jar`用于将相关cookie插入每个出站的Request，并使用每个入站的Response的cookie值进行更新。`Client`遵循的每个重定向都会咨询`Jar`。

        如果`Jar`为`nil`，则仅当在`Request`上显式设置cookie时，才发送cookie。
    */
    Jar CookieJar

    /*
        `Timeout`，指定此`Client`发出的请求的时间限制。超时包括连接时间，任何重定向和读取响应正文。 在`Get`，`Head`，`Post`或`Do`返回之后，计时器保持运行，并且将中断`Response.Body`的读取。

        `Timeout`为零值时表示没有超时。

        `Client`取消对底层传输的请求，就像`Request`的`Context`结束一样。

        为了兼容性，`Client`还将在`Transport`上使用不建议使用的`CancelRequest`方法（如果有）。新的`RoundTripper`实现应使用`Request`的`Context`进行取消，而不是实现`CancelRequest`。
    */
    Timeout time.Duration   // Go 1.3

}

```


## `RoundTripper`接口和`Transport`结构体

`RoundTripper`是表示`执行单个HTTP事务`，`获取给定请求的响应`的能力的接口。

`RoundTripper`必须可以同时被多个goroutine安全地使用。

```go

type RoundTripper interface {

    /*
        `RoundTrip`执行一个`HTTP事务`，为提供的Request返回一个Response。

        `RoundTrip`不应尝试解释该响应。特别是，如果`RoundTrip`获得响应，则无论响应的HTTP状态码如何，都必须返回`err == nil`。
        如果获取响应失败，则应保留一个`非nil的错误`。同样地，`RoundTrip`不应尝试处理更高层级的协议详细信息，例如重定向，身份验证或cookie。

        除了消耗和关闭`Request`的`Body`之外，`RoundTrip`不应修改请求。`RoundTrip`可以在单独的goroutine上读取请求的字段。在`Response`的`Body`关闭之前，调用者不应更改或重用该请求。

        `RoundTrip`必须始终关闭主体，包括发生错误时，但根据实现的不同，甚至在`RoundTrip`返回之后，也应该在单独的goroutine中关闭它。这意味着，希望给后续请求重用主体的调用者，必须安排在等待`Close`调用之后再这样做。

        `Request`的`URL`和`Header`字段必须是被初始化的。
    */
    RoundTrip(*Request) (*Response, error)

}

```

`DefaultTransport`是`Transport`的默认实现，由`DefaultClient`使用。它根据需要建立网络连接，并将其缓存以供后续调用进行重用。它使用`$HTTP_PROXY`和`$NO_PROXY`（或`$http_proxy`和`$no_proxy`）环境变量指示的`HTTP代理`。

```go

var DefaultTransport RoundTripper = &Transport{
    Proxy: ProxyFromEnvironment,
    DialContext: (&net.Dialer{
        Timeout: 30*time.Second,
        KeepAlive: 30*time.Second,
        DualStack: true,
    }).DialContext,
    ForceAttemptHTTP2: true,
    MaxIdleConns: 100,
    IdleConnTimeout: 90*time.Second,
    TLSHandshakeTimeout: 10*time.Second,
    ExpectContinueTimeout: 1*time.Second,
}

```

`Transport`是`RoundTripper`的一个实现，它支持`HTTP`，`HTTPS`和`HTTP代理`（对于`HTTP`或`带CONNECT的HTTPS`）。

默认情况下，`Transport`会`缓存连接`以供将来重用。访问许多主机时，这可能会留下许多打开的连接。可以使用`Transport`的`CloseIdleConnections`方法以及`MaxIdleConnsPerHost`和`DisableKeepAlives`字段来管理此行为。

`Transport`应该被重用，而不是根据需要就创建。`Transport`可以安全地同时被多个goroutine使用。

`Transport`是用于发出`HTTP`和`HTTPS`请求的`低层原语`。有关`Cookie`和`重定向`之类的`高层级功能`，请参阅`Client`结构体。

`Transport`为`HTTP URL`使用`HTTP/1.1`，为`HTTPS URL`使用`HTTP/1.1`或`HTTP/2`，这取决于服务器是否支持`HTTP/2`，以及`Transport`的配置方式。`DefaultTransport`支持`HTTP/2`。想要在一个`传输`上`显式`启用`HTTP/2`，请使用`golang.org/x/net/http2`包并调用`ConfigureTransport`。有关`HTTP/2`的更多信息，请参见其软件包文档。

`状态码`在`1xx`范围内的响应将自动处理（`100`--期望继续）或被忽略。一个例外是`HTTP状态码101`（`交换协议`），它被视为`终端状态`并由`RoundTrip`返回。若要查看被忽略的`1xx`响应，请使用`httptrace`跟踪包的`ClientTrace.Got1xxResponse`。

如果`请求是幂等的`且要么`没有主体`要么就定义了其`Request.GetBody`，则`Transport`仅在遇到网络错误时重试该请求。如果`HTTP请求`具有`HTTP方法`--`GET`，`HEAD`，`OPTIONS`或`TRACE`，则它们被认为是`幂等的`。或者其`头部键值对`包含`Idempotency-Key`或`X-Idempotency-Key`条目。如果该`幂等性的键值`为零长度的切片，那么该请求仍被视为`幂等`，但其`头部`不会在线上发送。

```go

type Transport struct {

    idleMu       sync.Mutex
	closeIdle    bool                                // user has requested to close all idle conns
	idleConn     map[connectMethodKey][]*persistConn // most recently used at end
	idleConnWait map[connectMethodKey]wantConnQueue  // waiting getConns
	idleLRU      connLRU
	reqMu       sync.Mutex
	reqCanceler map[*Request]func(error)
	altMu    sync.Mutex   // guards changing altProto only
	altProto atomic.Value // of nil or map[string]RoundTripper, key is URI scheme
	connsPerHostMu   sync.Mutex
	connsPerHost     map[connectMethodKey]int
	connsPerHostWait map[connectMethodKey]wantConnQueue // waiting getConns

    /*
        `Proxy`指定一个函数，用于为给定`Request`返回一个代理。如果该函数返回非空错误，则请求将因所提供的错误而中止。

        代理类型由`URL的方案`确定。支持`http`，`https`和`socks5`。如果方案为空，则假定为`http`。

        如果`Proxy`为空或返回空的`*URL`，则不使用任何代理。
    */
    Proxy func(*Request) (*url.URL, error)

    /*
        `DialContext`指定用于创建`未加密的TCP连接`的拨号函数。
        如果`DialContext`为空（并且下面的弃用的`Dial`字段也为空），则传输使用`net包`进行拨号。

       `DialContext`与`RoundTrip`的调用同时运行。
       当较早的连接在之后的`DialContext`完成之前处于空闲状态时，发起拨号的`RoundTrip`调用可能会使用先前拨打的连接结束。
    */
    DialContext func(ctx context.Context, network, addr string) (net.Conn, error)   // Go 1.7

    /*
        `Dial`指定用于创建`未加密的TCP连接`的拨号函数。

        `Dial`与`RoundTrip`的调用同时运行。
        当较早的连接在之后的`Dial`完成之前处于空闲状态时，发起拨号的`RoundTrip`调用可能会使用先前拨打的连接结束。

        弃用：请改用`DialContext`，`DialContext`允许`传输`在不再需要拨号时立即取消它们。
        如果两者都被设置了，则优先采用`DialContext`。
    */
    Dial func(network, addr string) (net.Conn, error)

    /*
        `DialTLSContext`指定一个可选的拨号函数，用于为`无代理的HTTPS请求`创建`TLS连接`。

        如果`DialTLSContext`为空（并且下面弃用的`DialTLS`也为空），则使用`DialContext`和`TLSClientConfig`。

        如果设置了`DialTLSContext`，则`Dial`和`DialContext`钩子不用于`HTTPS请求`，并且`TLSClientConfig`和`TLSHandshakeTimeout`将被忽略。返回的`net.Conn`被假定为已通过`TLS握手`。
    */
    DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)    // Go 1.14

    /*
        `DialTLS`指定用于为`无代理HTTPS请求`创建`TLS连接`的可选的拨号函数。

        弃用：改为使用`DialTLSContext`，`DialTLSContext`允许`传输`在不再需要拨号时立即取消它们。
        如果两者都被设置，则优先采用`DialTLSContext`。
    */
    DialTLS func(network, addr string) (net.Conn, error)    // Go 1.4

    /*
        `TLSClientConfig`指定要与`tls.Client`一起使用的`TLS配置`。
        如果为空，则使用默认配置。
        如果为非空，则默认情况下可能不会启用`HTTP/2`的支持。
    */
    TLSClientConfig *tls.Config

    /*
        `TLSHandshakeTimeout`指定等待TLS握手的最长时间。零表示没有超时限制。
    */
    TLSHandshakeTimeout time.Duration   // Go 1.3

    /*
        `DisableKeepAlives`（如果为true）将禁用`HTTP长连接状态`，并且仅将与服务器的连接用于单个HTTP请求。

        这与类似命名的`TCP长连接`无关。
    */
    DisableKeepAlives bool

    /*
        `DisableCompression`（如果为true），当请求没有已存在的`Accept-Encoding`值时，将阻止`Transport`使用`Accept-Encoding：gzip`请求头请求压缩。如果`Transport`本身请求gzip并获得gzip压缩的响应，则会在`Response.Body`中对其进行透明解码。但是，如果用户显式请求gzip，则不会自动将其解压缩。
    */
    DisableCompression bool

    /*
        `MaxIdleConns`控制所有主机之间的最大空闲（保持活动）连接数。零表示无限制。
    */
    MaxIdleConns int    // Go 1.7

    /*
        `MaxIdleConnsPerHost`（如果非零）控制与每个主机的最大空闲（保持长连接）连接数。如果为零，则使用`DefaultMaxIdleConnsPerHost`。
    */
    MaxIdleConnsPerHost int

    /*
        `MaxConnsPerHost`按需限制每个主机的连接总数，包括处于拨号，活动和空闲状态的连接。超出限制时，拨号将阻塞。

        零表示无限制。
    */
    MaxConnsPerHost int // Go 1.11

    /*
        `IdleConnTimeout`是保持活动状态的空闲连接在关闭自身之前能保持空闲状态的最长时间。

        零表示无限制。
    */
    IdleConnTimeout time.Duration   // Go 1.7

    /*
        `ResponseHeaderTimeout`（如果不为零）指定在完全写入请求（包括其主体（如果有））之后等待服务器的响应头的时间。此时间不包括读取响应正文的时间。
    */
    ResponseHeaderTimeout time.Duration // Go 1.1

    /*
        `ExpectContinueTimeout`（如果非零）指定如果请求具有`Expect: 100-continue`头部，则在完全写入请求头之后等待服务器的第一个响应头的时间。零表示没有超时限制，并导致正文立即发送，而无需等待服务器批准。
        此时间不包括发送请求头的时间。
    */
    ExpectContinueTimeout time.Duration // Go 1.6

    /*
        `TLSNextProto`指定在`TLS ALPN`协议协商后，`Transport`如何切换到备用协议（例如`HTTP/2`）。如果`Transport`使用非空协议名拨打`TLS连接`并且`TLSNextProto`包含该密钥的map条目（例如`h2`），则将以请求的权限（例如`example.com`或`example.com：1234`）和TLS连接来调用函数。该函数必须返回之后要处理请求的`RoundTripper`。
        如果`TLSNextProto`不为零，则不会自动启用`HTTP/2`支持。
    */
    TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper    // Go 1.6

    /*
        `ProxyConnectHeader`可以按需指定在`CONNECT请求`期间发送给代理的头部。
    */
    ProxyConnectHeader Header   // Go 1.8

    /*
        `MaxResponseHeaderBytes`指定对服务器的响应头中允许的响应字节数的限制。

        零表示使用默认限制。
    */
    MaxResponseHeaderBytes int64    // Go 1.7

    /*
        `WriteBufferSize`指定在写入`传输`时使用的写入缓冲区的大小。
        如果为零，则使用默认值（当前为4KB）。
    */
    WriteBufferSize int //Go 1.13

    /*
        `ReadBufferSize`指定从`传输`读取时使用的读取缓冲区的大小。
        如果为零，则使用默认值（当前为4KB）。
    */
    ReadBufferSize int  // Go 1.13

    /*
        nextProtoOnce guards initialization of TLSNextProto and h2transport (via onceSetNextProtoDefaults)
    */
	nextProtoOnce      sync.Once
	h2transport        h2Transport // non-nil if http2 wired up
	tlsNextProtoWasNil bool        // whether TLSNextProto was nil when the Once fired

    /*
        当提供非零`Dial`，`DialTLS`或`DialContext`函数或`TLSClientConfig`时，`ForceAttemptHTTP2`控制是否启用`HTTP/2`。
        默认情况下，保守地使用这些字段会禁用`HTTP/2`。
        想使用自定义`拨号程序`或`TLS配置`并仍尝试`HTTP/2升级`，请将该字段设置为true。
    */
    ForceAttemptHTTP2 bool  // Go 1.13

}

```


## `CookieJar`接口和`Cookie`结构体

`CookieJar`管理`HTTP请求`中`cookie的存储和使用`。

`CookieJar的实现`必须能被多个goroutine并发安全地使用。

`net/http/cookiejar`包提供了`CookieJar实现`。

```go

type CookieJar interface {

    /*
        `SetCookies`对给定URL的回复中的cookie的接收方式进行处理。它可能会或可能不会选择保存cookie，具体取决于jar的策略和实现。
    */
    SetCookies(u *url.URL, cookies []*Cookie)

    /*
        `Cookies`返回要在给定URL的请求中发送的cookie。具体的实现遵循标准的cookie使用限制（例如`RFC 6265`中的限制）。
    */
    Cookies(u *url.URL) []*Cookie

}

```

`Cookie`代表在`HTTP响应的Set-Cookie头部`或`HTTP请求的Cookie头部`中发送的`HTTP Cookie`。

有关详细信息，请参见[https://tools.ietf.org/html/rfc6265](https://tools.ietf.org/html/rfc6265)。

```go

type Cookie struct {
    Name string
    Value string

    Path string         // optional
    Domain string       // optional
    Expires time.Time   // optional
    RawExpires string   // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
    // MaxAge>0 means Max-Age attribute present and given in seconds.
    MaxAge int
    Secure bool
    HttpOnly bool
    SameSite SameSite // Go 1.11
    Raw string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}

```

`SameSite`允许`服务器`定义一个`cookie属性`，从而使浏览器`无法`将`此cookie`与`跨站点请求`一起发送。`主要目标是减少跨域信息泄漏的风险，并提供针对跨站点请求伪造攻击的某种保护`。

有关详细信息，请参见[https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00](https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00)。

```go

type SameSite int

const (
    SameSiteDefaultMode SameSite = iota + 1
    SameSiteLaxMode
    SameSiteStrictMode
    SameSiteNoneMode
)

```

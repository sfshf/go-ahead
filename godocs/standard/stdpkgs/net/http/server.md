
# `Server`结构体

`Server`定义了用于运行`HTTP服务器`的参数。`Server`的零值是有效配置。

```go

type Server struct {

    /*
        `Addr`可以选择性地以`"host:port"`的形式指定服务器监听的`TCP地址`。如果为空，则使用`":http"`（端口为`80`）。
        服务名称在`RFC 6335`中定义，并由IANA分配。
        有关`地址格式`的详细信息，请参见`net.Dial`。
    */
    Addr string

    Handler Handler // handler to invoke, `http.DefaultServeMux` if nil

    /*
        `TLSConfig`，可选地，提供`TLS配置`，以供`ServeTLS`和`ListenAndServeTLS`使用。请注意，此值是由`ServeTLS`和`ListenAndServeTLS`克隆的，因此无法使用`tls.Config.SetSessionTicketKeys`之类的方法修改配置。若要使用`SetSessionTicketKeys`，请使用带有`TLS Listener`的`Server.Serve`代替。
    */
    TLSConfig *tls.Config

    /*
        `ReadTimeout`是读取整个请求（包括正文）的最大持续时间。

        由于`ReadTimeout`不允许`处理程序`根据每个请求正文的可接受的截止时间或上载速率做出每个请求的决策，因此大多数用户将更喜欢使用`ReadHeaderTimeout`。同时使用它们是有效的。
    */
    ReadTimeout time.Duration

    /*
        `ReadHeaderTimeout`是允许读取请求头部的用时。读取头部后，将重置连接的读取截止日期，并且`处理程序`可以确定对主体来说哪部分会太慢。如果`ReadHeaderTimeout`为零，则使用`ReadTimeout`的值。如果两者均为零，则没有超时。
    */
    ReaderHeaderTimeout time.Duration   // Go 1.8

    /*
        `WriteTimeout`是响应的写入超时之前的最大持续时间。每当读取新请求的头部时，都会将其重置。 与`ReadTimeout`一样，它也不允许`处理程序`根据每个请求做出决策。
    */
    WriteTimeout time.Duration  

    /*
        `IdleTimeout`是启用保持活动状态后等待下一个请求的最长用时。如果`IdleTimeout`为零，则使用`ReadTimeout`的值。如果两者均为零，则没有超时。
    */
    IdleTimeout time.Duration   // Go 1.8

    /*
        `MaxHeaderBytes`控制服务器读取的最大字节数，以解析`请求头部`的键和值（包括`请求行`）。它不限制`请求主体`的大小。
        如果为零，则使用`DefaultMaxHeaderBytes`。
    */
    MaxHeaderBytes int

    /*
        `TLSNextProto`，可选地，指定函数，用于在`ALPN协议`升级发生后接管被提供的TLS连接的所有权。`map`的键，是协商的协议名称。`Handler`参数应用于处理HTTP请求，如果尚未设置，则将初始化`Request的TLS和RemoteAddr`。函数返回时，连接自动关闭。
        如果`TLSNextProto`不为空，则不会自动启用`HTTP/2支持`。
    */
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)   // Go 1.1

    /*
        `ConnState`指定一个可选的回调函数，当客户端连接更改状态时调用该函数。有关详细信息，请参见`ConnState`类型和关联的常量。
    */
    ConnState func(net.Conn, ConnState) // Go 1.3

    /*
        `ErrorLog`为接收连接的错误，来自处理程序的意外行为以及低层的`FileSystem`错误，指定了一个可选的记录器。
        如果为`nil`，则通过`log包的标准记录器`进行记录。
    */
    ErrorLog *log.Logger    // Go 1.3

    /*
        `BaseContext`，可选地，指定一个函数，该函数返回此服务器上传入的`请求的基础上下文`。
        提供的`Listener`是即将开始接受请求的指定监听器。
        如果`BaseContext`为`nil`，则默认值为`context.Background()`。如果为`非nil`，则它必须返回`非nil的上下文`。
    */
    BaseContext func(net.Listener) context.Context  // Go 1.13

    /*
        `ConnContext`，可选地，指定一个函数，该函数修改用于新连接`c`的上下文。提供的`ctx`派生自`基础上下文`，并且具有`ServerContextKey`值。
    */
    ConnContext func(ctx context.Context, c net.Conn) context.Context   // Go 1.13

    disableKeepAlives int32     // accessed atomically.
	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	doneChan   chan struct{}
	onShutdown []func()

}

```


## `Handler`接口

`Handler`响应`HTTP请求`。

`ServeHTTP`应将回复的`头部`和`数据`写入`ResponseWriter`，然后返回。返回会打信号表明请求已经完成；在`ServeHTTP`调用完成之后或与之并发地使用`ResponseWriter`或`从Request.Body中读取`都是无效的。

因为取决于`HTTP客户端软件`，`HTTP协议版本`以及`客户端和Go服务器之间的任何中间件`，可能无法在写入`ResponseWriter`之后从`Request.Body`中读取。谨慎的方式是，处理程序应先读取`Request.Body`，然后再进行回复。

除读取`请求体`外，处理程序不应修改提供的`Request`。

如果`ServeHTTP`出现`panic（宕机）`，则`服务器`（`ServeHTTP的调用者`）将假定`panic`的影响与`活动的请求`无关。它将从`panic`中`恢复（recover）`，将`堆栈跟踪`记录到`服务器错误日志`中，然后`关闭网络连接`或`发送HTTP/2 RST_STREAM`，具体取决于`HTTP协议`。`要中止处理程序`，以便客户端看到中断的响应，但服务器不会记录错误，可以使用带有`ErrAbortHandler`值的`panic（宕机）`。

```go

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

```


### `Request`结构体

`Request`表示由服务器接收或由客户端发送的`HTTP请求`。

该字段的语义在客户端和服务器使用情况之间略有不同。除了以下字段上的注释外，请参阅`Request.Write`和`RoundTripper`的文档。

```go

type Request struct {

    /*
        `Method`指定`HTTP方法`（`GET`，`POST`，`PUT`等）。
        对于`客户端请求`，`空字符串`表示`GET`。

        `Go的HTTP客户端`不支持使用`CONNECT方法`发送请求。有关详细信息，请参见`Transport`文档。
    */
    Method string

    /*
        `URL`指定`被请求的URI`（对于服务器请求）或`要访问的URL`（对于客户端请求）。

        对于服务器请求，将从`Request-Line（请求行）`（存储在`RequestURI`中）上提供的`URI`解析出`URL`。对于大多数请求，`Path`和`RawQuery`以外的其他字段将为空。（请参阅`RFC 7230`，第5.3节）

        对于客户端请求，`URL`的`Host`指定要连接的服务器，而`Request`的`Host`字段则可选地指定要在`HTTP请求`中发送的`Host`头部值。
    */
    URL *url.URL

    /*
        传入服务器的请求要使用的协议版本。

        对于客户请求，这些字段将被忽略。`HTTP客户端代码`始终使用`HTTP/1.1`或`HTTP/2`。
        有关详细信息，请参见`Transport`文档。
    */
    Proto string   // "HTTP/1.0"
    ProtoMajor int // 1
    ProtoMinor int // 0

    /*
        `Header`包含服务器接收或客户端发送的`请求头部字段`。

        如果服务器收到的请求带有如下`头部行`：

        Host: example.com
        accept-encoding: gzip, deflate
        Accept-Language: en-us
        fOO: Bar
        foo: two

        那么，`Header`的值为：

        Header = map[string][]string{
            "Accept-Encoding": {"gzip, deflate"},
            "Accept-Language": {"en-us"},
            "Foo": {"Bar", "two"},
        }

        对于传入的请求，`Host`头部会被传到`Request.Host`字段，并从`Header`的映射中删除。

        `HTTP`定义了`头部名称不区分大小写`。`请求解析器`通过使用`CanonicalHeaderKey`来让第一个字符以及连字符后紧跟的字符都变为大写，其余的都变为小写。

        对于客户端请求，某些头部（例如`Content-Length`和`Connection`）会在需要时自动写入，并且`Header`中的值可能会被忽略。请参阅文档中的`Request.Write`方法。
    */
    Header Header

    /*
        `Body`是请求的主体。

        对于客户端请求，`nil的主体`表示该请求没有主体，例如`GET请求`。`HTTP Client`的`Transport`负责调用`Close`方法。

        对于服务器请求，`Request.Body`始终为`非零`，但是当不存在主体时将立即返回`EOF`。
        `Server`将关闭请求体。不需要`ServeHTTP`处理程序来关闭请求体。
    */
    Body io.ReadCloser

    /*
        `GetBody`定义了一个可选的函数来返回`Body的新副本`。对于客户端请求，该字段用于当重定向需要多次读取请求体时。`GetBody`的使用仍然需要设置`Body`。

        对于服务器请求，不用该字段。
    */
    GetBody func() (io.ReadCloser, error) // Go 1.8

    /*
        `ContentLength`记录关联内容的长度。
        值`-1`，表示长度未知。
        值`>=0`，表示可以从`Body`中能读取到的字节数。

        对于客户请求，`非空Body`时该字段值为`0`，也会被视为未知长度。
    */
    ContentLength int64

    /*
        `TransferEncoding`列出从最外层到最内层的传输编码。空列表则表示`身份`编码（"identity" encoding）。
        通常可以忽略`TransferEncoding`。发送接收请求时，将根据需要自动添加和删除`分块编码（chunked encoding）`。
    */
    TransferEncoding []string

    /*
        `Close`用于指示，是在回复此请求之后（对于服务器），还是在发送此请求并读取其响应（对于客户端）之后关闭连接。

        对于服务器请求，`HTTP服务器`会自动处理此请求，并且`Handler`不需要此字段。

        对于客户端请求，设置此字段可防止在相同主机的请求之间重复使用`TCP连接`，就像设置了`Transport.DisableKeepAlives`一样。
    */
    Close bool

    /*
        对于服务器请求，`Host`指明在其上搜索URL的主机。对于`HTTP/1`（根据`RFC 7230`，第5.4节），这是`
        "Host"`头部的值或`URL本身`中提供的主机名。对于`HTTP/2`，它是`":authority"`伪头字段的值。
        它的形式可能是`"host:port"`。对于`国际域名`，`Host`可以采用`Punycode`或`Unicode`形式。如果需要，请使用`golang.org/x/net/idna`将其转换为这两种格式之一。
        为了防止`DNS重新绑定攻击`，服务器的处理程序应验证`"Host"`头部的值对自身来说是否有授权。随附的`ServeMux`会支持注册到特定主机名的`模式（pattern）`，从而保护其注册的`处理程序`。

        对于客户端请求，`Host`，可选地，覆盖要发送的`Host头部`。如果为空，则`Request.Write`方法使用`URL.Host`的值。`Host`可能包含一个`国际域名`。
    */
    Host string

    /*
        `Form`包含`已解析的表单数据`，包括`URL字段的查询参数`和`PATCH`，`POST`或`PUT`表单数据。

        该字段仅在调用`ParseForm`之后可用。

        `HTTP客户端`会忽略`Form`，而使用`Body`。
    */
    Form url.Values

    /*
        `PostForm`包含从`PATCH`，`POST`或`PUT`的`主体`参数中解析出的表单数据。

        该字段仅在调用`ParseForm`之后可用。

        `HTTP客户端`会忽略`PostForm`并改用`Body`。
    */
    PostForm url.Values

    /*
        `MultipartForm`是已解析的`多部分表单`，包括`文件上传`。

        仅在调用`ParseMultipartForm`之后，此字段才可用。

        `HTTP客户端`会忽略`MultipartForm`并改用`Body`。
    */
    MultipartForm *multipart.Form

    /*
        `Trailer`指定在`请求体`之后发送的`额外的头部字段`（可称为`尾部字段`）。

        对于服务器请求，`Trailer`映射最初仅包含`trailer`键，其值为`nil`。（客户端声明它将稍后发送的`尾部字段`。） 当处理程序从`Body`读取时，它不得引用`Trailer`。从`Body`读取返回`EOF`后，`Trailer`可以再次读取，并且包含`非nil值`（如果客户端发送了`尾部字段`的话）。

        对于客户端请求，必须将`Trailer`初始化为包含之后要发送的`尾部键`的映射。这些键的值可以为nil或它们的最终值。`ContentLength`必须为`0`或`-1`，以发送`分块的请求`。
        发送HTTP请求后，可以在读取`请求体`的同时更新`Trailer的映射值`。一旦`请求体`返回`EOF`，调用方就不得修改`Trailer`。

        很少有`HTTP客户端`，`服务器`或`代理`支持`HTTP尾部字段`。
    */
    Trailer Header

    /*
        `RemoteAddr`允许`HTTP服务器`和其他软件记录发来请求的网络地址，通常用于日志记录。`ReadRequest`不会填写此字段，并且没有被定义的格式。在调用处理程序之前，`net/http`包中的`HTTP服务器`会将`RemoteAddr`设置为`"IP:port"`地址。

        HTTP客户端将忽略此字段。
    */
    RemoteAddr string

    /*
        `RequestURI`是客户端发送到服务器的`请求行（Request-Line）`（`RFC 7230`，第3.1.1节）的未经修改的`请求目标`。通常应改用`URL字段`。

        在`HTTP客户端`请求中设置此字段是错误的。
    */
    RequestURI string

    /*
        `TLS`允许`HTTP服务器`和其他软件记录`TLS连接`的信息（在这些连接上有请求被接收）。`ReadRequest`不会填写此字段。

        `net/http`包中的`HTTP服务器`在调用处理程序之前，会为`启用TLS的连接`设置该字段。否则，该字段将为`nil`。

        `HTTP客户端`将忽略此字段。
    */
    TLS *tls.ConnectionState

    /*
        `Cancel`是一个可选通道，当其关闭时则表示客户的请求应被视为已取消。并非`RoundTripper`的所有实现都支持`Cancel`。

        对于服务器请求，此字段不适用。

        弃用的：使用`NewRequestWithContext`设置`Request`的上下文来替代。如果同时设置了`Request`的`Cancel`字段和上下文，则不确定是否应用`Cancel`。
    */
    Cancel <-chan struct{} // Go 1.5

    /*
        `Response`字段，指明导致创建此请求的`重定向响应`。

        仅在`客户端重定向期间`填充此字段。
    */
    Response *Response // Go 1.7

    /*
        `ctx`是`客户端上下文`或`服务器上下文`。它只能通过使用`WithContext`复制整个`Request`来进行修改。
        它不是导出的，以防止人们错误使用`Context`并更改由同一请求的调用者持有的上下文。
    */
    ctx context.Context

}

```


### `ResponseWriter`接口和`Response`结构体

`HTTP处理程序`使用`ResponseWriter`接口构造`HTTP响应`。

`Handler.ServeHTTP`方法返回后，`ResponseWriter`不能被使用。

```go

type ResponseWriter interface {

    /*
        `Header`返回将由`WriteHeader`发送的`头部映射`。`Header`映射也是`Handler`可以用来设置`HTTP尾部字段`的机制。

        除非修改后的头部要作为尾部字段，否则在调用`WriteHeader`（或`Write`）后对头部映射进行修改是无效的。

        设置`尾部字段`有两种方法。首选方法是在`头部`中预先声明您稍后将发送的`尾部字段`，方法是将`"Trailer"`头部的值设置为稍后将出现的`尾部字段的键`的名称。在这种情况下，头部映射中指明的键即将可能会被视为`尾部字段`。参见示例。第二种方法是，对于直到第一次`Write`之后才被`Handler`所知的尾随键，会在Header映射键之前加上`TrailerPrefix`常量值。请参阅`TrailerPrefix`。

        要禁止自动响应头（例如`"Date"`），请将其值设置为`nil`。
    */
    Header() Header

    /*
        `Write`将`数据`作为`HTTP回复`的一部分写入`连接`。

        如果尚未调用`WriteHeader`，则`Write`在写入数据之前会调用`WriteHeader(http.StatusOK)`。如果`Header`不包含`Content-Type`行，则`Write`会将值为`512`字节的`Content-Type`添加到`DetectContentType`中。此外，如果所有写入数据的总大小在几`KB`以下且没有`Flush`调用，则将自动添加`Content-Length`头部。

        根据`HTTP协议版本`和`客户端`的不同，调用`Write`或`WriteHeader`可能会阻止之后对`Request.Body`进行读取。对于`HTTP/1.x`请求，处理程序应在写入响应之前读取所有需要的请求体数据。刷新头部后（由于显式的`Flusher.Flush`调用或写入足够的数据以触发刷新），请求体可能会不可用。对于`HTTP/2`请求，`Go HTTP服务器`允许处理程序在同时写入响应的同时继续读取请求正文。但是，并非所有的`HTTP/2客户端`都支持这种行为。如果可能，处理程序应在写入之前先进行读取，以最大程度地实现兼容性。
    */
    Write([]byte) (int, error)

    /*
        `WriteHeader`发送带有`状态码`的`HTTP响应头部`。

        如果未显式调用`WriteHeader`，则对`Write`的第一次调用将触发一个隐式`WriteHeader(http.StatusOK)`。
        因此，对`WriteHeader`的显式调用主要用于发送`错误码`。

        提供的`状态码`必须是有效的`HTTP 1xx-5xx状态码`。只能写入一次头部。Go当前不支持发送`用户定义的1xx信息性的头部`，但`Server`会在读取`Request.Body`时自动发送的`100`状态码--继续响应头部。
    */
    WriteHeaders(statusCode int)

}

```

`Response`表示对`HTTP请求`的`响应`。

一旦收到`响应头`，`Client`和`Transport`将会返回来自服务器的`Response`。在读取`Body`字段时，将按需传输`响应体`。

```go

type Response struct {

    Status string       // e.g. "200 OK"
    StatusCode int      // e.g. 200
    Proto string        // e.g. "HTTP/1.0"
    ProtoMajor int      // e.g. 1
    ProtoMinor int      // e.g. 0

    /*
        `Header`将头部的键映射到值。如果响应中有多个具有相同键的头部，则可以使用`逗号`分隔符将它们连接在一起。 （`RFC 7230`第3.2.2节要求，多个头部在语义上等效于逗号分隔的序列。）当头部值被该结构中的其他字段（例如`ContentLength`，`TransferEncoding`，`Trailer`）复制时，该字段的值是权威性的。

        映射中的键已规范化（请参见`CanonicalHeaderKey`）。
    */
    Header Header

    /*
        `Body`代表`响应体`。

        在读取`Body`字段时，将按需传输`响应体`。如果网络连接失败或服务器终止了响应，则`Body.Read`的调用将返回`error`。

        http的`Client`和`Transport`保证了即使在没有主体的响应或主体长度为零的响应中，主体始终为`非零`。关闭`Body`是调用者的责任。如果主体没有读取完就关闭了，则`默认的HTTP客户端`的`Transport`可能不会重用`HTTP/1.x "keep-alive"的TCP连接`。

        如果服务器回复了`分块（chunked）`的`Transfer-Encoding`，则`Body`将自动分块。

        从`Go 1.12`开始，`Body`还将在成功的`"101 Switching Protocols"`响应（该响应会在`WebSocket`和`HTTP/2的"h2c"模式`里使用）上实现`io.Writer`。
    */
    Body io.ReadCloser

    /*
        `ContentLength`记录相关内容的长度。值`-1`表示长度未知。除非`Request.Method`为`"HEAD"`，否则值`>=0`表示可以从`Body`读取指定的字节数。
    */
    ContentLength int64

    /*
        包含从最外部到最内部的`传输编码`。值为nil时，表示使用`"identity"编码`。
    */
    TransferEncoding []string

    /*
        `Close`，用于标记在读取`Body`后，头部是否指示关闭连接。该值是给客户端的建议：`ReadResponse`和`Response.Write`都不要关闭连接。
    */
    Close bool

    /*
        `Uncompressed`，用于报告响应是否是压缩发送的，但已被`http包`解压缩。如果为true，则从`Body`读取将产生未压缩的内容，而不是来自服务器实际设置的压缩内容，`ContentLength`被设置为`-1`，并且从`responseHeader`中删除`"Content-Length"`和`"Content-Encoding"`字段。想要从服务器获取原始响应，请将`Transport.DisableCompression`设置为true。
    */
    Uncompressed bool

    /*
        `Trailer`将`尾部键`映射到与`Header`相同格式的值。

        `Trailer`最初仅包含`nil值`，服务器`"Trailer"`头部值中指定的每个键对应一个值。这些值不会添加到`Header`中。

        `Trailer`的访问不得与`Body`上的`Read`调用同时进行。

        在`Body.Read`返回`io.EOF`之后，`Trailer`将包含服务器发送的所有`尾部`值。
    */
    Trailer Header

    /*
        `Request`是指，为获取此响应而发送的请求。

        `Request`的`Body`为`nil`，则表示`请求体`已被消耗。

        该字段仅填充客户端请求。
    */
    Request *Request

    /*
        `TLS`包含有关在其上接收到响应的`TLS连接`的信息。对于`未加密的响应`，它为`nil`。

        指针在响应之间共享，不应修改。
    */
    TLS *tls.ConnectionState

}

```


### `ServeMux`结构体和`muxEntry`结构体

`ServeMux`是一个`HTTP请求多路复用器`。它根据`注册模式列表（a list of registered patterns）`将每个`传入的请求的URL`进行匹配，并调用`与URL最匹配`的`模式的处理程序（the handler for the pattern）`。

`模式（Pattern）`命名为`固定的、有根的路径`（例如`"/favicon.ico"`），或`有根的子树`（例如`"/images/"`）（请注意`尾部的斜杠`）。`较长的模式`优先于`较短的模式`，因此，如果同时分别为`"/images/"`和`"/images/thumbnails/"`注册了处理程序，则将为以`"/images/thumbnails/"`开头的路径调用后一个处理程序，将在`"/images/"`子树中接收对任何其他路径的请求。

请注意，由于`以斜杠结尾的模式`命名了一个`有根的子树`，因此`模式"/"`与`所有其他已注册模式不匹配的路径`匹配，而不仅仅是`Path=="/"`的URL。

如果已经注册了一个`子树`，并且接收到一个`指名了该子树的根而没有尾部斜杠的`请求，则`ServeMux`将该请求`重定向`到`该子树根（添加了尾部斜杠）`。可以`为该路径进行单独的注册`来覆盖此行为，而不必使用斜杠。例如，注册`"/images/"`会使`ServeMux`将对`"/images"`的请求重定向到`"/images/"`，除非已单独注册了`"/images"`。

`模式`可以用`主机名`开头，将匹配项只限制在该主机上。`特定于主机的模式`优先于`一般模式`，因此一个处理程序可以注册`"/codesearch"`和`"codesearch.google.com/"`这两种模式，而不必同时接收对`"http://www.google.com/"`的请求。

`ServeMux`还负责清理`URL请求路径`和`Host头部`，`清除端口号`以及`重定向任何包含的请求`。或`..`元素，或`重复的斜线`表示为等效的更清晰的URL。

```go

type ServeMux struct {
    mu sync.RWMutex
    m map[string]muxEntry
    es []muxEntry // slice of entries sorted from longest to shortest.
    hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
    h Handler
    pattern string
}

```

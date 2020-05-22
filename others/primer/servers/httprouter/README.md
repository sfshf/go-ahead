
## [httprouter: github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)


#### 一、简介

HttpRouter是一个用Go语言编写的轻量级高性能的HTTP请求路由器（或称作多路复用器--multiplexer/mux）。

与Go语言标准库中的`net/http`包下的默认多路复用器相比，HttpRouter的路由匹配模式支持变量并能匹配请求方法。

HttpRouter路由器具有高性能低内存的优化封装。它能很好地处理超长路径且超量的路由。它的高效匹配源于使用了压缩的动态查找树结构（radix tree）。


#### 二、特性

**唯一的明确匹配：** 其他的路由器，比如`http.ServeMux`，一个请求的URL路径可能匹配多个模式；因此，它们有些很尴尬的（笨拙的）模式优先级规则，例如最长匹配原则或最先注册最先匹配原则。而HttpRouter路由器的设计，一个请求只能准确地匹配一个路由或者空路由。因此，也没有意外的匹配，这使得它有利于搜索引擎优化（SEO）和改善用户体验。

**不用关心尾随斜杠：** 选择您喜欢的URL样式，如果缺少尾随斜杠或有多余的斜杠，路由器会自动重定向客户端。当然，只有当新路径有一个处理程序时，它才会这样做。如果你不喜欢，你可以[关掉这个行为](https://godoc.org/github.com/julienschmidt/httprouter#Router.RedirectTrailingSlash)。

**路径自动修正：** 除了无额外消耗地检查遗漏或多余的尾随斜杠外，HttpRouter路由器还可以修正错误的大小写和删除多余的路径元素（比如`../`或者`//`）。

**路由模式中的参数：** 停止解析请求的URL路径，只需给路径段命名，路由器就会将动态值传递给您。根据HttpRouter路由器的设计，路径参数消耗非常少。

**零垃圾：** 匹配和分派过程生成零字节的垃圾。唯一的堆分配是为路径参数构建键值对的切片，和构建新的上下文对象和请求对象（请求对象仅在标准`Handler/HandlerFunc`的API中）。在三参数的API中，如果请求路径不包含任何参数，则不需要进行单个堆分配。

**最优的性能：** [Benchmarks speak for themselves.](https://github.com/julienschmidt/go-http-routing-benchmark)

**不再服务器崩溃：** 您可以设置一个[Panic处理程序](https://godoc.org/github.com/julienschmidt/httprouter#Router.PanicHandler)来处理处理HTTP请求期间发生的Panic。然后路由器恢复并让PanicHandler记录发生的事情，并提供一个漂亮的错误页面。

**最优的API：** 路由器的设计支持构建合理的、分层的RESTful API。此外，它还内置了对[OPTIONS请求](http://zacstewart.com/2012/04/14/http-options-method.html)的原生支持以及可以应答`405 Method Not Allowed`。当然，您还可以自定义设置[`NotFound`](https://godoc.org/github.com/julienschmidt/httprouter#Router.NotFound)和[`MethodNotAllowed`](https://godoc.org/github.com/julienschmidt/httprouter#Router.MethodNotAllowed)处理器以及[供应静态文件](https://godoc.org/github.com/julienschmidt/httprouter#Router.ServeFiles)。


#### 三、基本用法

这里仅仅是简单的用法介绍，更多API细节查看[GoDoc](http://godoc.org/github.com/julienschmidt/httprouter)。

[简单示例代码](expls/usage)

**命名参数：** `命名参数`可以通过注册的`httprouter.Handle`类型函数的形参`httprouter.Params`访问，而`httprouter.Params`只是`httprouter.Param`的切片实例。可以通过参数在切片中的索引或使用`ByName()`方法来获取参数的值，例如：`:name`可以通过`ByName("name")`检索。

如果使用`http.Handler`（或者使用`router.Handler`，或者使用`http.HandlerFunc`）而不是使用`httprouter.Handle`时，命名参数存储在`request.Context`中，具体使用可参考[如何兼容`http.Handler` API](#五如何兼容httphandler-api)。

命名参数只匹配单个路径段：

```

模式: /user/:user

/user/gordon              匹配
/user/you                 匹配
/user/gordon/profile      不匹配
/user/                    不匹配

模式: /blog/:category/:post

/blog/go/request-routers            匹配: category="go", post="request-routers"
/blog/go/request-routers/           不匹配,但是路由器会重定向
/blog/go/                           不匹配
/blog/go/request-routers/comments   不匹配

```

`注意：`由于HttpRouter路由器采用唯一明确匹配的方式，因此`不能为同一路径段注册静态的路由和参数`。例如，不能同时为同一请求方法注册模式"/user/new"和"/user/:user"。`不同请求方法的路由是相互独立的`。

**全匹配参数：** 使用`*`符号的参数即全匹配参数，例如：`*name`。全匹配参数必须用于模式的结尾。

```

模式: /src/*filepath

 /src                      不匹配      路由器会重定向
 /src/                     匹配       filepath="/"
 /src/somefile.go          匹配       filepath="/somefile.go"
 /src/subdir/somefile.go   匹配       filepath="/subdir/somefile.go"

```


#### 四、HttpRouter多路复用器是如何工作的？

HttpRouter路由器所依赖的树结构充分利用了`公共前缀（common prefix）`，它基本上是一个紧凑的[前缀树（prefix tree）]()（或者仅仅是[根树（Radix tree）]()）。具有`公共前缀（common prefix）`的节点也共享一个公共的父节点。

下面是一个简单的GET请求方法的路由树示例：

```

优先级      路径              处理器
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>

```

每一个`*<num>`代表一个处理器函数（一个指针）的内存地址。如果沿着树从根到叶的，就会得到一个完整的路径，例如`blog\:post`，其中`:post`只是`实际post名称的占位符`（即：命名参数）。与散列映射不同，树结构还允许我们使用动态部分，如`:post`参数，因为我们实际上与路由模式匹配，而不仅仅是比较散列。正如[基准测试](https://github.com/julienschmidt/go-http-routing-benchmark)所显示的，这种方法非常有效。

由于URL路径具有层次结构，并且只使用有限的字符集（字节值），因此很可能有许多`公共前缀`。这样我们就可以轻松地将路由问题减少到更小的问题中。此外，路由器为每个请求方法管理一个单独的树。首先，它比在每个节点中持有一个`method->handle`映射的方式更节省空间。它还允许我们在开始查找前缀树之前大大减少路由问题。

为了获得更好的扩展性，每个树级别上的子节点都按优先级排序，其中优先级只是子节点（子节点、孙子节点等）中注册的处理器的数量。这有两方面的帮助：

1. 属于大多数路由路径的组成部分的节点会先被处理。这有助于使尽可能多的路由能够尽快处理。

2. 这是某种成本补偿。最长可达路径（成本最高）总是被最先处理。正如，下面的拟定表把树结构可视化；节点从上到下和从左到右进行计算。

```

├------------
├---------
├-----
├----
├--
├--
└-

```


#### 五、如何兼容`http.Handler` API？

HttpRouter是可以和标准包的`http.Handler`处理器接口兼容的。`httprouter.Router`本身就实现了`http.Handler`处理器接口。而且`httprouter.Router`还为`http.Handler`处理器接口和`http.HandlerFunc`处理器函数类型分别提供了方便的适配器[`Handler`](https://godoc.org/github.com/julienschmidt/httprouter#Router.Handler)和[`HandlerFunc`](https://godoc.org/github.com/julienschmidt/httprouter#Router.HandlerFunc)，在注册一个路由时与[`httprouter.Handle`](`httprouter.Handle`)效果一致。

`命名参数`可以通过`request.Context`来获取：

```go

func Hello(w http.ResponseWriter, r *http.Request) {
    params := httprouter.ParamsFromContext(r.Context())
    fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
}

```
作为选择，还可以使用`params := r.Context().Value(httprouter.ParamsKey)`替代助手函数`httprouter.ParamsFromContext()`。


#### 六、自动响应OPTIONS请求和跨域资源共享（CORS）

有时可能希望修改对OPTIONS请求的自动响应，例如，为了支持[CORS preflight requests](https://developer.mozilla.org/en-US/docs/Glossary/preflight_request)，或者设置其他的头部。这些都可以使用[`Router.GlobalOPTIONS`处理器](https://godoc.org/github.com/julienschmidt/httprouter#Router.GlobalOPTIONS)来达到。
```go

router.GlobalOPTIONS = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Access-Control-Request-Method") != "" {
        //Set CORS headers
        header := w.Header()
        header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
        header.Set("Access-Control-Allow-Origin", "*")
    }
    //Adjust status code to 204
    w.WriteHeader(http.StatusNoContent)
})

```

#### 七、如何兼容中间件？

`httprouter包`只是提供了一个非常有效的请求路由器和一些额外的功能。路由器只是一个`http.Handler`，您可以在路由器之前链接任何与`http.Handler`兼容的中间件，例如，[Gorilla handlers](http://www.gorillatoolkit.org/pkg/handlers)；或者还可以[自定义中间件](https://justinas.org/writing-http-middleware-in-go)。

再或者，可以使用基于`httprouter包`的[web框架](https://github.com/julienschmidt/httprouter#web-frameworks-based-on-httprouter)。


###### 多域/子域（Multi-domain / Sub-domains）

如果服务器需要服务多个域名或主机，或者需要使用子域名，可以采用每一个主机都定义一个路由的方式来解决。

[供参考的简例代码](expls/multidomain)


###### 基本身份验证（Basic Authentication）

[RFC 2617](https://www.ietf.org/rfc/rfc2617.txt)

[供参考的简例代码](expls/basicauth)


#### 八、如何链接NotFound处理器？

注意：可能需要将[`Router.HandleMethodNotAllowed`](https://godoc.org/github.com/julienschmidt/httprouter#Router.HandleMethodNotAllowed)设置为`false`以避免出现问题。

您可以使用另一个`http.Handler`，例如，另一个路由器，通过使用`router.NotFound`处理器来处理此路由器无法匹配的请求；而且还可以采用处理链方式。


###### 静态文件

`NotFound`处理器可用于从根路径`/`提供静态文件（比如：index.html文件及其他资源）：

```go

// Serve static files from the ./public directory.
router.NotFound = http.FileServer(http.Dir("public"))

```

但这种方法为了避免路由问题而偏离了路由器的严格的核心规则。更干净的方法是使用清楚明确的子路径来提供文件，比如`/static/*filepath`或`/files/*filepath`。


#### 九、[HttpRouter包内API说明](apis/HttpRouter.md)

#### 十、[作为比较：net/http标准包内API说明](apis/ServeMux.md)

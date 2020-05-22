
# What is Cookie?

- [Go语言标准包`net/http`的文档](https://golang.google.cn/pkg/net/http/#Cookie)

- 更多信息，请查看[HTTP State Management Mechanism](https://tools.ietf.org/html/rfc6265)


## `net/http`包对Cookie的实现

`Cookie`，是指`在HTTP响应的Set-Cookie头部`或`HTTP请求的Cookie头部`中发送的`HTTP Cookie`。


### `Cookie`

```go

type Cookie struct {
    Name string
    Value string

    Path string             // optional
    Domain string           // optional
    Expires time.Time       // optional
    RawExpires string       // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
    // MaxAge>0 means Max-Age attribute present and given in seconds.
    MaxAge int
    Secure bool
    HttpOnly bool
    SameSite SameSite       // Go 1.11
    Raw string
    Unparsed []string       // Raw text of unparsed attribute-value pairs
}

```


### `SameSite`

从`Go v1.11`开始，`Cookie`结构体有了`SameSite`字段，允许服务器定义cookie属性，从而使浏览器`无法`将cookie与跨站点请求一起发送。主要目标是减轻`跨域信息泄漏`的风险，并提供针对`跨站点请求伪造攻击`的某种保护。其定义如下：

```go

type SameSite int

const (
    SameSiteDefaultMode SameSite = iota + 1
    SameSiteLaxMode
    SameSiteStrictMode
    SameSiteNoneMode
)

```
有关`SameSite`详细阐述，请查看[RFC文档](https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00)。


### 有关`Cookie`的函数和方法

#### 1、`func (c *Cookie) String() string { ... }`

`String()`返回`cookie的序列化`，以在`Cookie头部`中使用（如果仅设置了`Name`和`Value`）或`Set-Cookie`响应头部（如果设置了其他字段）。

如果`c`为`nil`或`c.Name`无效，则返回`空字符串`。


#### 2、`SetCookie(w ResponseWriter, cookie *Cookie) { ... }`

`SetCookie`将`Set-Cookie头部`添加到被提供的`ResponseWriter的头部`中。

被提供的cookie`必须具有有效的名称`。`无效的cookie可能会被静默删除`。


#### 3、`func (r *Request) Cookie(name string) (*Cookie, error) { ... }`

`Cookie`返回`请求中提供的命名cookie`，如果找不到，则返回`ErrNoCookie`。如果多个Cookie与给定名称匹配，则`仅返回一个Cookie`。


#### 4、`func (r *Request) Cookies() []*Cookie { ... }`

`Cookies`解析并返回与请求一起发送的`HTTP cookie`。


### 有关`Cookie`的报错

#### 1、`var ErrNoCookie = errors.New("http: named cookie not present")`

未找到`Cookie`时，`Request`的`Cookie方法`将返回`ErrNoCookie`。


### `CookieJar`接口

`CookieJar`管理`HTTP请求中cookie`的`存储`和`使用`。

`CookieJar的实现必须安全，可以被多个goroutine并发使用`。

`net/http/cookiejar`包提供了`CookieJar`实现。

```go

type CookieJar interface {

    /*
        `SetCookies`在给定URL的回复中处理cookie的接收。它可能会或可能不会选择保存cookie，具体取决于`jar的策略和实现`。
    */
    SetCookies(u *url.URL, cookies []*Cookie)

    /*
        `Cookies`返回在给定URL的请求中发送的cookie。具体实现取决于遵守的`标准的cookie使用限制`，例如`RFC 6265`中的限制。
    */
    Cookies(u *url.URL) []*Cookie

}

```


#### `net/http/cookiejar`包

`cookiejar`包，实现了与`RFC 6265`兼容的在内存里的`http.CookieJar`。


## Gin框架对`net/http`包中Cookie功能的封装

### 1、`func (c *Context) Cookie(name string) (string, error) { ... }`

`Cookie`返回`请求中提供的命名cookie`，如果找不到，则返回`ErrNoCookie`。并返回`未转义的命名cookie`。如果多个Cookie与给定名称匹配，则`仅返回一个Cookie`。


### 2、`func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) { ... }`

`SetCookie`将`Set-Cookie头部`添加到`ResponseWriter的头部`中。被提供的cookie必须具有`有效的名称`。`无效的cookie可能会被静默删除`。


### 3、`func (c *Context) SetSameSite(samesite http.SameSite) { ... }`

`SetSameSite`用来设置cookie的`SameSite`字段。

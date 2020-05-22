**核心API:**

`Router -- 路由器/多路复用器:`

```go

/*
Router是一个可以通过可配置的路由来分发请求给不同处理器函数http.Handler实现。
*/
type Router struct {

    trees map[string]*node

	paramsPool sync.Pool
	maxParams  uint16

    /*
    如果当前路径不能匹配，但（没）有尾随斜杠的路径存在处理器，则启用自动重定向。
    例如，如果请求"/foo/"路径，但仅存在用于"/foo"的路由，则客户端会针对"GET请求"用"HTTP状态代码301"重定向到"/foo"，而对于所有的其他请求方法则使用"HTTP状态代码308"。
    */
    RedirectTrailingSlash bool

    /*
    如果启用，路由器将尝试修复当前请求路径（如果没有为该路径注册处理器）。首先，删除多余的路径元素，如"../"或"//"。之后，路由器对清理过的路径进行不区分大小写的查找。如果可以找到此路由的处理器，路由器将重定向到正确的路径，状态代码301用于GET请求，308用于所有其他请求方法。例如"/FOO"和"/..//FOO"可以重定向到"/foo"。`RedirectTrailingSlash`与此选项无关。
    */
    RedirectFixedPath bool

    /*
    如果启用，当路由器无法路由当前请求，路由器将检查当前路由是否允许使用其他方法；如果存在其他处理器方法，则用“Method Not Allowed”和HTTP状态代码405响应请求；如果不存在其他方法，则将请求委托给NotFound处理程序。
    */
    HandleMethodNotAllowed bool

    /*
    如果启用，路由器会自动回复OPTIONS请求。自定义OPTIONS处理器优先于自动答复。
    */
    HandleOPTIONS bool

    /*
    自动处理OPTIONS请求时要调用的可选的"http.Handler"。仅当"HandleOPTIONS"为"true"且未设置特定路径的OPTIONS处理器时，才会调用该处理器。在调用该处理器之前要设置“Allowed”头部。
    */
    GlobalOPTIONS http.Handler

    // Cached value of global (*) allowed methods
	globalAllowed string

    /*
    当找不到匹配路由时调用的可配置"http.Handler"。如果未设置，则使用"http.NotFound"。
    */
    NotFound http.Handler

    /*
    可配置的"http.Handler"，当请求无法路由且"HandleMethodNotAllowed"为"true"时调用。如果未设置，则调用传入"http.StatusMethodNotAllowed"参数的"http.Error"处理函数。在处理器被调用之前，要将“Allow”头部设置为所允许请求的方法。
    */
    MethodNotAllowed http.Handler

    /*
    用于处理从http处理器中捕获的宕机的函数。它应该用于生成错误页并返回http错误代码500（内部服务器错误）。处理程序可用于防止服务器因未捕获宕机而崩溃。
    */
    PanicHandler func(http.ResponseWriter, *http.Request, interface{})

}

```

`Param -- http请求路径的命名参数类型:`

```go

/*
Param，指单个的由一个键一个值组成的URL参数。
*/
type Param struct {
    Key   string
    Value string
}

```

`Params -- http请求路径的命名参数slice类型:`

```go
/*
Params，是由路由器返回的Param的一个切片。这个切片是有序的，第一个URL参数也就是切片的第一个值；因此，用索引取值也是安全的。
*/
type Params []Param

```

`Handle -- http请求的处理器函数类型:`

```go

/*
Handle，指一个可以注册到一个路由去处理HTTP请求的处理器函数。有点像"http.HandlerFunc"，但是有代表通配符的值（即：路径变量）的第三个参数。
*/
type Handle func(http.ResponseWriter, *http.Request, Params)

```

**其他API:**

`CleanPath -- :`

```go

func CleanPath(p string) string

```

`ParamsKey -- :`

```go

var ParamsKey = paramsKey{}

```

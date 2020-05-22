
**核心API:**

`Engine -- Gin框架的框架结构体`

```go

/*
Engine定义为Gin框架的实例。它包含了多路复用器（路由器），中间件和配置设置。
可以使用New()函数或者Default()函数来创建一个Engine实例。
*/
type Engine struct {

    RouterGroup

	RedirectTrailingSlash bool

	RedirectFixedPath bool

	HandleMethodNotAllowed bool

	ForwardedByClientIP    bool

	AppEngine bool

	UseRawPath bool

	UnescapePathValues bool

	MaxMultipartMemory int64

	RemoveExtraSlash bool

	delims           render.Delims

	secureJsonPrefix string

	HTMLRender       render.HTMLRender

	FuncMap          template.FuncMap

	allNoRoute       HandlersChain

	allNoMethod      HandlersChain

	noRoute          HandlersChain

	noMethod         HandlersChain

	pool             sync.Pool

	trees            methodTrees
}

```

`IRouter -- 路由接口`

```go

/*
IRouter接口，定义了所有路由的处理接口，包括单个的路由和组路由。
*/
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

```

`IRoutes -- 路由接口`

```go

/*
IRoutes接口，定义了所有路由的处理接口。
*/
type IRoutes interface {

	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes

}

```

`RouterGroup -- 路由组`

```go

/*
RouterGroup结构体，实现了IRouter和IRoutes两个接口，用于内部配置路由。一个路由组涉及一个前缀和一个处理器（或中间件）数组。   
*/
type RouterGroup struct {

    Handlers HandlersChain

	basePath string

    engine   *Engine

    root     bool
}

```

`HandlersChain -- 路由处理器链`

```go

/*
HandlersChain定义了一个HandlerFunc的切片。
*/
type HandlersChain []HandlerFunc

```

`HandlerFunc -- 路由处理器函数类型`

```go

/*
HandlerFunc定义了Gin中间件所要使用的处理器函数。
*/
type HandlerFunc func(*Context)

```

`Context -- 上下文`

```go

/*
Context结构体是Gin框架中最重要的部分，作用巨大。例如，它使得我们可以在中间件之间传递变量；管理流量；验证一个请求的JSON和渲染出一个JSON响应。
*/
type Context struct {

	writermem responseWriter

	Request   *http.Request

	Writer    ResponseWriter

	Params   Params

	handlers HandlersChain

	index    int8

	fullPath string

	engine *Engine

	Keys map[string]interface{}

	Errors errorMsgs

	Accepted []string

	queryCache url.Values

	formCache url.Values
}

```


**其他API:**


# Gin框架自带的`日志中间件`

**func Logger() HandlerFunc**

`Logger()`实例化一个`Logger中间件`，该中间件会将日志写入`gin.DefaultWriter`。默认情况下，`gin.DefaultWriter = os.Stdout`。


**func LoggerWithConfig(conf LoggerConfig) HandlerFunc**

`LoggerWithConfig`实例化一个`使用了指定配置`的`Logger中间件`。


**func LoggerWithFormatter(f LogFormatter) HandlerFunc**

`LoggerWithFormatter`实例化`具有指定日志格式函数`的`Logger中间件`。


**func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc**

`LoggerWithWriter`实例化一个`具有指定写入器缓冲区`的`Logger中间件`。例如：`os.Stdout`、以写模式打开的文件、套接字......


## `LoggerConfig`、`LogFormatter`、`LogFormatterParams`

### `LoggerConfig`定义`Logger中间件`的`配置`。

```go

type LoggerConfig struct {

    // Optional. Default value is gin.defaultLogFormatter
    Formatter LogFormatter

    // Output is a writer where logs are written.
    // Optional. Default value is gin.DefaultWriter.
    Output io.Writer

    // SkipPaths is a url path array which logs are not written.
    SkipPaths []string

}

```


### `LogFormatter`是传递给`LoggerWithFormatter`的`格式器函数`的`签名`。

```go

type LogFormatter func(params LogFormatterParams) string

```


#### `LogFormatterParams`是在记录时间到来时任何`格式器`都将要处理的结构。

```go

type LogFormatterParams struct {
    Request *http.Request

    // `TimeStamp`展示了服务器返回一个响应后的时间
    TimeStamp time.Time
    // `StatusCode`是HTTP响应码。
    StatusCode int
    // `Latency`是指服务器处理某个请求的花费的多少时间
    Latency time.Duration
    // `ClientIP`，等于`Context`的`ClientIP`方法。
    ClientIP string
    // `Method`是请求的HTTP方法
    Method string
    // `Path`是客户请求的路径。
    Path string
    // 如果在处理请求过程中出现报错，`ErrorMessage`会被设置。
    ErrorMessage string

    // `BodySize`，是响应主体的大小。
    BodySize int
    // `Keys`，是设置在请求上下文中的键。
    Keys map[string]interface{}

    // `isTerm`显示，gin的输出描述符是否指向终端。
    isTerm bool

}

```

##### `LogFormatterParams`的方法

**func (p *LogFormatterParams) StatusCodeColor() string**
**func (p *LogFormatterParams) MethodColor() string**
**func (p *LogFormatterParams) ResetColor() string**
**func (p *LogFormatterParams) IsOutputColor() bool**


#### `defaultLogFormatter`是`Logger中间件`使用的`默认日志格式函数`。

```go

var defaultLogFormatter = func(param LogFormatterParams) string {
    ... ...
}

```

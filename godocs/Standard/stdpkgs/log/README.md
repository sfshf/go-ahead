
# `log`包

`log`包，实现了一个`简单的`日志记录程序包。它定义了`Logger`类型，并带有`格式化输出`的方法。它还有可通过帮助函数`Print[f|ln]`，`Fatal[f|ln]`和`Panic[f|ln]`访问的预定义`标准记录器`，比手动创建记录器更容易使用。该记录器将写入`标准错误流`，并打印每条被记录的消息的日期和时间。`每条日志消息都在单独的行上输出：如果要打印的消息未以换行符结尾，则记录器将添加一个换行符`。`Fatal函数`在写入日志消息后调用`os.Exit(1)`。`Panic函数`会在写入日志消息后调用`panic`。


## 核心代码解析

### 1、`Logger`结构体

`Logger（记录器）`表示一个活动的记录对象，该对象会产生`成行的输出`到`io.Writer`。每个记录操作都会调用`Writer`的`Write`方法。一个`Logger`可以从多个goroutines中同时使用；它保证对`Writer`的序列化访问。

```go

type Logger struct {
    mu sync.Mutex   // ensures atomic writes; protects the following fields
    prefix string   // prefix on each line to identify the logger (but see Lmsgprefix)
    flag int        // properties
    out io.Writer   // destination for output
    buf []byte      // for accumulating text to write
}

```

### 2、`标识`常量


```go

const (
    Ldate       = 1 << iota     // the date in the local time zone: 2009/01/23
    Ltime                       // the time in the local time zone: 01:23:23
    Lmicroseconds               // microsecond resolution: 01:23:23.123123. assumes Ltime.
    Llongfile                   // full file name and line number: /a/b/c/d.go:23
    Lshortfile                  // final file name element and line number: d.go:23. overrides Llongfile.
    LUTC                        // if Ldate or Ltime is set, use UTC rather than the local time zone.
    Lmsgprefix                  // move the "prefix" from the beginning of the line to before the message.
    LstdFlags   = Ldate | Ltime // initial values for the standard logger.
)

```

`标识`定义了`Logger`生成的每个日志条目之前添加哪些文本。`位（Bits）`组合在一起以控制所打印的内容。除`Lmsgprefix`标识外，无法控制它们的显示顺序（此处列出的顺序）或它们显示的格式（如注释中所述）。仅当`Llongfile`或`Lshortfile`被指定时，前缀后才带有冒号。例如，标识`Ldate | Ltime (or LstdFlags)`产生：

```

2009/01/23 01:23:23 message

```

而标识`Ldate | Ltime | Lmicroseconds | Llongfile`产生：

```

2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message

```

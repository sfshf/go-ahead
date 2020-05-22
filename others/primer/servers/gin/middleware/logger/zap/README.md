
# [Zap](https://github.com/uber-go/zap)

用Go语言实现的，极速的、结构化的、分级的日志记录工具。


## Overview（概述）

`zap`包提供了极速的、结构化的、分级的日志记录工具。

由于使用`热路径进行记录`，`基于反射的序列化`和`字符串格式化`的应用程序成本过高，它们非常占用CPU且产生许多小的内存分配。换句话说就是，使用了`json.Marshal`和`fmt.Fprintf`来记录大量的`interface{}`，导致您的应用程序运行缓慢。

`Zap`采用了不同的方法。它包含一个`无反射`，`零分配`的`JSON编码器`，并且`基本的Logger尽力避免序列化开销和分配`；通过在该基础上构建`高级别`的`SugaredLogger`，`zap`让用户自行选择何时需要计算每个分配，以及何时需要更熟悉的，松散类型的API。


## Choosing a Logger（选择记录器）

`在性能要求不高的情况下`，请使用`SugaredLogger`；它比其他结构化日志记录包快4-10倍，并且支持`结构化的`和`printf样式的`日志记录功能。与`log15`和`go-kit`相同的是，`SugaredLogger`的结构化的日志记录`API都是松散类型的`，并且可以接收`可变数量的`键值对。（对于更高级的用例，它们还接受`强类型字段` -- 更多详细信息，请参见`SugaredLogger.With`的文档。）

```go

sugar := zap.NewExample().Sugar()
defer sugar.Sync()
sugar.Infow("failed to fetch URL",
       "url", "http://example.com",
       "attempt", 3,
       "backoff", time.Second)
sugar.Infof("failed to fetch URL: %s", "http://example.com")

```

`默认情况下，记录器是无缓冲的`。但是由于`zap`的`低级别API`允许缓冲，因此在退出进程之前调用`Sync`是一个好习惯。

`在每毫秒和每一次分配都很重要的罕见情况下`，请使用`Logger`；它比`SugaredLogger`更快，并且分配的资源少得多，`但它仅支持强类型的结构化日志`。

```go

logger := zap.NewExample()
defer logger.Sync()
logger.Info("failed to fetch URL",
    zap.String("url", "http://example.com"),
    zap.Int("attempt", 3),
    zap.Duration("backoff", time.Second))

```

在`Logger`和`SugaredLogger`之间进行的选择没必要成为整个应用程序的决定：在两者之间进行转换既简单又便宜。

```go

logger := zap.NewExample()
defer logger.Sync()
sugar := logger.Sugar()
plain := sugar.Desugar()

```


## Configuring Zap（配置`Zap`）

生成`Logger`的最简单方式是，使用`zap`自带的预设：`NewExample`，`NewProduction`和`NewDevelopment`。这些预设通过一个函数调用来构建记录器：

```go

logger, err := zap.NewProduction()
if err != nil {
    log.Fatalf("can't initialize zap logger: %v", err)
}
defer looger.Sync()

```

`这些预设适用于小型项目，而大型项目和组织自然需要更多的自定义设置`。对于大多数用户来说，`zap`的`Config`结构在灵活性和便利性之间取得了适当的平衡。示例代码，请参阅`包级别的BasicConfiguration`示例。

可以进行更特殊的配置（`在文件之间拆分输出`，`将日志发送到消息队列`等），但是需要直接使用`go.uber.org/zap/zapcore`包。示例代码，请参阅`包级别的AdvancedConfiguration`示例。


## Extending Zap（扩展`Zap`）

`zap`包本身是对`go.uber.org/zap/zapcore`包中接口的相当薄的一层包装。想要扩展`zap`来支持新的编码（例如`BSON`），新的日志接收器（例如`Kafka`）或其他奇特的东西（也许是`异常聚集服务`，例如`Sentry`或`Rollbar`），通常需要实现`zapcore.Encoder`，`zapcore.WriteSyncer `或`zapcore.Core`接口。有关详细信息，请参见`zapcore`文档。

同样地，包作者可以使用`zapcore`包中的高性能`Encoder`和`Core`实现来构建自己的记录器。


## Frequently Asked Questions（常被问及的问题）

从[https://github.com/uber-go/zap/blob/master/FAQ.md](https://github.com/uber-go/zap/blob/master/FAQ.md)，可获得涵盖从安装报错到设计决策的所有内容的FAQ。



## `go.uber.org/zap/zapcore`包API说明

## `go.uber.org/zap`包API说明

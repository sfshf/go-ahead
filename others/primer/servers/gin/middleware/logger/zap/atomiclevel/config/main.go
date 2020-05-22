package main

import (
    "encoding/json"

    "go.uber.org/zap"
)

func main() {

/*

    type Config struct {
        ... ...
    }

    `Config`提供了一种声明式的方式来构造记录器。它不能做任何连`New`，`Options`和各种`zapcore.WriteSyncer`和`zapcore.Core`包装器都无法完成的工作，但这是切换常见选项的一种更简单的方法。
    请注意，`Config`有意仅支持最常见的选项。可以进行更多不寻常的日志记录设置（记录到网络连接或消息队列，在多个文件之间拆分输出等），但是需要直接使用`zapcore`包。有关示例代码，请参见`包级别的BasicConfiguration`和`AdvancedConfiguration`示例。

    The zap.Config struct includes an AtomicLevel.
    To use it, keep a reference to the Config.
*/

    rawJSON := []byte(`{
        "level": "info",
        "outputPaths": ["stdout"],
        "errorOutputPaths": ["stderr"],
        "encoding": "json",
        "encoderConfig": {
            "messageKey": "message",
            "levelKey": "level",
            "levelEncoder": "lowercase"
        }
    }`)

    var cfg zap.Config
    if err := json.Unmarshal(rawJSON, &cfg); err != nil {
        panic(err)
    }

    logger, err := cfg.Build()
    if err != nil {
        panic(err)
    }
    defer logger.Sync()

    logger.Info("info logging enabled")

    cfg.Level.SetLevel(zap.ErrorLevel)
    logger.Info("info logging disabled")

}

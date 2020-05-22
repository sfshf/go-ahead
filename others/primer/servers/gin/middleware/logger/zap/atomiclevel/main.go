package main

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {

/*

    type AtomicLevel struct {
        // contains filtered or unexported fields.
    }

    `AtomicLevel`是原子可更改的动态日志记录级别。
    它使您可以在运行时安全地更改记录器树（`根记录器`和通过添加上下文创建的任何`子记录器`）的日志级别。

    `AtomicLevel`本身是一个`http.Handler`，它提供`JSON终端`来更改其级别。

    必须使用`NewAtomicLevel`构造函数来创建`AtomicLevel`才能分配其内部原子指针。

*/

    atom := zap.NewAtomicLevel()

    // To keep the example deterministic, disable timestamps in the output.
    encoderCfg := zap.NewProductionEncoderConfig()
    encoderCfg.TimeKey = ""

    logger := zap.New(zapcore.NewCore(
                zapcore.NewJSONEncoder(encoderCfg),
                zapcore.Lock(os.Stdout),
                atom))
    defer logger.Sync()

    logger.Info("info logging enabled")

    atom.SetLevel(zap.ErrorLevel)
    logger.Info("info logging disabled")

}

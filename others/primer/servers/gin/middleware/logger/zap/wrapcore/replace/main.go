package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {

/*
    `WrapCore`包装或替换`Logger`底层的`zapcore.Core`。
*/

    // Replacing a Logger's core can alter fundamental behaviors.
    // For example, it can convert a Logger to a no-op.
    nop := zap.WrapCore(func(zapcore.Core) zapcore.Core {
        return zapcore.NewNopCore()
    })

    logger := zap.NewExample()
    defer logger.Sync()

    logger.Info("working")
    logger.WithOptions(nop).Info("no-op")
    logger.Info("original logger still works")

}

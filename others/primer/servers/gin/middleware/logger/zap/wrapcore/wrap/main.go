package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {

    // Wrapping a Logger's core can extend its functionality.
    // As a trivial example, it can double-write all logs.
    doubled := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
        return zapcore.NewTee(c, c)
    })

    logger := zap.NewExample()
    defer logger.Sync()

    logger.Info("single")
    logger.WithOptions(doubled).Info("doubled")

}

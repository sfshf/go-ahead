package main

import (
    "go.uber.org/zap"
)

func main() {

/*
    `NewStdLog`返回一个`*log.Logger`，它将以`InfoLevel`写入到被提供的`zap Logger`。
    要重定向标准库的包级别的日志记录功能，请改用`RedirectStdLog`。
*/

    logger := zap.NewExample()
    defer logger.Sync()

    std := zap.NewStdLog(logger)
    std.Print("standard logger wrapper")

}

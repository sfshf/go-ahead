package main

import (
    "log"

    "go.uber.org/zap"
)

func main() {

/*
    `RedirectStdLog`将输出从标准库的包级别的记录器以`InfoLevel`级别重定向到被提供的记录器。
    由于zap已经处理了调用者注释，时间戳等，因此它将自动禁用标准库的注释和前缀。

    它返回一个函数来恢复原始前缀和标志，并将标准库的输出重置为`os.Stderr`。
*/

    log.Print("standard library original print format.\n")

    logger := zap.NewExample()
    defer logger.Sync()

    undo := zap.RedirectStdLog(logger)
    defer undo()

    log.Print("redirected standard library")

}

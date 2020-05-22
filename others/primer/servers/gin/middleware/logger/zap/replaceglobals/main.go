package main

import (
    "go.uber.org/zap"
)

func main() {

/*
    `ReplaceGlobals`替换全局`Logger`和`SugaredLogger`，并返回一个恢复原始值的函数。并发使用是安全的。
*/

    logger := zap.NewExample()
    defer logger.Sync()

    undo := zap.ReplaceGlobals(logger)
    defer undo()

    zap.L().Info("replaced zap's global loggers")

}

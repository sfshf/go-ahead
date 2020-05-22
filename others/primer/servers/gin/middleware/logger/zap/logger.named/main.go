package main

import (
    "go.uber.org/zap"
)

func main() {

/*
    `Named`将新的`路径段`添加到记录器的名称。`段`用句点（`.`）来连接。默认情况下，`Logger`未被命名。
*/

    logger := zap.NewExample()
    defer logger.Sync()

    // By default, Loggers are unnamed.
    logger.Info("no name")

    // The first call to Named sets the Logger name.
    main := logger.Named("main")
    main.Info("main logger")

    // Additional calls to Named create a period-separated path.
    main.Named("subpackage").Info("sub-logger")

}

package main

import (
    "go.uber.org/zap"
)

func main() {

/*
    如果启用了在指定级别记录消息，则`Check`返回`CheckedEntry`。
    这是完全可选的优化；在高性能应用程序中，`Check`可以帮助避免分配切片来容纳字段。
*/

    logger := zap.NewExample()
    defer logger.Sync()

    if ce := logger.Check(zap.DebugLevel, "debugging"); ce != nil {
        /*
            If debug-level log output isn't enabled or if zap's sampling would have
            dropped this log entry, we don't allocate the slice that holds these fields.
        */
        ce.Write(
            zap.String("foo", "bar"),
            zap.String("baz", "quux"),
        )
    }

}

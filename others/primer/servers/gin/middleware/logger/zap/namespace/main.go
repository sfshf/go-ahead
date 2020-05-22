package main

import (
    "go.uber.org/zap"
)

func main() {

/*
    `Namespace`在记录器的上下文中创建一个命名的隔离的`域`。所有后续字段都将添加到新的名称空间里。

    当将记录器注入子组件或第三方库时，这有助于防止键位冲突。
*/

    logger := zap.NewExample()
    defer logger.Sync()

    logger.With(
        zap.Namespace("metrics"),
        zap.Int("counter", 1),
    ).Info("tracked some metrics")


}

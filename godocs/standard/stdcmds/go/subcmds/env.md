
[Print Go environment information](https://golang.google.cn/cmd/go/#hdr-Print_Go_environment_information)

##### 一、用法

```

go env [-json] [-u] [-w] [var ...]

```

##### 二、env命令说明

`env命令`打印Go环境信息。

默认情况下，`env命令`以`shell脚本`的形式打印信息(在Windows上是批处理文件)。如果提供一个或多个环境变量名作为参数，则`env命令`将逐行打印每个指名的环境变量的值。

`-json`标志将以`JSON`格式替代`shell脚本`格式打印环境变量。

`-w`标志需要一个或多个格式为`NAME=VALUE`的参数，并且会将指名的环境变量的默认值改为给定的值。

`-u`标志需要一个或多个参数，并且会取消指名的环境变量的默认值设置，而改用`go env -w`设置的值。

使用`go help environment`查看更多关于环境变量的信息。


[The main module and the build list](https://golang.google.cn/cmd/go/#hdr-The_main_module_and_the_build_list)


##### 说明：

`"main module"`是包含用于运行`go命令`的目录的模块。`go命令`通过在`当前目录`中查找`go.mod`文件来找到模块根目录，或者在`当前目录的父目录`中，或者在`父目录的父目录`中，依此类推。

`主模块`的`go.mod`文件通过`requir`e、`replace`和`exclude`语句，定义了用于`go命令`的包的精确集合。依赖项模块（通过`require`语句找到）也有助于包集的定义，但只是通过它们的`go.mod`文件的`require`语句：依赖模块中的任何`replace`和`exclude`语句都将被忽略。因此，`replace`和`exclude`语句允许`主模块`完全控制自己的构建，而不受依赖项的完全控制。

为构建提供包的模块集称为`build list（构建列表）`。`构建列表`最初只包含`主模块`。然后`go命令`递归地将列表中已经存在的模块所需的模块版本添加到列表中，直到列表中没有需要添加的内容。如果某个特定模块的多个版本被添加到列表中，那么在最后只保留最新版本(根据语义版本顺序)，以便在构建中使用。

`go list`命令提供有关主模块和构建列表的信息。例如:

```

go list -m              # print path of main module
go list -m -f={{.Dir}}  # print root directory of main module
go list -m all          # print build list

```

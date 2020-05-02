
[Defining a module](https://golang.google.cn/cmd/go/#hdr-Defining_a_module)


##### 说明：

一个模块由一个Go源文件树定义，且在该源文件树的根目录中含有一个`go.mod`文件。包含go的目录。含有`go.mod`文件的目录被称为`模块根目录`。通常，`模块根目录`也对应于`源代码存储库根目录`(但通常不需要)。模块是模块根目录及其子目录中所有Go包的集合，但排除了具有自己的`go.mod`文件的子树。

`"module path"`是对应于模块根目录的导入路径前缀。`go.mod`文件定义了模块路径，并列出了其他模块的指定版本，其他模块是指在构建过程中解析导入项时可能被用到的模块，其他模块是通过它们的模块路径和版本号指明的。

举例：下面的`go.mod`文件声明了一个路径为`example.com/m`的模块的根目录，并且它还声明了依赖的特定版本的`golang.org/x/text`模块和`gopkg.in/yaml.v2`模块。

```

module example.com/m

require (
    golang.org/x/text v0.3.0
    gopkg.in/yaml v2 v2.1.0
)

```

`go.mod`文件还可以指定替换和排除的版本，只适用于直接构建模块；当模块被合并到更大的构建中时，它们将被忽略。了解更多关于`go.mod`文件的信息，见`go help go.mod`。

要启动一个新模块，只需在模块目录树的根目录下创建一个`go.mod`文件，`go.mod`文件里只包含一个模块语句。`go mod init`命令可以用来做这个：
`go mod init example.com/m`。

在一个使用了已存在的依赖管理工具（比如：`godep`、`glide`或`dep`）的项目中，`go mod init`工具也会添加`require`语句来匹配已存在的配置。

一旦`go.mod`文件已存在，不需要额外的步骤：`go命令（比如：'go build'、'go test'或'go list'）`会自动添加新依赖项来满足导入项的需要。

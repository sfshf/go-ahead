
[go](https://golang.google.cn/cmd/go/)

#####  一、用法

```

go <command> [arguments]

```


##### 二、子命令

| 命令名 | 概要 |
|--|--|
| [bug](subcmds/bug.md) | 开启一个漏洞报告。 |
| [build](subcmds/build.md) | 编译包及依赖。 |
| [clean](subcmds/clean.md) | 删除对象文件和缓存文件 |
| [doc](subcmds/doc.md) | 显示包或者symbol（标识）的说明文档。 |
| [env](subcmds/env.md) | 打印Go开发环境信息。 |
| [fix](subcmds/fix.md) | 更新包文件，使其使用新的API。 |
| [fmt](subcmds/fmt.md) | 格式化包资源。 |
| [generate](subcmds/generate.md) | 通过处理资源生成Go文件。 |
| [get](subcmds/get.md) | 添加依赖到当前模块并安装它们。 |
| [install](subcmds/install.md) | 编译并安装包和依赖。 |
| [list](subcmds/list.md) | 列出包或者模块。 |
| [mod](subcmds/mod.md) | 模块维护。 |
| [run](subcmds/run.md) | 编译并运行Go程序。 |
| [test](subcmds/test.md) | 测试包文件。 |
| [tool](subcmds/tool.md) | 运行指明的go工具。 |
| [version](subcmds/version.md) | 打印Go的版本。 |
| [vet](subcmds/vet.md) | 报告包文件中的错误。 |

##### 三、`go help <topic>`帮助命令工具

在命令行使用`go help`帮助命令工具可以查看`go`命令下各个子命令的用法说明，还可以查看其他列出的主题内容。

##### 四、几个重要主题的阐释

| 主题名 | 简单解释 |
|--|--|
| [Build modes](topics/buildmode.md) | 构建模式。 |
| [Calling between Go and C](topics/cgo.md) | Go语言与C/C++语言之间调用。 |
| [Build and test caching](topics/cache.md) | 构建和测试缓存。 |
| [Environment variables](topics/envar.md) | 环境变量。 |
| [File types](topics/filetype.md) | 文件类型。 |
| [The go.mod file](topics/gomod.md) | `go.mod`文件。 |
| [GOPATH environment variable](topics/gopath.md) | `GOPATH`环境变量。 |
| [GOPATH and Modules](topics/gopathmodule.md) | `GOPATH`和模块。 |
| [Internal Directories](topics/internal.md) | `internal`目录。 |
| [Vendor Directories](topics/vendor.md) | `vendor`目录。 |
| [Legacy GOPATH go get](topics/gopathmode.md) | `go get`的`legacy GOPATH`模式。 |
| [Module proxy protocol](topics/moduleproxy.md) | 模块代理协议。 |
| [Import path syntax](topics/importpath.md) | 导入路径的语法。 |
| [Relative import paths](topics/importpathrelative.md) | 相对的导入路径。 |
| [Remote import paths](topics/importpathremote.md) | 远程导入路径。 |
| [Import path checking](topics/importpathcheck.md) | 导入路径的检查。 |
| [Modules, module versions, and more](topics/module.md) | 模块、模块版本及其他。 |
| [Module support](topics/modulesupport.md) | 模块的使用。 |
| [Defining a module](topics/moduledefine.md) | 定义一个模块。 |
| [The main module and the build list](topics/modulelist.md) | 主模块和构建列表。 |
| [Maintaining module requirements](topics/modulerequire.md) | 维护模块需求。 |
| [Pseudo-versions](topics/modulepseudo.md) | 伪版本。 |
| [Module queries](topics/modulequery.md) | 模块查询。 |
| [Module compatibility and semantic versioning](topics/modulecompatibility.md) | 模块兼容性和语义化版本。 |
| [Module code layout](topics/modulelayout.md) | 模块代码布局。 |
| [Module downloading and verification](topics/moduledownload.md) | 模块下载和验证。 |
| [Modules and vendoring](topics/modulevendoring.md) | 模块和依赖供应。 |
| [Module authentication using go.sum](topics/modulegosum.md) | 使用`go.sum`文件来进行模块认证。 |
| [Module authentication failures](topics/moduleauthfail.md) | 有关模块认证失败。 |
| [Module configuration for non-public modules](topics/moduleconfig.md) | 非公共模块的模块配置。 |
| [Package lists and patterns](topics/packagelist.md) | 包列表和模式。 |
| [Testing flags](topics/testflag.md) | 测试标志。 |
| [Testing functions](topics/testfunc.md) | 测试函数。 |

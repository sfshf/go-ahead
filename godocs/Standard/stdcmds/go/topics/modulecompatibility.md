
[Module compatibility and semantic versioning](https://golang.google.cn/cmd/go/#hdr-Module_compatibility_and_semantic_versioning)


##### 说明：

`go命令`要求模块使用语义版本，并期望版本能够准确地描述兼容性：它假设`v1.5.4`是`v1.5.3`、`v1.4.0`甚至`v1.0.0`的向后兼容替代版本。一般来说，`go命令`期望包都遵循`import compatibility rule（导入兼容性规则）`，即：如果旧包和新包具有相同的导入路径，则新包必须向后兼容旧包。

因为`go命令`假设导入兼容规则，模块定义只能设置一个依赖项所需的最小版本:它不能设置最大值或排除选定的版本。尽管如此，导入兼容性规则并不是一个保证:可能是`v1.5.4`有bug，而且不是`v1.5.3`的向后兼容替代版本。因此，`go命令`不会在未请求的情况下从旧版本更新到新版本。

在语义版本控制中，更改主版本号表示缺乏与早期版本的向后兼容性。为了保持导入兼容性，`go命令`要求主版本`v2`或更高版本的模块要使用对应的主版本号作为最终元素的模块路径。举例，`example.com/m`模块的`v2.0.0`版本必须替换为`example.com/m/v2`模块路径，并且模块中的包要使用该模块路径作为导入路径的前缀，例如`example.com/m/v2/sub/pkg`。包含主版本号的模块路径和导入路径被称为`语义化导入版本控制（semantic import versioning）`。带有主版本号或更高版本号的模块的`伪版本（Pseudo-version）`会使用所携带的版本号替代`v0`，例如`v2.0.0-20180326061214-4fc5987536ef`。

有一个特殊情况，`gopkg.in/`开头的模块路径继续使用建立在其系统上的约定：朱版本号是显式的，并且使用`.`号代理`/`号，例如`gopkg.in/yaml.v1`和`gopkg.in/yaml.v2`，而不是`gopkg.in/yaml`和`gopkg.in/yaml/v2`。

`go命令`会认为不同模块路径的模块之间是没有关联的，例如，`example.com/m`和`example.com/m/v2`之间是没有联系的。不同主版本的模块可以在一次构建中被一起使用，并且它们的包使用不同的导入路径来保持区分。

在`语义版本控制`中，`主版本v0`用于初始开发，表示不期望稳定性或向后兼容性。`主版本v0`不会出现在模块路径中，因为这些版本是为`v1.0.0`准备的，而`v1`也不会出现在模块路径中。

在引入语义导入版本控制约定之前编写的代码可以使用`主版本v2`及其以后的版本来描述`v0`和`v1`中使用的未版本化的导入路径集。为了适应这样的代码，如果源代码存储库中的一个没有`go.mod`文件的文件树有一个`v2.0.0`或更高版本的标记，那么该版本被认为是`v1`模块可用版本的一部分，并在转换为模块版本时被赋予`+incompatible`后缀，如`v2.0.0+incompatible`。`+incompatible`标记也应用于从这些版本派生的伪版本，如`v2.0.1-0.yyyymmddhhmms -abcdefabcdef+incompatible`。

通常，在`v0版本`、`预发布版本`、`伪版本`或`+incompatible`版本的构建列表中有一个依赖项(如`go list -m all`所报告的)，这表明在升级该依赖项时更有可能出现问题，因为没有对这些依赖项的兼容性预期。

有关`语义化导入版本控制`的更多信息，请查看[https://research.swtch.com/vgo-import](https://research.swtch.com/vgo-import)；有关`语义化版本控制`的更多信息，请查看[https://semver.org/](https://semver.org/)。

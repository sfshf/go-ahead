
[The go.mod file](https://golang.google.cn/cmd/go/#hdr-The_go_mod_file)


##### 说明：

模块版本是由源文件树定义的,且其根目录中有`go.mod`文件。当`go命令`运行时，命令会在当前目录中查找用于标记主(当前)模块的根的`go.mod`文件，如果没找到，则会当前目录的父目录中查找；依次方式直到找到`go.mod`文件。

`go.mod`文件本身是`面向行（line-oriented）`的，支持有`//`注释，但是没有`/* */`注释。每一行包含一个指令，由`动词（verb）`和`参数`组成。例如:

```

module my/thing
go 1.12
require other/thing v1.0.2
require new/thing/v2 v2.3.4
exclude old/thing v1.2.3
replace bad/thing v1.4.5 => good/thing v1.4.5

```

`动词（verb）`有：

| 动词名 | 说明 |
|--|--|
| module | 定义模块路径。 |
| go | 设置期望的go语言版本。 |
| require | 引用一个特定的给定版本或更新版本的模块。 |
| exclude | 排除一个特定的模块版本。 |
| replace | 用一个不同的模块版本替代一个模块版本。 |

`exclude`和`replace`只应用于主模块的`go.mod`文件中，而在依赖项中被忽略。详见[https://research.swtch.com/vgo-mvs](https://research.swtch.com/vgo-mvs)。

每行领头的`动词`可以从相邻行中提取公因数出来，从而创建一个块，就像在Go包的导入方式那样，例如:

```

require (
	new/thing v2.3.4
	old/thing v1.2.3
)

```

`go.mod`文件既可以直接编辑，也可以通过工具轻松更新。`go mod edit`命令可用于解析和编辑来自从程序和工具的`go.mod`文件。参见`go help mod edit`。

`go命令`在每次使用的模块图（module graph）时，会自动更新`go.mod`文件，以确保`go.mod`文件总是能够准确地反映现实，并且格式正确。举例：

```

module M

require (
        A v1
        B v1.0.0
        C v1.0.0
        D v1.2.3
        E dev
)

exclude D v1.2.3

```

更新会将`非规范版本标识符`重写为`语义化格式`（semver form），因此`模块A`的`v1`会变成`v1.0.0`, `模块E`的`dev`会变成`dev`分支上最新提交的伪版本（pseudo-version），可能是`v0.0.0-20180523231146-b3f5c0f6e5f1`。

更新会为了遵守排除项而修改需求项，因此对被排除的`D v1.2.3`的需求进行了更新，会使用下一个可用的`模块D`版本，可能是`D v1.2.4`或`D v1.3.0`。

更新会删除冗余的或有歧义的需求项。例如，如果`A v1.0.0`本身需要`B v1.2.0`和`C v1.0.0`，那么`go.mod`文件中对`B v1.0.0`的要求是具有歧义的(被`模块A`对`v1.2.0`的需要所取代)，而`go.mod`文件中对`C v1.0.0`的要求则是多余的(由`模块A`对同一版本的需要所暗示)，因此两者都将被删除。如果`模块M`中有直接导入`模块B`或`模块C`中包的包，那么需求将被保留，但是会更新到实际使用的版本。

最后，更新会重新格式化`go.mod`文件为一个规范的格式，使得未来可能出现的无意识的变化只会引起最低限度的差异。

由于模块图（module graph）定义了`import语句`的含义，所以有关加载包的任何命令也会使用并更新`go.mod`文件，包括`go build`，`go get`，`go install`，`go list`，`go test`，`go mod graph`，`go mod tidy`, `go mod why`。

由`go指令`设置的期望的语言版本决定了在编译模块时哪些语言特性可用。该版本中可用的语言特性将可供使用。在早期版本中删除的或在后期版本中添加的语言特性将不可用。注意，语言版本不影响构建标记（build tags），构建标记由被使用的Go语言发行版本决定。

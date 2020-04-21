
[List packages or modules](https://golang.google.cn/cmd/go/#hdr-List_packages_or_modules)

##### 一、用法

```

go list [-f format] [-json] [-m] [list flags] [build flags] [packages]

```

##### 二、list命令说明

`list命令`会逐行列出被命名的包。最常用的标记是`-f`和`-json`，它们可以控制包的输出形式。下面记述了`list命令`的其他用来控制更多指定细节的标志。

默认的输出会显示包的导入路径，例如：

```

bytes
encoding/json
github.com/gorilla/mux
golang.org/x/net/html

```

`-f`标志使用包模板的语法为`list命令`指定了一种替代格式。默认输出相当于`-f '{{.ImportPath}}'`。传递给模板的结构是:

```go

type Package struct {
    Dir           string   // directory containing package sources
    ImportPath    string   // import path of package in dir
    ImportComment string   // path in import comment on package statement
    Name          string   // package name
    Doc           string   // package documentation string
    Target        string   // install path
    Shlib         string   // the shared library that contains this package (only set when -linkshared)
    Goroot        bool     // is this package in the Go root?
    Standard      bool     // is this package part of the standard Go library?
    Stale         bool     // would 'go install' do anything for this package?
    StaleReason   string   // explanation for Stale==true
    Root          string   // Go root or Go path dir containing this package
    ConflictDir   string   // this directory shadows Dir in $GOPATH
    BinaryOnly    bool     // binary-only package (no longer supported)
    ForTest       string   // package is only for use in named test
    Export        string   // file containing export data (when using -export)
    Module        *Module  // info about package's containing module, if any (can be nil)
    Match         []string // command-line patterns matching this package
    DepOnly       bool     // package is only a dependency, not explicitly listed

    // Source files
    GoFiles         []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
    CgoFiles        []string // .go source files that import "C"
    CompiledGoFiles []string // .go files presented to compiler (when using -compiled)
    IgnoredGoFiles  []string // .go source files ignored due to build constraints
    CFiles          []string // .c source files
    CXXFiles        []string // .cc, .cxx and .cpp source files
    MFiles          []string // .m source files
    HFiles          []string // .h, .hh, .hpp and .hxx source files
    FFiles          []string // .f, .F, .for and .f90 Fortran source files
    SFiles          []string // .s source files
    SwigFiles       []string // .swig files
    SwigCXXFiles    []string // .swigcxx files
    SysoFiles       []string // .syso object files to add to archive
    TestGoFiles     []string // _test.go files in package
    XTestGoFiles    []string // _test.go files outside package

    // Cgo directives
    CgoCFLAGS    []string // cgo: flags for C compiler
    CgoCPPFLAGS  []string // cgo: flags for C preprocessor
    CgoCXXFLAGS  []string // cgo: flags for C++ compiler
    CgoFFLAGS    []string // cgo: flags for Fortran compiler
    CgoLDFLAGS   []string // cgo: flags for linker
    CgoPkgConfig []string // cgo: pkg-config names

    // Dependency information
    Imports      []string          // import paths used by this package
    ImportMap    map[string]string // map from source import to ImportPath (identity entries omitted)
    Deps         []string          // all (recursively) imported dependencies
    TestImports  []string          // imports from TestGoFiles
    XTestImports []string          // imports from XTestGoFiles

    // Error information
    Incomplete bool            // this package or a dependency has an error
    Error      *PackageError   // error loading package
    DepsErrors []*PackageError // errors loading dependencies
}

```

在`vendor`目录下存储的包会报告一个包含了相对于`vendor`目录的`ImportPath`（例如："d/vendor/p"代替"p"），所以`ImportPath`独一无二地标识一个被给出的包的复制品。`Imports`、`Deps`、`TestImports`和`XTestImports`这几个列表也包含了展开的导入路径。

如果有任何错误信息，可以使用`PackageError`：

```go

type PackageError struct {
    ImportStack   []string // shortest path from package named on command line to this one
    Pos           string   // position of error (if present, file:line:col)
    Err           string   // the error itself
}

```

模块信息是使用了一个`Module`结构体。

模板函数`join`调用了`strings.Join`函数。

模板函数`context`返回一个构建上下文，定义如下：

```go

type Context struct {
    GOARCH        string   // target architecture
    GOOS          string   // target operating system
    GOROOT        string   // Go root
    GOPATH        string   // Go path
    CgoEnabled    bool     // whether cgo can be used
    UseAllFiles   bool     // use files regardless of +build lines, file names
    Compiler      string   // compiler to assume when computing target paths
    BuildTags     []string // build constraints to match in +build lines
    ReleaseTags   []string // releases the current release is compatible with
    InstallSuffix string   // suffix to use in the name of the install dir
}

```

想要了解更多关于`Context`结构体中字段的意义请查看`go/build`包里对`Context`的定义。

`-json`标志会替代模板格式，而使用JSON格式打印出包数据。

`-compiled`标志会导致`list命令`将`CompiledGoFiles`设置为呈现给编译器的Go源文件。通常这意味着它会重复`GoFiles`中列出的文件，然后添加通过处理`CgoFiles`和`SwigFiles`生成的Go代码。`Imports`列出了包含来自`GoFiles`和`CompiledGoFiles`里所有导入的并集。

`-deps`标志导致`list命令`不仅遍历指定的包，而且遍历它们的所有依赖项。它以深度优先（depth-first）的后序遍历（post-order）访问它们，因此包只在其所有依赖项之后列出。命令行中未显式列出的包将`DepOnly`字段设置为true。

`-e`标志改变了对错误包的处理，错误包是指不能被找到或者是畸形的的包。默认情况下，`list命令`为每个错误包打印一个错误到标准错误，并在通常的打印过程中忽略这些包。使用`-e`标志，`list命令`不会将错误打印到标准错误，而是使用通常的打印方式处理错误包。错误包将有一个非空的`ImportPath`和一个非空的`Error`字段;其他信息可能丢失，也可能没有丢失(归零)。

`-export`标志导致`list命令`将`Export`字段设置为包含了给定包的最新导出信息的文件的名字。

`-find`标志导致`list命令`标识指名的包，但不能解析它们的依赖:`Imports`和`Deps`两个字段将置空。

`-test`标志使`list命令`不仅报告指名的包，而且报告它们的测试二进制文件(对于带有测试的包)，以便向源代码分析工具准确地传达测试二进制文件是如何构造的。测试二进制文件的报告导入路径是包的导入路径，后面跟一个`.test`后缀，例如`math/rand.test`。在构建测试时，有时需要重建特定于测试的依赖项(最常见的是测试包本身)。为特定的测试二进制文件重新编译的包的报告导入路径后面是一个空格和测试二进制文件的名字，例如`math/rand [math/rand.test]`或`regexp [sort.test]`。`ForTest`字段也被设置为要测试的包的名字(在前面的例子中是`math/rand`或`sort`)。

`Dir`、`Target`、`Shlib`、`Root`、`ConflictDir`、`Export`各字段的文件路径都是绝对路径。

默认情况下，`GoFiles`、`CgoFiles`等等字段持有的是在`Dir`字段指示的目录里的文件的名字（也就是说，相对于`Dir`的路径，而不是绝对路径）。使用`-compile`和`-test`标记时添加的生成文件是指向生成的Go源文件的缓存副本的绝对路径。虽然它们是Go源文件，但路径可能不会以`.go`结尾。

`-m`标志导致`list命令`列出模块而不是包。

当列出模块时，`-f`标志仍然定义了一个格式模板应用于一个Go结构体上，只是现在是一个`Module`结构体：

```go

type Module struct {
    Path      string       // module path
    Version   string       // module version
    Versions  []string     // available module versions (with -versions)
    Replace   *Module      // replaced by this module
    Time      *time.Time   // time version was created
    Update    *Module      // available update, if any (with -u)
    Main      bool         // is this the main module?
    Indirect  bool         // is this module only an indirect dependency of main module?
    Dir       string       // directory holding files for this module, if any
    GoMod     string       // path to go.mod file for this module, if any
    GoVersion string       // go version used in module
    Error     *ModuleError // error loading module
}

type ModuleError struct {
    Err string // the error itself
}

```

默认输出是打印模块路径，然后是关于版本和替换的信息(如果有的话)。举例：`go list -m all`也许会打印出：

```

my/main/module
golang.org/x/text v0.3.0 => /tmp/text
rsc.io/pdf v0.1.1

```

`Module`结构体有一个`String`方法来格式化这一行输出，因此默认格式相当于`-f '{{.String}}'`。

注意，当一个模块被替换时，它的`Replace`字段将描述用于替换的模块，它的`Dir`字段将设置为用于替换的模块的源代码。(也就是说，如果`Replace`非空，则将`Dir`设置为`Replace.Dir`，无法访问被替换的源代码。)

`-u`标志添加关于可用升级的信息。当给定模块的最新版本比当前模块更新时，`list -u`将模块的`Update`字段设置为关于新模块的信息。模块的`String`方法通过在当前版本之后的方括号中格式化较新的版本来指示可用的升级。例如，`go list -m -u all`可能会打印:

```

my/main/module
golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
rsc.io/pdf v0.1.1 [v0.1.2]

```

（作为工具而言，`go list -m -u -json all`可能更便于解析。）

`-versions`标志使`list命令`将模块的`Versions`字段设置为该模块的所有已知版本的列表，按照语义版本的顺序排列，从最早到最新。该标志还更改了默认的输出格式，以显示模块路径后面跟着空格分隔的版本列表。

`list命令`的`-m`标志的参数被解释为模块列表，而不是包列表。主模块（the main module）是包含当前目录的模块。活动模块（the active modules）是主模块及其依赖项。在没有参数的情况下，`list命令`的`-m`标志显示主模块。在有参数的情况下，`list命令`的`-m`标志显示由参数指定的模块。任何活动模块都可以通过其模块路径指定。特殊模式`all`指定所有活动模块，第一个是主模块，然后是按模块路径排序的依赖项。包含`…`的模式指定其模块路径与模式匹配的活动模块。格式为`path@version`的查询指定了该查询的结果，该格式不限于活动模块。有关模块查询的更多信息，请参见`go help modules`。

模板函数`module`接收一个字符串参数，该参数必须是模块路径或查询，并以`Module`结构体返回指定的模块。如果发生错误，结果将是一个带有非空`Error`字段的`Module`结构体。

使用`go help build`查看更多关于构建标志的信息。

使用`go help packages`查看更多关于包的说明。

使用`go help modules`查看更多关于模块的信息。

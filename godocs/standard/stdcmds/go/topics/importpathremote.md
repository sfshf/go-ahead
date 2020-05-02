
[Remote import paths](https://golang.google.cn/cmd/go/#hdr-Remote_import_paths)


##### 说明：

某些导入路径还描述了如何使用修订控制系统获取包的源代码。

一些常见的代码托管站点有特定的语法:

```

Bitbucket (Git, Mercurial)

	import "bitbucket.org/user/project"
	import "bitbucket.org/user/project/sub/directory"

GitHub (Git)

	import "github.com/user/project"
	import "github.com/user/project/sub/directory"

Launchpad (Bazaar)

	import "launchpad.net/project"
	import "launchpad.net/project/series"
	import "launchpad.net/project/series/sub/directory"

	import "launchpad.net/~user/project/branch"
	import "launchpad.net/~user/project/branch/sub/directory"

IBM DevOps Services (Git)

	import "hub.jazz.net/git/user/project"
	import "hub.jazz.net/git/user/project/sub/directory"

```

对于托管在其他服务器上的代码，导入路径可以使用版本控制类型进行限定，或者go工具可以通过`https/http`动态获取导入路径，并从`HTML`中的`<meta>`标签中发现代码所在的位置。

想要声明代码位置，一个格式为`repository.vcs/path`的导入路径，就指明了使用指定的版本控制系统的存储库(带有或不带有`.vcs`后缀)，然后指定存储库中的路径。支持的版本控制系统有：

| 版本控制系统 | 后缀 |
|--|--|
| Bazaar | .bzr |
| Fossil | .fossil |
| Git | .git |
| Mercurial | .hg |
| Subversion | .svn |

举例：`import "example.org/user/foo.hg"`，表示了`Mercurial`存储库里`example.org/user/foo`或`foo.hg`的根目录；`import "example.org/repo.git/foo/bar"`，表示了`Git`存储库里`example.org/repo`或`repo.git`的`foo/bar`目录。

当一个版本控制系统支持多个协议（protocol）时，每一个协议会在下载时被轮流尝试。例如：`Git`下载会尝试`https://`，然后`git+ssh://`。

默认情况下，下载仅限于已知的安全协议(如`https`、`ssh`)。要覆盖`Git`下载的这个设置，可以设置G`IT_ALLOW_PROTOCOL`环境变量（更多细节参见:`go help environment`）。

如果导入路径不是一个已知的代码托管站点，并且缺少版本控制限定符，那么`go工具`将尝试通过`https/http`获取导入，并在`HTML`文档中的`<head>`标签里查找`<meta>`标签。

`<meta>`标签的格式为：`<meta name="go-import" content="import-prefix vcs repo-root">`。其中，`import-prefix`是与存储库根对应的导入路径；它必须是一个前缀或者是与可被`go get`获取的包完全匹配。如果不是完全匹配，则在前缀处发出另一个`http`请求来验证`<meta>`标签是否匹配。

`<meta>`标签应该尽可能早地出现在文件中。特别是，它应该出现在任何原始`JavaScript`或`CSS`之前，以避免混淆`go命令`的受限解析器。

`vcs（Virus Capture Scripter，集群服务器）`有：`bzr`、`fossil`、`git`、`hg`、`svn`。

`repo-root（仓库根目录）`是一个版本控制系统的根目录，它包含了一个`scheme（模式；方案）`，并不包含一个`.vcs`（例如：`.bzr`、`.git`等）限定符。

举例：`import "example.org/pkg/foo"`将导致下列的请求：

```

https://example.org/pkg/foo?go-get=1 (preferred)
http://example.org/pkg/foo?go-get=1  (fallback, only with -insecure)

```

如果获取的网页包含`<meta>`标签`<meta name="go-import" content="example.org git https://code.org/r/p/exproj">`，那么`go`工具会验证`https://example.org/?go-get=1`是否包含相同的`<meta>`标签，然后执行`git clone https://code.org/r/p/exproj`，将源码复制到`GOPATH/src/example.org`目录下。

当使用`GOPATH`时，被下载的包会被写到`GOPATH`环境变量中列出的第一个目录中。（详见：`go help gopath-get`和`go help gopath`）。

当使用模块时，名为`go-import`的`<meta>`标签的一个附加变体被识别，并且比那些列出的版本控制系统更受欢迎。该变量使用`mod`作为`content`值中的`vcs`，如:`<meta name="go-import" content="example.org mod https://code.org/moduleproxy">`。

这个`<meta>`标签意味着从模块代理站点（该模块代理站点可用URL--`https://code.org/moduleproxy`访问）获取模块，该模块使用以`example.org`开头的路径。更多关于代理协议（proxy protocol）的详细信息，可查看`go help goproxy`。

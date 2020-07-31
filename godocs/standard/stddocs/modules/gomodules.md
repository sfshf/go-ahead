
# [Go Mudules](https://github.com/golang/go/wiki/Modules#quick-start)

从`1.11`开始，`Go`就包含了对`版本化模块`的支持。最初的原型`vgo`于2018年2月发布。2018年7月，`版本化模块`在`主Go存储库`中落地。

在`Go 1.14`中，`模块支持`被认为已准备好用于生产，并且鼓励所有用户从其他`依赖项管理系统`迁移到`模块`。如果由于`Go工具链`中的问题而无法迁移，请确保[将该问题开放归档](https://github.com/golang/go/wiki/Modules#github-issues)。（如果问题不在`Go1.15`里程碑上，请留言为什么该问题阻止了您迁移，以便可以对该问题进行优先级排序）。您还可以[提供一份经历报告](https://github.com/golang/go/wiki/ExperienceReports)，来获得更详细的反馈。


## Recent Changes

### Go 1.14

可以查看`Go 1.14发行消息`获取细节信息。

- 当`主模块`包含顶级`vendor`目录并且其`go.mod`文件指定明`go 1.14`或更高版本时，`go命令`现在默认接收`-mod=vendor`标识来进行操作。
- 现在，当`go.mod`文件为`只读`并且不存在顶级`vendor`目录时，默认情况下会设置`-mod=readonly`标识。
- `-modcacherw`是一个新标识，它指示`go命令`以默认权限`将新创建的目录保留在模块高速缓存中`，而不是将其设置为只读。
- `-modfile=file`是一个新标识，指示`go命令`读取（并可能写入）`备用go.mod`文件，而不是`模块根目录`中的`go.mod`文件。
- 当显式启用`模块感知模式`（通过将`GO111MODULE=on`设置）时，如果不存在`go.mod`文件，则大多数模块命令的功能将受到更多限制。
- `go命令`现可在`模块模式`下支持`子版本存储库`。


### Go 1.13

可以查看`Go 1.13发行消息`获取细节信息。

- 现在，`go工具`默认是从`https://proxy.golang.org`上的`公共Go模块镜像`下载模块，并且还默认使用`https://sum.golang.org`上的`公共Go校验和数据库`来验证下载的模块（无论源如何）。
  - 如果您有`隐私代码`，那么您很可能应该配置`GOPRIVATE`设置（例如`go env -w GOPRIVATE=*.corp.com,github.com/secret/repo`），或者更精细的变体`GONOPROXY`或`GONOSUMDB`（不太频繁使用）。更多详细信息，请参见`go命令文档`。
- 如果发现了任何的`go.mod`文件，即使在`GOPATH`里，`GO111MODULE=auto`也会启用`模块模式`。（在`Go 1.13`之前，`GO111MODULE=auto`永远不会在`GOPATH`里启用`模块模式`）。
- `go get`的参数已被更改：
  - `go get -u`（不带任何参数），现在仅升级您当前程序包的直接和间接依赖项，而不再检查整个模块。
  - 来自您模块根目录的`go get -u ./...`，会升级您模块的所有直接和间接依赖项，并且现在不包括`测试依赖项`。
  - `go get -u -t ./...`作用同上，但是会升级`测试依赖项`。
  - `go get`不再支持`-m`（由于其他的更改，它会与`go get -d`在很大程度上重叠；通常可以将`go get -m foo`替换为`go get -d foo`）。


## Table of Contents

对于刚开始使用模块的人员，`Quick Start`和`New Concepts`部分特别重要。`How to...`部分介绍了有关`机制原理`的更多详细信息。此页面上的内容最多的是在`FAQ`部分，其回答了更具体的问题；浏览一下列出的FAQ是很值得的。

- [快速开始](#quick-start)
    - [示例](#example)
    - [日常工作流程](#daily-workflow)
- [新概念](#new-concepts)
    - [模块](#modules)
    - [`go.mod`](#gomod)
    - [版本选择](#version-selection)
    - [语义导入版本控制](#semantic-import-versioning)
- [如何使用模块](#how-to-use-modules)
    - [如何安装和激活模块支持](#how-to-install-and-activate-module-support)
    - [如何定义模块](#how-to-define-a-module)
    - [如何升级和降级依赖项](#how-to-upgrade-and-downgrade-dependencies)
    - [如何准备发行版（所有版本）](#how-to-prepare-for-a-release)
    - [如何准备发行版（v2或更高版本）](#releasing-modules-v2-or-higher)
    - [发布一个发行版](#publishing-a-release)
- [迁移到模块](#migrating-to-modules)
- [其他资源](#additional-resources)
- [自`最初的Vgo提案`以来的变化](#changes-since-the-initial-vgo-proposal)
- [GitHub问题](#github-issues)
- [常见问题](#faqs)
    - [版本如何被标记为不兼容？](#how-are-versions-marked-as-incompatible)
    - [什么时候使用旧行为，以及什么时候使用新的基于模块的行为？](#when-do-i-get-old-behavior-vs-new-module-based-behavior)
    - [为什么通过`go get`安装工具失败，并显示错误`cannot find main module`？](#why-does-installing-a-tool-via-go-get-fail-with-error-cannot-find-main-module)
    - [如何跟踪模块的工具依赖项？](#how-can-i-track-tool-dependencies-for-a-module)
    - [IDE，编辑器和标准工具（例如`goimports`，`gorename`等）中模块支持的状况如何？](#what-is-the-status-of-module-support-in-ides-editors-and-standard-tools-like-goimports-gorename-etc)
- [常见问题 -- 附加控制](#faqs--additional-control)
    - [存在哪些社区工具来使用模块？](#what-community-tooling-exists-for-working-with-modules)
    - [什么时候应该使用`replace`指令？](#when-should-i-use-the-replace-directive)
    - [在本地文件系统上可以完全脱离VCS来工作了吗？](#can-i-work-entirely-outside-of-vcs-on-my-local-filesystem)
    - [如何将供应控制与模块一起使用？供应控制会消失吗？](#how-do-i-use-vendoring-with-modules-is-vendoring-going-away)
    - [是否有`始终在线`的模块存储库和企业代理？](#are-there-always-on-module-repositories-and-enterprise-proxies)
    - [可以控制`go.mod`何时更新以及`go工具`何时使用网络来满足依赖项吗？](#can-i-control-when-gomod-gets-updated-and-when-the-go-tools-use-the-network-to-satisfy-dependencies)
    - [如何将模块与例如`Travis`或`CircleCI`等的`CI系统`一起使用？](#how-do-i-use-modules-with-ci-systems-such-as-travis-or-circleci)
    - [如何下载构建特定包或测试所需的模块？](#how-do-i-download-modules-needed-to-build-specific-packages-or-tests)
- [常见问题 -- `go.mod`和`go.sum`](#faqs--gomod-and-gosum)
    - [为什么`go mod tidy`会在`go.mod`中记录间接的和测试的依赖项？](#why-does-go-mod-tidy-record-indirect-and-test-dependencies-in-my-gomod)
    - [`go.sum`是锁定的文件吗？为什么`go.sum`会包含不再使用的模块版本的信息？](#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using)
    - [我应该提交`go.sum`和`go.mod`文件吗？](#should-i-commit-my-gosum-file-as-well-as-my-gomod-file)
    - [如果我没有任何依赖项，还应该添加一个`go.mod`文件吗？](#should-i-still-add-a-gomod-file-if-i-do-not-have-any-dependencies)
- [常见问题 -- 语义导入版本控制](#faqs--semantic-import-versioning)
    - [为什么`主要版本号`必须出现在导入路径中？](#why-must-major-version-numbers-appear-in-import-paths)
    - [为什么导入路径中省略了`主要版本v0，v1`？](#why-are-major-versions-v0-v1-omitted-from-import-paths)
    - [用`主要版本v0，v1`标记项目或`使用v2+进行重大更改`有什么含义？](#what-are-some-implications-of-tagging-my-project-with-major-version-v0-v1-or-making-breaking-changes-with-v2)
    - [模块可以使用`未选择模块的软件包`吗？](#can-a-module-consume-a-package-that-has-not-opted-in-to-modules)
    - [一个模块可以使用`未选择模块的v2+软件包`吗？ `+incompativle`是什么意思？](#can-a-module-consume-a-v2-package-that-has-not-opted-into-modules-what-does-incompatible-mean)
    - [如果未启用模块支持，在build中是如何处理v2+模块的？`1.9.7+`，`1.10.3+`和`1.11`中的`最小模块兼容性`是如何工作的？](#how-are-v2-modules-treated-in-a-build-if-modules-support-is-not-enabled-how-does-minimal-module-compatibility-work-in-197-1103-and-111)
    - [如果我创建`go.mod`但不将`semver标记`应用于存储库，会发生什么情况？](#what-happens-if-i-create-a-gomod-but-do-not-apply-semver-tags-to-my-repository)
    - [模块可以依赖于其自身的不同版本吗？](#can-a-module-depend-on-a-different-version-of-itself)
- [常见问题 -- 多模块存储库](#faqs--multi-module-repositories)
    - [什么是`多模块存储库`？](#what-are-multi-module-repositories)
    - [我应该在一个存储库中有多个模块吗？](#should-i-have-multiple-modules-in-a-single-repository)
    - [是否可以将模块添加到`多模块存储库`？](#is-it-possible-to-add-a-module-to-a-multi-module-repository)
    - [是否可以从`多模块存储库`中删除模块？](#is-it-possible-to-remove-a-module-from-a-multi-module-repository)
    - [一个模块可以依赖于`内部模块`吗？](#can-a-module-depend-on-an-internal-in-another)
    - [额外的`go.mod`可以排除不必要的内容吗？模块是否具有等效于`.gitignore`文件的功能？](#can-an-additional-gomod-exclude-unnecessary-content-do-modules-have-the-equivalent-of-a-gitignore-file)
- [常见问题 -- 最小版本选择](#faqs--minimal-version-selection)
    - [最少的版本选择是否会使开发人员无法获得重要的更新？](#wont-minimal-version-selection-keep-developers-from-getting-important-updates)
- [常见问题 -- 可能的问题](#faqs--possible-problems)
    - [如果我发现了问题，可以进行哪些常规检查？](#what-are-some-general-things-i-can-spot-check-if-i-am-seeing-a-problem)
    - [如果没有看到期望的依赖版本，该如何检查？](#what-can-i-check-if-i-am-not-seeing-the-expected-version-of-a-dependency)
    - [为什么会出现错误`cannot find module providing package foo`？](#why-am-i-getting-an-error-cannot-find-module-providing-package-foo)
    - [为什么`go mod init`给出错`cannot determine module path for source directory`？](#why-does-go-mod-init-give-the-error-cannot-determine-module-path-for-source-directory)
    - [我有一个尚未选择模块的复杂依赖性问题。我可以使用其当前`依赖项管理器`中的信息吗？](#i-have-a-problem-with-a-complex-dependency-that-has-not-opted-in-to-modules-can-i-use-information-from-its-current-dependency-manager)
    - [如何解决由在导入路径及声明的模块身份的缺失引起的`parsing go.mod: unexpected module path`和`error loading module requirements`错误？](#how-can-i-resolve-parsing-gomod-unexpected-module-path-and-error-loading-module-requirements-errors-caused-by-a-mismatch-between-import-paths-vs-declared-module-identity)
    - [为什么`go build`需要`gcc`，而诸如`net/http`之类的预构建软件包不需要？](#why-does-go-build-require-gcc-and-why-are-prebuilt-packages-such-as-nethttp-not-used)
    - [模块是否可以与相对导入（例如`import ./subdir`）一起使用吗？](#do-modules-work-with-relative-imports-like-import-subdir)
    - [使用的`vendor`目录中可能没有某些所需的文件](#some-needed-files-may-not-be-present-in-populated-vendor-directory)


## Quick Start

### Example

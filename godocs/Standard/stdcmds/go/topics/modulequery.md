
[Module queries](https://golang.google.cn/cmd/go/#hdr-Module_queries)


##### 说明：

`go命令`在命令行上和在主模块的`go.mod`文件中可以接受一个`module query（模块查询）`，以代替模块版本。(在估值主模块的`go.mod`文件中找到的查询之后，`go命令`会更新文件，用估值的结果替换查询。)

完全指明的语义版本，如`v1.2.3`，估值的结果为该指明的版本。

语义版本前缀，如`v1`或`v1.2`，估值为带有该前缀的最新可用标记版本。

语义版本比较，如`<v1.2.3`或`>=v1.5.6`，估值为最接近比较目标的可用标记版本(`<`和`<=`的最新版本、`>`和`>=`的最早版本)。

字符串`latest`会匹配最新的可用标记版本，或者匹配底层源存储库最新的未标记修订。

字符串`upgrade`与`latest`类似，但如果模块当前需要的是比`latest`选择的版本更新的版本(例如，一个较新的预发布版本)，`upgrade`将选择较新的版本。

字符串`patch`匹配模块的最新可用标记版本，主版本号和次版本号与当前需要的版本号相同。如果当前不需要版本，则`patch`等同于`latest`。

底层源存储库的修订标识符(如提交散列前缀、修订标记或分支名称)选择指明的代码修订。如果修订也使用语义版本进行标记，则查询将估值为该语义版本。否则，查询将估值为提交的伪版本。注意，不能以这种方式选择名称与其他查询语法匹配的分支和标记。例如，`查询"v2"`表示以`v2`开头的最新版本，而不是名为`v2`的分支。

与`预发布版本`相比，所有查询都优先选择`发布版本`；例如，`<v1.2.3`会优先选择`v1.2.2`而不是`v1.2.3-pre1`，即使`v1.2.3-pre1`更接近比较对象。

主模块`go.mod`文件中的`exclude`语句不允许使用的模块版本被认为是不可用的，不能被查询返回。

举例，下面的命令都是有效的：

```

go get github.com/gorilla/mux@latest    # same (@latest is default for 'go get')
go get github.com/gorilla/mux@v1.6.2    # records v1.6.2
go get github.com/gorilla/mux@e3702bed2 # records v1.6.2
go get github.com/gorilla/mux@c856192   # records v0.0.0-20180517173623-c85619274f5d
go get github.com/gorilla/mux@master    # records current meaning of master

```

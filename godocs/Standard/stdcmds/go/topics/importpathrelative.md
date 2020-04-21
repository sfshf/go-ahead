
[Relative import paths](https://golang.google.cn/cmd/go/#hdr-Relative_import_paths)


##### 说明：

以`./`或`..`开头的导入路径，称为相对路径。工具链以两种方式支持相对导入路径作为快捷方式。

第一种方式，相对路径可以用作命令行上的简写。如果您在包含有被导入为`unicode`包的代码的目录下工作，并且希望运行`unicode/utf8`包的测试，那么您可以输入`go test ./utf8`，而不需要指定完整的路径。同样，在相反的情况下，`go test ..`可以在`unicode/utf8`目录下测试`unicode`包。还允许使用相对模式（relative pattern），如`go test ./...`来测试所有子目录。有关模式语法的详细信息，请参阅`go help packages`。

第二种方式，如果您正在编译不在工作空间（一般为`$HOME/go`，即`GOPATH`）中的Go程序，则可以在该程序的`import`语句中使用相对路径来引用不在工作空间中的附近代码。这使得在通常的工作空间之外对小型多包程序进行试验变得很容易，但是这些程序不能通过`go install`(没有安装它们的工作空间)进行安装，因此每次构建它们时都要从头开始重新构建。为了避免歧义，Go程序不能在工作空间中使用相对导入路径。

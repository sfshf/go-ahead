
[Testing functions](https://golang.google.cn/cmd/go/#hdr-Testing_functions)


##### 说明：

`go test`命令会在指明的测试包的`*_test.go`文件里找到测试、基准测试和示例函数。

一个测试函数的签名格式如下（`Xxx`表示必须以大写字母开头）：

```go
func TestXxx(t *testing.T) { ... }
```

一个基准测试函数的签名格式如下：

```go
func BenchmarkXxx(b *testing.B) { ... }
```
一个示例函数与测试函数相似，但是不使用`*testing.T`来报告成功或失败，而是将输出打印到`os.Stdout`。如果示例函数中的最后一条注释以`Output:`开头，则将输出与注释进行精确比较（请参见下面的示例）。如果最后一条注释以`Unordered output:`开头，那么也会将输出与注释进行比较，但是会忽略行的顺序。没有此类注释的示例也会被编译但不执行。在`Output:`之后没有文本的示例将会被编译，执行，并且是不指望会产生任何输出的。

`Godoc`显示`ExampleXxx`的主体，以演示命名为`Xxx`的函数，常量或变量的用法。接收器类型为`T`或`*T`的方法`M`的示例函数名为`ExampleT_M`。对于给定的函数，常量或变量，可能有多个示例，以尾随的`_xxx`加以区分，其中`xxx`是不以大写字母开头的后缀。举例如下：

```go
func ExamplePrintln() {
    Println("The output of\nthis example.")
    // Output: The output of
    // this example.
}
```

下面是一个忽略输出顺序的示例函数：

```go
func ExamplePerm() {
    for _, value := range Perm(4) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

当整个测试文件包含单个示例函数，至少一个其他函数，类型，变量或常量声明，并且不包含测试或基准测试函数时，该测试文件将以示例形式显示。

查看`testing`包获取更多信息。

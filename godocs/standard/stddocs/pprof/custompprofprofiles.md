
# [CustomPprofProfiles](https://github.com/golang/go/wiki/CustomPprofProfiles)



`Go语言`提供了一些开箱即用的`pprof分析项`，可以从Go程序中`收集性能分析数据`。

`runtime/pprof`包提供的`内置分析项`：

- **profile（分析，剖面图）**：`CPU分析`，测定程序在活动消耗CPU周期时所花费的时间（而不是在睡眠或等待I/O时花费的时间）。
- **heap（堆）**：`堆分析`，报告了当前活动的内存分配；用于监视当前内存使用情况或检查内存泄漏。
- **threadcreate（线程创建）**：`线程创建分析`，报告了导致创建新`OS线程`的部分程序。
- **goroutine（go协程）**：`goroutine分析`，报告了所有当前goroutine的`栈跟踪`。
- **block（阻塞）**：`阻塞分析`，显示goroutine在哪为了等待同步原语（包括计时器通道）而阻塞。默认情况下不启用`阻塞分析`；使用`runtime.SetBlockProfileRate`可以启用它。
- **mutex（互斥量）**：`互斥量分析`，报告锁的争用情况。如果您认为由于`互斥锁争用`而无法充分利用CPU，请使用此项分析。`互斥量分析`默认情况下未启用，请参见`runtime.SetMutexProfileFraction`启用。

除了`内置分析项`之外，`runtime/pprof`包还允许您导出`自定义分析项`，并用您的代码来记录与该分析项有关的执行栈。

假设我们有一个`blob服务器`，并且正在为其编写`Go语言客户端`。我们的用户希望能够`在客户端上`分析打开的blob。我们可以创建一个`分析项`并记录blob打开和关闭的事件，以便用户可以随时知道有多少个打开的blob。

这是一个`blobstore`包，可让您打开一些blob。我们将创建一个新的`自定义分析项`，并开始记录与打开blob有关的执行栈：


```go

package blobstore

import "runtime/pprof"

var openBlobProfile = pprof.NewProfile("blobstore.Open")

// Open opens a blob, all opened blobs need to be closed when no longer in use.
func Open(name string) (*Blob, error) {
    blob := &Blob{name: name}
    // TODO: Initialize the blob...

    openBlobProfile.Add(blob, 2)  // add the current execution stack to the profile
    return blob, nil
}

```

一旦用户想关闭blob，我们需要删除与当前来自分析项的blob有关的执行栈：

```go

// Close closes the blob and frees the underlying resources.
func (b *Blob) Close() error {
    // TODO: Free other resources.
    openBlobProfile.Remove(b)
    return nil
}

```

现在，在使用了`blobstore`包的程序里，我们可以检索到`blobstore.Open`分析数据，并使用我们日常的`pprof工具`来检测并可视化它们。

让我们写一个简单的`main`程序，并开启一些blob：

```go

package main

import (
    "fmt"
    "math/rand"
    "net/http"
    _ "net/http/pprof"  // as a side effect, registers the pprof endpoints.
    "time"

    "myproject.org/blobstore"
)

func main() {
    for i := 0; i < 1000; i ++ {
        name := fmt.Sprintf("task-blob-%d", i)
        go func() {
            b, err := blobstore.Open(name)
            if err != nil {
                // TODO: Handle error.
            }
            defer b.Close()
            // TODO: Perform some work, write to the blob.
        }()
    }
    http.ListenAndServe("localhost:8080", nil)
}

```

开启服务器，然后使用`go tool`工具来读取和可视化该分析数据：

```sh

go tool pprof http://localhost:8080/debug/pprof/blobstore.Open

```

您可能会看到有800个开启的blob，所有的开启事件都来自`main.main.func1`。在这个小示例中，没有什么可看的了，但是在复杂的服务器中，您可以检查与开启的blob一起工作的最热点，并找出瓶颈或泄漏。

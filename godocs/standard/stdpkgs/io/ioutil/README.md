

# `ioutil`包

`ioutil`包实现了一些`I/O`实用函数。


## 函数

```go

func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, pattern string) (name string, err error)
func TempFile(dir, pattern string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm os.FileMode) error

```


## 变量

```go

// `Discard`是一个`io.Writer`，所有调用成功的`Write`均不执行任何操作。
var Discard io.Writer = devNull(0)

```



# `path`包

`path`包实现了对`斜杠分隔的路径`的实用操作函数。

`path`包应该仅为`以正斜杠分隔的路径`所使用，例如`URL中的路径`。该包`不处理带有驱动器号或反斜杠的Windows路径`。要处理操作系统路径，请使用`path/filepath`包。


## 接口


## 结构体


## 函数

```go

func Base(path string) string
func Clean(path string) string
func Dir(path string) string
func Ext(path string) string
func IsAbs(path string) bool
func Join(elem ...string) string
func Match(pattern, name string) (matched bool, err error)
func Split(path string) (dir, file string)

```


## 变量

```go

var ErrBadPattern = errors.New("syntax error in pattern")

```

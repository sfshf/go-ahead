

# `path/filepath`包

`path/filepath`包实现了兼容各操作系统的文件路径的实用操作函数。

`path/filepath`包使用`正斜杠`或`反斜杠`，具体取决于操作系统。要处理`诸如URL之类的路径`，无论使用什么操作系统，该路径始终使用`正斜杠`；请参阅`path`包。


## 接口


## 数据结构

```go

type WalkFunc func(path string, info os.FileInfo, err error) error

```


## 函数

```go

func Abs(path string) (string, error)
func Base(path string) string
func Clean(path string) string
func Dir(path string) string
func EvalSymlinks(path string)(string, error)
func Ext(path string) string
func FromSlash(path string) string
func Glob(pattern string) (matches []string, err error)
func HasPrefix(p, prefix string) bool  // RESERVED
func IsAbs(path string) bool
func Join(elem ...string) string
func Match(pattern, name string) (matched bool, err error)
func Rel(basepath, targpath string) (string, error)
func Split(path string) (dir, file string)
func SplitList(path string) []string
func ToSlash(path string) string
func VolumeName(path string) string
func Walk(root string, walkFn WalkFunc) error

```


## 常量

```go

const (
    Separator = os.PathSeparator
    ListSeparator = os.PathListSeparator
)

```


## 变量

```go

var ErrBadPattern = errors.New("syntax error in pattern")

var SkipDir = errors.New("skip this directory")

```

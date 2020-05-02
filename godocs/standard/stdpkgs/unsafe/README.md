

# `unsafe`包

`unsafe`包提供了一些`跳过go语言类型安全限制的操作`。

导入`unsafe`包的软件包`可能是不可移植的`，并且`不受Go 1兼容性准则的保护`。


## 结构体

```go

type ArbitraryType int

type Pointer *ArbitraryType

```


## 函数

```go

func Alignof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Sizeof(x ArbitraryType) uintptr

```

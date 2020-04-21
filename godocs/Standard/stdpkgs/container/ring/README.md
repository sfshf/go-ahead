

# `container/ring`包

`container/ring`包实现了`环形链表`的操作。


## 接口


## 结构体

```go

type Ring struct {
    next, prev *Ring
    Value interface{}  //for use by client; untouched by this library
}

```


## 函数

```go

func New(n int) *Ring

```

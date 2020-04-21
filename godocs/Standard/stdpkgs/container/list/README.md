

# `container/list`包

`container/list`包实现了了一个`双向链表（doubly linked list）`。


## 接口


## 结构体

```go

type Element struct {
    next, prev *Element
    list *List
    Value interface{}
}

type List struct {
    root Element
    len int
}

```


## 函数

```go

func New() *List

```

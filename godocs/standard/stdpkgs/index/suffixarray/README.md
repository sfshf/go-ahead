

# `index/suffixarray`包

`index/suffixarray`包通过使用内存中的`后缀树`实现了对`数级时间消耗`的`子字符串搜索`。


## 结构体

```go

type Index struct {

    data []byte
    sa ints  // suffix array for data; sa.len() == len(data)

}

```


## 函数

```go

func New(data []byte) *Index

```



# `container/heap`包

`container/heap`包提供了对任意类型（实现了`heap.Interface`接口）的`堆`操作。（最小）`堆`是具有`每个节点都是以其为根的子树中最小值`属性的`树`。

`树`的最小`元素`为其根元素，索引`0`的位置。

`heap`是常用的`实现优先队列`的方法。要创建一个`优先队列`，实现一个具有使用（负的）优先级作为比较的依据的`Less`方法的`Heap`接口，如此一来可用`Push`添加项目而用`Pop`取出队列最高优先级的项目。


## 接口

```go

type Interface interface {
    sort.Interface
    Push(x interface{})  // add x as element Len()
    Pop() interface{}    // remove the return element Len() - 1.
}

```


## 结构体


## 函数

```go

func Fix(h Interface, i int)
func Init(h Interface)
func Pop(h Interface) interface{}
func Push(h Interface, x interface{})
func Remove(h Interface, i int) interface{}

```

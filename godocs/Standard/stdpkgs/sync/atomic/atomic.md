

# `sync/atomic`包

`sync/atomic`包提供了低层的`原子级内存原语`，`对于同步算法的实现很有用`。

这些函数必须谨慎地保证正确使用。除了某些特殊的低层应用，使用`通道`或者`sync`包的功能设施来实现同步更好。`应通过通信来共享内存，而不通过共享内存实现通信。`

被`SwapT`系列函数实现的`交换操作`，在原子性上等价于：

```go

old = *addr
*addr = new
return old

```

`CompareAndSwapT`系列函数实现的`比较-交换操作`，在原子性上等价于：

```go

if *addr == old {
	*addr = new
	return true
}
return false

```

`AddT`系列函数实现`加法操作`，在原子性上等价于：


```go

*addr += delta
return *addr

```

`LoadT`和`StoreT`系列函数实现的`加载和存储操作`，在原子性上等价于：`"return *addr"`和`"*addr = val"`。


## 结构体

```go

type Value struct {
    v interface{}
}

```


## 函数

```go

func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)

//==============================================================================

func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)

//==============================================================================

func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)

//==============================================================================

func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)

//==============================================================================

func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

```

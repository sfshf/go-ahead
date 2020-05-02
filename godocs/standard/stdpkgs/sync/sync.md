

# `sync`包

`sync`包提供了基本的`同步`原语，如`互斥锁（mutual exclusion locks）`。除了`Once`和`WaitGroup`类型，大部分都供低层库例程使用。更高层的同步最好通过`通道（channel）`和`通信（communication）`来完成。

包含此包中定义的类型的值不应被复制。


## 核心代码解析


### 1、`Locker`接口

```go

type Locker interface {
    Lock()
    Unlock()
}

```

`Locker`，表示可以被锁定和被解锁的对象。


### 2、`Mutex`结构体

```go

type Mutex struct {
    state int32
    sema uint32
}

```

`Mutex`，是`互斥锁（mutual exclusion lock）`。`Mutex`的`零值`是`未锁定的互斥锁`。

`Mutex`决不能在首次使用后复制。


#### `Mutex`的方法

**func (m *Mutex) Lock() { ... }**

`Lock`会锁住`m`。如果`锁`已在使用中，则正想要调用它的goroutine会阻塞，直到`互斥锁`可用为止。

**func (m *Mutex) Unlock() { ... }**

`Unlock`解锁`m`。如果`m`在进行`Unlock`时未锁定，则是`运行时错误`。

`被锁定的Mutex与特定的goroutine没有关联`。允许一个goroutine锁定`Mutex`，然后安排另一个goroutine对其进行解锁。


### 3、`RWMutex`结构体

```go

type RWMutex struct {
    w Mutex             // held if there are pending writers.
    writerSem uint32    // semaphore for writers to wait for completing readers.
    readerSem uint32    // semaphore for readers to wait for completing writers.
    readerCount uint32  // number of pending readers.
    readerWait int32    // number of departing readers.
}

```

`RWMutex`是`读取器/写入器互斥锁`。锁可以由`任意数量`的读取器或`单个`写入器持有。`RWMutex`的`零值`是`未锁定的互斥锁`。

`RWMutex`在第一次使用后不能被复制。

当一个goroutine持有`RWMutex`进行读取而另一个goroutine可能调用了`Lock`时，则在释放`初始读取锁`之前，任何goroutine都不应该能够获取`读取锁`。尤其是，要禁用`递归读取锁`。这是为了确保`锁`最终可用；被阻塞的`Lock`调用将使`新的读取器`无法获得`锁`。


#### `RWMutex`的方法

**func (rw *RWMutex) Lock() { ... }**

`Lock`为`写入`锁住`rw`。如果该锁`已经`为读取或写入被锁住，则`Lock`会阻塞直到该锁可用。

**func (rw *RWMutex) Unlock() { ... }**

`Unlock`为`写入`解锁`rw`。如果`rw`在为`写入`进行`Unlock`时未锁定，则是`运行时错误`。

与`Mutex`一样，被锁住的`RWMutex`与特定的goroutine没有关联。一个goroutine可以`RLock`（`Lock`）一个`RWMutex`，然后安排另一个goroutine来`RUnlock`（`Unlock`）它。

**func (rw *RWMutex) RLock() { ... }**

`RLock`为`读取`锁住`rw`。

`不应将其用于递归读取锁定`。被阻塞的`Lock`调用将使新读取器无法获得`锁`。

**func (rw *RWMutex) RUnlock() { ... }**

`RUnlock`撤销单个`RLock`调用；它不会影响其他同时的读取器。如果`rw`在为读取进行`RUnlock`时未被锁住，则是`运行时错误`。

**func (rw *RWMutex) RLocker() Locker { ... }**

`RLocker`返回一个`Locker接口`，该接口通过调用`rw.RLock`和`rw.RUnlock`来实现`Lock`和`Unlock`方法。


### 4、`WaitGroup`结构体

```go

type WaitGroup struct {
    noCopy noCopy

    // 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
    //64-bit atomic operations require 64-bit alignment, but 32-bit compilers do not ensure it. So we allocate 12 bytes and then use the aligned 8 bytes in them as state, and the other 4 as storage for the sema.
    state1 [3]uint32
}

```

`WaitGroup`等待`goroutine的集合`完成。`主goroutine`调用`Add`以设置要等待的goroutine的数量。然后，每个goroutine都会运行并在完成后调用`Done`。同时，可以使用`Wait`来阻塞，直到所有goroutine完成。

`WaitGroup`在首次使用后不能被复制。


#### `WaitGroup`的方法

**func (wg *WaitGroup) Add(delta int) { ... }**

`Add`将`delta（可能为负数）`添加到`WaitGroup`计数器中。如果计数器`变为零`，则释放`Wait`阻塞的所有goroutine。如果计数器`变为负数`，请添加`panic`。

请注意，当计数器为零时发生的增量为正的调用必须在等待之前发生。 在计数器大于零时开始的负增量呼叫或正增量呼叫可能随时发生。 通常，这意味着对Add的调用应在创建goroutine或要等待的其他事件的语句之前执行。 如果重新使用WaitGroup来等待几个独立的事件集，则必须在所有先前的Wait调用返回之后再进行新的Add调用。 请参阅WaitGroup示例。

**func (wg *WaitGroup) Done() { ... }**

**func (wg *WaitGroup) Wait() { ... }**


### 5、`Pool`结构体

### 6、`Cond`结构体


### 7、`Map`接欧体

### 8、`Once`结构体

### 9、

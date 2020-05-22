
## package signal（signal包）

`signal包`实现了对传入的`信号`的访问。

`信号`主要用于`类Unix系统`。有关在`Windows`和`Plan 9`上`signal包`的用法，请参见下文。


### Types of signals（信号的类型）

信号`SIGKILL`和`SIGSTOP`可能不会被程序捕获，因此不会受到`signal包`的影响。

`同步信号`是由程序执行中的错误触发的信号：`SIGBUS`，`SIGFPE`和`SIGSEGV`。这些信号在`仅`由`程序执行`引起时才被认为是同步的，在使用`os.Process.Kill`或`kill程序`或`某种类似机制`发送时不会被视为同步。通常，除以下讨论的内容外，Go程序会将`同步信号`转换为`运行时宕机`。

`其余信号`是`异步信号`。它们`不是`由`程序错误`触发的，`而是`从`内核`或`其他程序`发送的。

在`异步信号`中，当程序丢失其`控制终端`时，将发送`SIGHUP`信号。当`控制终端`上的用户按下`中断字符`时，将发送`SIGINT`信号，默认情况下为`^C(Control-C)`。当`控制终端`上的用户按下`退出字符`时，将发送`SIGQUIT`信号，默认情况下为`^\(Control-Backslash)`。通常，您可以通过按`^C`简单地退出程序，还可以通过按`^\`使其堆栈溢出来退出。


### Defult behavior of signals in Go programs（Go程序中信号的默认行为）

默认情况下，`同步信号`会转换为`运行时宕机`。`SIGHUP`，`SIGINT`或`SIGTERM`信号导致程序退出。 `SIGQUIT`，`SIGILL`，`SIGTRAP`，`SIGABRT`，`SIGSTKFLT`，`SIGEMT`或`SIGSYS`信号会导致程序堆栈溢出而退出。`SIGTSTP`，`SIGTTIN`或`SIGTTOU`信号获得系统默认行为（shell将这些信号用于`作业控制`）。`SIGPROF`信号由`Go运行时`直接处理以实现`runtime.CPUProfile`。其他信号将被捕获，但不会采取任何措施。

如果Go程序在忽略`SIGHUP`或`SIGINT`的情况下启动（即`信号处理程序`设置为`SIG_IGN`），则它们将保持被忽略状态。

如果Go程序以`非空`的`信号掩码`启动，通常会被接受。但是，某些信号被明确解除阻塞：同步信号`SIGILL`，`SIGTRAP`，`SIGSTKFLT`，`SIGCHLD`，`SIGPROF`，以及在GNU/Linux上，信号`32`（`SIGCANCEL`）和`33`（`SIGSETXID`）（`glibc`在内部使用了`SIGCANCEL`和`SIGSETXID`）。由`os.Exec`或`os/exec`程序包启动的子进程将继承`修改后的信号掩码`。


### Changing the behavior of signals in Go programs（在Go程序中更改信号的行为）

`os/signal`包中的函数允许程序更改Go程序处理信号的方式。

`Notify``禁用`一组给定`异步信号`的`默认行为`，而是通过一个或多个`已注册通道`发送它们。具体来说，它适用于信号`SIGHUP`，`SIGINT`，`SIGQUIT`，`SIGABRT`和`SIGTERM`。它也适用于作业控制信号`SIGTSTP`，`SIGTTIN`和`SIGTTOU`，不会发生系统默认行为。它也适用于其他没有默认行为的信号：SIGUSR1，SIGUSR2，SIGPIPE，SIGALRM，SIGCHLD，SIGCONT，SIGURG，SIGXCPU，SIGXFSZ，SIGVTALRM，SIGWINCH，SIGIO，SIGPWR，SIGSYS，SIGINFO，SIGTHR，SIGWAITING，SIGLWP，SIGFREEZE，SIGTHAW，SIGLOST，SIGXRES，SIGJVM1，SIGJVM2，以及系统上使用的任何`实时信号`。请注意，并非所有这些信号在所有系统上都可用。

如果程序是在`忽略SIGHUP或SIGINT`的情况下启动的，并且为每个信号调用了`Notify`，则将为该信号安装信号处理程序，并且将不再忽略该信号处理程序。如果稍后再对该信号调用`Reset`或`Ignore`，或者在传给该信号的`Notify`的所有`通道`上调用`Stop`，则该信号将再次被忽略。`Reset`将恢复`信号的系统默认行为`，而`Ignore`将导致系统完全忽略信号。

如果程序以`非空信号掩码`启动，则如上所述，某些信号将被`显式解除阻塞`。如果为阻塞的信号调用了`Notify`，则该信号变为`非阻塞`。如果稍后再对该信号调用`Reset`，或者在传给该信号的`Notify`的所有通道上对该信号调用`Stop`，则该信号将再次被阻塞。


### SIGPIPE

当Go程序写入`损坏的管道`时，`内核`将引发`SIGPIPE`信号。

如果程序尚未调用`Notify`来接收`SIGPIPE`信号，则其行为取决于`文件描述符号`。从文件描述符`1`或`2`（`标准输出`或`标准错误`）上写入`损坏的管道`将导致程序以`SIGPIPE`信号退出。从其他文件描述符上写入`损坏的管道`将不对`SIGPIPE`信号执行任何操作，并且写入将失败并显示`EPIPE`错误。

如果程序已调用`Notify`来接收`SIGPIPE`信号，则文件描述符号将无关紧要。`SIGPIPE`信号将传递到`Notify`通道，并且写入将失败，并出现`EPIPE`错误。

这意味着默认情况下，`命令行程序`的行为将类似于`典型的Unix命令行程序`，而其他程序会在向关闭的网络连接写入时因`SIGPIPE`信号而崩溃。


### Go programs that use cgo or SWIG（使用cgo或SWIG的Go程序）

在包含`非Go代码`（通常是使用`cgo`或`SWIG`访问的`C/C++`代码）的`Go程序`中，`Go的启动代码`通常会首先运行。在`非Go启动代码`运行之前，它将按`Go运行时`的预期配置`信号处理程序`。如果`非Go启动代码`希望安装自己的`信号处理程序`，则必须采取某些步骤才能使Go正常运行。本节记录了这些步骤，并且`非Go代码`可以在`Go程序`上对`信号处理程序`设置进行整体更改。在极少数情况下，`非Go代码`可能会在Go代码之前运行，这种情况在下一节讲述。

如果`Go程序`调用的`非Go代码`没有更改任何`信号处理程序`或`掩码`，则其行为与`纯Go程序`相同。

如果`非Go代码`安装了任何`信号处理程序`，则必须使用带有`sigaction`的`SA_ONSTACK`标识。否则，如果收到信号，很可能导致程序崩溃。`Go程序`通常以`有限的堆栈`运行，因此设置了`备用`的`信号堆栈`。此外，`Go标准库`希望所有`信号处理程序`都将使用`SA_RESTART`标志。否则，可能导致某些库调用返回`"interrupted system call"`错误。

如果`非Go代码`为任何`同步信号`（`SIGBUS`，`SIGFPE`，`SIGSEGV`）安装了`信号处理程序`，则它应记录`现有`的`Go信号处理程序`。如果在执行Go代码时发生这些信号，则应调用`Go信号处理程序`（在执行Go代码时是否发生信号取决于被传给信号处理函数的PC）。否则，某些`Go运行时宕机`将不会按预期发生。

如果`非Go代码`为任何`异步信号`安装了`信号处理程序`，则它可能会选择调用`Go信号处理程序`，也可能不会调用它。通常，如果它不调用`Go信号处理程序`，则不会发生上述的Go行为。特别是这可能与`SIGPROF`信号有关。

`非Go代码`不应在`Go运行时`创建的任何线程上更改`信号掩码`。如果`非Go代码`启动了自己的新线程，则可以根据需要设置`信号掩码`。

如果`非Go代码`启动一个新线程，更改`信号掩码`，然后在该线程中调用Go函数，则`Go运行时`将自动取消阻塞某些信号：`同步信号`，`SIGILL`，`SIGTRAP`，`SIGSTKFLT`，`SIGCHLD`，`SIGPROF`，`SIGCANCEL`和`SIGSETXID`。当执行`Go函数`返回时，将恢复`非Go的信号掩码`。

如果在`未运行Go代码的非Go线程`上调用`Go信号处理程序`，则`该处理程序`通常将信号`转发`到非`Go代码`，如下所示。如果信号是`SIGPROF`，则`Go处理程序`不执行任何操作。否则，`Go处理程序`将删除自身，解除阻塞信号，然后再次引发信号，以调用任何`非Go处理程序`或`默认系统处理程序`。如果程序没有退出，则`Go处理程序`将重新安装自身并继续执行程序。


### Non-Go programs that call Go code（调用Go代码的非Go程序）

当使用`-buildmode=c-shared`之类的选项`构建`Go代码时，它将作为现有`非Go程序`的一部分运行。启动`Go代码`时，`非Go代码`可能已经安装了`信号处理程序`（在使用`cgo`或`SWIG`的异常情况下也可能会发生；在这种情况下，此处的讨论可适用）。对于`-buildmode=c-archive`，`Go运行时`将在`全局构造函数`时`初始化信号`。对于`-buildmode=c-shared`，当加载共享库时，`Go运行时`将初始化信号。

如果`Go运行时`看到`SIGCANCEL`或`SIGSETXID`信号的`现有信号处理程序`（`仅`在GNU/Linux上使用），它将打开`SA_ONSTACK`标志，否则`保留``信号处理程序`。

对于`同步信号`和`SIGPIPE`，`Go运行时`将安装`信号处理程序`。它将`保存`任何`现有的信号处理程序`。如果在执行`非Go代码`时`同步信号`到达，则`Go运行时`将调用`现有的信号处理程序`，而不是`Go信号处理程序`。

默认情况下，使用`-buildmode=c-archive`或`-buildmode=c-shared`构建的Go代码不会安装任何其他信号处理程序。如果存在`现有的信号处理程序`，则`Go运行时`将打开`SA_ONSTACK`标志，否则将`保留``信号处理程序`。如果为`异步信号`调用了`Notify`，则将为该信号安装`Go信号处理程序`。如果稍后再对该信号调用`Reset`，则将重新安装`该信号的原始处理`，以恢复`non-Go信号处理程序`（如果有）。

在不使用`-buildmode=c-archive`或`-buildmode=c-shared`的情况下构建的`Go代码`将为上面列出的`异步信号`安装`信号处理程序`，并保存任何`现有的信号处理程序`。如果将信号传递到`非Go线程`，则它将如上所述执行操作，不同之处在于，如果存在`现有的非Go信号处理程序`，则会在发出信号之前安装该处理程序。


### Windows

在Windows上，`^C（Control-C）`或`^BREAK（Control-Break）`通常会导致程序退出。如果为`os.Interrupt`调用了`Notify`，则`^C`或`^BREAK`将导致`os.Interrupt`在通道上发送，并且程序不会退出。如果调用了`Reset`，或者在传递给`Notify`的所有通道上调用了`Stop`，则将恢复默认行为。

此外，如果调用了`Notify`，并且Windows将`CTRL_CLOSE_EVENT`，`CTRL_LOGOFF_EVENT`或`CTRL_SHUTDOWN_EVENT`发送到该进程，则`Notify`将返回`syscall.SIGTERM`。与`Control-C`和`Control-Break`不同，当收到`CTRL_CLOSE_EVENT`，`CTRL_LOGOFF_EVENT`或`CTRL_SHUTDOWN_EVENT`时，`Notify`不会更改进程行为--除非退出，否则该进程仍将终止。但是接收`syscall.SIGTERM`将使进程有机会在终止之前进行清理。


### Plan 9

在Plan 9上，`信号的类型`为`syscall.Note`，这是一个`字符串`。使用`syscall.Note`类型调用`Notify`时会导致，当字符串作为消息发出时该值将在通道上发送。

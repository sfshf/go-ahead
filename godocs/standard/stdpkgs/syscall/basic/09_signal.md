
# 信号处理


## 概述

`信号`是一种`软中断`。它提供了一种通知进程某个事件发生的机制，是最简单的进程间通信的方式。它让一个进程被另一个进程（或内核）`异步地`中断处理某个事件，处理结束后，被中断的进程从中断点重新执行。`信号机制`是系统中关于信号`产生`、`传送`和`处理`的一套机构。

`信号机制`最早在`Unix SYSTEMV`中出现，它是不可靠的。`Unix BAD4.2`改善了`SYSTEMV`的信号机制，但与`SYSTEMV`的信号机制不兼容。以后，在它们各自派生出的各种UNIX系统中，这种不兼容性延续了下来。

为了给程序员提供一个一致的编程接口，`POSIX1003.1标准【IEEE90】`制定了一个规范，定义了一个`标准的信号接口`，所有UNIX版本的信号机制都必须支持该接口。Linux的信号机制也同样支持该接口。

在一个`信号的生命过程`中有`两个`阶段：`生成`和`传送`。当事件发生时，如果需要通知某个进程，内核就会产生一个信号。而当该进程发现信号到来，就以`预先定义好的方式`来处理信号（执行信号处理函数）。在`信号产生`和`对它进行处理`之间，`信号被"挂起"（pending）`。需要注意的是，信号`必需`有其`目的进程`，即它`必需`被发送给一个进程。

当出现下面几种情况时，可以产生信号。

（1）异常。当一个进程出现异常（如试图执行非法指令），内核通过向进程发送一个消息通知该事件的发生。

（2）用户用某些键（如`Ctrl+c`、`Ctrl+d`）对终端的控制。

（3）用户态进程之间的通信。用户可以使用`kill`命令向进程发信号，进程可以通过`kill`系统调用向另一进程发信号。

（4）进程等待的某个事件的发生，如I/O设备的就绪。

（5）进程可以调用某些函数向自己发送信号。如`abort`函数向它的调用者发送`SIGABRT`信号。

每个信号有一个缺省动作，当用户进程没有给这个信号规定处理程序时，内核执行缺省动作。有以下几种缺省动作：

（1）异常终止（`abort`）。在当前目录生成一个`core`映像文件后，终止进程。它以后可以被调试工具使用。

（2）退出（`exit`）。不产生`core`映像文件就终止进程。

（3）忽略（`ignore`）。忽略该信号。

（4）停止（`stop`）。挂起进程。

（5）继续（`continue`）。如果进程挂起，则恢复进程运行，否则忽略该信号。

与`硬件中断`相类似，一个信号也可以被屏蔽。如果传送给进程的信号已经被进程所屏蔽，则该信号被记录在进程的待处理信号集合中。在清除屏蔽前，进程对信号不会采取任何动作。Linux对中断的屏蔽遵循POSIX的规范。

系统不允许进程屏蔽信号`SIGKILL`、`SIGSTOP/SIGCONT`。


## Linux系统中的信号

Linux系统用一些常用表示信号，如下表所示。信号定义在头文件`/usr/include/bits/signum.h`中。

![Linux系统中的信号](imgs/signum.jpeg)

详细说明如下：

- `SIGHUP`：当终端断连时，如果它的`CLOCAL`标识没有被置位，这个信号被传给`终端`的`控制进程`（主会话进程）。如果主会话进程已退出，则该信号被送给该会话的每一个进程组的主进程。大多数进程接收到`SIGHUP`后终止，因为它表示用户已不在。该信号常被用来通知`守护进程`去重读它们的配置文件。因为守护进程没有自己的控制终端，所以只有用这种方法通知它们。这个信号同其他的终端结束信号不一样，那些信号通常只送给前台进程组。
- `SIGINT`：当用户按下`Ctrl+c`等`中断键`时，该信号被送给`前台进程组`中`所有的进程`。它常被用来`中断当前进程`。
- `SIGQUIT`：当用户按下`Ctrl+\`等`退出键`时，该信号被送给前台进程组中所有的进程。
- `SIGTRAP`：当进程遇到一个断点时，该信号被传给进程。`SIGTRAP`常被设置断点的调试器使用。
- `SIGABRT`：进程通过系统调用`abort`向自己发送该信号。它的缺省动作是使进程异常终止。
- `SIGBUS`：当一个进程违反了除了存储保护以外的其他的硬件限制时，该信号被发送。
- `SIGFPE`：当出现`除0`错、浮点运算下溢时，该信号被发送。
- `SIGALRM`：进程用系统调用`alarm`设置定时报警器，当经过指定的时间后，系统将该信号发送给进程。
- `SIGTERM`：`kill`命令产生该信号，进程在收到它后，必需迅速退出。
- `SIGCLD/SIGCHLD`：当进程的一个子进程退出或停止时，系统向它发送这个信号。在该信号的处理函数里，可以调用`wait函数`，以避免子进程僵死。如果进程用`wait`等待它的子进程结束，可以忽略该信号。
- `SIGSTOP`：作业控制信号，使进程无条件地停止。该信号与`SIGKILL`信号是不能被忽略地两个信号。
- `SIGTSTP`：当用户按下`暂停键`（如`Ctrl+z`），这个信号被送给`前台进程组`中的`所有进程`。常用于`作业控制`中。
- `SIGTTIN`：当一个`后台进程`试图从它的控制终端`读`的时候，核心向它发送这个信号。常用于`作业控制`中。
- `SIGTTOU`：当一个`后台进程`试图向它的控制终端`写`的时候，核心向它发送这个信号。常用于`作业控制`中。
- `SIGURG`：当从`socket`收到带外数据时，该信号被发送。
- `SIGVTALRM`：进程调用`setitimer`设置定时器，当经过规定的时间后，核心向进程发送该信号。
- `SIGPROF`：统计`信息计时器`的`时间到期`。该信号常被`统计信息程序`使用，这种程序检查`别的进程`的`运行时`的特性。我们可以通过这种程序找到进程执行的瓶颈，再优化进程的执行。
- `SIGWINCH`：当终端和伪终端的窗口大小改变时，核心将该信号发给终端的前台进程所有的进程。该信号用于`X Window系统`中。
- `SIGPOLL/SIGIO`：`异步I/O事件`发生。
- `SIGPWR`：当系统检测到`电源`将要出现问题时，`电源管理的守护进程`把该信号发给`初始进程init`，使机器能在断电前完成正常关机的各种步骤。


## 对信号的处理

如果一个进程没有屏蔽某个`信号`，当核心向它发送信号时，`信号处理函数`被调用。在Linux中，有几种方法可以给信号设置处理函数。本章先介绍比较传统的处理方法，并分析它们有可能产生的问题，然后再介绍POSIX规范的信号处理。


### 设置信号处理函数

最简单的设置信号处理函数的方式是使用系统调用`signal`。它的原型如下：

```c

#include <signal.h>
void (* signal(int signum, void (* handler)(int))) (int)

```

说明：系统调用`signal`有两个参数，`signum`和`handler`。

`signum`是信号名，即上一个表格中的常量，`handler`是新的信号处理函数的指针，它有一个参数，即激活它的信号。

返回：系统调用`signal`返回原来的信号处理函数的指针。

在第一次调用时，`signal`的返回值是缺省的信号处理函数的指针，`SIG_IGN（忽略）`，`SIG_DFL（缺省）`、`SIG_ERR（错误）`。这三个常量也在`/usr/include/bits/signum.h`中定义。

```c
typedef void (* _ _ sighandler_t) (int)
#define	SIG_ERR	 ((_ _ sighandler_t) -1)	/* Error return.  */
#define	SIG_DFL	 ((_ _ sighandler_t)  0)	/* Default action.  */
#define	SIG_IGN	 ((_ _ sighandler_t)  1)	/* Ignore signal.  */

```

对比地，`Go语言`的`syscall`包内提供了如下内容：

```go

// `Signal`是描述进程信号的数值。它实现了`os.Signal`接口。
type Signal int

func (s Signal) Signal()
func (s Signal) String() string

```

`os`包内提供了如下内容：

```go

// `Signal`代表操作系统的信号。通常底层实现依赖于操作系统：在Unix系统上的实现即`syscall.Signal`。
type Signal interface {
    String() string
    Signal()        // 为了与其他字符串生成器进行区分
}

// ...
var (
    Interrupt Signal = syscall.SIGINT
    Kill Signal = syscall.SIGKILL
)

```

`os/signal`包内封装了`Go语言`里信号机制的一些使用API。详见源码文档。

下面是两个示例代码：

```go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//fmt.Fprintln(os.Stdout, os.Getpid())
	curProc, err := os.FindProcess(os.Getpid())
	if err != nil {
		panic(err)
	}

    // 模拟主进程。主进程向子进程发送信号。
	go func(proc *os.Process) {

		time.Sleep(2*time.Second)

		err := proc.Signal(syscall.SIGUSR1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		}

		err = proc.Signal(syscall.SIGUSR2)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		}

		err = proc.Signal(os.Interrupt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		}

		time.Sleep(2*time.Second)
		err = proc.Signal(os.Kill)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		}

	}(curProc)

    // 模拟子进程。子进程接收信号，并进行处理。
	//ch := make(chan os.Signal)  
    //ch := make(chan os.Signal, 1)  
    ch := make(chan os.Signal, 3)  

	for {

		signal.Notify(ch)
		s := <- ch
        // 模拟信号处理函数
		fmt.Fprintf(os.Stdout, "Got signal: %v\n", s)

	}

}

```

示例代码：

```go

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	ch := make(chan os.Signal, 1)
	for {
		signal.Notify(ch)
		s := <- ch
		fmt.Fprintf(os.Stdout, "Got signal: %v\n", s)
		if s == os.Kill {
			os.Exit(0)
		}
	}

}

// Test it with:
// $ go build main.go
// $ ./main &
// $ kill -s SIGUSR1 <pid>
// $ kill <pid> （即kill -s TERM <pid>）
// $ kill -s KILL <pid>

```


### 系统对信号的处理

在缺省情况下，一个进程的信号处理是`default`或`ignore`。如果一个进程设置了自己的`信号处理函数`，然后又用`exec`加载运行`另一个程序`，则由于`新老程序`共用一个`进程空间`，而`老程序`的`进程上下文`被覆盖，因此对这些信号的设置`无效`。

对比地，由于`Go语言`的实现方式不一样，所以没有上述的影响。示例代码如下：

```go

// test.go
package main

func main() {
	for {}
}

// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1)

	go func(ch <-chan os.Signal) {
		s := <- ch
		fmt.Fprintf(os.Stdout, "Got signal: %v\n", s)
	}(ch)

	var attr os.ProcAttr
	_, err := os.StartProcess("./test", os.Args, &attr)
	if err != nil {
		panic(err)
	}
	for {}

}

// Test it with:
// $ go build test.go
// $ go build main.go
// $ ./main &
// $ kill -s SIGUSR1 <pid>
// $ kill <pid> （即kill -s TERM <pid>）
// $ kill -s KILL <pid>

```


### 不可靠的信号

当一个信号被送给进程，而该进程正在这个信号的处理函数中运行，系统将怎样处理呢？最简单的方法是：核心中断当前的处理程序并重新运行。但这会带来问题：当信号处理函数处理一些`全局数据结构`或`文件`时，容易出现错误。对于用户来说，可以用`加锁`的方法来`避免敏感数据的重复读写`。当`信号处理函数`第一次运行的时候，它锁住一个文件，当它又一次被调用的时候，已经不能再对这个文件加锁，因此它不能继续运行。当然最后的方法是编写`可重入`的`信号处理函数`。但这两种方法都比较麻烦，而且并`没有`提供`系统级`的控制。进程无法在`短时间`内处理`大量的`信号，因为`大量频繁的函数的调用`有可能使得`进程的栈溢出`。

对这一问题的最早的处理是：在`信号处理函数`被调用之前，系统自动把它的处理函数设为`SIG_DFL`，而在`信号处理函数`中，以适当的方式`重设`处理函数。尽管这使得编写信号处理函数很容易，但也使得`对信号的可靠处理`变得非常困难。考虑以下情况，`几乎同时`发生了`两个信号事件`，`系统`把信号处理函数`置为缺省`，然后调用`用户设置的`信号处理函数。如果在信号处理函数`重置`自己之前，系统根据`缺省的行为`来处理`第二个信号`，它实际上是被忽略掉或使进程结束。这样的信号系统被称为`不可靠的`。

`BSD4.2`提出了一种对于多个信号的简单的实现方法，那就是：系统`等待`第一个信号处理函数结束，`再发送`第二个信号。这种方法`保证了`信号不会丢失，而且`也避免了`用户进程的栈溢出。当核心保留一个信号以后再发送，这个信号被称为`悬挂（pending）`。

然而，当一个信号在悬挂的时候，如果它又一次产生，进程`只会`收到一个这样的信号。无法使进程得知有多少个信号发送给了它。一般来说，这不会产生太大的问题，因为`信号通常不带有任何信息`，在`短时间内连续发送两个信号`和`发送一个信号`没有多大的区别。只有在一种情况（即对第二个信号采用缺省的处理函数）下，才会导致错误。

在`BSD4.3`中，信号处理程序的建立是永久的，不需要每次重置。

在`BSD4.3`中，当调用用户的信号处理程序时，当前正在处理的信号自动地被屏蔽。

在`以前的Unix版本`中，每次只能屏蔽一个信号，因此，在一个信号传到进程之前，无法保证一组信号被屏蔽。在`BSD4.3`中，可以通过系统调用一次屏蔽若干个信号。

在`某些版本的Unix`中，当进程在一个`较慢的系统调用`中等待时（如网络或I/O操作），`信号的产生`会`中断`系统调用，使系统调用返回一个错误，把`errno`设为`EINTR`。这样做的`目的是`使`长时间等待的进程`有机会执行。在`BSD4.3`中，`被一个信号中断的`系统调用，可以`自动重新启动`，这使得用户的程序不必在每次系统调用返回后去检查返回值以判定这次调用是否应该再次执行。

在`以前的Unix版本`中，信号总是被传到一个`进程的用户栈`上，在`BSD4.3`中，通过`sigstack`系统调用可以为信号的传送指定`另外一个栈`。

除了`信号栈`以外，POSIX继承了`BSD4.3`的信号机构。


### 信号的阻塞

`信号的产生（generate）`是指导致信号的某些`事件`发生。

信号在产生后，系统将它`传递（deliver）`给`进程`。在产生和传递之间，信号被`悬挂（impending）`。

进程可以`通过屏蔽`来`阻塞（blocking）`信号的传递。如果一个`被阻塞的信号`产生了，而且`对信号的处理`不是`忽略`，则信号会`一直悬挂到解除阻塞`或者`信号处理函数被设成忽略`。需要注意的是，系统在`信号的传递`而不是`信号的产生`的时候决定如何处理`被阻塞的信号`，这就允许进程`在信号产生之后传递之前（即阻塞的时候）`改变它的处理函数。进程可以通过调用函数`sigpending`来获悉哪些信号被阻塞。

如果一个`阻塞的信号`在进程解除其阻塞之前再次产生，`POSIX.1`允许系统把它传送一次或两次。当信号被传递两次或两次以上时，称为`信号排队`。大部分的Unix系统`不支持`信号排队。

每个进程有一个`信号屏蔽（signalmask）`，它指出了当前哪些信号被屏蔽。进程可以通过调用`sigprocmask`获悉和改变信号屏蔽。


### 向进程发送信号

一般可以通过`kill`和`raise`向进程发送信号。

```c

#include <sys/types.h>
#include <signal.h>

int kill(int pid, int signo);

int raise(int signo);

```

`kill`说明：

（1）如果`pid>0`，信号被送给进程标识为`pid`的进程。

（2）如果`pid=0`，信号被送给`发送者的进程组`中的`所有进程`，只要发送者有权向该进程发送信号。这些进程不包括`交换进程`（`pid=0`），`初始进程`（`pid=1`），`页守护进程`（`pid=2`）。

（3）如果`pid<0`，信号被送给所有`组ID`等于`pid`的绝对值的进程，只要发送者有权向该进程发送信号。同样，这些进程不包括`交换进程`（`pid=0`），`初始进程`（`pid=1`），`页守护进程`（`pid=2`）。

（4）如果`pid=-1`，如果调用者是`超级用户`，把信号按照`pid`从大到小的顺序发给除了上述几个系统进程外的所有进程；如果是`普通用户`，把信号按照`pid`从大到小的顺序发给`UID`或`SID`等于调用者的`UID`或`有效的UID`的那些进程。

一个进程`并不能`给所有的进程发送信号。`超级用户的进程`可以给任一个进程发送信号。对于`其他用户的发送进程`，它的`真实或有效的用户ID`必需等于`接收进程的真实或有效的用户ID`。

有一个特殊情况，当信号是`SIGCONT`时，一个进程可以向同一会话中所有成员进程发送该信号。

`POSIX.1`定义了信号`0`，如果`signo`是`0`，`kill`会进行一般的错误检查，但不会发送信号。这经常用来判断一个进程是否存在。

如果给一个不存在的进程发送信号，`kill`返回`-1`，全局变量`errno`被设为`ESRCH`。值得注意的是，当进程结束一段时间之后，`Unix系统`会把它的`pid`重新分配给其他进程。因此，这种方法并不能保证想查询的进程存在。

如果用户无权向某些进程发送信号，`kill`返回`-1`，`errno`被置为`EPERM`。

`raise`向进程自身发送信号，相当于`kill(getpid(), signo)`。

对比地，`Go语言`的`syscall`包内提供了如下函数：

```go

func Kill(pid int, sig Signal) (err error)

```


### 用定时器使进程睡眠

```c

#include <unistd.d>

unsigned int alarm(unsigned int seconds);
返回：0或上一次设置的时钟剩下的秒数。

```

`seconds`表示在多少秒之后，一个信号`SIGALRM`会产生。这个信号是由核心产生，由于处理器的调度延迟，在进程捕捉到该信号之前，可能要经过更长的时间。

由于每个进程只有一个时钟，如果在调用一个`alarm`的时候，该进程已经登记了一个时钟还没有结束，`alarm`返回那个时钟剩下的秒数。老的时钟被新的所代替。

如果老的时钟还没有结束，对`alarm`的新调用的参数`seconds`是`0`，老的时钟被取消，它剩下的秒数被返回。

由于对信号`SIGALRM`的缺省处理是结束进程，通常需要捕获该信号，在进程完成一些必需的清理工作后再结束进程。

`pause`使进程悬挂，直到捕获任一个信号。

```c

#include <unistd.h>

int pause(void);
返回：-1，errno被设为EINTR。
当信号处理函数得到执行并返回后，pause才返回-1，设errno为EINTR。

```

示例代码：

```c

/*mysig3.c*/

#include <signal.h>
#include <unistd.h>

void alarm_handler(int signo)
{
    printf("catch SIGALRM\n")
    // DOING SOME WORK
}

int mysleep(int seconds)
{
    if (signal(SIGALRM, alarm_handler) == SIG_ERR)
    {
        return -1;
    }
    alarm(seconds);     // 启动时钟
    pause()             // 等待信号发生
    return (alarm(0))   // 关掉时钟，返回剩下的时间
}

```

这个函数`mysleep`的功能相当于系统提供的函数`sleep`，但它有以下一些问题：

（1）如果调用者已经设置了一个时钟，那么原来的时钟被取消。解决方法如下：

首先调用`alarm`，如果以前设置的时钟剩下的时间小于用户设置的时间，那么只需要等待以前设置的时钟到期；如果以前设置的时钟剩下的时间大于用户设置的时间，在函数返回之前，对`alarm`的调用的参数是这两个时间差。

（2）因为改变了对信号`SIGALRM`的处理，必需保存原来的信号处理函数，在退出前恢复。

（3）如果系统很忙，在第一个`alarm`的调用和`pause`的调用之间，也许会经过很长的时间，时钟可能会结束，`alarm_handler`执行并返回，这时，调动者会在`pause`调用中无限地等待下去。

早期的`sleep`的实现类似用户的程序。解决问题（3）有两种方法，一是使用`setjmp`，另一种是使用`sigprocmask`和`sigsuspend`。

使用`setjmp`解决冲突的方法如下例所示。

```c

/* mysig4.c */
#include <setjmp.h>
#include <signal.h>
#include <unistd.h>

jmp_buf env_alarm;

void alarm_handler(int signo)
{
    longjmp(env_alarm, 1);
}

int mysleep(int seconds)
{
    if (signal(SIGALRM, alarm_handler) == SIG_ERR)
    {
        return -1;
    }
    if (setjmp(env_alarm) == 0)
    {
        alarm(seconds);
        pause();
    }
    return (alarm(0));
}

```

如果在`pause`执行之前，时钟已结束，处理函数`alarm_handler`得到执行，`longjmp`跳到`if (setjmp(env_alarm) == 0)`，由于`longjmp`的值是`1`，不会执行`pause`，从而避免了冲突。

然而，函数`mysleep`仍然有问题。如果`SIGALRM`信号中断了其他的信号的处理函数，对于`longjmp`的调用会使其他的信号处理程序终止。以后的章节会介绍其他的方法来避免这样一些问题。

`alarm`除了用来实现`sleep`之外，还可以给某些操作设一个时间上限。例如在进行文件或网络读写的时候，可能需要在经过太长的时间后终止它们。

```c

/* mysig5.c */
#include <signal.h>

void signal_handler(int signo)
{
    return;
}

int main(void)
{
    if (signal(SIGALRM, signal_handler) == SIG_ERR)
        exit(-1);

    alarm(10);

    if (read(STDIN_FILENO, line, MAXLINE) < 0)
        printf("read error");
    else
        printf("read succedd");
}

```

该程序的问题如下：

（1）存在竞争，如果在`alarm`之后，进程被系统阻塞了很长的时间（大于10秒），在`read`开始执行之后，不会有`alarm`的中断发生，`read`有可能无限的等待下去。

（2）在有些系统中，系统调用（如`read`）在失败之后自动重新发启。因此，当`SIGALRM`的处理函数返回的时候，`read`不会被中断。

可以通过`longjmp`来解决这两个问题。

示例代码：

```c

/* mysig6.c */
#include <setjmp.h>
#include <signal.h>

static jmp_buf env_alarm;

void signal_handler(int signo)
{
    longjmp(env_alarm, 1);
}

int main(void)
{
    char line[MAXLINE];
    if (signal(SIGALRM, signal_handler) == SIG_ERR)
        exit(-1);
    if (setjmp(env_alarm) != 0)
        exit(-1);

    alarm(10);

    if (read(STDIN_FILENO, line, MAXLINE) < 0)
        printf("read error");
    else
        printf("read succeed");

}

```

只要发生了`SIGALRM`事件，程序就会被中断并退出。

如果想对I/O操作设一个时间上限，常可利用`setjmp`和`longjmp`。另一种方法是使用`高级I/O操作`章节中的`select`和`poll`。


### 信号与系统调用

进程在等待外部事件发生的时候，可以接受信号。例如，网络守护进程可能在`accept`上阻塞，等待另一个进程建立网络socket连接。当系统管理员向某个进程发送`SIGTERM`时（不带参数的`kill`命令），有以下几种处理方式。

- 不处理信号，让缺省的处理函数终止进程。
- 捕获信号，让用户的信号处理函数清理进程并退出。然而在程序很复杂的情况下，写出这样的程序难度很大。不推荐这种方法。
- 捕获信号，设一个`标识`表明信号发生，然后用某种方法使被阻塞的系统调用（如`accept`）退出。程序可以检查`标识`，判断是否发生了某个信号，并正确地处理它（包括正常退出）。

第三种方法实现比较容易，所谓地`慢系统调用`就是这样，当系统调用被信号中断时，它们返回，并置`errno`为`EINTR`。而`快系统调用`在信号被发送之前就返回。

`慢系统调用`指的是等待不可预测的事件发生（如其他进程或使用者的活动、外设）的系统调用，它们执行的时间不定。例如，系统调用`wait`等待一个子进程的退出。由于用户不可能知道这个事件何时才能发生，它被称为`慢系统调用`。

`慢系统调用`返回时，进程需要处理判断是否需要重启。这给程序员带来了很多不便，因为每次进行`慢系统调用`（如`read`对一个`慢文件描述符`进行操作）时，都要检查`errno`的值是不是`EINTR`。

为了减轻编程的负担，`BSD4.2`自动重启某些系统调用（如`read`、`write`）。在大多数情况下，用户进程不用操心系统调用被信号中断，因为它们会在信号处理程序运行结束之后继续执行。`POSIX的信号标准`没有规定系统应用如何对待被中断的系统调用，但大多数系统以这种方法处理：在缺省的情况下，系统调用不重启；但进程可以设一个标识，表示它希望被信号中断的系统调用自动重启。


### 信号集

`信号集`用来表示多个信号，通过它可以同时对多个信号进行处理。`POSIX.1`定义了一个数据类型`sigset_t`表示信号集。

```c

#include <signal.h>

int sigemptyset(sigset_t *set);
int sigfillset(sigset_t *set);
int sigaddset(sigset_t *set, int signo);
int sigdelset(sigset_t *set, int signo);
int sigismember(const sigset_t * set, int signo);

```

每一个信号集，在使用它之前必需先调用`sigemptyset`或`sigfillset`初始化。`sigemptyset`使信号集为空，`sigfillset`使信号集包括所有的信号。

当一个信号集被初始化之后，可以加入或删除某个信号。函数`sigaddset`用来加入，函数`sigdelset`用来删除。


### 使用信号集屏蔽信号

一个`进程的屏蔽信号集`是当前被屏蔽不能送往该进程的一组信号的集合。进程可以通过函数`sigprocmask`检查和改变当前的`信号屏蔽集`。函数`sigpending`通过参数`set`返回进程当前被阻塞的，正在悬挂的信号集。

```c

#include <signal.h>
int sigprocmask(int how, const sigset_t* set, sigset_t* oset);
返回值：成功：0；出错：-1。

#include <signal.h>
int sigpending(sigset_t* set);
返回值：成功：0；出错：-1。

```

当`oset`是一个非空的指针时，当前的`信号屏蔽集`在`oset`中返回。

当`set`是一个非空指针时，`how`决定了怎样修改当前的屏蔽集。如果`how=SIG_BLOCK`，`新的屏蔽集`是`原来的屏蔽集`和`set`的`并集`，用来`向屏蔽集中加信号`；如果`how=SIG_UNBLOCK`，`新的屏蔽集`是`原来的屏蔽集`和`set`的`差集`，用来`从屏蔽集中去掉信号`；如果`how=SIG_SETMASK`，信号屏蔽集被设为`set`。

如果`set`为空，当前的信号屏蔽集不变，`how`的值对它无影响。

如果在调用`sigprocmask`之后有正在悬挂的、未阻塞的信号，在`sigprocmask`返回之前其中至少有一个信号被送给进程。

下面是一个说明`sigprocmask`和`sigpending`用法的例子。

```c

/* mysig7.c */
#include <signal.h>

void sig_handler(int signo)
{
    printf("caught SIGUSR1");
}

int main()
{
    sigset_t newmask, oldmask, pendmask;
    if (signal(SIGUSR1, sig_handler) == SIG_ERR)
        exit(-1);
    sigempty(&newmask);
    sigaddset(&newmask, SIGUSR1);
    if (sigprocmask(SIG_BLOCK, &newmask, &oldmask) < 0)
        exit(-1);
    sleep(5);
    if (sigpending(&pendmask) < 0)
        exit(-1);
    if (sigismember(&pendmask, SIGUSR1))
        printf("signal SIGUSR1 pending");
    if (sigprocmask(SIG_SETMASK, &oldmask, NULL) < 0)
        exit(-1);
    printf("SIGUSR1 unblocked");
    sleep(5);
}

```


### 设置信号的处理函数

系统调用`sigaction`用来检查或修改与某一信号对应的动作。它可以取代`老版本的Unix`中的函数`signal`。

```c

#include <signal.h>

int sigaction(int signo, const struct sigaction* act, struct sigaction* oact);
返回：成功：0；出错：-1。

// sigaction结构
struct sigaction {
    void (*sa_handler) ();      // 信号处理函数的地址
    sigset_t sa_mask;           // 其他要屏蔽的信号
    int sa_flags;               // 其他选项
    void (*sa_restorer) (void);
}

```

`signo参数`指的是要修改或检查的信号，如果`act`非空，则把信号动作变为`act`，如果`oact`非空，则在`oact`中返回原来的动作。

当改变信号的动作时，如果`sa_handler`指向一个信号处理函数（不是`SIG_IGN`或`SIG_DFL`），那么，`sa_mask`表示了在`sa_handler`被调用`之前`，一组加到当前的`信号屏蔽集`中的信号；而当信号处理函数返回时，这些信号被去掉，原来的信号屏蔽集被恢复。通过这种方法，当调用`sa_handler`时，能够阻塞某些信号。

当`sa_handler`被调用时，系统自动阻塞被发送的信号。这样就可以保证当我们处理一个信号时，它的另一次发送被阻塞，直到`sa_handler`结束。如果不希望自动屏蔽被发送的信号，可以在`sa_flags`中设置`SA_NOMASK`。

系统中`不存在`信号队列保留多个被阻塞的信号，如果被阻塞的信号发生多次，当解除屏蔽时，信号处理函数只被调用一次。

在设定了一个信号的动作之后，它将一直有效，直到通过`sigaction`改变。

`sa_flags`规定了处理信号的不同方法，它可以包含下列一个或多个标识位（通过`or`，即`比特或`）。`POSIX.1`只定义了`SA_NOCLDSTOP`，其他的几个在Linux中定义，不一定能用于别的系统。

`SA_NOCLDSTOP`：正常情况下，当进程的子进程终止或停止时（通过`wait`可以得到它的返回状态），系统产生`SIGCHLD`并送给父进程。如果`SIGCHLD`的`sa_flags`包含`SA_NOCLDSTOP`，当子进程停止时（由`SIGSTOP`、`SIGTSTP`、`SIGTTIN`或`SIGTTOU`等作业控制信号引起），不产生`SIGCHLD`（仅当子进程终止时产生该信号）。

`SA_NOMASK`：当一个信号的处理函数被调用时，它并不被自动阻塞。使用这个标识会导致不可靠的信号，最好只在模拟不可靠的信号时才使用它。

`SA_ONESHOT`：当信号发送时，信号处理函数被恢复为`SIG_DFL`。该标识用来模拟老系统中的函数`signal`。

`SA_RESTART`：当一个进程正在一个`慢系统调用`中等待时，如果有一个信号发送给它，则系统调用在信号处理函数结束后重启。如果没有使用这个标识，系统调用返回错误，并把`errno`设为`EINTR`。

`sigaction结构`的最后一个域，`sa_restorer`不是`POSIX`的规定，用来保留给以后使用。程序可以忽略它，不用给它赋值。

可以用`sigaction`来实现`signal`。

```c

/* mysig8.c */
#include <signal.h>

void* signal(int signo, void *func)
{
    struct sigaction act, oact;
    act.sa_hander = func;
    sigemptyset(&act.sa_mask);
    act.sa_flags = 0;
    if (signo != SIGALRM)
        act.sa_flags |= SA_RESTART;
    if (sigaction(signo, &act, &oact) < 0)
        return (SIG_ERR);
    return (oact.sa_handler);
}

```

对于除了`SIGALRM`以外的其他信号，必需设置`SA_RESTART`，使被这些信号中断的系统调用重启。不让被`SIGALRM`中断的系统调用重启是为了给I/O操作设一个时间上限。


### 非局部跳转

`setjmp`和`longjmp`可以被用来做非局部的跳转，例如，使用`longjmp`来使信号处理函数直接返回到主进程中的某一点。`ANSI C`规定，信号处理函数可以使用`abort`、`exit`或`longjmp`中任一个退出。

在调用`longjmp`时有一个问题。当捕捉到一个信号时，在运行信号处理函数之前，该信号被自动加到进程的信号屏蔽集中，使另外的信号不会中断当前的信号处理函数。如果使用`longjmp`跳出信号处理函数，对于某些系统，当前的屏蔽集不一定会被改变，即该信号一直被屏蔽。

为了解决这个问题，`POSIX.1`提供了两个新的函数。

```c

#include <setjmp.h>

int sigsetjmp(sigjmp_buf env, int savemask);
返回值：成功：0；错误：-1。
void siglongjmp(sigjmp_buf env, int val);

```

函数`setjmp`和`longjmp`与函数`sigsetjmp`和`siglongjmp`的不同是，后者还有一个新的参数`savemask`。如果`savemask``非0`，`sigsetjmp`把进程的当前信号屏蔽集也存入`env`。当`siglongjmp`被调用时，如果进程的信号屏蔽集存在于`env`中，则它恢复被保存的信号屏蔽集。

如下示例为`sigsetjmp`和`siglongjmp`的用法。

```c

/* mysig9.c */
#include <signal.h>
#include <setjmp.h>
#include <time.h>

static void sig_usr1(int), sig_alarm(int);
static sigjmp_buf jmpbuf;
static volatile sig_atomic_t isjump;

int main()
{
    if (signal(SIGUSR1, sig_usr1) == SIG_ERR)
        exit(-1);
    if (signal(SIGALRM, sig_alarm) == SIG_ERR)
        exit(-1);
    if (sigsetjmp(jmpbuf, 1))
        exit(-1)
    isjump = 1;
    while (1)
        pause();
}

static void sig_usr1(int signo)
{
    time_t starttime;
    if (isjump == 0)
        return;
    alarm(3);
    starttune = time(NULL);
    while (1)
        if (time(NULL) > starttime+5)
            break;
    isjump = 0;
    siglongjmp(jmpbuf, 1);
}

static void sig_alarm(int signo)
{
    printf("in signal SIGALRM\n");
}

```

`isjump`的作用是判断是否应该执行信号处理函数。当执行完`sigsetjmp`之后，把它置为`1`。在信号处理程序中要判断这个变量，当它是`非0`时，表明`sigsetjmp`已经被执行，因此可以执行`siglongjmp`。这样可以避免在`sigsetjmp`还未初始化`jmpbuf`之前就执行信号处理函数。

数据类型`sig_atomic_t`（在Linux中是`atomic_t`，而且需要一些特殊的函数操纵这种类型的变量，参见`asm/atomic.h`），表示改变这种类型的数据的过程不会被中断。同时，使用了类型定义符`volatile`，不许编译程序对它进行寄存器优化，因为`isjump`被两个进程，`main`和`异步的信号处理函数`所访问。


### 屏蔽信号并使进程等待

当程序进入一些不能被信号中断的`临界区域`时，需要屏蔽信号，退出时，需要去掉屏蔽。有时，想去掉对一个信号的屏蔽，然后调用`pause`等待它的发生。如下：

```c

sigset_t newmask, oldmask;
sigempty(&newmask);
sigaddset(&newmask, SIGINT);
if (sigprocmask(SIG_BLOCK, &newmask, &oldmask) < 0)
    exit(-1);
临界区域
if (sigprocmask(SIG_SETMASK, &oldmask, NULL) < 0)
    exit(-1);
pause();

```

当信号`SIGINT`在`sigprocmask`去掉它的屏蔽和`pause`之间发生，它会丢失。为了解决这个问题，需要一种方法可以使`初始化信号屏蔽集`和`使进程睡眠`在一个原子操作中完成，函数`sigsuspend`满足了这个要求。

```c

#include <signal.h>

int sigsuspend(const sigset_t* sigmask);
返回值：-1，errno被设为EINTR。

```

新的信号屏蔽集由`sigmask`提供。然后进程悬挂直到某些信号被捕捉到或某些信号终止了进程。当一个信号被捕捉到，信号处理程序返回之后，`sigsuspend`退出，进程的信号屏蔽集恢复到调用`sigsuspend`之前。

注意：该函数没有成功的返回值，如果返回到调用者，返回值总是`-1`，`errno`被设为`EINTR（表示被信号中断）`。


### 使进程退出

函数`abort`使进程非正常退出。

```c

#include <stdlib.h>
void abort(void);
没有返回值

```

该函数向进程发送信号`SIGABRT`，进程不能忽略这个信号。

`ANSI C`要求，如果一个信号被捕获，信号处理程序返回，`abort`仍然不会返回到它的调用者。在用户的`SIGABRT`的信号处理函数中，它可以通过`exit`、`_exit`退出或通过`longjmp`和`siglongjmp`返回到调用它的主程序。`POSIX.1`规定了，即使进程阻塞了`SIGABRT`或忽略了它，`abort`也会使阻塞和忽略变得无效。

允许进程捕获`SIGABRT`的目的是使进程在退出之前，干一些清理的工作。如果进程在`SIGABRT`的处理函数中没有退出，`POSIX.1`规定，`abort`终止进程。

`ANSI C`没有规定了这个函数是否应该刷新输出流，是否应该删除临时文件。`POSIX.1`要求，如果对`abort`的调用终止了进程，它需要对于所有的打开的I/O流完成`fclose`的功能（显式或隐式）。如果没有终止进程，则不应该对打开的流有任何影响。


### 等待一个进程结束

用`wait`和`waitpid`等待一个进程结束。

```c

#include <sys/types.h>
#include <sys/wait.h>

pid_t wait(int* status);
pid_t waitpid(pid_t pid, int* status, int options);

```

系统调用`wait`使进程挂起，直到它的一个子进程终止或被一个信号中断。如果一个子进程在调用`wait`之前已经退出（所谓的`僵尸进程`），则`wait`立即返回，该子进程所占有的全部资源被释放。

`wait`的返回值：如果成功，返回子进程的`pid`；出错返回`-1`。

系统调用`waitpid`使进程挂起，直到它的一个进程号为`pid`的子进程终止或被一个信号中断。如果一个进程号为`pid`的子进程在调用`waitpid`之前已经退出，则`waitpid`立即返回，该子进程所占有的全部资源被释放。

- 如果`pid < -1`：等待所有`组ID号`等于`pid`的绝对值的子进程。
- 如果`pid = -1`：等待任一个子进程，类似`wait`。
- 如果`pid = 0`：等待所有组ID号等于调用进程的`pid`的那些子进程。
- 如果`pid > 0`：等待进程号为`pid`的子进程。
- `options`为下列一个或多个标识位（通过`or`，即`比特或`）。
- `WNOHANG`：没有子进程退出，立即返回。
- `WUNTRACED`：允许报告已经退出，但还未被报告的子进程。
- `status`，可以用来判断子进程的退出情况。
- `waitpid`的返回值：如果成功，返回子进程的`pid`；出错返回`-1`；如果在`options`中使用了`WNOHANG`，而没有子进程退出，返回`0`。

这两个系统调用被信号`SIGCHLD（由子进程退出导致）`中断返回时，`errno`被置为`EINTR`。


### 实现函数system的一种方法

```c

#include <stdlib.h>

int system(const char* string);  // 返回值：成功则返回加载程序的返回值；如果通过exec()调用shell失败则返回127；其他错误返回-1。

```

在`POSIX.2`中，规定了进程在调用函数`system`时，要忽略`SIGINT`和`SIGQUIT`，阻塞`SIGCHLD`。

因为在`system`中，要通过`fork`加载运行其他的子进程，而`SIGCHLD`是当子进程结束时，发送给父进程（即`system`的调用进程）的信号，所以在`system`还未结束的时候，为了避免中断`system`，进程不应该接收这个信号。

同时，如果`system`所加载的程序需要交互，常用`Ctrl+c`、`Ctrl+\`退出，按照规定，这些信号将送给所有的前台进程。而由于调用`system`的进程正在等待子进程的结束，它不应该因为接收信号`SIGINT`和`SIGQUIT`而退出，所以必需忽略这两个信号。

示例代码：略。


### 实现函数sleep的一种方法

```c

#include <unistd.h>

unsigned int sleep(unsigned int seconds);
返回值：0或剩下还未睡眠的秒数。

```

函数`sleep`使得调用它的进程`悬挂`，直到：

（1）已经过了`seconds`这么长的时间，此时返回值是`0`。

（2）进程捕获了一个信号，并且信号处理函数返回，此时返回值是剩下的秒数。

由于系统的其他活动（如进程上下文切换），`sleep`同`alarm`一样，实际返回所经过的时间有可能比要求的`seconds`长。

示例代码：略。


### 作业控制信号

`POSIX.1`规定，以下的几个信号是作业控制信号。

- `SIGCHLD`：子进程终止或停止。
- `SIGCONT`：如果进程停止，则使它继续。
- `SIGSTOP`：使进程停止。
- `SIGTSTP`：交互的停止信号。
- `SIGTTIN`：后台进程组的进程从控制终端读。
- `SIGTTOU`：后台进程组的进程向控制终端写。

大多数的应用程序不需要处理这些信号，因为shell已经做了几乎所有的工作。当按下`暂停键（Ctrl+z）`挂起当前的前台作业时，`SIGTSTP`送给所有的前台进程组的成员。当用户要shell恢复一个前台或后台的作业时，shell向该作业的所有进程发送信号`SOGCONT`。同样，如果信号`SIGTTIN`或`SIGTTOU`被送给一个进程，则缺省的行为是使进程停止，然后进行作业管理的shell通知用户。

当进程需要管理整个终端时（如`vi`等全屏幕程序所做的那样），它们必需知道自己什么时候被挂起，什么时候恢复运行。因为当它们进行`前后台切换`时，`必需`进行当前屏幕的保存和恢复。

这些信号有以下的相互关系。当对一个进程产生以下四个信号（`SIGTSTP`、`SIGSTOP`、`SIGTTIN`、`SIGTTOU`）中的任一个时，所有对该进程悬挂的`SIGCONT`被抛弃。同样，当对于一个进程产生`SIGCONT`时，这个进程所有悬挂的停止信号被抛弃。

`SIGCONT`的缺省处理是，如果该进程停止，则恢复其运行，否则忽略。正常情况下不用处理它。当它对于一个停止的进程产生时，即使这个信号被进程阻塞，它仍使进程继续执行。


# 目录操作


## 当前目录

在Linux中，每个进程都有一个`当前目录`。系统使用该目录来搜索`相对路径`（所有不以`/`开头的路径都是相对路径，系统认为都是相对于`当前目录`的）。当用户`刚刚`登录进系统时，`当前目录`是由`/etc/passwd`文件中与该用户对应的记录中`第六个域`指定的--用户的`home目录`。`当前目录`是`进程的一个属性`，而`home目录`是`用户的一个属性`，这两者是不同的。


### 获得当前目录

```c

#include <unistd.h>

char *getcwd(char *buf, size_t size);
char *get_current_dir_name(void);
char *getwd(char *buf);

```

`getcwd()函数`把当前目录的`绝对路径`拷贝到`buf`所指的缓冲区中，`buf`的大小为`size`字节。如果当前目录的路径名长度超过`size-1`，函数返回`NULL`，同时`errno`被设为`ERANGE`。应用程序应该检查这种情况以分配更大的空间。当`buf`是`NULL`时，`getcwd()`会调用`malloc()`分配所需的缓冲区。当`size`不为`0`时，所分配的缓冲区为`size`个字节，否则缓冲区为所需要的大小。应用程序应该负责用`free()`释放所得到的缓冲区。

只有在宏`_USE_GNU`被定义以后，`get_current_dir_name()函数`的函数原形才存在。该函数会首先检查`环境变量PWD`，如果`PWD`存在并且其值正确，函数就返回该环境变量的值。实际上该函数的实现类似于：

```c

char* get_current_dir_name(void)
{
    char* env = getenv("PWD");
    if (env) return strdup(env);
    else return getcwd(NULL, 0);
}

```

只有宏`__USE_BSD`被定义以后，`getwd()函数`的函数原形才存在。`getwd()`要求`buf`所指向的缓冲区至少有`PATH_MAX`个字节，它总是返回路径名的前`PATH_MAX`个字节。因此，当路径名大于`PATH_MAX`个字节时，其返回的结果是不正确的。这个函数仅仅是为了支持一些以前的程序而存在的，新的应用程序不应该使用它，应该使用更具有移植性的`getcwd()`。

对比地，`Go语言`的`syscall`包内提供了如下函数：

```go

func Getcwd(buf []byte) (n int, err error)
func Getwd() (wd string, err error)

```

`Go语言`的`os`包内提供了获取`当前目录`的`绝对路径`的函数如下：

```go

func Getwd() (dir string, err error)

```

`Go语言`的示例代码：

```go

package main

import (
    "fmt"
    "os"
    "syscall"
)

func main() {

    buf := make([]byte, 0)
    if _, err := syscall.Getcwd(buf); err == nil {
        fmt.Fprintf(os.Stdout, "Current working directory is %s.\n", buf)
    } else if err == syscall.ERANGE {
        fmt.Fprintln(os.Stderr, "buf too small.")
    } else {
        fmt.Fprintf(os.Stderr, "other error, errno = %d\n", err)
    }

}

```


### 设置当前目录

```c

#include <unistd.h>

int chdir(const char *path)
int fchdir(int fd)

```

`path参数`是希望成为当前目录的目录名，而`fd`是一个文件描述符。这两个函数在成功时返回`0`，可能返回的错误值有：

- `ENOTDIR`：所指定的路径中的一部分不是目录。
- `EACCESS`：对目录没有执行权限。
- `EBADF`：`fd`不是有效的文件描述符。

对比地，`Go语言`的`syscall`包内提供了如下函数：

```go

func Chdir(path string) (err error)
func Fchdir(fd int) (err error)

```

还有，`os`包内提供了如下函数：

```go

// `Chdir`将`当前工作目录`更改为指名的目录。如果有错误，则是`*PathError`类型。
func Chdir(dir string) error

// `Getwd`返回与`当前目录`对应的`根目录`路径名。如果可以通过多个路径访问当前目录（由于`符号链接`），则`Getwd`可能会返回其中任何一个。
func Getwd() (dir string, err error)

// `Chdir`改变`f`的当前工作目录，`f`必须是一个目录。如果有报错，则是`*PathError`类型。
func (f *File) Chdir() error

```


### 改变根目录

虽然`系统`只有一个根目录，但是`对每个进程而言`它的根目录却是可以改变的。这样做的目的是为了防止系统中一些不安全的进程存取整个文件系统。例如：如果`/home/ftp`被指定为`FTP服务器`的根目录，那么一个`FTP用户`执行`chdir /`会将其`当前目录`设为`/home/ftp`，而不是`/`，同时`getcwd()函数`则会返回`/`以保证一致性。为了安全起见，如果进程试图执行`chroot("/..")`，系统仍然会保持根目录为`/home/ftp`而不是`/home`。一旦进程执行了`chroot()`，以后所有与路径有关的操作都会以当前的根目录为基础进行解释。例如：进程先执行`chroot("/home/ftp")`，再执行`chroot("/")`仍然使根目录保持在`/home/ftp`下，因为此时的`/`是以`/home/ftp`为基础进行解释的。

```c

#include <unistd.h>

int chroot(const char *path);

```

需要注意的是，改变根目录`并不会`改变进程的当前目录！进程仍然可以通过相对当前目录的路径存取其他目录的文件。例如：假如当前目录为`/home`在执行`chroot("/home/ftp")`之后，进程仍然可以通过路径`../bin`存取`/bin`的文件。因此程序员应该在执行`chroot()`以后马上执行`chdir("/")`或者类似的调用来改变当前目录，否则程序会留下一个`安全漏洞`。

对比地，`Go语言`的`syscall`包里提供了如下函数：

```go

func Chroot(path string) (err error)

```


## 创建删除目录

```c

#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
#include <uinstd.h>

int mkdir(const char *dirname, mode_t mode)

```

`dirname`指定要创建的目录名，而`mode`则指定该目录的存取权限。最终的存取权限还要受到进程的`umask`的影响，即`mode &~ umask`。如果`dirname`已经存在或者`dirname`的任何一部分不是目录，函数将返回`-1`。因此，要创建多级目录，必需调用多次`mkdir()`。

`新创建目录的uid`将是`进程的有效uid`；`gid`则依赖于下列因素：`父目录的sgid位`，文件系统的安装方式。如果`父目录的sgid位`被设置或这目录所在文件系统是按照`BSD`的组语义安装的，`新目录的gid`将继承`父目录的gid`；否则，`进程的有效gid`将成为`新目录的gid`。另外，如果`父目录的sgid位`被设置，那么`新目录的sgid位`也被设置。

```c

#include <uinstd.h>

int rmdir(const char *pathname);

```

一个目录`只有`是空的时候，才能被删除，否则`rmdir()`会返回`-1`，同时`errno`会等于`ENOTEMPTY`。

对比地，`Go语言`的`syscall`包内提供了如下函数：

```go

func Mkdir(path string, mode uint32) (err error)
func Mkdirat(dirfd int, path string, mode uint32) (err error)

```

还有，`os`包内封装了如下函数：

```go

// `Mkdir`使用指定的名称和权限位（在`umask`之前）创建一个新目录。如果有错误，则是`*PathError`类型。
func Mkdir(name string, perm FileMode) error

// `MkdirAll`创建一个名为`path`的目录以及所有必要的父目录，并返回`nil`，否则返回错误。 权限位`perm`（在`umask`之前）用于`MkdirAll`创建的所有目录。如果`path`已经是目录，则`MkdirAl`l不执行任何操作并返回`nil`。
func MkdirAll(path string, perm FileMode) error

// `Remove`删除指名的文件或目录（空的）。如果有错误，则是`*PathError`类型。
func Remove(name string) error

// `RemoveAll`删除`path`及其包含的子项。该函数能删除它能删除的所有项，但是只返回第一个报错（如果发生的话）。如果`path`不存在，`RemoveAll`返回`nil`（即不会报错）。如果有报错，则是`*PathError`类型。
func RemoveAll(path string) error

```


## 浏览目录

应用程序经常需要知道包含在一个目录中的文件信息。由于Linux可以同时支持多种文件系统，为了使应用程序不必了解文件系统所采用的目录项的格式，Linux提供了一组函数帮助应用程序按照一种抽象的方式处理目录。

```c

#include <sys/types.h>
#include <dirent.h>

DIR* opendir(const char *dirname);
int closedir(DIR *dir);

```

`opendir()`打开一个`目录流`并返回一个指针，应用程序应该像使用`FIFE结构`一样只使用相应的库函数操纵该流。由于目录只能以`只读方式`打开，因此不必指定打开方式。

一旦目录被打开，就可以顺序读取目录项知道目录结束。

```c

#include <sys/types.h>
#include <dirent.h>

struct dirent* readdir(DIR *dir);

```

`readdir()`返回一个`指向静态缓冲区的指针`，以后用同样参数对`readdir()`的调用将覆盖该缓冲区，因此应用程序要自己保存有关的信息。`readdir()`不对返回的目录项的顺序做任何保证，要想对目录排序必需自己进行。

根据`POSIX`，虽然`struct dirent`包含很多域，但是只有一个域是可移植的：`d_name`。该域最多有`NAME_MAX`个字符（不包括结尾的空字符）。其他域是依赖于特定系统的，使用其他任何域都会损害程序的可移植性。

如果`重新读取`一个已经使用`opendir()`打开的目录的内容，可以使用`rewinddir()函数`。它将复位`DIR结构`中的数据，这样下次再调用`readdir()`将返回目录的第一项。

```c

#include <dirent.h>

int rewinddir(DIR *dir);

```

对比地，`Go语言`的`syscall`包里提供了如下函数：

```go

// `ParseDirent`从`buf`解析出`max`条目录条目，并将名称附加到`names`中。它返回从`buf`消耗的字节数，添加到`names`中的条目数以及新名称切片。
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)

func ReadDirent(fd int, buf []byte) (n int, err error)

type Dirent struct {
    Ino uint64
    Off int64
    Reclen uint16
    Type uint8
    Name [256]int8
    Pad_cgo_0 [5]byte
}

```

还有，`Go语言`的`os`包内提供了如下函数：

```go

/*
    `Readdir`读取与文件关联的目录的内容，并按目录顺序返回最多`n`个`FileInfo`值的切片，如`Lstat`将返回的那样。对同一文件的后续调用将产生进一步的`FileInfo`。

    如果`n>0`，则`Readdir`最多返回`n`个`FileInfo`结构。在这种情况下，如果`Readdir`返回一个空切片，它将返回一个非空报错来说明原因。在目录末尾，报错是`io.EOF`。

    如果`n<=0`，则`Readdir`在单个切片中返回目录中的所有`FileInfo`。在这种情况下，如果`Readdir`成功（一直读取到目录的末尾），它将返回切片和nil错误。如果它在目录末尾之前遇到错误，则`Readdir`返回读取的`FileInfo`直到该点为止，并且返回非nil错误。
*/
func (f *File) Readdir(n int) ([]FileInfo, error)

/*
    `Readdirnames`读取与文件关联的目录的内容，并按目录顺序返回目录中最多`n`个文件名的切片。在同一文件上的后续调用将产生其他名称。

    如果`n>0`，则`Readdirnames`最多返回`n`个名称。在这种情况下，如果`Readdirnames`返回一个空切片，它将返回一个非空报错来说明原因。在目录末尾，则报错为`io.EOF`。

    如果`n<=0`，则`Readdirnames`在单个切片中返回目录中的所有名称。在这种情况下，如果`Readdirnames`成功（一直读取到目录的末尾），它将返回切片和nil错误。如果在目录末尾遇到错误，`Readdirnames`将返回读取的名称，直到该点为止，并且返回非nil错误。
*/
func (f *File) Readdirnames(n int) (names []string, err error)

```

下面这个程序统计给定目录中不同类型文件的数量。

```go

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "syscall"
)

var (
    g_nFile int
    g_nDir int
    g_nBlock int
    g_nChar int
    g_nFIFO int
    g_nSymLink int
    g_nSock int
)

func CountFile(stat *syscall.Stat_t) {

    switch stat.Mode&syscall.S_IFMT {
    case syscall.S_IFREG:
        g_nFile ++
    case syscall.S_IFBLK:
        g_nBlock ++
    case syscall.S_IFCHR:
        g_nChar ++
    case syscall.S_IFIFO:
        g_nFIFO ++
    case syscall.S_IFLNK:
        g_nSymLink ++
    case syscall.S_IFSOCK:
        g_nSock ++
    }

}

func WalkDir(path string) {

    var stat syscall.Stat_t

    // 开启目录
    fd, err := syscall.Open(path, syscall.O_DIRECTORY|syscall.O_RDONLY, 0)
    if err != nil {
        panic(err)
    }


    buf := make([]byte, 256)
    names := make([]string, 0)

    for {
        n, err := syscall.ReadDirent(fd, buf)
        if err != nil {
            panic(err)
        }
        if n > 0 {
            _, _, names = syscall.ParseDirent(buf, -1, names)
        } else {
            break
        }
    }
    err = syscall.Close(fd)
    if err != nil {
        panic(err)
    }

    for _, name := range names {

        // 忽略"."和".."文件
        if strings.Compare(name, ".") == 0 ||
         strings.Compare(name, "..") == 0 {
             continue
        }

        // 目录内各文件的路径
        fpath := filepath.Join(path, name)

        err = syscall.Lstat(fpath, &stat)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        if stat.Mode&syscall.S_IFMT == syscall.S_IFDIR {
            // 当前文件是目录
            g_nDir ++
	        //fmt.Println(fpath)
            WalkDir(fpath)
        } else {
            // 当前文件是文件
 	        //fmt.Println(fpath)
            CountFile(&stat)
        }

    }

}

func main() {

    var stat syscall.Stat_t
    path := filepath.Dir(os.Args[1])

    if err := syscall.Lstat(path, &stat); err != nil {
        panic(err)
    }

    if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR {
        // 用户指定的不是目录
        fmt.Fprintf(os.Stderr, "您的输入值为：%s\n%s不是目录！\n", os.Args[1], path)
        return
    }

    WalkDir(path)

    nTotal := g_nFile + g_nDir + g_nBlock + g_nChar + g_nFIFO + g_nSymLink + g_nSock

    if nTotal == 0 {
        fmt.Fprintln(os.Stdout, "Empty direcotry!")
        return
    }

    fmt.Fprintf(os.Stdout, "regular file: %d\n", g_nFile)
    fmt.Fprintf(os.Stdout, "directory: %d\n", g_nDir)
    fmt.Fprintf(os.Stdout, "block device: %d\n", g_nBlock)
    fmt.Fprintf(os.Stdout, "char device: %d\n", g_nChar)
    fmt.Fprintf(os.Stdout, "FIFO: %d\n", g_nFIFO)
    fmt.Fprintf(os.Stdout, "symbolic link: %d\n", g_nSymLink)
    fmt.Fprintf(os.Stdout, "socket: %d\n", g_nSock)

}

```

上述代码是单线程运行，效率不高。`Go语言`的标准库进行了优异的模块化，而且其封装的上层代码也非常实用好用。下面示例代码的功能和上述代码一致。

```go

package main

import (
    "fmt"
    "os"
    "path/filepath"
)

var (
    g_nFile int
    g_nDir int
    g_nBlock int
    g_nChar int
    g_nFIFO int
    g_nSymLink int
    g_nSock int
)

func CountFile(mode os.FileMode) {

    switch {
    case mode.IsRegular():
        g_nFile ++
    case mode&os.ModeCharDevice != 0:
        g_nChar++
    case mode&os.ModeDevice != 0:
        g_nBlock ++
    case mode&os.ModeSymlink != 0:
        g_nSymLink ++
    case mode&os.ModeSocket != 0:
        g_nSock ++
    case mode&os.ModeNamedPipe != 0:
        g_nFIFO ++
    }

}

func WalkDir(dir *os.File) {

    finfos, err := dir.Readdir(-1)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    for _, finfo := range finfos {

        if finfo.IsDir() {
            g_nDir ++
            dpath := filepath.Join(dir.Name(), finfo.Name())
            newdir, err := os.Open(dpath)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                continue
            }
            defer newdir.Close()
            WalkDir(newdir)
        } else {
            CountFile(finfo.Mode())
        }

    }

}

func main() {

    dpath := os.Args[len(os.Args)-1]
    dir, err := os.Open(dpath)
    if err != nil {
        panic(err)
    }
    defer dir.Close()

    dinfo, err := dir.Stat()
    if err != nil {
        panic(err)
    }
    if dinfo.IsDir() {
        g_nDir ++
        WalkDir(dir)
    }

    fmt.Fprintf(os.Stdout, "regular file: %d\n", g_nFile)
    fmt.Fprintf(os.Stdout, "directory: %d\n", g_nDir)
    fmt.Fprintf(os.Stdout, "block device: %d\n", g_nBlock)
    fmt.Fprintf(os.Stdout, "char device: %d\n", g_nChar)
    fmt.Fprintf(os.Stdout, "FIFO: %d\n", g_nFIFO)
    fmt.Fprintf(os.Stdout, "symbolic link: %d\n", g_nSymLink)
    fmt.Fprintf(os.Stdout, "socket: %d\n", g_nSock)

}

```

通过运行比较，`Go语言`的标准库的上层代码非常实用好用，比自己控制底层接口操作状态要更好。

`Go语言`的`path/filepath`包提供了如下两个函数：

```go

/*
    `Walk`遍历以`root`为根的文件树，为树中的每个文件或目录（包括`root`）调用`walkFn`。`walkFn`会过滤访问文件和目录时出现的所有错误。这些文件`以词法顺序进行遍历`，这使输出具有确定性，但是这意味着对于非常大的目录，遍历可能效率不高。`Walk`不会跟踪`符号链接`。
*/
func Walk(root string, walkFn WalkFunc) error

/*
    `WalkFunc`是为`Walk`访问的每个文件或目录调用的函数的类型。`path`参数包含以`Walk`为前缀的参数；也就是说，如果使用`dir`（一个包含文件`a`的目录）调用`Walk`，则将使用参数`dir/a`来调用`WalkFunc`。`info`参数是指名路径的`os.FileInfo`。

    如果在遍历使用`path`指名的文件或目录时遇到问题，则传入的`error`将描述问题，并且该函数可以决定如何处理该错误（`Walk`不再继续进入该目录）。如果发生错误，`info`参数将为`nil`。 如果返回`error`，则处理停止。唯一的例外是当函数返回特殊值`SkipDir`时。如果函数`在目录上`调用时返回`SkipDir`，则`Walk`会完全跳过该目录的内容。如果该函数在`非目录文件`上调用时返回`SkipDir`，则`Walk`将跳过该目录中包含的其余文件。
*/
func WalkFunc func(path string, info os.FileInfo, err error) error

```

上述代码的功能的更实用的实现方式如下：

```go

package main

import (
    "fmt"
    "os"
    "path/filepath"
)

var (
    g_nFile int
    g_nDir int
    g_nBlock int
    g_nChar int
    g_nFIFO int
    g_nSymLink int
    g_nSock int
)

func CountFile(mode os.FileMode) {

    switch {
    case mode&os.ModeType == 0:
        g_nFile ++
    case mode&os.ModeCharDevice != 0:
        g_nChar ++
    case mode&os.ModeDevice != 0:
        g_nBlock ++
    case mode&os.ModeSocket != 0:
        g_nSock ++
    case mode&os.ModeNamedPipe != 0:
        g_nFIFO ++
    case mode&os.ModeSymlink != 0:
        g_nSymLink ++
    }

}

func WalkDir(path string, info os.FileInfo, err error) error {

    if err != nil {
        fmt.Fprintf(os.Stderr, "prevent panic by handling failure accessing a path %q: %v\n", path, err)
        return err
    }

    //fmt.Println(path)
    if info.IsDir() {
        g_nDir ++
    } else {
        CountFile(info.Mode())
    }

    return nil
}

func main() {

    path := os.Args[len(os.Args)-1]
    err := filepath.Walk(path, WalkDir)
    if err != nil {
        panic(err)
    }

    fmt.Fprintf(os.Stdout, "regular file: %d\n", g_nFile)
    fmt.Fprintf(os.Stdout, "directory: %d\n", g_nDir)
    fmt.Fprintf(os.Stdout, "block device: %d\n", g_nBlock)
    fmt.Fprintf(os.Stdout, "char device: %d\n", g_nChar)
    fmt.Fprintf(os.Stdout, "FIFO: %d\n", g_nFIFO)
    fmt.Fprintf(os.Stdout, "symbolic link: %d\n", g_nSymLink)
    fmt.Fprintf(os.Stdout, "socket: %d\n", g_nSock)

}

```

从上述的代码示例中可以看出，`Go`提供的API是很实用且很好用。


## 名字匹配

没有用户会认为执行`ls *.c`会显示名为`*.c`的文件信息，相反用户希望得到当前目录下所有以字母`c`结尾的文件的信息。将`*.c`扩展成`a.c`、`b.c`等的操作是有shell完成的。如果希望在程序中实现这种`文件名匹配`的操作，可以有两种方法。


### 使用子进程

最简单的办法就是`把shell作为子进程`运行，让shell来完成名字匹配。要完成这项工作，需要借助于`popen()函数`--通过`popen()`运行`ls *.c`，然后读入结果。这种方法虽然简单，但是具有很好的可移植性（这就是`Perl`使用这种方法的原因）。


### 内部匹配

如果需要匹配大量的文件名，那么`通过popen()调用子进程`来实现就显得效率很低了。`Linux`中有一个`glob()函数`可以在不需要子进程的情况下进行名字匹配。不过使用该函数会降低程序的可移植性。虽然`POSIX.2`中有关于`glob()函数`的内容，但是不少`Unix系统`并没有实现它。

```c

#include <glob.h>

int glob(const char *pattern, int flags, int errfunc(const char *epath, int errno), glob_t *pglob);

```

第一个参数`pattern`指定文件名必需匹配的模式，比如：`*.c`、`ab[cd]*c`。该函数支持`*`、`?`和`[]`，对它们处理的结果和shell相同。最后一个参数`pglob`指向一个用来保存函数结果的结构。该结构的定义如下：

```c

#include <glob.h>

typedef struct {
    int gl_pathc;       // gl_pathv中的元素数目
    char* gl_pathv[];     // 匹配的文件名
    int gl_offs;        // 在gl_pathv中为GLOB_DOOFS保留的位置
} glob_t;

```

`flags参数`是下面这些常量的`位或`：

- `GLOB_ERR`：遇到错误后返回（例如：由于没有权限无法访问某个目录）。
- `GLOB_MARK`：如果某个目录与模式匹配，在目录后加一个`/`。
- `GLOB_NOSORT`：通常情况下，返回的路径名是按照字母排序的。如果这个标志被设置，就不排序。
- `GLOB_DOOFS`：如果有该标识，调用者同时必需设置`pglob->gl_offs`域。该域指定在返回的`pglob->gl_pathv[]`数组中前多少个元素必需为空。也就是说`glob()`返回后，`pglob->gl_pathv[]`中的元素是这样的：`pglob->gl_offs`个`NULL`，然后是`pglob->gl_pathc`个匹配的路径名。这就可以用`glob()`构成一个可以直接传递给`execv()函数`的参数数组。
- `GLOB_NOCHECK`：如果被`置位`，在没有文件名匹配该模式时，`pattern`自身作为匹配项被返回（一般情况下是什么都不返回的）。而如果`pattern`不包括任何通配符（`*`、`?`、`[]`），无论是否被置位，`pattern`都会被返回。
- `GLOB_APPEND`：函数假设`pglob`包含了上次调用`glob()`得到的结果。这次调用的结果会被追加在以前的结果后面，这样就容易匹配多个模式。
- `GLOB_NOESCAPE`：一般情况下，如果在`*`、`?`、`[`前有一个`\`，该通配符就被认为是一个一般的字符从而失去其通配符的意义。例如：模式`a\*`仅会匹配名字为`a*`的文件。而如果指定了`GLOB_NOESCAPE`，`\`失去了它的特殊意义，`a\*`将匹配任何以`a\*`开头的文件。在这种情况下，`a\`和`a\bcd`都会被匹配，而`arach`不会被匹配。
- `GLOB_PERIOD`：大多数shell都不允许匹配以`.`开头的文件（因为这些文件是系统文件）。`glob()`在一般情况下也是这样。但是如果`GLOB_PERIOD`被指定，以`.`开头的文件也可以参与匹配。

通常情况下，`glob()`会遇到`进程没有权限存取的目录`，这会导致错误。有些情况下，应用程序希望处理这种错误。但是如果`glob()`返回这个错误（通常指定`GLOB_ERR`），程序就没有办法`从出错的位置继续进行匹配`。为了处理这种情况，`glob()`允许将错误报告给一个用户指定的函数。该函数在`glob()`的`第三个参数`指定。其原形如下：

```c

int globerr(const char *pathname, int globerrno);

```

`glob()`在出错时传递给它两个参数：`pathname`指定导致错误的路径名，`globerrno`指定具体的错误值。该错误值一般就是由`opendir()`、`readdir()`或`stat()`在出错时设置的`errno`的值。如果`globerr()`返回`非0值`或者指定了`GLOB_ERR`标识，`glob()`就结束并返回一个错误；否则，匹配操作继续进行。

匹配的结果保存在pglob指向的`glob_t`结构中。其中`gl_pathc`表示与模式匹配的数目，而`gl_pathv`则保存了匹配的文件名（或目录名）。

在应用程序使用完返回的`glob_t`结构后，应该通过`globfree()函数`释放其内存。

```c

#include <glob.h>

void globfree(globt *pglob);

```

下面是`glob()`遇到错误时的一些返回值：

- `GLOB_NOSPACE`表示它无法分配所需空间。
- `GLOB_ABORTED`表示遇到一个`读错误`。
- `GLOB_NOMATCH`表示没有找到匹配项。

对比地，`Go语言`的`filepath`包里提供了如下函数：

```go

/*
    `Glob`返回所有匹配`pattern`的文件名，如果没有匹配文件，则返回nil。`pattern`的语法与`Match`函数中的语法相同。

    `Glob`会忽略`文件系统错误`，例如读取目录的`I/O错误`。`pattern`格式错误时，唯一可能返回的错误是`ErrBadPattern`。
*/
func Glob(pattern string) (matches []string, err error)

/*

    `Match`报告`name`是否与`shell文件名称模式`匹配。模式语法为：

    pattern:
    	{ term }
    term:
    	'*'         matches any sequence of non-Separator characters
    	'?'         matches any single non-Separator character
    	'[' [ '^' ] { character-range } ']'
    	            character class (must be non-empty)
    	c           matches character c (c != '*', '?', '\\', '[')
    	'\\' c      matches character c

    character-range:
    	c           matches character c (c != '\\', '-', ']')
    	'\\' c      matches character c
    	lo '-' hi   matches character c for lo <= c <= hi

    `Match`要求使用`pattern`来匹配`name`整个字符串，而不仅仅是子字符串。`pattern`格式错误时，唯一可能返回的错误是`ErrBadPattern`。

    在Windows上，`转义`已禁用。而是将`\\`视为路径分隔符。

*/
func Match(pattern, name string) (matched bool, err error)

```

下面的例子，输出所有与参数相匹配的文件名（或目录名）。如果遇到错误，它会输出一条错误信息，但是匹配仍然继续进行。

```go

package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {

    pattern := os.Args[1]
    //fmt.Println(pattern)
    fnames, err := filepath.Glob(pattern)
    if err != nil {
        panic(err)
    }
    for _, fname := range fnames {
        fmt.Println(fname)
    }

}

// Test it with:
// go run main.go '*'   // 注意加引号

```

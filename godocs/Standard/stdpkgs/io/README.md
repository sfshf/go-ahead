

# `io`包

`io`包提供了`对I/O原语的基本接口`。本包的基本任务是包装这些`原语`的`已有实现`（如`os`包里的`已有实现`），使之成为`共享的公共接口`，这些公共接口抽象出了泛用的功能，并附加了一些`其它的相关原语`的操作。

由于这些`接口`和`原语`使用各种实现来包装低级操作，除非另行通知，否则客户机`不应假定它们对于并行执行是安全`的。


## 接口

```go

type ByteReader interface {
    ReadByte() (byte, error)
}

type ByteScanner interface {
    ByteReader
    UnreadByte() error
}

type ByteWriter interface {
    WriteByte() error
}

type Closer interface {
    Close() error
}

//=======================================================

type ReadCloser interface {
    Reader
    Closer
}

type ReaderSeeker interface {
    Reader
    Seeker
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

type ReadWriteSeeker interface {
    Reader
    Writer
    Seeker
}

type ReadWriter interface {
    Reader
    Writer
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

//=======================================================

type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}

type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}

//=======================================================

type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}

type RuneScanner interface {
    RuneReader
    UnreadRune() error
}

//=======================================================

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}

//=======================================================

type StringWriter interface {
    WriteString(s string) (n int, err error)
}

//=======================================================

type WriteCloser interface {
    Writer
    Closer
}

type WriteSeeker interface {
    Writer
    Seeker
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}

type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}

```


## 结构体

```go

type LimitedReader struct {
    R Reader    // underlying reader
    N int64     // max bytes remaining
}

//=======================================================

type PipeReader struct {
    p *pipe
}

type PipeWriter struct {
    p *pipe
}

//=======================================================

type SectionReader struct {
    r ReaderAt
    base int64
    off int64
    limit int64
}

```


## 函数

```go

func Copy(dst Writer, src Reader) (written int64, err error)
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
func Pipe() (*PipeReader, *PipeWriter)
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
func WriteString(w Writer, s string) (n int, err error)

func LimitReader(r Reader, n int64) Reader
func MultiReader(readers ...Reader) Reader
func TeeReader(r Reader, w Writer) Reader

func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

func MultiWriter(writers ...Writer) Writer

```


## 常量

```go

const (
    SeekStart = 0   // seek relative to the origin of the file
    SeekCurrent = 1 // seek relative to the current offset
    SeekEnd = 2     // seek relative to the end
)

```


##　变量

```go

var EOF = errors.New("EOF")
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
var ErrShortBuffer = errors.New("short buffer")
var ErrShortWrite = errors.New("short write")
var ErrUnexpectedEOF = errors.New("unexpected EOF")

```



# [Go database/sql tutorial](http://go-database-sql.org/index.html)

在Go语言中使用SQL或类SQL的数据库的符合go语言习惯的方式是通过`database/sql`包，该包提供了轻量级的接口来操作面向行的数据库。该文档提供了如何使用该包最常用内容的参考。

为什么需要本文档？`database/sql`包的文档提供的是如何操作数据库，而本文档并不告诉您如何使用`database/sql`包。因为我们很多人想要一个能快速参考的、面向入门的、类似讲故事的文档，而不是列述如何操作的文档。


## Overview

想要使用Go语言访问数据库，您要使用`sql.DB`。您可以使用该类型来创建`sql语句（statement）`、`事务（transaction）`、`执行查询（execute queries）`和`获取结果（fetch results）`。

首先，您需要明白的是，一个`sql.DB`对象并不是指一个`数据库连接`；而且也不映射任何特定的`数据库软件`的`"database"`或`"schema"`概念；它是一个`数据库`的接口和存在的一种抽象概念，该`数据库`会因`实际访问`的不同而不同，`实际访问`的可以是`一个本地文件`、`一个网络连接`或`在内存中的`和`在进程中的`。

`sql.DB`会在后台为您执行一些重要的任务：

- 它会通过`驱动`来打开和关闭与底层实际的数据库的`连接`。
- 它会`根据需求`来管理一个`连接池`，`需求`可以是之前提到的各种操作。

`sql.DB`的抽象设计是为了让您不同担心如何管理对底层数据库的并发访问。当您使用某个`连接`执行任务时，该连接会被标记为`使用中（in-use）`，当不再使用时，该连接会返回到`可用池`中。这样的后果之一是，如果无法将连接释放回池，则可能导致`sql.DB`打开很多连接，从而可能耗尽资源（太多的连接，太多的打开的文件句柄，缺少可用的网络端口等）。稍后我们将详细讨论。

创建`sql.DB`之后，可以使用它查询它代表的数据库以及创建语句和事务。


## Importing a Database Driver

要使用`database/sql`，您需要包本身，以及要使用的`特定数据库的驱动程序`。

通常，您`不应该`直接使用`驱动程序包`，尽管有些驱动程序鼓励您这样做（我们认为，通常这是个坏主意。）。相反，如果可能，您的代码应该仅仅引用`database/sql`中定义的类型。这有助于避免`您的代码`依赖于`驱动程序`，从而使您可以通过最少的代码更改来更改底层的驱动程序（从而更改正在访问的数据库）。它还会迫使您使用`Go语言惯用语`，而不是特定驱动程序作者可能提供的`专门工具的惯用语`。

在本文档中，我们将使用`@julienschmidt`和`@arnehormann`提供的出色的[MySQL驱动程序](https://github.com/go-sql-driver/mysql)作为示例。

将以下内容添加到`Go源文件`的顶部：

```go

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

```

请注意，我们正在`匿名地`加载驱动程序，将其程序包限定符别名命名为`_`，因此其`导出的命名`对我们的代码均不可见。 在幕后，`驱动程序`将自身`注册`为可用于`database/sql`包，但一般而言，除了运行`init函数`外，不会发生任何其他事情。

现在您可以访问数据库了。


## Accessing the Database

现在，您已经加载了`驱动包`，您已经可以创建`数据库对象`--`sql.DB`。

想要创建`sql.DB`对象，您要使用`sql.Open()`函数，该函数返回一个`*sql.DB`：

```go

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}

```

在示例中，我们要说明几件事：

1. `sql.Open`的`第一个参数`是`驱动程序名称`。这是`驱动程序`用于向`database/sql`注册自身的字符串，并且通常与程序包名称相同，以避免混淆。例如，它是`github.com/go-sql-driver/mysql`的`mysql`。某些驱动程序没有遵循该约定，而是使用数据库名称，例如用于`github.com/mattn/go-sqlite3`的`sqlite3`和用于`github.com/lib/pq`的`postgres`。

2. `第二个参数`是`特定于驱动程序的语法`，该语法告诉`驱动程序`如何访问`底层数据库`。在此示例中，我们将连接到`本地MySQL服务器实例`内的`"hello"数据库`。

3. 您应该（几乎）是总要检查并处理从所有`database/sql`操作返回的`报错`的。我们之后将讨论到一些没必要处理报错的`特殊情况`。

4. 如果`sql.DB的生命周期`不应超出某函数范围的话，习惯方式是`defer db.Close()`。

可能不怎么直观的是，`sql.Open()`没有建立与数据库的任何连接，也不会验证驱动程序的连接参数。取而代之的是，它仅仅准备了`数据库的抽象`以供之后使用。与`底层数据库`的`首次实际连接`将在第一次需要时`延迟建立`。如果要立即检查数据库是否可用和可访问（例如，检查是否可以建立网络连接并登录），请使用`db.Ping()`进行操作，并记住检查报错：

```go

err = db.Ping()
if err != nil {
    // do something here
}

```

尽管在完成数据库操作后将其进行`Close()`是惯用操作的，但是`sql.DB对象`被设计为`长寿命的`。不要频繁地`Open()`和`Close()`数据库；而要为您需要访问的每个不同的`数据库`创建一个`sql.DB对象`，并保留该对象，直到程序完成对该数据存储的访问为止。根据需要`传递`它，或以某种方式使其在`全局范围`内可用，但`保持打开状态`。而且不要在`短生命周期的函数`里进行`Open()`和`Close()`操作；而是将`sql.DB`作为`参数`传递给该`短生命周期的函数`。

如果不将`sql.DB`作为`长生命周期的对象`，则可能会遇到诸如`重用`和`连接共享`不良，`可用网络资源用尽`或由于大量TCP连接处于`TIME_WAIT`状态导致的偶发性连接失败等问题。这些问题表明您没有使用好`database/sql`的设计。

现在该使用`sql.DB对象`了。


## Retrieving Result Sets

从`数据库`中检索结果有`几种惯用的操作`。

1. 执行`返回多行`的一个`查询（query）`。
2. 准备一条`重复使用`的`语句（statement）`，`多次执行`，然后销毁它。
3. `一次性`执行一条`语句（statement）`，无需准备重复使用。
4. 执行`返回单行`的`查询（query）`。这种特殊情况有一个捷径。

Go语言的`database/sql`包内的`函数名称`是很重要的。如果函数名称包含`Query`，则该函数旨在询问一个`数据库问题`，并且即使它为空，也将返回`一组行`。`不返回多行`的`语句`不应使用`Query`函数；他们应该使用`Exec()`。


### Fetching Data from the Database

让我们看一个如何`查询（query）`数据库并处理结果的示例。我们将在`users`表中查询`id`为`1`的用户，并打印出该用户的`id`和`name`。我们将通过`row.Scan()`一次一行地将结果分配给变量。

```go

var (
    id int
    name string
)

rows, err := db.Query("SELECT id, name FROM users WHERE id=?", 1)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    err := rows.Scan(&id, &name)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(id, name)
}

err = rows.Err()
if err != nil {
    log.Fatal(err)
}

```

上面的代码中发生的事情：

1. 我们正在使用`db.Query()`将`查询（query）`发送到数据库。我们照常`检查错误`。
2. 使用`defer rows.Close()`。这个非常重要。
3. 我们使用`rows.Next()`遍历所有返回行。
4. 我们使用`rows.Scan()`将每一行中的`列`读入变量中。
5. 完成对所有行的迭代后，我们将`检查错误`。

这几乎是在`Go语言`中`唯一的方式`；例如，您无法获得`map格式的`一行结果；那是因为所有内容都是`强类型的`；正如示例所展示地，您需要创建`正确类型的变量`，并将`指针`传递给它们。

其中有几个部分很容易出错，并可能带来严重的后果。

- 您应该始终在`for rows.Next()`循环结束时`检查错误`。如果循环过程中出现错误，则需要知道该报错；不要认为循环部分只需要迭代您要处理所有行。
- 其次，只要有一个开放的`结果集`（用`rows`表示），那么`底层连接`就很忙，不能用于其他任何`查询`。这意味着它在`连接池`中为`不可用的`。如果您使用`rows.Next()`遍历`所有行`，在最终您将读取`最后一行`时，`rows.Next()`将遇到`内部的EOF错误`并为您调用`rows.Close()`。但是，如果您出于某种原因退出了循环（`提前返回`等），则`rows`不会关闭，连接保持打开状态。（但是，如果`rows.Next()`由于一个错误而返回`false`，它将`自动关闭`）。这种容易发生的情况将会耗尽资源，所以这一点必须要注意。
- `rows.Close()`在已经关闭的情况下是无害的空操作，因此您可以多次调用它。但是请注意，我们`首先检查错误`，如果没有错误，则仅调用`rows.Close()`，以避免运行时出现`宕机（panic）`。
- 即使您在循环结束时会显式调用了`rows.Close()`，也建议您应该总是写上`defer rows.Close()`；这不是一个坏主意。
- 不要在循环中`defer（延迟）`。在函数退出之前，不会执行`延迟语句`，因此`长时间运行的函数`不应使用该语句；如果您这样做，该函数将会慢慢积攒内存。如果您要`在循环中重复查询和使用结果集`，则在处理完每个结果后应`显式`调用`row.Close()`，而`不要`使用`defer`。


### How `Scan()` Works

当您遍历所有行并将其扫描到`目标变量`中时，`Go`会在后台执行`数据类型转换`。它是基于`目标变量的类型`。意识到这一点，可以清理您的代码，并有助于避免重复的工作。

例如，假设您从用`字符串`列定义的表中选择一些行，例如`VARCHAR(45)`或类似的字符串。但是您发现该表`始终`包含的是`数字`。如果将`指针`传递给`字符串`，`Go`会将`字节`复制到`字符串`中。现在，您可以使用`strconv.ParseInt()`或类似方法将值转换为`数字`。您必须检查`SQL操作`中的`错误`以及`解析整数的错误`。这是混乱而又乏味的。

或者，您可以仅将`Scan()`的`指针`传递给`整数`。`Go`将检测到`该错误`并为您调用`strconv.ParseInt()`。如果转换中出现错误，则`Scan()`的调用会将其返回。您的代码现在更整洁，更小了。这是使用`database/sql`的推荐方法。


### Preparing Queries

通常，您应该总是要`准备`一些`多次使用的查询（query）`。`某查询准备的结果`是一个`预备语句（prepared statement）`，该`语句`可以具有`占位符`（也称为`绑定值`），用于执行该语句时您提供参数。由于所有常见的原因（例如`避免SQL注入攻击`），这比`串联字符串`好得多。

在`MySQL`中，`参数占位符`为`?`，而在`PostgreSQL`中为`$N`，其中`N`为数字。`SQLite`接受其中任何一种。在`Oracle`中，`占位符`以`冒号`开头并`被命名`，例如`:param1`。因为我们以`MySQL`为例，我们将使用`?`。

```go

stmt, err := db.Prepare("SELECT id, name FROM users WHERE id=?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

rows, err := stmt.Query(1)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    // ...
}
if err = rows.Err(); err != nil {
    log.Fatal(err)
}

```

在内部，`db.Query()`实际上`准备`，`执行`和`关闭`一个`预备语句`；该过程对数据库进行了`三个往返`的访问。如果您不小心，会将应用程序进行的`数据库交互`的次数增加`三倍`！在某些情况下，`某些驱动程序`可以避免这种情况，但`并非`所有驱动程序都可以。有关更多信息，请参见[Using Prepared Statements](#using-prepared-statements)。


### Single-Row Queries

如果一个`查询（query）`返回最多一行，那么您可以在一些`冗长的样板代码`周围使用快捷方式。

```go

var name string
err = db.QueryRow("SELECT name FROM users WHERE id=?", 1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println(name)

```

`查询（query）`产生的错误会被延迟到`Scan()`被调用，然后从`Scan()`处返回。您还可以在一个`预备语句`上调用`QueryRow()`：

```go

stmt, err := db.Prepare("SELECT name FROM users WHERE id=?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println(name)

```


## Modifying Data and Using Transactions

现在，我们准备好看看如何`修改数据`和`处理事务`；如果您习惯于使用利用`"statement"对象`来获取行以及更新数据的编程语言，这样的区分似乎是人为的；但是在`Go语言`中，有一个`重要的原因`。


### Statements that Modify Data

最好将`Exec()`与`预备语句`一起使用，来完成`INSERT`，`UPDATE`，`DELETE`或`其他不返回行的语句`。以下示例显示如何`插入行`并`检查有关该操作的元数据`：

```go

stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
    log.Fatal(err)
}

res, err := stmt.Exec("Dolly")
if err != nil {
    log.Fatal(err)
}

lastId, err := res.LastInsertId()
if err != nil {
    log.Fatal(err)
}

rowCnt, err := res.RowsAffected()
if err != nil {
    log.Fatal(err)
}
log.Printf("ID=%d, affected=%d\n", lastId, rowCnt)

```

执行该语句将产生一个`sql.Result`，它可以访问`语句元数据`：`最后插入的ID`和`受影响的行数`。

如果您不关心结果，怎么办？如果您只想执行一条语句并检查是否有任何错误而忽略结果，怎么办？如下的两个语句不会做同一件事吗？

```go

_, err := db.Exec("DELETE FROM users")      // OK
_, err := db.Query("DELETE FROM users")     // BAD

```

答案是不。它们没有做同一件事情，而且您永远不应该以这样的方式使用`Query()`。`Query()`将返回一个`sql.Rows`，`sql.Rows将保留数据库连接`，直到关闭。由于可能存在未读数据（例如，更多的数据行），因此无法使用该连接。在上面的示例中，连接将不再被释放；垃圾收集器最终将为您关闭底层的`net.Conn`，但这可能需要很长时间。此外，`database/sql`包会在它的`连接池`中继续跟踪连接，指望您在某个点上释放它，以便可以再次使用该连接。因此，该`反面模式`是耗尽资源（例如，`连接过多`）的好方法。


### Working with Transactions

在`Go`中，`事务`本质上是一个保留`数据库连接`的对象。它可以让您执行到目前为止所看到的所有操作，但可以保证它们将`在同一连接上`执行。

您可以通过调用`db.Begin()`来开始事务，然后对作为结果的`Tx`类型的变量使用`Commit()`或`Rollback()`方法将其关闭。在内部，`Tx`从`连接池`中获取`连接`，`并保留该连接仅用于该事务`。`Tx`上的方法`一对一地`对应您可以在`数据库本身`上调用的方法，例如`Query()`等。

`在事务中创建的预备语句`专门`绑定`到该事务。有关更多信息，请参见[Using Prepared Statements](#using-prepared-statements)。

您`不应该`将`与事务相关的函数`（例如`Begin()`和`Commit()`）与`SQL语句`（例如您`SQL代码`中的`BEGIN`和`COMMIT`）混合使用。坏事可能导致：

- `Tx对象`可以保持打开状态，保留来自`连接池`的连接，而不返回它。
- `数据库的状态`可能与`代表它的Go变量的状态`不同步。
- 您可能会认为您是`在事务内部的单个连接`上执行查询，而实际上`Go`会为您无形地创建多个连接，并且`某些语句`不属于`事务`。

`在事务内部`进行操作时，应注意`不要`调用`db`变量；应该是对使用`db.Begin()`创建的`Tx变量`进行所有操作调用；`db`不是`事务`，只有`Tx对象`是`事务`。如果您进一步调用`db.Exec()`或类似方法，则这些调用将在`事务范围之外的其他连接`上发生。

如果您需要使用`多个会修改连接状态的语句`，则即使您本身不需要`事务`，但仍需要一个`Tx`。例如：

- 创建`临时表`，这些表仅对`一个连接`可见。
- 设置`变量`，例如`MySQL`的`SET @var := somevalue`语法。
- 更改`连接选项`，例如`字符集`或`超时`。

如果您需要执行这些操作中的任何操作，则需要将您的活动`绑定到单个连接`，而`在Go中`执行此操作的唯一方式是使用`Tx`。


## Using Prepared Statements

`预备语句`具有`Go`的所有常规好处：安全性，效率，便利性。但是它们的实现方式与您习惯的方式有些不同，尤其是在它们如何与`database/sql`内部组件交互方面。


### Prepared Statements And Connections

`在数据库层面上`，`预备语句`是绑定到`单个数据库连接`。典型流程是，客户端将`带有占位符的SQL语句`发送到服务器进行准备，服务器以`语句ID`进行响应，然后客户端通过发送`其ID和参数`来`执行`该语句。

但是，`在Go中`，`连接`不会直接向`database/sql`包的用户公开。`您无需在一个连接上准备一条语句`；您可以在`DB`或`Tx`上进行准备。而且`database/sql`有一些很方便的行为，例如`自动重试`。由于这些原因，`预备语句`和`连接`之间的底层关联（存在于`在驱动程序层面上`），对您的代码是隐藏的。

运作方式如下：

1. 准备语句时，它是`在连接池中的连接上`准备的。
2. `Stmt对象`会记住使用了哪个`连接`。
3. 当您执行`Stmt`时，它将尝试使用`该连接`。如果由于关闭或忙于执行其他操作而无法使用，它将从`连接池`中获取`另一个连接`，并在另一个连接上`重新准备该数据库的语句`。

由于在`原始连接`繁忙时会根据需要`重新准备语句`，因此`数据库的高并发使用`（可能会保持大量连接繁忙）会创建大量`预备语句`。这可能导致明显的`语句泄漏`，会比您能想象到的更频繁地准备和重新准备语句，甚至语句数量会达到服务器端的限制。


### Avoiding Prepared Statements

`Go`在内部会为您创建`预备语句`。例如，一个简单的`db.Query(sql, param1, param2)`，其工作内容是：`准备sql`，然后`使用参数执行该sql`，最后`关闭该语句`。

有时，准备好的语句不是您想要的。可能有如下原因：

1. 该数据库不支持`预备语句`。例如，使用`MySQL驱动程序`时，可以连接到`MemSQL`和`Sphinx`，因为它们支持`MySQL有线协议`。但是它们不支持包含`预备语句`的`二进制协议`，因此它们可能会造成混乱。
2. 语句的重用起不到意义，并且`安全问题`是以其他方式处理的，因此不希望有性能开销。可以在[VividCortex博客](https://www.solarwinds.com/database-performance-monitor)上看到一个示例。

如果您不想使用`预备语句`，则需要使用`fmt.Sprint()`或类似语句来组装SQL，并将其作为唯一参数传递给`db.Query()`或`db.QueryRow()`。并且您的`驱动程序`需要支持`纯文本查询（plaintext query）`的执行，这一点被通过`Execer`和`Queryer`接口添加在`Go 1.1`中，可查看[database/sql/driver中的文档](golang.org/pkg/database/sql/driver/#Execer)。


### Prepared Statements in Transactions

在`Tx`中创建的`预备语句`被专门绑定到该`Tx`中，因此，之前`关于重新准备的警告`不再适用。当您在`Tx对象`上进行操作时，您的操作将直接对应到`该对象下的唯一连接`。

这也意味着`在Tx中创建的预备语句`不能单独使用。同样，`在DB上创建的预备语句`不能在`事务`中使用，因为它们将绑定到其他连接。

要在`Tx`中使用`事务外部准备的预备语句`，可以使用`Tx.Stmt()`，它将一个`在事务外准备的语句`创建成一个`新的指定事务的语句`。为此，它采用`现有的预备语句`，将其连接设置为`事务的连接`，并在每次执行时重新准备所有语句。这种行为及其实现是不可取的，甚至在`database/sql`源代码中用`TODO`表示以后要改进它；我们建议不要使用它。

在事务中使用`预备语句`时必须谨慎。考虑以下示例：

```go

tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback()

stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()      // danger!

for i := 0; i < 10; i ++ {
    _, err := stmt.Exec(i)
    if err != nil {
        log.Fatal(err)
    }
}
err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
// stmt.Close()     runs here!

```

在`Go 1.4`之前，关闭`*sql.Tx`会将与其关联的连接释放回池中，但是在此之后，`对预备语句的Close的延迟调用`会被执行，这可能导致对底层连接的并发访问，从而导致连接状态不一致。如果使用`Go 1.4或更老的版本`，则应确保`在提交或回滚事务之前`将该语句关闭。[此问题](https://github.com/golang/go/issues/4459)已由[CR 131650043](https://codereview.appspot.com/131650043)在`Go 1.4`中修复。


### Parameter Placeholder Syntax

`预备语句`中的`占位符参数`的语法是`特定于数据库的`。例如，`MySQL`，`PostgreSQL`和`Oracle`的比较如下：

| MySQL | PostgreSQL | Oracle |
|--|--|--|
| WHERE col = ? | WHERE col = $1 | WHERE col = :col |
| VALUES(?, ?, ?) | VALUES($1, $2, $3) | VALUES(:val1, :val2, :val3) |


## Handling Errors

几乎所有与`database/sql`类型相关的操作都会返回一个`error`作为最后一个返回值。您应该一直要检查这些报错，绝不要忽略它们。

在有些地方，报错的行为有特殊情况，也就是说有另外一些东西您需要知道。


### Errors From Iterating Resultsets

考虑如下代码：

```go

for rows.Next() {
    // ...
}
if err = rows.Err(); err != nil {
    // handle the error here
}

```

来自`rows.Err()`的错误可能是`rows.Next()`循环中的不同的错误。该循环可能因为除了正常结束循环之外的一些原因而退出，所以您总需要取检查循环是否正常结束。异常结束会自动调用`rows.Close()`，尽管多次调用`rows.Close()`是无害的。


### Errors From Closing Resultsets

如前所述，如果过早退出循环，则应始终明确地关闭`sql.Rows`。如果循环正常退出或由于错误退出，它将自动关闭，但是您可能会错误地如下这样做：

```go

for rows.Next() {
    // ...
    break   // whoops, rows is not closed! memory leak ...
}

// do the usual "if err = rows.Err()" [omitted here]...
// it's always safe to [re?]close here:
if err = rows.Close(); err != nil {
    // but what should we do if there's an error?
    log.Println(err)
}

```

由`rows.Close()`返回的错误是一般规则的唯一例外，该常规规则最好是捕获并检查所有数据库操作中的错误。如果`rows.Close()`返回错误，则不清楚应该怎么做。`记录错误消息`或`引起宕机`可能是唯一明智的选择，如果这不够明智，那么也许您应该忽略该错误。


### Errors From `QueryRow()`

考虑如下`获取单行`的代码：

```go

var name string
err = db.QueryRow("SELECT name FROM users WHERE id=?", 1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println(name)

```

如果没有`id=1`的用户怎么办？然后返回的结果中将没有行，并且`.Scan()`不会将值扫描到`name`中。那会发生什么呢？

`Go`定义了一个`特殊的错误常量`，称为`sql.ErrNoRows`，当结果为空时从`QueryRow()`返回。在大多数情况下，这需要作为特殊情况进行处理。`应用程序代码`通常不会将`空结果`视为`错误`，并且，如果不检查错误是否是`此特殊常量`，则会导致您未曾想到的`应用程序代码错误`。

`查询中的错误`将推迟到调用`Scan()`之后返回出来。上述代码最好用如下代码替代：

```go

var name string
err = db.QueryRow("SELECT name FROM users WHERE id=?", 1).Scan(&name)
if err != nil {
    if err == sql.ErrNoRows {
        // there were no rows, but otherwise no error occurred
    } else {
        log.Fatal(err)
    }
}
fmt.Println(name)

```

有人可能会问，为什么将`空结果集`视为`错误`。`空集`是不含有任何错误的。原因是，`QueryRow()`方法需要使用`这种特殊情况`，以便让调用者区分`QueryRow()`实际上是否找到了一行。没有返回结果的话，`Scan()`将无法执行任何操作，您可能不会意识到您的`变量`其实没有从数据库中获得任何值。

您应该仅会在使用`QueryRow()`时遇到此错误。如果您在其他地方遇到此错误，则说明您在做错事。


### Identifying Specific Database Errors

很有可能写出如下代码：

```go

rows, err := db.Query("SELECT someval FROM sometable")
// err contains:
// ERROR 1045 (28000): Access denied for user 'foo'@'::1' (using password: NO)
if strings.Contains(err.Error(), "Access denied") {
    // Handle the permission-denied error
}

```

不过，这并不是最好的方法。例如，`报错的字符串值`可能会有所不同，具体取决于服务器使用哪种语言发送错误消息。比较`错误编号`以确定该特定的错误是什么要好得多。

但是，`驱动程序的执行机制`因`驱动程序`而异，因为它不是`database/sql`本身的一部分。在本教程重点关注的`MySQL驱动程序`中，您可以编写以下代码：

```go

if driverErr, ok := err.(*mysql.MySQLError); ok {   // Now the error number is accessible directly
    if driverErr.Number == 1045 {
        // Handle the permission-denied error
    }
}

```

同样，此处的`MySQLError`类型由此`特定的驱动程序`提供，`.Number字段`在驱动程序之间可能有所不同。但是，`数字的值`是从`MySQL的错误消息`中获取的，因此是`数据库特定的`，而不是`驱动程序特定的`。

这段代码仍然很难看。与`1045`相比，一个`魔术数字`就有代码味道了。一些`驱动程序`（尽管不是MySQL的驱动程序，由于此处不在主题之列）提供了`错误标识符列表`。例如，`Postgres`的`pq`驱动程序在[`error.go`](https://github.com/lib/pq/blob/master/error.go)中就有。还有[VividCortex维护的MySQL错误号的外部软件包](https://github.com/VividCortex/mysqlerr)。使用`这样的错误标识列表`，上述代码可以如下更好地编写：

```go

if driverErr, ok := err.(*mysql.MySQLError); ok {
    if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
        // Handle the permission-denied error
    }
}

```


### Handling Connection Errors

如果您`与数据库的连接`发生`丢失`、`被杀死`或`发生错误`，怎么办？

发生这些情况时，您无需实施任何逻辑来重试失败的语句。作为`database/sql`中的[连接池](#the-connection-pool)的一部分，`内置了处理失败的连接的功能`。如果执行查询或其他语句，而底层的连接失败，则`Go`将重新打开一个新连接（或仅从连接池中获得另一个连接）并重试，`最多10次`。

但是，可能会有一些意想不到的后果。当其他错误情况发生时，某些类型的错误可能会被重试。`这里也可能是特定于驱动程序的`。`MySQL驱动程序`中的一个示例是，使用`KILL`取消不需要的语句（例如长时间运行的查询），会导致该语句最多重试10次。


## Working with NULLs

`可空列`是很烦人的，还会导致很多难看的代码。如果可以，请避免使用它们。如果不行，那么您将需要使用`database/sql`包中的`特殊类型`来处理它们，或者定义自己的类型。

有`可空布尔型`，`可空字符串型`，`可空整型`和`可空浮点型`。使用方法如下：

```go

for rows.Next() {
    var s sql.NullString
    err := rows.Scan(&s)
    // check err
    if s.Valid {
        // use s.String
    } else {
        // NULL value
    }
}

```

`可空类型`具有局限性，以及您需要避免使用`可空列`的更有说服力的原因有：

1. 没有`sql.NullUint64`类型或`sql.NullYourFavoriteType`类型；您需要自定义。
2. `可空性`可能很棘手，并且不能适应未来。如果您认为某些内容不会为空，但是您错了，则您的程序将崩溃，崩溃可能很少发生，以至于您无法在解决错误之前发现错误。
3. `Go`的优点之一是每个`变量`都有一个有用的`默认零值`。这不是`可空内容`的工作方式。

如果需要自定义类型来处理`NULL`，则可以仿照`sql.NullString`的设计来实现。

如果您无法避免在数据库中使用`NULL值`，则另一种大多数数据库系统都支持的解决方法是，即`COALESCE()`。您可以使用如下类似的方式，而无需引入无数的`sql.Null*`类型。

```go

rows, err := db.Query(`
    SELECT
            name,
            COALESCE(other_field, '') as otherField
    WHERE id=?
`, 42)
for rows.Next() {
    err := rows.Scan(&name, &otherField)
    // ..
    // If `other_field` was NULL, `otherField` is now an empty string. This works with other data types as well.
}

```


## Working with Unknown Columns

`Scan()`函数要求您传递`正确数量`的`目标变量`。如果您不知道`查询`将返回什么，怎么办？

如果您不知道`查询`将返回多少`列`，则可以使用`Columns()`查找`列名列表`。您可以检查`此列表的长度`以查看有多少列，并且可以将`具有正确数量的值的接口切片`传递到`Scan()`中。 例如，`某些MySQL的分岔`会为`SHOW PROCESSLIST`命令返回不同的列，因此您必须为此做好准备，否则会导致错误。如下是一种实现方法；还有其他方式：

```go

cols, err := rows.Columns()
if err != nil {
    // handle the error
} else {
    dest := []interface{}{      // Standard MySQL columns
        new(uint64),    // id
        new(string),    // host
        new(string),    // user
        new(string),    // db
        new(string),    // command
        new(uint32),    // time
        new(string),    // state
        new(string),    // info
    }
    if len(cols) == 11 {
        // Percona Server
    } else if len(cols) > 8 {
        // Handler this case
    }
    err = rows.Scan(dest...)
    // Work with the values in dest
}

```

如果您不清楚有多少列或它们的类型，您应该使用`sql.RawBytes`。

```go

cols, err := rows.Columns()     // Remember to check err afterwards
vals := make([]interface{}, len(cols))
for i, _ := range cols {
    vals[i] = new(sql.RawBytes)
}
for rows.Next() {
    err = rows.Scan(vals...)
    // Now you can check each element of vals for nil-ness,
    // and you can use type introspection and type assertions
    // to fetch the column into a typed variable.
}


```


## The Connection Pool

`database/sql`包中有一个基本的`连接池`。没有很多`控制`或`检查`功能，但是一些您需要知道的对您有用的内容如下：

- `连接池`意味着，在单个数据库上执行两个连续的语句，可能会打开两个连接并分别执行它们。这非常容易让程序员感到困惑，为什么他们的代码行为不当。例如，后跟`INSERT`的`LOCK TABLES`可能会阻塞，因为`INSERT`是在一个`不持有表锁的连接`上。
- 当`连接池`中没有可用的连接，且在需要时，才会创建新连接。
- 默认情况下，连接的数量没有限制的。如果您想一次性执行很多操作，则可以创建任意数量的连接。这可能导致数据库返回错误，例如`"to many connections"`。
- 在`Go 1.1或更新的版本`中，可以使用`db.SetMaxIdleConns(N)`来限制`连接池`中`空闲连接的数量`。不过，这并不限制`连接池的大小`。
- 在`Go 1.2.1或更新的版本`中，可以使用`db.SetMaxOpenConns(N)`来限制访问数据库打开的连接总数。不幸的是，[死锁错误](https://groups.google.com/d/msg/golang-dev/jOTqHxI09ns/x79ajll-ab4J)（[修复程序](https://code.google.com/p/go/source/detail?r=8a7ac002f840)）阻止了`db.SetMaxOpenConns(N)`在`Go 1.2`中的安全使用。
- 连接会被相当快速地被回收利用。使用`db.SetMaxIdleConns(N)`设置大量的`空闲连接`可以减少这种混乱，并有助于保持连接可重复使用。
- `长时间保持一个连接空闲`会导致问题（例如在`Microsoft Azure`上的`MySQL`中出现的[问题](https://github.com/go-sql-driver/mysql/issues/257)）。如果由于连接空闲时间太长而导致连接超时，请尝试`db.SetMaxIdleConns(0)`。
- 您还可以通过设置`db.SetConnMaxLifetime(duration)`来指定`可重用连接的最长生存时间`，因为重用寿命长的连接可能会导致网络问题。这会延迟关闭不用的连接，即可能会延迟关闭过期的连接。


## Surprises, Antipatterns and Limitations

尽管一旦您习惯了`database/sql`包便觉得它很简单，但您应该会对它所支持的用例的精妙之处感到惊讶。这点精妙是`Go的核心库`所共有的。


### Resource Exhaustion

正如本文当所贯彻所述的，如果您没有按预期来使用`database/sql`，则肯定会给自己造成麻烦，通常是耗废一些资源或妨碍其有效地重复使用：

- 打开和关闭数据库可能会导致资源耗尽。
- 无法读取所有行或使用`rows.Close()`保留连接池中的连接。
- 对不返回行的语句使用`Query()`将保留来自连接池的连接。
- 不了解[预备语句](#using-prepared-statements)的工作方式，会导致大量额外的数据库活动。


### Large `uint64` Values

这里是一个令人惊讶的错误。如果`64位的无符号整数`的`高位`被设置，则不能将其作为参数传递给语句：

```go

_, err := db.Exec("INSERT INTO users(id) VALUES", math.MaxUint64)   // Error

```

这将引发错误。如果使用`uint64`值，请当心，因为它们开始时很小可以正常工作，但是随着时间的推移会增加并开始引发错误。


### Connection State Mismatch

有些事情会`改变连接状态`，这可能会引起问题，原因有两个：

1. 某些连接状态（例如您是否在事务中）应改为通过`Go类型`来处理。
2. 您可能会认为查询是在单个连接上运行的，但实际不是。

例如，对于许多人来说，使用`USE语句`设置当前数据库是很典型的事情。但是`在Go中`，它只能影响运行了`USE语句`的连接本身。`除非您处于事务中`，否则您认为在该连接上执行的其他语句实际上可能会在从连接池中获得的其他连接上运行，而其他连接不会看到这种变化的影响。

另外，您改变了连接后，它将返回到连接池中，并潜在地污染一些其他代码的状态。这也是为什么让您不要直接将`BEGIN`或`COMMIT`语句作为`SQL命令`直接发出的原因之一。


### Database-Specific Syntax

`database/sql`包的API提供了`面向行的数据库`的`抽象`，但是`特定的数据库和驱动程序`的`行为`和/或`语法`可能有所不同，例如[`预备语句占位符`](#parameter-placeholder-syntax)。


### Multiple Result Sets

`Go驱动程序`不以任何方式支持`单个查询`的`多个结果集`，也似乎没有任何要实现该功能的计划，尽管有[功能请求](https://github.com/golang/go/issues/5171)想要支持`批量操作`，例如`批量复制`。

这意味着，除了别的之外，连`返回多个结果集`的`存储过程`也是无法正常工作的。


### Invoking Stored Procedures

`调用存储过程`是`特定于驱动程序的`，但是在`MySQL驱动程序`中，目前尚无法完成。通过如下执行操作，您似乎可以调用`返回单个结果集`的简单`存储过程`：

```go

err := db.QueryRow("CALL mydb.myprocedure").Scan(&result)   // Error

```

实际上，`这是行不通的`。 您将得到错误：`Error 1312: PROCEDURE mydb.myprocedure can't return a result set in the given context`。这是因为`MySQL数据库`希望将`连接`设置为`多语句模式`，即使是单个结果也是如此，而`驱动程序`当前不支持这样做（尽管有人提出过[此问题](https://github.com/go-sql-driver/mysql/issues/66)）。


### Multiple Statement Support

`database/sql`没有明确地支持`多条语句`，这意味着其行为取决于其内部后台：

```go

_, err := db.Exec("DELETE FROM tbl1; DELETE FROM tbl2")     // Error/unpredictable result

```

这里允许服务器根据需要进行解释，可能包括`返回错误`，`仅执行第一条语句`或`执行两项`。

同样，没有办法在`事务`中`批处理语句`。`事务`中的每个语句必须按顺序执行，并且结果中的资源（例如，`Row（一行）`或`Rows（多行）`）必须进行`扫描`或`关闭`，来释放`底层的连接`以便`下一条语句`使用。这与您不使用`事务`时的通常行为不同。`在不使用事务的情况下`，完全有可能执行查询，循环遍历行，并在循环内对数据库进行查询（该查询将在新连接上发生）：

```go

rows, err := db.Query("SELECT * FROM tbl1")     // Uses connection 1
for rows.Next() {
    err = rows.Scan(&myvariable)
    // The following line will NOT use connection 1, which is already in-use
    db.Query("SELECT * FROM tbl2 WHERE id=?", myvariable)
}

```

但是`事务`只能绑定到一个连接，因此`事务`不可能做到这一点：

```go

tx, err := db.Begin()
rows, err := tx.Query("SELECT * FROM tbl1")     // Uses tx's connection
for rows.Next() {
    err = rows.Scan(&myvariable)
    // ERROR! tx's connection is already busy!
    tx.Query("SELECT * FROM tbl2 WHERE id=?", myvariable)
}

```

不过，`Go`并不会阻止您这样尝试。因此，如果您试图在第一个语句释放其资源并随后对其进行清理之前尝试执行另一个语句，则可能会导致`连接损坏`。这也就是说，`事务中的每个语句`都会引起一组单独的访问数据库的网络往返。


## Related Reading and Resources

以下是我们发现一些有帮助的外部信息源。

- [http://golang.org/pkg/database/sql/](golang.org/pkg/database/sql/)
- [http://jmoiron.net/blog/gos-database-sql/](jmoiron.net/blog/gos-database-sql/)
- [http://jmoiron.net/blog/built-in-interfaces/](jmoiron.net/blog/built-in-interfaces/)
- VividCortex博客，例如 [透明加密](https://vividcortex.com/blog/2014/11/11/encrypting-data-in-mysql-with-go/)
- 来自[golang github](https://github.com/golang/go/wiki/SQLDrivers)的SQL驱动程序概述

希望本文档对您有所帮助。如果您有任何改进建议，请在[https://github.com/VividCortex/go-database-sql-tutorial](https://github.com/VividCortex/go-database-sql-tutorial)发送请求或提出问题报告。

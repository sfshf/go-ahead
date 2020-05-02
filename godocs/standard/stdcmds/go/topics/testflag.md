
[Testing flags](https://golang.google.cn/cmd/go/#hdr-Testing_flags)


##### 说明：

`go test`命令可以同时接受应用于`go test`本身的标志和应用于将要生成的测试二进制文件的标志。

几个标志可以控制配置和写一个适合于`go tool pprof`的执行配置文件；运行`go tool pprof -h`可以获取更多信息。`pprof`的`--alloc_space`、`--alloc_objects`和`--show_bytes`选项可以控制信息的显示方式。

以下的标志可以被`go test`命令识别，并控制任何测试的执行：

| 标志 | 说明 |
|--|--|
| -bench regexp | 仅运行与一个正则表达式相匹配的那些基准测试。默认情况下，不运行任何基准测试。要运行所有基准测试，请使用"-bench ."或"-bench=."。正则表达式是被未使用方括号括起来的斜杠符号（"/"）分隔成一个正则表达式序列，并且一个基准测试标识符（如果有的话）的每一部分必须匹配序列中的对应元素。 匹配项的可能父级项能以"b.N=1"来运行，来标识子基准测试。例如，给定的`-bench=X/Y`，与"X"匹配的顶级基准测试以"b.N=1"来运行，来查找与"Y"匹配的任何子基准测试，然后被找到的子基准测试全部都会被运行。 |
| -benchtime t | 对每个基准测试迭代循环运行"t"时间，"t"时间被指定为一个"time.Duration"对象（例如，"-benchtime 1h30s"）。默认值为"1秒"（1s）。特殊语法"Nx"，意味着运行基准测试"N"次（例如，"-benchtime 100x"）。 |
| -count n | 运行每个测试和基准测试"n"次（默认是"1"次）。如果设置了"-cpu"标志，则对"GOMAXPROCS"的每个值都运行"n"次。 |
| -cover | 启用覆盖率分析（coverage analysis）。请注意，由于覆盖率是通过注释编译之前的源代码来起作用的，所以因为启用覆盖率分析而导致的编译和测试失败可能报告的行数字与原始的源文件不对应。 |
| -convermode set, count, atomic | 为测试包设置覆盖率分析的模式。默认是"set"模式；如果"-race"被启用，则是"atomic"模式。不同模式的值为："set"--bool，表示当前语句是否运行；"count"--int，表示当前语句要运行多少次；"atomic"--int，用来计数，在多线程测试中计算成功数；该模式开销相当大。需要启用"-cover"标志。 |
| -coverpkg pattern1,pattern2,pattern3 | 将每个测试中的覆盖率分析应用到与指定模式匹配的包。默认情况是，每个测试只分析被测试的包。"go help packages"可以查看有关包模式的描述。需要启用"-cover"标志。 |
| -cpu 1,2,4 | 为测试或基准测试指明一个"GOMAXPROCS"值的列表。默认情况下是"GOMAXPROCS"的当前值。 |
| -failfast | 在第一次测试失败之后，不再启动新的测试。 |
| -list regexp | 列出与正则表达式匹配的测试、基准测试或示例。不会运行测试、基准测试或示例。该标志只会列出顶层的测试，不会展示子测试或子基准测试。 |
| -parallel n | 允许并行执行调用了"t.Parallel"的测试函数。该标志的值是指要同时运行的最大测试数；默认情况下，它设置为"GOMAXPROCS"的值。请注意，"-parallel"仅适用于单个测试二进制文件。根据"-p"标志的设置，"go test"命令也可以并行地对不同的包运行测试（更多可查看"go help build"）。 |
| -run regexp | 只运行与正则表达式匹配的那些测试和示例。对于测试，"regexp"是用不带括号的斜杠符号（"/"）分隔为一个正则表达式序列，一个测试标识符的每个部分必须与序列中的相应元素匹配。 请注意，匹配项的父级也会被运行；因此，"-run=X/Y"会在即使没有匹配"Y"的子测试的情况下，也可能会匹配"X"的测试，然后运行并报告结果，因为必须运行它们才能查找那些子测试。 |
| -short | 要求长时间运行的测试缩短其运行时间。默认情况下它是关闭的，但是在"all.bash"文件运行时会被设置，所以安装"Go SDK包树"时会运行健全性检查，而不是花时间运行详尽的测试。 |
| -timeout d | 如果测试二进制文件的运行时间超过指明的持续时间"d"，则会宕机（"panic"）。如果"d"的值为"0"，则"timeout"标志被禁用。默认值为10分钟（10m）。 |
| -v | 详细的输出：打印出所有测试的日志。 还会打印所有来自"Log"和"Logf"调用的文本，即使测试成功。 |
| -vet list | 在"go test"期间，配置调用"go vet"命令，使用逗号分隔的审核清单。如果"list"为空，则"go test"会使用经过仔细挑拣且值得的检查来运行"go vet"命令。如果"list"为"off"，则"go test"根本不会运行"go vet"。 |

以下的标志也会被`go test`命令识别，并被用来在执行时配置测试：

| 标志 | 说明 |
|--|--|
| -benchmem | 打印出基准测试的内存分配统计。 |
| -blockprofile block.out | 在所有测试完成后，将一个goroutine阻塞分析（a goroutine bloking profile）写入指定的文件。像"-c"标志那样编写测试二进制文件。 |
| -blockprofilerate n | 通过调用传入"n"的"runtime.SetBlockProfileRate"函数来控制goroutine阻塞分析里的细节。详细信息请参阅"go doc runtime.SetBlockProfileRate"。分析器（profiler）旨在每当程序平均阻塞"n"纳秒时采样一次阻塞状态。 默认情况下，如果"-test.blockprofile"启用时不带"-blockprofilerate"标志，则所有阻塞事件都会被记录下来，相当于"-test.blockprofilerate=1"。 |
| -coverprofile cover.out | 所有测试通过后，将覆盖率分析写入文件。需要启用"-cover"标志。 |
| -cpuprofile cpu.out | 退出之前，将CPU分析写入指定的文件。像"-c"标志那样编写测试二进制文件。 |
| -memprofile mem.out | 所有测试通过后，将内存分配分析写入文件。像"-c"标志那样编写测试二进制文件。 |
| -memprofilerate n | 通过设置"runtime.MemProfileRate"来启用更精确的（同时开销更大）内存分配分析。详细信息请参阅"go doc runtime.MemProfileRate"。要分析所有的内存分配，请使用"-test.memprofilerate=1"。 |
| -mutexprofile mutex.out | 在所有测试完成后，将互斥锁争用分析写入指定的文件。像"-c"标志那样编写测试二进制文件。 |
| -mutexprofilefraction n | 对持有竞争互斥锁（a contended mutex）的goroutine的栈踪迹（stack trace）进行"n"分之一的采样。 |
| -outputdir directory | 将分析的输出文件放在指定目录中，默认情况下，放在运行"go test"命令的目录下。 |
| -trace trace.out | 退出之前，将执行跟踪记录写入指定的文件。 |.

以上的标志也能以另一个可选的方式被识别，即有一个"test."前缀，例如"-test.v"。但是，当直接调用生成的测试二进制文件（即："go test -c"的结果）时，"test."前缀是必需的。

"go test"命令会在调用测试二进制文件之前，视情况在可选软件包列表之前和之后重写或除去已识别的标志。举例如下：

```sh
go test -v -myflag testdata -cpuprofile=prof.out -x
```

该命令将编译测试二进制文件并如下运行：

```sh
pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out
```

通过上面举例的命令可见，`-x`标志被删除了，因为该标志只应用于`go命令`的运行，而不应用于测试本身。

生成分析文件（profile）的测试标志（除了覆盖率以外）还将测试二进制文件保留在`pkg.test`中，以便在解析分析文件时使用。

当`go test`运行测试二进制文件时，它将在相应程序包的源代码目录中运行。根据测试，直接调用生成的测试二进制文件时可能需要执行相同的操作。

命令行中指名的包列表（如果有的话）必须出现在`go test`命令未知的任何标志之前。继续上面的示例，包列表必须出现在`-myflag`之前，但可以出现在`-v`的任一侧。

当`go test`以包列表模式运行时，`go test`会缓存成功的包测试结果，以避免不必要的重复运行测试。要禁用测试缓存，请使用除可缓存标志以外的任何测试标志或参数。显式禁用测试缓存的惯用方式是使用`-count=1`。

为了防止传给测试二进制文件的参数被解释为同名的标志或包名称，请使用`-args`标志（请参阅`go help test`），它将命令行的剩余部分直接传递给测试二进制文件，不进行解释和更改。举例：

```sh
go test -v -args -x -v  # 第一个示例
```

该命令将编译测试二进制文件，并如下运行：

```sh
pkg.test -test.v -x -v
```

同样地：

```sh
go test -args math  # 第二个示例
```

该命令将编译测试二进制文件，并如下运行：

```sh
pkg.test math
```

在第一个示例中，`-x`和第二个`-v`不变地传递到测试二进制文件，并且对`go命令`本身也没有影响。 在第二个示例中，参数`math`传递给测试二进制文件，而不是被解释为包列表。

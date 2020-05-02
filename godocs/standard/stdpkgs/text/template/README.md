

[text/template](https://golang.google.cn/pkg/text/template/)


### 一、简介

`text/template`包为生成文本输出实现了数据驱动模板；`html/template`包与`text/template`包实现了相同的接口，可以生成HTML输出，而且还能自动保证HTML输出的安全足以抵御一些攻击。

模板通过应用于数据结构来被执行；模板中的注解指向数据结构的元素（通常是结构体的字段或一个map的键）来控制执行并导出被展示的值。模板的执行可以在结构上游走并设置`游标`，`游标`通过一个`.`符号（dot）来表示；随着执行的进行，游标指向结构当前位置上的值。

模板的输入文本可以是任何`UTF-8编码`的文本格式。`Action（操作）`--数据求值或控制结构--通过`{{`和`}}`来定界；所有在`action`之外的文本会不作改变地复制到输出中。`action`不能跨行，除非是原始字符串（raw string）。注释是可以换行的。

一旦解析（parse）完成，模板是可以安全地并行执行的，尽管，如果并行执行是共享一个编辑器（writer）会导致输出交错。

下面是一个简单示例，该示例会打印"17 items are made of wool"。

```go

type Inventory struct {
        Material string
        Count uint
}

sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil { panic(err) }
err = templ.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }

```


### 二、文本和空格

默认情况下，当模板被执行时，`action`之间的所有文本会被逐字地复制。正如上面的简例程序运行时，"items are made of"会出现在标准的输出中。

但是，为了格式化模板源代码，如果`action`的左定界符（默认是`{{`）后面紧跟着一个`-`号和`ASCII`编码的空格，即`{{- `，后面直到出现的文本之间的空白（white space）会被剪裁掉；相似地，如果右定界符（默认是`}}`）前面紧跟着一个`ASCII`编码的空格和`-`号，即` -}}`，前面直到出现的文本之间的空白（white space）会被剪裁掉。在这些剪裁标识中，必须要有ASCII空格；如`{{-3}}`会被解析为一个包含数字`-3`的`action`。举例，当执行源码为`{{23 -}} < {{- 45}}`的源码时，会生成输出为`23<45`。

这样的剪裁，空白（white space）字符的定义和`Go`语言里的一样，有`空格（space）`、`水平制表符（horizontal tab）`、`回车符（carriage return）`和`换行（newline）`。


### 三、操作（Actions）

下面是操作的列表。`Arguments`和`pipelines`是对数据的求值，后面有相应的节选进行详细说明。

| 操作 | 说明 |
|--|--|
| {{/* a comment */}} {{- /* a comment with white space trimmed from preceding and following text */ -}} | 一条注释；会被丢弃。可能包含换行符。注释不允许嵌套，并且在定界符内必须是完整的注释。 |
| {{pipeline}} | 流水线（pipeline）的结果的默认文本表示，会被复制到输出里。 |
| {{if pipeline}} T1 {{end}} | 如果流水线（pipeline）的值是空值（empty），则不会生成输出；否则，`T1`会被执行。空值可以是false、0、任何空指针、任何空接口值、任何空数组、任何空切片、任何空map或长度为零的字符串。游标（dot）不受影响。 |
| {{if pipeline}} T1 {{else}} T0 {{end}} | 如果流水线（pipeline）的值是空值（empty），`T0`会被执行；否则，`T1`会被执行。游标（dot）不受影响。 |
| {{if pipeline}} T1 {{else if pipeline}} T0 {{end}} | 为了简化`if-else`链的外观，因为一个`if`条件句的`else`操作里可能直接包含另一个`if`条件句；该`action`的效果和`{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}`一样。 |
| {{range pipeline}} T1 {{end}} | 流水线（pipeline）的值必须是一个数组、切片、map或channel。如果流水线（pipeline）的值的长度是零，则输出为空；否则，游标（.）会遍历数组、切片、map或channel的值，并执行`T1`。如果流水线（pipeline）的值为map且map的键是定义了顺序的基础类型（"comparable"），则按map的键的排序顺序来访问map元素。 |
| {{range pipeline}} T1 {{else}} T0 {{end}} | 流水线（pipeline）的值必须是一个数组、切片、map或channel。如果流水线（pipeline）的值的长度是零，游标不受影响，并且`T0`会被执行；否则，游标（.）会遍历数组、切片、map或channel的值，并且`T1`会被执行。 |
| {{template "name"}} | 被指名的模板会被空数据地（with nil data）执行。 |
| {{template "name" pipeline}} | 被指名的模板会被执行，且被指名的模板的游标（dot）被设置为流水线（pipeline）的结果。 |
| {{block "name" pipeline}} T1 {{end}} | `block（块）`是定义一个模板的简写，是`{{define "name"}} T1 {{end}}`的简写；然后像`{{template "name" pipeline}}`一样地执行。通常的用途是定义一组根模板，然后通过重新定义其中的块模板进行自定义。 |
| {{with pipeline}} T1 {{end}} | 如果流水线（pipeline）的值是空值（empty），则不会生成输出；否则，游标（dot）会被设置为流水线（pipeline）的结果，并且`T1`会被执行。 |
| {{with pipeline}} T1 {{else}} T0 {{end}} | 如果流水线（pipeline）的值是空值，游标（dot）不受影响，并且`T0`会被执行；否则游标（dot）会被设置为流水线（pipeline）的结果，并且`T1`会被执行。 |


### 四、形参（Arguments）

下面是形参的列表。

| 形参 | 说明 |
|--|--|
| 使用Go语言语法的布尔型、字符串型、字符型、整型、浮点型、虚数型或复数型的`常量`。这些`常量`与Go语言的`无类型常量`有同样的行为方式。 | 注意，就像在Go语言中，一个超大的整型常量在赋值或传递给一个函数时，是否会发生溢出取决于`主机`的整型是`32位`还是`64位`。 |
| 关键字`nil`。 | 表示一个Go语言的`无类型nil`。 |
| `.`字符。 | 其结果为游标（dot）的值。 |
| `变量名`，是一个`$`符号开头和字母数字组成的字符串，也可以是单独一个`$`符号。例如`$piOver2`或`$`。 | 其结果为该变量的值。`Variables`在之后的节选中有详细说明。 |
| `结构体的字段名`，前面加一个`.`号，例如`.Field`。 | 其结果为该字段的值。字段的调用可以是链式的，例如`.Field1.Field2`。字段还可以在`变量`上求值，同样可以是链式的，例如：`$x.Field1.Field2`。 |
| `map的键名`，前面加一个`.`号，例如`.Key`。 | 其结果为该map的键映射的map的值。map键的调用也可以是链式的，并且可以和`字段名`组合至任何深度，例如：`.Field1.Key1.Field2.Key2`。尽管map键必须是一个字母数字组成的标识符，但是不像`字段名`，它们不需要开头字母大写。map键还可以在`变量`上求值，同样可以是链式的，例如：`$x.key1.key2`。 |
| 数据的`无参数方法（a niladic method）`名，前面加一个`.`号，例如`.Method`。 | 其结果是调用以游标（dot）为接受者的方法获得的值，即`dot.Method()`。这样的方法必须有一个返回值（可以是任何类型）或两个返回值（第二个返回值是`error`类型），如果是两个返回值，且`error`不为空，执行将会终止并将`error`作为执行的值返回给调用者。方法的调用也可以是链式的或与`字段`和`map键`组合到任意深度，例如`.Field1.key1.Method1.Field2.key2.Method2`。方法还可以在`变量`上执行，同样可以是链式的，例如`$x.Method.Field`。 |
| `无参数函数（a niladic function）`名，例如`fun`。 | 其结果是调用函数获得的值，例如`fun()`。调用函数返回的值和类型与`无参数方法`的行为一致。后续有节选介绍函数和函数名的详细信息。 |
| 上面各种形参使用`()`括起来形成的组合实例。 | 其结果可以被`字段`或`map键`接收并调用，例如`print (.F1 arg1) (.F2 arg2)`，`(.StructValueMethod "arg").Field`。 |

`形参`可以被求值为任何类型；如果它们是指针，则实现将在需要时自动间接指向基本类型。如果求值为一个函数（例如，结构体的函数字段），则不会自动调用该函数，但可以将其用作`if`等操作的实参。如何调用函数，后续有节选详细阐述。


### 五、流水线（Pipelines）

流水线可以是链式的一系列命令（commands）。命令是指简单的值（形参）或函数和方法的调用（可以有多个形参）。

| 流水线 | 说明 |
|--|--|
| 形参（Argument） | 其结果为该形参的求值。 |
| .Method [Argument...] | 该方法可以是单独的或一个链式形参的最后一个元素，不像链式形参的中间的方法，该方法可以获取形参。其结果是该方法获取形参后的求值，例如`dot.Method(Argument1, etc.)`。 |
| functionName [Argument...] | 其结果是调用该函数获得的值，例如`function(Argument1, etc.)`。后续有节选介绍函数和函数名的详细信息。 |

流水线可以是链式的，即通过流水线字符`|`分隔的命令序列。在链式流水线中，每个命令的结果会被作为下一个命令的最后的形参；最后一个命令的输出结果就是该链式流水线的输出结果。

命令的输出可以是一个值或两个值（第二个值是`error`类型）。 如果存在第二个值且求值为非空，那么执行终止，并且将`error`返回给执行的调用方。


### 六、变量（Variables）

操作（action）里的流水线（pipeline）可以初始化一个变量来捕获其结果；初始化的语法如下：

```go
$variable := pipeline
```

其中`$variable`是该变量的名字。声明变量的操作（action）不会产生输出。

被定义过的变量还可以被赋值；其语法如下：

```go
$variable = pipeline
```

如果在`range`操作中初始化变量，那么该变量会遍历该迭代器的值；`range`操作还可以声明两个变量，要使用`,`隔开，其语法如下：

```go
range $index, $element := pipeline
```

其中`$index`和`$element`会分别地遍历数组、切片和map，`$index`遍历数组/切片的序号或map的键的值，`$element`会遍历数组/切片或map的元素的值。注意，如果只有一个变量，该变量将会遍历数组/切片或map的元素；这一点与`Go语言的range语句`相反。

控制结构（`if`、`with`或`range`）中声明的变量的作用域，直到该控制结构的`end`操作；如果不是控制结构中声明的变量，其作用域为当前模板；模板的调用不能从该模板继承出内部变量。

当执行开始后，`$`会被设置为数据形参，传递给执行方，也就是传给游标（dot）。


### 七、示例（Examples）

下面是一些单行的模板示例，用来演示流水线和变量。所有的输出都是带引号的`"output"`。

| 示例 | 说明 |
|--|--|
| {{"\"output\""}} | 一个字符串`常量`。 |
| {{`"output"`}} | 一个`原始`字符串`常量`。 |
| {{printf "%q" "output"}} | 一个函数调用。 |
| {{"output" | printf "%q"}} | 一个函数调用，其最后的形参来自于前一个命令。 |
| {{printf "%q" (print "out" "put")}} | 一个使用括号组合的形参。 |
| {{"put" | printf "%s%s" "out" | printf "%q"}} | 一个更复杂的调用方式。 |
| {{"output" | printf "%s" | printf "%q"}} | 一个链式流水线。 |
| {{with "output"}}{{printf "%q" .}}{{end}} | 一个使用了`游标（dot）`的`with`操作。 |
| {{with $x := "output" | printf "%q"}}{{$x}}{{end}} | 一个创建并使用变量的`with`操作。 |
| {{with $x := "output"}}{{printf "%q" $x}}{{end}} | 另一个创建并使用变量的`with`操作。 |
| {{with $x := "output"}}{{$x | printf "%q"}}{{end}} | 同上，但使用了链式流水线。 |


### 八、函数（Functions）

在执行期间，可以在两个函数map中找到所需函数：一个是在模板中，另一个是在全局函数map中。默认情况下，模板中没有定义任何函数；但可以使用`Funcs方法`来添加函数。

预定义的全局函数命名如下。

| 函数名 | 说明 |
|--|--|
| and | 布尔和（boolean AND）；返回该函数形参的布尔和，如果第一个形参为空，则返回第一个形参，否则返回第二个形参；也就是说，`and x y`和`if x then y else x`的行为一致。所有的形参都会被求值。 |
| call | 返回第一个形参的调用结果，该第一个形参必须是一个函数，之后的形参是该函数将接收的实参。举例，`call .X.Y 1 2`，用Go语言语法即表述为`dot.X.Y(1, 2)`，`Y`是一个函数类型字段、map项或其他可能的项。第一个形参必须是求值结果为函数类型，这里是为了与比如`print`这样的预定义的函数作区分。函数必须是一个或两个返回值，其中第二个返回值是`error`类型。如果形参与函数不匹配或返回的`error`值是非空的，执行会停止。 |
| html | 返回转义的HTML脚本，等价于该函数的形参的文本表述。这个函数在`html/template`包内是不可用，会有一些异常。 |
| index | 返回该函数的第一个形参中以后续形参为各级索引所获得的值。例如，`index x 1 2 3`，用Go语言语法即表述为`x[1][2][3]`。注意，各级被索引的对象必须是map、切片或数组。 |
| slice | 返回对该函数第一个形参使用后续形参进行切片的结果。例如，`slice x 1 2`，用Go语言语法即表述为`x[1:2]`，而`slice x`即为`x[:]`，`slice x 1`即为`x[1:]`，以及`slice x 1 2 3`即为`x[1:2:3]`。第一个形参必须是一个字符串、切片或者数组。 |
| js | 返回转义的JavaScript脚本，等价于该函数的形参的文本表述。 |
| len | 返回该函数的形参的长度（整型值）。 |
| not | 布尔非（boolean negation）；只能接收单个形参，并返回该形参的非值。 |
| or | 布尔或（boolean OR）；返回该函数形参的布尔或，如果第一个形参为非空，则返回第一个形参，否则返回第二个形参；也就是说，`or x y`和`if x then x else y`的行为一致。所有的形参都会被求值。 |
| print | `fmt.SPrint`函数的别名。 |
| printf | `fmt.SPrintf`函数的别名。 |
| println | `fmt.SPrintln`函数的别名。 |
| urlquery | 返回该函数形参的文本表述的转义值，该转义值用是一种适合嵌入到URL查询里的形式。这个函数在`html/template`包内是不可用，会有一些异常。 |

注意，布尔函数（boolean functions）会将任何零值当成`false`，以及非零值当成`true`。

下面是被定义为函数的二进制比较操作符（binary comparison operators）。

| 函数名 | 说明 |
|--|--|
| eq | 返回`arg1 == arg2`的布尔值。 |
| ne | 返回`arg1 != arg2`的布尔值。 |
| lt | 返回`arg1 < arg2`的布尔值。 |
| le | 返回`arg1 <= arg2`的布尔值。 |
| gt | 返回`arg1 > arg2`的布尔值。 |
| ge | 返回`arg1 >= arg2`的布尔值。 |

对于更简单的`多路相等性测试（multi-way equality tests）`，`eq（也只有eq）`函数可以接受两个或多个形参，并将第二个形参及后续形参分别与第一个形参进行比较，并返回有效值，就像Go语言里的`arg1==arg2 || arg1==arg3 || arg1==arg4 ...`，但与`||`不同的是，`eq`是一个函数调用，并且所有形参都会被求值。

比较函数仅适用于基本类型（或被命名的基本类型，例如`type Celsius float32`）。这些函数实现了Go语言的规则来进行值的比较，只是会忽略大小和精准类型，因此可以将任何符号或无符号的整数与任何其他整数值进行比较。（算数值的比较不是`位模式（bit pattern）`，因此所有负整数都小于所有无符号整数。）但是，像通用的一样，无法将`int类型`与`float32类型`进行比较，其他依此类推。


### 九、关联模板

每个模板均在被创建时指定的字符串命名。此外，每个模板都可以通过名称调用来与零个或多个其他模板相关联；这样的关联是可传递的，并形成模板的名称空间。

模板可以使用模板调用来实例化另一个关联的模板。请参阅之前节选中对`template`操作（action）的说明。 该名称必须是与包含调用的模板关联的模板的名称。


### 十、内嵌模板定义

当解析一个模板时，另一个模板可能被定义且与被解析的模板相关联。模板的定义必须出现在该模板的顶层，很类似于Go语言程序里的全局变量。

定义模板的语法是，使用一个`define`和`end`组合操作将模板的声明包裹其中。

`define`操作通过提供一个字符串常量来命名一个正被创建的模板。下面是一个简例：

```

`{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`

```

上面的简例先定义了两个模板，`T1`和`T2`，再定义了第三个`T3`模板，该模板在执行时会调用另两个模板；最后调用`T3`模板。如果执行该模板将会产生如下文本：

```
ONE TWO
```

通过构造，模板可以仅驻留在一个关联中。如果需要一个可从多个关联中寻址的模板，则必须多次解析模板定义以创建不同的`*Template`值，或者必须使用`Clone`或`AddParseTree`方法将其复制。

可以多次调用`Parse`来组装不同的相关模板。有关解析存储在文件中的相关模板的简单方法，请参见`ParseFiles`和`ParseGlob`函数和方法。

可以直接执行模板；也可以通过`ExecuteTemplate`执行模板，该模板执行方式可以执行一个由名称标识的关联模板。为了调用上面的模板简例，我们可以这样写：

```go

err := tmpl.Execute(os.Stdout, "no data needed")
if err != nil {
    log.Fatalf("execution failed: %s", err)
}

```

或者，显示地用名字调用一个特定的模板：

```go

err := tmpl.ExecuteTemplate(os.Stdout, "T2", "no data needed")
if err != nil {
    log.Fatalf("execution failed: %s", err)
}

```

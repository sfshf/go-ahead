
# `regexp`包

`regexp`包实现了`正则表达式搜索`。

`能被接受的正则表达式的语法`与`Perl`，`Python`和`其他语言`使用的常规语法`相同`。更准确地说，它是`RE2`能接受的语法，并在[https://golang.org/s/re2syntax](https://golang.org/s/re2syntax)中进行了描述，除了`\C`。有关语法的概述，请运行：

```sh

go doc regexp/syntax

```

本包提供的`regexp的实现`，保证运行是与输入大小的线性相关的用时。（这是大多数正则表达式的开源实现都不保证的属性。）有关此属性的更多信息，请参见[https://swtch.com/~rsc/regexp/regexp1.html](https://swtch.com/~rsc/regexp/regexp1.html)，或任何有关`自动机理论`的书。

`所有字符都是UTF-8编码的码点`。

`Regexp`结构体有`16`种匹配正则表达式并标识匹配文本的方法。它们的名称与如下正则表达式匹配：

```txt

Find(All)?(String)?(Submatch)?(Index)?

```

如果存在`'All'`，则例程将匹配`整个表达式`的`连续不重叠匹配`。与前一个匹配紧邻的空匹配将会被忽略。返回值是一个切片，其中包含相应的`非'All'例程`的连续返回值。这些例程使用一个额外的整型参数`n`。如果`n >= 0`，则该函数最多返回`n个匹配项/子匹配项`；否则，将返回所有匹配项/子匹配项。

如果存在`'String'`，则参数为字符串；否则，它是一个`字节切片`；返回值会适当调整。

如果存在`'Submatch'`，则返回值是一个切片，用于标识表达式的连续的子匹配项。`子匹配项`是正则表达式中`带括号的子表达式`（也称为`捕获组`）的匹配项，按左括号的顺序从左到右编号。`子匹配项0`是`整个表达式的匹配项`，`子匹配项1`是`第一个括号内的子表达式的匹配项`，依此类推。

如果存在`'Index'`，则`匹配和子匹配`由`输入字符串中的字节索引对`标识：`result[2*n:2*n+1]`标识了`第n个子匹配的索引`。`n == 0`的对表示`整个表达式的匹配`。如果不存在`'Index'`，则匹配由`匹配/子匹配的文本`标识。如果`索引为负`或`文本为nil`，则意味着`子表达式`与`输入中的任何字符串`都不匹配。对于`'String'`的版本，`空字符串`表示`不匹配`或`空匹配`。

还有一些方法的子集可应用于从`RuneReader`读取的文本：

```txt

MatchReader, FindReaderIndex, FindReaderSubmatchIndex

```

这个集合里的方法可能还会增加。请注意，正则表达式匹配可能需要检查匹配返回的文本之外的文本，因此匹配来自`RuneReader`的文本的方法可能会在返回之前`任意读取`输入内容。

（还有一些其他方法与此模式不匹配。）
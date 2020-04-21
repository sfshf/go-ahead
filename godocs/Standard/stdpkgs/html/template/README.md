

[html/template](https://golang.google.cn/pkg/html/template/)


### 一、简介

`html/template`包为生成HTML输出实现了数据驱动模板，生成的HTML输出是安全可靠的，可以抵御代码注入。该包与`text/template`包一样提供了同样的接口，并在输出为HTML文本时，可以使用该包替代`text/template`包。

当前文档主要阐述`html/template`包的安全特性；如果想要知道如何编写模板，请查看`text/template`包的文档。

`html/template`包是对`text/template`包的包装，所以您可以使用一样的模板API来安全地解析和执行HTML模板。例如：

```go
tmpl, err := template.New("name").Parse(...)
// Error checking elided
err = tmpl.Execute(out, data)
```

如果模板解析成功，则该模板就是注入安全（injection-safe）的了；否则，`err`为非空且是[`ErrorCode`](#)里定义的一种错误。

HTML模板处理数据就像需要被编码的无格式文本一样，所以模板可以安全地嵌入到一个HTML文档中。转义是上下文有关联的，所以操作（actions）可以出现在`JavaScript`、`CSS`和`URI上下文`中。

`html/template`包中使用到的安全模型是假定，模板的作者要是可信的，而执行（Execute）的数据实参可以是不可信的。后续节选有更详细的阐述。

举例：

```go
import "text/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

生成：

```
Hello, <script>alert('you have been pwned')</script>
```

而`html/template`包有`上下文自动转义功能（contextual autoescaping）`，如下：

```go
import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

生成安全的、被转义的HTML输出：

```
Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
```


### 二、上下文（Contexts）

`html/template`包能够识别HTML、CSS、JavaScript和URI；该包会给每个简单的操作流水线添加审查函数。举例：

```
<a href="/search?q={{.}}">{{.}}</a>
```

在解析时，如果需要的话，每个`{{.}}`会被重写，增加转义功能函数；这样的话，会变成如下所示：

```
<a href="/search?q={{. | urlescaper | attrescaper}}">{{. | htmlescaper}}</a>
```

其中，`urlescaper`、`attrescaper`和`htmlescaper`都是内部转移功能函数的别名。

对于这些内部转义功能函数，如果一个操作流水线求值为`nil`接口值，该值仍会像一个空字符串一样被处理。


### 三、错误（Errors）

可以查看[`ErrorCode`](https://golang.google.cn/pkg/html/template/#ErrorCode)源文档；以下是中文翻译。

`ErrorCode`是用来表示`error`种类的代码。定义如下

```go
type ErrorCode int
```

我们为在转义模板是出现的每个错误定义代码，但是转义好的模板在运行时也是有可能失败的。

举例，该示例会输出"ZgotmplZ"：

```
<img src="{{.X}}">
where {{.X}} evaluates to `javascript:...`
```

解释如下：

"ZgotmplZ"是一个特殊值，意味着运行时里CSS或URL上下文触及到不安全的内容。上面示例的输出将会是`<img src="#ZgotmplZ">`。如果数据来自一个受信任的源，可以使用内容类型（content types）来避免数据经过"URL\`javascript:...\`)"过滤。


### 四、有关上下文（Contexts）和错误（Errors）的详细阐述


#### 1、Contexts

假设`{{.}}`的传入值为`O'Reilly: How are <i>you</i>?`，下表展示了当使用左侧的上下文时`{{.}}`的实际值：

| Context | {{.}} After |
|--|--|
| `{{.}}` | O'Reilly: How are &lt;i&gt;you&lt;/i&gt;? |
| `<a title='{{.}}'>` | O&#39;Reilly: How are you? |
| `<a href="/{{.}}">` | O&#39;Reilly: How are %3ci%3eyou%3c/i%3e? |
| `<a href="?q={{.}}">` | O&#39;Reilly%3a%20How%20are%3ci%3e...%3f |
| `<a onx='f("{{.}}")'>` | O\x27Reilly: How are \x3ci\x3eyou...? |
| `<a onx='f({{.}})'>` | "O\x27Reilly: How are \x3ci\x3eyou...?" |
| `<a onx='pattern = /{{.}}/;'>` | O\x27Reilly: How are \x3ci\x3eyou...\x3f |

如果`{{.}}`被用于不安全的上下文，那么其值会被过滤出来：

| Context | {{.}} After |
|--|--|
| `<a href="{{.}}">` | #ZgotmplZ |

因为"O'Reilly:"是一个不被允许的协议（protocol），就像`http`。

如果`{{.}}`是无害化（innocuous）单词`left`，那么该单词可以更广泛地出现：

| Context | {{.}} After |
|--|--|
| `{{.}}` | left |
| `<a title='{{.}}'>` | left |
| `<a href='{{.}}'>` | left |
| `<a href='/{{.}}'>` | left |
| `<a href='?dir={{.}}'>` | left |
| `<a style="border-{{.}}: 4px">` | left |
| `<a style="align: {{.}}">` | left |
| `<a style="background: '{{.}}'>` | left |
| `<a style="background: url('{{.}}')>` | left |
| `<style>p.{{.}} {color:red}</style>` | left |

非字符串值可以被用在JavaScript上下文中。举例，如果`{{.}}`是`struct{A,B string}{ "foo", "bar" }`，且用在被转义的模板`<script>var pair = {{.}};</script>`中，那么该模板的输出为`&lt;script&gt;var pair = {"A": "foo", "B": "bar"};&lt;/script&gt;`。


#### 2、Typed Strings（被定义了类型的字符串）

默认情况下，`html/template`包假设，所有的流水线都是产生一个无格式（plain）文本字符串。`html/template`包添加了转义流水线的阶段过程，来正确和安全地在适当的上下文中嵌入无格式文本字符串。

当数据值不是纯文本时，可以通过将其标记为自身的类型来确保它不会被过度转义。

来自`content.go`的`HTML`、`JS`、`URL`等类型可以携带安全的内容，这些内容可以免于转义。

举例，一个模板`Hello, {{.}}!`可以用```tmpl.Execute(out, template.HTML(`<b>World</b>`))```来产生`Hello, <b>World</b>!`而不是`Hello, &lt;b&gt;World&lt;b&gt;!`，后者会在`{{.}}`是一个常规字符串时产生。


#### 3、Security Model

`html/template`使用的是[safetemplate](https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition)文档中定义的安全设计。

`html/template`包假定模板作者是受信任的，而不是执行的数据实参，并试图在面对不受信任的数据时保留以下属性：

结构保留属性：“ ...当模板作者以安全的模板语言编写HTML标签时，浏览器会将输出的相应部分解释为标签，而不管不受信任的数据的值如何；对于其他结构，例如，属性边界以及JS和CSS字符串边界。”

代码效果属性：“ ...只有模板作者指明的代码应该通过将模板输出注入到页面的方式来运行，且模板作者所指定的所有代码都应该以这种方式来运行。

最少意外属性：“熟悉HTML，CSS和JavaScript，当然也要理解会发生上下文自动转义的开发人员（或代码审阅者），应该有能力查看`{{.}}`并正确推断`审查（sanitization）`发生了什么。”



# [go-toml](https://github.com/pelletier/go-toml)


## 功能

`Go-toml`提供以下功能，用于使用从TOML文档解析的数据：

- 从`文件`和`字符串数据`加载TOML文档；
- 使用`Tree`轻松浏览TOML结构；
- 从数据结构`打包（Marshal）`成数据流，以及将数据流`解包（Unmarshal）`成数据结构；
- 所有被解析元素的`行`和`列`位置数据；
- [支持类似于`JSON-Path`的查询；](https://github.com/pelletier/go-toml/tree/master/query)
- 语法报错包含`行号`和`列号`；


## 使用示例

[**读取TOML文件** -- 示例代码](load/)

[**解包（Unmarshal）** -- 示例代码](unmarshal/)

[**使用查询** -- 示例代码](query/)


## [GoDoc 文档](http://godoc.org/github.com/pelletier/go-toml)

## 工具

`Go-toml`提供了两个方便的命令行工具：

**tomll** -- 读取TOML文件，并导出。

```sh

go install github.com/pelletier/go-toml/cmd/tomll
tomll --help

```

**tomljson** -- 读取TOML文件，并以JSON格式输出。

```sh

go install github.com/pelletier/go-toml/cmd/tomljson
tomljson --help

```

**jsontoml** -- 读取JSON文件，并以TOML格式输出。

```sh

go install github.com/pelletier/go-toml/cmd/jsontoml
jsontoml --help

```

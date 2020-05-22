

# [`github.com/gorilla/sessions`包](https://github.com/gorilla/sessions)

`gorilla/sessions`包为`自定义session后端`提供了`cookie和文件系统session`以及`基础结构`。

主要功能有：

- 简单的API：将其用作设置`签名cookie`（以及`加密cookie`，可选）的简便方法。
- `内置的后端`可将session存储在cookie或文件系统中。
- 闪存消息（flash message）：session值，可以一直持续到被读取的。
- 切换`session持久性`（又称为`remember me`）和设置其他属性的便捷方法。
- `循环身份验证和加密密钥的机制`。
- `每个请求有多个session`，甚至使用不同的后端也是如此。
- `自定义session后端的接口和基础结构`：可以使用通用API`检索`并`批量保存`来自不同存储区的session。

在[Gorilla网站](https://www.gorillatoolkit.org/pkg/sessions)可获取更多信息。


## 代码示例

- [nutshell](nutshell/)
- []()
- []()


# [gin: github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)


## Gin基础知识


### 一、安装Gin及快速开始

[快速开始 -- 示例代码](quickstart/)


### 二、请求路由

[**1、设置多种请求类型** -- 示例代码](route/methods/)

[**2、绑定静态文件** -- 示例代码](route/statics/)

意味着，可以将Gin当作`静态文件服务器`使用。

[**3、参数作为URL** -- 示例代码](route/url2params/)

也可以说成，将URL作为参数；这种方式多用于`RESTful`请求。

[**4、泛绑定** -- 示例代码](route/generic/)

所有请求都定向到一个资源中。

[**5、路由分组 -- 示例代码](route/group/)


### 三、`获取请求参数`和`验证请求参数`

#### 获取请求参数

[**1、获取`GET`请求参数** -- 示例代码](request/data_get/)

[**2、获取`POST`请求参数** -- 示例代码](request/data_post_body/)

[**3、获取`Body`内容** -- 示例代码](request/data_post_body/)

[**4、获取`url`上的请求参数** -- 示例代码](request/data_urlparams/)

[**4、获取参数来`绑定结构体`** -- 示例代码](request/data2struct/)

#### 验证请求参数

[Model binding and validation](https://github.com/gin-gonic/gin#model-binding-and-validation)

[**1、结构体绑定验证器** -- 示例代码](request/data_validate/binding/)

[**2、自定义验证器** -- 示例代码](request/data_validate/custom/)

[**3、升级验证-支持多语言错误信息** -- 示例代码](request/data_validate/multilang/)


### 四、中间件

#### Gin自带的中间件

[**1、Gin的Logger中间件**](middleware/logger/README.md)
    - [**自定义`格式器函数`** -- 示例代码](middleware/logger/custom/formatter)
    - [**自定义日志颜色输出** -- 示例代码](middleware/logger/custom/color)
    - [**自定义日志输出流** -- 示例代码](middleware/logger/custom/writer)
    - [**自定义Logger中间件** -- 示例代码](middleware/logger/custom/logger)

[**2、服务器宕机恢复**](middleware/recovery/README.md)


#### 自定义`IP白名单`中间件


#### 其他中间件

[**Cookie与Session验证** -- 示例代码](middleware/session/)


### 五、其他补充


#### 优雅关停


#### 模板渲染


#### 自动证书


#### 使用[`jsoniter`](https://github.com/json-iterator/go)

Gin的JSON包默认使用标准包`encoding/json`，但是你可以改用`jsoniter`包。

```sh

go build -tags=jsoniter .

```


#### [更多API示例](https://github.com/gin-gonic/gin#api-examples)



## Gin应用示例

1. [使用`AsciiJSON`生成纯ASCII编码的JSON对象](expls/bind/renders/json/asciijson)
2. [自定义结构体到表单数据请求的绑定](expls/bind/model/bfrwcs)
3. [HTML中的复选框的绑定](expls/bind/renders/html/bhc)
4. [查询字符串或post数据的绑定](expls/bind/model/bqsopd)
5. [Uri的绑定（路由中变量的使用）](expls/bind/path/banduri)
6. [绑定多个模板的单个二进制文件](expls/bind/renders/html/basbwt)
7. [控制日志的色彩输出](expls/middleware/logger/cloc)
8. [自定义HTTP配置](expls/engine/chc)
9. [自定义日志文件](expls/middleware/logger/clf)
10. [自定义中间件](expls/middleware/logger/cm)
11. [自定义校验器](expls/middleware/validator/cv)
12. [为路由日志定义格式](expls/middleware/logger/dfftlor)
13. [一个中间件内的多协程](expls/middleware/giam)
14. [优雅地重启或停止](expls/engine/graceful/gros)
15. [路由组](expls/bind/path/gr)
16. [如何打印日志文件](expls/middleware/logger/htwlf)
17. [渲染HTML](expls/bind/renders/html/hr)
18. [HTTP2服务推送](expls/static/hsp)
19. [JSONP](expls/bind/renders/jsonp)
20. [查询字符串或post格式参数中的Map](expls/bind/queryandform/maqopp)
21. [model的绑定和验证](expls/bind/model/mbav)
22. [Multipart/Urlencoded binding](expls/bind/model/mub)
23. [Multipart/Urlencoded form](expls/bind/model/muf)
24. [加载多模板文件](expls/bind/renders/html/m)
25. [只绑定查询字符串](expls/bind/model/obqs)
26. [路径中的参数（路由中变量的使用）](expls/bind/path/pip)
27. [`PureJSON`生成JSON字面量](expls/bind/renders/json/purejson)
28. [查询（query）和post表单](expls/bind/queryandform/qapf)
29. [查询（query）的字符串参数](expls/bind/queryandform/qsp)
30. [重定向](expls/bind/redirect)
31. [运行多服务](expls/engine/rms)
32. [使用`SecureJSON`阻止json劫持](expls/bind/renders/json/securejson)
33. [提供来自reader（输入流）的数据服务](expls/bind/renders/sdfr)
34. [提供静态文件服务](expls/static/ssf)
35. [设置和获取一个cookie](expls/bind/cookie/sagac)
36. [支持`LetsEncrypt HTTPS`](expls/engine/sle)
37. [尝试绑定body到不同的结构体](expls/bind/model/ttbbids)
38. [上传文件-多文件](expls/bind/uploadfile/mf)
39. [上传文件-单个文件](expls/bind/uploadfile/sf)
40. [使用`BasicAuth`中间件](expls/middleware/basicauth)
41. [使用HTTP方法](expls/bind/httpmethod)
42. [使用中间件](expls/middleware/README.md)
43. [`gin.New()`创建默认没有中间件的框架对象](expls/engine/README.md)
44. [XML/JSON/YAML/ProtoBuf渲染](expls/bind/renders/quickstart)
45. [如何编写Gin的测试案例](expls/test/README.md)

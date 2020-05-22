

# [spf13/viper](https://github.com/spf13/viper)


## What is Viper?

`Viper`是适用于`Go语言应用程序`（包括`12-Factor`应用程序）的`完整配置解决方案`。它旨在在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持：

- 设置默认值；
- 从`JSON`，`TOML`，`YAML`，`HCL`，`envfile`和`Java属性配置文件`中读取；
- `实时监视`和`重新读取`配置文件（可选）；
- 从`环境变量`中读取；
- 从远程配置系统（`etcd`或`Consul`）中读取，并监视更改；
- 从`命令行标识`读取；
- 从`缓冲区`读取；
- 设置明确的值；

可以将`Viper`视为能满足您所有应用程序配置需求的`注册表`。


## Why Viper?

在构建现代应用程序时，您不想担心配置文件格式；您想专注于构建出色的软件。`Viper`在这里为您提供帮助。

`Viper`为您执行以下操作：

1、`查找`，`加载`和`解包`格式为`JSON`，`TOML`，`YAML`，`HCL`，`INI`，`envfile`或`Java属性`的配置文件。
2、`提供一种机制`来为您的不同配置选项`设置默认值`。
3、`提供一种机制`来给通过`命令行标识`指明的选项`设置重写值`。
4、`提供别名系统`以轻松重命名参数而不会破坏现有代码。
5、易于区分用户何时提供了与默认值相同的命令行或配置文件。

`Viper`使用以下优先顺序。每个项目优先于其下面的项目：

- 显式调用`Set`
- 命令行标识（flag）
- 环境变量（env）
- 配置（config）
- 键/值存储（key/value store）
- 默认

`重要提示`：目前，`Viper`配置的`键`是`不区分大小写`的。关于使该选项为可选的正在进行的讨论。


## Putting Values into Viper （将`值`放入`Viper`）


### Establishing Defaults（创建`默认值`）

一个好的`配置系统`将`支持默认值`。`键`可以没有`默认值`，但在没有使用`配置文件`，`环境变量`，`远程配置`或`命令行标识`来设置`键`的情况下，则`默认值`是很有用的。

示例：

```go

viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

```


### Reading Config Files（读取配置文件）

`Viper`所需的配置最少，因此它知道在哪里可以找到配置文件。`Viper`支持`JSON`，`TOML`，`YAML`，`HCL`，`INI`，`envfile`和`Java属性`文件。`Viper`可以搜索多个路径，但是当前单个`Viper`实例仅支持单个配置文件。`Viper`不会默认使用任何配置搜索路径，而将`默认决定权`留给应用程序。

这是有关如何使用`Viper`搜索和读取配置文件的示例。不需要任何特定路径，但是应该在期望有配置文件的地方至少提供一个路径。

```go

viper.SetConfigName("config")  // name of config file (without extension)
viper.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
viper.AddConfigPath(".")  // optionally look for config in the working directory
err := viper.ReadInConfig()  // Find and read the config file
if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
}

```

对于`找不到配置文件`的特定情况，您可以如下处理：

```go

if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // Config file not found; ignore error if desired
    } else {
        // Config file was found but another error was produced
    }
}

// Config file found and successfully parsed

```

**注意[自1.6版本开始]**：您还可以使用`不带扩展名的文件`，并以编程方式指定格式；对于那些位于用户家中的，没有任何扩展名的配置文件，如`.bashrc`。


### Writing Config Files（存储配置文件）

从配置文件`读取`很有用，但有时您想`存储`在运行时所做的所有修改。为此，可以使用一系列命令，每个命令都有自己的用途：

- `WriteConfig` -- 将当前的`viper配置`写入到`预定义的`路径（如果存在）。 如果没有预定义的路径，则报错。`将重写`当前的配置文件（如果存在）。
- `SafeWriteConfig` -- 将当前的`viper配置`写入`预定义的`路径。如果没有预定义的路径，则报错。`不会覆盖`当前的配置文件（如果存在）。
- `WriteConfigAs` -- 将当前的`viper配置`写入`给定的`文件路径。`将重写`给定的文件（如果存在）。
- `SafeWriteConfigAs` -- 将当前的`viper配置`写入`给定的`文件路径。不会覆盖当前配置文件（如果存在）。

根据经验，标记为`safe`的所有内容都不会重写任何文件，而是直接创建（如果不存在），而`默认行为是创建或截断`。

一个简短示例：

```go

viper.WriteConfig()  // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config")  // will errors since it has already been written
viper.SafeWriteConfigAs("/path/to/my/.other_config")

```


### Watching and re-reading config files（监视和重新读取配置文件）

`Viper`支持在运行时让您的应用程序`实时读取配置文件`的功能。

`需要重新启动服务器以使配置生效的日子已经一去不复返了`，`viper驱动的`应用程序`可以在运行时读取配置文件的更新，而不会错过任何消息`。

只要简单地告知`viper实例`要`WatchConfig`即可。可选地，您可以为`Viper`提供一个在每次发生更改时运行的函数。

**请确保，在调用`WatchConfig()`之前，您已添加了所有的配置路径。**

```go

viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config file changed:", e.Name)
})

```


### Reading Config from `io.Reader`（从`io.Reader`读取配置）

`Viper`预定义了许多`配置源`，例如`文件`，`环境变量`，`命令行标识`和`远程K/V存储`，但您不会受其约束；您还可以`实现自己所需的配置源`并将其提供给`viper`。

```go

viper.SetConfigType("yaml")  // or viper.SetConfigType("YAML")

// any approach to require this configuration into ypur program.
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
    jacket: leather
    trousers: denim
age: 35
eyes: brown
beard: true`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name")  // this would be "steve"


```


### Setting Overrides（设置重写）

`设置重写`，可能来自`命令行标识`，或来自您自己`应用程序的逻辑`：

```go

viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)

```


### Registering and Using Aliases（注册和使用`别名`）

`别名`，可以让单个`值`被多个`键`引用。

```go

viper.RegisterAlias("loud", "Verbose")

viper.Set("verbose", true)  // same result as next line
viper.Set("loud", true)     // same result as prior line

viper.GetBool("loud")     // true
viper.GetBool("verbose")  // true

```


### Working with Environment Variables（使用`环境变量`）

`Viper`完全支持环境变量。`12-Factor应用程序`对其开箱即用。有`五种方法`可以帮助使用`ENV`：

- `AutomaticEnv()`
- `BindEnv(string...) : error`
- `SetEnvPrefix(string)`
- `SetEnvKeyReplacer(string...) *strings.Replacer`
- `AllowEmptyEnv(bool)`

使用`ENV变量`时务必要知道，`Viper`会以`区分大小写`对待`ENV变量`。

`Viper`提供了`一种机制`来尝试确保`ENV变量`是唯一的。通过使用`SetEnvPrefix`，可以告诉`Viper`在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用此前缀。

`BindEnv`使用一个或两个参数。第一个参数是键名称，第二个是环境变量的名称。环境变量的名称`区分大小写`。如果未提供`ENV变量`名称，则`Viper`将自动假定`ENV变量`与以下格式匹配：`prefix + "_" + the key name in ALL CAPS（键名称全大写）`。当您提供`ENV变量`名称（第二个参数）时，它`不会`自动添加前缀。例如，如果第二个参数是`"id"`，`Viper`将查找`ENV变量"ID"`。

使用`ENV变量`时要知道的一件重要的事是，`每次访问该值时都会读取该值`。调用`BindEnv`时，`Viper`不会固定该值。

`AutomaticEnv`是强大的帮助程序，尤其是与`SetEnvPrefix`结合使用时。调用时，`Viper`会在每次发出`viper.Get`请求时检查环境变量。它将应用以下规则。它将检查环境变量：该环境变量的名称是否与大写的键名相匹配，是否使用了`EnvPrefix`设置的前缀（如果设置了的话）。

`SetEnvKeyReplacer`允许您使用`strings.Replacer`对象来对`Env键`进行一定程度上的重写。如果您想使用`-`或`Get()调用中的某些内容`，但希望环境变量使用`_`分隔符，那么此函数很有用。在`viper_test.go`中可以找到使用该函数的示例。

另外一种方式是，您可以将`EnvKeyReplacer`与`NewWithOptions`工厂函数一起使用。与`SetEnvKeyReplacer`不同的是，这种方式接收了`StringReplacer`接口，允许您编写`自定义字符串替换逻辑`。

默认情况下，`空的环境变量被认为是未设置的`，并将退回到下一个`配置源`。要将空的环境变量视为已设置，请使用`AllowEmptyEnv`方法。

**Env示例**

```go

SetEnvPrefix("spf")  // will be uppercased automatically
BindEnv("id")

os.Setenv("SPF_ID", "13")  // typically done outside of the app

id := Get("id")  // 13

```


### Working with Flags（使用`命令行标识`）

`Viper`具有绑定到`命令行标识`的能力。具体来说，`Viper`支持[Cobra库](https://github.com/spf13/cobra)中使用的`Pflags`。

与`BindEnv`一样，在调用绑定方法时该值还未被设置，而是在访问该值时才进行设置。这意味着您可以尽早绑定，甚至在`init()`函数中进行绑定。

对于单个`命令行标识`，`BindPFlag()`方法提供了此功能。

示例：

```go

serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))

```

您还可以绑定一个已存在的`pflags`集（`pflag.FlagSet`）；示例：

```go

pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)
i := viper.GetInt("flagname")  // retrieve values from viper instead of pflag

```

在`Viper`中使用[`pflag`](https://github.com/spf13/pflag/)并不妨碍使用其他使用了标准库中`flag`包的软件包。`pflag`软件包可以通过`导入`来处理为`flag`包定义的标识。这是通过调用`pflag`包提供的便捷函数（被称为`AddGoFlagSet()`）来完成的。

示例：

```go

package main

import (
    "flag"
    "github.com/spf13/pflag"
)

func main() {

    // using standard library "flag" package
    flag.Int("flagname", 1234, "help message for flagname")

    pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
    pflag.Parse()
    viper.BindPFlags(pflag.CommandLine)

    i := viper.GetInt("flagname")  // retrieve value from viper

    ...

}

```

#### Flag interfaces

如果您不适用`Pflags`的话，`Viper`提供了两个Go语言接口来绑定其他`flag系统`。

`FlagValue`代表`单个flag`；下面是一个关于如何实现该接口的非常简单的示例：

```go

type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }

```

一旦您的`flag`实现了该接口，您可以很简单地让`Viper`来绑定它：

```go

viper.BindFlagValue("my-flag-name", myFlag{})

```

`FlagValueSet`代表`一组flag`；下面是一个关于如何实现该接口的非常简单的示例：

```go

type myFlagSet struct {
    flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
    for _, flag := range flags {
        fn(flag)
    }
}

```

一旦您的`flag集`实现了该接口，您可以很简单地让`Viper`来绑定它：

```go

fSet := myFlagSet{
    flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)

```


### Remote Key/Value Store Support（支持远程键值对存储）

要启用`Viper`的`远程支持`，需要`空白导入`其`viper/remote`包：

```go

import _ "github.com/spf13/viper/remote"

```

`Viper`将读取从`键值对存储区`（例如`etcd`或`Consul`）中的路径检索到的配置字符串（如`JSON`，`TOML`，`YAML`，`HCL`或`envfile`）。这些值优先于默认值，但会被从`磁盘`，`命令行标志`或`环境变量`检索的配置值重写。

`Viper`使用[`crypt`](https://github.com/bketelsen/crypt)从`键值对存储区`中检索配置，这意味着您可以存储已加密的配置值，并且在拥有正确的`gpg密钥环`的情况下将其自动解密。`加密是可选的`。

您可以将`远程配置`与`本地配置`结合使用，也可以独立使用。

`crypt`具有一个命令行帮助程序，您可以使用它来将配置放入`键值对存储区`中。`crypt`在`http://127.0.0.1:4001`上默认指向`etcd`。

```sh

go get github.com/bketelsen/crypt/bin/crypt
crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json

```

请确保，您已设置了值：

```sh

crypt get -plaintext /config/hugo.json

```

关于如何设置加密值，或如何使用`Consul`，请查看`crypt`文档。


#### Remote Key/Value Store Example - Unencrypted（远程键值对存储区示例 - 非加密）

##### etcd

```go

viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.json")
viper.SetConfigType("json")  // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()

```

##### Consul

您需要使用包含您想要的配置的`JSON值`给`Consul键值对存储区`设置一个`键`。例如，使用如下`值`创建一个`Consul键值对存储区`的`键`--`MY_CONSUL_KEY`：

```json

{
    "port": 8080,
    "hostname": "myhostname.com"
}

```

```go

viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json")  // Need to explicitly set this to json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port"))  // 8080
fmt.Println(viper.Get("hostname"))  // myhostname.com

```

##### Firestore

```go

viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json")  // Config's format: "json", "toml", "yaml", "yml"
err := viper.ReadRemoteConfig()

```

当然， 您还可以使用`SecureRemoteProvider`。


#### Remote Key/Value Store Example - Encrypted（远程键值对存储区示例 - 加密）

```go

viper.AddSecureRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.json", "/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json")  // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properies", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()

```


#### Watching Changes in etcd - Unencrypted（监视`etcd`中的更改 - 非加密）

```go

// alternatively, you can create a new viper instance.
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml")  // because there is no fiel extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

// read from remote config the first time.
err := runtime_viper.ReadRemoteConfig()

// unmarshal config
runtime_viper.Unmarshal(&runtime_conf)

// open a goroutine to watch remote changes forever
go func() {
    for {
        time.Sleep(time.Second * 5)  // delay after each request

        // currently, only tested with etcd support
        err := runtime_viper.WatchRemoteConfig()
        if err != nil {
            log.Errorf("Unable to read remote config: %v", err)
            continue
        }

        // unmarshal new config into our runtime config struct, you can also use channel to implement a signal to notify the system of the changes
        runtime_viper.Unmarshal(&runtime_conf)
    }
}()

```


## Getting Values From Viper（从`Viper`获取值）

`Viper`里有一些根据`值类型`来获取`值`的途径。如下函数和方法：

- `Get(key string) : interface{}`
- `GetBool(key string) : bool`
- `GetFloat64(key string) : float64`
- `GetInt(key string) : int`
- `GetIntSlice(key string) : []int`
- `GetString(key string) : string`
- `GetStringMap(key string) : map[string]interface{}`
- `GetStringMapString(key string) : map[string]string`
- `GetStringSlice(key string) : []string`
- `GetTime(key string) : time.Time`
- `GetDuration(key string) : time.Duration`
- `IsSet(key string) : bool`
- `AllSettings() : map[string]interface{}`

需要知道的一个重点是，上述Get函数`在没有找到目标值时`会返回一个`零值`。想要检查一个给定`键`是否存在，提供了`IsSet()`方法。

示例：

```go

viper.GetString("logfile")  // case-insensitive Setting & Getting
if viper.GetBool("verbose") {
    fmt.Println("verbose enabled")
}

```


### Accessing nested keys（访问内嵌`键`）

访问方法还可以接收格式化的路径，来深入到内嵌的`键`。例如，如下的JSON文件被加载：

```json

{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

```

`Viper`可以通过传递用`.`分隔的`键路径`，来访问内嵌字段：

```go

GetString("datastore.metric.host")  // (return "127.0.0.1")

```

这里也遵守上面建立的`优先级规则`；搜索路径将遍历其余配置注册表，直到找到为止。

举例，在给定此配置文件的情况下，`datastore.metric.host`和`datastore.metric.port`均已定义（并且可以被覆盖）。如果另外在默认值中定义了`datastore.metric.protocol`，`Viper`也会找到它。

但是，如果`datastore.metric`被`最接近的值`重写了（通过`命令行标识`，`环境变量`，`Set()`方法等），那么`datastore.metric`的所有`子键`都变为未定义状态，它们被`预示（shadowed）`为较高的优先级配置级别。

最后，如果存在与`分隔键路径`匹配的键，则将返回其值。例如：

```go

{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

GetString("datastore.metric.host") // returns "0.0.0.0"

```


### Extract sub-tree（提取`子树`）

从`Viper`提取`子树`。

举例，`viper`描述如下：

```yaml

app:
    cache1:
        max-items: 100
        item-size: 64
    cache2:
        max-items: 200
        item-size: 80

```

执行之后：

```go

subv := viper.Sub("app.cache1")

```

`subv`描述如下：

```yaml

max-items: 100
item-size: 64

```

假设我们有：

```go

func NewCache(cfg *Viper) *Cache {...}

```

该函数可以基于格式类似`subv`的配置信息创建`缓存`。现在，可以轻松地分别创建这两个`缓存`，如下所示：

```go

cfg1 := viper.Sub("app.cache1")
cache1 := NewCache(cfg1)

cfg2 := viper.Sub("app.cache2")
cache2 := NewCache(cfg2)

```


### Unmarshaling（解包，或称为`反序列化`）

您还可以选择解包所有内容，或将指定值解包为结构体、map等等。

有两种方法来实现：
- `Unmarshal(rawVal interface{}) : error`  
- `UnmarshalKey(key string, rawVal interface{}) : error`

示例：

```go

type config struct {
    Port int
    Name string
    PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
    t.Fatalf("unable to decode into struct, %v", err)
}

```

如果您想要对其`键本身`包含`.`符号（`.`是默认的`键分隔符`）的配置进行解包，您不得不改变`分隔符`：

```go

v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
    "ingress": map[string]interface{}{
        "annotations": map[string]interface{}{
            "traefik.frontend.rule.type": "PathPrefix",
            "traefik.ingress.kubernetes.io/ssl-redirect": "true",
        },
    },
})

type config struct {
    Chart struct{
        Values map[string]interface{}
    }
}

var C config

v.Unmarshal(&C)

```

`Viper`还支持解包到`内嵌结构体`：

```go

/*
Example config:

module:
    enabled: true
    token: 89h3f98hbwf987h3f98wenf89ehf
*/
type config struct {
    Module struct {
        Enabled bool
        moduleConfig `mapstructure:",squash"`
    }
}

// moduleConfig could be in a module specific package
type moduleConfig struct {
    Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
    t.Fatalf("unable to decode into struct, %v", err)
}

```

`Viper`在内部使用[github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)来解包使用了默认的`mapstructure`标记的值。


### Marshalling to string（打包（或称为`序列化`）到字符串）

您可能需要将`viper`中保存的所有设置打包为字符串，而不是将它们写入文件。您可以将自己喜欢的`格式打包器`与`AllSettings()`返回的`配置`一起使用。

```go

import (
    yaml "gopkg.in/yaml.v2"
    // ...
)

func yamlStringSettings() string {
    c := viper.AllSettings()
    bs, err := yaml.Marshal(c)
    if err != nil {
        log.Fatalf("unable to marshal config to YAML: %v", err)
    }
    return string(bs)
}

```


## Viper or Vipers?

`Viper`开箱即可使用。不需要配置或初始化即可开始使用`Viper`。由于大多数应用程序都希望使用单个中央存储库进行配置，因此`viper`包就提供了此功能。`它类似于单例`。

在以上所有示例中，它们都`以其单例样式`方法演示了使用`viper`。


### Working with multiple vipers

您还可以创建许多不同的`viper`以用于您的应用程序。每个组件都有其自己独特的一组配置和值。每个工具都可以从不同的`配置文件`，`键值存储区`等读取数据。`viper`包支持的所有`函数`都可以镜像为一个`viper`上的方法。

示例：

```go

x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

// ...

```

当使用多个`viper`时，如何跟踪不同的`viper`要取决于用户。

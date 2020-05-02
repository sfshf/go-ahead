
[Pseudo-versions](https://golang.google.cn/cmd/go/#hdr-Pseudo_versions)


##### 说明：

`go.mod`文件和`go命令`通常使用语义版本作为描述模块版本的标准形式，这样可以对版本进行比较，以确定哪个版本应该比另一个版本更早或更晚。像`v1.2.3`这样的模块版本是通过在底层源存储库中标记修订而引入的。未标记的修订可以使用`pseudo-version（伪版本）`引用，比如`v0.0.0-yyyymmddhhmmss-abcdefabcdef`，其中时间是提交时间的UTC表示，最后的后缀是提交散列的前缀。时间部分确保可以比较两个伪版本来确定哪个版本是较后产生的，提交散列标识了底层的提交，而前缀(本例中为`v0.0.0`)是派生自提交图中在本次提交之前的最新标记的版本。

有三种伪版本形式:

`vX.0.0-yyyymmddhhmmss-abcdefabcdef`，当目标提交之前没有更早的使用适当主版本的版本提交时被采用。(这本来是唯一的格式，所以一些老的`go.mod`文件使用这个格式，即使提交是跟在标签后面的。)

`vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef`，当目标提交之前最近的版本提交是`vX.Y.Z-pre`时被采用。

`vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef`，当目标提交之前最近的版本提交是`vX.Y.Z`时被采用。

伪版本不需要手动输入：`go命令`将接受普通的提交散列并自动将其转换为伪版本(如果可用，则转换为带标记的版本)。这个转换是一个[`module query（模块查询）`](mq.md)的例子。

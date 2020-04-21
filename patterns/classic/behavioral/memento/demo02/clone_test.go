package clone

import "fmt"

func ExampleMementoByClone() {

    //创建初始状态的发起人对象
    originator := &Originator{
        state: "初始状态...",
    }
    fmt.Printf("初始originator状态为：%s\n", originator.state)

    //建立备份
    originator.CreateMemento()

    //改变originator的状态
    originator.state = "修改后的状态..."
    fmt.Printf("修改后originator状态为：%s\n", originator.state)

    //回复原有状态
    originator.RestoreMemento()
    fmt.Printf("恢复后originator状态为：%s\n", originator.state)

    //Output:
    //初始originator状态为：初始状态...
    //修改后originator状态为：修改后的状态...
    //恢复后originator状态为：初始状态...

}

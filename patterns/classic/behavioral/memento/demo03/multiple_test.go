package multiple

import (
    "testing"
    "fmt"
)

func TestMultipleMemento(t *testing.T) {

    util := NewBackUtils()

    originator1 := NewOriginator(123, 1.23, "初始状态")
    originator1.Show()
    memento1 := util.BackupProp(originator1)

    //改变originator1的字段值
    originator1.State1 = 1
    originator1.State2 = 0.5
    originator1.State3 = "修改第一次后的状态"
    originator1.Show()
    memento2 := util.BackupProp(originator1)

    //改变originator1的字段值
    originator1.State1 = 456
    originator1.State2 = 2.15
    originator1.State3 = "修改第二后的状态"
    originator1.Show()

    fmt.Println("恢复到第一次修改节点:")
    util.RestoreProp(originator1, memento2.Backupid())
    originator1.Show()

    fmt.Println("恢复到初始节点:")
    util.RestoreProp(originator1, memento1.Backupid())
    originator1.Show()

}

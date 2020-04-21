package clone

/*

创建型的设计模式中有一个原型模式，是通过复制的方式进行新的对象创建；
而行为型设计模式中的备忘录模式也可以通过复制的方式来达到备忘的效果。

但是，对象体量很大的情况下，clone方式非常消耗资源，
因此，clone方式的备忘录模式适用于较简单的场景。

*/

type Cloneable interface {
    Clone() interface {}
}

type Originator struct {
    backup *Originator
    state string
}

func (o *Originator) CreateMemento() {
    copy := *o
    o.backup = &copy
}

func (o *Originator) RestoreMemento() {
    o.state = o.backup.state
}

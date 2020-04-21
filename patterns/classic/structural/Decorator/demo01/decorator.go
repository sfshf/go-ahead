package decorator

/*

装饰模式

装饰模式使用对象组合的方式动态改变或增加对象行为。

Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。

使用匿名组合，在装饰器中不必显式定义转调原对象方法。

*/

type Component interface {
    Calc() int
}

func NewComponent() Component {
    return &ConcreteComponent{}
}

type ConcreteComponent struct {}

func (*ConcreteComponent) Calc() int {
    return 0
}

type MulDecorator struct {
    Component
    num int
}

func NewMulDecorator(c Component, num int) Component {
    return &MulDecorator{
        Component: c,
        num: num,
    }
}

func (md *MulDecorator) Calc() int {
    return md.Component.Calc() * md.num
}

type AddDecorator struct {
    Component
    num int
}

func NewAddDecorator(c Component, num int) Component {
    return &AddDecorator{
        Component: c,
        num: num,
    }
}

func (ad *AddDecorator) Calc() int {
    return ad.Component.Calc() + ad.num
}

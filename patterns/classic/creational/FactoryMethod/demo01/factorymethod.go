package factorymethod

/*
工厂方法模式

工厂方法模式使用子类的方式延迟生成对象到子类中实现；什么工厂就生产什么产品；

Go中不存在继承 所以使用匿名组合来实现。
*/

//Operator是被封装的实际类接口
type Operator interface {
    SetA(int)
    SetB(int)
    Result() int
}

//OperatorFactory是工厂接口
type OperatorFactory interface {
    Create() Operator
}

//------------------------------------------------------------------------------

//OperatorBase是Operator接口的实现基类，封装公用方法
type OperatorBase struct {
    a, b int
}

func (ob *OperatorBase) SetA(a int) {
    ob.a = a
}

func (ob *OperatorBase) SetB(b int) {
    ob.b = b
}

//------------------------------------------------------------------------------

//PlusOperator：Operator的实际加法实现
type PlusOperator struct {
    *OperatorBase
}

func (po *PlusOperator) Result() int {
    return po.a + po.b
}

//PlusOperatorFactory是PlusOperator的工厂类
type PlusOperatorFactory struct {}

func (pof *PlusOperatorFactory) Create() Operator {
    return &PlusOperator{
        OperatorBase: &OperatorBase{},
    }
}

//------------------------------------------------------------------------------

//MinusOperator：Operator的实际减法实现
type MinusOperator struct {
    *OperatorBase
}

func (mo *MinusOperator) Result() int {
    return mo.a - mo.b
}

//MinusOperatorFactory是MinusOperator的工厂类
type MinusOperatorFactory struct {}

func (mof *MinusOperatorFactory) Create() Operator {
    return &MinusOperator{
        OperatorBase: &OperatorBase{},
    }
}

package human

import "fmt"

type Human interface {
    GetColor() string
    Talk() string
}

type HumanFactory interface {
    Create(a ...interface{}) Human
}

type HumanBase struct {
    height int
    weight float64
    age int
    color string
}

func (h *HumanBase) GetColor() string {
    return h.color
}

func (h *HumanBase) Talk() string {
    return fmt.Sprintf("我是一个身高=%d厘米，体重=%.2f斤，年龄=%d岁的%s人！",
        h.height, h.weight, h.age, h.GetColor())
}

//黄种人
type YellowHuman struct {
    *HumanBase
}

//白种人
type WhiteHuman struct {
    *HumanBase
}

//黑种人
type BlackHuman struct {
    *HumanBase
}

type YellowHumanFactory struct {}

func (yhf *YellowHumanFactory) Create(a ...interface{}) Human {
    return &YellowHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
    }
}

type WhiteHumanFactory struct {}

func (whf *WhiteHumanFactory) Create(a ...interface{}) Human {
    return &WhiteHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
    }
}

type BlackHumanFactory struct {}

func (bhf *BlackHumanFactory) Create(a ...interface{}) Human {
    return &BlackHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
    }
}

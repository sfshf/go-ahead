package human

import "fmt"

type Human interface {
    GetColor() string
    Talk() string
}

type Female interface {
    GetSex() string
}

type Male interface {
    GetSex() string
}

type HumanFactory interface {
    CreateAFemale(a ...interface{}) Human
    CreateAMale(a ...interface{}) Human
}

//-----------------------------------------------------------------------

type HumanBase struct {
    height int
    weight float64
    age int
    color string
}

func (hb *HumanBase) GetColor() string {
    return hb.color
}

func (hb *HumanBase) Talk() string {
    return fmt.Sprintf("我是一个身高=%d厘米，体重=%.2f斤，年龄=%d岁的%s人！",
        hb.height, hb.weight, hb.age, hb.color)
}

type YellowHuman struct {
    *HumanBase
    sex string
}

func (yh *YellowHuman) GetSex() string {
    return fmt.Sprintf("我是一个%v人！", yh.sex)
}

type WhiteHuman struct {
    *HumanBase
    sex string
}

func (wh *WhiteHuman) GetSex() string {
    return fmt.Sprintf("我是一个%v人！", wh.sex)
}

type BlackHuman struct {
    *HumanBase
    sex string
}

func (bh *BlackHuman) GetSex() string {
    return fmt.Sprintf("我是一个%v人！", bh.sex)
}

//-----------------------------------------------------------------------

type YellowHumanFactory struct {}

func (yhf *YellowHumanFactory) CreateAFemale(a ...interface{}) Human {
    return &YellowHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "female",
    }
}

func (yhf *YellowHumanFactory) CreateAMale(a ...interface{}) Human {
    return &YellowHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "male",
    }
}

type WhiteHumanFactory struct {}

func (whf *WhiteHumanFactory) CreateAFemale(a ...interface{}) Human {
    return &WhiteHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "female",
    }
}

func (whf *WhiteHumanFactory) CreateAMale(a ...interface{}) Human {
    return &WhiteHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "male",
    }
}

type BlackHumanFactory struct {}

func (bhf *BlackHumanFactory) CreateAFemale(a ...interface{}) Human {
    return &BlackHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "female",
    }
}

func (bhf *BlackHumanFactory) CreateAMale(a ...interface{}) Human {
    return &BlackHuman{
        HumanBase: &HumanBase{
            height: a[0].(int),
            weight: a[1].(float64),
            age: a[2].(int),
            color: a[3].(string),
        },
        sex: "male",
    }
}

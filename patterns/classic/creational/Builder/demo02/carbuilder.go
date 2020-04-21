package carbuilder

import "fmt"

// Model

type implement interface {
    start()
    stop()
    alarm()
    engineBoom()
}

type CarModel struct {
    implement
    sequence []string
}

func (cm *CarModel) Run() {
    for _, s := range cm.sequence {
        switch s {
        case "start" :
            cm.start()
        case "stop" :
            cm.stop()
        case "alarm" :
            cm.alarm()
        case "engineBoom" :
            cm.engineBoom()
        default :
            fmt.Println("ERROR: Invalid function string! --", s)
        }
    }
}

type BenzModel struct {
    *CarModel
}

func (benz *BenzModel) start() {
    fmt.Println("benz is starting ...")
}

func (benz *BenzModel) stop() {
    fmt.Println("benz is stopping ...")
}

func (benz *BenzModel) alarm() {
    fmt.Println("benz is biu biu biu ...")
}

func (benz *BenzModel) engineBoom() {
    fmt.Println("benz is booming ...")
}

type BMWModel struct {
    *CarModel
}

func (bmw *BMWModel) start() {
    fmt.Println("bmw is starting ...")
}

func (bmw *BMWModel) stop() {
    fmt.Println("bmw is stopping ...")
}

func (bmw *BMWModel) alarm() {
    fmt.Println("bmw is biu biu biu ...")
}

func (bmw *BMWModel) engineBoom() {
    fmt.Println("bmw is booming ...")
}

//Builder:

type CarBuilder interface {
    GetCarModel(seq []string) *CarModel
}

type BenzBuilder struct {}

func (bcb *BenzBuilder) GetCarModel(seq []string) *CarModel {
    model := &BenzModel{}
    model.CarModel = &CarModel{
                        implement: model,
                        sequence: seq,
                    }
    return model.CarModel
}

type BWMBuilder struct {}

func (bwmb *BWMBuilder) GetCarModel(seq []string) *CarModel {
    model := &BMWModel{}
    model.CarModel = &CarModel{
                        implement: model,
                        sequence: seq,
                    }
    return model.CarModel
}

//这里的Director没能体现效用。
type Director struct {
    CarBuilder
    // var builders []CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
    return &Director{
        CarBuilder: builder,
    }
}

package hummermodel

import "fmt"

type HummerModel interface {

    Run()

}

type implement interface {
    start()
    stop()
    alarm()
    engineBoom()
}

type template struct {
    implement
}

func (t *template) Run() {
    t.implement.start()
    t.implement.engineBoom()
    t.implement.alarm()
    t.implement.stop()
}

func newTemplate(impl implement) *template {
    return &template{
        implement: impl,
    }
}

type HummerH1Model struct {
    *template
}

func (h1 *HummerH1Model) start() {
    fmt.Println("h1 is starting ...")
}

func (h1 *HummerH1Model) engineBoom() {
    fmt.Println("h1's engine boom boom boom ...")
}

func (h1 *HummerH1Model) alarm() {
    fmt.Println("h1's alarm bell ringing ...")
}

func (h1 *HummerH1Model) stop() {
    fmt.Println("h1 is stopping ...")
}

func NewHummerH1Model() HummerModel {
    h1 := &HummerH1Model{}
    h1.template = newTemplate(h1)
    return h1
}

type HummerH2Model struct {
    *template
}

func (h2 *HummerH2Model) start() {
    fmt.Println("h2 is starting ...")
}

func (h2 *HummerH2Model) engineBoom() {
    fmt.Println("h2's engine boom boom boom ...")
}

func (h2 *HummerH2Model) alarm() {
    fmt.Println("h2's alarm bell ringing ...")
}

func (h2 *HummerH2Model) stop() {
    fmt.Println("h2 is stopping ...")
}

func NewHummerH2Model() HummerModel {
    h2 := &HummerH2Model{}
    h2.template = newTemplate(h2)
    return h2
}

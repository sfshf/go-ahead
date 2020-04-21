package spy

import (
    "fmt"
    "container/list"
)

type Observable interface {
    AddObserver(observer Observer)
    DeleteObserver(observer Observer)
    NotifyAll()
}

type Subject interface {
    Observable
    HaveBreakfast()
    HaveFun()
}

type HanFeiZi struct {
    observers *list.List
    context string
}

func NewSubject() Subject {
    return &HanFeiZi{
        observers: list.New(),
    }
}

func (hfz *HanFeiZi) AddObserver(observer Observer) {
    hfz.observers.PushBack(observer)
}

func (hfz *HanFeiZi) DeleteObserver(observer Observer) {
    for e := hfz.observers.Front(); e != nil; e = e.Next() {
        if e.Value == observer {
            hfz.observers.Remove(e)
        }
    }
}

func (hfz *HanFeiZi) NotifyAll() {
    for e := hfz.observers.Front(); e != nil; e = e.Next() {
        switch ins := e.Value.(type) {
        case *Spy:
            ins.UpdateState(hfz)
        }
    }
}

func (hfz *HanFeiZi) HaveBreakfast() {
    fmt.Println("韩非子开始吃饭！")
    hfz.context = "韩非子正在吃饭..."
    hfz.NotifyAll()
    fmt.Println("韩非子结束吃饭！")
    hfz.context = ""
}

func (hfz *HanFeiZi) HaveFun() {
    fmt.Println("韩非子开始娱乐！")
    hfz.context = "韩非子正在娱乐..."
    hfz.NotifyAll()
    fmt.Println("韩非子结束娱乐！")
    hfz.context = ""
}

type Observer interface {
    UpdateState(subject Subject)
}

type Spy struct {
    name string
    context string
    host string
}

func NewSpy(name string, host string) *Spy {
    return &Spy{
        name: name,
        host: host,
    }
}

func (s *Spy) UpdateState(subject Subject) {

    switch ins := subject.(type) {
    case *HanFeiZi:
        s.context = ins.context
        s.SendMessage()
    }

}

func (s *Spy) SendMessage() {
    fmt.Printf("间谍<%s>将【%s】消息传递给了<%s>。 \n", s.name, s.context, s.host)
}

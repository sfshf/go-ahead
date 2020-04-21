package handlers

import (
    "testing"
    "math/rand"
)

func NewInCharged(handler IHandler) *InCharged {
    return &InCharged{
        IHandler: handler,
    }
}

func TestHandlers(t *testing.T) {
    women := make([]*Woman, 0)
    for i := 0; i < 5; i ++ {
        women = append(women, &Woman{
            level: rand.Intn(3) + 1,
        })
    }
    for _, woman := range women {
        switch woman.level {
        case 1:
            woman.request = "女儿我想出去逛街"
        case 2:
            woman.request = "妻子我想出去逛街"
        case 3:
            woman.request = "老娘我想出去逛街"
        }
    }

    incharged1 := NewInCharged(&FatherHandler{})
    incharged2 := NewInCharged(&HusbandHandler{})
    incharged3 := NewInCharged(&SonHandler{})

    incharged1.SetNext(incharged2)
    incharged2.SetNext(incharged3)

    var handlers IHandler = incharged1

    for _, woman := range women {
        handlers.HandleRequest(woman)
    }

}

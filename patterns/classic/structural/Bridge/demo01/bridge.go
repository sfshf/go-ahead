package bridge

import "fmt"

/*

桥接模式

桥接模式分离抽象部分和实现部分。使得两部分独立扩展。

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。

桥接模式使抽象部分和实现部分分离，可以独立变化。

*/

type AbstractMessage interface {
    SendMessage(text, to string)
}

type CommonMessage struct {
    method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
    return &CommonMessage{
        method: method,
    }
}

func (cm *CommonMessage) SendMessage(text, to string) {
    cm.method.Send(text, to)
}

type UrgencyMessage struct {
    method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
    return &UrgencyMessage{
        method: method,
    }
}

func (um *UrgencyMessage) SendMessage(text, to string) {
    um.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

type MessageImplementer interface {
    Send(text, to string)
}

type MessageSMS struct {}

func ViaSMS() *MessageSMS {
    return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
    fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct {}

func ViaEmail() *MessageEmail {
    return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
    fmt.Printf("send %s to %s via Email", text, to)
}

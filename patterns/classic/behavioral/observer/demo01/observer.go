package observer

import "fmt"

/*

观察者模式

观察者模式用于触发联动。

一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。

*/

type Observable interface {
    Attach(observer Observer)
    Notify()
}

type Observer interface {
    Update(*Subject)
}

type Subject struct {
    observers []Observer
    context string
}

func (s *Subject) UpdateContext(context string) {
    s.context = context
    s.Notify()
}

func (s *Subject) Attach(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Subject) Notify() {
    for _, o := range s.observers {
        o.Update(s)
    }
}

type Reader struct {
    name string
}

func NewReader(name string) *Reader {
    return &Reader{
        name: name,
    }
}

func (r *Reader) Update(s *Subject) {
    fmt.Printf("%s receive %s\n", r.name, s.context)
}

package observer

import "testing"

func TestObserver(t *testing.T) {

    observer1 := NewReader("reader1")
    observer2 := NewReader("reader2")
    observer3 := NewReader("reader3")

    s := &Subject{}
    s.Attach(observer1)
    s.Attach(observer2)
    s.Attach(observer3)

    s.UpdateContext("今天要抓米国间谍")

}

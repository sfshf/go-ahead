package simplefactory

import "testing"

func TestType1(t *testing.T) {
    api := NewAPI(1)
    s := api.Say("Tom")
    if s != "Hi, Tom" {
        t.Fatal("Type1 test fail!")
    }
}

func TestType2(t *testing.T) {
    api := NewAPI(2)
    s := api.Say("Jerry")
    if s != "Hello, Jerry" {
        t.Fatal("Type2 test fail!")
    }
}

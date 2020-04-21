package adapter

import "fmt"

func ExampleAdapter() {

    adaptee := NewAdaptee()
    target := NewTarget(adaptee)
    res := target.Request()
    fmt.Println(res)
    //Output:
    //adaptee method

}

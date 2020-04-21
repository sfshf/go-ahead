package decorator

import "fmt"

func ExampleDecorator() {

    component := NewComponent()
    component = NewMulDecorator(component, 10)
    component = NewAddDecorator(component, 8)
    res := component.Calc()
    fmt.Printf("result: %d\n", res)
    //Output:
    //result: 8

}

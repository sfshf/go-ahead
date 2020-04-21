package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
    op := factory.Create()
    op.SetA(a)
    op.SetB(b)
    return op.Result()
}

func TestOperator(t *testing.T) {
    var (
        factory OperatorFactory
    )
    factory = &PlusOperatorFactory{}
    if compute(factory, 13, 12) != 25 {
        t.Fatal("error with factory method pattern! -- plus operator factory!")
    }
    factory = &MinusOperatorFactory{}
    if compute(factory, 13, 12) != 1 {
        t.Fatal("error with factory method pattern! -- minus operator factory!")
    }
}

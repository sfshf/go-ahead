package ring

type Element struct {
    Value interface{}
}

type Ring struct {
    elems []Element
    front int
    rear int
}

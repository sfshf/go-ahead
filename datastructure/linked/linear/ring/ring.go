package ring

type Element struct {
    Value interface{}
    next *Element
}

type Ring struct {

    root Element
    len int

}

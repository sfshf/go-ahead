package stack


type Element struct {
    Value interface{}
}

type Stack struct {
    elems []Element
    top int
}

func New() *Stack {

}

func (s *Stack) Init() *Stack {

}

// 入栈运算
func (s *Stack) Push(v interface{}) (*Element, error) {

    e := Element {
        Value: v
    }
    s.elems = append(s.elems, e)


}

// 退栈运算
func (s *Stack) Pop() (*Element, error) {

}

// 读取栈顶元素
func (s *Stack) GetTop() (*Element) {

}

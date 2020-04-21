package corporation

import "fmt"

type ICorp interface {
    PrintInfo()
    AddSubordinate(ICorp)
    SetSuperior(ICorp)
}

func NewICorp(kind int, name string, position string, salary int) ICorp {
    switch kind {
    case BRANCH:
        return &Branch{
            corp: &corp{
                name: name,
                position: position,
                salary: salary,
            },
            subordinates: make([]ICorp, 0),
        }
    case LEAF:
        return &Leaf{
            corp: &corp{
                name: name,
                position: position,
                salary: salary,
            },
        }
    default:
        return nil
    }
}

const (
    BRANCH = iota
    LEAF
)

type corp struct {
    name string
    position string
    salary int
    superior ICorp
}

func (c *corp) PrintInfo() {
    fmt.Printf("姓名：%s，职位：%s，薪水：%d，", c.name, c.position, c.salary)
    if obj, ok := c.superior.(*Branch); ok {
        fmt.Printf("直接上司是%s。\n", obj.name)
    } else {
        fmt.Printf("直接上司是%s。\n", "nil")
    }
}

func (c *corp) AddSubordinate(ic ICorp) {}

func (c *corp) SetSuperior(ic ICorp) {
    c.superior = ic
}

type Branch struct {
    *corp
    subordinates []ICorp
}

func (b *Branch) AddSubordinate(ic ICorp) {
    ic.SetSuperior(b)
    b.subordinates = append(b.subordinates, ic)
}

func (b *Branch) PrintInfo() {
    b.corp.PrintInfo()
    if len(b.subordinates) > 0 {
        for _, obj := range b.subordinates {
            obj.PrintInfo()
        }
    }

}

type Leaf struct {
    *corp
}

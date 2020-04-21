package buysalestock

import "fmt"

type IMediator interface {
    Execute(i interface{})
}

type Mediator struct {
    purchase *Purchase
    sale *Sale
    stock *Stock
}

func (m *Mediator) Execute(i interface{}) {
    switch ins := i.(type) {
    case *Purchase:
        m.sale.num_computer = ins.num_computer
        m.stock.num_computer = ins.num_computer
    case *Sale:
        m.purchase.num_computer = ins.num_computer
        m.stock.num_computer = ins.num_computer
    case *Stock:
        fmt.Printf("库存量还有%v件！\n", ins.num_computer)
    }
}

type Purchase struct {
    IMediator
    num_computer int
}

func (p *Purchase) BuyIBMcomputer(num int) {
    p.num_computer += num
    p.IMediator.Execute(p)
}

type Sale struct {
    IMediator
    num_computer int
}

func (s *Sale) SellIBMComputer(num int) {
    s.num_computer -= num
    s.IMediator.Execute(s)
}

type Stock struct {
    IMediator
    num_computer int
}

func (s *Stock) GetStockNumber() {
    s.IMediator.Execute(s)
}

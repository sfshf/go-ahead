package buysalestock

import "testing"

func TestMediator(t *testing.T) {

    mediator := &Mediator{}
    purchase := &Purchase{
        IMediator: mediator,
    }
    sale := &Sale{
        IMediator: mediator,
    }
    stock := &Stock{
        IMediator: mediator,
    }
    mediator.purchase = purchase
    mediator.sale = sale
    mediator.stock = stock

    purchase.BuyIBMcomputer(100)

    if sale.num_computer != 100 && stock.num_computer != 100 {
        t.Fatal("进货后，数据同步失败！")
    }

    sale.SellIBMComputer(27)

    if purchase.num_computer != 73 && stock.num_computer != 73 {
        t.Fatal("销售后，数据同步失败！")
    }

    stock.GetStockNumber()

}

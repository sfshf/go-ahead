package corp

import "fmt"

//将公司本身与主要产品进行解耦
type ICorp interface {
    MakeMoney()
}

type HouseCorp struct {
    IProduct
}

func NewHouseCorp(product IProduct) *HouseCorp {
    return &HouseCorp{
        IProduct: product,
    }
}

func (hc *HouseCorp) MakeMoney() {
    hc.IProduct.BeProducted()
    hc.IProduct.BeSold()
    fmt.Println("卖出去的房子有售后服务")
}

type ShanzhaiCorp struct {
    IProduct
}

func NewShanzhaiCorp(product IProduct) *ShanzhaiCorp {
    return &ShanzhaiCorp{
        IProduct: product,
    }
}

func (sc *ShanzhaiCorp) MakeMoney() {
    sc.IProduct.BeProducted()
    sc.IProduct.BeSold()
    fmt.Println("卖出去的山寨货没有售后服务")
}

type IProduct interface {
    BeProducted()
    BeSold()
}

type CommodityHouse struct {}

func NewCommodityHouse() *CommodityHouse {
    return &CommodityHouse{}
}

func (ch *CommodityHouse) BeProducted() {
    fmt.Println("商品房已经备好房源")
}

func (ch *CommodityHouse) BeSold() {
    fmt.Println("商品房已经卖出")
}

type ShanzhaiProduct struct {}

func NewShanzhaiProduct() *ShanzhaiProduct {
    return &ShanzhaiProduct{}
}

func (sp *ShanzhaiProduct) BeProducted() {
    fmt.Println("山寨货已经备好货源")
}

func (sp *ShanzhaiProduct) BeSold() {
    fmt.Println("山寨货已经卖出")
}

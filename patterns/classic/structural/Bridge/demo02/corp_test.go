package corp

func ExampleHouseProduct() {

    var corp ICorp = NewHouseCorp(NewCommodityHouse())
    corp.MakeMoney()
    //Output:
    //商品房已经备好房源
    //商品房已经卖出
    //卖出去的房子有售后服务

}

func ExampleShanzhaiProduct() {

    var corp ICorp = NewShanzhaiCorp(NewShanzhaiProduct())
    corp.MakeMoney()
    //Output:
    //山寨货已经备好货源
    //山寨货已经卖出
    //卖出去的山寨货没有售后服务

}

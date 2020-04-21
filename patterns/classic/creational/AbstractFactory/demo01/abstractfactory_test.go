package abstractfactory

import "testing"

func getMainAndDetail(factory DAOFactory) (s1, s2 string) {
    return factory.CreateOrderMainDAO().SaveOrderMain(),
    factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestRdbFactory(t *testing.T) {
    var factory DAOFactory
    factory = &RDBDAOFactory{}
    t.Log(getMainAndDetail(factory))
    //Output:
    // rdb main save!
    // rdb detail save!
}

func TestXmlFactory(t *testing.T) {
    var factory DAOFactory
    factory = &XMLDAOFactory{}
    t.Log(getMainAndDetail(factory))
    //Output:
    // xml main save!
    // xml detail save!
}

package abstractfactory

import "fmt"

/*
抽象工厂模式

抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。

如果抽象工厂退化成生成的对象无关联则成为工厂方法模式。

比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息
和订单详情信息。 如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的
类就能替换使用不同的存储方式了。
*/

//OrderMainDao为订单主记录
type OrderMainDAO interface {
    SaveOrderMain() string
}

//OrderDetailDao为订单详情记录
type OrderDetailDAO interface {
    SaveOrderDetail() string
}

//DAOFactory：DAO抽象模式工厂接口
type DAOFactory interface {
    CreateOrderMainDAO() OrderMainDAO
    CreateOrderDetailDAO() OrderDetailDAO
}

//-----------------------------------------------------------------------

//RDBMainDAO：为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {}

func (*RDBMainDAO) SaveOrderMain() string {
    return fmt.Sprintln("rdb main save!")
}

//RDBDetailDAO：为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct {}

func (*RDBDetailDAO) SaveOrderDetail() string {
    return fmt.Sprintln("rdb detail save!")
}

//RDBDAOFactory：是RDB抽象工厂实现
type RDBDAOFactory struct {}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
    return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
    return &RDBDetailDAO{}
}

//-----------------------------------------------------------------------

//XMLMainDAO：XML存储
type XMLMainDAO struct {}

//SaveOrderMain：...
func (*XMLMainDAO) SaveOrderMain() string {
    return fmt.Sprintln("xml main save!")
}

//XMLDetailDAO：XML存储
type XMLDetailDAO struct {}

//SaveOrderDetail：...
func (*XMLDetailDAO) SaveOrderDetail() string {
    return fmt.Sprintln("xml detail save!")
}

//XMLDAOFactory：是RDB抽象工厂实现
type XMLDAOFactory struct {}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
    return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
    return &XMLDetailDAO{}
}

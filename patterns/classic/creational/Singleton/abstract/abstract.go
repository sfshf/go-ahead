package singleton

/*
抽象工厂单例模式 -- 只生产唯一对象的抽象工厂模式
*/
type MonkeyKing interface {
    FanJingDou() string
    QiShiErBian() string
}

type GuanYinPuSa interface {
    ZiJinPing() string
    LianHuaTai() string
}

type GodFactory interface {
    MonkeyKing() MonkeyKing
    GuanYinPuSa() GuanYinPuSa
}

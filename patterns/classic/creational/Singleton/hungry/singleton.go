package singleton

// https://blog.csdn.net/weixin_41709624/article/details/109022900

// 饿汉单例模式
var _singleton = new(singleton)

type singleton struct{}

func Singleton() *singleton {
    return _singleton
}

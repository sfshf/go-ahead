package singleton

import "sync"

// 懒汉单例模式
var (
    _singleton *singleton
    _once sync.Once
)

/*
Singleton是单例模式类；私有化，防止package外直接创建；

如果是空结构体，建议使用`饿汉单例模式`；

单例模式的类，如果有必要的属性，在多线程（多协程）读写应用时注意加锁；
建议：单例类的方法内尽量避免出现并发共用的属性；
*/
type singleton struct {}

// Singleton用户获取单例模式对象
func Singleton() *singleton {
    _once.Do(func() {
        _singleton = &singleton{}
    })
    return _singleton
}

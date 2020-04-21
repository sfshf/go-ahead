package singleton

import "sync"

/*
单例模式

使用懒惰模式的单例模式，使用双重检查加锁保证线程安全*/

//Singleton是单例模式类
type singleton struct{}  //私有化，防止package外直接创建

var singleIns *singleton
var once sync.Once

//GetInstance用户获取单例模式对象
func GetInstance() *singleton {
    once.Do(func() {
        singleIns = &singleton{}
    })
    return singleIns
}

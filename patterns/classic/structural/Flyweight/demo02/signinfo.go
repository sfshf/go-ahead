package signinfo

import "fmt"

type signInfoFactory struct {}

var Factory *signInfoFactory

func GetSignInfoFactory() *signInfoFactory {
    if Factory == nil {
        Factory = &signInfoFactory{}
    }
    return Factory
}

//原始工厂创建对象的函数，大量调用会产生大量对象，占用大量内存
// func (factory *signInfoFactory) GetSignInfo() *SignInfo {
//     return &SignInfo{}
// }

//使用享元模式，有效控制对象的数量
func (factory *signInfoFactory) GetSignInfo(keystr string) *SignInfo {

    for key, value := range Pool.pool {
        if key == keystr {
            fmt.Printf("直接从对象池中获取[key=%s]对象。\n", keystr)
            return value
        }
    }
    Pool.pool[keystr] = &SignInfo{}
    fmt.Printf("对象池中没有[key=%s]对象，生成新对象并加入对象池。\n", keystr)
    return Pool.pool[keystr]

}

func init() {
    Pool = GetPool4SignInfo()
    Factory = GetSignInfoFactory()
}

//后修复程序，另写的对象池结构体；此处代码设计只是突显出对象池;代码还可以优化
type Pool4SignInfo struct {
    pool map[string]*SignInfo
}

var Pool *Pool4SignInfo

func GetPool4SignInfo() *Pool4SignInfo {
    if Pool == nil {
        Pool = &Pool4SignInfo{
            pool: make(map[string]*SignInfo),
        }
        for i := 0; i < 4; i ++ {
            subject := fmt.Sprintf("科目%d",i)
            for j := 0; j < 30; j ++ {
                key := fmt.Sprintf("%s考试地点%d", subject, j)
                Pool.pool[key] = &SignInfo{}
            }
        }
    }

    return Pool

}

//原始结构体
type SignInfo struct {
    id string
    location string
    subject string
    postAddress string
}

func (si *SignInfo) SetId(id string) {
    si.id = id
}

func (si *SignInfo) Id() string {
    return si.id
}

func (si *SignInfo) SetLocation(location string) {
    si.location = location
}

func (si *SignInfo) Location() string {
    return si.location
}

func (si *SignInfo) SetSubject(subject string) {
    si.subject = subject
}

func (si *SignInfo) Subject() string {
    return si.subject
}

func (si *SignInfo) SetPostAddress(address string) {
    si.postAddress = address
}

func (si *SignInfo) PostAddress() string {
    return si.postAddress
}

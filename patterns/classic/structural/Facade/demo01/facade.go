package facade

import "fmt"

/*

外观模式

API为facade模块的外观接口，大部分代码使用此接口简化对facade类的访问。

facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

*/

func NewAPI() API {
    return &apiImpl{
        a: NewAModuleAPI(),
        b: NewBModuleAPI(),
    }
}

//API is facade interface of facade package
type API interface {
    Test() string
}

type apiImpl struct {
    a AModuleAPI
    b BModuleAPI
}

func (api *apiImpl) Test() string {
    aRet := api.a.TestA()
    bRet := api.b.TestB()
    return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAModuleAPI() AModuleAPI {
    return &aModuleImpl{}
}

type AModuleAPI interface {
    TestA() string
}

type aModuleImpl struct {}

func (a *aModuleImpl) TestA() string {
    return "A module running"
}

func NewBModuleAPI() BModuleAPI {
    return &bModuleImpl{}
}

type BModuleAPI interface {
    TestB() string
}

type bModuleImpl struct {}

func (b *bModuleImpl) TestB() string {
    return "B module running"
}

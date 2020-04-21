package chain

import "fmt"

/*

职责链模式

职责链模式用于分离不同职责，并且动态组合相关职责。

Golang实现职责链模式时候，因为没有继承的支持，使用链对象包涵职责的方式，即：

 · 链对象包含当前职责对象以及下一个职责链。
 · 职责对象提供接口表示是否能处理对应请求。
 · 职责对象提供处理函数处理相关职责。

同时可在职责链类中实现职责接口相关函数，使职责链对象可以当做一般职责对象使用。

*/

type Handler interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type ChainNode struct {
	Handler
	next *ChainNode
}

func NewChainNode(handler Handler) *ChainNode {
	return &ChainNode{
		Handler: handler,
	}
}

func (cn *ChainNode) SetNext(n *ChainNode) {
	cn.next = n
}

func (*ChainNode) HaveRight(money int) bool {
	return true
}

func (cn *ChainNode) HandleFeeRequest(name string, money int) bool {
	if cn.Handler.HaveRight(money) {
		return cn.Handler.HandleFeeRequest(name, money)
	}
	if cn.next != nil {
		return cn.next.HandleFeeRequest(name, money)
	}
	return false
}

type ProjectHandler struct{}

func (*ProjectHandler) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectHandler) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project handler permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project handler don't permit %s %d fee request\n", name, money)
	return false
}

type DepHandler struct{}

func (*DepHandler) HaveRight(money int) bool {
	return money > 500 && money < 5000
}

func (*DepHandler) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep handler permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep handler don't permit %s %d fee request\n", name, money)
	return false
}

type GeneralHandler struct{}

func (*GeneralHandler) HaveRight(money int) bool {
	return money > 5000
}

func (*GeneralHandler) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General handler permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General handler don't permit %s %d fee request\n", name, money)
	return false
}

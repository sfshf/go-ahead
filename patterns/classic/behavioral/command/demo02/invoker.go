package invoker

import "fmt"

type Invoker struct {
    ICommand
}

func (i *Invoker) Action() {
    i.ICommand.Execute()
}

type ICommand interface {
    Execute()
}

type DeletePageCommand struct {
    pg *PageGroup
}

func (dpc *DeletePageCommand) Execute() {
    dpc.pg.Find()
    dpc.pg.Delete()
}

type AddRequirementCommand struct {
    rg *RequirementGroup
}

func (arc *AddRequirementCommand) Execute() {
    arc.rg.Find()
    arc.rg.Add()
}

type IGroup interface {
    Find()
    Add()
    Delete()
    Change()
    Plan()
}

type RequirementGroup struct {}

func (rg *RequirementGroup) Find() {
    fmt.Println("找到需求组...")
}

func (rg *RequirementGroup) Add() {
    fmt.Println("客户要求增加一项需求...")
}

func (rg *RequirementGroup) Delete() {
    fmt.Println("客户要求删除一项需求...")
}

func (rg *RequirementGroup) Change() {
    fmt.Println("客户要求修改一项需求...")
}

func (rg *RequirementGroup) Plan() {
    fmt.Println("客户要求需求变更计划...")
}

type PageGroup struct {}

func (pg *PageGroup) Find() {
    fmt.Println("找到美工组...")
}

func (pg *PageGroup) Add() {
    fmt.Println("客户要求增加一个页面...")
}

func (pg *PageGroup) Delete() {
    fmt.Println("客户要求删除一个页面...")
}

func (pg *PageGroup) Change() {
    fmt.Println("客户要求修改一个页面...")
}

func (pg *PageGroup) Plan() {
    fmt.Println("客户要求页面变更计划...")
}

type CodeGroup struct {}

func (cg *CodeGroup) Find() {
    fmt.Println("找到代码组...")
}

func (cg *CodeGroup) Add() {
    fmt.Println("客户要求增加一项功能...")
}

func (cg *CodeGroup) Delete() {
    fmt.Println("客户要求删除一项功能...")
}

func (cg *CodeGroup) Change() {
    fmt.Println("客户要求修改一项功能...")
}

func (cg *CodeGroup) Plan() {
    fmt.Println("客户要求功能变更计划...")
}

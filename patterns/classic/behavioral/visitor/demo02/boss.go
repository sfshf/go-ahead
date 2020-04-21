package boss

import "fmt"

type IEmployee interface {
    Report(boss IBoss)
}

type IBoss interface {
    Visit(ee IEmployee)
}

type Gender string

const (
    MALE Gender = "男"
    FEMALE Gender = "女"
)

type BasicEmployee struct {
    Name string
    Sex Gender
    Salary int
}

func (be *BasicEmployee) ShowBasicInfo() {
    fmt.Printf("[Basic Info] name=%s, sex=%#v, salary=%d\n", be.Name, be.Sex, be.Salary)
}

type CommonEmployee struct {
    *BasicEmployee
    Job string  `普工谈工作`
}

func NewCommonEmployee(name string, sex Gender, salary int, job string) *CommonEmployee {
    basic := &BasicEmployee{
        Name: name,
        Sex: sex,
        Salary: salary,
    }
    common := &CommonEmployee{
        BasicEmployee: basic,
        Job: job,
    }
    return common
}

func (ce *CommonEmployee) Report(boss IBoss) {
    boss.Visit(ce)
}

type ManagerEmployee struct {
    *BasicEmployee
    Performance string  `管理者谈绩效`
}

func NewManagerEmployee(name string, sex Gender, salary int, performance string) *ManagerEmployee {
    basic := &BasicEmployee{
        Name: name,
        Sex: sex,
        Salary: salary,
    }
    manager := &ManagerEmployee{
        BasicEmployee: basic,
        Performance: performance,
    }
    return manager
}

func (me *ManagerEmployee) Report(boss IBoss) {
    boss.Visit(me)
}

type BossA struct {}

func (boss *BossA) Visit(ee IEmployee) {
    switch em := ee.(type) {
    case *CommonEmployee:
        em.ShowBasicInfo()
        fmt.Printf("[BossA/Common Employee] job is %s\n", em.Job)
    case *ManagerEmployee:
        em.ShowBasicInfo()
        fmt.Printf("[BossA/Manager Employee] performance is %s\n", em.Performance)
    }
}

type BossB struct {}

func (boss *BossB) Visit(ee IEmployee) {
    switch em := ee.(type) {
    case *ManagerEmployee:
        em.ShowBasicInfo()
        fmt.Printf("[BossB/Manager Employee] performance is %s\n", em.Performance)
    }
}

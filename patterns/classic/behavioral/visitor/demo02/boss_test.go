package boss

import "testing"

func TestBossA(t *testing.T) {

    employee := make([]IEmployee, 0)
    employee = append(employee, NewCommonEmployee("张三", MALE, 2100, "Java工程师"))
    employee = append(employee, NewCommonEmployee("李花花", FEMALE, 2700, "界面美工"))
    employee = append(employee, NewManagerEmployee("王五", MALE, 13000, "老板狗腿子，很会吹牛逼"))
    employee = append(employee, NewManagerEmployee("顺六", FEMALE, 23000, "老板秘书，老板的公司贤内助"))

    boss := &BossA{}
    for _, em := range employee {
        em.Report(boss)
    }

}

func TestBossB(t *testing.T) {

    employee := make([]IEmployee, 0)
    employee = append(employee, NewCommonEmployee("张三", MALE, 2100, "Java工程师"))
    employee = append(employee, NewCommonEmployee("李花花", FEMALE, 2700, "界面美工"))
    employee = append(employee, NewManagerEmployee("王五", MALE, 13000, "老板狗腿子，很会吹牛逼"))
    employee = append(employee, NewManagerEmployee("顺六", FEMALE, 23000, "老板秘书，老板的公司贤内助"))

    boss := &BossB{}
    for _, em := range employee {
        em.Report(boss)
    }

}

package corporation

import "testing"

func TestCorporation(t *testing.T) {

    ceo := NewICorp(BRANCH, "王大麻子", "总经理", 150000)
    cto := NewICorp(BRANCH, "刘大瘸子", "技术部经理", 30000)
    groupleader1 := NewICorp(BRANCH, "杨三乜斜", "技术一组组长", 15000)
    groupleader2 := NewICorp(BRANCH, "吴大棒槌", "技术二组组长", 12000)
    groupleader3 := NewICorp(BRANCH, "郑老六", "技术三组组长", 23000)
    developer1 := NewICorp(LEAF, "员工A", "开发员工", 8500)
    developer2 := NewICorp(LEAF, "员工B", "开发员工", 7500)
    developer3 := NewICorp(LEAF, "员工C", "开发员工", 6500)
    developer4 := NewICorp(LEAF, "员工D", "开发员工", 9000)
    developer5 := NewICorp(LEAF, "员工E", "开发员工", 8600)
    developer6 := NewICorp(LEAF, "员工F", "开发员工", 5500)
    marketmanager := NewICorp(BRANCH, "马二拐子", "销售部经理", 100000)
    saler1 := NewICorp(LEAF, "员工G", "销售员工", 20000)
    saler2 := NewICorp(LEAF, "员工H", "销售员工", 18000)
    financialmanager := NewICorp(BRANCH, "赵三驼子", "财务部经理", 300000)
    accountant := NewICorp(LEAF, "员工I", "财务员工", 4000)
    secretary := NewICorp(LEAF, "员工J", "总经理秘书", 4000)

    ceo.AddSubordinate(cto)
    ceo.AddSubordinate(marketmanager)
    ceo.AddSubordinate(financialmanager)
    ceo.AddSubordinate(secretary)

    cto.AddSubordinate(groupleader1)
    cto.AddSubordinate(groupleader2)
    cto.AddSubordinate(groupleader3)
    groupleader1.AddSubordinate(developer1)
    groupleader1.AddSubordinate(developer2)
    groupleader1.AddSubordinate(developer3)
    groupleader2.AddSubordinate(developer4)
    groupleader2.AddSubordinate(developer5)
    groupleader2.AddSubordinate(developer6)

    marketmanager.AddSubordinate(saler1)
    marketmanager.AddSubordinate(saler2)

    financialmanager.AddSubordinate(accountant)

    ceo.PrintInfo()

}

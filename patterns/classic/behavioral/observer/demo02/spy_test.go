package spy

import "testing"

func TestObservers(t *testing.T) {

    spy1 := NewSpy("间谍甲", "秦国")
    spy2 := NewSpy("间谍乙", "楚国")
    spy3 := NewSpy("间谍丙", "魏国")

    hanfeizi := NewSubject()
    hanfeizi.AddObserver(spy1)
    hanfeizi.AddObserver(spy2)
    hanfeizi.AddObserver(spy3)

    hanfeizi.HaveBreakfast()
    hanfeizi.HaveFun()

    hanfeizi.DeleteObserver(spy1)

    hanfeizi.HaveBreakfast()
    hanfeizi.HaveFun()


}

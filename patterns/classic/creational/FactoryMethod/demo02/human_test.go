package human

import "testing"

func CreateAHuman(creator HumanFactory, a ...interface{}) Human {
    return creator.Create(a...)
}

func TestHuman(t *testing.T) {

    var (
        creator HumanFactory
        person Human
    )

    creator = &YellowHumanFactory{}
    person = CreateAHuman(creator, 175, 157.6, 32, "黄皮肤的")
    t.Log(person.Talk())

    creator = &WhiteHumanFactory{}
    person = CreateAHuman(creator, 187, 187.6, 41, "白皮肤的")
    t.Log(person.Talk())

    creator = &BlackHumanFactory{}
    person = CreateAHuman(creator, 207, 197.6, 27, "黑皮肤的")
    t.Log(person.Talk())

}

package human

import "testing"

var (
    factory HumanFactory
    h Human
    f Female
    m Male
)

func TestYellowHumanFactory(t *testing.T) {
    factory = &YellowHumanFactory{}
    h = factory.CreateAFemale(155, 88.6, 23, "黄皮肤的")
    t.Log(h.Talk())
    f, ok := h.(Female)
    if !ok {
        t.Fatal("Invalid type assertion (Female)!")
    }
    t.Log(f.GetSex())

    h = factory.CreateAMale(177, 152.3, 27, "黄皮肤的")
    t.Log(h.Talk())
    m, ok := h.(Male)
    if !ok {
        t.Fatal("Invalid type assertion (Male)!")
    }
    t.Log(m.GetSex())
}

func TestBlackHumanFactory(t *testing.T) {
    factory = &BlackHumanFactory{}
    h = factory.CreateAFemale(165, 108.6, 21, "黑皮肤的")
    t.Log(h.Talk())
    f, ok := h.(Female)
    if !ok {
        t.Fatal("Invalid type assertion (Female)!")
    }
    t.Log(f.GetSex())

    h = factory.CreateAMale(197, 182.3, 29, "黑皮肤的")
    t.Log(h.Talk())
    m, ok := h.(Male)
    if !ok {
        t.Fatal("Invalid type assertion (Male)!")
    }
    t.Log(m.GetSex())
}

func TestWhiteHumanFactory(t *testing.T) {
    factory = &WhiteHumanFactory{}
    h = factory.CreateAFemale(175, 130.6, 25, "白皮肤的")
    t.Log(h.Talk())
    f, ok := h.(Female)
    if !ok {
        t.Fatal("Invalid type assertion (Female)!")
    }
    t.Log(f.GetSex())

    h = factory.CreateAMale(187, 182.3, 31, "白皮肤的")
    t.Log(h.Talk())
    m, ok := h.(Male)
    if !ok {
        t.Fatal("Invalid type assertion (Male)!")
    }
    t.Log(m.GetSex())
}

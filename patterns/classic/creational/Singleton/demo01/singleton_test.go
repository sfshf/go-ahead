package singleton

import "testing"

func TestSingleton(t *testing.T) {
    ins1 := GetInstance()
    ins2 := GetInstance()
    if ins1 != ins2 {
        t.Fatal("ins1 is not equal to ins2!")
    }
    ins3 := &singleton{}
    if ins1 != ins3 {
        t.Fatal("ins1 is not equal to ins3!")
    }
}

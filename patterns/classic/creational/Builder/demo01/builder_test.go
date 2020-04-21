package builder

import "testing"

func TestBuilder1(t *testing.T) {
    builder := &Builder1{}
    director := NewDirector(builder)
    director.Construct()
    res := builder.GetResult()
    if res != "-->Part1-->Part2-->Part3" {
        t.Fatalf("Builder1 fail expect -->Part1-->Part2-->Part3 acture %s", res)
    }
}

func TestBuilder2(t *testing.T) {
    builder := &Builder2{}
    director := NewDirector(builder)
    director.Construct()
    res := builder.GetResult()
    if res != 6 {
        t.Fatalf("Builder2 fail expect 6 acture %d", res)
    }
}

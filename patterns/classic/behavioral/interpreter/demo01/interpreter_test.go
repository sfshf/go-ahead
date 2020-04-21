package interpreter

import "testing"

func TestInterpreter(t *testing.T) {

    p := &Parser{}
    p.Parse("1 + 2 - 3 + 4 + 5 - 6")
    res := p.Result().Interpret()
    expect := 3
    if res != expect {
        t.Fatalf("Error: expect %d but got %d\n", expect, res)
    }

}

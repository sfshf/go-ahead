package main

import (
    "fmt"
)

func main() {

    strs := []string{"X17FS6", "PL4ZQ2", "JI8FR9", "XL8FQ6", "PY2ZR5", "KV7WS9", "JL2ZV3", "KI4WR2"}
    strs = CardinalSort(strs)
    fmt.Println(strs)

}

func CardinalSort(strs []string) []string {

    // 0-0,1-1,2-2,...,10-A,11-B...35-Z


    return nil
}

func countingsort(ints []int) []int {

    max := 0
    for i := 1; i < len(ints); i ++ {
        if ints[max] < ints[i] {
            max = i
        }
    }
    m := ints[max]+1
    equal := countKeysEqual(ints, m)
    less := countKeysLess(equal, m)
    return reArrange(ints, less, m)

}

func countKeysEqual(ints []int, m int) []int {

    equal := make([]int, m)
    for _, v := range ints {
        equal[v] ++
    }
    return equal

}

func countKeysLess(equal []int, m int) []int {

    less := make([]int, m)
    less[0] = 0
    for i := 1; i < len(equal); i ++ {
        less[i] = less[i-1]+equal[i-1]
    }
    return less

}

func reArrange(ints, less []int, m int) []int {

    sorted := make([]int, len(ints))
    next := less
    for i := 0; i < len(ints); i ++ {
        key := ints[i]
        index := next[key]
        sorted[index] = ints[i]
        next[key] ++
    }
    return sorted

}

package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 2, 12, 4, 6, 3, 7, 19, 13, 17, 23, 11, 31, 11}
    MergeSort(ints)
    fmt.Println(ints)

}

func MergeSort(ints []int) {

    mergeSort(ints)

}

func mergeSort(ints []int) {

    if len(ints) <= 1 {
        return
    }

    q := len(ints)/2
    mergeSort(ints[:q])
    mergeSort(ints[q:])
    merge(ints, q)

}

func merge(ints []int, q int) {

    tmp1 := make([]int, q)
    copy(tmp1, ints[:q])
    tmp2 := make([]int, len(ints)-q)
    copy(tmp2, ints[q:])

    i := 0
    p1 := 0
    p2 := 0
    end1 := len(tmp1)
    end2 := len(tmp2)

    for {

        if tmp1[p1] > tmp2[p2] {
            ints[i] = tmp2[p2]
            p2 ++
        } else {
            ints[i] = tmp1[p1]
            p1 ++
        }
        i ++
        if p1 == end1 {
            copy(ints[i:], tmp2[p2:])
            break
        } else if p2 == end2 {
            copy(ints[i:], tmp1[p1:])
            break
        }

    }

}

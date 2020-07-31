package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 2, 12, 4, 6, 3, 7, 19, 13, 17, 23, 11, 31, 11}
    QuickSort(ints)
    fmt.Println(ints)

}

// 确定快速排序法
func QuickSort(ints []int) {

    quicksort(ints)

}

func quicksort(ints []int) {

    if len(ints) <= 1 {
        return
    }
    q := partition(ints)
    quicksort(ints[:q])
    quicksort(ints[q+1:])

}

func partition(ints []int) int {

    q := 0
    r := len(ints)-1
    for u := q; u < r; u ++ {
        if ints[u] <= ints[r] {
            ints[q], ints[u] = ints[u], ints[q]
            q ++
        }
    }
    ints[q], ints[r] = ints[r], ints[q]
    return q

}

// TODO -- 随机快速排序法

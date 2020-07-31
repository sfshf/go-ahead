package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 2, 12, 4, 6, 3, 7, 19, 13, 17, 23, 11, 31, 11}
    ints = CountingSort(ints)
    fmt.Println(ints)

    ints = []int{11, 1, 3, 13, 15, 11, 7, 4, 91, 56, 43, 31, 27, 21, 78, 56, 23, 77, 21, 19, 61, 51, 12, 91, 93, 15, 16, 11, 7, 3, 1}
    ints = CountingSort(ints)
    fmt.Println(ints)

}

// 计数排序 -- 适用于数值（0和正整数）较小的序列进行排序；一般作为`基数排序`的内部排序算法；
func CountingSort(ints []int) []int {

    // 找出ints中最大的数
    max := 0
    for i := 1; i < len(ints); i ++ {
        if ints[max] < ints[i] {
            max = i
        }
    }
    m := ints[max]+1

    equal := countKeysEqual(ints, m)

    less := countKeysLess(equal, m)

    sorted := reArrange(ints, less, m)

    return sorted
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
    for i := 1; i < m; i ++ {
        less[i] = less[i-1] + equal[i-1]
    }
    return less

}

func reArrange(ints, less []int, m int) []int {

    sorted := make([]int, len(ints))
    next := less    // 由于ints的索引从0开始，所以next的索引也要从从0开始，所以`next := less`
    for i := 0; i < len(ints); i ++ {
        key := ints[i]
        index := next[key]
        sorted[index] = ints[i]
        next[key] ++
    }
    return sorted

}

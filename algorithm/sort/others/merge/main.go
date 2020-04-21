package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    mergesort(ints1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    mergesort(ints2)
    fmt.Println(ints2)


}

// 排序的元素越多，其效率优势越明显。
func mergesort(ints []int) {

    if len(ints) < 2 {
        return
    }

    p := 0
    r := len(ints)-1
    signal := 1<<63-1

    binary(ints, p, r, signal)

}

func binary(ints []int, p, r, signal int) {

    if p < r {                      // 此处判断不能忘记，否则会报错`stack overflow`
        q := (p+r)/2

        // 拆分
        binary(ints, p, q, signal)    // 递归调用
        binary(ints, q+1, r, signal)  // 递归调用
        // 归并排序
        merge(ints, p, q, r, signal)
    }

}

func merge(ints []int, p, q, r, signal int) {

    // 复制出来
    sub1 := make([]int, 0)
    sub2 := make([]int, 0)

    for _, v := range ints[p:q+1] {
        sub1 = append(sub1, v)
    }
    sub1 = append(sub1, signal)


    for _, v := range ints[q+1:r+1] {
        sub2 = append(sub2, v)
    }
    sub2 = append(sub2, signal)

    // 排序回去
    i := 0
    j := 0
    for k := p; k < r+1; k ++ {
        if sub1[i] <= sub2[j] {
            ints[k] = sub1[i]
            i ++
        } else {
            ints[k] = sub2[j]
            j ++
        }
    }

}

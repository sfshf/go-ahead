package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    heapsort(ints1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    heapsort(ints2)
    fmt.Println(ints2)

}

func heapsort(ints []int) {

    // 建堆
    buildheap(ints)

    fmt.Println(ints)

    for i := len(ints)-1; i > 0; i -- {

        ints[i], ints[0] = ints[0], ints[i]

        // fmt.Println(ints)

        fixheap(ints[:i])

    }

}

func fixheap(ints []int) {

    changed := true

    for i := 0; i < len(ints) && changed ; {

        changed = false
        i, changed = minfixheap(ints, i, len(ints))
        // fmt.Println(ints)
    }

}

func buildheap(ints []int) {

    for i := 0; i < len(ints)/2; i ++ {
        for j := len(ints)/2; j > i-1; j -- {
            minfixheap(ints, j, len(ints))
        }
    }

}

// 从i节点开始调整,n为节点总数 从0开始计算 i节点的子节点为 2*i+1, 2*i+2
// 完全二叉树 -- 没有左子结点，就不会有右子节点
func minfixheap(ints []int, i, n int) (max int, changed bool) {

    j := 2 * i + 1  // 左子树根结点
    max = i
    changed = false

    if j < n {

        // 与左子树根结点比较
        if ints[max] < ints[j] {
            max = j
        }

        // 与右子树根结点比较
        if j + 1 < n && ints[max] < ints[j+1] {
            max = j + 1
        }

    }

    if i != max {
        ints[i], ints[max] = ints[max], ints[i]
        changed = true
    }

    return

}

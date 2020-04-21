package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 3, 43, 23, 11, 6, 7, 21, 2, 78, 56, 47}
    fmt.Println(ints)
    bubblesort(ints)
    fmt.Println(ints)

    arr := [12]int{1, 3, 43, 23, 11, 6, 7, 21, 2, 78, 56, 47}
    fmt.Println(arr)
    bubblesort2(arr)
    fmt.Println(arr)  // 未排序

    /*
        在Go语言中，一定要注意`数组`和`切片`的区别：
            `切片`底层利用了`数组`，其有三个属性：指针、长度和容量；所以，在函数（或方法）传值时，与`数组`不同，
        `数组`是`传值`（即传递的是数组的副本），`切片`是`传址`（即传递的是底层数组地址的副本）。
    */


}

//冒泡排序法（从小到大）
func bubblesort(ints []int) {

    for i, changed := len(ints)-1, true; i > 0 && changed; i -- {
        changed = false
        for j := 0; j < i; j ++ {
            if ints[j] > ints[j+1] {
                ints[j], ints[j+1] = ints[j+1], ints[j]
                changed = true
            }
        }
    }

}

func bubblesort2(arr [12]int) {

    for i, changed := len(arr)-1, true; i > 0 && changed; i -- {
        changed = false
        for j := 0; j < i; j ++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                changed = true
            }
        }
    }

}

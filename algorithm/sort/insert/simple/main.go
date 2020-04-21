package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    insertsort2(ints1)
    // insertsort(ints1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    insertsort2(ints2)
    // insertsort(ints2)
    fmt.Println(ints2)

}

func insertsort(ints []int) {

    for i := 1; i < len(ints); i ++ {

        key := ints[i]

        for j := i-1; ; j -- {

            if j >= 0 && ints[j] > key {  // 插入排序，内层循环的迭代次数取决于外层循环的索引i和数组元素值。
                ints[j+1] = ints[j]
            } else {
                ints[j+1] = key
                break
            }

        }

    }

}

// 简单插入排序法（从小到大）
func insertsort2(ints []int) {

    for i := 1; i < len(ints); i ++ {

        for j := i; j > 0; j -- {

            if ints[j-1] > ints[j] {

                ints[j-1], ints[j] = ints[j], ints[j-1]

            } else {  // 假定前i个元素是有序的

                break

            }

        }

    }

}

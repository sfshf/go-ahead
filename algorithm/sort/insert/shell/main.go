package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    shellsort(ints1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    shellsort(ints2)
    fmt.Println(ints2)

}

// 希尔排序算法（从小到大）
func shellsort(ints []int) {

    // 变化的偏移量
    for offset := len(ints)/2; offset > 0; offset -- {

        // 根据偏移量产生的分组
        for i := 0; i + offset < len(ints); i ++ {

            // 对每个分组进行`简单插入排序`
            for j := i + offset; j < len(ints); j += offset {

                for k := j; k > j - offset; k -= offset {

                    if ints[k] < ints[k - offset] {
                        ints[k], ints[k - offset] = ints[k - offset], ints[k]
                    }

                }

            }

        }

    }

}

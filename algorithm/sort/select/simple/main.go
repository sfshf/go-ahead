package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    selectsort(ints1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    selectsort(ints2)
    fmt.Println(ints2)

}

func selectsort(ints []int) {

    for i := 0; i < len(ints); i ++ {

        min := i

        for j := i+1; j < len(ints); j ++ {

            if ints[j] < ints[min] {
                min = j
            }

        }

        if i != min {
            ints[i], ints[min] = ints[min], ints[i]
        }

    }

}

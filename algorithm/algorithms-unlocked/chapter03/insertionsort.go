package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 2, 12, 4, 6, 3, 7, 19, 13, 17, 23, 11}
    InsertionSort(ints)
    fmt.Println(ints)

}

func InsertionSort(ints []int) {

    for i := 1; i < len(ints); i ++ {
        key := ints[i]
        for j := i-1; j > 0; j -- {
            if ints[j] > key {
                ints[j+1] = ints[j]
            } else {
                ints[j+1] = key
                break
            }
        }
    }

}

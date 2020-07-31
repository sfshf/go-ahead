package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 2, 12, 4, 6, 3, 7, 19, 13, 17, 23, 11}
    SelectionSort(ints)
    fmt.Println(ints)

}

func SelectionSort(ints []int) {

    n := len(ints)
    for i := 0; i < n; i ++ {
        smallest := i
        for j := i+1; j < n; j ++ {
            if ints[smallest] > ints[j] {
                smallest = j
            }
        }
        ints[i], ints[smallest] = ints[smallest], ints[i]
    }

}

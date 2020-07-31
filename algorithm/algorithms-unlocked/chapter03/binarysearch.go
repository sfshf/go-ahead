package main

import (
    "fmt"
)

func main() {

    ints := []int{1, 3, 4, 6, 7, 9, 11, 13, 16, 18, 32, 37, 41} // 有序
    fmt.Println(BinarySearch(ints, 13))
    fmt.Println(BinarySearch(ints, 31))
    println()
    fmt.Println(RecursiveBinarySearch(ints, 13))
    fmt.Println(RecursiveBinarySearch(ints, 31))

}

func BinarySearch(ints []int, x int) int {

    p := 0                      // 左指针
    r := len(ints)-1            // 右指针
    for p <= r {
        i := (p+r)/2
        if ints[i] > x {
            r = i-1
        } else if ints[i] < x {
            p = i+1
        } else if ints[i] == x {
            return i
        }
    }
    return -1

}

func RecursiveBinarySearch(ints []int, x int) int {

    return recursiveBinarySearch(ints, 0, len(ints)-1, x)

}

func recursiveBinarySearch(ints []int, p, r, x int) int {

    if p < 0 || r < 0 {
        return -1
    }
    if p <= r {
        i := (p+r)/2
        if ints[i] > x {
            r = i-1
        } else if ints[i] < x {
            p = i+1
        } else if ints[i] == x {
            return i
        }
        return recursiveBinarySearch(ints, p, r, x)
    }
    return -1

}

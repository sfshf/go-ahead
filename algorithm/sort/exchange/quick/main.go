package main

import (
    "fmt"
)

func main() {

    ints1 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37 }
    fmt.Println(ints1)
    quicksort(ints1, 0, len(ints1)-1)
    fmt.Println(ints1)

    ints2 := []int{ 21, 13, 57, 93, 71, 83, 19, 7, 11, 37, 111, 34, 17, 41, 59 }
    fmt.Println(ints2)
    quicksort(ints2, 0, len(ints2)-1)
    fmt.Println(ints2)

}

// 快速排序法（从小到大）
func quicksort(ints []int, i, j int) {

    if i >= j {
        return
    }

    front := i
    rear := j
    key := ints[front]

    for rear > front {

        for rear > front && ints[rear] >= key {
            rear --
        }

        if front < rear {
            ints[front] = ints[rear]
            front ++
        }

        for front < rear && ints[front] <= key {
            front ++
        }

        if front < rear {
            ints[rear] = ints[front]
            rear --
        }

    }

    // front == rear
    ints[front] = key
    quicksort(ints, i, front-1)
    quicksort(ints, rear+1, j)



}

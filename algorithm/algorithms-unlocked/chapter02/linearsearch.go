package main

import (
    "fmt"
    "time"
)

func main() {

    ints := make([]int, 100000000)
    ints[99999999] = 99999999

    t1 := time.Now()
    fmt.Println(LinearSearch(ints, 99999999))
    t2 := time.Now()
    fmt.Println(t2.Sub(t1))
    fmt.Println(BetterLinearSearch(ints, 99999999))
    t3 := time.Now()
    fmt.Println(t3.Sub(t2))
    fmt.Println(SentinelLinearSearch(ints, 99999999))
    t4 := time.Now()
    fmt.Println(t4.Sub(t3))
    ints[999999] = 999999
    fmt.Println(RecursiveLinearSearch(ints, 999999))
    t5 := time.Now()
    fmt.Println(t5.Sub(t4))
    println()
    fmt.Println(Factorial(20))

}

func LinearSearch(ints []int, x int) int {

    answer := -1
    // 完全遍历了数组
    for i := 0; i < len(ints); i ++ {
        if ints[i] == x {
            answer = x
        }
    }
    return answer

}

func BetterLinearSearch(ints []int, x int) int {

    // 找到第一个符合的值，则返回
    for i := 0; i < len(ints); i ++ {
        if ints[i] == x {
            return i
        }
    }

    return -1

}

func SentinelLinearSearch(ints []int, x int) int {

    n := len(ints)-1
    tmp := ints[n]      // 暂存数组最右边元素值
    ints[n] = x         // 设置假盒子
    i := 0
    for {               // 无需考虑数组是否会索引越界
        if ints[i] == x {
            break
        }
        i ++
    }
    if i != n {
        return i
    } else {
        if tmp == x {   // 判断数组最右边元素值是否和要查找的值相等
            return i
        } else {
            return -1
        }
    }

}

func RecursiveLinearSearch(ints []int, x int) int {

    n := len(ints)
    return recursiveLinearSearch(ints, 0, n, x)

}

func recursiveLinearSearch(ints []int, i, n, x int) int {
    if i >= n || i < 0 {
        return -1
    }
    if ints[i] != x {
        i ++
        return recursiveLinearSearch(ints, i, n, x)
    } else {
        return i
    }
}

func Factorial(n int) int {

    return factorial(n)

}

func factorial(n int) int {
    if n < 0 {
        return -1
    }
    if n == 0 {
        return 1
    }
    return n*factorial(n-1)
}

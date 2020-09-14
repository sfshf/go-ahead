package main

import (
    "fmt"
)

func main() {

    strs := []string{"XI7FS6", "PL4ZQ2", "JI8FR9", "XL8FQ6", "PY2ZR5", "KV7WS9", "JL2ZV3", "KI4WR2"}
    strs = CardinalSort(strs)
    //fmt.Println(strs)

}

var (
    // 0-0,1-1,2-2,...,10-A,11-B...35-Z
    tmpMap = map[rune]int {
        '0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
        'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
        'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
        'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35,
    }
)

func CardinalSort(strs []string) []string {

    if len(strs) <= 1 {
        return strs
    }
    // 假定元素规则一致，如每个元素的长度均为6
    // 从右向左
    for i := 5; i >= 0; i -- {
        datas := make([]Data, 0)
        for j := 0; j < len(strs); j ++ {
            var data Data
            data.Key = tmpMap[rune(strs[j][i])]
            data.real = strs[j]
            datas = append(datas, data)
        }
        sorted := countingsort(datas)
        strs = make([]string, 0)
        for _, v := range sorted {
            strs = append(strs, v.real)
        }
        fmt.Println(strs)
    }

    return strs
}

type Data struct {
    Key int
    real string
}

func countingsort(datas []Data) []Data {

    max := 0
    for i := 1; i < len(datas); i ++ {
        if datas[max].Key < datas[i].Key {
            max = i
        }
    }
    m := datas[max].Key+1
    equal := countKeysEqual(datas, m)
    less := countKeysLess(equal, m)
    return reArrange(datas, less, m)

}

func countKeysEqual(datas []Data, m int) []int {

    equal := make([]int, m)
    for _, v := range datas {
        equal[v.Key] ++
    }
    return equal

}

func countKeysLess(equal []int, m int) []int {

    less := make([]int, m)
    less[0] = 0
    for i := 1; i < len(equal); i ++ {
        less[i] = less[i-1]+equal[i-1]
    }
    return less

}

func reArrange(datas []Data, less []int, m int) []Data {

    sorted := make([]Data, len(datas))
    next := less
    for i := 0; i < len(datas); i ++ {
        key := datas[i]
        index := next[key.Key]
        sorted[index] = datas[i]
        next[key.Key] ++
    }
    return sorted

}

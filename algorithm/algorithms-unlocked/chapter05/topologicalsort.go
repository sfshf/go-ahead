package main

import (
    "fmt"
)

// TODO
func main() {

    // 数据序列
    data := []string{"undershorts", "socks", "compression shorts", "hose", "cup",
                    "pants", "skates", "leg pads", "T-shirts", "chest pad",
                    "sweater", "mask", "catch glove", "blocker"}
    // 有向图（关系图）
    adjacencyList := [][]int{[]int{2}, []int{3}, []int{3, 4}, []int{5}, []int{5},
                            []int{6, 10}, []int{7}, []int{12},[]int{9}, []int{10},
                            []int{11}, []int{12}, []int{13}, nil}

    sorted := TopologicalSort(data, adjacencyList)
    fmt.Println(sorted)

}

func TopologicalSort(data []string, adjacencyList [][]int) []string {

    indegree := make([]int, len(data))
    sorted := make([]string, len(data))
    for _, vs := range adjacencyList {
        for _, v := range vs {
            indegree[v] ++
        }
    }

    next := InitStack()
    for u, v := range indegree {
        if v == 0 {
            next.Push(&data{
                uIndex: u,
            })
        }
    }

    for next.HasNext() {

    }

}

// 栈
type stack struct {
    root data
}

type data struct {
    uIndex int
    next *data
}

func InitStack() *stack {
    return &stack{
        root: data{
            uIndex: -1,
        },
    }
}

func (s *stack) HasNext() bool {
    return s.root.next != nil
}

func (s *stack) Pop() *data {
    pop := s.root.next
    if pop != nil { s.root.next = pop.next }
    return pop
}

func (s *stack) Push(data *data) {
    data.next = s.root.next
    s.root.next = data
}

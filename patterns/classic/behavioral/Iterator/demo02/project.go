package project

import (
    "container/list"
    "fmt"
)

type IProjectIterator interface {
    First() interface{}
    IsDone() bool
    Next() interface{}
}

type Aggregate interface {
    Add(value interface{})
    Iterator() IProjectIterator
}

type Projects struct {
    context *list.List
}

func NewProjects() *Projects {
    return &Projects{
        context: &list.List{},
    }
}

func (projects *Projects) Add(value interface{}) {
    projects.context.PushBack(value)
}

func (projects *Projects) Iterator() IProjectIterator {
    copyprojects := *projects
    return &ProjectIterator{
        projects: &copyprojects,
    }
}

type ProjectIterator struct {
    projects *Projects
    nextele *list.Element
    next interface{}
    count int
}

func (pi *ProjectIterator) First() interface{} {
    pi.count = 0
    return pi.Next()
}

func (pi *ProjectIterator) IsDone() bool {
    return pi.count == pi.projects.context.Len()
}

func (pi *ProjectIterator) Next() interface{} {
    if pi.count == 0 {
        pi.nextele = pi.projects.context.Front()
        pi.count = 1
    } else if !pi.IsDone() {
        pi.nextele = pi.nextele.Next()
        pi.count += 1
    }
    pi.next = pi.nextele.Value
    return pi.next
}

func IteratorPrint(iterator IProjectIterator) {
    for !iterator.IsDone()  {
        value := iterator.Next()
        fmt.Printf("%#v\n", value)
    }
}

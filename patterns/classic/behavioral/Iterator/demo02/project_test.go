package project

func ExampleProjectIterator() {

    projects := NewProjects()
    projects.Add("先打开冰箱；")
    projects.Add("再将大象塞进冰箱；")
    projects.Add("最后关上冰箱门。")

    iterator := projects.Iterator()

    IteratorPrint(iterator)
    //Output:
    //"先打开冰箱；"
    //"再将大象塞进冰箱；"
    //"最后关上冰箱门。"


}

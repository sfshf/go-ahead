package signinfo


func ExampleFlyweight() {

    Factory.GetSignInfo("科目1考试地点1")
    Factory.GetSignInfo("科目4考试地点100")
    //Output:
    //直接从对象池中获取[key=科目1考试地点1]对象。
    //对象池中没有[key=科目4考试地点100]对象，生成新对象并加入对象池。

}

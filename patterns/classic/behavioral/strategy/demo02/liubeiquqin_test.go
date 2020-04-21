package liubeiquqin

func ExampleLiubeiQuQin() {

    var ctx *Context = &Context{
        IStrategy: &BackDoor{},
    }
    ctx.Execute()

    ctx.IStrategy = &GivenGreenLight{}
    ctx.Execute()

    ctx.IStrategy = &BlockEnemy{}
    ctx.Execute()

    //Output:
    //[BackDoor]找乔国老帮忙，让吴国太给孙权施加压力。
    //[GivenGreenLight]求吴国太开绿灯，放行！
    //[BlockEnemy]孙夫人断后，挡住追兵！

}

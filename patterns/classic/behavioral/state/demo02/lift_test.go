package lift

func ExampleLift() {

    liftctx := NewLiftContext()
    liftctx.SetLiftState(liftctx.CloseState)
    liftctx.Open()
    liftctx.Close()
    liftctx.Run()
    liftctx.Stop()
    //Output:
    //电梯门开启...
    //电梯门关闭...
    //电梯正在运行...
    //电梯处于停止状态...

}

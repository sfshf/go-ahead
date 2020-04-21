package lift

import "fmt"

type Lift interface {
    open()
    close()
    run()
    stop()
}

type LiftContext struct {
    Lift
    *OpenState
    *CloseState
    *RunState
    *StopState
}

func NewLiftContext() *LiftContext {
    lc := &LiftContext{}
    open := &OpenState{
        LiftContext: lc,
    }
    close := &CloseState{
        LiftContext: lc,
    }
    run := &RunState{
        LiftContext: lc,
    }
    stop := &StopState{
        LiftContext: lc,
    }
    lc.OpenState = open
    lc.CloseState = close
    lc.RunState = run
    lc.StopState = stop
    return lc
}

func (lc *LiftContext) SetLiftState(liftstate Lift) {
    lc.Lift = liftstate
}

func (lc *LiftContext) Open() {
    lc.Lift.open()
}

func (lc *LiftContext) Close() {
    lc.Lift.close()
}

func (lc *LiftContext) Run() {
    lc.Lift.run()
}

func (lc *LiftContext) Stop() {
    lc.Lift.stop()
}

type OpenState struct {
    *LiftContext
}

func (os *OpenState) open() {
    fmt.Println("电梯门开启...")
}

func (os *OpenState) close() {
    os.LiftContext.SetLiftState(os.LiftContext.CloseState)
    os.LiftContext.Close()
}

func (os *OpenState) run() {}

func (os *OpenState) stop() {}

type CloseState struct {
    *LiftContext
}

func (cs *CloseState) open() {
    cs.LiftContext.SetLiftState(cs.LiftContext.OpenState)
    cs.LiftContext.Open()
}

func (cs *CloseState) close() {
    fmt.Println("电梯门关闭...")
}

func (cs *CloseState) run() {
    cs.LiftContext.SetLiftState(cs.LiftContext.RunState)
    cs.LiftContext.Run()
}

func (cs *CloseState) stop() {
    cs.LiftContext.SetLiftState(cs.LiftContext.StopState)
    cs.LiftContext.Stop()
}

type RunState struct {
    *LiftContext
}

func (rs *RunState) open() {}

func (rs *RunState) close() {}

func (rs *RunState) run() {
    fmt.Println("电梯正在运行...")
}

func (rs *RunState) stop() {
    rs.LiftContext.SetLiftState(rs.LiftContext.StopState)
    rs.LiftContext.Stop()
}

type StopState struct {
    *LiftContext
}

func (ss *StopState) open() {
    ss.LiftContext.SetLiftState(ss.LiftContext.OpenState)
    ss.LiftContext.Open()
}

func (ss *StopState) close() {
    ss.LiftContext.SetLiftState(ss.LiftContext.CloseState)
    ss.LiftContext.Close()
}

func (ss *StopState) run() {
    ss.LiftContext.SetLiftState(ss.LiftContext.RunState)
    ss.LiftContext.Run()
}

func (ss *StopState) stop() {
    fmt.Println("电梯处于停止状态...")
}

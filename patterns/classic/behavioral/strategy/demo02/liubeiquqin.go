package liubeiquqin

import "fmt"

type Context struct {
    IStrategy
}

func (ctx *Context) Execute() {
    ctx.IStrategy.Operate()
}

type IStrategy interface {
    Operate()
}

type BackDoor struct {}

func (*BackDoor) Operate() {
    fmt.Println("[BackDoor]找乔国老帮忙，让吴国太给孙权施加压力。")
}

type GivenGreenLight struct {}

func (*GivenGreenLight) Operate() {
    fmt.Println("[GivenGreenLight]求吴国太开绿灯，放行！")
}

type BlockEnemy struct {}

func (*BlockEnemy) Operate() {
    fmt.Println("[BlockEnemy]孙夫人断后，挡住追兵！")
}

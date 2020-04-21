package gameplayer

import "fmt"

type IGamePlayer interface {
    Login(user, password string)
    KillBoss()
    Upgrade()
}

type GamePlayer struct {
    user string
    password string
}

func (gp *GamePlayer) Login(user, password string) {
    fmt.Printf("[Player]登录名为<%s>的用户<%s>登录成功！\n", user, gp.user)
}

func (gp *GamePlayer) KillBoss() {
    fmt.Printf("[Player]<%s>在打怪...\n", gp.user)
}

func (gp *GamePlayer) Upgrade() {
    fmt.Printf("[Player]<%s>又升了一级！\n", gp.user)
}

type GamePlayerProxy struct {
    *GamePlayer
}

func (gpp *GamePlayerProxy) Login(user, password string) {
    fmt.Printf("[Proxy]登录名为<%s>的用户<%s>登录成功！\n", user, gpp.user)
}

func (gpp *GamePlayerProxy) KillBoss() {
    fmt.Printf("[Proxy]<%s>在打怪...\n", gpp.user)
}

func (gpp *GamePlayerProxy) Upgrade() {
    fmt.Printf("[Proxy]<%s>又升了一级！\n", gpp.user)
}

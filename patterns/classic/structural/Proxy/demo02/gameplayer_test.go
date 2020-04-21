package gameplayer

import "testing"

func TestGamePlayerProxy(t *testing.T) {
    player := &GamePlayer{
        user: "张三",
        password: "123456",
    }
    proxy := &GamePlayerProxy{
        GamePlayer: player,
    }
    proxy.Login(proxy.user, proxy.password)
    proxy.KillBoss()
    proxy.Upgrade()
}

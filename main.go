package main

import (
	"github.com/jodealter/mmo_game_zinx/apis"
	"github.com/jodealter/mmo_game_zinx/core"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)
	player.SynPid()
	player.BroadCastStartPosition()
	core.WorldMgrObj.AddPlayer(player)
	conn.SetProperty("pid", player.Pid)
}
func OnConnectionLost(conn ziface.IConnection) {
	pid, _ := conn.GetProperty("pid")
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.Offline()
}
func main() {
	s := znet.NewServr("MMo Game Zinx")

	s.SetOnConnStart(OnConnectionAdd)
	s.AddRouter(2, &apis.WorldChatApi{})
	s.Serve()
}

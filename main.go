package main

import (
	"github.com/jodealter/mmo_game_zinx/core"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)
	player.SynPid()
	player.BroadCastStartPosition()
}

func main() {
	s := znet.NewServr("MMo Game Zinx")
	s.Serve()
}

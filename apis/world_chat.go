package apis

import (
	"fmt"
	"github.com/jodealter/mmo_game_zinx/core"
	"github.com/jodealter/mmo_game_zinx/pb"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (w *WorldChatApi) Hadnle(request ziface.IRequest) {
	protomsg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), protomsg)
	if err != nil {
		fmt.Println("unmarshal errors", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.Talk(protomsg.Content)
}

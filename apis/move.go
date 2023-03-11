package apis

import (
	"fmt"
	"github.com/jodealter/mmo_game_zinx/core"
	"github.com/jodealter/mmo_game_zinx/pb"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type Moveapi struct {
	znet.BaseRouter
}

func (m *Moveapi) Handle(request ziface.IRequest) {
	proto_msg := &pb.Position{}

	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("unmarshal err ", err)
		return
	}
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty erroe :", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)
}

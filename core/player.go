package core

import (
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"sync"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var PidGen int32 = 1  //用来生成玩家
var IdLock sync.Mutex //保护上边的锁

func NewPlayer(conn ziface.IConnection) *Player {
	//生成玩家id
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	//创建一个玩家对象
	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(20)),
		V:    0,
	}
	return p
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal error", err)
		return
	}
	if p.Conn == nil {
		fmt.Println("connection is nil")
		return
	}
	if err = p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("player sendmsg error")
		return
	}
}
func (p *Player) SynPid() {
	proto_msg := &pb.SyncPid{Pid: p.Pid}
	p.SendMsg(1, proto_msg)
}
func (p *Player) BroadCastStartPosition() {
	proto_msg := &pb.Broadcast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.Broadcast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		}},
	}
	p.SendMsg(200, proto_msg)
}

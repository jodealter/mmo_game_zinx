package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	//格子左边的边界坐标
	GID int
	//格子右边的边界坐标
	MinX int
	//格子左边的边界坐标
	MaxX int
	//格子上边的边界坐标
	MinY int
	//格子下边的边界坐标
	MaxY int
	//当前格子内的玩家或者物体的id集合
	playeIDs map[int]bool
	//保护当前集合的锁
	pIDLock sync.RWMutex
}

// 初始化当前的格子的方法
func NewGrid(gID, Minx, MaxX, MinY, MaxY int) *Grid {
	return &Grid{
		GID:      gID,
		MinX:     Minx,
		MaxX:     MaxX,
		MaxY:     MaxY,
		MinY:     MinY,
		playeIDs: make(map[int]bool),
		pIDLock:  sync.RWMutex{},
	}
}

// 给格子添加一个玩家
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playeIDs[playerID] = true
}

//从格子中删除一个玩家

func (g *Grid) Remove(playID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playeIDs, playID)
}

// 得到当前格子中的所有玩家
func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	for k, _ := range playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}

// 调式使用，打印格子的基本信息
func (g *Grid) String() string {
	return fmt.Sprintf("Grid id : %d, minX : %d, maxX:%d,minY : %d, maxX:%d, playerIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playeIDs)
}

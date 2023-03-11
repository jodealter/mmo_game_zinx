package core

import "sync"

type WorldManager struct {
	AoiMgr *AOIManager
	Player map[int32]*Player
	plock  sync.RWMutex
}

var WorldMgrObj *WorldManager

// 初始化方法
func init() {

}

package core

import "fmt"

const (
	AOI_MIN_X_int  = 85
	AOI_MAX_X_int  = 410
	AOI_CNTS_X_int = 10
	AOI_MIN_Y_int  = 75
	AOI_MAX_Y_int  = 400
	AOI_CNTS_Y_int = 20
)

type AOIManager struct {
	//区域的左边界
	MinX int
	//区域的有边界
	MaxX int
	//x轴上格子的数量
	CntsX int
	//区域的上边界
	MinY int
	//区域的下边界
	MaxY int
	//y轴上的格子的数量
	CntsY int
	//区域内格子的集合
	grids map[int]*Grid
}

func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			gid := y*cntsX + x
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(), aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridWidth(), aoiMgr.MinY+(y+1)*aoiMgr.gridWidth())
		}
	}
	return aoiMgr
}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\n MinX:%d, MaxX:%d,cntsX:%d,MinY:%d, MaxY:%d,cntsY:%d\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断是否在manager中
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//初始化切片值

	grids = append(grids, m.grids[gID])
	//需要gid的左边是否有格子？右边是否有格子
	//需要通过gid得到当前格子x轴的编号 idx = id % nx
	idx := gID % m.CntsX
	//判断idx左边是否有空格，如果有就加入进gidsx集合
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	//判断idx右边是否有空格，如果有就加入进gidsx集合
	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}
	//便利x轴格子上下是否有格子
	//得到当前x轴格子的集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}
	//便利gidsx 集合中每个格子的gid

	for _, v := range gidsX {
		idy := v / m.CntsX
		//判断上下是否存在格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}

	}
	return
}

func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLength()
	return idy*m.CntsX + idx

}
func (m *AOIManager) GetPidByPos(x, y float32) (playerIDs []int) {
	//获取当前玩家所在格子的id
	gID := m.GetGidByPos(x, y)
	grids := m.GetSurroundGridsByGid(gID)
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Printf("===>grid ID : %d, pids : %v===", grid.GID, grid.GetPlayerIDs())
	}
	return
}

// 通过pid将player添加到格子（给定gid）中
func (m *AOIManager) AddPidToGrid(pID, Gid int) {
	m.grids[Gid].Add(pID)
}

// 从格子中一处某个player，通过pid与gid
func (m *AOIManager) RemovePidFromGrid(pid int, Gid int) {
	m.grids[Gid].Remove(pid)
}
func (m *AOIManager) GetPidsByGid(gid int) (playIDs []int) {
	playIDs = m.grids[gid].GetPlayerIDs()
	return
}
func (m *AOIManager) AddToGridByPos(pid int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.Add(gid)
}
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.Remove(pID)
	return
}

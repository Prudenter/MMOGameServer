package core

import (
	"sync"
	"fmt"
)

/*
	AOI兴趣点,格子的模块,相关操作
*/

type Grid struct {
	//格子ID
	GID int
	//格子的左边边界坐标
	MinX int
	//格子的右边边界的坐标
	MaxX int
	//格子的上边边界的坐标
	MinY int
	//格子的下边边界的坐标
	MaxY int
	//当前格子内 玩家/物体 成员的ID集合 map[玩家/物体ID]
	playerIDs map[int]interface{}
	//保护当前格子内容map的锁
	pIDLock sync.RWMutex
}

//初始化格子的方法
func NewGrid(gId, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gId,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]interface{}),
	}
}

//给格子添加一个玩家
func (g *Grid) Add(playerId int, player interface{}) {
	//添加写锁
	g.pIDLock.Lock()
	//解锁
	defer g.pIDLock.Unlock()
	g.playerIDs[playerId] = player
}

//从格子中删除一个玩家
func (g *Grid) Remove(playerId int) {
	//添加写锁
	g.pIDLock.Lock()
	//解锁
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerId)
}

//得到当前格子内的所有玩家的Id
func (g *Grid) GetPlayerIds() (playerIds []int) {
	//添加读锁
	g.pIDLock.RLock()
	//解锁
	defer g.pIDLock.RUnlock()
	//遍历map集合,将key封装成一个slice并返回
	for playerId,_ := range g.playerIDs{
		playerIds = append(playerIds,playerId)
	}
	return
}

//调试打印格子信息方法
func (g *Grid) String()string {
	return fmt.Sprintf("Grid id : %d, minX:%d, maxX:%d , minY:%d, maxY:%d, playerIDs:%v\n",
		g.GID,g.MinX,g.MaxX,g.MinY,g.MaxY,g.playerIDs)
}

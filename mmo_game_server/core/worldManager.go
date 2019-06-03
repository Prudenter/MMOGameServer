/**
* @Author: ASlowPerson  
* @Date: 19-6-1 上午11:23
*/

package core

import (
	"sync"
)

/*
	当前世界地图的边界参数
*/
const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

/*
	当前场景的世界管理模块
*/
type WorldManager struct {
	//当前全部在线的Player集合
	Players map[int32]*Player
	//保护Player集合的锁
	pLock sync.RWMutex
	//AOIManager当前的地图管理器
	AorMgr *AOIManager
}

//对外提供一个全局的世界管理模块指针
var WorldMgrObj *WorldManager

//定义init方法,只要导入当前包,就会执行init方法
func init() {
	//创建一个全局的世界管理对象
	WorldMgrObj = NewWorldManager()
}

//初始化方法
func NewWorldManager() *WorldManager {
	wm := &WorldManager{
		AorMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		Players: make(map[int32]*Player),
	}
	return wm
}

//添加一个玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	//将玩家加入世界管理器中
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()
	//将玩家加入到世界地图中
	wm.AorMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

//删除一个玩家
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	wm.pLock.Lock()
	//先通过pid,从世界管理器得到player对象
	player := wm.Players[pid]
	//从世界地图中删除
	wm.AorMgr.RemoveFromGridbyPos(int(pid), player.X, player.Z)
	//从世界管理中删除
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

//通过一个玩家Id得到一个Player对象
func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.pLock.RLock()
	p := wm.Players[pid]
	wm.pLock.RUnlock()
	return p
}

//获取全部的在线玩家集合
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	players := make([]*Player, 0)
	//将世界管理器的player对象加入到返回的切片中
	for _, player := range wm.Players {
		players = append(players, player)
	}
	return players
}

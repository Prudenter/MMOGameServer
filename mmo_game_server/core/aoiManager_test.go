/**
* @Author: ASlowPerson  
* @Date: 19-5-28 下午12:37
*/

package core

import (
	"testing"
	"fmt"
)

func TestNewAOIManager(t *testing.T) {
	//初始化AOIManager
	aoiManager := NewAOIManager(0, 250, 5, 0, 250, 5)
	//打印信息
	fmt.Println(aoiManager)
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	//初始化AOIManager
	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)
	//求出每个格子周边的九宫格信息
	for gId, _ := range aoiMgr.grids {
		grids := aoiMgr.GetSurroundGridsByGid(gId)
		fmt.Println("gid:", gId, "grids num=", len(grids))

		//当前九宫格的ID集合
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}

		fmt.Println("grids IDs are ", gIDs)
	}

	fmt.Println("-------------------------")
	playerIds := aoiMgr.GetSurroundPidsByPos(175,68)
	fmt.Println("playerIds:",playerIds)
}

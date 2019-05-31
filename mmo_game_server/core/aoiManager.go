/**
* @Author: ASlowPerson  
* @Date: 19-5-28 下午12:17
*/

package core

import "fmt"

/*
	AOi地图类
*/
type AOIManager struct {
	//区域的左边边界
	MinX int
	//区域的右边边界
	MaxX int
	//X轴方向格子的数量
	CntsX int
	//区域的上边边界
	MinY int
	//区域的下边边界
	MaxY int
	//Y轴方向的格子的数量
	CntsY int
	//整体区域(地图中)拥有哪些格子map:key格子ID,value:格子对象
	grids map[int]*Grid
}

//得到每个格子在X轴方向的宽度
func (m *AOIManager) GridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//得到每个格子在Y轴方向的高度
func (m *AOIManager) GridHeight() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

//初始化一个地图 AOIManager
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiManager := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}
	//隶属于当期地图的全部格子,也一并进行出初始化
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//初始化一个格子
			//格子ID:= cntsX*y + x
			gid := y*cntsX + x
			//给aoiManager添加一个格子
			aoiManager.grids[gid] = NewGrid(gid,
				aoiManager.MinX+x*aoiManager.GridWidth(),
				aoiManager.MinX+(x+1)*aoiManager.GridWidth(),
				aoiManager.MinX+y*aoiManager.GridHeight(),
				aoiManager.MinX+(y+1)*aoiManager.GridHeight())
		}
	}
	return aoiManager
}

//打印当前的地图信息
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager : \n MinX:%d,MaxX:%d,cntsX:%d, minY:%d, maxY:%d,cntsY:%d, Grids inManager:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	//打印全部的格子
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

//添加一个PlayerId到AOI格子中
func (m *AOIManager) AddPidToGrid(pId, gId int) {
	m.grids[gId].Add(pId, nil)
}

//从一个AOI格子中移除一个PlayerId
func (m *AOIManager) RemovePidFromGrid(pId, gId int) {
	m.grids[gId].Remove(pId)
}

//通过格子Id获取当前格子的全部PlayerId
func (m *AOIManager) GetPidsByGid(gId int) (playerIds []int) {
	playerIds = m.grids[gId].GetPlayerIds()
	return
}

//通过一个格子Id得到当前格子周边的九宫格的格子Id集合
func (m *AOIManager) GetSurroundGridsByGid(gId int) (grids []*Grid) {
	//判断gId是否存在在AOI中
	if _, ok := m.grids[gId]; !ok {
		fmt.Println("gId不存在!")
		return
	}
	//将当前格子ID放入九宫格切片中
	grids = append(grids, m.grids[gId])
	//判断gId左右两边是否有格子
	//1.通过gId得到当前X轴编号
	idX := gId % m.CntsX
	//2.根据X轴编号判断左右是否有格子
	if idX > 0 {
		//将左边的格子加入到grids中
		grids = append(grids, m.grids[gId-1])
	}
	if idX < m.CntsX-1 {
		//将右边的格子加入到grids中
		grids = append(grids, m.grids[gId+1])
	}
	//得到一个X轴的格子集合后,遍历这个集合,判断每个元素上下是否有格子
	gIdsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gIdsX = append(gIdsX, v.GID)
	}
	//遍历所有的gId,判断每个子元素上下是否有格子
	for _, gId := range gIdsX {
		//通过gId得到当期gId的Y轴编号
		idY := gId / m.CntsX
		//2.根据Y轴编号判断上下是否有格子
		if idY > 0 {
			//将左边的格子加入到grids中
			grids = append(grids, m.grids[gId-m.CntsX])
		}
		if idY < m.CntsY-1 {
			//将右边的格子加入到grids中
			grids = append(grids, m.grids[gId+m.CntsX])
		}
	}
	return
}

//通过x,y坐标得到对应的格子Id
func (m *AOIManager) GetGidByPos(x, y float32) int {
	if x < 0 || x > float32(m.MaxX) {
		fmt.Println("X坐标超出地图边界!")
		return -1
	}
	if y < 0 || y > float32(m.MaxY) {
		fmt.Println("Y坐标超出地图边界!")
		return -1
	}
	//根据坐标,得到当前玩家所在的格子的X轴和Y轴编号
	idX := (int(x) - m.MinX) / m.GridWidth()
	idY := (int(y) - m.MinY) / m.GridHeight()
	//根据X轴和Y轴编号得到当前玩家所在格子ID
	gId := idY*m.CntsX + idX
	return gId
}

//根据一个坐标,得到其周边九宫格之内的全部的玩家Id集合
func (m *AOIManager) GetSurroundPidsByPos(x, y float32) (playerIds []int) {
	//调用函数,获取坐标对应的格子Id
	gId := m.GetGidByPos(x, y)
	//通过格子Id,得到周边九宫格集合
	grids := m.GetSurroundGridsByGid(gId)
	fmt.Println("gid=", gId)
	//分别将九宫格内的全部玩家添加到playerIds中
	for _, grid := range grids {
		playerIds = append(playerIds, grid.GetPlayerIds()...)
	}
	return
}

//通过坐标,将pId加入到一个格子中
func (m *AOIManager) AddToGridByPos(pId int, x, y float32) {
	//调用函数,获取坐标对应的格子Id
	gId := m.GetGidByPos(x, y)
	//调用函数,将pid加入到对应的格子中
	m.AddPidToGrid(pId,gId)
}

//通过坐标,从一个AOI格子中移除一个PlayerId
func (m *AOIManager) RemoveFromGridbyPos(pId int, x, y float32) {
	//调用函数,获取坐标对应的格子Id
	gId := m.GetGidByPos(x, y)
	//调用函数,将pid从对应的格子中移除
	m.RemovePidFromGrid(pId,gId)
}
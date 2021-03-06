/**
* @Author: ASlowPerson  
* @Date: 19-5-31 下午5:11
*/
package main

import (
	"ZinxServerFramework/zinx/zinxNet"
	"ZinxServerFramework/zinx/zinxInterface"
	"fmt"
	"MMOGameServer/mmo_game_server/core"
	"MMOGameServer/mmo_game_server/apis"
)

//当前客户端建立链接之后触发Hook函数
func OnConnectionAdd(conn zinxInterface.InterfaceConnection) {
	fmt.Println("conn Add..")
	//创建一个玩家 将链接和玩家模块绑定
	p := core.NewPlayer(conn)
	//给客户端发送一个msgID:1的消息
	p.ReturnPid()
	//给客户端发送一个msgID:200的消息
	p.ReturnPlayerPosition()

	//玩家上线成功了,将玩家对象添加到世界管理器中
	core.WorldMgrObj.AddPlayer(p)

	//给conn添加一个属性,pid属性
	conn.SetProperty("pid",p.Pid)

	fmt.Println("----> player ID = ", p.Pid, "Online...", ", Player num = ", len(core.WorldMgrObj.Players))
}

func main() {
	//创建一个zinx server对象
	server := zinxNet.NewServer()

	//注册一些创建链接之后或者销毁链接之前的Hook钩子函数
	server.AddOnConnStart(OnConnectionAdd)

	//针对MsgID:2建立路由业务
	server.AddRouter(2,&apis.WorldChat{})

	//启动服务
	server.Run()
}
